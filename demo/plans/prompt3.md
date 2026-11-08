# AGENT C — QSG-T001
## Independent Adversarial Audit of Renormalized Stress-Energy Conservation

**Role:** Independent Mathematical Reviewer  
**Programme:** Quantum Fields on Classical Curved Spacetime  
**Priority:** Critical  
**Target:** QSG-T001 — Covariant Conservation of Hadamard-Renormalized Stress-Energy

---

## 1. Your responsibility

You are an independent mathematical physicist reviewing Agent B's research:

`Quantum Stress Energy Conservation Research.pdf`

Your task is to independently establish whether its central mathematical result is correct:

\[
\nabla^\mu
\langle\hat T_{\mu\nu}\rangle_{\mathrm{ren}}=0.
\]

You must reconstruct the mathematical argument from primary sources and first principles.

Do not accept Agent B's DERIVED classifications.

Do not assume its equations, intermediate steps, citations or conclusions are correct.

Do not try to make its derivation work.

Identify exactly what is established, what is conditional, what is incorrect and what remains unresolved.

You are authorized to reject the proposed theorem or substantially narrow its scope.

A demonstrated mathematical error or counterexample is a successful research outcome.

---

## 2. Research isolation

Work exclusively inside:

`research/experiments/qsg-preflight-001/agents/agent-c-conservation-audit/`

Use a separate working directory.

Do not modify:

- Agent B's PDF or deliverables.
- Agent A's infrastructure deliverables.
- PhysMath.
- PhysLib.
- PhysArtifact.
- Phys-Pseudo.
- Existing GR theorem implementations.
- Canonical QSG or BHC research records.

You may inspect the supplied PDF and independently access primary research sources.

Do not rely on previous reviewers' conclusions. Your audit must stand independently.

---

## 3. Frozen mathematical model

Use the following initial conventions unless a demonstrated inconsistency requires a documented correction.

Spacetime:

- Four-dimensional.
- Smooth Lorentzian manifold.
- Globally hyperbolic.
- Metric signature \((-+++)\).
- Natural units \(c=\hbar=1\).

Matter:

- Real scalar field.
- Mass \(m>0\).
- Minimal curvature coupling \(\xi=0\).

Classical action:

\[
S_m[\phi,g]
=
-\frac12\int d^4x\sqrt{-g}
\left[
g^{\mu\nu}\nabla_\mu\phi\nabla_\nu\phi
+
m^2\phi^2
\right].
\]

Classical field equation:

\[
(\Box_g-m^2)\phi=0.
\]

Classical stress-energy tensor:

\[
T_{\mu\nu}
=
\nabla_\mu\phi\nabla_\nu\phi
-\frac12g_{\mu\nu}
\left[
\nabla_\alpha\phi\nabla^\alpha\phi
+
m^2\phi^2
\right].
\]

The expected off-shell identity is:

\[
\nabla^\mu T_{\mu\nu}
=
(\Box_g\phi-m^2\phi)\nabla_\nu\phi.
\]

Verify this independently before accepting it.

The quantum state is an appropriately specified Hadamard state of the free scalar quantum field on the fixed classical background.

Do not introduce semiclassical gravitational backreaction into the initial conservation proof. The background is fixed for this investigation.

---

## 4. Audit A — Classical foundation

Independently derive:

1. Field equation from the action.
2. Stress-energy tensor through metric variation.
3. Covariant divergence.
4. Off-shell Noether identity.
5. On-shell conservation.

Check:

- Metric signature.
- Signs.
- Index placement.
- Mass dimensions.
- Variation conventions.
- Boundary assumptions.

Pay particular attention to the distinction between:

\[
\nabla^\mu T_{\mu\nu}=0
\]

as a dynamical on-shell identity and the off-shell identity obtained from diffeomorphism invariance.

Determine whether Agent B's classical reconstruction is correct.

---

## 5. Audit B — Quantum field and Hadamard foundations

Independently reconstruct the required quantum-field foundations.

Verify:

1. Operator-valued distribution formulation.
2. Canonical commutation relations.
3. Two-point Wightman function.
4. Hadamard state definition.
5. Microlocal spectrum condition.
6. Local Hadamard parametrix.
7. Smooth remainder after subtraction.

Determine precisely what is guaranteed by the Hadamard condition.

Distinguish:

