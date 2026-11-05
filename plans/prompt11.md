The adversarial review found real weaknesses. I agree with its overall FAIL verdict, but not with every finding or its severity. The most important problems are the unsafe composition of baseline distribution, symlink handling in filesystem operations, and the difference between operational worktree separation and actual access isolation.

## 1. My assessment of the eight findings

| Finding                                         | My assessment                                                                                    | Priority                 |
| ----------------------------------------------- | ------------------------------------------------------------------------------------------------ | ------------------------ |
| CR-CA-001: `CreateRunDir` symlink attack        | Valid concern. `EvalSymlinks` alone would not close the race.                                    | High                     |
| CR-CA-002: Partial baseline distribution        | Valid. The explanation of tar's exit behavior is incorrect for GNU tar.                          | High                     |
| CR-CA-003: Tar path escape                      | Not established against the actual archive-generation procedure. Needs a realistic reproduction. | Medium                   |
| CR-CA-004: Paper-fetch symlink                  | Valid cache-integrity concern.                                                                   | High                     |
| CR-CA-005: Ambiguous run IDs                    | Overstated. `.run` is actually rejected by the stated regex.                                     | Low                      |
| CR-CA-006: `go.work` in manifest                | Valid if machine-local and excluded from distribution.                                           | Medium                   |
| CR-CA-007: Untracked files absent from manifest | Valid. Distribution needs its own complete payload manifest.                                     | Medium                   |
| CR-CA-008: No enforced worktree boundary        | Correct. Worktrees are not filesystem isolation.                                                 | Architectural limitation |

### Important correction: CR-CA-002

I tested `tar --keep-old-files` using GNU tar 1.35. With an existing destination file, it:

- Preserved the existing file.
- Reported `Cannot open: File exists`.
- Exited with status `2`.

Therefore, the report's assertion that it exits successfully and silently skips the file is incorrect for GNU tar.

However, the underlying finding is still valid: the tracked patch can be applied before the tar extraction fails. The destination is then partially updated.

The solution should be transactional staging, not merely replacing the tar option.

### Important correction: CR-CA-001

The proposed `EvalSymlinks` check is not sufficient by itself. A directory can be replaced after the check.

Use filesystem operations anchored to a trusted directory, preventing symlink traversal. Otherwise, explicitly document the remaining race rather than claiming it is fixed.

### Important correction: CR-CA-008

No Go helper can prevent an agent with unrestricted filesystem access from directly writing to another worktree.

We can make accidental cross-agent writes substantially less likely through worktrees, working-directory instructions and create-only operations. Genuine enforcement against arbitrary writes requires OS-level isolation.

That limitation should remain explicit.

## 2. What remains safe

The review reports that all three agent worktrees are clean and registered, and that the Branch C research baseline remains intact.&#x20;

adv_review_cross_agent.md



Do not reset, recreate or overwrite those worktrees while addressing these findings.

I would also avoid implementing a universal file-management framework. We only need to make the specific operations reliable.

## 3. Prompt for the implementation agent

# Targeted Remediation: Cross-Agent Output Integrity — Review 3

## Objective

Address the confirmed findings from `adv_review_cross_agent.md` while preserving all existing research outputs and agent worktrees.

This is a narrow remediation task. Do not redesign the architecture.

## 1. Preservation first

Before making changes:

- Record the current Git status, branches and worktree list.
- Verify the existing research baseline manifest.
- Preserve all existing agent worktrees and their contents.
- Preserve all Branch C artifacts, equations, D-001, H-002, H-003, E-002 and dossier.
- Do not reset, clean, restore, delete, rename, commit or merge existing work.
- Never overwrite an existing agent output.

## 2. Fix CreateRunDir path safety — CR-CA-001

Inspect `internal/record/record.go`.

Requirements:

- Prevent symlink traversal through agent and run directory paths.
- Validate agent IDs and run IDs as individual path components.
- Prevent concurrent path replacement from redirecting directory creation.
- Preserve create-only behavior.
- Existing run directories must always produce an explicit collision error.
- A failed attempt must never reuse an existing directory.

Do not rely exclusively on `EvalSymlinks` followed by ordinary `Mkdir`; that leaves a time-of-check/time-of-use race.

Prefer standard-library filesystem operations anchored to a trusted directory where supported by the project's Go version.

If complete protection cannot be implemented portably within existing constraints, document the precise limitation.

## 3. Fix paper-fetch symlink handling — CR-CA-004

Inspect `internal/paper/fetch.go`.

Requirements:

- Reject symlinks at the PDF destination.
- Reject symlinked or otherwise unsafe manifest destinations.
- Accept only an existing regular PDF whose content matches the downloaded content.
- Preserve the existing SHA-256 collision behavior.
- Preserve existing files on every failure.
- Never accept a symlink to another agent's cache as an equivalent local cached PDF.
- Ensure temporary files and manifest creation cannot overwrite existing valid outputs.

