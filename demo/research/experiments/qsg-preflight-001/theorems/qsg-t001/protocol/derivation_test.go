// Package protocol_test is the executable derivation/protocol test for the
// QSG-T001 fixture. It exercises Build + Replay + gates; it never merely
// checks static files. The JSON sidecars are metadata, never derivation
// premises: this file reads them, fixture.go does not.
//
// Boundary (binding): every test here proves fixture integrity,
// arithmetic replay, or regression stability only. No test verifies
// QSG-T001's physics: a changed coefficient trips a snapshot (data
// integrity), it is not validated as mathematically correct here;
// replay checks derivation-codec consistency, not the proof (SK-1).
// See manifest.json limitations and AGENTS.md.
package protocol_test

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"testing"

	"github.com/PithomLabs/phys-artifact"
	"github.com/PithomLabs/phys-math"
	qsgt001 "github.com/PithomLabs/qsg-t001"
)

func areaFile(t *testing.T, name string) []byte {
	t.Helper()
	_, thisFile, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("runtime.Caller failed")
	}
	raw, err := os.ReadFile(filepath.Join(filepath.Dir(thisFile), "..", name))
	if err != nil {
		t.Fatal(err)
	}
	return raw
}

func resolve(t *testing.T, store physmath.Store, ref artifact.ArtifactRef) []byte {
	t.Helper()
	_, payload, err := store.Lookup(ref)
	if err != nil {
		t.Fatal(err)
	}
	return payload
}

func TestDerivationReplays(t *testing.T) {
	fx, err := qsgt001.Build()
	if err != nil {
		t.Fatal(err)
	}
	if err := qsgt001.Verify(fx); err != nil {
		t.Fatal(err)
	}
}

// goldenSnapshots pins every golden payload byte-for-byte, captured from
// an accepted build (v2 goldens: THIRD/CANCEL/TRACE/FLAT recalculated
// per doc 08 N2/N4; COMBO-6 and shapes unchanged). A source modification
// that silently changes a previously accepted golden trips this test.
// Stability here is data integrity, not physical correctness.
var goldenSnapshots = map[string]string{
	"combo":  "02000000000106000000000101",
	"cancel": "020000000000000000000101",
	"third":  "02000000000102000000000101",
	"trace":  "02000000000101000000000104",
	"flat":   "02010000000103000000000120",
	"v0":     "030000000204000000020200000000010100000000010201000000026d3201441dea1ba54f62ebb2b35102e065053627fae7cd7c2e4ade11f00c677041254a000400000002050200000000010100000000010c0100000001520156837812e7bd1107312e28e62d06acce8cc5ea8819e01e1ee9b5f7a7a36b038400",
	"h1":     "03000000040400000003020000000001020000000001010100000001520156837812e7bd1107312e28e62d06acce8cc5ea8819e01e1ee9b5f7a7a36b0384000100000003526d6e01557d1b4d83fe022c627f723f27e7e6ee9ce0adbd8b4e9662edcf171d1908acca000400000003050200000000010100000000010201000000016701eb2e43735b675acc026248fead387086f479545981d7e722f35913a38381b031000100000002523201478261f4e0dcd007c4ebd30747b904dd4ddab25a807116cb492f350cf75d042e0004000000030200000000010200000000010101000000016701eb2e43735b675acc026248fead387086f479545981d7e722f35913a38381b031000100000004426f78520191bc7b50606db9139019d6502ffef018fa275ee92d3ffa63e03344cbd5a001aa000400000002050200000000010200000000010101000000034e4e5201c686375d14ca22a58ef9809336150666824fdef66fa783facbe04ddc707f0dec00",
	"h2":     "0300000006050100000007444e7552616d75010339f831c5635f6fdbc26c6ba26cc7cb282301824e0a72c6eeac8f16b9e0911700050100000007444d7552616e7501b2cfce28af08ff12302e0c67f3d274fafb2b15da2640a53df2254e2f3039d7cd000100000006426f78526d6e01739c3bd157288cd02830200ac1efb75899e95e7a98f933b16ee63d7e8ca3adfa0004000000030200000000010100000000010201000000016701eb2e43735b675acc026248fead387086f479545981d7e722f35913a38381b031000100000004426f78520191bc7b50606db9139019d6502ffef018fa275ee92d3ffa63e03344cbd5a001aa000400000003050200000000010100000000010201000000016701eb2e43735b675acc026248fead387086f479545981d7e722f35913a38381b031000100000005526162737101de1bfb0e23f8a6d1487b1424127cdcf4def95c0c0927d97bb0901d77a2c4a38f000400000002020000000001020000000001010100000002525201fa5f4039a9333ceca7ceab0b363756b3b82a9260aca66e66ce567f316feb325c00",
	"flrw":   "040000000202000000000112000000000101030000000304000000020200000000010200000000010101000000044864644801ceb1ced3d062ae29159b1fe35e20406def5f2dae8f89fe7d3a5c2bca6eb94c750005010000000348643201f0e5ba61beb43b17ee6453b88b0a383986ff21c7c41db2f389c5ef514661ebba000400000002020000000001060000000001010100000004486448320166cf348da89310f850c476904e171e432fb5961d52bbe2043d310dfe964e5c2200",
}

