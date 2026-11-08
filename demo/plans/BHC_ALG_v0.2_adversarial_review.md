# Adversarial review: BHC + ALG cores and MRC coverage matrix v0.2

**Review date:** 2026-11-06  
**Reviewed artifact:** `CORES_AND_COVERAGE_v0.2.md`  
**Review status:** **NOT READY FOR RESEARCH-AGENT EXECUTION**  
**Scope:** mathematical validity, source fidelity, internal consistency, pre-registration quality, branch comparability, and coverage-rung justification. This is a review of the draft as written, not an endorsement of BHC or ALG.

## Executive verdict

The v0.2 draft has a strong research-design idea: compare two genuinely different cores against common physical targets, preserve negative outcomes, record assumptions and silence classes, and prevent post-hoc changes to the test. The interface-contract proposal is also promising.

However, the current draft has several **blocking defects**. Most importantly:

1. Several target results are already included as branch assumptions and then counted as explanations by the coverage matrix. This makes parts of the forecast circular.
2. The matter–spacetime–dynamics interface contracts conflate different mathematical maps and use incorrect or overstrong limit statements.
3. BHC's actual-configuration source tensor is not specified sufficiently to be tested or falsified, and its conservation condition is still an unresolved structural obstacle.
4. ALG merges distinct crossed-product constructions and overgeneralizes their type and entropy conclusions.
5. The measurement and clock pilots conflate probability representation, Born-rule justification, outcome actualization, and relational evolution.
6. Several pilots are compound claims but PR4 requires exactly one terminal state per pilot.
7. The two branches do not yet have comparable auxiliary-dependence accounting.

**Recommendation:** retain v0.2 as a planning draft, but do not hand it to agents as a frozen execution contract. First produce v0.2.1 with the repairs below. Do not change BHC's historical record or silently promote any branch.

---

## 1. Critical blocker: circularity between core assumptions and coverage claims

PR1 says each branch core is exactly its postulate list. The matrix then uses consequences already stated in those lists as evidence that the branch explains a pilot.

### ALG circularities

- **Q2** assumes that the crossed-product algebra is Type II and that generalized entropy is recovered up to a constant. R4 then proposes to count the generalized-entropy result as something ALG explains from Q2. If Q2 remains a core postulate, the entropy relation is postulated, not derived or explained.
- **Q3** assumes a Gleason-type uniqueness result. R3 then proposes to count Born-form weights on the invariant algebra as explained by Q3. The mathematical extension theorem can reconstruct a linear functional from an already supplied bounded finitely additive projection measure; it does not derive that laboratory outcome frequencies obey that measure.
- **Q4** assumes evolution is conditioning on the internal clock. R2 then tests the same Page–Wootters-style relation and forecasts EXPLAINS-to-DERIVES because the clock is built in. A premise cannot count as a derived result merely because it is in the core.
- **Q1** already declares the invariant observable algebra to be a crossed product. That is acceptable as a proposed construction, but the precise action, group, representation, and invariant subalgebra must be specified before the construction can support Q2.

### BHC circularity

- **P2** includes equivariance of the Psi-density. R3 then proposes to derive Born weights using equivariance plus typicality. If P2 already assumes the equilibrium density and its equivariance, an important part of the Born route is already loaded into the core. The continuity/equivariance result should instead be derived from a specific wave equation and guidance law.
- **P5** assumes equilibrium record statistics are independent of foliation and clock within K, while R2 is intended to test that proposition. This can be a falsifiable core commitment, but it must be labelled explicitly as a hypothesis under test, not treated as a fact available to the derivation.

### Required repair

Separate four classes of statements:

1. **Core postulates:** physical commitments defining the branch.
2. **Imported mathematical theorems:** established results with exact hypotheses and citations.
3. **Bridge conjectures / implementation hypotheses:** propositions the pilot is meant to test.
4. **Pilot targets:** conclusions not already assumed by the branch.

Every theorem used by a pilot must have its hypotheses checked in the selected model. If a target is retained as a core postulate, its terminal status cannot be DERIVES or EXPLAINS; record POSTULATES (or test whether the core is internally inconsistent).

---

## 2. Interface contracts: promising idea, but the mathematics is currently mis-typed

### I_MS: matter ↔ spacetime

