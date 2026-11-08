# AGENT A — QSG PREFLIGHT 001
## Phys Research Infrastructure Reconnaissance and TDD Implementation Contract

**Role:** Senior Go Engineer / Phys Research Infrastructure Architect

**Research programme:** QSG — Quantum Fields on Classical Spacetime

**Phase:** Infrastructure reconnaissance

**Priority:** High

**Objective:** Establish the exact implementation path from Phys-Pseudo exploration to an executable, independently verifiable mathematical derivation for our first QSG theorem.

---

# 1. Strategic change

We are temporarily moving away from Bohmian matter coupled to classical gravity as our primary research direction.

Our new starting point is conventional quantum field theory on classical curved spacetime.

The initial research progression is:

1. Quantum field theory on a fixed classical spacetime.
2. Semiclassical Einstein equations with renormalized quantum stress-energy.
3. Self-consistent gravitational backreaction.
4. State-dependent conditions for the validity of the semiclassical approximation.

The initial matter model is a real scalar quantum field.

A possible gravitational equation is:

\[
G_{\mu\nu}+\Lambda g_{\mu\nu}
=
8\pi G\langle\hat T_{\mu\nu}\rangle_{\mathrm{ren}}.
\]

This is a research starting point, not an accepted theorem or a declaration that semiclassical gravity is fully consistent.

The immediate objective is to determine how our existing Phys Research infrastructure can support a precise mathematical investigation of this system.

## Important boundaries

- BHC remains an archived candidate research direction.
- Preserve the BHC preflight investigation and its reported conservation obstruction.
- Do not modify or delete existing BHC research artifacts.
- Do not assume that BHC findings establish anything about conventional QFT-based semiclassical gravity.
- Do not introduce Bohmian mechanics into QSG-T001 unless a later research decision explicitly requires it.
- Do not introduce Einstein–Cartan torsion, stochastic gravity or quantum spacetime into the initial QSG theorem.
- Do not implement the theorem during this assignment.

We are simplifying the physics problem, not redesigning Phys Research.

---

# 2. Your role

You are the coding agent responsible for understanding the existing Phys Research infrastructure and determining how it supports a test-driven theoretical physics workflow.

You have access to:

- Phys infrastructure libraries.
- Phys-Pseudo.
- PhysMath.
- PhysLib.
- PhysArtifact.
- Phys research tooling and CLI.
- Existing GR theorem formalization, which remains the canonical reference implementation.

Inspect actual source code and repository instructions.

Do not infer implementation capabilities from documentation alone when the implementation can be inspected.

Your working directory is:

`research/experiments/qsg-preflight-001/agents/agent-a-infra-recon/`

Use a dedicated Git worktree and branch.

Your work must not modify shared infrastructure or canonical research repositories.

---

# 3. Workstream A — Repository and architecture reconnaissance

Locate the actual workspace and inspect:

1. Workspace-level AGENTS.md.
2. All relevant package-level AGENTS.md files.
3. Authoritative PSEUDO.md.
4. Existing GR theorem implementation.
5. Relevant Go modules and workspace configuration.
6. Phys research CLI and artifact lifecycle.

Determine the actual responsibility of:

### Phys-Pseudo

Establish:

- The authoritative seven-construct syntax.
- Its grammar and validation behavior.
- How definitions, assumptions, claims, derivations and unresolved obligations are represented.
- How competing hypotheses can be expressed.
- How an accepted exploration is handed to formal derivation.
- What is actually implemented versus merely documented.

Do not invent additional syntax.

### Phys infrastructure

Determine how existing components support:

- Mathematical research records.
- Assumptions and hypotheses.
- Source references and provenance.
- Derivation artifacts.
- Validation.
- Executable verification.
- Independent reviews.
- Artifact lifecycle and promotion.

### PhysMath, PhysLib and PhysArtifact

Determine their existing responsibilities and identify any relevant mathematical representations.

Do not propose new abstractions unless an actual requirement cannot be represented by the existing architecture.

---

# 4. Workstream B — Canonical GR implementation

The existing GR theorem repository remains the canonical execution reference.

Inspect the complete relevant implementation:

- AGENTS.md.
- manifest.json.
- Mathematical definitions.
- protocol/derivation_test.go.
- Supporting Go packages.
- Verification commands.
- Test fixtures.
- Provenance and review conventions.

Answer:

