# research-brief.md — GSST-INFO-001

Governed by AGENTS.md. This brief fixes scope; it does not prescribe a conclusion.

## 1. Subject
Mass-parameter sensitivity and quantum distinguishability of Hawking radiation in the RST–GSST holographic-CFT evaporating black hole (GSST, JHEP 05 (2020) 091, arXiv:2004.00598).
Starting artifact: JT-ISLAND-002 Canonical Equation Ledger v1.2 (read it; its tags carry over, including open item OQ-3 and the B1/B2 conditions).

## 2. Central question
Can the model support a rigorous, independently defined relationship between the initial mass parameter M, the radiation state, its distinguishability from nearby radiation states, and the Page transition? If not, say precisely why.

## 3. Starting facts (treat as [A], leading order; re-derive before use)
With tau = t_A - sigma_A (retarded anchor time):
- S_no = (c/12) tau + O(c) constants; S_island = 2M - (c/24) tau + O(c) constants; crossing tau_P = 16M/c at leading order.
- At fixed tau: dS_no/dM = 0, dS_island/dM = 2, at leading order only.
These are entropy-branch statements under the QRT prescription. They are NOT relative entropy, Fisher information, distinguishability, or recoverability. Keep those distinctions.

## 4. Prerequisite checks (do these before any "distinguishability" computation)
P1. Source of the "2M". In the island branch it comes from the area term 2(Omega_I - Omega_crit) (verify in the ledger). The matter state is a coherent state on the vacuum (ledger import [32]); for coherent states the matter entropy equals the vacuum value. Determine what a matter-algebra relative entropy can and cannot see, and whether the area term enters any state-level quantity at all. This may split the problem into "what the matter algebra distinguishes" versus "what the gravitational term carries".
P2. Common algebra across M. Geometry, anchor Omega_A, and the QES all move with M. Establish whether a family omega_R(M) on ONE fixed radiation algebra A_R exists, and whether its mass dependence is explicit, derivable or imported. A negative answer is a first-class result.
P3. Algebra type and regulator. Local continuum subregion algebras are expected to be type III_1; the cutoff delta in (4.6) does not automatically make them type I. Establish what is well defined (Araki relative entropy, density-matrix relative entropy, information metric, modular Hamiltonian).
P4. Reference state. Define it; show both states live on the same algebra.
P5. Missing quantum-gravity input. If an encoding isometry, complementary factorization, or the central dogma is needed, name it. Do not construct it.
P6. Regime mismatch. Recoverability thresholds concern nonperturbatively small errors; GSST controls O(c) quantities at leading order in epsilon. State what the model can and cannot resolve.

## 5. Seed hypotheses (not endorsed; generate 3–7 including a null; replace freely)
H0 (null): the model supplies no state-level distinguishability beyond the entropy-branch switch.
H1: matter-algebra relative entropy between nearby masses is controlled by coherent-state data and is blind to the area term.
H2: mass dependence of radiation entropy after tau_P is carried entirely by the island area term, so it is not a matter-state distinguishability statement.
H3: no well-defined common algebra/regulator exists for the comparison (obstruction).
H4: any smoothness or non-analyticity at the Page time is an artifact of the branch minimum, not of state distinguishability.
Score each on physical significance, well-definedness, compatibility with GSST, independent computability, falsifiability, novelty, expected contribution. Do not pick a hypothesis because it is easy to compute.

## 6. Prior art (starting points; access NOT yet verified; fetch and pin before relying)
- Zhong, "An algebraic description of the Page transition", arXiv:2601.11363, JHEP 04 (2026) 160. Reported scope: explicit relative-entropy-difference probes for type I/II factors; direct type III proof open. Assumes a complete quantum-gravity description (encoding isometry, complementary recovery, central dogma).
- Cited in the search results for that paper, not yet read: Zhong arXiv:2408.15104; Xu–Zhong arXiv:2411.06361; Gomez arXiv:2403.09165, arXiv:2405.08435; van der Heijden–Verlinde arXiv:2408.00071; Engelhardt–Liu arXiv:2311.04281; Witten arXiv:2112.12828; Chandrasekaran–Penington–Witten arXiv:2209.10454.
- Also search: Fisher information / entanglement first law in evaporating or RST/JT models; any prior mass-sensitivity or distinguishability computation in GSST-like models; Geng et al. JHEP 01 (2022) 182 (islands in long-range gravity; interpretation-level risk R6).
A novelty claim must survive comparison with the closest existing result.

## 7. Scope
In: mass-parameter response; radiation state families; relative entropy and justified information metrics; modular/algebraic structure; regulator and continuum issues; entropy–Page comparisons; necessary quantum-gravity inputs; theorem / proposition / conditional / no-go discovery.
Out: reopening model selection; rebuilding the ledger; proving QRT or replica wormholes; microscopic unitarity; 4D Schwarzschild; endpoint microphysics; full information-paradox resolution; unrelated programmes; production PhysMath/PhysLib/PhysArtifact/PhysTrace changes; swapping in HSS/DREH/JT.
Forbidden claims: unitarity proof; recoverability established; entropy derivative = Fisher information; any invented matter Lagrangian or microscopic Hilbert space.

## 8. Expectation (calibration, not a target)
Honest prior: N0–N2 most likely, with a precisely stated obstruction or known-result finding a probable outcome. That is acceptable. Do not inflate.

## 9. Budgets (runtime must enforce; fill before launch)
max_hypotheses=7; max_cycles_per_hypothesis=6; adversarial_passes=1; correction_passes=1; stagnation_cycles=3; wall_clock_cap=<set>; token_cap=<set>; literature_search_cap=<set>.
If the runtime cannot enforce a cap, the final record must say so.

## 10. Deliverables
Mandatory: `research-ledger/claims.jsonl`, `final-decision.md` (one terminal outcome, N-level, next question).
Conditional on content (create only if substantive): hypothesis-register, source-and-prior-art ledger (with stored source text + hashes), derivation ledger, reality-check register, `reproducibility/` (only necessary scripts), research report.
Analytic derivations stay primary; scripts are checks, and numerical results are never theorem proofs.
