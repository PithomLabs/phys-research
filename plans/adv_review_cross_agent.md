# SCOPED ADVERSARIAL CODE REVIEW: Cross-Agent Output Integrity and Baseline Distribution

## A. Executive Verdict

**FAIL** — The cross-agent output isolation and baseline distribution procedure has critical gaps that can silently overwrite or lose another agent's research output. While individual components (CreateRunDir, paper fetch collision, Git worktree isolation) are well-implemented, their composition and documented procedures contain fail-open behaviors.

---

## B. Findings

### CR-CA-001: CreateRunDir Symlink TOCTOU Allows Cross-Agent Overwrite
**Severity:** CRITICAL  
**Where:** `internal/record/record.go:241-263` (`CreateRunDir`)  
**Attack:** An attacker with write access to `research/proposed/<agentID>/` can replace the agent directory with a symlink to another agent's proposed directory between `os.MkdirAll(parent)` and `os.Mkdir(leaf)`. The function follows symlinks for both operations, causing the run directory to be created in the target agent's space.  
**Reproduction:**
```bash
# As attacker with write access to research/proposed/muse/
mkdir -p research/proposed/muse
# Create run dir for muse
# Between MkdirAll and Mkdir, replace with symlink:
rm -rf research/proposed/muse
ln -s ../codex research/proposed/muse
# Next CreateRunDir("muse", "run-1") creates directory in codex's space
```
**Actual Behavior:** No symlink validation; `filepath.Join` + `os.MkdirAll`/`os.Mkdir` follow symlinks.  
**Impact:** Agent A's run directory can be created inside Agent B's proposed directory, enabling silent overwrite of B's outputs.  
**Minimal Correction:** After `MkdirAll`, verify `filepath.EvalSymlinks(parent)` is within `filepath.EvalSymlinks(filepath.Join(root, ProposedDir))`. Use `os.Mkdir` with `O_NOFOLLOW` equivalent (not available in stdlib; must check post-creation).

### CR-CA-002: Baseline Distribution Not Fail-Closed (Partial Update on Mixed Conflict)
**Severity:** CRITICAL  
**Where:** `docs/AGENT_CONTRACT.md:100-125` (documented procedure)  
**Attack:** Tracked patch applies cleanly (`git apply` succeeds), but untracked tar extraction has a conflicting file. `tar --keep-old-files` **does not error** — it warns and skips the file, leaving the worktree with tracked changes applied but untracked baseline files missing.  
**Reproduction:**
```bash
# In main checkout
git diff HEAD > /tmp/baseline-tracked.patch
git ls-files --others --exclude-standard -z | tar --null -cf /tmp/baseline-untracked.tar --files-from=-

# In agent worktree, create conflicting untracked file
echo "agent's own file" > research/sources/P-999.md

# Run procedure
git apply --check /tmp/baseline-tracked.patch && git apply /tmp/baseline-tracked.patch  # SUCCEEDS
tar --keep-old-files -xf /tmp/baseline-untracked.tar  # WARNS, SKIPS P-999.md, EXITS 0
sha256sum -c docs/BASELINE-2026-11-05.sha256  # FAILS (P-999.md not in manifest, but other files missing)
```
**Actual Behavior:** `tar --keep-old-files` exits 0 on conflict (only warns). The procedure does not check tar's exit code or verify all files extracted.  
**Impact:** Destination left partially updated — tracked changes applied, untracked baseline files silently omitted. Subsequent manifest check fails but damage is done.  
**Minimal Correction:** Replace `tar --keep-old-files` with a Go implementation that: (1) reads tar entries, (2) verifies no absolute/`../` paths, (3) verifies no symlinks, (4) fails atomically if any destination exists, (5) only extracts on full success. Or use `tar -x --no-overwrite-dir` (GNU tar) with explicit error checking, but prefer Go for portability.

### CR-CA-003: Tar Extraction Allows Path Escape and Symlinks
**Severity:** HIGH  
**Where:** `docs/AGENT_CONTRACT.md:117` (`tar --keep-old-files -xf`)  
**Attack:** Tar archive contains entries with absolute paths (`/etc/passwd`), parent traversal (`../../../etc/passwd`), or symlinks pointing outside destination. `tar -x` follows these by default.  
**Reproduction:** Create tar with `../../../etc/passwd` entry or symlink to `/etc`. Extract in worktree — writes outside worktree.  
**Actual Behavior:** No `--no-absolute-names`, `--no-recursion`, or symlink rejection in documented command.  
**Impact:** Arbitrary file write outside agent worktree, potentially overwriting system files or other agents' data.  
**Minimal Correction:** Use Go's `archive/tar` with explicit path validation: reject absolute paths, `..` components, symlinks, hard links, non-regular files.

