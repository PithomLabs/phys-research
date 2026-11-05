Use OpenCode as your research workspace. You do not need to operate the agents independently, manage a queue, or manually maintain a database. You give OpenCode a research task, select the appropriate agent, inspect its output, and move to the next stage.

Here is the actual operating procedure for your MICROSCOPE MVP.

# 1. Start the research project

Open a terminal:

```
cd /path/to/phys-research

git status --short --branch
git branch --list
```

Make sure you are on the working branch (`main`), not the protected `accepted` branch.

Build and test the Go tool:

```
go test ./...
go build -o ./bin/physctl ./cmd/physctl
```

If the implementation requires your developer-local Go workspace, use the `GOWORK` path established during implementation:

```
GOWORK=/absolute/path/to/go.work go test ./...
GOWORK=/absolute/path/to/go.work go build -o ./bin/physctl ./cmd/physctl
```

Check that the CLI works:

```
./bin/physctl --help
./bin/physctl status --json
```

Use the actual `--help` output as authoritative if the implemented command syntax differs.

# 2. Open OpenCode

From the project root:

```
opencode
```

OpenCode will load the project configuration and discover the Markdown agents.

Configuration detail worth checking: OpenCode V2 uses `opencode.json` or `opencode.jsonc` in the project root for project configuration. Agent definitions belong under `.opencode/agents/`. Your implementation notes previously placed configuration at `.opencode/opencode.json`; make sure the permissions are actually in the root configuration file before relying on them.&#x20;

