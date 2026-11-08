# PhysTrace

**A traceable, auditable way to do theoretical physics research — so that claims cannot outrun their evidence.**

---

## What this is

PhysTrace is a research methodology for theoretical physics. It is not a journal, not a preprint server, and not a chatbot. It is a **discipline**: a set of tools, rules, and records that forces every physics claim to carry its own audit trail, from the paper it came from, through the derivation that establishes it, to the experiment (or lack of one) that could kill it.

The project has two halves:

- **Phys Research** — a constitution and a workflow for how physics claims are proposed, challenged, derived, and accepted.
- **Phys Infrastructure** — the software (`phys`, `phys-lib`, `phys-math`, `phys-artifact`, `phys-pseudo`, and a small Go command-line tool called `physctl`) that makes the workflow checkable by machine.

The first public demonstration is a reconstruction of a celebrated result in black hole physics. The physics is old. What is new is the audit trail.

---

## The problem we are trying to solve

Theoretical physics has a credibility problem that has nothing to do with intelligence and everything to do with **process**.

A typical paper contains: a claim, a derivation, some citations, and a conclusion. Three things usually go wrong:

1. **The derivation has gaps.** A step says "it follows that" and the reader is expected to fill in the algebra. Sometimes the algebra is wrong.
2. **The claim grows in the telling.** A calculation in a toy model becomes a press release about "solving" a famous paradox.
3. **The provenance is lost.** Five years later, nobody can tell which version of which equation was actually verified, or by whom.

Every field has this problem. Physics is unusual only in how hard the math makes it to catch.

The analogy is a forensic audit. A company can produce a profit statement in an afternoon. A real auditor reconstructs every transaction, checks every receipt, and only then signs the accounts. PhysTrace is trying to make theoretical physics auditable in that sense.

---

## The Five Grand Problems

The project is built around five open problems in theoretical physics. They are not new — some have been open for sixty years. They are the reason a discipline like PhysTrace is needed: without a structured way to track what is established, what is assumed, and what is still open, progress stalls and claims drift from evidence.

| # | Problem | In one line | Status |
|---|---|---|---|
| **R1** | Semiclassical consistency | Do classical gravity and quantum matter keep honest books? | **Active** (demo) |
| **R2** | Problem of time | What happens when time disappears from the equation? | Back burner |
| **R3** | Born rule | Why do quantum outcomes have probabilities at all? | Input (taken as given) |
| **R4** | Black hole information | Does the fireplace burn the books? | **Active** (demo) |
| **R5** | Cosmological constant | Why is the vacuum so light? | Back burner |

The first demo (**JT-ISLAND-001**) tackles **R4** as its primary target and **R1** as its cross-check. The other three are acknowledged but explicitly out of scope for this proof of concept. Taking R3 (the Born rule) as an input is not a dismissal — it is a declaration. The Born rule is the one postulate no structure in three thousand years of mathematics has been able to derive.

---

### R1 — Semiclassical Consistency: Do Gravity and Quantum Matter Keep Honest Books?

**The question.** If matter is quantum and gravity is classical, can they be consistently coupled? Does the matter's energy drive gravity in a way that conserves energy, respects the symmetries of spacetime, and matches what we see?

**The analogy.** You hire a classical accountant (gravity) to track the spending of a quantum workforce (matter). The workforce exists in superpositions — it might be in two places at once. The accountant must produce a single ledger. Does the ledger balance? Does it match the bank statement?

**Why it's hard.** The Eppley–Hannah argument (1977) and the Page–Geilker experiment (1981) suggest the simplest hybrid theory is inconsistent — it either violates the Born rule or allows faster-than-light signaling. Whether this is a fatal objection or an artifact of assuming quantum mechanics is exact at every scale is open. If quantum mechanics is itself emergent, the objection loses its force.

**Status.** Active. R1 is the cross-check half of the first demo: does the matter energy used in the Page-curve calculation actually satisfy the conservation and energy-balance conditions that a consistent semiclassical theory requires?

---

### R2 — The Problem of Time: What Happens When Time Disappears?

