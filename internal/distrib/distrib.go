// Package distrib implements helpers for distributing a reviewed research
// baseline to agent workspaces. Payloads are validated before anything is
// written and verified afterwards; any violation fails closed. Standard
// library only; no network access and no git interaction.
package distrib

import (
	"archive/tar"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"
)

// ValidateTar lists the regular-file entries of a tarball, rejecting
// absolute paths, traversal components, symlinks, hard links, duplicate
// names and every non-regular, non-directory entry type. Directories are
// permitted structurally but omitted from the returned list. The tarball
// itself is only read.
func ValidateTar(tarPath string) ([]string, error) {
	f, err := os.Open(tarPath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var out []string
	seen := map[string]bool{}
	tr := tar.NewReader(f)
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("read tar %s: %w", tarPath, err)
		}
		name := hdr.Name
		switch hdr.Typeflag {
		case tar.TypeDir:
			if _, err := cleanEntryPath(name); err != nil {
				return nil, err
			}
			continue
		case tar.TypeReg, tar.TypeRegA:
			// Accepted below after path validation.
		case tar.TypeSymlink, tar.TypeLink:
			return nil, fmt.Errorf("tar %s: rejecting link entry %q (links never distribute)", tarPath, name)
		default:
			return nil, fmt.Errorf("tar %s: rejecting unsupported entry %q (type %q)", tarPath, name, rune(hdr.Typeflag))
		}
		rel, err := cleanEntryPath(name)
		if err != nil {
			return nil, fmt.Errorf("tar %s: %w", tarPath, err)
		}
		if seen[rel] {
			return nil, fmt.Errorf("tar %s: duplicate entry %q", tarPath, name)
		}
		seen[rel] = true
		out = append(out, rel)
	}
	sort.Strings(out)
	return out, nil
}

// cleanEntryPath rejects absolute paths, traversal components and empty
// names, returning the cleaned slash-separated relative path.
func cleanEntryPath(name string) (string, error) {
	if name == "" {
		return "", fmt.Errorf("empty entry name")
	}
	if path.IsAbs(name) || filepath.IsAbs(name) {
		return "", fmt.Errorf("rejecting absolute entry %q", name)
	}
	for _, comp := range strings.Split(name, "/") {
		if comp == ".." {
			return "", fmt.Errorf("rejecting traversal entry %q", name)
		}
	}
	clean := path.Clean(name)
	if clean == "." || clean == "" {
		return "", fmt.Errorf("empty entry name %q", name)
	}
	return clean, nil
}

// WriteManifest returns manifest bytes ("<sha256hex>  <relpath>" per line,
// sorted) for the named files beneath dir. Every name must already be a
// clean relative path; anything else is an error.
func WriteManifest(dir string, rels []string) ([]byte, error) {
	names := append([]string(nil), rels...)
	sort.Strings(names)
	var b strings.Builder
	for _, rel := range names {
		if _, err := cleanEntryPath(rel); err != nil {
			return nil, err
		}
		raw, err := os.ReadFile(filepath.Join(dir, filepath.FromSlash(rel)))
		if err != nil {
			return nil, fmt.Errorf("manifest: read %s: %w", rel, err)
		}
		sum := sha256.Sum256(raw)
		fmt.Fprintf(&b, "%s  %s\n", hex.EncodeToString(sum[:]), rel)
	}
	return []byte(b.String()), nil
}

// VerifyManifest checks dir against manifest bytes. Every listed file must
// exist as a regular file with a matching digest; every file present under
// dir must be listed, except explicitly ignored paths and the .git
// directory (git metadata is never payload). Duplicate, malformed or
// unsafe manifest lines are errors. A nil ignore list ignores nothing
// besides .git.

