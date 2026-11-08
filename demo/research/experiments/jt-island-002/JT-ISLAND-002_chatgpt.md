# JT-ISLAND-002 — Model Selection Research Report v1.0

Research cutoff: November 8, 2026 · Scope: R4 black-hole information and R1 semiclassical matter–geometry consistency

## Preliminary research finding

The investigation has identified two particularly strong competitors:

1. RST with holographic CFT matter: An established, analytically tractable semiclassical evaporation model with a published QRT Page-curve calculation. Its published result places the Page time at approximately one-third of the black-hole lifetime, at leading order in semiclassical corrections.&#x20;

   [image](https://www.google.com/s2/favicons?domain=https://link.springer.com\&sz=32)

   link.springer.com+1

2. Dimensionally reduced Einstein–Hilbert/Schwarzschild dilaton gravity (DREH): A newer model with an unusually direct connection to four-dimensional spherical gravity. Both its collapse/backreaction study and its Page-curve study appeared in Physical Review D in May 2026.&#x20;

   [image](https://www.google.com/s2/favicons?domain=https://journals.aps.org\&sz=32)

   journals.aps.org+1



An important distinction is already apparent: a model's relationship to four-dimensional Schwarzschild geometry is not the same thing as having the strongest independently reproducible semiclassical and entropy calculations.

The decisive question is whether the same semiclassical solution, matter state, boundary conditions and approximation scheme actually feed both R1 and R4. Merely finding a Page curve in a model that also possesses semiclassical field equations is insufficient.

## 1. Candidate C: the DREH model under closer examination

Strong scientific candidate

Primary papers

- Stefan Đorđević and Voja Radovanović, Collapse Scenario and Final State of Evaporation for Schwarzschild Black Hole in Dimensionally-Reduced Model of Dilaton Gravity, arXiv:2506.09946; Physical Review D 113, 105010 (May 14, 2026). DOI: [10.1103/q2n7-qnc1](https://doi.org/10.1103/q2n7-qnc1?utm_source=chatgpt.com).
- Stefan Đorđević and Voja Radovanović, Page Curve for an Evaporating Schwarzschild Black Hole in Dimensionally-Reduced Model of Dilaton Gravity, arXiv:2507.17855; Physical Review D 113, 105020 (May 29, 2026). DOI: [10.1103/nc2z-hx6s](https://doi.org/10.1103/nc2z-hx6s?utm_source=chatgpt.com).&#x20;

  [image](https://www.google.com/s2/favicons?domain=https://journals.aps.org\&sz=32)

  Phys. Rev. D+2



### 1.1 Why the model is attractive

The gravitational action comes directly from spherical reduction of the four-dimensional Einstein–Hilbert action:

\\[ S\_{\phi}=\frac{1}{4G}\int d^2x\sqrt{-g} \left[ e^{-2\phi}\left(R+2(\nabla\phi)^2\right) +2\lambda^2 \right]. \\]

The reduced matter sector contains a minimally coupled massless scalar:

\\[ S_m=-\frac12\int d^2x\sqrt{-g}(\nabla f)^2. \\]

The quantum correction is introduced through the Polyakov–Liouville effective action, whose local auxiliary-field representation is

\\[ S\_{\rm PL}=-\frac{\hbar}{96\pi} \int d^2x\sqrt{-g} \left[2R\psi+(\nabla\psi)^2\right], \qquad \Box\psi=R. \\]

These are not equations borrowed from an unrelated JT or CGHS model. They belong to the same spherically reduced gravitational framework.&#x20;

[image](https://www.google.com/s2/favicons?domain=https://arxiv.org\&sz=32)

arxiv.org+2



The dilaton has a direct geometric interpretation: it encodes the radius of the spherical orbits removed in the reduction. This makes the relation between the two-dimensional field equations and four-dimensional spherical geometry particularly useful for the infrastructure demonstration.

### 1.2 R1 is unusually well connected to the original collapse problem

The backreaction paper supplies:

- The reduced gravitational equations.
- The classical scalar-field source.
- The Polyakov–Liouville quantum correction.
- State-dependent stress-energy terms.
- Quantum-corrected equations of motion.
- A perturbative collapse and evaporation solution.

Its state choice is also identifiable. The initial geometry is Minkowski spacetime, and the quantum state is chosen as the Unruh state, with \\(t\_\pm(\sigma^\pm)=0\\) in the initial asymptotically flat coordinates. The later coordinate transformation produces the nonzero outgoing energy flux.&#x20;

[image](https://www.google.com/s2/favicons?domain=https://arxiv.org\&sz=32)

arxiv.org+1



One particularly useful consistency chain is:

\\[ \text{Polyakov effective action} \rightarrow \langle T\_{\mu\nu}\rangle \rightarrow \text{semiclassical field equations} \rightarrow \text{evaporating geometry}. \\]

The paper gives the state-dependent stress-tensor components in Eqs. (57)–(58), quantum-corrected field equations in Eqs. (59)–(61), and the first-order expansion in Eqs. (62)–(65).&#x20;

[image](https://www.google.com/s2/favicons?domain=https://arxiv.org\&sz=32)

arxiv.org+2



This makes DREH a promising R1 fixture. It also gives us a way to test whether the geometric and matter equations agree independently of the entropy calculation.

### 1.3 R4 has a significant qualification

The Page-curve paper explicitly compares two cases.

First, the classical collapse geometry without quantum backreaction. Its island calculation produces an entropy that saturates at the Bekenstein–Hawking entropy, rather than the desired complete Page curve.

Second, the quantum-corrected collapse geometry, where the backreaction changes the late-time solution. The authors obtain the rising and falling fine-grained entropy branches and calculate a Page transition.&#x20;

[image](https://www.google.com/s2/favicons?domain=https://arxiv.org\&sz=32)

arxiv.org+1



This distinction is important: the Page curve is not simply a consequence of applying an island formula to any Schwarzschild-like solution. The particular quantum-corrected background matters.

There are also two unusually important audit points:

First: the endpoint regularization. The paper explicitly says that its selected regularization ensures the fine-grained entropy vanishes at the evaporation endpoint. This is a legitimate prescription to investigate, but the endpoint purity must not subsequently be presented as independent evidence when that condition was used to select the prescription. The POC must reproduce the result and separately challenge its renormalization dependence.&#x20;

[image](https://www.google.com/s2/favicons?domain=https://arxiv.org\&sz=32)

arxiv.org+1



Second: the approximation. The quantum-corrected geometry and island position are treated perturbatively. The QES position is only established to the order supported by that solution. We cannot describe the resulting Page curve as an exact solution of the complete quantum theory.&#x20;

[image](https://www.google.com/s2/favicons?domain=https://arxiv.org\&sz=32)

arxiv.org+2



My assessment: DREH has the strongest physical connection to four-dimensional Schwarzschild gravity among the required candidates, but its entropy prescription deserves a more demanding independent audit than the abstract alone suggests.

## 2. Candidate B: RST gravity with holographic CFT matter

Leading POC candidate

The strongest case for RST is not simply that it is analytically solvable. It is that one can trace a continuous chain from its semiclassical equations to the geometry used by its Page-curve calculation, without importing the gravitational dynamics from another model.

### 2.1 Canonical model

The starting CGHS dilaton-gravity action is

\\[ I\_{\mathrm{CGHS}}= \frac{1}{2\pi}\int d^2x\sqrt{-g}\\, e^{-2\phi} \left[ R+4(\nabla\phi)^2+4\lambda^2 \right]. \\]

The matter sector is a strongly coupled two-dimensional holographic CFT with central charge \\(c\\), assumed to have an AdS\\(\_3\\) gravitational dual.

The semiclassical correction contains the Polyakov term,

\\[ I_Q=-\frac{c}{12\pi}\int d^2x\\, \partial\_+\rho\\,\partial\_-\rho, \\]

and the RST counterterm,

\\[ I\_{\mathrm{RST}}= \frac{c}{48\pi}\int d^2x\sqrt{-g}\\,\phi R. \\]

The RST term is added to make the semiclassical equations analytically tractable. It is generally covariant, contributes at the same order as the Polyakov term, and does not change the classical limit. This is an explicit model choice, not an identity that follows from the original CGHS action alone.&#x20;

[image](https://www.google.com/s2/favicons?domain=https://arxiv.org\&sz=32)

arxiv.org+1



The semiclassical equations simplify after defining

\\[ \Omega=e^{-2\phi}+\frac{c}{24}\phi. \\]

In Kruskal coordinates, they become

\\[ \partial\_+\partial\_-\Omega+1=0, \qquad -\partial\_\pm^2\Omega =\frac{c}{24}t\_\pm. \\]

The functions \\(t\_\pm\\) encode the quantum state and transform under coordinate changes with the Schwarzian contribution.&#x20;

[image](https://www.google.com/s2/favicons?domain=https://arxiv.org\&sz=32)

arxiv.org+1



### 2.2 R1: a reproducible consistency chain

The CFT trace anomaly and the state-dependent stress tensor are given by

\\[ \langle T^\mu{}\_\mu\rangle=\frac{c}{12}R, \\]

\\[ \langle T\_{+-}\rangle =-\frac{c}{6}\partial\_+\partial\_-\rho, \\]

\\[ \langle T\_{\pm\pm}\rangle =\frac{c}{12} \left[ 2\partial\_\pm^2\rho -2(\partial\_\pm\rho)^2-t\_\pm \right]. \\]

The paper explicitly derives these components by integrating the two-dimensional conservation equation with the trace anomaly as input. The coordinate transformation law is also specified, so the state-dependent terms cannot be treated as arbitrary constants.&#x20;

[image](https://www.google.com/s2/favicons?domain=https://arxiv.org\&sz=32)

arxiv.org+1



For a collapse geometry, a short incoming matter pulse is sent into the linear-dilaton vacuum. The semiclassical exterior solution is

\\[ ds^2=-M e^{2\phi}\\,dv\\,du, \\]

\\[ \Omega=M\left[ 1-v(u+1)-\epsilon\log(-Mvu) \right]. \\]

This is particularly useful for R1/R4 integration: the same quantum-corrected solution is the background subsequently used for the Page calculation.&#x20;

[image](https://www.google.com/s2/favicons?domain=https://arxiv.org\&sz=32)

arxiv.org



The asymptotic outgoing flux is explicitly calculated:

\\[ T\_{--}(\sigma^-) = \frac{c}{24} \left[ 1-\frac{1}{(1+e^{\sigma^-})^2} \right]. \\]

The authors identify it with Hawking radiation and relate the flux to black-hole mass loss. The leading lifetime is \\(1/\epsilon\\).&#x20;

[image](https://www.google.com/s2/favicons?domain=https://arxiv.org\&sz=32)

arxiv.org



R1 verdict: strong. The action, trace anomaly, stress tensor, state functions, semiclassical field equations, collapse solution and outgoing flux are all available in one primary source. The POC still needs to reconstruct the Ward identity and energy balance independently rather than merely quote them.

### 2.3 R4: the Page-curve calculation

The authors use

\\[ S\_{\mathrm{gen}} = \frac{\mathrm{Area}(I)}{4G_N} + S\_{\mathrm{Bulk}}. \\]

The gravitational contribution in the RST model is

\\[ \frac{\mathrm{Area}(I)}{4G_N} = 2\left[\Omega(I)-\Omega\_{\mathrm{crit}}\right]. \\]

The subtraction ensures that the island contribution vanishes when the island is absent.&#x20;

[image](https://www.google.com/s2/favicons?domain=https://arxiv.org\&sz=32)

arxiv.org



The bulk entropy is supplied by the holographic matter sector. The CFT entropy is calculated through the AdS\\(\_3\\)/CFT\\(\_2\\) Ryu–Takayanagi construction. The calculation compares the no-island branch against the branch with an island and extremizes the generalized entropy.&#x20;

[image](https://www.google.com/s2/favicons?domain=https://arxiv.org\&sz=32)

arxiv.org+1



For the dynamical black hole, the island saddle is approximated near the horizon. The calculation uses a large logarithm to simplify the QES equations and checks that the approximation is self-consistent. The resulting leading-order island entropy is

\\[ S\_{\mathrm{gen}}^{\mathrm{island}} = 2M-\frac{c}{24}(t_A-\sigma_A)+\cdots. \\]

The competing no-island entropy grows with retarded time. Their crossing gives

\\[ t\_{\mathrm{Page}} = \frac{1}{3\epsilon} = \frac{16M}{c}. \\]

This is one-third of the leading-order evaporation lifetime.&#x20;

[image](https://www.google.com/s2/favicons?domain=https://arxiv.org\&sz=32)

arxiv.org+2



The entropy is calculated for radiation on both sides of the black hole. The endpoint count and normalization therefore need to be preserved exactly; halving the radiation region without changing the formulas would invalidate the reconstruction.

### 2.4 What is assumed rather than independently derived?

There are three crucial limitations:

1. QRT is an input. The authors explicitly state that they obtain the Page curve by assuming the Quantum Ryu–Takayanagi prescription. The POC can reconstruct and challenge its application, but should not claim to derive QRT from the RST model.
2. The matter sector is specialized. The strongly coupled large-\\(c\\) CFT and its AdS\\(\_3\\) dual are part of the chosen model. This is not equivalent to replacing it with \\(N\\) free scalar fields.
3. The endpoint is outside the reliable semiclassical regime. The authors say their calculation is reliable only while the black-hole mass remains large compared with the scale set by the CFT central charge.&#x20;

   [image](https://www.google.com/s2/favicons?domain=https://arxiv.org\&sz=32)

   arxiv.org+1



That last limitation matters: the Page transition is derived in the controlled semiclassical regime, but the complete endpoint of evaporation is not established by this calculation.

## 3. Evaluation of the remaining candidates

### Candidate A — JT gravity with holographic CFT matter

The original JT research is important, but the two cited papers should not be treated as interchangeable descriptions of a single evaporation experiment.

AEMM, The entropy of bulk quantum fields and the entanglement wedge of an evaporating black hole, JHEP 12 (2019) 063, arXiv:1905.08762, supplies a detailed JT evaporation setup, including boundary energy evolution and a coupling to an external bath. Its construction uses a near-extremal black hole initially in a Hartle–Hawking state. The near-extremal hierarchy, including the large extremal-entropy contribution, is part of the physical setup.&#x20;

[image](https://www.google.com/s2/favicons?domain=https://arxiv.org\&sz=32)

arxiv.org+3



AMMZ, The Page curve of Hawking radiation from semiclassical geometry, JHEP 03 (2020) 149, arXiv:1908.10996, develops the holographic treatment of matter entropy and the QES construction, using a higher-dimensional gravitational dual. Its bath and holographic construction must be matched to the specific evaporation protocol before its entropy results are combined with AEMM's energy and state calculations.&#x20;

[image](https://www.google.com/s2/favicons?domain=https://arxiv.org\&sz=32)

arxiv.org+2



Assessment: JT remains theoretically valuable, but it is not the best next POC foundation under this assignment. The JT-ISLAND-001 incompatibility cannot be resolved by simply importing AEMM's semiclassical hierarchy and AMMZ's Page-curve machinery as though they were already a common calculation. A single canonical JT setup would first need to demonstrate a compatible parameter window, state, bath, and endpoint normalization. I therefore place this candidate on hold rather than declare JT invalid.

### Candidate D — Fabbri–Russo generalized dilaton gravity

Yu and Ge, Entanglement Islands in Generalized Two-dimensional Dilaton Black Holes, arXiv:2208.01943, study a one-parameter family of exactly solvable models. Their action contains \\(N\\) massless scalar fields, and the \\(n=1\\) member is the RST model. The paper derives a free-field representation, identifies a critical dilaton boundary and calculates island Page curves.&#x20;

[image](https://www.google.com/s2/favicons?domain=https://arxiv.org\&sz=32)

arxiv.org+1



Its main advantage is mathematical generality. Its disadvantage for this POC is that the extra parameter \\(n\\) adds model space without solving the central R1/R4 interoperability question better than the established RST construction.

Assessment: use \\(n=1\\) as a useful RST consistency cross-check, not as the primary model. Do not tune \\(n\\) to obtain a desired entropy curve.

### Additional alternatives

Two recent developments deserve attention:

- Hartman, Shaghoulian and Strominger, Islands in Asymptotically Flat 2D Gravity, JHEP 07 (2020) 022, arXiv:2004.13857, give an RST calculation with \\(N\\) free scalar fields. This is a distinct matter realization from the holographic-CFT RST model. Its most important warning is that the directly computed asymptotic entropy follows the Hawking curve, while the island-rule entropy follows a Page curve. The authors also do not construct the replica-wormhole saddles globally. This is an excellent adversarial fixture, but it must not be merged with the holographic-CFT equations.&#x20;

  [image](https://www.google.com/s2/favicons?domain=https://arxiv.org\&sz=32)

  arxiv.org+3

- Easson, Fate of Schwarzschild–de Sitter black holes: Nonequilibrium evaporation, PRD 113, 084014 (April 7, 2026), DOI 10.1103/trlf-d6sk, is a strong new R1 alternative. It derives a spherical reduction of four-dimensional Schwarzschild–de Sitter gravity, includes Polyakov backreaction, and studies the Unruh–de Sitter state, conserved flux and mass evolution. However, its Page curve is described as a thermodynamic estimate with QES/island features, rather than the fully reconstructed collapse QES calculation required here. It does not yet replace the required R4 chain.&#x20;

  [image](https://www.google.com/s2/favicons?domain=https://journals.aps.org\&sz=32)

  Phys. Rev. D+1



A 2026 finite-\\(N\\) JT paper by Snigdha Saha also claims to smooth the sharp Page transition through nonperturbative saddle interference. It may be relevant to a later R4 refinement, but it does not by itself supply the independent semiclassical matter–geometry consistency chain required by R1.&#x20;

[image](https://www.google.com/s2/favicons?domain=https://repo.scoap3.org\&sz=32)

repo.scoap3.org



## 4. Comparative decision

The weighted scores below are research judgments on the assignment's six criteria, not objective measurements. A candidate still has to pass the scientific-coherence gate regardless of its total.

| Candidate                   | R1/R4 coherence /30 | POC speed /25 | Math breadth /15 | Sources /10 | Audience /10 | Reproducibility /10 | Total |
| --------------------------- | ------------------- | ------------- | ---------------- | ----------- | ------------ | ------------------- | ----- |
| A — JT + holographic CFT    | 19                  | 18            | 13               | 9           | 8            | 6                   | 73    |
| B — RST + holographic CFT   | 28                  | 22            | 13               | 9           | 8            | 8                   | 88    |
| C — DREH Schwarzschild      | 26                  | 17            | 15               | 8           | 10           | 6                   | 82    |
| D — Fabbri–Russo family     | 21                  | 17            | 12               | 7           | 7            | 6                   | 70    |
| E — Schwarzschild–de Sitter | 15                  | 17            | 14               | 9           | 8            | 5                   | 68    |

Decision: select Candidate B — RST gravity with a large-central-charge holographic CFT.

It has the best combination of a reproducible semiclassical collapse solution, an explicitly calculated pair of competing entropy branches, an analytically identified Page transition, and a bounded mathematical workload. Its main limitation—reliance on QRT—is explicit, well localized and suitable for an adversarial test.&#x20;

[image](https://www.google.com/s2/favicons?domain=https://link.springer.com\&sz=32)

Springer Nature Link+4



Candidate C is the runner-up, not a rejected theory. It has the more compelling Schwarzschild connection and the richer direct bridge to four-dimensional spherical geometry. But its perturbative QES, endpoint regularization and loss of semiclassical control near the final evaporation point make it a less secure first demonstration. The scalar field is also a minimally coupled two-dimensional matter model; the gravitational action's spherical reduction does not, by itself, establish a complete four-dimensional matter reduction.&#x20;

[image](https://www.google.com/s2/favicons?domain=https://journals.aps.org\&sz=32)

Phys. Rev. D+4



The selection is therefore not “RST is more fundamental than DREH.” It is: RST is the better first model for demonstrating a connected and independently challengeable R1/R4 workflow.

## 5. Canonical model definition

The selected model is specifically the asymptotically flat, evaporating RST black hole coupled to a large-\\(c\\), strongly coupled holographic CFT, as treated by Gautason et al. It is not generic JT gravity, free-scalar RST, or DREH.

### Physical specification

| Element                              | Canonical choice                                                                  |
| ------------------------------------ | --------------------------------------------------------------------------------- |
| Gravity                              | Two-dimensional CGHS dilaton gravity with the RST correction                      |
| Matter                               | Strongly coupled holographic CFT with central charge \\(c\\)                      |
| Quantum backreaction                 | Polyakov effective action                                                         |
| Geometry                             | Asymptotically flat, two-dimensional dilaton black hole                           |
| Initial state                        | Linear-dilaton vacuum                                                             |
| Collapse                             | Short incoming matter pulse                                                       |
| Matter state for entropy calculation | Coherent state built on the inertial vacuum at past null infinity                 |
| Entropy prescription                 | QRT, with an AdS\\(\_3\\) gravitational dual                                      |
| Radiation observation                | Timelike anchor curve at large, approximately constant \\(\Omega_A\\)             |
| Regime                               | Large initial mass, large central charge and controlled semiclassical corrections |

The model uses the hierarchy

\\[ c\gg 24, \qquad \epsilon=\frac{c}{48M}\ll1, \qquad M\gg c. \\]

The characteristic scale is set to \\(\lambda=1\\) by rescaling coordinates. The published calculation retains the leading nontrivial semiclassical corrections in \\(\epsilon\\).&#x20;

[image](https://www.google.com/s2/favicons?domain=https://arxiv.org\&sz=32)

arxiv.org+1



The initial geometry is the linear-dilaton vacuum. A short incoming pulse at \\(v=1\\) forms the black hole. The matter CFT is not simply in an arbitrary vacuum on the evolving geometry: the authors use its coherent-state construction and coordinates in which the incoming CFT is in its vacuum state to calculate its entropy. The entropy equivalence of that coherent state and the vacuum is an explicit assumption in the construction.&#x20;

[image](https://www.google.com/s2/favicons?domain=https://arxiv.org\&sz=32)

arXiv+2



### What this model is not

It is not a full four-dimensional Schwarzschild calculation. The gravitational sector is two-dimensional dilaton gravity. Its matter sector has a holographic AdS\\(\_3\\) dual, used to evaluate entanglement entropy.

The extra dimension in the holographic construction should not be confused with a claim that the black hole itself has been solved as a complete higher-dimensional gravitational system.

This distinction is essential to the eventual public explanation.

## 6. R1 feasibility — the minimum independent semiclassical reconstruction

The R1 investigation should not begin by copying the final semiclassical solution. It should reconstruct the connections in the following order.

### R1.1 Action to field equations

Start with the classical dilaton-gravity action, matter sector, Polyakov contribution and RST correction.

The matter quantum effects enter through the trace anomaly. The RST correction is what makes the semiclassical field equations analytically tractable.

The resulting equations in Kruskal coordinates are

\\[ \partial\_+\partial\_-\Omega+1=0, \\]

\\[ -\partial\_\pm^2\Omega = \frac{c}{24}t\_\pm. \\]

These equations should be independently recovered by varying the effective action, including the constraints. The state functions \\(t\_\pm\\) must remain explicit rather than being absorbed into arbitrary constants.&#x20;

[image](https://www.google.com/s2/favicons?domain=https://arxiv.org\&sz=32)

arXiv+1



### R1.2 Anomaly to stress-energy consistency

The stress tensor must be reconstructed with both its trace and its state-dependent components.

The essential checks are:

\\[ \langle T^\mu{}\_\mu\rangle = \frac{c}{12}R, \\]

and, in the paper's conventions,

\\[ \langle T\_{+-}\rangle = -\frac{c}{6}\partial\_+\partial\_-\rho, \\]

\\[ \langle T\_{\pm\pm}\rangle = \frac{c}{12} \left[ 2\partial\_\pm^2\rho -2(\partial\_\pm\rho)^2-t\_\pm \right]. \\]

The reconstruction must establish:

- The trace anomaly.
- The diffeomorphism Ward identity.
- The conservation equation.
- The transformation of \\(t\_\pm\\) under conformal coordinate changes.
- The Schwarzian contribution to the stress tensor.

These are separate mathematical obligations. The anomaly is not itself the conservation equation, and the state-dependent functions are not arbitrary additions to the theory.&#x20;

[image](https://www.google.com/s2/favicons?domain=https://arxiv.org\&sz=32)

arxiv.org+1



### R1.3 Backreaction and evaporation

Use the initial vacuum and incoming pulse to reproduce the semiclassical collapse solution, including the exterior geometry:

\\[ \Omega = M\left[ 1-v(u+1)-\epsilon\log(-Mvu) \right]. \\]

Then independently calculate the asymptotic Hawking flux and verify the leading lifetime.

The paper obtains

\\[ T\_{--}(\sigma^-) = \frac{c}{24} \left[ 1-\frac{1}{(1+e^{\sigma^-})^2} \right]. \\]

The energy accounting should connect the outgoing flux to the loss of black-hole energy. The paper gives the mass-energy relation

\\[ \mathcal M=\frac{\lambda M}{\pi} \\]

and the entropy and temperature relations

\\[ S\_{\mathrm{BH}}=2M, \qquad T=\frac{\lambda}{2\pi}. \\]

Together with the first law, these provide an independent route to the leading evaporation timescale.&#x20;

[image](https://www.google.com/s2/favicons?domain=https://arxiv.org\&sz=32)

arXiv+2



R1 assessment: The required core is reconstructable. The remaining issue is how to treat the strong-coupling boundary at \\(\Omega=\Omega\_{\mathrm{crit}}\\) and the very final stage of evaporation. The calculation should not claim complete endpoint energy accounting from a semiclassical solution that has already lost control.

## 7. R4 feasibility and the R1/R4 interoperability test

This is the decisive part of the assignment.

### R4.1 Define the entropy problem

Use the dynamical collapse solution, not the eternal black-hole solution.

Define a radiation region \\(R\\) collected at the distant anchor \\(A\\). Compare:

No-island configuration

The generalized entropy contains no gravitational island contribution. The bulk CFT entropy is calculated using the no-island region.

Island configuration

The island \\(I\\) is included, with its endpoint determined by quantum extremality. The generalized entropy is

\\[ S\_{\mathrm{gen}} = \frac{\mathrm{Area}(I)}{4G_N} + S\_{\mathrm{Bulk}}(I\cup R). \\]

For the selected RST normalization, the gravitational contribution is

\\[ \frac{\mathrm{Area}(I)}{4G_N} = 2\left[\Omega(I)-\Omega\_{\mathrm{crit}}\right]. \\]

This is the normalization that must be retained. It must not be replaced by the JT entropy expression or by an independently chosen constant entropy term.&#x20;

[image](https://www.google.com/s2/favicons?domain=https://arxiv.org\&sz=32)

arXiv+1



### R4.2 Endpoint count and normalization

For the dynamical collapse calculation, there is one island and one gravitating QES endpoint.

The matter entropy is calculated for the interval connecting the island and the radiation anchor. That interval has two matter-theory endpoints, but this does not mean that there are two independent gravitational area contributions.

The no-island branch instead reaches the semiclassical boundary at

\\[ \Omega=\Omega\_{\mathrm{crit}}. \\]

There is no independent JT-style topological \\(S_0\\) term in the selected CGHS/RST action. The Page-curve reconstruction must not import one from Candidate A.&#x20;

[image](https://www.google.com/s2/favicons?domain=https://arxiv.org\&sz=32)

arXiv+1



### R4.3 QES conditions and Page transition

The QES position is obtained by extremizing the island generalized entropy with respect to its coordinates.

The relevant published calculation gives equations (59)–(60), reduces them using the large-logarithm approximation, and obtains the near-horizon island solution in equations (61)–(65).

The two entropy branches are then compared. Their crossing yields

\\[ t\_{\mathrm{Page}} = \frac{1}{3\epsilon} = \frac{16M}{c}. \\]

With the leading lifetime \\(t\_{\mathrm{lifetime}}=1/\epsilon\\),

\\[ \boxed{ t\_{\mathrm{Page}} = \frac{1}{3}t\_{\mathrm{lifetime}} } \\]

The result is a leading-order Page transition, not an exact solution at arbitrary mass or at the final endpoint.&#x20;

[image](https://www.google.com/s2/favicons?domain=https://arxiv.org\&sz=32)

arxiv.org+2



### R1/R4: the actual mathematical bridge

R1 — Validated semiclassical inputs

\\((g\_{\mu\nu},\phi,\Omega)\\), CFT central charge \\(c\\), stress tensor, state functions \\(t\_\pm\\), collapse solution and boundary conditions.

Shared physical configuration

The same quantum-corrected collapse geometry and same incoming matter state are used to construct the entropy calculation.

R4 — Generalized-entropy inputs

The dilaton value at the island determines the gravitational contribution. The CFT state and conformal coordinates determine the bulk entropy through the AdS\\(\_3\\)/CFT\\(\_2\\) prescription.

QES and branch selection

Extremize each candidate entropy, compare no-island and island branches, and reconstruct the Page transition.

The substantive interoperability is in the shared geometry and state. The R1 solution supplies the dilaton and metric that enter the gravitational entropy term and the conformal transformation used to evaluate CFT matter entropy.

The connection preserves the central charge, the chosen state correlations and the geometry used in the entropy calculation. It does not preserve every microscopic property of the holographic CFT: the entropy calculation uses its universal holographic description.

The bridge is therefore genuine but conditional. It requires the QRT prescription and the specified matter state. The result does not follow from the semiclassical field equations alone.&#x20;

[image](https://www.google.com/s2/favicons?domain=https://arxiv.org\&sz=32)

arXiv+3



## 8. Three Reality Checks

### RC1 — Mathematical landscape

The selected model exercises several genuinely different mathematical structures.

| Mathematical structure                   | Actual role                                        |
| ---------------------------------------- | -------------------------------------------------- |
| Differential geometry                    | Metric, curvature and covariant field equations    |
| Dilaton gravity                          | Gravitational action and black-hole geometry       |
| Variational calculus                     | Derivation of field equations                      |
| Quantum field theory in curved spacetime | Matter stress-energy and Hawking flux              |
| Conformal field theory                   | Matter sector and state-dependent entropy          |
| Trace anomaly                            | Quantum contribution to the stress tensor          |
| Polyakov effective action                | Semiclassical backreaction                         |
| Schwarzian derivatives                   | State transformation between conformal coordinates |
| Differential equations                   | Collapse solution and evaporation geometry         |
| AdS\\(\_3\\)/CFT\\(\_2\\)                | Holographic calculation of matter entropy          |
| Ryu–Takayanagi/QRT                       | Bulk entropy prescription                          |
| Quantum extremal surfaces                | Extremization of generalized entropy               |
| Entanglement entropy                     | Competing radiation entropy branches               |

These structures play different roles. For example, the Polyakov action contributes to the semiclassical dynamics, whereas the RT prescription evaluates the matter entropy. They are related through the same CFT, not interchangeable mathematical objects.

### RC2 — Mathematical intersection

The minimal demonstrated intersections are:

1. Gravity and matter: The stress-energy tensor sources the gravitational equations.
2. Quantum fields and geometry: The trace anomaly and its conservation identity determine the state-dependent quantum stress tensor.
3. Coordinates and states: The Schwarzian transformation connects the quantum state description in different conformal frames.
4. Geometry and entropy: The dilaton at the QES determines the gravitational entropy contribution.
5. CFT and holography: The matter entropy is calculated through the AdS\\(\_3\\) dual.
6. Entropy and variational calculus: Extremization and minimization select the QES and dominant branch.

These intersections are necessary for the demonstration, but I would not yet claim they are the mathematically minimal possible set.

The removal tests are particularly informative:

- Remove the anomaly: the Hawking backreaction chain is lost.
- Remove the RST correction: the simple analytic semiclassical solution is lost.
- Remove the holographic CFT assumption: the selected RT calculation no longer follows.
- Change the incoming state: the stress tensor and entropy inputs must be recalculated.
- Remove the island branch: the calculation retains the monotonically growing no-island result.
- Relax the large-mass hierarchy: the published semiclassical approximation is no longer controlled.

### RC3 — Mathematical interoperability

Demonstrated at leading semiclassical order

The bridge between geometry, matter and entropy is valid within the stated model and approximations.

However, it does not establish:

- A microscopic reconstruction of the CFT state from the Page curve.
- A derivation of QRT from the RST field equations.
- A general theorem that any semiclassical evaporating geometry produces a Page curve.
- A complete four-dimensional account of black-hole information recovery.

Those are outside the evidence supplied by the selected model.

## 9. Risk and limitation register

| Risk                                                      | Severity | Required treatment                                                                               |
| --------------------------------------------------------- | -------- | ------------------------------------------------------------------------------------------------ |
| QRT prescription is imported                              | High     | Label it as an assumption and challenge its application                                          |
| Holographic large-\\(c\\) matter sector                   | High     | Keep the matter model fixed                                                                      |
| Strong-coupling boundary at \\(\Omega\_{\mathrm{crit}}\\) | High     | Explicitly identify boundary prescription                                                        |
| Semiclassical endpoint                                    | High     | Do not claim control of final evaporation                                                        |
| Large-\\(M\\), small-\\(\epsilon\\) hierarchy             | High     | Check every approximation against it                                                             |
| Matter state and coordinate choices                       | High     | Preserve the state functions and transformations                                                 |
| Two-sided versus collapse normalization                   | High     | Use the dynamical collapse branch only                                                           |
| Greybody factors                                          | Medium   | State that the two-dimensional conformal treatment omits them                                    |
| Microscopic unitarity                                     | High     | Treat Page-curve reconstruction as a result of the prescription, not proof of complete unitarity |

A particularly important adversarial comparison is Hartman–Shaghoulian–Strominger's free-scalar RST result.

Their direct asymptotic entropy calculation follows the Hawking curve, even though their separately defined island-rule entropy follows the Page curve. This demonstrates why a POC must identify exactly which entropy is being computed. It must not silently identify every island-rule Page curve with a directly measured microscopic radiation entropy.&#x20;

[image](https://www.google.com/s2/favicons?domain=https://arxiv.org\&sz=32)

arXiv+2



The selected holographic-CFT RST model is distinct from that free-scalar model. The latter is a valuable adversarial comparison, not an additional source of equations for the selected model.

## 10. Minimum credible POC specification

I would constrain the POC to four deliverables.

Deliverable A — R1 action and consistency chain

Reconstruct the effective action, field equations, trace anomaly, stress tensor, Ward identity, state transformation, collapse solution, asymptotic flux and mass-loss relation.

Deliverable B — R4 entropy chain

Reconstruct the radiation region, no-island configuration, generalized-entropy expression, gravitational entropy normalization, matter entropy, QES equations, island branch and Page transition.

Deliverable C — R1/R4 cross-validation

Demonstrate that the geometry, CFT state, central charge and coordinate conventions used in R4 are exactly those validated by R1. Any imported result must be explicitly identified.

Deliverable D — adversarial controls

At minimum, challenge:

- The sign and normalization of the trace anomaly and stress tensor.
- The Schwarzian transformation and state functions.
- The mass-loss calculation and evaporation timescale.
- The single-QES area contribution and UV divergence cancellation.
- The approximation used to solve the island equations.
- The no-island boundary prescription at \\(\Omega\_{\mathrm{crit}}\\).
- The sensitivity of the Page transition to the anchor and boundary choices.

The POC should first reconstruct the analytical results. Numerical plotting or symbolic checks can support the work, but should not substitute for the derivations.

## 11. Handoff recommendation

### Muse — formalization assignment

Muse should formalize only the selected Gautason et al. RST holographic-CFT construction.

The formalization should preserve the paper's distinction between:

- Classical gravitational action.
- Semiclassical effective corrections.
- Matter stress-energy.
- Quantum state and boundary functions.
- Collapse geometry.
- Bulk entropy.
- Gravitational entropy.
- QES conditions.
- Entropy branch selection.

Every formalized claim should identify its source equation, assumptions, approximation order, state and applicable region.

Do not import the DREH equations, free-scalar RST formulas, or the JT near-extremal hierarchy into this specification.

### OpenCode — adversarial verification assignment

OpenCode should challenge the formalization by independently:

1. Varying the action and recovering the gravitational equations.
2. Deriving the stress tensor from the trace anomaly and conservation law.
3. Checking the conformal transformation and state functions.
4. Reconstructing the collapse solution from its initial conditions.
5. Reproducing the outgoing flux and mass-loss relationship.
6. Verifying the generalized entropy normalization and endpoint count.
7. Solving the QES equations independently.
8. Checking the no-island and island branches and their crossing.
9. Testing the approximation and boundary assumptions.
10. Identifying every conclusion that depends on imported QRT results.

A calculation should fail verification if it silently changes the matter sector, state, endpoint count, normalization or physical regime.

### Work to reserve for later

Do not expand this first POC into a nonperturbative RST completion, exact microscopic reconstruction, finite-\\(c\\) holographic correction, or full four-dimensional generalization.

The 2026 finite-\\(N\\) JT work is potentially relevant to future Page-transition refinements, but it is not needed to establish this first R1/R4 demonstration.&#x20;

[image](https://www.google.com/s2/favicons?domain=https://repo.scoap3.org\&sz=32)

repo.scoap3.org



### Prior-art and novelty status

N0 — Reconstruction of known results

The Page curve, semiclassical RST equations, QRT application and one-third-lifetime result are established results in the cited literature.

The platform's achievement would be a reproducible, linked and independently challengeable formalization of those results.

A later machine-checkable reconstruction could qualify as N1 — More explicit or machine-checkable formalization, provided it genuinely adds independent verification and not simply a machine-readable transcription.

There is no basis here for claiming novel physics.

## 12. Primary-source register

These are the central papers and the most useful equation/section locators for the selected research programme.

| Source                                                                                                                                                    | Publication and identifier                                                                                                                                                                        | Key material                                                          |
| --------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------- |
| Gautason, Schneiderbauer, Sybesma and Thorlacius, Page Curve for an Evaporating Black Hole                                                                | JHEP 05 (2020) 091; DOI: [10.1007/JHEP05(2020)091](https://doi.org/10.1007/JHEP05\(2020\)091?utm_source=chatgpt.com); [arXiv:2004.00598](https://arxiv.org/abs/2004.00598?utm_source=chatgpt.com) | Eqs. (4), (26)–(35), (39)–(42), (58)–(68)                             |
| Hartman, Shaghoulian and Strominger, Islands in Asymptotically Flat 2D Gravity                                                                            | JHEP 07 (2020) 022; DOI: [10.1007/JHEP07(2020)022](https://doi.org/10.1007/JHEP07\(2020\)022?utm_source=chatgpt.com); [arXiv:2004.13857](https://arxiv.org/abs/2004.13857?utm_source=chatgpt.com) | RST free-scalar comparison; Eqs. (2.1)–(2.10), (3.1)–(3.17)           |
| Almheiri, Engelhardt, Marolf and Maxfield, The Entropy of Bulk Quantum Fields and the Entanglement Wedge of an Evaporating Black Hole                     | JHEP 12 (2019) 063; DOI: [10.1007/JHEP12(2019)063](https://doi.org/10.1007/JHEP12\(2019\)063?utm_source=chatgpt.com); [arXiv:1905.08762](https://arxiv.org/abs/1905.08762?utm_source=chatgpt.com) | JT near-extremal model, bath coupling and boundary energy balance     |
| Almheiri, Mahajan, Maldacena and Zhao, The Page Curve of Hawking Radiation from Semiclassical Geometry                                                    | JHEP 03 (2020) 149; DOI: [10.1007/JHEP03(2020)149](https://doi.org/10.1007/JHEP03\(2020\)149?utm_source=chatgpt.com); [arXiv:1908.10996](https://arxiv.org/abs/1908.10996?utm_source=chatgpt.com) | JT/holographic-CFT and QES construction                               |
| Đorđević and Radovanović, Collapse Scenario and Final State of Evaporation for Schwarzschild Black Hole in Dimensionally-Reduced Model of Dilaton Gravity | PRD 113, 105010 (2026); DOI: [10.1103/q2n7-qnc1](https://doi.org/10.1103/q2n7-qnc1?utm_source=chatgpt.com); [arXiv:2506.09946](https://arxiv.org/abs/2506.09946?utm_source=chatgpt.com)           | Spherical reduction, semiclassical collapse and energy flux           |
| Đorđević and Radovanović, Page Curve for an Evaporating Schwarzschild Black Hole in Dimensionally-Reduced Model of Dilaton Gravity                        | PRD 113, 105020 (2026); DOI: [10.1103/nc2z-hx6s](https://doi.org/10.1103/nc2z-hx6s?utm_source=chatgpt.com); [arXiv:2507.17855](https://arxiv.org/abs/2507.17855?utm_source=chatgpt.com)           | Perturbative Page curve; Eqs. (114)–(116)                             |
| Yu and Ge, Entanglement Islands in Generalized Two-dimensional Dilaton Black Holes                                                                        | PRD 107, 066020 (2023); [arXiv:2208.01943](https://arxiv.org/abs/2208.01943?utm_source=chatgpt.com)                                                                                               | Fabbri–Russo family and RST special case                              |
| Easson, Fate of Schwarzschild–de Sitter Black Holes: Nonequilibrium Evaporation                                                                           | PRD 113, 084014 (2026); DOI: [10.1103/trlf-d6sk](https://doi.org/10.1103/trlf-d6sk?utm_source=chatgpt.com)                                                                                        | Recent spherical-reduction and semiclassical-backreaction alternative |

## Final recommendation

GO — Candidate B, with explicit QRT limitation

Proceed with RST gravity coupled to a large-\\(c\\) holographic CFT as the foundation for JT-ISLAND-002's first proof of concept.

It offers the strongest balance between model coherence, mathematical breadth, primary-source traceability, and implementation-independent reproducibility.

Keep DREH as the strongest competing model. Its direct connection to four-dimensional spherical gravity makes it particularly attractive, but it should not displace the RST model until its perturbative QES, entropy regularization and endpoint assumptions have survived independent reconstruction.

The scientific claim of the POC should be limited to:

We can reconstruct and independently challenge a published semiclassical black-hole evaporation model, connect its matter–geometry consistency equations to its QES entropy calculation, and reproduce its leading-order Page transition without mixing different physical models.

That is already a substantial and defensible demonstration.
