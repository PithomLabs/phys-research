## Goal

Execute `plans/prompt4.md` (Agent A — QSG-T001 resumption): finish the independent conservation audit, pass the theorem acceptance gate, and produce the minimal formalization of QSG-T001 — covariant conservation of the Hadamard-renormalized stress-energy for a free, real, minimally coupled, massive scalar field — without repeating completed work.

## Success Criteria

- All seven audit documents exist under `research/experiments/qsg-preflight-001/agents/agent-c-conservation-audit/`, each read back, using only the six dispositions (DERIVED, SOURCE-REPORTED, INDEPENDENTLY REPRODUCED, CONDITIONAL, FALSIFIED, BLOCKED; never PASS).
- The conservation-restoring local correction is derived with explicit coefficient and sign under one consistent convention, and the corrected tensor's divergence is explicitly shown to vanish (or a BLOCKED obstruction is precisely identified).
- Agent B's defects are documented with exact PDF page/equation locations and dispositions.
- The strongest defensible QSG-T001 proposition is stated with explicit assumptions, exclusions, and ambiguities; the Phys-Pseudo representation in `06-theorem-contract.md` uses only the seven permitted constructs.
- Any Go fixture follows the GR reference architecture; build/test/vet/gofmt results are recorded; unsupported analytical steps are never represented as executed proofs.
- Agent B's PDF, completed Agent A reports, and existing audit docs `01`/`02` are untouched.

## Approach

Resume, don't restart. Verified this session: audit docs `01` (532 lines, complete through §8) and `02` (538 lines, complete through §9) already exist and are conclusive — `01` fixes Décanini–Folacci as the load-bearing source (Wald full texts BLOCKED, abstracts verified) and `02` reproduces the classical foundation. Agent A recon (8 docs + G1 closure) is complete, and G1 is closed (`go.work` exists at `/home/chaschel/Documents/go/go.work`). The missing work is docs `03`–`07`, the acceptance gate, and conditional minimal formalization. No broad survey, no PDF re-extraction beyond single targeted inspections, no code before the gate.

Key decisions:

- **Trust boundary for `01`/`02`:** accept as done after readback verification; do not rewrite. Rationale: both end with explicit conclusions and disposition tables; prompt4 Rule 1 forbids repeated work.
- **Load-bearing source:** Décanini–Folacci (2008) full text, already equation-verified in `01`; Wald 1977/1978 stay SOURCE-REPORTED-at-abstract-level/BLOCKED for details. Rationale: paywalled full texts cannot be independently reproduced; DF Eq. (70) already covers the needed coefficient.
- **Convention freeze before calculation:** DF conventions (signature −+++, □ = g^μν∇_μ∇_ν, units ħ=c=1) plus explicit Wightman/symmetric two-point-function normalization, frozen in writing at the start of doc `03` work. Rationale: prompt4 Rule 5; mixing conventions is the failure mode behind the g·v₁ coefficient dispute.
- **Formalization location:** default to the standalone area `research/experiments/qsg-preflight-001/theorems/qsg-t001/` per the implementation contract (`07-implementation-contract.md` §1–3), not a new directory under shared `phys-lib/gr/`. Rationale: prompt4 §7 asks for "the smallest appropriate new location under `phys-lib/gr/`", but the binding recon contract (§5) forbids theory-specific meaning in shared modules and mandates the standalone derivation area; writing into the shared repo also needs explicit authorization. Flagged as the one prompt-vs-contract conflict; standalone wins unless you direct otherwise.
- **Quantum representation bound:** c-number results as pinned goldens + per-instance evaluated obligations (Bianchi pattern from `phys-math/obligation.go`); operator algebra, state semantics, and renormalization procedure stay human obligations via `ProtocolSketch` UNSUPPORTED. Rationale: feasibility matrix (`04-qsg-mathematical-feasibility.md`) — procedure is NO, result-pinning is YES.

## Steps

