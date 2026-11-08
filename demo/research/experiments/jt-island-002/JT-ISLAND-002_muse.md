# JT-ISLAND-002 Model Selection Research Report v1.0

**Research programme:** Phys Research / Phys Infrastructure
**Primary target:** R4 — Black-hole information
**Cross-validation target:** R1 — Semiclassical matter–geometry consistency
**Research mode:** Comparative literature investigation, mathematical feasibility assessment and model selection
**Implementation authorization:** NONE
**Research through:** 8 November 2026

## 1. Executive recommendation

**Select exactly one model: RST gravity in its non-holographic, large-$N$ free-matter formulation, as used by Thomas Hartman, Edgar Shaghoulian and Andrew Strominger.**

Do not make the Gautason holographic-CFT / QRT specialization the canonical POC model. Do not select JT, DREH, or the general Fabbri– Russo family.

RST wins because it is the only candidate where the same action, same matter, same state, same coordinates and same evaporating solution feed both R1 and R4 without a change of regime, a holographic dual, or a perturbative redefinition of the background. Its limitation is also the most transparent and controllable: it is a two-dimensional toy model, not four-dimensional Schwarzschild gravity. That limitation can be stated on page one of the POC and does not move during the calculation.

Hartman et al. define the large-$N$ limit of asymptotically flat two-dimensional dilaton gravity coupled to $N$ free matter fields as a toy model for semiclassical black holes, and show that its island-rule entropy follows the unitary Page curve, in contrast to the Hawking curve obtained from asymptotic flux alone. At leading order, the island-rule entropy reproduces the Page curve for unitary information flux.

## 2. Candidate comparison

Scores are out of 10 per criterion, weighted as specified. Coherence is a hard gate.

| Candidate | Coherence 30% | Speed 25% | Breadth 15% | Traceability 10% | Clarity 10% | Reproducibility 10% | **Weighted** |
| --- | --- | --- | --- | --- | --- | --- | --- |
| **A — JT + CFT / bath** | 6 | 7 | 7 | 9 | 6 | 8 | **6.90** |
| **B — Gautason RST + holographic CFT + QRT** | 6 | 5 | 8 | 8 | 5 | 6 | **6.15** |
| **B′ — RST + $N$ free scalars, Hartman formulation — SELECTED** | 9 | 9 | 7 | 9 | 8 | 8 | **8.50** |
| **C — DREH / dimensionally reduced Schwarzschild** | 7 | 4 | 9 | 5 | 7 | 4 | **6.05** |
| **D — Fabbri– Russo / Yu– Ge generalized dilaton** | 5 | 6 | 7 | 7 | 6 | 6 | **5.95** |
| **E — Step-function Vaidya 4D** | 4 | 6 | 6 | 7 | 7 | 6 | **5.60** |

### Candidate A — JT gravity with holographic CFT matter

**Sources:** Almheiri, Engelhardt, Marolf, Maxfield (AEMM), arXiv:1905.08762; Almheiri et al. (AMMZ), arXiv:1908.10996; Penington, arXiv:1905.08255; Penington et al., arXiv:1911.11977.

AEMM examine bulk quantum effects on QESs and the entanglement wedge in a two-boundary 2D bulk defined by JT gravity coupled to a 1+1 CFT. Penington shows a phase transition in the quantum Ryu-Takayanagi surface at precisely the Page time when absorbing boundary conditions evaporate a black hole in AdS/CFT.

**Verdict: internally consistent as a toy model, not as a Schwarzschild model.** JT captures the near-horizon AdS$_2$ throat: the near-horizon geometry of extremal black holes is AdS$_2$, and the AdS$_2 \times S^{d-2}$ near-horizon geometry of near-extremal black holes is well approximated by JT gravity. JT is described as capturing dimensionally reduced near-extremal black holes through the dilaton as transverse-area fluctuation above extremality.

That is the JT-ISLAND-001 failure mode. If the POC strips the near-extremal hierarchy to keep the AMMZ Page curve, it deletes the physical justification for using JT at all. If it keeps the hierarchy, it cannot claim Schwarzschild evaporation. JT also needs a bath: effective 1+1D models such as JT coupled to a nongravitating CFT bath and RST provide calculable setups. The bath, the topological $S_0$ term, and the endpoint count are then extra structures that must be audited, not imported silently. JT remains a strong second choice as a pure toy model, but it does not beat RST on R1/R4 coherence for an evaporating, asymptotically flat demonstration.

