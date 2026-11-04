// Package validate performs STRUCTURAL validation of research records
// against the conventions in AGENTS.md §6.
//
// It never parses Phys-Pseudo, never evaluates physics or mathematics, and
// its result always carries an explicit disclaimer: structural validation
// is not scientific verification. Unresolved `?` lines are obligations
// reported by status, never errors here.
package validate

import (
	"bytes"
	"fmt"
	"os/exec"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"github.com/PithomLabs/phys-research/internal/check"
	"github.com/PithomLabs/phys-research/internal/frontmatter"
	"github.com/PithomLabs/phys-research/internal/record"
	"github.com/PithomLabs/phys-research/internal/store"
)

// Disclaimer is attached to every result, text and JSON.
const Disclaimer = "Structural validation only: no mathematical or physical correctness is implied."

// Severities.
const (
	SeverityError   = "error"
	SeverityWarning = "warning"
)

// Finding is one validation finding.
type Finding struct {
	Rule     string `json:"rule"`
	Severity string `json:"severity"`
	File     string `json:"file"`
	Line     int    `json:"line,omitempty"`
	Message  string `json:"message"`
}

// Result is the complete validation output.
type Result struct {
	Findings   []Finding `json:"findings"`
	Errors     int       `json:"errors"`
	Warnings   int       `json:"warnings"`
	Disclaimer string    `json:"disclaimer"`
}

// Run validates every record under root.
func Run(root string) (Result, error) {
	recs, err := store.Load(root)
	if err != nil {
		return Result{}, err
	}
	strays, err := store.Strays(root)
	if err != nil {
		return Result{}, err
	}

	var findings []Finding
	add := func(rule, severity, file string, line int, msg string) {
		findings = append(findings, Finding{Rule: rule, Severity: severity, File: file, Line: line, Message: msg})
	}

	byID := map[string][]*store.Record{}
	for i := range recs {
		r := &recs[i]
		if r.ParseErr != nil {
			line := 0
			if se, ok := r.ParseErr.(*frontmatter.SyntaxError); ok {
				line = se.Line
			}
			add("frontmatter", SeverityError, r.Rel, line, r.ParseErr.Error())
			continue
		}
		byID[r.Header.ID] = append(byID[r.Header.ID], r)
	}

	for _, s := range strays {
		add("location", SeverityError, s, 0, "record file outside its designated directory (AGENTS.md §6.3)")
	}

	// Duplicate IDs.
	var dupIDs []string
	for id := range byID {
		if id != "" && len(byID[id]) > 1 {
			dupIDs = append(dupIDs, id)
		}
	}
	sort.Strings(dupIDs)
	for _, id := range dupIDs {
		files := []string{}
		for _, r := range byID[id] {
			files = append(files, r.Rel)
		}
		sort.Strings(files)
		add("duplicate", SeverityError, files[1], 0,
			fmt.Sprintf("duplicate id %s also declared in %s", id, files[0]))
	}

	for i := range recs {
		r := &recs[i]
		if r.ParseErr != nil {
			continue
		}
		validateRecord(r, add)
	}

	// Cross-record reference checks and depends_on cycles.
	for i := range recs {
		r := &recs[i]
		if r.ParseErr != nil {
			continue
		}
		for _, ref := range r.Header.Refs {
			if ref.Relation == "" || !contains(record.Relations, ref.Relation) {
				add("refs", SeverityError, r.Rel, 0,
					fmt.Sprintf("invalid relation %q (closed vocabulary: %s)", ref.Relation, strings.Join(record.Relations, ", ")))
				continue
			}
			if !record.IDPattern.MatchString(ref.ID) {
				add("refs", SeverityError, r.Rel, 0,
					fmt.Sprintf("invalid reference id %q (expected e.g. P-001)", ref.ID))
				continue
			}
			if ref.ID == r.Header.ID {
				add("refs", SeverityError, r.Rel, 0, "self-reference is not allowed")
				continue
			}
			if _, ok := byID[ref.ID]; !ok {
				add("refs", SeverityError, r.Rel, 0,
					fmt.Sprintf("reference target %s not found (%s %s)", ref.ID, ref.Relation, ref.ID))
			}
		}
	}
	validateDependsOnCycles(recs, add)
	validateReviews(recs, byID, root, add)

	res := Result{Findings: findings, Disclaimer: Disclaimer}
	for _, f := range findings {
		if f.Severity == SeverityError {
			res.Errors++
		} else {
			res.Warnings++
		}
	}
	sort.Slice(res.Findings, func(i, j int) bool {
		a, b := res.Findings[i], res.Findings[j]
		if a.File != b.File {
			return a.File < b.File
		}
		if a.Line != b.Line {
			return a.Line < b.Line
		}
		if a.Rule != b.Rule {
			return a.Rule < b.Rule
		}
		return a.Message < b.Message
	})
	return res, nil
}

