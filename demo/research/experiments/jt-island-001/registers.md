# JT-ISLAND-001 Supporting Registers v0.1 (prompt9 §6 artifacts 2,3,4,5,6,8)

Companion to `phys-pseudo-r1r4-v0.1.md`. Status vocabulary: ESTABLISHED /
DERIVED / SOURCE-REPORTED / CONDITIONAL / POSTULATED / APPROXIMATE /
CONJECTURAL / DISPUTED / UNSUPPORTED / FALSIFIED (dossier-only statuses from
audit v1.0 are superseded by addendum v1.1 where they differ).

## 1. Assumption and imported-result register

| ID | Category | Statement | Scope | Source | If altered | Status |
|---|---|---|---|---|---|---|
| A-DEF-1 | Defining | Canonical model = AMMZ holographic extension of AEMM two-sided JT+bath | Whole spec | Director Decision 1 | Different model = different spec | POSTULATED |
| A-DEF-2 | Defining | One-endpoint island I_t=(−∞,a(t)] | Part B | Config adopted | Endpoint count changes area term (Decision 2 rule) | POSTULATED |
| A-PHY-1 | Physical postulate | No diffeomorphism anomaly | R1 | Standard QFT | Ward identity fails; R1 void | POSTULATED |
| A-PHY-2 | Physical postulate | Holographic CFT matter (large c, strong coupling) | R4 route | AMMZ §2 | RT route unavailable → G(η) fallback | POSTULATED |
| A-PHY-3 | Physical postulate | RT/HRT correctness for von Neumann entropy | R4 route | AMMZ line 469; [5,6] | Island rule loses derivation | POSTULATED |
| A-PHY-4 | Physical postulate | Quench state (HH/eternal + right coupling + E_S) | R1/R4 | AEMM §3.1-3.2 | Different state = different flux/entropy | POSTULATED |
| A-MATH-1 | Regularity | Smooth manifold/dirichlet BCs for variation | R1 derivations | Standard | Variation ill-defined | POSTULATED |
| A-MATH-2 | Regularity | S_gen smooth in a^± (stationarity meaningful) | QES step | Assumed | No QES via stationarity | CONDITIONAL |
| A-BC-1 | Boundary/initial | Cutoff z=ϵf′(u); φ_b=φ̄_r/ϵ; f(0)=0,f′(0)=1 | R1/R4 | AEMM(12),(64)ff | UV scheme change; S_gen split shifts | POSTULATED |
| A-BC-2 | Boundary/initial | S₀≪S^i_Bek (island can dominate) | R4 | AMMZ §3.3 2S₀ argument | No Page transition (large_S0_regime branch) | CONDITIONAL |
| A-IMP-1 | Imported theorem | JT action/equations AEMM(6)-(11) | R1 | AEMM | — | SOURCE-REPORTED |
| A-IMP-2 | Imported theorem | Anomaly + stress forms AP(5.1)-(5.2), AEMM(42) | R1 | AP/AEMM | — | SOURCE-REPORTED |
| A-IMP-3 | Imported theorem | Schwarzian/ADM/evaporation AEMM(13)-(25) | R1 | AEMM | — | SOURCE-REPORTED |
| A-IMP-4 | Imported theorem | QES position AMMZ(12)←[8,9] | R4 | AMMZ/Penington/AEMM §4 | — | SOURCE-REPORTED |
| A-IMP-5 | Imported theorem | RT/HRT minimality → branch selection | R4 | AMMZ §3.3←[5,6] | — | SOURCE-REPORTED |
| A-SC-1 | Semiclassical approx | T_ab=<T_ab>; fluctuations suppressed; k≪1 | R1 | AEMM; hierarchy | Backreaction uncontrolled | APPROXIMATE |
| A-SC-2 | Semiclassical approx | Slow-evaporation in branch intersection (F1-F5) | t_P derivation | Baseline §7 | t_P shifts at O(1) | APPROXIMATE |
| A-REN-1 | Renormalizer | S_gen bare/Ren split; φ₀^Ren; N-endpoint log | R4 | AEMM(69)-(73) | Scheme dependence | SOURCE-REPORTED |

