package record

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

// TestWorktreeConflictingEditsPreserved documents the git layer of the
// isolation story: two agents editing the same tracked file in separate
// worktrees cannot silently overwrite each other — the merge conflicts and
// both sides stay present in the conflict markers. Skipped when git is
// unavailable; it needs a real git binary, not a mock.
func TestWorktreeConflictingEditsPreserved(t *testing.T) {
	git, err := exec.LookPath("git")
	if err != nil {
		t.Skip("git binary unavailable; worktree isolation unverified here")
	}
	run := func(dir string, args ...string) (string, error) {
		cmd := exec.Command(git, args...)
		cmd.Dir = dir
		cmd.Env = append(os.Environ(),
			"GIT_AUTHOR_NAME=test", "GIT_AUTHOR_EMAIL=test@test",
			"GIT_COMMITTER_NAME=test", "GIT_COMMITTER_EMAIL=test@test",
		)
		out, err := cmd.CombinedOutput()
		return string(out), err
	}

	root := t.TempDir()
	for _, args := range [][]string{
		{"init", "-q", "-b", "main"},
		{"config", "commit.gpgsign", "false"},
	} {
		if out, err := run(root, args...); err != nil {
			t.Fatalf("git %v: %v\n%s", args, err, out)
		}
	}
	record := filepath.Join(root, "research", "hypotheses", "H-001.md")
	if err := os.MkdirAll(filepath.Dir(record), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(record, []byte("baseline\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	for _, args := range [][]string{
		{"add", "."},
		{"commit", "-qm", "baseline"},
		{"branch", "agent-a"},
		{"branch", "agent-b"},
		{"worktree", "add", "-q", "wt-a", "agent-a"},
		{"worktree", "add", "-q", "wt-b", "agent-b"},
	} {
		if out, err := run(root, args...); err != nil {
			t.Fatalf("git %v: %v\n%s", args, err, out)
		}
	}

	// Each agent edits the same line differently and commits on its branch.
	edits := map[string]string{"wt-a": "agent A revision\n", "wt-b": "agent B revision\n"}
	for wt, content := range edits {
		p := filepath.Join(root, wt, "research", "hypotheses", "H-001.md")
		if err := os.WriteFile(p, []byte(content), 0o644); err != nil {
			t.Fatal(err)
		}
		for _, args := range [][]string{{"add", "."}, {"commit", "-qm", wt}} {
			if out, err := run(filepath.Join(root, wt), args...); err != nil {
				t.Fatalf("git %v in %s: %v\n%s", args, wt, err, out)
			}
		}
	}

	// Merging the second branch into the first must conflict, preserving
	// both revisions in the markers instead of picking a winner.
	if out, err := run(filepath.Join(root, "wt-a"), "merge", "-q", "agent-b"); err == nil {
		t.Fatalf("expected merge conflict, merge succeeded:\n%s", out)
	}
	merged, err := os.ReadFile(filepath.Join(root, "wt-a", "research", "hypotheses", "H-001.md"))
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{"<<<<<<<", "=======", ">>>>>>>", "agent A revision", "agent B revision"} {
		if !strings.Contains(string(merged), want) {
			t.Errorf("conflicted file missing %q:\n%s", want, merged)
		}
	}
	if out, err := run(filepath.Join(root, "wt-a"), "merge", "--abort"); err != nil {
		t.Fatalf("merge --abort: %v\n%s", err, out)
	}
}