### CR-CA-004: Paper Fetch Collision Vulnerable to Symlink Replacement
**Severity:** HIGH  
**Where:** `internal/paper/fetch.go:52-94` (`placeDownload`)  
**Attack:** If `paper-cache/<id>/<version>/source.pdf` is a symlink to another agent's PDF or arbitrary location, `os.ReadFile` follows it (reading target), `exclWrite` fails with `EEXIST` (symlink exists), then re-reads the target and compares digests. If attacker crafts target with matching digest, fetch succeeds but manifest written to original location — inconsistent state.  
**Reproduction:**
```bash
ln -s /home/attacker/fake.pdf paper-cache/2209.15487v1/v1/source.pdf
# fake.pdf has same SHA-256 as legitimate download
physctl paper fetch 2209.15487v1 v1 ...
# Reads fake.pdf, exclWrite fails (symlink exists), re-reads fake.pdf, digests match
# Manifest written to paper-cache/.../manifest.json (not following symlink)
# PDF effectively "placed" but is actually a symlink to attacker-controlled location
```
**Actual Behavior:** No `EvalSymlinks` + boundary check on `pdfPath`/`manifestPath` before operations.  
**Impact:** Cache integrity compromised; source records cite SHA-256 of symlink target, not actual cached file.  
**Minimal Correction:** In `Fetch`, after `MkdirAll(dir)`, verify `filepath.EvalSymlinks(pdfPath)` and `manifestPath` are within `filepath.EvalSymlinks(paperCacheDir(root))`. Reject if not.

### CR-CA-005: Run ID Validation Allows Ambiguous Identifiers
**Severity:** MEDIUM  
**Where:** `internal/record/record.go:239` (`runIDPattern = ^[A-Za-z0-9][A-Za-z0-9._-]*$`)  
**Attack:** Run IDs like `a..b`, `a.-b`, `run.`, `.run` are valid per regex but create ambiguous filesystem paths. `a..b` is not `..` but contains double-dot; `.run` starts with dot (hidden file).  
**Reproduction:** `CreateRunDir(root, "muse", "a..b")` succeeds, creates `research/proposed/muse/a..b`.  
**Impact:** Potential confusion with `..` traversal, hidden directories, or tools that normalize paths.  
**Minimal Correction:** Tighten regex: `^[A-Za-z0-9][A-Za-z0-9_-]*(\.[A-Za-z0-9_-]+)*$` (no leading dot, no consecutive dots, no trailing dot).

### CR-CA-006: Baseline Manifest Includes Machine-Local go.work
**Severity:** MEDIUM  
**Where:** `docs/BASELINE-2026-11-05.sha256:65` (`go.work` hash)  
**Attack:** Manifest verification (`sha256sum -c`) fails on `go.work` because each machine has different paths. Procedure says "go.work is machine-local and never distributes" but manifest includes it, causing false failure.  
**Actual Behavior:** `sha256sum -c` reports `go.work: FAILED` on every machine with different paths.  
**Impact:** Manifest check cannot pass cleanly; operators must mentally filter the failure.  
**Minimal Correction:** Remove `go.work` from manifest, or add explicit `--ignore-missing`/`--status` filtering in verification step.

### CR-CA-007: Untracked Baseline Files Not Inventoried in Manifest
**Severity:** MEDIUM  
**Where:** `docs/BASELINE-2026-11-05.sha256` vs `git ls-files --others --exclude-standard`  
**Attack:** New untracked files created after baseline freeze (e.g., `internal/record/worktree_test.go`, `internal/paper/transport_test.go`) are distributed via tar but not in manifest. Manifest check ignores them (not in manifest), but they're still distributed. If an agent already has a different version, `tar --keep-old-files` silently skips — agent keeps their version, not baseline.  
**Impact:** Baseline distribution incomplete for post-freeze files; agent worktrees diverge silently.  
**Minimal Correction:** Either (a) include all distributed files in manifest, or (b) document that only pre-freeze files are baselined and post-freeze files are agent-responsibility.

### CR-CA-008: Cross-Agent Ownership Not Enforced by Go Code
**Severity:** HIGH (Documentation vs Implementation Gap)  
**Where:** `docs/AGENT_CONTRACT.md:64-71` vs `internal/record/record.go`  
**Claim:** "Each coding-agent process must be launched with its assigned worktree as its working directory, and must not access another agent's worktree. Checkouts never share files..."  
**Reality:** No Go code enforces worktree boundary. An agent with `PHYS_RESEARCH_ROOT` set to another agent's worktree path, or using absolute paths, can read/write any worktree. The `CreateRunDir` uses `root` parameter — if caller passes another agent's worktree root, it creates run dirs there.  
**Enforcement Matrix:** See Section C.  
**Impact:** Operational safeguard only; misbehaving or compromised agent can cross worktrees.  
**Minimal Correction:** Document clearly as "operational convention, not enforced." Consider adding `PHYS_AGENT_ID` validation in `CreateRunDir` (compare with directory name).

---

## C. Enforcement Matrix

