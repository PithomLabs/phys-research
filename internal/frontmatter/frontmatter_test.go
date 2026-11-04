package frontmatter

import (
	"strings"
	"testing"
)

func parseErr(t *testing.T, content string) *SyntaxError {
	t.Helper()
	_, _, _, err := Parse([]byte(content))
	if err == nil {
		t.Fatal("expected a parse error, got nil")
	}
	se, ok := err.(*SyntaxError)
	if !ok {
		t.Fatalf("expected *SyntaxError, got %T: %v", err, err)
	}
	return se
}

func TestParseValid(t *testing.T) {
	content := "---\nid: I-001\ntype: issue\nstate: open\nbranch: C\npriority: high\nrefs:\n  - relation: motivated_by\n    id: P-001\n  - relation: addresses\n    id: H-002\n---\n\n# Body\n\nSome text.\n"
	h, body, _, err := Parse([]byte(content))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if h.ID != "I-001" || h.Type != "issue" || h.State != "open" || h.Branch != "C" || h.Priority != "high" {
		t.Errorf("header mismatch: %+v", h)
	}
	if len(h.Refs) != 2 {
		t.Fatalf("expected 2 refs, got %d", len(h.Refs))
	}
	if h.Refs[0].Relation != "motivated_by" || h.Refs[0].ID != "P-001" {
		t.Errorf("ref[0] mismatch: %+v", h.Refs[0])
	}
	if h.Refs[1].Relation != "addresses" || h.Refs[1].ID != "H-002" {
		t.Errorf("ref[1] mismatch: %+v", h.Refs[1])
	}
	if string(body) != "\n# Body\n\nSome text.\n" {
		t.Errorf("body mismatch: %q", body)
	}
}

