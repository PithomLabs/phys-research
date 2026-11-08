Agent A should be treated as the Senior Go Engineer / Phys Research Infrastructure Architect. Unlike Agent B, its responsibility is to establish how our existing infrastructure supports the research workflow, from Phys-Pseudo exploration through executable Phys derivation.

I would give it the following assignment.

# AGENT A — BHC PREFLIGHT 001

## Phys Research Infrastructure Reconnaissance and TDD Implementation Contract

**Role:** Senior Go Engineer / Software Architect\
**Research programme:** BHC — Bohmian Matter + Classical Gravity\
**Priority:** Highest\
**Phase:** Infrastructure reconnaissance\
**Primary objective:** Establish the exact implementation path from Phys-Pseudo hypothesis formulation to an executable, independently verifiable Phys derivation.

---

# 1. Your role

You are the coding agent responsible for understanding and evaluating the existing Phys Research infrastructure.

You have access to the relevant repositories and libraries, including:

- Phys infrastructure libraries.
- Phys-Pseudo.
- PhysMath.
- PhysLib.
- PhysArtifact.
- Phys research tooling and CLI.
- Existing GR theorem formalization, which is the canonical reference implementation.

Your responsibility is not to design a new research framework.

Your responsibility is to establish how the existing infrastructure should be used to implement our next mathematical research experiment.

You must inspect the actual repositories and source code. Do not infer implementation details from README descriptions or previous conversations when the source can be inspected directly.

## Research objective

We want to resume BHC research using a strict test-driven theoretical physics methodology.

The intended pipeline is:

**Research question → Phys-Pseudo exploration → Math Reality Check → theorem contract → Phys derivation → executable verification → independent adversarial review → promotion.**

Phys-Pseudo is an exploratory mathematical representation tool. It is not the final deliverable.

The ultimate deliverable is an executable, inspectable and independently reviewable mathematical theorem artifact, following the established GR reference implementation.

---

# 2. Operational and repository boundaries

Before doing anything:

1. Locate the actual workspace and relevant repositories.
2. Read the authoritative workspace-level AGENTS.md.
3. Read all applicable package-level AGENTS.md files.
4. Identify the authoritative PSEUDO.md.
5. Identify the canonical GR theorem example.
6. Identify the existing Go modules, workspace configuration and dependency relationships.

Respect repository boundaries.

In particular:

- Do not modify canonical research artifacts.
- Do not modify PhysMath, PhysLib, PhysArtifact or Phys-Pseudo.
- Do not modify the frozen legacy `phys` system or its authoritative instructions.
- Do not change the existing GR reference implementation.
- Do not modify the BHC candidate core.
- Do not create or promote BHC-T001.
- Do not modify shared Go modules, manifests, protocols or documentation.

Your only writable location is your dedicated experiment directory:

`research/experiments/bhc-preflight-001/agents/agent-a-infra-recon/`

Use a separate Git worktree and branch.

If the actual folder structure differs, identify the discrepancy and follow the authoritative repository instructions.

You may inspect source files and execute existing read-only verification commands.

If a necessary experiment requires modifying code, use a disposable copy inside your assigned directory and clearly identify it as a temporary experiment.

Do not assume worktree separation alone provides filesystem-level isolation.

---

# 3. Workstream A — Phys Research architecture reconnaissance

Inspect the actual Go modules and establish their responsibilities.

At minimum, investigate:

### A. Phys-Pseudo

Determine:

- Where its authoritative specification resides.
- Its seven permitted constructs.
- The actual grammar and parsing behavior.
- Validation rules.
- How unresolved obligations are represented.
- Whether hypotheses and competing derivations can be represented without inventing additional syntax.
- How its output is consumed by downstream tooling.
- What is implemented versus merely described in documentation.

Do not modify its syntax.

### B. Phys infrastructure

Determine which existing capabilities support:

- Research records.
- Hypothesis management.
- Derivation artifacts.
- References and provenance.
- Validation.
- CLI operations.
- Agent workspaces and isolation.
- Artifact distribution and verification.

Identify the exact Go packages, modules, interfaces and relevant source files.

### C. PhysMath, PhysLib and PhysArtifact

Determine:

- Their respective responsibilities.
- Existing mathematical representations.
- How mathematical objects are represented and referenced.
- How physical models and theory-specific assumptions are represented.
- How derivations connect to source evidence and physical claims.
- Whether existing types can support the proposed BHC scalar-field investigation.