[image](https://www.google.com/s2/favicons?domain=https://docs.opencode.ai\&sz=32)

OpenCode+1



Your three agents should appear in OpenCode's agent discovery:

- `researcher`
- `adversary`
- `formalizer`

They should be configured as subagents. You can invoke them by mentioning them with `@` in the primary session. Each subagent gets its own context and configured permissions.&#x20;

[image](https://www.google.com/s2/favicons?domain=https://dev.opencode.ai\&sz=32)

OpenCode+1



# 3. First session: ingest the MICROSCOPE paper

Start with the Researcher.

Type this into OpenCode:

Researcher task

First run

```
@researcher

Process the MICROSCOPE final results paper.

DOI: 10.1103/PhysRevLett.129.121102
arXiv: 2209.15487v1

Create the MVP 1 research records:

P-001: Source record.
I-001: Outstanding research issue.
H-001: Phys-Pseudo hypothesis.
E-001: Experimental evidence.

Requirements:
- Verify the paper and its exact source locators.
- Preserve statistical and systematic uncertainties separately.
- Keep theoretical expectation separate from experimental results.
- Do not invent an independent derivation.
- Include the stated assumptions and experimental regime.
- Preserve the two unresolved ? obligations.
- Include the falsification condition in the research record, not as a new Phys-Pseudo construct.
- Use physctl to create and validate records.
- Do not approve or merge anything.

Report created files, IDs, source locators and unresolved issues.
```

The Researcher creates the proposed records. You inspect them before moving on.

The Researcher should use its configured research capabilities to access the paper. Alternatively, place the PDF in the implementation's ignored local paper cache and give it the path.

# 4. Inspect the results

Return to your main OpenCode session and ask:

```
Show the current research status.

Run:
- physctl status --json
- physctl validate

Show the references for I-001.

Identify missing source locators, broken references,
unsupported claims and unresolved obligations.

Do not modify accepted scientific artifacts.
```

You should see something conceptually like:

```
P-001  SOURCE       Created
I-001  ISSUE        Open
H-001  HYPOTHESIS   Proposed
E-001  EVIDENCE     Recorded

Open obligations: 2
Broken references: 0
Stale reviews: 0
```

This is illustrative output, not a claim that your implementation has already produced it.

# 5. Second session: challenge the hypothesis

Now invoke the Adversary in a fresh context:

Adversary task

```
@adversary

Independently review H-001 against P-001 and E-001.

Focus on:
1. The delta-to-eta identification.
2. Thermal systematic uncertainty.
3. Glitch masking and M-ECM reconstruction.
4. The interpretation of statistical and systematic errors.
5. Whether the falsification condition is operational.
6. Whether the hypothesis generalizes beyond the
   actual MICROSCOPE experimental regime.

Use the MRC review protocol.

Record findings in a new, append-only review record.
Record the exact SHA-256 of H-001 that you reviewed.

Do not modify H-001.
Do not approve or merge.
```

Once the review is complete, inspect its findings.

If you change H-001, the previous review becomes stale. Ask the Adversary to review the revised version again.

This gives you an actual independent challenge instead of an agent merely approving its own work.

# 6. Human acceptance

Once you are satisfied:

Run:

```
./bin/physctl status --json
./bin/physctl validate
git status --short
git diff
```

Review the complete research record and the Adversary's findings.

Then, outside OpenCode, perform the human-controlled acceptance merge into `accepted`, using the protected branch and separate human credentials established during implementation.

This completes MVP 1.

# 7. When to use the Formalizer

You do not need the Formalizer for ordinary paper ingestion or initial hypothesis creation.

Use it when you have agreed on a mathematical question and actually performed the derivation.

Your workflow then becomes:

1. You formulate the mathematical argument on paper.
2. You write a short `sketch.md` containing its steps.
3. Invoke `@formalizer`.
4. The Formalizer checks whether PhysMath and PhysLib support the required operations.
5. It maps your steps to existing mathematical representations.
6. It prepares the derivation artifact and test fixtures.
7. You independently review and accept the result.

Example:

```
@formalizer

Read:
- H-001
- its accepted review
- derivations/C-001/sketch.md

Perform the capability audit.

Transcribe only the derivation steps explicitly
provided in the sketch.

Use existing PhysMath rules and PhysLib conventions.

Identify unsupported operations and unresolved proof
obligations. Do not invent missing mathematical steps.

Prepare the artifacts and run available tests.
Do not claim that passing tests proves physical truth.
```

OpenCode V2 supports assigning models and permissions separately to each agent, so the Adversary can use a different model from the Researcher. Make sure that model is actually available in your configured providers.&#x20;

[image](https://www.google.com/s2/favicons?domain=https://dev.opencode.ai\&sz=32)

OpenCode+1



# 8. Your everyday operating routine

Once the initial records exist, this is essentially all you need to do.

| Your instruction                                         | What happens                                       |
| -------------------------------------------------------- | -------------------------------------------------- |
| `@researcher Investigate this paper.`                    | Adds scientific knowledge and outstanding problems |
| `@researcher Develop a hypothesis for I-001.`            | Proposes a Phys-Pseudo hypothesis                  |
| `@adversary Challenge H-001.`                            | Independently attacks the hypothesis               |
| `@formalizer Formalize this accepted derivation sketch.` | Prepares mathematical artifacts                    |
| `Show outstanding problems.`                             | Retrieves current research status                  |
| `Trace dependencies of H-001.`                           | Shows linked papers, hypotheses and derivations    |

You can also use OpenCode's non-interactive `opencode run` for later scripted tasks, but I recommend the interactive TUI for your initial research sessions because you can inspect and approve each transition.&#x20;

[image](https://www.google.com/s2/favicons?domain=https://opencode.ai\&sz=32)

OpenCode+1



## The most useful mental model

## You only need to manage three things

Research direction

Tell the Researcher what question deserves investigation.

Scientific disagreement

Let the Adversary identify what is wrong, missing or unsupported.

Mathematical acceptance

Perform the derivation and decide what deserves formalization and acceptance.

Everything else—paper records, references, issue creation, version tracking and validation—is clerical work that OpenCode and Go should handle.

My advice: Start with one complete MICROSCOPE cycle. Do not invoke all three agents simultaneously. Researcher → inspect → Adversary → revise if necessary → human acceptance. Only then move to mathematical formalization.
