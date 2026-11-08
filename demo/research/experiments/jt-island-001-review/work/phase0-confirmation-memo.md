# Phase 0 Confirmation Memo — prompt10 execution (`go`, 2026-11-08)

## Decision: VERIFY-FIRST — no overwrite of the existing v1.0 review

The target deliverable `research/experiments/jt-island-001/reviews/opencode-review-v1.0.md`
already exists (1853 lines). Per the approved plan, this execution verified it
instead of blindly overwriting it. No file in the package or in
`/home/chaschel/Documents/go` was modified. This memo is the only new artifact,
written to scratch (`research/experiments/jt-island-001-review/work/`).

## 1. Provenance (start hashes, this run)

All package + source hashes match the prior start manifest
(`work/package-hashes-start.txt`) and `sources/README.md` exactly:

- `dossier-audit-note-v1.0.md`: `c18c6002…` (BYTE-INTACT anchor) — match
- AMMZ v2 PDF: `f68be764…` (22 pp) — match
- AEMM v3 PDF: `f51a59cb…` (55 pp) — match
- AP v3 PDF: `37e78a45…` — match
- Penington PDF: `fa7cf02d…` — match
- All 10 root package `.md` files + 3 HTML cross-checks — match

Conclusion: package and pinned sources are byte-identical to the reviewed state.
Prior extraction artifacts (txt/png/info) remain in `work/`.

## 2. Structural compliance of existing v1.0 vs prompt10 §10

Present and in order: §§1–15 (executive verdict; methodology+versions;
reconstructions; findings; OB-1–10; R1; R4; interop; RC1–3; conformance;
limits/counterexamples; corrections; accepted; rejected/narrowed; final verdict)
plus Appendices A–G (hashes, locator map, symbols, ledger, bridges, obligations,
finding index). Finding count 24 = 3 BLOCKER + 14 MATERIAL + 7 MINOR, matching
the §1.3 summary. Spot-checked finding F-01 carries all nine prompt10 §8 fields
plus the plan-required Disclosure field.

## 3. Independent spot verification of the three BLOCKERs (this run)

- **F-01 (hierarchy contradiction) — CONFIRMED.** Spec line 12 asserts both
  `phi0>>phi` and `S_0<<S^i_Bek` in one line; baseline §7 lines 223–224 assert
  `φ₀ ≪ φ_hor-scale` and `φ₀ ≫ φ` adjacently. AEMM PDF p.8 (text after (2.3)):
  dilaton `φ0 + φ`, "implicitly imposing the restriction φ0 >> φ". AMMZ PDF
  p.14 ties island viability to small S₀. Since S_Bek=(φ₀+φ_hor)/4G_N, the two
  conditions are φ₀≫φ_hor and φ₀≪φ_hor: no overlap.
- **F-04 (endpoint/2S₀ inconsistency) — CONFIRMED.** Spec adopts one-endpoint
  `I_t=(-∞,a(t)]`, `∂I={a}` (lines 52–57) but invokes the `2S₀` island price
  (line 74, citing AMMZ §3.3). AMMZ PDF p.14 verbatim: "we need to pay a 2S0
  price to create the 'island' (because of the two endpoints of the island)".
  The cited price explicitly requires two endpoints.
- **F-09 (circular/restatement links) — CONFIRMED by direct read of spec
  lines 88–92.** `link_matter_sgen` concludes `sgen_functional`'s content from
  `sgen_functional`; `link_sgen_qes` concludes about `qes_position` from
  `qes_position`; `link_branches_rad` concludes `page_curve` from `page_curve`;
  `link_flux_mass` concludes about `global_balance` from `global_balance`;
  `link_dilaton_area` restates `dilaton_term` + the endpoint postulate.

Verdict logic holds: ≥1 unresolved BLOCKER → overall BLOCKED; the three
governing questions (model coherence, Page-curve consistency, genuine interop)
are each answered with evidence.

## 4. What this run did NOT do

Full independent re-derivation of every R1/R4 equation (Phases 1–6 of the plan)
was not repeated: the existing v1.0 already contains that reconstruction
record (§3, §§6–11), its load-bearing source quotes were re-verified above
against the pinned PDFs, and re-running the full pass would duplicate rather
than test it. A fresh full re-execution or a v1.1 requires separate director
authorization per the plan's stop conditions.

## 5. End-hash QA

Re-run `sha256sum` over the same file set after this memo: all package and
source hashes unchanged (this memo lives outside the package, in scratch).
Zero writes under `/home/chaschel/Documents/go`. Review-only boundaries held:
no code, no freeze, no implementation authorization, no novelty claim.
