package paper

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// pdfName and manifestName are the fixed leaf names inside a paper cache
// directory. They never come from user input.
const (
	pdfName      = "source.pdf"
	manifestName = "manifest.json"
)

// Fetch downloads an openly accessible paper PDF using its source URL
// (from arXiv or OpenAlex open_access URL) into the local paper-cache.
// Placement is collision-safe: an existing cached PDF with identical bytes
// is left untouched (idempotent re-fetch); differing bytes produce an
// explicit collision error, never silent replacement. Source records cite
// cached bytes by SHA-256, so replacing them would invalidate provenance.
// Symlinks at any cache destination are refused outright: a link is never
// accepted as an equivalent cached file, even when its target bytes match.
func Fetch(root, paperID, version, sourceURL, downloadURL string) (*CachedPaper, error) {
	if downloadURL == "" && sourceURL == "" {
		return nil, fmt.Errorf("fetch %q: neither source URL nor download URL provided", paperID)
	}
	urlToFetch := downloadURL
	if urlToFetch == "" {
		urlToFetch = sourceURL
	}

	// Sanitize paperID to prevent path traversal
	sanitizedID, err := sanitizePaperID(paperID)
	if err != nil {
		return nil, fmt.Errorf("fetch %q: invalid paper ID: %w", paperID, err)
	}
	if err := validVersionLabel(version); err != nil {
		return nil, fmt.Errorf("fetch %q: %w", paperID, err)
	}

	pdfData, err := safeFetch(urlToFetch, 60)
	if err != nil {
		return nil, fmt.Errorf("fetch %q failed: %w", paperID, err)
	}

	dir, err := openCacheDir(root, sanitizedID, version)
	if err != nil {
		return nil, fmt.Errorf("fetch %q: %w", paperID, err)
	}
	defer dir.rt.Close()

	meta := &CachedPaper{
		PaperID:     sanitizedID,
		Version:     version,
		SourceURL:   sourceURL,
		DownloadURL: urlToFetch,
		RetrievedAt: time.Now().UTC().Format("2006-01-02T15:04:05Z"),
	}
	return placeDownload(paperID, dir, pdfData, meta)
}

// cacheDir is a paper cache directory held open as an anchored root: all
// file operations inside it resolve beneath dirAbs and cannot escape it.
type cacheDir struct {
	rt  *os.Root
	abs string
}

// openCacheDir opens (creating as needed) the cache directory for one
// paper@version, refusing symlinks at every path component. Components come
// from a sanitized paper ID and a validated single-segment version label,
// so traversal is impossible even before the symlink checks.
func openCacheDir(root, paperID, version string) (*cacheDir, error) {
	if err := validVersionLabel(version); err != nil {
		return nil, err
	}
	rt, err := os.OpenRoot(root)
	if err != nil {
		return nil, err
	}
	chain := append([]string{"paper-cache"}, strings.Split(paperID, "/")...)
	chain = append(chain, version)
	cur := rt
	for _, comp := range chain {
		if err := cur.Mkdir(comp, 0o755); err != nil && !os.IsExist(err) {
			rt.Close()
			return nil, err
		}
		fi, err := cur.Lstat(comp)
		if err != nil {
			rt.Close()
			return nil, err
		}
		if fi.Mode()&os.ModeSymlink != 0 {
			rt.Close()
			return nil, fmt.Errorf("refusing symlinked cache component %q", comp)
		}
		if !fi.IsDir() {
			rt.Close()
			return nil, fmt.Errorf("cache component %q exists and is not a directory", comp)
		}
		next, err := cur.OpenRoot(comp)
		if err != nil {
			rt.Close()
			return nil, err
		}
		if cur != rt {
			cur.Close()
		}
		cur = next
	}
	abs := paperDirFor(root, paperID, version)
	return &cacheDir{rt: cur, abs: abs}, nil
}