### Candidate B — RST with holographic CFT matter, Gautason et al.

**Source:** Gautason, Schneiderbauer, Sybesma and Thorlacius, *Page Curve for an Evaporating Black Hole*, arXiv:2004.00598, JHEP 05 (2020) 091, DOI 10.1007/JHEP05(2020)091.

Gautason et al. compute a Page curve for an evaporating black hole in asymptotically flat spacetime by adapting the Quantum Ryu-Takayanagi (QRT) proposal to an analytically solvable semiclassical 2D dilaton gravity theory, finding the Page time to be one third of the black hole lifetime at leading order in semiclassical corrections. They also obtain a Page curve for a semiclassical eternal black hole with incoming flux balancing evaporation.

**Verdict: valuable cross-check, wrong canonical model.** The QRT step imports a three-dimensional gravitational dual. The POC would then have to formalize two theories (2D RST + 3D dual + holographic CFT) and a QRT theorem, while R1 would have to validate a holographic matter sector that is not the same explicit $N$ free scalars whose stress-energy feeds the RST backreaction. That is manufactured interoperability, not genuine interoperability. Keep Gautason as an independent QRT cross-check on the Page time, not as the definition of the model.

### Candidate C — DREH, dimensionally reduced Schwarzschild

**Sources:** Đorđević and Radovanović, *Page Curve for an Evaporating Schwarzschild Black Hole in Dimensionally-Reduced Model of Dilaton Gravity*, arXiv:2507.17855; *Collapse Scenario and Final State of Evaporation for Schwarzschild Black Hole in Dimensionally-Reduced Model of Dilaton Gravity*, arXiv:2506.09946.

**Citation caution:** The assignment cites the Page-curve paper as Phys. Rev. D 113, 105020 (2026). That volume/page could not be verified from APS search results in this investigation. Cite it as arXiv:2507.17855 until the APS record is checked directly. Do not present the PRD citation as verified.

DREH is the most physically ambitious candidate. It derives a 2D dilaton model from 4D Einstein-Hilbert by spherical reduction, adds Polyakov– Liouville backreaction, and studies collapse with islands. Its headline result is real: without backreaction the fine-grained entropy saturates near Bekenstein-Hawking, as for an eternal black hole. With quantum-corrected collapse, the authors claim a decreasing branch and a Page curve.

**Why it loses despite the Schwarzschild story:**

- **The decreasing branch is perturbative.** The island position at late times is $x_I = 1 - (\tilde\varepsilon/4)(\Delta\tau/2\bar a)$, linear in the small parameter, and the area term is $(1/12\tilde\varepsilon)(1 - (\tilde\varepsilon/4)(\Delta\tau/2\bar a))^2$. The Page curve therefore leans on an $\mathcal{O}(\tilde\varepsilon)$ motion. The final fine-grained entropy formula is the $\min$ of a Hawking branch and that perturbative island branch.
- **The Page time is approximation-sensitive.** It is obtained by equating the two branches, giving $\kappa\tau_P = [4(\lambda a)^2/\varepsilon](3-2\sqrt2)$.
- **The island hits the singularity.** The island appears behind the horizon, mirrors the apparent horizon, and crosses the singularity shortly before the evaporation endpoint.
- **The endpoint is outside semiclassical validity.** The authors state that near the endpoint the distance becomes comparable to the Planck length, quantum-gravity effects take place, and the semiclassical approximation breaks, so the reproduced Page curve is certain only up to its final few points.
- **Regularization is imposed, not derived.** The regularization is chosen so that fine-grained entropy vanishes, aligning with the island intersecting the singularity near the endpoint. That is a legitimate modelling choice, but for a POC it is a hidden assumption that must be carried explicitly.
- **Reproducibility is low.** Single group, 2025–2026, long perturbative expansions in $\tilde\varepsilon$, transcendental tortoise-coordinate equations, and cutoff conditions $b_*\tilde\varepsilon \gg 1$. POC effort is LARGE.

