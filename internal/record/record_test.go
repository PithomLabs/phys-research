package record

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"testing"

	"github.com/PithomLabs/phys-research/internal/frontmatter"
	"github.com/PithomLabs/phys-research/internal/store"
)

// fileDigest returns the SHA-256 hex of a file, matching store.Record.Digest.
func fileDigest(t *testing.T, root, rel string) string {
	t.Helper()
	raw, err := os.ReadFile(filepath.Join(root, filepath.FromSlash(rel)))
	if err != nil {
		t.Fatal(err)
	}
	sum := sha256.Sum256(raw)
	return hex.EncodeToString(sum[:])
}

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
			body := template(tc.opts, "X-001", "deadbeef", nil)
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

// buildGraph creates the canonical one-hop graph used by the material-input
// tests: P-001 (source), I-001 (issue → P-001), H-001 (hypothesis → I-001,
// P-001) and E-001 (evidence → H-001).
func buildGraph(t *testing.T) string {
	t.Helper()
	root := tempRoot(t)
	if _, err := Create(root, Options{Type: TypeSource, Title: "Source"}); err != nil {
		t.Fatal(err)
	}
	if _, err := Create(root, Options{Type: TypeIssue, Title: "Issue",
		Refs: []frontmatter.Ref{{Relation: "motivated_by", ID: "P-001"}}}); err != nil {
		t.Fatal(err)
	}
	if _, err := Create(root, Options{Type: TypeHypothesis, Title: "Hyp", Class: "physical",
		Refs: []frontmatter.Ref{{Relation: "addresses", ID: "I-001"}, {Relation: "depends_on", ID: "P-001"}}}); err != nil {
		t.Fatal(err)
	}
	if _, err := Create(root, Options{Type: TypeEvidence, Title: "Ev",
		Refs: []frontmatter.Ref{{Relation: "supports", ID: "H-001"}}}); err != nil {
		t.Fatal(err)
	}
	return root
}

func TestReviewTemplatePinsMaterialInputs(t *testing.T) {
	root := buildGraph(t)
	c, err := Create(root, Options{Type: TypeReview, Title: "R", Review: "H-001"})
	if err != nil {
		t.Fatalf("create review: %v", err)
	}
	raw, err := os.ReadFile(filepath.Join(root, filepath.FromSlash(c.Rel)))
	if err != nil {
		t.Fatal(err)
	}
	body := string(raw)

	hDigest := fileDigest(t, root, "research/hypotheses/H-001.md")
	if !strings.Contains(body, "reviewed: H-001") {
		t.Errorf("review does not record its target:\n%s", body)
	}
	if !strings.Contains(body, "digest: sha256:"+hDigest) {
		t.Errorf("review does not record digest sha256:%s:\n%s", hDigest, body)
	}

	// One input line per one-hop neighbour (both directions, reviews
	// excluded), sorted by ID, each pinning that file's exact bytes.
	want := []string{
		"input: E-001 sha256:" + fileDigest(t, root, "research/evidence/E-001.md"),
		"input: I-001 sha256:" + fileDigest(t, root, "research/issues/I-001.md"),
		"input: P-001 sha256:" + fileDigest(t, root, "research/sources/P-001.md"),
	}
	lines := strings.Split(body, "\n")
	var got []string
	digestAt, reviewerAt := -1, -1
	for i, ln := range lines {
		switch {
		case strings.HasPrefix(ln, "input: "):
			got = append(got, ln)
		case strings.HasPrefix(ln, "digest: sha256:"):
			digestAt = i
		case strings.HasPrefix(ln, "reviewer:"):
			reviewerAt = i
		}
	}
	if len(got) != len(want) {
		t.Fatalf("input lines = %v, want %v", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("input line %d = %q, want %q", i, got[i], want[i])
		}
	}
	// Ordering: inputs sit between the digest and the reviewer identity.
	digestOff := strings.Index(body, "digest: sha256:")
	reviewerOff := strings.Index(body, "reviewer:")
	firstInputOff := strings.Index(body, want[0])
	lastInputOff := strings.Index(body, want[len(want)-1])
	if digestAt < 0 || reviewerAt < 0 || digestOff < 0 || reviewerOff < 0 ||
		!(digestOff < firstInputOff && lastInputOff < reviewerOff) {
		t.Errorf("input lines are not placed between digest and reviewer (digest line %d, reviewer line %d)", digestAt, reviewerAt)
	}
}

