// Package check runs ONLY the mathematical checks that the existing Phys
// infrastructure actually supports, expressed as a fixed allowlist of
// `go test` profiles over the developer-local Go workspace (AGENTS.md §10).
//
// A missing workspace, a missing Phys repository or an unknown profile
// produces an explicit UNAVAILABLE diagnostic — never a PASS. Nothing in
// this package establishes mathematical or physical correctness: a
// passing profile means the named tests passed, nothing more.
package check

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"
)

// Profile is one allowlisted mathematical check. Args are the arguments
// after "go" (always beginning with "test"). Callers cannot supply
// arbitrary commands: profiles are fixed in this file.
type Profile struct {
	Name    string
	Module  string // Go module path the profile must resolve to in go.work
	Args    []string
	Timeout time.Duration
}

// Profiles is the complete allowlist. Extend only with checks backed by
// the actual Phys infrastructure, and only after a capability audit
// (DERIVATION_SPEC_GUIDELINES.md §9).
var Profiles = []Profile{
	{
		Name:    "phys-artifact",
		Module:  "github.com/PithomLabs/phys-artifact",
		Args:    []string{"test", "-count=1", "./..."},
		Timeout: 3 * time.Minute,
	},
	{
		Name:    "phys-math-fast",
		Module:  "github.com/PithomLabs/phys-math",
		Args:    []string{"test", "-count=1", "-run", "TestRegistryCompleteness|TestNoPhysicalVocabulary|TestImportDiscipline", "."},
		Timeout: 3 * time.Minute,
	},
	{
		Name:    "phys-lib-equivariance",
		Module:  "github.com/PithomLabs/phys-lib",
		Args:    []string{"test", "-count=1", "./gr/equivariance/..."},
		Timeout: 3 * time.Minute,
	},
	{
		Name:    "phys-lib-ltb",
		Module:  "github.com/PithomLabs/phys-lib",
		Args:    []string{"test", "-count=1", "./gr/ltb/..."},
		Timeout: 10 * time.Minute,
	},
	{
		Name:    "phys-lib-bohm",
		Module:  "github.com/PithomLabs/phys-lib",
		Args:    []string{"test", "-count=1", "./bohm/..."},
		Timeout: 15 * time.Minute,
	},
}

// Statuses.
const (
	StatusPass        = "PASS"
	StatusFail        = "FAIL"
	StatusUnavailable = "UNAVAILABLE"
)

// Result is the outcome of one profile run.
type Result struct {
	Profile    string `json:"profile"`
	Status     string `json:"status"`
	Command    string `json:"command,omitempty"`
	ExitCode   int    `json:"exit_code,omitempty"`
	DurationMS int64  `json:"duration_ms"`
	Output     string `json:"output,omitempty"`     // tail of combined output
	Diagnostic string `json:"diagnostic,omitempty"` // why UNAVAILABLE / FAIL detail
}

// Names returns the allowlisted profile names, sorted.
func Names() []string {
	var out []string
	for _, p := range Profiles {
		out = append(out, p.Name)
	}
	sort.Strings(out)
	return out
}

// ProfileByName looks up a profile; ok is false when the name is not in
// the allowlist.
func ProfileByName(name string) (Profile, bool) {
	for _, p := range Profiles {
		if p.Name == name {
			return p, true
		}
	}
	return Profile{}, false
}

// ProfilesFromBody extracts `- profile: <name>` declarations from a
// derivation record body, in order of appearance.
func ProfilesFromBody(body string) []string {
	var out []string
	inChecks := false
	for _, ln := range strings.Split(body, "\n") {
		if strings.HasPrefix(ln, "## ") {
			inChecks = ln == "## Checks"
			continue
		}
		if !inChecks {
			continue
		}
		t := strings.TrimSpace(ln)
		if strings.HasPrefix(t, "- profile:") {
			name := strings.TrimSpace(strings.TrimPrefix(t, "- profile:"))
			if name != "" && !strings.Contains(name, "<") && !strings.Contains(name, "TODO") {
				out = append(out, name)
			}
		}
	}
	return out
}

