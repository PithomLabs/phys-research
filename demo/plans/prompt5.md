# AGENT A — QSG-T001 NORMALIZATION REPAIR

## Mission

Repair the mathematical foundation of QSG-T001 following adversarial review of the completed audit.

Do not restart the research programme.

Do not rewrite the existing infrastructure reconnaissance.

Do not introduce new mathematical or Go infrastructure.

The objective is to establish a correctly normalized Hadamard subtraction prescription, reconstruct the conservation calculation and determine whether the previously reported correction coefficient survives.

## 1. Critical defect

The current definition is:

\[
G^{(1)}=G^++G^-,
\]

\[
H^{(1)}=\frac{1}{8\pi^2}
\left(\frac{U}{\sigma}+V\ln|\bar\sigma|\right),
\]

\[
S=\frac12(G^{(1)}-H^{(1)}).
\]

These definitions are inconsistent with the claimed smoothness of S.

For the stated normalization, the singular part of \(G^{(1)}\) is twice the singular part of \(H^{(1)}\).

The existing S therefore retains a singular contribution.

Do not accept P2 as established.

## 2. Resolve the normalization first

Determine the correct prescription using the primary-source definitions and the Minkowski two-point function.

Evaluate these alternatives:

A. Keep H^(1) with its existing 1/(8π²) prefactor and define:

\[
S=\frac12G^{(1)}-H^{(1)}.
\]

B. Keep the original S structure but redefine H^(1) with the normalization required to subtract the full symmetric two-point function.

Choose only after independently verifying the singular coefficients, Wightman normalization, symmetric-function normalization and Feynman correspondence.

Do not assume either alternative automatically validates the conservation coefficient.

## 3. Reconstruct the proof

Once the subtraction prescription is frozen:

1. Establish smoothness of the resulting remainder.
2. Re-derive its inhomogeneous Klein–Gordon equation.
3. State and derive the required Hadamard transport equations.
4. Establish the relevant coincidence limits.
5. Reconstruct the point-split divergence using Synge's rule.
6. Derive the residual geometric term.
7. Derive the correction from that residual.
8. Calculate the correction's divergence.
9. Demonstrate explicit cancellation.
10. Reconcile the result with Décanini–Folacci and Wald.

Do not reuse the old numerical coefficient without recalculation.

## 4. Audit the unresolved framework difference

Revisit the claim:

\[
W=8\pi^2S+A.
\]

Establish whether A is genuinely smooth under the corrected definitions.

Calculate its field-equation coincidence limit.

If it cannot be established, mark the comparison BLOCKED.

Do not use an unspecified smooth term to hide a normalization discrepancy.

## 5. Repair provenance and metadata

Keep the original audit documents and their locked sections intact.

Append explicit errata and correction records.

Reconcile O1–O5 across:

- `06-theorem-contract.md`
- `07-independent-handoff.md`
- `research_metadata.json`
- `manifest.json`

Distinguish theorem proof obligations from machine-verification limitations.

## 6. Implementation

Do not change the Go fixture until the corrected mathematical coefficients have been established.

Then update only the affected mathematical goldens, anti-goldens and metadata.

Keep universal conservation steps explicitly unsupported by the machine.

Retain the existing architecture and dependency boundaries.

## 7. Acceptance criteria

The task is complete only when:

- The symmetric two-point function and parametrix singularities match exactly.
- The subtracted remainder is genuinely smooth.
- The point-split coincidence limit is mathematically well-defined.
- The remainder equation is independently derived.
- The residual divergence is explicitly calculated.
- The correction coefficient is derived rather than imported.
- The correction divergence is calculated.
- The cancellation is demonstrated.
- The source correspondence is reconciled or explicitly BLOCKED.
- All theorem and fixture metadata agree.
- Tests demonstrate the actual limited capabilities claimed.

If a mathematical gate fails, report BLOCKED and do not promote the theorem.

**The goal is to repair the mathematics, not preserve the previous conclusion.**
