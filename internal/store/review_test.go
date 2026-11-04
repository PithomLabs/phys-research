package store

import (
	"strings"
	"testing"
)

const (
	hexA = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	hexB = "bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb"
	hexC = "cccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccc"
	hexZ = "0000000000000000000000000000000000000000000000000000000000000000"
)

// lookupOf builds a resolver over "current" digests.
func lookupOf(digests map[string]string) func(string) (*Record, bool) {
	return func(id string) (*Record, bool) {
		d, ok := digests[id]
		if !ok {
			return nil, false
		}
		return &Record{Digest: d}, true
	}
}

func provenance(body string) ReviewProvenance {
	return ParseReview(&Record{Body: []byte(body), BodyLine: 10})
}

func TestParseReview(t *testing.T) {
	body := "\n## Review target\n\nreviewed: H-001\ndigest: sha256:" + hexA +
		"\ninput: P-001 sha256:" + hexB +
		"\ninput: E-001 sha256:" + hexC +
		"\nreviewer: adversary\nmodel: some-model\n"
	p := provenance(body)

	if p.Reviewed != "H-001" {
		t.Errorf("Reviewed = %q, want H-001", p.Reviewed)
	}
	// Lines are file-absolute: body starts at line 10.
	if p.ReviewedLine != 13 {
		t.Errorf("ReviewedLine = %d, want 13", p.ReviewedLine)
	}
	if p.DigestLine != 14 {
		t.Errorf("DigestLine = %d, want 14", p.DigestLine)
	}
	if len(p.Inputs) != 2 {
		t.Fatalf("Inputs = %+v, want 2", p.Inputs)
	}
	if p.Inputs[0].ID != "P-001" || p.Inputs[0].Line != 15 || p.Inputs[0].Digest != "sha256:"+hexB {
		t.Errorf("Inputs[0] = %+v", p.Inputs[0])
	}
	if p.Inputs[1].ID != "E-001" || p.Inputs[1].Line != 16 {
		t.Errorf("Inputs[1] = %+v", p.Inputs[1])
	}
}

func TestParseReviewMalformedInputLines(t *testing.T) {
	cases := []struct {
		name string
		line string
	}{
		{"missing digest", "input: P-001"},
		{"bad digest", "input: P-001 sha256:nothex"},
		{"empty", "input:"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			p := provenance("reviewed: H-001\ndigest: sha256:" + hexA + "\n" + tc.line + "\n")
			if len(p.Inputs) != 1 {
				t.Fatalf("Inputs = %+v, want the line recorded", p.Inputs)
			}
			if p.Inputs[0].Err == "" {
				t.Errorf("malformed line not flagged: %+v", p.Inputs[0])
			}
		})
	}
}

func TestReviewIssuesFresh(t *testing.T) {
	p := provenance("reviewed: H-001\ndigest: sha256:" + hexA +
		"\ninput: P-001 sha256:" + hexB + "\n")
	issues, unpinned := p.ReviewIssues(lookupOf(map[string]string{
		"H-001": hexA, "P-001": hexB,
	}))
	if len(issues) != 0 {
		t.Errorf("issues = %+v, want none", issues)
	}
	if unpinned {
		t.Error("review pins an input but was reported unpinned")
	}
}

func TestReviewIssuesTargetStale(t *testing.T) {
	p := provenance("reviewed: H-001\ndigest: sha256:" + hexA + "\n")
	issues, _ := p.ReviewIssues(lookupOf(map[string]string{"H-001": hexZ}))
	if len(issues) != 1 || issues[0].Kind != IssueTargetStale {
		t.Fatalf("issues = %+v, want one target-stale", issues)
	}
	if issues[0].Want != "sha256:"+hexA || issues[0].Have != hexZ {
		t.Errorf("want/have = %q/%q", issues[0].Want, issues[0].Have)
	}
	if !strings.Contains(issues[0].Message(), "STALE") {
		t.Errorf("message = %q", issues[0].Message())
	}
}

