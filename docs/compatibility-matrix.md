# Compatibility Matrix: Coding-Agent Operation of phys-research

Same fixture, isolated copies. Fixture slice: read instructions, run
`physctl validate` + `status --branch C`, run `paper search --limit 3`,
create and validate one probe source record. Fixture prompt stored with
the run notes (not in this repo). Date: 2026-11-05.

## Core capabilities (mandatory per docs/AGENT_CONTRACT.md)

| Capability | Muse | Codex | OpenCode |
| ---------- | ---- | ----- | -------- |
| Read shared research instructions | VERIFIED | NOT VERIFIED | NOT VERIFIED |
| Execute physctl | VERIFIED | NOT VERIFIED | NOT VERIFIED |
| Create research artifacts | VERIFIED | NOT VERIFIED | NOT VERIFIED |
| Validate research records | VERIFIED | NOT VERIFIED | NOT VERIFIED |
| Preserve provenance | VERIFIED | NOT VERIFIED | NOT VERIFIED |

## Enhancement capabilities (optional)

| Capability | Muse | Codex | OpenCode |
| ---------- | ---- | ----- | -------- |
| Literature discovery (live search) | VERIFIED | NOT VERIFIED | NOT VERIFIED |
| Complete Researcher workflow | VERIFIED | NOT VERIFIED | NOT VERIFIED |

## Evidence

- Muse: fixture `/tmp/wt-muse` (isolated worktree + working-tree overlay,
  prebuilt `./bin/physctl`). Observed: `validate` exit 0; `status --branch C`
  exit 0 with 12 `?` lines; `paper search` exit 0 with 3 results; probe record
  P-008 created via `physctl new`, filled, final `validate` exit 0 (1
  pre-existing LEGACY_UNPINNED warning). Full Branch C program (source
  qualification, equations, D-001, H-002/H-003, dossier) was additionally
  executed by this Muse session in the real tree. No `muse exec`
  subprocess was used (`muse` CLI reports "not logged in" for subprocess
  use); the Muse column reflects the interactive Muse session.
- Codex: NOT VERIFIED — cross-platform runs discontinued per user
  instruction 2026-11-05 before any model call. No compatibility claimed.
- OpenCode: NOT VERIFIED — same; no compatibility claimed.

## Setup notes for a future attempt (observed, not verified end-to-end)

- Codex: needs `CODEX_HOME` redirected to a writable dir with a writable
  COPY of `auth.json` (a symlink breaks token refresh: "Read-only file
  system"); the observed ChatGPT token was expired (401 token_expired) and
  refresh could not persist. Invocation shape:
  `codex exec -C <fixture> -s workspace-write --skip-git-repo-check <prompt>`.
- OpenCode: under a filesystem sandbox, redirect `XDG_DATA_HOME`,
  `XDG_STATE_HOME` (server writes `service.json` under `~/.local/state`)
  and `XDG_CACHE_HOME`; symlink `auth.json`/`mcp-auth.json` read-only into
  the redirected data dir. Invocation shape:
  `opencode run --auto <prompt>` with cwd at the fixture.
- Muse subprocess: `muse exec` requires `muse login` (interactive); the
  interactive session path is verified instead.
