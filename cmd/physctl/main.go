// Command physctl is the deterministic tooling layer of the phys-research
// workbench: record commands (new, status, refs, validate, check), paper
// acquisition (paper search|related|fetch|extract) and distribution checks
// (dist check-tar|check-staged).
//
// It is not an agent, not a scheduler and not a scientific authority.
// Structural validation never establishes mathematical or physical
// correctness (see internal/validate), and check runs only allowlisted
// tests of the existing Phys infrastructure (see internal/check).
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/PithomLabs/phys-research/internal/check"
	"github.com/PithomLabs/phys-research/internal/distrib"
	"github.com/PithomLabs/phys-research/internal/frontmatter"
	"github.com/PithomLabs/phys-research/internal/paper"
	"github.com/PithomLabs/phys-research/internal/record"
	"github.com/PithomLabs/phys-research/internal/refs"
	"github.com/PithomLabs/phys-research/internal/status"
	"github.com/PithomLabs/phys-research/internal/store"
	"github.com/PithomLabs/phys-research/internal/validate"
)

const usage = `physctl — deterministic research-record tooling (structural only)

Usage:
  physctl new      --type <type> [--title t] [--branch Q|C] [--priority high|medium|low]
                   [--class physical|mathematical] [--ref relation=ID]... [--review ID] [--json]
  physctl status   [--branch C] [--json]
  physctl refs     <ID> [--depth N] [--in] [--json]
  physctl validate [--json]
  physctl check    [<profile> | derivations/<file> | <derivation-id>] [--json]
  physctl paper search --query <q> [--limit N] [--json]
  physctl paper related <paperID> [--limit N] [--json]
  physctl paper fetch --version v1 <paperID>
  physctl paper extract --version v1 <paperID>
  physctl dist check-tar <tarball> [--json]
  physctl dist check-staged --dir <dir> --manifest <file> --base <file> [--ignore <rel,rel>] [--json]

Exit codes: 0 clean/pass, 1 findings or test failures,
2 usage errors or unavailable capabilities.

Validation is structural only: it never establishes scientific correctness.
`

func main() {
	os.Exit(run(os.Args[1:], os.Stdout, os.Stderr))
}

func run(args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		fmt.Fprint(stderr, usage)
		return 2
	}
	switch args[0] {
	case "new":
		return cmdNew(args[1:], stdout, stderr)
	case "status":
		return cmdStatus(args[1:], stdout, stderr)
	case "refs":
		return cmdRefs(args[1:], stdout, stderr)
	case "validate":
		return cmdValidate(args[1:], stdout, stderr)
	case "check":
		return cmdCheck(args[1:], stdout, stderr)
	case "paper":
		return cmdPaper(args[1:], stdout, stderr)
	case "dist":
		return cmdDist(args[1:], stdout, stderr)
	case "-h", "--help", "help":
		fmt.Fprint(stdout, usage)
		return 0
	default:
		fmt.Fprintf(stderr, "physctl: unknown command %q\n\n", args[0])
		fmt.Fprint(stderr, usage)
		return 2
	}
}

// findRoot locates the repository root: $PHYS_RESEARCH_ROOT, otherwise the
// nearest ancestor containing a research/ directory, otherwise the cwd.
func findRoot() (string, error) {
	if v := os.Getenv("PHYS_RESEARCH_ROOT"); v != "" {
		return v, nil
	}
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	dir := wd
	for {
		if fi, err := os.Stat(filepath.Join(dir, "research")); err == nil && fi.IsDir() {
			return dir, nil
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return wd, nil
		}
		dir = parent
	}
}

// refFlags collects repeatable --ref relation=ID flags.
type refFlags []string

func (r *refFlags) String() string { return strings.Join(*r, ",") }
func (r *refFlags) Set(v string) error {
	*r = append(*r, v)
	return nil
}

