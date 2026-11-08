# 02 — Phys-Pseudo contract

Authoritative source: `/home/chaschel/Documents/go/phys-pseudo/PSEUDO.md`
(sole spec; `CONVENTIONS.md` is a pointer note). Operational guide:
`phys-pseudo/AGENTS.md`. Review method: `MRC_REVIEW_PROTOCOL.md`
(8 responsibilities). Implementation gate:
`conformance/DERIVATION_SPEC_GUIDELINES.md` (byte-identical to the
`phys-trace/` copy; `diff -q` verified).

## 1. The seven constructs (exact syntax, PSEUDO.md §1)

```text
given NAME: <content>
define NAME := <meaning>
assume NAME: <premise>
claim NAME: <proposition>
derive NAME: <explicit result> from [NAMES] | <reasoning>
if <condition>:
    then <consequence>
    else <alternative>
? <what is unresolved and what would resolve it>
```

Nothing else is a construct. `# ...` lines are prose. No status, provenance,
or eighth construct exists. Status words (`PROVEN`, `VERIFIED`, …) as labels
are non-conforming (§7).

## 2. Binding rules (operationally load-bearing ones)

- **Names** (§2): plain identifiers, no argument lists; declaration before
  use; exact-match references; premise lists take ≥1 name (`from []`
  is not part of the language); branch-local names die with the branch.
- **Roles** (§3): `define` = meaning only (substitution test: replace RHS
  everywhere — if any proposition changes, it was not a definition);
  `assume` = adopted premise, not endorsed; `claim` = proposition under
  examination, never promoted by proximity; `given` = starting input.
  Confusing roles is a reportable defect.
- **`derive`** (§4): name + explicit result + non-empty premise list +
  reasoning sketch all required. The premise list is *proposed*, not
  complete — omitted premises are invisible by construction; finding them
  is the reviewer's central job. A `derive` here is asserted, never
  validated (validation = independent checking or later implementation).
- **Conditionals** (§5): branch statements hold only under the condition;
  branch names/conclusions are NOT referenceable outside (deliberate
  expressiveness boundary); continuing under a condition requires staying
  in the branch or restating it as a new top-level `assume`.
- **`?`** (§6): unresolved/underspecified ONLY, never "known false".
  Attachment is mechanical: statement-level = immediately follows target,
  indented exactly one deeper; document-level = column 0. Known-false is
  refuted via counterexample/contradiction `derive`, never `?`.
- **Invariant** (§7): representation is not endorsement, in either
  direction. Structure carries no verdicts.
- **Source fidelity** (§8): attribution is prose/`#` comments/citation
  handles only; `from` appears solely in `derive`; handles never in
  premise lists; no automatic `?` for uncited paper assertions.

## 3. Competing hypotheses, obligations, handoff

- Competing hypotheses: representable — see corpus fixture
  `conformance/08_competing_hypotheses/fixture.md` (dir verified present;
  fixture bodies intentionally contain non-conforming forms as adversarial
  inputs, per `conformance/README.md`). No new syntax needed.
- Unresolved obligations: `?` with what-would-resolve-it content (§6).
- Handoff to formalization: exploratory doc → MRC review (protocol's 8
  responsibilities: reconstruct target, separate roles, inspect
  derivations counterexample-first, proportional dimension/domain/limit
  checks, separate math-validity from applicability, `?`/refute
  discipline, findings with minimum correction, evidential-boundary
  statement, no global verdict) → if implementation candidate,
  `DERIVATION_SPEC_GUIDELINES.md` normative requirements on top.
  Syntactic clarity never implies implementation readiness (PSEUDO.md §10).

## 4. Implemented vs documented (per adjustment A: inspected, not inferred)

| Capability | Status | Evidence |
|---|---|---|
| Language definition (7 constructs) | IMPLEMENTED (as spec) | `PSEUDO.md`, `AGENTS.md` |
| MRC review method | IMPLEMENTED (as human/agent process) | `MRC_REVIEW_PROTOCOL.md` (8 responsibilities, read) |
| Implementation-ready requirements | IMPLEMENTED (as normative doc) | `DERIVATION_SPEC_GUIDELINES.md` (both copies) |
| Adversarial example corpus | IMPLEMENTED | `conformance/{01..08,S_cosmology_sketch}/fixture.md` |
| Grammar/parser/checker/linter/registry | MISSING **by design** | PSEUDO.md §10: "No checker, linter, or parser exists or is implied"; corpus README: "explicitly out of bounds" |
| Automated pseudo→derivation consumption | MISSING | No consumer in any `phys-*` module (no importer of pseudo; downstream tooling = `phystrace-mcp` review of Go fixtures, not of pseudo) |
| Status/promotion engine | MISSING by design | §7 bans author-typed status; no engine specified |

For QSG-T001 this means: Phys-Pseudo can EXPRESS renormalization
assumptions, competing source prescriptions, and conservation obligations
(no syntax gap), but every downstream step — review, handoff,
formalization — is human/agent-executed process, not automation. The
pseudo layer contributes expressiveness and discipline, not verification.
