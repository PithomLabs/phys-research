# Plan 4 — Post-MVP-1 corrections: source fix, review input pinning, fresh review, protected remote

Status: **APPROVED — A → B → C authorised for execution; D deferred.**
Scope: `phys-research/` only. Supersedes nothing in `plan3.md`; extends it.

Human approval decisions on the four open questions:

| # | Question | Decision |
|---|---|---|
| B.1 | Material-input rule | **Approved**: one-hop neighbours in both directions, excluding review records |
| B.2 | Legacy R-001 with no input digests | **Approved**: warning, exit 0; label `LEGACY_UNPINNED`, never "fully fresh" |
| B.3 | Documenting the rule in `AGENTS.md` | **Approved** to document; agent permission restrictions must not be weakened — a **human-applied patch** is supplied instead of an agent edit |
| D | Remote setup | **Deferred**, human action required; no remote creation, credential change, `git config`, or acceptance merge |

Two additional mandatory safeguards added at approval:

- **Safeguard 5 — never silently skip a material input.** `physctl new
  --type review` must **fail with an explicit diagnostic** if any material
  neighbour cannot be read or parsed. A review that appeared complete while
  missing one of its sources is exactly the provenance failure this plan
  exists to prevent.
- **Safeguard 6 — one freshness rule.** `physctl validate` and `physctl
  status` must share the same review-freshness logic. A changed target digest
  *or* a changed material-input digest must produce `STALE` in both.

Authorizing decisions (from the human, verbatim in intent):

1. Leave the local Git hook disabled. Do not enable it or modify Git
   configuration. It is optional defense in depth, not a security boundary.
2. Accept the MICROSCOPE research records as a first-pass research case with
   open adversarial findings. Do not mark H-001 as an established hypothesis
   or physical result.
3. Correct F6: the eta definition in P-001 omits the factor of two in Eq. (1).
   Preserve the original record's history and record the correction
   transparently.
4. Preserve R-001 unchanged. Its findings F1–F9 remain part of the research
   history. Obtain a new adversarial review when the corrected source context
   is available.
5. Ensure reviews identify the exact versions of all material inputs they rely
   on, not just the hypothesis digest. At minimum, record the relevant source
   and evidence digests so that a changed input cannot silently leave a review
   appearing current.
6. Leave F1 and F2 open. Do not invent or automatically resolve a numerical
   falsification threshold or model-specific force constraints.
7. Set up a credential-separated remote and protect the accepted branch.
   Agents must have no ability to bypass its protection. Human acceptance and
   merging happen outside the agent-controlled environment.
8. Do not expand MVP 1 or begin Branch C implementation. Preserve the existing
   Go implementation and avoid unnecessary refactoring.

Governing distinction, to be preserved in every record and report:

> Accepted research record ≠ accepted scientific hypothesis ≠ established
> physical result.

H-001 remains **proposed and contested**, with its open findings preserved.

---

## 0. Facts verified read-only before planning

These were checked against the repository and the extracted paper text; no
file was modified while planning.

- **Eq. (1) carries the factor 2.** `paper-cache/2209.15487v1.txt` renders
  Eq. (1) as `η_{A,B} = 2(a_A − a_B)/(a_A + a_B)` (p. 2). The paper
  separately states, on p. 5: *"we identified the measured, approximated
  Eötvös ratio δ with the exact one η"*. With `a ≈ g·(mg/mi)`, the factor 2
  is exactly what makes `η ≈ δ`; without it `η ≈ δ/2` and the identification
  fails. F6's source-side diagnosis is therefore correct.
- **The defect lives at P-001 line 24**: `- The Eötvös ratio is defined as
  η_A,B = (aA − aB)/(aA + aB) (locator: Eq. (1), p. 2)`.
- **R-001 F6 already authorises this fix**: its minimum fix reads *"Restore
  the factor 2 in the definition, or state the convention explicitly and
  reconcile it with the δ ≈ η identification; correct the source record's
  transcription if appropriate."*
- **Editing P-001 does not stale R-001**: R-001 pins only H-001's digest
  (`sha256:ee747151…746cd`). Editing H-001 *would* stale it — which is why
  H-001 stays byte-identical.
- **No validator rule rejects extra `##` sections** (`checkSection` only
  requires that named sections exist and are substantive), so a
  `## Corrections` block in P-001 is safe.
- **Warnings do not fail validation**: `cmd/physctl` returns exit 1 only when
  `res.Errors > 0`. A missing-input-digest warning can therefore be visible
  without breaking the clean run.
