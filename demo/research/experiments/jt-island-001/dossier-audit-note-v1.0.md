# JT-ISLAND-001 Dossier Audit Note v1.0 (Step 2 deliverable)

**Scope:** mathematical and source-integrity audit of "JT Black Hole Research
Dossier v1.0" (25 pp, plans/) against primary sources, per prompt9 §3.
**Method:** every displayed dossier equation read from rendered PDF pages
(`pdftotext` loses all equations); every locator checked against the cited
source; independent algebra where flagged. **Result: 8 BLOCKER-class and
14 MATERIAL-class findings. Load-bearing contradictions confirmed — see
§8 (pause-and-report).**

**Primary sources inspected (all in `sources/`, SHA-256 recorded there):**
AMMZ arXiv:1908.10996v2 (22 pp, 15 numbered eqs, single-sequence);
AEMM arXiv:1905.08762v3 (55 pp, 127 numbered eqs, single-sequence);
AP arXiv:1402.6334v3 (tiebreak for anomaly coefficients and 16π/8π pattern);
Penington arXiv:1905.08255 (obtained; QES-origin chain, not yet deep-read).
Polyakov 1981: canonical citation recorded; original NOT obtained (no trusted
open copy). Necessity determined (review amendment 2): NOT necessary for
issues A/F resolution — AP (5.1)-(5.2) + AEMM suffice. Necessary ONLY if the
spec retains the explicit nonlocal S_Polyakov formula (currently unverified).

**Severity key:** BLOCKER = load-bearing falsehood/inconsistency; must be
corrected/narrowed before dependent formalization. MATERIAL = weakens scope;
must be qualified. MINOR = recorded, non-blocking.

---

## 1. Issues A–H verdicts

### A. Trace anomaly vs covariant conservation — BLOCKER (conceptual falsehood)

Dossier RC2 (p.17): "Removing the trace anomaly breaks covariant conservation
for CFTs in curved space (∇Left = 0 ≠ ∇Right)". FALSE, three witnesses:
(1) the dossier's OWN p.13 derivation assumes ∇T = 0 (Ward identity) and uses
the anomaly only to supply T_+− — conservation and anomaly are joint
INDEPENDENT inputs there; RC2 contradicts the dossier's own derivation.
(2) AP §5.1: anomaly gives T_+− (5.1), then "Conservation of T_ab implies
that" T_±± take form (5.2) — conservation imposed independently.
(3) A traceless conserved tensor (classical CFT) satisfies both; removing the
anomaly changes the SOLUTION (kills Hawking-flux terms), not the identity.
Correct statement: anomaly + conservation jointly determine the flux; dropping
either changes predictions; neither causes the other. The p.13 chiral-conclusion
derivation structure itself is SOUND given p.11 inputs.

### B. No-island entropy and Page time — BLOCKER (three mutually inconsistent numbers)

Dossier p.8 prints S_no-island = (c/6)ln[(β/πε)sinh(πt/β)] + S₀ — the LOG
FORMULA IS CORRECT (standard Cardy-Tonni; implies slope πcT/6, agreeing with
AMMZ(14) and AEMM(112)). But its TWO printed linearizations are wrong:
(cπ/3β)t is 2× the formula-implied (cπ/6β)t, and (2πc/3)T_H·t disagrees with
BOTH (4× the implied value given βT_H = 1). Page time p.10, t_P ≈ 3βS_BH/πc:
independent derivation from correct branches gives 6βS_BH/πc (2× off); the
printed t_P equals (correct 1× island)/(wrong 2× slope) — mixing inconsistent
normalizations — and is inconsistent with the dossier's own 2× island value.
+S₀ term unsourced. Neither AMMZ (Page time implicit, equality of (13)=(14))
nor AEMM (order estimate t_P ∼ k⁻¹(T₁−T₀)/T₀ in a DIFFERENT two-sided+shock
setup) contains the printed formula.

### C. Island normalization and endpoint counting — BLOCKER (denominator error)

Dossier p.10: S_island = (φ₀+φ_h)/2G_N = 2S_BH. The 2× comes from a 2G_N
denominator contradicting: (a) dossier's own p.8 S_grav (4G_N, ONE endpoint
∂I = {a}); (b) dossier's own p.9 S_BH = (φ₀+φ_h)/4G_N; (c) AMMZ(13)
S = S_Bek(T(t)) + logs (1×); (d) AEMM S_gen = (φ₀+φ)/4G_N + S (1×).
Single-sided one-endpoint island gives 1×; the 2× matches NOTHING in the
stated model (AMMZ's 2× items are the EARLY-branch asymptote (14)→2S^i_Bek
and the eternal-case 2S₀ island price — likely conflation sources).
Additionally the p.10 Page-time equality writes (φ_h/4G_N) = S_BH, silently
DROPPING φ₀ vs p.9's definition. "Saturates" misdescribes a decreasing
branch; O(c⁰) misstates AEMM's O(c·log) matter terms.

