# Adversarial review of `plan6`

Verdict: approve the direction, but amend the plan before execution. It correctly treats the normalization defect as real, withdraws the old smoothness claim, and makes reconstruction—not preservation of the old coefficient—the basis for proceeding. It also keeps the existing programme and fixture architecture intact.&#x20;

plan6(6).md



The plan is substantially better than simply patching the factor of two. I see four important amendments and a few execution details to tighten.

## 1. Make Step 1 a complete distribution-level normalization audit

The proposed prescription,

\\[ S\_{\mathrm{new}}=\frac12G^{(1)}-H^{(1)}, \\]

is the right candidate under the conventions previously documented. It matches the leading singularity if \\(H^{(1)}\\) has the standard \\(1/(8\pi^2)\\) normalization.

But matching the leading \\(1/\sigma\\) coefficient is necessary, not sufficient. Step 1 should explicitly verify:

- The sign and coefficient under the declared metric signature and \\(\sigma\\) convention.
- The complete singular structure, including the \\(U/\sigma\\) and \\(V\ln|\bar\sigma|\\) terms.
- The distributional prescriptions: Wightman boundary values, their symmetric combination, and the Feynman parametrix.
- Which statements concern the real/symmetric part and which concern the full complex distribution.

This distinction matters for the anticipated Agent B reassessment. If \\(H^{(1)}\\) is the real symmetric parametrix, then \\(G^+-H^{(1)}\\) may have a smooth real part while retaining a distributional imaginary part. That is not equivalent to proving that the full complex difference is smooth.

Suggested addition: Require a convention table with the exact definitions and relations among \\(G^+\\), \\(G^-\\), \\(G^{(1)}\\), \\(H^{(1)}\\), and the Feynman object. Do not allow the same symbol \\(H\\) to silently mean different parametrix objects.

## 2. Resolve A-versus-B algebraically, not informally

The plan says the two alternatives appear notationally equivalent. That is correct only under a precise condition:

\\[ S_A=\frac12G^{(1)}-H_A^{(1)} \\]

and

\\[ S_B=\frac12\left(G^{(1)}-H_B^{(1)}\right) \\]

are identical precisely when

\\[ H_B^{(1)}=2H_A^{(1)}. \\]

That factor must apply to the entire parametrix, not merely the leading \\(1/\sigma\\) coefficient. The logarithmic term and all other relevant normalization conventions must follow it.

I agree with choosing Alternative A if the primary-source verification confirms the standard \\(H^{(1)}\\) normalization. It is the smaller change to the existing definitions. But record the equivalence as an algebraic identity after verifying the conventions, rather than as an expectation.&#x20;

plan6(6).md



## 3. Remove the anticipated C2 reversal

This is the clearest methodological weakness.

The plan says that C2's reversal is anticipated and that C3–C7 and D–J are expected to be unaffected. Although it says these must be verified, those expectations can bias the audit toward preserving the anticipated result.&#x20;

plan6(6).md

plan6(6).md



I recommend replacing that language with:

> All prior Agent B and Agent A findings are provisional wherever they depend directly or indirectly on the defective definition of S. No finding is presumed to reverse or remain unchanged. Re-derive the relevant claims from the corrected definitions, then document the result and its dependency chain.

In particular, distinguish these questions:

1. Does Agent B's subtraction correctly remove the singularity?
2. Is the resulting real part smooth?
3. Is the full complex remainder smooth?
4. Does the remainder satisfy the claimed homogeneous or inhomogeneous field equation?

They are separate claims. A positive answer to one does not establish the others.

## 4. State the correction's uniqueness limits

Step 3 correctly requires the correction to be derived from the residual rather than imported.&#x20;

plan6(6).md



One mathematical qualification needs to be explicit: a calculated divergence does not, by itself, uniquely determine a tensor.

For example, finding

\\[ \nabla^\mu R\_{\mu\nu}=-a\nabla\_\nu F \\]

motivates the metric-proportional correction

\\[ C\_{\mu\nu}=a g\_{\mu\nu}F. \\]

That demonstrates cancellation, but uniqueness requires additional constraints. One can generally add an identically conserved tensor without changing the divergence.

The derivation should therefore distinguish:

- The minimal local metric-proportional conservation-restoring term.
- The freedom to add separately conserved local curvature terms.
- The assumptions under which the complete renormalization ambiguity is restricted.

Do not describe the correction as uniquely determined unless those additional assumptions establish uniqueness.

## Other execution amendments

| Area                 | Required tightening                                                                                                                                                                                                       |
| -------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Step 2: smoothness   | Explicitly state the Hadamard-state premise and show that the complete singular structure cancels. Derive the transport equations and the order needed for the point-split derivatives to have finite coincidence limits. |
| Source \\(J\\)       | State the exact differential operator, which argument it acts on, and the sign convention connecting the field equation to \\(J\\). Do not import the old source coefficient.                                             |
| O1–O5                | Define the exact significance of O1. The plan specifies O2 and O5, but O1's effect on theorem status still needs an explicit decision.                                                                                    |
| Errata and revisions | Clarify which original files or sections remain byte-identical and which receive revised versions. “Append errata” and “rewrite 06” should not leave ambiguity about what is protected.                                   |
| Build validation     | Include the exact commands, working directory, and `GOWORK`/`GOCACHE` environment recipe, rather than referring to a recipe elsewhere.                                                                                    |

The test plan is also moving in the right direction: fixed snapshots are stronger than comparing two builds of the same code, and injecting unknown fields tests strict decoding.&#x20;

plan6(6).md

plan6(6).md



One important boundary: a test that catches a changed coefficient in a manifest or golden payload proves that the stored value changed. It does not prove that the physical coefficient is mathematically correct. The fixture must continue to distinguish data integrity and arithmetic replay from verification of QSG-T001.

## Final assessment



Sound recovery strategy

Amend before execution



The most important feature of plan6 is that it no longer treats the old \\(3/(16\pi^2)\\) coefficient as an established result. It makes the normalization check the load-bearing gate, requires an independent reconstruction, and allows a mathematically honest `BLOCKED` outcome.&#x20;

plan6(6).md

plan6(6).md



My recommendation is to make the four amendments above, especially the distribution-level normalization audit and the neutral treatment of C2. Then execute the plan in its existing order.

QSG-T001 remains mathematically unpromoted. Neither the corrected prescription nor the expected cancellation has been independently established by this plan itself; those are the results Agent A still needs to derive.
