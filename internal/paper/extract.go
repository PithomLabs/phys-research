package paper

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
)

// Extract runs the optional `pdftotext` utility on the cached PDF for a
// paper. It writes fulltext.txt and updates the manifest. If `pdftotext`
// is unavailable, it updates the manifest with UNAVAILABLE and returns
// an explicit diagnostic (never claims success without verification).
func Extract(root, paperID, version string) error {
	if err := validVersionLabel(version); err != nil {
		return fmt.Errorf("extract %q: %w", paperID, err)
	}
	sanitizedID, err := sanitizePaperID(paperID)
	if err != nil {
		return fmt.Errorf("extract %q: invalid paper ID: %w", paperID, err)
	}
	pdfPath := SourcePDFPath(root, sanitizedID, version)
	if err := requireRegular(pdfPath, "source PDF"); err != nil {
		return fmt.Errorf("extract %q: %w", paperID, err)
	}
	if err := requireNotSymlink(ManifestPath(root, sanitizedID, version), "manifest"); err != nil {
		return fmt.Errorf("extract %q: %w", paperID, err)
	}
	if err := requireNotSymlink(FullTextPath(root, sanitizedID, version), "full-text output"); err != nil {
		return fmt.Errorf("extract %q: %w", paperID, err)
	}

	if _, err := exec.LookPath("pdftotext"); err != nil {
		// Report UNAVAILABLE explicitly in manifest and return diagnostic.
		manifestPath := ManifestPath(root, sanitizedID, version)
		if _, statErr := os.Stat(manifestPath); statErr == nil {
			// Update existing manifest with UNAVAILABLE extraction status.
			data, readErr := os.ReadFile(manifestPath)
			if readErr == nil {
				var cached CachedPaper
				if jsonErr := json.Unmarshal(data, &cached); jsonErr == nil {
					cached.Extraction = "unavailable"
					cached.Extractor = "pdftotext not found in PATH"
					updated, _ := json.Marshal(cached)
					_ = os.WriteFile(manifestPath, updated, 0o644)
				}
			}
		}
		return fmt.Errorf("extract %q: UNAVAILABLE — pdftotext not found in PATH (optional dependency missing); no silent success claimed", paperID)
	}

	fullTextPath := FullTextPath(root, sanitizedID, version)
	cmd := exec.Command("pdftotext", "-layout", pdfPath, fullTextPath)
	cmd.Stdout = nil
	cmd.Stderr = nil
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("extract %q: pdftotext failed: %w", paperID, err)
	}

	// Verify non-empty output.
	info, err := os.Stat(fullTextPath)
	if err != nil || info.Size() == 0 {
		return fmt.Errorf("extract %q: pdftotext produced empty or missing output at %s", paperID, fullTextPath)
	}

	// Update manifest with full extraction status.
	manifestPath := ManifestPath(root, sanitizedID, version)
	data, readErr := os.ReadFile(manifestPath)
	if readErr == nil {
		var cached CachedPaper
		if jsonErr := json.Unmarshal(data, &cached); jsonErr == nil {
			cached.Extraction = "full"
			cached.Extractor = "pdftotext -layout"
			updated, _ := json.Marshal(cached)
			_ = os.WriteFile(manifestPath, updated, 0o644)
		}
	}

	return nil
}

// requireRegular requires path to exist as a regular file. Symlinks are
// refused without being followed, so extraction never reads a PDF through
// a link planted by another writer.
func requireRegular(path, what string) error {
	fi, err := os.Lstat(path)
	if err != nil {
		return fmt.Errorf("%s not found at %s", what, path)
	}
	if fi.Mode()&os.ModeSymlink != 0 {
		return fmt.Errorf("refusing symlinked %s at %s", what, path)
	}
	if !fi.Mode().IsRegular() {
		return fmt.Errorf("%s at %s is not a regular file", what, path)
	}
	return nil
}

// requireNotSymlink refuses a destination that exists as a symlink. A
// missing destination is fine and returns nil.
func requireNotSymlink(path, what string) error {
	fi, err := os.Lstat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	if fi.Mode()&os.ModeSymlink != 0 {
		return fmt.Errorf("refusing symlinked %s at %s", what, path)
	}
	return nil
}
