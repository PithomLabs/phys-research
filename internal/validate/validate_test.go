package validate

import (
	"crypto/sha256"
	"encoding/hex"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/PithomLabs/phys-research/internal/frontmatter"
)

// writeRec writes a record file at rel with the given header fields and body.
func writeRec(t *testing.T, root, rel string, header map[string]string, body string) {
	t.Helper()
	var b strings.Builder
	b.WriteString("---\n")
	// Stable order for required fields.
	for _, k := range []string{"id", "type", "state", "branch", "priority", "class"} {
		if v, ok := header[k]; ok {
			b.WriteString(k + ": " + v + "\n")
		}
	}
	if refs, ok := header["refs"]; ok {
		b.WriteString(refs)
	}
	b.WriteString("---\n\n")
	b.WriteString(body)
	path := filepath.Join(root, filepath.FromSlash(rel))
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, []byte(b.String()), 0o644); err != nil {
		t.Fatal(err)
	}
}

func findRule(res Result, rule, substr string) bool {
	for _, f := range res.Findings {
		if f.Rule == rule && strings.Contains(f.Message, substr) {
			return true
		}
	}
	return false
}

// digestOf returns the SHA-256 hex of a file, as store.Record.Digest does.
func digestOf(t *testing.T, path string) string {
	t.Helper()
	raw, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	sum := sha256.Sum256(raw)
	return hex.EncodeToString(sum[:])
}

func cleanIssueBody() string {
	return `# T

## Question

Is the equivalence principle violated?

## Theoretical expectation

- Claim or prediction: WEP holds
- Assumptions: local Lorentz invariance
- Applicable regime: solar system
- Source or derivation reference: none

## Independent theoretical result

NOT APPLICABLE — no independent theoretical derivation is required for this experimental result.

## Experimental expectation

- Predicted observable: eta(Ti,Pt)
- Units and conditions: dimensionless
- Uncertainty: 1e-15
- Falsification condition: |eta| > 0 at 3 sigma

## Experimental actual

- Measured result: eta = -1.5e-15
- Units and uncertainty: dimensionless, stat 2.3e-15, syst 1.5e-15
- Methodology and experimental conditions: torsion balance
- Exact source locator: Eq. (5) of P-001
`
}

func cleanPhysicalHypothesis() string {
	return `# T

given apparatus MICROSCOPE torsion balance
define eta(Ti,Pt) as the differential free-fall acceleration parameter
assume the WEP holds for the Ti/Pt composition pair
claim wep_prediction: under WEP, eta(Ti,Pt) = 0
derive measured_eta from the published fit of Eq. (5)
if measured_eta differs from 0 beyond the declared total uncertainty then the WEP is falsified
?  is the systematic budget of the 2022 run fully characterized?

## Falsification condition

A statistically significant non-zero eta(Ti,Pt), using a declared total-uncertainty
prescription (stat and syst combined in quadrature), would falsify the WEP prediction.
`
}

func TestRunClean(t *testing.T) {
	root := t.TempDir()
	writeRec(t, root, "research/sources/P-001.md", map[string]string{
		"id": "P-001", "type": "source", "state": "open",
	}, `# T

## Bibliographic identity

title: T
authors: A
doi: 10.1103/PhysRevLett.129.121102
arxiv: 2209.15487v1
sha256: `+strings.Repeat("a", 64)+`

## Claims

- WEP tested with Ti and Pt (locator: Eq. (1))
- eta(Ti,Pt) measured (locator: Fig. 1)
`)
	writeRec(t, root, "research/issues/I-001.md", map[string]string{
		"id": "I-001", "type": "issue", "state": "open", "branch": "C", "priority": "high",
		"refs": "refs:\n  - relation: motivated_by\n    id: P-001\n",
	}, cleanIssueBody())
	writeRec(t, root, "research/hypotheses/H-001.md", map[string]string{
		"id": "H-001", "type": "hypothesis", "state": "open", "class": "physical",
		"refs": "refs:\n  - relation: addresses\n    id: I-001\n",
	}, cleanPhysicalHypothesis())

	res, err := Run(root)
	if err != nil {
		t.Fatalf("Run: %v", err)
	}
	for _, f := range res.Findings {
		t.Errorf("unexpected finding: [%s] %s: %s", f.Severity, f.Rule, f.Message)
	}
	if res.Errors != 0 {
		t.Errorf("errors = %d, want 0", res.Errors)
	}
	if res.Disclaimer != Disclaimer {
		t.Errorf("disclaimer = %q", res.Disclaimer)
	}
}