DREH should be the follow-up candidate if the programme later wants a Schwarzschild-connected demonstration and is willing to pay LARGE effort and to audit every $\mathcal{O}(\tilde\varepsilon)$ step. It is not the fastest credible POC.

### Candidate D — Fabbri– Russo generalized dilaton, Yu– Ge

**Source:** Yu and Ge, *Entanglement Islands in Generalized Two-dimensional Dilaton Black Holes*, arXiv:2208.01943, Phys. Rev. D 107, 066020 (2023), DOI 10.1103/PhysRevD.107.066020.

The Fabbri– Russo model is a generalized 2D dilaton theory with parameters $n$; RST fits $n=1$. Islands cause the rise of radiation entanglement entropy to decelerate after Page time. For eternal black holes the fine-grained entropy saturates at twice Bekenstein-Hawking; for evaporating black holes it finally reaches zero. The $n$ dependence is significant at extremely early times but subleading and exponentially suppressed at late times and large distance, so Page curves are $n$-independent at leading order.

**Verdict: useful robustness argument, not a standalone POC.** The $n$-independence result actually supports selecting $n=1$ (RST) as the canonical case. The generalized family does not provide a cleaner backreacted evaporation solution than RST, and choosing general $n$ adds a parameter without adding a POC result. Use Yu– Ge as evidence that the RST Page-curve shape is not a fine-tuned accident of $n$, not as the POC model.

### Candidate E — Other alternatives

**Step-function Vaidya, 4D:** Guo, Gan and Shu, arXiv:2302.02379, JHEP 05 (2023) 042. They choose the “in” vacuum and apply QES construction to the step-function Vaidya model of evaporating black holes in four dimensions produced by spherical null shells, obtaining metrics for three regions and studying very-late-time islands. When the cutoff is in Minkowski region III after complete evaporation, an island still exists with $\partial I$ on an equal-time Cauchy surface for $r^2(A) \ge 64 G_N\kappa_c$.

**Verdict: reject as canonical.** The metric is prescribed by shells, not solved from semiclassical backreaction. R1 cannot independently validate the same background that R4 consumes. Island location is strongly cutoff-dependent. Useful as a 4D contrast case, not as the POC.

**Bousso– Penington far-outside-horizon islands, Shu et al. one-sided asymptotically flat, and Antonini et al. apologia:** These are important reality-check literature, not POC models. Antonini et al. argue islands can appear with massless gravitons and no external reservoir, including radiation at null infinity in asymptotically flat spacetimes. That weakens the claim that islands are artefacts of a bath, but it does not supply a bounded, solvable POC action.

## 3. Canonical model definition — RST, Hartman formulation

**Action.** The RST action is the one-loop quantum-corrected version of the CGHS dilaton-gravity action. The RST theory is the quantum-corrected completion of CGHS, where the Polyakov action accounts for the conformal anomaly and the additional RST counterterm restores solvability. The key point of RST is the addition of a kinetic counterterm to the one-loop semiclassical CGHS action so that an extra symmetry is maintained at quantum level, implying exact solubility.

Schematically, in conformal gauge and with $N$ scalars $f_i$:

$$S = S_{\text{CGHS}}[\phi,g,f_i] + S_{\text{Polyakov}}[g] + S_{\text{RST}}[\phi,g]$$

where $S_{\text{CGHS}}$ contains $e^{-2\phi}(R + 4(\nabla\phi)^2 + 4\lambda^2) - \tfrac12\sum_i (\nabla f_i)^2$, $S_{\text{Polyakov}}$ is the nonlocal anomaly action, and $S_{\text{RST}}$ is the local counterterm that makes the semiclassical equations solvable. Exact coefficients, field redefinitions and gauge choices must be taken from Hartman et al. §2, not reconstructed from memory. Do not import Gautason’s holographic CFT coefficients into this action.

**Matter content.** $N$ free massless scalars $f_i$, $i=1,\dots,N$. No holographic CFT, no extra 3D bulk, no bath. The matter is the same field whose stress-energy drives backreaction and whose entanglement entropy enters R4.

