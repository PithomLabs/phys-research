# 07 — Independent handoff

Audit: QSG-T001. Author: Agent A (resumption). Date: 2026-11-06.

Final mathematical verdict, implementation-readiness decision, and answers
to the ten handoff questions. This document closes the audit; no fixture
exists yet (Steps 9–10 follow if authorized).

Dispositions follow `01`. `PASS` is not used.

---

## 1. Verdict

Agent B's conservation CONCLUSION is true under the corrected
prescription; Agent B's DERIVATION is invalid at every load-bearing step.
The independent blind derivation establishes:

- Remainder: \((\Box_x-m^2)S=J\), \([J]=-3[v_1]/(8\pi^2)\)
  (doc 03 (1.8)). DERIVED.
- Residual: \(\nabla^\mu R_{\mu\nu}=-3\nabla_\nu[v_1]/(16\pi^2)\)
  (doc 04 (1.3)). DERIVED, cross-checked.
- Correction: \(C_{\mu\nu}=+3g_{\mu\nu}[v_1]/(16\pi^2)\), explicitly
  not independently conserved (doc 04 (2.1)–(2.2)). DERIVED, cross-checked.
- Conservation: \(\nabla^\mu(R+C+\Theta)_{\mu\nu}=0\) (doc 04 (2.3),
  doc 06 (T)). DERIVED, cross-checked.
- Ambiguities: polynomial basis (3.3), CONDITIONAL on A6.
- Trace: (3.4 corrected by E2), nonzero, compatible with (T). DERIVED.

Agent B's specific errors (doc 05): singular \(W\) presented as smooth
(C2); false homogeneous equation (E4); dropped residual source (E5);
spurious transport mechanism (E6); contradictory "conserved" correction
(E7) with no formula; wrong \([V_0]\), \([V_1]\), \(a_2\) (C5–C7); wrong
\(H^{(1)}\), \(H^{(2)}\) (D2–D3); dimensional errors (D1); wrong trace
formula (F1); wrong Friedmann coefficients (G4); non-conforming
Phys-Pseudo (I1–I9); unsupported verification labels (A3, H1–H8, J1).

## 2. Readiness decisions

- Statement-level formalization: READY. Theorem (T) precisely stated;
  assumptions A1–A8 explicit; exclusions listed; obligations O1–O5
  registered; Phys-Pseudo representation conforming (doc 06 §6).
- Executable mathematical verification: READY FOR BOUNDED CHECKS ONLY.
  Per the feasibility study, the engine supports representation, pinned
  c-number results, per-instance evaluated obligations, and provenance —
  not universal identities, operator algebra, state semantics, or
  renormalization procedure. The fixture may pin (T)'s c-number content
  (coefficients, FLRW reductions, ambiguity basis) as goldens with
  hand derivations cited, and record the universal proof as
  ProtocolSketch UNSUPPORTED steps. It must never present replay as
  proof of (T).
- Verdict on the programme question: QSG-T001 is a genuine,
  independently established theorem at statement level with bounded
  executable regression support — not a machine-verified proof.

## 3. Handoff answers (prompt4 §11)

1. Was Agent B's conservation conclusion correct? YES — under the
   corrected prescription (2.2), \(\nabla^\mu\langle T_{\mu\nu}\rangle_
   {\rm ren}=0\) holds.
2. Was Agent B's derivation mathematically correct? NO — invalid at
   C2, E4–E7 (load-bearing), plus C5–C7, D1–D3, F1, G4.
3. What correction is necessary? \(C_{\mu\nu}=+3g_{\mu\nu}[v_1]/
   (16\pi^2)\) in §0 conventions (doc 04 (2.1)); DF-coefficient
   cross-match CONDITIONAL on \(A\) (O2).
4. Strongest theorem responsibly statable? Doc 06 (T) with A1–A8,
   exclusions §3, obligations O1–O5.
5. Which results independently reproduced? Classical core (`02`);
   \(T\)-operator, \([U]\), \([v_0]\), flat \([v_1]\); DF structure +
   equation transcriptions; H&W/Wald-abstract structure;
   coexistence statement; scale-absorption structure. (Doc 05 §§B–D,
   F, K; doc 06 §7.)
6. Which assumptions and obligations remain? A1–A8 (A6 conditional);
   O1–O5 (none blocking; O4 blocked-by-paywall but unneeded).
7. Ready for statement-level formalization? YES (§2).
8. Ready for executable mathematical verification? BOUNDED ONLY (§2).
9. What was actually implemented? NOTHING yet — audit complete (docs
   01–07), no fixture. Implementation is Steps 9–10.
10. Smallest remaining task? Packaging preflight (Step 9): determine
    the Go module arrangement for `theorems/qsg-t001/` (expected: new
    local module + disposable test workspace, zero committed changes),
    then the minimal fixture + validation (Step 10).

## 4. Guidance for Steps 9–10 (binding)

- Location: standalone `research/experiments/qsg-preflight-001/theorems/
  qsg-t001/` (doc 05 provenance + recon contract forbid shared-repo
  writes). No committed file may change without explicit approval.
- Content: manifest + research_metadata sidecar + fixture
  (Build/Verify) + protocol tests per the implementation contract;
  goldens for (1.8)-coefficient, (1.3)/(2.1)-coefficients, (3.1)/(3.2)
  (with hand derivations cited), FLRW \(H^{(1)}_{00}\) reduction,
  (3.4-trace) structure; anti-goldens for Agent B's falsified formulas
  (B-sign \(H^{(1)}\), homogeneous-\(W\), \([V_0]=R/12\));
  ProtocolSketch UNSUPPORTED for (T), (D1), TI, recursion, subtraction.