- Smoothness of the subtracted two-point function.
- Satisfaction of the exact Klein-Gordon equation.
- Satisfaction of a Klein-Gordon equation modulo smooth local terms.
- Distributional identities.
- Coincidence-limit identities.

Do not assume that the Hadamard parametrix is an exact bisolution.

Inspect every use of this distinction in Agent B's conservation argument.

---

## 6. Audit C — Point-splitting renormalization

This is the central assignment.

Independently reconstruct the Hadamard point-splitting prescription for the minimally coupled massive scalar field in four dimensions.

Define:

- The two-point function \(G^+(x,x')\).
- The Hadamard parametrix \(H(x,x')\).
- The smooth remainder \(W(x,x')\).
- The appropriate bi-differential stress-energy operator.
- The parallel-transport structures required for tensorial coincidence limits.
- The finite local geometric correction.

Establish the precise expression:

\[
\langle\hat T_{\mu\nu}\rangle_{\mathrm{ren}}
=
\lim_{x'\rightarrow x}
D_{\mu\nu}(x,x')W(x,x')
+
C_{\mu\nu}(x).
\]

This is schematic notation. Define the actual prescription and conventions.

### Required calculations

1. Apply the bi-differential operator before taking the coincidence limit.
2. Calculate its covariant divergence.
3. Identify the residual local geometric contribution.
4. Derive the necessary finite geometric correction.
5. Establish whether the corrected expression is conserved.
6. Verify the relevant coincidence limits.
7. Identify all finite renormalization ambiguities.

Do not simply state that the correction tensor restores conservation.

Calculate its divergence and demonstrate the cancellation.

### Critical primary-source comparison

Reconstruct and compare:

- Wald's 1977 and 1978 work on renormalized stress-energy.
- Décanini and Folacci (2008), particularly their explicit point-splitting formulas.
- Hollands and Wald's work on local covariance and renormalized quantum fields.

Identify the exact equations and assumptions relevant to this particular scalar-field model.

If the papers use different prescriptions or conventions, reconcile them explicitly.

Do not combine equations from different prescriptions without establishing their mathematical relationship.

---

## 7. Audit D — Conservation theorem

The precise proposition under investigation is:

**For a real, minimally coupled, massive scalar quantum field in a Hadamard state on a smooth, globally hyperbolic four-dimensional Lorentzian spacetime, the appropriately renormalized stress-energy tensor is covariantly conserved.**

Determine whether this statement is correct as written.

If not, formulate the strongest precisely justified version.

Establish:

1. Required state regularity.
2. Required field equation.
3. Required renormalization prescription.
4. Required local geometric counterterms.
5. Conservation identity.
6. Finite renormalization ambiguities.
7. Any assumptions concerning locality and covariance.

Pay particular attention to whether the conservation result is:

- A universal identity for the specified prescription.
- A state-dependent identity.
- Conditional on a local covariance requirement.
- Dependent on a particular choice of finite geometric terms.

Do not conflate a trace anomaly with a failure of covariant conservation.

Verify the distinction explicitly.

---

## 8. Audit E — Finite renormalization ambiguity

Independently determine the permitted finite geometric ambiguity in the renormalized stress-energy tensor.

For four-dimensional spacetime, investigate the relevant local conserved curvature tensors and their relationship to variations of curvature-squared actions.

Determine:

- Which terms are allowed.
- Their mass dimensions.
- Which gravitational couplings they redefine.
- Whether each ambiguity preserves conservation.
- Whether the subtraction scale changes the physical coupling conventions.

Check Agent B's parameter count and dimensional assignments.

Do not call coefficients dimensionless unless their definitions and units justify that statement.

---

## 9. Audit F — Adversarial examination of Agent B's PDF

Inspect the entire supplied report, but concentrate on:

- Pages 4–6: classical action and conservation.
- Pages 7–11: quantum field and renormalization construction.
- Pages 12–13: main conservation derivation.
- Pages 14–16: backreaction and existence claims.
- Pages 17–20: MRC classifications and Phys-Pseudo theorem proposal.

For every consequential claim, identify:

1. Exact location in the PDF.
2. Agent B's assertion.
3. Independent mathematical result.
4. Missing or incorrect intermediate step.
5. Strongest correction.
6. Final disposition.

Explicitly audit:

- Whether the subtracted two-point function is treated correctly.
- Whether the point-split divergence is calculated correctly.
- Whether the local curvature correction is derived or merely asserted.
- Whether the final stress-energy tensor is actually conserved.
- Whether Wald's axioms and finite ambiguities are accurately described.
- Whether trace-anomaly statements are consistent.
- Whether the proposed theorem scope exceeds what the calculations establish.

Do not expand this into a review of every literature claim unrelated to the conservation theorem.

---

## 10. Excluded research scope

Do not attempt to resolve:

- General semiclassical gravitational backreaction.
- Global existence or uniqueness of cosmological solutions.
- Black-hole information.
- Quantum measurement.
- Cosmological constant naturalness.
- Stochastic gravity.
- Bohmian mechanics.

You may identify mathematical dependencies on these subjects, but record them as separate obligations.

The only central deliverable is a correct, independent conservation derivation.

---

## 11. Math Reality Check

For every consequential mathematical step, check:

| Check | Requirement |
|---|---|
| Definitions | Every operator, distribution and tensor defined |
| Sign conventions | Consistent with \((-+++)\) |
| Dimensions | Correct engineering dimensions |
| Tensor structure | Correct indices and transformations |
| State conditions | Hadamard requirements established |
| Renormalization | Explicit subtraction and coincidence prescription |
| Conservation | Actual divergence calculated |
| Local correction | Its divergence explicitly verified |
| Ambiguities | Correct finite geometric terms |
| Trace anomaly | Distinguished from conservation |
| Applicability | Scope and limitations explicitly stated |

Use:

- DERIVED.
- SOURCE-REPORTED.
- INDEPENDENTLY REPRODUCED.
- CONDITIONAL.
- FALSIFIED.
- BLOCKED.

Do not use PASS as a substitute for an explicit mathematical disposition.

---

## 12. Phys-Pseudo

Use Phys-Pseudo only for the final concise representation of the reviewed proposition.

Follow the actual repository `PSEUDO.md`.

Use exactly its seven permitted constructs:

- `given`
- `define`
- `assume`
- `claim`
- `derive`
- `if/then/else`
- `?`

Do not invent status constructs.

Do not claim that writing the Phys-Pseudo representation proves the theorem.

If the authoritative specification is inaccessible, report BLOCKED and preserve the mathematical derivation in conventional notation.

---

## 13. Required deliverables

Produce only these documents in your assigned folder.

### 01-primary-source-verification.md

Verified primary bibliography, exact equations, pages, conventions and source-supported claims.

### 02-classical-independent-derivation.md

Independent derivation of the scalar-field action, stress-energy and off-shell conservation identity.

### 03-hadamard-point-splitting-audit.md

Complete reconstruction of the subtraction prescription and the appropriate coincidence-limit machinery.

### 04-conservation-and-counterterm-derivation.md

Explicit divergence calculation, residual local geometric terms, correction and finite ambiguities.

### 05-agent-b-findings.md

A concise defect-by-defect assessment of the supplied PDF, with locations and corrections.

### 06-theorem-contract.md

The strongest mathematically defensible QSG-T001 statement, exact assumptions, proof obligations and exclusions.

### 07-independent-handoff.md

Short final disposition answering:

1. Is Agent B's conservation derivation correct?
2. What must be corrected?
3. What is the strongest theorem we can responsibly state?
4. Which claims are independently reproduced?
5. Which assumptions remain conditional?
6. Is the theorem ready for formalization through the GR-style protocol?

Avoid unnecessary historical exposition.

Every section must support a mathematical decision.

---

## 14. Stop conditions

Stop a derivation and identify the blocker if:

- A required coincidence limit cannot be justified.
- A renormalization prescription is insufficiently specified.
- A curvature correction cannot be derived.
- A cited source does not support the claimed result.
- Conventions cannot be reconciled.
- A universal claim is supported only by a restricted calculation.

Do not invent a missing correction tensor.

Do not repair Agent B's derivation silently.

Do not modify its PDF or source documents.

Do not begin theorem implementation.

---

## 15. Definition of success

Your assignment is complete when an independent researcher can use your derivation to determine exactly why:

\[
\nabla^\mu
\langle\hat T_{\mu\nu}\rangle_{\mathrm{ren}}
=0
\]

holds under the declared assumptions, including the role of the local geometric correction and finite renormalization ambiguities.

A finding that Agent B's theorem statement needs narrowing is acceptable.

A finding that the derivation contains a fundamental error is acceptable.

An independently reconstructed established theorem is a successful result.

**The objective is mathematical verification, not agreement with Agent B.**
