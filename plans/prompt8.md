Use this as the implementation agent's task. It deliberately prioritizes a working scientific research cycle over additional infrastructure.

# Implementation Task: Researcher Scientific Capability Demonstration

## Objective

Make `phys-research` demonstrate meaningful physics research capability.

The next milestone is not another infrastructure or security milestone. It is a complete, reproducible research cycle in which the existing Researcher can move from a physics question and published literature to a mathematically inspectable research dossier.

**Primary success criterion:** The Researcher must produce scientifically useful, traceable research outputs—not merely paper summaries.

## 1. Operating constraints

Work in the existing repository:

`/home/chaschel/Desktop/biz/TOE-v5/phys-research`

Before making changes:

1. Read the applicable `AGENTS.md` instructions.
2. Inspect the current repository structure, implementation, Researcher agent definition, `physctl` commands, existing research artifacts and validation rules.
3. Review the existing Branch C research, including P-002 through P-006, I-002, H-002, E-002 and D-001 where present.
4. Identify what the current system can already do and what demonstrably prevents a complete research cycle.

Do not assume previous implementation reports accurately describe the current repository. Verify the actual code and artifacts.

### Architectural constraints

- OpenCode remains the LLM runtime and orchestration layer.
- Use the existing Researcher role. Do not introduce additional agent roles.
- Go standard library only for custom Go implementation.
- Git and existing research artifacts remain canonical storage.
- Preserve Phys-Pseudo's seven constructs:\
  `given`, `define`, `assume`, `claim`, `derive`, `if/then/else`, `?`.
- Preserve existing repository architecture and artifact formats unless a demonstrated structural blocker requires a minimal change.
- No database, message bus, custom scheduler, MCP server, additional service or new orchestration framework.
- Do not change existing model assignments or introduce a new model without demonstrated necessity.

## 2. Research demonstration question

Use the existing Branch C investigation as the demonstration subject:

**Can classical Riemann–Cartan geometry be consistently coupled to quantum matter without quantizing spacetime, and what mathematical conditions must such a theory satisfy?**

Keep these three theoretical frameworks explicitly separate:

1. Semiclassical Einstein–Cartan gravity.
2. Stochastic or hybrid classical–quantum Einstein–Cartan dynamics.
3. Quantum Einstein–Cartan gravity.

Do not presume that mathematical results from one framework transfer to another.

Do not presume that the existence of a mathematically interesting formalism establishes its physical validity.

The existing Branch C artifacts are starting material, not authoritative conclusions. Reassess their actual contents and provenance.

## 3. Required Researcher capabilities

First determine whether each capability already exists. Improve only what is missing or demonstrably inadequate.

### A. Literature discovery and source qualification

The Researcher must:

- Find relevant primary literature.
- Distinguish foundational, directly relevant, related and tangential sources.
- Identify original papers rather than relying exclusively on reviews.
- Record exact source identifiers, publication status, versions and equation/section/page locators.
- Retrieve and inspect available full texts using the existing paper acquisition workflow.
- Distinguish source-author claims from independent analysis.

Keyword overlap is not evidence of scientific relevance.

### B. Mathematical claim and equation representation

For important claims, capture:

- The mathematical statement.
- Its equations and exact source locations.
- Definitions of variables and mathematical objects.
- Assumptions and boundary conditions.
- Applicable physical regime and effective-theory scale.
- Whether the claim is asserted by an author, derived in the source, or independently reproduced.

Do not turn an informal description into a theorem without a derivation.

### C. Assumption and framework tracking

For each theoretical framework, explicitly identify:

- Geometric variables and dynamical degrees of freedom.
- Matter content and coupling assumptions.
- Field equations and relevant constraints.
- Conservation requirements.
- Initial and boundary conditions.
- Approximation schemes and validity regimes.
- Mathematical assumptions that are not established by the cited sources.

Keep assumptions belonging to different frameworks separate.

### D. Cross-paper comparison

Produce a comparison that identifies:

- Shared mathematical structures.
- Different physical assumptions.
- Compatible and incompatible constraints.
- Different treatments of matter and geometry.
- Results that can legitimately be transferred.
- Results that cannot be transferred without additional derivation.
- Conflicts, unresolved questions and missing mathematical steps.

Every proposed connection between papers must explain the mapping from the source framework to the target framework.

### E. Derivation obligations

Convert unresolved mathematical steps into explicit, inspectable obligations using the existing artifact conventions and Phys-Pseudo.

Use `?` for unresolved steps.

Examples of questions worth investigating include:

- Under what conditions do the Hamiltonian and momentum constraints close?
- Which conservation laws follow from the proposed coupling?
- Does the proposed stochastic or hybrid evolution preserve the required constraints?
- What additional assumptions are necessary to establish consistency?
- Does a proposed matter coupling introduce additional degrees of freedom or instabilities?

These are research questions, not assumed findings.

Do not claim that an obligation has been resolved merely because a paper asserts the corresponding result.

### F. Competing hypotheses

Where the literature permits multiple interpretations, formulate competing hypotheses.

For each hypothesis, state:

- Its assumptions.
- Its mathematical commitments.
- What it explains.
- What remains unexplained.
- What derivation or calculation could distinguish it from alternatives.
- What observable consequences, if any, follow.

Keep all hypotheses proposed or contested until independently established.

### G. Experimental correspondence

