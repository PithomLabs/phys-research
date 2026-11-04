// Package store loads research records from the repository layout defined
// in AGENTS.md §6.3. Every .md file directly inside a designated record
// directory must carry a valid frontmatter header; parse failures are kept
// on the record so validation can report them instead of crashing.
package store

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/PithomLabs/phys-research/internal/frontmatter"
)

// RecordDirs are the directories under the repository root that hold
// research records directly (one level, no subdirectories).
var RecordDirs = []string{
	"research/sources",
	"research/issues",
	"research/hypotheses",
	"research/reviews",
	"research/evidence",
}

// DerivationsDir holds derivation records at its top level. Subdirectories
// of derivations/ hold free-form human-authored sketches and are never
// parsed as records.
const DerivationsDir = "derivations"

// Record is one loaded research record.
type Record struct {
	Rel      string // repository-relative path, slash-separated
	Path     string // absolute path
	Raw      []byte // exact file bytes (digest is computed over these)
	Header   frontmatter.Header
	Body     []byte // bytes after the closing frontmatter delimiter
	BodyLine int    // 1-based file line where Body begins (for diagnostics)
	ParseErr error  // non-nil when the frontmatter could not be parsed
	Digest   string // SHA-256 hex of Raw (review snapshot provenance)
}

// Load reads every record in the known directories, sorted by relative
// path for deterministic output. Missing directories are tolerated.
func Load(root string) ([]Record, error) {
	var out []Record

	dirs := append([]string{}, RecordDirs...)
	for _, d := range dirs {
		entries, err := os.ReadDir(filepath.Join(root, d))
		if os.IsNotExist(err) {
			continue
		}
		if err != nil {
			return nil, fmt.Errorf("store: read %s: %w", d, err)
		}
		for _, e := range entries {
			if e.IsDir() || strings.HasPrefix(e.Name(), ".") || !strings.HasSuffix(e.Name(), ".md") {
				continue
			}
			rel := d + "/" + e.Name()
			rec, err := read(root, rel)
			if err != nil {
				return nil, err
			}
			out = append(out, rec)
		}
	}

	entries, err := os.ReadDir(filepath.Join(root, DerivationsDir))
	if err == nil {
		for _, e := range entries {
			if e.IsDir() || strings.HasPrefix(e.Name(), ".") || !strings.HasSuffix(e.Name(), ".md") {
				continue
			}
			rel := DerivationsDir + "/" + e.Name()
			rec, err := read(root, rel)
			if err != nil {
				return nil, err
			}
			out = append(out, rec)
		}
	} else if !os.IsNotExist(err) {
		return nil, fmt.Errorf("store: read %s: %w", DerivationsDir, err)
	}

	sort.Slice(out, func(i, j int) bool { return out[i].Rel < out[j].Rel })
	return out, nil
}

func read(root, rel string) (Record, error) {
	path := filepath.Join(root, filepath.FromSlash(rel))
	raw, err := os.ReadFile(path)
	if err != nil {
		return Record{}, fmt.Errorf("store: read %s: %w", rel, err)
	}
	h, body, bodyLine, perr := frontmatter.Parse(raw)
	sum := sha256.Sum256(raw)
	return Record{
		Rel:      rel,
		Path:     path,
		Raw:      raw,
		Header:   h,
		Body:     body,
		BodyLine: bodyLine,
		ParseErr: perr,
		Digest:   hex.EncodeToString(sum[:]),
	}, nil
}

// Strays returns repository-relative paths of .md files under research/
// that sit outside the designated record directories. Files under
// derivations/ subdirectories (human sketches) are intentionally ignored.
func Strays(root string) ([]string, error) {
	base := filepath.Join(root, "research")
	if _, err := os.Stat(base); os.IsNotExist(err) {
		return nil, nil
	}
	allowed := map[string]bool{}
	for _, d := range RecordDirs {
		allowed[d] = true
	}
	var out []string
	err := filepath.WalkDir(base, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		if strings.HasPrefix(d.Name(), ".") || !strings.HasSuffix(d.Name(), ".md") {
			return nil
		}
		rel, rerr := filepath.Rel(root, path)
		if rerr != nil {
			return rerr
		}
		rel = filepath.ToSlash(rel)
		dir := filepath.ToSlash(filepath.Dir(rel))
		if !allowed[dir] {
			out = append(out, rel)
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("store: scan research/: %w", err)
	}
	sort.Strings(out)
	return out, nil
}