1. **State inspection and convention freeze.** Read back `01`/`02` fully; confirm audit-dir contents; open Agent B's PDF only at pages needed for `03`/`04` (pp. 10–13 first, once each); write the frozen convention record (two-point normalization, parametrix, operator, curvature conventions) as the preamble to doc `03`.
2. **Doc `03`: Hadamard / point-splitting audit.** Derive (□−m²)H explicitly; establish the actual smooth source term for W = G⁺−H (falsify or correct Agent B's (□−m²)W = 0); audit coincidence-limit derivative handling (total vs single-point differentiation); pin subtraction prescription, bi-differential operator, parallel transport, coincidence limits. Disposition each claim.
3. **Doc `04`: conservation and counterterm derivation (critical path).** Compute the point-split divergence; identify the residual local geometric term; derive the correction (candidate ∼ g_μνv₁ — reconcile coefficient/sign against DF's general-dimensional formula under the frozen conventions); compute its divergence; show explicit cancellation; verify finite-ambiguity basis (Tichy–Flanagan vs Salcedo conditionality per `01` §8) and Agent B's H⁽¹⁾/H⁽²⁾; keep trace anomaly separate with the corrected formula.
4. **Doc `05`: Agent B findings.** Build the page-anchored defect table for prompt4 §4 locations (pp. 1–3, 10–11, 11, 12–13, 13 ×3, 14–16, 16–18, 19–20), each with: page/equation, assertion, correct expression or reason unestablishable, missing step, minimum correction, disposition. No silent repairs.
5. **Doc `06`: theorem contract + gate.** State the strongest defensible proposition with assumptions, exclusions, ambiguities, proof obligations; answer the 10 acceptance-gate questions explicitly; include the concise Phys-Pseudo representation (given/define/assume/claim/derive/if-then-else/? only, per `02-phys-pseudo-contract.md` binding rules).
6. **Doc `07`: independent handoff.** Final verdict: was Agent B's conclusion/derivation correct, what correction is necessary, what was reproduced, what remains, readiness for statement-level formalization vs executable verification, smallest remaining task. If any gate question fails, report BLOCKED with the exact obstruction and stop before Step 7.
7. **Minimal formalization (only if gate passes).** Scaffold `theorems/qsg-t001/` per the implementation contract (AGENTS.md, manifest.json, research_metadata.json, fixture.go, protocol/derivation_test.go); represent established results, ProtocolSketch UNSUPPORTED for analytical steps; per-instance witnesses + goldens/anti-goldens + mutation negatives; run build/test/vet/gofmt under the workspace recipe (`GOWORK` + writable `GOCACHE`); record actual results; second-checkout `phystrace-mcp run_verification` if available.
8. **Final handoff report.** Answer prompt4 §11's ten questions concisely in chat.

Non-goals: rewriting `01`/`02`, recon docs, or Agent B's survey; backreaction, FLRW, stochastic gravity, noise kernel, Bohmian mechanics; any engine growth, parser/validator, new dependencies, or shared-repo edits.

## Validation Plan

- Per doc (`03`–`07`): full readback after writing; disposition-vocabulary check (grep for PASS/status words — must be absent); every equation cites frozen convention + source equation or derivation step.
- Gate check: the 10 questions in `06` each answered DERIVED/CONDITIONAL/BLOCKED with pointers; no theorem claim stronger than its evidence.
- Phys-Pseudo check: manual conformance pass against the seven constructs and binding rules (names declared before use, non-empty `from`, `?` only for unresolved, known-false via refuting `derive`); no status/provenance constructs.
- Fixture checks (Step 7 only): `go build ./...`, `go test -count=1 ./...` (fixture + protocol), `go vet`, `gofmt -l` under the workspace recipe — all exit 0/clean with results recorded; manifest strict-decode + unknown-field rejection tests; every golden cites its hand derivation; every mutation class fails. Highest-risk validation: the Step 3 cancellation calculation itself — a plausible final equation is not evidence; the explicit divergence-of-correction computation is the gate's load-bearing check.

## Risks / Open Questions

- **Wald full texts BLOCKED (paywalled).** Mitigation: already contained — DF is the load-bearing source; Wald claims stay SOURCE-REPORTED at abstract level. No action unless DF proves insufficient for a needed identity.
- **Correction coefficient/sign reconciliation.** The DF general-dimensional formula must be rechecked against the frozen two-point normalization; a mismatch here is the most likely BLOCKED outcome. Per Rule 10, report it early rather than extending the investigation.
- **Prompt-vs-contract location conflict** (`phys-lib/gr/` vs standalone `theorems/qsg-t001/`). Default: standalone. Open question only if you prefer the shared-repo location — that needs your explicit authorization.
- **Session-file naming:** prompt4 cites `session-c4a1a27d-...(...)1).md`; on disk it is `plans/session-c4a1a27d-d96d-4646-80c2-734cec802d77.md` (2.8 MB). Treated as the same record; non-blocking.
- External research: none required for this plan — no new libraries, engines, or APIs are recommended; all approach decisions trace to inspected workspace evidence (recon docs, `01`/`02`, prompt3/review3, Agent B PDF presence).
