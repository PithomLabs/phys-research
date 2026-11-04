package status

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// digestOf returns the SHA-256 hex of a file, matching store.Record.Digest.
func digestOf(t *testing.T, path string) string {
	t.Helper()
	raw, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	sum := sha256.Sum256(raw)
	return hex.EncodeToString(sum[:])
}

func writeRec(t *testing.T, root, rel, header, body string) {
	t.Helper()
	path := filepath.Join(root, filepath.FromSlash(rel))
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, []byte("---\n"+header+"---\n\n"+body), 0o644); err != nil {
		t.Fatal(err)
	}
}

func buildRepo(t *testing.T) string {
	t.Helper()
	root := t.TempDir()
	writeRec(t, root, "research/sources/P-001.md",
		"id: P-001\ntype: source\nstate: open\n",
		"# P-001\n\n## Bibliographic identity\n\nx\n\n## Claims\n\n- c (locator: p.1)\n")
	writeRec(t, root, "research/issues/I-001.md",
		"id: I-001\ntype: issue\nstate: open\nbranch: C\npriority: high\nrefs:\n  - relation: motivated_by\n    id: P-001\n",
		"# I-001\n\n## Theoretical expectation\n\nx\n\n## Independent theoretical result\n\nNOT APPLICABLE — no independent theoretical derivation is required for this experimental result.\n\n## Experimental expectation\n\nx\n\n## Experimental actual\n\nx\n")
	writeRec(t, root, "research/issues/I-002.md",
		"id: I-002\ntype: issue\nstate: blocked\npriority: low\n",
		"# I-002\n\n## Theoretical expectation\n\nx\n\n## Independent theoretical result\n\nx\n\n## Experimental expectation\n\nx\n\n## Experimental actual\n\nx\n")
	writeRec(t, root, "research/issues/I-003.md",
		"id: I-003\ntype: issue\nstate: closed\n",
		"# I-003\n\n## Theoretical expectation\n\nx\n\n## Independent theoretical result\n\nx\n\n## Experimental expectation\n\nx\n\n## Experimental actual\n\nx\n")
	// A physical hypothesis with two unresolved ? obligations.
	writeRec(t, root, "research/hypotheses/H-001.md",
		"id: H-001\ntype: hypothesis\nstate: open\nclass: physical\n",
		"# H-001\n\nclaim wep_prediction: under WEP, eta(Ti,Pt) = 0\n?  is the systematic budget complete?\n?  is the null-chamber drift corrected?\n\n## Falsification condition\n\nx\n")
	// A stale review: digest does not match H-001.
	writeRec(t, root, "research/reviews/R-001.md",
		"id: R-001\ntype: review\nstate: open\n",
		"# R-001\n\n## Review target\n\nreviewed: H-001\ndigest: sha256:"+strings.Repeat("0", 64)+"\nreviewer: adversary\nmodel: m\n\n## Blocking findings\n\nnone\n\n## Nonblocking findings\n\nnone\n\n## Evidential boundary\n\nx\n")
	// A fresh review with a blocking finding.
	writeRec(t, root, "research/hypotheses/H-002.md",
		"id: H-002\ntype: hypothesis\nstate: open\nclass: physical\n",
		"# H-002\n\nclaim x: y\n\n## Falsification condition\n\nz\n")
	sum := digestOf(t, filepath.Join(root, "research", "hypotheses", "H-002.md"))
	writeRec(t, root, "research/reviews/R-002.md",
		"id: R-002\ntype: review\nstate: open\n",
		"# R-002\n\n## Review target\n\nreviewed: H-002\ndigest: sha256:"+sum+"\nreviewer: adversary\nmodel: m\n\n## Blocking findings\n\n- The falsification condition does not declare a total-uncertainty prescription.\n\n## Nonblocking findings\n\nnone\n\n## Evidential boundary\n\nx\n")
	return root
}

func TestCollectShape(t *testing.T) {
	root := buildRepo(t)
	out, err := Collect(root, "")
	if err != nil {
		t.Fatal(err)
	}

	if len(out.Issues) != 2 {
		t.Errorf("issues = %+v, want I-001 and I-002 (closed excluded)", out.Issues)
	}
	if len(out.Obligations) != 2 {
		t.Errorf("obligations = %+v, want 2", out.Obligations)
	}
	for _, ob := range out.Obligations {
		if ob.ID != "H-001" {
			t.Errorf("obligation from unexpected record: %+v", ob)
		}
	}
	if len(out.StaleReviews) != 1 || out.StaleReviews[0].ID != "R-001" {
		t.Errorf("stale = %+v, want R-001", out.StaleReviews)
	}
	if len(out.BlockingFindings) != 1 || out.BlockingFindings[0].Review != "R-002" {
		t.Errorf("blocking = %+v, want R-002", out.BlockingFindings)
	}
}