- **Remote setup is not executable as things stand**: no remote is configured;
  `gh` is installed but its stored token is **invalid** (account `ibmendoza`);
  a single SSH key exists (`~/.ssh/id_ed25519.pub`). Credential separation
  requires two distinct credentials, which only the human can mint.

---

## A. Correct F6 in the source record (P-001 only)

Goal: the source record stops mis-transcribing Eq. (1), transparently, with
history preserved.

1. Rewrite the line-24 claim so the definition reads
   `η_A,B = 2(aA − aB)/(aA + aB)`, locator unchanged
   (`locator: Eq. (1), p. 2`).
2. Add a `## Corrections` section **after** `## Claims` containing:
   - the date of correction;
   - the previous text and the corrected text, quoted;
   - attribution to `R-001`, finding F6;
   - a sentence noting that the paper's δ ≈ η identification (p. 5) is what
     the factor 2 makes consistent, so the correction removes an inconsistency
     rather than changing any measured number.
   Git history preserves the original bytes; the section makes the correction
   visible to any reader of the current file.
3. **Do not touch H-001.** Its own `define eta_tipt` still omits the factor 2.
   That half of F6 stays open, exactly as R-001 recorded it. Editing H-001
   would invalidate R-001's digest and is forbidden by decision 4.
4. **Do not touch R-001.**
5. Commit separately and descriptively, e.g.
   `fix(source): P-001 restores factor 2 in Eq. (1) transcription (R-001 F6)`.

### A. Verification

- `physctl validate` → `0 error(s), 0 warning(s)` (or, if decision B.2 selects
  the warning, `0 error(s), 1 warning(s)`, exit 0).
- `sha256sum research/hypotheses/H-001.md` still
  `ee747151ab3b7ab3ede6b96fbcfd4115a83167b35fecd790912f0991c3d746cd`.
- `sha256sum research/reviews/R-001.md` unchanged from its pre-A value
  (record that value before starting).
- `git show` of the commit shows exactly one claim line changed plus the new
  section.

---

## B. Pin material input versions in reviews (decision 5)

Goal: a changed input can no longer leave a review silently appearing current.

### B.1 Which records count as a material input

Recommended rule: **every one-hop neighbour of the review target, in both
directions, excluding `type: review` records.**

For H-001 that resolves to exactly:

| Input | Role |
|---|---|
| `P-001` | source the hypothesis depends on (`depends_on`) |
| `I-001` | issue it addresses (`addresses`) |
| `E-001` | evidence pointing at it (`supports`) |

This satisfies "at minimum, the relevant source and evidence digests" and also
catches a rewritten problem statement. Excluding reviews prevents
digest-of-digest chains. *Alternative (not selected unless the human says
otherwise): source and evidence types only.*

### B.2 Behaviour for reviews created before this change (**approved**)

R-001 must remain byte-identical and therefore has no input lines. The
approved behaviour is to surface the gap rather than stay silent:

- `input:` lines present → each is checked strictly (below).
- no `input:` lines at all → **warning**, exit 0, and the review is classified
  **`LEGACY_UNPINNED`**:

  `LEGACY_UNPINNED: review records only the target digest; material input versions are not pinned`

`LEGACY_UNPINNED` means: a valid historical review of the exact target version
it examined, whose original source context was **not** fully pinned. It must
never be described as fully fresh, and the warning must never be described as
a clean review. R-002 provides the complete provenance.

### B.3 Code changes (minimal, no refactoring)

1. `internal/record/record.go` — the review template (`case TypeReview`)
   gains one line per material input immediately after `digest:`:

   ```
   reviewed: H-001
   digest: sha256:<target hex>
   input: P-001 sha256:<hex>
   input: I-001 sha256:<hex>
   input: E-001 sha256:<hex>
   reviewer: adversary (OpenCode agent role)
   model: <self-reported model id used for this review>
   ```

   Inputs are resolved at `physctl new --type review --review <ID>` time,
   alongside the existing target-digest resolution (same `store.Load` pass).

   **Safeguard 5 (mandatory): never silently skip a material input.** If a
   material neighbour cannot be read or parsed, creation **fails** with an
   explicit diagnostic, for example:

   - `material input P-001 not found` (no such record);
   - `material input P-001 could not be parsed: <parser message>` (file exists,
     frontmatter rejected);
   - `material input P-999 could not be resolved` (ref points nowhere).

   Only records explicitly excluded by B.1 (reviews) are omitted, and that
   omission is a rule, not a silent failure.

