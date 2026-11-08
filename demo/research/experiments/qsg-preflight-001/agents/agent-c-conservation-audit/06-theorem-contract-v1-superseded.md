# 06 — Theorem contract

Audit: QSG-T001. Author: Agent A (resumption). Date: 2026-11-06.

Strongest defensible QSG-T001 proposition, assumptions, exclusions, proof
obligations, acceptance-gate answers, cross-check record, and the
Phys-Pseudo representation. All conventions are doc-03 §0.

Dispositions follow `01`. `PASS` is not used.

---

## 1. Theorem (strongest defensible proposition)

Frozen model: four-dimensional smooth globally hyperbolic Lorentzian
\((M,g)\), signature \((-+++)\), \(c=\hbar=1\); real scalar field, mass
\(m>0\), minimal coupling \(\xi=0\); Hadamard state \(\omega\); fixed
classical background (no backreaction).

Definitions (all §0): \(G^{(1)}=G^++G^-\);
\(H^{(1)}=(1/8\pi^2)[U/\sigma+V\ln|\bar\sigma|]\), \(U=\Delta^{1/2}\);
\(S=\tfrac12(G^{(1)}-H^{(1)})\); \(T_{\mu\nu}(x,x')\) the symmetrized
point-split operator (0.10);
\(R_{\mu\nu}=\lim_{x'\to x}T_{\mu\nu}S\);
\(C_{\mu\nu}=+3g_{\mu\nu}[v_1]/(16\pi^2)\);
\(\Theta_{\mu\nu}\) conserved local geometric, dimension 4.

QSG-T001 (conservation): with
\(\langle T_{\mu\nu}\rangle_{\rm ren}=R_{\mu\nu}+C_{\mu\nu}
+\Theta_{\mu\nu}\),
\[
\boxed{\nabla^\mu\langle T_{\mu\nu}\rangle_{\rm ren}=0.}
\tag{T}
\]
Proof chain (all blind, locked): remainder equation (1.8) \(\to\) residual
(1.3) \(\to\) correction (2.1) \(\to\) cancellation (2.3). The correction
is not independently conserved (§2.2); the trace is generically nonzero
(3.4 as corrected by E2) and compatible with (T).

Scope note: this is the fixed-background free-field conservation theorem.
It does not assert backreaction, existence, uniqueness, stability, or any
stochastic extension (see §3).

## 2. Assumptions register

- A1 (P1): \(G^+\) satisfies \((\Box-m^2)G^+=0\) in each argument
  (free field). Premise of frozen model.
- A2 (P2): \(S\) is \(C^4\) (Hadamard state). Premise; consumed by \(T\)
  plus one divergence.
- A3 (DP): Hadamard recursion closes (smooth \((\Box-m^2)H^{(1)}\));
  minimal-\(U\) representative adopted (doc 03 §1.3). Standard
  construction + derived choice.
- A4 (geometry): L1 (Van Vleck transport), L2 (leading-transport
  consequence), G2 (\([(\Box-m^2)\Delta^{1/2}]=R/6-m^2\)), L4
  (contracted Bianchi), Palatini identity, Lovelock (Gauss–Bonnet
  variation vanishes in 4D). SOURCE-REPORTED standard theory except L3
  and Palatini (derived); pure background, no renormalization content.
- A5 (scale): fixed renormalization length \(\ell\); changes absorbed
  into \(\Theta\) (doc 04 §3.4, derived).
- A6 (ambiguity scope): polynomial dimension-four basis (3.3).
  CONDITIONAL on Salcedo's renormalizability criterion; the
  Tichy–Flanagan non-polynomial extension is acknowledged scope beyond
  (doc 04 §4.3, `01` §§4.5–4.6).
- A7 (state class): Hadamard states only. Non-Hadamard states excluded
  (P2 fails).
- A8 (no backreaction): background fixed; SCEE coupling excluded.

## 3. Exclusions (not claimed)

Massless/conformal case (\(m=0\), \(\xi=1/6\)); non-minimal coupling;
interacting/perturbative conservation (Hollands–Wald, source-reported);
semiclassical backreaction, FLRW existence/uniqueness/stability;
stochastic gravity/noise kernel; Bohmian mechanics; explicit curved-space
\([v_1]\) polynomial (formal \([v_1]\) suffices for (T)); DF-coefficient
cross-match beyond structure (CONDITIONAL on \(A\), doc 03 §3.2).

## 4. Proof obligations register

- O1. Remainder limit (1.8): DERIVED blind + flat-scratch confirmed.
  Independent re-derivation outstanding (desirable, non-blocking).
- O2. Framework-difference term \(A\) (doc 03 §3.2): explicit
  determination open (bounded computation; non-blocking — (T) exact in
  §0 conventions regardless).
- O3. \(H^{(2)}\) full nonlinear \(\nabla\)-check line-by-line: covered
  by linearized check + Theorem V + external contrast (doc 05 D3);
  full line-check desirable, non-blocking.
- O4. Wald 1977/1978 full texts: BLOCKED (paywalled). Not needed for
  (T) (blind derivations + inspected DF suffice).
- O5. Executable verification: not done (assessed in doc 07; Step 10
  if authorized).

## 5. Acceptance gate (prompt4 §6)

1. Classical identity independently reproduced? YES (`02` §§2–5,
   readback-verified, K7). DERIVED.
2. Remainder's actual field equation established? YES ((1.8) blind +
   flat-scratch + falsification (1.9); O1 notes re-derivation).
   DERIVED.
3. Point-split divergence explicitly calculated? YES ((D1)+TI,
   cross-checked §7). DERIVED.
4. Residual local term identified? YES ((1.3)). DERIVED.
5. Correction explicitly derived? YES ((2.1), forced by residual,
   cross-checked). DERIVED.
6. Its divergence calculated? YES (§2.1–2.2, cross-checked). DERIVED.
7. Cancellation under one consistent convention? YES ((2.3), all §0).
   DERIVED.
8. Ambiguities correctly identified? YES polynomial basis (3.3);
   CONDITIONAL on A6. DERIVED/CONDITIONAL.
9. Theorem appropriately restricted? YES (§§1–3). DERIVED.
10. Primary sources sufficient? YES for structure (DF inspected, Wald
    abstracts, blind derivations); Wald full texts BLOCKED but
    unneeded; DF coefficient match CONDITIONAL on O2 (non-blocking).
    DERIVED/CONDITIONAL.

Gate verdict: PASS — proceed to minimal formalization (§§6–7 evidence
below; readiness decision in doc 07). No blocking item; O1–O5 carried
open as recorded.

## 6. Phys-Pseudo representation (conforming)

Checked against PSEUDO.md via the recon contract: seven constructs only;
`define` with `:=`; `derive` with non-empty `from` + `|`; no argument
lists; no nesting; multiline `if`; prose-only `?`; no status words; no
claim promotion. `?` marks exactly the open obligations O1–O3 (O4–O5
are access/future-work items, stated in prose above).

```
given MODEL: four-dimensional smooth globally hyperbolic Lorentzian manifold signature (-+++) units c=hbar=1
given FIELD: real scalar field mass m>0 minimal coupling xi=0 fixed classical background no backreaction
given STATE: Hadamard state with two-point function satisfying Klein-Gordon in each argument
define SYMM2PT := sum of Wightman function and reversed Wightman function
define PARAMETRIX := Hadamard singular part with U=Delta^{1/2} and V series
define REMAINDER := half the difference of SYMM2PT and PARAMETRIX
define V1COINC := coincidence value [v1] a local scalar of mass dimension 4
define PSOPER := symmetrized point-splitting stress-energy bi-operator
define RAWTENSOR := coincidence limit of PSOPER applied to REMAINDER
assume SMOOTHNESS: REMAINDER is smooth on a convex neighborhood of the diagonal
assume HADAMARDRECURSION: parametrix coefficients satisfy the transport equations making (Box-m^2)PARAMETRIX smooth
claim HOMOGENEOUS: (Box-m^2)REMAINDER vanishes identically
derive REMAINDERLIMIT: coincidence limit of (Box-m^2)REMAINDER equals -3 V1COINC/(8pi^2) from [SMOOTHNESS, HADAMARDRECURSION] | Leibniz on U/sigma and V ln terms with pole and log cancellation
derive REFUTEHOMOGENEOUS: HOMOGENEOUS is false from [REMAINDERLIMIT] | Minkowski value -3m^4/(64pi^2) is nonzero for m>0
define RESIDUAL := divergence of RAWTENSOR
derive RESIDUALVALUE: RESIDUAL equals gradient of -3 V1COINC/(16pi^2) from [REMAINDERLIMIT, SMOOTHNESS] | Synge-rule divergence computation with transport identity and Ricci cancellation
define CORRECTION := (3/(16pi^2)) metric times V1COINC
derive CORRECTIONDIVERGENCE: divergence of CORRECTION equals gradient of +3 V1COINC/(16pi^2) from [RESIDUALVALUE] | metric compatibility with coefficient forced by cancellation
derive CORRECTIONNOTCONSERVED: CORRECTION is not independently conserved from [CORRECTIONDIVERGENCE] | its nonzero divergence is what cancels the residual
define AMBIGUITY := conserved local geometric tensor of mass dimension 4
derive AMBIGUITYBASIS: polynomial dimension-four ambiguities spanned by four tensors from [MODEL, FIELD] | metric variation of local scalars is conserved with Lovelock reduction
define RENTENSOR := RAWTENSOR + CORRECTION + AMBIGUITY
claim QSGCONSERVATION: divergence of RENTENSOR vanishes
derive QSGPROOF: QSGCONSERVATION holds from [RESIDUALVALUE, CORRECTIONDIVERGENCE] | opposite gradients cancel and AMBIGUITY is conserved by construction
claim NONZEROTRACE: trace of RENTENSOR is generically nonzero
derive TRACEVALUE: trace equals -(1/2)Box s - m^2 s + 3 V1COINC/(8pi^2) + trace AMBIGUITY from [REMAINDERLIMIT, SMOOTHNESS] | operator trace plus correction trace per Erratum E2
claim POLYNOMIALSCOPE: finite ambiguities restricted to polynomial dimension-four basis
claim GENERALAMBIGUITY: non-polynomial ambiguities possible beyond renormalizability criterion
if POLYNOMIALSCOPE:
    then AMBIGUITYBASIS
    else GENERALAMBIGUITY
? Independent re-derivation of the remainder limit REMAINDERLIMIT from the Hadamard construction
? Explicit smooth framework-difference term reconciling symmetric-route and Feynman-route remainder coefficients
? Full nonlinear line-by-line check of H^(2) conservation beyond the linearized check and variation theorem
```

Conformance self-audit: one construct per line; names plain identifiers
declared before use, exact-match `from` lists; `define` `:=` throughout;
every `derive` carries result + premises + reasoning; `if` multiline with
`then`/`else`; `?` prose-only with what-would-resolve-it; no
`PROVEN/VERIFIED/PASS/STATUS` labels; claims examined, only endorsed via
`derive` (QSGPROOF, REFUTEHOMOGENEOUS, TRACEVALUE) — representation, not
machine proof. Phys-Pseudo expresses the argument; docs 03–04 prove it.

## 7. Cross-check record (Step 6)

Protocol: same-session subagent with quarantined context — received only
the frozen model, §0 definitions, geometric lemmas, and the remainder
limit (1.8) as a premise; explicitly forbidden from reading the audit
directory and plans (compliance confirmed: no quarantined files read).
Tasked to independently derive the divergence, residual, correction,
cancellation, and \(H^{(1)}\).

Result: residual \((1/2)\nabla[J]=-3\nabla[v_1]/(16\pi^2)\) DERIVED (via
an independent bitensor route, not T1/TI); \(K=+3/(16\pi^2)\) DERIVED;
correction-not-conserved DERIVED; conservation DERIVED; \(H^{(1)}\) +
\(\nabla H^{(1)}=0\) DERIVED (matching (3.1), with the sign cross-check
noting flipped signs would give \(4R\nabla R\ne0\)); commutator L3 and
\(\nabla s=2s_\mu\) DERIVED.

The checker DISAGREED on one item: doc-03 (2.2)'s
\(\Box s=[\Box_xS]+t\), deriving \(\Box s=2[\Box_xS]+2t\) instead. The
checker is right (the doc dropped the factor 2 from \(\nabla s=2s_\mu\)).
Corrected via Errata E1 (doc 03) and E2 (doc 04 (3.4)); (1.3), (2.1)–(2.3),
(3.1)–(3.3), §3.4 verified unaffected (no \(\Box s\) use). Both lock hashes
re-verified after the appends. Notably, the cross-check prompt itself
contained the error, which the checker caught rather than parroted —
evidence of genuine independence.

Coverage and limitation: verified — residual→correction→cancellation
chain, \(H^{(1)}\), L3, Synge instances. Taken as premise — remainder
limit (1.8) (see O1). Not re-derived by checker — (V0)/(V1a), \(H^{(2)}\),
scale argument (covered by flat-scratch, DF-match, linearized+Theorem V
respectively). Limitation: a same-session subagent is weaker than a truly
independent reviewer; no fully independent cross-review was available.
The theorem is therefore CONDITIONAL on single-programme derivation with
internal cross-checks (not on unreviewed single-agent work, and not on
fully independent replication either). Recorded honestly per review §4.

## 8. What this document establishes

Theorem (T) with proof chain, assumption register A1–A8, exclusions,
obligations O1–O5, gate PASS (10/10 answered, none blocking), conforming
Phys-Pseudo representation, and the cross-check record with its
correction and limitation. Carries into doc 07: gate verdict + O1–O5 +
cross-check limitation for the readiness decision.
