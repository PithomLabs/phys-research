I reviewed the attached session. There is an important distinction between what is finished and what is still pending.

Agent A has already completed the infrastructure reconnaissance. Agent B has delivered the 24-page physics report, but its central conservation proof has not been established. The subsequent independent audit contains substantial research and identifies mathematical defects, but the session ends before the seven required audit documents are completed.&#x20;

session-c4a1a27d-d96d-4646-80c2-734cec802d77(1).md



### What Agent A should not repeat

| Area                      | Existing work                                                        | Instruction                              |
| ------------------------- | -------------------------------------------------------------------- | ---------------------------------------- |
| Repository reconnaissance | Eight Agent A reports and G1 closure work                            | Reuse; do not repeat                     |
| Phys-Pseudo               | Authoritative specification and implementation boundaries identified | Follow existing contract                 |
| GR implementation pattern | Existing `phys-lib/gr/equivariance` fixture inspected                | Reuse as reference                       |
| Agent B's research        | 24-page QSG report                                                   | Audit, don't reproduce                   |
| Conservation proof        | Independent analysis partially completed                             | Finish the mathematical derivation       |
| QSG-T001 formalization    | Not yet accepted                                                     | Proceed only after the mathematical gate |

### The most important problems inherited from Agent B

1. The central mathematical error: Agent B assumes that the smooth remainder \\(W=G^+-H\\) satisfies the homogeneous Klein–Gordon equation. Smoothness does not establish this. The Hadamard parametrix is generally not an exact solution.
2. The missing correction: The point-split stress-energy divergence contains a residual local geometric contribution. Agent B does not correctly derive its cancellation.
3. A logical contradiction: Agent B describes its correction tensor as conserved while simultaneously using its nonzero divergence to cancel another term.
4. Additional mathematical defects: The Hadamard coefficients, finite renormalization ambiguities, curvature tensors and trace-anomaly formula require independent verification.
5. Unsupported verification claims: Several literature results are labelled independently reproduced without corresponding derivations.
6. Formalization defects: The proposed Phys-Pseudo representation does not conform to the authoritative language specification.

The important distinction is that the conservation theorem may still be correct even though Agent B's proof is invalid. Agent A must establish that independently, not simply repair Agent B's text.

## Prompt for Agent A

# AGENT A — QSG-T001 RESUMPTION

## Independent Conservation Audit, Theorem Formalization and Minimal Implementation

**Role:** Senior Go Architect, Theoretical Physics Research Lead and Mathematical Verification Engineer

**Priority:** Critical

**Programme:** Quantum Fields on Classical Curved Spacetime

**Target:** QSG-T001 — Covariant Conservation of Hadamard-Renormalized Stress-Energy

## 1. Mission

Continue the existing QSG-T001 programme from the completed infrastructure reconnaissance.

Your objective is to take the work from Agent B, the existing Agent A infrastructure findings and the unfinished independent conservation audit, and bring the first QSG theorem to a mathematically defensible implementation state.

Do not restart the investigation.

Do not repeat Agent A's completed infrastructure research.

Do not reproduce Agent B's literature survey.

Do not assume Agent B's conclusions are correct.

The priority is to finish the missing mathematical verification, establish the exact theorem contract and implement the smallest representation supported by the existing architecture.

**Mathematical correctness takes priority over implementation completeness.**

A correctly identified obstruction is a successful result.

---

## 2. Existing workspace and authoritative materials

Workspace:

`/home/chaschel/Desktop/biz/TOE-v6/`

### Existing Agent A work — already completed

Read and reuse:

`research/experiments/qsg-preflight-001/agents/agent-a-infra-recon/`

Particularly:

- `01-repository-map.md`
- `02-phys-pseudo-contract.md`
- `03-gr-reference-analysis.md`
- `04-qsg-mathematical-feasibility.md`
- `05-tdd-research-workflow.md`
- `06-infrastructure-gaps.md`
- `07-implementation-contract.md`
- `08-handoff.md`
- `g1-closure-report.md`

These documents represent completed reconnaissance. Do not regenerate or rewrite them.

### Existing physics research

Agent B's source:

`plans/Quantum Stress Energy Conservation Research.pdf`

Previous independent review:

`plans/review3.md`

Original independent audit assignment:

`plans/prompt3.md`

Use these as the research record and requirements.

### In-progress independent audit

`research/experiments/qsg-preflight-001/agents/agent-c-conservation-audit/`

Inspect its current state before writing anything. The previous session found this directory empty, despite substantial mathematical investigation already having taken place.

Do not assume that the previous session's scratch findings are final mathematical results. Verify consequential equations independently.

### Authoritative implementation specifications