func TestQuestionMarkNeverAnError(t *testing.T) {
	root := t.TempDir()
	writeRec(t, root, "research/sources/P-001.md", map[string]string{
		"id": "P-001", "type": "source", "state": "open",
	}, "# T\n\n## Bibliographic identity\n\ndoi: 10.1/x\nsha256: "+strings.Repeat("b", 64)+"\n\n## Claims\n\n- c (locator: p. 1)\n")
	writeRec(t, root, "research/issues/I-001.md", map[string]string{
		"id": "I-001", "type": "issue", "state": "open",
	}, cleanIssueBody()+"\n?  an unresolved obligation line\n?  another one\n")

	res, err := Run(root)
	if err != nil {
		t.Fatal(err)
	}
	for _, f := range res.Findings {
		if strings.Contains(f.Message, "?") {
			t.Errorf("a ? line produced a finding: %+v", f)
		}
	}
	if res.Errors != 0 {
		t.Errorf("errors = %d, want 0 (? lines are obligations, not errors)", res.Errors)
	}
}

func TestBrokenReference(t *testing.T) {
	root := t.TempDir()
	writeRec(t, root, "research/issues/I-001.md", map[string]string{
		"id": "I-001", "type": "issue", "state": "open",
		"refs": "refs:\n  - relation: motivated_by\n    id: P-999\n",
	}, cleanIssueBody())

	res, err := Run(root)
	if err != nil {
		t.Fatal(err)
	}
	if !findRule(res, "refs", "P-999 not found") {
		t.Errorf("missing broken-reference finding: %+v", res.Findings)
	}
}

func TestDependsOnCycleIsError(t *testing.T) {
	root := t.TempDir()
	writeRec(t, root, "research/issues/I-001.md", map[string]string{
		"id": "I-001", "type": "issue", "state": "open",
		"refs": "refs:\n  - relation: depends_on\n    id: I-002\n",
	}, cleanIssueBody())
	writeRec(t, root, "research/issues/I-002.md", map[string]string{
		"id": "I-002", "type": "issue", "state": "open",
		"refs": "refs:\n  - relation: depends_on\n    id: I-001\n",
	}, cleanIssueBody())

	res, err := Run(root)
	if err != nil {
		t.Fatal(err)
	}
	if !findRule(res, "cycle", "depends_on cycle") {
		t.Errorf("depends_on cycle not reported: %+v", res.Findings)
	}
}

func TestNonDependsOnCycleIsNotAnError(t *testing.T) {
	root := t.TempDir()
	writeRec(t, root, "research/issues/I-001.md", map[string]string{
		"id": "I-001", "type": "issue", "state": "open",
		"refs": "refs:\n  - relation: challenges\n    id: I-002\n",
	}, cleanIssueBody())
	writeRec(t, root, "research/issues/I-002.md", map[string]string{
		"id": "I-002", "type": "issue", "state": "open",
		"refs": "refs:\n  - relation: challenges\n    id: I-001\n",
	}, cleanIssueBody())

	res, err := Run(root)
	if err != nil {
		t.Fatal(err)
	}
	if res.Errors != 0 {
		t.Errorf("mutual challenge must not be an error, got: %+v", res.Findings)
	}
}

func TestInvalidEnum(t *testing.T) {
	root := t.TempDir()
	writeRec(t, root, "research/issues/I-001.md", map[string]string{
		"id": "I-001", "type": "issue", "state": "urgent", "priority": "critical", "branch": "Z",
	}, cleanIssueBody())

	res, err := Run(root)
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{"invalid state", "invalid priority", "invalid branch"} {
		if !findRule(res, "enum", want) {
			t.Errorf("missing enum finding %q: %+v", want, res.Findings)
		}
	}
}