Do not invent new abstractions merely because a desired capability has not yet been located.

### D. Phys research workflow

Determine how the existing research record lifecycle works.

Identify how an exploratory Phys-Pseudo hypothesis should become a formal research artifact without prematurely promoting conjectures into accepted mathematical results.

---

# 4. Workstream B — GR reference implementation

Treat the existing GR theorem repository as the canonical execution reference.

Inspect the complete relevant implementation, including:

- AGENTS.md.
- manifest.json.
- PSEUDO.md references, where applicable.
- protocol/derivation_test.go.
- Mathematical definitions and supporting packages.
- Test fixtures and verification behavior.
- Required commands and build environment.

Answer these questions:

1. What makes a theorem artifact executable?
2. What exactly is verified by derivation_test.go?
3. What assumptions are encoded in the manifest?
4. How are equations, definitions and intermediate derivation steps represented?
5. How are mathematical identities tested?
6. How are limiting cases and mathematical invariants checked?
7. How does the implementation distinguish a proven result from a conditional result?
8. How is an independent reviewer expected to reproduce the result?
9. Which parts of the GR implementation are reusable?
10. Which parts are specific to General Relativity?

Do not interpret a passing Go test as proof of an arbitrary mathematical theorem.

Clearly distinguish:

- Mathematical proof obligations.
- Executable verification.
- Numerical or symbolic checks.
- Research artifact validation.

---

# 5. Workstream C — Test-driven theoretical physics workflow

Our intended methodology adapts test-driven software engineering to theoretical physics.

Establish exactly how the existing infrastructure can support the following process.

### RED — Theorem contract

Before attempting a derivation, specify:

- The mathematical problem.
- Inputs and definitions.
- Physical assumptions.
- Applicable regime and scale.
- Target proposition.
- Expected mathematical properties.
- Failure conditions.
- Required Math Reality Check obligations.

Identify how these requirements can be recorded and verified using existing infrastructure.

### EXPLORE — Phys-Pseudo

Determine how agents can independently formulate:

- Alternative hypotheses.
- Definitions.
- Assumptions.
- Mathematical propositions.
- Derivation attempts.
- Counterexamples.
- Unresolved obligations.

The seven permitted Phys-Pseudo constructs must remain authoritative.

### MRC — Mathematical acceptance

Determine how the existing infrastructure supports checks for:

- Dimensions.
- Tensor and index structure.
- Covariance.
- Conservation.
- Mathematical invariants.
- Constraint closure.
- Limiting cases.
- Applicable physical regime.
- Strongest known obstructions.

Identify which checks are already automated and which remain human mathematical obligations.

### GREEN — Phys derivation

Explain the exact procedure for translating an accepted Phys-Pseudo hypothesis into an executable theorem artifact.

Use the existing GR implementation.

Identify the required source structure, equations, definitions, intermediate steps, Go representations, test protocol and manifest information.

Do not invent a new derivation language.

### REFACTOR — Independent verification

Determine how the infrastructure can support:

- Reproducibility.
- Independent mathematical reconstruction.
- Adversarial review.
- Regression tests.
- Versioned theorem revisions.
- Preservation of failed and superseded derivations.

---

# 6. Workstream D — BHC-T001 feasibility mapping

Do not implement BHC-T001.

Instead, establish whether the existing infrastructure can represent and verify its proposed mathematical requirements.

The candidate theorem concerns:

**Existence and consistency of a Bohmian configuration-dependent gravitational source.**

The target equation is:

8\pi G S\_{\mu\nu}[Q,\Psi,g].\
]

The main mathematical requirement is:

[\
\nabla^\mu S\_{\mu\nu}=0\
]

under appropriately specified coupled matter and gravitational equations.

The initial matter model is a real scalar field coupled to classical spacetime.

Assess whether the existing infrastructure can accommodate:

1. A scalar-field matter action.
2. Euler–Lagrange equations.
3. Canonical field variables.
4. A Bohmian wave functional.
5. Quantum Hamilton–Jacobi equations.
6. Guidance dynamics.
7. Quantum-potential contributions.
8. A configuration-dependent stress-energy tensor.
9. Covariant conservation.
10. Appropriate classical and semiclassical limits.

For each requirement, report:

- Existing support.
- Exact supporting package and source file.
- Reusable implementation.
- Missing capability.
- Whether the missing capability is fundamental or can be expressed through the existing derivation protocol.