func cmdNew(args []string, stdout, stderr io.Writer) int {
	fs := flag.NewFlagSet("new", flag.ContinueOnError)
	fs.SetOutput(stderr)
	typ := fs.String("type", "", "record type: source|issue|hypothesis|review|evidence|derivation")
	title := fs.String("title", "", "record title")
	branch := fs.String("branch", "", "Q or C")
	priority := fs.String("priority", "", "high|medium|low")
	class := fs.String("class", "", "physical|mathematical (hypotheses only)")
	review := fs.String("review", "", "target hypothesis ID (required for --type review)")
	jsonOut := fs.Bool("json", false, "machine-readable output")
	var refList refFlags
	fs.Var(&refList, "ref", "relationship as relation=ID, e.g. motivated_by=P-001 (repeatable)")

	if err := fs.Parse(args); err != nil {
		return 2
	}
	if *typ == "" {
		fmt.Fprintln(stderr, "physctl new: --type is required")
		return 2
	}

	var rf []frontmatter.Ref
	for _, s := range refList {
		parts := strings.SplitN(s, "=", 2)
		if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
			fmt.Fprintf(stderr, "physctl new: invalid --ref %q (expected relation=ID)\n", s)
			return 2
		}
		rf = append(rf, frontmatter.Ref{Relation: parts[0], ID: parts[1]})
	}

	root, err := findRoot()
	if err != nil {
		fmt.Fprintf(stderr, "physctl new: %v\n", err)
		return 2
	}

	created, err := record.Create(root, record.Options{
		Type:     *typ,
		Title:    *title,
		Branch:   *branch,
		Priority: *priority,
		Class:    *class,
		Review:   *review,
		Refs:     rf,
	})
	if err != nil {
		fmt.Fprintf(stderr, "physctl new: %v\n", err)
		return 2
	}

	if *jsonOut {
		return writeJSON(stdout, map[string]string{"id": created.ID, "path": created.Rel})
	}
	fmt.Fprintf(stdout, "created %s (id %s)\n", created.Rel, created.ID)
	return 0
}

func cmdStatus(args []string, stdout, stderr io.Writer) int {
	fs := flag.NewFlagSet("status", flag.ContinueOnError)
	fs.SetOutput(stderr)
	branch := fs.String("branch", "", "filter issues by Q/C branch")
	jsonOut := fs.Bool("json", false, "machine-readable output")
	if err := fs.Parse(args); err != nil {
		return 2
	}
	root, err := findRoot()
	if err != nil {
		fmt.Fprintf(stderr, "physctl status: %v\n", err)
		return 2
	}
	out, err := status.Collect(root, *branch)
	if err != nil {
		fmt.Fprintf(stderr, "physctl status: %v\n", err)
		return 2
	}
	if *jsonOut {
		return writeJSON(stdout, out)
	}
	fmt.Fprint(stdout, status.Text(out))
	return 0
}

func cmdRefs(args []string, stdout, stderr io.Writer) int {
	fs := flag.NewFlagSet("refs", flag.ContinueOnError)
	fs.SetOutput(stderr)
	depth := fs.Int("depth", 1, "maximum traversal depth")
	in := fs.Bool("in", false, "also follow incoming references")
	jsonOut := fs.Bool("json", false, "machine-readable output")
	if err := fs.Parse(args); err != nil {
		return 2
	}
	if fs.NArg() != 1 {
		fmt.Fprintln(stderr, "physctl refs: exactly one record ID is required")
		return 2
	}
	if *depth < 1 {
		fmt.Fprintln(stderr, "physctl refs: --depth must be >= 1")
		return 2
	}
	root, err := findRoot()
	if err != nil {
		fmt.Fprintf(stderr, "physctl refs: %v\n", err)
		return 2
	}
	out, err := refs.Traverse(root, fs.Arg(0), *depth, *in)
	if err != nil {
		fmt.Fprintf(stderr, "physctl refs: %v\n", err)
		if strings.Contains(err.Error(), "not found") {
			return 1
		}
		return 2
	}
	if *jsonOut {
		return writeJSON(stdout, out)
	}
	fmt.Fprintf(stdout, "%s (max depth %d)\n", out.Start, out.Depth)
	fmt.Fprintf(stdout, "  nodes:   %s\n", strings.Join(out.Nodes, ", "))
	fmt.Fprintf(stdout, "  edges:\n")
	if len(out.Edges) == 0 {
		fmt.Fprintln(stdout, "    (none)")
	}
	for _, e := range out.Edges {
		fmt.Fprintf(stdout, "    %s %s -[%s]-> %s  (%s)\n", e.Direction, e.From, e.Relation, e.To, e.File)
	}
	fmt.Fprintf(stdout, "  cycles:  %s\n", orNone(out.Cycles))
	fmt.Fprintf(stdout, "  missing: %s\n", orNone(out.Missing))
	return 0
}