func TestIDTypeMismatch(t *testing.T) {
	root := t.TempDir()
	// A source ID inside an issue record (wrong prefix for its type).
	writeRec(t, root, "research/issues/P-001.md", map[string]string{
		"id": "P-001", "type": "issue", "state": "open",
	}, cleanIssueBody())

	res, err := Run(root)
	if err != nil {
		t.Fatal(err)
	}
	if !findRule(res, "identity", "expected prefix I-") {
		t.Errorf("missing prefix finding: %+v", res.Findings)
	}
}

func TestRecordInWrongDirectory(t *testing.T) {
	root := t.TempDir()
	// Correct ID and type, but stored in the wrong directory.
	writeRec(t, root, "research/sources/I-001.md", map[string]string{
		"id": "I-001", "type": "issue", "state": "open",
	}, cleanIssueBody())

	res, err := Run(root)
	if err != nil {
		t.Fatal(err)
	}
	if !findRule(res, "location", "must live at research/issues/I-001.md") {
		t.Errorf("missing location finding: %+v", res.Findings)
	}
}

func TestFilenameIDMismatch(t *testing.T) {
	root := t.TempDir()
	// File name and header ID disagree.
	writeRec(t, root, "research/issues/I-002.md", map[string]string{
		"id": "I-001", "type": "issue", "state": "open",
	}, cleanIssueBody())

	res, err := Run(root)
	if err != nil {
		t.Fatal(err)
	}
	if !findRule(res, "location", "must live at research/issues/I-001.md") {
		t.Errorf("missing location finding: %+v", res.Findings)
	}
}

func TestDuplicateID(t *testing.T) {
	root := t.TempDir()
	writeRec(t, root, "research/issues/I-001.md", map[string]string{
		"id": "I-001", "type": "issue", "state": "open",
	}, cleanIssueBody())
	writeRec(t, root, "research/issues/I-001-copy.md", map[string]string{
		"id": "I-001", "type": "issue", "state": "open",
	}, cleanIssueBody())

	res, err := Run(root)
	if err != nil {
		t.Fatal(err)
	}
	if !findRule(res, "duplicate", "duplicate id I-001") {
		t.Errorf("missing duplicate finding: %+v", res.Findings)
	}
}

