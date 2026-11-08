## Goal

Author and get approval for Phys-Pseudo R3-W v0.1 in TOE-v6 only: for a specified nonrelativistic Bohmian system, which assumptions establish equivariance of the configuration-space distribution, and what additional assumptions connect that result to observed measurement frequencies — keeping probability weights, outcome actualization, and stable records distinct. No implementation in this plan; implementation is authorized only after you freeze the spec.

## Success Criteria

- A Deep Research dossier exists first (primary sources, established results, exact hypotheses, competing explanations, unresolved questions, novelty screen), with a source-and-claim ledger; no invented premises.
- Phys-Pseudo R3-W v0.1 uses only the seven constructs (`given`, `define`, `assume`, `claim`, `derive`, `if/then/else`, `?`), with no Go/packages/APIs/ASTs.
- Assumption + imported-result register, derivation dependency graph, and unresolved proof obligations are explicit.
- Limiting cases, failure conditions, counterexamples, and independently-verifiable claims are listed.
- All Three Reality Checks are applied; weights vs. actualization vs. records are never conflated.
- OpenCode adversarial review is dispositioned (BLOCKER/MATERIAL/MINOR); unresolved BLOCKERs narrow or block claims.
- You approve and freeze a versioned spec. No novelty claimed; benchmark tests not treated as proof of the physical result.

## Approach

Build directly on the existing Bohmian benchmark rather than re-deriving it. The TOE-v6 audit (`research/experiments/bohm-preflight-001/audit/01–03`) already freezes the system (S1: N spinless particles, Q=R^3N, Schrödinger + guidance law), registers B1–B9, derives continuity as a lemma, and leaves local/global equivariance conditional on flow/transport/no-loss premises. The standalone fixture (`theorems/bohm-equiv/`) explicitly does not execute flow, pushforward, PDE uniqueness, nodes, or frequencies. R3-W starts where that stops: formalize the equivariance→frequencies bridge and its missing premises.

Key decisions:

- **System scope:** reuse the S1 freeze and B1–B9 register as the R3-W starting point, narrowed to the same nonrelativistic spinless case. Rationale: prompt6 says build on the benchmark; changing the system now reopens settled premises. Alternative (fresh system definition) rejected as rework without dossier justification.
- **Dossier gates drafting:** no spec prose until the dossier lands. Rationale: prompt6 §1 and Constitution Articles III–V forbid inventing scientific premises; workflow.md §6 requires traceable sources/scopes/competing explanations. No R3-W dossier exists in TOE-v6 today (only mentions in `workflow.md`, `prompt6.md`, `roadmap.md`).
- **Spec home:** new TOE-v6 experiment directory (proposed `research/experiments/r3w-001/`), leaving `bohm-preflight-001/` byte-intact and never touching `/home/chaschel/Documents/go/*` (phys, phys-math, phys-lib, phys-trace, phys-pseudo, phys-research, phys-system, phys-gr, phys-docs, phys-demo, phys-artifact, phys-gemini, physmath, phys-calc surveyed read-only for this plan).
- **Spec discipline:** PSEUDO.md conformance (declaration-before-use, no assumption-smuggling through `define`, every `derive` has explicit result + non-empty premise list + reasoning, branch scope respected, `?` only for unresolved) plus MRC review readiness and DERIVATION_SPEC source-to-result chain (source→quantities→premises→operations→result→evidence→limitations), without writing code.
- **Reality Checks:** RC1 inventory structures (Schrödinger dynamics, guidance/current, ODE/flow, measure/typicality, decoherence/records); RC2 intersection/minimality (show what fails when each ingredient is removed, per `reality_check2.md`); RC3 bridge contracts (source/target/map/assumptions/preserved/lost/regime/obligations/failure/evidence, per `reality_check3.md` and `workflow.md` §4E). Gleason-type results treated with their exact hypotheses (dim≥3, additivity, non-contextuality) and never as solving actualization.
- **Review efficiency:** one full OpenCode pass + one targeted revision pass; unresolved blockers → BLOCKED, per `workflow.md` §7.

## Steps