A locally covariant QFT is represented by a functor from an appropriately defined category of spacetimes (typically oriented, time-oriented, globally hyperbolic Lorentzian spacetimes with causally convex isometric embeddings) to a category of algebras. Functoriality alone does **not** automatically supply Einstein causality, the time-slice axiom, a preferred state, or self-consistent metric backreaction; these are additional axioms or structures.

The draft's statement that the functor preserves locality and covariance is too compressed. Specify at least:

- the spacetime category and its morphisms;
- the algebra category and *-homomorphisms;
- isotony/locality (Einstein causality);
- the time-slice property;
- the state-space assignment, which is naturally contravariant;
- how a metric perturbation or variation is represented;
- whether and how the stress tensor is obtained from metric variation.

The claim that the interface "loses the state-dependence needed for backreaction" is not right as a statement about locally covariant QFT. A state may be supplied on each algebra. What the fixed-background functor does not by itself provide is a **dynamical feedback equation** determining the metric from the state or its renormalized stress tensor.

The proposed roundtrip is not actually a roundtrip: a spacetime-to-algebra functor need not have an inverse that reconstructs the spacetime. Rename `roundtrip` to `reference_case` or `consistency_fixture` unless an actual inverse/reconstruction theorem is supplied.

### I_SD: spacetime ↔ dynamics

The ADM hypersurface-deformation constraints do not simply mean "evolution is gauge." Their Poisson brackets close with structure functions, and lapse/shift encode the embedding of spatial hypersurfaces. The contract should distinguish:

- gauge transformations generated by constraints;
- changes of foliation/embedding;
- physical relational evolution with respect to a clock;
- the initial-value problem and reconstruction of spacetime solutions.

The ADM reference case needs explicit assumptions: suitable initial data satisfying the constraints, lapse/shift, regularity, boundary conditions, and the relevant equivalence to the Einstein equations.

### I_MD: matter ↔ dynamics

Stone's theorem is narrower than the current contract suggests. It relates a strongly continuous one-parameter unitary group indexed by real time to a self-adjoint generator. It does not identify every algebra automorphism flow with a Hamiltonian in the chosen representation, and it does not turn a Hamiltonian constraint into physical time evolution.

The line "lost: anything not generated by a self-adjoint operator" should be removed. It also excludes open-system dynamics, which are often described by completely positive maps, instruments, or semigroups rather than reversible unitary groups. Expand `LANG_D` to distinguish:

- reversible unitary dynamics;
- algebra automorphism groups (including potentially outer automorphisms);
- constrained/gauge dynamics;
- open-system CP maps and measurement instruments.

### G2: overlap limits need correction

- `G_N -> 0` does **not** automatically give flat-space QFT. With a fixed curved background it gives a decoupling/fixed-background limit: QFT on that specified background. Minkowski QFT follows only if the background is separately chosen to be Minkowski (or another limit is derived).
- `hbar -> 0` does **not** generically give GR with the same matter. A controlled classical limit usually needs specified coherent/WKB states, scaling of fields and parameters, and a regime in which fluctuations are suppressed.
- Since `G_N` and `hbar` are dimensional in ordinary units, state the dimensionless ratios and scale hierarchy used in each limit. In units `c=hbar=1`, still declare the characteristic scale against which `G_N` is compared.

Do not describe the two limits as automatic inverse translations. Test each limit in a specified model.

---

## 3. BHC core: P4 is too unconstrained to be a falsifiable source rule

P4 states

`G_mu_nu + Lambda g_mu_nu = 8 pi G S_mu_nu[Q, Psi]`

and describes S as the stress-energy of the actual configuration, including "any guidance contribution." This does not define a determinate mathematical family. The guidance contribution could be local or nonlocal, state-dependent, metric-dependent, constrained, or an arbitrary compensating tensor. Without a fixed ansatz or admissible class, P4 can be changed to repair almost any conservation failure.

The condition

`CONSISTENT := ∇^mu S_mu_nu = 0 on solutions of P2`

is necessary for the displayed Einstein equation with constant Lambda, by the contracted Bianchi identity. But merely defining the condition does not establish it.

For a scalar field, the ordinary classical stress tensor has the off-shell identity

`∇^mu T_mu_nu = (□_g φ - V'(φ)) ∇_nu φ`

