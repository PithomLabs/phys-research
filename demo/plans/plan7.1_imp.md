Plan 7.1 execution complete. The Bohmian artifact is built, green, and handed off with the evidence-based coverage assessment below.

## What was delivered

**Executable artifact** (Benchmark 2): standalone module `github.com/PithomLabs/bohm-equiv` at [bohm-equiv](/home/chaschel/Desktop/biz/TOE-v6/research/experiments/bohm-preflight-001/theorems/bohm-equiv) — `fixture.go` (four 1D test pairs, 30 one-step DIFFERENTIATE/SIMPLIFY derivations, represented-only Flow/Pushforward/Equivariance relations), `manifest.json` + `source.json` + `source_mapping.json`, `manifest_test.go` (5 tests), `protocol/derivation_test.go` (15 tests). All 20 tests pass; `go vet` and `gofmt` clean.

**Transparent derivation**: [audit/01](/home/chaschel/Desktop/biz/TOE-v6/research/experiments/bohm-preflight-001/audit/01-premises-assumptions.md) (S1 freeze: N spinless particles, Schrödinger, B1–B7 register), [audit/02](/home/chaschel/Desktop/biz/TOE-v6/research/experiments/bohm-preflight-001/audit/02-continuity-derivation.md) (continuity DERIVED, with sign/dimensional self-checks), [audit/03](/home/chaschel/Desktop/biz/TOE-v6/research/experiments/bohm-preflight-001/audit/03-flow-equivariance.md) (transport + uniqueness + local equivariance `(Φ)#ρ₀=ρ_t` DERIVED; global SOURCE-REPORTED via inspected Teufel–Tumulka with a.e. qualification; nodes excluded).

**Placement deviation (your order, recorded):** plan7.1 prescribed `phys-lib/bohm/equivariance/`. Per your "do not modify any phys codebase," I removed that tree mid-Step-4 and rebuilt the artifact as a TOE-v6 standalone module on the qsg-t001 pattern. All three shared repos verify zero diff; the removal is covered by recovery ref `refs/tbh/recovery/before-discard/20261106T182751Z-39065`, and doc 01 §5 carries a dated addendum.

## The five platform questions

1. **Reused from QSG-T001:** the household, not the math — colocated-`go.work` standalone pattern, Build/Verify + derivation replay, manifest/sidecar conventions, golden/anti-golden/tamper/condition-binding test shapes, canonical-compare discipline. Representation layer reused: expression/relation/condition/derivation families, SIMPLIFY-001, SUBSTITUTE-001.
2. **Genuinely new:** executed DIFFERENTIATE-001 chains (QSG executed none), including nested second derivatives; density/current/quotient assembly on test inputs; point-substitution verdicts with idempotence-proved totality; represented-only claim relations with footprint-scan enforcement.
3. **PhysMath unpolluted:** yes — zero PhysMath changes, zero shared-repo writes of any kind. Imports are phys-artifact + phys-math only; all Bohmian semantics live in TOE-v6 docs, names, and comments.
4. **Stays theory-specific:** wavefunction/density/current/guidance/flow/equivariance semantics, test-pair definitions, condition vocabularies, claim-relation linkage — and anything measure/ODE/PDE if ever built (belongs in theory or analysis libraries, not the core).
5. **Demonstrated deficiencies (prioritized):** (i) **no decision procedure for rational-tree equality** — the #1 finding: SIMPLIFY-001 is non-confluent (executed `Mul(t,…)` vs manual `Mul(q,t,…)` normal forms), with no distributivity, like-term collection, or const-folding; even substituted trees don't fold (`simp(sub(Res4))≠0`, both recorded as negative results, limitation `equality_not_decided`); (ii) **no complex arithmetic** — forced the real (R,S) test-units formulation, Schrödinger residual unrepresentable; (iii) **no integration/ODE/measure/pushforward execution** — flow, transport, uniqueness, (E-loc)/(E-glob) represented-only; (iv) raw-form-only differentiation needing hand-staged partials; (v) substitution evaluates but can't fold; (vi) no node/domain semantics beyond NEQ conditions.

## QSG-vs-Bohm comparison

QSG-T001 pins curved-spacetime c-number goldens with SIMPLIFY-only execution and a sketch for the proof remainder; Bohmian equivariance executes differential chains (DIF/SIMPLIFY/SUBSTITUTE) on probability-flow fragments with represented-only relations for the theorem remainder. Both confirm the same joint result: the platform reliably **pins and replays explicit computations** with honest supported-subset tests, but **cannot close equational proofs** (gap i) or **represent analysis content** (gaps ii–iii). Per safeguard S2, this is an **initial cross-theory coverage assessment**, not a universality claim.

## Verification

QSG v2 gate re-run green (both scopes healthy); bohm-equiv 20/20 green; all JSON sidecars strict-valid; shared repos clean. Honesty audit performed manually: every proof obligation is dispositioned, and no test smuggles a conclusion — the two hand-authorized identities (product-rule EQ, pair-4 cancellation) are fenced by linkage/totality machine checks plus adjacent hand algebra, with the failed machine attempts kept as documented negative results. Hashes (sha256): fixture.go `5c0455a1…`, derivation_test.go `cfa766fb…`, manifest_test.go `b7daa66a…`, manifest.json `32bb4919…`, source_mapping.json `106d4700…`, audit 01/02/03 `bd595a78…`/`c917003d…`/`892158bc…`.

Step 0 bookkeeping verified present (O1 split, provisional wording, J/J′ note, reading list). Not claimed: external replication, general R^{3N} coverage, or anything beyond the frozen S1 scope.
