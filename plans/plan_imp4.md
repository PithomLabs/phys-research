Follow-up complete. Stopped after verification — no commits, merges, or deletions.

## Count reconciliation (issue 2C)

12 = 2 (H-001, pre-existing Branch Q/MICROSCOPE) + 7 (H-002) + 3 (H-003). The report's "10" counted Branch C only. Verified per-record via `status`; no duplicated or orphaned obligations.

## Follow-up execution

1. **Inventory**: [docs/BASELINE-2026-11-05.sha256](/home/chaschel/Desktop/biz/TOE-v5/phys-research/docs/BASELINE-2026-11-05.sha256) — 63 files (research outputs, methodology, agent config, Go sources, your AGENTS.md), self-check clean.
2. **No disturbance**: nothing reset, cleaned, restored, or deleted.
3. **Worktrees + branches**: `agent/muse`, `agent/codex`, `agent/opencode` at HEAD with sibling checkouts under `phys-research-agents/`. They start at pre-freeze HEAD; the frozen Branch C baseline lives only in the main checkout + manifest. Integration is a future human merge.
4. **Create-only**: `record.Create` was already O_EXCL with race-retry (verified, not assumed). The real silent-replacement bug was `paper fetch` (`os.Rename` clobbered the cache) — now collision-safe: identical re-fetch is idempotent and untouched; differing bytes yield an explicit `collision` error naming both digests.
5. **Canonical preservation**: records stay create-only; revisions go to `research/proposed/<agent>/<path>`, which tooling ignores (not a scanned dir) so proposals can't disturb validation.
6. **Concurrency**: racers get distinct IDs (records) or exactly-one-winner + collision errors (fetch); worktree separation is structural.
7. **Tests**: 4 new — unowned-file preservation, 6 fetch-placement cases (fresh/idempotent/collision/concurrent), git worktree merge-conflict preserving both sides (skips without git). Plus 2 pre-existing (concurrent IDs, overwrite refusal). All pass.
8. **Hash verification**: 60/63 manifest entries byte-identical after all work. The 3 diffs are exactly the intended edits (`fetch.go`, `record_test.go`, `AGENT_CONTRACT.md`). All Branch C research, P/H/E/I/D records, dossier, and your AGENTS.md untouched.
9. **Branch C preserved**: equation transcriptions, D-001, H-002, H-003, E-002, dossier all hash-identical.
10. **No unrelated changes**: no conclusion or methodology edits. Full suite 9/9 green; `validate` 0 errors.

## Enforced / documented / unresolved

- **Enforced in code**: O_EXCL record creation with concurrent-ID absorption; fetch bytes-same-or-error; merge conflicts preserve both sides (git's own behavior, test-pinned).
- **Merely documented**: proposed-version path convention; worktree ownership protocol; agent-workspace start state.
- **Unresolved**: nothing technical within this scope. Residual truth: no Go code can stop an agent from editing files directly with its own tools — that vector is fenced by worktree separation + human merge review, not by tooling. `hooks/` could harden it but is constitution-unmodifiable; left for you.
