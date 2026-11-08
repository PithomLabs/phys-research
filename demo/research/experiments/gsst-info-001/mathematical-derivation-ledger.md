# Mathematical Derivation Ledger

## D-001 — FPST [32] scope adjudication (H6 crux) [SELF-CHECK, unverified]

Sources opened: P-003 text lines 270–300, 650–690, 915–935, 1085–1130, 310–330.
Findings:
1. FPST assumes "the matter fields are free massless scalar fields" (line ~1090) for the density-matrix/entropy construction. GSST's matter is a holographic CFT, NOT free scalars.
2. FPST's coherent-state assumption serves its GSL proof + free-field entropy computation (lines ~276–290); exotic states would violate their GSL.
3. GSST cites [32] ONLY for "S_vN(coherent) = S_vN(vacuum)" to justify RT evaluation in vacuum coordinates (ledger C-008/B1).
CONCLUSION (self-check): [32] licenses S_vN-equality for entropy evaluation ONLY. It does NOT promote t± + coherent data to full radiation states on a common algebra for GSST's holographic CFT. Additional gap: GSST applies a free-field entropy-equality result to holographic-CFT matter — assumed, not shown in [32] (already carried as B1 conditional).
Effect: H6's cheapest-test crux resolves TOWARD H6 (core gap confirmed); H6 itself still needs the underdetermination construction or missing-lemma cut. AWAITING independent verification (spawn V-001).

## D-002 — H5 δ/ρ bookkeeping audit (cheapest test) [SELF-CHECK, unverified]

Restored-δ form of (4.7): S_bulk = (c/6)[log d² + ρ(A) + ρ(I) − 2log δ] (from (4.6) z=δe^{−ρ} + BH c=3L₃/2G₃; algebra shown in orchestrator notes).
M-derivative at fixed τ: (c/6)[d(log d²)/dM + dρ(A)/dM + dρ(I)/dM] − (c/3)(1/δ)(dδ/dM).
Results:
1. With M-independent δ: δ-term drops out of d/dM; M-dependence flows ONLY through d, ρ(A), ρ(I) (visible, via B2 substitution + QES shift). δ-bookkeeping SAFE for derivatives.
2. WITHOUT fixed δ (absolute Q): ∂S/∂δ = −(c/3)(1/δ) diverges as δ→0 ≫ O(1) mass signal — absolute-Q mass claims are regulator-dominated. H5 mechanism CONFIRMED in this direction.
3. δ(M) M-dependent cutoff is nowhere forbidden by GSST — would inject arbitrary M-dependence via −(c/3)log δ(M). Hence a STATED M-independent-δ prescription (or difference quantities at common δ) is NECESSARY for any regulator-independent mass signal.
4. Area term 2(Ω−Ω_crit): Ω_crit M-independent (c only) → d/dM = 2(dΩ/dM), no δ anywhere. Regulator-clean. ✓
5. Relative/difference quantities at common fixed δ: δ cancels — admissible form for H1–H4 signals.
CONCLUSION (self-check): H5 SUBSTANTIATED as a constraint (not standalone discovery): H1–H4 signals MUST be formulated as difference quantities with stated M-independent δ; absolute-Q mass signals are inadmissible. AWAITING independent verification (spawn V-001).
VERIFICATION (V-001, complete:true): all sub-items CORRECT; hidden assumptions surfaced (c,δ M-independent; differentiability; no UV/IR mixing; fixed-τ/frame disambiguated; GSST-vs-FPST Ω_crit distinction). → C-011 [R]. H5 SUPPORTED-AS-CONSTRAINT.
ADVERSARIAL REPAIR (ADV-001 wound 2a-ii): the "∂S/∂δ→∞ ≫ O(1) signal" comparison in D-002 item 2 is dimensionally mixed and is WITHDRAWN as argument. REPLACEMENT (finite-scheme-change formulation): an O(1) scheme rescaling δ→e^α·δ shifts S by −(c/3)α=O(c), which dwarfs the O(1) per-unit-M island signal (dS_isl/dM=2) since c≫1; an unstated M-dependent δ(M) injects arbitrary −(c/3)log δ(M). The δ-domination CONCLUSION is unchanged and strengthened (O(c) vs O(1)); only the presentation is repaired. Algebra itself re-derived independently by ADV-001 from GSST (4.6)+BH relation (MISS on 2a-i).

## D-003 — H4 RC3 mapping tables (Zhong + JLMS → GSST) [VERIFIED]

Full tables: `h4-rc3-tables.md` (this ledger records the derivation; the tables are the artifact).
Table Z (P-004 → GSST): Z1 no tensor factorization (type III) FAIL; Z2 no code/isometry FAIL; Z3 no finite-DOF BH space + unitarity scope-forbidden FAIL; Z4 type I/II only, type-III extension unproved (P-004 §5) FAIL; Z5 Araki quantity exists but criterion unproved PARTIAL. Mapping verdict BLOCKED → C-012 [R] (V-002).
Table J (P-007 → GSST, corrected): J1 wrong asymptotics FAIL; J2 no reference-σ/nearby-family FAIL (JLMS has no literal H_code axiom — V-003); J3 island outside matter-duality scope FAIL; J4 no boundary/bulk map FAIL; J5 EWR is JLMS implication, not premise (direction corrected) FAIL. Mapping verdict BLOCKED → C-013 [R] (V-002 verdicts + V-003 premises).
Correction history: J-source first misattributed to 1503.06237 (HaPPY, now P-006); caught by V-002, repaired by fetching P-007 + V-003 re-verification. No conclusion rested on the misattribution (H4 stayed proposed throughout).