**Geometry.** Asymptotically flat 2D dilaton gravity. In CGHS/RST models the vacuum is a linear-dilaton vacuum with flat metric, with dilaton charge fixed by asymptotic flatness. The evaporating solution is formed by an incoming shockwave / null shell, develops an apparent horizon, and evaporates by Hawking flux. The endpoint is a thunderpop-like final burst; the POC must stop before claiming semiclassical control at the endpoint.

**State.** “In” vacuum for the collapsing shell, with outgoing Hawking radiation at future null infinity. This is the same state for R1 flux and R4 radiation entropy. Do not switch to a thermofield-double or eternal state for R4.

**Boundary / initial conditions.** Asymptotically flat at past and future null infinity; reflecting / regular conditions at the origin in the CGHS coordinates as specified in Hartman et al. §2; cutoff surface for radiation placed at large radius, with cutoff dependence stated explicitly.

**Physical regime and parameter hierarchy.**

| Hierarchy | Meaning |
| --- | --- |
| $N \gg 1$ | Large-$N$ semiclassical expansion; $1/N$ corrections suppressed |
| $M/\lambda \gg 1$ | Large initial mass; semiclassical evaporation lasts many dynamical times |
| $N\hbar$ fixed in semiclassical limit | Backreaction is $\mathcal{O}(1)$, not $\mathcal{O}(1/N)$ |
| $\kappa \sim \lambda/N$ scale | Hawking temperature / evaporation rate set by $\lambda$ and $N$ |

**Approximations.** Large-$N$, semiclassical (quantum matter on classical geometry with backreaction), s-wave / 2D reduction, leading-order Page curve. No claim of exact all-orders unitarity.

**Relevant observables.** Dilaton $\phi$ (gravitational entropy / area), metric conformal factor $\rho$, matter stress-energy $\langle T_{\pm\pm}\rangle$, Hawking flux, radiation entropy $S(R)$, island endpoint $x_I$, Page time $t_P$.

## 4. R1 research feasibility — Semiclassical consistency

**Minimal defensible chain (all in one model):**

1. **Action to field equations.** Vary $S_{\text{RST}}$ with respect to $g_{\mu\nu}$ and $\phi$ to obtain dilaton-gravity equations + matter equations. Source: Hartman et al. §2, “Review of the RST model.” Status: independently derivable in Phys-Pseudo.
2. **Matter stress-energy.** Vary with respect to $g_{\mu\nu}$ for the $f_i$ sector. Status: derivable.
3. **Trace anomaly.** Polyakov action gives $\langle T^\mu{}_\mu\rangle \propto N R$. Status: imported theorem (anomaly), but its coefficient and use in RST must be checked against Hartman et al. §2.
4. **Diffeomorphism Ward identity / covariant conservation.** $\nabla_\mu \langle T^{\mu\nu}\rangle = 0$ up to anomaly terms, in the same coordinates as the field equations. Status: derivable as an identity, not a new physical input.
5. **Boundary / matching conditions.** Shockwave matching, asymptotically flat falloff, origin regularity. Status: must be stated explicitly; these are assumptions, not derivations.
6. **Matter energy flux.** Compute $\langle T_{++}\rangle$, $\langle T_{--}\rangle$ on the evaporating solution. Status: derivable from anomaly + conservation + boundary conditions.
7. **Mass / gravitational energy evolution.** Dilaton at apparent horizon tracks $M(u)$; $dM/du = -\text{flux}$. Status: derivable, and this is the key R1 output that R4 consumes.
8. **Global energy balance.** Initial shell energy = final radiation energy + thunderpop burst + residual, within semiclassical accuracy. Status: checkable as a consistency fixture, not assumed.

**Effort:** SMALL–MEDIUM. The RST equations are exactly solvable, so the POC does not need numerical PDE solving for the headline chain. The main work is formalizing the anomaly and the RST counterterm correctly.

## 5. R4 research feasibility — Page-curve reconstruction

**Minimal defensible chain (same model, same state):**

