## Goal

Implement `plans/prompt5.md` (Agent A — QSG-T001 normalization repair) in light of `plans/plan5_imp_review.md`: fix the incorrectly normalized Hadamard subtraction, reconstruct the conservation calculation from corrected definitions, determine whether the reported correction coefficient survives, and reconcile all theorem/fixture metadata — without restarting the programme, rewriting recon, or adding infrastructure.

## Success Criteria

Exactly the prompt5 §7 checklist — the task is complete only when: singularities match exactly; the remainder is genuinely smooth; the point-split limit is well-defined; the remainder equation is independently derived; the residual is explicitly calculated; the correction is derived, not imported; its divergence is calculated; the correction divergence is calculated; cancellation is demonstrated; source correspondence is reconciled or explicitly BLOCKED; all metadata agree; tests demonstrate the limited capabilities claimed. Any failed math gate → BLOCKED, theorem not promoted.

## Approach

Triage first (my session check, not a result): I verified the reviewer's singularity arithmetic two independent ways — the Minkowski Wightman mode integral gives G⁺ ~ 1/(8π²σ), and DF's Feynman form (inspected in doc 03) gives the same — so G⁽¹⁾ ~ 1/(4π²σ) vs H⁽¹⁾ ~ 1/(8π²σ), and S as written is singular. The defect is real; P2 is withdrawn. The full independent verification is Step 1 itself.

Key decisions:

- **Prescription choice:** recommend Alternative A (S = (1/2)G⁽¹⁾ − H⁽¹⁾, standard H kept) pending Step-1 proof — minimal diff to locked §0.3, matches DF/BD H-normalization. A-vs-B looks notationally equivalent (same S either way); Step 1 must prove that or find the difference. Decided by evidence in-step, not assumed.
- **What transfers:** the D1/TI divergence proof *structure* (a cross-checked conditional: *if* S smooth with equation J, *then* residual = (1/2)∇[J]) may be reused only after step-by-step re-verification against the new S — the old steps are vacuous (old S singular), so nothing is copy-pasted. L2 must be derived fresh (transport equations stated and proved), per review §3 / prompt5 §3.3 — no import at the crucial point.
- **C2 reversal anticipated:** with correct normalizations, Agent B's W = G⁺ − H has smooth *real* part (B's subtraction captures the right singularity); B's defect restates as the unaddressed distributional imaginary part plus the false homogeneous equation — not singularity. Step 5 must verify this explicitly; the reversal is recorded openly, never silent.
- **No coefficient preview in this plan:** any scaling guess from the old (vacuous) derivation is forbidden as a target. Reconstruction is the only route; the old 3/(16π²) is treated as unestablished ink.
- **Blocking semantics** per review §4: O1 open with explicitly decided significance; O2 blocks DF-equivalence claims, not the internal theorem; O5 never blocks the mathematical theorem. Reconciled across 06/07/JSON.
- **Fixture frozen** until corrected math gates pass; then goldens/metadata only, architecture intact, SK-steps unsupported.

## Steps

1. **Normalization verification (prompt5 §§1–2).** From the Minkowski mode integral plus DF correspondence (re-inspected, not recalled): verify Wightman, symmetric-function, and Feynman singular coefficients independently; prove or refute A⟷B equivalence; freeze the corrected prescription with derivations. Gate: singularities match exactly.
2. **Remainder reconstruction ("03-repair", prompt5 §3.1–3.4).** Prove S_new smooth; re-derive its inhomogeneous KG equation; state and derive the required Hadamard transport equations (closes L2); establish coincidence limits. Gate: well-defined point-split limit.
3. **Divergence reconstruction ("04-repair", prompt5 §3.5–3.9).** Re-execute D1/TI for S_new with every Synge instance shown; derive residual, force correction from residual, compute its divergence, demonstrate cancellation. No old coefficient reused. Gate: explicit residual → correction → cancellation.
4. **Source + A-term audit (prompt5 §3.10, §4).** DF/Wald correspondence under corrected definitions; A smoothness and field-equation limit computed — reconcile or mark BLOCKED. Either outcome acceptable; unspecified smooth terms forbidden as cover.
5. **Finding repair.** Correct doc-05 C2 (reversed arithmetic, restated defect) and C1; check every other 05 verdict for normalization-dependence (expected unaffected: C3–C7, D–J — verify each, assume none); restate E4 for S_new.
6. **Provenance + metadata repair (prompt5 §5, review §§1/4/5).** Append errata (locks byte-intact); rewrite 06 (theorem, coefficients, re-answered gate, O1–O5 per §4 distinctions); fix 07's status/chronology contradiction and verdict; reconcile research_metadata.json + manifest.json blocking semantics; rewrite Phys-Pseudo with complete premise lists (field-equation, geometric, operator, Synge, transport premises all named).
7. **Fixture repair (prompt5 §6, review §5).** Update affected goldens/anti-goldens/metadata only; tighten tests (inject unknown fields, fixed golden snapshots, correct mutation descriptions); retain architecture and boundaries.
8. **Final validation.** Full gate per-command (build/test/vet/gofmt-empty); non-modification audit over the extended protected set (locked sections byte-identical, before/after hashes); §7 checklist item-by-item; BLOCKED report if any math gate fails.

Non-goals: new surveys, recon rewrites, new math/Go infrastructure, stochastic/backreaction/FLRW extensions, touching shared repos or the committed go.work.

## Validation Plan

- Steps 1–4: each gate is a derived equation with its derivation shown (singular-coefficient match, smoothness proof, transport derivation, residual/correction/cancellation); reviewer can check each line against definitions.
- Steps 5–6: errata appended (never silent edits — verified by lock-hash re-check over original scopes); 06/07/JSON agreement checked field-by-field (O-status table must match in all four places); Phys-Pseudo premise-completeness self-audit per PSEUDO §7.
- Step 7: `go build`, `go test -count=1`, `go vet` exit 0, `gofmt -l` empty under the GOWORK/GOCACHE recipe; mutation tests kill wrong-coefficient and tampered-byte mutants; manifest/sidecar tests assert the tightened properties. Highest-risk validation: the Step 2–3 reconstruction itself — a formally clean derivation on wrongly normalized objects is exactly the failure being repaired, so Step 1's singularity match is the load-bearing check everything else stands on.

## Risks / Open Questions

- Reconstruction may fail on its own terms (residual not a pure gradient, correction non-local) — then BLOCKED per prompt5, which is a valid outcome, not a plan failure.
- L2-closure depth: bounded to stating and deriving the *required* transport equations plus order-of-vanishing, not a Hadamard-theory treatise.
- Anchoring: the old coefficient must not leak in as a target — mitigated by derive-first ordering and by this plan stating no preview value.
- C2 reversal changes a published audit verdict — mitigated by explicit erratum with before/after arithmetic, cross-referenced from 06/07.
- External research: none — no new libraries or methods; DF text already inspected; Minkowski verification is analytic derivation.
- Open questions for you: none — prompt5 is fully prescriptive.
