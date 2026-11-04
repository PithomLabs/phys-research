// Package status reports the current state of the research store: open and
// blocked issues, unresolved `?` obligations, stale reviews and blocking
// findings. It reports; it does not decide anything.
package status

import (
	"fmt"
	"sort"
	"strings"

	"github.com/PithomLabs/phys-research/internal/record"
	"github.com/PithomLabs/phys-research/internal/store"
)

// Issue is one open or blocked issue.
type Issue struct {
	ID       string   `json:"id"`
	State    string   `json:"state"`
	Priority string   `json:"priority,omitempty"`
	Branch   string   `json:"branch,omitempty"`
	File     string   `json:"file"`
	Refs     []string `json:"refs,omitempty"` // "relation id"
}

// Obligation is one unresolved `?` line.
type Obligation struct {
	ID   string   `json:"id"`
	File string   `json:"file"`
	Line int      `json:"line"`
	Text string   `json:"text"`
	Refs []string `json:"refs,omitempty"`
}

// StaleReview is a review whose recorded digest no longer matches its
// target (or whose target is gone).
type StaleReview struct {
	ID       string `json:"id"`
	Reviewed string `json:"reviewed"`
	File     string `json:"file"`
	Reason   string `json:"reason"`
}

// Blocking is one non-"none" line under a review's Blocking findings.
type Blocking struct {
	Review string `json:"review"`
	File   string `json:"file"`
	Line   int    `json:"line"`
	Text   string `json:"text"`
}

// Output is the machine-readable status snapshot.
type Output struct {
	Issues           []Issue       `json:"issues"`
	Obligations      []Obligation  `json:"obligations"`
	StaleReviews     []StaleReview `json:"stale_reviews"`
	BlockingFindings []Blocking    `json:"blocking_findings"`
}

