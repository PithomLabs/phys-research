Agent A's plan is good. I would approve it with a few important adjustments, rather than ask the agent to redesign anything.

It correctly recognizes that QSG preflight is primarily an infrastructure reconnaissance exercise, and that the existing Phys architecture should be reused wherever possible.&#x20;

plan4(20261106-044448).md



## 1. Decisions on Agent A's questions

| Question                                               | My decision                                                                                                                                                | Reason                                                                                         |
| ------------------------------------------------------ | ---------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| Q1: No Git repository in TOE-v6                        | Use a standalone assigned folder. Do not run `git init`.                                                                                                   | Avoid changing the workspace simply to satisfy a worktree requirement.                         |
| Q2: Full PhysMath suite takes approximately 30 minutes | Authorize targeted tests and shard C only. Defer heavy shards.                                                                                             | Reconnaissance does not justify unnecessary expensive testing.                                 |
| Q3: Which GR implementation is canonical?              | Preserve `phys-gr/solution` as the canonical protocol reference. Inspect `phys-lib/gr/equivariance` as the current implementation and packaging reference. | Maintain the existing GR theorem standard while benefiting from newer implementation patterns. |

One detail needs resolving: the plan mentions both `phys-lib/gr/equivalence` and `phys-lib/gr/equivariance`. Agent A must verify the exact directory names before proceeding.

Do not modify TOE-v6's Git structure. The assigned folder provides logical separation, while filesystem permissions and read-only repository access provide the actual safeguards.

## 2. Four adjustments to the plan

These are important for maintaining our engineering discipline.

A. Do not classify missing infrastructure based solely on keyword searches.

The plan reports zero matches for `renormal`, `Hadamard`, `two-point` and related terms. That is useful initial evidence, but absence of a particular name does not prove absence of an underlying capability.

Require inspection of the actual mathematical types, interfaces, packages and documentation before marking something MISSING.

B. Separate representation from execution.

For QFT, this distinction is critical.

For example:

- Can PhysMath represent a two-point function?
- Can it represent an operator-valued quantity?
- Can Phys-Pseudo express a renormalization assumption?
- Can Phys Derivation actually verify a renormalized composite operator identity?

These are four different capabilities.

A mathematical object being expressible as an equation does not mean our infrastructure can verify it.

C. Preserve the standard-library constraint.

Do not introduce symbolic algebra dependencies or a new mathematics engine during reconnaissance.

Agent A should identify what can be handled through the existing Go infrastructure and what must remain an analytical proof obligation.

D. Do not implement missing functionality yet.

The deliverable is an implementation contract, not an implementation.

The reported functional-derivative limitation and apparent renormalization gaps must be accurately classified before recommending changes.&#x20;

plan4(20261106-044448).md



## 3. Revised instruction to Agent A

Approved. Proceed with Plan4, subject to these decisions and adjustments.

1. Q1: TOE-v6 has no Git repository. Create the assigned standalone directory under TOE-v6. Do not initialize Git or alter workspace repository structure. Maintain strict read-only access to shared repositories. Record the limitation of folder-based isolation.
2. Q2: Run the planned targeted builds, focused tests, PhysMath shard C and purity/import checks. Defer the approximately 30-minute full PhysMath suite unless your findings establish that it is necessary to resolve a material uncertainty. Record exact commands, exit codes and test results.
3. Q3: Treat the existing `phys-gr/solution` structure and protocol as the canonical theorem-format reference. Inspect `phys-lib/gr/equivariance` as the current implementation and packaging reference. Compare them explicitly. Verify the actual directory names; the plan currently mentions both `equivalence` and `equivariance`. Do not independently redefine the canonical theorem protocol.

Additional requirements:

4. A zero-result keyword search is not sufficient evidence that a capability is MISSING. Inspect actual types, interfaces, implementations and documentation. Use NOT VERIFIED where evidence is insufficient.
5. Strictly distinguish mathematical representability from executable verification. In particular, assess operator-valued objects, two-point functions, renormalization, functional derivatives and nonlocal kernels separately.
6. Do not add dependencies, implement mathematical capabilities, modify shared repositories or redesign Phys Research during reconnaissance.
7. Preserve all BHC artifacts unchanged.
8. Complete all eight deliverables and perform the final non-modification audit.

The objective remains an accurate, minimal implementation contract establishing whether the existing infrastructure can support QSG-T001 without architectural changes.

Proceed.

## 4. What I expect at the end

I want Agent A's final report to answer one fundamental engineering question:

Can our existing Phys infrastructure take a conventional QFT hypothesis, express its mathematical obligations in Phys-Pseudo, and turn an accepted derivation into an executable GR-style theorem artifact?

In particular, I want a clear verdict on these three capabilities:

1. Representation: Can we express the mathematical objects required by QSG?
2. Verification: Can the derivation protocol check the properties we need?
3. Provenance: Can we trace every result to its assumptions, sources and derivation?

The distinction is essential. We might discover that our infrastructure can represent QSG-T001 perfectly well but cannot yet execute some of its mathematical verification.

That would be a precise engineering finding, not a reason to redesign everything.

My decision: authorize Agent A to execute Plan4 with these amendments. Keep Agent B's physics investigation independent. Their findings should meet when we formulate the QSG-T001 theorem contract—not before.
