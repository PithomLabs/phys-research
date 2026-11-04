package paper

import (
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"testing"
)

// testCacheDir stages an empty cache dir in a temp root and returns it.
func testCacheDir(t *testing.T) *cacheDir {
	t.Helper()
	dir := t.TempDir()
	rt, err := os.OpenRoot(dir)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { rt.Close() })
	return &cacheDir{rt: rt, abs: dir}
}

func placeMeta() *CachedPaper {
	return &CachedPaper{
		PaperID:     "9999.99999",
		Version:     "v1",
		SourceURL:   "https://arxiv.org/abs/9999.99999",
		DownloadURL: "https://arxiv.org/pdf/9999.99999.pdf",
		RetrievedAt: "2026-11-05T00:00:00Z",
	}
}

func TestPlaceDownloadFresh(t *testing.T) {
	dir := testCacheDir(t)
	data := []byte("%PDF-1.4 fresh bytes")

	got, err := placeDownload("9999.99999", dir, data, placeMeta())
	if err != nil {
		t.Fatalf("fresh placement: %v", err)
	}
	if got.SHA256PDF != digestString(data) {
		t.Errorf("manifest sha %s, want digest of placed bytes", got.SHA256PDF)
	}
	onDisk, err := os.ReadFile(filepath.Join(dir.abs, "source.pdf"))
	if err != nil || string(onDisk) != string(data) {
		t.Fatalf("placed PDF mismatch: %v", err)
	}
}

func TestPlaceDownloadIdempotent(t *testing.T) {
	dir := testCacheDir(t)
	data := []byte("%PDF-1.4 same bytes")
	first, err := placeDownload("9999.99999", dir, data, placeMeta())
	if err != nil {
		t.Fatalf("first placement: %v", err)
	}
	before, err := os.ReadFile(filepath.Join(dir.abs, "manifest.json"))
	if err != nil {
		t.Fatal(err)
	}
	second, err := placeDownload("9999.99999", dir, append([]byte(nil), data...), placeMeta())
	if err != nil {
		t.Fatalf("idempotent re-place must succeed: %v", err)
	}
	after, err := os.ReadFile(filepath.Join(dir.abs, "manifest.json"))
	if err != nil {
		t.Fatal(err)
	}
	if string(before) != string(after) {
		t.Error("idempotent re-fetch rewrote the manifest; it must leave the entry untouched")
	}
	if second.SHA256PDF != first.SHA256PDF {
		t.Error("idempotent re-fetch returned a different digest")
	}
}

func TestPlaceDownloadCollision(t *testing.T) {
	dir := testCacheDir(t)
	original := []byte("%PDF-1.4 original bytes")
	if _, err := placeDownload("9999.99999", dir, original, placeMeta()); err != nil {
		t.Fatalf("first placement: %v", err)
	}
	_, err := placeDownload("9999.99999", dir, []byte("%PDF-1.4 DIFFERENT bytes"), placeMeta())
	if err == nil {
		t.Fatal("differing bytes must cause a collision error, never silent replacement")
	}
	if !strings.Contains(err.Error(), "collision") {
		t.Errorf("error must say collision, got: %v", err)
	}
	onDisk, rerr := os.ReadFile(filepath.Join(dir.abs, "source.pdf"))
	if rerr != nil || string(onDisk) != string(original) {
		t.Fatal("cached PDF was modified despite the collision")
	}
	mdata, rerr := os.ReadFile(filepath.Join(dir.abs, "manifest.json"))
	if rerr != nil {
		t.Fatal(rerr)
	}
	var m CachedPaper
	if json.Unmarshal(mdata, &m) != nil || m.SHA256PDF != digestString(original) {
		t.Fatal("manifest no longer matches the preserved bytes")
	}
}

func TestPlaceDownloadConcurrentSameBytes(t *testing.T) {
	dir := testCacheDir(t)
	data := []byte("%PDF-1.4 raced bytes")
	const n = 8
	var wg sync.WaitGroup
	errs := make([]error, n)
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			_, errs[i] = placeDownload("9999.99999", dir, data, placeMeta())
		}(i)
	}
	wg.Wait()
	for i, err := range errs {
		if err != nil {
			t.Fatalf("racer %d: %v", i, err)
		}
	}
	onDisk, err := os.ReadFile(filepath.Join(dir.abs, "source.pdf"))
	if err != nil || string(onDisk) != string(data) {
		t.Fatal("raced placement left inconsistent bytes")
	}
}