(up to the action/signature conventions). A Bohmian guidance equation with a quantum-force term does not generally make this divergence vanish. Quantum-potential corrections or extra exchange/source terms require their own derivation and may not automatically restore conservation.

### Required BHC repair

Before R1 begins, specify:

1. the matter ontology and configuration space (particle, scalar field, or another definite choice);
2. the wave equation and its Hilbert/state space;
3. the guidance law and foliation dependence;
4. the exact stress/source tensor, including every guidance or quantum-potential term;
5. the allowed locality, covariance, differentiability, and dependence on `g` and `Psi`;
6. whether additional fields or exchange currents are allowed;
7. the exact conservation identity and whether it is off-shell or on-shell.

Freeze one concrete source family or a finite pre-registered family of candidates. Do not let an agent invent a new compensating tensor after seeing a failed divergence calculation. If the family is not fixed, R1 is BLOCKED before any mathematical verdict; "no source in P4 family" is not a valid falsifier when the family itself is undefined.

Also clarify whether `S_mu_nu` depends on `g`; the current notation omits it even though a covariant stress tensor and its divergence are metric-dependent.

P5 is a strong additional empirical/mathematical requirement. Define "record," "equilibrium statistics," foliation changes, and admissible clock class K operationally. Statistics may be foliation-independent even when individual Bohmian trajectories depend on a foliation; the claim must be phrased and tested at the level of specified observable distributions.

---

## 4. ALG core: three distinct crossed-product constructions are being merged

The current Q1/Q2 wording treats Witten, CLPW, and JSS as one generic construction. They should be kept separate.

- **Witten (2021/2022):** in the large-N setting described in the paper, an emergent Type III_1 algebra is crossed with its modular automorphism group to obtain a Type II_infinity algebra.
- **CLPW (2023):** a specific observer-dressed static-patch algebra in de Sitter is Type II_1, with a maximum-entropy state; the entropy of semiclassical states agrees with generalized entropy up to a state-independent additive constant.
- **Jensen–Sorce–Speranza (2023):** a construction for compact gravitational subregions introduces an observer and argues for Type II_1; it relies on well-motivated but conjectural states whose modular flow is geometric at an instant in time. More general subregions/asymptotic settings can lead to Type II algebras that are not all Type II_1.

These are related but not interchangeable. In particular, "cross product of A_loc with Gdress" is not a sufficient definition. A crossed product needs a specified von Neumann algebra, a specified group, a specified action of that group on the algebra, and a representation/construction. A gravitational constraint/dressing group is not automatically the modular automorphism group used in Witten's theorem.

### Required ALG repair

Make Q1 a construction registry with separately named cases, e.g.:

- `ALG-W`: Type III_1 crossed product by modular `R`-action (Type II_infinity target);
- `ALG-CLPW`: de Sitter static-patch observer-dressed construction (Type II_1 target);
- `ALG-JSS`: compact-subregion observer construction (Type II_1 under stated assumptions), with conjectural modular-flow input explicitly recorded.

Do not combine them into a single theorem until an equivalence or common generalization is actually proved.

### Trace and entropy correction

A Type II factor has a faithful normal semifinite trace. Only Type II_1 has a finite normalized tracial state with `tau(1)=1`; Type II_infinity has an infinite semifinite trace. The existence of a trace does not by itself make every state's entropy finite. The state must be represented by an appropriate density with respect to the trace, and finiteness/integrability of the entropy expression must be checked.

The generalized-entropy correspondence is a result of the specific constructions and assumptions—not a generic consequence of any Type II crossed product.

### ALG frame-register correction

`ALG_F1` should not say that ALG discards a primitive state: Q3 explicitly retains states as positive normalized functionals. The intended change is that the Hilbert-space state vector is no longer the sole primitive representation; the observable algebra and its states are treated together. Likewise, `ALG_F2` should say that ALG rejects an external/preferred background time as fundamental, not that it discards time as a parameter altogether: its automorphism flows and clock-conditioned descriptions remain parameterized.

### Q3 correction

The Mackey–Gleason–Bunce–Wright theorem says, under its hypotheses, that a bounded finitely additive measure on the projection lattice of a von Neumann algebra with no Type I_2 direct summand has a unique bounded linear extension. This is a state-reconstruction/extension theorem. It does not derive the empirical Born rule from no probability assumption.