- The fixture asserts structural/regression facts only. Its limitation
  record must state: universal conservation remains a human proof
  obligation; replay is not proof.
- Validation: per-command recorded build/test/vet/gofmt-empty;
  second-checkout PhysTrace or recorded BLOCKED; before/after hashes.

## 5. Audit close

Seven documents complete: 01 (sources), 02 (classical), 03 (Hadamard +
Lock 1 + E1), 04 (conservation + Lock 2 + E2), 05 (findings +
dependency recheck), 06 (contract + gate PASS + Phys-Pseudo +
cross-check record), 07 (this verdict). Locks 1–2 hash-verified after
all appends. Protected files untouched (verified in Step 10
after-hashes). The mathematics is defensible; the implementation may
proceed minimally.

---

## 6. Steps 9–10 execution record (2026-11-06)

Preflight (Step 9): new module `github.com/PithomLabs/qsg-t001` in the
assigned area + colocated `go.work` (absolute `use` paths; different
filesystem root forces absolute) over phys-artifact + phys-math only
(no phys-lib: GR-closed categories unfit for QFT content). Hello-world
probe (family registry + store + artifact refs) green. Zero
committed-file changes; no approval needed.

Fixture (Step 10): `theorems/qsg-t001/` — `fixture.go` (Build/Verify:
5 replayed SIMPLIFY derivations, 4 shape goldens, 3 anti-golden
separations, nonzero gate, 6-step unsupported sketch),
`manifest.json`, `research_metadata.json`, `AGENTS.md`,
`protocol/derivation_test.go` (7 tests). Observed exits: `go build`
0, `go test -count=1` 7/7 PASS, `go vet` 0, `gofmt -l` empty
(GOWORK/GOCACHE/GOPROXY=off recipe; default GOCACHE read-only).

Second checkout: BLOCKED. phystrace-mcp builds and runs but is a
case-scoped review harness (`run_verification` serves demo-arc /
coverage-audit / full-suite profiles over registered demo workspaces
only); no selector or profile addresses a standalone theorem fixture.
Local acceptance: full Go gate green (observed). Independent
acceptance: outstanding.

Non-modification audit: before/after SHA-256 identical for Agent B's
PDF, all eight recon reports + G1 closure, and docs 01/02; all eight
shared repos byte-identical in git status (clean except pre-existing
untracked `phys-lib/.kilo/`). Locks 1–2 re-verified after all appends.

---

## 7. Normalization-repair superseding verdict (2026-11-06)

This section supersedes §§1–3/answer-numbers above (original lines
intact per append-only discipline). Adversarial review found the
v1 remainder singular; the repair (doc 08 N1–N5, E3/E4/E5,
revised-06) froze Alternative A, re-derived the remainder
(\([J]=-3[v_1]/(4\pi^2)\)), corrected TI (J/J'-distinction;
residual \([\nabla'J]=\tfrac13\nabla[J]\) via two routes),
derived v1-transport + F2-weak/F3 + Q1-split, and DERIVED the
residual \(-\nabla[v_1]/(4\pi^2)\) with forced correction
\(+g[v_1]/(4\pi^2)\) — matching DF exactly (DF chain re-inspected
verbatim; `01`-(67) sign corrected in doc 08 + revised-06),
not imported. Cancellation (N4.7); trace (N4.8); A-data
\((0,0)\) derived with explicit \(A\) open (O2 refined: blocks
only full DF framework identification). O1 closed in full (O1a:
N2 + checker #1; O1b: checker #1(d) + N4.2/N4.4 + checker #2
returned agreement — revised-06 §7).
Revised gate: PASS (repaired); statement-level formalization READY;
executable verification BOUNDED ONLY (repaired fixture, Step 7).

Re-answered handoff (prompt4 §11): 1. YES — (N4.7) with DERIVED
DF-matched coefficient. 2. NO — invalid at C2, E4–E7 (re-derived
C2/E5-numbers). 3. \(C_{\mu\nu}=+g_{\mu\nu}[v_1]/(4\pi^2)\) (N4.6)
+ uniqueness layers N3.3. 4. Revised-06 (T) with A1–A9,
exclusions, O1-closed/O2-refined/O3–O5. 5. Adds: N1-conventions,
N2-remainder, N4-coefficient (DF-matched), re-derived C2/E5,
corrected TI. 6. A1–A9 (A6 conditional, A9 source-reported);
O1 closed; O2-refined/O3 open non-blocking; O4 blocked-unneeded;
O5 bounded-fixture. 7. YES. 8. BOUNDED ONLY (repaired fixture).
9. v1 fixture SUPERSEDED — repaired fixture (Step 7) with
recalculated goldens (THIRD/CANCEL/TRACE/FLAT/NONZERO/COMBO-6 +
shapes), tightened tests, boundary statements. 10. Done: repair
record + errata + revised contract + this verdict + fixture.

Chronology reconciliation (review item 1): Q9's "NOTHING yet"
was true when §§1–5 were written (pre-implementation); §6
recorded the v1 implementation; this section records the repair
and v1-fixture supersession. Read §§1–5 as the v1 verdict
(history), §6 as the v1 implementation record (history), and
this section as the current verdict. No original line was altered
to fix the chronology; supersession is explicit.
