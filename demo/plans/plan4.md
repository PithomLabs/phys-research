## Goal

Execute `prompt2.md` (Agent A — QSG Preflight 001) as a read-only infrastructure reconnaissance for the pivoted programme: conventional QFT on classical spacetime toward QSG-T001, with BHC preserved as an archived negative result. Produce the 8 required deliverables and answer the 6 completion questions.

## Success Criteria

- All 8 files (`01`–`08` per §9) exist in the assigned folder, each capability tagged IMPLEMENTED / DOCUMENTED ONLY / PARTIALLY IMPLEMENTED / MISSING / NOT VERIFIED with exact paths or command output.
- `07-implementation-contract.md` enforces the §8 trichotomy (mathematical proof vs executable verification vs research provenance) and states what must exist before the first QSG Phys derivation is accepted.
- `08-handoff.md` answers all 6 completion questions, including whether QSG-T001 can begin without architectural changes.
- Nothing outside the assigned folder is modified; BHC artifacts (both PDFs, `BHC_CORE_AND_COVERAGE.md`, prior BHC prompts) untouched.

## Approach

Reuse this session's verified infra map, then go deep only where QSG differs from BHC. Already established: `phys-pseudo/PSEUDO.md` is spec-only (no parser/checker — "no tooling is specified"); two GR references exist (frozen `phys-gr/solution/` + conformant `phys-lib/gr/equivariance/`); `phys-math` has 9 closed rules; `phys-trace` is a read-only MCP review server. New evidence gathered for this plan:

- `renormal`, `Hadamard`, `two-point`/`expectation value`/`operator-valued` have **zero matches** in `phys-math`, `phys-lib`, `phys-gr` — renormalization machinery is almost certainly MISSING (execution must confirm with a broader sweep including docs).
- General functional derivatives are explicitly UNSUPPORTED (`bohm/minisuperspace/manifest.json` limitations) — directly hits Workstream C's "functional derivatives / renormalized composite operators" questions.
- Stdlib-only confirmed for `phys`, `phys-gr`, `phys-math`, `phys-lib`, `phys-artifact`; only `phys-trace` has external deps (MCP SDK) — bounds the §8 dependency assessment.
- TOE-v6 still contains only `plans/` (no `research/` tree, no `.git`), so the assigned `research/experiments/qsg-preflight-001/agents/agent-a-infra-recon/` path and "dedicated Git worktree" cannot apply literally (see Q1).

## Steps

1. **Set up the assigned directory (sole write action).** Create the `qsg-preflight-001/.../agent-a-infra-recon/` path under TOE-v6 per assumption A1; log the TOE-v6 discrepancies (no repo, no `research/` tree) at the top of `01`.
2. **Workstream A → `01-repository-map.md`, `02-phys-pseudo-contract.md`.** Module map (`go.mod` graph, `AGENTS.md` hierarchy, frozen boundaries), 7-construct contract with `PSEUDO.md` §refs, pseudo→`DERIVATION_SPEC_GUIDELINES.md` handoff, CLI/MCP lifecycle inventory. Verify-no-parser with a quoted empty search.
3. **Workstream B → `03-gr-reference-analysis.md`.** Answer the 9 questions against `phys-gr/solution/` (manifest, 14-stage firewall `derivation_test.go`) and `phys-lib/gr/equivalence`+`bohm/minisuperspace` (nested layout, `Build`/`Verify`, goldens/anti-goldens); separate reusable protocol from GR-specific curvature machinery.
4. **Workstream C → `04-qsg-mathematical-feasibility.md`.** Score the 10 object classes (§5) plus the 6 hard probes (operator-valued quantities, functional derivatives, renormalized composites, state-dependent expectation values, curved-spacetime operators, nonlocal kernels) against `phys-math` rules/`families.go`/`store.go` and `phys-lib/gr` semantics. Keep representability strictly separate from verifiability; expected headline: renormalization/state machinery MISSING.
5. **Workstream D → `05-tdd-research-workflow.md`.** Ground RED–EXPLORE–MRC–GREEN–REFACTOR in real mechanisms; classify each MRC check AUTOMATED / PARTIALLY AUTOMATED / HUMAN OBLIGATION / NOT SUPPORTED with the test or gap that justifies it.
6. **Synthesize → `06-infrastructure-gaps.md`, `07-implementation-contract.md`, `08-handoff.md`.** Demonstrated gaps only (failing command or absent-symbol search + minimal remedy); contract with exact paths/commands/acceptance gates incl. stdlib-only impact; handoff with the 6 answers and smallest next task. Final read-back audit of all 8 files.

## Validation Plan

- Fast read-only gates: `go build ./...` + focused `go test -count=1` on `phys-artifact`, `phys-lib/gr/equivariance`, `phys-lib/bohm/minisuperspace`, `phys-gr/solution`; record exit codes and test names in `03`/`04`.
- phys-math shard C + purity/import gates only, per its `AGENTS.md` (heavy obligation/divergence shards only if Q2 says yes).
- Quoted empty searches as MISSING evidence (renormalization, Hadamard, two-point, operator-valued, nonlocal kernels) across code + docs.
- Non-modification proof: `git status --porcelain` in each `phys-*` repo before/after; BHC files in `plans/` checksummed before/after (`sha256sum`).

## Risks / Open Questions

- **Q1 (blocks step 1): TOE-v6 has no git repo, so "dedicated Git worktree and branch" is impossible as written.** Default A1: plain directory under TOE-v6, no worktree. Alternative: you `git init` TOE-v6 first. Which?
- **Q2 (cost): phys-math full suite is ~30 min serial.** Default: fast shard-C + targeted suites only. Authorize heavy shards if you want executed verdicts on checker depth?
- Assumption A2: the contract targets the `phys-lib` nested `protocol/` pattern as the canonical QSG-T001 shape (frozen `phys-gr/solution` as historical reference). Object now if you want `phys-gr/solution` as the anchor instead.