The draft must distinguish:

- original Gleason assumptions (including dimension/countable-additivity structure);
- effect/POVM extensions;
- Mackey–Gleason–Bunce–Wright bounded finite-additivity result;
- the additional physical premise identifying the measure with laboratory outcome frequencies.

A Type II or III **factor** has no Type I_2 direct summand, but this does not mean it has no embedded qubit subalgebra. It means the direct-summand obstruction in that theorem is absent. Do not say that the qubit exception or qubit subsystems disappear.

### D3 correction

"Type II or III algebras have no pure states" is too broad. They have no minimal projections and no normal pure states; that does not rule out non-normal pure states on the algebra considered as a C*-algebra. More importantly, a Bohmian beable need not be a minimal projection. The actual compatibility question is whether the theory supplies a suitable commutative configuration algebra/PVM (possibly diffuse), a configuration space, a probability measure, and a guidance law compatible with the gravitational constraints and observer dressing.

Replace D3 with that question. Do not treat the absence of minimal projections as a no-go result.

---

## 5. R2: the observer clock and Page–Wootters are not yet the same construction

Q4 says evolution is conditioning on an internal clock. The R2 pilot then studies a constraint `C = H_clock + H_sys`, a physical inner product, and a map to clock-conditioned states. That is a Page–Wootters-style construction. It is not automatically supplied by the observer degree of freedom used in JSS/CLPW.

The bounded-below energy condition is precisely where ideal-clock assumptions need care. A semibounded clock generally cannot have an ideal self-adjoint time operator canonically conjugate to its Hamiltonian on the same Hilbert space. POVM time observables and approximate/quasi-ideal clocks are possible, but their conditional evolution can have deviations and must be calculated.

The R2 test also cannot be literally identical for BHC and ALG:

- BHC starts with a foliation/time parameter and Bohmian trajectories. It needs a test of whether **specified equilibrium record statistics** are invariant under a change of foliation/clock. It does not automatically possess the same physical-inner-product construction.
- ALG needs a physical Hilbert space/representation, a constraint, a clock POVM or other time observable, conditional states, and a theorem or error bound for relational evolution.

The phrase "Both branches share the test" should mean a shared operational question only. Give each branch its own mathematical test and pass criterion.

Also, failure of essential self-adjointness on one chosen dense domain does not automatically mean no self-adjoint realization exists; it may mean multiple self-adjoint extensions. Conversely, even a self-adjoint constraint may have continuous zero spectrum and no normalizable kernel, requiring a specified physical inner product/group-averaging or rigged-Hilbert-space construction. BLOCKED(domain) should refer to the absence of a specified well-defined realization/physical inner product, not merely a failed essential-self-adjointness check on an arbitrary core.

The ALG R2 forecast "EXPLAINS to DERIVES (semi-bounded clock is built in)" is not supported. Bounded energy is an input, not a proof of good relational evolution or clock-independence.

---

## 6. R3: probability weights, outcomes, and records must be three separate targets

The attachment and the matrix correctly identify that the Born-weight question and the actualization question are different. The v0.2 matrix still needs a third distinction: stable records.

### Born weights

- For BHC, derive the continuity/equivariance equation from the actual wave equation and guidance law. Do not put equivariance of the Born density into P2 if R3 is meant to test it. Then separately state the quantum-equilibrium/typicality measure and derive the frequency result under the specified measurement model.
- For ALG, MGBW reconstructs a linear state functional from a bounded finitely additive projection measure. It does not itself identify that measure with experimental frequencies. The physical probability postulate must be declared, or a separate argument supplied.
- A Bohmian nonequilibrium distribution on configuration space is not, by itself, a violation of the Gleason/MGBW hypotheses on a projection lattice. It changes the relation between the wavefunction and observed configuration frequencies. R3_q2 currently mixes these levels.

### Outcomes and records

- BHC's actual configuration gives an ontology in which a pointer has an actual configuration. But P1 alone does not derive a measurement model, pointer stability, effective subsystem wavefunction, decoherence, or the mapping from configurations to recorded outcomes. Classify the bare definiteness as a core commitment/postulate; test stable record formation separately.
- ALG remains silent on actualization unless an ontology or collapse mechanism is added. The absence of minimal projections is not the reason for that silence.
- The measurement trilemma needs explicit definitions of completeness, linear/unitary evolution, and single definite outcomes, plus the measurement-interaction assumptions. Everett is better described as retaining unitary dynamics and the universal state while rejecting a unique single outcome as fundamental (or treating outcomes as branch-relative), rather than simply "denying B."

