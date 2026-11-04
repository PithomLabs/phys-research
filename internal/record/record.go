// Package record defines record types, ID allocation and creation of new
// records from templates. IDs are readable and allocated by this package
// (never content hashes); creation is atomic and never overwrites an
// existing file.
package record

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/PithomLabs/phys-research/internal/frontmatter"
	"github.com/PithomLabs/phys-research/internal/store"
)

// Supported record types, in ID-prefix order.
const (
	TypeSource     = "source"
	TypeIssue      = "issue"
	TypeHypothesis = "hypothesis"
	TypeReview     = "review"
	TypeEvidence   = "evidence"
	TypeDerivation = "derivation"
)

// PrefixByType maps each record type to its ID letter (AGENTS.md §6.2).
var PrefixByType = map[string]string{
	TypeSource:     "P",
	TypeIssue:      "I",
	TypeHypothesis: "H",
	TypeReview:     "R",
	TypeEvidence:   "E",
	TypeDerivation: "D",
}

// DirByType maps each record type to its directory (AGENTS.md §6.3).
// Derivation records live at the top level of derivations/.
var DirByType = map[string]string{
	TypeSource:     "research/sources",
	TypeIssue:      "research/issues",
	TypeHypothesis: "research/hypotheses",
	TypeReview:     "research/reviews",
	TypeEvidence:   "research/evidence",
	TypeDerivation: "derivations",
}

// Closed vocabularies (AGENTS.md §6.1). Validation of values happens in
// the validate package; these exist for creation-time fail-fast checks
// and for command help.
var (
	States     = []string{"open", "blocked", "closed"}
	Priorities = []string{"high", "medium", "low"}
	Branches   = []string{"Q", "C"}
	Classes    = []string{"physical", "mathematical"}
	Relations  = []string{"motivated_by", "supports", "challenges", "addresses", "depends_on", "supersedes"}
)

// Types lists supported record types.
func Types() []string {
	return []string{TypeSource, TypeIssue, TypeHypothesis, TypeReview, TypeEvidence, TypeDerivation}
}

// IDPattern matches readable record IDs such as P-001.
var IDPattern = regexp.MustCompile(`^([A-Z])-([0-9]{3,})$`)

// Options are the creation parameters accepted by physctl new.
type Options struct {
	Type     string
	Title    string
	Branch   string
	Priority string
	Class    string
	Review   string // target record ID; required for type review
	Refs     []frontmatter.Ref
}

// Created reports what was written.
type Created struct {
	ID  string
	Rel string
}

// Create allocates the next ID for the record type and writes the record
// from its template. It never overwrites an existing file: the final write
// uses an atomic no-clobber operation and retries ID allocation on race.
func Create(root string, opts Options) (Created, error) {
	prefix, ok := PrefixByType[opts.Type]
	if !ok {
		return Created{}, fmt.Errorf("unknown record type %q (supported: %s)", opts.Type, strings.Join(Types(), ", "))
	}
	if opts.Class != "" && opts.Type != TypeHypothesis {
		return Created{}, errors.New("class is only valid on hypothesis records")
	}
	if opts.Type == TypeHypothesis {
		if opts.Class == "" {
			return Created{}, errors.New("hypothesis records require --class physical|mathematical")
		}
		if !contains(Classes, opts.Class) {
			return Created{}, fmt.Errorf("invalid class %q (supported: %s)", opts.Class, strings.Join(Classes, ", "))
		}
	}
	if opts.Branch != "" && !contains(Branches, opts.Branch) {
		return Created{}, fmt.Errorf("invalid branch %q (supported: %s)", opts.Branch, strings.Join(Branches, ", "))
	}
	if opts.Priority != "" && !contains(Priorities, opts.Priority) {
		return Created{}, fmt.Errorf("invalid priority %q (supported: %s)", opts.Priority, strings.Join(Priorities, ", "))
	}
	for _, r := range opts.Refs {
		if !contains(Relations, r.Relation) {
			return Created{}, fmt.Errorf("invalid relation %q (supported: %s)", r.Relation, strings.Join(Relations, ", "))
		}
		if !IDPattern.MatchString(r.ID) {
			return Created{}, fmt.Errorf("invalid reference id %q (expected e.g. P-001)", r.ID)
		}
	}

	if opts.Type == TypeReview {
		if opts.Review == "" {
			return Created{}, errors.New("review records require --review <hypothesis id>")
		}
	} else if opts.Review != "" {
		return Created{}, errors.New("--review is only valid for --type review")
	}

	// Resolve the review target and its current digest before allocating.
	reviewDigest := ""
	if opts.Type == TypeReview {
		recs, err := store.Load(root)
		if err != nil {
			return Created{}, err
		}
		var target *store.Record
		for i := range recs {
			if recs[i].ParseErr == nil && recs[i].Header.ID == opts.Review {
				target = &recs[i]
				break
			}
		}
		if target == nil {
			return Created{}, fmt.Errorf("review target %q not found", opts.Review)
		}
		if target.Header.Type != TypeHypothesis {
			return Created{}, fmt.Errorf("review target %q is a %s, not a hypothesis", opts.Review, target.Header.Type)
		}
		reviewDigest = target.Digest
	}

	for attempt := 0; attempt < 1000; attempt++ {
		id, err := nextID(root, prefix)
		if err != nil {
			return Created{}, err
		}
		rel := expectedRel(opts.Type, id)
		data := render(opts, id, reviewDigest)
		err = writeNew(filepath.Join(root, filepath.FromSlash(rel)), data)
		if err == nil {
			return Created{ID: id, Rel: rel}, nil
		}
		if errors.Is(err, os.ErrExist) {
			continue // raced with another writer; rescan for the next ID
		}
		return Created{}, err
	}
	return Created{}, errors.New("could not allocate an ID after repeated collisions")
}

