// Package paper implements the paper-retrieval layer for physctl:
// search (arXiv + OpenAlex), related (OpenAlex citation/reference),
// fetch (open-access PDF download), and extract (optional pdftotext).
//
// It uses only the Go standard library, respects rate limits, never bypasses
// access controls, and treats all downloaded content as untrusted input.
// The local paper-cache/ directory (gitignored) holds PDFs and manifests;
// original source records stay in Git using the existing P- namespace.
package paper

import (
	"context"
	"crypto/sha256"
	"crypto/tls"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

const userAgent = "phys-research/1.0 (paper-retrieval; open-access only)"

// SearchResult is one paper discovered by the search operation.
type SearchResult struct {
	SourceType string `json:"source_type"` // arxiv | openalex | crossref
	PaperID    string `json:"paper_id"`    // arXiv ID (e.g. 2209.15487v1) or OpenAlex ID or DOI
	Title      string `json:"title"`
	Authors    string `json:"authors,omitempty"`
	DOI        string `json:"doi,omitempty"`
	ArXivID    string `json:"arxiv_id,omitempty"`
	ArXivVer   string `json:"arxiv_version,omitempty"`
	URL        string `json:"url"`              // source URL (PDF or abstract)
	OAURL      string `json:"oa_url,omitempty"` // open-access full-text URL
	IsOA       bool   `json:"is_oa,omitempty"`
	PubYear    int    `json:"publication_year,omitempty"`
	Abstract   string `json:"abstract,omitempty"`
}

// CachedPaper records the local retrieval state of a downloaded paper.
// It lives in paper-cache/<paper-id>/v<version>/manifest.json.
type CachedPaper struct {
	PaperID     string `json:"paper_id"`
	Version     string `json:"version"`
	DOI         string `json:"doi,omitempty"`
	ArXivID     string `json:"arxiv_id,omitempty"`
	Title       string `json:"title"`
	SourceURL   string `json:"source_url"`
	DownloadURL string `json:"download_url"`
	SHA256PDF   string `json:"sha256_pdf"`
	RetrievedAt string `json:"retrieved_at"`
	License     string `json:"license,omitempty"`
	Extraction  string `json:"extraction"` // full | partial | unavailable
	Extractor   string `json:"extractor,omitempty"`
}

// paperCacheDir returns the paper-cache root relative to the repo root.
func paperCacheDir(root string) string {
	return filepath.Join(root, "paper-cache")
}

// paperDirFor returns the versioned directory for a given paper ID.
func paperDirFor(root, paperID, version string) string {
	return filepath.Join(paperCacheDir(root), paperID, version)
}

// ManifestPath returns the manifest.json path for a cached paper.
func ManifestPath(root, paperID, version string) string {
	return filepath.Join(paperDirFor(root, paperID, version), "manifest.json")
}

// SourcePDFPath returns the downloaded PDF path for a cached paper.
func SourcePDFPath(root, paperID, version string) string {
	return filepath.Join(paperDirFor(root, paperID, version), "source.pdf")
}

// FullTextPath returns the extracted full-text path.
func FullTextPath(root, paperID, version string) string {
	return filepath.Join(paperDirFor(root, paperID, version), "fulltext.txt")
}

// isPrivateIP checks if an IP address is private, loopback, link-local, multicast, or otherwise prohibited.
func isPrivateIP(ip net.IP) bool {
	if ip.IsLoopback() || ip.IsLinkLocalUnicast() || ip.IsLinkLocalMulticast() || ip.IsMulticast() {
		return true
	}
	// RFC 1918 private ranges
	if ip4 := ip.To4(); ip4 != nil {
		// 10.0.0.0/8
		if ip4[0] == 10 {
			return true
		}
		// 172.16.0.0/12
		if ip4[0] == 172 && ip4[1] >= 16 && ip4[1] <= 31 {
			return true
		}
		// 192.168.0.0/16
		if ip4[0] == 192 && ip4[1] == 168 {
			return true
		}
		// 169.254.0.0/16 (link-local)
		if ip4[0] == 169 && ip4[1] == 254 {
			return true
		}
		// 100.64.0.0/10 (CGNAT)
		if ip4[0] == 100 && ip4[1] >= 64 && ip4[1] <= 127 {
			return true
		}
		// 192.0.0.0/24 (IETF protocol assignments)
		if ip4[0] == 192 && ip4[1] == 0 && ip4[2] == 0 {
			return true
		}
		// 192.0.2.0/24 (TEST-NET-1), 198.51.100.0/24 (TEST-NET-2), 203.0.113.0/24 (TEST-NET-3)
		if ip4[0] == 192 && ip4[1] == 0 && ip4[2] == 2 {
			return true
		}
		if ip4[0] == 198 && ip4[1] == 51 && ip4[2] == 100 {
			return true
		}
		if ip4[0] == 203 && ip4[1] == 0 && ip4[2] == 113 {
			return true
		}
	}
	// IPv6 unique local addresses (fc00::/7)
	if ip16 := ip.To16(); ip16 != nil && len(ip16) == 16 {
		if ip16[0]&0xfe == 0xfc {
			return true
		}
	}
	return false
}

// resolveAndValidateHost resolves a hostname and validates that all resolved IPs are safe.
// Returns the first safe resolved address, or an error if any resolved IP is unsafe.
func resolveAndValidateHost(host string) (string, error) {
	ips, err := net.LookupIP(host)
	if err != nil {
		return "", fmt.Errorf("DNS lookup failed for %q: %w", host, err)
	}
	if len(ips) == 0 {
		return "", fmt.Errorf("no IP addresses resolved for %q", host)
	}
	for _, ip := range ips {
		if isPrivateIP(ip) {
			return "", fmt.Errorf("unsafe resolved IP %s for host %q", ip.String(), host)
		}
	}
	// Return the first safe IP as the validated address to connect to
	return ips[0].String(), nil
}

// isSafeHost checks if a host string (hostname or IP) is safe for the initial request.
// For hostnames, we can't fully validate until we resolve, but we can reject obvious bad patterns.
func isSafeHost(host string) error {
	// Reject explicit IP addresses that are private
	if ip := net.ParseIP(host); ip != nil {
		if isPrivateIP(ip) {
			return fmt.Errorf("unsafe IP address %q", host)
		}
		return nil
	}
	// For hostnames, reject obvious bad patterns
	lower := strings.ToLower(host)
	if lower == "localhost" || strings.HasPrefix(lower, "localhost.") {
		return fmt.Errorf("unsafe host %q", host)
	}
	return nil
}

// safeTransport builds the HTTP transport for safeFetch. When the environment
// configures an egress proxy (HTTPS_PROXY and friends), requests honor it:
// dialing goes to the operator-provided proxy while TLS certificate
// verification against the target host stays enforced, so a proxy-only
// environment can still fetch. Without a configured proxy the transport
// keeps the direct path with per-connection DNS resolution and IP
// validation (no private/loopback/link-local/multicast targets, no DNS
// rebinding to unvalidated addresses).
func safeTransport() *http.Transport {
	tlsCfg := &tls.Config{
		ServerName:         "",    // Set per request from the target host
		InsecureSkipVerify: false, // Enforce cert verification
	}
	if envProxyForHTTPS() {
		return proxyTransport(tlsCfg)
	}
	return directTransport(tlsCfg)
}

// proxyTransport routes through the operator-provided egress proxy while
// keeping TLS certificate verification against the target host enforced.
func proxyTransport(tlsCfg *tls.Config) *http.Transport {
	return &http.Transport{
		Proxy:           http.ProxyFromEnvironment,
		TLSClientConfig: tlsCfg,
	}
}

// directTransport dials targets directly with per-connection DNS resolution
// and IP validation.
func directTransport(tlsCfg *tls.Config) *http.Transport {
	return &http.Transport{
		DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
			// Parse the address to get host:port
			host, port, err := net.SplitHostPort(addr)
			if err != nil {
				return nil, fmt.Errorf("invalid address %q: %w", addr, err)
			}
			// Resolve and validate the host
			validatedIP, err := resolveAndValidateHost(host)
			if err != nil {
				return nil, err
			}
			// Connect to the validated IP with the original port
			validatedAddr := net.JoinHostPort(validatedIP, port)
			d := &net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}
			return d.DialContext(ctx, network, validatedAddr)
		},
		TLSClientConfig: tlsCfg,
	}
}

