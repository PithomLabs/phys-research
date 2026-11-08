The OpenCode assignment should be deliberately adversarial. Muse has completed the package and passed its mechanical self-audit; now OpenCode must determine whether the physics and mathematical connections are actually correct, independently of Muse's conclusions.

I would use the following prompt.

# OpenCode Adversarial Review: JT-ISLAND-001

## Independent Scientific and Mathematical Verification — Review v1.0

**Role:** Independent adversarial physics and mathematics reviewer\
**Project:** JT-ISLAND-001\
**Primary target:** R4 — Page-curve reconstruction\
**Cross-validation target:** R1 — Semiclassical matter–geometry consistency\
**Specification under review:** Phys-Pseudo R1/R4 v0.1\
**Governing authority:** Phys Research Workflow Protocol v1.0\
**Workspace:** TOE-v6\
**Review mode:** Independent reconstruction and attempted falsification\
**Authorization:** Review only. NO IMPLEMENTATION.

---

## 1. Mission

Independently determine whether Muse's JT-ISLAND-001 research package is mathematically sound, physically coherent, faithful to the primary literature and compliant with Phys Research Workflow Protocol v1.0.

Your responsibility is not to confirm that Muse followed instructions or produced a well-organized specification.

Your responsibility is to challenge whether the underlying mathematics and physics are correct.

**Attempt to disprove the central results.**

Reconstruct important equations independently from the pinned primary sources. Check the assumptions, signs, normalization, mathematical dependencies, physical regimes and interoperability between R1 and R4.

Do not simply repeat Muse's derivations and declare them verified.

The research package is a candidate formalization, not an approved scientific result.

## 2. Review package

Review the complete indexed directory:

`TOE-v6/research/experiments/jt-island-001/`

Start with `index.md` and follow its documented reading order.

The principal artifacts are:

- `phys-pseudo-r1r4-v0.1.md`
- `working-baseline.md`
- `branch-reconstruction.md`
- `audit-corrections-v1.1.md`
- `registers.md`
- `reality-checks.md`
- `opencode-checklist.md`
- `summary.md`
- `index.md`
- `sources/`

Inspect the original pinned sources, including their versions and hashes.

The dossier's equations and earlier audit corrections are historical inputs, not ground truth. The primary literature and independently reconstructed mathematics take precedence.

The original dossier v1.0 must remain unchanged.

## 3. Independent reconstruction requirements

### A. Canonical model and source compatibility

Independently verify the selected physical model:

The two-sided evaporating JT gravity setup coupled to a nongravitating bath, with the AMMZ holographic-matter specialization.

Determine whether the combination of AEMM and AMMZ equations is actually valid.

Specifically investigate:

- The physical configuration of the gravitational region and bath.
- The number and location of boundaries.
- Initial quantum state.
- Shock or injection conditions, if applicable.
- Boundary trajectory and coordinate transformations.
- Matter content and central charge.
- Metric and conformal-factor conventions.
- Cutoff surfaces and regulator transformations.
- Which equations are imported from AEMM.
- Which equations are specialized or derived by AMMZ.
- Which results require separate compatibility arguments.

Do not assume that equations from the two papers can be combined merely because AMMZ builds on the AEMM framework.

Confirm every cross-paper mapping used in the specification.

### B. R1 — Semiclassical matter–geometry consistency

Independently reconstruct:

1. JT gravitational action, including its topological contribution and boundary terms.
2. Dilaton equation and curvature constraint.
3. Metric field equations.
4. Polyakov effective action and stress-energy expectation tensor.
5. Conformal trace-anomaly relation.
6. Diffeomorphism Ward identity.
7. Covariant stress-energy conservation.
8. Relevant conformal light-cone expressions.
9. Transparent-interface matching conditions.
10. Schwarzian boundary dynamics and ADM mass.
11. Energy exchange and global energy balance.

Check signs, numerical coefficients, coordinate conventions, dimensions and boundary contributions.

Pay special attention to the distinction between:

- The trace anomaly.
- A diffeomorphism anomaly.
- The diffeomorphism Ward identity.
- Covariant conservation of the stress-energy expectation tensor.

Determine whether each conclusion follows from the stated mathematical assumptions.

### C. R4 — Page-curve reconstruction

Independently derive the no-island and island branches.

Check:

- The generalized-entropy functional.
- The gravitational entropy contribution per QES endpoint.
- The topological dilaton contribution.
- The dynamical dilaton contribution.
- The matter CFT entanglement entropy.
- The assumptions behind the two-interval entropy expression.
- Purity and state assumptions.
- Cross-ratio dependence and conformal-block approximations.
- QES stationarity conditions.
- QES solution and its approximation regime.
- On-shell island entropy.
- No-island asymptotic growth.
- Island-branch behavior.
- Branch intersection and Page-time estimate.
- The order of extremization and entropy minimization.

