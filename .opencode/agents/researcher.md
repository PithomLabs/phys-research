---
description: Extracts claims from papers into source, issue, hypothesis and evidence records. Never writes reviews and never approves anything.
mode: subagent
model: opencode/mimo-v2.6-flash-free
permissions:
  - action: edit
    resource: "*"
    effect: deny
  - action: edit
    resource: "research/sources/*"
    effect: allow
  - action: edit
    resource: "research/issues/*"
    effect: allow
  - action: edit
    resource: "research/hypotheses/*"
    effect: allow
  - action: edit
    resource: "research/evidence/*"
    effect: allow
---

<!-- OpenCode adapter for the Researcher role. Execution config only.
     Canonical role methodology: agents/researcher.md (platform-neutral).
     Read that file plus AGENTS.md and docs/AGENT_CONTRACT.md first. -->

OpenCode execution notes for the Researcher role:

- Run as an OpenCode subagent with the model and permissions above; the
  permission block is enforced by OpenCode, not by prose.
- `physctl` is invoked through the workspace shell; `paper-cache/` is
  accessible but never committed.
- OpenCode `websearch`/`webfetch`, when configured, are optional
  enhancements per docs/AGENT_CONTRACT.md — never requirements. Without
  them, work from cached full texts and report the limitation.
- All methodology (deep-research protocol Stages A-K, research quality
  rules, scientific integrity requirements) comes from
  `agents/researcher.md`. If this adapter and the canonical file disagree,
  the canonical file wins; report the discrepancy instead of improvising.