// nextID returns the lowest unused ID of the form PREFIX-NNN by scanning
// record filenames (which validate enforces to equal header IDs).
func nextID(root, prefix string) (string, error) {
	seen := map[int]bool{}
	scan := func(dir string) error {
		entries, err := os.ReadDir(dir)
		if err != nil {
			if os.IsNotExist(err) {
				return nil
			}
			return err
		}
		for _, e := range entries {
			if e.IsDir() || !strings.HasSuffix(e.Name(), ".md") || strings.HasPrefix(e.Name(), ".") {
				continue
			}
			m := IDPattern.FindStringSubmatch(strings.TrimSuffix(e.Name(), ".md"))
			if m == nil || m[1] != prefix {
				continue
			}
			n, err := strconv.Atoi(m[2])
			if err != nil {
				continue
			}
			seen[n] = true
		}
		return nil
	}
	for _, d := range store.RecordDirs {
		if err := scan(filepath.Join(root, d)); err != nil {
			return "", fmt.Errorf("record: scan %s: %w", d, err)
		}
	}
	if err := scan(filepath.Join(root, store.DerivationsDir)); err != nil && !os.IsNotExist(err) {
		return "", fmt.Errorf("record: scan derivations: %w", err)
	}
	for n := 1; n <= 1_000_000; n++ {
		if !seen[n] {
			return fmt.Sprintf("%s-%03d", prefix, n), nil
		}
	}
	return "", errors.New("record: ID space exhausted")
}

// expectedRel is the canonical location for a record of the given type.
func expectedRel(typ, id string) string {
	return DirByType[typ] + "/" + id + ".md"
}

// writeNew writes data to path atomically without replacing an existing
// file: a temp file in the same directory is linked into place (O_EXCL
// semantics), and falls back to O_CREATE|O_EXCL where hard links are
// unavailable.
func writeNew(path string, data []byte) error {
	dir := filepath.Dir(path)
	// Record directories are not tracked while empty, so a fresh clone may
	// lack them. Paths only ever come from DirByType, so this is safe.
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}
	tmp, err := os.CreateTemp(dir, ".physctl-new-*")
	if err != nil {
		return err
	}
	tmpName := tmp.Name()
	cleanup := func() { _ = tmp.Close(); _ = os.Remove(tmpName) }

	if _, err := tmp.Write(data); err != nil {
		cleanup()
		return err
	}
	if err := tmp.Chmod(0o644); err != nil {
		cleanup()
		return err
	}
	if err := tmp.Close(); err != nil {
		_ = os.Remove(tmpName)
		return err
	}
	defer os.Remove(tmpName)

	if err := os.Link(tmpName, path); err != nil {
		if errors.Is(err, os.ErrExist) {
			return os.ErrExist
		}
		// Filesystem without usable hard links: fall back to O_EXCL.
		f, oerr := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0o644)
		if oerr != nil {
			if errors.Is(oerr, os.ErrExist) {
				return os.ErrExist
			}
			return oerr
		}
		if _, werr := f.Write(data); werr != nil {
			_ = f.Close()
			return werr
		}
		return f.Close()
	}
	return nil
}

// render builds the complete record file: frontmatter then body.
func render(opts Options, id, reviewDigest string) []byte {
	var b strings.Builder
	b.WriteString("---\n")
	fmt.Fprintf(&b, "id: %s\n", id)
	fmt.Fprintf(&b, "type: %s\n", opts.Type)
	b.WriteString("state: open\n")
	if opts.Branch != "" {
		fmt.Fprintf(&b, "branch: %s\n", opts.Branch)
	}
	if opts.Priority != "" {
		fmt.Fprintf(&b, "priority: %s\n", opts.Priority)
	}
	if opts.Class != "" {
		fmt.Fprintf(&b, "class: %s\n", opts.Class)
	}
	if len(opts.Refs) > 0 {
		b.WriteString("refs:\n")
		for _, r := range opts.Refs {
			fmt.Fprintf(&b, "  - relation: %s\n    id: %s\n", r.Relation, r.ID)
		}
	}
	b.WriteString("---\n\n")
	b.WriteString(template(opts, id, reviewDigest))
	return []byte(b.String())
}

