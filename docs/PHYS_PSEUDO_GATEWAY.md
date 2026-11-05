# Phys-Pseudo Formalization Gateway

Status: documented convention. Until the human applies the proposed
`AGENTS.md` patch (`docs/AGENTS_GATEWAY_PATCH.proposal.md`), this
gateway is a workflow rule supported by review discipline — not a
technically enforced prohibition. Do not describe it as enforced.

## Pipeline order

Researcher → Phys-Pseudo investigation → Adversarial review →
Human authorization → Formalization readiness audit →
PhysMath / PhysLib / PhysArtifact.

No stage may be skipped. A failed adversarial review returns the
investigation to Phys-Pseudo; it never proceeds automatically.

## Investigation rules

- Canonical `PSEUDO.md` syntax only: named constructs, non-empty
  `from [NAMES]` premise lists, explicit `derive` results.
- `given` for source-supported inputs (locators in prose/comments).
- `assume` for conditional premises. `claim` for propositions under
  examination. `derive` only where an actual inference and explicit
  result are supplied. `?` for unresolved calculations, with what
  would resolve them.
- A source-transcribed bracket is a premise, never an independently
  reproduced derivation. Undefined objects (kernels, operators,
  transports) must stay in `claim`/`?`, never in `derive` results.

## Review rules

- Every consequential `derive` gets its own disposition in the review
  record: `VERIFIED`, `CONDITIONAL`, `UNSUBSTANTIATED`, or `FALSIFIED`.
- Dispositions live only in the review record, never in the
  investigation document (`PSEUDO.md` §7). No global correct/incorrect
  verdict (`MRC_REVIEW_PROTOCOL.md` §8).
- `VERIFIED` means the reviewer independently reproduced the derivation
  within the stated checks. It does not mean the physical model is
  correct, empirically supported, or universally valid.
- Each finding uses Where/What/Why/Minimum-fix; each review closes
  with an evidential boundary. Reviews are append-only with pinned
  digests; a changed target stales the review.

## Formalization rules

- The Formalizer transcribes only human-authorized, reviewed
  derivations, and audits capability first: every required rule or
  structure is available or `UNAVAILABLE`, never assumed.
- PhysMath takes only theory-neutral mathematics (brackets, transport
  algebra, substitution/evaluation machinery). EC or any physical
  interpretation belongs in a phys-lib framework specialization, never
  in PhysMath or `core/`.
- PhysArtifact mints identities and provenance links only for
  formalized objects. `phys` is frozen legacy; never a destination.
- `physctl validate` is structural only; `physctl check` runs only
  allowlisted profiles backed by real infrastructure (needs the
  untracked local `go.work`, else `UNAVAILABLE`). Neither establishes
  mathematical correctness.

## Human-authorization boundary

Agents may propose, challenge, revise, and test. Only the human
declares a derivation ready for formalization and only the human
accepts scientific results (merge into `accepted`). A derivation
marked `UNSUBSTANTIATED` or `FALSIFIED` must not be formalized;
`CONDITIONAL` derivations formalize only with their conditions
carried explicitly, after human authorization.

## Worked example: D-002 readiness audit (2026-11-05)

Target: `derivations/D-002.md` (RULE-T1 investigation) as examined in
`research/reviews/R-003.md`. Method: read-only inspection of
`github.com/PithomLabs/phys-math` (rule IDs, geometry), `phys-lib`
(`gr/`, `core/`, classifications), `phys-artifact` (identity API).
No library file was modified.

| D-002 derive | Disposition (R-003) | Formalization need | Availability |
|---|---|---|---|
| pic_trigger | CONDITIONAL | Table substitution into fixed combinations | ABLE IN PRINCIPLE (SUBSTITUTE-001 exists); not authorized, mapping conditional |
| pair_classes | CONDITIONAL | Same substitution machinery | ABLE IN PRINCIPLE; same conditions |
| ec_hamiltonian | CONDITIONAL | Term evaluation incl. singular-function semantics | PARTIAL: evaluation machinery exists; constraint/λ semantics need phys-lib structures that do not exist |
| hbar_governs | CONDITIONAL | None (modus ponens on source rule, not a computation) | NOT A FORMALIZATION TARGET |
| ruleT1_form | UNSUBSTANTIATED | Poisson/commutator algebra + CQ kernel structures | UNAVAILABLE: no Poisson, Dirac-bracket, Lindblad/CPTP, or commutator machinery in phys-math/phys-lib (verified by search) |
| no_free_closure | CONDITIONAL | Same bracket machinery + reduced-transcription premise | UNAVAILABLE (same) |
| gen_constraint | UNSUBSTANTIATED | Above plus defined ClassicalTransport operator | UNAVAILABLE |
| obstruction | UNSUBSTANTIATED | None (no checkable result) | NOT A FORMALIZATION TARGET |

EC physical meaning: no EC framework exists in phys-lib (`gr/` is
metric GR; `core/` holds only `doc.go`; torsion appears solely as
torsion-free assumptions/notes). A new phys-lib specialization would
be required — not started, not authorized.

PhysArtifact: identity minting is mechanically available, but there
is nothing authorized to mint.

Audit verdict: NOTHING in D-002 is ready for formalization. The
three CONDITIONAL table-evaluation derives are the closest; they
await (a) human authorization, (b) EC-corner mapping confirmation,
(c) a formalization target decision. Everything else is blocked on
missing library machinery (brackets, CQ structures), missing
mathematics (coefficients, operators, kernel), or review findings
F1–F4.
