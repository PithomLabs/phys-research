# 07 — Implementation contract for QSG-T001

Minimal engineering requirements for the first QSG theorem. Reuses the
GR execution protocol wherever it applies; assigns everything else to
its correct §8 category. No architectural changes required to begin.

## 1. Prerequisites (before any derivation work)

1. Commit root workspace `/home/chaschel/Documents/go/go.work` (gap G1):
   `go 1.25.7` + `use` over `phys-artifact`, `phys-math`, `phys-lib`,
   `phys`, `phys-gr`. Until then, use the verified recipe:
   `GOWORK=<file> GOCACHE=<writable> go build|test` (`01 §2`).
2. Create the QSG-T001 derivation area (standalone folder; no git init):
   `research/experiments/qsg-preflight-001/theorems/qsg-t001/`.
3. Freeze the physics input from Agent B: matter action, state class,
   renormalization prescription, regime, target proposition — as
   pseudo + sidecar claims (RED, `05`). No derivation starts without it.

## 2. Claim assignment (the §8 trichotomy, per QSG area)

| Area (§7) | Mathematical proof (human) | Executable verification (machine) | Research provenance (artifacts) |
|---|---|---|---|
| A. Matter foundations | EL derivation from action; metric-variation identity (general) | Chart-bound curvature/differential witnesses; pinned classical identities as goldens | action/params as `expression/1`+`function/1`; assumptions as `assumption/1`; pseudo `assume/derive/?` |
| B. Quantum formulation | Operator identities; state/Hadamard analysis | c-number consequences (kernels, mode identities) as replayed derivations + goldens | state/regularity as assumptions + pseudo; `ProtocolSketch` UNSUPPORTED for unexecuted steps |
| C. Renormalization | Prescription justification; subtraction derivation; anomaly analysis | Pinned renormalized-result goldens + anti-goldens; curvature-invariant inputs computed | prescription as assumption; result as `equation/1`; hand-derivation cited per golden |
| D. Conservation | Universal ∇⟨T⟩=0 identity | Per-instance evaluated obligations (copy `obligation.go` Bianchi pattern: instance witnesses + replayed closing derivations) | obligation records + sidecar claim status |
| E. Backreaction | Existence/stability analysis | Coupled-system instance checks; regression suite | `equation_set/1`; limitations recorded |

Rule: any claim in column 2 without a column-3 witness is `open-conjecture`
in the sidecar and MUST NOT appear in `expected_result`.

## 3. Required derivation package (per theorem)

```
qsg-t001/
  AGENTS.md                  # local contract (copy equivariance shape)
  manifest.json              # 16-key GR schema; strict decode; unknown fields rejected
  research_metadata.json     # claims[] + obligations[] (ltb shape)
  fixture.go                 # Build mints into physmath.Store
  protocol/derivation_test.go# trace + goldens/anti-goldens + mutations + TestManifest*
```

## 4. Commands and acceptance gates

```bash
GOWORK=<root/go.work> GOCACHE=<writable> go build ./...          # exit 0
GOWORK=<root/go.work> GOCACHE=<writable> go test -count=1 ./...  # fixture + protocol, exit 0
go vet ./... ; gofmt -l .                                        # clean
```

Acceptance (all must hold): manifest validates against
`manifestschema`; every derivation replays (`Verify`); every golden has
an independent hand derivation cited; every mutation class fails;
`research_metadata.json` has no claim in a stronger status than its
evidence; limitations list is non-vacuous (states what was NOT shown);
`phystrace-mcp run_verification` (coverage-audit or full-suite) passes
in a second checkout.

## 5. Constraints (must NOT)

- No new external dependencies (stdlib-only derivation path preserved).
- No engine growth (families/rules) without a bounded, approved
  proposal — and only after a concrete derivation demonstrates the need.
- No `expected_result` for unexecuted claims; no promotion by
  computation; no prose-as-math (gated by tests).
- No modification of frozen repos (`phys`, `phys-gr`), shared modules,
  manifests, or BHC artifacts. Theory-specific meaning lives in the new
  QSG derivation area (and, if needed, a future framework package — not
  in `phys-math` or `gr`).

## 6. Entry criterion (single sentence)

QSG-T001 formalization may begin when Agent B's frozen physics input
(§1.3) exists and G1 is closed; everything else in this contract is
then executable with existing infrastructure.
