# Implementation Task: Minimal Go-Based Physics Research Tracker

## Objective

Implement a minimal, Git-native physics research workbench orchestrated by OpenCode.

The purpose is to track physics papers, outstanding problems, competing hypotheses, theoretical expectations, experimental observations and mathematical derivations, eventually supporting the development of robust physical theories.

Prioritize simplicity, inspectability and scientific integrity. Do not introduce infrastructure or abstractions without a demonstrated need.

## 1. Reconnaissance — mandatory before implementation

Inspect the existing repositories and documentation:

- `phys-pseudo`
- `phys`
- `phys-math`
- `phys-lib`
- `phys-artifact`
- `phys-trace`

Read the applicable `AGENTS.md`, Phys-Pseudo specification, MRC review protocol, derivation guidelines and relevant repository-local instructions.

Establish:

1. The exact seven supported Phys-Pseudo constructs.
2. Existing artifact identity and reference conventions.
3. Existing mathematical operations and derivation-test capabilities.
4. Current Go module and repository conventions.
5. The installed OpenCode version and its supported agent configuration and permission syntax.

Do not assume a capability exists because a research hypothesis needs it.

Do not modify existing specifications or introduce competing artifact identity systems.

If a required capability is missing, document the gap rather than inventing an implementation.

## 2. Architecture

Implement:

- One Go executable: `physctl`.
- Three OpenCode agent roles: Researcher, Adversary and Formalizer.
- Git and plain files as the canonical research store.
- Phys-Pseudo as the scientific reasoning language.
- Existing Phys infrastructure for mathematical formalization.

OpenCode is the orchestrator and LLM runtime.

Go is the deterministic tooling layer.

Do not implement:
- An LLM client.
- A custom agent orchestrator.
- Background Go workers or daemons.
- A scheduler or message bus.
- An MCP server.
- SQLite or another database.
- A graph database.
- A web UI or dashboard.

## 3. Repository structure

Implement a minimal structure consistent with existing repository conventions:

```text
phys-research/
├── AGENTS.md
├── .opencode/
│   └── agents/
│       ├── researcher.md
│       ├── adversary.md
│       └── formalizer.md
├── research/
│   ├── sources/
│   ├── issues/
│   ├── hypotheses/
│   ├── reviews/
│   └── evidence/
├── derivations/
└── cmd/
    └── physctl/
```

Original PDFs should not be committed repeatedly. Research records should identify their source through DOI, arXiv ID/version, source hash and precise page, section, equation, figure or table locators.

## 4. Minimal research-record metadata

Use a small machine-readable metadata header, separate from Phys-Pseudo.

Proposed fields:

```yaml
---
id: C-001
type: issue
state: open
branch: C
priority: high
refs:
  - relation: motivated_by
    id: P-001
---
```

Supported record types:

- `source`
- `issue`
- `hypothesis`
- `review`
- `evidence`
- `derivation`

Workflow states:

- `open`
- `blocked`
- `closed`

The Q/C branch field is optional when relevant.

Use a small documented relationship vocabulary, initially:

- `motivated_by`
- `supports`
- `challenges`
- `addresses`
- `depends_on`
- `supersedes`

IDs must be readable and allocated by Go without collisions. Use `P-001`, `I-001`, `H-001`, etc., or the existing repository's compatible convention.

Do not use content hashes as filenames. Use PhysArtifact's existing identity conventions where canonical artifact identity is required.

Do not turn metadata into a second scientific language. Keep it minimal.

## 5. Go CLI

Implement only these five commands:

### `physctl new`

Create a record from an appropriate template, allocate its unique ID and write it safely.

### `physctl status`

List open and blocked issues, unresolved obligations and their references.

Support machine-readable JSON output.

### `physctl refs`

Traverse explicitly declared relationships. Do not infer relationships by searching arbitrary prose.

### `physctl validate`

Check:
- Required metadata fields.
- Valid record types and workflow states.
- Unique IDs.
- Reference existence and relationship validity.
- Source locator presence for extracted paper claims.
- Required falsification information for explicitly testable physical hypotheses.
- Record and artifact integrity.

Do not implement a general Phys-Pseudo parser or invent additional syntax. Respect the existing specification, including legitimate unresolved `?` statements.

Structural validation must never be presented as mathematical or physical verification.

### `physctl check`

Run only mathematical checks supported by the actual Phys infrastructure.

Use existing derivation fixtures and test conventions where applicable. Do not invent derivation rules, test oracles or unsupported mathematical capabilities.

A missing capability must produce an explicit diagnostic, not a false PASS.

All commands should be deterministic, have useful errors, and support JSON output where useful. Prefer the Go standard library and existing repository dependencies.

## 6. Research-record requirements

A research issue should distinguish:

**Theoretical expectation**
- What is claimed or predicted.
- Assumptions.
- Applicable regime.
- Source or derivation reference.

**Independent theoretical result**
- What has actually been reconstructed or derived.
- Its premises and scope.
- Derivation artifact or source reference.

**Experimental expectation**
- Predicted observable.
- Units and conditions.
- Uncertainty where applicable.
- A falsification condition.

**Experimental actual**
- Measured result.
- Units and uncertainty.
- Methodology and experimental conditions.
- Exact source locator.