### Required repair

Split R3 into at least:

- `R3-W`: probability weights / Born-form reconstruction;
- `R3-O`: outcome definiteness / actualization;
- `R3-R`: stable records and pointer dynamics.

Give each its own terminal result. Do not count all three as one "measurement" victory.

---

## 7. R4: generalized entropy does not by itself derive the island formula or Page curve

The R4 distinction between entropy and the Page curve is good and should be retained. But several points need tightening.

1. A Type II trace/entropy result in ALG is not the island formula, quantum-extremal-surface extremization, replica-wormhole saddle, or microscopic reconstruction. These must be derived separately.
2. CLPW/JSS entropy results and island/Page-curve calculations may involve different algebras, models, and assumptions. Do not assume that a generalized-entropy expression from one automatically provides the generalized entropy functional or extremization rule required by another.
3. JT gravity plus a CFT bath is a lower-dimensional dilaton-gravity benchmark, not automatically a direct test of a four-dimensional BHC or generic ALG core. State the dimension, dilaton, matter, bath, state, and boundary conditions; construct an explicit map/embedding from each branch to the benchmark.
4. State precisely which steps use a gravitational path integral over replica geometries and which are valid only within the chosen semiclassical/large-parameter approximation. A formal power series in `G_N` may not contain nonperturbative saddle contributions; specify the approximation being tested.
5. The BHC condition "if QES formula is needed as AUX then POSTULATES" does not match the rung definitions. If a standard framework result is imported as one allowed AUX, the result may be EXPLAINS under the stated rule; if it is independently assumed as the target formula, it is POSTULATES. Choose one and apply it consistently.
6. Failure to reproduce one island result does not automatically refute every classical-metric branch. REFUTES(P3) requires a demonstration that the specified benchmark and its agreed assumptions force a feature incompatible with P3, and that no alternative route within the frozen BHC core can reproduce the target.

ALG's R4 forecast should be split into `R4-E` (algebraic/generalized entropy), `R4-I` (island/QES extremization), `R4-P` (Page curve), and optionally `R4-M` (microscopic reconstruction). Only the first is even plausibly supported by the current Q2 literature, and Q2 currently assumes that result anyway.

---

## 8. R5: the unimodular-gravity target is incorrectly stated

The step

"derive that a constant vacuum energy density does not source curvature"

is not the correct classical unimodular-gravity result.

In the usual classical formulations, the trace-free field equations and the Bianchi identity recover Einstein's equations with a cosmological constant as an integration constant. A constant vacuum stress shifts the effective cosmological-constant combination. The classical theory is locally equivalent to GR with that effective constant, so it does not make a distinct classical lensing prediction or simply cause vacuum energy not to gravitate.

Replace the R5 steps with:

1. derive the trace-free equations from the selected unimodular action/formulation;
2. use the Bianchi identity and stress-energy conservation to obtain the integration constant;
3. derive how a constant vacuum-energy shift enters the effective cosmological constant;
4. state precisely which quantity is unobservable/absorbed and which curvature observables still depend on the effective constant;
5. identify what, if anything, changes at the quantum level in the selected formulation.

The conditional DESI falsifier must target **a pure constant-Λ dark-energy account**, not unimodular gravity as a formalism. Confirmed evolving dark energy would challenge the Λ-only account; unimodular gravity can be extended with additional dynamical sectors. The statistical significance is dataset-dependent, and 2026 reanalyses report materially different values depending on the supernova compilation. A 5-sigma threshold alone is not a complete model-selection or systematics criterion.

ALG can legitimately record Λ as an input to the de Sitter background and its horizon scale. The Type II_1 algebra does not determine the observed value of Λ.

---

## 9. PR4 and the unit of adjudication: one pilot currently has multiple answers

PR4 requires exactly one terminal state per branch and per pilot. But:

- R3 scores weights and outcomes separately (and records need a third result);
- R4 scores entropy and Page curve separately;
- R1 contains an additional GIE identifiability exercise;
- some pilots can both refute one core postulate and leave another subclaim blocked.

