I consolidated the four adversarial reviews. They identify several genuine weaknesses, but some proposed fixes introduce more machinery than the problems justify.

My recommendation is to retain the OpenCode + Go + Git architecture, with a few important safeguards added.

The guiding rule should be:

> Every safeguard must prevent a specific, credible failure. If it merely makes the system look more rigorous, remove it.

# 1. Consolidated adversarial findings

I would classify the findings into three groups.

## A. Must fix before autonomous research begins

Required

| Finding                                               | My assessment             | Consolidated fix                                                          |
| ----------------------------------------------------- | ------------------------- | ------------------------------------------------------------------------- |
| Human approval can be bypassed                        | Valid, critical           | Agents cannot merge into the accepted branch. Human controls the merge.   |
| Provenance is underspecified                          | Valid, critical           | Explicit references between research artifacts and derivation steps.      |
| Agent reviews can become stale                        | Valid                     | Each review identifies the exact version it reviewed.                     |
| Paper claims lack precise sources                     | Valid, critical           | Every extracted claim has a page, section or equation locator.            |
| Research status and scientific standing are conflated | Valid                     | Separate workflow state from epistemic standing.                          |
| Hypotheses lack falsification conditions              | Valid, with qualification | Every testable physical hypothesis specifies what would count against it. |

These are the important findings from Claude and DeepSeek.&#x20;

plan_review(6).md

plan_review(6).md



Human approval deserves particular attention.

A command that agents are instructed not to call is not a security boundary.

My solution is simple:

- Agents work on a research branch.
- The accepted branch is protected against agent writes and merges.
- Agents cannot access the human's merge credentials.
- You review and merge accepted derivations yourself.

A signed commit can supplement this, but a signature alone is not sufficient if the agent can use the signing key.

Do not require a formal pull request for every source note. Reserve the protected merge boundary for accepted scientific artifacts.

## B. Fix before the first actual derivation

Important

| Finding                                       | My assessment   | Consolidated fix                                                                |
| --------------------------------------------- | --------------- | ------------------------------------------------------------------------------- |
| `refs` is undefined                           | Valid           | Explicit typed references in file metadata.                                     |
| Formalizer may generate its own test oracle   | Valid, critical | Human-authored or independently constructed derivation goldens.                 |
| No capability audit                           | Valid, critical | Check PhysMath/PhysLib support before formalizing a theory.                     |
| Pen-and-paper work has no inspectable record  | Valid           | A short human-authored derivation sketch.                                       |
| Branch Q/C is implicit                        | Valid           | An optional `branch: Q` or `branch: C` field.                                   |
| Adversarial reviews may share the same errors | Valid           | Fresh context and preferably a different model.                                 |
| Unresolved questions can disappear silently   | Valid           | Append-only reviews and explicit resolution references.                         |
| Experimental evidence is underspecified       | Valid           | Record the source, measurement, uncertainty, units and experimental conditions. |

Claude's capability-audit objection is particularly important. The existing infrastructure must not be presumed to support torsion, density operators, integration or stochastic hybrid dynamics simply because the research hypothesis requires them.&#x20;

plan_review(6).md



DeepSeek correctly identifies the missing provenance chain and the importance of a human-authored derivation sketch.&#x20;

plan_review(6).md



I would also adopt Qwen's append-only review approach. An adversary creates a new review artifact rather than modifying the hypothesis it is challenging.&#x20;

plan_review(6).md



## C. Do not implement yet

Deferred or rejected

| Proposal                           | Decision         | Reason                                                           |
| ---------------------------------- | ---------------- | ---------------------------------------------------------------- |
| SQLite index from day one          | Defer            | Git is enough at the initial scale.                              |
| Graph database                     | Reject           | Explicit relationships in files suffice.                         |
| Full Phys-Pseudo AST parser        | Reject initially | Phys-Pseudo is not intended to impose a complete grammar.        |
| Automatic proof correctness report | Reject           | Structural checks cannot establish mathematical truth.           |
| Non-stop autonomous agent loop     | Reject           | One bounded research task per run.                               |
| Mandatory PR for every file        | Reject           | Human-controlled merge only where scientific acceptance matters. |
| Separate dependency graph service  | Reject           | References can be read from files.                               |
| Predeclared future SQLite schema   | Reject           | No need to design infrastructure we may never need.              |
| Ten or more agent roles            | Reject           | Three are enough to begin.                                       |

