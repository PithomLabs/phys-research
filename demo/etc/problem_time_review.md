# Adversarial review of problem_time.md against the Math Reality Check

This reads as the Wikipedia article on the problem of time, so it is a tertiary survey, not a theory. I reviewed it in two ways: whether its claims are accurate and weighted fairly, and how each proposal fares against the maps (entropy, DoF, symmetry, algebra, spectra, emergence, Born-rule wall). I read the same map sections as last time. I haven't checked the references themselves, and the physics below is from memory, so verify before citing.

**Verdict:** The article is a reasonable map of the territory but overclaims in several places. It treats mainstream relational approaches and fringe proposals as equals, and it blurs two different problems. The frozen-formalism problem is about whether time *exists*. The arrow-of-time problem is about why it has a *direction*. Almost none of the listed proposals touch the second.

## Accuracy problems

1. **The lead mischaracterizes quantum mechanics.** "QM regards time as universal and absolute" is true only of non-relativistic QM. Relativistic QFT is covariant on a fixed background. The conflict appears only when the background is dynamical, as in canonical quantum gravity.
2. **"Closed universe with no external time" overstates the case.** The frozen formalism is sharp for compact spatial slices. Asymptotically flat and AdS spacetimes have a boundary Hamiltonian, which supplies a time. Nobody has established that our universe is closed.
3. **The Wheeler–DeWitt (WdW) treatment omits the momentum (diffeomorphism) constraints.** It also omits the inner-product and Hilbert-space problem. "Constant wave function" is loose, since the state is annihilated by the constraints.
4. **Isham's taxonomy is missing.** The article covers only the frozen formalism. Absent are the problem of observables (Dirac observables are hard to construct), the multiple-choice problem (different clock choices give inequivalent quantum theories), and the global problem (a clock can fail globally).
5. **The 2013 INRIM experiment is overclaimed.** The article says it "confirmed" that time is emergent. The paper's own title calls it an illustration. It is a two-photon polarization model with no gravity, and it cannot discriminate PW from other readings. The follow-up on multitime correlations is more substantive because it addresses Kuchař's objection, and the article doesn't say so.
6. **"Provides a solution" is too strong for consistent discretizations.** The article itself concedes the discrete parameter is not physically accessible. Calling the Montevideo interpretation a solution to the measurement problem and the black hole information paradox presents contested claims as settled.
7. **The standard route is absent.** The Born–Oppenheimer/WKB semiclassical derivation of time from WdW (Zeh, Kiefer) is the most mainstream emergence story. Its link to decoherence and the arrow, via a small-scale-factor initial condition, is also missing. So is any discussion of the Past Hypothesis.

## Proposal by proposal

| Proposal | Map lens | Result |
|---|---|---|
| **Frozen formalism** | Spectra: a spectrum is a photograph of the engine. Calculus: the engine (the flow) is separate. | The zero eigenvalue of the constraint is the photograph. The problem is how to recover the flow. The framing is sound. |
| **Page–Wootters** | DoF: the tensor factorization is chosen. Born wall: conditional probabilities assume Born. | Clean for an ideal clock. For realistic clocks, Pauli's theorem and the clock-choice problem remain. It gives no arrow. |
| **Consistent discretizations and Montevideo** | Entropy map: decoherence does not give single outcomes. | Still needs a clock choice. Fundamental decoherence yields improper mixtures, not outcomes. Contested, with no decisive test that I know of. |
| **Reduced phase space** | Emergence map: a perturbative approximation is language. | Dittrich's scheme gives Dirac observables only as an approximate series. Technically serious, but incomplete. |
| **Elitzur–Dolev** | Maps are silent: no structure and no dividing line. | The experiments cited appear consistent with standard QM, so they don't discriminate. Interpretive, not a physical proposal. |
| **Smolin thick present / Cortês–Smolin** | Entropy and symmetry maps: the arrow is a boundary condition. | As I read the paper, the point is that fundamentally irreversible limit-cycle dynamics can look time-symmetric at the emergent level, which reverses the usual story. Toy models only. |
| **Weyl time (Wang)** | Symmetry: Noether needs an exact continuous symmetry. | Scale invariance is broken by the Higgs mass and by anomalies. A conserved current gives a monotone "harmonic time" only under extra assumptions. It is effectively another clock choice, so the multiple-choice problem applies. |
| **Thermal time (Connes–Rovelli)** | Algebra: active when the product structure does physical work. | The strongest entry. Tomita–Takesaki gives a canonical flow, unique up to inner automorphisms. But the state is an input, and the flow is generally not geometric except in special cases such as Bisognano–Wichmann. It gives a flow, not an arrow. |

