# JT-ISLAND-002 Model Selection Research Report v1.0

**Governing authority:** Phys Research Workflow Protocol v1.0  
**Research mode:** Comparative literature investigation, mathematical feasibility assessment and model selection  
**Implementation authorization:** NONE  
**Literature cut-off:** 8 November 2026


## 1. Executive recommendation

**Preferred model: Candidate C — Dimensionally Reduced Einstein–Hilbert / Schwarzschild Dilaton Gravity (DREH).**

The DREH model is the only candidate that combines a direct structural link to four-dimensional Schwarzschild geometry with an explicit semiclassical backreaction mechanism and a published island/QES Page-curve calculation in the same evaporating background. Its two-dimensional action is obtained by spherical reduction of the four-dimensional Einstein–Hilbert action, scalar matter is coupled, and the Polyakov–Liouville effective action supplies the quantum stress-energy responsible for backreaction. The companion collapse paper constructs the evaporating solution perturbatively from a Minkowski initial state and tracks the geometry to a static end-state. The Page-curve paper applies the island rule via the quantum extremal surface approach in that same collapse scenario. This is the cleanest available realization of the requirement that R1 and R4 belong to the same physical model, because the geometry consumed by the entropy calculation is the geometry produced by the semiclassical evolution.

The recommendation is conditional. DREH is a perturbative effective two-dimensional theory. The Page curve is obtained at a definite order in the semiclassical expansion, the dilaton is treated classically in the existing calculations, and the reduction is an s-wave truncation of four-dimensional gravity, not a full four-dimensional quantum-gravity calculation. These limitations are transparent and controllable, which is exactly what the assignment asks for when every candidate has a material limitation. Candidate B (RST with holographic CFT matter) is a close second and is arguably faster to a working POC; it is discussed below and should remain the fallback if the DREH perturbative audit reveals an unbridgeable gap between the background used in R1 and the background used in R4.

**No manufactured interoperability.** The recommendation does not combine DREH with RST, JT, or any other model. It selects one model and identifies the mathematical objects that genuinely cross the R1/R4 boundary within that model.


## 2. Candidate comparison

The weighted framework is applied with scientific coherence as a hard gate. Scores are given on a 0–5 scale before weighting. A candidate that fails the coherence gate is marked NO-GO regardless of aggregate score.

| Criterion | Weight | A: JT/AMMZ | B: RST + holographic CFT | C: DREH | D: Fabbri–Russo |
|---|---:|---:|---:|---:|---:|
| Genuine R1/R4 coherence | 30% | 2.5 | 4.0 | **4.5** | 3.5 |
| Speed / feasibility of POC | 25% | 3.5 | **4.5** | 3.5 | 3.5 |
| Mathematical breadth | 15% | 4.0 | 4.0 | **4.5** | 3.5 |
| Primary-source completeness | 10% | 4.0 | **4.5** | 4.0 | 3.5 |
| Clarity for technical audience | 10% | 4.0 | 4.0 | **4.5** | 3.0 |
| Reproducibility / robustness | 10% | 3.0 | **4.0** | 3.5 | 3.5 |
| **Weighted aggregate** | | **3.23** | **4.18** | **4.20** | **3.48** |

**Hard-gate note on Candidate A (JT/AMMZ).** JT gravity with a holographic CFT bath is the historical source of the island Page curve, but the AEMM near-extremal hierarchy and the AMMZ construction are not obviously the same parameter regime. The previous JT-ISLAND-001 investigation already returned BLOCKED on this point. The assignment explicitly forbids repairing that specification by deleting an inconvenient assumption. Candidate A therefore remains a reference point, not a selection.

**Why DREH narrowly edges RST.** The aggregate scores are close. DREH wins on coherence and breadth because the two-dimensional dilaton is tied to the four-dimensional spherical geometry, and because the same Polyakov–Liouville stress-energy appears in both the backreaction equations and the generalized-entropy construction. RST wins on speed and reproducibility because the Gautason–Schneiderbauer–Sybesma–Thorlacius calculation is analytically solvable at leading semiclassical order and the Page time is explicitly derived as one third of the black-hole lifetime. The choice of DREH is driven by the assignment’s instruction to prioritize mathematical and physical coherence first and to treat the four-dimensional connection as a genuine scientific asset, not merely a presentational one.

**Candidate E — additional alternatives.** The independent search did not identify a stronger alternative that satisfies all eight non-negotiable requirements. Recent work on Wald-like entropy in DREH (arXiv:2607.24883) provides a Lorentzian complement to the Euclidean island calculation and strengthens DREH’s R1/R4 bridge, but it is a refinement of Candidate C rather than a separate model. The Fabbri–Russo family contains RST as the \(n=1\) special case and produces \(n\)-independent Page curves at leading order, but its parameter \(n\) introduces an extra knob without adding a qualitatively new R1/R4 bridge. No other asymptotically flat two-dimensional model in the searched literature offers a cleaner combination of explicit action, backreaction, and published Page curve.