## 2. Claim register (consequential claims only; full chain is the spec)

| Claim (spec name) | Premises | Source | Status | Limitation |
|---|---|---|---|---|
| curvature_constraint, metric_eqs | jt_action, stressT | AEMM(9),(11) | DERIVED (rebuilt) | Correct signs per audit; OpenCode to re-derive |
| ward_identity | no_diffeo_anomaly | Standard | SOURCE-REPORTED | Dossier citation voided |
| chiral_flux | conservation + anomaly | AP/AEMM structure | DERIVED (rebuilt) | Given p.11 inputs (verified) |
| schwarzian_dynamics, evap_law | jt_action + quench | AEMM(13)-(25) | DERIVED (rebuilt) | Sign verified via E₀>0 |
| global_balance | energy_balance + flux | AEMM(22) + standard | CONDITIONAL | ?OB-4 derivation owed |
| noisland_slope | noisland_branch | Verified form + AMMZ(14) | DERIVED | UV const non-universal |
| qes_position | import chain | AMMZ(12)←[8,9] | SOURCE-REPORTED | ?OB-9 re-derivation not done |
| island_value (1×) | Decision 2 + AMMZ(13) | AMMZ §3.1.2 | DERIVED (rebuilt) | Logs subleading only |
| page_condition | branches + selection | AMMZ §3.3 | DERIVED | — |
| pagetime_value | intersection | Baseline derivation | CONDITIONAL (UNFROZEN) | Footprint F1-F5 (?OB-7) |
| page_curve | min-selection | AMMZ §3.3 structure | DERIVED (rebuilt) | Inherits t_P footprint |
| island_rule_status | AMMZ §4 | AMMZ(15) | CONDITIONAL prescription | RT/HRT-dependent |
| interop_complete | 10 links | This spec Part C | CONDITIONAL | Carriers OB-4/5/9/10 |

## 3. Derivation dependency graph (text form; cycle check: acyclic)

```
jt_action ─┬─→ curvature_constraint ─→ (background for all R4)
           └─→ metric_eqs ─→ schwarzian_dynamics ─→ energy_balance ─→ evap_law
no_diffeo_anomaly + trace_anomaly + stress_forms ─→ ward_identity
  ─→ lightcone_conservation ─→ chiral_flux ─→ (flux into energy_balance)
quench_state ─→ shock_stress ─→ evap_law; quench_state ─→ holographic_entropy_route
R1 outputs (metric, state, stress, f-trajectory, cutoff) ──→ R4 inputs (Part C links)
radiation_region + island_region + sgen_functional ─→ noisland_branch ─→ noisland_slope
qes_stationarity + qes_position + dilaton_term ─→ island_value
branches + branch_selection ─→ page_condition ─→ pagetime_value(C) ─→ page_curve
```
No edge points backward: R1 uses no R4 result; R4 consumes R1 outputs only
via Part C links. The qes_position import (AMMZ(12)←AEMM §4) depends on AEMM
dynamics compatible with R1's — recorded, not circular (import, not premise).

## 4. Bridge register (ten-field; BR-IDs rebuilt from dossier)