func TestCreateReviewFailsOnMissingMaterialInput(t *testing.T) {
	root := tempRoot(t)
	// H-001 declares a neighbour that does not exist anywhere.
	if _, err := Create(root, Options{Type: TypeHypothesis, Title: "H", Class: "physical",
		Refs: []frontmatter.Ref{{Relation: "depends_on", ID: "P-999"}}}); err != nil {
		t.Fatal(err)
	}
	_, err := Create(root, Options{Type: TypeReview, Title: "R", Review: "H-001"})
	if err == nil {
		t.Fatal("review creation must fail when a material input cannot be read")
	}
	if !strings.Contains(err.Error(), "material input P-999 not found") {
		t.Errorf("diagnostic = %q, want an explicit 'material input P-999 not found'", err)
	}
	entries, _ := os.ReadDir(filepath.Join(root, "research", "reviews"))
	if len(entries) != 0 {
		t.Errorf("a review file was written despite the failure: %v", entries)
	}
}

func TestCreateReviewFailsOnUnparsableMaterialInput(t *testing.T) {
	root := tempRoot(t)
	// P-001 exists but its frontmatter is rejected by the parser.
	path := filepath.Join(root, "research", "sources", "P-001.md")
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, []byte("---\nid: P-001\ncolor: red\n---\n\n# broken\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if _, err := Create(root, Options{Type: TypeHypothesis, Title: "H", Class: "physical",
		Refs: []frontmatter.Ref{{Relation: "depends_on", ID: "P-001"}}}); err != nil {
		t.Fatal(err)
	}
	_, err := Create(root, Options{Type: TypeReview, Title: "R", Review: "H-001"})
	if err == nil {
		t.Fatal("review creation must fail when a material input cannot be parsed")
	}
	if !strings.Contains(err.Error(), "material input P-001 could not be parsed") {
		t.Errorf("diagnostic = %q, want an explicit 'could not be parsed'", err)
	}
	entries, _ := os.ReadDir(filepath.Join(root, "research", "reviews"))
	if len(entries) != 0 {
		t.Errorf("a review file was written despite the failure: %v", entries)
	}
}

func TestReviewNeighbourExcludedFromMaterialInputs(t *testing.T) {
	root := buildGraph(t)
	// An earlier review declaring a header ref at H-001: excluded by rule,
	// and its presence must not fail creation of the next review.
	if _, err := Create(root, Options{Type: TypeReview, Title: "R1", Review: "H-001",
		Refs: []frontmatter.Ref{{Relation: "challenges", ID: "H-001"}}}); err != nil {
		t.Fatalf("create first review: %v", err)
	}
	c, err := Create(root, Options{Type: TypeReview, Title: "R2", Review: "H-001"})
	if err != nil {
		t.Fatalf("create second review: %v", err)
	}
	raw, err := os.ReadFile(filepath.Join(root, filepath.FromSlash(c.Rel)))
	if err != nil {
		t.Fatal(err)
	}
	body := string(raw)
	if strings.Contains(body, "input: R-001") {
		t.Errorf("a review neighbour must be excluded from material inputs:\n%s", body)
	}
	for _, id := range []string{"E-001", "I-001", "P-001"} {
		if !strings.Contains(body, "input: "+id+" sha256:") {
			t.Errorf("material input %s missing:\n%s", id, body)
		}
	}
}

func TestMaterialInputsBothDirections(t *testing.T) {
	target := &store.Record{
		Rel:    "research/hypotheses/H-001.md",
		Header: frontmatter.Header{ID: "H-001", Type: TypeHypothesis, Refs: []frontmatter.Ref{{Relation: "addresses", ID: "I-001"}, {Relation: "depends_on", ID: "P-001"}}},
		Digest: "hh",
	}
	recs := []store.Record{
		*target,
		{Rel: "research/sources/P-001.md", Header: frontmatter.Header{ID: "P-001", Type: TypeSource}, Digest: "pp"},
		{Rel: "research/issues/I-001.md", Header: frontmatter.Header{ID: "I-001", Type: TypeIssue, Refs: []frontmatter.Ref{{Relation: "motivated_by", ID: "P-001"}}}, Digest: "ii"},
		{Rel: "research/evidence/E-001.md", Header: frontmatter.Header{ID: "E-001", Type: TypeEvidence, Refs: []frontmatter.Ref{{Relation: "supports", ID: "H-001"}}}, Digest: "ee"},
		{Rel: "research/reviews/R-001.md", Header: frontmatter.Header{ID: "R-001", Type: TypeReview, Refs: []frontmatter.Ref{{Relation: "challenges", ID: "H-001"}}}, Digest: "rr"},
		{Rel: "research/issues/I-002.md", Header: frontmatter.Header{ID: "I-002", Type: TypeIssue}, Digest: "nn"}, // unrelated
	}

	got, err := materialInputs(t.TempDir(), target, recs)
	if err != nil {
		t.Fatalf("materialInputs: %v", err)
	}
	want := []store.PinnedInput{
		{ID: "E-001", Digest: "ee"},
		{ID: "I-001", Digest: "ii"},
		{ID: "P-001", Digest: "pp"},
	}
	if len(got) != len(want) {
		t.Fatalf("inputs = %+v, want %+v", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("input %d = %+v, want %+v", i, got[i], want[i])
		}
	}
}
