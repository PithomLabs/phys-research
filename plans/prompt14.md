# Adversarial Review: D-004 — Corrected Einstein–Cartan Specialization of Yo–Nester Classification

## 1. Objective

Independently review `derivations/D-004.md`, which supersedes D-003 following findings F1–F3 in R-004.

Determine whether:

1. The previous errors were actually corrected.
2. The EC-corner substitutions and resulting primary if-constraints (PICs) are mathematically correct.
3. The nine linear second-class pair classifications follow from Yo–Nester's equations.
4. The exceptional \(2^-\) pair is correctly characterized as unresolved.
5. The conditional \(\bar H_\perp\) conclusion follows from the available premises.
6. D-004 accurately distinguishes source-reported results, independent calculations, assumptions and unresolved questions.

This is an independent mathematical review, not a document-quality review.

**Do not assume D-004 is correct because it addresses the previous review findings. Reconstruct every consequential calculation independently.**

## 2. Review scope and source hierarchy

Primary target:

- `derivations/D-004.md`

Historical context:

- `derivations/D-003.md`
- `research/reviews/R-004.md`
- `derivations/D-002.md`
- `research/reviews/R-003.md`

Primary mathematical source:

- Yo–Nester, `gr-qc/0112030`, particularly Eqs. (36)–(39), (43)–(48), Table 1, and Eqs. (93)–(94).

Additional source:

- P-007, for the classical Einstein–Cartan baseline.

Inspect the actual cached full texts and PDF wherever equation rendering, coefficients or signs matter. In particular, visually inspect the rendered source equation rather than trusting potentially defective text extraction.

Use exact equation numbers, page numbers and source locators.

Keep inherited D-002 assumptions explicitly conditional. Do not silently promote the EC-corner mapping to a source-established fact.

## 3. Reassess previous findings F1–F3

Review each previous finding separately.

### F1 — Adjacent spin-\(2^-\) comparison

D-003 previously claimed that the spin-\(2^-\) PIC list in Eq. (94) equalled the EC-corner list minus \(T\phi\).

D-004 now additionally identifies the absence of \(\widetilde{\phi}_{\mu\nu}^{\perp}\).

Independently:

1. Reconstruct the complete EC-corner PIC list.
2. Reconstruct the complete spin-\(2^-\) PIC list from Eq. (94).
3. Compare every entry, including homogeneous and inhomogeneous terms.
4. Verify the conditions involving \(b_1+b_2\) and \(b_1+b_4\).
5. Verify the surviving inhomogeneous \(\phi^\perp\) term.
6. Check whether any additional entry or coefficient was overlooked.

Determine whether F1 is fully corrected.

Do not infer equivalence between different parameter points merely because some equations resemble one another. Treat this as a consistency comparison, not a transfer of classification.

### F2 — Eq. (46) constant

D-003's constant was corrected using the rendered PDF.

Independently inspect Eq. (46), including its exact coefficient, sign, normalization and powers of \(l\).

Verify the resulting EC specialization:

\[
\phi^\perp =
\frac{\pi^\perp}{J}
-\frac{3a_0}{l^2}
\approx 0.
\]

Check the transcription directly against the PDF and surrounding conventions.

Determine whether F2 is resolved.

### F3 — Unused premise

Verify that the unused `given` premise from D-003 was removed from D-004.

Check that the remaining premises are actually used by the derivations or explicitly serve as contextual source information.

Do not introduce unnecessary changes merely for stylistic consistency.

## 4. Independently reproduce the EC-corner substitutions

Starting from Yo–Nester's original equations, independently substitute:

\[
a_0\neq 0,\qquad
a_1=a_2=a_3=b_1=b_2=b_3=b_4=b_5=b_6=0.
\]

Do not copy D-004's resulting expressions.

### A. Torsion PICs

Recalculate Eqs. (36)–(39).

Verify every surviving term, coefficient and homogeneous condition.

### B. Curvature PICs

Recalculate Eqs. (43)–(48).

Verify:

- Every homogeneous PIC.
- The sole claimed inhomogeneous PIC.
- The PDF-confirmed constant in Eq. (46).
- All signs, indices and normalization factors.

Produce the complete enumerated list, rather than merely confirming that the equations simplify.

### C. Kinetic critical combinations

Independently evaluate every Table-1 kinetic combination at the EC parameter point.

Show the original combination, substitution and resulting value.

Confirm whether all ten combinations actually vanish.

Do not confuse kinetic degeneracy with a completed classification of the resulting constraints.

## 5. Independently review the pair classification

D-004 claims nine pairs are second-class at linear level and that the exceptional \(2^-\) pair has a vanishing linear mass term.

Reproduce:

1. Every PIC-to-\(\chi\) partner correspondence in Table 1.
2. Every corresponding mass combination.
3. Every substitution at the EC corner.
4. The resulting nine nonzero linear-level pair brackets.
5. The exceptional \(T\phi/T\chi\) linear-level bracket.

Then distinguish carefully between:

- A nonzero linearized Poisson bracket.
- A nondegenerate constraint matrix.
- A genuinely second-class constraint pair.
- A nonlinear constraint pair whose classification depends on phase-space variables.

Do not conclude that nine complete second-class pairs have been established unless the relevant constraint-matrix conditions justify that conclusion.

For the exceptional \(2^-\) pair, determine precisely what follows from the vanishing linear mass combination and what remains unknown.

Consider the possibility of:
- Nonzero nonlinear Poisson-bracket terms.
- First-class behavior.
- Tertiary constraints.
- Phase-space-dependent classification.
- Further bifurcation.

Do not select among these possibilities without sufficient evidence.

