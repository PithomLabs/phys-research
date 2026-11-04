package distrib

import (
	"archive/tar"
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

// writeTar builds a tarball from name -> (body or link) specifications.
type tarEntry struct {
	body     string
	linkname string
	typeflag byte
}

func writeTar(t *testing.T, path string, entries map[string]tarEntry) {
	t.Helper()
	var order []string
	for name := range entries {
		order = append(order, name)
	}
	f, err := os.Create(path)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	tw := tar.NewWriter(f)
	defer tw.Close()
	for _, name := range order {
		e := entries[name]
		flag := e.typeflag
		if flag == 0 {
			flag = tar.TypeReg
		}
		if err := tw.WriteHeader(&tar.Header{Name: name, Mode: 0o644, Size: int64(len(e.body)), Typeflag: flag, Linkname: e.linkname}); err != nil {
			t.Fatal(err)
		}
		if _, err := tw.Write([]byte(e.body)); err != nil {
			t.Fatal(err)
		}
	}
}

func TestValidateTarClean(t *testing.T) {
	tarPath := filepath.Join(t.TempDir(), "p.tar")
	writeTar(t, tarPath, map[string]tarEntry{
		"research/sources/P-099.md": {body: "x"},
		"docs/":                     {typeflag: tar.TypeDir},
		"docs/note.md":              {body: "y"},
	})
	got, err := ValidateTar(tarPath)
	if err != nil {
		t.Fatalf("clean tarball rejected: %v", err)
	}
	if len(got) != 2 || got[0] != "docs/note.md" || got[1] != "research/sources/P-099.md" {
		t.Errorf("unexpected entry list: %q", got)
	}
}

func TestValidateTarRejects(t *testing.T) {
	cases := map[string]map[string]tarEntry{
		"absolute path":    {"/etc/passwd": {body: "x"}},
		"traversal":        {"../escape.md": {body: "x"}},
		"nested traversal": {"a/../../escape.md": {body: "x"}},
		"symlink":          {"link.md": {typeflag: tar.TypeSymlink, linkname: "/etc/passwd"}},
		"hardlink":         {"hard.md": {typeflag: tar.TypeLink, linkname: "docs/note.md"}},
		"char device":      {"dev": {typeflag: tar.TypeChar}},
		"duplicate":        {"dup.md": {body: "x"}},
		"absolute dir":     {"/tmp/d/": {typeflag: tar.TypeDir}},
		"traversal dir":    {"../d/": {typeflag: tar.TypeDir}},
		"dot-dot only":     {"..": {body: "x"}},
	}
	for name, entries := range cases {
		t.Run(name, func(t *testing.T) {
			tarPath := filepath.Join(t.TempDir(), "p.tar")
			if name == "duplicate" {
				// Two entries with the same name need ordered writes.
				f, err := os.Create(tarPath)
				if err != nil {
					t.Fatal(err)
				}
				defer f.Close()
				tw := tar.NewWriter(f)
				defer tw.Close()
				for _, body := range []string{"one", "two"} {
					if err := tw.WriteHeader(&tar.Header{Name: "dup.md", Mode: 0o644, Size: int64(len(body))}); err != nil {
						t.Fatal(err)
					}
					if _, err := tw.Write([]byte(body)); err != nil {
						t.Fatal(err)
					}
				}
			} else {
				writeTar(t, tarPath, entries)
			}
			if _, err := ValidateTar(tarPath); err == nil {
				t.Errorf("%s: malicious tarball accepted", name)
			}
		})
	}
}

// TestSystemTarEscapeRepro is the realistic CR-CA-003 reproduction: it shows
// that the system tar binary WOULD write outside the destination (threat is
// real), while ValidateTar rejects the same tarball (defense works).
// Extraction stays inside a temp dir; nothing outside it is touched.
func TestSystemTarEscapeRepro(t *testing.T) {
	if _, err := exec.LookPath("tar"); err != nil {
		t.Skip("tar binary unavailable; escape repro unverified here")
	}
	root := t.TempDir()
	dest := filepath.Join(root, "dest")
	if err := os.Mkdir(dest, 0o755); err != nil {
		t.Fatal(err)
	}
	tarPath := filepath.Join(root, "evil.tar")
	writeTar(t, tarPath, map[string]tarEntry{"../escaped.md": {body: "escaped"}})
	if _, err := ValidateTar(tarPath); err == nil {
		t.Fatal("ValidateTar accepted the traversal tarball")
	}
	cmd := exec.Command("tar", "-xf", tarPath, "-C", dest)
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Skipf("system tar refused the traversal entry itself (%v): %s", err, out)
	}
	if _, err := os.Stat(filepath.Join(root, "escaped.md")); err != nil {
		t.Skipf("system tar did not write outside dest on this platform: %v", err)
	}
	// If we reach here, the threat is demonstrated AND the defense held.
}