## 3. Canonical model definition (DREH)

**Action.** The two-dimensional dilaton-gravity action is obtained by spherical reduction of the four-dimensional Einstein–Hilbert action. Schematically,

\[
S_{\text{DREH}} = \frac{1}{16\pi G_2}\int d^2x\sqrt{-g}\,e^{-2\phi}\left[R + 2(\nabla\phi)^2 + \frac{2}{a^2}\right] + S_{\text{matter}} + S_{\text{PL}},
\]

where \(\phi\) is the dilaton, \(a\) is a length scale inherited from the reduction, and \(S_{\text{PL}}\) is the Polyakov–Liouville effective action. The dilaton is related to the four-dimensional spherical radius; the effective two-dimensional theory is the s-wave sector of four-dimensional Einstein gravity.

**Matter content.** Scalar matter is coupled to gravity. The quantum correction is incorporated through the Polyakov–Liouville action for a two-dimensional conformal field theory with large central charge \(c\). The dilaton is treated as a purely classical field in the existing calculations; the Polyakov–Liouville term corrects the matter stress-energy but not the dilaton equation.

**Geometry and state.** The model admits eternal and evaporating Schwarzschild-like black-hole solutions. The evaporating solution is constructed perturbatively from a Minkowski initial state; infalling matter forms a black hole that shrinks and eventually meets a shockwave, after which the metric is continuously matched to a static end-state geometry that is Minkowski to first order in the perturbation.

**Physical regime and parameter hierarchy.** The calculation is perturbative in the semiclassical expansion. The central charge of the matter CFT is large, and the backreaction is retained at leading order. The black hole is taken to be large compared with the Planck scale so that the semiclassical approximation is under control. The Page curve is obtained within this perturbative regime, not as an exact non-perturbative statement.

**Observables.** The fine-grained entropy of Hawking radiation, the quantum extremal surface location, the Page time, and the Page curve are the primary R4 observables. The matter stress-energy tensor, the trace anomaly, the diffeomorphism Ward identity, and the global energy balance are the primary R1 observables.


## 4. R1 research feasibility — Semiclassical consistency

The minimum credible R1 chain for DREH is:

1. **Gravitational and matter actions.** The reduced action and the Polyakov–Liouville action are explicit in the source literature.
2. **Gravitational field equations.** Variation of the reduced action with respect to the metric and dilaton gives the two-dimensional field equations. The dilaton equation remains classical; the metric equation receives a quantum correction from the Polyakov–Liouville stress tensor.
3. **Matter stress-energy tensor.** The Polyakov–Liouville action yields a quantum correction to the energy-momentum tensor. This is the object that drives backreaction and enters the semiclassical geometry.
4. **Trace anomaly.** The Polyakov–Liouville action encodes the conformal anomaly of the two-dimensional CFT. The anomaly is state-independent and is the origin of the Hawking flux in the reduced theory.
5. **Diffeomorphism Ward identity and conservation.** The quantum stress tensor satisfies a modified conservation law inherited from the four-dimensional conservation equation after reduction. This is not identical to the classical conservation law; the modification is part of what R1 must formalize.
6. **Boundary and initial conditions.** The collapse paper specifies a Minkowski initial state and tracks the solution until the black hole forms and evaporates. The end-state matching is performed continuously across the shockwave hypersurface.
7. **Matter energy flux and black-hole mass evolution.** The flux of Hawking radiation and the shrinking of the horizon are computed from the same semiclassical solution. The black-hole mass is tied to the dilaton profile.
8. **Global energy balance.** The energy carried by the radiation is balanced against the change in the gravitational energy, including the end-state matching. This is the most delicate step and should be formalized explicitly rather than assumed from the existence of a Page curve.

The R1 chain is independently derivable at the level of the action, the field equations, and the stress tensor. The conservation identity and the global energy balance require careful attention to the reduction and to the renormalization scheme. These are imported approximations, not independent derivations, and must be flagged as such in Phys-Pseudo.


## 5. R4 research feasibility — Page-curve reconstruction

The minimum credible R4 chain for DREH is:

