# Agent Contract: Operating phys-research From Any Coding Agent

`phys-research` is a platform-independent research core: Go CLI, research
artifacts, Phys-Pseudo methodology, validation and provenance. Any coding
agent meeting the required capabilities below can operate it. OpenCode,
Codex and Muse are execution environments (adapters), not platform
dependencies. No core research capability depends on any one of them.

## Required capabilities (mandatory core)

1. Read repository instructions (`AGENTS.md`, `agents/<role>.md`, this file).
2. Read and write research artifacts within the authorized workspace.
3. Execute `physctl` commands and inspect their output and exit codes.
4. Handle errors explicitly (`UNAVAILABLE` stays `UNAVAILABLE`; never a PASS).
5. Follow the research methodology and scientific integrity requirements
   (canonical role files under `agents/`; acceptance boundary below).
6. Maintain traceability between sources, claims, assumptions, equations
   and derivations (exact locators; `refs` edges; SHA-256 provenance).
7. Respect human-controlled scientific acceptance: no agent declares any
   theorem or hypothesis established.

## Optional capabilities (enhancements, never hidden requirements)

- Web search and web fetching (platform-provided).
- PDF inspection beyond `physctl paper extract`.
- Parallel subagents.
- Native MCP integration or specialized research tools.

When an optional capability is unavailable, the agent works from cached
full texts and repository artifacts, and reports the limitation explicitly.

## Stable interface

- CLI: `physctl new | status | refs | validate | check | paper` (see
  `physctl help`). Exit codes: 0 clean, 1 findings/failures, 2 usage or
  unavailable capability. Standard library only; plain-text output.
- Records: Markdown with restricted frontmatter under `research/*` and
  `derivations/` (see `AGENTS.md` §6). IDs from `physctl new`, never by hand.
- Storage: Git working tree; `paper-cache/` and `go.work` stay untracked.
- Phys-Pseudo: exactly the seven constructs (`given`, `define`, `assume`,
  `claim`, `derive`, `if/then/else`, `?`).

## Acceptance boundary (all platforms)

`main` (or `research/*`) is the working branch. Scientific acceptance is a
merge into `accepted` performed by the human from a human-controlled
environment. Agents never merge into `accepted`, never push, and never
claim acceptance. Platform permission systems may additionally enforce
this; the human-controlled merge is the boundary, not any local hook.

## Platform adapters

- OpenCode: `.opencode/opencode.json` (permissions, model pins) plus
  `.opencode/agents/*.md` (thin pointers to `agents/*.md`). Canonical
  methodology lives in `agents/*.md`; adapter/model drift is resolved in
  favour of the canonical files.
- Codex: no repository adapter file; observed invocation `codex exec`
  (non-interactive) reading repo instructions directly. Behaviour notes in
  `docs/compatibility-matrix.md` after matrix testing.
- Muse: no repository adapter file; observed invocation `muse exec`
  (headless) reading repo instructions directly. Behaviour notes in
  `docs/compatibility-matrix.md` after matrix testing.

## Output isolation (multi-agent non-overwrite)

Each coding-agent process must be launched with its assigned worktree as
its working directory, and must not access another agent's worktree.
Checkouts never share files, so one agent cannot overwrite another's bytes
under normal operation. This is an operational safeguard, not filesystem
isolation: it does not technically prevent a misbehaving process with
broader filesystem access from reaching another checkout.

Canonical records are create-only: `physctl new` refuses existing
destinations (`record.writeNew`, O_EXCL; concurrent racers receive distinct
IDs), and `physctl paper fetch` refuses to replace cached bytes that differ
(explicit `collision` error naming both digests; identical re-fetch is
idempotent and leaves the entry untouched). Conflicting edits to the same
tracked file surface as a Git merge conflict preserving both sides — never
a silent winner.

Revisions to canonical records are proposed versions, never in-place
rewrites: `research/proposed/<agent-id>/<run-id>/<artifact-path>`,
mirroring the canonical relative path (for example
`research/proposed/codex/2026-11-05T01-30/hypotheses/H-002.r2.md`). Each
run uses a fresh run-id (timestamp or UUID); reusing a run-id fails
explicitly (`record.CreateRunDir`), so separate runs never share a
destination. Tooling ignores `research/proposed/` (it is not a scanned
record dir), so proposals cannot disturb validation; a human promotes a
proposal by review. `paper-cache/` entries are content-addressed by
(paper, version) and fenced by the fetch collision rule above;
`fulltext.txt` is a deterministic re-derivation of its cached PDF, not an
agent output.