1. What makes the theorem executable?
2. What mathematical claims are actually verified?
3. How are assumptions represented?
4. How are intermediate derivation steps recorded?
5. How are limiting cases and invariants tested?
6. How are conditional results represented?
7. What is the distinction between a proof obligation and an executable test?
8. Which implementation components can be reused for QSG-T001?
9. Which GR-specific components must not be generalized?

Do not confuse a successful Go test with a formal mathematical proof.

Explicitly distinguish:

- Analytical derivation.
- Symbolic computation.
- Numerical verification.
- Executable regression testing.
- Mathematical proof.

---

# 5. Workstream C — QSG mathematical representation feasibility

Assess whether the existing infrastructure can represent the mathematical objects likely to arise in QSG-T001.

The initial research involves:

- A smooth globally hyperbolic Lorentzian spacetime.
- A real scalar quantum field.
- A specified matter action.
- Quantum states and field operators.
- A renormalized stress-energy tensor.
- Covariant conservation.
- Semiclassical gravitational backreaction.

Potential mathematical objects include:

1. Classical metric and curvature tensors.
2. Scalar field and canonical variables.
3. Quantum field operators.
4. Quantum states and two-point functions.
5. Hadamard or otherwise appropriately specified state conditions.
6. Renormalized stress-energy.
7. Renormalization prescriptions and their ambiguities.
8. Semiclassical Einstein equations.
9. Coupled matter-gravity equations.
10. Conservation and constraint identities.

Determine which objects can be represented through existing PhysMath and PhysLib structures, which can be represented within a Phys derivation, and which require additional mathematical conventions.

In particular, investigate whether our current derivation infrastructure can adequately represent:

- Operator-valued quantities.
- Functional derivatives.
- Renormalized composite operators.
- State-dependent expectation values.
- Curved-spacetime differential operators.
- Nonlocal kernels, where relevant.

Do not claim support merely because an expression can be written as a string or stored in Phys-Pseudo.

Identify the difference between representability and verifiability.

---

# 6. Workstream D — Test-driven research workflow

Establish the exact workflow we should use for QSG-T001.

## RED — Theorem contract

Before derivation, define:

- Exact mathematical problem.
- Physical and mathematical inputs.
- Assumptions and conventions.
- Applicable spacetime and energy regime.
- Target proposition.
- Expected identities.
- Failure conditions.
- MRC obligations.

Determine how existing infrastructure represents and validates these requirements.

Do not decide the final mathematical proposition yourself. It will be established from the physics investigation.

## EXPLORE — Phys-Pseudo

Determine how agents can formulate and compare alternative hypotheses using the seven authorized constructs.

Ensure that:

- Proposals are not confused with established results.
- Definitions are explicit.
- Assumptions are declared.
- Derivation obligations are testable.
- Unresolved mathematical questions remain visible.

## MRC — Mathematical Reality Check

Map existing infrastructure support for:

- Dimensional consistency.
- Tensor and index structure.
- Covariance.
- Renormalization assumptions.
- Conservation.
- Mathematical invariants.
- Limiting cases.
- Constraint consistency.
- Physical regime and effective scale.
- Known mathematical obstructions.

Classify every check:

- AUTOMATED.
- PARTIALLY AUTOMATED.
- HUMAN MATHEMATICAL OBLIGATION.
- NOT SUPPORTED.

## GREEN — Phys derivation

Establish how a validated QSG hypothesis becomes an executable mathematical theorem artifact.

Follow the GR reference implementation, including:

- AGENTS.md.
- manifest.json.
- protocol/derivation_test.go.

Identify how analytical derivations, supporting equations and verification tests are connected.

## REFACTOR — Independent verification

Determine how our infrastructure supports:

- Independent reproduction.
- Adversarial mathematical review.
- Versioned theorem changes.
- Failed and superseded derivations.
- Provenance and evidence.
- Regression checks following changes to mathematical assumptions.

Do not introduce a new workflow where existing infrastructure already provides the capability.

---

# 7. QSG-T001 feasibility assessment

Our working target is:

**QSG-T001 — Conservation and consistency of renormalized quantum stress-energy as a source for classical gravity.**

The initial investigation will use a specified real scalar quantum field on classical curved spacetime.

The possible starting relation is:

\[
G_{\mu\nu}+\Lambda g_{\mu\nu}
=
8\pi G\langle\hat T_{\mu\nu}\rangle_{\mathrm{ren}}.
\]