1. **Radiation region.** The radiation region is defined in the asymptotically flat portion of the two-dimensional spacetime. In the DREH model there is no separate non-gravitational bath; the radiation lives in the black-hole spacetime itself.
2. **Candidate island.** The island is a region behind the horizon whose endpoints are quantum extremal surfaces.
3. **Generalized entropy.** The generalized entropy is the sum of the gravitational entropy and the matter entanglement entropy. In DREH the gravitational entropy is a Wald-like Noether charge of the combined DREH–PL action; a recent Lorentzian prescription shows that this coincides with the Euclidean replica-method result.
4. **No-island branch.** At early times the island is absent and the radiation entropy follows the Hawking curve, growing monotonically.
5. **Island branch.** After the Page time the island appears and the entropy is limited by the black-hole entropy, producing the downward branch of the Page curve.
6. **QES conditions.** The quantum extremal surface is determined by extremizing the generalized entropy. The endpoint count and normalization must be checked against the dilaton profile and the topological contribution, if any.
7. **Page transition.** The transition occurs when the island branch becomes dominant. The Page time is derived or estimated within the perturbative regime.
8. **Domain of validity.** The Page curve is a leading-order semiclassical result. It is not an exact statement about the full quantum-gravity theory, and it does not by itself prove complete microscopic unitarity.

The DREH Page-curve paper demonstrates that the fine-grained entropy follows the Page curve within the island/QES prescription. The R4 chain is therefore feasible, but the entropy prescription is an imported theorem (the island formula) rather than a derived result. R4 must be formalized as a conditional reconstruction: given the island formula and the DREH background, the Page curve follows.


## 6. R1/R4 interoperability

The genuine shared objects between R1 and R4 in DREH are:

| Shared object | R1 role | R4 role | Mathematical mapping | Preserved | Lost / approximated |
|---|---|---|---|---|---|
| Dilaton profile \(\phi(x)\) | Determines gravitational field equations and black-hole mass | Determines gravitational entropy and QES locations | Same background solution | Geometry, horizon location | Quantum corrections to dilaton are neglected |
| Polyakov–Liouville stress tensor \(\langle T_{ab}\rangle\) | Drives backreaction and energy flux | Sets the matter entropy input through the conformal anomaly | Same one-loop effective action | Anomaly coefficient, central charge | State-dependent finite terms may be treated differently in R1 and R4 |
| Causal structure | Defines the evaporating geometry and end-state matching | Defines the radiation region and island endpoints | Same Penrose diagram | Horizon, singularity, asymptotic region | Perturbative matching at the shockwave is approximate |
| Central charge \(c\) | Controls backreaction strength | Controls matter entanglement entropy normalization | Same CFT parameter | Large-\(c\) scaling | Finite-\(c\) corrections are dropped |

The bridge adds mathematical information in two places. First, the dilaton profile that solves the R1 backreaction equations is the same object that enters the R4 gravitational entropy. Second, the Polyakov–Liouville stress tensor that determines the R1 energy flux is the same stress tensor whose anomaly determines the R4 matter entropy. These are not shared notation or circular assertions; they are the same field configurations and the same effective action evaluated in the same state.

The unresolved obligation is the **state-dependence of the matter entropy**. R1 computes the stress tensor as a one-loop expectation value; R4 computes the entanglement entropy of the radiation. In a conformal field theory the two are related through the anomaly and the modular Hamiltonian, but the precise map between the R1 stress tensor and the R4 entropy must be established explicitly. The DREH literature treats this map as part of the island/QES prescription, not as an independent derivation. Phys-Pseudo must record this as an imported theorem, not as a derived bridge.

A second obligation is the **perturbative order count**. R1 and R4 must use the same order in the semiclassical expansion. If the backreaction is retained at one order and the entropy at another, the Page transition may be an artifact of mismatched approximations. The DREH papers are perturbative; the formalization must make the order counting explicit.


## 7. Three Reality Checks

### RC1 — Mathematical landscape

DREH involves differential geometry (two-dimensional metric, dilaton, curvature), variational calculus (action variation), quantum field theory in curved spacetime (Polyakov–Liouville effective action, trace anomaly), conformal field theory (large central charge, entanglement entropy), operator and state theory (quantum state of the scalar field), entanglement entropy (fine-grained radiation entropy), quantum extremal surfaces (island formula), and differential equations (backreaction and collapse solution). This is a broad but bounded inventory. It is narrower than a full four-dimensional quantum-gravity calculation, but wide enough to exercise the infrastructure across several mathematical domains.

### RC2 — Mathematical intersection

The minimal demonstrated intersection needed for R1 and R4 is the pair (dilaton profile, Polyakov–Liouville stress tensor). Removing the dilaton profile removes the gravitational entropy and the QES construction. Removing the Polyakov–Liouville stress tensor removes the backreaction and the anomaly-controlled matter entropy. The intersection is not claimed to be minimal without evidence; it is the set of objects that explicitly appear in both the R1 field equations and the R4 generalized entropy. Other objects, such as the central charge and the causal structure, are supporting parameters rather than the core bridge.

### RC3 — Mathematical interoperability

The interoperability is genuine because the bridge carries field configurations and effective-action data from R1 to R4. The dilaton profile is a solution of the R1 equations, and it is inserted into the R4 entropy functional. The Polyakov–Liouville stress tensor is computed from the same effective action in both investigations. The limits are equally clear: the bridge does not carry a full quantum state of the radiation, and it does not derive the island formula from the DREH action. The island formula remains an imported assumption. The interoperability is therefore substantive but conditional.


