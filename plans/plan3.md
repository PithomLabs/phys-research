# MVP 1 paper selection — for your approval

## 1. Identification

- **Title:** *MICROSCOPE mission: final results of the test of the Equivalence Principle*
- **Authors:** P. Touboul, G. Métris, M. Rodrigues, J. Bergé, A. Robert, Q. Baghi, et al. (MICROSCOPE Collaboration)
- **Publication:** Phys. Rev. Lett. **129**, 121102 (2022)
- **DOI:** `10.1103/PhysRevLett.129.121102` · **arXiv:** `2209.15487v1` (gr-qc), open access, HTML/PDF verified readable

## 2. Why it is suitable

- Full text freely accessible (I already read the arXiv HTML to confirm locators below).
- Clean **theoretical expectation** (WEP ⇒ η = 0 exactly, GR, with stated assumptions) vs **experimental actual** (η(Ti,Pt) = [−1.5 ± 2.3(stat) ± 1.5(syst)] × 10⁻¹⁵, dimensionless, 1σ).
- Explicit regime (low-Earth orbit, Ti/Pt alloys, free-fall sessions) and explicit assumptions (δ ≈ η identification, thermal/glitch systematics treatment).
- A straightforward, two-sided falsification condition (a reproducible ≥3σ nonzero η falsifies the null hypothesis; a re-analysis showing |η| above the stated bound falsifies the bound claim).
- Zero new mathematical infrastructure needed — MVP 1 is pure tracking.
- Naturally adversarial material: the δ→η approximation (Eq. 2), the thermal-systematics extrapolation, and glitch masking are all attackable.

## 3. Proposed research issue (I-001)

> **What bounds does the MICROSCOPE final η(Ti,Pt) constraint place on long-range composition-dependent new forces, and which stated assumptions (δ ≈ η identification, thermal systematic upper bound, glitch masking) limit that bound?**
> State: open · Priority: high · ref: `motivated_by P-001`

with the four required sections (theoretical expectation / independent theoretical result / experimental expectation / experimental actual) — the *independent theoretical result* section will explicitly say **"not established — no independent derivation has been performed"**, keeping the paper's assertion from being laundered into one.

## 4. Proposed Phys-Pseudo hypothesis (H-001) — draft

```text
# H-001 — WEP null at 10^-15 for Ti/Pt. Source: Touboul et al. 2022, arXiv:2209.15487v1.

given eta_def: eta(A,B) = 2 (a_A - a_B)/(a_A + a_B)  # paper Eq. (1)
given eta_data: measured eta(Ti,Pt) = [-1.5 +/- 2.3 (stat) +/- 1.5 (syst)] x 10^-15 at 1 sigma  # paper Eq. (5)
assume wep: universality of free fall holds for TiAlV(90/6/4) and PtRh(90/10) in the Earth-orbit regime
claim null: the reported data are consistent with eta(Ti,Pt) = 0 within the stated uncertainties  # paper assertion, not an independent derivation
claim bound: any long-range mediator producing a Ti/Pt differential acceleration must yield |eta| below the reported sensitivity in this regime
? whether the thermal systematic upper bound (factor-500 attenuation extrapolated above 10^-3 Hz) is sufficient to support the bound claim; would be resolved by the thermal transfer-function analysis of Touboul et al. 2022b
```

