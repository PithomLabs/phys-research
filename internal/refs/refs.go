// Package refs traverses explicitly declared relationships between research
// records. It follows only frontmatter `refs` edges — never prose — and
// reports cycles it encounters without treating them as errors (only
// depends_on cycles are validation errors; see the validate package).
package refs

import (
	"fmt"
	"sort"

	"github.com/PithomLabs/phys-research/internal/record"
	"github.com/PithomLabs/phys-research/internal/store"
)

// Edge is one traversed relationship. For incoming edges From is the
// neighbouring record and To is the start record; Direction says which.
type Edge struct {
	From      string `json:"from"`
	To        string `json:"to"`
	Relation  string `json:"relation"`
	Direction string `json:"direction"` // "out" or "in"
	File      string `json:"file"`
}

// Output is the traversal result.
type Output struct {
	Start   string   `json:"start"`
	Depth   int      `json:"depth"`
	Nodes   []string `json:"nodes"`   // discovery order, deterministic
	Edges   []Edge   `json:"edges"`   // sorted
	Cycles  []string `json:"cycles"`  // cycle paths encountered, deduplicated
	Missing []string `json:"missing"` // referenced IDs with no record, sorted
}

// Traverse follows declared refs from startID out to maxDepth.
// includeIn adds incoming edges to the traversal.
func Traverse(root, startID string, maxDepth int, includeIn bool) (Output, error) {
	if !record.IDPattern.MatchString(startID) {
		return Output{}, fmt.Errorf("invalid record id %q (expected e.g. H-001)", startID)
	}
	recs, err := store.Load(root)
	if err != nil {
		return Output{}, err
	}

	type frontmatterRef struct {
		relation string
		target   string
	}
	type entry struct {
		id    string
		rel   string
		file  string
		edges []frontmatterRef
	}

	byID := map[string]*entry{}
	var ids []string
	for i := range recs {
		r := &recs[i]
		if r.ParseErr != nil {
			continue
		}
		id := r.Header.ID
		if id == "" {
			continue
		}
		if _, ok := byID[id]; !ok {
			e := &entry{id: id, rel: r.Rel, file: r.Rel}
			for _, ref := range r.Header.Refs {
				e.edges = append(e.edges, frontmatterRef{relation: ref.Relation, target: ref.ID})
			}
			byID[id] = e
			ids = append(ids, id)
		}
	}
	sort.Strings(ids)

	start, ok := byID[startID]
	if !ok {
		return Output{}, fmt.Errorf("record %s not found", startID)
	}
	_ = start

	out := Output{Start: startID, Depth: maxDepth}
	visited := map[string]bool{startID: true}
	missingSet := map[string]bool{}
	cycleSet := map[string]bool{}

	// DFS with an explicit path stack for cycle reporting.
	type frame struct {
		id    string
		depth int
	}
	var walk func(id string, depth int, path []string)
	walk = func(id string, depth int, path []string) {
		if depth >= maxDepth {
			return
		}
		e := byID[id]
		if e == nil {
			return
		}
		type nbr struct {
			rel, target, direction string
		}
		var nbrs []nbr
		for _, ef := range e.edges {
			nbrs = append(nbrs, nbr{ef.relation, ef.target, "out"})
		}
		if includeIn {
			for _, otherID := range ids {
				if otherID == id {
					continue
				}
				for _, ef := range byID[otherID].edges {
					if ef.target == id {
						nbrs = append(nbrs, nbr{ef.relation, otherID, "in"})
					}
				}
			}
		}
		sort.Slice(nbrs, func(i, j int) bool {
			if nbrs[i].direction != nbrs[j].direction {
				return nbrs[i].direction < nbrs[j].direction
			}
			if nbrs[i].rel != nbrs[j].rel {
				return nbrs[i].rel < nbrs[j].rel
			}
			return nbrs[i].target < nbrs[j].target
		})

		for _, n := range nbrs {
			if n.direction == "out" {
				out.Edges = append(out.Edges, Edge{From: id, To: n.target, Relation: n.rel, Direction: "out", File: e.file})
			} else {
				out.Edges = append(out.Edges, Edge{From: n.target, To: id, Relation: n.rel, Direction: "in", File: byID[n.target].file})
			}

			if _, exists := byID[n.target]; !exists {
				if n.direction == "out" && !missingSet[n.target] {
					missingSet[n.target] = true
				}
				continue
			}

			// Cycle if the target is already on the current path.
			onPath := false
			for _, p := range path {
				if p == n.target {
					onPath = true
					break
				}
			}
			if onPath {
				cycle := append(append([]string{}, path...), n.target)
				key := canonical(cycle[:len(cycle)-1])
				if !cycleSet[key] {
					cycleSet[key] = true
					out.Cycles = append(out.Cycles, fmt.Sprintf("%s -> %s", join(cycle[:len(cycle)-1]), n.target))
				}
				continue
			}
			if visited[n.target] {
				continue
			}
			visited[n.target] = true
			out.Nodes = append(out.Nodes, n.target)
			walk(n.target, depth+1, append(path, n.target))
		}
	}

	out.Nodes = append(out.Nodes, startID)
	walk(startID, 0, []string{startID})

	sort.Slice(out.Edges, func(i, j int) bool {
		a, b := out.Edges[i], out.Edges[j]
		if a.From != b.From {
			return a.From < b.From
		}
		if a.To != b.To {
			return a.To < b.To
		}
		if a.Relation != b.Relation {
			return a.Relation < b.Relation
		}
		return a.Direction < b.Direction
	})
	sort.Strings(out.Missing)
	sort.Strings(out.Cycles)
	for m := range missingSet {
		out.Missing = append(out.Missing, m)
	}
	sort.Strings(out.Missing)
	return out, nil
}

func canonical(cycle []string) string {
	if len(cycle) == 0 {
		return ""
	}
	min := 0
	for i, s := range cycle {
		if s < cycle[min] {
			min = i
		}
	}
	rot := append(append([]string{}, cycle[min:]...), cycle[:min]...)
	return join(rot)
}

func join(parts []string) string {
	s := ""
	for i, p := range parts {
		if i > 0 {
			s += " -> "
		}
		s += p
	}
	return s
}