// envProxyForHTTPS reports whether the environment routes HTTPS requests
// through a proxy. It performs no network I/O.
func envProxyForHTTPS() bool {
	probe := &http.Request{URL: &url.URL{Scheme: "https", Host: "example.invalid"}}
	proxyURL, err := http.ProxyFromEnvironment(probe)
	return err == nil && proxyURL != nil
}

// safeFetch performs a bounded, safe HTTP GET with identifying User-Agent,
// timeout, redirect limits, and content-type checks. It validates all redirect
// destinations, verifies PDF magic bytes, enforces streaming size limits,
// and validates Content-Type as an additional signal (not sole test).
// On direct (non-proxied) connections it additionally blocks
// private/loopback/link-local/multicast IPs and prevents DNS rebinding by
// connecting only to validated resolved addresses.
func safeFetch(urlStr string, timeoutSec int) ([]byte, error) {
	parsedURL, err := url.Parse(urlStr)
	if err != nil {
		return nil, fmt.Errorf("invalid URL %q: %w", urlStr, err)
	}
	if parsedURL.Scheme != "https" {
		return nil, fmt.Errorf("only HTTPS URLs are allowed, got %q", parsedURL.Scheme)
	}

	// Validate initial host
	if err := isSafeHost(parsedURL.Host); err != nil {
		return nil, fmt.Errorf("unsafe initial host %q: %w", parsedURL.Host, err)
	}

	transport := safeTransport()

	client := &http.Client{
		Timeout:       time.Duration(60) * time.Second,
		Transport:     transport,
		CheckRedirect: makeCheckRedirect(""),
	}

	req, err := http.NewRequest("GET", urlStr, nil)
	if err != nil {
		return nil, fmt.Errorf("invalid URL %q: %w", urlStr, err)
	}
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Accept", "application/pdf")

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("fetch %q failed: %w", urlStr, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("fetch %q returned status %d", urlStr, resp.StatusCode)
	}

	// Validate Content-Type as an additional signal (not sole test)
	ct := resp.Header.Get("Content-Type")
	if strings.Contains(ct, "text/html") || strings.Contains(ct, "application/json") {
		// Read a bit to confirm it's an error page
		sniff := make([]byte, 512)
		n, _ := io.ReadFull(resp.Body, sniff)
		if n > 0 && (strings.Contains(string(sniff[:n]), "<html") || strings.Contains(string(sniff[:n]), "<!DOCTYPE")) {
			return nil, fmt.Errorf("fetch returned HTML instead of PDF (status %d)", resp.StatusCode)
		}
	}

	// Read with bounded size limit (50MB), streaming to detect PDF magic
	const maxSize = 50 * 1024 * 1024
	var out []byte
	var magicBuf [4]byte
	magicRead := 0

	limit := io.LimitedReader{R: resp.Body, N: int64(maxSize)}
	for {
		buf := make([]byte, 32*1024)
		n, err := limit.Read(buf)
		if n > 0 {
			if magicRead < 4 {
				copy(magicBuf[magicRead:], buf[:min(4-magicRead, n)])
				magicRead += min(4-magicRead, n)
				if magicRead == 4 {
					// Check PDF magic bytes (%PDF)
					if string(magicBuf[:4]) != "%PDF" {
						return nil, fmt.Errorf("downloaded content is not a valid PDF (missing %%PDF magic bytes)")
					}
				}
			}
			out = append(out, buf[:n]...)
		}
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, fmt.Errorf("fetch read failed: %w", err)
		}
	}

	// Verify PDF magic bytes were found
	if magicRead < 4 {
		return nil, fmt.Errorf("downloaded content too small to be a valid PDF")
	}
	if string(magicBuf[:4]) != "%PDF" {
		return nil, fmt.Errorf("downloaded content is not a valid PDF (missing %%PDF magic bytes)")
	}

	if limit.N <= 0 {
		return out, fmt.Errorf("download exceeded %d bytes", maxSize)
	}
	return out, nil
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// makeCheckRedirect creates a redirect validation function that validates each redirect target.
func makeCheckRedirect(originalHost string) func(*http.Request, []*http.Request) error {
	return func(req *http.Request, via []*http.Request) error {
		if len(via) >= 5 {
			return fmt.Errorf("redirect loop exceeded 5 hops")
		}
		// Validate redirect target host
		if err := isSafeHost(req.Host); err != nil {
			return fmt.Errorf("redirect to unsafe host %q: %w", req.Host, err)
		}
		// Validate the redirect URL scheme
		if req.URL.Scheme != "https" {
			return fmt.Errorf("redirect to non-HTTPS URL %q", req.URL.String())
		}
		return nil
	}
}

