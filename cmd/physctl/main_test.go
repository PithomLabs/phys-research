package main

import (
	"archive/tar"
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

// buildBin compiles physctl once for the whole test binary.
func buildBin(t *testing.T) string {
	t.Helper()
	bin := filepath.Join(t.TempDir(), "physctl")
	cmd := exec.Command("go", "build", "-o", bin, ".")
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("go build: %v\n%s", err, out)
	}
	return bin
}

// cli runs the built binary in dir with PHYS_RESEARCH_ROOT set to dir.
func cli(t *testing.T, bin, dir string, args ...string) (int, string) {
	t.Helper()
	cmd := exec.Command(bin, args...)
	cmd.Dir = dir
	cmd.Env = append(os.Environ(), "PHYS_RESEARCH_ROOT="+dir)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out
	err := cmd.Run()
	code := 0
	if err != nil {
		ee, ok := err.(*exec.ExitError)
		if !ok {
			t.Fatalf("run %v: %v", args, err)
		}
		code = ee.ExitCode()
	}
	return code, out.String()
}

func TestCLIRoundTrip(t *testing.T) {
	bin := buildBin(t)
	root := t.TempDir()

	// --- new: source ---
	code, out := cli(t, bin, root, "new", "--type", "source", "--title", "MICROSCOPE 2022")
	if code != 0 {
		t.Fatalf("new source: exit %d\n%s", code, out)
	}
	if !strings.Contains(out, "created research/sources/P-001.md") {
		t.Errorf("unexpected output: %s", out)
	}

	// --- new: issue, referencing the source ---
	code, out = cli(t, bin, root, "new", "--type", "issue", "--title", "WEP test",
		"--branch", "C", "--priority", "high", "--ref", "motivated_by=P-001", "--json")
	if code != 0 {
		t.Fatalf("new issue: exit %d\n%s", code, out)
	}
	var created struct {
		ID   string `json:"id"`
		Path string `json:"path"`
	}
	if err := json.Unmarshal([]byte(out), &created); err != nil {
		t.Fatalf("new --json output not JSON: %v\n%s", err, out)
	}
	if created.ID != "I-001" || created.Path != "research/issues/I-001.md" {
		t.Errorf("created = %+v", created)
	}

	// --- new: physical hypothesis ---
	code, out = cli(t, bin, root, "new", "--type", "hypothesis", "--title", "WEP",
		"--class", "physical", "--ref", "addresses=I-001")
	if code != 0 {
		t.Fatalf("new hypothesis: exit %d\n%s", code, out)
	}

	// --- hypothesis requires class (exit 2, usage) ---
	code, out = cli(t, bin, root, "new", "--type", "hypothesis", "--title", "no class")
	if code != 2 {
		t.Errorf("missing --class: exit %d, want 2\n%s", code, out)
	}

	// --- fill in the templates so validation can pass ---
	fillTemplates(t, root)

	// --- validate: clean (exit 0) ---
	code, out = cli(t, bin, root, "validate")
	if code != 0 {
		t.Errorf("validate: exit %d, want 0\n%s", code, out)
	}
	if !strings.Contains(out, "Structural validation only") {
		t.Errorf("validate output missing disclaimer:\n%s", out)
	}

	// --- refs (flags must precede the positional ID: Go's flag package) ---
	code, out = cli(t, bin, root, "refs", "--depth", "2", "H-001")
	if code != 0 {
		t.Errorf("refs: exit %d\n%s", code, out)
	}
	for _, want := range []string{"H-001", "I-001", "P-001"} {
		if !strings.Contains(out, want) {
			t.Errorf("refs output missing %s:\n%s", want, out)
		}
	}

	// --- status --json ---
	code, out = cli(t, bin, root, "status", "--json")
	if code != 0 {
		t.Errorf("status: exit %d\n%s", code, out)
	}
	var st struct {
		Issues []struct {
			ID string `json:"id"`
		} `json:"issues"`
	}
	if err := json.Unmarshal([]byte(out), &st); err != nil {
		t.Fatalf("status --json not JSON: %v\n%s", err, out)
	}
	if len(st.Issues) != 1 || st.Issues[0].ID != "I-001" {
		t.Errorf("status issues = %+v", st.Issues)
	}

	// --- review creation records the hypothesis digest ---
	code, out = cli(t, bin, root, "new", "--type", "review", "--title", "R", "--review", "H-001")
	if code != 0 {
		t.Fatalf("new review: exit %d\n%s", code, out)
	}
	reviewRaw, err := os.ReadFile(filepath.Join(root, "research", "reviews", "R-001.md"))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(reviewRaw), "reviewed: H-001") {
		t.Errorf("review does not record target:\n%s", reviewRaw)
	}
	if !strings.Contains(string(reviewRaw), "digest: sha256:") {
		t.Errorf("review does not record digest:\n%s", reviewRaw)
	}
}

