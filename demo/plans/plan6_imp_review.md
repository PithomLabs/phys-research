# Adversarial review: plan6 implementation report

Verdict: substantial progress, with one important status inconsistency to resolve before treating the independent-review gate as closed. The implementation report describes a genuine mathematical reconstruction, not merely a normalization patch. It also identifies a second defect in the earlier derivation that was not part of my original review: confusing \\(J=(\Box-m^2)S\\) with \\(J'=(\Box'-m^2)S\\). That is potentially the most consequential new finding.&#x20;

plan6_imp(1).md



## 1. The new coefficient is internally consistent

The report gives:

\\[ [J]=-\frac{3[v_1]}{4\pi^2} \\]

and the revised residual:

\\[ [\nabla'\_\nu J] =\frac13\nabla\_\nu[J] =-\frac{\nabla\_\nu[v_1]}{4\pi^2}. \\]

It then derives:

\\[ C\_{\mu\nu} =\frac{1}{4\pi^2}g\_{\mu\nu}[v_1], \\]

so that

\\[ \nabla^\mu C\_{\mu\nu} =\frac{\nabla\_\nu[v_1]}{4\pi^2}. \\]

Consequently,

\\[ \nabla^\mu(R\_{\mu\nu}+C\_{\mu\nu})=0. \\]

Algebraically consistent

This is a coherent replacement for the old \\(3/(16\pi^2)\\) result. Importantly, the coefficient is no longer being carried forward from the defective calculation. The reported reconstruction derives it from the repaired prescription and the revised Synge calculation.&#x20;

plan6_imp(1).md



But I am not yet independently certifying the \\(1/3\\) identity. The report summarizes the adversarial Synge-route result; it does not show the actual index-level calculation. That calculation is now the most important mathematical object to inspect.

In particular, I would verify:

- Exactly how \\(J\\) and \\(J'\\) are defined, including their differential operators and argument indices.
- How the primed derivative is transported to the coincidence limit.
- Every use of Synge's rule and the coincidence limits of the relevant bitensors.
- Why the earlier \\(1/2\\) calculation fails, rather than simply replacing it with \\(1/3\\).

The result may be correct, but the report alone does not demonstrate it.

## 2. The strongest new finding: the J/J′ distinction

This deserves special attention.

The earlier calculation apparently treated the field-equation source and the primed-argument source as interchangeable. The report now says the residual is determined by the latter and that the correction changes from the previous result accordingly.&#x20;

plan6_imp(1).md



This is exactly the kind of hidden mathematical error that our adversarial workflow should catch: two expressions that look equivalent before coincidence need not have identical differentiated coincidence limits.

I would make this a permanent methodological rule for Phys:

> Never exchange an unprimed and primed differential operator inside a point-split expression without explicitly deriving the relationship between their coincidence limits.

That principle is more valuable than the particular numerical correction. It can prevent similar errors in other bitensor calculations.

## 3. The C2 reversal is handled more carefully

The report now states:

\\[ W_B=S+\frac{i}{2}\Delta \\]

and distinguishes the smooth real part from the full complex distribution.&#x20;

plan6_imp(1).md



That is the right distinction to preserve.

A smooth real remainder does not establish that the entire complex difference is smooth. Nor does smoothness establish that the remainder satisfies the homogeneous Klein–Gordon equation. These must remain separate findings.

I would retain this correction in the research record as an explicit revision of the previous verdict, with the original finding preserved as history.

## 4. One important inconsistency: O1 closure versus independence

This is my principal objection to the completion report.

It says:

- O1 is closed in full through checker #1, executor cross-routes, and checker #2.
- The two checkers are from the same session, and replication is not fully independent.

plan6_imp(1).md

Those statements can coexist only if O1 means internal adversarial cross-checking, rather than independent replication.

The earlier research record treated O1 as an independent second derivation. Under that interpretation, two same-session checkers do not establish full closure.

I recommend splitting the status:

| Verification level                                                    | Recommended status                                                                                                |
| --------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------- |
| Independent derivation within the current execution                   | Closed, if the checker genuinely derived the result independently rather than reviewing the executor's derivation |
| Adversarial internal cross-check                                      | Closed, subject to inspection                                                                                     |
| Independent external replication                                      | Open                                                                                                              |
| Reproduction by a separate agent/session with no access to the answer | Open                                                                                                              |

This does not invalidate the derived result. It prevents the project from claiming a stronger verification standard than it has met.

## 5. The remaining open questions are appropriately disclosed

The report identifies four important qualifications:

| Item                               | Status          | My interpretation                                            |
| ---------------------------------- | --------------- | ------------------------------------------------------------ |
| Explicit \\(A\\)-term              | Open (O2)       | DF framework identification must remain qualified            |
| Nonlinear \\(H^{(2)}\\) line-check | Open (O3)       | Do not claim that calculation is fully verified              |
| Wald source access                 | Open (O4)       | Do not claim detailed equation-by-equation Wald verification |
| F2-weak                            | Source-reported | Keep its evidential status visible                           |

plan6_imp(1).md

The O2 qualification is particularly important. The report says the coefficient agrees with DF by derivation, but explicit \\(A\\) remains open. I would phrase the result as:

The conservation-restoring coefficient has been independently derived within the repaired prescription and numerically agrees with the DF coefficient under the stated convention mapping. Full identification of the prescriptions remains open pending the explicit \\(A\\)-term audit.

That is more precise than saying the entire framework has been shown equivalent to DF.

## 6. What the implementation tests establish—and what they do not

The engineering work appears improved:

- Fixed byte snapshots.
- Unknown-field injection tests.
- More accurate descriptions of mutation tests.
- Updated boundary statements.
- Build, test, vet, formatting, and JSON validation all reported successful.

plan6_imp(1).md

Those are meaningful improvements. But the tests verify the integrity and behavior of the fixture, not the physical theorem. The report does acknowledge this distinction, and it should remain explicit in the manifest.

Also, because the working directory has no Git repository, the hash manifest is the modification record. That is useful, but it is not equivalent to a reviewable Git history with independently inspectable diffs. Preserve the hash evidence and the byte-identical v1 files.

## 7. My recommended next action

I would not restart the research or reopen the entire normalization exercise. The next step should be a narrow mathematical audit of the new derivation.

The minimum evidence needed is:

1. `08-normalization-repair-record.md`, especially N2, N3, N4, and the checker rebuttal.
2. The revised conservation derivation in doc 04, including the exact \\(J/J'\\) treatment and all Synge steps.
3. The revised theorem contract in doc 06, including the complete premise list.
4. The revised test file and manifest, so the engineering claims can be checked against the actual implementation.

The most important question is not whether \\(1/(4\pi^2)\\) matches DF. It is whether the repaired bitensor derivation genuinely forces that value without importing DF's answer.

Final assessment: This is a promising and materially improved result. The original normalization defect has been addressed, and the newly discovered \\(J/J'\\) distinction may have corrected a second independent mathematical error. I would provisionally accept the repaired theorem as a derived human result under its declared prescription, while keeping independent replication and the explicit \\(A\\)-term correspondence open. I would not yet call the full mathematical audit independently closed.
