package refs

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// buildGraph writes a small record graph.
func buildGraph(t *testing.T) string {
	t.Helper()
	root := t.TempDir()
	write := func(rel, header, body string) {
		t.Helper()
		path := filepath.Join(root, filepath.FromSlash(rel))
		if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(path, []byte("---\n"+header+"---\n\n"+body), 0o644); err != nil {
			t.Fatal(err)
		}
	}

	write("research/sources/P-001.md",
		"id: P-001\ntype: source\nstate: open\n",
		"# P-001\n")
	write("research/sources/P-002.md",
		"id: P-002\ntype: source\nstate: open\n",
		"# P-002\n")
	write("research/issues/I-001.md",
		"id: I-001\ntype: issue\nstate: open\nrefs:\n  - relation: motivated_by\n    id: P-001\n  - relation: depends_on\n    id: P-002\n",
		"# I-001\n")
	write("research/hypotheses/H-001.md",
		"id: H-001\ntype: hypothesis\nstate: open\nclass: physical\nrefs:\n  - relation: addresses\n    id: I-001\n",
		"# H-001\n")
	write("research/evidence/E-001.md",
		"id: E-001\ntype: evidence\nstate: open\nrefs:\n  - relation: supports\n    id: H-001\n",
		"# E-001\n")
	// A record with a broken outgoing edge.
	write("research/issues/I-002.md",
		"id: I-002\ntype: issue\nstate: open\nrefs:\n  - relation: motivated_by\n    id: P-999\n",
		"# I-002\n")
	// A mutual-challenge cycle (legal).
	write("research/issues/I-003.md",
		"id: I-003\ntype: issue\nstate: open\nrefs:\n  - relation: challenges\n    id: I-004\n",
		"# I-003\n")
	write("research/issues/I-004.md",
		"id: I-004\ntype: issue\nstate: open\nrefs:\n  - relation: challenges\n    id: I-003\n",
		"# I-004\n")
	return root
}

func TestTraverseDepth1(t *testing.T) {
	root := buildGraph(t)
	out, err := Traverse(root, "I-001", 1, false)
	if err != nil {
		t.Fatal(err)
	}
	if out.Start != "I-001" || out.Depth != 1 {
		t.Errorf("start/depth = %s/%d", out.Start, out.Depth)
	}
	// Outgoing only: P-001 and P-002, not H-001 (incoming).
	if len(out.Nodes) != 3 {
		t.Errorf("nodes = %v, want I-001 + 2 outgoing", out.Nodes)
	}
	for _, n := range out.Nodes {
		if n == "H-001" {
			t.Error("depth-1 outgoing traversal must not include incoming H-001")
		}
	}
	if len(out.Edges) != 2 {
		t.Errorf("edges = %+v, want 2", out.Edges)
	}
	if len(out.Missing) != 0 {
		t.Errorf("I-001 has no broken edges, missing = %v", out.Missing)
	}
}

func TestTraverseIncoming(t *testing.T) {
	root := buildGraph(t)
	out, err := Traverse(root, "I-001", 1, true)
	if err != nil {
		t.Fatal(err)
	}
	found := false
	for _, e := range out.Edges {
		if e.Direction == "in" && e.To == "I-001" && e.From == "H-001" {
			found = true
		}
	}
	if !found {
		t.Errorf("incoming edge H-001 -> I-001 missing: %+v", out.Edges)
	}
}

func TestTraverseDepthChain(t *testing.T) {
	root := buildGraph(t)

	// depth 1 from E-001 reaches H-001 only.
	out1, err := Traverse(root, "E-001", 1, false)
	if err != nil {
		t.Fatal(err)
	}
	if len(out1.Nodes) != 2 {
		t.Errorf("depth1 nodes = %v", out1.Nodes)
	}

	// depth 3 reaches E-001 -> H-001 -> I-001 -> P-001, P-002.
	out3, err := Traverse(root, "E-001", 3, false)
	if err != nil {
		t.Fatal(err)
	}
	want := map[string]bool{"E-001": true, "H-001": true, "I-001": true, "P-001": true, "P-002": true}
	for _, n := range out3.Nodes {
		if !want[n] {
			t.Errorf("unexpected node %s", n)
		}
		delete(want, n)
	}
	if len(want) != 0 {
		t.Errorf("missing nodes: %v (got %v)", want, out3.Nodes)
	}
}

func TestTraverseMissing(t *testing.T) {
	root := buildGraph(t)
	out, err := Traverse(root, "I-002", 2, false)
	if err != nil {
		t.Fatal(err)
	}
	if len(out.Missing) != 1 || out.Missing[0] != "P-999" {
		t.Errorf("missing = %v, want [P-999]", out.Missing)
	}
}

func TestTraverseCyclesReportedNotFatal(t *testing.T) {
	root := buildGraph(t)
	out, err := Traverse(root, "I-003", 5, false)
	if err != nil {
		t.Fatal(err)
	}
	if len(out.Cycles) == 0 {
		t.Fatalf("mutual challenge cycle not reported: %+v", out)
	}
	joined := strings.Join(out.Cycles, "; ")
	if !strings.Contains(joined, "I-003") || !strings.Contains(joined, "I-004") {
		t.Errorf("cycles = %v, want I-003/I-004", out.Cycles)
	}
}

func TestTraverseDeterministic(t *testing.T) {
	root := buildGraph(t)
	a, err := Traverse(root, "I-001", 3, true)
	if err != nil {
		t.Fatal(err)
	}
	for i := 0; i < 5; i++ {
		b, err := Traverse(root, "I-001", 3, true)
		if err != nil {
			t.Fatal(err)
		}
		if strings.Join(a.Nodes, ",") != strings.Join(b.Nodes, ",") {
			t.Fatalf("node order differs between runs:\n%v\n%v", a.Nodes, b.Nodes)
		}
		if len(a.Edges) != len(b.Edges) {
			t.Fatalf("edge count differs: %d vs %d", len(a.Edges), len(b.Edges))
		}
		for i := range a.Edges {
			if a.Edges[i] != b.Edges[i] {
				t.Fatalf("edge %d differs: %+v vs %+v", i, a.Edges[i], b.Edges[i])
			}
		}
	}
}

func TestTraverseInvalidID(t *testing.T) {
	root := buildGraph(t)
	for _, id := range []string{"", "h-001", "P-1", "NOTANID", "../../etc/passwd"} {
		if _, err := Traverse(root, id, 1, false); err == nil {
			t.Errorf("id %q accepted, want error", id)
		}
	}
}

func TestTraverseUnknownID(t *testing.T) {
	root := buildGraph(t)
	if _, err := Traverse(root, "H-999", 1, false); err == nil {
		t.Error("expected error for unknown record")
	} else if !strings.Contains(err.Error(), "not found") {
		t.Errorf("error = %v", err)
	}
}