**The question.** In general relativity, time is part of the dynamical geometry — it curves and flows with the universe. In quantum mechanics, time is a fixed external parameter — it ticks outside the system. When you try to quantize gravity (the Wheeler–DeWitt equation), time disappears from the equation entirely. What does that mean?

**The analogy.** A clockmaker builds a clock that measures everything except its own ticking. How do you know it is working?

**Why it's hard.** No agreed solution. The problem of time is one of the deepest conceptual issues in quantum gravity, and it is unresolved after sixty years. It is a problem about what time *is*, not just how to calculate with it.

**Status.** Back burner. Not part of the current demo. R2 would require a genuinely relational construction of time, not merely a change of coordinates.

---

### R3 — The Born Rule: Why Do Quantum Outcomes Have Probabilities?

**The question.** Quantum mechanics predicts amplitudes — complex numbers that can interfere. Experiments see outcomes — one of many possible results. The Born rule says the probability of an outcome is the squared modulus of its amplitude. But why? No one has derived the Born rule from something more fundamental.

**The analogy.** A casino advertises the odds of each game. But the rule that converts the casino's internal ledger into what the player actually experiences is posted on a wall nobody can read.

**Why it's hard.** The Born rule is a postulate, not a theorem. Every attempt to derive it (Everett, Bohm, GRW, decoherence, QBism) either changes the theory or assumes something equivalent to the Born rule in disguise. The Math Reality Check calls this "the wall all fifteen structures share": no mathematical discipline derives it, and no physics theory avoids assuming it.

**Status.** Input. Taken as given, not investigated in this demo. Taking R3 as an input is not a dismissal — it is a declaration that this project does not attempt to derive the Born rule, and will not pretend to.

---

### R4 — Black Hole Information: Does the Fireplace Burn the Books?

**The question.** Hawking's 1974 calculation says a black hole radiates featurelessly — as if the information of what fell in is destroyed. Quantum mechanics says information cannot be destroyed. In 2019, a team of physicists showed that in a simplified two-dimensional model, the information is preserved by a small region *inside* the black hole — the "island." Does this hold in general? Does it solve the paradox, or merely show that a toy model is consistent?

**The analogy.** A fireplace that burns books. The smoke looks featureless. But a tiny ember still inside the fireplace carries the full text. Count the ember as part of the smoke, and the books are not lost.

**Why it's hard.** The island calculation is a **consistency check**, not a microscopic mechanism. It shows that the entropy of radiation rises, peaks, and falls (the "Page curve"), but it does not construct the unitary S-matrix that carries the information out. Reproducing the Page curve is not the same as solving the information paradox. The distinction matters, and the demo says so.

**Status.** Active. R4 is the primary half of the first demo: reconstruct the Page curve from primary sources, step by step, with every assumption written down.

---

### R5 — Cosmological Constant: Why Is the Vacuum So Light?

**The question.** Quantum field theory predicts the vacuum energy should be enormous — about 10¹²⁰ times larger than what we observe. Why is it so small? And why is it not exactly zero?

**The analogy.** A room that should weigh as much as a mountain weighs as much as a feather. Either the room is not really there, or the way we count weight is wrong.

**Why it's hard.** No agreed solution. The cosmological constant problem is widely regarded as the worst fine-tuning problem in physics. Every attempt to explain it (anthropic, sequestering, unimodular gravity, swampland) either requires new physics or changes the rules of quantum field theory.

**Status.** Back burner. Not part of the current demo. R5 is a problem about the *value* of a parameter, and the Math Reality Check is clear that values are inputs, not outputs. Any resolution will require either a new symmetry or a new principle, and neither is in scope.

---

### Why these five, and why in this order?

The five problems were chosen because they are: (1) genuinely open, (2) mathematically well-posed, (3) connected to each other, and (4) tractable enough that a structured research infrastructure can make progress on at least the classification of what is known and unknown.

The order (R4 primary, R1 cross-check, R2/R3/R5 deferred) is not a ranking of importance. It is a ranking of **deliverability**. R4 and R1 can be reconstructed from existing literature, with a clear falsifier (does the energy balance hold? does the Page curve turn around?), in a controlled toy model, within months. R2, R3, and R5 would each require either new physics or a new foundational principle — and the project is not in the business of claiming new physics.