Add tests for:

- Symlink pointing to an identical PDF.
- Symlink pointing to different contents.
- Symlink pointing outside the cache.
- Concurrent placement.
- Failure during temporary-file placement.

Verify actual file contents and hashes.

## 4. Make baseline distribution transactional — CR-CA-002

The current procedure applies tracked changes before extracting untracked files. A later failure can leave a partially updated destination.

Replace this with a staged procedure:

1. Create a new, uniquely named staging worktree or staging directory.
2. Apply the tracked patch there.
3. Validate the untracked payload before extraction.
4. Reject absolute paths, traversal components and unsupported file types.
5. Reject symlinks and hard links unless explicitly required and safely handled.
6. Extract without replacing existing files.
7. Verify the complete payload manifest.
8. Run the relevant validation checks.
9. Only after every check succeeds, designate this as a new agent workspace.
10. On failure, leave every pre-existing agent worktree untouched.

Do not attempt to repair an existing agent worktree in place.

Preserve prior agent outputs even when baseline distribution fails.

### Important tar clarification

Independently verify the actual behavior of the target platform's tar implementation.

GNU tar 1.35 with `--keep-old-files` returns exit status 2 on an existing-file collision. Do not repeat the previous review's incorrect claim that GNU tar silently exits successfully.

Nevertheless, a nonzero exit does not roll back a patch already applied. The solution must address transactionality, not just exit codes.

Use the Go standard library if an extraction helper is genuinely necessary. Do not add external dependencies.

## 5. Correct baseline manifest handling — CR-CA-006 and CR-CA-007

Separate:

- Canonical source baseline.
- Machine-local configuration.
- Exact distribution payload.

Requirements:

- Exclude machine-local `go.work` from the portable source manifest.
- Document how each fresh worktree obtains its local `go.work`.
- Generate a payload manifest covering every tracked and untracked file actually distributed.
- Verify file contents and hashes.
- Detect missing, unexpected, changed and duplicate payload entries.
- Do not allow a successful verification when files were silently omitted.

Preserve the existing historical baseline manifest rather than rewriting its history without explanation.

## 6. Run-ID convention — CR-CA-005

Keep the convention:

`research/proposed/<agent-id>/<run-id>/<artifact-path>`

Confirm that:

- `..`, absolute paths and separators are rejected.
- Empty and invalid IDs are rejected.
- Run IDs cannot be reused.
- Concurrent creation is safe.

Do not treat `a..b` as path traversal merely because it contains consecutive dots. The current regex already rejects `.run` because its first character must be alphanumeric.

Only tighten the naming convention where it materially improves portability or prevents path ambiguity.

## 7. Clarify the cross-agent security boundary — CR-CA-008

Update `docs/AGENT_CONTRACT.md` to distinguish:

- Go-enforced create-only operations.
- Git/worktree separation.
- Agent-instruction requirements.
- Limitations arising from unrestricted filesystem access.

Explicitly state that worktree separation is an operational safeguard, not an OS-level security boundary.

Do not claim that `PHYS_AGENT_ID`, `PHYS_RESEARCH_ROOT` or working-directory instructions prevent a process from directly accessing another worktree.

Do not introduce OS sandboxing in this task.

## 8. Required adversarial tests

Add or update tests for:

- Symlink replacement and path traversal in `CreateRunDir`.
- Existing run directory refusal.
- Concurrent run creation.
- Paper-fetch symlink rejection.
- Preservation of valid cached PDFs and manifests.
- Mixed baseline conflict: tracked patch succeeds, untracked extraction fails.
- Payload path traversal and unsupported file types.
- Missing and extra payload files.
- `go.work` exclusion from portable verification.
- Failed distribution leaves existing agent worktrees byte-identical.

Run tests only in disposable directories or isolated new worktrees.

## 9. Verification

Run:

- `go test -count=1 ./internal/record`
- `go test -count=1 ./internal/paper`
- `go test -count=1 ./...`
- `go vet ./...`
- `go build ./...`
- `physctl validate`
- Complete baseline and payload hash verification.

Preserve the standard-library-only requirement.

## 10. Reporting

For each finding, report:

- FIXED
- NOT REPRODUCED
- PARTIALLY FIXED
- UNRESOLVED
- NOT TESTED

Include exact files, functions, tests and results.

Verify that all three existing agent worktrees and every existing Branch C artifact remain intact.

Do not commit, merge or delete anything.

Stop after verification.

**Success means a failed operation preserves existing research outputs, a successful new run receives an exclusive destination, and baseline distribution never partially modifies an existing agent workspace.**

## 4. Recommended decision

Hold baseline distribution

Continue to preserve the three existing worktrees, but do not distribute the updated uncommitted research baseline into them using the current procedure.

The most important architectural rule is now clear:

A new agent run must receive a new, verified workspace. Never update an existing agent's workspace in place.

That gives us a much simpler and safer operating model without introducing unnecessary infrastructure.