// Run executes one allowlisted profile against the developer-local
// workspace under root.
func Run(root, name string) Result {
	prof, ok := ProfileByName(name)
	if !ok {
		return Result{
			Profile:    name,
			Status:     StatusUnavailable,
			Diagnostic: fmt.Sprintf("unknown profile %q; allowlisted profiles: %s", name, strings.Join(Names(), ", ")),
		}
	}

	goWork, err := findGoWork(root)
	if err != nil {
		return Result{Profile: name, Status: StatusUnavailable, Diagnostic: err.Error()}
	}

	moduleDir, err := resolveModule(goWork, prof.Module)
	if err != nil {
		return Result{Profile: name, Status: StatusUnavailable, Diagnostic: err.Error()}
	}

	if _, err := exec.LookPath("go"); err != nil {
		return Result{Profile: name, Status: StatusUnavailable, Diagnostic: "go toolchain not found on PATH"}
	}

	ctx, cancel := context.WithTimeout(context.Background(), prof.Timeout)
	defer cancel()

	cmd := exec.CommandContext(ctx, "go", prof.Args...)
	cmd.Dir = moduleDir
	cmd.Env = append(os.Environ(), "GOWORK="+goWork)
	var buf bytes.Buffer
	cmd.Stdout = &buf
	cmd.Stderr = &buf

	start := time.Now()
	runErr := cmd.Run()
	elapsed := time.Since(start)

	res := Result{
		Profile:    name,
		Command:    fmt.Sprintf("cd %s && GOWORK=%s go %s", moduleDir, goWork, strings.Join(prof.Args, " ")),
		DurationMS: elapsed.Milliseconds(),
		Output:     tail(buf.String(), 40),
	}

	if errors.Is(ctx.Err(), context.DeadlineExceeded) {
		res.Status = StatusFail
		res.Diagnostic = fmt.Sprintf("timeout after %s", prof.Timeout)
		return res
	}
	if runErr != nil {
		var ee *exec.ExitError
		if errors.As(runErr, &ee) {
			res.Status = StatusFail
			res.ExitCode = ee.ExitCode()
			return res
		}
		res.Status = StatusUnavailable
		res.Diagnostic = fmt.Sprintf("could not execute profile: %v", runErr)
		return res
	}
	res.Status = StatusPass
	res.ExitCode = 0
	return res
}

// findGoWork locates the developer-local, untracked go.work at the
// repository root.
func findGoWork(root string) (string, error) {
	p := filepath.Join(root, "go.work")
	fi, err := os.Stat(p)
	if err != nil || fi.IsDir() {
		return "", fmt.Errorf("go workspace not found at %s — create a local, untracked go.work as described in AGENTS.md §10 (physctl check reports UNAVAILABLE without it; it never reports PASS)", p)
	}
	abs, err := filepath.Abs(p)
	if err != nil {
		return "", fmt.Errorf("cannot resolve %s: %w", p, err)
	}
	return abs, nil
}

// useLineRe matches `use ./path` or a bare path inside a use ( ... ) block.
var useLineRe = regexp.MustCompile(`^\s*(?:use\s+)?(\S+)\s*$`)

// resolveModule maps a Go module path to its directory using the go.work
// `use` directives and each directory's go.mod.
func resolveModule(goWork, modulePath string) (string, error) {
	data, err := os.ReadFile(goWork)
	if err != nil {
		return "", fmt.Errorf("cannot read %s: %w", goWork, err)
	}

	base := filepath.Dir(goWork)
	var candidates []string
	inBlock := false
	for _, ln := range strings.Split(string(data), "\n") {
		trimmed := strings.TrimSpace(ln)
		if trimmed == "" || strings.HasPrefix(trimmed, "//") {
			continue
		}
		if strings.HasPrefix(trimmed, "use (") || trimmed == "use (" {
			inBlock = true
			continue
		}
		if inBlock && trimmed == ")" {
			inBlock = false
			continue
		}
		if strings.HasPrefix(trimmed, "use ") || inBlock {
			m := useLineRe.FindStringSubmatch(trimmed)
			if m != nil && m[1] != "use" {
				p := m[1]
				if !filepath.IsAbs(p) {
					p = filepath.Join(base, p)
				}
				candidates = append(candidates, p)
			}
		}
	}

	for _, dir := range candidates {
		gomod := filepath.Join(dir, "go.mod")
		gd, err := os.ReadFile(gomod)
		if err != nil {
			continue
		}
		for _, ln := range strings.Split(string(gd), "\n") {
			if strings.HasPrefix(ln, "module ") {
				if strings.TrimSpace(strings.TrimPrefix(ln, "module ")) == modulePath {
					return dir, nil
				}
			}
		}
	}
	return "", fmt.Errorf("module %s is not present in %s — add its directory to the local go.work use directives (AGENTS.md §10)", modulePath, goWork)
}

// tail returns the last n lines of s, trimmed.
func tail(s string, n int) string {
	lines := strings.Split(strings.TrimRight(s, "\n"), "\n")
	if len(lines) > n {
		lines = lines[len(lines)-n:]
	}
	out := strings.Join(lines, "\n")
	if len(out) > 8000 {
		out = out[len(out)-8000:]
	}
	return out
}