Independently verify the reported no-island slope:

[\
\frac{\pi cT}{6}.\
]

Independently reconstruct the leading Page-time estimate:

[\
t_P\sim\frac{6\beta S}{\pi c}.\
]

Do not accept either expression merely because Muse reports multiple checks.

Verify the exact assumptions and approximation footprint, particularly the slow-evaporation condition and the reported correction of order one.

Determine whether the result is exact, leading-order, approximate or conditional.

### D. Highest-risk unresolved obligations

Give special attention to:

**OB-4 — Black-hole mass evolution**

Reconstruct the ADM mass relation and energy-flux equation line by line. Verify the sign, normalization, flux orientation and boundary-time convention.

**OB-5 — Shock versus smooth evolution**

Determine whether the selected shock or smooth evolution assumptions are consistent across the canonical model and the derivations using it.

**OB-7 — Page-time approximation**

Reconstruct the five-item approximation footprint. Determine which assumptions produce the leading estimate and where it stops being reliable.

**OB-9 — Quantum extremal surface**

Challenge the imported QES result. Establish the exact source theorem or derivation, its hypotheses and applicability to the selected model.

**OB-10 — Ryu–Takayanagi contribution**

Challenge the imported RT result. Establish why the selected higher-dimensional holographic construction applies to the radiation calculation and what mathematical content is imported rather than derived.

Do not promote an imported result to an independently derived result.

### E. R1–R4 interoperability

This is a central acceptance requirement.

Verify that R1 and R4 genuinely share:

- The same physical model.
- Compatible geometry and coordinate conventions.
- The same relevant matter state.
- Compatible central charge and parameter hierarchy.
- Compatible boundary conditions.
- Compatible cutoff and regulator definitions.
- Consistent stress-energy and entropy calculations.

In particular, investigate the claimed AEMM–AMMZ evaporation-rate identity:

[\
\kappa=k.\
]

Reconstruct it independently from the cited equations.

Determine whether the identity is exact or conditional.

Verify that the geometry and matter-state inputs established in R1 are precisely the inputs consumed by R4.

The dependency must be one-way:

R1 validated background and matter assumptions → R4 entropy reconstruction → shared cross-validation result.

Identify any circular dependencies, unsupported transitions or hidden assumptions.

## 4. Audit all specification registers

Review every consequential claim and its epistemic classification.

Inspect:

- Claims C-R4-01 through C-R4-05.
- Claims C-R1-01 through C-R1-04.
- Cross-validation claim C-INT-01.
- Bridges BR-01 through BR-10.
- Assumptions ASM-01 through ASM-06.
- Unresolved obligations ?OB-1 through ?OB-10.

Check whether each item has:

- A precise statement.
- Correct premises.
- An applicable source and locator.
- A physical regime.
- An explicit derivation or imported result.
- Correct epistemic status.
- Relevant limitations and failure conditions.

Challenge the dependency graph rather than accepting Muse's assertion that it is acyclic.

Check whether definitions conceal assumptions, whether imported theorems retain their original hypotheses and whether a claim is being used as its own premise.

## 5. Apply the Three Reality Checks independently

The Three Reality Checks are mandatory constitutional requirements.

Do not merely check whether Muse included three sections with those headings.

### RC1 — Mathematical Landscape

Determine whether the mathematical structures are correctly identified, scoped and connected to their actual roles.

Check whether important structures have been omitted or mathematical capabilities overstated.

### RC2 — Mathematical Intersection

Examine whether the proposed intersections genuinely establish the claimed phenomena.

Challenge unsupported minimality claims. Test what changes when relevant mathematical ingredients are removed or altered.

### RC3 — Mathematical Interoperability

Inspect all ten-field bridge contracts:

- Source structure.
- Target structure.
- Mapping.
- Assumptions.
- Preserved properties.
- Losses or approximations.
- Applicable regime.
- Proof obligations.
- Failure conditions.
- Evidence.

Determine whether each bridge is mathematically valid.

A successful calculation in one mathematical language does not automatically establish a valid translation into another.

## 6. Phys-Pseudo conformance

Independently inspect the specification for:

- Exclusive use of the seven authorized constructs.
- Declaration-before-use.
- Definitions that do not smuggle in assumptions.
- Complete premise lists for every `derive`.
- Explicit reasoning for each derivation.
- Correct conditional branch scope.
- Correct attachment of unresolved `?` obligations.
- No hidden implementation details.
- No unapproved software abstractions.
- Complete and traceable mathematical dependencies.

Verify that the specification is implementation-independent.

Identify mathematical operations that are represented but not actually derived or verified.

Do not confuse a syntactically valid Phys-Pseudo document with a mathematically valid physical theory.

## 7. Limiting cases and falsification attempts

Challenge the claimed limiting cases:

1. Vanishing gravitational coupling.
2. Reflecting or decoupled bath boundary.
3. Breakdown of the semiclassical parameter hierarchy.

