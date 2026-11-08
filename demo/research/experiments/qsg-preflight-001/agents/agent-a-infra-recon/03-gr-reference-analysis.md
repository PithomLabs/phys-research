# 03 — GR reference analysis

Per Q3: `phys-gr/solution` = canonical protocol reference (frozen);
`phys-lib/gr/equivariance` = current implementation/packaging reference;
`phys-lib/bohm/minisuperspace` = nearest-physics precedent (executed
fragments + declared limits). All test names below from `go test -list`
runs on 2026-11-06; exits in `01 §6` (all ok).

## 1. The 9 questions (prompt2 §4)

**1. What makes the theorem executable?** An ordered derivation trace that
the test re-executes and re-verifies step by step — not a stored result.
`phys-gr/solution`: `Primary()` emits 14 `GRStep` stages;
`verifyPrimaryTrace` asserts stage order, golden intermediate states
(9 Christoffel classes, `R_thth` general form, exact `A(r)*B(r)` diff
input, `k2` first-appearing at stage 12), and byte-identity linkage
between stages. `equivariance`: 6 transformations
(APPLY,APPLY,CANON,APPLY,APPLY,CANON) with byte-linkage of staged Calls
and conclusion gate (`Verify`+`VerifyClaim`). `minisuperspace`:
`Fixture.Verify` replays every stored derivation (`Derivs` +
`FRWDerivRef`) and re-checks construction relations.

**2. What mathematical claims are actually verified?** Structural/
computational facts only: trace order and golden intermediates
(`TestPrimaryDerivationTrace`, `TestEquivarianceTrace`,
`TestDerivationInventory`); operator-composition goldens and
anti-goldens (`TestOperatorCompositionVsSquare`,
`TestHistoricalWrongPath`); replay of every evidence derivation
(`TestVerifyReplaysAllEvidence`); claim→conclusion linkage
(`TestClaimTracePasses`); manifest↔code consistency
(`TestManifestExactValues`, `TestManifestExpectedResult`,
`TestLimitationRecord`). No test asserts a general theorem, physical
truth, or empirical content.

**3. How are assumptions represented?** Three layers: (a) manifest
`premises[]`/`assumptions[]`/`conventions{}` (metadata, strict typed
decode, unknown fields rejected — `TestManifestRejectsUnknownFields`);
(b) content-addressed `PhysicalAssumption` artifacts with closed
categories (convention/regularity/model) and `condition/1` refs
(`phys-lib/gr/assumption.go`); (c) per-step conditions threaded through
`Apply`/`Invocation.ConditionRefs` and preserved (e.g. `NEQ(a,0)` in
minisuperspace). Manifest is metadata only — never a premise
(`TestManifestNoDerivationShortcut`, equivariance firewall).

**4. How are intermediate derivation steps recorded?** As a replayable
chain of minted artifacts: each step names producer, rule, exact input
refs, conditions, parameters, output ref (`physmath.Derivation`;
`replay.go` fails closed on missing/wrong-family/corrupted evidence).
`phys-gr` uses its own `GRStep` trace + `replay.Step` with canonical
codecs. Anti-shortcut: mutation tests corrupt/skip/swap stages and must
fail (`TestMutation*`, 8 in solution, `TestMutationNegatives`,
`TestFreshVerifierMutations`, `TestVerifyRejectsMissingEvidence`).
Any passing mutation is a fixture bug, not tolerance.

**5. How are limiting cases and invariants tested?** As golden/anti-golden
pairs and invariant re-checks: symmetry/antisymmetry re-verified from
outputs (Christoffel `VerifySymmetry`, Riemann `VerifyAntisymmetry`);
boundary behavior pinned (asymptotic-flatness `A,B→1` substitutions,
stage 8); factored-vs-collapsed forms pinned with anti-goldens
(`a²∂²_a` vs `(a∂_a)²`, wrong-pressure path); vacuum certification
counts (16 inverse `zero_test`s, all `ZERO`); falsification conditions
enumerated in manifest (`nonzero_vacuum_component`,
`trace_replay_failure`, `closed_form_mismatch`).

**6. How are conditional results represented?** Conditions travel with the
artifacts (`condition/1` refs, `NEQ`/`EQ` relations), and open conjectures
use the `ltb` sidecar pattern: `research_metadata.json` with `claims[]`
(statuses `open-conjecture`, `candidate-construction-incomplete`, …)
blocked by explicit `obligations[]` (`O1`–`O7`) — never promoted to
`expected_result` (see GAP-EXPECTED-RESULT). Non-executed reasoning is
preserved as `ProtocolSketch` steps with `SUPPORTED`/`UNSUPPORTED` status
and deliberately NO Verify/Replay method.

**7. Proof obligation vs executable test.** The architecture enforces the
split mechanically: replay proves *the recorded computation reproduces
from its premises*; goldens prove *byte-identity with an independently
derived expectation*; mutations prove *the test is not vacuous*. None of
these is a proof of a general proposition — generality lives in the
human derivation plus per-instance obligation records (see Bianchi
precedent in `04`). Prose is never math (`TestManifestProseIsNotMath`,
`TestManifestExpectedResult`).

**8. Reusable for QSG-T001.** Everything protocol-shaped: manifest schema
+ strict decode + unknown-field rejection; `protocol/derivation_test.go`
pattern (positive trace + mutation negatives + manifest tests);
`Build`/`Verify` fixture shape; golden/anti-golden discipline;
`research_metadata.json` claims/obligations sidecar;
`ProtocolSketch` for preserved-but-unexecuted steps; `physmath`
replay + `Check*` validators + evaluated-obligation records;
`phystrace-mcp` independent-review tooling (read-only inspect + run).

**9. GR-specific, must not be generalized.** The Schwarzschild ansatz and
14-stage derivation; MTW/`-+++`/`Λ=0` conventions as content (reusable
only as *slots*); `phys-gr`'s `symbolic/metric/connection/curvature`
packages (frozen, GR-shaped — QSG curvature needs go through
`phys-math`'s chart-bound evaluators instead); the `k2=-2μ`
correspondence (HYPOTHESIS provenance, not a premise); equivariance's
schematic `f(v)=2v` content.

## 2. Distinctions (prompt2 §4, mapped to artifacts)

| Category | What it is here | Example |
|---|---|---|
| Analytical derivation | Human proof, recorded in prose + pseudo | minisuperspace hand-traced goldens |
| Symbolic computation | Bounded rule execution | 9 rules, `EvaluateRiemann*`, `LeviCivita` |
| Numerical verification | NOT used (no numerics; `no_numerics` limitation) | — |
| Executable regression testing | `derivation_test.go` suites above | 19 + 9 + 21 tests listed |
| Mathematical proof | Human obligation; machine checks structure | Bianchi per-instance obligation, never universal |

## 3. Reproduction contract (independent reviewer)

Build under a workspace (`01 §2` recipe), run the fixture suite
(`go test -count=1 ./...` in the fixture + `protocol/`), read the
manifest for premises/limitations/falsifiers, and inspect goldens
against their hand derivations. `phystrace-mcp run_verification`
profiles (`demo-arc`, `coverage-audit`, `full-suite`) automate the
mechanical part; mathematical judgment stays human.