## 6. Review the conditional Hbar conclusion

D-004 claims the Hbar branch follows if the exceptional pair is second-class.

Challenge this implication.

Specifically:

1. Does the presence of additional constraints justify replacing \(H_\perp\) with \(\bar H_\perp\) under Blagojević's stated rule?
2. Is the complete EC constraint inventory established?
3. Does the unresolved \(2^-\) pair prevent uniform elimination?
4. Are the additional \(C_P F C\) contributions known?
5. Is the resulting Hbar expression explicit or still conditional?
6. Are the inherited D-002 assumptions correctly propagated?

Distinguish:
- Established source rule.
- Independent EC-corner substitution.
- Conditional inference.
- Unresolved nonlinear classification.

The review must not turn a conditional Hbar result into an unconditional conclusion.

## 7. Review the EC-corner mapping itself

D-004 explicitly assumes that the Einstein–Cartan action corresponds to the Yo–Nester parameter point.

The source does not explicitly state this mapping in the examined material.

Evaluate whether:

- P-007's EC action supports the identification.
- Yo–Nester's coupling conventions match the conventions used in P-007.
- The absence of all other couplings is justified.
- Normalization, signs and parameter definitions are compatible.
- The mapping is source-established, independently derived or merely plausible.

If the mapping cannot be independently established, retain it as a conditional assumption.

Do not reject valid conditional calculations merely because their premise remains unverified.

## 8. Review Phys-Pseudo compliance

Inspect the canonical `PSEUDO.md` syntax and the D-004 investigation.

Verify:

- Named constructs are valid.
- Each `from [NAMES]` list refers to existing premises or intermediate results.
- Every consequential `derive` has an explicit result.
- Every calculation identifies its inputs and mathematical operation.
- Assumptions are explicitly labeled as assumptions.
- Source-transcribed equations are not presented as independently derived results.
- `?` obligations accurately capture unresolved work.
- Conditional branches do not silently promote their antecedents into established facts.

Flag any derivation that is syntactically valid but mathematically unsupported.

Also check the consistency between the formal Obligations section and the Phys-Pseudo obligations, particularly D2', D2prime and D2map.

## 9. Per-derivation adversarial dispositions

For each consequential derivation, provide an individual disposition:

- VERIFIED: independently reproduced within the explicitly stated scope.
- CONDITIONAL: reproduced only under identified assumptions.
- UNSUBSTANTIATED: insufficient inputs or calculation to establish the result.
- FALSIFIED: an explicit mathematical contradiction or counterexample is demonstrated.

Do not issue a global correct/incorrect verdict for D-004.

For every finding, provide:

1. Exact derivation name.
2. Relevant D-004 lines.
3. Source equation and locator.
4. Independent calculation or counterexample.
5. Mathematical consequence.
6. Minimum correction required.
7. Evidential boundary.

For conditional findings, identify exactly which premise controls the condition.

## 10. MRC requirements

Apply `docs/MRC_REVIEW_PROTOCOL.md`.

Explicitly evaluate:

- Mathematical well-formedness.
- Dimensional and index consistency.
- Constraint and phase-space compatibility.
- Applicability of the Yo–Nester parameter specialization.
- Linear versus nonlinear classification.
- Hidden assumptions.
- Limiting cases.
- Whether the stated conclusions exceed the evidence.

State the applicable regime: classical Hamiltonian Poincaré gauge theory at the proposed Einstein–Cartan parameter corner.

Do not infer empirical validity from mathematical consistency.

## 11. Review artifact

Create the next review record using the repository's existing conventions and `physctl new`.

Review records may target hypotheses only. Therefore:

- Target H-002.
- Pin D-004 as a material input.
- Explicitly state that the substantive review scope is D-004.
- Reference D-003 and R-004 as historical context.
- Record each consequential derivation's disposition and evidence.
- Do not imply that the review independently validates the entire H-002 hypothesis.

Use a unique run identifier and preserve all prior records.

## 12. Preservation and verification

This is a read-only review of D-004.

- Do not edit D-004.
- Do not edit D-003 or R-004.
- Do not modify D-002 or H-002.
- Do not modify existing research conclusions.
- Do not touch PhysMath, PhysLib, PhysArtifact or Phys.
- Do not commit or merge.

Before and after the review, verify existing research artifacts and hashes.

Run:
- `physctl validate`
- Relevant record/reference checks.
- Relevant tests if any tooling changes are unavoidable.

If a correction is required, record it in R-005 and recommend a new D-005 version. Do not repair D-004 yourself.

## 13. Final report

Provide:

### A. Executive disposition
PASS / PASS WITH LIMITATIONS / FAIL.

### B. F1–F3 resolution
Report whether each previous finding is genuinely resolved.

### C. Substitution audit
Show the complete independently reproduced EC PIC and kinetic-combination results.

### D. Pair-classification audit
Show all nine linear pairs and the exceptional \(2^-\) result, with the linear/nonlinear distinction explicit.

### E. Hbar assessment
State precisely what is established, conditional and unresolved.

### F. Per-derivation disposition table
Include every consequential `derive` in D-004.

### G. Remaining obligations
List the minimum unresolved mathematical work, without inventing additional tasks.

### H. Preservation and validation
Report exact commands, results and hash checks.

### I. Recommended next mathematical calculation
Name one highest-impact calculation that can proceed without assuming unresolved results.

## Review principle

The goal is not to make D-004 pass.

The goal is to determine exactly which EC-corner statements survive independent mathematical reconstruction, which remain conditional on the parameter mapping, and which require further Dirac–Bergmann analysis.

**Preserve valid conditional mathematics. Reject unsupported conclusions. Never manufacture a successful review.**