// validateRecord checks one record's header values, identity, location and
// required body sections.
func validateRecord(r *store.Record, add func(rule, severity, file string, line int, msg string)) {
	h := r.Header

	// Required fields.
	for _, f := range []struct{ name, val string }{
		{"id", h.ID}, {"type", h.Type}, {"state", h.State},
	} {
		if f.val == "" {
			add("field", SeverityError, r.Rel, 0, fmt.Sprintf("missing required field %q", f.name))
		}
	}

	// Enums.
	typOK := contains(record.Types(), h.Type)
	if h.Type != "" && !typOK {
		add("enum", SeverityError, r.Rel, 0,
			fmt.Sprintf("invalid type %q (supported: %s)", h.Type, strings.Join(record.Types(), ", ")))
	}
	if h.State != "" && !contains(record.States, h.State) {
		add("enum", SeverityError, r.Rel, 0,
			fmt.Sprintf("invalid state %q (supported: %s)", h.State, strings.Join(record.States, ", ")))
	}
	if h.Priority != "" && !contains(record.Priorities, h.Priority) {
		add("enum", SeverityError, r.Rel, 0,
			fmt.Sprintf("invalid priority %q (supported: %s)", h.Priority, strings.Join(record.Priorities, ", ")))
	}
	if h.Branch != "" && !contains(record.Branches, h.Branch) {
		add("enum", SeverityError, r.Rel, 0,
			fmt.Sprintf("invalid branch %q (supported: %s)", h.Branch, strings.Join(record.Branches, ", ")))
	}
	if h.Class != "" && h.Type == record.TypeHypothesis && !contains(record.Classes, h.Class) {
		add("enum", SeverityError, r.Rel, 0,
			fmt.Sprintf("invalid class %q (supported: %s)", h.Class, strings.Join(record.Classes, ", ")))
	}
	if h.Class != "" && h.Type != record.TypeHypothesis {
		add("enum", SeverityError, r.Rel, 0, "class is only valid on hypothesis records")
	}
	if h.Type == record.TypeHypothesis && h.Class == "" {
		add("field", SeverityError, r.Rel, 0, "hypothesis records require class: physical|mathematical")
	}

	// Identity: ID pattern, prefix/type agreement, filename and directory.
	idOK := record.IDPattern.MatchString(h.ID)
	if h.ID != "" && !idOK {
		add("identity", SeverityError, r.Rel, 0,
			fmt.Sprintf("id %q does not match the <PREFIX>-NNN form", h.ID))
	}
	if idOK && typOK {
		prefix := record.PrefixByType[h.Type]
		if string(h.ID[0]) != prefix {
			add("identity", SeverityError, r.Rel, 0,
				fmt.Sprintf("id %s does not match type %s (expected prefix %s-)", h.ID, h.Type, prefix))
		}
	}
	if idOK && typOK {
		expected := record.DirByType[h.Type] + "/" + h.ID + ".md"
		if r.Rel != expected {
			add("location", SeverityError, r.Rel, 0,
				fmt.Sprintf("record must live at %s", expected))
		}
	}

	// Required body sections per type (AGENTS.md §6.4).
	switch h.Type {
	case record.TypeIssue:
		for _, s := range []string{
			"Theoretical expectation",
			"Independent theoretical result",
			"Experimental expectation",
			"Experimental actual",
		} {
			checkSection(r, s, add)
		}
	case record.TypeHypothesis:
		switch h.Class {
		case "physical":
			checkSection(r, "Falsification condition", add)
		case "mathematical":
			for _, s := range []string{"Premises", "Conclusion", "Domain", "Proof obligations"} {
				checkSection(r, s, add)
			}
		}
	case record.TypeEvidence:
		for _, s := range []string{
			"Measured value", "Units and uncertainty", "Methodology", "Source locator",
		} {
			checkSection(r, s, add)
		}
	case record.TypeDerivation:
		for _, s := range []string{
			"Premises", "Dependencies", "Intermediate results", "Mathematical rules", "Obligations",
		} {
			checkSection(r, s, add)
		}
		validateCheckProfiles(r, add)
	case record.TypeSource:
		validateSource(r, add)
	case record.TypeReview:
		for _, s := range []string{"Review target", "Blocking findings", "Nonblocking findings", "Evidential boundary"} {
			checkSection(r, s, add)
		}
	}
}