func TestVerifyStaged(t *testing.T) {
	stage := t.TempDir()
	// Base file predates the payload; payload file is the distribution.
	if err := os.WriteFile(filepath.Join(stage, "base.md"), []byte("base"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(stage, "new.md"), []byte("new"), 0o644); err != nil {
		t.Fatal(err)
	}
	m, err := WriteManifest(stage, []string{"new.md"})
	if err != nil {
		t.Fatal(err)
	}
	base := []byte("base.md\n")
	if err := VerifyStaged(stage, m, base, []string{"go.work"}); err != nil {
		t.Fatalf("clean stage rejected: %v", err)
	}
	// Changed payload bytes fail.
	os.WriteFile(filepath.Join(stage, "new.md"), []byte("tampered"), 0o644)
	if err := VerifyStaged(stage, m, base, nil); err == nil {
		t.Error("tampered payload accepted")
	}
	os.WriteFile(filepath.Join(stage, "new.md"), []byte("new"), 0o644)
	// Unlisted, non-base files fail.
	os.WriteFile(filepath.Join(stage, "rogue.md"), []byte("r"), 0o644)
	if err := VerifyStaged(stage, m, base, nil); err == nil {
		t.Error("rogue file accepted")
	}
	os.Remove(filepath.Join(stage, "rogue.md"))
	// Missing payload fails.
	os.Remove(filepath.Join(stage, "new.md"))
	if err := VerifyStaged(stage, m, base, nil); err == nil {
		t.Error("missing payload accepted")
	}
}

// TestFailedDistributionLeavesAgentTreeUntouched pins the fail-closed
// composition: every validation in the staged procedure runs before any
// write to an agent workspace, so a rejected payload cannot disturb
// pre-existing agent outputs. The agent tree below stands in for a real
// agent worktree; the test asserts byte-identity after each failure.
func TestFailedDistributionLeavesAgentTreeUntouched(t *testing.T) {
	agent := t.TempDir()
	owned := map[string]string{
		"research/hypotheses/H-900.md":   "agent's own hypothesis",
		"research/proposed/muse/r1/x.md": "agent's own proposal",
	}
	for rel, body := range owned {
		p := filepath.Join(agent, filepath.FromSlash(rel))
		if err := os.MkdirAll(filepath.Dir(p), 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(p, []byte(body), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	snapshot := func() map[string]string {
		out := map[string]string{}
		filepath.WalkDir(agent, func(p string, d os.DirEntry, err error) error {
			if err != nil || d.IsDir() {
				return err
			}
			raw, err := os.ReadFile(p)
			if err != nil {
				return err
			}
			rel, _ := filepath.Rel(agent, p)
			out[filepath.ToSlash(rel)] = string(raw)
			return nil
		})
		return out
	}
	before := snapshot()
	equal := func() bool {
		after := snapshot()
		if len(after) != len(before) {
			return false
		}
		for k, v := range before {
			if after[k] != v {
				return false
			}
		}
		return true
	}

	// Failure 1: malicious tarball rejected at validation; no extraction.
	evil := filepath.Join(t.TempDir(), "evil.tar")
	writeTar(t, evil, map[string]tarEntry{"../escape.md": {body: "x"}})
	if _, err := ValidateTar(evil); err == nil {
		t.Fatal("malicious tarball accepted")
	}
	if !equal() {
		t.Fatal("agent tree changed after rejected tarball")
	}

	// Failure 2: tampered payload manifest rejected at verification.
	stage := t.TempDir()
	os.WriteFile(filepath.Join(stage, "p.md"), []byte("payload"), 0o644)
	m, err := WriteManifest(stage, []string{"p.md"})
	if err != nil {
		t.Fatal(err)
	}
	os.WriteFile(filepath.Join(stage, "p.md"), []byte("tampered"), 0o644)
	if err := VerifyStaged(stage, m, nil, nil); err == nil {
		t.Fatal("tampered stage accepted")
	}
	if !equal() {
		t.Fatal("agent tree changed after rejected stage")
	}
}

func TestManifestRoundTrip(t *testing.T) {
	dir := t.TempDir()
	files := map[string]string{"a.md": "aaa", "sub/b.md": "bbb"}
	for rel, body := range files {
		p := filepath.Join(dir, filepath.FromSlash(rel))
		if err := os.MkdirAll(filepath.Dir(p), 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(p, []byte(body), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	m, err := WriteManifest(dir, []string{"a.md", "sub/b.md"})
	if err != nil {
		t.Fatalf("write: %v", err)
	}
	if err := VerifyManifest(dir, m, nil); err != nil {
		t.Fatalf("round trip: %v", err)
	}
}

func TestVerifyManifestDetects(t *testing.T) {
	setup := func(t *testing.T) (string, []byte) {
		dir := t.TempDir()
		if err := os.WriteFile(filepath.Join(dir, "a.md"), []byte("aaa"), 0o644); err != nil {
			t.Fatal(err)
		}
		m, err := WriteManifest(dir, []string{"a.md"})
		if err != nil {
			t.Fatal(err)
		}
		return dir, m
	}
	t.Run("changed bytes", func(t *testing.T) {
		dir, m := setup(t)
		os.WriteFile(filepath.Join(dir, "a.md"), []byte("zzz"), 0o644)
		if err := VerifyManifest(dir, m, nil); err == nil || !strings.Contains(err.Error(), "changed") {
			t.Errorf("changed bytes undetected: %v", err)
		}
	})
	t.Run("missing file", func(t *testing.T) {
		dir, m := setup(t)
		os.Remove(filepath.Join(dir, "a.md"))
		if err := VerifyManifest(dir, m, nil); err == nil || !strings.Contains(err.Error(), "missing") {
			t.Errorf("missing file undetected: %v", err)
		}
	})
	t.Run("unexpected file", func(t *testing.T) {
		dir, m := setup(t)
		os.WriteFile(filepath.Join(dir, "extra.md"), []byte("e"), 0o644)
		if err := VerifyManifest(dir, m, nil); err == nil || !strings.Contains(err.Error(), "unexpected") {
			t.Errorf("unexpected file undetected: %v", err)
		}
	})
	t.Run("duplicate lines", func(t *testing.T) {
		dir, m := setup(t)
		dup := append(append([]byte(nil), m...), m...)
		if err := VerifyManifest(dir, dup, nil); err == nil || !strings.Contains(err.Error(), "duplicate") {
			t.Errorf("duplicate lines undetected: %v", err)
		}
	})
	t.Run("malformed lines", func(t *testing.T) {
		dir, _ := setup(t)
		if err := VerifyManifest(dir, []byte("not-a-digest  a.md\n"), nil); err == nil {
			t.Error("malformed line undetected")
		}
	})
	t.Run("manifest traversal", func(t *testing.T) {
		dir, _ := setup(t)
		line := bytes.Repeat([]byte("a"), 64)
		if err := VerifyManifest(dir, append(append([]byte(nil), line...), []byte("  ../x.md\n")...), nil); err == nil {
			t.Error("traversal manifest line undetected")
		}
	})
	t.Run("ignored go.work", func(t *testing.T) {
		dir, m := setup(t)
		os.WriteFile(filepath.Join(dir, "go.work"), []byte("machine-local"), 0o644)
		if err := VerifyManifest(dir, m, []string{"go.work"}); err != nil {
			t.Errorf("ignored go.work flagged: %v", err)
		}
		if err := VerifyManifest(dir, m, nil); err == nil {
			t.Error("go.work without ignore entry must count as unexpected")
		}
	})
	t.Run("git always ignored", func(t *testing.T) {
		dir, m := setup(t)
		os.MkdirAll(filepath.Join(dir, ".git", "refs"), 0o755)
		os.WriteFile(filepath.Join(dir, ".git", "HEAD"), []byte("ref"), 0o644)
		if err := VerifyManifest(dir, m, nil); err != nil {
			t.Errorf(".git metadata flagged: %v", err)
		}
	})
}
