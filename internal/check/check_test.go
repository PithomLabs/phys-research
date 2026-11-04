package check

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// workspace builds a repo root with a go.work pointing at a fake Phys
// module directory whose go.mod declares modulePath.
func workspace(t *testing.T, modulePath string) (root, moduleDir string) {
	t.Helper()
	root = t.TempDir()
	moduleDir = filepath.Join(root, "fake", "phys")
	if err := os.MkdirAll(moduleDir, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(moduleDir, "go.mod"), []byte("module "+modulePath+"\n\ngo 1.25.7\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	work := "go 1.25.7\n\nuse (\n\t./fake/phys\n)\n"
	if err := os.WriteFile(filepath.Join(root, "go.work"), []byte(work), 0o644); err != nil {
		t.Fatal(err)
	}
	return root, moduleDir
}

// shimGo installs a fake `go` executable at the front of PATH.
// The script appends "$GOWORK" and its arguments to record, then exits
// with the code given by exitCode.
func shimGo(t *testing.T, exitCode int) (record string, restore func()) {
	t.Helper()
	bin := t.TempDir()
	record = filepath.Join(bin, "record.txt")
	script := "#!/bin/sh\n" +
		"echo \"GOWORK=$GOWORK\" >> " + record + "\n" +
		"echo \"ARGS=$*\" >> " + record + "\n" +
		"echo \"PWD=$(pwd)\" >> " + record + "\n" +
		"echo \"fake go ran\"\n" +
		"exit " + itoa(exitCode) + "\n"
	goPath := filepath.Join(bin, "go")
	if err := os.WriteFile(goPath, []byte(script), 0o755); err != nil {
		t.Fatal(err)
	}
	oldPath := os.Getenv("PATH")
	if err := os.Setenv("PATH", bin+string(os.PathListSeparator)+oldPath); err != nil {
		t.Fatal(err)
	}
	return record, func() { _ = os.Setenv("PATH", oldPath) }
}

func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	var b []byte
	for n > 0 {
		b = append([]byte{byte('0' + n%10)}, b...)
		n /= 10
	}
	return string(b)
}

func TestProfileAllowlist(t *testing.T) {
	names := Names()
	if len(names) == 0 {
		t.Fatal("empty allowlist")
	}
	// Names must be sorted and unique.
	for i := 1; i < len(names); i++ {
		if names[i-1] >= names[i] {
			t.Errorf("names not sorted/unique: %v", names)
		}
	}
	for _, want := range []string{"phys-artifact", "phys-math-fast", "phys-lib-equivariance"} {
		found := false
		for _, n := range names {
			if n == want {
				found = true
			}
		}
		if !found {
			t.Errorf("allowlist missing %q: %v", want, names)
		}
	}
	if _, ok := ProfileByName("rm -rf /"); ok {
		t.Error("arbitrary command accepted as a profile")
	}
}

func TestUnknownProfileUnavailable(t *testing.T) {
	root, _ := workspace(t, "github.com/PithomLabs/phys-artifact")
	res := Run(root, "not-a-profile")
	if res.Status != StatusUnavailable {
		t.Errorf("status = %s, want UNAVAILABLE", res.Status)
	}
	if !strings.Contains(res.Diagnostic, "allowlisted") {
		t.Errorf("diagnostic must list the allowlist: %q", res.Diagnostic)
	}
}

func TestMissingGoWorkUnavailable(t *testing.T) {
	root := t.TempDir() // no go.work
	res := Run(root, "phys-artifact")
	if res.Status != StatusUnavailable {
		t.Errorf("status = %s, want UNAVAILABLE", res.Status)
	}
	if !strings.Contains(res.Diagnostic, "go.work") {
		t.Errorf("diagnostic must mention go.work: %q", res.Diagnostic)
	}
	if !strings.Contains(res.Diagnostic, "never reports PASS") {
		t.Errorf("diagnostic must state it never reports PASS: %q", res.Diagnostic)
	}
}

func TestModuleNotInWorkspaceUnavailable(t *testing.T) {
	root := t.TempDir()
	// go.work exists but uses a different module.
	work := "go 1.25.7\n\nuse (\n\t./somewhere\n)\n"
	if err := os.WriteFile(filepath.Join(root, "go.work"), []byte(work), 0o644); err != nil {
		t.Fatal(err)
	}
	res := Run(root, "phys-artifact")
	if res.Status != StatusUnavailable {
		t.Errorf("status = %s, want UNAVAILABLE", res.Status)
	}
	if !strings.Contains(res.Diagnostic, "not present") {
		t.Errorf("diagnostic = %q", res.Diagnostic)
	}
}

func TestPassWithShim(t *testing.T) {
	root, moduleDir := workspace(t, "github.com/PithomLabs/phys-artifact")
	record, restore := shimGo(t, 0)
	defer restore()

	res := Run(root, "phys-artifact")
	if res.Status != StatusPass {
		t.Fatalf("status = %s (diagnostic %q), want PASS", res.Status, res.Diagnostic)
	}
	if res.ExitCode != 0 {
		t.Errorf("exit = %d, want 0", res.ExitCode)
	}

	data, err := os.ReadFile(record)
	if err != nil {
		t.Fatal(err)
	}
	text := string(data)
	// The profile must run in the module directory with GOWORK set.
	if !strings.Contains(text, "PWD="+moduleDir) {
		t.Errorf("wrong working directory:\n%s", text)
	}
	absWork, _ := filepath.Abs(filepath.Join(root, "go.work"))
	if !strings.Contains(text, "GOWORK="+absWork) {
		t.Errorf("GOWORK not exported to the child:\n%s", text)
	}
	// Arguments must match the fixed profile, starting with test.
	if !strings.Contains(text, "ARGS=test -count=1 ./...") {
		t.Errorf("wrong arguments:\n%s", text)
	}
}

func TestFailWithShim(t *testing.T) {
	root, _ := workspace(t, "github.com/PithomLabs/phys-artifact")
	_, restore := shimGo(t, 1)
	defer restore()

	res := Run(root, "phys-artifact")
	if res.Status != StatusFail {
		t.Fatalf("status = %s (diagnostic %q), want FAIL", res.Status, res.Diagnostic)
	}
	if res.ExitCode != 1 {
		t.Errorf("exit = %d, want 1", res.ExitCode)
	}
}

func TestPassNeverWithoutShimAuthority(t *testing.T) {
	// A profile result can only be PASS if the go command actually ran
	// and exited 0: with no go on PATH, it must be UNAVAILABLE.
	root, _ := workspace(t, "github.com/PithomLabs/phys-artifact")
	oldPath := os.Getenv("PATH")
	if err := os.Setenv("PATH", t.TempDir()); err != nil { // empty dir: no go
		t.Fatal(err)
	}
	defer func() { _ = os.Setenv("PATH", oldPath) }()

	res := Run(root, "phys-artifact")
	if res.Status != StatusUnavailable {
		t.Errorf("status = %s, want UNAVAILABLE when go is absent", res.Status)
	}
}

func TestProfilesFromBody(t *testing.T) {
	body := `# D

## Checks

- profile: phys-artifact
- profile: phys-math-fast
- profile: <allowlisted physctl check profile name>
- profile: TODO later

## Something else

- profile: phys-lib-ltb
`
	got := ProfilesFromBody(body)
	want := []string{"phys-artifact", "phys-math-fast"}
	if len(got) != len(want) {
		t.Fatalf("got %v, want %v", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("got %v, want %v", got, want)
		}
	}
}

func TestResolveModuleParsesUseBlock(t *testing.T) {
	root, moduleDir := workspace(t, "github.com/PithomLabs/phys-math")
	goWork := filepath.Join(root, "go.work")
	got, err := resolveModule(goWork, "github.com/PithomLabs/phys-math")
	if err != nil {
		t.Fatalf("resolveModule: %v", err)
	}
	if got != moduleDir {
		t.Errorf("got %q, want %q", got, moduleDir)
	}
	if _, err := resolveModule(goWork, "github.com/other"); err == nil {
		t.Error("unknown module resolved without error")
	}
}

func TestResolveModuleSingleUseLine(t *testing.T) {
	root := t.TempDir()
	dir := filepath.Join(root, "m")
	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "go.mod"), []byte("module example.com/m\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	work := "go 1.25.7\n\nuse ./m\n"
	workPath := filepath.Join(root, "go.work")
	if err := os.WriteFile(workPath, []byte(work), 0o644); err != nil {
		t.Fatal(err)
	}
	got, err := resolveModule(workPath, "example.com/m")
	if err != nil {
		t.Fatalf("resolveModule: %v", err)
	}
	if got != dir {
		t.Errorf("got %q, want %q", got, dir)
	}
}
