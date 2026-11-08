# 05 — Agent B findings

Audit: QSG-T001. Author: Agent A (resumption). Date: 2026-11-06.

Concise adversarial defect table for Agent B's `Quantum Stress Energy
Conservation Research.pdf` (24 pages), with corrections and mathematical
dispositions. Reference results are the blind derivations of docs 03–04,
the readback-verified records of docs 01–02, and the DF full text inspected
in doc 03 §3.

Evidence basis (this session, post-lock): rendered-page reads of PDF pp. 1,
2, 3, 10, 11, 12, 13, 14, 19, 20; text extraction of pp. 10–20 and 22–24;
prior-session extraction index (session file) used for navigation only.

Entry format: Location — Assertion — Correct result / reason — Missing
step — Minimum correction — Disposition.

Dispositions follow `01`. `PASS` is not used.

---

## A. Ledger and bibliography (pp. 1–3, 24)

- A1. p. 1, ledger row Wald: "Phys. Rev. D 17, 1477" — Correct:
  1477–1484. Missing: page range. Correction: complete the range. —
  FALSIFIED (bibliographic, minor).
- A2. p. 1, ledger row Wald: "four fundamental axioms" — Correct: Wald
  1977 lists five conditions (abstract verified, `01` §4.4). Missing: axiom
  5 (controls the \(\Box R\) ambiguity). Correction: five conditions. —
  FALSIFIED.
- A3. pp. 1–3, ledger status column: "INDEPENDENTLY REPRODUCED" for Wald,
  Radzikowski, Hollands–Wald, Décanini–Folacci — No reproduction of any of
  these results appears in the paper; the DF-derived formulas attempted are
  incorrect (C5–C7). Missing: every derivation so labeled. Correction:
  relabel SOURCE-REPORTED (or derive). — FALSIFIED (all four labels).
- A4. p. 2: Radzikowski "179, 529", H&W "223, 289" / "17, 227" —
  Correct: 179, 529–553; 223, 289–326; 17, 227–312. Correction: complete
  ranges. — FALSIFIED (bibliographic, minor).
- A5. p. 24, Works cited (31 entries): no Wald 1977/1978 entry, no
  Radzikowski 1996 entry — Correct: both are load-bearing primaries.
  Correction: add them. — FALSIFIED (missing primaries; confirms `01` §6).
- A6. Citation keys: Wald claims keyed to [1] (Salcedo), conservation to
  [7] (self-similar spacetimes, arXiv:2509.10360), μSC to [12]
  (arXiv:1911.01829) — Correct: keys must point at the works actually
  supporting the claims. Correction: re-key (Wald→Wald entries;
  conservation→DF/Wald; μSC→Radzikowski). — FALSIFIED (misattribution;
  confirms `01` §6; bibliography verified by extraction this session).

## B. Classical reconstruction (pp. 4–6)

Carried from `02` §8 (readback-verified; non-load-bearing; pp. 4–6 not
re-rendered): final results (EOM, \(T_{\mu\nu}\), off-shell identity,
on-shell conservation, Noether argument) are correct —
INDEPENDENTLY REPRODUCED. Intermediate defects (p. 4 canonical-momentum
sign; p. 4 metric-variation intermediate; p. 5 Hessian-justification
wording; p. 6 symmetry-statement blank) — FALSIFIED as written, all
non-load-bearing. No re-verification against renders (recorded choice;
classical foundation independently re-derived in full in `02`).

## C. Hadamard setup (p. 10)

- C1. p. 10, \(H(x,x')\) defining formula: fraction numerators and
  denominators are missing glyphs in the PDF (illegible at 150 and 300
  DPI; text extraction also blanks it; prior session confirms). The
  normalization cannot be verified from p. 10. — BLOCKED (document
  defect). Resolved for substance by C2a (p. 19 fixes intent).
- C2. p. 10, \(W\equiv G^+-H\) "smooth regularized biscalar", with
  \(\langle T\rangle_{\rm reg}=\lim TW\) — Correct: with Agent B's own
  \(H=(1/8\pi^2)[U/\sigma+V\ln]\) (p. 19 pseudo-code, §I; the
  \(G^{(1)}\)-parametrix),
  \(W=G^+-H=S-\tfrac12H^{(1)}+(i/2)\Delta+\)(\(i\epsilon\)-terms):
  singular (residual \(-\tfrac12H^{(1)}\) pole+log), complex,
  non-symmetric. So \(\lim TW\) diverges; the "regularized" tensor is not
  regularized. Missing: symmetric function \((1/2)G^{(1)}\); the
  \(i\Delta/2\) term; any normalization justification. Minimum correction:
  subtract from \((1/2)G^{(1)}\) per (0.5)/(0.11), or halve \(H\) with full
  re-derivation. — FALSIFIED (prescription; load-bearing — steps 1–5 all
  assume smooth \(W\)). Note: under the alternative half-\(H\) reading,
  \([V_0]\) would have to be half-standard; it is not (C5) — both readings
  defective.
- C3. p. 10, \({\cal T}_{\mu\nu}(x,x')\) operator — Matches blind (0.10)
  exactly. — INDEPENDENTLY REPRODUCED (Agent B got this right).
- C4. p. 10, \([U]=1\) — Matches \([\Delta^{1/2}]=1\). —
  INDEPENDENTLY REPRODUCED.
- C5. p. 10, \([V_0]=\tfrac12(\tfrac16R-m^2)+\tfrac12m^2=\tfrac1{12}R\) —
  Correct: \([v_0]=m^2/2-R/12\) (blind (V0); DF (108a)). Agent B's \(m^2\)
  cancels spuriously and the \(R\)-sign is flipped. Missing: correct
  recursion/conversion step. Correction: adopt (V0). — FALSIFIED.
- C6. p. 10, \([V_1]=\tfrac18m^4-\tfrac1{24}m^2R+\tfrac12a_2\) —
  Correct (DF (109), \(\xi=0\), inspected): curvature part
  \((1/720)R_{abcd}^2-(1/720)R_{ab}^2+(1/288)R^2+(1/120)\Box R\).
  Agent B's \(\tfrac12a_2\) gives twice DF's Riemann/Ricci/\(R^2\)
  coefficients with a wrong-sign \(\Box R\), even before C7. Missing:
  correct \(V_1\leftrightarrow a_2\) conversion. Correction: adopt DF
  (109). — FALSIFIED.
- C7. p. 10, \(a_2=(1/180)R_{abcd}^2-(1/180)R_{ab}^2+(1/72)R^2
  -(1/30)\Box R\) — Correct minimal-coupling coefficient of \(\Box R\) is
  \(+1/36\) (\(-\tfrac1{180}+\tfrac16\cdot\tfrac15\); Riemann/Ricci/\(R^2\)
  terms agree). Correction: \(+1/36\,\Box R\). — FALSIFIED (\(\Box R\)
  sign and value).

## D. Ambiguities (p. 11; rendered-page evidence)

- D1 (=B3). "\(c_1..c_4\) real dimensionless" for
  \(c_1g+c_2G+c_3H^{(1)}+c_4H^{(2)}\) — Correct: \([c_1]=4\), \([c_2]=2\)
  (blind (3.3) carries \(m^4\), \(m^2\)). Correction: insert mass factors.
  — FALSIFIED (dimensional).
- D2 (=B1). \(H^{(1)}=2\nabla\nabla R-2g\Box R+2RR-\tfrac12gR^2\) —
  Correct: blind (3.1) (derivative signs opposite); Agent B's version has
  \(\nabla^\mu B^{(1)}_{\mu\nu}=+4R_\nu{}^\sigma\nabla_\sigma R\ne0\),
  contradicting its "conserved" label. Correction: adopt (3.1). —
  FALSIFIED (sign; load-bearing for §G).
- D3 (=B2). \(H^{(2)}=\nabla\nabla R-g\Box R-\Box R_{\mu\nu}
  +2R^{\alpha\beta}R_{\alpha\mu\beta\nu}-\tfrac12gR^2\) — Linearized
  divergence \(-\tfrac12\Box\partial_\nu R\ne0\) (blind (3.2) gives 0).
  Correction: adopt (3.2). — FALSIFIED (derivative structure).
- D4 (=B4). "Four axioms" — FALSIFIED (see A2).
- D5 (=B8). Scale shift "\(\propto[V_1]g_{\mu\nu}\)" — Correct: the shift
  \(\lim T\Delta S\) involves \(V\)'s full coincidence data (including
  \(v_0\)-Hessians with \(\nabla\nabla R\) pieces), not a pure trace;
  blind §3.4. The structural claim (shift is a conserved basis element)
  is INDEPENDENTLY REPRODUCED by §3.4; the specific form is FALSIFIED and
  the basis mapping unproven (CONDITIONAL as Agent B's claim).

## E. Conservation derivation (pp. 12–13; rendered-page evidence)

- E1. p. 12, step 1 (regularized form \(\lim TW\)): rests on smooth \(W\)
  — FALSIFIED as applied (C2: \(W\) singular; the limit diverges).
- E2. p. 12, step 2 (commutation \(\nabla\lim=\lim\nabla\)): valid for
  smooth objects; inapplicable to Agent B's singular \(W\). —
  FALSIFIED as applied (false premise). (The commutation principle
  itself is standard and used correctly in blind §1.5.)
- E3. p. 12, step 3 (derivative expansion): algebra essentially correct;
  drops the \((\nabla^\mu g_\nu{}^{\nu'})\) term while keeping
  \((\nabla^\mu g_\mu{}^{\mu'})\) — inconsistent bookkeeping, harmless in
  the limit (\([\nabla g]=0\)). — INDEPENDENTLY REPRODUCED (algebra),
  minor inconsistency noted.
- E4. p. 12, step 4: "Both \(G^+\) and \(H\) satisfy KG up to smooth
  terms absorbed into \(W\). Therefore \(W\) obeys the homogeneous
  equation identically: \((\Box_x-m^2)W=0\)" — FALSIFIED (central):
  non sequitur ("up to smooth terms" entails \((\Box-m^2)W=-(\Box-m^2)H
  \ne0\), not \(0\)); contradicts blind (1.8) (\([(\Box-m^2)S]
  =-3[v_1]/(8\pi^2)\ne0\), Minkowski counterexample). No minimum repair
  within Agent B's route — requires the §1 remainder computation.
- E5 (=B5). p. 13, \(m^2\)-cancellation-to-zero — FALSIFIED (uses E4;
  drops \(\tfrac12[\nabla_{\nu'}J]=\tfrac14\nabla_\nu[J]\ne0\)).
- E6 (=B6). p. 13, residual from "\(\nabla^\mu g_\mu{}^{\mu'}\)" —
  FALSIFIED (\([\nabla g]=0\); true residual from \(J\), blind (1.3)).
- E7 (=B7). p. 13, \(C^{(0)}_{\mu\nu}\) "local conserved curvature
  tensor" with \(\nabla^\mu\langle T\rangle_{\rm reg}
  =-\nabla^\mu C^{(0)}_{\mu\nu}\) cancelled by \(+\nabla^\mu C^{(0)}\) —
  FALSIFIED (logical contradiction: a conserved tensor cannot supply the
  cancelling divergence; cf. blind §2.2). No explicit formula ever given
  (unsubstantiated as stated). Minimum correction: adopt blind (2.1)
  with non-conservation explicit.

## F. Trace anomaly (p. 13; rendered-page evidence)

- F1 (=B9). \(\langle T^\mu{}_\mu\rangle=(1/16\pi^2)a_2
  =(1/2880\pi^2)[C^2-E_4+\tfrac13\Box R]\) — FALSIFIED: (i)
  \(C^2-E_4=2R_{ab}^2-\tfrac23R^2\), matching neither Wald 1978
  (\((2880\pi^2)^{-1}[C^2+R_{ab}^2-\tfrac13R^2]+c\Box R\)) nor DF (111);
  (ii) minimal-coupling \(a_2\) content used for the conformal case;
  (iii) inapplicable to the frozen minimal massive model, whose trace
  (3.4) carries state-dependent pieces.
- F2 (=B10). "Non-zero trace does not imply non-conservation" (prose) —
  statement INDEPENDENTLY REPRODUCED (blind §3.5); Agent B's argument
  FALSIFIED (E4–E7, F1).

## G. Backreaction (pp. 14–16; p. 14 rendered)

- G1. p. 14, SCEE \(G+\Lambda g+\alpha H^{(1)}+\beta H^{(2)}
  =8\pi G\langle T\rangle_{\rm ren}\) as written — Structural form
  standard; as written uses falsified \(H^{(1)}/H^{(2)}\) (D2–D3). —
  FALSIFIED as written (correction: insert (3.1)–(3.2)).
- G2. p. 14, "mathematically consistent set" via "\(\nabla\langle T\rangle
  =0\), \(\nabla G=0\), \(\nabla g=0\), \(\nabla H^{(1,2)}=0\)" —
  FALSIFIED as Agent B's argument: first premise unproven by Agent B
  (E4–E7), last false for Agent B's tensors
  (\(\nabla B^{(1)}=4R\nabla R\ne0\); \(\nabla B^{(2)}\) fails
  linearized). (True with blind (2.3) + correct \(H\)-tensors — not
  Agent B's result.)
- G3. p. 14, "conservation guaranteed by general covariance" —
  FALSIFIED: covariance does not imply conservation (the raw covariant
  tensor (1.1) has divergence (1.3)); true justification is (2.3).
- G4. p. 14, modified Friedmann \(3H^2=\Lambda+8\pi G\rho_{\rm ren}
  -16\pi G[3\alpha(2\ddot HH-\dot H^2+6\dot HH^2)
  +\beta(\ddot HH-\tfrac12\dot H^2+3\dot HH^2)]\) — FALSIFIED:
  dimensionally inconsistent (stray \(16\pi G\): \(16\pi G\cdot\alpha
  \cdot X\) has dimension \(-4\ne2\)); \(\alpha\)-term matches neither
  Agent B's own \(H^{(1)}_{00}\) (\(-18(2H\ddot H+\dot H^2+10H^2\dot H)\))
  nor the correct one (\(+18(2\ddot HH-\dot H^2+6\dot HH^2)\),
  derived by FLRW reduction of (3.1)); \(\beta\)-term inherits D3 plus
  the same stray factor. Correction: \(3H^2=\Lambda+8\pi G\rho_{\rm ren}
  -\alpha H^{(1)}_{00}-\beta H^{(2)}_{00}\) with (3.1)–(3.2).
- G5. pp. 15–16, FLRW "operational flow" steps 1–5 (generic prose) —
  No independent defect as generic description. — INDEPENDENTLY
  REPRODUCED (generic); specifics inherit C–E defects.
- G6. pp. 15–16 + table item 10, Pinamonti/Meda/Siemssen existence
  summary ("local existence/uniqueness via Banach fixed-point; global
  needs state constraints"), labeled CONDITIONAL — Honest labeling;
  carried SOURCE-REPORTED (real results, not verified here, no
  derivation claimed). — No defect.

## H. Noise, validity, obstruction matrix (pp. 16–19)

Summaries (noise-kernel definition/conservation, Kuo–Ford criterion,
limitations, Einstein–Langevin, items 10–11 CONDITIONAL): standard
material, no derivation shown, honestly labeled where labeled — carried
SOURCE-REPORTED, no independent defect as summaries. Verdicts:
- H1. Table item 7 "Covariant Conservation: PASS — DERIVED from Hadamard
  KG wave equations" — FALSIFIED (E4–E7; the homogeneous premise is
  false).
- H2. Table item 9 "Constraint Compatibility: PASS" — FALSIFIED (rests
  on H1 + false \(\nabla H=0\)).
- H3. Table item 6 "Renormalization: PASS — eliminates divergences
  unambiguously" — FALSIFIED (C2: Agent B's subtraction leaves \(W\)
  singular; ambiguities §3.3 exist regardless).
- H4. Table item 5 "State Validity: PASS" — statement partially true
  (\(\mu\)SC \(\Leftrightarrow\) Hadamard, source-reported); PASS
  verdict FALSIFIED (no subtraction analysis; Agent B's subtraction
  defective per C2).
- H5. Table item 8 "Trace Anomalies: PASS" — statement (F2) true;
  verdict FALSIFIED (rests on wrong formula F1).
- H6. Table item 2 "Dimensional Compatibility: PASS" — FALSIFIED by D1
  (Agent B's own ambiguity violates dimensions).
- H7. Table item 1 "Mathematical Definitions: PASS — rigorously
  defined" — FALSIFIED (no operator-algebra/state-semantics definitions
  appear; verdict exceeds content).
- H8. Table items 3–4 (tensor structure, covariance PASS) —
  CONDITIONAL (true for a correct subtraction; Agent B's defective per
  C2).

## I. Phys-Pseudo (pp. 19–20; rendered-page evidence)

Checked against `phys-pseudo/PSEUDO.md` via the recon contract
(`02-phys-pseudo-contract.md`: seven constructs, exact syntax). Every
`derive` lacks `from [...]` and `|`; every `define` uses `=` for `:=`;
names carry argument lists; constructs nest; `?` takes a binding.
- I1. `given manifold M signature ...` (no colon; several per line) —
  NON-CONFORMING (PSEUDO §1–2: `given NAME: <content>`, one per line).
- I2. `define action S_m = ...`, `define eom = ...` (`=` for `:=`) —
  NON-CONFORMING (§1).
- I3. `G_plus(x, x_prime)`, `T_op(x, x_prime)`, `T_ren[g, omega]`
  (argument lists in names) — NON-CONFORMING (§2: plain identifiers).
- I4. All `derive ... :` without `from [NAMES] |` — NON-CONFORMING
  (§4: result + non-empty premise list + reasoning required).
- I5. `derive qsg_t001_conservation: assume ... assume ... derive
  step1: ... derive step2: ... claim QSG_T001_THEOREM: ...` (nesting) —
  NON-CONFORMING (no construct nests; a `derive` is one result + premises
  + reasoning).
- I6. Single-line `if ... then ... else ...` — NON-CONFORMING (§5:
  `if <cond>:` with `then`/`else` branch lines).
- I7. `? transition_to_stochastic_gravity: G_munu + ... = ...`
  (binding form) — NON-CONFORMING (§6: `?` takes unresolved-prose only).
- I8. `claim ... : true`, `claim QSG_T001_THEOREM` inside a derive
  (truth-endorsement; promotion by proximity) — NON-CONFORMING (§3, §7:
  claims never promoted).
- I9. `derive step1: (box_g-m^2)_x W = 0` — mathematically FALSE (E4),
  in addition to I4–I5. The pseudo-code asserts the falsified equation.
- I10. `define hadamard_parametrix H = (1/(8pi^2))*(...)`,
  `W = G_plus - H` (p. 19) — confirms C2 is Agent B's intent (not a
  rendering artifact): with the full \(1/(8\pi^2)\) parametrix
  subtracted from \(G^+\), \(W\) is singular. Cross-reference making C2
  definitive.
- I11. Stochastic content (`stochastic_noise_xi`, Einstein–Langevin,
  `transition_to_stochastic_gravity`) inside a QSG-T001 contract —
  scope defect (QSG-T001 is fixed-background conservation; no
  stochastic extension). Recorded, not a syntax verdict.
- I12. No `?` marks any open obligation (correction derivation, residual
  computation, \(H\)-tensor fixes) — representation hides obligations
  (review note, not a §-violation).

## J. Handoff and completion answers (pp. 20–24; extraction evidence)

- J1. Deliverables table ("01-ledger: INDEPENDENTLY REPRODUCED",
  "03-reconstruction: INDEPENDENTLY REPRODUCED", ...) — FALSIFIED (same
  basis as A3; the ledger's labels are unsupported and its formulas
  wrong).
- J2. "02-classical: DERIVED" — the classical derivation's finals are
  correct (`02` §8). — INDEPENDENTLY REPRODUCED (appropriate label).
- J3. Completion answer 2 ("conservation holds provided: Hadamard state,
  linear KG, local covariance") — FALSIFIED (incomplete: omits the
  conservation-restoring correction; (2.2)–(2.3) show Hadamard + KG +
  covariance leave the residual (1.3) uncancelled). This omission is the
  assumption-level root of E4–E7.
- J4. Completion answer 4 ("\(H^{(1)}\), \(H^{(2)}\) indispensable...")
  — FALSIFIED as written (inherits D2–D3; true with (3.1)–(3.2)).
- J5. Completion answer 7 (remaining obligations: higher-derivative
  coincidence limits in non-homogeneous spacetimes; state bounds) —
  Honest but INCOMPLETE (omits correction derivation, residual
  computation, \(H\)-tensor fixes). Not falsified; superseded by doc 06's
  obligation register.

---

## K. Dependency recheck (docs 01/02 → 03/04)

Per `plan5_review.md` §4: `01`/`02` frozen as records; every mathematical
dependency carried into docs 03–04 independently rechecked (not merely
accepted). Carried items only — nothing else from `01`/`02` is used.

- K1. `01` DF equation numbers/content as navigation — Re-verified
  against the ar5iv full text at equation level (doc 03 §§3.1–3.3:
  (23), (29a), (52), (55), (60), (61), (63), (67)–(70), (108a), (109)).
  — INDEPENDENTLY REPRODUCED. Navigation use sound.
- K2. `01` Wald-abstract structural claims — Carried SOURCE-REPORTED
  (readback only; full texts still BLOCKED). Mathematical dependence:
  structural comparison only (doc 04 §4.1). — No unrechecked load-bearing
  use.
- K3. `01` Tichy–Flanagan/Salcedo records — Carried SOURCE-REPORTED
  (readback only). Blind (3.3) derived without them; comparison only
  (doc 04 §4.3). — No unrechecked load-bearing use.
- K4. `01` \([v_0]\) formula — Re-derived blind (V0) from DP+G2; matched
  exactly. — INDEPENDENTLY REPRODUCED.
- K5. `01` \([v_1]\) explicit polynomial — NOT carried (stayed formal in
  blind work; flat value \(m^4/8\) re-derived in doc 03 §1.7). — No
  dependence.
- K6. `01` §8 conclusions (corrected-prescription shape, falsification
  verdicts) — QUARANTINED during blind stages (Locks 1–2 + quote checks
  §3/§4); used only post-lock for comparison. — No blind-stage use.
- K7. `02` action/EOM/\(T_{\mu\nu}\)/signs/off-shell identity —
  Rechecked by full readback (derivations §§2–5 verified line by line in
  execution Step 1, including flat-space, positivity, dimension, and
  Noether-cross checks). — INDEPENDENTLY REPRODUCED (recheck).
- K8. `02` (7.1) point-splitting operator — Rechecked V-02-71 (doc 03
  §0.4: explicit coincidence computation \(\lim TF=(3.4)\)). —
  DERIVED (recheck).
- K9. `02` §8 classical-defect records — Carried into §B above
  (readback; pp. 4–6 not re-rendered — recorded choice; non-load-bearing
  and `02` re-derives the classical core in full). — No load-bearing use.
- K10. Session-file scratch (prior extraction index, \(H\)-formula
  discussion) — Navigation only; zero mathematical dependence (all
  claims re-derived or re-inspected from primary evidence). — Clean.

Verdict: every carried dependency is either re-derived, re-inspected
against primary evidence, or carried as explicitly-marked
SOURCE-REPORTED with no load-bearing blind use. The Agent A/C provenance
requirement is satisfied: `01`/`02` served as records and navigation
aids; docs 03–04 stand on their own derivations.

---

## L. What Agent B got right (preserved contributions)

1. Classical foundation finals (EOM, \(T_{\mu\nu}\), off-shell/on-shell,
   Noether) — correct (§B).
2. Point-splitting operator \({\cal T}_{\mu\nu}\) — correct (C3).
3. \([U]=1\) — correct (C4).
4. Hadamard/\(\mu\)SC framework selection — correct strategy.
5. Coexistence statement (nonzero trace \(\not\Rightarrow\)
   non-conservation) — true (F2), though unproven by Agent B.
6. Scale-shift-is-conserved-ambiguity (structural) — true (D5), though
   Agent B's specific form is wrong.
7. Noise/backreaction/existence summaries as summaries (G5–G6, §H
   summaries) — standard material, honestly labeled where labeled
   CONDITIONAL/SOURCE-REPORTED.
8. Partial obligation acknowledgment (J5) — honest, incomplete.
9. The conservation CONCLUSION \(\nabla^\mu\langle T_{\mu\nu}\rangle_{\rm
   ren}=0\) is TRUE under the corrected prescription (blind (2.3)) —
   even though Agent B's derivation is invalid. The two questions stay
   separate: derivation FALSIFIED (§E), theorem DERIVED (doc 04 §2.3).

---

## M. Summary counts and carries

Counts (this document): FALSIFIED verdicts on every Agent B formula and
derivation step in the conservation chain (C2, C5–C7, D1–D4, E1, E4–E7,
F1, G1–G4, H1–H7, I1–I9, J1, J3–J4); INDEPENDENTLY REPRODUCED on classical
finals, the \(T\)-operator, \([U]\), generic summaries, and two structural
statements (D5-idea, F2-statement); CONDITIONAL on items 3–4 of the
obstruction matrix, Agent B's basis mapping, and the half-\(H\) reading
(closed by I10); BLOCKED on C1-rendering (substance resolved by I10).
Non-conforming: I1–I8 (every pseudo construct misused at least once).

Carries into doc 06: §L (preserved contributions) as the reuse list;
correction (2.1), basis (3.3), trace (3.4) as the theorem's content;
TF/Salcedo scope flag (doc 04 §4.3); open items \(A\) (doc 03 §3.2) and
nonlinear \(\nabla H^{(2)}\) (doc 04 §5) for the obligation register.

---

## Erratum E5 (normalization repair, 2026-11-06)

Neutral re-audit of every \(S\)-dependent entry (Amendment 3; full
record doc 08 N5, provisional register N5.1, four separate questions
N5.2). C2: FALSIFIED CONFIRMED with replaced reasoning
(\(W_B=S+(i/2)\Delta\), light-cone-singular + complex +
non-symmetric; real part \(=S\) smooth — answers (1)–(3)
separately). E1/E2/E4/E6/E7/H1/H3/H4: FALSIFIED CONFIRMED
(refutations re-pointed: E4 to N2.5, E6 to (N4.5), E7-minimum to
(N4.6)). E5: FALSIFIED CONFIRMED with corrected number
(drops \(\tfrac16\nabla[J]\), not \(\tfrac14\nabla[J]\)). F2/G2/G3/
H8/J3: CONFIRMED with re-pointed citations ((N4.5)–(N4.8)). L9:
UPGRADED — conclusion now DERIVED with DF-matched coefficient
(N4.7). D5-structural: CONFIRMED (re-proved via v1-transport);
specific form still FALSIFIED. K1: AMENDED (01-(67) sign corrected,
N4.1; DF chain re-verified verbatim). All other entries
(A–B, C1, C3–C7, D1–D4, F1, G1, G4–G6, H2, H5–H7, I1–I12,
J1–J2, J4–J5, K2–K10, L1–L8): S-INDEPENDENT on inspection —
verdicts STAND (I9/I10 re-point to N2.5/re-derived-C2). No
original line above is altered; this erratum is appended.