Gemini's performance concern is valid in principle, but its assertion that ordinary file scanning becomes problematic after only a few dozen papers is not substantiated. A simple scan of a modest research corpus is acceptable.&#x20;

plan_review(6).md



Qwen also correctly points out that Git does not automatically resolve simultaneous edits to the same file.&#x20;

plan_review(6).md



The solution is to avoid concurrent writers and use separate review files, not introduce a database and orchestration service.

# 2. The consolidated architecture

This is what I would freeze.

## Minimal physics research system

OpenCode

Human-directed orchestration; three specialized research roles.

Researcher

Adversary

Formalizer

One Go executable: `physctl`

ID allocation, structural validation, references, status and mathematical test execution.

Git research repository

All research records, hypotheses, reviews, evidence and derivation sketches.

Existing Phys infrastructure

Phys-Pseudo · PhysMath · PhysLib · PhysArtifact

Human-controlled scientific acceptance

No database, message bus, MCP server or custom orchestrator.

## 3. Minimal file organization

```
phys-research/
│
├── AGENTS.md
│
├── .opencode/
│   └── agents/
│       ├── researcher.md
│       ├── adversary.md
│       └── formalizer.md
│
├── research/
│   ├── sources/
│   ├── issues/
│   ├── hypotheses/
│   ├── reviews/
│   └── evidence/
│
├── derivations/
│   └── C-001/
│       ├── sketch.md
│       └── derivation.pseudo
│
└── cmd/
    └── physctl/
```

`AGENTS.md` contains common project rules. The individual OpenCode files contain short, role-specific instructions. Neither should duplicate the other's responsibilities.

One important rule: agents do not edit the same record simultaneously.

Issues and hypotheses can evolve through version-controlled changes, but reviews and evidence observations are append-only. A review must identify the exact hypothesis version it examined.

# 4. Minimal record format

Claude and DeepSeek correctly identify the absence of a formal relationship mechanism. Qwen correctly identifies why text searching is inadequate for dependable traversal.

I would introduce a small metadata header, entirely separate from Phys-Pseudo.

For example:

```
---
id: C-001
type: issue
state: open
branch: C
priority: high
refs:
  - relation: motivated_by
    id: P-001
  - relation: challenged_by
    id: P-002
---
```

The actual research content follows underneath.

The allowed metadata should be deliberately small:

| Field      | Purpose                                                  |
| ---------- | -------------------------------------------------------- |
| `id`       | Readable, stable identifier                              |
| `type`     | Paper, issue, hypothesis, review, evidence or derivation |
| `state`    | Open, blocked or closed                                  |
| `branch`   | Q or C, where relevant                                   |
| `priority` | Human-set task priority                                  |
| `refs`     | Explicit typed relationships                             |

The relation vocabulary should be small and controlled, for example:

`motivated_by`, `supports`, `challenges`, `addresses`, `depends_on`.

`physctl new` allocates an ID and creates a template. Use readable IDs, not content hashes. A content digest can still be recorded to identify the exact reviewed file version, without pretending it is the artifact's canonical identity.

This gives us a simple, machine-readable knowledge graph while retaining files as the source of truth.

# 5. Expected versus actual: one important refinement

I would preserve Gemini's strongest suggestion, but require evidence for each assertion.

Every comparison should identify:

| Field       | Meaning                                                   |
| ----------- | --------------------------------------------------------- |
| Expected    | What the paper or hypothesis predicts                     |
| Actual      | What was independently derived or experimentally measured |
| Assumptions | Conditions under which the comparison is valid            |
| Regime      | Energy, scale, curvature or matter conditions             |
| Source      | Paper, equation, page, figure or experimental dataset     |
| Origin      | Source claim, independent derivation or observation       |
| Difference  | Agreement, discrepancy or unresolved                      |

For experiments, include the measured value, units, uncertainty and relevant methodology.

For theory, do not present a copied paper assertion as an independent derivation.