Connect theoretical predictions to actual experiments and observations only when a defensible mathematical mapping exists.

Distinguish:

- A theoretical prediction.
- An actual measured quantity.
- An interpretation of the measurement.
- A proposed experimental test.
- An experiment that does not directly test the hypothesis.

Do not claim experimental support from superficial similarity between measured quantities and theoretical variables.

## 4. Required integrated research dossier

The Researcher must produce one coherent research dossier using existing artifact conventions.

It must contain:

1. **Research question and scope** — precise question, physical regime and boundaries.
2. **Literature map** — qualified primary and supporting sources with exact locators.
3. **Framework descriptions** — mathematical objects, assumptions and equations for each framework.
4. **Comparison matrix** — agreements, differences, compatibility and non-transferability.
5. **Established results** — what the sources actually demonstrate.
6. **Independent derivations** — what the Researcher can reproduce, with inspectable steps.
7. **Open obligations** — missing derivations and unresolved assumptions represented explicitly.
8. **Competing hypotheses** — alternatives and the mathematical results that could distinguish them.
9. **Experimental correspondence** — predictions and actual evidence where a valid connection exists.
10. **Research assessment** — strongest result, most important unresolved problem and recommended next investigation.

Use existing records and references wherever possible. Do not create redundant documents merely to satisfy a checklist.

The dossier must distinguish clearly between:

- Source-reported statements.
- Independently verified mathematics.
- Researcher conjectures.
- Empirical evidence.
- Unresolved questions.

Absence of evidence must remain absence of evidence.

## 5. Implementation strategy

Work incrementally.

**Step 1 — Reconnaissance**

Map the current Researcher workflow from question intake through source acquisition, analysis and artifact creation.

Identify concrete failures or gaps. Do not propose new abstractions based on hypothetical future needs.

**Step 2 — Prioritize**

Classify findings as:

- Functional blocker: prevents a meaningful research cycle.
- Research-quality defect: permits a cycle but damages scientific reliability.
- Enhancement: useful but not necessary for the demonstration.
- Accidental complexity: adds machinery without improving research capability.

Prioritize functional blockers and research-quality defects.

**Step 3 — Minimal implementation**

Modify only the components required to address demonstrated problems.

Prefer improvements to the existing Researcher instructions, artifact workflow and `physctl` functionality over new infrastructure.

Preserve existing interfaces and artifacts where feasible.

**Step 4 — Demonstrate**

Run the Researcher against the Branch C question.

Produce the integrated dossier and the associated research artifacts.

Do not fabricate a completed Researcher run if the current OpenCode environment cannot execute it. In that case, provide the exact invocation and clearly identify what remains untested.

**Step 5 — Evaluate**

Assess the actual output against the required capabilities and dossier.

Identify weaknesses honestly. Do not declare the system scientifically capable merely because it generates plausible-looking documents.

## 6. What is explicitly out of scope

Do not spend this task on:

- OpenCode permission hardening.
- OS-level agent isolation.
- Extensive security documentation.
- Additional security architecture.
- New agent roles.
- Autonomous adversarial research review.
- Autonomous theorem formalization or acceptance.
- Database or graph infrastructure.
- Custom scheduling and orchestration.
- Generic dashboards or user interfaces.
- Performance optimization without a demonstrated bottleneck.
- Refactoring unrelated packages.

Preserve existing validation and provenance safeguards. Fix a structural or security issue only if it actually blocks the required research cycle or threatens existing research records.

Do not enable unrestricted autonomous PDF fetching as part of this task.

## 7. Scientific integrity requirements

The human remains responsible for accepting scientific claims, derivations and hypotheses.

No agent may:

- Declare a hypothesis established without the required mathematical and evidential support.
- Treat a paper's assertion as an independently verified derivation.
- Hide incompatible assumptions.
- Conflate semiclassical, hybrid and fully quantum frameworks.
- Invent missing equations, citations or experimental results.
- Silently rewrite existing research conclusions to make the demonstration look successful.

Preserve the Math Reality Check discipline: mathematical consistency is necessary but does not by itself establish correspondence with physical reality.

## 8. Verification and deliverables

At completion, provide:

1. **Reconnaissance report:** What already worked and what was genuinely missing.
2. **Implementation report:** Exact files changed and why.
3. **Functional demonstration:** Research question, sources inspected, outputs produced and exact artifact paths.
4. **Scientific quality assessment:** Which derivations were independently reproduced, which remain assertions, and which obligations remain open.
5. **Validation results:** Existing tests and relevant `physctl` validation commands.
6. **Remaining blockers:** Only issues that materially prevent or compromise the next research cycle.
7. **Next recommended research action:** One highest-impact mathematical or literature investigation.

Preserve existing research-record contents and provenance unless a specific correction is necessary and explicitly justified.

Do not merge into an acceptance branch or claim human scientific approval.

## Definition of success

The milestone succeeds when the existing Researcher can take the Branch C question, work through relevant primary literature, compare the three frameworks without conflating them, identify mathematically meaningful unresolved obligations, and produce a traceable research dossier that a human physicist can continue investigating.

**The measure of progress is the quality of the physics research—not the amount of code, infrastructure, agents or documentation added.**

Start with reconnaissance. Implement only what the evidence says is necessary.

The key instruction is the last one: make the Researcher prove its capability through an actual research cycle. Do not let the implementation agent turn this into another architecture or infrastructure project.
