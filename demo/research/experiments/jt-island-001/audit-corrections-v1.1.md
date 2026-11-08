# JT-ISLAND-001 Audit Corrections Addendum v1.1

**Parent record:** `dossier-audit-note-v1.0.md` (SHA-256
`c18c6002a3fe54a9fb1b7fdd0ff83750a9238db57801d3008bf76ba83503b24c`),
preserved byte-intact as the pause record. This addendum records the
director's rulings from `plans/plan_prompt9_imp_review.md`, applies the
approved corrections to every disposition, and defines the adopted working
model. It does not edit v1.0; where this addendum and v1.0 differ, this
addendum supersedes for all downstream formalization.

**Authority chain:** prompt9 (assignment) → audit v1.0 (findings + pause) →
imp_review (director rulings) → this addendum (applied corrections) →
`working-baseline.md` (formalization input) → Phys-Pseudo v0.1.

---

## 1. Director decisions (recorded verbatim in substance)

### Decision 1 — Canonical model: ADOPT (with correction)

Adopt the **AMMZ holographic-matter extension of the AEMM two-sided
evaporating JT gravity and nongravitating bath setup**. Do NOT characterize
it as a single-sided-collapse model (the auditor's paused recommendation is
corrected on this point: AMMZ explicitly builds on the AEMM setup,
specializing the matter to a holographic CFT₂).

Component → source mapping (binding):

| Component | Primary source |
|---|---|
| JT gravity and dynamical background | AEMM, with AMMZ compatibility checks |
| Matter stress-energy and evaporation dynamics | AEMM, subject to matching assumptions |
| Holographic matter and higher-dimensional dual | AMMZ |
| Radiation entropy and Page-curve construction | AMMZ |
| Additional QES and information-theoretic interpretation | Penington, where relevant |

Critical rule: AEMM equations may be imported ONLY with a per-equation
demonstration of compatibility with the AMMZ holographic-CFT₂ specialization.

### Decision 2 — Island normalization: ADOPT ONE TERM PER ENDPOINT