And not every mathematical theorem needs an experimental falsifier. A testable physical hypothesis does need a failure condition; a mathematical theorem needs explicit premises and must withstand counterexamples and derivation review.

# 6. The five Go commands

I would keep the CLI extremely small:

| Command            | Function                                       |
| ------------------ | ---------------------------------------------- |
| `physctl new`      | Allocate an ID and create a record             |
| `physctl status`   | List open issues and outstanding obligations   |
| `physctl refs`     | Traverse explicit relationships                |
| `physctl validate` | Check metadata, references and source locators |
| `physctl check`    | Run available PhysMath derivation tests        |

`validate` must not pretend to parse or prove all of Phys-Pseudo.

It checks the research record's structure. Unresolved `?` statements are reported as obligations, not treated as errors.

`check` can replay only operations actually supported by the current Phys infrastructure. If no applicable implementation exists, it reports that the capability is unavailable.

The Formalizer cannot manufacture an independent test oracle. Before formalization, a capability audit identifies supported operations and missing functionality. The human supplies a concise derivation sketch; the Formalizer transcribes it into the supported representation.

This is a substantive safeguard, not bureaucracy.

# 7. Revised agent responsibilities

| Agent      | Can do                                                        | Cannot do                                              |
| ---------- | ------------------------------------------------------------- | ------------------------------------------------------ |
| Researcher | Discover papers, extract claims, create issues and hypotheses | Accept a theorem                                       |
| Adversary  | Independently challenge hypotheses and append findings        | Modify the hypothesis under review                     |
| Formalizer | Transcribe agreed mathematical work and prepare artifacts     | Invent missing derivation steps or certify correctness |

The Adversary should review a fresh snapshot, preferably using a different model. Its review records the target's commit or content digest. If the hypothesis changes, the previous review becomes stale.

The Researcher must cite precise source locations. A paper is untrusted input: any instructions embedded in its text are content to analyze, not instructions for the agent to execute.

OpenCode is not a persistent background scheduler. You or OpenCode explicitly initiate each research pass. I would limit each pass to one issue and one intended transition.

# 8. The first milestone must be smaller

DeepSeek makes a good point that attempting the whole paper-to-theorem lifecycle as the first MVP is too ambitious.&#x20;

plan_review(6).md



I would split it into two milestones.

MVP 1

Research tracking

One paper → one issue → one hypothesis → one testable prediction or stated limitation → one independent adversarial review → one human-controlled merge.

MVP 2

Mathematical formalization

Agreed hypothesis → capability audit → human derivation sketch → PhysMath representation → independent golden/anti-golden checks → human acceptance.

For MVP 1, I would use a manageable published result or a finite-dimensional hybrid toy model. I would not make the first implementation depend on successfully solving relativistic Einstein–Cartan hybrid dynamics.

Then use that system to investigate Branch C.

# 9. The final decisions

Here is the consolidated decision register.

| Decision              | Final choice                                                 |
| --------------------- | ------------------------------------------------------------ |
| Orchestration         | OpenCode, human-directed                                     |
| Agent implementation  | OpenCode roles; Go tools                                     |
| Number of roles       | Three                                                        |
| Go application        | One CLI                                                      |
| Storage               | Git and plain files                                          |
| Database              | None initially                                               |
| Relationships         | Minimal typed metadata                                       |
| IDs                   | Readable, Go-allocated                                       |
| Hashing               | Snapshot verification only; PhysArtifact owns identity       |
| Phys-Pseudo           | Existing seven constructs, unchanged                         |
| Validation            | Structural only                                              |
| Mathematical checking | Existing supported PhysMath operations                       |
| Reviews               | Fresh-context, append-only                                   |
| Experimental evidence | Explicit measurement and uncertainty                         |
| Human authority       | Protected accepted branch                                    |
| First milestone       | One paper, one issue, one hypothesis, one adversarial review |

The one important distinction I would not sacrifice for simplicity is the boundary between an agent's proposed result and a human-accepted scientific result.

Everything else should earn its complexity.

The aim is not to create a physics version of Jira. It is to create a small Go-powered research notebook that remembers every important scientific question, connects it to its evidence, identifies the next unresolved mathematical step, and lets OpenCode do the clerical and investigative work.

That is the design I would now adopt.
