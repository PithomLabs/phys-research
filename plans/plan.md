After comparing all four proposals, I would take a hybrid of Claude's minimal Go tracker, DeepSeek's deterministic command-line interface, Gemini's expected-versus-actual tracking, and Qwen's Git-native storage.

But I would remove more than I would retain.

My final recommendation is:

# OpenCode + Go + Git

## No database initially. No custom orchestrator. No graph database.

| Component           | Responsibility                                   |
| ------------------- | ------------------------------------------------ |
| OpenCode            | Orchestrate research agents                      |
| Agent instructions  | Small OpenCode role files                        |
| Go CLI              | Deterministic validation and research operations |
| Git and files       | Single source of truth                           |
| Phys-Pseudo         | Hypotheses and scientific reasoning              |
| Phys infrastructure | Mathematical derivation and formalization        |

## 1. What I would take from each proposal

| Source   | Keep                                                                                         | Remove                                                                                             |
| -------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| Claude   | Small Go tracker, file storage, atomic writes, validation, explicit human approval           | Separate tracker architecture, unnecessary workflow automation                                     |
| DeepSeek | One Go CLI, structured JSON output, deterministic gates, human-controlled theorem acceptance | SQLite as mandatory storage, excessive commands and status transitions                             |
| Gemini   | Expected-versus-actual comparisons, issue tracking, relationship traversal                   | Database-dependent workflow and custom orchestration                                               |
| Qwen     | Git as the canonical record, content history, Go validation, branch-based review             | JSON wrappers for everything, forced no-prose rules, hash-named files, invented Phys-Pseudo syntax |

### Claude: the best operational simplicity

I particularly like Claude's proposal that Go owns deterministic operations while OpenCode owns reasoning.

Its `trk` design has useful commands for finding work, validating records, inspecting relationships and auditing derivations. Its distinction between workflow state and scientific standing is also worth retaining.&#x20;

claude(1).md



I would simplify the command set and avoid creating another tracking convention alongside the existing Phys infrastructure.

### DeepSeek: the best separation of responsibility

DeepSeek makes the important distinction:

> The Go binary is not an agent. It is a deterministic CLI tool.

Exactly.

OpenCode performs research judgments. Go validates files, calculates identifiers, checks references and runs reproducible tests.&#x20;

deepseek.md



I also like machine-readable JSON output and nonzero exit codes when validation fails.

I would not retain its SQLite requirement or its elaborate automatic promotion rules.

### Gemini: the best research-management feature

Gemini correctly emphasizes the expected-versus-actual comparison.

This is essential because we need to distinguish what a paper claims from what we independently derive, and what a theory predicts from what an experiment measures.&#x20;

gemini(1).md



Keep this feature, but store it in research files rather than requiring database tables.

### Qwen: the best storage simplification

Qwen makes the strongest argument for Git-native research records. OpenCode naturally works with files, directories, diffs and commits.&#x20;

qwen.md



I agree with its storage direction.

However, I would reject several implementation details:

- Phys-Pseudo must not be restricted to four starting keywords. Its seven constructs must remain intact.
- Natural-language source attribution and explanations must not be banned.
- `FalsificationCondition` must not be invented as a new Phys-Pseudo construct.
- Physical terms such as gravity must not be indiscriminately prohibited from research hypotheses.
- Hash-named files are unnecessary when PhysArtifact already owns artifact identity.
- Passing a Go validator does not establish scientific correctness.

The proposed validator also has a Go implementation bug: its local `filepath` variable shadows the imported package, preventing the subsequent package-qualified call.&#x20;

qwen.md



# 2. The simplified implementation I would actually build

One Go executable: `physctl`.

