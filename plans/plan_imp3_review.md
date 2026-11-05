I reviewed `plan_imp3.md`. Plan 8 has produced meaningful progress on the physics research workflow and agent portability, but it has not addressed the cross-agent overwrite problem we discussed immediately before this report.

My assessment follows.

## 1. What was accomplished

Substantial progress

The implementation report documents several valuable improvements:

- Reconciled the Researcher methodology into platform-neutral instructions while preserving the deep-research stages.
- Repaired D-001 and added explicit framework-transfer mappings.
- Added mathematical content and equation locators to P-003 through P-006.
- Corrected questionable DOI and publication metadata.
- Converted seven H-002 research gaps into explicit `?` obligations.
- Created H-003 as a competing hypothesis.
- Produced the ten-section Branch C dossier.
- Fixed a demonstrated proxy issue in the paper-fetching pipeline.&#x20;

  plan_imp3.md



These are meaningful functional improvements, not infrastructure for its own sake.

## 2. Three issues I would not overlook

A. Scientific capability is not yet demonstrated at the derivation level.

The report explicitly says that no mathematical derivation was independently reproduced. The equations are source-reported or conjectural.&#x20;

plan_imp3.md



Therefore, the current achievement is a traceable literature investigation and mathematical problem decomposition—not yet an independently validated physics result.

The recommended O5 investigation is appropriate:

> Extend the P-003 Dirac-argument method to Einstein–Cartan variables for the momentum constraint.

This is the highest-impact next research activity because it could expose a genuine mathematical obstruction or establish a necessary consistency condition.

B. Portability is only partially demonstrated.

Muse was tested in an isolated worktree, but Codex and OpenCode remain unverified.&#x20;

plan_imp3.md



Do not advertise full cross-platform compatibility yet.

C. There is an unexplained obligation-count discrepancy.

The report describes ten open obligations—seven from H-002 and three from H-003—but says `physctl status` shows twelve `?` lines.&#x20;

plan_imp3.md



Have the implementation agent reconcile those counts and ensure the extra two lines are not duplicated or orphaned obligations.

## 3. Critical omission: cross-agent non-overwrite protection

This is the most important issue relative to our immediately preceding instruction.

The report establishes that Muse was tested in an isolated worktree, but does not demonstrate:

- Unique ownership of agent-generated outputs.
- Create-only file operations.
- Protection against writing over another agent's existing artifact.
- Collision handling between concurrent processes.
- Preservation of both results when two agents investigate the same question.
- Adversarial tests verifying that another agent's files cannot be overwritten.

plan_imp3.md

A separate worktree is a good operational safeguard, but it is not a complete non-overwrite policy.

Do not assume the previous guardrail task was implemented.

## 4. My recommended next action

Freeze the current uncommitted work for review. Before launching another coding agent against the shared research directory, establish its baseline and protect the outputs already created.

Then give the implementation agent this additional instruction:

# Follow-up: Cross-Agent Research Output Integrity

The previous implementation report does not demonstrate completion of the cross-agent non-overwrite requirement.

Implement and verify this requirement without disturbing the completed Branch C work.

1. Inventory the entire current repository state, including uncommitted and untracked files. Record SHA-256 hashes of all existing research outputs.
2. Do not reset, clean, restore, replace or delete existing work.
3. Establish separate Git worktrees and branches for Muse, Codex and OpenCode sessions.
4. Implement create-only behavior for agent-generated outputs. Existing destinations must cause an explicit collision error, never silent replacement.
5. Preserve existing canonical research records. Revisions must create separately identifiable proposed versions.
6. Ensure concurrent agents can produce competing hypotheses and derivations without overwriting one another.
7. Add tests for concurrent creation, existing-file collisions, separate worktrees, preservation of unowned files and conflicting Git changes.
8. Verify file contents and hashes before and after the tests.
9. Preserve all Branch C changes, equation transcriptions, D-001, H-002, H-003, E-002 and the dossier.
10. Make no unrelated changes to scientific conclusions or research methodology.

Do not introduce a database, coordination server, new runtime or distributed locking system.

Report exactly what is protected, what is enforced, what is merely documented and what remains unresolved.

Stop after verification. Do not commit, merge or delete any existing worktree without explicit authorization.

## Final priority

I would now sequence the work as follows:

1. Preserve the current research outputs and enforce non-overwrite protection.
2. Continue the O5 mathematical investigation.
3. Demonstrate the same research workflow in Codex and OpenCode.
4. Expand literature coverage only when it helps resolve the mathematical obligations.

The key distinction is that we now have a research workflow capable of organizing a serious investigation. The next meaningful scientific advance should come from resolving a mathematical obligation, not building more infrastructure.
