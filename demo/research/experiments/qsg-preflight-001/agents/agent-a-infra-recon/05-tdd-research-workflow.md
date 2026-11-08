# 05 — TDD research workflow (RED–EXPLORE–MRC–GREEN–REFACTOR)

Every phase below names an existing mechanism. Nothing here is proposed;
all of it was read or executed during this recon.

## RED — Theorem contract

Record the contract before deriving, using three existing artifacts:

1. `manifest.json` (16-key GR schema; `equivariance/manifest.json` is the
   conformant model): `derivation_id`, `regime`, `coordinates`,
   `conventions{}`, `premises[]`, `assumptions[]`, `expected_result`,
   `verification{}`, `limitations[]`, `falsification_conditions[]`.
   Enforced by `internal/manifestschema/manifest_conformance_test.go`.
2. `research_metadata.json` (ltb precedent): `target`, `domain`,
   `claims[]` with epistemic statuses (`open-conjecture`,
   `candidate-construction-incomplete`, …), `obligations[]` with
   `blocks[]` + `evidence_required`. This is where failure conditions
   and MRC obligations live as data, not prose.
3. Phys-Pseudo `?` lines for obligations still too vague for (2).

The final QSG-T001 proposition comes from the physics investigation
(Agent B); this workflow only fixes where each contract element lives.
Do NOT use `expected_result` for unexecuted claims (GAP-EXPECTED-RESULT:
that field asserts executed, verified closed form).

## EXPLORE — Phys-Pseudo

Seven constructs (`02`), with these load-bearing separations:

- Proposals vs results: the §7 invariant (representation ≠ endorsement)
  plus the sidecar claim-status vocabulary — a `claim` in pseudo stays a
  hypothesis until a derivation artifact + `Verify` + review promote the
  sidecar status. No promotion by computation exists anywhere (verified:
  no promotion API in `phys`, `phys-math`, or `phys-lib`).
- Competing hypotheses: corpus fixture `08_competing_hypotheses`
  demonstrates the form; no syntax extension needed.
- Testability: every `derive` needs an explicit result + premises;
  MRC review checks premise-completeness counterexample-first.

## MRC — check classification (prompt2 §MRC list)

| Check | Class | Mechanism / gap |
|---|---|---|
| Dimensional consistency | HUMAN | New arch has no dimension concept (`phys-math` purity gates forbid physical vocabulary; verified PASS). Frozen `phys` checks dimensions but is SR/mechanics-scoped and frozen — precedent only, not reusable. |
| Tensor / index structure | PARTIALLY AUTOMATED | `Check*` validators (affine-connection shape, variance/slot discipline, `CoordParamMap` checks); `tensor/1` sparse semantics deferred. |
| Covariance | PARTIALLY AUTOMATED | Chart-bound machinery enforces internal consistency (identity scoping, disjoint checks); general covariance semantics are human. |
| Renormalization assumptions | HUMAN | `assumption/1` + pseudo `assume` record them; no engine check. |
| Conservation | PARTIALLY AUTOMATED | Per-instance evaluated obligations (Bianchi precedent, `obligation.go`); universal identities human. |
| Mathematical invariants | PARTIALLY AUTOMATED | Symmetry/antisymmetry re-verification + golden pinning; invariant selection human. |
| Limiting cases | PARTIALLY AUTOMATED | Golden/anti-golden pinning machinery; limits hand-derived (no general series engine — recorded GR limitation). |
| Constraint consistency | PARTIALLY AUTOMATED | Per-instance checks; propagation analysis human. |
| Physical regime / scale | HUMAN | `regime` + assumption categories record it; no engine notion of scale. |
| Known obstructions | HUMAN + process | `obligations[]` sidecar + `ProtocolSketch` UNSUPPORTED steps + adversarial review (below). |

## GREEN — Phys derivation

Follow the equivariance packaging contract (`gr/equivariance/AGENTS.md`,
read): fixture dir with `AGENTS.md` + `manifest.json` +
`protocol/derivation_test.go`; `Build` mints artifacts into a
`physmath.Store`; `Verify` replays every derivation and re-checks
construction relations; manifest embedded + strict-typed decode +
unknown-field rejection; golden/anti-golden pairs; mutation negatives
(skipped/swapped/corrupted stages must fail); `TestManifest*` suite
(exact values, prose-is-not-math, no-derivation-shortcut).
Analytical derivation ↔ equations ↔ tests connect through references:
`PhysicalTheorem.MathematicalTheoremRef`,
`AssumptionRefs`/`InterpretationRefs`, derivation output refs — never
through prose.

## REFACTOR — independent verification

| Need | Mechanism (status) |
|---|---|
| Independent reproduction | Workspace recipe (`01 §2`) + `go test -count=1` fixture suites; `phystrace-mcp run_verification` profiles (`demo-arc` 120s, `coverage-audit` 180s, `full-suite` 600s). IMPLEMENTED. |
| Adversarial mathematical review | `MRC_REVIEW_PROTOCOL.md` (8 responsibilities, findings + evidential boundary, no global verdict) + `inspect_evidence` selectors (manifest, inventory, q-relation, anti-goldens, covariant-fragment, limitations, file) + `phys-trace/protocol/prompts/*` review templates. IMPLEMENTED as process. |
| Versioned theorem changes | PARTIALLY: `schema_version` + new-manifest-revision process; no theorem-versioning machinery in code (NOT VERIFIED beyond schema field — no version type found in `phys-lib/gr`). |
| Failed / superseded derivations | `research_metadata.json` obligations + `ProtocolSketch` (UNSUPPORTED steps, no Verify by design) + preserved fixtures (ltb keeps its open conjecture visible). IMPLEMENTED. |
| Provenance / evidence | `ArtifactRef` digest chains, derivation input/output refs, `source_mapping.json` sidecar (structural-only, enforced by `TestSourceMappingStructuralConsistency`). IMPLEMENTED. |
| Regression after assumption changes | Fixture `derivation_test.go` suites + `TestVerifyCoverageAudit`-style coverage guards. IMPLEMENTED. |

No new workflow is introduced. The single process gap is
theorem-versioning vocabulary (see `06`, minor).