- `phys-system/AGENTS.md`
- `phys-lib/gr/equivariance/AGENTS.md`
- `phys-pseudo/PSEUDO.md`
- `phys-pseudo/AGENTS.md`
- `phys-pseudo/MRC_REVIEW_PROTOCOL.md`
- `phys-pseudo/conformance/DERIVATION_SPEC_GUIDELINES.md`

The actual repository specifications take precedence over examples, previous agent outputs and inferred conventions.

---

## 3. First priority: close the independent mathematical audit

Investigate only:

**Conservation of the renormalized stress-energy tensor for a free, real, minimally coupled, massive scalar field on a fixed classical curved spacetime.**

Frozen initial model:

- Four-dimensional smooth Lorentzian spacetime.
- Globally hyperbolic.
- Signature (-+++).
- Natural units c = ℏ = 1.
- Real scalar field.
- Mass m > 0.
- Minimal curvature coupling ξ = 0.
- Appropriately specified Hadamard quantum state.
- No gravitational backreaction in the conservation proof.

Target proposition:

[\
\nabla^\mu\langle\hat T\_{\mu\nu}\rangle\_{\mathrm{ren}}=0.\
]

### A. Classical foundation

Independently derive:

1. Euler-Lagrange equation from the action.
2. Stress-energy tensor from metric variation.
3. Covariant divergence.
4. Off-shell Noether identity.
5. On-shell conservation.

Verify the expected identity:

(\Box_g\phi-m^2\phi)\nabla\_\nu\phi.\
]

Check signs, conventions, dimensions and boundary assumptions.

Do not spend time rewriting Agent B's classical derivation if independent verification confirms it.

### B. Hadamard foundations

Establish the exact mathematical relationship between:

- State two-point function.
- Local Hadamard parametrix.
- Smooth remainder.
- Klein-Gordon operator.
- Coincidence-limit operations.

The central defect to investigate is Agent B's assertion:

[\
(\Box_g-m^2)W=0.\
]

Do not assume this follows from smoothness.

For W = G⁺ - H, explicitly establish:

-(\Box_g-m^2)H\
]

when G⁺ satisfies the field equation.

Calculate the actual smooth source term from the Hadamard construction.

Also check whether the derivative of a coincidence limit has been handled correctly. Do not silently replace a total derivative with differentiation in only one point.

Maintain consistent conventions for the Wightman and symmetric two-point functions.

### C. Point-splitting and local correction

This is the critical mathematical deliverable.

Independently establish:

1. The Hadamard subtraction prescription.
2. The bi-differential stress-energy operator.
3. Required parallel-transport structures.
4. Relevant coincidence limits.
5. The covariant divergence of the regularized tensor.
6. The residual local geometric term.
7. The finite correction required for conservation.
8. The divergence of that correction.
9. The explicit cancellation.

The previous audit identified a candidate correction proportional to:

[\
g\_{\mu\nu}v_1.\
]

Under the Décanini-Folacci conventions, check the coefficient and sign obtained from their general-dimensional formula.

Do not copy the candidate without reconciling the precise definitions of the Hadamard coefficients and two-point functions.

**Important:** Do not describe the conservation-restoring correction as independently conserved if its nonzero divergence is required to cancel the residual.

Distinguish this required prescription term from the freedom to add separately conserved finite geometric tensors.

### D. Primary-source reconciliation

Concentrate on:

- Wald (1977, 1978).
- Décanini and Folacci (2008).
- Radzikowski (1996).
- Hollands and Wald (2001, 2005), where relevant.

Prioritize the actual equations governing the scalar-field conservation argument.

Use exact equation and page references.

Verify bibliographic details.

If a primary paper is inaccessible, explicitly identify the limitation. Do not claim independent reproduction based on an abstract or secondary summary.

Reconcile all normalization and sign differences before combining results from different prescriptions.

### E. Finite ambiguities

Verify:

- Allowed local conserved curvature tensors.
- Their relationship to variations of curvature-dependent gravitational actions.
- Mass dimensions of their coefficients.
- Their effects on gravitational couplings.
- Dependence on the renormalization scale.
- Differences between massive and massless cases.

Check Agent B's expressions for H^(1) and H^(2) independently.

Do not reproduce incorrect curvature tensors merely because they appear in the PDF.

### F. Trace anomaly

Keep conservation and the trace anomaly separate.

Audit Agent B's formula and its mixing of minimal and conformal coupling.

A nonzero trace does not itself indicate non-conservation.

Do not turn this into a separate research programme.

---

## 4. Specific Agent B defects to investigate

The following are leads from the previous adversarial review, not accepted conclusions. Verify each one.

