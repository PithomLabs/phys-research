# Scoped Adversarial Code Review: Cross-Agent Output Integrity and Baseline Distribution

## Mission

Independently audit the latest cross-agent output-integrity changes in `phys-research`.

Your task is to find ways the implementation or documented procedure could accidentally or silently overwrite, replace, expose, or lose another coding agent's research output.

Be adversarial. Do not assume the implementation report is correct. Inspect the actual code, tests, documentation, Git state and filesystem behavior.

**This is a read-only review. Do not fix anything.**

## 1. Scope

Review only the latest changes and their directly related code paths:

- `docs/AGENT_CONTRACT.md`
- `internal/record/record.go`
- `internal/record/record_test.go`
- `internal/paper/fetch.go`
- `docs/BASELINE-2026-11-05.sha256`
- The relevant `AGENTS.md` instructions, especially §10 concerning `go.work`.

Inspect related functions and tests only where necessary to establish whether these changes are safe.

Do not perform a general security audit or redesign the architecture.

Do not modify:
- Existing research records.
- Branch C artifacts or conclusions.
- Other agents' worktrees.
- Existing branches or commits.
- The baseline manifest.
- Any existing uncommitted changes.

Do not commit, merge, reset, clean, restore, delete or rename anything.

## 2. Verify the actual baseline first

Inspect:

- `git status --short`
- `git diff`
- `git diff --check`
- `git worktree list`
- Relevant branch references.
- The baseline SHA-256 manifest.

Independently verify the report that:

- The baseline covers 63 files.
- 59 files remain byte-identical.
- The four differences are exactly the intended changes:
  - `docs/AGENT_CONTRACT.md`
  - `internal/paper/fetch.go`
  - `internal/record/record.go`
  - `internal/record/record_test.go`
- All Branch C records, equations, dossier and user-owned `AGENTS.md` remain byte-identical.
- The three agent worktrees and branches remain registered and untouched.

Flag any unexpected modification, missing file, untracked file or unexplained hash discrepancy.

Do not repair discrepancies.

## 3. Adversarial review: CreateRunDir

The new convention is:

`research/proposed/<agent-id>/<run-id>/<artifact-path>`

The implementation reportedly provides:

- Parent directory creation.
- Fail-closed leaf creation using `Mkdir`.
- Explicit `never reuse` error when the run directory already exists.
- Rejection of traversal and multi-segment run IDs.
- Fresh run identifiers using timestamp/UUID conventions.

Challenge every one of these claims.

### Required attack scenarios

A. Run ID validation:
- `..`
- `../other`
- Absolute paths.
- Embedded forward slashes.
- Embedded backslashes.
- Empty identifiers.
- Whitespace-only identifiers.
- Dot and dot-dot components.
- Null bytes.
- Platform-specific path separators.
- Unicode or unusual characters that could produce ambiguous paths.

B. Run freshness and reuse:
- Create the same run ID twice.
- Create two run IDs simultaneously.
- Simulate timestamp collisions.
- Simulate UUID collisions where practical.
- Restart an agent with stale environment variables.
- Attempt to reuse an existing run directory from another agent.

C. Filesystem attacks:
- Existing run directory.
- Existing regular file where a directory is expected.
- Symlinked parent directory.
- Symlinked run directory.
- Parent directory replacement during creation.
- Permission-denied conditions.
- Concurrent processes attempting the same destination.

D. Failure handling:
- Parent creation succeeds but leaf creation fails.
- Partial directory creation.
- Process interruption.
- Filesystem errors.
- Retry behavior after a collision.

Determine whether any failure can result in an existing run being reused, an outside path being created, or another agent's files being modified.

Pay particular attention to time-of-check/time-of-use races. Do not confuse `Mkdir` failing on an existing leaf with complete protection against symlinked parents.

## 4. Adversarial review: Cross-agent ownership

Verify what is actually enforced and what is only documented.

The contract states that each coding agent must operate from its assigned worktree and must not access another agent's worktree.

Check:

- Whether each agent has a unique worktree and branch.
- Whether launch instructions explicitly establish the correct working directory.
- Whether agent identity and run ID can be confused or reused.
- Whether output paths are consistently scoped by agent and run.
- Whether existing artifacts without ownership metadata remain protected.
- Whether a run can write into another agent's proposed-output directory through supported CLI operations.
- Whether direct filesystem writes can bypass the convention.

Explicitly distinguish:

1. Enforced by Go code.
2. Enforced by Git/worktree behavior.
3. Enforced only by agent instructions.
4. Not enforced.

Do not describe worktree separation as filesystem isolation. An agent with unrestricted filesystem access may still address another worktree through an absolute path.

## 5. Adversarial review: Paper-fetch collision safety

The `paper fetch` path reportedly changed from ordinary rename replacement to:

- Identical bytes: idempotent success without replacing the existing file.
- Different bytes: explicit collision error showing both digests.
- Concurrent fetches: exactly one winner.

Verify this at the filesystem-operation level.