func TestGoldenPayloadsStable(t *testing.T) {
	a, err := qsgt001.Build()
	if err != nil {
		t.Fatal(err)
	}
	b, err := qsgt001.Build()
	if err != nil {
		t.Fatal(err)
	}
	gotA := map[string]artifact.ArtifactRef{
		"combo": a.ComboResult, "cancel": a.CancelResult,
		"third": a.ThirdResult, "trace": a.TraceResult,
		"flat": a.FlatResult, "v0": a.V0Shape,
		"h1": a.H1Shape, "h2": a.H2Shape, "flrw": a.FLRW001,
	}
	gotB := map[string]artifact.ArtifactRef{
		"combo": b.ComboResult, "cancel": b.CancelResult,
		"third": b.ThirdResult, "trace": b.TraceResult,
		"flat": b.FlatResult, "v0": b.V0Shape,
		"h1": b.H1Shape, "h2": b.H2Shape, "flrw": b.FLRW001,
	}
	for tag, snapHex := range goldenSnapshots {
		want, err := hex.DecodeString(snapHex)
		if err != nil {
			t.Fatalf("%s: bad snapshot hex: %v", tag, err)
		}
		if raw := resolve(t, a.Store, gotA[tag]); !bytes.Equal(raw, want) {
			t.Fatalf("golden %s payload differs from fixed snapshot (regression break)", tag)
		}
		if rawB := resolve(t, b.Store, gotB[tag]); !bytes.Equal(resolve(t, a.Store, gotA[tag]), rawB) {
			t.Fatalf("golden %s payload differs across rebuilds (nondeterminism)", tag)
		}
	}
}

func TestAntiGoldenSeparation(t *testing.T) {
	fx, err := qsgt001.Build()
	if err != nil {
		t.Fatal(err)
	}
	// Build already gates separation; re-assert from stored refs so the
	// suite, not just the constructor, owns the falsification regression.
	for tag, pair := range map[string][2]artifact.ArtifactRef{
		"v0": {fx.V0Shape, fx.V0AntiB},
		"h1": {fx.H1Shape, fx.H1AntiB},
		"h2": {fx.H2Shape, fx.H2AntiB},
	} {
		if bytes.Equal(resolve(t, fx.Store, pair[0]), resolve(t, fx.Store, pair[1])) {
			t.Fatalf("%s: anti-golden payload equals golden", tag)
		}
	}
}