var placeholderRe = regexp.MustCompile(`<[A-Za-z][^<>]*>`)

// checkSection requires a level-2 heading with substantive, placeholder-free
// content. Guidance comments (#) do not count as content.
func checkSection(r *store.Record, heading string, add func(rule, severity, file string, line int, msg string)) {
	want := "## " + heading
	lines := strings.Split(string(r.Body), "\n")
	start := -1
	for i, ln := range lines {
		if ln == want {
			start = i
			break
		}
	}
	if start < 0 {
		add("section", SeverityError, r.Rel, 0, fmt.Sprintf("missing required section %q", want))
		return
	}
	end := len(lines)
	for i := start + 1; i < len(lines); i++ {
		if strings.HasPrefix(lines[i], "## ") {
			end = i
			break
		}
	}
	substantive := 0
	for i := start + 1; i < end; i++ {
		ln := strings.TrimSpace(lines[i])
		if ln == "" || strings.HasPrefix(ln, "#") {
			continue
		}
		substantive++
		if strings.Contains(ln, "TODO") || placeholderRe.MatchString(ln) {
			add("section", SeverityError, r.Rel, r.BodyLine+i,
				fmt.Sprintf("section %q contains an unfilled placeholder: %s", want, ln))
		}
	}
	if substantive == 0 {
		add("section", SeverityError, r.Rel, r.BodyLine+start, fmt.Sprintf("section %q is empty", want))
	}
}

// validateSource checks bibliographic identity and claim locators.
func validateSource(r *store.Record, add func(rule, severity, file string, line int, msg string)) {
	checkSection(r, "Bibliographic identity", add)
	checkSection(r, "Claims", add)

	lines := strings.Split(string(r.Body), "\n")
	var doiOK, arxivOK, shaOK bool
	claims := []int{}
	for i, ln := range lines {
		trimmed := strings.TrimSpace(ln)
		switch {
		case strings.HasPrefix(trimmed, "doi:"):
			if valueOK(strings.TrimSpace(strings.TrimPrefix(trimmed, "doi:"))) {
				doiOK = true
			}
		case strings.HasPrefix(trimmed, "arxiv:"):
			if valueOK(strings.TrimSpace(strings.TrimPrefix(trimmed, "arxiv:"))) {
				arxivOK = true
			}
		case strings.HasPrefix(trimmed, "sha256:"):
			v := strings.TrimSpace(strings.TrimPrefix(trimmed, "sha256:"))
			if hex64Re.MatchString(v) {
				shaOK = true
			} else if valueOK(v) {
				add("bibliographic", SeverityError, r.Rel, r.BodyLine+i,
					"sha256 must be exactly 64 hex characters (hash of the exact file examined)")
			}
		case strings.HasPrefix(trimmed, "- "):
			claims = append(claims, i)
		}
	}

	if !doiOK && !arxivOK {
		add("bibliographic", SeverityError, r.Rel, 0,
			"bibliographic identity requires a non-placeholder doi: or arxiv: line")
	}
	if !shaOK {
		add("bibliographic", SeverityError, r.Rel, 0,
			"bibliographic identity requires a sha256: line with a 64-hex digest")
	}

	withLocator := 0
	for _, i := range claims {
		ln := strings.TrimSpace(lines[i])
		idx := strings.Index(ln, "locator:")
		if idx < 0 {
			add("source-locator", SeverityError, r.Rel, r.BodyLine+i,
				"claim has no 'locator:' (page, section, equation, figure or table)")
			continue
		}
		loc := strings.TrimSpace(ln[idx+len("locator:"):])
		loc = strings.TrimSuffix(loc, ")")
		loc = strings.TrimSpace(loc)
		if valueOK(loc) && !placeholderRe.MatchString(loc) {
			withLocator++
		}
	}
	if len(claims) == 0 {
		add("source-locator", SeverityError, r.Rel, 0, "source record declares no claims under '## Claims'")
	} else if withLocator == 0 {
		add("source-locator", SeverityError, r.Rel, 0, "no claim carries a concrete source locator")
	}
}

var hex64Re = regexp.MustCompile(`^[0-9a-f]{64}$`)