// Collect builds the status snapshot. branch != "" filters issues by branch.
func Collect(root, branch string) (Output, error) {
	recs, err := store.Load(root)
	if err != nil {
		return Output{}, err
	}

	byID := map[string]*store.Record{}
	for i := range recs {
		r := &recs[i]
		if r.ParseErr == nil && r.Header.ID != "" {
			if _, ok := byID[r.Header.ID]; !ok {
				byID[r.Header.ID] = r
			}
		}
	}

	out := Output{
		// Non-nil so --json emits [] rather than null for empty sections.
		Issues:           []Issue{},
		Obligations:      []Obligation{},
		StaleReviews:     []StaleReview{},
		BlockingFindings: []Blocking{},
	}

	for i := range recs {
		r := &recs[i]
		if r.ParseErr != nil {
			continue
		}
		h := r.Header

		if h.Type == record.TypeIssue && (h.State == "open" || h.State == "blocked") {
			if branch == "" || h.Branch == branch {
				issue := Issue{ID: h.ID, State: h.State, Priority: h.Priority, Branch: h.Branch, File: r.Rel}
				for _, ref := range h.Refs {
					issue.Refs = append(issue.Refs, ref.Relation+" "+ref.ID)
				}
				out.Issues = append(out.Issues, issue)
			}
		}

		// Unresolved ? obligations (line-level, no Phys-Pseudo parsing).
		for j, ln := range strings.Split(string(r.Body), "\n") {
			trimmed := strings.TrimSpace(ln)
			if strings.HasPrefix(trimmed, "?") {
				text := strings.TrimSpace(strings.TrimPrefix(trimmed, "?"))
				ob := Obligation{ID: h.ID, File: r.Rel, Line: r.BodyLine + j, Text: text}
				for _, ref := range h.Refs {
					ob.Refs = append(ob.Refs, ref.Relation+" "+ref.ID)
				}
				out.Obligations = append(out.Obligations, ob)
			}
		}

		if h.Type == record.TypeReview {
			prov := store.ParseReview(r)
			if prov.Reviewed == "" {
				continue
			}
			// One shared freshness rule with the validate command.
			issues, _ := prov.ReviewIssues(func(id string) (*store.Record, bool) {
				rec, ok := byID[id]
				return rec, ok
			})
			if len(issues) > 0 {
				for _, is := range issues {
					out.StaleReviews = append(out.StaleReviews, StaleReview{
						ID: h.ID, Reviewed: prov.Reviewed, File: r.Rel, Reason: is.Reason(),
					})
				}
				continue
			}
			// Non-"none" blocking finding lines.
			lines := strings.Split(string(r.Body), "\n")
			inBlocking := false
			for j, ln := range lines {
				if strings.HasPrefix(ln, "## ") {
					inBlocking = ln == "## Blocking findings"
					continue
				}
				if !inBlocking {
					continue
				}
				t := strings.TrimSpace(ln)
				if t == "" || strings.HasPrefix(t, "#") || strings.EqualFold(t, "none") {
					continue
				}
				out.BlockingFindings = append(out.BlockingFindings, Blocking{
					Review: h.ID, File: r.Rel, Line: r.BodyLine + j, Text: t,
				})
			}
		}
	}

	priorityRank := map[string]int{"high": 0, "medium": 1, "low": 2}
	sort.Slice(out.Issues, func(i, j int) bool {
		a, b := out.Issues[i], out.Issues[j]
		ra, oka := priorityRank[a.Priority]
		if !oka {
			ra = 3
		}
		rb, okb := priorityRank[b.Priority]
		if !okb {
			rb = 3
		}
		if ra != rb {
			return ra < rb
		}
		return a.ID < b.ID
	})
	sort.Slice(out.Obligations, func(i, j int) bool {
		if out.Obligations[i].File != out.Obligations[j].File {
			return out.Obligations[i].File < out.Obligations[j].File
		}
		return out.Obligations[i].Line < out.Obligations[j].Line
	})
	sort.Slice(out.StaleReviews, func(i, j int) bool { return out.StaleReviews[i].ID < out.StaleReviews[j].ID })
	sort.Slice(out.BlockingFindings, func(i, j int) bool {
		if out.BlockingFindings[i].Review != out.BlockingFindings[j].Review {
			return out.BlockingFindings[i].Review < out.BlockingFindings[j].Review
		}
		return out.BlockingFindings[i].Line < out.BlockingFindings[j].Line
	})
	return out, nil
}

// Text renders the human-readable report.
func Text(out Output) string {
	var b strings.Builder

	b.WriteString("Open and blocked issues:\n")
	if len(out.Issues) == 0 {
		b.WriteString("  (none)\n")
	}
	for _, is := range out.Issues {
		prio := is.Priority
		if prio == "" {
			prio = "unspecified"
		}
		br := ""
		if is.Branch != "" {
			br = ", branch " + is.Branch
		}
		fmt.Fprintf(&b, "  %s [%s, priority %s%s] %s\n", is.ID, is.State, prio, br, is.File)
		if len(is.Refs) > 0 {
			fmt.Fprintf(&b, "      refs: %s\n", strings.Join(is.Refs, ", "))
		}
	}

	b.WriteString("\nUnresolved ? obligations:\n")
	if len(out.Obligations) == 0 {
		b.WriteString("  (none)\n")
	}
	for _, ob := range out.Obligations {
		fmt.Fprintf(&b, "  %s %s:%d: ? %s\n", ob.ID, ob.File, ob.Line, ob.Text)
	}

	b.WriteString("\nStale reviews:\n")
	if len(out.StaleReviews) == 0 {
		b.WriteString("  (none)\n")
	}
	for _, sr := range out.StaleReviews {
		fmt.Fprintf(&b, "  %s (%s): %s — %s\n", sr.ID, sr.Reviewed, sr.File, sr.Reason)
	}

	b.WriteString("\nBlocking findings:\n")
	if len(out.BlockingFindings) == 0 {
		b.WriteString("  (none)\n")
	}
	for _, bl := range out.BlockingFindings {
		fmt.Fprintf(&b, "  %s %s:%d: %s\n", bl.Review, bl.File, bl.Line, bl.Text)
	}
	return b.String()
}
