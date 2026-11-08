# JT-ISLAND-001 resumption report (review package complete)

**Top finding: the complete OpenCode review package is assembled and self-audit gates pass — resumption plan fully executed, stopping here as instructed.**

## What was delivered

All in `research/experiments/jt-island-001/` (TOE-v6 only; zero writes to any phys repo):

| File | Content |
|---|---|
| `phys-pseudo-r1r4-v0.1.md` | **The specification** — Parts A (R1), B (R4), C (10 interop links), 63 names, seven constructs only |
| `working-baseline.md` | Corrected equations, coordinate/convention/symbol dictionaries, hierarchy, per-class AEMM↔AMMZ compatibility, ?OB-1..10 |
| `branch-reconstruction.md` | Independent branch reconstruction + intersection derivation, t_P derived-not-frozen with footprint F1–F5 |
| `audit-corrections-v1.1.md` | Applied director rulings + corrected C/BR/ASM dispositions (v1.0 byte-intact, hash-verified `c18c6002…`) |
| `registers.md` | Assumption/import, claim, dependency-graph, ten-field bridge, obligations, limitations registers |
| `reality-checks.md` | RC1 landscape, RC2 intersections, RC3 bridges (RC2 removal claim replaced; no minimality claims) |
| `opencode-checklist.md` | 18-item independent-verification checklist + BLOCKER/MATERIAL/MINOR protocol |
| `summary.md` + `index.md` | Seven-bullet §10 summary; indexed package handoff (read order, entry points) |
| `sources/` | AMMZ v2, AEMM v3, AP v3 (pdf+html, hashed) + Penington (pdf, bounded-read recorded) |

## Key resolutions along the way

- **Director's model correction honored:** two-sided AEMM + AMMZ holographic specialization (not the paused single-sided recommendation). This forced the Step-4 re-resolution: the pure-state complement argument is unavailable two-sided, so R4 entropy goes through **AMMZ's 3D-RT route**, with AEMM's G(η) machinery as documented fallback. Decisive evidence: AEMM's own bath entropy (112) grows forever — only AMMZ constructs a radiation Page curve.
- **Genuine interop derived, not assumed:** AMMZ↔AEMM evaporation-rate identity (κ=k) proven from AEMM(20)+(25); slope πcT/6 triple-verified (log expansion, AMMZ(14), AEMM(112)).
- **Page time honestly open:** leading estimate 6βS/πc reproduced with a five-item footprint (F1 shows t_P is O(k⁻¹), so slow-evaporation is an O(1) approximation) — frozen neither at 3βS/πc nor 6βS/πc per Decision 2.

## Verification evidence

- Mechanical PSEUDO check: **CLEAN** (63 names, premise closure, declaration order, no duplicates, derive completeness, `?` attachment, branch-scope isolation all pass).
- Dossier-independence grep: every "dossier" mention in the spec is a rejection/void note or the director-adopted configuration postulate — no dossier equation is a premise.
- Implementation-leakage grep: empty. Dependency graph acyclic by construction (R1→R4 one direction).
- Constitution check: all derives carry premises + reasoning; imports/conditionals labeled; scope held to N0/N1 with explicit must-not-claim list.

## What's next (not done here)

OpenCode review (invoke separately against the indexed package), then the freeze decision. Open carriers for review: ?OB-4 (dM/dt line-by-line), ?OB-5 (shock-vs-smooth), ?OB-7 (t_P footprint), ?OB-9/?OB-10 (QES/RT wholesale imports). No code exists; nothing to build or run.
