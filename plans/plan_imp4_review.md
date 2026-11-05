The follow-up is successful, with one important qualification: cross-agent overwrite protection is now partially enforced, not universally guaranteed.

I consider this a good result without introducing unnecessary infrastructure.

## 1. What is actually protected

Verified improvements

| Protection                 | Result                                                                       |
| -------------------------- | ---------------------------------------------------------------------------- |
| Existing research records  | `O_EXCL` prevents replacement through `record.Create`                        |
| Concurrent record creation | Distinct IDs through race handling                                           |
| Existing downloaded PDFs   | Identical downloads are idempotent; different bytes produce collision errors |
| Separate agent workspaces  | Three worktrees and branches created                                         |
| Git merge conflicts        | Both sides remain available for explicit resolution                          |
| Existing Branch C records  | Hash-identical after remediation                                             |

The baseline covers 63 files. Sixty remained byte-identical; the three changes were the intended implementation edits. Branch C records and your `AGENTS.md` remained untouched.&#x20;

plan_imp4.md



The test suite also passed, with nine packages and zero validation errors.&#x20;

plan_imp4.md



## 2. Three remaining qualifications

A. Worktree isolation is not filesystem isolation.

The report correctly acknowledges that an agent with direct filesystem access can modify files outside its assigned worktree. Worktree separation prevents accidental collisions under normal operation, but does not technically prevent a misbehaving agent from accessing another worktree.&#x20;

plan_imp4.md



For our current workflow, I consider this acceptable as an operational safeguard, provided every agent is explicitly assigned its own working directory.

B. Proposed outputs are protected by convention, not enforcement.

The proposed-version location is `research/proposed/<agent>/<path>`, but tooling ignores it and the policy is documented rather than enforced.&#x20;

plan_imp4.md



I recommend one small improvement:

`research/proposed/<agent>/<run-id>/<artifact-path>`

This prevents separate runs by the same agent from accidentally sharing destinations.

C. The worktrees start from the old HEAD.

This is the most immediate operational consideration. The report says the new worktrees start at pre-freeze HEAD, while the latest Branch C baseline remains uncommitted in the main checkout.&#x20;

plan_imp4.md



Consequently, Muse, Codex and OpenCode will not automatically see the latest research work in their new worktrees.

Do not assume those worktrees contain the completed Branch C dossier. Before assigning further research, establish a reviewed snapshot of the current baseline and deliberately make it available to each agent.

## 3. My decision

Accept as interim implementation

Do not expand this into a workspace-management framework, database, locking service or agent runtime.

The existing solution is proportionate to our needs:

- Git worktrees for independent work.
- Unique branches for independent changes.
- Create-only record operations.
- Collision-safe paper acquisition.
- Human-controlled integration.
- Preservation of competing scientific outputs.

I would authorize the next research cycle after ensuring the agents receive the correct baseline.

## 4. One small follow-up instruction

Make only these final adjustments to cross-agent output handling:

1. Change the proposed-output convention to `research/proposed/<agent-id>/<run-id>/<artifact-path>`.
2. Ensure a new run never reuses or silently replaces an existing run directory.
3. Document that each coding-agent process must be launched with its assigned worktree as its working directory.
4. Document that newly created worktrees do not automatically include uncommitted changes from the main checkout.
5. Provide a repeatable procedure for distributing a human-reviewed research baseline to agent worktrees without overwriting existing agent outputs.

Do not alter existing research records, redo the Branch C investigation, add dependencies, or introduce new infrastructure.

Run relevant tests and stop.

Then return to physics. The next high-value task remains the O5 investigation of the momentum constraint in Einstein–Cartan theory. We should now spend more effort resolving that mathematical question than improving the mechanics of agent coordination.
