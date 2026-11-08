# 01 — Repository map

Recon date: 2026-11-06. All paths absolute. TOE-v6 discrepancies and the
plan4.md typo correction are recorded first, per plan4_review.md §3.3.

## 0. TOE-v6 discrepancies (authoritative-repo instructions absent)

- TOE-v6 (`/home/chaschel/Desktop/biz/TOE-v6`) contains only `plans/`. No
  `AGENTS.md`, no `research/` tree (created by this task, §5), no `.git`.
- Consequence: "dedicated Git worktree and branch" (prompt2 §2) cannot apply.
  Per plan4_review Q1: standalone assigned folder, no `git init`, read-only
  access to shared repos. Folder isolation is logical only, not a sandbox.
- Typo correction: plan4.md step 3 wrote `phys-lib/gr/equivalence`. Verified
  2026-11-06: `phys-lib/gr/equivalence` does NOT exist; the fixture is
  `phys-lib/gr/equivariance/` (AGENTS.md + manifest.json + protocol/ read).
  plan4.md itself left frozen; this note is the correction.

## 1. Instruction hierarchy (verified by reading)

| Level | File | Authority |
|---|---|---|
| Workspace entry | `/home/chaschel/Documents/go/AGENTS.md` | Pointer only: cross-repo work reads phys-system first |
| Collective manual | `phys-system/AGENTS.md` (`system.md` identical copy) | Cross-repo architecture, layer separation, agent ops |
| New-arch packages | `phys-math/AGENTS.md` (test shards), `phys-lib/gr/equivariance/AGENTS.md` (fixture contract) | Local details authoritative |
| Frozen legacy | `phys/AGENTS.md` | SR/mechanics workflow; never modified, never superseded |
| Frozen GR history | `phys-gr/AGENTS.md` (§8.1 artifact contract) | Historical reference only |
| Pseudo language | `phys-pseudo/AGENTS.md` + `PSEUDO.md` | Sole syntax authority; no tooling specified |
| Normative derivation spec | `DERIVATION_SPEC_GUIDELINES.md` (identical bytes in `phys-trace/` root and `phys-pseudo/conformance/`; verified with `diff -q`) | Implementation-ready requirements |

No root `AGENTS.md` in `phys-lib`, `phys-artifact` (contract = `doc.go` +
boundary tests), `phys-trace`, `phys-docs` (not a git repo, docs only).

## 2. Go module graph (all `go 1.25.7`, verified `go.mod` reads)

| Module | Dir | Requires | Replace | External deps |
|---|---|---|---|---|
| `github.com/PithomLabs/phys` | `phys/` | — | — | none (stdlib-only) |
| `github.com/PithomLabs/phys-gr` | `phys-gr/` | `phys v0.0.0` | `phys => ../phys` | none |
| `github.com/PithomLabs/phys-artifact` | `phys-artifact/` | — | — | none |
| `github.com/PithomLabs/phys-math` | `phys-math/` | — (imports `phys-artifact` without a require) | — | none, but **does not build standalone** |
| `github.com/PithomLabs/phys-lib` | `phys-lib/` | `phys-math v0.0.0` | — (M-4 DEFERRED, D17) | none, needs workspace |
| `github.com/PithomLabs/phys-trace` | `phys-trace/` | MCP SDK v1.8.0 + indirect | — | yes (MCP, oauth2, x/sync…) |
| `phys-pseudo` | `phys-pseudo/` | no `go.mod` (docs + fixtures only) | — | n/a |

No committed top-level `go.work` (only per-model files under
`phys-demo/problem_8a/*/` and generated ones from `phys-trace/scripts/`).
Demonstrated 2026-11-06: `go build ./...` in `phys-math` fails with
`no required module provides package github.com/PithomLabs/phys-artifact`.
All builds/tests in this report ran under a disposable workspace
`/tmp/qsg-recon/go.work` (`use` over the five modules) with
`GOCACHE=/tmp/qsg-recon/gocache` (default cache is read-only in this
environment). Recipe reproduced in `07`. Status: PARTIALLY IMPLEMENTED
workspace story; see gap G1 in `06`.

## 3. Package responsibilities (one line each, verified)

- `phys-artifact/`: `ArtifactRef` identity, canonical primitive framing
  (`primitives.go`), `BigInt`, digest formula. Owns no math/physics
  (`doc.go`, enforced by `boundary_test.go`). IMPLEMENTED.
- `phys-math/`: theory-neutral representation (34 artifact families, §4) +
  9 closed executable rules + replay. No physical vocabulary (gates PASS,
  §6). IMPLEMENTED (bounded by design).
- `phys-lib/gr/`: GR framework semantics — `FrameworkArtifact`,
  `PhysicalAssumption` (closed categories convention/regularity/model),
  `PhysicalDefinition/Equation/Theorem/Interpretation`, `ProtocolSketch`
  (non-executed step records, deliberately no Verify/Replay). IMPLEMENTED.