func cmdValidate(args []string, stdout, stderr io.Writer) int {
	fs := flag.NewFlagSet("validate", flag.ContinueOnError)
	fs.SetOutput(stderr)
	jsonOut := fs.Bool("json", false, "machine-readable output")
	if err := fs.Parse(args); err != nil {
		return 2
	}
	root, err := findRoot()
	if err != nil {
		fmt.Fprintf(stderr, "physctl validate: %v\n", err)
		return 2
	}
	res, err := validate.Run(root)
	if err != nil {
		fmt.Fprintf(stderr, "physctl validate: %v\n", err)
		return 2
	}
	if *jsonOut {
		code := 0
		if res.Errors > 0 {
			code = 1
		}
		if err := writeJSON(stdout, res); err != 0 {
			return err
		}
		return code
	}
	for _, f := range res.Findings {
		loc := f.File
		if f.Line > 0 {
			loc = fmt.Sprintf("%s:%d", f.File, f.Line)
		}
		fmt.Fprintf(stdout, "%s [%s] %s: %s\n", loc, f.Severity, f.Rule, f.Message)
	}
	fmt.Fprintf(stdout, "%d error(s), %d warning(s)\n", res.Errors, res.Warnings)
	fmt.Fprintln(stdout, validate.Disclaimer)
	if res.Errors > 0 {
		return 1
	}
	return 0
}

func cmdCheck(args []string, stdout, stderr io.Writer) int {
	fs := flag.NewFlagSet("check", flag.ContinueOnError)
	fs.SetOutput(stderr)
	jsonOut := fs.Bool("json", false, "machine-readable output")
	if err := fs.Parse(args); err != nil {
		return 2
	}
	root, err := findRoot()
	if err != nil {
		fmt.Fprintf(stderr, "physctl check: %v\n", err)
		return 2
	}

	// No argument: list the allowlist (never runs anything, never PASSes).
	if fs.NArg() == 0 {
		if *jsonOut {
			type item struct {
				Name   string `json:"name"`
				Module string `json:"module"`
			}
			var items []item
			for _, p := range check.Profiles {
				items = append(items, item{p.Name, p.Module})
			}
			return writeJSON(stdout, map[string]any{"profiles": items})
		}
		fmt.Fprintln(stdout, "Allowlisted check profiles:")
		for _, p := range check.Profiles {
			fmt.Fprintf(stdout, "  %-22s %s\n", p.Name, p.Module)
		}
		return 0
	}

	arg := fs.Arg(0)

	// Derivation record: run the profiles it declares (allowlisted only).
	names, ok := derivationProfiles(root, arg)
	if ok {
		if len(names) == 0 {
			fmt.Fprintf(stderr, "physctl check: %s declares no check profiles; nothing was run\n", arg)
			return 2
		}
		var results []check.Result
		for _, n := range names {
			results = append(results, check.Run(root, n))
		}
		return reportResults(stdout, stderr, results, *jsonOut)
	}

	// Single allowlisted profile.
	res := check.Run(root, arg)
	return reportResults(stdout, stderr, []check.Result{res}, *jsonOut)
}

// derivationProfiles resolves arg as a derivation record (relative path or
// record ID) and returns the profiles it declares. ok is false when arg is
// not a derivation record.
func derivationProfiles(root, arg string) ([]string, bool) {
	clean := strings.TrimPrefix(arg, "./")
	recs, err := store.Load(root)
	if err != nil {
		return nil, false
	}
	for i := range recs {
		r := &recs[i]
		if r.ParseErr != nil {
			continue
		}
		if (strings.HasPrefix(clean, "derivations/") && r.Rel == clean) ||
			(r.Header.Type == record.TypeDerivation && r.Header.ID == arg) {
			return check.ProfilesFromBody(string(r.Body)), true
		}
	}
	return nil, false
}