1. **Radiation region.** $R = [\sigma_A, \infty)$ on a late-time slice, outside a cutoff. Source: Hartman et al. §3.
2. **Candidate island.** $I$ with one QES endpoint $\partial I$ in 2D. Do not invent a second endpoint.
3. **Generalized entropy.** $S_{\text{gen}} = S_{\text{grav}}(\partial I) + S_{\text{matter}}(R \cup I)$, where $S_{\text{grav}}$ is the dilaton value at $\partial I$ (the 2D analogue of Area$/4G$).
4. **Gravitational entropy contribution.** From the same $\phi$ that appears in R1 mass evolution. This is the strongest R1→R4 bridge.
5. **Matter entropy contribution.** 2D CFT entropy formula for $R \cup I$ on the same conformal factor $\rho$ that appears in R1 stress-energy. Status: imported CFT entropy formula, but applied to the same $\rho$.
6. **No-island branch.** $S_{\text{no-island}} \sim$ Hawking growth, linear in time at leading order. This is the Hawking curve.
7. **Island branch.** $S_{\text{island}} \sim$ dilaton at QES + matter term, decreasing as the black hole shrinks.
8. **QES conditions.** Extremize $S_{\text{gen}}$ with respect to $\partial I$ location. Status: derivable from the explicit $S_{\text{gen}}$.
9. **Endpoint count and normalization.** One QES endpoint; no JT $S_0$ topological term. The POC must state this explicitly and not import $S_0$ from JT.
10. **Branch selection.** $S(R) = \min\{S_{\text{no-island}}, S_{\text{island}}\}$.
11. **Page transition.** Crossing of branches. Gautason’s QRT calculation gives Page time at one third of lifetime at leading order, as an independent cross-check, not as the canonical derivation.
12. **Domain of validity.** Leading-order, large-$N$, large-mass, before the thunderpop endpoint.

**Effort:** SMALL–MEDIUM. The island calculation is analytic at leading order. The POC must not claim an exact all-orders Page curve.

## 6. R1/R4 interoperability

### Shared-object and dependency map

| Shared object | R1 produces | R4 consumes | Bridge status |
| --- | --- | --- | --- |
| Action $S_{\text{RST}}$ | Field equations, anomaly | $S_{\text{gen}}$ background | **Genuine** — same action |
| Dilaton $\phi$ | Horizon mass, $M(u)$ | Gravitational entropy at QES | **Genuine** — same field, same solution |
| Metric conformal factor $\rho$ | Stress-energy, flux | Matter entropy on $R \cup I$ | **Genuine** — same $\rho$ |
| Matter state (in vacuum + shockwave) | $\langle T_{\mu\nu}\rangle$ | $S_{\text{matter}}$ | **Genuine** — same state, same coordinates |
| Evaporating geometry $x^\pm$ | Apparent horizon, flux | Island location, Page time | **Genuine** — same solution |
| Cutoff $\sigma_A$ | Flux at cutoff | Radiation region $R$ | **Genuine, but cutoff-dependent** — must be stated |
| Island prescription $S(R)=\min\ \text{ext}\ S_{\text{gen}}$ | — | Defines R4 | **Imported** — not derived from RST action in POC |

**Genuine bridges (add mathematical information):**

- **Bridge 1 — Dilaton = area = mass.** The same $\phi$ that tracks $M(u)$ in R1 gives $S_{\text{grav}}$ in R4. If R1 mass evolution is wrong, R4 island branch is wrong. That is a testable dependency.
- **Bridge 2 — Conformal factor = flux = entropy.** The same $\rho$ that determines $\langle T_{\pm\pm}\rangle$ in R1 determines the CFT entropy in R4. Changing $\rho$ moves both flux and Page time coherently.
- **Bridge 3 — State = flux = radiation.** The same “in” vacuum + shockwave that produces Hawking flux in R1 produces the radiation whose entropy R4 computes. No state switch between R1 and R4.

**Unresolved obligations:**

- The island formula itself is an imported prescription. The replica-wormhole derivation (Goto– Hartman– Tajdini for JT, Almheiri et al. for replica wormholes) is not re-derived in the POC. It must be labelled “imported.”
- The QRT cross-check is imported holography, not a derivation from RST.
- The thunderpop endpoint is outside semiclassical control. The POC must stop the Page curve before the endpoint and state that the final burst is not semiclassically resolved.
- Cutoff dependence: Page time and island location depend on $\sigma_A$. The POC must show that dependence, not hide it.