- `phys-lib/bohm/minisuperspace/`: Dürr–Struyve §6 fixture; executed
  fragments + declared limitations. IMPLEMENTED (toy scope).
- `phys-lib/gr/{equivariance,ltb,maxwellbox}`: equivariance = only
  manifest-CONFORMANT fixture; ltb/maxwellbox carry registered
  PROTOCOL GAPs (see `PROTOCOL_GAPS.md`). IMPLEMENTED.
- `phys/relativity/`, `phys/mechanics/`: frozen SR/classical-mechanics
  corpus + `derivation_test.go` E=mc² demo. IMPLEMENTED, frozen.
- `phys-gr/`: frozen Schwarzschild-vacuum derivation (`solution/`) +
  `metric/connection/curvature/tensor/symbolic/replay` machinery.
  IMPLEMENTED, frozen at `ea5ec2d` (matches the PROTOCOL_GAPS baseline).
- `phys-pseudo/`: language spec + MRC review protocol + conformance
  corpus (fixtures `01`–`08`, `S_cosmology_sketch`). Spec DOCUMENTED ONLY
  as process; automation explicitly MISSING by design (no parser/checker).
- `phys-trace/`: MCP review server (`cmd/phystrace-mcp`, tools
  `case_overview`, `inspect_evidence`, `run_verification`,
  `repository_status`) + protocol templates + e2e. IMPLEMENTED (demo scope;
  adapter on local branch only per its README).
- `phys-docs/`, `phys-system/`: documentation only. `phys-calc/`: empty
  directory. `phys-demo/problem_8a/`: multi-model comparison with
  per-model manifest + `protocol/derivation_test.go` + `go.work`.

## 4. Key source index (entry points used by this report)

- Families: `phys-math/families.go` (`Family*`, `AllPhysMathFamilies`)
- Rules: `phys-math/rules.go` (`Apply`, 9 `Rule*` IDs), `replay.go`
- Geometry: `riemann.go`, `riccicurv.go`, `riccitrace.go`, `coveval.go`,
  `covgeom.go`, `coordexpr.go`, `contraction_eval.go`, `levicivita.go`,
  `kronecker.go`, `obligation.go` (Bianchi evaluated-obligation)
- Operators: `phys-math/operator.go` (`ValidOperatorID`, `Executable`)
- Semantics: `phys-lib/gr/{framework,assumption,definition,equation,
  theorem,interpretation,conservation}.go`
- Manifest schema gate: `phys-lib/internal/manifestschema/*_test.go`
- Claim-status precedent: `phys-lib/gr/ltb/research_metadata.json`
  (`claims` + `obligations`, sidecar per GAP-CLAIM-STATUS)

## 5. This assignment's footprint (only writable location)

`/home/chaschel/Desktop/biz/TOE-v6/research/experiments/qsg-preflight-001/agents/agent-a-infra-recon/`
— the 8 deliverable files, nothing else.

## 6. Verification evidence (2026-11-06, `GOWORK=/tmp/qsg-recon/go.work`)

| Command (from module root) | Result |
|---|---|
| `go build ./...` (artifact, math, lib, gr) | exit 0 ×4 |
| `go test -count=1 ./...` `phys-artifact` | ok, 2.3s |
| `go test -count=1 ./...` `gr/equivariance` (+protocol) | ok |
| `go test -count=1 ./...` `bohm/minisuperspace` (+protocol) | ok |
| `go test -count=1 .` `phys-gr/solution` | ok, 3.2s |
| `go test -count=1 -timeout 20m -skip 'TestObligation\|TestDivergence\|TestWitness' .` `phys-math` (shard C) | ok, 9.8s |
| `go test -count=1 -run 'TestRegistryCompleteness\|TestNoPhysicalVocabulary\|TestImportDiscipline' .` | 3/3 PASS |
| `go vet` (math; lib equivariance+bohm) | exit 0 |
| `gofmt -l` (math, equivariance, bohm, gr/solution) | clean |

Heavy shards A/B deferred per Q2. Full `phys-lib/...` suite not run (scope:
fixtures above + gates).

## 7. Git baselines (before; after-state re-checked in `08`)

| Repo | HEAD | `status --porcelain` (before) |
|---|---|---|
| phys | `f9f5f7d` | clean |
| phys-gr | `ea5ec2d` | clean |
| phys-math | `7f3e0a7` | clean |
| phys-lib | `0f6c1cc` | `?? .kilo/` (pre-existing untracked, not mine) |
| phys-artifact | `31c21fe` | clean |
| phys-trace | `e1ee833` | clean |
| phys-pseudo | `b9bb9a4` | clean |
| phys-system | `efffa0e` | clean |
| phys-docs | n/a | not a git repo |

BHC artifacts (sha256, before; re-checked in `08`): BHC_CORE_AND_COVERAGE.md
`239a24c7…`, No-Go PDF `e7ad3a77…`, Preflight PDF `7e8e86d9…`,
prompt_agentB.md `f5ba7f6c…`, prompt_agentB2.md `9c59e163…`.