Do not add infrastructure simply to accommodate a hypothetical formulation.

---

# 7. TDD acceptance contract

Recommend a minimal, concrete implementation contract for the first BHC theorem.

It must establish:

| Requirement                | Acceptance                                                                      |
| -------------------------- | ------------------------------------------------------------------------------- |
| Mathematical specification | All relevant objects and assumptions are explicitly defined.                    |
| Derivation representation  | Intermediate mathematical steps are inspectable.                                |
| Verification               | Mathematical properties have executable checks where appropriate.               |
| MRC                        | All applicable mathematical and physical obligations are recorded.              |
| Failure behavior           | An invalid or incomplete derivation cannot be promoted as a successful theorem. |
| Reproducibility            | Another investigator can execute the prescribed verification.                   |
| Provenance                 | Sources, assumptions, equations and code revisions are traceable.               |
| Independence               | Review can occur without modifying the original derivation.                     |

Determine how to implement this using existing infrastructure.

Identify the smallest changes that are genuinely necessary.

If the current infrastructure cannot support an essential requirement, explain exactly why.

Do not propose architectural changes without a demonstrated blocker.

---

# 8. Required deliverables

Create only these files inside your assigned directory.

### 01-repository-map.md

Exact repository paths, Go modules, package responsibilities, dependencies and relevant instructions.

### 02-phys-pseudo-contract.md

Authoritative seven-construct syntax, validation rules, parser/consumer behavior and the path from exploratory hypothesis to formalization.

### 03-gr-reference-analysis.md

Detailed analysis of the canonical GR implementation, including manifest, derivation protocol, tests and reusable patterns.

### 04-tdd-research-workflow\.md

Concrete RED–EXPLORE–MRC–GREEN–REFACTOR workflow using actual infrastructure capabilities.

### 05-bhc-t001-feasibility.md

Requirement-by-requirement feasibility assessment for BHC-T001.

### 06-gaps-and-recommendations.md

Only demonstrated implementation gaps, their impact, recommended minimal remedies and priorities.

### 07-implementation-contract.md

A concise, actionable contract for the next coding phase.

Include exact file paths, commands, test procedures and acceptance conditions where established.

### 08-handoff.md

A compact decision report:

- What can be used immediately.
- What already works.
- What is missing.
- What must not be changed.
- Whether the existing infrastructure is ready for BHC-T001.
- The smallest next implementation task.

---

# 9. Reporting discipline

This is an engineering reconnaissance task.

Do not produce an architectural manifesto or lengthy narrative.

Prefer:

- Exact paths.
- Actual Go symbols.
- Interfaces and types.
- Existing commands.
- Test names.
- Source references.
- Reproducible observations.
- Concrete blockers.

Every claim about existing infrastructure must be supported by an actual repository path or executable verification.

Distinguish:

- IMPLEMENTED.
- DOCUMENTED ONLY.
- PARTIALLY IMPLEMENTED.
- MISSING.
- NOT VERIFIED.

Do not claim a capability exists merely because the architecture describes it.

Do not modify infrastructure to make the assessment look complete.

---

# 10. Completion conditions

Your assignment is complete when we can answer five questions:

1. How exactly does Phys-Pseudo feed the existing theorem formalization workflow?
2. What is the canonical path from a hypothesis to an executable Phys derivation?
3. What must BHC-T001 provide mathematically to pass the existing verification protocol?
4. Which existing infrastructure components can be reused without modification?
5. What is the smallest genuine blocker preventing us from proceeding?

Do not implement the theorem.

Do not launch the other BHC pilots.

Do not modify shared infrastructure.

**Your objective is to give us an accurate engineering contract so the next agent can implement the first theorem without inventing another research framework.**

A concise, accurate assessment that identifies a blocker is preferable to an elaborate redesign.

## My expectation from Agent A

The most valuable output is `07-implementation-contract.md`.

I want Agent A to tell us, with actual repository evidence, precisely how this transition works:

Phys-Pseudo hypothesis → MRC acceptance → theorem specification → Phys derivation → `derivation_test.go` → independent verification.

If the infrastructure already supports this, we should use it immediately. If something is missing, we should implement only that missing piece—not redesign Phys Research.

Agent A and Agent B should work independently. Agent A owns infrastructure reconnaissance; Agent B owns the physics investigation. Their findings meet only when we formulate BHC-T001.
