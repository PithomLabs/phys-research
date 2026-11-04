package record

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"testing"

	"github.com/PithomLabs/phys-research/internal/frontmatter"
	"github.com/PithomLabs/phys-research/internal/store"
)

func tempRoot(t *testing.T) string {
	t.Helper()
	return t.TempDir()
}

func TestCreateAllocatesIDs(t *testing.T) {
	root := tempRoot(t)

	c1, err := Create(root, Options{Type: TypeSource, Title: "First"})
	if err != nil {
		t.Fatalf("create 1: %v", err)
	}
	if c1.ID != "P-001" || c1.Rel != "research/sources/P-001.md" {
		t.Errorf("got %+v, want P-001 at research/sources/P-001.md", c1)
	}

	c2, err := Create(root, Options{Type: TypeSource, Title: "Second"})
	if err != nil {
		t.Fatalf("create 2: %v", err)
	}
	if c2.ID != "P-002" {
		t.Errorf("got %s, want P-002", c2.ID)
	}

	// Different types have independent prefixes.
	ci, err := Create(root, Options{Type: TypeIssue, Title: "Issue"})
	if err != nil {
		t.Fatalf("create issue: %v", err)
	}
	if ci.ID != "I-001" {
		t.Errorf("got %s, want I-001", ci.ID)
	}
}

func TestCreateSkipsExistingID(t *testing.T) {
	root := tempRoot(t)
	dir := filepath.Join(root, "research", "sources")
	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Fatal(err)
	}
	// Occupy P-001 manually; allocation must skip to P-002.
	if err := os.WriteFile(filepath.Join(dir, "P-001.md"), []byte("occupied\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	c, err := Create(root, Options{Type: TypeSource, Title: "Second"})
	if err != nil {
		t.Fatalf("create: %v", err)
	}
	if c.ID != "P-002" {
		t.Errorf("got %s, want P-002 (collision must be skipped)", c.ID)
	}
}

func TestCreateConcurrentDistinctIDs(t *testing.T) {
	root := tempRoot(t)
	const n = 8
	var wg sync.WaitGroup
	ids := make([]string, n)
	errs := make([]error, n)
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			c, err := Create(root, Options{Type: TypeEvidence, Title: fmt.Sprintf("E%d", i)})
			if err != nil {
				errs[i] = err
				return
			}
			ids[i] = c.ID
		}(i)
	}
	wg.Wait()

	seen := map[string]bool{}
	for i := 0; i < n; i++ {
		if errs[i] != nil {
			t.Fatalf("goroutine %d: %v", i, errs[i])
		}
		if seen[ids[i]] {
			t.Errorf("duplicate ID allocated: %s", ids[i])
		}
		seen[ids[i]] = true
	}
	if len(seen) != n {
		t.Errorf("expected %d distinct IDs, got %d", n, len(seen))
	}
}

func TestCreateOverwriteRefused(t *testing.T) {
	root := tempRoot(t)
	path := filepath.Join(root, "research", "sources", "P-001.md")
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, []byte("original"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := writeNew(path, []byte("replacement")); err == nil {
		t.Fatal("writeNew overwrote an existing file; it must refuse")
	} else if !os.IsExist(err) && err != os.ErrExist {
		t.Errorf("expected ErrExist, got %v", err)
	}
	got, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if string(got) != "original" {
		t.Errorf("file content changed: %q", got)
	}
}

func TestCreateValidation(t *testing.T) {
	root := tempRoot(t)
	cases := []struct {
		name string
		opts Options
		want string
	}{
		{"unknown type", Options{Type: "paper"}, "unknown record type"},
		{"hypothesis without class", Options{Type: TypeHypothesis}, "require --class"},
		{"hypothesis bad class", Options{Type: TypeHypothesis, Class: "empirical"}, "invalid class"},
		{"class on non-hypothesis", Options{Type: TypeSource, Class: "physical"}, "only valid on hypothesis"},
		{"bad branch", Options{Type: TypeIssue, Branch: "Z"}, "invalid branch"},
		{"bad priority", Options{Type: TypeIssue, Priority: "urgent"}, "invalid priority"},
		{"bad relation", Options{Type: TypeIssue, Refs: []frontmatter.Ref{{Relation: "likes", ID: "P-001"}}}, "invalid relation"},
		{"bad ref id", Options{Type: TypeIssue, Refs: []frontmatter.Ref{{Relation: "motivated_by", ID: "P-1"}}}, "invalid reference id"},
		{"review without target", Options{Type: TypeReview}, "require --review"},
		{"review flag on other type", Options{Type: TypeIssue, Review: "H-001"}, "only valid for --type review"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := Create(root, tc.opts)
			if err == nil {
				t.Fatal("expected an error, got nil")
			}
			if !strings.Contains(err.Error(), tc.want) {
				t.Errorf("error %q does not contain %q", err, tc.want)
			}
		})
	}
}