Seven constructs only; attribution is prose/`#` comments per PSEUDO.md §8; the `?` is a legitimate §6 obligation. The **falsification condition lives in record prose/metadata, not as an eighth construct** (Qwen's `FalsificationCondition` syntax stays rejected).

## 5. Falsification condition

- **Physical (H-001):** falsified by a reproducible, independently analyzed measurement of |η(Ti,Pt)| ≥ 3σ ≠ 0 with systematics below the statistical error, **or** by a re-analysis of MICROSCOPE data yielding |η| > 4.5×10⁻¹⁵ (≈ combined 1σ envelope) with the thermal/glitch systematics shown to be misestimated.
- Mathematical obligations (if any later) instead require premises/conclusion/domain/proof obligations — no empirical falsifier.

## 6. Precise source locators (all verified in the arXiv HTML today)

| Claim | Locator |
|---|---|
| Result value | Abstract; **Eq. (5)** |
| η definition | **Eq. (1)**; approximation δ(A,B) **Eq. (2)** |
| Measurement model | **Eq. (3)**, fitted model **Eq. (4)** |
| Per-segment estimates, δ(Pt,Pt)=(0.0±1.1)×10⁻¹⁵, δ(Ti,Pt)=(−1.5±2.3)×10⁻¹⁵ | **Fig. 1** (upper/lower panels) |
| Test-mass compositions/masses/densities | **Table 1** |
| Spin/orbital/EP frequencies | **Table 2** |
| Off-centerings & [Ad] matrix | **Table 3** |
| Systematics budget (SUEP bound 1.5×10⁻¹⁵ vs spec 0.2×10⁻¹⁵; temperature dominant; factor-500 attenuation) | paragraph beginning "We found that the overall systematics upper bound…" |
| Glitch treatment (4.5σ/3σ clipping, M-ECM) | paragraph beginning "As with glitches…" / "We tackled this difficulty…" |
| Prior bounds (2×10⁻¹³, 7.1×10⁻¹⁴, 2×10⁻¹⁴) | Introduction, paragraph 3 |
| 10⁻¹⁷ upgrade target | Closing paragraph |

Exact page/section numbers of the published PRL pagination will be pinned against the PDF during implementation (the source hash of the PDF will be recorded; the PDF itself will not be committed).

## 7. What the Adversary should independently challenge

1. Whether H-001 merely restates the paper's claim (citation laundering) rather than standing as an independent statement.
2. The **δ ≈ η identification** (Eq. 2 → Eq. 5) and its regime validity.
3. The thermal systematic upper bound: factor-500 attenuation extrapolated above 10⁻³ Hz where probes are below noise.
4. Whether glitch masking + M-ECM could bias a signal at f_EP (the paper's own mock-signal check is the counter-evidence to test).
5. Operability of the falsifier (what precision, what experiment).
6. Regime mismatch: Ti/Pt in low-Earth orbit vs the generality of any "new force bound" wording.
7. Missing `?` obligations (e.g., the unpublished MCMC cross-check referenced in the text).

---

# Implementation plan

## Phase 0 — Reconnaissance (completed, findings recorded)

Verified facts the implementation must respect:

| Finding | Consequence |
|---|---|
| PSEUDO.md §10: "No checker, linter, or parser exists or is implied" | `validate` never parses Phys-Pseudo; only metadata + line-level `?` obligation counting (a `?` line is reported as an obligation, never an error) |
| phys-artifact has **no artifact type** for research records or `.pseudo` | Records use readable IDs; SHA-256 only as a *review snapshot digest*. Documented as a gap — no new artifact family invented |
| No research-ID convention exists anywhere in the Phys repos | Adopt the task's `P-/I-/H-/R-/E-/D-` type prefixes (source→P, issue→I, hypothesis→H, review→R, evidence→E, derivation→D); `branch: Q\|C` stays a separate field |
| `phys-math`/`phys-lib` do **not** build standalone (missing `require`s, no `go.work`) | I ran the experiment: a `go.work` + `GOWORK` env makes phys-math's fast gate tests and `phys-lib/gr/equivariance` tests **pass**. `physctl check` uses a committed `go.work` and allowlisted `go test` profiles (modeled on `phys-trace/cmd/phystrace-mcp/verify.go`) |
| GUIDELINES §9/§10 (capability audit; oracle must not come solely from the implementation under test) | `check` runs only pre-declared profiles; unknown/missing capability → explicit `UNAVAILABLE` diagnostic, never a PASS. Formalizer's MVP-2 duty includes a capability audit |
| OpenCode v2.2.5 docs: agent frontmatter supports `description/mode/model/permissions`; **subagents use their own permissions**; last matching rule wins; shell resources are raw command text | Permissions are configured per-agent, not inherited; `git merge *`/`git push *` denies are expressible |
| All existing Go repos are stdlib-only (enforced by a test in phys-artifact) | Restricted hand-written frontmatter parser, zero dependencies (per your instruction) |

## Phase 1 — Repository scaffold: `TOE-v5/phys-research/` (own git repo)

```text
phys-research/
├── AGENTS.md                  # constitution: shared rules, metadata subset spec, ID/vocab, branch policy, "validation ≠ truth"
├── go.work                    # use ../../../Documents/go/{phys-artifact,phys-math,phys-lib}  (verified working)
├── .gitignore                 # local PDF cache, build artifacts
├── hooks/
│   └── pre-merge-commit       # refuses merges into `accepted` unless PHYS_HUMAN_MERGE=1 (second layer; human runs `git config core.hooksPath hooks` once)
├── .opencode/
│   ├── opencode.json          # project-wide permission denies (applies to every agent incl. build)
│   └── agents/{researcher,adversary,formalizer}.md
├── research/{sources,issues,hypotheses,reviews,evidence}/
├── derivations/
└── cmd/physctl/
```

**Git policy:** `main` = working branch (agents commit here / on `research/*`); `accepted` = protected. Agent-facing enforcement, in layers: (1) project `opencode.json` denies `shell: git push *`, `git merge *`, `git config *` and `edit` on `AGENTS.md`, `.opencode/*`, `cmd/*`, `go.mod`, `go.work`, `hooks/*`; (2) the committed hook refuses agent merges into `accepted` (requires `PHYS_HUMAN_MERGE=1`, set only in your terminal); (3) you perform the actual acceptance merge outside OpenCode with your own credentials. No formal PRs for routine notes.

**Permission rules per agent (frontmatter, last-match-wins):**

| Agent | edit permissions | model |
|---|---|---|
| researcher | allow `research/sources/*`, `research/issues/*`, `research/hypotheses/*`, `research/evidence/*`; deny `research/reviews/*`, `derivations/*` | `opencode/mimo-v2.6-flash-free` |
| adversary | deny `*`, then allow `research/reviews/*` only — cannot touch the hypothesis it reviews | `opencode/deepseek-v4.1-flash-free` (fresh context, different model) |
| formalizer | deny `*`, then allow `derivations/*` (MVP 2) | inherits session model |

## Phase 2 — Go CLI: one binary, five commands (stdlib only)

Module `phys-research` (module path — approval item below). Packages: `cmd/physctl` + `internal/{frontmatter,record,store,validate,refs,check}`.

**Record format** — `<TYPE>-NNN.md` filenames, frontmatter subset exactly as you specified (`id`, `type`, `state`, `branch`, `priority`, `refs[{relation,id}]`), line-based parser rejecting unknown fields, duplicates, bad indentation, anchors/aliases, multiline strings, nested objects — with explicit errors, no guessing. Body untouched.

| Command | Behavior |
|---|---|
| `new` | `physctl new --type issue [--branch C] [--priority high] [--ref motivated_by=P-001]`. Allocates next ID by scanning existing IDs of that type (max+1, collision-safe via `O_EXCL` + atomic temp-file rename). Writes type template: issue → four expectation/actual sections; hypothesis → Phys-Pseudo scaffold + falsification section; source → DOI/arXiv/source-hash + claim+locator table; review → auto-computed `reviewed_sha256` of the target + reviewer identity/model + blocking/nonblocking; evidence → value/units/uncertainty/methodology/locator; derivation → premises/deps/intermediate/rules/obligations. Refuses overwrite. |
| `status` | Open+blocked issues sorted by priority then ID; unresolved `?` obligations (line numbers); stale reviews; blocking review findings. `--json`, `--branch C`. |
| `refs` | BFS over declared `refs` only (never prose search), `--depth`, `--in` for incoming edges, cycle detection with the cycle path reported, deterministic ordering. `--json`. |
| `validate` | Frontmatter subset; required fields; enum validity (type/state/priority/branch); ID↔type↔filename agreement; unique IDs; refs exist + relation ∈ closed vocabulary {`motivated_by`,`supports`,`challenges`,`addresses`,`depends_on`,`supersedes`} + no self-refs + cycle detection; source-locator presence on every extracted claim and evidence record; falsification section required for physical hypotheses / premises+conclusion+domain+proof-obligations required for `class: mathematical`; review digest match → **STALE** if the hypothesis changed; append-only check via `git log --diff-filter=M` on `reviews/` when git history is available; derivation reference integrity. Output ends with an explicit statement: *structural validation only — no mathematical or physical correctness is implied.* Exit 0/1/2 (clean / findings / usage-IO). `--json` emits findings as `{rule, severity, file, message}`. |
| `check` | Allowlisted profiles only (fixed in Go, e.g. `phys-math-fast` = the three gate tests, `phys-lib-equivariance`, `phys-lib-ltb`, `phys-lib-bohm`), each with pinned dir, `-run` pattern, `-count=1`, timeout — run with `GOWORK=<repo>/go.work`. `physctl check derivations/<id>` runs only the profiles a derivation record declares against that allowlist. Missing repo/workspace/test failure → `UNAVAILABLE`/`FAIL` with the exact command and exit code; **never a false PASS**; never claims a test ran that didn't. |

## Phase 3 — MVP 1 record set (using the approved paper)

`P-001` (source, claims with the locators above) → `I-001` (issue §3) → `H-001` (hypothesis §4, explicit falsifier §5) → `E-001` (evidence: value, units, stat/syst uncertainty, M-ECM/Adam methodology, 18 SUEP sessions, Eq. (5)/Fig. 1 locators) → `R-001` (Adversary review of the exact SHA-256 of `H-001`, MRC §8 finding format — Where/What/Why/Minimum fix — plus closing evidential-boundary paragraph) → **you** perform the acceptance merge into `accepted`.

## Phase 4 — Tests (in `t.TempDir()`, adversarial inputs included)

- Frontmatter: valid / empty refs / multi refs / malformed indentation / duplicate keys / unknown fields / missing fields / anchors, `|`, `>` rejected / CRLF, tabs, unicode, `id: ../../etc` path-traversal attempt.
- ID allocation: sequential, collision skip, concurrent `O_EXCL` (two goroutines).
- Record creation: templates, overwrite refusal, atomic write.
- Validation: broken ref, cyclic refs, enum violations, ID-type mismatch, missing locator, missing falsifier, stale review, modified (non-append) review.
- refs: depth limiting, cycles, incoming direction, deterministic output.
- status/JSON: shape assertions, priority ordering, nonzero exit codes.
- check: unknown profile → diagnostic + nonzero; workspace absent → `UNAVAILABLE`; **no test that just restates an implementation assumption as its expected result**.
- Integration: build the binary, drive `new → validate → refs → status --json` end-to-end against a temp git repo.

## Phase 5 — Execution & report

`go build`, `gofmt -l .`, `go vet ./...`, `go test ./...`, then inspect the final git diff; report per §13 (implemented / files changed / tests actually run with results / limitations / decisions needing your approval).

---

# Decisions I need from you

1. **The paper** (§1–7 above) — approve or replace.
2. **A 7th optional metadata field `class: physical|mathematical`** on hypotheses — required so `validate` can demand a falsifier only for physical hypotheses without inventing inference. (Alternative: no new field and validate requires *both* sections always — sloppier but strictly within the six listed fields.)
3. **Module path:** `github.com/PithomLabs/phys-research` (matches existing repos) or plain `phys-research`.
4. **Priority enum:** `high|medium|low`?
5. Confirm the **one-time human setup**: `git init` + branch layout, `git config core.hooksPath hooks`, and that *you* make the acceptance merge (agents are permission-denied from it).

Once you approve (or amend), I'll switch out of Plan mode and implement phases 0→5.
