# AGENTS.md — theorems/bohm-equiv

Bohmian equivariance (Benchmark 2) evidence fixture. Audit:
`../../audit/` (docs 01–03; S1 freeze, B-register, continuity DERIVED,
local equivariance DERIVED, global SOURCE-REPORTED via TT). This area
is representation + regression only: it executes fragments of the
audit chain on explicit 1D test inputs and replays the evidence. It
proves nothing about the general R^{3N} theory.

PLACEMENT NOTE: plan7.1 prescribed `phys-lib/bohm/equivariance/`; a
direct user order ("do not modify any phys codebase") relocated the
artifact here as a standalone module on the qsg-t001 pattern. The
shared repos are read-only inputs, never targets.

## Layout

- `go.mod`, `go.work` — new module `github.com/PithomLabs/bohm-equiv`;
  colocated workspace over phys-artifact + phys-math (absolute
  `use` paths: this module lives outside `/home/chaschel/Documents/go`).
  No phys-lib import: the fixture uses only the representation layer.
- `fixture.go` — `Build` (four 1D test pairs, 30 one-step
  DIFFERENTIATE/SIMPLIFY derivations) and `(*Fixture).Verify`
  (replays + linkage gates). Gates compare resolved payload bytes
  or canonical forms, never bare refs.
- `manifest.json` — premises/assumptions/limitations manifest.
- `source.json`, `source_mapping.json` — TT import identity +
  structural equation-to-fixture linkage (linkage only, never
  semantic equivalence; enforced by test).
- `protocol/derivation_test.go` — executable suite (15 tests).

## Commands (GOCACHE is read-only; always override)

```
export GOWORK=$PWD/go.work GOCACHE=/tmp/bohm-equiv-gocache GOPROXY=off
go build ./...
go test -count=1 ./...
go vet ./...
gofmt -l .   # must print nothing
```

## What replay proves (and does not)

Proves: 30 single-step derivations replay; executed partials match
hand-built raw engine forms (power rule decrements: d_t(q^2) is raw
Mul(0,2,q^1)); solution-pair residual is engine-zero; non-solution
residual separates from zero and from the solution residual;
product-rule relation links the stored executed divergence and both
sides substitute totally at (t,q)=(1,1); traveling-pair substitution
is consistent, deterministic, and nontrivial; two engine-built wrong
paths separate from the corrected residual and from zero; guidance
verdicts require their NEQ(rho,0) conditions; flow/pushforward/
equivariance relations appear in zero derivation footprints; only
SIMPLIFY-001 and DIFFERENTIATE-001 appear; rebuilds are byte-identical.

Does NOT prove: the product-rule EQ agreement or the pair-4
cancellation — both rest on adjacent hand algebra because
SIMPLIFY-001 cannot decide rational-tree equality (no
distributivity, like-term collection, or const-folding; recorded as
limitation `equality_not_decided`, demonstrated by the
Mul(t,...)-vs-Mul(q,t,...) normal-form split and the unfolded
substituted trees). Does not prove flow existence, the transport
identity, uniqueness, node behavior, (E-loc), or (E-glob) — those
live in the audit record only, as derived, sourced, or excluded.

## Golden policy

Every golden cites the hand derivation that produced it (adjacent
protocol comments): raw engine normal forms, pair algebras
(R2: 4tR(1+q); R4: -12Ru/+12Ru at u=q-3t), and the audit sections
(02:C, 02:j-def, 03:guide, 03:T, 03:E-loc, 03:flow, 01:S1).
Changing a golden requires changing the cited audit section first.

## No-go rules

- No imports beyond phys-artifact + phys-math.
- No writes outside this directory and `../../audit/`: shared repos
  are read-only; the committed `/home/chaschel/Documents/go/go.work`
  is untouched; nothing is committed anywhere.
- No test may assert what the engine cannot prove: a failed machine
  attempt becomes a recorded negative result, never a weakened gate.
- No weakening a failing gate: fix the fixture/test or record why
  the expectation is wrong in the manifest limitations.
- No relativistic/QFT/typicality/measurement content (doc 01 §4).