func TestStrayFile(t *testing.T) {
	root := t.TempDir()
	path := filepath.Join(root, "research", "notes.md")
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, []byte("stray\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	res, err := Run(root)
	if err != nil {
		t.Fatal(err)
	}
	if !findRule(res, "location", "outside its designated directory") {
		t.Errorf("missing stray finding: %+v", res.Findings)
	}
}

func TestMissingIssueSections(t *testing.T) {
	root := t.TempDir()
	writeRec(t, root, "research/issues/I-001.md", map[string]string{
		"id": "I-001", "type": "issue", "state": "open",
	}, "# T\n\n## Question\n\nx\n")

	res, err := Run(root)
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{"Theoretical expectation", "Independent theoretical result", "Experimental expectation", "Experimental actual"} {
		if !findRule(res, "section", want) {
			t.Errorf("missing section finding %q: %+v", want, res.Findings)
		}
	}
}

func TestPlaceholderInSectionIsError(t *testing.T) {
	root := t.TempDir()
	// Replace a filled value in a *required* section with a placeholder.
	body := strings.Replace(cleanIssueBody(),
		"- Claim or prediction: WEP holds",
		"- Claim or prediction: <what is claimed or predicted>", 1)
	if body == cleanIssueBody() {
		t.Fatal("test setup: replacement did not match")
	}
	writeRec(t, root, "research/issues/I-001.md", map[string]string{
		"id": "I-001", "type": "issue", "state": "open",
	}, body)

	res, err := Run(root)
	if err != nil {
		t.Fatal(err)
	}
	if !findRule(res, "section", "unfilled placeholder") {
		t.Errorf("placeholder not flagged: %+v", res.Findings)
	}
}

func TestPhysicalHypothesisRequiresFalsifier(t *testing.T) {
	root := t.TempDir()
	body := "# T\n\nclaim wep_prediction: under WEP, eta(Ti,Pt) = 0\n"
	writeRec(t, root, "research/hypotheses/H-001.md", map[string]string{
		"id": "H-001", "type": "hypothesis", "state": "open", "class": "physical",
	}, body)

	res, err := Run(root)
	if err != nil {
		t.Fatal(err)
	}
	if !findRule(res, "section", "Falsification condition") {
		t.Errorf("missing falsification condition not flagged: %+v", res.Findings)
	}
}

func TestMathematicalHypothesisSections(t *testing.T) {
	root := t.TempDir()
	writeRec(t, root, "research/hypotheses/H-001.md", map[string]string{
		"id": "H-001", "type": "hypothesis", "state": "open", "class": "mathematical",
	}, "# T\n\nclaim x: some proposition\n")

	res, err := Run(root)
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{"Premises", "Conclusion", "Domain", "Proof obligations"} {
		if !findRule(res, "section", want) {
			t.Errorf("missing %q finding: %+v", want, res.Findings)
		}
	}
}

func TestHypothesisRequiresClass(t *testing.T) {
	root := t.TempDir()
	writeRec(t, root, "research/hypotheses/H-001.md", map[string]string{
		"id": "H-001", "type": "hypothesis", "state": "open",
	}, "# T\n")

	res, err := Run(root)
	if err != nil {
		t.Fatal(err)
	}
	if !findRule(res, "field", "require class") {
		t.Errorf("missing class finding: %+v", res.Findings)
	}
}

func TestClassOnlyOnHypotheses(t *testing.T) {
	root := t.TempDir()
	writeRec(t, root, "research/issues/I-001.md", map[string]string{
		"id": "I-001", "type": "issue", "state": "open", "class": "physical",
	}, cleanIssueBody())

	res, err := Run(root)
	if err != nil {
		t.Fatal(err)
	}
	if !findRule(res, "enum", "only valid on hypothesis") {
		t.Errorf("missing class-scope finding: %+v", res.Findings)
	}
}

func TestSourceLocators(t *testing.T) {
	root := t.TempDir()
	writeRec(t, root, "research/sources/P-001.md", map[string]string{
		"id": "P-001", "type": "source", "state": "open",
	}, `# T

## Bibliographic identity

doi: 10.1/x
sha256: `+strings.Repeat("c", 64)+`

## Claims

- a claim with no locator at all
- another claim (locator: <page, section, equation, figure or table>)
`)

	res, err := Run(root)
	if err != nil {
		t.Fatal(err)
	}
	if !findRule(res, "source-locator", "no 'locator:'") {
		t.Errorf("missing locator finding: %+v", res.Findings)
	}
	if !findRule(res, "source-locator", "no claim carries a concrete source locator") {
		t.Errorf("placeholder locator not flagged: %+v", res.Findings)
	}
}

func TestSourceRequiresDOIORArxivAndSHA(t *testing.T) {
	root := t.TempDir()
	writeRec(t, root, "research/sources/P-001.md", map[string]string{
		"id": "P-001", "type": "source", "state": "open",
	}, `# T

## Bibliographic identity

title: T

## Claims

- c (locator: p. 1)
`)

	res, err := Run(root)
	if err != nil {
		t.Fatal(err)
	}
	if !findRule(res, "bibliographic", "doi: or arxiv:") {
		t.Errorf("missing doi/arxiv finding: %+v", res.Findings)
	}
	if !findRule(res, "bibliographic", "sha256:") {
		t.Errorf("missing sha256 finding: %+v", res.Findings)
	}
}

func TestShortSHARejected(t *testing.T) {
	root := t.TempDir()
	writeRec(t, root, "research/sources/P-001.md", map[string]string{
		"id": "P-001", "type": "source", "state": "open",
	}, `# T

## Bibliographic identity

doi: 10.1/x
sha256: abc123

## Claims

- c (locator: p. 1)
`)

	res, err := Run(root)
	if err != nil {
		t.Fatal(err)
	}
	if !findRule(res, "bibliographic", "64 hex characters") {
		t.Errorf("short sha not flagged: %+v", res.Findings)
	}
}

func TestUnknownCheckProfile(t *testing.T) {
	root := t.TempDir()
	writeRec(t, root, "derivations/D-001.md", map[string]string{
		"id": "D-001", "type": "derivation", "state": "open",
	}, `# T

## Premises

p

## Dependencies

none

## Intermediate results

n

## Mathematical rules

none

## Obligations

none

## Checks

- profile: not-a-real-profile
`)

	res, err := Run(root)
	if err != nil {
		t.Fatal(err)
	}
	if !findRule(res, "check-profile", "unknown check profile") {
		t.Errorf("unknown profile not flagged: %+v", res.Findings)
	}
}

func TestFrontmatterParseErrorIsFinding(t *testing.T) {
	root := t.TempDir()
	path := filepath.Join(root, "research", "sources", "P-001.md")
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, []byte("---\nid: P-001\ncolor: red\n---\n"), 0o644); err != nil {
		t.Fatal(err)
	}

	res, err := Run(root)
	if err != nil {
		t.Fatal(err)
	}
	if !findRule(res, "frontmatter", "unknown field") {
		t.Errorf("frontmatter error not surfaced: %+v", res.Findings)
	}
	if res.Errors == 0 {
		t.Error("parse error must count as an error")
	}
}