## D-004 — H6 missing-lemma cut (H6 cycle 2) [VERIFIED]

Full statement: `cycle-screens.md` D-004. Exhaustive audit of ledger-v1.2 state-relevant inputs: (a) c-number functions, (b) single numbers, (c) [32] S_vN-license, (d) RT geometry→number, (e) QES saddle-selection rule. NOWHERE: algebra A_R, state ω_R, Hilbert space, or input→state construction rule. Missing lemmas exactly N1/N2/N3. Corroborating parable: S(ρ)=S(σ₁)=S(σ₂)=1.5 bits with S(ρ‖σ₁)=0, S(ρ‖σ₂)=0.25 bits>0 — entropy data underdetermines relative entropy.
VERIFICATION (V-004, complete:true, no unresolved): (a)–(e) EXHAUSTIVE over full 296-line ledger + regex; NOWHERE clause confirmed §-by-§ (zero Hilbert/operator/density-matrix/expectation-functional hits); parable arithmetic executed. → C-018 [R]. H6 SUPPORTED (proven-negative form). Numeric parable reproduced orchestrator-side: `reproducibility/check_toys.py`.

## D-005 — H1 prerequisite reduction (H1 cycle 1) [VERIFIED]

Full statement: `cycle-screens.md` D-005. General first law δS=δ⟨K_σ⟩ needs: (i) ω_R(M) on common algebra with M-differentiability, (ii) modular K_σ (cyclic+separating vector, Tomita–Takesaki — unestablished for GSST R), (iii) BCHM setting mismatched (balls-in-vacuum vs half-line-R-in-evaporating-BH). Existence premises = H6's N2 + differentiability + cyclic/separating.
VERIFICATION (V-002, reduction logic): → C-014 [R]. H1 CONDITIONAL-ON-N2 (dormant).

## D-006 — Toy screens H2/H3/H0 (cycle 1 each) [VERIFIED]

Full statements: `cycle-screens.md` D-006.
H2: "flat entropy ⟹ zero Fisher" REFUTED — |ψ(θ)⟩=cosθ|0⟩+sinθ|1⟩ has S_vN=0 ∀θ, QFI F=4(⟨∂ψ|∂ψ⟩−|⟨ψ|∂ψ⟩|²)=4≠0 (derived + numeric θ∈{0,0.3,1.0,2.5}). V-004 item C → C-017 [R]. H2 REJECTED-as-stated.
H3: "value-min switch ⟹ distance kink" REFUTED — numbers don't fix states. As-stated qubit literal impossible (V-002 K-flag); REPAIRED qutrit toy diag(λ,(1−λ)/2,(1−λ)/2): S=1 at λ*≈0.7729, dS/dλ≈−2.767≠0, Bures smooth (~1e−8 2nd-deriv agreement) while min(S1,S2) kinks. V-002 → C-015 [R]. H3 REJECTED-as-stated. Numbers reproduced orchestrator-side: `reproducibility/check_toys.py` (λ*=0.772908, dS/dλ=−2.767018, |Δd2|=6.66e−08).
H0: coherence toy PASSED — value branches + min-rule compatible with M-independent states (zero QFI). V-003 item B (6-step CONSISTENT) → C-016 [R]. H0 toy-coherent, GSST content dormant behind N2.

## D-007 — Missing-lemma triple: definitions + necessity argument (adversarial repair, ADV-001 wounds 1a-ii/1a-vii) [ORCHESTRATOR ARGUMENT, verifier-readable]

The construction task: DERIVE the radiation state family ω_R(M) on a common algebra A_R from ledger v1.2 + pinned imports (P-001…P-008) — i.e., determine GSST's state from GSST's collapse data, not merely posit an arbitrary state.
- N1 (interior data beyond cited [15]): the semiclassical interior evolution / honourable specification of how the incoming I⁻ coherent data propagates to I⁺ through the evaporating geometry. [15] is cited-not-pinned AND covers N free scalars, not holographic-CFT matter — doubly unavailable.
- N2 (stated algebra+regulator prescription): an explicit outgoing-radiation algebra A_R (factor type, region anchoring across M-spacetimes) plus an M-independent regulator scheme (δ prescription, subtraction rule) under which states/functionals and info quantities are defined.
- N3 (plan11 §5E QG inputs): whichever additional quantum-gravity structures the island rule's state-level content needs — encoding isometry / complete QG Hilbert space / complementary-subsystem factorization / central-dogma assumption — to fix what the island encodes and which replica density matrix (if any) lives on the gravitational region.
Necessity (argued, per lemma): N2 necessary by definition of the task (no states without a stated algebra+regulator). N1 necessary for the DERIVATION task (conceded to ADV-001: bare SPECIFICATION of some state needs only N2; DETERMINING GSST's state from its collapse data needs propagation through the interior). N3 necessary because the ledgered QRT rule is extremize-then-min over numbers — its state-level content (encoded DOF, replica rule on gravitating regions) is nowhere fixed semiclassically. Joint sufficiency: SKETCHED only (N1 evolve → N2 restrict → N3 encode); not proved, and sufficiency is not claimed as a verified result.
Consequence: D-004/C-018/H6 "exactly N1/N2/N3" is DOWNGRADED to "at least N1, N2, N3 (necessity argued; sufficiency sketched)". The unconstructibility conclusion is unchanged.