Agent workspaces (branches `agent/muse`, `agent/codex`, `agent/opencode`;
sibling `phys-research-agents/<agent>` checkouts) start at the pre-freeze
HEAD. Newly created worktrees do NOT include uncommitted changes from the
main checkout: never assume an agent worktree contains the latest research
until the baseline below has been deliberately distributed to it.

### Baseline distribution (staged, never in place)

Rule: a new agent run receives a new, verified workspace. Existing agent
workspaces are never updated in place — distribution assembles a fresh
staging worktree, verifies it completely, and only then designates it as
an agent workspace. On any failure, every pre-existing agent worktree is
left untouched. The current agent worktrees are therefore ON HOLD for
baseline updates until a human reviews and designates fresh staged
workspaces for them.

Canonical source baseline, machine-local configuration and distribution
payload are three separate things:

- The historical freeze manifest `docs/BASELINE-2026-11-05.sha256` is
  preserved as the review record; it is not rewritten.
- Machine-local `go.work` is excluded from every portable manifest. Each
  fresh worktree recreates its own `go.work` per AGENTS.md §10 (`physctl
  check` reports UNAVAILABLE without it — never PASS).
- Each distribution carries its own payload manifest covering every
  tracked and untracked file actually distributed. Verification fails on
  missing, unexpected, changed or duplicated payload entries; files can
  never be silently omitted.

Procedure (human-executed, every step fails closed):

1. In the main checkout, record the export base and export the reviewed
   baseline: `git rev-parse HEAD > /tmp/dist.base-commit`
   `git diff HEAD > /tmp/dist-tracked.patch`
   `git ls-files --others --exclude-standard -z | tar --null -cf
   /tmp/dist-untracked.tar --files-from=-`
   `git ls-tree -r --name-only HEAD > /tmp/dist-base.list`
2. Validate the tarball before anything is extracted (rejects absolute
   paths, traversal, links, unsupported types, duplicates):
   `physctl dist check-tar /tmp/dist-untracked.tar`
   Then generate the payload manifest over exactly the distributed set
   (NUL-safe file lists; `go.work` is ignored and never listed):
   `(git diff --name-only -z HEAD; git ls-files --others
   --exclude-standard -z) | sort -zu | (cd <main> && xargs -0 sha256sum)
   > /tmp/dist-payload.sha256`
3. Create a new staging worktree at the recorded export base
   (a fresh, uniquely named directory — never an existing agent
   worktree) and apply the tracked patch there:
   `git apply --check` first; `git apply` refuses on conflict.
4. Extract the tarball into the staging worktree, then verify the
   assembly against the payload manifest from step 2:
   `physctl dist check-staged --dir <stage> --manifest
   /tmp/dist-payload.sha256 --base /tmp/dist-base.list --ignore go.work`
   `check-staged` fails on any missing, changed, unexpected or
   duplicated file.
5. Run `physctl validate` inside the staging worktree.
6. Only after every check succeeds, designate the staging worktree as
   the agent's new workspace (rename/repoint by human decision). On any
   failure, discard the staging area; pre-existing agent worktrees are
   byte-identical by construction because nothing ever wrote to them.

Tar note (verified on this platform): GNU tar 1.35 `--keep-old-files`
exits 2 on collision and preserves existing bytes — it does not silently
succeed. The staged procedure above does not rely on that behavior
either way: collisions in a fresh stage indicate a base mismatch and
abort the run.

### Security boundary (explicit)

- Go-enforced: create-only record/run-dir creation, fetch byte-identity
  collision rule, tarball entry validation, staged-tree verification.
- Git-enforced: worktree file separation under normal operation, merge
  conflicts preserving both sides.
- Instruction-only: launching each agent in its assigned worktree, not
  accessing other worktrees, writing revisions as proposed versions.
- Not a boundary: worktree separation is an operational safeguard, not
  OS-level isolation. No environment variable, working-directory
  instruction or Go helper prevents a process with unrestricted
  filesystem access from directly writing to another worktree. Genuine
  enforcement against arbitrary writes requires OS-level isolation,
  which is out of scope for this task.

Integration of agent work back into the main line remains a human merge
decision in all cases.

## Known coupling (human-owned)

`AGENTS.md` §§2–4 still frames the instruction hierarchy, roles and
acceptance boundary in OpenCode terms. Rewriting those sections in
platform-neutral language requires a human edit (agents may not modify
`AGENTS.md`) and is tracked as a remaining item, not done here.
