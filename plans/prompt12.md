# Task: Draft the Root README.md for phys-research

## Objective

Draft a clear, technically accurate and concise root-level `README.md` that explains what `phys-research` is, why it exists, how it works, and how to use it.

Position `phys-research` as an **independent, coding-agent-agnostic physics research system**. OpenCode, Codex, Muse and other compatible coding agents are execution environments, not architectural dependencies.

The README should help a new contributor or physics researcher understand the project and successfully begin using it.

## 1. Reconnaissance before writing

Inspect the actual repository before drafting.

Read the relevant:

- `AGENTS.md`
- `docs/AGENT_CONTRACT.md`
- `agents/` shared instructions
- `docs/compatibility-matrix.md`
- `docs/dossier-branch-c.md`
- Existing architecture and methodology documentation
- `go.mod`
- `physctl` CLI usage and help output
- `.gitignore`
- Relevant research artifact examples

Verify all commands, paths, dependencies, Go version requirements and capabilities against the current implementation.

Do not assume previous implementation reports describe the current repository accurately.

## 2. README structure

Use the following organization, adapting it where the actual repository warrants it.

### A. Project introduction

Explain:

- What problem `phys-research` addresses.
- Why existing AI-assisted paper summarization is insufficient for serious physics research.
- How it connects literature, mathematical claims, assumptions, derivations, hypotheses and empirical evidence.
- Why traceability and scientific inspectability matter.

Make the distinction clear:

**The system assists scientific investigation; it does not automatically establish the truth of a physical theory.**

### B. Core principles

Explain the project's guiding principles:

- Mathematics is central to representing physics research.
- Source claims must be distinguished from independently reproduced derivations.
- Assumptions, physical regimes and effective-theory scales must be explicit.
- Relationships between theoretical frameworks require justified mathematical mappings.
- Unresolved questions remain explicit obligations.
- Hypotheses remain provisional until scientifically established.
- Human researchers retain scientific acceptance authority.

### C. Architecture

Explain the actual architecture, including:

- Go-based `physctl` deterministic operations.
- Plain-text research artifacts and Git-based provenance.
- Phys-Pseudo and its seven constructs:
  `given`, `define`, `assume`, `claim`, `derive`, `if/then/else`, `?`.
- Paper discovery, acquisition and extraction capabilities.
- Researcher methodology and platform-neutral agent instructions.
- Validation and cross-reference mechanisms.
- Agent-specific adapters.

Include a simple Mermaid architecture diagram if it accurately represents the implementation.

Do not invent components or imply that planned capabilities already exist.

### D. Coding-agent compatibility

Explain that the research core is independent of the coding-agent vendor.

Describe the shared agent contract and the role of platform-specific adapters.

Document the actual compatibility status of:

- OpenCode
- Codex
- Muse

Use the compatibility matrix as the source of truth. Clearly distinguish verified, partially verified and unverified capabilities.

Explain mandatory core capabilities versus optional integrations such as web search, PDF inspection, subagents and MCP.

### E. Research artifact model

Explain the existing artifact types and their relationships.

Show a small, real example of how a paper, issue, hypothesis, evidence item and derivation relate to each other.

Explain provenance, references, validation and the meaning of unresolved `?` obligations.

Use actual repository examples. Do not fabricate records.

### F. Current research demonstration: Branch C

Introduce the investigation:

"Can classical Riemann–Cartan geometry be consistently coupled to quantum matter without quantizing spacetime?"

Explain the three frameworks being investigated separately:

1. Semiclassical Einstein–Cartan gravity.
2. Stochastic or hybrid classical–quantum Einstein–Cartan dynamics.
3. Quantum Einstein–Cartan gravity.

Describe the existing literature map, comparison work, competing hypotheses and outstanding mathematical obligations.

Be explicit that the existing dossier has not independently reproduced its load-bearing mathematical derivations. Do not present the investigation as a proven theory or completed scientific result.

### G. Getting started

Provide verified installation and setup instructions.

Include:

- Prerequisites and minimum Go version.
- Repository checkout.
- Build instructions.
- `physctl` availability.
- Required environment configuration, if any.
- How to inspect CLI help.
- How to validate the repository.
- How to run tests.

Use commands that have been confirmed against the current implementation.

### H. Working with a coding agent

Explain how to:

- Choose a supported coding-agent environment.
- Load the shared research instructions.
- Assign a dedicated Git worktree and branch.
- Establish a unique agent/run identity.
- Execute the Researcher workflow.
- Create and validate research records.
- Preserve other agents' outputs.
- Submit work for human integration and review.

Explain that worktree separation is an operational safeguard, not OS-level filesystem isolation.

Include the actual baseline distribution procedure, referencing its authoritative documentation rather than duplicating a fragile procedure.

### I. Development and contribution

Explain:

- Repository organization.
- How to add research artifacts.
- How to validate changes.
- How to preserve provenance.
- How to propose revisions without overwriting existing research.
- How competing hypotheses and derivations should be retained.

### J. Current limitations and roadmap

Separate implemented capabilities from planned work.

Mention only limitations supported by the current repository.

Do not imply that autonomous theorem proving, independent mathematical verification, experimental evaluation or unrestricted multi-agent isolation are already implemented.

### K. License and citation

Inspect the repository for an actual license and citation metadata.

Do not invent a license, DOI, formal publication status or citation format. State clearly when these are not yet established.

## 3. Writing requirements

- Write for technically literate researchers and software contributors.
- Explain physics-research concepts without assuming the reader knows the project's internal architecture.
- Be precise rather than promotional.
- Prefer concrete examples over abstract claims.
- Keep sections concise and avoid repeated explanations.
- Use tables only when they improve comprehension.
- Use relative links to existing repository documents.
- Ensure links and commands are valid.
- Do not include private machine paths, user-specific directories or machine-local configuration.
- Do not include secrets, credentials or personal information.

## 4. Accuracy requirements

For every capability described, determine whether it is:

- IMPLEMENTED
- VERIFIED
- PARTIALLY VERIFIED
- PLANNED
- NOT IMPLEMENTED

Do not describe a planned capability as implemented.

Do not claim that a paper's result has been independently derived unless an actual derivation exists.

Do not imply that coding-agent compatibility has been verified merely because an agent can read the shared instructions.

Do not claim complete filesystem isolation.

Do not invent benchmarks, success rates, scientific conclusions or supported integrations.

## 5. Scope and preservation

This is a documentation-only task.

- Create or revise only the root `README.md`.
- Preserve every existing research artifact and its hash.
- Preserve all agent instructions and configurations.
- Do not modify Go source code.
- Do not change research conclusions.
- Do not modify `AGENTS.md`.
- Do not add dependencies, features or infrastructure.
- Do not commit or merge.

If a root `README.md` already exists, inspect it and preserve useful accurate material. Do not silently discard existing content.

## 6. Verification

After drafting:

1. Check every relative link.
2. Verify every command against actual CLI help or execution.
3. Verify every documented path exists.
4. Verify the architecture description against the actual implementation.
5. Verify compatibility claims against the compatibility matrix.
6. Check that implemented and planned capabilities are not conflated.
7. Run `git diff --check`.
8. Confirm that only `README.md` changed.

## 7. Deliverables

Provide:

1. The completed root `README.md`.
2. A concise explanation of its structure.
3. Any facts that could not be verified and were therefore omitted or qualified.
4. Verification results.
5. Confirmation that no research artifacts or implementation files were changed.

**Success criterion:** A new contributor should be able to understand the purpose of `phys-research`, understand its scientific and architectural boundaries, install it, select a coding-agent environment and begin a traceable research workflow without needing this conversation.
