Next: use QSG-T001 to validate the platform, then move to Bohmian equivariance. Do not spend another cycle reviewing QSG-T001.

The goal now is to establish that PhysMath and PhysTrace can handle genuinely different mathematical structures—not just one successful example of stress-energy conservation.

## The next milestone: a second physics theorem

Benchmark 1 — QSG-T001

Working result

Covariant conservation of renormalized stress-energy for a free scalar field on fixed curved spacetime.

Mathematical status: derived human result under the repaired prescription. Machine verification remains limited.

Benchmark 2 — Bohmian equivariance

Next

Equivariance of the quantum-equilibrium distribution under Bohmian configuration-space dynamics.

Mathematical structure: probability continuity equations, configuration-space flow, and preservation of distributions.

Why this is the right next experiment:

QSG-T001 involves curved spacetime, bitensors, Hadamard subtraction, coincidence limits and covariant differentiation.

Bohmian equivariance involves a different structure: wavefunctions, probability currents, continuity equations and dynamical flows.

Both should be representable using shared mathematical foundations, while their physical interpretations remain in their respective theory-specific packages.

## What Agent A should do

Give Agent A this mission:

# Mission: PhysMath / PhysLib Second Vertical Slice — Bohmian Equivariance

## Objective

Formalize and implement the Bohmian equivariance theorem as the second major physics benchmark, following the existing Phys / PhysLib / PhysMath / PhysTrace research-artifact conventions.

The strategic objective is to test whether the shared mathematical foundation can represent substantially different physical theories without introducing theory-specific semantics into PhysMath.

## Scope

Use nonrelativistic Bohmian mechanics for a finite-particle system governed by the Schrödinger equation.

Formalize the relationship between:

- The wavefunction and its probability density.
- The Schrödinger probability current.
- The continuity equation.
- The Bohmian guidance equation.
- The evolution of the configuration-space probability distribution.
- The equivariance result.

Explicitly state the assumptions needed for the result, including regularity, boundary conditions, and existence of the relevant configuration-space flow.

Distinguish local-in-time results from global results. Do not claim global existence merely from the continuity equation.

## Architectural constraints

1. Inspect the existing repositories and AGENTS.md instructions before changing anything.
2. Reuse existing PhysMath and PhysLib abstractions wherever mathematically appropriate.
3. Keep PhysMath theory-neutral.
4. Keep Bohmian physical commitments inside phys-lib/bohm.
5. Do not introduce abstractions merely to accommodate this theorem.
6. Do not modify frozen legacy repositories or unrelated packages.
7. Preserve the established manifest.json, protocol/derivation_test.go and research-artifact conventions.

## Mathematical requirements

Derive the theorem from its actual premises.

Do not import the theorem conclusion as a golden value.

Make the relationship between the continuity equation and the guidance flow explicit. Show which mathematical assumptions establish preservation of the probability distribution.

Identify any steps that are mathematically derived, source-reported, conditionally established, or still unsupported.

Apply the Math Reality Check. State the physical regime and effective-theory scope explicitly.

## Platform evaluation

After implementation, answer these questions:

1. Which mathematical structures from QSG-T001 can be reused?
2. Which structures are genuinely new?
3. Can the new structures be represented without polluting PhysMath with Bohmian interpretation?
4. What abstractions should remain theory-specific?
5. What genuine deficiencies in the shared mathematical foundation have now been demonstrated?

Do not refactor the architecture preemptively. Propose changes only when an actual mathematical requirement exposes a deficiency.

## Deliverables

- An executable research artifact for Bohmian equivariance.
- Explicit assumptions and a transparent derivation.
- Tests that verify only what the implementation actually supports.
- A concise comparison of the QSG-T001 and Bohmian representations.
- A prioritized list of demonstrated gaps in the shared mathematical foundation.

## Stop condition

Do not expand into quantum field theory, relativistic Bohmian mechanics, measurement theory, or a general interpretation comparison.

The deliverable is one complete, bounded theorem formalization that tests the common mathematical foundation.

Success means a second independently scoped physics theorem can be represented using the shared platform without compromising mathematical fidelity.

## What happens after that?

We compare the two artifacts and produce a Mathematical Coverage Map.

It should show which mathematical capabilities are genuinely common across theories, which belong to specific physical frameworks, and where the current architecture fails to represent something important.

That becomes the evidence-based foundation for the next development of PhysMath—not an abstract attempt to design a universal mathematics framework in advance.

The strategic sequence is therefore:

1. QSG-T001 establishes the first working research artifact.
2. Bohmian equivariance tests cross-theory generality.
3. The coverage map identifies the true shared mathematical core.
4. Only then do we expand the platform.

This is the shortest credible path toward the larger goal: a platform that can represent, investigate, and eventually compare competing physical theories using a common mathematical language.&#x20;