2. `internal/validate/validate.go` — `validateReviews` additionally parses
   `input:` lines:
   - unknown input ID → **error** (`review`, "input P-999 not found");
   - input digest ≠ current digest of that record → **stale** error naming
     the input, e.g.
     `STALE: input P-001 digest sha256:… does not match current sha256:…; write a new review`;
   - malformed line (bad ID or non-64-hex) → **error**;
   - zero `input:` lines → **warning** per B.2.

3. **Safeguard 6 (mandatory): one shared freshness rule.** `internal/validate`
   and `internal/status` must not each keep their own copy of the freshness
   logic. A single function (in `internal/store`, which both already depend on)
   parses `reviewed:` / `digest:` / `input:` lines and reports:

   - target digest mismatch → stale;
   - any recorded input digest mismatch → stale, **naming that input**;
   - recorded input that no longer resolves → stale/unknown, naming it;
   - no `input:` lines → `LEGACY_UNPINNED` (warning, not stale).

   Both commands call it, so `physctl status` and `physctl validate` can never
   disagree: a changed target **or** a changed material input produces `STALE`
   in both.

4. `AGENTS.md` — documentation is approved, but `AGENTS.md` sits on the
   project's edit-deny list and those restrictions must not be weakened.
   The required wording is delivered as a **human-applied patch** file outside
   the repository instead of an agent edit.

### B.4 Tests

- `internal/record`:
  - review template emits one `input:` line per material input, with the
    correct digests, and still emits `reviewed:`/`digest:`;
  - **safeguard 5**: creation fails with an explicit diagnostic when a
    material neighbour is missing or unparsable — asserted for both cases;
  - a `type: review` neighbour is excluded by rule and does not fail creation.
- `internal/store`: shared freshness function — target mismatch, input
  mismatch (names the input), unresolvable input, legacy no-input case.
- `internal/validate`: (a) fresh review with correct inputs → no findings;
  (b) changed input → stale error naming the input; (c) unknown input ID →
  error; (d) review with no input lines → exactly one `LEGACY_UNPINNED`
  warning, 0 errors; (e) target-only change still reported stale as before.
- `internal/status`: stale report appears when **only** an input changed,
  matching `validate` for the same fixture (safeguard 6).
- Existing tests must keep passing unchanged; adjust only assertions that
  hard-code the old review template.

---

## C. Fresh review of the corrected source context

Goal: an Adversary pass that sees the corrected P-001 and records the corrected
source digest.

1. **Order: A → B → C.** Both the source correction and the input-pinning code
   must land first, otherwise R-002 would pin the pre-correction P-001 or carry
   no input lines at all.
2. `physctl new --type review --title "…" --review H-001` creates `R-002`,
   which now records H-001's digest plus `input:` lines for P-001
   (corrected), I-001 and E-001.
3. Spawn the Adversary subagent on `opencode/deepseek-v4.1-flash-free`, with
   these explicit instructions:
   - review H-001 against the **corrected** P-001;
   - R-001 is research history and must not be edited in any way;
   - **F1 (falsification threshold is a free parameter) and F2 (`eta_tipt`
     conflates parameter and estimate) remain open blockers** — do not resolve,
     soften or numerically instantiate them;
   - the source half of F6 is corrected in P-001; the H-001 half (its `define`
     still omits the factor 2) is **still open** and should be recorded as such;
   - F3–F9 findings from R-001 carry forward as context; re-assess rather than
     assume;
   - never state that H-001 is correct, verified or established; no global
     verdict; close with `## Evidential boundary`.
4. Fill only R-002's placeholder sections. Do not edit H-001, I-001, P-001,
   E-001 or R-001 during the review.

### C. Verification

- Before C: record `sha256sum` of H-001 and R-001.
- After C: both hashes identical; only `R-002.md` and the A/B code changes are
  new in `git status`.
- `physctl validate` → 0 errors. R-002 must be fully fresh (target **and**
  all inputs pinned). R-001 must be reported exactly `LEGACY_UNPINNED` — never
  "fresh", never "clean"; the warning is not a clean review.
- `physctl status` shows F1/F2 as blocking findings from whichever review
  records them, and H-001's two `?` obligations still listed.

---

## D. Protected remote (decision 7) — **DEFERRED, human action required**

**Do not attempt D.** No remote creation, no credential changes, no
`git configuration changes`, no acceptance merge. The agent must not create or
borrow human credentials. Acceptance merges stay on hold until remote
protection is actually established **and** both the positive and negative
credential tests below pass.