func valueOK(v string) bool {
	if v == "" || v == "none" {
		return false
	}
	if strings.HasPrefix(v, "<") || strings.HasSuffix(v, ">") {
		return false
	}
	return !placeholderRe.MatchString(v) && !strings.Contains(v, "TODO")
}

// validateCheckProfiles verifies declared derivation check profiles against
// the allowlist (placeholder values are unfilled, not invalid).
func validateCheckProfiles(r *store.Record, add func(rule, severity, file string, line int, msg string)) {
	lines := strings.Split(string(r.Body), "\n")
	inChecks := false
	for i, ln := range lines {
		if strings.HasPrefix(ln, "## ") {
			inChecks = ln == "## Checks"
			continue
		}
		if !inChecks {
			continue
		}
		trimmed := strings.TrimSpace(ln)
		if !strings.HasPrefix(trimmed, "- profile:") {
			continue
		}
		name := strings.TrimSpace(strings.TrimPrefix(trimmed, "- profile:"))
		if name == "" || placeholderRe.MatchString(name) || strings.Contains(name, "TODO") {
			continue // unfilled optional placeholder
		}
		if !contains(check.Names(), name) {
			add("check-profile", SeverityError, r.Rel, r.BodyLine+i,
				fmt.Sprintf("unknown check profile %q (allowlisted: %s)", name, strings.Join(check.Names(), ", ")))
		}
	}
}

// validateReviews enforces review provenance: target existence, digest
// match (staleness), and reviewer identity fields.
func validateReviews(recs []store.Record, byID map[string][]*store.Record, root string, add func(rule, severity, file string, line int, msg string)) {
	for i := range recs {
		r := &recs[i]
		if r.ParseErr != nil || r.Header.Type != record.TypeReview {
			continue
		}
		lines := strings.Split(string(r.Body), "\n")
		var reviewed, digest string
		var reviewedLine, digestLine int
		var reviewerOK, modelOK bool
		for i, ln := range lines {
			trimmed := strings.TrimSpace(ln)
			switch {
			case strings.HasPrefix(trimmed, "reviewed:") && reviewed == "":
				reviewed = strings.TrimSpace(strings.TrimPrefix(trimmed, "reviewed:"))
				reviewedLine = r.BodyLine + i
			case strings.HasPrefix(trimmed, "digest:") && digest == "":
				digest = strings.TrimSpace(strings.TrimPrefix(trimmed, "digest:"))
				digestLine = r.BodyLine + i
			case strings.HasPrefix(trimmed, "reviewer:"):
				reviewerOK = valueOK(strings.TrimSpace(strings.TrimPrefix(trimmed, "reviewer:")))
			case strings.HasPrefix(trimmed, "model:"):
				modelOK = valueOK(strings.TrimSpace(strings.TrimPrefix(trimmed, "model:")))
			}
		}

		if reviewed == "" {
			add("review", SeverityError, r.Rel, 0, "review must record 'reviewed: <hypothesis id>'")
		}
		if !reviewerOK {
			add("review", SeverityError, r.Rel, 0, "review must record a non-placeholder 'reviewer:' identity")
		}
		if !modelOK {
			add("review", SeverityError, r.Rel, 0, "review must record a non-placeholder 'model:' (self-reported)")
		}
		if digest == "" || !strings.HasPrefix(digest, "sha256:") || !hex64Re.MatchString(strings.TrimPrefix(digest, "sha256:")) {
			add("review", SeverityError, r.Rel, 0, "review must record 'digest: sha256:<64 hex>' of the reviewed version")
		}

		if reviewed == "" {
			continue
		}
		targets, ok := byID[reviewed]
		if !ok || len(targets) == 0 {
			add("review", SeverityError, r.Rel, reviewedLine,
				fmt.Sprintf("reviewed target %s not found", reviewed))
			continue
		}
		target := targets[0]
		if target.Header.Type != record.TypeHypothesis {
			add("review", SeverityError, r.Rel, reviewedLine,
				fmt.Sprintf("reviewed target %s is a %s; reviews target hypotheses", reviewed, target.Header.Type))
			continue
		}
		if strings.HasPrefix(digest, "sha256:") && hex64Re.MatchString(strings.TrimPrefix(digest, "sha256:")) {
			want := "sha256:" + target.Digest
			if digest != want {
				add("stale", SeverityError, r.Rel, digestLine,
					fmt.Sprintf("STALE: reviewed digest %s does not match current %s (sha256:%s); write a new review",
						digest, target.Header.ID, target.Digest))
			}
		}
	}
	validateReviewsAppendOnly(root, add)
}

