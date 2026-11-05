Agreed. We have been spending too much attention on engineering safeguards before demonstrating the platform's scientific value. The priority should now be research capability, not infrastructure expansion.

I would reclassify Plan 6 as deferred maintenance. We should address only issues that actually prevent the research workflow from functioning.

## The highest-impact next milestone: a complete physics research cycle

The system must demonstrate that it can move from published literature to a mathematically inspectable research problem.

I would prioritize these capabilities:

| Priority | Capability                                    | Research value                                                           |
| -------- | --------------------------------------------- | ------------------------------------------------------------------------ |
| P0       | Literature discovery and source qualification | Find the right papers, not merely keyword matches                        |
| P0       | Equation and claim extraction                 | Represent what papers actually establish                                 |
| P0       | Assumption and regime tracking                | Identify where a result applies and where it does not                    |
| P0       | Cross-paper comparison                        | Expose contradictions, incompatible assumptions and missing connections  |
| P0       | Derivation obligations                        | Turn unresolved mathematical steps into explicit, inspectable work       |
| P1       | Competing hypothesis development              | Formulate alternative explanations without prematurely accepting one     |
| P1       | Experimental correspondence                   | Connect theoretical predictions to actual measurements                   |
| P2       | Advanced autonomous orchestration             | Automate more of the workflow after its research quality is demonstrated |

The distinction is important: a paper summary is not a physics research result. The valuable output is a chain connecting assumptions, mathematical representations, derivations, predictions and evidence.

## What I would build next

Use the existing Branch C investigation as the first serious functional test.

Research question:

Can classical Einstein–Cartan geometry consistently couple to quantum matter without quantizing spacetime, and what mathematical conditions must such a theory satisfy?

Keep three possibilities separate:

1. Semiclassical Einstein–Cartan gravity.
2. Stochastic or hybrid classical–quantum Einstein–Cartan dynamics.
3. Quantum Einstein–Cartan gravity.

Do not assume one framework can inherit the results of another.

Have Researcher produce one integrated research dossier containing:

- Primary sources, relevance and exact equation/section locators.
- The mathematical structure of each framework.
- Explicit assumptions and physical regimes.
- A comparison of their constraint structures and conservation requirements.
- Established derivations versus claims merely asserted by authors.
- Missing derivation steps, represented as explicit `?` obligations.
- Competing hypotheses and what mathematical results would distinguish them.

The output should be usable by a human physicist to continue the investigation without having to repeat the literature review.

## What we should defer

| Current engineering issue             | Decision                                                 |
| ------------------------------------- | -------------------------------------------------------- |
| OpenCode permission hardening         | Defer                                                    |
| OS-level agent isolation              | Defer until broader autonomous operation is needed       |
| Extensive rate-limiter stress testing | Defer unless it blocks actual research acquisition       |
| TLS investigation                     | Resolve only if real paper retrieval fails               |
| Cache symlink protection              | Retain as a gate before unrestricted autonomous fetching |
| Security documentation expansion      | Defer                                                    |

Maintain source provenance, existing validation and human control over scientific acceptance. Those are foundational, not accidental complexity.

## The rule going forward

Every proposed feature must answer:

Does this materially improve our ability to discover, understand, compare, derive, verify or experimentally evaluate physics?

If not, it waits.

Also, no new abstraction, agent role, service, database or orchestration mechanism should be introduced until an actual research cycle demonstrates the need.

My recommendation is to pause Plan 6 and make the next milestone Researcher Scientific Capability Demonstration. We should measure success by the quality of the mathematical questions and derivations it produces—not by how much infrastructure we have built.
