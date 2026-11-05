That changes my last answer: the LLM client code in Go goes away. OpenCode becomes the LLM runtime, and Go keeps everything that must be trusted: state, gates, audit and storage. The split is simple:

| Piece | Lives in | Why |
|---|---|---|
| Role behavior (extract, formulate, review) | `.opencode/agents/*.md` | An agent is a prompt, model and permissions, and OpenCode runs it. Your repos already use `AGENTS.md` and `.opencode/plans`. |
| State, gates, digests, audit, scout | Go binary `trk` | Deterministic, table-testable, and no LLM can self-promote |
| Records | Plain files in `track/` | Agents read and write files natively, so no driver is needed. This reinforces the file-store recommendation. |

## `trk` (stdlib only)

```
trk scout            arXiv → src stubs           trk lint <hyp>     shallow structure check
trk next             next ticket + role + files   trk put <type>    validate, digest, atomic write
trk show|ls|refs     read                         trk advance <id>   gate-checked transition
trk audit <obl>      run go test for the fixture  trk approve <id>   human-only
```

- **`trk next`** replaces LLM planning. It emits the top actionable ticket, the role to run, and the exact files that role must read (PSEUDO.md, MRC_REVIEW_PROTOCOL.md, the manifest).
- **`trk lint`** is tracker-side only. It checks names, `?` placement and explicit `derive` results. It does not change the phys-pseudo spec.

## Gates enforced in code, not prompts

| Transition | Requires |
|---|---|
| open → hyp | `hyp` record references the issue, and `trk lint` passes |
| hyp → agreed | A review by a different agent ID **and** model, zero blocking findings, then `trk approve` |
| agreed → proved | An `audit` pass event for every `obl` of the hypothesis, then `trk approve` |
| any → blocked | Any agent, with a reason |

`trk approve` is denied in every agent's permissions, so only you can run it. The log records agent and model as self-reported strings. That is integrity, not authenticity, the same caveat your phys README already states.

## Status fields

I'd follow your existing two axes and not invent new ones. Each `iss` and `hyp` carries `state` (workflow), plus the math-versus-physical/empirical split from the derivation guidelines. Corpus status (`NONE`, `ESTABLISHED`, `CONTESTED`, `SUPERSEDED`, `FALSIFIED`) is human-curated only, matching phys's rule that the library never computes it.

## OpenCode wiring

```
.opencode/agents/orchestrator.md   primary: loop `trk next` → invoke role → repeat; stop at human gates
.opencode/agents/extractor.md      subagent
.opencode/agents/formulator.md     subagent
.opencode/agents/reviewer.md       subagent, different model, edit denied
AGENTS.md                          roles, "never claim a check you didn't run", two status axes
```

```markdown
---
description: Adversarial MRC reviewer. Findings only, no verdicts.
mode: subagent
model: <provider>/<model-other-than-formulator>
---
Read PSEUDO.md and MRC_REVIEW_PROTOCOL.md. Review the hypothesis from `trk show`.
Record findings with `trk put rev`. Never edit hyp files. Never run `trk approve`.
```

Two version cautions:
- The docs show frontmatter with `description`, `mode` (`primary`, `subagent`, `all`) and `model` as `provider/model`, with an optional variant.
- Permission syntax differs by version. The v2 docs use an ordered permissions array and deprecate legacy fields like `permission`. Check your build's schema.
- Lock down each agent's own permissions: a subagent currently uses its own configured permissions, not a restricted copy of its parent's.
- Headless runs vary by build, so check your build's flag before scheduling `trk scout` and the orchestrator loop.

## Repo shape

```
trk/  cmd/trk  store/ (Rec, Store, FileStore)  gate/  scout/  audit/
track/{src,iss,hyp,obl}/  log.jsonl
```

I can generate the skeleton: the `trk` module, the four agent files and `AGENTS.md`.
