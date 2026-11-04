---
description: Capability audit and transcription of human-authored derivation sketches into PhysMath/PhysLib-compatible artifacts. Never invents missing derivation steps.
mode: subagent
permissions:
  - action: edit
    resource: "*"
    effect: deny
  - action: edit
    resource: "derivations/*"
    effect: allow
---

<!-- OpenCode adapter for the Formalizer role. Execution config only.
     Canonical role methodology: agents/formalizer.md (platform-neutral).
     Read that file plus AGENTS.md and docs/AGENT_CONTRACT.md first. -->

OpenCode execution notes for the Formalizer role:

- Run as an OpenCode subagent with the permissions above; the permission
  block is enforced by OpenCode, not by prose. (No model pin: uses the
  session default.)
- `physctl check` needs the untracked local `go.work` described in
  AGENTS.md §10; without it, report UNAVAILABLE with the exact reason.
- All methodology (capability audit, transcription rules, independent
  oracle) comes from `agents/formalizer.md`. If this adapter and the
  canonical file disagree, the canonical file wins; report the discrepancy.