| Guarantee | Go Code | Git/Worktree | Instructions Only | Unresolved |
|---|---|---|---|---|
| Unique run ID | ✅ `CreateRunDir` rejects reuse | | | |
| No run-directory reuse | ✅ `os.Mkdir` fails on exist | | | Symlink TOCTOU (CR-CA-001) |
| No cross-agent overwrite (canonical) | | ✅ Git merge conflicts | | |
| No cross-agent overwrite (proposed) | ❌ Symlink TOCTOU | | ✅ Contract says "operational safeguard" | CR-CA-001, CR-CA-008 |
| Fetch collision preservation | ✅ `placeDownload` with digests + O_EXCL | | | Symlink escape (CR-CA-004) |
| Baseline distribution safety | ❌ Shell procedure | | ✅ Documented | CR-CA-002, CR-CA-003, CR-CA-006, CR-CA-007 |
| Correct launch working directory | | | ✅ Contract §12.2 | CR-CA-008 |
| Research baseline integrity | | ✅ Baseline manifest (63 files) | | go.work false positive (CR-CA-006), missing files (CR-CA-007) |

---

## D. Test Evidence

| Command | Result | Notes |
|---|---|---|
| `go test -count=1 ./internal/record` | PASS | All 22 tests pass, including `TestCreateRunDir` |
| `go test -count=1 ./internal/paper` | PASS | 9 tests pass, including collision tests |
| `go test -count=1 ./...` | PASS | All 9 packages OK |
| `go vet ./...` | PASS | No issues |
| `go build ./...` | PASS | |
| `physctl validate` | PASS | 0 errors, 1 LEGACY_UNPINNED warning |
| `sha256sum -c docs/BASELINE-2026-11-05.sha256` | 4 FAILED | `docs/AGENT_CONTRACT.md`, `internal/paper/fetch.go`, `internal/record/record.go`, `internal/record/record_test.go` — exactly the 4 intended post-freeze changes |
| `git worktree list` | 4 worktrees | Main + 3 agents, all at a007dc3 |
| `git status --short` (all worktrees) | Clean | No uncommitted changes in agent worktrees |

**Missing Tests:**
- No test for `CreateRunDir` symlink attack
- No test for `placeDownload` symlink escape
- No test for baseline distribution procedure (shell-only)
- No concurrent cross-process `CreateRunDir` test with same agentID/runID
- No test for tar extraction path escape/symlinks

---

## E. Final Disposition

### Explicit Answers

1. **Can two agents independently create research outputs without silently replacing one another's files through the supported workflow?**
   → **NO.** `CreateRunDir` has symlink TOCTOU (CR-CA-001). Baseline distribution can partially update (CR-CA-002). Paper fetch has symlink escape (CR-CA-004).

2. **Can an existing run directory ever be reused?**
   → **NO** via direct reuse (explicit error). **YES** via symlink replacement of parent directory between `MkdirAll` and `Mkdir`.

3. **Can a distribution failure leave the destination partially updated?**
   → **YES.** Tracked patch applies, tar extraction has conflict → tracked changes remain, untracked files silently skipped. Procedure is not atomic.

4. **Can baseline distribution overwrite an existing agent output?**
   → **NO** for tracked files (`git apply` refuses on conflict). **YES** for untracked files — `tar --keep-old-files` silently skips, leaving agent's version (not baseline). If agent's version differs from baseline, divergence is silent.

5. **Are the three agent worktrees and existing Branch C baseline intact?**
   → **YES.** All three worktrees exist at `phys-research-agents/{codex,muse,opencode}`, all at commit `a007dc3`, all clean (`git status` empty). Branches `agent/codex`, `agent/muse`, `agent/opencode` registered. Baseline manifest covers 63 files; 59 match, 4 differ exactly as intended (post-freeze changes).

6. **Which guarantees depend only on agent instructions?**
   → Cross-agent worktree isolation (CR-CA-008), correct launch directory, not accessing another agent's worktree, not modifying canonical records in-place. No Go code enforces these.

7. **What is the smallest remaining change needed to make the workflow reliable?**
   → **Three minimal changes:**
   1. Add symlink boundary check in `CreateRunDir` (after `MkdirAll`, verify resolved path within `ProposedDir`).
   2. Replace shell baseline distribution with Go implementation that: validates tar entries, extracts atomically (all-or-nothing), verifies manifest, excludes `go.work`.
   3. Add symlink boundary check in `Fetch`/`placeDownload` (verify `EvalSymlinks(pdfPath)` within cache root).

---

## Summary

The implementation correctly handles the **common case** (no symlinks, no conflicts, clean distribution). The **adversarial cases** (symlinks, partial conflicts, path escape) are not protected. The baseline distribution procedure is the weakest link — it is a documented shell script with known fail-open behaviors, not implemented code.

**Recommendation:** Do not authorize autonomous multi-agent operation until CR-CA-001, CR-CA-002, and CR-CA-004 are fixed. The Git worktree layer provides real isolation for canonical records; the proposed-output layer and baseline distribution do not.
