# Can We Tell Black Holes Apart by Their Radiation?
### A plain-English guide to the GSST-INFO-001 research mandate and the harness that keeps it honest

---

## The question, in one paragraph

When a black hole evaporates, it radiates away its mass as Hawking radiation. If you form two black holes with slightly different initial masses, do they produce radiation that is *distinguishable* — meaning, in principle, you could tell which black hole made which radiation? And if so, is that difference connected to the famous "Page curve," the fingerprint of information coming back out of a black hole? This is the question an autonomous AI research agent is being asked to investigate, within a specific toy model of an evaporating black hole called GSST. The answer is not known in advance. The agent might prove a theorem, might discover a mathematical obstruction, or might find that the question is already answered in the literature. All three are valid outcomes.

---

## Why this is hard: the missing pieces

Imagine you have a recipe for a cake (the black hole model), and you want to know whether changing the amount of sugar (the initial mass) changes the *flavor* of the cake (the quantum state of the radiation) in a way a taster could detect.

The GSST model gives you the recipe. It tells you how the black hole forms, how it evaporates, and what the entropy of the radiation is at every moment. It even gives you the Page curve — the entropy rises, peaks, and falls, exactly as quantum mechanics demands.

But here is the catch. The model gives you the **entropy** of the radiation — a single number that measures uncertainty. It does **not** give you the **quantum state** of the radiation — the full description from which you could, in principle, compute whether two radiation samples are distinguishable.

**Analogy:** knowing the temperature of a room is not the same as knowing the exact positions and velocities of every molecule in it. The temperature is a summary. The state is the full picture. GSST gives you the summary. The research question asks about the full picture. Whether the full picture is even *definable* within GSST is part of what the agent must investigate.

There is a second catch. In quantum field theory — the mathematical language GSST uses for the radiation — the "algebra of observables" (the collection of measurements you can make on the radiation) is of a type called **type III**. This is a technical term, but its consequence is simple: you cannot just write down a density matrix for the radiation and compute relative entropy between two states the naive way. You need a reference state and a piece of machinery called the **modular Hamiltonian** to even define the entropy of a region.

**Analogy:** imagine trying to compare the "information content" of two books, but the books have infinitely many pages and no page numbers, and the only way to define what "page 1" means is to pick a reference book and measure everything relative to it. That is what type III algebras are like. The comparison is possible, but only with the right scaffolding.

---

## The bridge that might work: the modular Hamiltonian

The most promising route the agent might take is through the **modular Hamiltonian**. In plain terms, this is an operator that encodes how a region's quantum state responds to small perturbations. For a region of space in the vacuum, the modular Hamiltonian is known exactly — it is an integral of the stress-energy tensor against a geometric weight.

Here is why this matters for the research question. The GSST ledger already establishes (in R1) the **stress tensor** of the radiation — how the quantum matter pushes on the geometry. If the agent can show that a small change in the initial mass produces a small change in the stress tensor, and that this change is related to the change in entanglement entropy via the **first law of entanglement** (a linearized relation, δS = δ⟨K⟩), then there is a genuine mathematical bridge from "mass" to "entropy" to "distinguishability."

**Analogy:** think of the stress tensor as the *pressure* a gas exerts on its container, and the entanglement entropy as the *uncertainty* about where the gas molecules are. The modular Hamiltonian is the conversion factor between the two. If you know how the pressure changes, and you know the conversion factor, you know how the uncertainty changes. The first law of entanglement is the statement that this conversion is exact for small perturbations. Whether it applies in GSST, and at what order in the semiclassical expansion, is exactly what the agent must determine.

This is the most likely *positive* outcome: a conditional result that says, roughly, "under the coherent-state identification used in the GSST ledger, and at leading order in the semiclassical expansion, a change in initial mass δM produces a change in radiation entropy δS that is bounded by the modular response — but this is not the same as full state distinguishability, and here is exactly what additional structure would be needed to make that stronger claim."

That would be a genuine scientific contribution. It would be a bridge from the ledger's R1 objects to an information-theoretic statement, with the non-identifications made explicit.

---

## The other likely outcome: a precise obstruction