func TestReviewStaleDigest(t *testing.T) {
	root := t.TempDir()
	writeRec(t, root, "research/hypotheses/H-001.md", map[string]string{
		"id": "H-001", "type": "hypothesis", "state": "open", "class": "physical",
	}, cleanPhysicalHypothesis())
	writeRec(t, root, "research/reviews/R-001.md", map[string]string{
		"id": "R-001", "type": "review", "state": "open",
	}, `# T

## Review target

reviewed: H-001
digest: sha256:`+strings.Repeat("0", 64)+`
reviewer: adversary
model: opencode/deepseek-v4.1-flash-free

## Summary

x

## Blocking findings

none

## Nonblocking findings

none

## Evidential boundary

checked structurally only
`)

	res, err := Run(root)
	if err != nil {
		t.Fatal(err)
	}
	if !findRule(res, "stale", "STALE") {
		t.Errorf("stale review not flagged: %+v", res.Findings)
	}
}

func TestReviewFreshDigestNoStale(t *testing.T) {
	root := t.TempDir()
	writeRec(t, root, "research/hypotheses/H-001.md", map[string]string{
		"id": "H-001", "type": "hypothesis", "state": "open", "class": "physical",
	}, cleanPhysicalHypothesis())

	// Compute the digest the same way store does.
	sum := digestOf(t, filepath.Join(root, "research", "hypotheses", "H-001.md"))
	writeRec(t, root, "research/reviews/R-001.md", map[string]string{
		"id": "R-001", "type": "review", "state": "open",
	}, `# T

## Review target

reviewed: H-001
digest: sha256:`+sum+`
reviewer: adversary
model: opencode/deepseek-v4.1-flash-free

## Summary

x

## Blocking findings

none

## Nonblocking findings

none

## Evidential boundary

checked structurally only
`)

	res, err := Run(root)
	if err != nil {
		t.Fatal(err)
	}
	if findRule(res, "stale", "STALE") {
		t.Errorf("fresh review wrongly flagged stale: %+v", res.Findings)
	}
}

func TestReviewRequiresIdentity(t *testing.T) {
	root := t.TempDir()
	writeRec(t, root, "research/hypotheses/H-001.md", map[string]string{
		"id": "H-001", "type": "hypothesis", "state": "open", "class": "physical",
	}, cleanPhysicalHypothesis())
	sum := digestOf(t, filepath.Join(root, "research", "hypotheses", "H-001.md"))
	writeRec(t, root, "research/reviews/R-001.md", map[string]string{
		"id": "R-001", "type": "review", "state": "open",
	}, `# T

## Review target

reviewed: H-001
digest: sha256:`+sum+`
reviewer: <self-reported model id used for this review>
model: <self-reported model id used for this review>

## Blocking findings

none

## Nonblocking findings

none

## Evidential boundary

x
`)

	res, err := Run(root)
	if err != nil {
		t.Fatal(err)
	}
	if !findRule(res, "review", "reviewer:") {
		t.Errorf("placeholder reviewer not flagged: %+v", res.Findings)
	}
	if !findRule(res, "review", "model:") {
		t.Errorf("placeholder model not flagged: %+v", res.Findings)
	}
}