func TestMutations(t *testing.T) {
	fx, err := qsgt001.Build()
	if err != nil {
		t.Fatal(err)
	}
	// Mutant 1: wrong rational input (3+4 instead of 2+4) must not match
	// the COMBO-6 expected payload.
	mutantIn, err := physmath.PutExpression(fx.Store, physmath.Add(
		physmath.MustRational(3, 1), physmath.MustRational(4, 1)))
	if err != nil {
		t.Fatal(err)
	}
	res, err := physmath.Apply(fx.Store, physmath.Invocation{
		ProducerNamespace: "physmath",
		RuleID:            physmath.RuleExprSimplify,
		InputRefs:         []artifact.ArtifactRef{mutantIn},
	})
	if err != nil {
		t.Fatal(err)
	}
	if bytes.Equal(resolve(t, fx.Store, res.OutputRef), resolve(t, fx.Store, fx.ComboExpected)) {
		t.Fatal("mutant 3+4 matched the 2+4 golden")
	}
	// Mutant 2 (INTEGRITY tripwire, not a replay failure): tampered
	// derivation bytes must fail Verify. What this proves: Verify's
	// stored-bytes-vs-re-encode comparison rejects a swapped ref. What
	// it does NOT prove: replay behavior on a decoded malicious
	// derivation — there is no codec to decode tampered bytes into a
	// Derivation struct, so a decoded-malicious replay is untested.
	_, raw, err := fx.Store.Lookup(fx.DCombo6)
	if err != nil {
		t.Fatal(err)
	}
	bad := bytes.Clone(raw)
	bad[len(bad)-1] ^= 0xff
	badRef, err := fx.Store.Put(physmath.FamilyDerivation, bad)
	if err != nil {
		t.Fatal(err)
	}
	_ = badRef
	fx2, err := qsgt001.Build()
	if err != nil {
		t.Fatal(err)
	}
	fx2.DCombo6 = badRef
	if err := qsgt001.Verify(fx2); err == nil {
		t.Fatal("Verify accepted a tampered derivation ref")
	}
}

func TestSketchUnsupported(t *testing.T) {
	fx, err := qsgt001.Build()
	if err != nil {
		t.Fatal(err)
	}
	if len(fx.Sketch.Steps) == 0 {
		t.Fatal("empty sketch")
	}
	for _, s := range fx.Sketch.Steps {
		if s.Status != qsgt001.StatusSketchUnsupported {
			t.Fatalf("sketch %s leaves unsupported status", s.ID)
		}
		if s.Evidence == "" {
			t.Fatalf("sketch %s lacks evidence pointer", s.ID)
		}
	}
	if _, ok := reflect.TypeOf(qsgt001.QSGSketch{}).MethodByName("Verify"); ok {
		t.Fatal("QSGSketch must not grow a Verify method")
	}
	if _, ok := reflect.TypeOf(qsgt001.QSGSketch{}).MethodByName("Replay"); ok {
		t.Fatal("QSGSketch must not grow a Replay method")
	}
}

type manifestFile struct {
	SchemaVersion  string   `json:"schema_version"`
	ArtifactType   string   `json:"artifact_type"`
	DerivationID   string   `json:"derivation_id"`
	FrameworkID    string   `json:"framework_id"`
	Name           string   `json:"name"`
	EntryPoints    []string `json:"entry_points"`
	Regime         string   `json:"regime"`
	Coordinates    []string `json:"coordinates"`
	Conventions    []string `json:"conventions"`
	Premises       []string `json:"premises"`
	Assumptions    []string `json:"assumptions"`
	ExpectedResult []struct {
		Statement  string `json:"statement"`
		ExecutedBy string `json:"executed_by"`
	} `json:"expected_result"`
	Verification       []string `json:"verification"`
	Limitations        []string `json:"limitations"`
	Anomalies          []string `json:"anomalies"`
	FalsificationConds []string `json:"falsification_conditions"`
}

