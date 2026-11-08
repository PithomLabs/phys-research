## Goal

Execute `plans/prompt9.md` (Muse Assignment: JT-ISLAND-001): first audit the 25-page JT Black Hole Research Dossier v1.0 for mathematical and source-integrity defects, then author implementation-independent Phys-Pseudo R1/R4 v0.1 (R4 Page-curve reconstruction primary, R1 semiclassical consistency cross-validation) plus all supporting registers, and stop with a review-ready package for OpenCode. No code, no freeze, no implementation contract in this plan.

## Success Criteria

- Every dossier item (C-R4-01–05, C-R1-01–04, C-INT-01, BR-01–10, ASM-01–06) is dispositioned against primary sources: confirmed, conditional, incomplete, inconsistent, unsupported, or incorrect.
- All eight flagged issues (A: trace anomaly vs. conservation, B: no-island slope vs. Page time, C: island normalization/factor-of-two, D: two-interval formula scope, E: R1–R4 compatibility, F: JT action/field equations, G: interface energy balance, H: epistemic classifications) are resolved from primary literature, narrowed, or explicitly blocked.
- Phys-Pseudo v0.1 (Parts A/B/C) uses only the seven constructs, with every `derive` carrying explicit result + non-empty premises + reasoning, and R1-validated geometry/state/stress-energy as R4's actual inputs.
- RC1/RC2/RC3 applied; bridge register uses the ten-field contract; N0/N1 contribution honestly bounded (no unitarity/microstate/4D/Bohmian promotion).
- All eight supporting artifacts (§6) delivered; OpenCode verification checklist identifies what must be independently reconstructed.
- Nothing written outside TOE-v6; no phys-repo modifications; no Go/packages/goldens.

## Approach

Audit before formalizing, single model throughout, primary sources over dossier prose. The dossier is research input, not authority (prompt9 §3), and it already shows symptoms that justify this ordering: its RC2 claims removing the trace anomaly breaks covariant conservation (issue A — a trace anomaly is not a diffeomorphism anomaly), and prompt9 flags coefficient/factor-of-two inconsistencies across the entropy branches and Page time (issues B, C).

Key decisions:

- **Equation ground truth is rendered pages + arXiv, not PDF text extraction.** I extracted the dossier with `pdftotext` (1078 lines): all displayed equations are lost (rendered images), leaving only prose and structure. Equation-level audit (signs, coefficients, factors of two, cross-ratio dependence) must therefore be done against the rendered PDF pages and the canonical primary sources, never against the text extraction.
- **Fetch AMMZ + AEMM fresh into TOE-v6, version-pinned and hashed.** The local phys-research paper cache does not contain `1908.10996` or `1905.08762`. Fetch AMMZ v2 (the dossier cites the arXiv HTML v2, which preserves exact equations) and AEMM from arXiv into TOE-v6 workspace or scratch; record version + SHA-256. Prefer arXiv canonical over the dossier's ResearchGate/eScholarship mirrors. Never write fetched files into `/home/chaschel/Documents/go` (paper-cache writes included).
- **R1 outputs are R4 inputs by verification, not by naming.** C-INT-01 (shared state/geometry/central charge) must be independently established per issue E (state definitions, coordinate transforms, conformal factors, cutoffs, regulators, conventions); a shared model name proves nothing.
- **Dossier handoff type names do not enter Phys-Pseudo.** The dossier's "Mathematical Objects to Represent" (`SpacetimeManifold2D`, `DilatonField`, `CFTState`, …) are software-flavored; prompt9 §4 requires mathematical declarations, not software abstractions. Translate the mathematics; leave the type names behind.
- **Spec home:** `research/experiments/jt-island-001/` in TOE-v6 (audit note + spec + registers) — approved by director review; existing experiments byte-intact.
- **Capability gaps recorded, physics never bent.** For each proposed operation, classify per §8 (definition/derivation/import/symbolic-check/numerical/unsupported); unsupported stays an explicit gap, never a silent spec change.

## Steps