If any of these obligations are hidden, the POC fails the cross-test.

## 7. Three Reality Checks

### RC1 — Mathematical Landscape

Genuine structures: differential geometry, dilaton gravity, variational calculus, QFT in curved spacetime, 2D CFT, Polyakov effective action, renormalization / conformal anomaly, entanglement entropy, QES, differential equations and backreaction.

Not genuinely needed for the canonical POC: 3D holography, QRT, SYK, ensemble averaging, replica wormholes as a derivation. Those are imported or cross-check structures. Do not count them as POC breadth.

### RC2 — Mathematical Intersection

Minimal intersections needed:

- **Variational calculus ∩ anomaly:** Without the Polyakov term, there is no backreaction and no evaporation. Removing it collapses R1 to CGHS without Hawking flux.
- **Anomaly ∩ conservation:** Without conservation + anomaly + boundary conditions, flux is undetermined. Removing any one makes $dM/du$ ambiguous.
- **Dilaton ∩ entropy:** Without $\phi$ as area, there is no island branch. Removing the dilaton–area identification removes R4.
- **CFT entropy ∩ geometry:** Without the same $\rho$ in flux and entropy, R1 and R4 decouple.

Removing the holographic dual does not break the POC. Removing the RST counterterm does: it breaks exact solvability.

### RC3 — Mathematical Interoperability

**Pass, with one imported bridge.** R1 and R4 share action, field, state, coordinates and solution. The bridges in §6 are genuine derivations, not shared notation. The only non-genuine bridge is the island prescription itself, which is explicitly labelled imported. That is acceptable for a reconstruction POC, provided the label is kept. It would not be acceptable to claim the POC derives the island rule from RST.

## 8. Risk and limitation register

| Risk | Severity | Mitigation |
| --- | --- | --- |
| **RST is a toy model, not 4D Schwarzschild** | High for impact, low for coherence | State on page one; do not claim 4D Schwarzschild |
| **Large-$N$ limit** | Medium | State $1/N$ corrections are outside POC; do not claim exactness |
| **Island formula imported** | Medium | Label as imported; do not claim derivation from RST |
| **QRT / holographic CFT conflation** | High if allowed | Keep Gautason as cross-check only; canonical matter is $N$ free scalars |
| **Cutoff dependence** | Medium | Show Page time vs. cutoff; do not quote a single cutoff-free Page time |
| **Thunderpop endpoint** | Medium | Stop Page curve before endpoint; do not claim semiclassical endpoint resolution |
| **Gauss-law / massless-gravity objection** | Medium | Acknowledge Geng et al.: islands are disconnected wedges, and in long-range gravity the lack of an asymptotic region creates a Gauss-law puzzle, suggesting islands might not be consistent wedges in standard massless gravity. Counter with Antonini et al. apologia, but do not claim the debate is settled. |
| **DREH temptation** | Medium | Do not switch to DREH mid-POC for Schwarzschild impact; DREH is a separate, LARGE follow-up |
| **JT $S_0$ / bath conflation** | Medium | Do not import JT $S_0$, bath, or near-extremal hierarchy into RST |
| **DREH PRD citation unverified** | Low for POC, high for bibliography | Verify APS record before citing DREH as published |


## 9. Minimum credible POC specification

**Effort: SMALL–MEDIUM for analytic reconstruction. LARGE if QRT holography or DREH perturbative audit is added. Do not add them.**

**R1 demonstration (3 items):**

- One complete action-to-field-equations chain for RST.
- One stress-energy / anomaly / conservation chain, yielding Hawking flux.
- One backreaction / energy-balance relationship, $dM/du = -\text{flux}$, with explicit boundary conditions.

**R4 demonstration (4 items):**

- One generalized-entropy construction, $S_{\text{gen}} = S_{\text{grav}} + S_{\text{matter}}$.
- One no-island branch (Hawking growth).
- One island / QES result (extremization, island location).
- One branch-selection / Page-curve result, with domain of validity and cutoff dependence.

**Interoperability demonstration (1 item):**