// validateReviewsAppendOnly flags review files that were modified after
// creation, using git history and the working tree when a repository is
// present. Absent git is not an error: the check is simply skipped.
func validateReviewsAppendOnly(root string, add func(rule, severity, file string, line int, msg string)) {
	if _, err := exec.LookPath("git"); err != nil {
		return
	}
	modified := map[string]bool{}

	run := func(args ...string) []byte {
		cmd := exec.Command("git", append([]string{"-C", root}, args...)...)
		var out bytes.Buffer
		cmd.Stdout = &out
		if err := cmd.Run(); err != nil {
			return nil
		}
		return out.Bytes()
	}

	// Committed modifications.
	out := run("log", "--diff-filter=M", "--name-only", "--pretty=format:", "--", "research/reviews/")
	for _, ln := range strings.Split(string(out), "\n") {
		ln = strings.TrimSpace(ln)
		if ln != "" {
			modified[ln] = true
		}
	}
	// Uncommitted modifications in the working tree. Porcelain status is
	// "XY path": only an M in the staged or unstaged column is a
	// modification; new files (??, A) are legitimate creations.
	out = run("status", "--porcelain", "--", "research/reviews/")
	for _, ln := range strings.Split(string(out), "\n") {
		ln = strings.TrimRight(ln, "\r")
		if len(ln) > 3 && strings.Contains(ln[:2], "M") {
			modified[filepath.ToSlash(strings.TrimSpace(ln[3:]))] = true
		}
	}

	var files []string
	for f := range modified {
		files = append(files, f)
	}
	sort.Strings(files)
	for _, f := range files {
		add("append-only", SeverityError, f, 0,
			"review file modified after creation; reviews are append-only — create a new review instead")
	}
}

// validateDependsOnCycles reports directed cycles in the depends_on graph
// only. Cycles among other relations (mutual challenges, etc.) are
// legitimate science and are never validation errors (AGENTS.md §8).
func validateDependsOnCycles(recs []store.Record, add func(rule, severity, file string, line int, msg string)) {
	graph := map[string][]string{}
	var nodes []string
	for _, r := range recs {
		if r.ParseErr != nil {
			continue
		}
		id := r.Header.ID
		if id == "" {
			continue
		}
		nodes = append(nodes, id)
		for _, ref := range r.Header.Refs {
			if ref.Relation == "depends_on" && ref.ID != "" && ref.ID != id {
				graph[id] = append(graph[id], ref.ID)
			}
		}
	}
	sort.Strings(nodes)

	// Depth-first search reporting each directed cycle once.
	const (
		white = 0
		grey  = 1
		black = 2
	)
	color := map[string]int{}
	var stack []string
	reported := map[string]bool{}
	relFile := map[string]string{}
	for _, r := range recs {
		if r.ParseErr == nil && r.Header.ID != "" {
			if _, ok := relFile[r.Header.ID]; !ok {
				relFile[r.Header.ID] = r.Rel
			}
		}
	}

	var dfs func(n string)
	dfs = func(n string) {
		color[n] = grey
		stack = append(stack, n)
		edges := append([]string{}, graph[n]...)
		sort.Strings(edges)
		for _, m := range edges {
			switch color[m] {
			case white:
				if _, exists := relFile[m]; exists {
					dfs(m)
				}
			case grey:
				// Cycle: from m to end of stack.
				idx := -1
				for i, s := range stack {
					if s == m {
						idx = i
						break
					}
				}
				if idx < 0 {
					continue
				}
				cycle := append(append([]string{}, stack[idx:]...), m)
				key := canonicalCycle(cycle[:len(cycle)-1])
				if !reported[key] {
					reported[key] = true
					file := relFile[cycle[0]]
					add("cycle", SeverityError, file, 0,
						fmt.Sprintf("depends_on cycle: %s", strings.Join(cycle, " -> ")))
				}
			}
		}
		stack = stack[:len(stack)-1]
		color[n] = black
	}

	for _, n := range nodes {
		if color[n] == white {
			dfs(n)
		}
	}
}

// canonicalCycle rotates a cycle so it starts at its lexicographically
// smallest member, for deterministic dedup.
func canonicalCycle(cycle []string) string {
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
	return strings.Join(rot, ">")
}

func contains(list []string, v string) bool {
	for _, s := range list {
		if s == v {
			return true
		}
	}
	return false
}
