Plan: Deep-research extension to phys-research (sequential, Step A first — user clarification confirmed)
Status: Plan approved; Phase 0 reconnaissance to begin; no file modifications outside /home/chaschel/.opencode/plan/ in Plan mode.

=== Context ===
Plan 4 A→B→C executed and committed in phys-research:
- A (13d76b8): P-001 Eq. (1) corrected (factor 2 restored) + ## Corrections
- B (3415b6a): material-input pinning (store.ReviewIssues shared rule) + safeguard 5 (fail loud on unreadable/unparseable) + safeguard 6 (one freshness logic in validate/status)
- C (df41122): R-002 fresh adversarial review (3 pinned inputs); F1/F2 blocking upheld independently; F6 split (P-001 source half closed by 13d76b8, H-001 half still open); F3–F9 reassessed; no numeric threshold proposed; no long-range claim made
- AGENTS.md patch applied (28 insertions, edit-deny preserved)
- D deferred (no remote, hooksPath unset, no borrowed credentials, no acceptance merge)

=== User clarification answers (used) ===
- Plan file: /home/chaschel/.opencode/plan/ (Plan-mode required path — directory; writing inside as deep-research-plan.md)
- Priority: Step A first — sequential (not full protocol in parallel)
- API verification: Mandatory Phase 0 (verify arXiv, OpenAlex, Crossref, Unpaywall from actual environment before Go code)
- Previous work: Extend completed A/B/C (not independent restart)
- No new agents/databases/schedulers; standard library only; preserve seven Phys-Pseudo constructs; no new artifact family

=== Phase sequence ===
Phase 0 (MANDATORY before Phase 1):
  Verify arXiv, OpenAlex, Crossref, Unpaywall APIs from this environment.
  Record actual HTTP results, rate limits, auth requirements, endpoint availability.
  Deliver evidence in this plan directory (updated file or separate verification note).
  Do not proceed to Phase 1 until Phase 0 reviewed by user.

Phase 1 (Step A — paper acquisition):
  Implement `physctl paper` command group (`search`, `related`, `fetch`, `extract`) using only standard library.
  Add local `paper-cache/` structure (gitignored); manifest rebuildable; no committed PDFs.
  Implement optional `pdftotext` invocation (UNAVAILABLE if missing — never silent success).

Phase 2 (Step B — Researcher enhancement):
  Update `.opencode/agents/researcher.md` with iterative deep-research protocol.
  Preserve existing permissions and scientific boundaries.

Phase 3 (End-to-end demonstration):
  Topic: Classical Riemann-Cartan gravity coupled to quantum matter.
  Must demonstrate full cycle: search → discover → download → read → compare → identify gaps → formulate competing Phys-Pseudo hypotheses → compact synthesis.

=== Preservation requirements ===
- Preserve all existing source/issue/hypothesis/review/evidence records (including corrected P-001, LEGACY_UNPINNED R-001, pinned R-002)
- Preserve corrected P-001 digest; do not revert correction or alter R-001 byte-identically
- Preserve `internal/store/review.go` shared freshness rule; do not break validate/status agreement
- Preserve `AGENTS.md` edit-deny; do not weaken agent permission rules
- Preserve `core.hooksPath` unset; no `git config` changes
- Preserve standard-library-only Go convention; no new dependencies
- Preserve `paper-cache/` gitignored status; no PDF or full-text committed
- Preserve the distinction: structural validation ≠ scientific correctness; no agent declares hypothesis established
- Preserve research budget defaults (30/20/3/5/2) as operational limits, not scientific standards

=== Phase 0 — Completed (2026-11-04 UTC) ===
Actual HTTP verification performed from this environment (no invented results). Evidence saved: /home/chaschel/.opencode/plan/API_VERIFICATION.md

Results:
- arXiv: VERIFIED (200 Atom/XML, 3 results, 6017 bytes, 0.71s; sequential delay guidance noted)
- OpenAlex: VERIFIED_WITH_LIMITATIONS (200 JSON, budget note $0.10/day without key; optional API key available; 13450 results for test query; 5 returned with open_access fields verified; DOI/open_access/oa_url present)
- Crossref: VERIFIED (200 JSON, DOI 10.1103/PhysRevLett.129.121102 resolves; 19452 bytes; polite pool required; caching required)
- Unpaywall: REQUIRES_CONFIGURATION (422 without email; endpoint responds correctly; email parameter mandatory; endpoint retired for search 2026-09-18; optional supplementary only — not mandatory for pipeline)
- pdftotext: VERIFIED (version 24.02.0 at /usr/bin/pdftotext; optional, not required dependency)

No Phase 1 Go code implemented yet. User approval required before proceeding.

=== Unanswered clarification points (resolved by user confirmation) ===
1. `physctl paper` namespace: reuse existing source-ID prefix (`P-`) — CONFIRMED YES
2. `pdftotext` availability: VERIFIED present (v24.02.0) — CONFIRMED optional, never required dependency
3. Researcher model reference: confirm `opencode/deepseek-v4.1-flash-free` from R-002? — CONFIRMED preserve `opencode/mimo-v2.6-flash-free` (DeepSeek was adversary only)
4. Should this plan document be updated as phases complete? — CONFIRMED YES
5. Is the user ready to begin Phase 0? — CONFIRMED YES (completed; review evidence and approve Phase 1)