Independently derive their stated consequences.

Additionally, search for:

- Dimensional inconsistencies.
- Sign errors.
- Missing endpoint contributions.
- Inappropriate coordinate transformations.
- Invalid limiting operations.
- Missing regularity conditions.
- Incompatible state assumptions.
- Boundary terms that have been dropped.
- Counterexamples to overgeneralized claims.
- Incorrect transitions between local and global statements.

Where practical, construct a concrete counterexample or demonstrate the failed derivation.

If a claim cannot be disproved, do not automatically classify it as proven.

## 8. Review classifications

Assign each finding one severity:

**BLOCKER**

Invalidates a central conclusion, creates a circular derivation, uses an incompatible physical model, has a material mathematical error or makes a false scientific claim.

Must be resolved or the affected claim must be narrowed or blocked.

**MATERIAL**

Weakens the scope or support of a conclusion, exposes a missing assumption or requires a significant qualification.

Must be resolved or explicitly documented.

**MINOR**

Editorial, organizational or non-load-bearing technical issue.

Record it without unnecessarily blocking progress.

For each finding provide:

- Finding ID.
- Severity.
- Affected file and exact section or line.
- Mathematical or physical issue.
- Independent evidence or derivation.
- Relevant primary source and exact locator.
- Consequence for the research conclusion.
- Minimum correction required.
- Whether the issue affects R1, R4 or their interoperability.

Do not recommend architectural expansion to solve a mathematical defect.

## 9. Final verdict

Provide exactly one overall review outcome:

- **PASS:** No unresolved BLOCKER or MATERIAL findings. The specification is suitable for project-director approval.
- **CONDITIONAL:** No unresolved BLOCKERs, but MATERIAL qualifications or limited scope restrictions remain. State exactly what is permitted and what remains restricted.
- **BLOCKED:** At least one unresolved BLOCKER prevents reliable approval of the affected specification or central conclusion.

Include separate verdicts for:

1. R1 mathematical validity.
2. R4 mathematical validity.
3. R1–R4 interoperability.
4. Phys-Pseudo conformance.
5. Source fidelity and epistemic integrity.
6. Overall JT-ISLAND-001 readiness.

A PASS does not authorize implementation. Only the project director may approve and freeze the specification.

## 10. Required review deliverable

Produce a self-contained report:

`research/experiments/jt-island-001/reviews/opencode-review-v1.0.md`

The report must contain:

1. Executive verdict.
2. Review methodology and source versions.
3. Independent mathematical reconstructions.
4. Findings classified BLOCKER/MATERIAL/MINOR.
5. Detailed assessment of OB-1 through OB-10.
6. R1 assessment.
7. R4 assessment.
8. Interoperability assessment.
9. Three Reality Checks assessment.
10. Phys-Pseudo conformance assessment.
11. Limiting-case and counterexample assessment.
12. Required corrections.
13. Accepted claims and their scope.
14. Rejected or narrowed claims.
15. Final verdict and approval recommendation.

Reference the pinned source files and exact equation or theorem locators. Preserve enough mathematical work for another reviewer to reproduce your conclusions.

Do not silently modify Muse's artifacts. Do not rewrite the specification on Muse's behalf.

Keep all review outputs within TOE-v6. Leave the original research package and pinned sources unchanged.

## 11. Strict execution boundaries

This is one full adversarial review.

Do not:

- Write Go code.
- Modify PhysMath, PhysLib, PhysTrace or other shared repositories.
- Create implementation contracts.
- Add unsupported mathematical primitives.
- Change the physical model silently.
- Freeze the specification.
- Authorize implementation.
- Claim scientific novelty.
- Treat green mechanical checks as scientific proof.

If a critical defect is found, explain exactly why it blocks the relevant conclusion and what minimum correction is required.

Do not initiate endless review cycles. After submitting the report, stop.

A subsequent targeted review may inspect only the corrections Muse makes in response to your findings.

**The standard is independent mathematical reconstruction, not agreement with Muse. Your duty is to protect scientific correctness, even when that means blocking an attractive result.**

## What I particularly want OpenCode to decide

Three questions matter more than everything else:

1. Is the AEMM–AMMZ combination genuinely one coherent physical model?
2. Are the entropy normalization, Page-curve branches and Page-time approximation mathematically consistent?
3. Does R1 genuinely validate the background and matter assumptions consumed by R4, or is the claimed interoperability still partly assumed?

The completion report identifies these as central achievements and outstanding risks.&#x20;

plan_prompt9_imp2.md



The review should follow the package's existing 18-item checklist, but OpenCode must independently verify it rather than simply mark its boxes.

My desired outcome is an honest PASS, CONDITIONAL or BLOCKED—not a PASS at any cost. Even a BLOCKED result is valuable if it identifies exactly where the scientific reasoning fails.