| Location    | Suspected defect                                                         |
| ----------- | ------------------------------------------------------------------------ |
| Pages 1–3   | Bibliographic errors and unsupported independent-verification labels     |
| Pages 10–11 | Hadamard normalization and coincidence coefficients                      |
| Page 11     | Incorrect finite ambiguity dimensions and curvature tensors              |
| Pages 12–13 | False homogeneous equation for W                                         |
| Page 13     | Undefined local correction and contradictory conservation claim          |
| Page 13     | Trace-anomaly formula and coupling mismatch                              |
| Pages 14–16 | Downstream backreaction claims relying on questionable curvature tensors |
| Pages 16–18 | Noise-kernel and validity claims exceeding their mathematical support    |
| Pages 19–20 | Phys-Pseudo non-conformance                                              |

For every consequential defect, provide:

1. Exact PDF page and equation.
2. Agent B's assertion.
3. Correct mathematical expression or reason it cannot be established.
4. Missing intermediate step.
5. Minimum necessary correction.
6. Final disposition.

Do not silently repair Agent B's derivation.

Do not rewrite the entire research paper.

---

## 5. Required audit deliverables

Produce these seven documents under:

`research/experiments/qsg-preflight-001/agents/agent-c-conservation-audit/`

1. `01-primary-source-verification.md` — Bibliography, exact equation references, conventions and source verification.
2. `02-classical-independent-derivation.md` — Independent classical derivation and off-shell/on-shell identities.
3. `03-hadamard-point-splitting-audit.md` — Quantum field foundations, subtraction prescription, smoothness and coincidence-limit analysis.
4. `04-conservation-and-counterterm-derivation.md` — Explicit divergence, local correction, cancellation and finite ambiguities.
5. `05-agent-b-findings.md` — Concise adversarial defect table with corrections and mathematical dispositions.
6. `06-theorem-contract.md` — Strongest defensible QSG-T001 proposition, assumptions, exclusions and proof obligations.
7. `07-independent-handoff.md` — Final mathematical verdict and implementation-readiness decision.

Every document must advance a mathematical decision. Avoid lengthy historical exposition and duplicated explanations.

Use the following dispositions in the audit documents:

- DERIVED
- SOURCE-REPORTED
- INDEPENDENTLY REPRODUCED
- CONDITIONAL
- FALSIFIED
- BLOCKED

Never substitute a generic PASS for a mathematical disposition.

---

## 6. Theorem acceptance gate

Do not begin theorem implementation until the following questions have been answered.

1. Is the classical identity independently reproduced?
2. Is the Hadamard remainder's actual field equation established?
3. Is the point-split divergence explicitly calculated?
4. Is the residual local term identified?
5. Is the conservation-restoring correction explicitly derived?
6. Is its divergence calculated?
7. Does the cancellation work under one consistent convention?
8. Are finite renormalization ambiguities correctly identified?
9. Is the resulting theorem statement appropriately restricted?
10. Are the primary sources sufficient to support the mathematical claims?

The expected research possibility is that the conservation theorem remains valid under the appropriately corrected renormalization prescription, even though Agent B's proof is invalid.

Do not assume this outcome.

If the calculations support a narrower theorem, adopt the narrower theorem.

If a necessary calculation cannot be completed, report BLOCKED and identify the exact obstruction.

---

## 7. Second priority: minimal QSG-T001 formalization

Only after the mathematical acceptance gate, use the existing GR implementation architecture.

Reference:

`phys-lib/gr/equivariance/`

Follow the established package-local AGENTS.md, manifest.json, Build/Verify conventions and `protocol/derivation_test.go` structure.

Use the existing infrastructure rather than designing another one.

Inspect the existing types:

- PhysicalAssumption
- PhysicalDefinition
- PhysicalEquation
- PhysicalTheorem
- ProtocolSketch

Respect the separation between mathematical representation and physical interpretation.

### Formalization boundary

The current architecture supports mathematical expressions, artifact references and bounded executable derivations.

It does not provide a general implementation of:

- Quantum operator algebras.
- Hadamard distributions.
- General point-splitting renormalization.
- State semantics.
- Arbitrary functional differentiation.
- General coincidence-limit calculations.

Do not invent these capabilities.

Represent the complete mathematical proof in conventional mathematics in the research documents.

Use ProtocolSketch for non-executable proof steps.

Mark a step SUPPORTED only when an actual replayed derivation or executable artifact supports it.

Use UNSUPPORTED for analytical steps that the current engine cannot execute.

A structurally valid theorem artifact is not equivalent to a machine-verified proof.

### Implementation constraints

- Do not redesign PhysMath, PhysLib or PhysArtifact.
- Do not introduce external dependencies.
- Do not modify existing GR fixtures.
- Do not create a Phys-Pseudo parser or validator.
- Do not introduce generic quantum-field abstractions without a demonstrated requirement.
- Do not change the module graph simply to make the fixture convenient.
- Follow the existing temporary Go workspace recipe for validation where necessary.