Split composite pilots into atomic claims/subpilots, each with its own frozen target and one terminal result. Suggested minimum split:

- R1-S: semiclassical/source consistency;
- R1-G: GIE model identifiability;
- R2-C: clock/relational evolution per branch;
- R3-W, R3-O, R3-R;
- R4-E, R4-I, R4-P, R4-M;
- R5-L: selected cosmological-constant/adjustment proposal.

Alternatively, define a vector-valued terminal record per pilot, but then PR4 must say one terminal state per **atomic target**, not one per pilot.

---

## 10. Rung definitions and cohesion audit are not yet operational

### Rungs overlap

`EXPLAINS` currently includes both an AUX-based derivation and a qualitative account that does not fix a numerical value. `ACCOMMODATES` covers selecting a free element. `POSTULATES` covers adding an assumption. These can overlap.

Specify precedence and proof requirements. For example:

- DERIVES: explicit derivation from frozen core plus imported mathematical theorems whose hypotheses are checked; no additional physical AUX.
- EXPLAINS: explicit derivation using a predeclared, counted AUX; no target-specific fit parameter is chosen after seeing the target.
- ACCOMMODATES: reproduces the target only after selecting a free parameter/function/sector not fixed by the core.
- POSTULATES: target claim is directly included as a branch axiom or added assumption rather than derived.
- REFUTES: an explicit contradiction/counterexample to a named frozen core postulate under the declared domain.
- BLOCKED: the target cannot be adjudicated because a named mathematical/physical prerequisite or derivation is missing. Distinguish `BLOCKED-MATH`, `BLOCKED-MODEL`, `BLOCKED-DATA`, and `BLOCKED-COMPUTATION` in metadata.

A standard framework must not be countable as one giant AUX. AUX granularity, source, scope, and whether it is physical or mathematical must be fixed before a pilot begins.

### INPUT_ALG and other background premises are not counted

`INPUT_ALG` lists QFT algebra existence, convergence of the `G_N` expansion, a bounded-below clock, and an observer. These are not cost-free merely because they are grouped under a `define`. They need to be classified as:

- mathematical hypotheses of an imported theorem;
- physical branch assumptions;
- background/model-fixture inputs;
- conjectural convergence assumptions;
- experimental inputs.

Count or report them separately. Apply the same accounting to BHC's unspecified matter content, action, field equation, and source ansatz. Equal AUX counts do not make branches comparable if one branch hides a free tensor functional or a major framework in a single item.

### Footprint and assumption graph

`FOOTPRINT(pilot) := the minimal set ...` is too strong unless minimality is demonstrated by removal tests. Rename to `observed_dependency_footprint` until each ingredient is removed in a controlled variant and the derivation is shown to fail.

The pairwise co-occurrence graph loses higher-order dependencies. Preserve each derivation's full assumption set as a hyperedge (or store a derivation-to-assumption bipartite graph), then derive pairwise views as summaries. Distinguish core postulates, AUX, shared framework theorems, fixture assumptions, and empirical inputs. An assumption used by multiple pilots is not automatically an AUX if it is already a core postulate.

### Cohesion threshold

`STRONG_N >= 2` and `AUX_TOTAL <= 3` are draft-chosen heuristics, not an established measure of theoretical coherence. Retain them only as a screening statistic; do not let `COHESIVE_SO_FAR` imply mathematical consistency, explanatory unity, predictive success, or truth. Add explicit contradiction checks and cost/complexity reporting for free functions and unconstrained sectors, not just an AUX count.

---

## 11. GIE subpilot: retain the model-identifiability question, but isolate it

The draft is right not to treat the Aziz–Howl calculation as an uncontested premise. The literature remains actively disputed. Trillo–Navascués argue that the DP model can generate GIE with a model-dependent separation threshold and time behavior; Feng–Vedral–Marletto (2026) argue that the DP mechanism violates the locality assumption of the gravitational witness and attribute the entanglement to hidden quantum-like monitoring degrees of freedom. Angeli–Carlesso also analyze entangling Markovian hybrid models and identify nonlocal features.

This is not one generic "classical gravity" alternative. Each model must be specified by equations, mediator degrees of freedom, locality assumptions, and operational channels.

Required GIE subpilot additions:

- define the exact information-theoretic locality condition (e.g. local operations plus classical communication) being tested;
- distinguish spacetime locality, mediator locality, and LOCC/local-channel assumptions;
- specify whether matter-field propagation or hidden monitoring degrees of freedom mediate correlations;
- compute the full output state and chosen witness, not infer the measured witness directly from a perturbative phase parameter;
- include noise, electromagnetic interactions, binding forces, environmental couplings, and material response as nuisance/alternative channels;
- keep the composite-matter dependence proposal labelled **novelty unverified** until prior-art review and a realistic derivation are completed.

BHC should not yet be called a distinct calculable row: P4 is not sufficiently specified to predict the empty-branch contribution or the response surface. Freeze the BHC source rule first, or mark that row BLOCKED.

---

## 12. Forecasts: current rung predictions are mostly too optimistic

These are pre-review recommendations for the forecast fields, not terminal judgments. No pilot has been executed.

| Target | Current forecast problem | Safer pre-execution forecast |
|---|---|---|
| BHC R1 | P4 source family undefined; conservation not established | **BLOCKED** until source ansatz and conservation proof obligation are fixed |
| ALG R1 | Crossed-product type/entropy does not provide a state-dependent semiclassical-validity criterion | **BLOCKED** until a fluctuation/response criterion and controlled approximation are specified |
| BHC R2 | P5 is assumed but clock/foliation invariance is not operationally defined | **BLOCKED / test pending**; define record statistics and clock class first |
| ALG R2 | Bounded-below observer energy does not itself prove Page–Wootters evolution | **BLOCKED** until the clock POVM, physical inner product, conditional map, and error bound are derived |
| BHC R3-W | Equivariance already appears in P2 | Rework P2; then **conditional EXPLAINS** only if the equilibrium/typicality route is derived with declared measure and measurement model |
| BHC R3-O/R | Actual configuration postulates definite configuration, not automatically stable records | **POSTULATES** definiteness; record formation **BLOCKED** pending apparatus model |
| ALG R3-W | MGBW reconstructs a functional from a measure; does not derive empirical Born frequencies | **POSTULATES** probability identification, or **BLOCKED** if the pilot requires it to be derived |
| ALG R3-O | No ontology or outcome rule supplied | **BLOCKED / genuine silence** |
| ALG R4-E | Q2 already assumes the generalized-entropy relation | **POSTULATES** unless Q2 is rewritten as a theorem target to derive from a specified construction |
| ALG R4-I/R4-P | Q2 does not derive replica wormholes, QES extremization, or a Page curve | **BLOCKED** pending independent derivation |
| BHC R4 | JT benchmark not embedded in a specified BHC model | **BLOCKED** until a branch-specific benchmark model is built |
| R5 | The vacuum-energy non-gravitation target is false as stated | Rewrite target; then **POSTULATES** the value of effective Lambda in both branches |

---

## 13. Source-intake corrections

The current source registry is a useful start, but the following distinctions must be preserved.

- Witten's `Type II_infinity` result is a modular crossed product in a specific large-N setting; it is not the CLPW Type II_1 static-patch construction.
- CLPW's Type II_1 and generalized-entropy relation are specific to their de Sitter static-patch construction.
- JSS's compact-subregion Type II_1 claim includes the conjectural geometric-modular-flow state assumption; record that hypothesis in every pilot borrowing the result.
- Bunce–Wright concerns bounded finitely additive measures on projections and unique bounded linear extension in the absence of a Type I_2 direct summand. Do not cite it as a derivation of the Born rule's empirical applicability.
- Unimodular gravity's classical equivalence to GR with an integration constant is well established in the standard formulation. Any claimed quantum inequivalence must specify the path-integral measure, boundary conditions, and quantization prescription.
- The GIE literature remains disputed. Add the 2026 PRD paper by Feng, Vedral, and Marletto and the 2025 Angeli–Carlesso paper to the DP/hybrid debate register. Do not summarize the dispute as settled in either direction without comparing assumptions.
- DESI evolving-dark-energy significance remains dataset-dependent; include the 2026 DES-SN5Y/Dovekie reanalysis in the source register and phrase the result as a conditional cosmological-model challenge, not a falsifier of unimodular gravity generally.

---

## 14. Minimum repair plan for v0.2.1