func reportResults(stdout, stderr io.Writer, results []check.Result, jsonOut bool) int {
	if jsonOut {
		if err := writeJSON(stdout, map[string]any{"results": results}); err != 0 {
			return err
		}
	} else {
		for _, r := range results {
			fmt.Fprintf(stdout, "%-12s %s", r.Status, r.Profile)
			if r.Command != "" {
				fmt.Fprintf(stdout, "  (exit %d, %dms)", r.ExitCode, r.DurationMS)
			}
			fmt.Fprintln(stdout)
			if r.Diagnostic != "" {
				fmt.Fprintf(stdout, "  ! %s\n", r.Diagnostic)
			}
			if r.Command != "" {
				fmt.Fprintf(stdout, "  $ %s\n", r.Command)
			}
			for _, ln := range strings.Split(r.Output, "\n") {
				if ln != "" {
					fmt.Fprintf(stdout, "  | %s\n", ln)
				}
			}
		}
		fmt.Fprintln(stdout, "A PASS means the named tests passed; it does not establish scientific correctness.")
	}

	fail, unavailable := false, false
	for _, r := range results {
		switch r.Status {
		case check.StatusFail:
			fail = true
		case check.StatusUnavailable:
			unavailable = true
		}
	}
	switch {
	case fail:
		return 1
	case unavailable:
		return 2
	}
	return 0
}

func cmdDist(args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		fmt.Fprintln(stderr, "physctl dist: requires a subcommand (check-tar, check-staged)")
		return 2
	}
	switch args[0] {
	case "check-tar":
		return cmdDistCheckTar(args[1:], stdout, stderr)
	case "check-staged":
		return cmdDistCheckStaged(args[1:], stdout, stderr)
	default:
		fmt.Fprintf(stderr, "physctl dist: unknown subcommand %q\n", args[0])
		return 2
	}
}

// cmdDistCheckTar validates a distribution tarball before extraction.
// Exit 0: clean (entries listed); 1: violations; 2: usage or IO errors.
func cmdDistCheckTar(args []string, stdout, stderr io.Writer) int {
	fs := flag.NewFlagSet("check-tar", flag.ContinueOnError)
	fs.SetOutput(stderr)
	jsonOut := fs.Bool("json", false, "machine-readable output")
	if err := fs.Parse(args); err != nil {
		return 2
	}
	if fs.NArg() != 1 {
		fmt.Fprintln(stderr, "physctl dist check-tar: exactly one tarball path is required")
		return 2
	}
	entries, err := distrib.ValidateTar(fs.Arg(0))
	if err != nil {
		if os.IsNotExist(err) || strings.Contains(err.Error(), "read tar") {
			fmt.Fprintf(stderr, "physctl dist check-tar: %v\n", err)
			return 2
		}
		if *jsonOut {
			writeJSON(stdout, map[string]any{"ok": false, "violation": err.Error()})
		} else {
			fmt.Fprintf(stdout, "REJECTED: %v\n", err)
		}
		return 1
	}
	if *jsonOut {
		return writeJSON(stdout, map[string]any{"ok": true, "entries": entries})
	}
	fmt.Fprintf(stdout, "OK: %d entries\n", len(entries))
	for _, e := range entries {
		fmt.Fprintf(stdout, "  %s\n", e)
	}
	return 0
}