func TestJSONShape(t *testing.T) {
	root := buildRepo(t)
	out, err := Collect(root, "")
	if err != nil {
		t.Fatal(err)
	}
	raw, err := json.Marshal(out)
	if err != nil {
		t.Fatal(err)
	}
	var m map[string]json.RawMessage
	if err := json.Unmarshal(raw, &m); err != nil {
		t.Fatal(err)
	}
	for _, key := range []string{"issues", "obligations", "stale_reviews", "blocking_findings"} {
		if _, ok := m[key]; !ok {
			t.Errorf("JSON missing key %q: %s", key, raw)
		}
	}
	// Round-trip must not lose entries.
	var back Output
	if err := json.Unmarshal(raw, &back); err != nil {
		t.Fatal(err)
	}
	if len(back.Issues) != len(out.Issues) || len(back.Obligations) != len(out.Obligations) {
		t.Errorf("round-trip lost data: %+v", back)
	}
}

func TestPriorityOrdering(t *testing.T) {
	root := t.TempDir()
	body := "# T\n\n## Theoretical expectation\n\nx\n\n## Independent theoretical result\n\nx\n\n## Experimental expectation\n\nx\n\n## Experimental actual\n\nx\n"
	// Create in reverse priority order to prove sorting.
	writeRec(t, root, "research/issues/I-003.md", "id: I-003\ntype: issue\nstate: open\npriority: low\n", body)
	writeRec(t, root, "research/issues/I-002.md", "id: I-002\ntype: issue\nstate: open\npriority: medium\n", body)
	writeRec(t, root, "research/issues/I-001.md", "id: I-001\ntype: issue\nstate: open\npriority: high\n", body)
	writeRec(t, root, "research/issues/I-004.md", "id: I-004\ntype: issue\nstate: open\n", body)

	out, err := Collect(root, "")
	if err != nil {
		t.Fatal(err)
	}
	var ids []string
	for _, is := range out.Issues {
		ids = append(ids, is.ID)
	}
	want := []string{"I-001", "I-002", "I-003", "I-004"}
	for i, id := range want {
		if ids[i] != id {
			t.Errorf("order = %v, want %v", ids, want)
			break
		}
	}
}

func TestBranchFilter(t *testing.T) {
	root := buildRepo(t)
	out, err := Collect(root, "C")
	if err != nil {
		t.Fatal(err)
	}
	if len(out.Issues) != 1 || out.Issues[0].ID != "I-001" {
		t.Errorf("branch C filter = %+v, want only I-001", out.Issues)
	}
	out, err = Collect(root, "Q")
	if err != nil {
		t.Fatal(err)
	}
	if len(out.Issues) != 0 {
		t.Errorf("branch Q filter = %+v, want none", out.Issues)
	}
}

func TestTextReport(t *testing.T) {
	root := buildRepo(t)
	out, err := Collect(root, "")
	if err != nil {
		t.Fatal(err)
	}
	text := Text(out)
	for _, want := range []string{
		"Open and blocked issues:",
		"Unresolved ? obligations:",
		"Stale reviews:",
		"Blocking findings:",
		"I-001",
		"R-001",
		"R-002",
		"? is the systematic budget complete?",
	} {
		if !strings.Contains(text, want) {
			t.Errorf("text report missing %q:\n%s", want, text)
		}
	}
}

func TestObligationLineNumbers(t *testing.T) {
	root := t.TempDir()
	body := "# H\n\nline one\n?  first obligation\n\nmore\n?  second obligation\n"
	writeRec(t, root, "research/hypotheses/H-001.md",
		"id: H-001\ntype: hypothesis\nstate: open\nclass: physical\n", body)
	out, err := Collect(root, "")
	if err != nil {
		t.Fatal(err)
	}
	if len(out.Obligations) != 2 {
		t.Fatalf("obligations = %+v", out.Obligations)
	}
	// File lines: 1 ---, 2-5 header, 6 ---, 7 blank, 8 # H, 9 blank,
	// 10 line one, 11 first ?, 12 blank, 13 more, 14 second ?.
	if out.Obligations[0].Line != 11 {
		t.Errorf("first obligation line = %d, want 11", out.Obligations[0].Line)
	}
	if out.Obligations[1].Line != 14 {
		t.Errorf("second obligation line = %d, want 14", out.Obligations[1].Line)
	}
}

func TestNoReviewsNoStaleOrBlocking(t *testing.T) {
	root := t.TempDir()
	out, err := Collect(root, "")
	if err != nil {
		t.Fatal(err)
	}
	if len(out.Issues) != 0 || len(out.Obligations) != 0 ||
		len(out.StaleReviews) != 0 || len(out.BlockingFindings) != 0 {
		t.Errorf("empty repo should report nothing: %+v", out)
	}
	text := Text(out)
	if !strings.Contains(text, "(none)") {
		t.Errorf("empty sections should say (none):\n%s", text)
	}
}