There is a real chance the answer is: **GSST cannot support the full question.** Not because the model is wrong, but because it does not contain the structures required.

To claim that two radiation states are distinguishable, you need:

- A **Hilbert space** for the radiation — the full space of possible quantum states.
- An **encoding isometry** — a map from the black hole's initial state to the radiation state, telling you how information gets out.
- A **common algebra** on which both ω_R(M₁) and ω_R(M₂) are defined — so you can compare them.

GSST supplies none of these. It supplies a semiclassical geometry, a holographic matter sector, and a coherent-state identification of the radiation. That is enough to compute entropy branches. It is not enough to define the full quantum state of the radiation as a function of mass.

**Analogy:** imagine you have a machine that produces a statistical summary of a baseball game — final score, hits, errors. You want to know whether two games with different starting pitchers are "distinguishable" in terms of the full play-by-play. The summary tells you the games were different. It does not give you the play-by-play, and it does not give you a way to reconstruct it. The honest answer is: "The summary depends on the pitcher, but the play-by-play is not available from this machine."

A precise obstruction statement — "GSST does not define ω_R(M) on a common algebra; mass-distinguishability requires an encoding isometry and a complete radiation Hilbert space, neither of which is present; here is the minimum additional structure that would be needed" — is a legitimate and valuable scientific result. It is not a failure. It is the research finding.

The harness is designed to make this outcome legible and honest, rather than treating it as a disappointing near-miss that the agent should keep pushing past.

---

## The harness: how do you stop an AI from making things up?

This is the other half of the project, and it is arguably the more important half. The research is being conducted by an autonomous AI agent. The question is: how do you let an AI do science without it hallucinating results, fabricating citations, or paper-milling its way to a predetermined conclusion?

The answer is a **research harness** — a set of rules and structures that control what enters the accepted scientific record. The design rests on several key ideas.

### The claim ledger: the only memory that counts

Every substantive claim the agent makes is recorded in an append-only ledger, tagged with an evidence class:

- **[S]** — verified against a source, with an exact page and equation locator.
- **[R]** — independently reconstructed, with a reproducible derivation and independent verification.
- **[I]** — imported from another paper, with its assumptions and applicability stated.
- **[A]** — approximation, with the regime of validity and the neglected terms specified.
- **[U]** — unverified.
- **[X]** — contradicted, with the counterexample recorded.

A claim cannot be promoted from [U] to [R] just because the agent is confident. It needs a derivation that an independent reviewer can reproduce, and an independent route that agrees.

**Analogy:** think of the claim ledger as a laboratory notebook with a strict rule: you can add pages, but you cannot erase them. If you correct an entry, you add a new page explaining what was wrong and why. At the end, anyone can audit the entire history of what was claimed, what was checked, and what was corrected.

### Pre-registration: write the hypothesis before you test it

Before any computation, the agent must write down the hypothesis, its assumptions, its predicted consequence, its falsifier, and the cheapest discriminating test. This is what separates research from post-hoc storytelling.

**Analogy:** it is the difference between a detective who writes down their theory before examining the evidence, and one who examines the evidence and then constructs a theory that fits it. The first can be wrong in a checkable way. The second can always be right, which means it can never be trusted.

### Generator/verifier separation: no agent grades its own work

The agent that generates a hypothesis and derives a result is not the agent that verifies it. The verifier sees only the claim, the definitions, the assumptions, and the evidence artifacts. It does not see the generator's reasoning or preferred conclusion. Its job is to independently reconstruct the result and try to break it.

If separate agents are not available, the self-review must be explicitly labeled as such — not presented as independent verification.

**Analogy:** it is the difference between a student grading their own exam and a second teacher grading it blind, without knowing what the first teacher expected. The second teacher might still be wrong, but at least the errors are independent.

### The three reality checks as gates, not prose

The research protocol requires three checks at every stage:

- **RC1 — the toolbox:** what mathematical structures are being used, in what regime, and what do they actually describe?
- **RC2 — the intersection:** which structures must meet, and what happens when you remove one?
- **RC3 — the bridge:** for every mapping between frameworks, what is preserved, what is lost, and what are the failure conditions?

