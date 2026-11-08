Muse's next step should be to author the first research specification in Phys-Pseudo—not to write code yet.

The Constitution establishes what must govern our research, but it deliberately does not prescribe a rigid implementation sequence. Our immediate operational choice should be to turn one well-researched physics question into a specification that can survive independent scrutiny.

## Recommended first assignment: R3-W

I recommend starting with R3-W: the mathematical foundations of quantum probability weights in Bohmian mechanics.

This is the strongest first candidate because we already have a Bohmian equivariance benchmark. We can build on that work while investigating the boundary between what equivariance establishes and what additional assumptions are needed to connect it to observed measurement frequencies.

R1-S remains important, but it is not ready for specification until the Bohmian gravitational source tensor and its governing assumptions are explicitly defined.

## Muse's immediate responsibilities

1\. Receive the research dossier.

The Deep Research Agent must establish the literature baseline: primary sources, established results, exact hypotheses, competing explanations, unresolved questions and potential novelty. Muse must not invent missing scientific premises to fill gaps.

2\. Produce Phys-Pseudo v0.1.

Translate the research question and its mathematical foundations into implementation-independent reasoning.

It should make explicit:

- The physical system and applicable regime.
- The mathematical structures involved and why they are needed.
- Definitions, assumptions and imported results.
- The actual derivation, including its logical dependencies.
- What follows mathematically and what requires additional physical assumptions.
- What remains unresolved and what could falsify the proposed conclusions.

Use our established Phys-Pseudo vocabulary: `given`, `define`, `assume`, `claim`, `derive`, `if/then/else` and `?`.

Do not introduce Go code, package structures, APIs, ASTs or implementation-specific concepts into this specification.

3\. Submit it to OpenCode.

OpenCode must independently reconstruct and challenge the reasoning, applying all Three Reality Checks. It must look particularly for hidden assumptions, circular derivations, invalid mathematical bridges, unsupported conclusions and overlooked prior results.

4\. Revise and submit for approval.

Muse addresses the review findings. Critical unresolved defects mean the specification is blocked or its claims must be narrowed. The specification is not frozen until you approve it.

5\. Implement only after the specification is frozen.

Muse then produces the implementation contract and implements the approved mathematics using PhysMath and PhysLib. Implementation tests must distinguish what is actually executed from what is merely represented or assumed.

## The instruction I would give Muse now

# Muse — Next Assignment: Phys-Pseudo R3-W v0.1

Your immediate assignment is to prepare an implementation-independent Phys-Pseudo specification for R3-W:

**Research question:**

For a specified nonrelativistic Bohmian system, which mathematical assumptions establish equivariance of the configuration-space probability distribution, and what additional assumptions are needed to connect that result to observed measurement frequencies?

Distinguish carefully between probability weights, outcome actualization and stable measurement records. Do not assume that establishing one establishes the others.

### Governing requirements

1. Follow Phys Research Workflow Protocol v1.0.
2. Apply all Three Reality Checks: Mathematical Landscape, Mathematical Intersection and Mathematical Interoperability.
3. Use the Deep Research Agent's verified research dossier as the scientific starting point.
4. Preserve the distinction between established results, imported theorems, derived claims, physical assumptions and conjectures.
5. Make the logical dependencies and unresolved obligations explicit.
6. Use the established seven Phys-Pseudo constructs.
7. Do not introduce implementation details or modify any existing physics codebase.

### Deliverables

- Phys-Pseudo R3-W v0.1.
- An explicit assumption and imported-result register.
- Derivation dependencies and unresolved proof obligations.
- Relevant limiting cases, failure conditions and possible counterexamples.
- A list of claims that must be independently verified.

### Restrictions

Do not implement the specification yet. Do not claim novelty. Do not treat existing benchmark tests as proof of the complete physical result.

Submit the specification for independent adversarial review by OpenCode. Address the findings before requesting approval to freeze the specification.

Your responsibility is to make the physics explicit, inspectable and challengeable—not to make an incomplete argument look complete through implementation.

The immediate handoff is therefore: Deep Research Agent → Muse → OpenCode → your approval.

The important boundary is that Muse does not get to decide that its own specification is scientifically correct. It authors the formal reasoning; independent review challenges it; you authorize the transition to implementation.

That gives us a practical first application of the Constitution without turning the Constitution itself into a rigid production pipeline.&#x20;
