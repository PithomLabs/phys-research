# 06 — Theorem contract (REVISED, normalization repair)

Audit: QSG-T001. Author: Agent A (repair). Date: 2026-11-06.

REVISED version. The v1 contract is preserved byte-identical as
`06-theorem-contract-v1-superseded.md` (SHA-256
`9d12c5ee6bc6a70896d24d5cad738fca80289af7d6c0a8ea2541347910d1eb76`).
Changes: repaired prescription (doc 08 N1), remainder (N2),
corrected residual/correction/trace (N4.5/N4.8), TI correction
(N4.2), v1-transport (N4.3), refined O-register (N5.4), re-answered
gate, revised Phys-Pseudo, extended cross-check record.
Unchanged: frozen model, \(H^{(1)}\), point-split operator,
ambiguity basis (3.3), classical core, S-independent findings.
Also records the `01`-(67) DF-transcription sign correction
(N4.1); `01` itself stays byte-identical.

Dispositions follow `01`. `PASS` is not used.

---

## 1. Theorem (strongest defensible proposition)

Frozen model: four-dimensional smooth globally hyperbolic Lorentzian
\((M,g)\), signature \((-+++)\), \(c=\hbar=1\); real scalar field, mass
\(m>0\), minimal coupling \(\xi=0\); Hadamard state \(\omega\); fixed
classical background (no backreaction).

