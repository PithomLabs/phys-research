## Goal

Implement `plan7.md`: establish Bohmian equivariance as Benchmark 2 (second physics vertical slice), reusing the shared platform without polluting PhysMath, then evaluate what the two benchmarks jointly prove about cross-theory generality. No new QSG-T001 review cycle.

## Success Criteria

- A bounded equivariance derivation exists from stated premises (wavefunction → density/current → continuity → guidance → flow → equivariance), with an assumptions register, local-vs-global separation, Math Reality Check regime/scope, and per-step dispositions (derived / source-reported / conditional / unsupported).
- An executable artifact exists following minisuperspace + qsg-t001 conventions (manifest, protocol tests, honest supported-subset tests, no conclusion-imported goldens), with Bohmian commitments inside `phys-lib/bohm` and PhysMath theory-neutral.
- The five platform-evaluation questions are answered with evidence, plus a QSG-vs-Bohm representation comparison and a prioritized demonstrated-gap list. No preemptive refactoring.
- Benchmark 1's status inconsistency from `plan6_imp_review.md` §4 is resolved as bookkeeping (O1 verification-level split + provisional wording), with zero QSG re-derivation.
- Shared-repo diff audited and left uncommitted; frozen repos (`phys`, `phys-gr`) untouched.

## Approach

Three discoveries shape this plan. First, `phys-lib/bohm/doc.go` already owns exactly this remit ("Bohmian-mechanics framework semantics… mathematical execution stays in phys-math"), and `phys-lib/bohm/minisuperspace/` is a prior in-repo artifact (fixture, manifest, protocol, tests) — so placement and conventions are precedented, not designed fresh. Second, `plan7.md` (01:50) is newer than `plan6_imp_review.md` (01:46) and explicitly forbids another QSG review cycle — so the review is dispositioned as bounded status bookkeeping, and its narrow-audit evidence becomes a packaged reading list for future external replication, not new work. Third, unlike QSG (TOE-v6 only), this mission writes to the shared repo — so the plan carries a git-diff audit and no-commit boundary throughout.

## Steps

1. **Baseline + status alignment (bookkeeping only, no re-derivation).** Run the QSG v2 gate as a health check; record the O1 four-level split (in-execution derivation / internal cross-check closed; external replication / separate-session reproduction open) and the provisional-acceptance wording in the QSG sidecar/manifest/06-O1; record the J/J′ non-exchange rule as a *proposed* permanent methodology note; package the narrow-audit reading list (08 N2–N4, doc 04 J/J′, doc 06, v2 tests).
2. **Read-first recon.** `phys-system/AGENTS.md`, `phys-math/AGENTS.md`, phys-lib README + PROTOCOL_GAPS, `bohm/doc.go` + full minisuperspace (check for existing continuity/flow/equivariance coverage to avoid duplication); map reusable PhysMath abstractions vs genuine gaps; confirm placement (recommend: Go in `phys-lib/bohm/equivariance/` per minisuperspace precedent, math record in new `research/experiments/bohm-preflight-001/` mirroring qsg-preflight-001); record shared-repo `git status` baseline.
3. **Math derivation.** Nonrelativistic finite-particle Schrödinger chain per mission scope; assumptions register (regularity, node singularities of `v=j/ρ`, boundary/decay, flow existence); local-in-time first, global only under explicit existence premises (expect CONDITIONAL/unsupported — honest, not forced); MRC regime/scope statement reusing the BHC MRC convention; dispositions per result. Stop conditions enforced: no QFT, relativistic, measurement, or interpretation-comparison content (BHC R3 typicality/Born routes explicitly out).
4. **Implementation.** Theory commitments in `phys-lib/bohm`; no new PhysMath abstractions unless the math requires (deficiency-driven only); manifest + `protocol/derivation_test.go` + sketch pattern mirroring qsg-t001 v2; tests verify only the supported subset; derivation conclusion never a golden input.
5. **Platform evaluation + handoff.** Answer the five mission questions with file-level evidence; write the QSG-vs-Bohm comparison and prioritized gap list; full gates (build/test/vet/gofmt both scopes, JSON validity); shared-repo diff audit, everything left uncommitted; TOE-v6 hash manifest; report.

## Validation Plan

- QSG v2 gate green before starting (baseline health): `go build/test/vet`, `gofmt -l` empty in `theorems/qsg-t001/` with the colocated `GOWORK/GOCACHE/GOPROXY=off` recipe.
- New package + artifact: `go build ./...`, `go test -count=1 ./...`, `go vet ./...`, `gofmt -l` empty; manifest/sidecar strict-decode + unknown-field-injection tests; snapshot/determinism tests; sketch exposes no Verify/Replay (mirror v2 suite properties).
- Highest-risk check (manual, non-automatable): audit that every proof obligation is either derived-from-premises or explicitly marked conditional/unsupported, and that no test or golden smuggles the conclusion — the same honesty audit that caught the QSG defects.
- Boundary checks: `git status`/diff in `/home/chaschel/Documents/go` shows only intended `phys-lib/bohm` additions; frozen `phys`, `phys-gr`, committed root `go.work` untouched; nothing committed.

## Risks / Open Questions

- Shared-repo writes are new for this programme (QSG wrote only TOE-v6): mitigated by baseline + final diff audit and the no-commit rule.
- Global Bohmian flow existence is genuinely hard (nodes, measure-theoretic premises): treated as expected-CONDITIONAL scope, not a failure; local result is the deliverable floor.
- No open questions blocking planning: artifact naming (`bohm-preflight-001`) and package layout are reversible execution details; the mission is otherwise prescriptive.
