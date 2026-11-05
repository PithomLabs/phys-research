Plan 8 is substantially aligned with the direction we established. I would approve its overall approach, subject to four corrections before implementation.

It gets the priorities right: Branch C research first, portability second, and no new runtime, database, agent roles, MCP server or orchestration framework.&#x20;

plan8(4).md



## Four corrections

### 1. Do not automatically make the stale `agents/` directory canonical

This is the most important correction.

Plan 8 proposes making `agents/*.md` canonical and treating `.opencode/agents/*.md` as adapters. But it explicitly describes `agents/` as a stale mirror.&#x20;

plan8(4).md



Do not simply promote the stale copy.

Require the implementation agent to:

- Compare both versions of every agent instruction.
- Reconcile differences against Git history and the latest Researcher methodology.
- Preserve the most complete and up-to-date research instructions.
- Move model pins and platform-specific settings into the relevant adapters.
- Make the reconciled, platform-neutral instructions canonical.

The shared Researcher instructions must not lose the iterative deep-research methodology or its scientific integrity requirements.

### 2. Test Muse as well as OpenCode and Codex

The plan says Muse is the current execution environment, but its cross-platform matrix only explicitly requires OpenCode and Codex.&#x20;

plan8(4).md



Since all three CLIs are reported installed, test all three.

Use separate Git worktrees and the same research fixture. Verify:

- Instruction loading.
- `physctl` execution.
- Research artifact creation.
- Artifact validation.
- Literature lookup.
- Completion of the same Researcher workflow.

Do not mark a platform VERIFIED merely because it can read the instructions.

### 3. Separate core portability from optional capabilities

The plan correctly identifies web search, PDF inspection, parallel subagents and MCP as optional capabilities.&#x20;

plan8(4).md



Preserve that distinction in testing.

Mandatory core: Read shared instructions, execute `physctl`, create and validate research records, preserve provenance.

Optional enhancements: Web discovery, PDF retrieval, subagents and specialized integrations.

A coding agent must be able to conduct meaningful research using available sources and cached papers even when it lacks a particular web-search integration.

This prevents vendor-specific tools from becoming hidden requirements.

### 4. Correct the justification for deferring TLS

Plan 8 says six cached manifests prove retrieval worked and therefore TLS can remain deferred.&#x20;

plan8(4).md



Those manifests establish that previous retrievals occurred. They do not establish that current retrieval works.

The correct decision is:

- Use cached full texts for the Branch C demonstration if they adequately cover the required sources.
- Do not make network retrieval a blocker when those sources are sufficient.
- Mark live retrieval as unverified if the current probe fails.
- Do not enable unrestricted fetching on the strength of historical cache manifests.

This keeps us focused without confusing historical success with current verification.

## Scientific priorities: retain these as written

I particularly support the following parts:

| Research task                           | Assessment                    |
| --------------------------------------- | ----------------------------- |
| P-001–P-006 source qualification        | Essential                     |
| Load-bearing equation transcription     | Essential                     |
| Three frameworks kept separate          | Essential                     |
| D-001 transfer-mapping repair           | Highest-value comparison task |
| Seven explicit `?` obligations          | Essential                     |
| H-003 competing hypothesis              | Important                     |
| No unsupported `supports` relationships | Non-negotiable                |
| Ten-section traceable dossier           | Correct demonstration output  |

These requirements make the demonstration about actual physics research rather than document generation.&#x20;

plan8(4).md



One additional scientific condition: H-003 must be a genuinely distinguishable competing hypothesis, with its own assumptions and falsification condition. It must not exist merely to satisfy a record-count requirement.

## Additional instruction for the implementation agent

Append this to Plan 8:

## Final Amendments to Plan 8

Implement the following corrections:

1. **Canonical instruction reconciliation**
   - Do not promote `agents/*.md` as canonical merely because it exists.
   - Compare the shared instruction files and OpenCode copies.
   - Reconcile them against Git history and the latest Researcher methodology.
   - Preserve all current research capabilities and corrections.
   - Make the reconciled platform-neutral files canonical.
   - Keep model pins, permission syntax and execution-specific settings in platform adapters.
2. **Three-platform compatibility**
   - Test OpenCode, Codex and Muse because all three are reported available.
   - Use isolated Git worktrees and the same research fixture.
   - Record actual commands, outputs and limitations.
   - Never claim compatibility without executing the workflow.
   - Mark unavailable capabilities honestly.
3. **Portability contract**
   - Separate mandatory core capabilities from optional integrations.
   - The core research workflow must not require vendor-specific web search, MCP or subagents.
   - Demonstrate that each agent can operate the research core through shared instructions, repository artifacts and `physctl`.
4. **Cached literature and network retrieval**
   - Use adequate cached full texts for the Branch C demonstration.
   - Distinguish successful historical retrieval from verified current retrieval.
   - Do not enable unrestricted PDF fetching as part of this task.
5. **Scientific acceptance**
   - H-003 must represent a substantive competing hypothesis with distinct assumptions and a falsification condition.
   - Every load-bearing equation must have an exact source locator.
   - Every framework transfer must state the mapping and its limitations.
   - Keep unresolved mathematical questions as explicit `?` obligations.
   - Do not claim independent verification where only a source assertion exists.

Preserve Plan 8's execution order: Branch C demonstration first, portability second.

Do not expand the architecture to solve hypothetical future integration problems.

## Final decision

Approve with amendments

Proceed with Plan 8 after incorporating these corrections.

The central objective remains unchanged: build an independent physics research system whose scientific methodology survives changes in coding agents, models and vendors.

The strongest evidence of progress will be a meaningful Branch C mathematical investigation successfully reproduced through OpenCode, Codex and Muse—not another increase in infrastructure.