// readBounded reads up to maxBytes from r, returning an error if exceeded.
func readBounded(r io.Reader, maxBytes int) ([]byte, error) {
	limit := io.LimitedReader{R: r, N: int64(maxBytes)}
	data, err := io.ReadAll(&limit)
	if limit.N <= 0 {
		return data, fmt.Errorf("download exceeded %d bytes", maxBytes)
	}
	return data, err
}

// digestString computes the SHA-256 hex of a byte slice.
func digestString(data []byte) string {
	h := sha256.New()
	h.Write(data)
	return fmt.Sprintf("%x", h.Sum(nil))
}

// cacheSegmentPattern admits only single-segment, filesystem-safe labels.
var cacheSegmentPattern = regexp.MustCompile(`^[A-Za-z0-9][A-Za-z0-9._-]*$`)

// validVersionLabel rejects version labels that could escape the paper
// cache directory. Every entry point that joins a version into a path
// must call it before touching the filesystem.
func validVersionLabel(version string) error {
	if version == "." || version == ".." || !cacheSegmentPattern.MatchString(version) {
		return fmt.Errorf("invalid version label %q (single path segment, letters/digits/dot/underscore/dash)", version)
	}
	return nil
}

// sanitizePaperID validates and sanitizes a paper ID for use in filesystem paths.
// Rejects path traversal attempts and returns a sanitized version.
func sanitizePaperID(paperID string) (string, error) {
	if paperID == "" {
		return "", fmt.Errorf("empty paper ID")
	}
	// Only allow alphanumeric, dots, hyphens, underscores, colons, slashes
	// This is permissive for arXiv IDs (e.g., "2209.15487v1"), DOIs, OpenAlex IDs
	valid := regexp.MustCompile(`^[a-zA-Z0-9._:/-]+$`)
	if !valid.MatchString(paperID) {
		return "", fmt.Errorf("paper ID %q contains invalid characters", paperID)
	}
	// Check for path traversal
	if strings.Contains(paperID, "..") || strings.HasPrefix(paperID, "/") || strings.HasPrefix(paperID, "./") {
		return "", fmt.Errorf("paper ID %q contains path traversal", paperID)
	}
	// Prevent absolute paths
	if filepath.IsAbs(paperID) {
		return "", fmt.Errorf("paper ID %q is an absolute path", paperID)
	}
	// Clean and verify it doesn't escape
	clean := filepath.Clean(paperID)
	if strings.HasPrefix(clean, "..") || strings.HasPrefix(clean, "/") {
		return "", fmt.Errorf("paper ID %q resolves outside cache directory", paperID)
	}
	return clean, nil
}
