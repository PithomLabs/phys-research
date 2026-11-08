# 06 — Infrastructure gaps (demonstrated only)

Each gap: evidence (command output or inspected symbol), impact on
QSG-T001, minimal remedy, priority. Anything not demonstrated below is
absent by design or NOT VERIFIED — not a gap.

## G1. No committed multi-module workspace — HIGHEST (execution blocker)

- Evidence: `go build ./...` in `phys-math` fails standalone:
  `no required module provides package github.com/PithomLabs/phys-artifact`.
  `phys-lib/go.mod` carries the M-4 DEFERRED note (no replace;
  workspace-only resolution). No `go.work` is committed at
  `/home/chaschel/Documents/go/` (only per-model files under
  `phys-demo/` and generated ones).
- Impact: no independent reviewer can build without inventing the
  workspace recipe; every verification run needs the `/tmp` workaround.
- Remedy (minimal, 5 lines): commit a root `go.work` with `use` over
  `phys-artifact`, `phys-math`, `phys-lib`, `phys`, `phys-gr` — or add
  `replace` directives. Recipe in `07 §1`.
- Note: G1 is the smallest genuine blocker in the whole report, and it
  is procedural, not mathematical.

## G2. No quantum/renormalization representation families — HIGH (scoping)

- Evidence: full 34-family inventory (`phys-math/families.go`, all
  `Family*` constants read): expression/function/tensor/geometry/
  derivation/theorem/condition/equation families exist; no
  state/two-point/operator-algebra/renormalization/expectation family.
  AST tags (`ast.go`) and `operator.go` vocabulary (exactly
  {differentiate, covariant_derivative}) confirm: differential
  operators only, no quantum operators.
- Impact: QFT semantics (states, Hadamard, renormalization procedure,
  expectation operation) cannot be engine-represented; c-number
  consequences can. This bounds §7B/§7C verification, not expression.
- Remedy: none during recon (adjustment D). When a concrete derivation
  needs it: add representational (reserved, non-executable) families
  first, following the `operator.go` reservation precedent; execution
  only through bounded, approved proposals. The architecture anticipates
  new frameworks (`phys-system/AGENTS.md` §2.1) — this is growth, not
  redesign. Fundamental? No: representational.
- Until then: quantum semantics live in pseudo + `assumption/1` +
  hand derivation; engine pins c-number results as goldens.

## G3. General functional derivatives unrepresented — HIGH (scoping)

- Evidence: rule inventory shows only `DIFFERENTIATE-001` (partials via
  declaring function) and `TOTAL-DERIVATIVE-001`; `variation_reduction/1`
  is a bounded ansatz projection (minisuperspace `reduced_variation`
  premise + `reduced_variation_bounded` limitation, both read).
- Impact: deriving EL equations / metric variation from the action is a
  human obligation (or ansatz-bounded). Affects §7A.
- Remedy: same growth path as G2. Fundamental? No: bounded.

## G4. Engine identities are per-instance, never universal — discipline, not code

- Evidence: `obligation.go` (contracted Bianchi): "Replay success alone
  never promotes an instance result into a universal claim"; witnesses
  are pointwise/record-scoped with replayed closing derivations.
- Impact: a team expecting engine-proved conservation identities will be
  blocked. The correct pattern exists — QSG conservation must copy it
  (per-instance witnesses + human universal proof). No code remedy;
  the `07` contract assigns each identity accordingly.

## G5. Claim-status/promotion machinery lives in a sidecar — LOW

- Evidence: `phys-lib/PROTOCOL_GAPS.md` (read): GAP-CLAIM-STATUS,
  GAP-EXPECTED-RESULT, GAP-FALSIFICATION open; only `equivariance` is
  CONFORMANT. `ltb/research_metadata.json` demonstrates the working
  pattern (claims + obligations + evidence refs).
- Impact: RED-phase contract is split across two files. Workaround
  exists and is demonstrated — reuse the ltb pattern as-is.
- Remedy: GR-schema extension only via authorized decision; not recon
  work, not blocking.

## G6. No theorem-versioning vocabulary — LOW (NOT fully VERIFIED)

- Evidence: `schema_version` manifest field exists; no versioning type
  found in the seven `phys-lib/gr/*.go` files inspected
  (framework/assumption/definition/equation/theorem/interpretation/
  conservation). Full-package sweep NOT VERIFIED.
- Impact: versioned theorem changes are process-level (new manifest
  revision + preserved prior), not typed.
- Remedy: process suffices for QSG-T001; typed versioning only if
  revision churn demands it.

## Accepted non-gaps (by design, no remedy proposed)

- No Phys-Pseudo parser/checker/automation (PSEUDO.md §10; corpus
  README "explicitly out of bounds"). Human/agent process suffices.
- No solver, integrator, series engine, numerics (bounded substrate;
  recorded in GR manifest limitations). Hand derivation + goldens.
- Stdlib-only constraint (`01 §2`): all derivation-path modules have
  zero external deps. Preserve; no symbolic-algebra dependency.

## Priority order

G1 (commit workspace; unblocks all execution) → G2/G3 scoping decision
per QSG-T001 claim (represent-only vs bounded growth; a decision, not
code) → proceed. G4 is discipline. G5/G6 ride along.