The point of the demo is not to solve the five problems. It is to show that PhysTrace can take a hard problem, reconstruct what is known, check it from a second angle, and refuse to upgrade the claim. That is a discipline. The five problems are the landscape. The demo is the proof that the discipline works.

---

## How it works

Every research claim passes through five stages. Nothing can skip a stage. Nothing can jump ahead.

```
   ┌─────────────┐
   │ 1. PROBLEM  │  What physical question are we answering?
   └──────┬──────┘
          ▼
   ┌─────────────┐
   │ 2. HYPOTHESIS│ Written in a strict format, with every gap marked "?"
   └──────┬──────┘
          ▼
   ┌─────────────┐
   │ 3. REVIEW   │  Adversarial: what breaks if a premise is removed?
   └──────┬──────┘
          ▼
   ┌─────────────┐
   │ 4. DERIVE   │  Every step must map to an existing rule, or be flagged
   └──────┬──────┘
          ▼
   ┌─────────────┐
   │ 5. SEAL     │  A human signs. The machine cannot.
   └─────────────┘
```

The important rules:

- **A hypothesis is allowed to be incomplete.** A `?` in the middle of a derivation is not a bug, it is the research question. The linter checks *syntax*, not *truth*. Pretending a hypothesis is complete is the actual failure.
- **A derivation step either maps to an existing rule or becomes an obligation.** If a step has no rule, it is not a derivation, it is a wish.
- **The machine can check structure, not physics.** Passing a validator proves the file is well-formed. It does not prove the physics is right.
- **Only a human can seal a theorem.** A command-line flag is not a permission boundary. A signed Git commit is.

The analogy: a legal contract is different from a handshake because a contract has witnesses and a paper trail. PhysTrace is the contract, the witnesses, and the paper trail, applied to mathematics.

---

## The Three Reality Checks

Every substantive investigation must pass three questions. They are deliberately simple.

| Check | The question | The analogy |
|---|---|---|
| **1. Landscape** | Which branches of mathematics are needed, and what can each one actually do? | Which tools are in the toolbox? |
| **2. Intersection** | Which combination of tools produces the result, and what fails if you remove one? | Which ingredients in the recipe are essential? |
| **3. Interoperability** | When one branch of mathematics hands a result to another, what is lost in translation? | What are the unit conversions? |

The third check is the one most physics papers skip. A number can mean different things in geometry, field theory, and information theory. PhysTrace requires every "bridge" between two branches to be written down explicitly: source, target, assumptions, preserved properties, approximations, and failure conditions.

---

## The first demonstration

The first public demo is called **JT-ISLAND-001**. It reconstructs a well-known 2019 result about black hole evaporation, in a simplified two-dimensional model of gravity.

The physics, in one paragraph: if a black hole evaporates and the radiation carries away no information, quantum mechanics is broken. If information is preserved, the entropy of the radiation should rise, peak, and fall. The 2019 calculation showed that this turnaround happens in the simplified model, because a small region *inside* the black hole, called an "island," is secretly entangled with the radiation outside. The entropy at any moment is whichever is smaller: the entropy computed with the island, or the entropy computed without it. The result is a curve that rises, peaks, and falls.

The fireplace analogy: the smoke looks featureless, but a tiny ember still inside the fireplace carries the full information. Count the ember as part of the smoke, and the books are not lost.

The point of the demo is **not** to solve the black hole information paradox. It is to show that PhysTrace can:

- trace every equation to a primary source,
- represent the calculation in five different mathematical languages at once,
- cross-check the matter-energy bookkeeping against the entropy calculation,
- and refuse to upgrade the claim beyond what was proven.

A second investigation (R1) acts as an independent auditor: does the energy leaving the black hole exactly match the energy arriving in the surrounding space? If the two fixtures (R4 and R1) use the same metric, the same vacuum state, and the same central charge, then the calculation is internally consistent. If they do not, the demo fails, and we say so.