func TestPlaceDownloadConcurrentDifferingBytes(t *testing.T) {
	dir := testCacheDir(t)
	const n = 4
	var wg sync.WaitGroup
	errs := make([]error, n)
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			data := []byte{0x25, 0x50, 0x44, 0x46, byte('0' + i)} // distinct per racer
			_, errs[i] = placeDownload("9999.99999", dir, data, placeMeta())
		}(i)
	}
	wg.Wait()
	wins := 0
	for _, err := range errs {
		if err == nil {
			wins++
		} else if !strings.Contains(err.Error(), "collision") {
			t.Fatalf("loser error must say collision, got: %v", err)
		}
	}
	if wins != 1 {
		t.Fatalf("exactly one racer must win, got %d", wins)
	}
	onDisk, err := os.ReadFile(filepath.Join(dir.abs, "source.pdf"))
	if err != nil {
		t.Fatal(err)
	}
	mdata, err := os.ReadFile(filepath.Join(dir.abs, "manifest.json"))
	if err != nil {
		t.Fatal(err)
	}
	var m CachedPaper
	if json.Unmarshal(mdata, &m) != nil || m.SHA256PDF != digestString(onDisk) {
		t.Fatal("winner's manifest does not match the placed bytes")
	}
}

func TestExclWriteRefusesExisting(t *testing.T) {
	dir := testCacheDir(t)
	if err := exclWrite(dir.rt, "f.txt", []byte("first")); err != nil {
		t.Fatalf("fresh exclWrite: %v", err)
	}
	if err := exclWrite(dir.rt, "f.txt", []byte("second")); !os.IsExist(err) {
		t.Fatalf("second exclWrite must fail with EEXIST, got %v", err)
	}
	got, err := os.ReadFile(filepath.Join(dir.abs, "f.txt"))
	if err != nil || string(got) != "first" {
		t.Fatal("exclWrite modified an existing file")
	}
}

func TestPlaceDownloadRefusesSymlinkPDF(t *testing.T) {
	for _, target := range []string{"identical", "different", "outside"} {
		t.Run(target, func(t *testing.T) {
			dir := testCacheDir(t)
			var targetBytes []byte
			var linkTarget string
			switch target {
			case "identical":
				targetBytes = []byte("%PDF-1.4 victim bytes")
				linkTarget = filepath.Join(dir.abs, "victim.pdf")
			case "different":
				targetBytes = []byte("%PDF-1.4 other bytes")
				linkTarget = filepath.Join(dir.abs, "victim.pdf")
			case "outside":
				outside := t.TempDir()
				targetBytes = []byte("%PDF-1.4 outside bytes")
				linkTarget = filepath.Join(outside, "outside.pdf")
			}
			if err := os.WriteFile(linkTarget, targetBytes, 0o644); err != nil {
				t.Fatal(err)
			}
			if err := os.Symlink(linkTarget, filepath.Join(dir.abs, "source.pdf")); err != nil {
				t.Fatal(err)
			}
			// Even byte-identical targets must be refused: a link is never
			// accepted as an equivalent cached file.
			download := []byte("%PDF-1.4 victim bytes")
			_, err := placeDownload("9999.99999", dir, download, placeMeta())
			if err == nil {
				t.Fatalf("symlinked PDF (%s target) must be refused", target)
			}
			if !strings.Contains(err.Error(), "symlink") {
				t.Errorf("error must name the symlink refusal, got: %v", err)
			}
			if _, serr := os.Lstat(filepath.Join(dir.abs, "manifest.json")); !os.IsNotExist(serr) {
				t.Error("refused placement must not write a manifest")
			}
			got, rerr := os.ReadFile(linkTarget)
			if rerr != nil || string(got) != string(targetBytes) {
				t.Error("symlink target disturbed by the refused placement")
			}
		})
	}
}