func TestCreateHypothesisClass(t *testing.T) {
	root := tempRoot(t)
	for _, class := range Classes {
		c, err := Create(root, Options{Type: TypeHypothesis, Title: class, Class: class})
		if err != nil {
			t.Fatalf("%s: %v", class, err)
		}
		raw, err := os.ReadFile(filepath.Join(root, filepath.FromSlash(c.Rel)))
		if err != nil {
			t.Fatal(err)
		}
		if !strings.Contains(string(raw), "class: "+class) {
			t.Errorf("%s: header missing class line:\n%s", class, raw)
		}
	}
	// Prefixes must differ per number: H-001 and H-002.
}

func TestCreateReviewRecordsDigest(t *testing.T) {
	root := tempRoot(t)
	h, err := Create(root, Options{Type: TypeHypothesis, Title: "H", Class: "physical"})
	if err != nil {
		t.Fatal(err)
	}
	recs, err := store.Load(root)
	if err != nil {
		t.Fatal(err)
	}
	var wantDigest string
	for _, r := range recs {
		if r.Header.ID == h.ID {
			wantDigest = r.Digest
		}
	}
	if wantDigest == "" {
		t.Fatal("hypothesis digest not found")
	}

	r, err := Create(root, Options{Type: TypeReview, Title: "R", Review: h.ID})
	if err != nil {
		t.Fatalf("create review: %v", err)
	}
	raw, err := os.ReadFile(filepath.Join(root, filepath.FromSlash(r.Rel)))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(raw), "reviewed: "+h.ID) {
		t.Errorf("review does not record target:\n%s", raw)
	}
	if !strings.Contains(string(raw), "digest: sha256:"+wantDigest) {
		t.Errorf("review does not record digest sha256:%s:\n%s", wantDigest, raw)
	}
}

func TestCreateReviewTargetNotFound(t *testing.T) {
	root := tempRoot(t)
	if _, err := Create(root, Options{Type: TypeReview, Title: "R", Review: "H-999"}); err == nil {
		t.Fatal("expected error for missing target")
	}
}

func TestCreateReviewTargetNotHypothesis(t *testing.T) {
	root := tempRoot(t)
	if _, err := Create(root, Options{Type: TypeSource, Title: "P"}); err != nil {
		t.Fatal(err)
	}
	if _, err := Create(root, Options{Type: TypeReview, Title: "R", Review: "P-001"}); err == nil {
		t.Fatal("expected error when review target is not a hypothesis")
	}
}

func TestTemplatesContainRequiredSections(t *testing.T) {
	cases := []struct {
		opts     Options
		sections []string
	}{
		{Options{Type: TypeIssue, Title: "T"}, []string{
			"## Theoretical expectation", "## Independent theoretical result",
			"## Experimental expectation", "## Experimental actual",
			"NOT APPLICABLE — no independent theoretical derivation is required for this experimental result.",
		}},
		{Options{Type: TypeHypothesis, Title: "T", Class: "physical"}, []string{"## Falsification condition"}},
		{Options{Type: TypeHypothesis, Title: "T", Class: "mathematical"}, []string{
			"## Premises", "## Conclusion", "## Domain", "## Proof obligations",
		}},
		{Options{Type: TypeSource, Title: "T"}, []string{"## Bibliographic identity", "doi:", "sha256:", "## Claims", "locator:"}},
		{Options{Type: TypeEvidence, Title: "T"}, []string{
			"## Measured value", "## Units and uncertainty", "## Methodology", "## Source locator",
		}},
		{Options{Type: TypeDerivation, Title: "T"}, []string{
			"## Premises", "## Dependencies", "## Intermediate results",
			"## Mathematical rules", "## Obligations", "## Checks",
		}},
		{Options{Type: TypeReview, Title: "T", Review: "H-001"}, []string{
			"## Review target", "## Blocking findings", "## Nonblocking findings", "## Evidential boundary",
		}},
	}
	for _, tc := range cases {
		name := tc.opts.Type
		if tc.opts.Class != "" {
			name += "/" + tc.opts.Class
		}
		t.Run(name, func(t *testing.T) {
			body := template(tc.opts, "X-001", "deadbeef")
			for _, s := range tc.sections {
				if !strings.Contains(body, s) {
					t.Errorf("template for %s missing %q", name, s)
				}
			}
		})
	}
}

func TestSortedIDs(t *testing.T) {
	root := tempRoot(t)
	for _, typ := range Types() {
		opts := Options{Type: typ, Title: typ}
		if typ == TypeHypothesis {
			opts.Class = "physical"
		}
		if typ == TypeReview {
			opts.Review = "H-001" // hypotheses are created before reviews
		}
		if _, err := Create(root, opts); err != nil {
			t.Fatalf("%s: %v", typ, err)
		}
	}
	ids, err := SortedIDs(root)
	if err != nil {
		t.Fatal(err)
	}
	if len(ids) != len(Types()) {
		t.Fatalf("got %v, want %d ids", ids, len(Types()))
	}
	if !sorted(ids) {
		t.Errorf("ids not sorted: %v", ids)
	}
}

func sorted(ids []string) bool {
	for i := 1; i < len(ids); i++ {
		if ids[i-1] > ids[i] {
			return false
		}
	}
	return true
}
