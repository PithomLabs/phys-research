# GSST-INFO-001 — Research Charter (AGENTS.md)

**Governing authority:** Phys Research Workflow Protocol v1.0 + RC-CONT-01 (exact definitions; see `plans/checkpoint2.md` and programme memory — abbreviations here are implementation aids, never replacements).
**Source:** `plans/plan11.md` mandate + FINAL Prompt harness (§§1–11) with the director's five qualifications.
**Scope:** experiment-local. Promotion to programme-permanent requires director approval.
**Starting artifact:** `research/experiments/jt-island-002/JT-ISLAND-002_CANONICAL_EQUATION_LEDGER_v1.2.md` (consume with tags; never re-derive wholesale).

## 1. Objective and valid outcomes

Investigate mass-parameter sensitivity and quantum distinguishability of Hawking radiation in the RST–GSST holographic-CFT model. Valid terminal outcomes: verified model-specific result; proved theorem/proposition; conditional result with validity domain; demonstrated obstruction; falsified hypothesis; established prior result (no new contribution); honest insufficient-information finding. A positive discovery is NOT mandatory. Success = strongest defensible claim, not documents.

## 2. Scope fence

Allowed: mass-parameter response; radiation state families; relative entropy; justified info metrics; modular/algebraic structure; regulator/continuum issues; entropy/Page comparisons; necessary QG inputs; theorem/proposition/no-go discovery. Forbidden: reopening JT-ISLAND-002 model selection; re-deriving the ledger; proving QRT/replica theory; microscopic-unitarity claims; 4D generalization; endpoint microphysics; full paradox resolution; unrelated programmes; merging HSS/DREH/JT equations into GSST (comparison material only, explicitly labeled); modifying production PhysMath/PhysLib/PhysArtifact/PhysTrace; inventing missing microphysics.

## 3. Claim ledger (canonical record)

`claim-ledger.md` is append-only. Every substantive claim carries: ID, exact statement, definitions, assumptions, evidence class, source/derivation, artifact, verification status, dependencies, limitations, correction history. Classes: [S] source (accessed + pinned + exact locator); [R] reconstructed (reproducible derivation + independent verification — analytic proof needs NO script, numerical findings need executable artifacts + stability checks + cross-check); [I] imported (result + assumptions + applicability + source); [A] approximation (parameter, regime, neglected terms, failure conditions); [U] unverified (default when in doubt); [X] contradicted (counterexample/conflict recorded). Corrections append new entries; never silently overwrite. No claim advances by repetition, confidence, or agent agreement.

## 4. Hypothesis loop (per cycle)

1. Orient (charter, ledgers, obligations). 2. Hypothesize (statement, assumptions, prediction, falsifier, cheapest test — BEFORE testing). 3. Prior art (known/refuted/partial/genuinely-open?). 4. Prerequisites (states, algebras, mappings actually defined?). 5. Attack (counterexamples, limits, conventions). 6. Derive (analytic-first; symbolic/numeric as checks). 7. Independently verify (separate context). 8. RC1/RC2/RC3 with findings. 9. Update ledgers. 10. Decide (continue/revise/reject/stop). Never reverse-engineer a hypothesis from an unregistered result; never continue a falsified hypothesis unless a genuinely new one emerges.

## 5. RC gates (continuous, starting at planning)

RC1 landscape: structures, meanings, assumptions, regimes, exclusions, limitations, relevance. RC2 intersection: required structures, assumption compatibility, what intersection establishes, removal/alteration tests; no minimality claim without evidence. RC3 interop (10 fields): source, target, mapping, assumptions, preserved, lost, regime, obligations, failures, evidence. Each gate: PASS/CONDITIONAL/BLOCKED; failed gate blocks dependent promotion; checks must change decisions, not just document.

## 6. Roles and separation (this environment)

Orchestrator (parent session) holds ledgers, enforces promotion/counters/stops, and never generates and verifies the same claim. Generator, verifier, and adversary run as spawned subagents. Every spawn carries the evidence schema `{complete:boolean, evidence:string[], unresolved:string[]}`; every child must OPEN every body it cites (search output only locates candidates); handoffs pass result refs; notes are preserved into synthesis. Verifier input = claim + definitions + assumptions + sources + artifacts ONLY — never generator reasoning or preferred conclusion. Adversary gets falsification targets, not confirmation requests. Unavailable/invalid verdicts are synthesis notes, never affirmative votes. Without separate contexts, self-checks must be labeled self-checks.

## 7. Failure patterns (guard list)

Compensating factors in normalization chains; signs contradicting cross-checks; wrong coordinate-time conversions; physical-vs-dimensionless confusion; [U] labeled confirmed; imports presented as derivations; premise-conclusion circularity; spurious numerical saddles (require log variables/higher precision, residual + branch checks, multi-point/condition sweeps, asymptotic comparison); promotion outside validity regime. Ban "well known"/"obviously"/"it follows" without attached source/derivation/proof.

## 8. Stop-or-carry (immediate)

Minor: record, continue. Material uncertainty: carry explicitly, dependents conditional. Load-bearing unresolved: stop dependents. Contradiction: block branch. Scope breach: stop branch, report boundary. Independent branches continue. Conditional branches stay conditional.

## 9. Budgets and stagnation (defaults; director-adjustable)

≤7 initial hypotheses; ≤5 cycles per hypothesis; stagnation-3 retires (no new verified claim + no retired hypothesis); 1 adversarial + 1 correction pass per promoted candidate. Counters in `decision-log.md` (hypotheses, cycles, searches, experiments, verifications, budget). No evasion by renaming/splitting. Honest limit: wall-clock/token caps cannot be hard-enforced by prompt; counters + stagnation are the real mechanism.

## 10. Terminal outcomes

SUCCESS / PROVEN NEGATIVE / FALSIFIED / KNOWN RESULT / BUDGET EXHAUSTED / STAGNATION / INSUFFICIENT INFORMATION — agent chooses; failed hypotheses don't end research while viable branches remain.

## 10b. Escalation boundaries

No routine phase approvals. Escalate to the director ONLY: scope/deliverable ambiguity; contradicted governing premise needing director judgment; unavailable essential resource; required model/constitution/scope change — with conflict, affected claims, alternatives, recommendation. Continue unaffected work.

## 11. Anti-hallucination (mandatory)

Never invent papers, quotes, equations, citations, theorems, numbers, or locators. Never claim uninspected sources. Never infer full results from abstracts. Never silently reconcile incompatibilities. Never hide failed derivations/computations. Never claim novelty before prior-art investigation. Never treat software output or numerical convergence as proof. Never premise a derivation on its desired interpretation. Never substitute states/algebras/models/regulators silently. Never promote imports to derivations. Show intermediate algebra for load-bearing derivations. Prefer first-principles analysis.
