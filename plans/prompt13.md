This is the right next step. In fact, I think this is the most important methodological improvement we have made to `phys-research` so far.

We are finally putting a hard boundary between discovering a mathematical possibility, establishing a mathematical result, and implementing that result.

I particularly agree with the decision to reuse the existing Researcher, Adversary and Formalizer roles rather than build more infrastructure.

## What I like about this plan

1\. Phys-Pseudo becomes the mandatory mathematical reasoning layer.

The O5a investigation must be rewritten using the actual seven-construct grammar, with named premises and explicit derivation results.

This matters because the existing notes contain many useful mathematical templates, but some of those templates are still proposed conditions rather than derived results.

2\. The adversarial review becomes substantive.

The requirement to review every consequential `derive` individually is much stronger than asking an adversarial agent whether an entire document looks correct.

I particularly support:

- Independent reproduction.
- Phase-space and bracket compatibility.
- Index, density-weight and dimensional consistency.
- Limiting cases.
- Detecting smuggled assumptions.
- Counterexample attempts.
- Explicit evidential boundaries.

The restriction against global correct/incorrect verdicts is also sensible. A derivation can be verified within a specific scope without establishing the physical theory.

3\. Formalizer remains idle.

This is essential. The Formalizer should not convert plausible-looking equations into authoritative library objects merely because they are expressed in structured syntax.

4\. Library separation is preserved.

The proposed mapping is appropriate:

| Destination  | What belongs there                                                             |
| ------------ | ------------------------------------------------------------------------------ |
| PhysMath     | Pure mathematical structures, brackets and transport algebra                   |
| PhysLib      | Einstein–Cartan physical interpretation and theory-specific structures         |
| PhysArtifact | Provenance and identities connecting research, derivations and implementations |
| Phys         | Frozen legacy system                                                           |

The first formalization should happen only after mathematical review and human authorization.

## Three safeguards I would emphasize

There are three details I would make especially explicit to Muse.

First: do not force derivations into existence.

The existing O5a work contains numerous proposed symbolic expressions.

For example, an expression involving an undefined \\(W\_{\mathrm{EC}}\\) is not a completed derivation simply because it appears beneath `derive`.

Use:

- `given` for source-supported inputs.
- `assume` for conditional mathematical premises.
- `claim` for the proposition being investigated.
- `derive` only when an actual inference and result are supplied.
- `?` for unresolved calculations.

A source-transcribed bracket is a premise, not an independently reproduced derivation.

Second: distinguish mathematical review from scientific acceptance.

A reviewer marking a derivation `VERIFIED` means that the reviewer independently reproduced it within the stated checks. It does not mean the physical model is correct, empirically supported or universally valid.

Third: the gateway is currently conventional, not technically enforced.

The screenshot correctly identifies that `AGENTS.md` requires a human-applied patch. Until that happens, the gateway is a documented workflow rule supported by review discipline, not a hard technical prohibition.

That is acceptable at this stage, provided everyone understands the distinction.

## What I expect at the end of this phase

Four concrete deliverables

1. RULE-T1 Phys-Pseudo investigation: Canonical syntax, explicit assumptions, mathematical targets, actual derivations and unresolved obligations.
2. Adversarial review record: A separate disposition for each consequential derivation, with reproduction evidence, objections and limitations.
3. Formalization readiness audit: An inventory of what PhysMath and PhysLib can represent, and what is still missing. No library modifications.
4. Gateway documentation and proposed AGENTS.md patch: The formalization rule, review requirements and human-authorization boundary.

All existing library repositories remain untouched during this phase.

One additional success condition: the investigation must be allowed to conclude that none of its proposed derivations currently qualifies for independent verification. That is better than manufacturing a successful result to satisfy the plan.

## My verdict

Approve this approach

I would proceed exactly in this order:

Researcher → Phys-Pseudo → Adversarial review → Human authorization → Formalization readiness → PhysMath / PhysLib / PhysArtifact.

The important change is that the system is no longer being judged by how much physics material it generates.

It is being judged by whether it can produce mathematical reasoning that survives independent reconstruction and criticism.

That is the standard we need if `phys-research` is going to contribute something scientifically meaningful rather than become a sophisticated physics-paper factory.