// VerifyStaged checks an assembled staging tree. Payload files (the
// distribution manifest) must exist with matching digests; every other
// file present must belong to the base listing (the checkout the stage
// started from) or the ignore list, else it is unexpected. .git is always
// ignored. baseList holds one clean relative path per line, as produced
// by `git ls-files` at the export base.
func VerifyManifest(dir string, manifest []byte, ignore []string) error {
	skip := map[string]bool{}
	for _, ig := range ignore {
		skip[path.Clean(ig)] = true
	}
	want := map[string]string{}
	for i, line := range strings.Split(string(manifest), "\n") {
		if strings.TrimSpace(line) == "" {
			continue
		}
		parts := strings.SplitN(line, "  ", 2)
		if len(parts) != 2 || len(parts[0]) != 64 {
			return fmt.Errorf("manifest line %d: malformed %q", i+1, line)
		}
		if _, err := hex.DecodeString(parts[0]); err != nil {
			return fmt.Errorf("manifest line %d: bad digest %q", i+1, parts[0])
		}
		rel, err := cleanEntryPath(parts[1])
		if err != nil {
			return fmt.Errorf("manifest line %d: %w", i+1, err)
		}
		if _, dup := want[rel]; dup {
			return fmt.Errorf("manifest line %d: duplicate entry %q", i+1, rel)
		}
		want[rel] = parts[0]
	}
	for rel, digest := range want {
		raw, err := os.ReadFile(filepath.Join(dir, filepath.FromSlash(rel)))
		if err != nil {
			return fmt.Errorf("payload missing %s: %w", rel, err)
		}
		sum := sha256.Sum256(raw)
		if hex.EncodeToString(sum[:]) != digest {
			return fmt.Errorf("payload changed %s: digest mismatch", rel)
		}
	}
	allowed := map[string]bool{}
	for rel := range want {
		allowed[rel] = true
	}
	return checkUnexpected(dir, allowed, skip)
}

// VerifyStaged checks an assembled staging tree against a payload manifest
// plus the base listing it was assembled over. See VerifyManifest for the
// manifest rules; baseList excuses pre-existing checkout files.
func VerifyStaged(dir string, manifest, baseList []byte, ignore []string) error {
	skip := map[string]bool{}
	for _, ig := range ignore {
		skip[path.Clean(ig)] = true
	}
	base := map[string]bool{}
	for _, line := range strings.Split(string(baseList), "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		rel, err := cleanEntryPath(line)
		if err != nil {
			return fmt.Errorf("base list: %w", err)
		}
		base[rel] = true
	}
	want := map[string]string{}
	for i, line := range strings.Split(string(manifest), "\n") {
		if strings.TrimSpace(line) == "" {
			continue
		}
		parts := strings.SplitN(line, "  ", 2)
		if len(parts) != 2 || len(parts[0]) != 64 {
			return fmt.Errorf("manifest line %d: malformed %q", i+1, line)
		}
		if _, err := hex.DecodeString(parts[0]); err != nil {
			return fmt.Errorf("manifest line %d: bad digest %q", i+1, parts[0])
		}
		rel, err := cleanEntryPath(parts[1])
		if err != nil {
			return fmt.Errorf("manifest line %d: %w", i+1, err)
		}
		if _, dup := want[rel]; dup {
			return fmt.Errorf("manifest line %d: duplicate entry %q", i+1, rel)
		}
		want[rel] = parts[0]
	}
	for rel, digest := range want {
		raw, err := os.ReadFile(filepath.Join(dir, filepath.FromSlash(rel)))
		if err != nil {
			return fmt.Errorf("payload missing %s: %w", rel, err)
		}
		sum := sha256.Sum256(raw)
		if hex.EncodeToString(sum[:]) != digest {
			return fmt.Errorf("payload changed %s: digest mismatch", rel)
		}
	}
	allowed := map[string]bool{}
	for rel := range want {
		allowed[rel] = true
	}
	for rel := range base {
		allowed[rel] = true
	}
	return checkUnexpected(dir, allowed, skip)
}

func checkUnexpected(dir string, allowed, skip map[string]bool) error {
	var extra []string
	err := filepath.WalkDir(dir, func(p string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		rel, err := filepath.Rel(dir, p)
		if err != nil {
			return err
		}
		if rel == "." {
			return nil
		}
		rel = filepath.ToSlash(rel)
		if rel == ".git" || strings.HasPrefix(rel, ".git/") {
			if d.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}
		if skip[rel] {
			if d.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}
		if d.IsDir() {
			return nil
		}
		if _, ok := allowed[rel]; !ok {
			extra = append(extra, rel)
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("payload walk: %w", err)
	}
	if len(extra) > 0 {
		sort.Strings(extra)
		return fmt.Errorf("payload has %d unexpected file(s): %s", len(extra), strings.Join(extra, ", "))
	}
	return nil
}