### D. Two-interval CFT entropy — BLOCKER-adjacent (unstated purity hypothesis)

Dossier p.9: S(R∪I) = (c/6)ln[d(a,b)²/(ϵ_aϵ_b e^ρe^ρ)] is the SINGLE-interval
formula applied to the COMPLEMENT [a,b] — valid ONLY via S(R∪I) = S([a,b])
for a PURE total state. Purity is NEVER stated/justified; it holds (arguably)
in AMMZ's single-sided collapse setup but FAILS in AEMM's setup (R∪B mixed,
entangled with L — hence AEMM's G(η) cross-ratio machinery, Eqs. (56)-(60),
with G undetermined except η→0,1 limits). NEITHER source contains the printed
formula (AMMZ computes via 3D RT; AEMM keeps log G(η) explicit). Further
defects: notation flip e^ρ → z_a z_b between the two p.9 lines (bath point has
no AdS z); log argument (a⁺−b⁺)(a⁻−b⁻) < 0 for spacelike separation needs
modulus/ordering (cf. AEMM(67)). Formula usable ONLY after purity justification
or replacement by AEMM (60)/(66)-(68).

### E. R1–R4 compatibility — BLOCKER (setups incompatible as specified)

C-INT-01 ("identical state/geometry/charge") is ASSERTED; evidence shows
INCOMPATIBLE blend: AMMZ single-sided collapse/pulse narrative (pure-state,
holographic matter, 3D-RT entropy) vs AEMM two-sided eternal + quench
(E_S shock, mixed subsystem, replica entropy, t = f(u) time-matching).
Specific incompatibilities: (1) formation history unspecified (AMMZ: late-time
results history-independent — never engaged); (2) holographic-matter
requirement (AMMZ §2) NEVER MENTIONED in dossier (grep: zero hits);
(3) E_S shock (AEMM (21),(44)-(46),(30)) ABSENT (zero hits) — yet R1 uses
AEMM's shock-derived evaporation law; (4) cutoff z = ϵ const vs AEMM z = ϵf'(u)
(the f' factor is load-bearing in AEMM (64),(68)); (5) direct time-matching
vs AEMM t = f(u); (6) coordinates x^±/y^±/t/u/x̄ used interchangeably, NO
coordinate dictionary; (7) φ_h means THREE things (p.3 boundary const φ̄/ϵ;
p.9 horizon dynamical; p.10 equality ambiguous); (8) κ collision (dossier p.8
surface gravity 2πT_H vs AMMZ/AEMM evaporation rate); (9) hierarchy
φ₀ ∼ φ_h contradicts AEMM φ₀ ≫ φ and tensions AMMZ S₀ ≪ S^i_Bek (the 2S₀
island price could kill the island the dossier claims); (10) renormalization
bare-vs-Ren (AEMM §4.1) unaddressed. A single consistent model is NOT
established. REQUIRES DIRECTOR SCOPE DECISION (see §8).

### F. JT action and field equations — BLOCKER (sign + coefficient defects)

(1) Action p.3: topological coefficient φ₀/4πG_N is 4× AEMM(6)'s φ₀/16πG_N
(implies S₀ = φ₀/G_N, contradicting dossier's own S_BH on p.9); boundary
counterterm (K−1) vs AEMM's K, unexplained. (2) Covariant metric equation
p.12: (g∇²φ − ∇∇φ + gφ)/8πG_N has WRONG SIGN on gφ term — four witnesses:
AEMM(11a) trace, vacuum-solution test (AEMM(16) fails dossier's trace,
passes AEMM's), AP (5.3)/(5.5)-(5.6) 16π-vs-8π pattern, and the dossier's OWN
p.16 RC2 form (∇∇φ − g□φ + gφ = −8πG<T>, which has the CORRECT trace).
(3) +− light-cone expansion p.12 ([2∂∂φ + e^{2ρ}φ]/8πG_N): 2× AEMM(11a) AND
inconsistent with the dossier's own covariant form. (4) NEVER-PRINTED action
variation: dossier writes "S_JT + S_CFT" symbolically at the variation step.
VERIFIED CLEAN: R + 2 = 0 ✓; ++/−− constraints match AEMM(11b-c) EXACTLY ✓.

### G. Interface and global energy balance — BLOCKER-adjacent (sign error)

(1) Matching p.13 (T^grav = T^bath at interface): oversimplified — no
f-reparametrization (AEMM §3.1's whole point), no E_S shock (AEMM quench has
a δ-shock, NOT smooth continuity); matches AMMZ §2.1 PROSE only.
(2) ADM mass M(t) = +(φ̄/8πG){x⁺,t}: SIGN ERROR vs AEMM(14)
E = −(φ̄_r/8πG){f,u} — gives NEGATIVE mass for the BH solution ({tanh} < 0).
(3) dM/dt = −(πc/12)T_H² + F_in: standard-correct FORM (2D Stefan-Boltzmann;
AEMM(22) structure) but UNDERIVED ("differentiating M(t) yields" — asserted;
inconsistent with the mis-signed M(t)). (4) E_S transient absent. Actual
sources: AEMM (14),(20),(22) — dossier cites AMMZ (contains none of these).

### H. Epistemic classifications — MATERIAL (pervasive mislabeling)

DERIVED claimed for: imported results (QES position AMMZ(12)←[8,9]),
defective derivations (BR-01 field eqs), never-performed computations
(BR-05 replica). POSTULATED claimed for imported-standard results (BR-06 area
law ← EW[7]; BR-09 minimization ← RT/HRT minimality, AMMZ line 555).
ESTABLISHED claimed for prescriptions under test (BR-07 S_gen additivity).
Circular self-citation: minimization rule cited to rationale.md ("2"), Ward
identity to protocol.md ("1"). Penington (co-origin of QES position) NEVER
CITED. RT/HRT-dependence of AMMZ's derivation (AMMZ line 469) unacknowledged.
All H-labels must be redone in the baseline (prompt9 §6 ten-label scheme).

---

## 2. Claim dispositions (summarized; full register in Step 7)

| ID | Printed status | Disposition |
|---|---|---|
| C-R4-01 (S_gen) | ESTABLISHED | CONDITIONAL: formula form ✓ (AMMZ(10)/AEMM(69)); locators invalid; bare/Ren gap |
| C-R4-02 (no-island) | DERIVED | SPLIT: log formula SOURCE-REPORTED ✓; linearizations INCORRECT; locator invalid |
| C-R4-03 (QES) | DERIVED | UNSUPPORTED as printed; correct content at AMMZ(12)←[8,9] |
| C-R4-04 (island) | DERIVED | INCORRECT (denominator); locator invalid |
| C-R4-05 (Page) | DERIVED | SPLIT: rule CONDITIONAL (sourcing); t_P INCORRECT; Fig. 7 wrong (→ Fig. 11) |
| C-R1-01 (JT eq) | ESTABLISHED | INCORRECT (action 4×); locator invalid; AP ID mistyped (1405.4251→1402.6334) |
| C-R1-02 (anomaly) | ESTABLISHED | SPLIT: <T>=cR/24π CONDITIONAL ✓-form; T_+−/T_±± VERIFIED (AP(5.1)-(5.2)); S_Polyakov coeff UNVERIFIED; locator invalid |
| C-R1-03 (conserv.) | DERIVED | CONDITIONAL: structure sound-given-inputs; unsourced; inherits coeff uncertainty |
| C-R1-04 (mass/flux) | DERIVED | SPLIT: M(t) INCORRECT (sign); matching INCOMPLETE; dM/dt CONDITIONAL (standard, underived) |
| C-INT-01 (shared) | DERIVED | UNSUPPORTED (asserted); evidence points INCOMPATIBLE |

## 3. Bridge dispositions (summarized)

BR-01 DERIVED→DISPUTED (derivation defective, redo). BR-02 DERIVED→CONDITIONAL
(coeff unsourced). BR-03 DERIVED→CONDITIONAL (standard result, bogus citation).
BR-04 POSTULATED→definition OK / equation INCOMPLETE. BR-05 DERIVED→IMPORTED-
CONDITIONAL (missing purity/G(η)). BR-06 POSTULATED→mislabeled, should be
IMPORTED (AMMZ(10)←EW[7]). BR-07 ESTABLISHED→overstated, CONDITIONAL prescription.
BR-08 DERIVED→split (conditions OK / solution UNSUPPORTED). BR-09 POSTULATED→
mislabeled, IMPORTED-from-RT/HRT + circular citation. BR-10 DERIVED→UNSUPPORTED
(central integration claim).

## 4. Assumption dispositions (summarized)

ASM-01 (JT action): postulate-form OK; printed action INCORRECT (4×).
ASM-02 (hierarchy): INCORRECT as printed (φ₀∼φ_h; ambiguous φ_h).
ASM-03 (transparent BC): definition OK; implementation INCOMPLETE.
ASM-04 (CFT entropy formula): scope OVERSTATED (missing purity/η hypotheses).
ASM-05 (island prescription): defensible label; missing holographic-matter condition.
ASM-06 (pure vacuum state): matches NEITHER source (AMMZ formed-BH; AEMM mixed
subsystem) — INCOMPATIBLE as stated.

## 5. Corrected locator map (dossier → actual)

- "AMMZ Eq. (2.9)" → AMMZ Eq. (10); AEMM Eq. (69). "AEMM Eq. (2.4)" → AEMM (69).
- "AMMZ Sec. 3.2.1, Eq. (3.4)" → AMMZ Eq. (14) (§3.2.2/3.3); §3.2.1 has NO entropy formula.
- "AMMZ Eq. (3.15)" → AMMZ Eq. (12) ← Penington + AEMM §4. "AEMM Sec. 3" → AEMM §4.
- "AMMZ Sec. 3.1.2, Eq. (3.18)" → AMMZ Eq. (13) (§3.1.1); NO Eq. 18 exists (15 total).
- "AMMZ Sec. 3.3, Fig. 7" → §3.3 ✓; Fig. 7 WRONG (formation diagram) → Fig. 11.
- "AMMZ Eq. (2.1)" → AEMM (6)-(8),(11a-c); AMMZ has NO JT equations.
- "AMMZ Eq. (2.5)" → AEMM fn.2 + Eq. (42); AP (5.1)-(5.2); AMMZ has none.
- "AMMZ Sec. 2.1, Eq. (2.8)" → AEMM (14),(20),(22); AMMZ has none.
- GENERAL: AMMZ uses single-sequence (1)-(15); AEMM single-sequence; AP's (X.Y)
  contents do NOT match dossier attributions. ALL dossier (X.Y) locators for
  AMMZ/AEMM are FABRICATED numbering — treat every printed locator as unverified.
- "AP arXiv:1405.4251" → TRUE ID arXiv:1402.6334 (1405.4251 is an unrelated
  statistics paper; my own first guess 1402.1509 was also wrong — resolved via
  AEMM ref (35)).

## 6. What verified CLEAN (fairness record)

Poincaré metric (AEMM(10)) ✓; R+2=0 ✓; ++/−− constraints = AEMM(11b-c) ✓;
S_gen one-endpoint form ✓; single-interval entropy + Weyl rule ✓;
no-island LOG formula + implied slope πcT/6 ✓; chiral-conservation derivation
STRUCTURE (given inputs) ✓; AP-verified T_+−/T_±± forms ✓; p.16 RC2
backreaction form (correct trace) ✓; dM/dt standard form ✓; hierarchy PURPOSE
(c≪S_BH) ✓ (as printed with ∼, ✗); scope-boundary honesty (pp.22-23 N0/N1,
exclusions) ✓ — the dossier's SCOPE discipline is good; its EQUATIONS are not.

## 7. Dossier AP-ID error (C-R1-01)

Printed "AP arXiv:1405.42514" (= 1405.4251 + footnote 4): 1405.4251 is an
unrelated statistics paper. True AP (via AEMM ref 35): 1402.6334, obtained
and verified. The dossier's AP leg is VOID as printed.

## 8. Pause-and-report items (director decision required)

Per the approved plan: the following load-bearing contradictions block
dependent formalization. Independent work continues (§9); Parts B/C and Part C
(interop) drafting PAUSES until decided.
1. MODEL FORK (E): single-sided-collapse (AMMZ-pure: complement trick valid,
   holographic matter required, RT/HRT entropy) vs two-sided+shock (AEMM-full:
   E_S transient, G(η) machinery, t=f(u)). Recommendation: ADOPT AMMZ
   single-sided as the working model (matches dossier narrative + R4 formulas'
   only valid regime); import AEMM equations ONLY with per-equation
   setup-compatibility justification; REJECT silent blending. Please confirm
   or redirect.
2. ISLAND NORMALIZATION (C): adopt 1× (AMMZ(13)) → voids C-R4-04/05 as printed;
   Page time becomes 6βS_BH/πc (from correct branches). Please acknowledge the
   rewrite before I formalize Part B on corrected values.
3. FIELD-EQUATION SIGN (F): adopt AEMM-consistent form (dossier p.16 form).
   Unambiguous (4 witnesses) — proceeding with corrected form unless directed
   otherwise; flagging here for visibility, no pause needed for Part A draft.
4. RC2 REMOVAL CLAIM (A): replace with correct joint-inputs statement. Same
   handling as (3).

## 9. Independent work continuing during pause

AP/Penington registration in sources record; Step-3 baseline skeleton with
CORRECTED R1 equations (flagged provisional pending (1)); register scaffolding
(RC1 inventory, assumption/import seeds); OpenCode checklist seeds. NO Part B/C
derivations on disputed normalizations; NO interop claims until (1) decided.