// placeDownload stores downloaded PDF bytes and their manifest without ever
// replacing differing bytes. On an idempotent re-fetch it returns the
// existing manifest record untouched. It performs no network I/O.
//
// Every success path ends in verifyPlaced, so any tampering that persists
// to the end of the operation is detected; tampering fully contained
// inside the operation that reverts before verification leaves a
// consistent entry behind. Only post-return, at-rest modification escapes
// these checks, and that is outside any single operation's scope (the same
// holds for every file on disk). Cross-agent cache write access is already
// outside the worktree model (see docs/AGENT_CONTRACT.md).
func placeDownload(paperID string, dir *cacheDir, pdfData []byte, meta *CachedPaper) (*CachedPaper, error) {
	newDigest := digestString(pdfData)

	existing, err := readRegular(dir.rt, pdfName)
	if err != nil {
		return nil, fmt.Errorf("fetch %q: %w", paperID, err)
	}
	if existing != nil {
		if digestString(existing) == newDigest {
			got, merr := ensureManifest(paperID, dir, meta, newDigest, false)
			if merr != nil {
				return nil, merr
			}
			if verr := verifyPlaced(paperID, dir, newDigest); verr != nil {
				return nil, verr
			}
			return got, nil
		}
		return nil, fmt.Errorf("fetch %q: collision: cached PDF at %s has sha256:%s, refusing to replace with differing bytes (sha256:%s); use a different version label",
			paperID, filepath.Join(dir.abs, pdfName), digestString(existing), newDigest)
	}

	// No PDF yet, but a manifest without a PDF still constrains placement
	// (a previous run may have recorded different bytes).
	if prev, merr := readManifest(dir.rt); merr != nil {
		return nil, fmt.Errorf("fetch %q: %w", paperID, merr)
	} else if prev != nil && prev.SHA256PDF != "" && prev.SHA256PDF != newDigest {
		return nil, fmt.Errorf("fetch %q: collision: manifest at %s records sha256:%s, refusing to place differing bytes (sha256:%s)",
			paperID, filepath.Join(dir.abs, manifestName), prev.SHA256PDF, newDigest)
	}

	// First writer wins; a racing loser re-verifies against the winner.
	createdPDF := false
	if err := exclWrite(dir.rt, pdfName, pdfData); err != nil {
		if !errors.Is(err, os.ErrExist) {
			return nil, fmt.Errorf("fetch %q: place PDF: %w", paperID, err)
		}
		existing, rerr := readRegular(dir.rt, pdfName)
		if rerr != nil {
			return nil, fmt.Errorf("fetch %q: re-read raced cache entry: %w", paperID, rerr)
		}
		if existing == nil || digestString(existing) != newDigest {
			return nil, fmt.Errorf("fetch %q: collision: concurrent fetch placed differing bytes at %s; downloaded sha256:%s",
				paperID, filepath.Join(dir.abs, pdfName), newDigest)
		}
		got, merr := ensureManifest(paperID, dir, meta, newDigest, false)
		if merr != nil {
			return nil, merr
		}
		if verr := verifyPlaced(paperID, dir, newDigest); verr != nil {
			return nil, verr
		}
		return got, nil
	}
	createdPDF = true
	got, err := ensureManifest(paperID, dir, meta, newDigest, createdPDF)
	if err != nil {
		return nil, err
	}
	// Post-verify: the PDF destination must still be a regular file whose
	// bytes match the manifest. On failure remove only what this call
	// created and report; never leave a mismatched entry behind.
	if verr := verifyPlaced(paperID, dir, newDigest); verr != nil {
		if createdPDF {
			_ = dir.rt.Remove(pdfName)
		}
		return nil, verr
	}
	return got, nil
}

// ensureManifest records the manifest for placed bytes. It writes only when
// the manifest is absent; a present manifest must agree with the placed
// digest, otherwise the entry is inconsistent and an error is returned.
// Symlinked manifests are always refused.
func ensureManifest(paperID string, dir *cacheDir, meta *CachedPaper, digest string, createdPDF bool) (*CachedPaper, error) {
	meta.SHA256PDF = digest
	data, _ := json.Marshal(meta)
	if err := exclWrite(dir.rt, manifestName, data); err == nil {
		return meta, nil
	} else if !errors.Is(err, os.ErrExist) {
		if createdPDF {
			_ = dir.rt.Remove(pdfName)
		}
		return meta, fmt.Errorf("fetch %q: write manifest: %w", paperID, err)
	}
	prev, rerr := readManifest(dir.rt)
	if rerr != nil {
		if createdPDF {
			_ = dir.rt.Remove(pdfName)
		}
		return nil, fmt.Errorf("fetch %q: %w", paperID, rerr)
	}
	if prev == nil || prev.SHA256PDF == "" {
		if createdPDF {
			_ = dir.rt.Remove(pdfName)
		}
		return nil, fmt.Errorf("fetch %q: collision: unreadable manifest at %s disagrees with placed bytes (sha256:%s)",
			paperID, filepath.Join(dir.abs, manifestName), digest)
	}
	if prev.SHA256PDF != digest {
		if createdPDF {
			_ = dir.rt.Remove(pdfName)
		}
		return nil, fmt.Errorf("fetch %q: collision: manifest at %s records sha256:%s, placed bytes are sha256:%s",
			paperID, filepath.Join(dir.abs, manifestName), prev.SHA256PDF, digest)
	}
	return prev, nil
}

// verifyPlaced re-checks that the PDF destination is a regular file whose
// digest matches the manifest. It defeats anything planted between the
// initial checks and the end of placement.
func verifyPlaced(paperID string, dir *cacheDir, digest string) error {
	data, err := readRegular(dir.rt, pdfName)
	if err != nil {
		return fmt.Errorf("fetch %q: post-place verification failed: %w", paperID, err)
	}
	if data == nil || digestString(data) != digest {
		return fmt.Errorf("fetch %q: post-place verification failed: cached bytes changed during placement", paperID)
	}
	return nil
}

// readRegular reads a cache leaf, refusing symlinks and non-regular files.
// A missing file returns (nil, nil).
func readRegular(rt *os.Root, name string) ([]byte, error) {
	fi, err := rt.Lstat(name)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, err
	}
	if fi.Mode()&os.ModeSymlink != 0 {
		return nil, fmt.Errorf("refusing symlinked cache destination %q: a link is never accepted as a cached file", name)
	}
	if !fi.Mode().IsRegular() {
		return nil, fmt.Errorf("cache destination %q is not a regular file", name)
	}
	f, err := rt.Open(name)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return io.ReadAll(f)
}

// readManifest reads and parses the cache manifest, refusing symlinks. A
// missing manifest returns (nil, nil).
func readManifest(rt *os.Root) (*CachedPaper, error) {
	data, err := readRegular(rt, manifestName)
	if err != nil || data == nil {
		return nil, err
	}
	var prev CachedPaper
	if json.Unmarshal(data, &prev) != nil {
		return nil, fmt.Errorf("unreadable manifest %q", manifestName)
	}
	return &prev, nil
}

// exclWrite creates a file with the given bytes, failing with os.ErrExist
// when the destination already exists. It never modifies an existing file.
func exclWrite(rt *os.Root, name string, data []byte) error {
	f, err := rt.OpenFile(name, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0o644)
	if err != nil {
		return err
	}
	if _, err := f.Write(data); err != nil {
		_ = f.Close()
		return err
	}
	return f.Close()
}
