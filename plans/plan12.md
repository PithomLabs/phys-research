## Goal

Implement `prompt15.md` in light of `adv_review_d004.md`: correct D-004 into a new superseding D-005, reconcile the pair classification (including the 0− mixed-pair question) and obligations, and commission a fresh R-006 — without editing D-004, without touching R-005, and without proceeding to PhysMath/PhysLib.

## Success Criteria

- D-005 exists as a new record superseding D-004, with the φ⊥ coefficient and adjacent-case comparison set by fresh primary evidence, the 0− sector explicitly accounted for, and obligations reconciled between the formal list and Phys-Pseudo.
- R-006 independently re-examines every D-005 derive, adjudicates each `adv_review_d004.md` claim, and issues per-derivation dispositions with an evidential boundary.
- D-004, both R-005 versions' content, and all prior hashes are byte-identical afterward; `physctl validate` stays at 0 errors.
- No library repo is modified; no formalization is attempted.

## Approach

Verify before correcting. Research surfaced three conflicts that make blind implementation unsafe, so evidence comes first:

- **Conflicting φ⊥ readings.** My session's PDF render read Eq.(46) as −3a0/l²; the adversarial review parses 3a0/(2l) from extraction plus a dimensional argument. One reading is wrong.
- **Internal tension in the review.** It claims Eq.(46) gives 3a0/(2l) while Eq.(94) gives 3a0/l — but (94) is (46) evaluated at different b-parameters, and the a0 term cannot differ between them. At least one of those two readings must be wrong, which is itself a verification target.
- **Two divergent R-005s.** My worktree's R-005 (hash a95eb322) verified D-004; `adv_review_d004.md` is a different agent run's failing R-005 (its D-004 hash d9291ba4 ≠ mine 196a3b88, though quoted lines match mine substantively). R-006 must adjudicate both, not assume either.

The Table 1 0− row extraction is also garbled, so the mixed-pair claim needs the same render-or-source treatment. If fresh evidence contradicts prompt15's prescribed values, D-005 records what the evidence shows instead of forcing the prescription.

## Steps

1. **Baseline and evidence re-verification.**
   - Record `git status`, hashes of D-002/D-003/D-004/R-003/R-004/R-005, and `physctl validate` output.
   - Fresh high-resolution PDF renders of Yo–Nester Eq.(46) region (p.8) and Eq.(94) region, read independently; settle each constant (2l vs l² vs l) and whether the two equations' a0 terms agree.
   - Re-read Table 1's 0− row (mass combos, pφ/pχ pairing); render if the extraction cannot settle it.
   - Dimensional cross-check of π⊥/J vs the candidate constants as supporting (not deciding) evidence.

2. **Author D-005 (supersedes D-004).**
   - `physctl new --type derivation` with `supersedes D-004`, `depends_on D-002`, I-002/P-007 refs as before.
   - Set the φ⊥ coefficient and adjacent-case comparison strictly from step-1 evidence: if the renders confirm prompt15's table (3a0/(2l) EC vs 3a0/l spin-2−, non-matching), implement it; if they show anything else, record the evidenced values and flag the prompt15 discrepancy explicitly.
   - Rewrite `pair_structure` to account for the 0− sector as the evidence shows (mixed pair vs clean pair), keeping the 2− nonlinear-PB obligation intact.
   - Reconcile obligations: one consistent D2′/D2map/D2-comp (+0− as needed) set, identical between the formal Obligations section and the `?` list.
   - D-004 untouched; no edits to any existing record.

3. **Adversarial review R-006.**
   - `physctl new --type review --review H-002`, pin D-005 digest; scope statement names `adv_review_d004.md` (TOE-v5 root, left in place, cited by path) as an input claim-set alongside D-005.
   - Re-perform every D-005 substitution against the source; rule explicitly on each adv-review claim (F2 coefficient, F1 partial resolution, 0− accounting, Hbar/conditionality points) as confirmed, refuted, or still open.
   - Per-derivation dispositions, Where/What/Why/Minimum-fix findings, evidential boundary distinguishing render evidence from extraction reads. No global verdict.

4. **Close-out.**
   - `physctl validate`, `physctl status --branch C` (D-005 `?` lines surface), hash re-check of all frozen files, library-repo cleanliness check.

## Validation Plan

- `git status --short` / `git diff --stat`: only D-005 and R-006 added; D-001/H-002 diffs unchanged in content; D-002/D-003/D-004/R-003/R-004/R-005 hashes identical to baseline.
- `GOCACHE=/tmp/muse-gocache go run ./cmd/physctl validate` → 0 errors; pre-existing R-001 warning only.
- PSEUDO conformance checklist on D-005 (named premises, explicit results, `?` placement, no in-document statuses, obligations list consistent with `?` lines).
- Review checklist: every D-005 derive has a disposition; every adv-review claim has an explicit ruling; digests recompute; reviewed records unmodified.
- Highest-risk validation: the dual-render read of Eqs.(46)/(94) — the entire correction hinges on settling 2l vs l² vs l from primary pixels, not from extraction text or prior claims.

## Risks / Open Questions

- If renders are ambiguous at any resolution, D-005 must record the constant as disputed with both readings and the dimensional argument, rather than picking one; R-006 then dispositions it UNSUBSTANTIATED.
- If evidence contradicts prompt15's prescribed corrections, the plan prioritizes evidence over prescription and reports the divergence plainly.
- The other agent's divergent record versions (differing D-002/R-004/R-005 hashes) are out of scope; this plan reconciles at the claim level via R-006, not by merging worktrees.
- `physctl validate` remains structural-only; `physctl check` without the untracked local `go.work` reports `UNAVAILABLE`.
