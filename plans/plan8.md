## Goal

Two linked outcomes: (1) a complete, reproducible Branch C research cycle (Riemann–Cartan + quantum matter, three frameworks kept separate) yielding a traceable dossier; (2) `phys-research` established as a coding-agent-agnostic research core — OpenCode becomes one adapter among others, not a platform dependency. Physics quality first; portability must not delay the Branch C demonstration unnecessarily.

## Success Criteria

- Branch C: `physctl status` shows Branch C `?` obligations; load-bearing equations transcribed with locators (asserted/derived/reproduced marked); D-001 repaired with explicit transfer mappings; competing H-003 exists; all 10 dossier sections traceable; `validate` clean; `go test ./...` passes.
- Portability: shared methodology lives in platform-neutral files; OpenCode files are a thin adapter; a documented agent contract (required/optional capabilities) exists; compatibility matrix filled by real tests with at least OpenCode + one independent environment (Codex preferred) VERIFIED on the core workflow; no core capability depends exclusively on OpenCode.

## Approach

Reconnaissance findings (verified this session): build green, `validate` 0 errors, 6 papers cached, Branch C records honest-but-thin. Coupling points found: model pins (`opencode/mimo-…`, `opencode/deepseek-…`) in both `.opencode/agents/` and `agents/` mirrors; permission/acceptance boundary expressed as OpenCode rules (`opencode.json`, AGENTS.md §3–4); `opencode` mentions in `internal/paper/paper.go`, `internal/record/record.go`, one test. All three CLIs (`opencode`, `codex`, `muse`) are installed — cross-platform testing is feasible without new setup.

- **Key decisions**:
  - *Refactor shape*: extract shared methodology into platform-neutral `agents/*.md` (making the stale mirror the canonical source), slim `.opencode/agents/*.md` + `opencode.json` to adapter-only (model pins, permission syntax, subagent mode); add `docs/AGENT_CONTRACT.md` with the required/optional capability contract. No new runtime, plugins, MCP, or per-platform logic.
  - *Dossier form*: no new record type (closed header set). Dossier = repaired record set + assessment; thin `docs/dossier-branch-c.md` index only if needed.
  - *Researcher patch target*: the shared neutral file, not the OpenCode adapter (resolves the earlier authority question; OpenCode copy keeps only config).
  - *Ordering*: Branch C demonstration first (steps 1–6), portability refactor + matrix second (steps 7–8), so decoupling never blocks physics. Cross-platform test uses a Git worktree, never the same checkout simultaneously.
  - *Muse identity*: this session runs in Muse Code CLI — the Codex/Muse columns get tested here directly where possible; anything una runnable is marked NOT VERIFIED with setup notes, never invented.

## Steps

1. **Recon verification**: re-probe `paper search` with visible exit code; confirm D-001 parse behavior; baseline `validate`/`status`/`refs`; catalog every OpenCode-only dependency (APIs, config formats, permission behavior).
2. **Researcher methodology upgrade** (in shared neutral file): equation-transcription template, 4-level source qualification, per-framework assumption checklist, transfer-mapping rule, `?`-obligation rule.
3. **`physctl` minimal fix**: document `paper` subcommand in usage; fix or honestly mark `UNAVAILABLE` per probe result. No new commands.
4. **Sources + equations**: qualify P-001–P-006; transcribe load-bearing equations only (P-003 constraints, P-004 stochastic action, P-005 CPTP generator, P-006 EC equations + spin-density); add one foundational EC source (cap: 1–2 new).
5. **Comparison repair**: rewrite D-001 in place as justified correction (real refs, dedupe, typo fixes, variables/field equations/conservation/regime per framework, transfer/non-transfer mappings).
6. **Obligations + competitor + dossier**: H-002's 7 gaps → Phys-Pseudo `?` lines; H-003 via `physctl new` (`challenges` H-002); minimal E-002 touch-up; assemble 10-section dossier from records.
7. **Portability refactor**: canonicalize shared instructions in `agents/` (+ contract doc), slim OpenCode files to adapter, write minimal Codex/Muse adapter notes from observed CLI behavior (no invented config formats).
8. **Cross-platform matrix**: run the same Researcher workflow slice (read instructions → `physctl` run → create/validate one record → literature lookup) in OpenCode and Codex worktrees from this machine; fill matrix with VERIFIED / PARTIALLY / NOT VERIFIED / UNSUPPORTED plus setup notes.
9. **Evaluation**: 7 deliverables (recon, implementation, demonstration, quality assessment, validation, blockers, one next action) covering both physics and portability.

## Validation Plan

- `go build ./... && go test ./...`; `physctl validate` (0 errors); `status` (Branch C `?` visible); `refs I-002 --in --depth 2`.
- Record checks: every equation has a locator; every transfer names its mapping; H-003 has falsification condition; no `supports` without derivation.
- Portability checks: `grep` proves no methodology text remains only in `.opencode/`; matrix cells cite actual command runs (worktree paths + outputs); OpenCode-only acceptance-boundary text in AGENTS.md flagged for human rewrite (agents can't edit AGENTS.md).
- Highest-risk check: the `paper search` re-probe — if retrieval is broken, Step 4 uses cached fulltexts only and the matrix marks literature-discovery accordingly.

## Risks / Open Questions

- AGENTS.md and `.opencode/**` are constitution-restricted; proceeding on your explicit authorization, all such edits staged for your review.
- Live-model behavior in OpenCode/Codex subagents may differ from this session; matrix reports observed behavior only.
- Non-goals: permission hardening, isolation, dashboards, new roles/services/databases, orchestration frameworks, unrelated refactors. TLS stays deferred (cached fulltexts cover the Branch C demonstration; historical manifests do not prove current retrieval works — see Final Amendments §4).

## Final Amendments to Plan 8

Implement the following corrections:

1. **Canonical instruction reconciliation**
   - Do not promote `agents/*.md` as canonical merely because it exists.
   - Compare the shared instruction files and OpenCode copies.
   - Reconcile them against Git history and the latest Researcher methodology.
   - Preserve all current research capabilities and corrections.
   - Make the reconciled platform-neutral files canonical.
   - Keep model pins, permission syntax and execution-specific settings in platform adapters.
2. **Three-platform compatibility**
   - Test OpenCode, Codex and Muse because all three are reported available.
   - Use isolated Git worktrees and the same research fixture.
   - Record actual commands, outputs and limitations.
   - Never claim compatibility without executing the workflow.
   - Mark unavailable capabilities honestly.
3. **Portability contract**
   - Separate mandatory core capabilities from optional integrations.
   - The core research workflow must not require vendor-specific web search, MCP or subagents.
   - Demonstrate that each agent can operate the research core through shared instructions, repository artifacts and `physctl`.
4. **Cached literature and network retrieval**
   - Use adequate cached full texts for the Branch C demonstration.
   - Distinguish successful historical retrieval from verified current retrieval.
   - Do not enable unrestricted PDF fetching as part of this task.
5. **Scientific acceptance**
   - H-003 must represent a substantive competing hypothesis with distinct assumptions and a falsification condition.
   - Every load-bearing equation must have an exact source locator.
   - Every framework transfer must state the mapping and its limitations.
   - Keep unresolved mathematical questions as explicit `?` obligations.
   - Do not claim independent verification where only a source assertion exists.

Preserve Plan 8's execution order: Branch C demonstration first, portability second.

Do not expand the architecture to solve hypothetical future integration problems.