// cmdDistCheckStaged verifies an assembled staging tree against a payload
// manifest and the base listing it was assembled over.
// Exit 0: verified; 1: mismatch; 2: usage or IO errors.
func cmdDistCheckStaged(args []string, stdout, stderr io.Writer) int {
	fs := flag.NewFlagSet("check-staged", flag.ContinueOnError)
	fs.SetOutput(stderr)
	dir := fs.String("dir", "", "staging tree to verify")
	manifest := fs.String("manifest", "", "payload manifest file")
	base := fs.String("base", "", "base listing file (one path per line)")
	ignore := fs.String("ignore", "", "comma-separated relative paths to excuse (e.g. go.work)")
	jsonOut := fs.Bool("json", false, "machine-readable output")
	if err := fs.Parse(args); err != nil {
		return 2
	}
	if *dir == "" || *manifest == "" || *base == "" {
		fmt.Fprintln(stderr, "physctl dist check-staged: --dir, --manifest and --base are required")
		return 2
	}
	mdata, err := os.ReadFile(*manifest)
	if err != nil {
		fmt.Fprintf(stderr, "physctl dist check-staged: read manifest: %v\n", err)
		return 2
	}
	bdata, err := os.ReadFile(*base)
	if err != nil {
		fmt.Fprintf(stderr, "physctl dist check-staged: read base list: %v\n", err)
		return 2
	}
	var ignoreList []string
	if *ignore != "" {
		ignoreList = strings.Split(*ignore, ",")
	}
	if err := distrib.VerifyStaged(*dir, mdata, bdata, ignoreList); err != nil {
		if *jsonOut {
			writeJSON(stdout, map[string]any{"ok": false, "violation": err.Error()})
		} else {
			fmt.Fprintf(stdout, "REJECTED: %v\n", err)
		}
		return 1
	}
	if *jsonOut {
		return writeJSON(stdout, map[string]any{"ok": true})
	}
	fmt.Fprintln(stdout, "OK: staged tree verified")
	return 0
}

func orNone(list []string) string {
	if len(list) == 0 {
		return "(none)"
	}
	return strings.Join(list, ", ")
}

func writeJSON(w io.Writer, v any) int {
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	if err := enc.Encode(v); err != nil {
		fmt.Fprintf(os.Stderr, "physctl: json output: %v\n", err)
		return 2
	}
	return 0
}

func cmdPaper(args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		fmt.Fprintln(stderr, "physctl paper: requires a subcommand (search, related, fetch, extract)")
		return 2
	}
	switch args[0] {
	case "search":
		return cmdPaperSearch(args[1:], stdout, stderr)
	case "related":
		return cmdPaperRelated(args[1:], stdout, stderr)
	case "fetch":
		return cmdPaperFetch(args[1:], stdout, stderr)
	case "extract":
		return cmdPaperExtract(args[1:], stdout, stderr)
	default:
		fmt.Fprintf(stderr, "physctl paper: unknown subcommand %q\n", args[0])
		return 2
	}
}

func cmdPaperSearch(args []string, stdout, stderr io.Writer) int {
	fs := flag.NewFlagSet("search", flag.ContinueOnError)
	fs.SetOutput(stderr)
	query := fs.String("query", "", "search query")
	limit := fs.Int("limit", 10, "max results (1-30)")
	jsonOut := fs.Bool("json", false, "machine-readable output")
	if err := fs.Parse(args); err != nil {
		return 2
	}
	if *query == "" {
		fmt.Fprintln(stderr, "physctl paper search: --query required")
		return 2
	}
	if *limit < 1 || *limit > 30 {
		fmt.Fprintln(stderr, "physctl paper search: --limit must be 1-30")
		return 2
	}
	results, err := paperSearch(*query, *limit)
	if err != nil {
		fmt.Fprintf(stderr, "physctl paper search: %v\n", err)
		return 2
	}
	if *jsonOut {
		enc := json.NewEncoder(stdout)
		enc.SetIndent("", "  ")
		enc.Encode(results)
		return 0
	}
	fmt.Fprintf(stdout, "Results (%d) for %q\n", len(results), *query)
	for _, r := range results {
		fmt.Fprintf(stdout, "  [%s] %s\n", r.SourceType, r.PaperID)
		fmt.Fprintf(stdout, "         %s\n", r.Title)
		if r.DOI != "" {
			fmt.Fprintf(stdout, "         DOI: %s\n", r.DOI)
		}
		if r.IsOA {
			fmt.Fprintf(stdout, "         open-access: yes\n")
		}
	}
	return 0
}