func TestSelfReferenceRejected(t *testing.T) {
	root := t.TempDir()
	writeRec(t, root, "research/issues/I-001.md", map[string]string{
		"id": "I-001", "type": "issue", "state": "open",
		"refs": "refs:\n  - relation: motivated_by\n    id: I-001\n",
	}, cleanIssueBody())

	res, err := Run(root)
	if err != nil {
		t.Fatal(err)
	}
	if !findRule(res, "refs", "self-reference") {
		t.Errorf("self-reference not flagged: %+v", res.Findings)
	}
}

func TestInvalidRelationVocabulary(t *testing.T) {
	root := t.TempDir()
	writeRec(t, root, "research/issues/I-001.md", map[string]string{
		"id": "I-001", "type": "issue", "state": "open",
		"refs": "refs:\n  - relation: likes\n    id: I-002\n",
	}, cleanIssueBody())

	res, err := Run(root)
	if err != nil {
		t.Fatal(err)
	}
	if !findRule(res, "refs", "closed vocabulary") {
		t.Errorf("invalid relation not flagged: %+v", res.Findings)
	}
}

func TestAppendOnlyReviewModification(t *testing.T) {
	if _, err := exec.LookPath("git"); err != nil {
		t.Skip("git not available")
	}
	root := t.TempDir()
	run := func(args ...string) {
		t.Helper()
		cmd := exec.Command("git", append([]string{"-C", root}, args...)...)
		cmd.Env = append(os.Environ(),
			"GIT_AUTHOR_NAME=t", "GIT_AUTHOR_EMAIL=t@t",
			"GIT_COMMITTER_NAME=t", "GIT_COMMITTER_EMAIL=t@t",
		)
		if out, err := cmd.CombinedOutput(); err != nil {
			t.Fatalf("git %v: %v\n%s", args, err, out)
		}
	}
	run("init", "-b", "main")

	writeRec(t, root, "research/hypotheses/H-001.md", map[string]string{
		"id": "H-001", "type": "hypothesis", "state": "open", "class": "physical",
	}, cleanPhysicalHypothesis())
	sum := digestOf(t, filepath.Join(root, "research", "hypotheses", "H-001.md"))
	reviewBody := `# T

## Review target

reviewed: H-001
digest: sha256:` + sum + `
reviewer: adversary
model: opencode/deepseek-v4.1-flash-free

## Summary

x

## Blocking findings

none

## Nonblocking findings

none

## Evidential boundary

x
`
	writeRec(t, root, "research/reviews/R-001.md", map[string]string{
		"id": "R-001", "type": "review", "state": "open",
	}, reviewBody)

	run("add", "-A")
	run("commit", "-m", "initial records")

	// Unmodified review: no append-only finding.
	res, err := Run(root)
	if err != nil {
		t.Fatal(err)
	}
	if findRule(res, "append-only", "append-only") {
		t.Errorf("unmodified review flagged: %+v", res.Findings)
	}

	// Modify the review after creation.
	path := filepath.Join(root, "research", "reviews", "R-001.md")
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, append(data, []byte("\nedited after creation\n")...), 0o644); err != nil {
		t.Fatal(err)
	}

	res, err = Run(root)
	if err != nil {
		t.Fatal(err)
	}
	if !findRule(res, "append-only", "append-only") {
		t.Errorf("modified review not flagged: %+v", res.Findings)
	}
}

func TestDisclaimerAlwaysPresent(t *testing.T) {
	root := t.TempDir()
	res, err := Run(root)
	if err != nil {
		t.Fatal(err)
	}
	if res.Disclaimer == "" {
		t.Error("empty disclaimer")
	}
	if !strings.Contains(strings.ToLower(res.Disclaimer), "structural") {
		t.Errorf("disclaimer must say it is structural only: %q", res.Disclaimer)
	}
}

func TestSyntaxErrorLineReported(t *testing.T) {
	// The parser's line numbers flow through to the finding.
	se := &frontmatter.SyntaxError{Line: 42, Msg: "boom"}
	if !strings.Contains(se.Error(), "42") {
		t.Errorf("line number missing from error: %v", se)
	}
}
