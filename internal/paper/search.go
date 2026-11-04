package paper

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// arXivSearch calls the arXiv Atom API for a given query (bounded).
func arXivSearch(query string, maxResults int) ([]SearchResult, error) {
	if maxResults <= 0 || maxResults > 30 {
		maxResults = 10
	}

	// Apply arXiv rate limiting (1 req/3s) - cross-process persistent
	if err := WaitForArXivRateLimit(); err != nil {
		return nil, fmt.Errorf("arXiv rate limit: %w", err)
	}

	base := "https://export.arxiv.org/api/query"
	q := fmt.Sprintf("search_query=all:%s&start=0&max_results=%d", url.QueryEscape(query), maxResults)
	urlStr := base + "?" + q

	req, err := http.NewRequest("GET", urlStr, nil)
	if err != nil {
		return nil, fmt.Errorf("arXiv search: invalid URL %q: %w", urlStr, err)
	}
	req.Header.Set("User-Agent", userAgent)

	resp, err := doWithRetry(req, &http.Client{Timeout: 15 * time.Second}, DefaultRetryConfig())
	if err != nil {
		return nil, fmt.Errorf("arXiv search failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("arXiv search returned status %d", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("arXiv search read failed: %w", err)
	}

	return parseArXivAtom(data)
}

// Search searches both arXiv and OpenAlex, deduplicates, and returns results.
func Search(query string, maxResults int) ([]SearchResult, error) {
	// Initialize persistent rate limiters on first use
	if err := initPersistentRateLimiters(); err != nil {
		return nil, fmt.Errorf("init persistent rate limiters: %w", err)
	}

	arXivResults, arXivErr := arXivSearch(query, maxResults/2+1)
	openAlexResults, openErr := openAlexSearch(query, maxResults/2+1, "")

	var combined []SearchResult
	seen := map[string]bool{}

	for _, r := range arXivResults {
		key := r.SourceType + ":" + r.PaperID
		if !seen[key] {
			seen[key] = true
			combined = append(combined, r)
		}
	}
	for _, r := range openAlexResults {
		key := r.SourceType + ":" + r.PaperID
		if !seen[key] {
			seen[key] = true
			combined = append(combined, r)
		}
	}

	if len(combined) > maxResults {
		combined = combined[:maxResults]
	}

	if arXivErr != nil && openErr != nil {
		return combined, fmt.Errorf("search partially failed: arXiv=%v; openalex=%v", arXivErr, openErr)
	}
	return combined, nil
}

func httpQuery(q string) string {
	// Minimal query encoding for the standard library (no external URL-encoding needed for basic queries).
	q = fmt.Sprintf("all:%s", q)
	return q
}

// Minimal Atom parser: only extracts the fields we care about.
type arXivFeed struct {
	XMLName xml.Name     `xml:"feed"`
	Entries []arXivEntry `xml:"entry"`
}

type arXivLink struct {
	XMLName xml.Name `xml:"link"`
	Title   string   `xml:"title,attr"`
	Href    string   `xml:"href,attr"`
}

type arXivAuthor struct {
	XMLName xml.Name `xml:"author"`
	Name    string   `xml:"name"`
}

type arXivEntry struct {
	XMLName xml.Name      `xml:"entry"`
	Title   string        `xml:"title"`
	Summary string        `xml:"summary"`
	ID      string        `xml:"id"`
	Updated string        `xml:"updated"`
	Authors []arXivAuthor `xml:"author"`
	Links   []arXivLink   `xml:"link"`
}

func parseArXivAtom(data []byte) ([]SearchResult, error) {
	var feed arXivFeed
	if err := xml.Unmarshal(data, &feed); err != nil {
		return nil, fmt.Errorf("arXiv Atom parse failed: %w", err)
	}
	var out []SearchResult
	for _, entry := range feed.Entries {
		arXivID := ""
		if entry.ID != "" {
			// ID is like https://arxiv.org/abs/2209.15487v1
			parts := strings.Split(entry.ID, "/")
			if len(parts) > 0 {
				arXivID = parts[len(parts)-1]
			}
		}
		pdfLink := ""
		absLink := ""
		for _, l := range entry.Links {
			if l.Title == "pdf" {
				pdfLink = l.Href
			}
			if strings.HasPrefix(l.Href, "https://arxiv.org/abs/") {
				absLink = l.Href
			}
		}
		authors := ""
		for _, a := range entry.Authors {
			if authors != "" {
				authors += ", "
			}
			authors += a.Name
		}
		out = append(out, SearchResult{
			SourceType: "arxiv",
			PaperID:    arXivID,
			Title:      strings.TrimSpace(entry.Title),
			Authors:    authors,
			Abstract:   strings.TrimSpace(entry.Summary),
			ArXivID:    arXivID,
			URL:        absLink,
			OAURL:      pdfLink,
			IsOA:       true,
		})
	}
	return out, nil
}

// initPersistentRateLimiters initializes the persistent rate limiters.
// It reads the cache root from the environment or finds it relative to the current directory.
func initPersistentRateLimiters() error {
	root := ""
	if v := os.Getenv("PHYS_RESEARCH_ROOT"); v != "" {
		root = v
	} else {
		wd, err := os.Getwd()
		if err != nil {
			return err
		}
		// Walk up to find the research directory
		for {
			if fi, err := os.Stat(filepath.Join(wd, "research")); err == nil && fi.IsDir() {
				root = wd
				break
			}
			parent := filepath.Dir(wd)
			if parent == wd {
				break
			}
			wd = parent
		}
		if root == "" {
			root = "."
		}
	}
	return InitPersistentRateLimiters(root)
}