func cmdPaperRelated(args []string, stdout, stderr io.Writer) int {
	fs := flag.NewFlagSet("related", flag.ContinueOnError)
	fs.SetOutput(stderr)
	limit := fs.Int("limit", 10, "max results (1-30)")
	jsonOut := fs.Bool("json", false, "machine-readable output")
	if err := fs.Parse(args); err != nil {
		return 2
	}
	if fs.NArg() != 1 {
		fmt.Fprintln(stderr, "physctl paper related: exactly one paper ID required")
		return 2
	}
	if *limit < 1 || *limit > 30 {
		fmt.Fprintln(stderr, "physctl paper related: --limit must be 1-30")
		return 2
	}
	results, err := paperRelated(fs.Arg(0), "references", *limit)
	if err != nil {
		fmt.Fprintf(stderr, "physctl paper related: %v\n", err)
		return 2
	}
	if *jsonOut {
		enc := json.NewEncoder(stdout)
		enc.SetIndent("", "  ")
		enc.Encode(results)
		return 0
	}
	fmt.Fprintf(stdout, "Related papers (%d):\n", len(results))
	for _, r := range results {
		fmt.Fprintf(stdout, "  [%s] %s (%s)\n", r.SourceType, r.PaperID, r.Title)
	}
	return 0
}

func cmdPaperFetch(args []string, stdout, stderr io.Writer) int {
	fs := flag.NewFlagSet("fetch", flag.ContinueOnError)
	fs.SetOutput(stderr)
	version := fs.String("version", "v1", "version label")
	if err := fs.Parse(args); err != nil {
		return 2
	}
	if fs.NArg() != 1 {
		fmt.Fprintln(stderr, "physctl paper fetch: exactly one paper ID required")
		return 2
	}
	root, err := findRoot()
	if err != nil {
		fmt.Fprintf(stderr, "physctl paper fetch: %v\n", err)
		return 2
	}
	paperID := fs.Arg(0)
	// Use arXiv download URL for arXiv IDs; for other IDs, attempt search-based URL.
	downloadURL := fmt.Sprintf("https://arxiv.org/pdf/%s.pdf", paperID)
	sourceURL := fmt.Sprintf("https://arxiv.org/abs/%s", paperID)
	cached, err := paper.Fetch(root, paperID, *version, sourceURL, downloadURL)
	if err != nil {
		fmt.Fprintf(stderr, "physctl paper fetch: %v\n", err)
		if strings.Contains(err.Error(), "UNAVAILABLE") {
			return 2
		}
		return 1
	}
	fmt.Fprintf(stdout, "Fetched %s version %s (sha256:%s)\n", cached.PaperID, cached.Version, cached.SHA256PDF)
	return 0
}

func cmdPaperExtract(args []string, stdout, stderr io.Writer) int {
	fs := flag.NewFlagSet("extract", flag.ContinueOnError)
	fs.SetOutput(stderr)
	version := fs.String("version", "v1", "version label")
	if err := fs.Parse(args); err != nil {
		return 2
	}
	if fs.NArg() != 1 {
		fmt.Fprintln(stderr, "physctl paper extract: exactly one paper ID required")
		return 2
	}
	root, err := findRoot()
	if err != nil {
		fmt.Fprintf(stderr, "physctl paper extract: %v\n", err)
		return 2
	}
	err = paper.Extract(root, fs.Arg(0), *version)
	if err != nil {
		fmt.Fprintf(stderr, "physctl paper extract: %v\n", err)
		if strings.Contains(err.Error(), "UNAVAILABLE") {
			return 2
		}
		return 1
	}
	fmt.Fprintf(stdout, "Extracted %s version %s: %s\n", fs.Arg(0), *version, paper.FullTextPath(root, fs.Arg(0), *version))
	return 0
}

func paperSearch(query string, maxResults int) ([]paper.SearchResult, error) {
	return paper.Search(query, maxResults)
}

func paperRelated(paperID, direction string, maxResults int) ([]paper.SearchResult, error) {
	return paper.Related(paperID, direction, maxResults)
}