## 8. Risk and limitation register

| Risk | Description | Severity | Mitigation |
|---|---|---|---|
| Perturbative truncation | Page curve is obtained at a definite order in the semiclassical expansion; higher-order corrections could modify the transition | HIGH | Formalize the order counting explicitly; restrict POC to the leading-order result |
| Dilaton treated classically | Quantum corrections to the dilaton equation are neglected | MEDIUM | State as a model assumption; do not claim full quantum-gravity consistency |
| Dimensional reduction | Spherical reduction is an s-wave truncation of 4D gravity; it is not a full 4D quantum-gravity calculation | MEDIUM | Explicitly label the POC as an effective 2D demonstration |
| State-dependence of entropy | R1 stress tensor and R4 entanglement entropy are not identical objects; the map is part of the island prescription | HIGH | Record as an imported theorem; do not present as a derived bridge |
| Endpoint count and normalization | The number of QES endpoints and the gravitational entropy normalization must be checked against the dilaton profile | MEDIUM | Audit against the source paper’s equations before formalization |
| End-state matching | The shockwave matching to Minkowski is first-order in perturbation theory | MEDIUM | Treat the end-state as an approximate boundary condition, not an exact result |
| Unitarity claims | The Page curve does not prove complete microscopic unitarity or resolve information recovery | HIGH | Classify the achievement as N0/N1, not as a resolution of the information paradox |

The most serious scientific risk is not that DREH is wrong; it is that the R1/R4 bridge may be overclaimed. The assignment explicitly forbids treating an imported QES theorem as an independently derived result. The formalization must distinguish the DREH-derived geometry and stress tensor from the imported island formula.


## 9. Minimum credible POC specification

The smallest meaningful scientific demonstration that exercises DREH across its mathematical structures is:

**R1 demonstration:**
- One complete action-to-field-equations chain: vary the DREH action with respect to the metric and dilaton and obtain the two-dimensional field equations.
- One stress-energy/anomaly/conservation chain: derive the Polyakov–Liouville stress tensor, identify the trace anomaly, and write the modified conservation identity.
- One backreaction or energy-balance relationship: show how the stress tensor sources the dilaton profile and how the Hawking flux changes the black-hole mass.

**R4 demonstration:**
- One generalized-entropy construction: write \(S_{\text{gen}}\) as the sum of the DREH gravitational entropy and the matter entanglement entropy.
- One no-island branch: compute the early-time entropy growth.
- One island/QES result: extremize \(S_{\text{gen}}\) and locate the quantum extremal surface.
- One branch-selection/Page-curve result: show that the island branch dominates after the Page time and produces the downward branch.

**Interoperability demonstration:**
- At least one substantive chain connecting the validated R1 dilaton profile and stress tensor to the R4 generalized entropy, with explicit assumptions and failure conditions.

This is a bounded POC. It does not require solving the full four-dimensional theory, constructing a complete quantum state of the radiation, or proving unitarity. It demonstrates breadth across meaningful mathematical structures rather than the number of equations formalized.


## 10. Handoff recommendation

**Muse (Phys-Pseudo formalization).** Formalize the DREH action, the field equations, the Polyakov–Liouville stress tensor, and the modified conservation identity as the R1 core. Formalize the generalized entropy, the no-island branch, the QES condition, and the branch-selection rule as the R4 core. Make the imported status of the island formula explicit. Represent the dilaton profile and the stress tensor as shared objects with a documented mapping and documented loss of information.

**OpenCode (challenge).** Challenge the order counting: confirm that the R1 backreaction and the R4 entropy are evaluated at the same order in the semiclassical expansion. Challenge the state-dependence of the entropy: confirm whether the R1 stress tensor and the R4 entanglement entropy are evaluated in the same quantum state. Challenge the endpoint count and the gravitational entropy normalization against the source paper. Challenge any claim that the Page curve constitutes a resolution of the information paradox.

**Reserved for later implementation.** Symbolic verification of the field equations, numerical evaluation of the Page curve, and any extension beyond the perturbative regime. The POC should not attempt to build a full quantum-gravity research system.


## 11. Prior-art and novelty status

**N0 — Reconstruction of known results.** The DREH Page curve has been published in Physical Review D. The collapse scenario and end-state matching have been published. The Wald-like entropy interpretation has been published. The POC will reconstruct these results, not discover new physics.

**N1 — More explicit or machine-checkable formalization.** The expected contribution is a formalization that makes the R1/R4 dependency chain explicit, distinguishes derived objects from imported theorems, and exposes the approximation boundaries. This is a methodological contribution to the infrastructure, not a novel physics claim. The report does not claim to solve the black-hole information problem.