1. **Confirm inputs (dossier + scope + drop points).** Obtain the Deep Research R3-W dossier with ledger (paper+version, exact equation/theorem/section, what it establishes, required assumptions, derived-vs-imported, closest prior, falsifier). Confirm: reuse S1/B1–B9 scope; spec path `research/experiments/r3w-001/`; how OpenCode receives the spec. If the dossier is missing, this step commissions/requests it and drafting waits.
2. **Assemble the reusable baseline (read-only).** Extract from audit 01–03: S1 freeze, B1–B9, continuity derivation consumption (B1–B3 only), flow/transport gaps (B6–B9, Teufel–Tumulka import), dispositions (DERIVED/SOURCE-REPORTED/CONDITIONAL/UNSUPPORTED). Extract from `manifest.json` limitations: 1D fragment, test units, no flow/pushforward/uniqueness/node/frequency execution. Output: a short reuse-and-gaps note the spec cites, not a rewrite.
3. **Draft Phys-Pseudo R3-W v0.1.** Cover: physical system + regime; mathematical structures with domains; equivariance chain (continuity lemma → velocity on {ρ>0} → local flow → transport/pushforward → local vs. global equivariance); typicality/measurement-model bridge to frequencies; weights/actualization/records separation; `if/then/else` for regimes (nodes, singular potentials excluded per B3/B5, decay B4 only where integrated); `?` for every undischarged obligation. Include falsifiers and at least one attempted counterexample + limiting case per major claim.
4. **Registers and graphs.** Assumption/import register (core postulates, regularity, imported theorems, approximations, boundary/initial, empirical inputs, each with ID); derivation dependency graph (premise→result edges for circularity check); unresolved-obligations list; claims-needing-independent-verification list; epistemic status per claim (DERIVED/SOURCE-REPORTED/CONDITIONAL/POSTULATED/CONJECTURAL/UNSUPPORTED/FALSIFIED).
5. **Self-audit before submission.** PSEUDO.md checklist (names, premise lists, substitution test on `define`, branch scope, `?` attachment); RC1/RC2/RC3 trace (each structure/intersection/bridge mapped to spec lines); Constitution check (no claim exceeds derivation+evidence; negative results allowed). Fix defects in a new revision, never by silent edit after submission.
6. **OpenCode adversarial review.** Submit spec + registers + dossier pointers. Review covers: mathematical validity, hidden assumptions/circularity, physical fidelity (units/regime/limits), counterexamples, literature fidelity — with BLOCKER/MATERIAL/MINOR outcomes.
7. **Revise to v1.0 candidate.** Address findings; route disputed literature back to Deep Research for verification; narrow claims or mark BLOCKED on unresolved critical defects. One targeted re-review of changed lines only.
8. **Approval package and stop.** Version, hash, and freeze request with: spec, registers, dependency graph, review dispositions, known limitations, and explicit non-authorization of implementation. End here; the implementation contract is a separate future plan after your freeze.

Non-goals: R1-S sourcing tensor definition; any Go/PhysMath/PhysLib work; phys-repo edits; novelty assessment beyond N0/N1 honesty (Deep Research adjudicates N2+ later per `workflow.md` §8).

## Validation Plan

- Step 1: dossier ledger complete (every imported premise has paper+version+locator+scope); missing dossier → plan parks, no drafting.
- Step 2: reuse note cites exact audit sections and manifest limitation IDs; no benchmark test cited as proving frequencies.
- Steps 3–4: PSEUDO conformance self-audit passes (premise closure, no forward refs, no empty `from`, no cross-branch refs); dependency graph has no cycles through `claim`-as-premise without flagging; RC trace table complete.
- Step 6: every BLOCKER has a fix or a narrowed claim; MATERIALs resolved or explicitly qualified; MINORs recorded.
- Step 8: frozen bytes + hash recorded; approval request lists what is DERIVED vs. CONDITIONAL vs. UNSUPPORTED. Highest-risk check: the equivariance→frequencies bridge — a clean equivariance derivation on top of an unexamined typicality/measurement premise is exactly the failure this spec exists to prevent.

## Risks / Open Questions

- `reality_check.md` (11,965 lines / 15 maps) was sampled at head only due to size; RC1 grounding relies on `workflow.md`'s summary plus `reality_check2/3.md` (fully read). Risk: low for R3-W scoping; flag if you want full-RC1 ingestion first.
- No git repo in TOE-v6 (`git status` fails), so freeze uses file version + recorded hash rather than a commit until you say otherwise.
- Open questions for you:
  1. Who produces the R3-W dossier, and where should it land (or point me to it if it exists outside TOE-v6)?
  2. How do I hand the spec to OpenCode from this session (paste, file drop, command)?
  3. Confirm the new spec path `research/experiments/r3w-001/` and reuse of the S1/B1–B9 scope, or name alternatives.