func TestCLIExitCodes(t *testing.T) {
	bin := buildBin(t)
	root := t.TempDir()

	// No args: usage, exit 2.
	if code, out := cli(t, bin, root); code != 2 || !strings.Contains(out, "Usage:") {
		t.Errorf("no args: exit %d, want 2 with usage\n%s", code, out)
	}
	// Unknown command: exit 2.
	if code, out := cli(t, bin, root, "frobnicate"); code != 2 {
		t.Errorf("unknown command: exit %d, want 2\n%s", code, out)
	}
	// Unknown profile: exit 2 (capability problem).
	if code, out := cli(t, bin, root, "check", "not-a-profile"); code != 2 {
		t.Errorf("unknown profile: exit %d, want 2\n%s", code, out)
	}
	if code, _ := cli(t, bin, root, "check"); code != 0 {
		t.Errorf("check with no args lists the allowlist: exit %d, want 0", code)
	}
	// refs with no ID: exit 2.
	if code, _ := cli(t, bin, root, "refs"); code != 2 {
		t.Errorf("refs without ID: exit %d, want 2", code)
	}
	// refs with a bad ID: exit 2.
	if code, out := cli(t, bin, root, "refs", "not-an-id"); code != 2 {
		t.Errorf("refs bad ID: exit %d, want 2\n%s", code, out)
	}
}

func TestCLIBrokenValidationExit1(t *testing.T) {
	bin := buildBin(t)
	root := t.TempDir()

	// A record with a broken reference must fail validation with exit 1.
	path := filepath.Join(root, "research", "issues", "I-001.md")
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatal(err)
	}
	body := `---
id: I-001
type: issue
state: open
refs:
  - relation: motivated_by
    id: P-999
---

# T

## Theoretical expectation

x

## Independent theoretical result

NOT APPLICABLE — no independent theoretical derivation is required for this experimental result.

## Experimental expectation

x

## Experimental actual

x
`
	if err := os.WriteFile(path, []byte(body), 0o644); err != nil {
		t.Fatal(err)
	}

	code, out := cli(t, bin, root, "validate")
	if code != 1 {
		t.Errorf("broken record: exit %d, want 1\n%s", code, out)
	}
	if !strings.Contains(out, "P-999") {
		t.Errorf("output does not name the broken ref:\n%s", out)
	}
	// JSON form must carry the disclaimer too.
	code, out = cli(t, bin, root, "validate", "--json")
	if code != 1 {
		t.Errorf("validate --json: exit %d, want 1", code)
	}
	var res struct {
		Errors     int    `json:"errors"`
		Disclaimer string `json:"disclaimer"`
	}
	if err := json.Unmarshal([]byte(out), &res); err != nil {
		t.Fatalf("validate --json not JSON: %v\n%s", err, out)
	}
	if res.Errors == 0 || res.Disclaimer == "" {
		t.Errorf("validate --json = %+v", res)
	}
}

func TestCLIUnknownCommandHelp(t *testing.T) {
	bin := buildBin(t)
	code, out := cli(t, bin, t.TempDir(), "--help")
	if code != 0 {
		t.Errorf("--help exit %d, want 0", code)
	}
	for _, want := range []string{"physctl new", "physctl status", "physctl refs", "physctl validate", "physctl check"} {
		if !strings.Contains(out, want) {
			t.Errorf("help missing %q", want)
		}
	}
}