Definitions (doc 08 N1): \(G^{(1)}=G^++G^-\);
\(H^{(1)}=(1/8\pi^2)[U/\sigma+V\ln|\bar\sigma|]\), \(U=\Delta^{1/2}\);
\(S=\tfrac12G^{(1)}-H^{(1)}\) (Alternative A, N1.6);
\(T_{\mu\nu}(x,x')\) the symmetrized point-split operator (0.10);
\(R_{\mu\nu}=\lim_{x'\to x}T_{\mu\nu}S\);
\(C_{\mu\nu}=+g_{\mu\nu}[v_1]/(4\pi^2)\) (N4.6);
\(\Theta_{\mu\nu}\) conserved local geometric, dimension 4.

QSG-T001 (conservation): with
\(\langle T_{\mu\nu}\rangle_{\rm ren}=R_{\mu\nu}+C_{\mu\nu}
+\Theta_{\mu\nu}\),
\[
\boxed{\nabla^\mu\langle T_{\mu\nu}\rangle_{\rm ren}=0.}
\tag{T}
\]
Proof chain (doc 08): prescription (N1.6) \(\to\) remainder
equation (N2.1/N2.2) \(\to\) residual \([\nabla'J]\) (N4.2, two
routes) \(\to\) \(1/3\)-split (N4.4) \(\to\) residual (N4.5)
\(\to\) correction (N4.6) \(\to\) cancellation (N4.7). The
correction is not independently conserved (N3.4); the trace is
generically nonzero (N4.8) and compatible with (T); DF's numbers
are matched by derivation, not import (N4.1/N4.5).

Scope note: fixed-background free-field conservation theorem.
No backreaction, existence, uniqueness, stability, or stochastic
extension (see §3). Explicit-\(A\) framework identification is
CONDITIONAL (O2 refined); the internal theorem is exact.

## 2. Assumptions register

- A1 (P1): \(G^+\) satisfies \((\Box-m^2)G^+=0\) in each argument
  (free field). Premise of frozen model.
- A2 (P2-repaired): \(S\) is \(C^4\) (Hadamard state; N2.1 shows
  the premise coherent with the verified singular match).
- A3 (DP): Hadamard recursion closes (smooth \((\Box-m^2)H^{(1)}\));
  minimal-\(U\) representative (doc 03 §1.3, R5); exact
  \(v_0\)-transport (N2.3) and \(v_1\)-transport (N4.3); coefficient
  uniqueness standard.
- A4 (geometry): L1 (Van Vleck transport), G2
  (\([(\Box-m^2)\Delta^{1/2}]=R/6-m^2\)), L4 (contracted Bianchi),
  Palatini identity, Lovelock (Gauss–Bonnet), Synge rule, F1
  (\([\nabla g]=0\), DF-(66b)-verified), F3
  (\([\nabla\tau]=[\nabla'\tau]=0\), derived N4.4); L3 derived;
  L2 DERIVED (N2.4). Pure background + standard theory.
- A5 (scale): fixed renormalization length \(\ell\); changes absorbed
  into \(\Theta\) (re-proved via v1-transport, N5.3-D5).
- A6 (ambiguity scope): polynomial dimension-four basis (3.3).
  CONDITIONAL on Salcedo's renormalizability criterion; the
  Tichy–Flanagan non-polynomial extension is acknowledged scope
  beyond (doc 04 §4.3, `01` §§4.5–4.6); uniqueness layers N3.3.
- A7 (state class): Hadamard states only. Non-Hadamard states excluded.
- A8 (no backreaction): background fixed; SCEE coupling excluded.
- A9 (F2-weak): \([\nabla v_1]=[\nabla'v_1]\) (v1 first-derivative
  coincidence halves agree). SOURCE-REPORTED standard Hadamard
  theory (symmetric \(V\)-coefficients); only this weak form is
  consumed (N4.4).

## 3. Exclusions (not claimed)

Massless/conformal case (\(m=0\), \(\xi=1/6\)); non-minimal coupling;
interacting/perturbative conservation (Hollands–Wald, source-reported);
semiclassical backreaction, FLRW existence/uniqueness/stability;
stochastic gravity/noise kernel; Bohmian mechanics; explicit curved-space
\([v_1]\) polynomial (formal \([v_1]\) suffices for (T)); explicit
framework-difference term \(A\) (CONDITIONAL on O2-refined; A-data
\((0,0)\) derived, N4.5).

## 4. Proof obligations register

- O1. INTERNAL LEVELS CLOSED, EXTERNAL OPEN (level split per
  `plan6_imp_review.md` §4; no new QSG cycle per plan7). O1a
  (remainder limit + \(v_0\)-transport + L2): executor N2 +
  checker #1 (a–c,e) agree exactly, both blind. O1b-core
  (\(\delta\)/\(J'\)-structure, \([\nabla'J]=\tfrac13\nabla[J]\)):
  checker #1 (d, Synge-A) + executor N4.2/N4.4 (revised-T1/TI +
  Q1-split, different route) agree. O1b-transport (v1-transport
  N4.3): executor-derived + checker #2's bounded blind
  re-derivation RETURNED full agreement (TI-correction,
  \(R_0\equiv0\), F3-from-identity, Q1-split, \(1/(4\pi^2)\);
  all flags closed — doc 08 N2.6). Verification levels:
  in-execution derivation CLOSED; adversarial internal
  cross-check CLOSED (same-session); independent external
  replication OPEN; separate-session reproduction OPEN. Theorem
  status: provisionally accepted derived human result under the
  declared prescription. Blocks nothing.
- O2. Explicit-\(A\) determination: OPEN (refined). Blocks ONLY
  full DF framework identification (claim QSG-A-TERM), NOT the
  internal theorem: (N4.7) is exact in N1 conventions with
  DF-matched numbers, and A-data \(([(\Box-m^2)A],[\nabla'J_A])
  =(0,0)\) is derived (N4.5).
- O3. \(H^{(2)}\) full nonlinear \(\nabla\)-check line-by-line: OPEN.
  Covered by linearized check + Theorem V + external contrast;
  blocks shape-claims (QSG-SHAPES) only.
- O4. Wald 1977/1978 full texts: BLOCKED (paywalled). Not needed
  for (T). Blocks nothing.
- O5. Executable verification: repaired fixture (Step 7) with
  bounded scope (SK-records + boundary statements). Non-blocking
  for the mathematical theorem by design. Blocks nothing.

## 5. Acceptance gate (prompt4 §6, re-answered)

1. Classical identity independently reproduced? YES (`02` §§2–5,
   readback-verified, K7). DERIVED.
2. Remainder's actual field equation established? YES ((N2.1/N2.2)
   + falsification (N2.5); O1a closed). DERIVED.
3. Point-split divergence explicitly calculated? YES ((N4.2)
   \([\nabla'J]\), two routes: revised-T1/TI + Synge-A).
   DERIVED.
4. Residual local term identified? YES ((N4.5)). DERIVED.
5. Correction explicitly derived? YES ((N4.6), forced by residual
   + uniqueness layers N3.3). DERIVED.
6. Its divergence calculated? YES (N4.6 + non-conservation N3.4).
   DERIVED.
7. Cancellation under one consistent convention? YES ((N4.7), all
   N1). DERIVED.
8. Ambiguities correctly identified? YES polynomial basis (3.3);
   CONDITIONAL on A6 (+ N3.3 layers). DERIVED/CONDITIONAL.
9. Theorem appropriately restricted? YES (§§1–3, A9 explicit).
   DERIVED.
10. Primary sources sufficient? YES: DF chain re-inspected verbatim
    (N4.1); `01`-(67) sign corrected; DF numbers DERIVED-match
    (N4.5); Wald abstracts structural; Wald full texts BLOCKED but
    unneeded; explicit-\(A\) CONDITIONAL on O2 (non-blocking).
    DERIVED/CONDITIONAL.

Gate verdict: PASS (repaired) — proceed with repaired fixture
(§§6–7 evidence; readiness re-decided in doc 07 §7). O2/O3 open
as recorded (non-blocking for the internal theorem).

## 6. Phys-Pseudo representation (conforming, revised)

Checked against PSEUDO.md via the recon contract: seven constructs
only; `define` with `:=`; `derive` with non-empty `from` + `|`;
no argument lists; no nesting; multiline `if`; prose-only `?`;
no status words; no claim promotion. `?` marks exactly the open
obligations O2-refined and O3 (O1 closed, O4–O5 access/future-work
items stated in prose above). Premise lists are complete per
review §5 (no omitted-premise hiding: every transport/Synge/
curvature/J-premise is named).

```
given MODEL: four-dimensional smooth globally hyperbolic Lorentzian manifold signature (-+++) units c=hbar=1
given FIELD: real scalar field mass m>0 minimal coupling xi=0 fixed classical background no backreaction
given STATE: Hadamard state with two-point function satisfying Klein-Gordon in each argument
define SYMM2PT := sum of Wightman function and reversed Wightman function
define PARAMETRIX := Hadamard singular part with U=Delta^{1/2} and V series
define REMAINDER := half of SYMM2PT minus PARAMETRIX
define V1COINC := coincidence value [v1] a local scalar of mass dimension 4
define JSOURCE := (Box-m^2) applied to REMAINDER
define JPRIME := primed (Box-m^2) applied to REMAINDER
define PSOPER := symmetrized point-splitting stress-energy bi-operator
define RAWTENSOR := coincidence limit of PSOPER applied to REMAINDER
assume SMOOTHNESS: REMAINDER is smooth on a convex neighborhood of the diagonal
assume HADAMARDRECURSION: parametrix coefficients satisfy the transport equations making (Box-m^2)PARAMETRIX smooth
assume V0TRANSPORT: exact v0-transport equation killing the N1 pole part identically
assume V1TRANSPORT: exact v1-transport equation making the NL log term C1 with zero gradient
assume SYNGE: coincidence derivatives follow the transported-primed-derivative rule
assume JPRIMESWAP: primed-source JPRIME satisfies JPRIME(x,x')=JSOURCE(x',x)
assume RICCIID: commutator and contracted-Bianchi curvature identities
assume TAUDATA: first derivatives of the dAlembertian world-function remainder vanish at coincidence
assume VSYMMWEAK: v1 first-derivative coincidence halves agree
assume TRANSPORTVANISH: parallel-transport first derivatives vanish at coincidence
claim HOMOGENEOUS: JSOURCE vanishes identically
derive REMAINDERLIMIT: coincidence limit of JSOURCE equals -3 V1COINC/(4pi^2) from [SMOOTHNESS, HADAMARDRECURSION, V0TRANSPORT] | Leibniz on U/sigma and V ln terms with pole and log cancellation plus exact transport
derive REFUTEHOMOGENEOUS: HOMOGENEOUS is false from [REMAINDERLIMIT] | Minkowski value -3m^4/(32pi^2) is nonzero for m>0
derive SPLITTHIRD: primed-half coincidence derivative of JSOURCE equals one third of the total gradient from [REMAINDERLIMIT, V1TRANSPORT, VSYMMWEAK, TAUDATA] | Q1-split 4w over 12w with vanishing log-gradient
define RESIDUAL := divergence of RAWTENSOR
derive RESIDUALVALUE: RESIDUAL equals primed-half coincidence derivative and equals gradient of -V1COINC/(4pi^2) from [REMAINDERLIMIT, SPLITTHIRD, SMOOTHNESS, SYNGE, JPRIMESWAP, RICCIID, TRANSPORTVANISH] | corrected transport identity with J/J' distinction plus independent Synge-A route
define CORRECTION := (1/(4pi^2)) metric times V1COINC
derive CORRECTIONDIVERGENCE: divergence of CORRECTION equals gradient of +V1COINC/(4pi^2) from [RESIDUALVALUE] | metric compatibility with coefficient forced by cancellation
derive CORRECTIONNOTCONSERVED: CORRECTION is not independently conserved from [CORRECTIONDIVERGENCE] | its nonzero divergence is what cancels the residual
define AMBIGUITY := conserved local geometric tensor of mass dimension 4
derive AMBIGUITYBASIS: polynomial dimension-four ambiguities spanned by four tensors from [MODEL, FIELD] | metric variation of local scalars is conserved with Lovelock reduction
define RENTENSOR := RAWTENSOR + CORRECTION + AMBIGUITY
claim QSGCONSERVATION: divergence of RENTENSOR vanishes
derive QSGPROOF: QSGCONSERVATION holds from [RESIDUALVALUE, CORRECTIONDIVERGENCE] | opposite gradients cancel and AMBIGUITY is conserved by construction
claim NONZEROTRACE: trace of RENTENSOR is generically nonzero
derive TRACEVALUE: trace equals -(1/2)Box s - m^2 s + V1COINC/(4pi^2) + trace AMBIGUITY from [REMAINDERLIMIT, SMOOTHNESS] | operator trace plus correction trace per N4.8
claim POLYNOMIALSCOPE: finite ambiguities restricted to polynomial dimension-four basis
claim GENERALAMBIGUITY: non-polynomial ambiguities possible beyond renormalizability criterion
if POLYNOMIALSCOPE:
    then AMBIGUITYBASIS
    else GENERALAMBIGUITY
? Explicit smooth framework-difference term fully identifying the symmetric route with the Feynman route beyond the derived equation data
? Full nonlinear line-by-line check of H^(2) conservation beyond the linearized check and variation theorem
```

Conformance self-audit: one construct per line; names plain identifiers
declared before use, exact-match `from` lists; `define` `:=` throughout;
every `derive` carries result + premises + reasoning; `if` multiline with
`then`/`else`; `?` prose-only with what-would-resolve-it (O2-refined,
O3); no `PROVEN/VERIFIED/PASS/STATUS` labels; claims examined, only
endorsed via `derive` (QSGPROOF, REFUTEHOMOGENEOUS, TRACEVALUE,
SPLITTHIRD, RESIDUALVALUE) — representation, not machine proof.
Phys-Pseudo expresses the argument; doc 08 proves it.

## 7. Cross-check record (repair: checkers #1 and #2)

v1 record (same-session subagent, remainder-limit-as-premise,
residual→correction→cancellation + \(H^{(1)}\) + E1-catch):
carried as history in the superseded file; its J-symmetric
framework is corrected by N4.2.

Checker #1 (normalization repair): quarantined context — frozen
model, N1 prescription, DP, Synge data, L1/G2, operator, notation;
explicitly forbidden audit/plan reads (compliance confirmed: none
read). Tasked derive-don't-receive: remainder equation + \([J]\),
transport closure + L2-content, coincidence calculus, divergence
structure, flat check. Returned: \([J]=-3[v_1]/(4\pi^2)\) (integer
\(6=4+2\)), exact \(v_0\)-transport with \(M\equiv0\), E1-form
\(\Box s\), flat values — ALL AGREEING with N2 (O1a closed).
Adversarial finding: via independent Synge-A route, residual
\([\nabla'J]=\tfrac13\nabla[J]\) (refuting the executor's interim
\(1/2\)), with flags F1 ([\nabla g]=0), F2 (v1-halves), F3
([\nabla\tau]=[\nabla'\tau]=0), F4 (exact \(R_0\equiv0\) needed).
Executor verified the refutation (N4.2 J/J', two routes) and
closed F1 (DF-66b), F3 (derived), F4 (N4.3 v1-transport); F2-weak
stands source-reported (A9). O1b-core closed by the two routes.

Checker #2 (residual-factor closure): bounded blind brief (premises
N2-results-taken-as-given + J/J'-distinction instruction; tasks:
TI-correction, v1-transport, F3, Q1-split, 1/3-combination).
RETURNED full agreement: residual \([\nabla'J]\) via third route
(B/C bitensor divergence), \(R_0\equiv0\) with vanishing
log-gradients, \([\nabla\tau]=[\nabla'\tau]=0\) derived from the
frozen \(s^{;u}s_{;u}=2s\) identity, Q1-split \(8w/4w/12w\),
residual \(-\nabla[v_1]/(4\pi^2)\), correction \(+1/(4\pi^2)\);
all nine flags closed (doc 08 N2.6). O1b-transport CLOSED.

Limitation: same-session subagents (two) plus executor
cross-route verification; stronger than v1's single check but
still not fully independent replication. Recorded honestly.

## 8. What this document establishes

Revised theorem (T) with repaired proof chain (N1–N5),
assumption register A1–A9, exclusions, refined obligations
(O1 closed, O2-refined/O3 open non-blocking, O4–O5 unblocking),
gate PASS (repaired, 10/10), conforming revised Phys-Pseudo with
complete premises, and the two-checker cross-check record. The v1
file is preserved byte-identical; `01`'s (67)-sign is corrected
here and in doc 08 while `01` stays byte-identical. Carries into
doc 07 §7: gate verdict + O-register + checker limitation for the
re-readiness decision.
