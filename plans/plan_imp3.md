Plan 8 (as amended) is executed. No commits made; all work is uncommitted in the tree for your review.

## 1. Reconnaissance report

What already worked: record/validation/status tooling, paper acquisition (6 cached PDFs with manifests), provenance discipline, and honest conditional claims in I-002/H-002/E-002 after the Nov-4 corrective pass. What was genuinely missing: no equations in any P-record; zero Branch C `?` obligations surfaced; D-001 a malformed comparison essay (pseudo-header refs invisible to `physctl refs`, duplicated bullets, typos); single hypothesis; `physctl paper` undocumented; live fetch broken. Two deeper finds: P-004/P-005 carried false DOIs behind bogus "verified against Crossref" claims, and prior "verification" had been fabricated — both corrected with live API evidence.

## 2. Implementation report

- [agents/researcher.md](/home/chaschel/Desktop/biz/TOE-v5/phys-research/agents/researcher.md) (+adversary/formalizer): reconciled per amendment 1 — `.opencode` copies held the latest methodology, so their content became the neutral canonical base (deep-research Stages A-K preserved verbatim), plus 5 research-quality rules (qualification scale, equation template, framework checklist, transfer rule, `?` rule).
- [cmd/physctl/main.go](/home/chaschel/Desktop/biz/TOE-v5/phys-research/cmd/physctl/main.go): usage text documents `paper` subcommands. (Note: `cmd/**` is constitution-restricted; touched only for this, under your explicit tasking.)
- [internal/paper/paper.go](/home/chaschel/Desktop/biz/TOE-v5/phys-research/internal/paper/paper.go): `safeFetch` bypassed the egress proxy (custom Transport, `Proxy: nil`) so **all** fetching failed in proxy environments — demonstrated blocker, fixed by honoring `HTTPS_PROXY` with TLS verification intact; direct path keeps IP validation. Neutralized the `opencode-research` user agent.
- [internal/paper/transport_test.go](/home/chaschel/Desktop/biz/TOE-v5/phys-research/internal/paper/transport_test.go): new hermetic test (both transport constructors + dispatch consistency; no network).
- [internal/record/record.go](/home/chaschel/Desktop/biz/TOE-v5/phys-research/internal/record/record.go): review scaffold `reviewer:` line neutralized.
- Sources: P-007 (Trautman 2006, foundational, fetched via repaired pipeline); P-003–P-006 gained `## Mathematical content` with load-bearing equations, locators, and asserted/derived status; P-004 false DOI removed (belongs to Oppenheim's PRX 2023 paper), P-005 DOI corrected to Quantum 7, 891, P-006 venue downgraded to UNVERIFIED. P-001 deliberately untouched (editing it staled R-002; reverted; its "related" qualification lives in the dossier).
- [derivations/D-001.md](/home/chaschel/Desktop/biz/TOE-v5/phys-research/derivations/D-001.md): repaired — real header refs (now visible in `physctl refs`), deduped O1–O7, typos fixed, per-framework variables/equations/conservation/regime, explicit transfer mappings, P-007 baseline.
- [research/hypotheses/H-002.md](/home/chaschel/Desktop/biz/TOE-v5/phys-research/research/hypotheses/H-002.md): 7 prose gaps → `?` lines O1–O7; applied the documented-but-unexecuted removal of `assume stochastic_extension_applies`.
- [research/hypotheses/H-003.md](/home/chaschel/Desktop/biz/TOE-v5/phys-research/research/hypotheses/H-003.md) (new, `challenges` H-002): semiclassical-sufficiency null with distinct assumptions and a two-part falsification condition; H-002 vs H-003 is decided by O3/O5/O6.
- [research/evidence/E-002.md](/home/chaschel/Desktop/biz/TOE-v5/phys-research/research/evidence/E-002.md): wired to H-003, fixed stale "P-002 stochastic framework" framing.
- [docs/dossier-branch-c.md](/home/chaschel/Desktop/biz/TOE-v5/phys-research/docs/dossier-branch-c.md): 10-section dossier index over record IDs.
- Portability: `.opencode/agents/*.md` slimmed to adapters (config only, canonical pointer); [docs/AGENT_CONTRACT.md](/home/chaschel/Desktop/biz/TOE-v5/phys-research/docs/AGENT_CONTRACT.md) (mandatory core vs optional); [docs/compatibility-matrix.md](/home/chaschel/Desktop/biz/TOE-v5/phys-research/docs/compatibility-matrix.md); `.gitignore` covers `.ratelimit/`.

## 3. Functional demonstration

Branch C question → 7 examined full texts → qualified map (1 foundational, 4 directly relevant, 1 related, 1 tangential) → transcribed equations (P-003 Eqs. 4.18–5.37; P-004 Eqs. 67–68 + trade-off; P-005 Eqs. 4–5, 28; P-006 Eqs. 15–22; P-007 Eqs. 11–33) → D-001 comparison with transfer mappings → 10 open `?` obligations (7 H-002 + 3 H-003, all in `status`) → H-002/H-003 competition → dossier. No OpenCode subagent run exists in this environment, so the cycle was executed by this session following the Researcher protocol — labeled as such throughout, never as an agent run.

## 4. Scientific quality assessment

Independently reproduced: **nothing** — stated plainly in dossier §6. All mathematics is source-reported (transcribed, PDF-verification caveat attached) or conjectured. Assertions corrected, not laundered: 3 false/unverified provenance claims fixed, LISA Pathfinder figure refused transcription (ambiguous render), P-001 left pristine. Open: O1–O7 + H-003's three questions. Strongest result: P-006 semiclassical EC on the P-007 baseline. The dossier distinguishes source-reported / verified-provenance-only / conjectured / empirical / unresolved.

## 5. Validation results

`go build ./...` clean; `go test ./...` 9/9 packages pass (incl. new transport test); `physctl validate` 0 errors (1 pre-existing R-001 LEGACY_UNPINNED warning); `status --branch C` shows 12 `?` lines; `refs` traverses D-001/I-002/H-002/H-003 correctly. Compatibility: Muse column VERIFIED in an isolated worktree (since removed); Codex/OpenCode NOT VERIFIED — discontinued per your instruction before any model call, setup notes preserved in the matrix doc.

## 6. Remaining blockers

1. `AGENTS.md` §§2–4 still OpenCode-framed — needs your edit (untouchable by agents). Note: you added §12 in parallel during this session; I never touched that file.
2. No independent derivation exists anywhere in Branch C — the next cycle needs paper-and-pen (or Formalizer-transcribed) mathematics, starting below.
3. P-006's journal publication is unverified; P-004's LISA figure needs PDF inspection before citation.

## 7. Next recommended research action

Attempt O5 on paper: extend the P-003 Dirac-argument method to EC variables for the momentum constraint. A clean failure (non-closure/over-constraining) is as valuable as success — it decides H-002 vs H-003.