func TestParseMinimal(t *testing.T) {
	h, body, _, err := Parse([]byte("---\nid: P-001\ntype: source\nstate: open\n---\nbody\n"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if h.ID != "P-001" || h.Branch != "" || h.Priority != "" || h.Class != "" || len(h.Refs) != 0 {
		t.Errorf("header mismatch: %+v", h)
	}
	if string(body) != "body\n" {
		t.Errorf("body mismatch: %q", body)
	}
}

func TestParseEmptyRefs(t *testing.T) {
	h, _, _, err := Parse([]byte("---\nid: P-001\ntype: source\nstate: open\nrefs:\n---\n\nbody\n"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(h.Refs) != 0 {
		t.Errorf("expected no refs, got %d", len(h.Refs))
	}
}

func TestParseCommentsAndBlankLines(t *testing.T) {
	h, _, _, err := Parse([]byte("---\n# a comment\n\nid: P-001\n\ntype: source\nstate: open\n---\nbody\n"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if h.ID != "P-001" {
		t.Errorf("id mismatch: %q", h.ID)
	}
}

func TestParseQuotedValues(t *testing.T) {
	h, _, _, err := Parse([]byte("---\nid: \"P-001\"\ntype: 'source'\nstate: open\n---\nbody\n"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if h.ID != "P-001" || h.Type != "source" {
		t.Errorf("quoted values not stripped: %+v", h)
	}
}

func TestParseErrors(t *testing.T) {
	cases := []struct {
		name    string
		content string
		line    int
	}{
		{"no opening delimiter", "id: P-001\n", 1},
		{"unterminated", "---\nid: P-001\n", 1},
		{"unknown field", "---\nid: P-001\ntype: source\nstate: open\ncolor: red\n---\n", 5},
		{"duplicate field", "---\nid: P-001\nid: P-002\ntype: source\nstate: open\n---\n", 3},
		{"tab", "---\nid: P-001\n\ttype: source\nstate: open\n---\n", 3},
		{"anchor", "---\nid: &x P-001\ntype: source\nstate: open\n---\n", 2},
		{"alias", "---\nid: *x\ntype: source\nstate: open\n---\n", 2},
		{"multiline pipe", "---\nid: |\n  P-001\ntype: source\nstate: open\n---\n", 2},
		{"multiline gt", "---\nid: >\n  P-001\ntype: source\nstate: open\n---\n", 2},
		{"flow list", "---\nid: P-001\ntype: source\nstate: open\nrefs: [P-002]\n---\n", 5},
		{"flow map", "---\nid: P-001\ntype: source\nstate: open\nrefs: {a: b}\n---\n", 5},
		{"missing space after colon", "---\nid:P-001\ntype: source\nstate: open\n---\n", 2},
		{"not a key value", "---\nid: P-001\ngarbage line\ntype: source\nstate: open\n---\n", 3},
		{"indented outside refs", "---\nid: P-001\n  type: source\nstate: open\n---\n", 3},
		{"item indent not 2", "---\nid: P-001\ntype: source\nstate: open\nrefs:\n   - relation: supports\n     id: P-002\n---\n", 6},
		{"key indent not 4", "---\nid: P-001\ntype: source\nstate: open\nrefs:\n  - relation: supports\n    id: P-002\n      extra: 1\n---\n", 8},
		{"unknown refs field", "---\nid: P-001\ntype: source\nstate: open\nrefs:\n  - color: blue\n    id: P-002\n---\n", 6},
		{"refs item missing id", "---\nid: P-001\ntype: source\nstate: open\nrefs:\n  - relation: supports\n---\n", 6},
		{"refs item missing relation", "---\nid: P-001\ntype: source\nstate: open\nrefs:\n  - id: P-002\n---\n", 6},
		{"duplicate relation", "---\nid: P-001\ntype: source\nstate: open\nrefs:\n  - relation: supports\n    relation: challenges\n    id: P-002\n---\n", 7},
		{"malformed quoted value", "---\nid: \"P-001\ntype: source\nstate: open\n---\n", 2},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			se := parseErr(t, tc.content)
			if se.Line != tc.line {
				t.Errorf("expected line %d, got %d (%s)", tc.line, se.Line, se.Msg)
			}
		})
	}
}

func TestParseCRLF(t *testing.T) {
	h, body, _, err := Parse([]byte("---\r\nid: P-001\r\ntype: source\r\nstate: open\r\n---\r\n\r\nbody\r\n"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if h.ID != "P-001" {
		t.Errorf("id mismatch: %q", h.ID)
	}
	if strings.Contains(string(body), "\r") {
		t.Errorf("body still contains CR: %q", body)
	}
}

func TestParseUnicodeBodyPreserved(t *testing.T) {
	_, body, _, err := Parse([]byte("---\nid: P-001\ntype: source\nstate: open\n---\n\nη(Ti,Pt) = −1.5 ± 2.3 × 10⁻¹⁵\n"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(string(body), "η(Ti,Pt)") {
		t.Errorf("unicode body mangled: %q", body)
	}
}

func TestParseBodyUntouched(t *testing.T) {
	bodyText := "\n# H-001\n\nclaim wep_prediction: under WEP, eta(Ti,Pt) = 0\n?  is the Eöt-Wash reanalysis systematic budget compatible?\n"
	_, body, _, err := Parse([]byte("---\nid: H-001\ntype: hypothesis\nstate: open\nclass: physical\n---\n" + bodyText))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if string(body) != bodyText {
		t.Errorf("body was modified:\n got %q\nwant %q", body, bodyText)
	}
}

func TestBodyLineIsFileAbsolute(t *testing.T) {
	// The closing --- is on file line 5, so the body starts on line 6.
	_, _, bodyLine, err := Parse([]byte("---\nid: P-001\ntype: source\nstate: open\n---\nbody\n"))
	if err != nil {
		t.Fatal(err)
	}
	if bodyLine != 6 {
		t.Errorf("bodyLine = %d, want 6", bodyLine)
	}

	// An empty header shifts nothing: closing --- on line 2, body on line 3.
	_, _, bodyLine, err = Parse([]byte("---\n---\nbody\n"))
	if err != nil {
		t.Fatal(err)
	}
	if bodyLine != 3 {
		t.Errorf("bodyLine = %d, want 3", bodyLine)
	}
}

func TestAllowedKeysClosed(t *testing.T) {
	want := []string{"id", "type", "state", "branch", "priority", "class", "refs"}
	if len(AllowedKeys) != len(want) {
		t.Fatalf("AllowedKeys = %v, want %v", AllowedKeys, want)
	}
	for i, k := range want {
		if AllowedKeys[i] != k {
			t.Errorf("AllowedKeys[%d] = %q, want %q", i, AllowedKeys[i], k)
		}
	}
}