These are not reporting sections. They are gates. A failed gate blocks promotion of dependent claims. Independent work may continue.

**Analogy:** think of them as quality-control stations on an assembly line. A part that fails inspection does not move to the next station. It gets flagged, and the downstream work that depends on it is halted until the issue is resolved.

### Stop conditions: when is the research finished?

The agent decides when to stop. There are six terminal outcomes:

- **Success:** a verified result meeting the deliverable requirements.
- **Proven negative:** a precise mathematical obstruction has been demonstrated.
- **Falsified:** the hypothesis is disproved and no worthwhile replacement remains.
- **Known result:** prior art already establishes the result.
- **Budget exhausted:** an execution limit has been reached.
- **Insufficient information:** a missing input prevents further inference.

A failed initial hypothesis does not end the research. The agent is authorized to pivot to adjacent questions within the same model, documenting the scope reduction.

**Analogy:** it is the difference between a project manager who says "we failed, stop everything" and one who says "the original approach didn't work, but here is a narrower question we can still answer, and here is why it is still worth answering."

---

## The pivot protocol: how to descend gracefully

For GSST-INFO-001 specifically, there is a natural ladder of questions, from strongest to weakest:

1. **Full state distinguishability:** can we compare ω_R(M₁) and ω_R(M₂) as states on a common algebra? (Likely blocked: GSST does not supply the full state.)
2. **Relative entropy under coherent-state identification:** can we define a meaningful relative entropy using the [32] coherent state and the vacuum modular Hamiltonian? (Possible, conditional on the identification.)
3. **Modular-Hamiltonian linearized response:** can we relate δM to δS via the first law of entanglement, at leading order? (Most likely positive result.)
4. **Entropy derivative as a semiclassical proxy:** can we at least establish that ∂S_island/∂M = 2, while explicitly *not* identifying this with distinguishability? (Almost certainly yes, and valuable as a negative clarification.)
5. **Obstruction:** GSST does not support mass-distinguishability at any level; here is the minimum additional structure needed. (Possible, and a legitimate result.)

The agent is free to descend this ladder. Each descent is a scope reduction that must be recorded, not a silent substitution. The harness makes the descent legible.

---

## What a "win" actually looks like

The most likely overall outcome is a **mix**: a conditional positive result on the modular-response side (level 3), plus an obstruction on the full state-distinguishability side (level 1). That is the honest shape of what GSST can support.

The single most valuable positive result would be a conditional theorem: "Under the coherent-state identification and at leading order in the semiclassical expansion, a change in initial mass δM produces a change in radiation entropy δS bounded by the modular response, with the following explicit assumptions and failure conditions." That would be an N1–N2 contribution — a formalization or a new mathematical bridge — not a claim to have solved the information paradox.

The single most valuable negative result would be a precise obstruction: "GSST does not define ω_R(M) on a common algebra; mass-distinguishability requires an encoding isometry and a complete radiation Hilbert space; here is the minimum additional structure needed." That would be an N1 contribution — a clarification of what the model can and cannot do.

Both are legitimate. The harness is designed so that the agent does not have to force a positive result to justify the investigation.

---

## Why this matters beyond the black hole

The black hole is the test case. The real product is a **template for honest autonomous research**.

Physics has a credibility problem that is not about dishonesty but about *traceability*. A headline says "black holes preserve information." A reader has no way to check. Even a specialist often cannot tell, from a summary, which parts of a result are derived, which are assumed, and which are imported from another paper with different assumptions.

The harness is an attempt to build a small, honest machine for answering that question. It takes one well-studied corner of physics, gives an AI agent the authority to investigate it autonomously, and enforces a set of rules that make every claim auditable, every assumption explicit, and every failure legible.

If it works, it is a template: a way to make AI-assisted physics claims checkable, challengeable, and reproducible — one bounded investigation at a time. The black hole is the test case. The real product is **trust with a paper trail**.

**Analogy for the whole project:** think of it as building a flight recorder for scientific reasoning. The recorder does not fly the plane. It does not guarantee the plane will not crash. But it ensures that if the plane crashes, you can reconstruct exactly what happened, which assumptions failed, and what would need to change. That is what the harness is for.