Choose the smallest appropriate new location under `phys-lib/gr/`, following the existing fixture conventions. Document the path and rationale.

The implementation must make the distinction between mathematically established results and unsupported analytical steps explicit.

---

## 8. Phys-Pseudo requirements

Use only the actual seven permitted constructs from `phys-pseudo/PSEUDO.md`:

- `given`
- `define`
- `assume`
- `claim`
- `derive`
- `if/then/else`
- `?`

Follow exact declaration, binding and conditional rules.

Remember:

- A `claim` is not automatically true.
- An `assume` is not an endorsed mathematical result.
- A `derive` requires named premises and explicit reasoning.
- `?` represents an unresolved issue, not a known false statement.
- Known false statements must be contradicted through a derivation.
- Do not introduce status or provenance constructs.

Put the concise QSG-T001 representation in `06-theorem-contract.md`.

Phys-Pseudo expresses the mathematical argument. It does not prove it.

---

## 9. Efficiency and execution discipline

This section is mandatory.

Previous research spent too much effort on broad exposition without completing its central mathematical calculation.

Do not repeat that pattern.

**Rule 1 — No repeated reconnaissance.**

Reuse Agent A's completed repository map, implementation contract and G1 closure report.

Only inspect a specific file again when an implementation decision requires it.

**Rule 2 — No broad literature survey.**

Start with the primary sources identified in Section 3D.

Expand the source list only when an unresolved equation or convention requires it.

**Rule 3 — No repeated PDF extraction.**

Use text extraction to locate relevant passages. Use rendered PDF pages to verify equations and mathematical notation.

Inspect each relevant page once, record the exact finding and reference it thereafter.

**Rule 4 — No premature conclusions.**

Never declare a theorem DERIVED before explicitly completing the necessary calculation.

A plausible final equation is not a derivation.

**Rule 5 — One convention at a time.**

Freeze the two-point function normalization, Hadamard parametrix, differential operator and curvature conventions before calculating.

Do not assemble equations from incompatible conventions.

**Rule 6 — Separate the research questions.**

Do not expand QSG-T001 into:

- General semiclassical backreaction.
- FLRW global existence.
- Stochastic gravity.
- Noise-kernel validity.
- Black-hole information.
- Bohmian mechanics.

Those belong to separate investigations.

**Rule 7 — Minimize implementation.**

Do not build a symbolic engine to accommodate one theorem.

Represent what is mathematically established and preserve the remaining analytical proof obligations.

**Rule 8 — No invented infrastructure.**

Do not create a new parser, status engine, generic framework or abstraction unless a demonstrated requirement cannot be met using existing capabilities.

**Rule 9 — Preserve research provenance.**

Never modify Agent B's PDF, completed Agent A reports or canonical research records.

Record corrections separately.

**Rule 10 — Prefer an early blocker over a long unsupported investigation.**

When a required source, coincidence limit, correction or mathematical identity cannot be established, identify the blocker immediately.

---

## 10. Acceptance criteria

The assignment is complete when:

1. All seven audit documents exist and have been read back.
2. The conservation-restoring local correction is derived, not asserted.
3. The corrected stress-energy tensor's divergence is explicitly shown to vanish.
4. Agent B's central defects are documented with exact locations.
5. The strongest defensible QSG-T001 theorem is stated precisely.
6. All assumptions and renormalization ambiguities are explicit.
7. The Phys-Pseudo representation conforms to the actual specification.
8. Any Go fixture follows the existing GR reference architecture.
9. Tests and build commands are recorded with actual results.
10. Unsupported analytical operations are never represented as executed proofs.
11. Existing repositories and completed Agent A deliverables remain untouched except for explicitly authorized minimal implementation changes.

## 11. Final handoff

Provide a concise report answering:

1. Was Agent B's conservation conclusion correct?
2. Was Agent B's derivation mathematically correct?
3. What correction is necessary?
4. What is the strongest theorem we can responsibly state?
5. Which results were independently reproduced?
6. Which assumptions and obligations remain?
7. Is QSG-T001 ready for statement-level formalization?
8. Is it ready for executable mathematical verification?
9. What was actually implemented?
10. What is the smallest remaining task?

**Final instruction: Do not optimize for the appearance of completion. Optimize for a result that another mathematical physicist can independently verify and another engineer can faithfully represent.**

### My recommendation

Give Agent A this prompt as the continuation task.

The critical sequencing is:

Independent derivation → theorem acceptance → minimal formalization → executable checks where supported.

Do not let Agent A begin by writing Go code. The existing architecture is already adequate for representing the research; the unresolved issue is whether the mathematical theorem has been correctly derived.

That is precisely where the previous work stopped.