func TestPlaceDownloadRefusesSymlinkManifest(t *testing.T) {
	dir := testCacheDir(t)
	victim := filepath.Join(dir.abs, "victim.json")
	if err := os.WriteFile(victim, []byte(`{"paper_id":"x"}`), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.Symlink(victim, filepath.Join(dir.abs, "manifest.json")); err != nil {
		t.Fatal(err)
	}
	_, err := placeDownload("9999.99999", dir, []byte("%PDF-1.4 data"), placeMeta())
	if err == nil {
		t.Fatal("symlinked manifest must be refused")
	}
	if !strings.Contains(err.Error(), "symlink") {
		t.Errorf("error must name the symlink refusal, got: %v", err)
	}
	// The PDF placed before the manifest refusal must be cleaned up.
	if _, serr := os.Lstat(filepath.Join(dir.abs, "source.pdf")); !os.IsNotExist(serr) {
		t.Error("refused placement left a PDF behind; only consistent entries may remain")
	}
}

func TestPlaceDownloadRefusesNonRegularPDF(t *testing.T) {
	dir := testCacheDir(t)
	if err := os.Mkdir(filepath.Join(dir.abs, "source.pdf"), 0o755); err != nil {
		t.Fatal(err)
	}
	if _, err := placeDownload("9999.99999", dir, []byte("%PDF-1.4 data"), placeMeta()); err == nil {
		t.Fatal("directory at the PDF destination must be refused")
	}
}

func TestOpenCacheDirRefusesSymlinkComponent(t *testing.T) {
	root := t.TempDir()
	cache := filepath.Join(root, "paper-cache", "9999.99999")
	if err := os.MkdirAll(cache, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.Symlink(cache, filepath.Join(root, "paper-cache", "evil")); err != nil {
		t.Fatal(err)
	}
	if _, err := openCacheDir(root, "evil", "v1"); err == nil {
		t.Fatal("symlinked cache component must be refused")
	} else if !strings.Contains(err.Error(), "symlink") {
		t.Errorf("error must name the symlink refusal, got: %v", err)
	}
}

func TestFetchRejectsBadVersionBeforeNetwork(t *testing.T) {
	for _, bad := range []string{"../x", "a/b", "..", ".", "", "v 1"} {
		if _, err := Fetch(t.TempDir(), "9999.99999", bad, "https://invalid.invalid/abs", "https://invalid.invalid/pdf"); err == nil {
			t.Errorf("version %q must be rejected", bad)
		} else if !strings.Contains(err.Error(), "version label") {
			t.Errorf("version %q: error must name the version label, got: %v", bad, err)
		}
	}
}

func TestPlaceDownloadFailureCreatesNothing(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("permission-based failure needs POSIX semantics")
	}
	dir := t.TempDir()
	sub := filepath.Join(dir, "ro")
	if err := os.Mkdir(sub, 0o555); err != nil {
		t.Fatal(err)
	}
	rt, err := os.OpenRoot(sub)
	if err != nil {
		t.Fatal(err)
	}
	defer rt.Close()
	d := &cacheDir{rt: rt, abs: sub}
	if _, err := placeDownload("9999.99999", d, []byte("%PDF-1.4 data"), placeMeta()); err == nil {
		t.Fatal("placement into a read-only directory must fail")
	}
	entries, err := os.ReadDir(sub)
	if err != nil {
		t.Fatal(err)
	}
	if len(entries) != 0 {
		t.Errorf("failed placement left %d entries behind", len(entries))
	}
}

func TestExtractRefusesSymlinkPDF(t *testing.T) {
	root := t.TempDir()
	dir := filepath.Join(root, "paper-cache", "9999.99999", "v1")
	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Fatal(err)
	}
	target := filepath.Join(dir, "real.pdf")
	if err := os.WriteFile(target, []byte("%PDF-1.4 x"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.Symlink(target, filepath.Join(dir, "source.pdf")); err != nil {
		t.Fatal(err)
	}
	if err := Extract(root, "9999.99999", "v1"); err == nil {
		t.Fatal("symlinked source PDF must be refused")
	} else if !strings.Contains(err.Error(), "symlink") {
		t.Errorf("error must name the symlink refusal, got: %v", err)
	}
}

func TestExtractRejectsBadVersion(t *testing.T) {
	if err := Extract(t.TempDir(), "9999.99999", "../x"); err == nil {
		t.Fatal("traversal version must be rejected")
	}
	if err := Extract(t.TempDir(), "9999.99999/../x", "v1"); err == nil {
		t.Fatal("traversal paper ID must be rejected")
	}
}