// fillTemplates replaces every placeholder in the generated records with
// concrete MVP-1 content so that structural validation can pass.
func fillTemplates(t *testing.T, root string) {
	t.Helper()

	replace := func(rel string, from, to string) {
		t.Helper()
		p := filepath.Join(root, filepath.FromSlash(rel))
		raw, err := os.ReadFile(p)
		if err != nil {
			t.Fatal(err)
		}
		if !strings.Contains(string(raw), from) {
			t.Fatalf("%s: expected to find %q", rel, from)
		}
		out := strings.ReplaceAll(string(raw), from, to)
		if err := os.WriteFile(p, []byte(out), 0o644); err != nil {
			t.Fatal(err)
		}
	}

	sha := strings.Repeat("0123456789abcdef", 4)

	replace("research/sources/P-001.md",
		"title: <full title>",
		"title: Test of the Weak Equivalence Principle with MICROSCOPE")
	replace("research/sources/P-001.md",
		"authors: <authors>",
		"authors: Touboul et al.")
	replace("research/sources/P-001.md",
		"doi: <DOI, or \"none\">",
		"doi: 10.1103/PhysRevLett.129.121102")
	replace("research/sources/P-001.md",
		"arxiv: <arXiv ID with version, or \"none\">",
		"arxiv: 2209.15487v1")
	replace("research/sources/P-001.md",
		"sha256: <SHA-256 of the exact file examined; the file itself is never committed>",
		"sha256: "+sha)
	replace("research/sources/P-001.md",
		"- <claim text> (locator: <page, section, equation, figure or table>)",
		"- eta(Ti,Pt) = [-1.5 +- 2.3(stat) +- 1.5(syst)] x 10^-15 (locator: Eq. (5))")

	replace("research/issues/I-001.md",
		"<the outstanding problem, stated precisely>",
		"Does the MICROSCOPE 2022 result confirm the WEP for Ti/Pt?")
	replace("research/issues/I-001.md",
		"- Claim or prediction: <what is claimed or predicted>",
		"- Claim or prediction: WEP implies eta(Ti,Pt) = 0")
	replace("research/issues/I-001.md",
		"- Assumptions: <conditions the claim rests on>",
		"- Assumptions: local Lorentz invariance, no composition-dependent fifth force")
	replace("research/issues/I-001.md",
		"- Applicable regime: <energy, scale, curvature or matter conditions>",
		"- Applicable regime: solar-system gravity, Ti/Pt composition pair")
	replace("research/issues/I-001.md",
		"- Source or derivation reference: <record ID plus locator, or \"none\">",
		"- Source or derivation reference: P-001, Eq. (5)")
	replace("research/issues/I-001.md",
		"- Predicted observable: <quantity that could be measured>",
		"- Predicted observable: eta(Ti,Pt)")
	replace("research/issues/I-001.md",
		"- Units and conditions: <units, experimental conditions>",
		"- Units and conditions: dimensionless, drag-free satellite")
	replace("research/issues/I-001.md",
		"- Uncertainty: <expected uncertainty where applicable>",
		"- Uncertainty: stat 2.3e-15, syst 1.5e-15 (recorded separately)")
	replace("research/issues/I-001.md",
		"- Falsification condition: <what would count against the claim>",
		"- Falsification condition: |eta| exceeding the declared total uncertainty")
	replace("research/issues/I-001.md",
		"- Measured result: <measured value, or \"not yet measured\">",
		"- Measured result: eta(Ti,Pt) = [-1.5 +- 2.3(stat) +- 1.5(syst)] x 10^-15")
	replace("research/issues/I-001.md",
		"- Units and uncertainty: <units; statistical and systematic uncertainties>",
		"- Units and uncertainty: dimensionless; stat and syst kept separate as published")
	replace("research/issues/I-001.md",
		"- Methodology and experimental conditions: <how it was measured>",
		"- Methodology and experimental conditions: MICROSCOPE torsion balance, 2016-2017 data")
	replace("research/issues/I-001.md",
		"- Exact source locator: <equation, figure, table, page and source ID>",
		"- Exact source locator: Eq. (5), Fig. 1 of P-001")

	replace("research/hypotheses/H-001.md",
		`<Phys-Pseudo content above this line: exactly the seven constructs of
phys-pseudo/PSEUDO.md §1. Attribution is prose or # comments (§8). Never
write a falsification construct into the pseudo content: the condition
belongs in the section below.>`,
		`given apparatus MICROSCOPE torsion balance, composition pair (Ti, Pt)
define eta(Ti,Pt) as the differential free-fall acceleration parameter
assume the WEP holds for the Ti/Pt composition pair
claim wep_prediction: under WEP, eta(Ti,Pt) = 0
derive measured_eta from the published fit of Eq. (5)
if measured_eta differs from 0 beyond the declared total uncertainty then the WEP prediction is falsified
?  is the 2022 systematic budget fully characterized independently of the fit?
?  does the drag-free control loop couple differently to Ti and Pt?`)
	replace("research/hypotheses/H-001.md",
		`<falsification condition: the observation, measurement or re-analysis that
would count against this hypothesis, with units and regime where
applicable. Declare any total-uncertainty prescription explicitly.>`,
		`The WEP prediction for the Ti/Pt pair is falsified if a re-measurement or
re-analysis of the MICROSCOPE data yields |eta(Ti,Pt)| greater than 3 sigma
of a declared total uncertainty, where the total uncertainty prescription
must be declared explicitly before evaluation. The 2022 published value keeps
stat (2.3e-15) and syst (1.5e-15) separate; combining them in quadrature
gives approximately 2.7e-15, but that combination is a declared choice, not
a published number.`)
}