---

## The Constitution, in plain English

The project is governed by an eight-article constitution. It is not decoration. It is the rulebook that every tool, agent, and researcher must obey.

| Article | In one sentence |
|---|---|
| **I. Reality is the subject** | Physics, not math, not software, is the point. |
| **II. Three Reality Checks** | Every investigation must show its tools, its recipe, and its translations. |
| **III. Epistemic integrity** | Facts, assumptions, derivations, and conjectures must be labeled as such. |
| **IV. Research and discovery** | Novelty claims require prior-art comparison and honest classification. |
| **V. Independent challenge** | Every consequential claim must be open to adversarial review. |
| **VI. Freedom of implementation** | Tools and workflows can change; the principles cannot. |
| **VII. Scientific promotion** | Claims cannot exceed their evidence. Negative results are results. |
| **VIII. Constitutional continuity** | No agent, no tool, and no convenience can override the principles. |

The constitution is short on purpose. A rulebook that is longer than the research it governs is a rulebook nobody reads.

---

## What this project does not claim

PhysTrace is deliberately conservative about what it can prove. It does not claim:

- to solve the black hole information paradox,
- to construct a complete theory of quantum gravity,
- to replace human judgment with machine verification,
- to certify that a physics claim is *true*.

It claims only this: **if a claim passes through PhysTrace, the reasoning is inspectable, the assumptions are explicit, and the provenance is intact.** Whether nature agrees is a separate question, answered by experiment.

The analogy: a proofreader can guarantee that a manuscript has no typos. A proofreader cannot guarantee that the manuscript is a masterpiece. PhysTrace is the proofreader.

---

## Repository layout

```
phys-research/
├── protocol.md              # The Constitution (v1.0, 8 articles)
├── AGENTS.md                # Shared rules for all AI agents
│
├── .opencode/
│   └── agents/
│       ├── researcher.md    # Reads papers, drafts hypotheses
│       ├── adversary.md     # Challenges assumptions, finds gaps
│       └── formalizer.md    # Transcribes human derivations into checked artifacts
│
├── research/
│   ├── sources/             # Paper records: DOI, arXiv ID, claims, equations
│   ├── issues/              # Outstanding problems, with severity and status
│   ├── hypotheses/          # Phys-Pseudo hypotheses
│   ├── reviews/             # Adversarial findings
│   └── evidence/            # Experiments, observations, expected vs actual
│
├── derivations/             # Human-authored, machine-replayable proof steps
│
└── cmd/
    └── physctl/             # The Go CLI (lint, status, refs, check, hash)
```

There is no database. There is no message queue. There is no orchestration server. Git is the ledger. Files are the records. The Go tool is the gatekeeper.

The analogy: a library does not need a database to know where the books are. It needs a card catalog, and a librarian who insists that every book has a call number. Git is the card catalog. `physctl` is the librarian.

---

## How to read the first demo

Start with these three files, in this order:

1. **`protocol.md`** — the rules. Ten minutes.
2. **`research/issues/JT-ISLAND-001.md`** — the question. Five minutes.
3. **`research/hypotheses/JT-ISLAND-001.pseudo`** — the working hypothesis, with every unresolved gap marked `?`. Ten minutes.

If those three documents are clear, the rest of the repository is details. If they are not clear, the rest of the repository is not worth reading, and that is exactly the point.

---

## Status

| Milestone | Status |
|---|---|
| Constitution v1.0 | Adopted |
| POC direction (R4 + R1 in JT gravity) | Locked |
| Deep Research dossier v1.0 | Delivered, under review |
| `physctl` CLI specification | In progress |
| Muse formalization | Not started |
| First human-sealed theorem | Not yet |

The first proof of concept is a single research issue, taken end to end: paper ingestion, adversarial review, a human-signed derivation, and a sealed theorem. Nothing else is in scope until that one loop works.

---

## Guiding principle

> **Automate the clerical work. Make the reasoning inspectable. Keep mathematical proof and physical acceptance rigorous.**

Everything in this repository, from the constitution to the Go binary, is an attempt to take that sentence seriously.
