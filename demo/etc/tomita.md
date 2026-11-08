**The idea in plain English**

Think of a quantum system as a set of questions you can ask about it, which is its algebra of observables. A state is what you currently know. Tomita–Takesaki theory says that if the state is generic enough to probe every question (cyclic) and none of the questions is invisible to it (separating), then the state automatically produces a time evolution for that algebra. You don't have to supply a Hamiltonian.

The recipe has three steps. Take an operator A, act on the state to get AΩ, and then reverse the question to get A†Ω. That reversing map isn't well-behaved on its own. Splitting it into two pieces gives:
- a **mirror flip** (J) that swaps the algebra with its complement, the "other side";
- a **stretching operator** (Δ) whose powers Δ^{it} form a one-parameter flow, the **modular flow**.

The theorem says this flow always maps the algebra to itself. It is a symmetry of the algebra built from the state alone. The state is also exactly thermal with respect to that flow, so every state "looks like equilibrium" for its own modular time.

A finite-dimensional case makes it concrete. If the state is a density matrix ρ, the modular flow is A → ρ^{it} A ρ^{−it}. The generator is −log ρ, the "entanglement Hamiltonian."

**What goes with it (from recall)**
- Connes showed the flow's non-trivial part doesn't depend on which state you chose. That structure classifies algebras into types I, II and III.
- Type I algebras have pure states. Type II has a trace but no pure states. Type III has no trace, and local QFT regions are type III.
- If the state is a trace, the modular flow is trivial.
- **Bisognano–Wichmann:** for the vacuum and a Rindler wedge, the modular flow is a Lorentz boost, which is the Unruh effect.
- **Takesaki:** adding the modular flow back into a type III algebra (a crossed product) gives a type II algebra.

**Why it matters in this thread**

1. **ALG's engine (Q1 and Q2).** The crossed-product step in the Witten and CLPW constructions is the Takesaki move. Dressing a local algebra with the observer's clock turns type III into type II. That gives a trace, and a trace is what makes entropy well defined. Without this theory ALG has no entropy formula.
2. **The arrow-of-time check.** The de Sitter maximum-entropy state is tracial, so its modular flow is trivial. That's why one of DeepSeek's discriminators fails.
3. **The common-language question.** Modular flow is the cleanest example of the matter↔dynamics interface. A state yields a flow with no Hamiltonian supplied. Bisognano–Wichmann is the matter↔spacetime version, where a state on wedge algebras yields geometry. Both are working examples of the Rosetta idea, but only in special cases.
4. **R2 (time).** The "thermal time" proposal says physical time is the modular flow of the state. It is a proposal, not a theorem. It is state-dependent, and in general the flow isn't geometric.
5. **R4 (islands and reconstruction).** Relative entropy is well defined even for type III, and it is written with modular operators. It shows up in the entanglement-wedge reconstruction results (JLMS, QNEC). The missing microscopic reconstruction is partly a question about what modular theory can and can't say in gravity.
6. **The BHC–ALG fork (D3).** The theory works with faithful states on algebras with no pure states, and a Bohmian beable has nowhere obvious to attach.

**What it does not do.** It doesn't select outcomes, fix the physical clock, or work without a faithful state. In ALG it only applies in the G_N expansion.

**Plan changes.** I'd add two round-trip cases to the interface contracts: the finite-dimensional ρ → −log ρ → flow example for matter↔dynamics, and Bisognano–Wichmann for matter↔spacetime. I'd also add "is the state faithful?" to the sanity-check rule. Should I update the file?
