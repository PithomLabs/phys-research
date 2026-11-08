# 04 — QSG mathematical feasibility

Method per adjustment A: every cell below rests on inspected types/rules,
not keyword counts. Inventories: all 34 `Family*` constants in
`phys-math/families.go` (full list read 2026-11-06 — no
quantum/state/two-point/renormalization/expectation family exists); all 9
`Rule*` IDs in `phys-math/rules.go`; AST tags in `ast.go`
(symbol/rational/add/mul/neg/pow/sqrt/call/relation/branch —
`TagCall` targets classical `function/1` artifacts only);
`operator.go` vocabulary = exactly
{differentiate, covariant_derivative}, only `differentiate` executes.

Columns: **R** = representable in PhysMath/PhysLib · **P** = expressible
in Phys-Pseudo · **V** = verifiable by the derivation protocol.
Scale: YES / BOUNDED / NO.

## 1. Object matrix (prompt2 §5 + §7 A–E)

| # | Object (§5) | R | P | V | Evidence / note |
|---|---|---|---|---|---|
| 1 | Classical metric, curvature | YES | YES | BOUNDED | `symmetric_bilinear/1`, `tensor/1` (sparse; general semantics deferred per families.go), `tensor_field/1`, `coordinate_system/1`, `affine_connection/1`; executable chart-bound `EvaluateRiemann`, `EvaluateRiemannContraction/Scalar/TraceReverse`, `LeviCivita`, `KroneckerSelector`, `EvaluateCovariantDerivative` (doc comments read: torsion never assumed, singular metric fails closed, replayable derivations). Covers §7E curvature side. |
| 2 | Scalar field, canonical vars | YES | YES | BOUNDED | `expression/1` over `function/1`-declared params; EL equation as `equation/1`. Deriving EL from the action needs functional variation → bounded `variation_reduction/1` (ansatz projection; minisuperspace `reduced_variation` premise) or human. §7A metric variation: same bound. |
| 3 | Quantum field operators | NO | YES | NO | No operator-algebra family; `OperatorApplication` is differential-operator machinery (`operator.go` inspected), not quantum operators. Closest honest form: c-number identities over `expression/1`. MISSING (representational). |
| 4 | States, two-point functions | BOUNDED | YES | NO | A two-point kernel as a plain two-argument `expression/1` is representable; state semantics (positivity, Hadamard, GNS) have no vocabulary. No executable state check. §7B regularity = assumption + pseudo. |
| 5 | Hadamard / state conditions | BOUNDED | YES | NO | As `PhysicalAssumption` description + pseudo `assume` (needs a new framework's category vocabulary — `gr`'s three categories are GR-closed per assumption.go). No engine check. |
| 6 | Renormalized stress-energy | BOUNDED | YES | BOUNDED | The renormalized *result expression* is representable (`expression/1`+`equation/1`); the *procedure* (regularization, subtraction, coincidence limits) has no rule support. Verification only via hand-derived goldens pinned in tests (minisuperspace precedent: `TestHistoricalWrongPath`, factored goldens). §7C. |
| 7 | Renormalization ambiguities | YES | YES | BOUNDED | Local curvature invariants computable (§7C `EvaluateRiemannScalar` etc.); ambiguity parametrization as `expression/1` YES. Which terms appear = human derivation. |
| 8 | Semiclassical Einstein eqs. | YES | YES | BOUNDED | As `equation_set/1` linking computable Einstein tensor to expectation expressions. No solver (substrate is bounded: `SELECT-BRANCH-001` single-pattern only; no integration/series/numerics). §7E existence/stability = human + per-instance checks. |
| 9 | Coupled matter-gravity eqs. | YES | YES | BOUNDED | Same as 8; time evolution is represent-only (precedent: minisuperspace `no_time_evolution` limitation). |
| 10 | Conservation / constraint identities | YES | YES | BOUNDED | Reusable pattern EXISTS: `obligation.go` contracted-Bianchi evaluated obligation — per-instance zero witnesses, each component with a replayed closing derivation to literal zero; replay never promotes to universal claim. QSG conservation must follow this pattern. Constraint propagation: no automation seen. §7D anomalies/geometric terms: representable as expressions; derivation human. |

## 2. The six hard probes (prompt2 §5, adjustment B)

| Probe | Verdict | Ground |
|---|---|---|
| Operator-valued quantities | NO / NO | No family, no rule; c-number shadows only. |
| Functional derivatives | BOUNDED (ansatz) / NO (general) | `variation_reduction/1` bounded projection exists; general δ/δφ unrepresented (minisuperspace `reduced_variation_bounded` limitation, read). |
| Renormalized composite operators | result YES / procedure NO | Engine cannot regularize/subtract; can pin results as goldens. |
| State-dependent expectation values | operation NO / result YES | No expectation rule; expectation *outputs* are ordinary expressions. |
| Curved-spacetime differential operators | BOUNDED-executable | Partials via `DIFFERENTIATE-001` + declaring function; covariant via chart-bound evaluators; wave operators assembled by hand composition (precedent: `(a∂_a)²` via `a·∂_a(a·∂_a·)` + anti-golden vs `a²∂²_a`). Single-chart, explicit `CoordParamMap` discipline. |
| Nonlocal kernels | represent YES / execute NO | `expression/1` or `integral/1` (families.go, geometry.go) — but no integration rule among the 9, so `integral/1` is representational only. |

## 3. What this means for QSG-T001 (§7 A–E coverage)

- **A. Matter foundations**: action/field eq/operator/metric variation —
  representable; classical EL + metric variation executable only in
  bounded/ansatz form; full functional calculus is a human proof
  obligation with per-instance machine witnesses.
- **B. Quantum formulation**: field equation as operator identity — NO;
  as c-number consequences (mode functions, kernels) — YES. State,
  two-point semantics, Hadamard — assumption + pseudo + hand derivation.
- **C. Renormalization**: prescription/ambiguities — assumption +
  expressions YES; procedure — NO; result identities — golden-pinned
  regression YES.
- **D. Conservation**: follows the Bianchi-obligation pattern —
  per-instance replayed witnesses YES; universal identity — human.
- **E. Backreaction**: coupled system representable; existence/stability
  — human; regression checks on pinned instances — YES.

Headline: the engine covers classical curved-spacetime computation
(chart-bound) and result-pinning very well; quantum semantics and
renormalization procedure are outside it. Nothing here requires
redesign — it requires assigning each QSG claim to the correct one of
the three §8 categories (done in `07`).