For one QES endpoint `a`: S_grav(a) = (φ₀ + φ(a))/4G_N. Multiple endpoints
sum individual contributions (cf. AMMZ's 2S₀ two-endpoint island cost).
REJECT the dossier's (φ₀+φ(a))/2G_N = 2S_BH for its one-endpoint configuration.
**Page time: freeze NEITHER 3βS/(πc) NOR 6βS/(πc).** Reconstruct both
corrected branches and derive their intersection independently; the result is
recorded as derived-not-frozen.

### Decision 3 — Metric equation and trace anomaly: PROCEED WITH CORRECTION

- Approve AEMM-consistent metric-equation signs and coefficients.
- Approve correct JT action normalization including the topological term.
- Replace the RC2 assertion "removing the trace anomaly breaks covariant
  conservation" with the correct distinction: diffeomorphism invariance gives
  the stress-energy Ward identity (absent a diffeomorphism anomaly); the
  conformal trace anomaly gives the nonzero trace. Distinct, coexisting
  statements; the trace anomaly is not a diffeomorphism anomaly. Both
  relationships must be reconstructed, neither derived from the other.

### Resumption authorization

Formalization is resumed (corrected R1 baseline, R4, R1/R4 interop) under the
adopted model. Each disputed result must first be independently reconstructed.
Unresolved critical dependencies stay BLOCKED/`?`. The three Reality Checks
remain mandatory. Deliverable: corrected audit + Phys-Pseudo R1/R4 v0.1 +
registers + OpenCode checklist. No Go code, no repo modifications, no
implementation contract, no freeze. Stop after the review package.

---

## 2. Corrected dispositions

Corrections applied to v1.0 §2–§4. Status vocabulary per prompt9 §6.

### Claims

| ID | Corrected disposition |
|---|---|
| C-R4-01 (S_gen) | SOURCE-REPORTED (AMMZ(10)/AEMM(69), one-endpoint form); bare-vs-Ren split handled in baseline per AEMM §4.1; dossier locators VOID |
| C-R4-02 (no-island) | SPLIT: log formula SOURCE-REPORTED (verified; slope πcT/6); dossier linearizations FALSIFIED; +S₀ dropped unless sourced; re-source to AMMZ(14) |
| C-R4-03 (QES) | Dossier formula FALSIFIED/UNSUPPORTED; adopt AMMZ(12)←Penington/AEMM §4 chain as SOURCE-REPORTED import; derivation obligation on QES-solver step |
| C-R4-04 (island) | Dossier 2× expression FALSIFIED; adopt 1×/endpoint (Decision 2); branch value DERIVED-in-baseline from corrected inputs |
| C-R4-05 (Page) | Minimization rule SOURCE-REPORTED (RT/HRT minimality, AMMZ §3.3); t_P OPEN — derive from corrected branches, do not freeze |
| C-R1-01 (JT eq) | Dossier action FALSIFIED (4×); adopt AEMM(6)-(8) action + (11a-c) equations as SOURCE-REPORTED; AP leg re-sourced to 1402.6334 |
| C-R1-02 (anomaly) | <T>=cR/24π SOURCE-REPORTED; T_+−/T_±± forms VERIFIED (AP(5.1)-(5.2)); R-form equality SUSPENDED pending convention pin (recommend dropping); S_Polyakov coefficient UNSUPPORTED (obligation; baseline uses anomaly-transformation route instead) |
| C-R1-03 (conserv.) | Ward identity SOURCE-REPORTED (standard); light-cone expansion CONDITIONAL (verify in baseline); chiral-conclusion argument DERIVED-in-baseline from verified inputs |
| C-R1-04 (mass/flux) | M(t) dossier sign FALSIFIED → adopt AEMM(14); matching INCOMPLETE → re-derive with f-trajectory + shock; dM/dt CONDITIONAL (standard form, derivation owed in baseline) |
| C-INT-01 (shared) | UNSUPPORTED as printed; REBUILT under adopted model via per-equation compatibility demonstrations (Decision 1 rule) |

### Bridges

BR-01 DISPUTED→ re-derive (corrected action/variation). BR-02 CONDITIONAL
(anomaly route; Polyakov-action leg suspended). BR-03 CONDITIONAL→ re-source
(standard Ward identity; drop protocol.md citation). BR-04 definition
RETAINED as postulate; equation REBUILT (f + shock). BR-05 CONDITIONAL
(pending issue-D re-resolution under two-sided model). BR-06 re-label
IMPORTED (AMMZ(10)←EW[7]). BR-07 re-label CONDITIONAL prescription.
BR-08 split: conditions RETAINED / solution REBUILT from AMMZ(12) chain.
BR-09 re-label IMPORTED (RT/HRT minimality); drop rationale.md citation.
BR-10 REBUILT under Decision 1 compatibility rule.

### Assumptions

ASM-01 formulation CORRECTED (AEMM(6)-(8)). ASM-02 REWRITTEN as
c≪S₀≪S^i_Bek with single-meaning symbols (φ_h collision resolved in
baseline dictionary). ASM-03 definition RETAINED; implementation REBUILT.
ASM-04 scope NARROWED to single-interval + stated hypotheses; disjoint-union
use governed by issue-D resolution. ASM-05 RETAINED as postulate WITH the
holographic-matter condition made explicit. ASM-06 REWRITTEN for the adopted
two-sided Hartle-Hawking + quench state (AEMM §3); single-sided purity claims
dropped.

---

## 3. Corrected locator map (binding for formalization)

AMMZ(10) S_gen; AMMZ(12) QES position (←[8,9]); AMMZ(13) late-time entropy;
AMMZ(14) growth branch; AMMZ(15) island rule; AMMZ §3.3 + Fig. 11 (NOT Fig. 7)
Page construction. AEMM(6)-(8) action; AEMM(9) R+2=0; AEMM(11a-c) metric
equations; AEMM(13)-(14) Schwarzian/ADM; AEMM(20) E₀; AEMM(21)-(25) flux and
evaporation law; AEMM(42),(44)-(46) stress-tensor transformation + shock;
AEMM(56)-(60) cross-ratio/G(η); AEMM(61)-(68) AdS entropy formulas;
AEMM(69)-(73) S_gen renormalization; AEMM §4 QES derivations;
AEMM(108),(112) scrambling/bath-entropy checks. AP(5.1)-(5.2) anomaly +
conservation forms; AP(5.3)/(5.5) coefficient pattern. Penington: TBC in
bounded read (Step 2). All dossier (X.Y) locators for AMMZ/AEMM are VOID;
the dossier's "AP 1405.4251" is VOID (true ID 1402.6334).

## 4. Remaining obligations by issue (director's table, adopted)

B: recalculate both branches + intersection (t_P unfrozen). C: endpoint
counting + 1× normalization. D: exact state/purity/cross-ratio/block
assumptions under the adopted two-sided model. E: coordinates, states, shock,
boundary trajectory, regulators reconciled per-equation. F: action
normalization, metric variation, coordinate expansions corrected. G: ADM sign
corrected; flux matching independently derived. H: citations, locators,
imports fixed; circular self-certification removed.
