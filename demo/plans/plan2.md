Yes. I remember the approach, and I think this is exactly where we need to become much stricter.

We previously established an important architectural principle:

The canonical unit of research is an executable mathematical derivation, not a research document.

The General Relativity repository was our reference implementation. A theorem or hypothesis formalization must have:

- `AGENTS.md`
- `manifest.json`
- `protocol/derivation_test.go`

It must be executable and independently inspectable. Otherwise, it is still a research proposal or mathematical prose, not a completed formalization.

I would now combine that architecture with test-driven development.

# 1. The fundamental division of responsibility

We need three distinct layers.

Exploration

Phys-Pseudo

Allows agents to formulate competing hypotheses, definitions, assumptions, claims and derivations from first principles.

Exactly seven constructs:

`given`, `define`, `assume`, `claim`, `derive`, `if/then/else`, `?`

Scientific validation

Math Reality Check

Determines whether the hypothesis is mathematically meaningful, internally consistent, properly scoped and worth formalizing.

Executable research

Phys Derivation

Converts the selected hypothesis into a precise mathematical proposition, explicit derivation, executable verification protocol and independently reviewable theorem artifact.

Phys-Pseudo is deliberately flexible. We must not constrain an agent's ability to discover alternative mathematical structures.

But flexibility during exploration must not become flexibility in the acceptance criteria.

# 2. Apply test-driven development to theoretical physics

I would adapt the familiar RED–GREEN–REFACTOR cycle.

| Software engineering        | Theoretical physics equivalent                                                                  |
| --------------------------- | ----------------------------------------------------------------------------------------------- |
| Write a failing test        | Specify the theorem and its acceptance conditions before deriving it.                           |
| Implement the functionality | Derive the mathematical proposition from the declared assumptions.                              |
| Run tests                   | Execute mathematical consistency and regression checks.                                         |
| Refactor                    | Simplify the derivation and eliminate unnecessary assumptions without changing the proposition. |
| Code review                 | Independent adversarial mathematical derivation.                                                |

The workflow becomes:

1. Define the problem

   One precise unresolved mathematical question. No extensive literature narrative.
2. RED: Define the theorem contract

   Specify inputs, assumptions, expected proposition, failure conditions and mathematical checks. Initially, the verification is incomplete.
3. Explore using Phys-Pseudo

   Agents develop competing hypotheses independently. They may challenge the original assumptions or propose alternative mathematical formulations.
4. Run Math Reality Check

   Reject undefined objects, unsupported assumptions, invalid mathematical transformations and missing physical regimes.
5. GREEN: Phys Derivation

   Translate the strongest surviving hypothesis into an explicit mathematical derivation with intermediate steps.
6. Execute verification

   Run derivation tests, mathematical identities, limiting-case checks and applicable consistency conditions.
7. REFACTOR

   Improve the derivation, simplify assumptions and strengthen the result without changing the theorem contract.
8. Independent adversarial review

   Another agent reproduces or challenges the mathematics independently.
9. Promote the MVP theorem

   Publish the executable, reviewed mathematical artifact with explicit limitations.

One important qualification: Go tests do not automatically prove arbitrary mathematical propositions. They verify executable calculations, identities and specified properties. The actual mathematical derivation remains the proof obligation.

# 3. Make Phys-Pseudo flexible, but prohibit endless prose

I would establish these rules:

Rule 1 — One investigation, one mathematical target.

An agent cannot submit a literature review as completion of a mathematical assignment.

Rule 2 — Every `claim` must have a purpose.

It must either be supported by a derivation, explicitly conditional, falsifiable or unresolved.

Rule 3 — Every `derive` must calculate something.

No symbolic placeholders masquerading as mathematical progress.

Rule 4 — Every `?` must identify a specific obligation.

Not vague questions such as "investigate consistency."

Rule 5 — Every assumption has an owner and a reason.

Undeclared assumptions are defects.

Rule 6 — Negative results count.

A failed theorem attempt is useful when the obstruction is explicit and reproducible.

Rule 7 — No document promotion without executable mathematical evidence.

A beautiful Phys-Pseudo file is not an MVP theorem.

# 4. What I would do for BHC immediately

I would temporarily set aside all five validation pilots and concentrate on one mathematical target.

First theorem candidate

# BHC-T001

Existence and consistency of a Bohmian configuration-dependent gravitational source

Given a specified quantum matter action, wave functional, Bohmian guidance dynamics and classical spacetime geometry:

Can we derive an effective gravitational source tensor that is compatible with the gravitational field equations and whose covariant divergence vanishes on the coupled equations of motion?

\\[ G\_{\mu\nu}+\Lambda g\_{\mu\nu} = 8\pi G S\_{\mu\nu}[Q,\Psi] \\]

Required result: derive whether such an \\(S\_{\mu\nu}\\) exists, under which assumptions, and whether those assumptions are physically and mathematically defensible.

This is a proposed theorem target, not an assertion that the source exists or that the theorem is true.

The first test is not to make the equation work.

It is to determine whether the proposed construction is possible at all.

I would assign separate agents to:

- Reconstruct the Bohmian scalar-field dynamics.
- Independently derive candidate stress-energy sourcing.
- Investigate the conservation requirement.
- Search for a counterexample or no-go result.
- Challenge the entire construction.

Each gets its own directory and worktree.

# 5. Definition of an MVP theorem

I would establish the following acceptance contract.

| Requirement                 | Acceptance condition                                         |
| --------------------------- | ------------------------------------------------------------ |
| Mathematical specification  | All objects and assumptions defined                          |
| First-principles derivation | Explicit intermediate steps                                  |
| MRC                         | All applicable obligations addressed                         |
| Executability               | `protocol/derivation_test.go` runs successfully              |
| Mathematical verification   | Relevant identities and consistency checks pass              |
| Reproducibility             | Independent researcher can inspect and reproduce             |
| Adversarial review          | No unresolved blocking mathematical defect                   |
| Scope                       | Limitations and applicable physical regime explicitly stated |

A theorem may be conditional. That is acceptable.

It may also establish a no-go result. That is acceptable.

What is not acceptable is presenting an unproved conjecture as a successful theorem.

# 6. The operating principle I would lock

I would make this the governing principle of the next Phys Research phase:

> Explore freely, specify rigorously, derive explicitly, test continuously, and promote only verified mathematical results.

And one additional rule:

The unit of progress is not the number of documents produced. It is the number of mathematical uncertainties eliminated.

For our immediate work, I recommend we now prepare a single implementation specification for the BHC TDD workflow, anchored to the existing GR example and the actual Phys-Pseudo protocol.

Before that specification is written, the coding agent should inspect the existing repository instructions and GR reference implementation rather than inventing a new derivation format.

That gives us a disciplined path from hypothesis to executable MVP theorem without turning Phys Research into another documentation-production system.