- One substantive chain: R1 dilaton + conformal factor + state → R4 $S_{\text{grav}}$ + $S_{\text{matter}}$ → Page transition, with assumptions, preserved properties, lost information (island prescription imported, $1/N$ corrections dropped, endpoint excluded), and failure conditions.

**Explicitly out of scope:** Replica-wormhole derivation, QRT 3D dual derivation, SYK / ensemble duality, DREH perturbative audit, 4D Vaidya backreaction, thunderpop microphysics, firewall / monogamy resolution.

**What can be reconstructed analytically:** RST field equations, flux, $S_{\text{gen}}$, QES location, leading-order Page curve.
**What must be imported:** Anomaly coefficient, CFT entropy formula, island prescription, Gautason Page-time cross-check.
**What is suitable for later symbolic / computational verification:** Flux vs. $dM/du$ balance, QES extremization, branch crossing as a function of cutoff.

## 10. Handoff recommendation

**Muse should formalize in Phys-Pseudo:**

- RST action, field redefinitions, and field equations in one fixed gauge and coordinate system.
- Anomaly + conservation + flux chain, with boundary conditions as explicit assumptions.
- $S_{\text{gen}}$, QES extremization, no-island and island branches, and branch selection.
- The §6 dependency map, with imported items labelled imported.

**OpenCode must challenge:**

- Is the RST counterterm coefficient correct and is exact solvability actually used, or merely asserted?
- Is the matter state in R1 identical to the matter state in R4, in the same coordinates?
- Is $S_{\text{grav}}$ in R4 the same dilaton that tracks $M(u)$ in R1, with the same normalization?
- Where does cutoff dependence enter, and does the Page time move with it?
- Is the island prescription labelled imported, and is any conclusion used as a premise?
- Is any JT, Gautason holographic, or DREH equation smuggled in?

**Reserve for later implementation:**

- Symbolic verification of flux vs. mass loss.
- Numerical QES extremization as a function of cutoff.
- QRT cross-check as a separate, optional module, not as canonical POC.
- DREH audit as a separate research task, with its own perturbative-order tracking.

## 11. Prior-art and novelty status

**Classification: N0 — Reconstruction of known results, plus N1 — More explicit or machine-checkable formalization.**

No novel physics is claimed. The POC reconstructs Hartman et al.’s RST island Page curve and Gautason et al.’s Page-time cross-check, with R1 and R4 made explicit and interoperable in one model. The novelty, if any, is in the explicitness, traceability and challengeability of the formalization, not in the physics.

**Bottom line:** One physical model — RST with $N$ free scalars. Two connected investigations — R1 semiclassical consistency and R4 island Page curve — sharing action, dilaton, conformal factor, state and solution. Multiple mathematical structures — variational calculus, anomaly, CFT entropy, QES — without manufactured interoperability.



ar5iv.labs.arxiv.org
Page Curve for an Evaporating Schwarzschild Black Hole in Dimensionally-Reduced Model of Dilaton Gravity

librarysearch.hillsdale.edu
The entropy of bulk quantum fields and the entanglement wedge of an evaporating black hole - Hillsdale College

arxiv.org
[2004.00598] Page Curve for an Evaporating Black Hole

arxiv.org
Page curves and Entanglement Islands for the Step-Function Vaidya Model of Evaporating Black Holes

docslib.org
Arxiv:2004.13857V1 [Hep-Th] 28 Apr 2020 Contents

arxiv.org
Entanglement Revivals and Scrambling for Evaporating Black Holes

arxiv.org
Black hole mining in the RST model

arxiv.org
An apologia for islands

ouci.dntb.gov.ua
OUCI

arxiv.org
[2208.01943] Entanglement Islands in Generalized Two-dimensional Dilaton Black Holes

arxiv.org
[2107.03390] Inconsistency of Islands in Theories with Long-Range Gravity

arxiv.org
[2004.13857] Islands in Asymptotically Flat 2D Gravity

ncatlab.org
extremal black hole in nLab

readkong.com

researchgate.net
Page curve (assuming a large initial mass M 1).

arxiv.org
Contents

archive.org
Full text of "AdS_2 and quantum stability in the CGHS model"

arxiv.org
arXiv:hep-th/9606097v3 13 Oct 1997
