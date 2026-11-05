This is an important architectural correction to the previous prompt: OpenCode must be a supported execution environment, not a platform dependency. `phys-research` should be portable across coding agents, with the research system itself remaining independent of whichever agent runs it.

Append this prompt to the implementation task:

# Additional Architecture Requirement: Coding-Agent-Agnostic Portability

## Objective

Ensure `phys-research` can operate through different coding-agent environments, including OpenCode, Codex, Muse and other compatible agents.

OpenCode must NOT be treated as the mandatory runtime, architectural foundation or permanent platform dependency.

The research system must remain independently usable regardless of which coding agent provides reasoning, file access, tool execution and orchestration.

## 1. Separate the research system from the agent platform

Establish a clear separation between:

**A. Research core**

- Research artifacts and their schemas.
- Phys-Pseudo and research methodology.
- `physctl` CLI and deterministic operations.
- Literature acquisition and extraction.
- Validation, provenance and research-record conventions.
- Scientific integrity and human acceptance requirements.

**B. Agent execution environment**

- LLM and model selection.
- Agent instructions and role configuration.
- Tool availability and invocation.
- File editing and shell execution.
- Web search and document retrieval.
- Workflow orchestration.

The research core must not depend on OpenCode-specific APIs, configuration formats, agent definitions or execution behavior.

## 2. Establish a portable agent contract

Define the minimum capabilities that a coding agent must provide to operate `phys-research`.

Required capabilities:

1. Read and understand repository instructions.
2. Read and write research artifacts within the authorized workspace.
3. Execute the existing Go `physctl` commands.
4. Inspect command output and handle errors.
5. Follow the research methodology and scientific integrity requirements.
6. Maintain traceability between sources, claims, assumptions, equations and derivations.
7. Respect human-controlled scientific acceptance.

Optional capabilities:

- Web search and web fetching.
- PDF inspection.
- Parallel subagents.
- Native MCP integration.
- Specialized research tools.

Optional capabilities must not be silently treated as mandatory. Where unavailable, the agent should use existing repository functionality or explicitly report the limitation.

Document this contract in a platform-neutral location within the repository.

## 3. Make agent instructions platform-neutral

Identify and separate:

- Shared research instructions.
- Researcher role instructions.
- Platform-specific execution configuration.

The shared instructions must define what the agent should accomplish, how it should reason about physics, what artifacts it should produce and which boundaries it must respect.

Platform-specific files should contain only the configuration necessary to make a particular coding agent execute those instructions.

Avoid duplicating scientific methodology across different agent configurations.

OpenCode-specific files may remain as an adapter. They must not become the authoritative source of research methodology.

Do not assume every coding agent supports the same configuration format, subagent mechanism or permission system.

Inspect each platform's actual capabilities before proposing integration.

## 4. Preserve a stable CLI and artifact interface

The existing Go CLI should remain the primary deterministic interface between the coding agent and the research system.

Prefer:

- Stable CLI commands.
- Predictable exit codes.
- Clear, inspectable output.
- Explicit artifact paths.
- Plain-text research records.
- Git-compatible workflows.

Do not introduce MCP servers, additional APIs, SDKs or abstraction layers simply to achieve portability.

A coding agent capable of reading repository files and executing Go CLI commands should be able to perform the fundamental research workflow without a proprietary integration.

## 5. Demonstrate cross-platform operation

Evaluate at least:

- OpenCode.
- OpenAI Codex.
- Muse, subject to identifying the exact product and its actual capabilities.

Create a compatibility matrix:

| Capability                        | OpenCode | Codex | Muse |
| --------------------------------- | -------- | ----- | ---- |
| Read shared research instructions | Test     | Test  | Test |
| Execute physctl                   | Test     | Test  | Test |
| Create research artifacts         | Test     | Test  | Test |
| Validate research records         | Test     | Test  | Test |
| Conduct literature discovery      | Test     | Test  | Test |
| Complete Researcher workflow      | Test     | Test  | Test |

Use these statuses:

- VERIFIED
- PARTIALLY VERIFIED
- NOT VERIFIED
- UNSUPPORTED

Do not claim compatibility just because an instruction file can be opened by a platform.

Where an environment is unavailable, document the required setup and mark it unverified rather than inventing test results.

At minimum, demonstrate that the same Researcher workflow and artifact conventions work in OpenCode and one independent coding-agent environment, preferably Codex if available.

Use an isolated working copy or Git worktree for the cross-platform test. Do not allow simultaneous agents to modify the same research records.

## 6. Avoid accidental complexity

Do NOT:

- Build a universal agent orchestration framework.
- Create a custom agent runtime.
- Implement separate research logic for each platform.
- Add platform-specific dependencies to the Go core.
- Introduce a plugin architecture without a demonstrated need.
- Require all platforms to support MCP or subagents.
- Create a new agent role simply to facilitate portability.

Prefer one shared research methodology and minimal platform-specific adapters.

## 7. Definition of success

`phys-research` is coding-agent-agnostic when:

1. Its research methodology is independent of any vendor.
2. Its canonical artifacts are readable and writable by different coding agents.
3. Its deterministic operations are accessible through the stable Go CLI.
4. Different agents can execute the same research workflow without changing the scientific methodology.
5. Results remain compatible with the existing validation and provenance system.
6. Platform-specific limitations are explicitly documented.
7. No core research capability depends exclusively on OpenCode.

## 8. Implementation priority

First perform reconnaissance to identify existing platform coupling.

Then make only the changes necessary to establish the portable contract and demonstrate operation in an independent coding-agent environment.

Do not let cross-platform support delay the Branch C Scientific Capability Demonstration unnecessarily.

The ultimate objective is not to make `phys-research` an OpenCode application, a Codex application or a Muse application.

**It is to build an independent physics research system that can use whichever capable coding agent is available.**

## One architectural decision to preserve

I recommend this separation:

phys-research — platform-independent research core

Go CLI · Research artifacts · Phys-Pseudo · Validation · Provenance

Shared instructions and stable CLI contract

OpenCode

Adapter

Codex

Adapter

Muse

Adapter

This also preserves an important strategic advantage: you can change models, coding agents and providers without rewriting the research infrastructure or losing the scientific record. The implementation agent should establish this boundary without turning it into a universal orchestration framework.