The precise theorem target remains provisional.

Assess the infrastructure's ability to express and verify requirements involving:

### A. Matter foundations

- Scalar-field action.
- Classical field equation.
- Stress-energy operator.
- Metric variation.

### B. Quantum formulation

- Quantum field equation.
- Quantum state.
- Two-point functions.
- State regularity requirements.

### C. Renormalization

- Definition of the renormalized stress-energy tensor.
- Choice of prescription.
- Local curvature ambiguities.
- Dependence on renormalization conventions.

### D. Conservation

- Covariant conservation of the renormalized expectation value.
- Assumptions required for the relevant identity.
- Possible anomalies or additional geometric terms.

### E. Semiclassical backreaction

- Coupled metric and matter equations.
- Consistency of gravitational constraints.
- Conditions for existence and stability of the resulting solutions.

For each requirement, record:

- Existing implementation support.
- Exact repository path and relevant symbols.
- Reusable components.
- Missing capabilities.
- Whether the missing capability is fundamental or merely representational.

Do not implement any of these mathematical structures during reconnaissance.

---

# 8. TDD implementation contract

Prepare a minimal implementation contract for QSG-T001.

The contract must distinguish three things:

**Mathematical proof:** The actual derivation and argument establishing the proposition.

**Executable verification:** Code that checks defined identities, special cases, numerical examples and regression conditions.

**Research provenance:** The sources, assumptions, representations and derivations supporting the result.

Every theorem must make clear which claims belong to each category.

Identify what must be present before the first Phys derivation can be accepted.

The goal is to reuse the existing GR execution protocol wherever possible.

Do not introduce symbolic algebra or external dependencies merely because they are convenient. Assess the existing standard-library-only constraints and determine how they affect the proposed mathematical representation.

---

# 9. Required deliverables

Produce only these files inside your assigned folder.

### 01-repository-map.md

Actual repositories, modules, package responsibilities, dependencies and authoritative instructions.

### 02-phys-pseudo-contract.md

Authoritative syntax, parsing, validation and transition from exploratory Phys-Pseudo to formal derivation.

### 03-gr-reference-analysis.md

The canonical GR theorem structure, manifest, protocol, tests and reusable implementation components.

### 04-qsg-mathematical-feasibility.md

Assessment of whether the existing infrastructure can represent the mathematical objects required for QSG-T001.

### 05-tdd-research-workflow.md

Concrete RED–EXPLORE–MRC–GREEN–REFACTOR workflow using existing capabilities.

### 06-infrastructure-gaps.md

Only demonstrated gaps, with exact evidence and minimal remedies.

### 07-implementation-contract.md

The precise engineering requirements for implementing the first QSG theorem.

### 08-handoff.md

A short decision report specifying:

- What is ready now.
- What is already implemented.
- What is documented but not implemented.
- What cannot currently be represented or verified.
- What must remain human mathematical work.
- What the smallest next engineering task should be.
- Whether the infrastructure is ready for QSG-T001 formalization.

---

# 10. Reporting and quality requirements

This is an engineering reconnaissance task.

Avoid lengthy architectural prose.

Prefer:

- Exact repository paths.
- Actual Go symbols.
- Existing interfaces and types.
- Test names.
- Executable commands.
- Reproducible observations.
- Concrete blockers.

Use these implementation statuses:

- IMPLEMENTED.
- DOCUMENTED ONLY.
- PARTIALLY IMPLEMENTED.
- MISSING.
- NOT VERIFIED.

Do not claim a feature exists merely because it appears in a design document.

Do not implement new infrastructure to make the assessment look complete.

Do not use the QSG research change as an excuse to redesign the entire Phys Research architecture.

---

# 11. Completion criteria

The assignment is complete when we can answer:

1. Can Phys-Pseudo represent the hypotheses and mathematical obligations of QSG-T001?
2. What is the exact path from Phys-Pseudo to the executable GR-style theorem artifact?
3. Can the current representation handle renormalized quantum-field quantities adequately?
4. Which mathematical claims must remain human proof obligations?
5. What is the smallest demonstrated infrastructure gap?
6. Can we begin QSG-T001 without modifying the existing architecture?

**The objective is to prepare the infrastructure for rigorous theoretical physics research, not to produce another research framework.**

Inspect, verify, report. Do not implement the theorem.