No Go daemons. No LLM client. No scheduler. No database server.

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
│
├── cmd/
│   └── physctl/
│
└── go.mod
```

Each directory has one purpose.

| Directory     | Contents                                                  |
| ------------- | --------------------------------------------------------- |
| `sources`     | Paper identity, references, relevant claims and equations |
| `issues`      | Unresolved physical or mathematical problems              |
| `hypotheses`  | Phys-Pseudo formulations                                  |
| `reviews`     | Adversarial findings and unanswered objections            |
| `evidence`    | Experimental results and theory comparisons               |
| `derivations` | Agreed mathematical work using the Phys infrastructure    |

Original PDFs can be held in a local cache rather than committed repeatedly into Git. Keep DOI/arXiv identifiers, version and source hash in the research record.

No separate JSON artifact envelope is needed for every Pseudo file. Use the existing PhysArtifact conventions where canonical artifact identity is required.

## 3. Only five Go commands

I would start with these.

| Command          | Responsibility                                                                        |
| ---------------- | ------------------------------------------------------------------------------------- |
| `physctl lint`   | Validate Phys-Pseudo structure and artifact references                                |
| `physctl status` | List outstanding problems and unresolved derivation obligations                       |
| `physctl refs`   | Follow relationships between papers, problems, hypotheses and theorems                |
| `physctl check`  | Run available mathematical replay and derivation tests                                |
| `physctl hash`   | Calculate or verify artifact identities using the applicable PhysArtifact conventions |

Every command should support machine-readable JSON output.

For example:

```
physctl status --open --json
physctl refs C-001 --depth 3 --json
physctl lint research/hypotheses/C-001.pseudo
physctl check derivations/C-001/
```

No `scout` command is needed initially: OpenCode can discover and analyze papers.

No `approve` command is needed initially: human acceptance happens through review and Git integration.

No `next` command is necessary until choosing the next issue becomes genuinely tedious.

## 4. The only tracking concept that really matters

I would make every research issue answer four questions:

1. What is established?
2. What remains unresolved?
3. What do competing hypotheses predict?
4. What actually happens mathematically or experimentally?

For example:

C-001: Relativistic hybrid dynamics

Open

Question

Can classical Riemann–Cartan geometry couple consistently to quantum matter through stochastic dynamics?

Expected — Theory

A hybrid evolution preserving complete positivity, constraint closure and finite renormalized sources.

Actual — Independent derivation

Not established. The relativistic Markovian construction has a known divergence problem.

Expected — Experiment

A physically distinguishable decoherence and metric-diffusion prediction.

Actual — Experiment

Existing measurements constrain portions of the parameter space; the proposed Einstein–Cartan-specific prediction has not been established.

Next mathematical question

Does a non-Markovian evolution admit a consistent local Lorentz and diffeomorphism constraint structure?

This is enough to keep the research focused without creating elaborate ticket workflows.

Importantly, `Expected` and `Actual` have different meanings in theory and experiment. We should never put them into one undifferentiated pass/fail field.

## 5. Keep the agents to three

OpenCode remains the orchestrator.

Researcher

Reads papers, identifies claims, creates issues and proposes Phys-Pseudo hypotheses.

Adversary

Challenges hypotheses, checks assumptions, identifies missing derivations and searches for counterexamples. Findings only; no scientific approval.

Formalizer

After human agreement, prepares the mathematical derivation artifacts using PhysMath, PhysLib and PhysArtifact.

The roles are small OpenCode agent configurations, not independent Go programs. All custom executable tooling is Go.

I would use the existing `AGENTS.md` conventions rather than invent a new `.opencode.md` constitution. Keep each agent's instructions short and task-specific.

# 6. What I would deliberately NOT build

This is where we save most of the effort.

| Proposed feature               | Decision         | Reason                                                  |
| ------------------------------ | ---------------- | ------------------------------------------------------- |
| SQLite                         | Remove initially | Git already holds the canonical records                 |
| Graph database                 | Remove           | Typed relationships and simple traversal are sufficient |
| JSON for every artifact        | Remove           | Phys-Pseudo files are more useful to agents and humans  |
| Custom message bus             | Remove           | OpenCode coordinates                                    |
| Go agent daemons               | Remove           | OpenCode provides the model runtime                     |
| Separate orchestration service | Remove           | Duplicates OpenCode                                     |
| Automatic theorem approval     | Remove           | Scientific acceptance belongs to humans                 |
| Ten or more specialist agents  | Remove           | Three roles cover the initial work                      |
| Mandatory GitHub Actions       | Defer            | Local Go validation is sufficient initially             |
| Custom dashboard               | Remove           | Inspect records through OpenCode and CLI                |
| Mandatory PR for every change  | Defer            | Local branches and ordinary Git reviews suffice         |

SQLite can be added later as a disposable, rebuildable search index if repository scanning becomes measurably slow. It should never become a competing source of truth.

# 7. One essential safeguard: distinguish validation from truth

I would retain Claude's and DeepSeek's strongest idea: gates must be deterministic, not just instructions in an agent prompt.

But the gates should be small.

| Operation                                     | Agent allowed? |
| --------------------------------------------- | -------------- |
| Create a source record                        | Yes            |
| Create an issue                               | Yes            |
| Propose a hypothesis                          | Yes            |
| Record an adversarial finding                 | Yes            |
| Run a mathematical test                       | Yes            |
| Mark a proof obligation as unresolved         | Yes            |
| Declare a theorem physically established      | No             |
| Accept a derivation as independently verified | No             |

There is a subtle security point here: a human-only command in a Go executable is not genuinely human-only if an agent has unrestricted access to that executable and the filesystem. A real acceptance boundary needs a human-controlled Git merge or permissions boundary.

Also, a hypothesis containing `?` is not necessarily invalid. The question mark is useful precisely because research has unresolved obligations. The linter should detect malformed syntax, not force every hypothesis to pretend it is complete.

# 8. The final research workflow

## One command-driven research cycle

1. Give OpenCode a paper or research question.
2. Researcher creates the source record and extracts claims.
3. Researcher creates or updates the relevant outstanding issue.
4. Adversary challenges the assumptions and proposed hypothesis.
5. You agree on the question and premises.
6. You derive the result using pen and paper.
7. Formalizer prepares the executable Phys research artifact.
8. Go validates structure and runs reproducible checks.
9. You review and accept the scientific result through Git.

## Final decision

I would adopt:

Claude's lightweight Go tracker + DeepSeek's deterministic CLI + Gemini's expected/actual discipline + Qwen's Git-native storage.

But simplify the result further:

- One Go executable.
- Three OpenCode roles.
- Five CLI commands.
- One Git repository.
- No database initially.
- No custom orchestration.
- No invented Phys-Pseudo syntax.
- No automatic scientific approval.

The guiding principle should be:

Automate the clerical work. Make the reasoning inspectable. Keep mathematical proof and physical acceptance rigorous.

The first implementation should demonstrate one complete research issue from paper ingestion through adversarial review to a human-approved derivation. Nothing else is necessary until that workflow proves useful.
