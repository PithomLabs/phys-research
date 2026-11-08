# G1 closure report — root Go workspace

Date: 2026-11-06. Task: close gap G1 (`06-infrastructure-gaps.md`) from
QSG Preflight 001 Agent A recon. Environment note: default `GOCACHE` is
read-only in this session, so all Go commands ran with
`GOCACHE=/tmp/g1-closure/gocache`; `GOWORK` set explicitly per command
below. Creating `/home/chaschel/Documents/go/go.work` required a
one-time sandbox escalation (approved); content is exactly as specified.

## 1. Change applied (only change)

Created `/home/chaschel/Documents/go/go.work` (80 bytes). No prior
`go.work*` existed at that root (verified before writing — nothing
overwritten). No `go.mod` touched. No `go.work.sum` created (none
needed: zero external dependencies across the five modules).

```text
go 1.25.7

use (
	./phys
	./phys-gr
	./phys-artifact
	./phys-math
	./phys-lib
)
```

`go work sync` was NOT run: unjustified — it rewrites `go.mod` files
(forbidden by constraints) and exists to sync external-module hashes,
of which there are none.

## 2. Modules and dependency resolution

| Module | Requirement | Resolution under workspace |
|---|---|---|
| phys | none | self-contained |
| phys-gr | `phys v0.0.0` (+ `replace => ../phys`) | local (replace also works standalone) |
| phys-artifact | none | self-contained |
| phys-math | imports `phys-artifact`, no require | workspace `use` (previously unbuildable) |
| phys-lib | `phys-math v0.0.0`, no replace (M-4 deferred) | workspace `use` (previously unbuildable) |

External packages introduced: none. Dependency versions changed: none.

## 3. Build results (`GOWORK=<root>/go.work`, per-module `go build ./...`)

| Module | Exit |
|---|---|
| phys | 0 |
| phys-gr | 0 |
| phys-artifact | 0 |
| phys-math | 0 |
| phys-lib | 0 |

## 4. Focused regression tests (`go test -count=1`)

| Suite | Result |
|---|---|
| `phys-gr/solution` | ok (4.6s) |
| `phys-lib/gr/equivariance` + `protocol` | ok |
| `phys-lib/bohm/minisuperspace` + `protocol` | ok (3.4s protocol) |
| `phys-lib/gr/ltb` + `protocol` | ok |
| `phys-artifact` | ok (0.6s) |
| `phys-math` shard C (`-skip 'TestObligation\|TestDivergence\|TestWitness'`, `-timeout 20m`) | ok (7.8s) |
| `phys-math` registry/purity/import gates | ok |

`go vet`: exit 0 on phys-math (`./...`), phys-lib (`gr/... bohm/...
internal/...`), phys-gr, phys-artifact. (`phys` vet not run: frozen
repo, out of scope; no change there to validate.)
`gofmt -l` over all five module trees: clean except 5 files under
frozen `phys/` (`core/corpus.go`, `core/expr_test.go`,
`core/object_test.go`, `hypothesis/candidate.go`,
`mechanics/primitives.go`) — pre-existing formatting drift in a frozen
repo, verified untouched (git clean), NOT fixed per constraints.

## 5. Standalone behavior unchanged (`GOWORK=off`)

| Module | Standalone result | Matches pre-G1 state |
|---|---|---|
| phys-artifact | exit 0 | yes |
| phys-gr | exit 0 (via its replace) | yes |
| phys-math | exit 1, same `no required module provides package` error | yes |
| phys-lib | exit 1, `missing go.sum entry` for phys-math | yes |

Workspace file only *adds* a resolution path; module-mode behavior is
byte-identical to before (no `go.mod`/`go.sum` modified). Existing
repository instructions untouched (no `AGENTS.md` edited anywhere).

## 6. Non-modification confirmation

- `git status --porcelain`: clean in phys, phys-gr, phys-math,
  phys-artifact, phys-trace, phys-pseudo, phys-system; phys-lib shows
  only the pre-existing untracked `.kilo/` (present before this task,
  preserved).
- BHC artifacts: sha256 re-checked — `239a24c7…` (core),
  `e7ad3a77…` (No-Go PDF), `7e8e86d9…` (Preflight PDF) — all match.
- Protected sets (PhysMath/PhysLib/PhysArtifact/Phys-Pseudo sources,
  GR fixtures, BHC artifacts): zero modifications.
- No Git initialized, no workspace restructured, no untracked files
  removed.

## 7. Remaining blockers

None for G1: acceptance met (5/5 builds, focused regressions green, no
protected files modified). G1 is CLOSED. Open items belonging to later
work, not this task: G2/G3 per-claim scoping (awaits Agent B physics
input); pre-existing `gofmt` drift in frozen `phys/` (fixing it would
violate the freeze — left as found).