**P0 — must fix before any agent starts**

1. Separate core postulates from imported theorems, bridge conjectures, and pilot targets. Remove circular coverage claims.
2. Freeze a concrete BHC source ansatz and conservation obligation; otherwise mark BHC R1/GIE as BLOCKED.
3. Split ALG-W, ALG-CLPW, and ALG-JSS constructions and state their distinct algebra types and hypotheses.
4. Correct G2 and the I_MS/I_SD/I_MD contracts; define reference cases instead of fake roundtrips.
5. Split composite pilots into atomic claims so PR4 is satisfiable.
6. Correct R5's unimodular-gravity target.
7. Correct Q3's MGBW interpretation and D3's pure-state/minimal-projection language.

**P1 — must fix before interpreting pilot results**

8. Rewrite R2 separately for BHC and ALG while retaining a shared operational comparison.
9. Split R3 into weights, outcomes, and records.
10. Specify the JT benchmark and embedding for each branch before R4.
11. Redefine AUX accounting, `FOOTPRINT`, the assumption hypergraph, and the meaning of `COHESIVE_SO_FAR`.
12. Complete source intake for all [RECALL] and [UNVERIFIED] claims.

**P2 — implementation and presentation**

13. Convert the surface form to the authoritative seven-construct Phys-Pseudo syntax before handing it to agents.
14. Add exact citations, versions, assumptions, equations, and source snapshots to each imported theorem record.
15. Add a per-pilot MRC sheet: regime/scale, dimensions, covariance/gauge constraints, conservation, limiting cases, mathematical domain, null model, strongest obstruction, and empirical status.

---

## 15. Final judgment

**Keep the two-branch experiment. Do not freeze this v0.2 version.**

The most valuable part of the draft is the commitment to compare different mathematical/ontological cores using shared observable targets while recording their different footprints and silences. The new interface contracts could become a useful foundation for the mathematical-bridge programme.

But the current document overstates what its cores already entail. ALG's Type II/entropy and probability results must be reconstructed from specific hypotheses rather than imported as a generic package. BHC's source tensor must be fixed before conservation can be adjudicated. The clock and measurement pilots must distinguish their mathematical targets. The rung and AUX rules must not reward assumptions that were already inserted into the core.

The appropriate v0.2.1 status is **research-ready only after P0 repairs**. Until then, the document is a valuable research hypothesis and planning map, not an executable pre-registration.

## Selected verification sources

- Witten, *Gravity and the Crossed Product* (2022): https://arxiv.org/abs/2112.12828
- Chandrasekaran, Longo, Penington & Witten, *An algebra of observables for de Sitter space* (2023): https://link.springer.com/article/10.1007/JHEP02%282023%29082
- Jensen, Sorce & Speranza, *Generalized entropy for general subregions in quantum gravity* (2023): https://doi.org/10.1007/JHEP12(2023)020
- Bunce & Wright, *The Mackey–Gleason Problem* (1992): https://arxiv.org/abs/math/9204228
- Fewster, *Locally covariant quantum field theory and the problem of formulating the same physics in all space-times* (2015): https://doi.org/10.1098/rsta.2014.0238
- Struyve, *Semi-classical approximations based on Bohmian mechanics* (2015): https://arxiv.org/abs/1507.04771
- Struyve, *Towards a novel approach to semi-classical gravity* (2019): https://arxiv.org/abs/1902.02188
- Padilla & Saltas, *A note on classical and quantum unimodular gravity* (2015): https://link.springer.com/article/10.1140/epjc/s10052-015-3767-0
- Trillo & Navascués, *Diósi–Penrose model of classical gravity predicts gravitationally induced entanglement* (2025): https://doi.org/10.1103/PhysRevD.111.L121101
- Angeli & Carlesso, *Entanglement in Markovian hybrid classical-quantum theories of gravity* (2025): https://doi.org/10.1103/PhysRevD.112.024047
- Feng, Vedral & Marletto, *Collapse-based models for gravity do not violate the entanglement-based witness of nonclassicality* (2026): https://doi.org/10.1103/83rl-nygv
- DESI DR2 cosmological constraints: https://arxiv.org/abs/2503.14738 (plus the 2026 DES-SN5Y/Dovekie reanalysis, https://doi.org/10.1140/epjc/s10052-026-16010-6)