Never label a paper's assertion as an independent derivation.

Never label a theoretical target as an experimental measurement.

Mathematical theorems do not require empirical falsifiers. They require explicit premises, conclusions, domains and mathematical proof obligations.

## 7. Append-only reviews and provenance

Reviews must be separate, append-only records.

Every review must identify:
- The hypothesis ID.
- The exact reviewed version or SHA-256 digest.
- Reviewer agent identity and model, recorded as self-reported metadata.
- Findings and supporting references.
- Blocking and nonblocking issues.

If the hypothesis changes, its previous review must be reported as stale.

Do not let the Adversary modify the hypothesis it reviews.

Every derivation must identify its premises, dependencies, intermediate results, mathematical rules and resulting obligations.

Preserve human-authored derivation sketches. The Formalizer must not invent missing pen-and-paper reasoning.

Use atomic creation and writes where appropriate. Avoid concurrent edits to the same record.

## 8. OpenCode agents

Create three concise role definitions.

### Researcher

Responsibilities:
- Identify relevant papers.
- Extract claims and source locators.
- Identify outstanding problems.
- Draft Phys-Pseudo hypotheses.
- Record theoretical and experimental expectations.

Cannot approve scientific results.

### Adversary

Responsibilities:
- Independently challenge assumptions and hypotheses.
- Search for counterexamples.
- Identify missing derivations, unsupported claims and regime mismatches.
- Create append-only review records.

Must use a fresh context and, preferably, a different model from the Researcher.

Cannot modify the reviewed hypothesis or accept the theorem.

### Formalizer

Responsibilities:
- Inspect agreed hypotheses and human-authored derivation sketches.
- Perform a capability audit.
- Transcribe supported derivations into PhysMath and PhysLib-compatible artifacts.
- Connect artifacts through PhysArtifact conventions.
- Run existing tests and report unresolved proof obligations.

The Formalizer is a transcriber and implementation assistant, not an independent mathematical authority.

Do not allow an agent to modify `AGENTS.md`, agent definitions, validation code or accepted scientific artifacts.

Inspect OpenCode's actual permission model before configuring these restrictions. Do not assume subagents inherit their parent's permissions.

## 9. Scientific acceptance boundary

Agents may propose, challenge, revise and test research records.

They cannot declare a theorem or physical hypothesis scientifically established.

The accepted Git branch must be protected from agent writes and merges. Human-controlled credentials and merge permissions must enforce this boundary.

Do not rely solely on prompts or a CLI command that agents can invoke.

Avoid requiring formal pull requests for routine source notes. Apply the protected merge boundary to accepted scientific artifacts.

## 10. Implementation milestones

### MVP 1 — Research tracking

Demonstrate:

One paper → one source record → one outstanding issue → one Phys-Pseudo hypothesis → one explicit falsification condition where applicable → one adversarial review → one human-controlled acceptance of the research record.

Use a manageable, already published scientific example.

Ensure the source claims have precise locators and that the adversarial review identifies the exact hypothesis version it examined.

### MVP 2 — Mathematical formalization

After MVP 1 works:

Human agreement on the research target → capability audit → human-authored derivation sketch → PhysMath representation → independent golden and anti-golden checks → reproducible derivation tests → human-controlled acceptance.

Goldens must not be generated solely by the implementation under test. The human or an independent reviewer must establish their correctness.

Do not pretend that passing tests proves the physical theory.

## 11. Testing requirements

Provide Go unit and integration tests for:

- ID allocation and collision handling.
- Record creation.
- Metadata validation.
- Broken and cyclic references.
- Relationship traversal.
- Source locator validation.
- Stale review detection.
- Append-only review behavior.
- Invalid state transitions, if implemented.
- Existing mathematical test integration.
- JSON output and meaningful nonzero exit codes.

Use temporary repositories and test files.

Test malformed, incomplete and adversarial inputs.

Do not add tests that merely repeat implementation assumptions as their own expected results.

## 12. Definition of done

MVP 1 is complete when:

1. The Go executable builds and its tests pass.
2. OpenCode can invoke the three configured roles using the installed version's valid configuration format.
3. A paper can be recorded with verifiable source locators.
4. An issue can be created and linked to source records.
5. A hypothesis can be expressed using the existing Phys-Pseudo specification.
6. A separate Adversary review can challenge the exact hypothesis version.
7. Expected and actual theoretical and experimental results remain distinguishable.
8. Structural validation catches broken records without claiming scientific correctness.
9. Agents cannot merge into the protected accepted branch.
10. The complete research trail can be inspected through ordinary files and Git.

## 13. Execution instructions

Work recon-first, then implement.

Do not expand scope to solve theoretical physics problems during this tooling implementation.

Do not rewrite the existing Phys infrastructure.

Do not introduce unnecessary dependencies, abstraction layers or configuration systems.

Do not silently assume the existing specifications support functionality that has not been verified.

Run tests and inspect the final Git diff.

At completion, report only:

1. What was implemented.
2. Files created or changed.
3. Tests actually executed and their results.
4. Remaining limitations.
5. Any decisions that require human approval.

**Primary engineering principle: build the smallest tool that reliably records scientific reasoning, preserves provenance, exposes unresolved problems and protects human scientific authority.**