The Weyl-time and Elitzur–Dolev entries get equal billing with relational QM and thermal time, though they differ greatly in maturity and testability.

## What this does to dual_kernel.md

The article confirms the main review point. PW, discretizations, reduced phase space and Weyl time are all time-symmetric, so none supplies an arrow. The dual-kernel essay conflates the two problems in the same way the article's lead does.

The thermal time hypothesis also gives the essay a concrete route:

- The accessible subsystem is a von Neumann algebra A. The inaccessible part is its commutant A′.
- The state's modular Hamiltonian, −log ρ in finite dimensions, generates the flow.
- The modular conjugation J maps A to A′ (JAJ = A′). This is the closest thing available to a genuine "dual" in the duality map's sense.
- Read this way, "K₀ = the commutant" is a defined object, where "kernel" currently is not.

Two limits remain. The modular flow runs both ways, so the arrow still comes from the state's preparation. And the state is chosen, which is the same observer-relativity I flagged earlier.

I can draft a reformulation of the essay that uses modular theory and Zeh–Kiefer decoherence for the arrow, if you'd like.


I don't resolve it. Nobody has, and I'd distrust anyone who says they have. But I think the problem is really three problems, and they are in very different shape.

## 1. Does time exist when the constraint says H|ψ⟩ = 0?

This is close to resolved in principle. The resolution is to stop looking for time outside the system and put a clock inside it. Time is then relational: the physical state is the correlation between clock and system. Page–Wootters is the toy version, and the modern version is the quantum-reference-frame or "perspective-neutral" programme (Höhn et al.). Its main gain is on the multiple-choice problem. Different clock choices give descriptions related by explicit transformations, instead of inequivalent theories.

The most striking recent progress is in the algebraic treatment (Witten, Chandrasekaran–Longo–Penington–Witten). Including an observer with a Hamiltonian turns the type III algebra into type II. This makes entropy well defined even in de Sitter. By the algebra map's own test, this is a case where the multiplication table does physical work.

Where the problem was always weaker than advertised, it dissolves:
- In asymptotically flat and AdS spacetimes the Hamiltonian is a boundary term, and boundary time is real time.
- The frozen formalism bites only for compact, closed slicings.

## 2. Why does classical time emerge?

Here the mainstream answer is semiclassical. The WKB expansion of WdW gives an effective Schrödinger equation for matter, and decoherence makes the gravitational sector classical. The emergence map's caution applies: if this is only a perturbative expansion, it is a language rather than a cause. The honest status is that it works as an approximation, with no clean control parameter in the Planck regime.

## 3. Why does time have a direction?

This is a separate problem, and the formalism above is time-symmetric. The entropy map puts it plainly. Entropy defines the arrow but doesn't explain why the initial state was low. The arrow needs a boundary condition (the Past Hypothesis), and no mechanism has been found that selects it. Both the relational and thermal-time routes leave this untouched.

## Alternatives I'd weigh

| Option | What it buys | Cost against the maps |
|---|---|---|
| **Relational time plus quantum reference frames** | Clock choice becomes a transformation, not an ambiguity | Which clock and which factorization are still inputs (DoF map). Clocks can fail globally. |
| **Fundamental foliation** (Hořava–Lifshitz, shape dynamics, Bohmian with a preferred foliation) | Time is real and the formalism is unfrozen | Lorentz invariance must be emergent, and it is tested to extreme precision. Which foliation is an input. |
| **Boundary time** (AdS/CFT, ADM) | Exact where it applies | Silent for cosmology, which is where the problem matters most. |
| **Thermal time** | Canonical flow from the state | The state is an input, and it gives a flow, not an arrow. |

My bet is that layer 1 is solved up to one residual input, namely which factorization or clock. The DoF map flags exactly that as the deepest choice in the theory, one that can be described but not derived. So the problem of time reduces to the problem of the observer and the factorization, which is progress but not a derivation.

A preferred foliation is the cleanest way to remove the problem entirely. Its price is giving up fundamental Lorentz invariance, which is why it must reproduce emergent invariance very precisely. I'd rank it as the live alternative to the relational view, not a rival to dismiss.

## What stays open under every option

- The Hamiltonian constraint itself is input. The calculus map's "engine" is never derived here.
- The Born rule is assumed everywhere, including in conditional probabilities across the clock. That is the wall every map shares.
- The low-entropy start is unexplained.
- Closed cosmologies and global clock failure remain hard.
- No experiment I know of discriminates between these options. Lorentz-violation bounds constrain the foliation route and nothing else does.

So the problem splits into one mostly tamed piece, one approximate piece, and one that is a boundary-condition problem and not a problem of time at all. Would you like me to go further on the foliation route or on the observer-algebra route?