| ID | Source → Target | Mapping | Assumptions | Preserved | Lost/approx | Regime | Obligations | Failure | Status |
|---|---|---|---|---|---|---|---|---|---|
| BR-01 | JT action → field eqs | Variation | A-MATH-1 | EOMs, symmetries | Boundary terms (bulk only) | Semiclassical | Re-derive variation | Sign/coeff mismatch | DERIVED (rebuilt) |
| BR-02 | Anomaly → stress forms | AP/AEMM transformation | 2D CFT, fixed c | Trace value | State piece separate | Any CFT₂ | Convention pin (?OB-2) | Wrong c/normalization | CONDITIONAL |
| BR-03 | Diffeo inv → Ward ID | Noether | A-PHY-1 | Local conservation | None (identity) | No diffeo anomaly | None | Diffeo anomaly present | SOURCE-REPORTED |
| BR-04 | Grav ↔ bath interface | Quench + t=f(u) + shock | A-PHY-4 | Stress-energy flow | Smooth-continuity picture | Quench model | ?OB-5 branch | Reflecting BC instead | CONDITIONAL |
| BR-05 | State → matter entropy | 3D-RT (route i) | A-PHY-2/3/4 | Fine-grained S (given RT) | 2D-CFT explicitness | Holographic | ?OB-10 RT reconstruction | Non-holographic matter | CONDITIONAL |
| BR-06 | Dilaton → grav entropy | Area law/endpoint | A-DEF-2 | 1×/endpoint count | Multi-endpoint sum needed | JT | Endpoint count | Wrong count (dossier 2×) | SOURCE-REPORTED |
| BR-07 | Area+matter → S_gen | Additivity + Ren split | A-REN-1 | Regulator independence | Separate UV pieces | Semiclassical | Ren scheme stated | Scheme mismatch | CONDITIONAL |
| BR-08 | S_gen → QES | Stationarity | A-MATH-2 | Extremum location | Non-extremal candidates | Smooth S_gen | ?OB-9 re-derivation | No critical point | CONDITIONAL |
| BR-09 | Candidates → S(R) | Min-selection | A-IMP-5 | Page transition | Losing branch | RT/HRT regime | None beyond BR-05 | S₀-large (no transition) | SOURCE-REPORTED |
| BR-10 | R1↔R4 integration | Part C 10 links | Per-link | Single-model use | Per-link gaps (?OBs) | Adopted model | OB-4/5/9/10 | Setup mismatch (E) | CONDITIONAL |

## 5. Unresolved obligations (?OB-1..10 + carriers)

OB-1 (K−1) counterterm: shortlisted not-needed; revisit for on-shell finiteness.
OB-2 R-form trace equality: pin or drop (dropped downstream). OB-3 S_Polyakov
coefficient: suspended; anomaly route used. OB-4 dM/dt line-by-line: OWED in
Part A scope. OB-5 shock-vs-smooth branch: decide via derivation. OB-6 S₀ term:
re-add only if sourced. OB-7 t_P footprint F1-F5: close before any freeze.
OB-8 route decision: CLOSED (route i adopted; fallback documented). OB-9 AEMM
§4 re-derivation: explicitly deferred, import recorded. OB-10 RT
reconstruction: deferred, import recorded. Interop carriers: OB-4/5/9/10.

## 6. Limitations and failure conditions (re-derived, not copied)

(L1) Vanishing coupling G_N→0: area term diverges → min-selection keeps
no-island branch for all t; NO Page transition (classical limit). Dossier's
direction-correct case, re-derived from Decision-2 normalization.
(L2) Reflecting bath BC: no energy exchange → E const → static BH, no
evaporation, no Page curve (AEMM setup with coupling off).
(L3) Hierarchy breakdown c∼S₀ or k∼1: QES delocalizes, semiclassical control
lost; S_gen extremization unreliable (AEMM near-extremal breakdown).
(L4) S₀≳S^i_Bek: 2S₀ island price blocks transition (AMMZ §3.3; spec
large_S0_regime branch). (L5) G_N→∞: strong quantum gravity — dossier RC2's
"removal" reading CORRECTED to breakdown (baseline §7). (L6) Non-holographic
matter: route (i) void → fallback route (ii) with full G(η) hypotheses owed.
SCOPE (retained from dossier, verified honest): N0/N1 reconstruction only; NO
microscopic S-matrix, NO interior reconstruction, NO 4D extension, NO Bohmian
gravity, NO general QG solution.
