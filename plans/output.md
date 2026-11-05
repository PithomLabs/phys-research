The output is a continuously maintained, traceable physics research knowledge base that helps us formulate and validate new physical theories.

It is not a physics simulation platform, a paper summarizer, or a new gravity theory by itself.

There are two levels of output.

## 1. Immediate output: the research workbench

The implementation produces:

- A Go CLI (`physctl`) for managing research records.
- OpenCode agents that investigate, formulate and challenge scientific hypotheses.
- A Git repository containing structured research artifacts.
- A connected record of papers, claims, outstanding problems, hypotheses, reviews and experimental evidence.

For the MICROSCOPE MVP, the concrete output is:

## One completed research case

1. Source: MICROSCOPE 2022 paper, with precise citations and equations.
2. Problem: What does its experimental result establish about the equivalence principle?
3. Hypothesis: A compact Phys-Pseudo formulation of the theoretical expectation and experimental result.
4. Evidence: Measured values, uncertainties, experimental conditions and limitations.
5. Adversarial review: Independent challenges to assumptions, systematics and interpretation.
6. Accepted research record: Human-reviewed and preserved in Git.

## 2. Ultimate output: formally developed physical theories

This is the more important objective.

The workbench should progressively produce:

| Output                       | What it gives us                                                |
| ---------------------------- | --------------------------------------------------------------- |
| Research hypotheses          | Explicit, compact Phys-Pseudo formulations                      |
| Outstanding problems         | Precisely identified unresolved mathematics and physics         |
| Mathematical derivations     | Reproducible derivation artifacts                               |
| Definitions and equations    | Theory-specific PhysLib components                              |
| Mathematical representations | Theory-neutral PhysMath objects                                 |
| Theorems                     | Derived results with explicit assumptions and proof obligations |
| Predictions                  | Expected physical observables                                   |
| Experimental comparisons     | Agreement, contradiction or constraints                         |
| Theory specifications        | An integrated, inspectable physical framework                   |

Ultimately, this feeds into our existing infrastructure:

Phys-Pseudo → PhysMath → PhysLib → PhysArtifact → Phys evaluation.

## 3. The actual end goal

For example, our Branch C research might eventually produce:

A classical Einstein–Cartan theory consistently coupled to quantum matter through stochastic dynamics, provided that the mathematical construction succeeds and survives adversarial and empirical evaluation.

The output would not merely be a document claiming that the theory works. It would contain its assumptions, equations, derivations, consistency results, predictions, falsification conditions and known limitations.

Importantly, the system must also be capable of concluding that Branch C does not work.

In one sentence:

We are building an autonomous-agent-assisted scientific research and theory-construction workbench whose ultimate output is not summaries of physics papers, but rigorously derived, inspectable and experimentally testable physical theories.

The Go application manages the work. OpenCode performs the research. Phys provides the scientific infrastructure. Humans establish mathematical and physical acceptance.
