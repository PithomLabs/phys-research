# AGENTS.md — theorems/qsg-t001

Minimal QSG-T001 theorem fixture. Audit: `../agents/agent-c-conservation-audit/`
(docs 01–08; repair record doc 08 N1–N5). Theorem: revised doc 06 (T).
This area is representation + regression only: it pins the audit's
c-number results and replays their rational-arithmetic skeleton. It
proves nothing. v1 goldens (3/(16pi^2) family) are SUPERSEDED; v2
goldens below are recalculated per doc 08.

## Layout

- `go.mod`, `go.work` — new module `github.com/PithomLabs/qsg-t001`;
  colocated workspace over phys-artifact + phys-math (absolute
  `use` paths: this module lives outside `/home/chaschel/Documents/go`).
  No phys-lib import: GR-closed framework categories must not carry
  QFT content.
- `fixture.go` — `Build` (mints goldens, five one-step SIMPLIFY
  derivations, six-step sketch) and `Verify` (replays + payload-byte
  gates). All gates compare resolved payload bytes, never bare refs.
- `manifest.json` — 16-key package manifest (equivariance shape, QSG
  content, locally validated; the GR manifestschema gate does not
  apply to QFT content).
- `research_metadata.json` — claims/obligations sidecar (ltb shape).
  `math_status` vs `machine_status` are separate fields: machine
  execution never upgrades a mathematical disposition.
- `protocol/derivation_test.go` — executable suite (this file reads the
  JSON sidecars; fixture.go does not).

## Commands (GOCACHE is read-only; always override)

```
export GOWORK=$PWD/go.work GOCACHE=/tmp/qsg-preflight/gocache GOPROXY=off
go build ./...
go test -count=1 ./...
go vet ./...
gofmt -l .   # must print nothing
```

## What replay proves (and does not)

Proves: five rational identities (2+4=6; (1/3)*6=2; -2+2=0;
-3/4+1=+1/4; (1/8)*(-3/4)=-3/32), one canon-inequality (-3/32 != 0),
fixed-snapshot match of nine golden payloads, separation of three
anti-goldens, strict JSON decoding with unknown-field rejection, and
all-unsupported sketch status. These are the audit's arithmetic
skeleton plus regression pins.

Does NOT prove: the remainder limit, the corrected TI, the Q1-split,
the residual, the correction, the cancellation, or (T) — those are
human proof obligations SK-1..SK-6 (see `BuildSketch`), each pointing
at its doc-08 section. A snapshot trip proves a stored value changed,
never that a physical coefficient is correct; replay checks
derivation-codec consistency, never the proof. Symbolic factors (pi
powers, m^4, curvature polynomials) are documented, never computed:
e.g. CANCEL checks -2+2=0 given the shared symbolic denominator
8pi^2; THIRD checks (1/3)*6=2 per the (N4.4) split.

## Golden policy

Every golden cites the hand derivation that produced it (fixture.go
comments): (1.7), (N4.5)+(N4.6), (N4.4), (N4.8), (N2.5), (V0), (3.1),
(3.2), doc-05 G4. Anti-goldens cite the falsified Agent B formula (C5,
D2, D3) and the rendered page inspected. Changing a golden requires
changing the cited audit section first; the fixed-snapshot test fails
otherwise (regression tripwire, not a physics check).

## No-go rules

- No imports beyond phys-artifact + phys-math.
- No universal claim in `expected_result` (enforced by test).
- No `Verify`/`Replay` method on `QSGSketch`, ever (enforced by test).
- No edits outside this directory (shared repos read-only; the
  committed `/home/chaschel/Documents/go/go.work` is untouched).
- No weakening a failing gate: fix the fixture or record why the
  expectation is wrong in the manifest anomalies.
