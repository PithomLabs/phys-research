# Proposed AGENTS.md patch: Phys-Pseudo formalization gateway

For human application. Agents must not modify `AGENTS.md` themselves.
Append the following as a new top-level section at the end of
`AGENTS.md` (after the current final section):

---

## 13. Phys-Pseudo formalization gateway

No derivation may be formalized into PhysMath, PhysLib, or
PhysArtifact except through this pipeline, in order:

Researcher investigation → Phys-Pseudo record → Adversarial review
→ Human authorization → Formalization readiness audit →
library formalization.

1. The investigation record uses canonical `PSEUDO.md` syntax:
   named constructs, non-empty `from [NAMES]` lists, explicit
   `derive` results. `derive` appears only where an actual
   inference and result are supplied; proposed forms and undefined
   objects stay in `claim`/`?`.
2. The adversarial review gives each consequential `derive` its own
   disposition (`VERIFIED`, `CONDITIONAL`, `UNSUBSTANTIATED`, or
   `FALSIFIED`) in the review record only, with Where/What/Why/
   Minimum-fix findings and an evidential boundary. No global
   verdict. `VERIFIED` means independently reproduced within the
   stated checks — never physical truth or empirical support.
3. A failed review returns the investigation to Phys-Pseudo.
   `UNSUBSTANTIATED` or `FALSIFIED` derivations must not be
   formalized. `CONDITIONAL` derivations formalize only with
   conditions carried explicitly and only after human authorization.
4. The Formalizer transcribes only human-authorized derivations and
   audits capability first; missing rules or structures are
   reported `UNAVAILABLE`, never assumed. PhysMath takes only
   theory-neutral mathematics; physical interpretation belongs in a
   phys-lib framework specialization. `phys` is frozen legacy.
5. Until a technical gate exists, this gateway is conventional,
   enforced by review discipline and human acceptance — see
   `docs/PHYS_PSEUDO_GATEWAY.md`. Do not describe it as
   technically enforced.

---