func TestCLIDistChecks(t *testing.T) {
	bin := buildBin(t)
	work := t.TempDir()

	writeFile := func(rel, body string) {
		p := filepath.Join(work, filepath.FromSlash(rel))
		if err := os.MkdirAll(filepath.Dir(p), 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(p, []byte(body), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	writeTar := func(name string, entries map[string]string) string {
		p := filepath.Join(work, name)
		f, err := os.Create(p)
		if err != nil {
			t.Fatal(err)
		}
		defer f.Close()
		tw := tar.NewWriter(f)
		defer tw.Close()
		for ename, body := range entries {
			if err := tw.WriteHeader(&tar.Header{Name: ename, Mode: 0o644, Size: int64(len(body))}); err != nil {
				t.Fatal(err)
			}
			if _, err := tw.Write([]byte(body)); err != nil {
				t.Fatal(err)
			}
		}
		return p
	}

	clean := writeTar("clean.tar", map[string]string{"docs/n.md": "n"})
	evil := writeTar("evil.tar", map[string]string{"../escape.md": "x"})

	if code, out := cli(t, bin, work, "dist", "check-tar", clean); code != 0 {
		t.Errorf("check-tar clean: exit %d\n%s", code, out)
	}
	if code, out := cli(t, bin, work, "dist", "check-tar", evil); code != 1 {
		t.Errorf("check-tar evil: exit %d, want 1\n%s", code, out)
	}
	if code, _ := cli(t, bin, work, "dist", "check-tar", filepath.Join(work, "missing.tar")); code != 2 {
		t.Errorf("check-tar missing: exit %d, want 2", code)
	}

	// Staged-tree check: payload manifest over one file, base listing empty.
	writeFile("stage/p.md", "payload")
	sum := sha256.Sum256([]byte("payload"))
	writeFile("payload.sha256", hex.EncodeToString(sum[:])+"  p.md\n")
	writeFile("base.list", "")
	stage := filepath.Join(work, "stage")
	manifest := filepath.Join(work, "payload.sha256")
	base := filepath.Join(work, "base.list")
	if code, out := cli(t, bin, work, "dist", "check-staged", "--dir", stage, "--manifest", manifest, "--base", base); code != 0 {
		t.Errorf("check-staged clean: exit %d\n%s", code, out)
	}
	writeFile("stage/p.md", "tampered")
	if code, out := cli(t, bin, work, "dist", "check-staged", "--dir", stage, "--manifest", manifest, "--base", base); code != 1 {
		t.Errorf("check-staged tampered: exit %d, want 1\n%s", code, out)
	}
	if code, _ := cli(t, bin, work, "dist", "check-staged", "--dir", stage); code != 2 {
		t.Errorf("check-staged missing flags: exit %d, want 2", code)
	}
}
