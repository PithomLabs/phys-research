My recommendation: yes, conduct deep research first, but make it narrowly scoped and preceded by repository reconnaissance. Do not start coding or launch all five pilots yet.

We need to establish two things before writing our first theorem test:

1. What does our existing Phys Research architecture actually require?
2. What mathematics and physical assumptions are needed for the first BHC theorem candidate?

I would run two independent agents in parallel.

## 1. Agent A — Repository reconnaissance

Read-only assignment

Ask the agent to inspect the actual repository, not invent a new workflow.

It should inspect:

- Workspace and package-level `AGENTS.md`.
- `PSEUDO.md` and the seven permitted constructs.
- The existing GR theorem implementation.
- Its `manifest.json`.
- Its `protocol/derivation_test.go`.
- Existing Math Reality Check and verification requirements.
- Current BHC research artifacts.

Deliverable: A short implementation contract explaining exactly how a new BHC theorem must be represented, tested, reviewed and promoted.

No modifications to existing files.

## 2. Agent B — Focused physics deep research

Primary-source investigation

The research question should be:

> Can Bohmian quantum matter interacting with classical dynamical spacetime produce a mathematically consistent, configuration-dependent gravitational source?

I would restrict this first investigation to a simple scalar-field model, while allowing the agent to recommend a better starting model if the literature demonstrates a problem with that choice.

It must investigate:

1. Bohmian field ontology and guidance dynamics.
2. Quantum potential and quantum Hamilton–Jacobi formulation.
3. Stress-energy tensor derived from the matter action.
4. Configuration-dependent gravitational sourcing.
5. Covariant conservation and the Bianchi identity.
6. Foliation dependence and relativistic consistency.
7. Existing attempts at Bohmian semiclassical gravity.
8. Known obstructions, counterexamples and alternative formulations.

The most important question is:

\\[ \boxed{ \nabla^\mu S\_{\mu\nu}[Q,\Psi,g]=0\\;? } \\]

Can this be derived from the coupled equations of motion, or does it require additional assumptions?

The agent must not assume that the actual-configuration source works. It must investigate whether it works.

Deliverable: A compact primary-source ledger, mathematical reconstruction of the relevant existing results, and a precise list of unresolved mathematical obligations.

The attached BHC draft itself identifies source verification and the conservation condition as unresolved weaknesses.&#x20;

BHC_CORE_AND_COVERAGE.md



## 3. Then we write the theorem contract

Only after receiving both reports would I authorize a third agent to prepare:

`BHC-T001 — Configuration-Dependent Gravitational Source`

The contract must specify:

| Item               | Requirement                                                                      |
| ------------------ | -------------------------------------------------------------------------------- |
| Given              | Exact matter action, spacetime geometry, quantum state and Bohmian configuration |
| Definitions        | All dynamical variables and mathematical objects                                 |
| Assumptions        | Explicit, numbered, justified                                                    |
| Claim              | Existence and consistency of the proposed source                                 |
| Derivation         | Actual first-principles calculation                                              |
| Failure conditions | Conservation failure, incompatible dynamics, counterexample                      |
| MRC                | Dimensions, covariance, limiting cases, constraints and physical regime          |
| Tests              | Executable checks against the mathematical claims                                |

This becomes the RED phase of our TDD process: we define what must pass before attempting to derive the result.

## 4. Execution sequence

1. Now: Two independent investigations

   Repository reconnaissance and focused physics research.
2. Next: Freeze BHC-T001 contract

   Agree on the mathematical model, proposition and acceptance tests.
3. Then: Independent derivation

   Separate agents investigate the source tensor, conservation and possible counterexamples.
4. Finally: Executable theorem

   Implement the derivation using the established GR reference structure, run tests and commission adversarial review.

## 5. The instruction I would give both agents

# BHC Preflight 001 — Research Instructions

We are preparing the first theorem-driven investigation of BHC.

This is a narrowly scoped preflight. Do not attempt to solve all five research problems.

## Shared rules

- Work exclusively in your assigned directory and separate Git worktree.
- Do not inspect other agents' outputs.
- Do not modify canonical research artifacts, PhysMath, PhysLib or PhysArtifact.
- Do not produce lengthy narrative reports.
- Preserve source claims, independent derivations and hypotheses as separate categories.
- Use primary sources wherever possible.
- Apply the Math Reality Check.
- Report negative findings and mathematical obstructions.
- Do not claim that BHC is mathematically consistent.

## Agent A: Repository reconnaissance

Determine the authoritative implementation contract for a new theorem using the existing Phys GR reference implementation.

Identify:

- Required repository instructions.
- Phys-Pseudo syntax and constraints.
- Required theorem artifact structure.
- manifest.json requirements.
- protocol/derivation_test.go requirements.
- Verification and adversarial-review requirements.
- Any conflicts or missing implementation instructions.

Deliver:

1. implementation-contract.md
2. reference-file-map.md
3. unresolved-protocol-questions.md

Read-only investigation. No implementation changes.

## Agent B: Focused physics research

Investigate whether Bohmian quantum matter can consistently source classical dynamical spacetime through its actual configuration.

Start with an explicitly defined scalar-field model unless the literature provides a compelling reason to choose otherwise.

Investigate:

- Matter action and stress-energy tensor.
- Bohmian wave functional and guidance equation.
- Quantum potential and continuity equation.
- Candidate configuration-dependent gravitational source.
- Covariant conservation.
- Gravitational constraints and foliation.
- Existing proposals and their limitations.

Do not assume the source tensor exists.

Determine whether its divergence can be derived to vanish from the coupled matter and gravitational equations.

Deliver:

1. primary-source-ledger.md
2. mathematical-reconstruction.md
3. candidate-model.md
4. unresolved-obligations.md
5. strongest-objections.md

Every consequential mathematical statement must identify its source or provide its own derivation.

## Completion criterion

Stop when the implementation requirements and physics foundations needed to formulate BHC-T001 are sufficiently clear.

Do not expand into the other four pilots.

Do not generate a proposed solution merely to complete the assignment.

The objective is to establish whether a precise, testable theorem can be formulated.

My decision: Start with these two agents, independently. Do not launch a broad deep-research exercise across all five problems.

The first milestone is not a new theory document. It is a theorem contract that is mathematically specific enough to fail.

That is where our TDD-driven Phys Research approach should resume.