Current state: **no remote exists**, the only `gh` token is invalid, and one
SSH key exists. Nothing here can be executed by an agent today, and
`.opencode/opencode.json` denies agents `git push *` and `git config *` by
design.

Required from the human before any of D can run:

1. A hosting target (module path implies `github.com/PithomLabs/phys-research`,
   but no repository exists yet) and a public/private decision.
2. A working **human** credential, refreshed (`gh auth login` is currently
   failing).
3. A **separate agent credential** — a second, distinct identity, never a
   reuse of the human's. Two distinct credentials are the minimum that makes
   "credential-separated" true.
4. A decision on who executes: the human runs the pushes, or the human
   explicitly amends the agent permission rules for named commands. Amending
   the rules is a human act; it must not happen implicitly.

Target configuration:

- `main` — working branch, agents may push.
- `accepted` — protected: no direct pushes, merging requires the human's
  approving review, agent credentials cannot satisfy or bypass it.
- Agent credential: no administration rights, cannot modify the protection.
- Human credential: authorised to review and merge.

### D. Verification required before anyone may claim the remote is protected

1. Read back the protection object for `accepted` from the host API and record
   its actual fields (required reviews, push restrictions, admin enforcement).
2. **Negative test**: attempt `git push` of `main:accepted` using the *agent*
   credential → must be **rejected**.
3. **Positive test**: push to `main` using the agent credential → must
   **succeed**, proving the failure in (2) is protection and not a broken
   credential.
4. Confirm the agent credential has no administration scope.

Only after all four pass may the report say `accepted` is protected. Until
then the correct statement is: *no protected remote exists; the acceptance
merge has not been performed.*

---

## E. Reporting format

Four separate sections, never blended:

1. **Source correction** — what changed in P-001, the old and new text, the
   commit, and the verification hashes (H-001 and R-001 byte-identical).
2. **Code changes** — the input-pinning implementation, the two safeguards,
   the shared freshness rule, tests actually run and their results, and the
   human-applied `AGENTS.md` patch.
3. **Fresh review** — R-002's target and input digests, its findings, and an
   explicit statement of which R-001 findings remain open (F1, F2 blocking;
   the H-001 half of F6; F3–F5, F7–F9 requiring independent reassessment).
   State R-001's classification as `LEGACY_UNPINNED`.
4. **Deferred security setup** — exactly what is and is not enforced, with no
   claim of protection until D's four checks pass.

Every report must restate: H-001 is proposed and contested; structural
validation is not scientific correctness; `physctl check` PASS means only that
the named tests passed. **Never describe a warning as a clean or fully fresh
review.**

---

## Explicit non-goals

- No Branch C work, no new features, commands, roles, agents, databases,
  services or dependencies.
- No edits to H-001 or R-001.
- No numeric falsification threshold, no model-specific force constraint, no
  resolution of F1 or F2.
- The local hook stays disabled; no `git config` changes; no enabling of
  `core.hooksPath`.
- No refactoring of the existing Go implementation beyond what B requires.

## Resolved decisions (formerly open)

| # | Question | Decision |
|---|---|---|
| 1 | Material-input rule | **Approved**: one-hop neighbours both directions, reviews excluded |
| 2 | Reviews with no input lines (R-001) | **Approved**: warning, exit 0; classified `LEGACY_UNPINNED`, never "fully fresh" |
| 3 | Documenting the rule in `AGENTS.md` §6.4/6.5 | **Approved** to document via a human-applied patch; agent permission restrictions not weakened |
| 4 | Remote host, public/private, and who executes | **Deferred** — human action required; D not attempted |

## Human baseline hashes (recorded before execution)

```
ee747151ab3b7ab3ede6b96fbcfd4115a83167b35fecd790912f0991c3d746cd  research/hypotheses/H-001.md
7ee684efba8deb9704d707e4511a2ad0493b4028eb2f4e8e4e1283a3e787aaed  research/reviews/R-001.md
749770a1da211351ba7723fb2fa02ce139582aaad773f05794df1890cdf1ff61  research/sources/P-001.md   (pre-correction)
ba671029bb0cd9b43e835e9a78eb461c144011e4b0060928ac8291a89ede485a  research/issues/I-001.md
c3cb71fe45c2a93a12781b9f1fff05c07c72d1f74f8444a605134c36924044ac  research/evidence/E-001.md
```

H-001 and R-001 must match these values at the end of A → B → C.