Test or reason through:

- Existing file with identical contents.
- Existing file with different contents.
- Two concurrent downloads with identical contents.
- Two concurrent downloads with different contents.
- Process termination between temporary-file creation and final placement.
- Temporary-file cleanup.
- Existing destination symlinks.
- Existing destination directories.
- Rename and link behavior on supported operating systems.
- Whether error handling can accidentally remove or replace a previously valid PDF.

Verify that the collision path preserves the original bytes and manifest.

Do not access live external services. Use isolated temporary fixtures.

## 6. Adversarial review: Four-step baseline distribution procedure

The documented procedure is:

1. Export a patch for tracked changes and a tar archive for untracked files.
2. Run `git apply --check`, then `git apply`.
3. Extract untracked files using `tar --keep-old-files`.
4. Verify the destination against the manifest.

The reported dry run established:

- A conflicting tracked edit causes `git apply` to refuse.
- A pre-existing agent file causes tar extraction to refuse.
- The existing agent file survives byte-identically.
- A clean destination receives the expected files.
- The temporary worktree is removed.
- The three agent worktrees and branches remain registered and untouched.
- Fresh worktrees require a machine-local `go.work` recreated according to `AGENTS.md` §10.

Do not stop at repeating those individual tests.

### Critical combined-failure test

Test a mixed payload where:

- The tracked patch applies successfully.
- An untracked destination file already exists and conflicts with the tar archive.

Determine whether the procedure applies the tracked changes before tar refuses.

If so, does the destination remain partially updated?

This is a particularly important test: individually fail-closed operations do not necessarily compose into a fail-closed multi-step procedure.

Also investigate:

- Whether patch and tar payloads share a verified file manifest.
- Whether absolute paths or `../` paths in tar entries can escape the destination.
- Whether symlinks, hard links, duplicate paths or unusual file types are accepted.
- Whether extraction can follow a pre-existing symlink.
- Whether file permissions or executable bits are altered unexpectedly.
- Whether the manifest catches missing, extra, changed or replaced files.
- Whether a failed verification leaves partial changes behind.
- Whether the procedure can safely be repeated.
- Whether an existing `go.work` is preserved rather than overwritten.
- Whether the documented machine-local `go.work` recreation procedure is accurate.

Determine whether the distribution procedure is genuinely fail-closed, or merely prevents individual files from being overwritten.

Do not execute the distribution procedure against any real agent worktree. Reproduce it only in a disposable temporary directory.

## 7. Regression verification

Run only relevant, non-destructive verification:

- `go test ./internal/record`
- Relevant paper-fetch tests.
- `go test ./...`
- `go vet ./...`
- `go build ./...`
- `physctl validate`

Check that:

- Existing tests exercise real collision behavior.
- Tests do not merely verify return codes.
- File contents and hashes are checked.
- Concurrency tests are deterministic enough to be meaningful.
- Tests do not depend on external network access.
- Temporary test files and directories are cleaned up safely.

Do not change the implementation to make tests pass.

## 8. Required report format

Produce a concise adversarial report with these sections:

### A. Executive verdict
Choose one:
- PASS
- PASS WITH LIMITATIONS
- FAIL

### B. Findings

For each finding include:

- Finding ID.
- Severity: CRITICAL / HIGH / MEDIUM / LOW / INFORMATIONAL.
- Exact file and line/function.
- Attack or failure scenario.
- Reproduction command or test.
- Actual observed behavior.
- Impact on other agents' research outputs.
- Recommended minimal correction.

Do not report hypothetical vulnerabilities as confirmed findings. Label untested scenarios clearly.

### C. Enforcement matrix

| Guarantee | Go code | Git/worktree | Instructions only | Unresolved |
|---|---|---|---|---|
| Unique run ID | | | | |
| No run-directory reuse | | | | |
| No cross-agent overwrite | | | | |
| Fetch collision preservation | | | | |
| Baseline distribution safety | | | | |
| Correct launch working directory | | | | |
| Research baseline integrity | | | | |

### D. Test evidence

Include exact commands and results, including skipped or unavailable tests.

### E. Final disposition

Answer explicitly:

1. Can two agents independently create research outputs without silently replacing one another's files through the supported workflow?
2. Can an existing run directory ever be reused?
3. Can a distribution failure leave the destination partially updated?
4. Can baseline distribution overwrite an existing agent output?
5. Are the three agent worktrees and existing Branch C baseline intact?
6. Which guarantees depend only on agent instructions?
7. What is the smallest remaining change needed to make the workflow reliable?

## 9. Review boundaries

Do not implement fixes.

Do not broaden this into TLS remediation, OpenCode permission hardening, scientific methodology review, database design or agent orchestration.

Preserve all evidence and all work in progress.

The purpose of this review is to establish whether the latest cross-agent output protection and baseline distribution procedure actually work under adversarial conditions.

**A collision must preserve both sides. A failed distribution must not silently leave a half-integrated baseline. Existing research output must always survive.**