1. **Assemble audit inputs.** Fetch AMMZ arXiv:1908.10996v2 (PDF + HTML), AEMM arXiv:1905.08762, and Polyakov's original work (for the trace-anomaly/Ward-identity audit, where necessary) into TOE-v6; record versions/hashes. Confirm rendered-PDF reading path for dossier equations. Create `research/experiments/jt-island-001/` skeleton.
2. **Dossier audit (§3, issues A–H).** Claim-by-claim assessment of all C/BR/ASM identifiers against primary-source equations: trace-anomaly/Ward-identity separation consulting Polyakov's original argument, not only AMMZ's restatement (A); no-island expansion vs. stated slope, independent Page-time derivation (B); per-endpoint dilaton terms vs. BH entropy definition and endpoint count (C); two-interval formula hypotheses — cross-ratio, channel, state, approximation (D); R1↔R4 state/coordinate/cutoff/convention reconciliation (E); independent action variation with signs/coefficients/boundary terms (F); bulk-vs-interface energy accounting, Schwarzian/ADM conventions (G); island-formula/QES/minimization epistemics (H). Output: dossier-audit note (confirmed/corrected/qualified/blocked per claim).
3. **Establish the audited working model baseline.** Provisional, mathematically corrected foundation for v0.1 — not an approved or frozen spec. Canonical JT+bath model, domains, metric charts, dilaton, CFT matter, state, hierarchy `1 ≪ c ≪ φ₀/G_N` with verified meaning, radiation/island regions, conventions. Anything issue E/F/G leaves unresolved stays an explicit `?`/BLOCKED, never an assumption by default. Formal approval and freezing remain the project director's responsibility after OpenCode's review.
4. **Author Phys-Pseudo v0.1 Part A (R1 chain, §4 ¶242–257).** Action → domains/metric/BCs → field equations → Polyakov/stress-energy → trace anomaly → Ward identity/conservation → light-cone forms → interface matching → Schwarzian/ADM → energy balance. Keep bulk conservation, interface matching, and global accounting as separate derivations.
5. **Author Part B (R4 chain, §4 ¶259–280).** R_t, I_t, generalized-entropy functional, dilaton + matter contributions, no-island branch, island branch, QES stationarity, QES solution in documented approximation, on-QES evaluation, branch comparison, Page transition from intersection, branch-selection prescription, Page curve. Extremization and branch minimization as separate operations; no assumed QES existence outside the validated regime.
6. **Author Part C (interoperability, §4 ¶282–301).** The ten R1↔R4 connections as explicit derivations or exposed obligations; unestablished links become `?` items, not prose bridges.
7. **Reality Checks + registers (§5–§6).** RC1 structure inventory with role/regime/limits; RC2 intersections with removal-failure arguments (no minimality claims without proof); RC3 ten-field bridge register; assumption/import register (seven categories); claim register (ten epistemic labels); dependency graph R1→R4→cross-validation with circularity check; obligations list; OpenCode verification checklist; verified limiting cases (vanishing coupling, reflecting bath, hierarchy breakdown — consequences re-derived, not copied).
8. **Self-audit + review package (§9–§10).** PSEUDO conformance (declaration-before-use, substitution test on `define`, premise closure, branch scope, `?` attachment), Constitution check (claims ≤ derivation+evidence), final summary (established/derived/imported/conditional/unresolved/demonstrable/forbidden). Assemble the package as a directory with an index identifying every file and its purpose, and stop; no implementation. (OpenCode is invoked against that package separately.)

Non-goals: OpenCode review itself, approval/freeze, implementation contract, Go fixtures, R2/R3/R5, 4D or Bohmian extensions, novelty claims.

## Validation Plan

- Step 1: AMMZ, AEMM (and Polyakov original where the audit requires it) present in TOE-v6 with pinned versions + hashes; rendered-dossier equation access confirmed (text extraction alone fails this gate by construction).
- Step 2: every C/BR/ASM ID has a disposition with a primary-source locator; issues B/C include independent algebra (expansion, endpoint count, intersection solve), not dossier quotation; issue A states the two identities separately with their distinct derivations.
- Steps 4–6: PSEUDO self-audit passes (no forward refs, no empty `from`, no cross-branch refs, no Go/AST leakage, no handoff type names as constructs); Part B consumes only Part A–validated inputs (traceable by name).
- Step 7: dependency graph acyclic except flagged conditional edges; each limiting case re-derived to its stated consequence; bridge register complete on all ten fields.
- Step 8: final summary's seven bullets each trace to registers; package contains all §9 items. Highest-risk validation: the B+C+E triangle — if branch normalizations, Page time, and R1/R4 state compatibility do not mutually agree, the spec must narrow or block, not average the disagreement.

## Risks / Decisions

- Dossier works-cited mixes canonical arXiv with mirrors and unrelated entries (Hugging Face daily papers, theses); audit follows arXiv-canonical only and treats the rest as unvetted.
- The dossier asserts "no circular dependencies" and "ESTABLISHED (N0/N1)" scope outcomes for its own chain; both are audit targets, not accepted premises (prompt9 §§3H, 9).
- Decisions (recorded from director review):
  1. Spec path `research/experiments/jt-island-001/` — APPROVED; audit note + spec kept together as separate files with supporting registers in the same experiment directory.
  2. OpenCode handoff: complete review package as a directory with an index identifying every file and its purpose; OpenCode invoked against that package separately.
  3. Load-bearing contradiction: PAUSE AND REPORT before proceeding with dependent derivations; independent unaffected research may continue; never build downstream conclusions on an unresolved contradiction.
