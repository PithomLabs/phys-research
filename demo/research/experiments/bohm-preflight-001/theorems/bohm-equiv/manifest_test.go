// Package equivariance sidecar tests: strict structural validation of
// source.json and source_mapping.json plus manifest limitation pins.
// These tests validate linkage structure only, never semantic equivalence.
package equivariance

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"testing"

	"github.com/PithomLabs/phys-artifact"
)

type mappingEntry struct {
	SourceEq          string   `json:"source_eq"`
	SourceQuantity    string   `json:"source_quantity"`
	Meaning           string   `json:"meaning"`
	DefiningRelation  string   `json:"defining_relation"`
	FixtureRef        string   `json:"fixture_ref"`
	Dependencies      []string `json:"dependencies"`
	AllowUnrepresentd bool     `json:"allow_unrepresented"`
}

type mappingDoc struct {
	SourceMapping []mappingEntry `json:"source_mapping"`
}

type sourceDoc struct {
	Source struct {
		ArxivID        string `json:"arxiv_id"`
		JournalRef     string `json:"journal_ref"`
		ScopeNote      string `json:"scope_note"`
		Representation string `json:"representation"`
	} `json:"source"`
}

type manifestLimits struct {
	Limitations []struct {
		ID        string `json:"id"`
		Statement string `json:"statement"`
	} `json:"limitations"`
}

func sidecarPath(t *testing.T, name string) string {
	t.Helper()
	_, thisFile, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("runtime.Caller failed")
	}
	return filepath.Join(filepath.Dir(thisFile), name)
}

func loadMapping(t *testing.T) mappingDoc {
	t.Helper()
	raw, err := os.ReadFile(sidecarPath(t, "source_mapping.json"))
	if err != nil {
		t.Fatal(err)
	}
	var doc mappingDoc
	dec := json.NewDecoder(bytes.NewReader(raw))
	dec.DisallowUnknownFields()
	if err := dec.Decode(&doc); err != nil {
		t.Fatalf("source_mapping strict decode: %v", err)
	}
	return doc
}

func resolveFixtureRef(f *Fixture, name string) bool {
	name = strings.TrimSpace(name)
	v := reflect.ValueOf(f).Elem()
	if fv := v.FieldByName(name); fv.IsValid() && fv.Type() == reflect.TypeOf(artifact.ArtifactRef{}) {
		return true
	}
	if _, ok := f.Conds[name]; ok {
		return true
	}
	if _, ok := f.Symbols[name]; ok {
		return true
	}
	return false
}

func inScope(eq string) bool {
	for _, p := range []string{"01:", "02:", "03:", "fixture:"} {
		if strings.HasPrefix(eq, p) && len(eq) > len(p) {
			return true
		}
	}
	return false
}

func validateMapping(f *Fixture, doc mappingDoc) error {
	for i, e := range doc.SourceMapping {
		if e.SourceEq == "" || e.SourceQuantity == "" || e.Meaning == "" || e.DefiningRelation == "" || e.FixtureRef == "" {
			return errf("entry %d missing required field", i)
		}
		if !inScope(e.SourceEq) {
			return errf("entry %d source_eq %q out of scope", i, e.SourceEq)
		}
		for _, r := range strings.Split(e.FixtureRef, ",") {
			if !resolveFixtureRef(f, r) {
				return errf("entry %d fixture_ref %q unresolvable", i, r)
			}
		}
		for _, d := range e.Dependencies {
			if !resolveFixtureRef(f, d) {
				return errf("entry %d dependency %q unresolvable", i, d)
			}
		}
	}
	return nil
}

type terr string

func (e terr) Error() string { return string(e) }

func errf(format string, args ...any) error { return terr(sprintf(format, args...)) }

func sprintf(format string, args ...any) string {
	// Minimal formatting without importing fmt into test helpers twice.
	out := format
	for _, a := range args {
		out = strings.Replace(out, "%d", itoa(a), 1)
		out = strings.Replace(out, "%q", quote(a), 1)
	}
	return out
}

func itoa(a any) string {
	switch v := a.(type) {
	case int:
		if v == 0 {
			return "0"
		}
		neg := v < 0
		if neg {
			v = -v
		}
		digits := ""
		for v > 0 {
			digits = string(rune('0'+v%10)) + digits
			v /= 10
		}
		if neg {
			return "-" + digits
		}
		return digits
	default:
		return "?"
	}
}

func quote(a any) string {
	if s, ok := a.(string); ok {
		return "\"" + s + "\""
	}
	return "?"
}

func TestSourceMappingStructuralConsistency(t *testing.T) {
	f := Build()
	if err := validateMapping(f, loadMapping(t)); err != nil {
		t.Fatal(err)
	}
}

func TestSourceMappingRejectsCorruptions(t *testing.T) {
	f := Build()
	doc := loadMapping(t)
	mut := doc
	mut.SourceMapping[0].FixtureRef = "NoSuchRef"
	if err := validateMapping(f, mut); err == nil {
		t.Fatal("accepted unresolvable fixture_ref")
	}
	mut = doc
	mut.SourceMapping[1].SourceEq = "99:bogus"
	if err := validateMapping(f, mut); err == nil {
		t.Fatal("accepted out-of-scope source_eq")
	}
	mut = doc
	mut.SourceMapping[2].Meaning = ""
	if err := validateMapping(f, mut); err == nil {
		t.Fatal("accepted entry with missing field")
	}
}

func TestSourceMappingRejectsUnknownFields(t *testing.T) {
	raw, err := os.ReadFile(sidecarPath(t, "source_mapping.json"))
	if err != nil {
		t.Fatal(err)
	}
	injected := bytes.Replace(raw, []byte(`"source_eq"`), []byte(`"bogus_field": 1, "source_eq"`), 1)
	if bytes.Equal(injected, raw) {
		t.Fatal("injection anchor missing")
	}
	var doc mappingDoc
	dec := json.NewDecoder(bytes.NewReader(injected))
	dec.DisallowUnknownFields()
	if err := dec.Decode(&doc); err == nil {
		t.Fatal("accepted unknown field in source_mapping.json")
	}
}

func TestSourceSidecar(t *testing.T) {
	raw, err := os.ReadFile(sidecarPath(t, "source.json"))
	if err != nil {
		t.Fatal(err)
	}
	var doc sourceDoc
	dec := json.NewDecoder(bytes.NewReader(raw))
	dec.DisallowUnknownFields()
	if err := dec.Decode(&doc); err != nil {
		t.Fatalf("source strict decode: %v", err)
	}
	if doc.Source.ArxivID == "" || doc.Source.ScopeNote == "" || doc.Source.Representation == "" {
		t.Fatal("source sidecar missing required identity field")
	}
}

func TestManifestLimitations(t *testing.T) {
	raw, err := os.ReadFile(sidecarPath(t, "manifest.json"))
	if err != nil {
		t.Fatal(err)
	}
	var m manifestLimits
	if err := json.Unmarshal(raw, &m); err != nil {
		t.Fatal(err)
	}
	have := map[string]bool{}
	for _, l := range m.Limitations {
		have[l.ID] = true
	}
	for _, id := range []string{"no_flow_execution", "no_pushforward", "no_uniqueness_pde", "no_global_existence", "no_node_behavior", "no_schrodinger_residual", "test_units_only", "fragment_1d"} {
		if !have[id] {
			t.Fatalf("manifest limitations missing %s", id)
		}
	}
}