func template(opts Options, id, reviewDigest string) string {
	title := opts.Title
	if title == "" {
		title = "Untitled"
	}
	switch opts.Type {
	case TypeIssue:
		return fmt.Sprintf(`# %s

## Question

<the outstanding problem, stated precisely>

## Theoretical expectation

- Claim or prediction: <what is claimed or predicted>
- Assumptions: <conditions the claim rests on>
- Applicable regime: <energy, scale, curvature or matter conditions>
- Source or derivation reference: <record ID plus locator, or "none">

## Independent theoretical result

NOT APPLICABLE — no independent theoretical derivation is required for this experimental result.
# Replace the line above only if an independent derivation actually exists:
# state what was reconstructed or derived, its premises and scope, and the
# derivation artifact or source reference. Never label a paper's assertion
# as an independent derivation.

## Experimental expectation

- Predicted observable: <quantity that could be measured>
- Units and conditions: <units, experimental conditions>
- Uncertainty: <expected uncertainty where applicable>
- Falsification condition: <what would count against the claim>

## Experimental actual

- Measured result: <measured value, or "not yet measured">
- Units and uncertainty: <units; statistical and systematic uncertainties>
- Methodology and experimental conditions: <how it was measured>
- Exact source locator: <equation, figure, table, page and source ID>
`, title)

	case TypeHypothesis:
		if opts.Class == "mathematical" {
			return fmt.Sprintf(`# %s

<Phys-Pseudo content above this line: exactly the seven constructs of
phys-pseudo/PSEUDO.md §1. Attribution is prose or # comments (§8). A
mathematical hypothesis carries proof obligations, not empirical
falsifiers.>

## Premises

<premises this hypothesis adopts>

## Conclusion

<the proposition to be established>

## Domain

<the mathematical domain in which the conclusion is claimed>

## Proof obligations

<what must be proved, and by which derivation artifacts>
`, title)
		}
		return fmt.Sprintf(`# %s

<Phys-Pseudo content above this line: exactly the seven constructs of
phys-pseudo/PSEUDO.md §1. Attribution is prose or # comments (§8). Never
write a falsification construct into the pseudo content: the condition
belongs in the section below.>

## Falsification condition

<falsification condition: the observation, measurement or re-analysis that
would count against this hypothesis, with units and regime where
applicable. Declare any total-uncertainty prescription explicitly.>
`, title)

	case TypeSource:
		return fmt.Sprintf(`# %s

## Bibliographic identity

title: <full title>
authors: <authors>
doi: <DOI, or "none">
arxiv: <arXiv ID with version, or "none">
sha256: <SHA-256 of the exact file examined; the file itself is never committed>

## Claims

- <claim text> (locator: <page, section, equation, figure or table>)
# Add one bullet per extracted claim. Every claim needs a locator.
# Attribution of what the paper states vs. what is derived is prose (PSEUDO.md §8).
`, title)

	case TypeReview:
		return fmt.Sprintf(`# %s

## Review target

reviewed: %s
digest: sha256:%s
reviewer: adversary (OpenCode agent role)
model: <self-reported model id used for this review>

## Summary

<one paragraph: what was examined, from what vantage point>

## Blocking findings

none

## Nonblocking findings

none

## Finding details

### F1 — <short title>

- Where: <the affected statement or step, named>
- What: <the defect, classified>
- Why: <supporting reasoning or counterexample>
- Minimum fix: <the smallest correction or evidence that would resolve it>

## Evidential boundary

<what was actually checked and how; what could not be checked; what
remains unresolved. No global correct/incorrect verdict, and nothing
called "verified" beyond the checks actually performed
(MRC_REVIEW_PROTOCOL.md §8).>
`, title, opts.Review, reviewDigest)

	case TypeEvidence:
		return fmt.Sprintf(`# %s

## Measured value

<the measured quantity and value exactly as published>

## Units and uncertainty

<units; statistical and systematic uncertainties recorded separately as published>

## Methodology

<method, experimental conditions, sessions, data treatment>

## Source locator

<exact locator: equation, figure, table, page and version of the source record>
`, title)

	case TypeDerivation:
		return fmt.Sprintf(`# %s

## Premises

<premises this derivation adopts>

## Dependencies

<record IDs and artifact references this derivation depends on>

## Intermediate results

<intermediate results and where they come from>

## Mathematical rules

<the exact PhysMath/PhysLib rule identities used, e.g. DIFFERENTIATE-001>

## Obligations

<proof obligations that remain unresolved>

## Checks

- profile: <allowlisted physctl check profile name>
`, title)
	}
	return fmt.Sprintf("# %s\n", title)
}

func contains(list []string, v string) bool {
	for _, s := range list {
		if s == v {
			return true
		}
	}
	return false
}

// SortedIDs returns the IDs of all records, sorted, for deterministic
// allocation diagnostics.
func SortedIDs(root string) ([]string, error) {
	recs, err := store.Load(root)
	if err != nil {
		return nil, err
	}
	var ids []string
	for _, r := range recs {
		if r.ParseErr == nil && r.Header.ID != "" {
			ids = append(ids, r.Header.ID)
		}
	}
	sort.Strings(ids)
	return ids, nil
}
