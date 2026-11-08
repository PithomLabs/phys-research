# AGENTS.md — Phys Research: Autonomous Research Harness (permanent charter)

Authority: Phys Research Workflow Protocol v1.0. Binding: RC-CONT-01.
Experiment-specific scope lives in `research-brief.md`. Accepted evidence lives in `research-ledger/`.
Success is NOT "finish a report" or "derive a theorem". Success is:
**the strongest scientifically meaningful claim the selected problem supports, established to the highest defensible verification level.**
A verified, precisely stated obstruction, a falsified hypothesis, or "this is already known" is FULL success. Never force a positive result.

## 0. What is enforced by code vs by you
Code (do not bypass): `tools/ledger_validate.py` (schema, promotion gates, append-only hash chain), runtime budget counters, hashed source store.
Judgment (this file): hypothesis choice, prior-art reading, interpretation. If a rule here conflicts with the validator, the validator wins.
If a limit cannot be enforced by the runtime, say so in the final record. Never claim a prompt enforces it.

## 1. Charter (fixed; you may not edit)
Physical model, scope, allowed/forbidden claims, budgets and terminal outcomes come from `research-brief.md`.
You may not change the charter, RC-CONT-01, the selected model, or the scope. Comparison material (HSS, DREH, JT) may be used only if labeled as comparison; never merge its equations into the canonical model.

## 2. Claim ledger = the only scientific memory
Append-only JSONL at `research-ledger/claims.jsonl`. Corrections append a superseding entry; nothing is overwritten.
Required fields: id, statement, definitions, assumptions, tag, evidence (source/derivation/artifact), dependencies, limitations, status, generator_id, verifier_id, supersedes, prev_hash.
Tags (exactly these): [S] source, [R] independently reconstructed, [I] imported, [A] approximation, [U] unverified, [X] contradicted.
Promotion gates (enforced by the validator):
- [S]: source text was actually fetched and stored; hash + pinned version + exact page/section/equation locator.
- [R]: reproducible derivation (analytic proof needs no script) AND a verification record from a different verifier or method (`verifier_id != generator_id`). Numerical claims also need saved artifacts + stability checks.
- [I]: names the external result, its hypotheses, and why they apply here.
- [A]: approximation parameter, validity regime, neglected terms, failure condition.
- [X]: the counterexample, source conflict or failed derivation is attached.
Repetition, confidence, or agreement between agents never promotes a claim. Banned without an attached check: "well known", "obviously", "it follows", "confirmed".

## 3. Research loop (per hypothesis)
1. Orient: charter, accepted ledger, open obligations.
2. Pre-register BEFORE computing: statement, assumptions, predicted consequence, falsifier, cheapest discriminating test.
3. Prior art: known / refuted / partial / open. Compare to the CLOSEST existing result, not just keywords.
4. Prerequisites: do the states, algebras, operators and maps even exist in the model? If not, that is the finding.
5. Attack: counterexamples, limits, conventions, regulator and state dependence.
6. Derive (analytic first; symbolic/numeric as checks).
7. Independent verification (section 5).
8. RC1/RC2/RC3 gates (section 4).
9. Update ledger and hypothesis register.
10. Decide: continue, revise, reject, stop.
Do not reverse-engineer a hypothesis from a result you did not pre-register. Do not rename an unchanged hypothesis to evade budgets.

## 4. RC-CONT-01 gates
Use the exact RC1/RC2/RC3 definitions from Phys Research Workflow Protocol v1.0 (not any abbreviation). Minimum content per phase:
- RC1 landscape: structures, physical meaning, assumptions, regime, excluded phenomena, whether each quantity is used for what it actually describes.
- RC2 intersection: required structures, compatibility of assumptions, what the intersection establishes, removal/alteration tests. No minimality claim without a removal argument.
- RC3 interoperability, per mapping: source, target, mapping, assumptions, preserved, lost, regime, proof obligations, failure conditions, evidence. Shared notation or analogous equations are not interoperability.
Each gate returns PASS / CONDITIONAL / BLOCKED and must change a decision when it is not PASS.
Stop-or-carry: minor = record; material uncertainty = carry, dependents conditional; load-bearing unresolved premise = stop dependents; confirmed contradiction = block branch; independent branches continue.

## 5. Verification and roles
- Generator: hypotheses, prior art, derivations, candidate claims.
- Verifier: receives ONLY claim, definitions, assumptions, sources, artifacts. Not the generator's reasoning or preferred conclusion. Reconstructs independently and tries to break it.
- Adversary: is it novel, meaningful, consistent with competing explanations?
- Prefer a different model family for the verifier. The harness, not an agent report, executes numerical checks.
- Same-model or same-context review must be labeled "self-check" and cannot promote a claim to [R].
- Calibration gate: the verifier setup must pass `calibration/` (catch all known-bad claims, accept the controls) before the real run starts. Failure means the harness is not ready.
- Human expert scrutiny is required before anything is called a new theorem or N2+ result.

## 6. Known failure patterns (guard explicitly)
Compensating factors of 2; sign conventions contradicting their own cross-check; wrong time section (fixed-sigma+ is a null direction); physical vs dimensionless mass; "confirmed" tags on unchecked items; imports presented as derivations; circular premises; float-cancellation saddles (use log variables or high precision, check stationarity residuals and branch identity); approximations used outside their regime; entropy derivative treated as relative entropy or Fisher information.
Numerical rules: several parameter points and starting guesses, asymptotic cross-check, convergence is not proof.

## 7. Budgets, stagnation, terminal outcomes
Budgets are set in the brief and enforced by the runtime. Defaults: at most 7 initial hypotheses; one adversarial pass plus one targeted correction pass per promoted candidate; hard overall wall-clock/token cap; stagnation = 3 consecutive cycles with no new verified claim and no retired hypothesis.
Terminal outcomes (you decide when): SUCCESS, PROVEN NEGATIVE (a precise obstruction, not "literature is thin"), FALSIFIED, KNOWN RESULT, INSUFFICIENT INFORMATION, BUDGET EXHAUSTED, STAGNATION. A failed first hypothesis does not end the work while other viable hypotheses remain inside budget and scope.

## 8. Escalation (only these)
Ambiguity that changes the physical problem or deliverable; a governing premise contradicted; a missing resource or permission; a needed change to model, constitution or scope. Provide the conflict, affected claims, alternatives, recommendation. Continue unaffected work. No routine phase approvals.

## 9. Final record
Established / verified / imported / conditional / unverified; surviving and rejected hypotheses; strongest proved statement and its regime; prior-art and novelty (N0–N5; no N2+ without demonstrated novelty); RC findings; failed approaches; why stopped; what would change the conclusion; single highest-impact next question. A theorem candidate without an acceptable proof is a conjecture.