func TestReviewIssuesInputStaleNamesTheInput(t *testing.T) {
	p := provenance("reviewed: H-001\ndigest: sha256:" + hexA +
		"\ninput: P-001 sha256:" + hexB +
		"\ninput: E-001 sha256:" + hexC + "\n")
	issues, unpinned := p.ReviewIssues(lookupOf(map[string]string{
		"H-001": hexA, "P-001": hexB, "E-001": hexZ, // only E-001 changed
	}))
	if unpinned {
		t.Error("review pins inputs but was reported unpinned")
	}
	if len(issues) != 1 || issues[0].Kind != IssueInputStale {
		t.Fatalf("issues = %+v, want one input-stale", issues)
	}
	if issues[0].ID != "E-001" {
		t.Errorf("issue names %q, want E-001", issues[0].ID)
	}
	if !strings.Contains(issues[0].Message(), "material input E-001") {
		t.Errorf("message = %q", issues[0].Message())
	}
	if issues[0].Rule() != "stale" {
		t.Errorf("rule = %q, want stale", issues[0].Rule())
	}
}

func TestReviewIssuesInputMissing(t *testing.T) {
	p := provenance("reviewed: H-001\ndigest: sha256:" + hexA +
		"\ninput: P-999 sha256:" + hexB + "\n")
	issues, _ := p.ReviewIssues(lookupOf(map[string]string{"H-001": hexA}))
	if len(issues) != 1 || issues[0].Kind != IssueInputMissing {
		t.Fatalf("issues = %+v, want one input-missing", issues)
	}
	if issues[0].ID != "P-999" {
		t.Errorf("issue names %q, want P-999", issues[0].ID)
	}
}

func TestReviewIssuesTargetMissing(t *testing.T) {
	p := provenance("reviewed: H-999\ndigest: sha256:" + hexA + "\n")
	issues, _ := p.ReviewIssues(lookupOf(map[string]string{}))
	found := false
	for _, is := range issues {
		if is.Kind == IssueTargetMissing && is.ID == "H-999" {
			found = true
		}
	}
	if !found {
		t.Errorf("issues = %+v, want target-missing for H-999", issues)
	}
}

func TestReviewIssuesLegacyUnpinnedIsNotStale(t *testing.T) {
	// R-001's shape: correct target digest, no input lines at all.
	p := provenance("reviewed: H-001\ndigest: sha256:" + hexA + "\n")
	issues, unpinned := p.ReviewIssues(lookupOf(map[string]string{"H-001": hexA}))
	if len(issues) != 0 {
		t.Errorf("issues = %+v, want none: an unpinned review is not stale", issues)
	}
	if !unpinned {
		t.Error("review with no inputs must be reported unpinned")
	}
	if !strings.Contains(LegacyUnpinnedWarning, "LEGACY_UNPINNED") {
		t.Errorf("warning = %q", LegacyUnpinnedWarning)
	}
}

func TestReviewIssuesMalformedInput(t *testing.T) {
	p := provenance("reviewed: H-001\ndigest: sha256:" + hexA + "\ninput: P-001\n")
	issues, _ := p.ReviewIssues(lookupOf(map[string]string{"H-001": hexA}))
	if len(issues) != 1 || issues[0].Kind != IssueInputMalformed {
		t.Fatalf("issues = %+v, want one input-malformed", issues)
	}
	if issues[0].Msg == "" {
		t.Error("malformed input issue carries no detail")
	}
	if issues[0].Rule() != "review" {
		t.Errorf("rule = %q, want review", issues[0].Rule())
	}
}

func TestReasonIsHumanReadable(t *testing.T) {
	for _, kind := range []ReviewIssueKind{
		IssueTargetMissing, IssueTargetMalformed, IssueTargetStale,
		IssueInputMalformed, IssueInputMissing, IssueInputStale,
	} {
		is := ReviewIssue{Kind: kind, ID: "P-001"}
		if is.Reason() == "" || is.Reason() == "unknown provenance problem" {
			t.Errorf("%s: reason = %q", kind, is.Reason())
		}
	}
}