func TestManifestStrict(t *testing.T) {
	raw := areaFile(t, "manifest.json")
	var mf manifestFile
	dec := json.NewDecoder(bytes.NewReader(raw))
	dec.DisallowUnknownFields()
	if err := dec.Decode(&mf); err != nil {
		t.Fatalf("manifest strict decode: %v", err)
	}
	// Unknown-field injection: strict decoding must REJECT a manifest
	// carrying a field the struct does not declare.
	var m map[string]any
	if err := json.Unmarshal(raw, &m); err != nil {
		t.Fatal(err)
	}
	m["unknown_probe_field"] = 1
	bad, err := json.Marshal(m)
	if err != nil {
		t.Fatal(err)
	}
	var mf2 manifestFile
	dec2 := json.NewDecoder(bytes.NewReader(bad))
	dec2.DisallowUnknownFields()
	if err := dec2.Decode(&mf2); err == nil {
		t.Fatal("manifest strict decode accepted an unknown field")
	}
	if len(mf.ExpectedResult) == 0 || len(mf.Limitations) < 3 || len(mf.Assumptions) == 0 {
		t.Fatal("manifest missing expected_result/limitations/assumptions")
	}
	joinedExp := ""
	for _, e := range mf.ExpectedResult {
		if e.ExecutedBy == "" {
			t.Fatalf("expected_result entry without executed_by: %q", e.Statement)
		}
		joinedExp += e.Statement + "\n"
	}
	if strings.Contains(joinedExp, "universal") {
		t.Fatal("expected_result must not claim the universal identity")
	}
	joinedLim := strings.Join(mf.Limitations, "\n")
	if !strings.Contains(joinedLim, "universal") {
		t.Fatal("limitations must state the universal identity is unexecuted")
	}
	joinedAnom := strings.Join(mf.Anomalies, "\n")
	if !strings.Contains(joinedAnom, "E1") || !strings.Contains(joinedAnom, "E2") {
		t.Fatal("anomalies must record errata E1 and E2")
	}
	if !strings.Contains(joinedAnom, "E3") {
		t.Fatal("anomalies must record repair errata E3/E4/E5")
	}
}

type sidecarFile struct {
	ProvenanceNote string `json:"provenance_note"`
	Claims         []struct {
		ID            string   `json:"id"`
		Title         string   `json:"title"`
		Status        string   `json:"status"`
		MathStatus    string   `json:"math_status"`
		MachineStatus string   `json:"machine_status"`
		EvidenceRefs  []string `json:"evidence_refs"`
	} `json:"claims"`
	Obligations []struct {
		ID               string   `json:"id"`
		Statement        string   `json:"statement"`
		Blocks           []string `json:"blocks"`
		EvidenceRequired string   `json:"evidence_required"`
	} `json:"obligations"`
}

func TestSidecar(t *testing.T) {
	raw := areaFile(t, "research_metadata.json")
	var sc sidecarFile
	dec := json.NewDecoder(bytes.NewReader(raw))
	dec.DisallowUnknownFields()
	if err := dec.Decode(&sc); err != nil {
		t.Fatalf("sidecar strict decode: %v", err)
	}
	// Unknown-field injection: strict decoding must REJECT a sidecar
	// carrying a field the struct does not declare.
	var m map[string]any
	if err := json.Unmarshal(raw, &m); err != nil {
		t.Fatal(err)
	}
	m["unknown_probe_field"] = 1
	bad, err := json.Marshal(m)
	if err != nil {
		t.Fatal(err)
	}
	var sc2 sidecarFile
	dec2 := json.NewDecoder(bytes.NewReader(bad))
	dec2.DisallowUnknownFields()
	if err := dec2.Decode(&sc2); err == nil {
		t.Fatal("sidecar strict decode accepted an unknown field")
	}
	seenTheorem := false
	for _, c := range sc.Claims {
		if strings.Contains(c.Status, "machine-proved") || strings.Contains(c.Status, "machine-verified") {
			t.Fatalf("claim %s asserts machine proof: %q", c.ID, c.Status)
		}
		if c.ID == "QSG-T001" {
			seenTheorem = true
			if !strings.Contains(c.Status, "machine-unexecuted") {
				t.Fatalf("QSG-T001 must be machine-unexecuted, got %q", c.Status)
			}
		}
	}
	if !seenTheorem {
		t.Fatal("QSG-T001 claim missing")
	}
	wantO := map[string]bool{"O1": false, "O2": false, "O3": false, "O4": false, "O5": false}
	for _, o := range sc.Obligations {
		if _, ok := wantO[o.ID]; !ok {
			t.Fatalf("unexpected obligation %s", o.ID)
		}
		wantO[o.ID] = true
		if o.EvidenceRequired == "" {
			t.Fatalf("obligation %s lacks evidence_required", o.ID)
		}
	}
	for id, seen := range wantO {
		if !seen {
			t.Fatalf("obligation %s missing", id)
		}
	}
}
