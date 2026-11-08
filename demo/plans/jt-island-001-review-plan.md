# Plan: OpenCode Adversarial Review v1.0 for JT-ISLAND-001

## 1. Scope and deliverable

**Mode:** independent reconstruction and attempted falsification; review only.

**Final deliverable (when execution is authorized):**

`/home/chaschel/Desktop/biz/TOE-v6/research/experiments/jt-island-001/reviews/opencode-review-v1.0.md`

**Scratch workspace:**

`/home/chaschel/Desktop/biz/TOE-v6/research/experiments/jt-island-001-review/`

**Hard constraints**

- No Go code, no implementation contract, no architectural proposal.
- No modification of any file under `/home/chaschel/Documents/go`.
- No modification of Muse's package or pinned sources. The only addition inside the package is the required `reviews/` report.
- Dossier, audit v1.0, sources, and all existing package files remain byte-intact; verify by SHA-256 at start and end.
- One full review pass. No revision pass unless the director separately authorizes it.
- Stop after submitting the report. Do not freeze, approve implementation, or claim novelty.

---

## 2. Authority and evidence order

1. `prompt10.md` and the Phys Research Workflow Protocol v1.0.
2. `PSEUDO.md` and `MRC_REVIEW_PROTOCOL.md` for language conformance and review method.
3. Pinned primary sources: AMMZ v2, AEMM v3, AP v3, Penington.
4. Director rulings: `plan_prompt9_imp_review.md`, `audit-corrections-v1.1.md`.
5. Muse's specification, baseline, branch reconstruction, registers, and reality checks — candidate formalization, not ground truth.
6. Dossier and audit v1.0 — historical inputs only.

The package's own checklist is an input, not an authority. Verify its 18 items independently.

**Reading coverage:** normative Phys-Pseudo documents are read in full. The advisory `plan/` history, including the 12k-line `reality_check.md`, is inventoried and consulted for failure modes—dimensional analysis, boundary conditions, limits, symmetries, entropy, holography—but never treated as normative. Phase 0 completes any remaining advisory coverage needed for the report's methodology section.

---

## 3. The three governing questions

1. **Model identity:** Is the AEMM–AMMZ combination genuinely one coherent physical model, or a blend of compatible-looking but distinct setups?
2. **R4 mathematics:** Are the entropy normalization, island/no-island branches, QES input, and Page-time approximation mutually consistent?
3. **Interoperability:** Does R1 genuinely validate the background and matter assumptions consumed by R4, or is the claimed interoperability still partly assumed?

---

## 4. Execution phases

| Phase | Work | Main output | Gate |
|---|---|---|---|
| 0 | Provenance, hashes, PDF extraction, locator index, symbol and ID crosswalk | Evidence index | Hashes match; every locator resolved or flagged |
| 1 | AEMM–AMMZ model compatibility | Compatibility matrix | Every mapping classified: same / specialized / imported / incompatible / unproven |
| 2 | R1 independent reconstruction | R1 equation ledger | Every load-bearing equation derived, source-verified, or explicitly marked imported |
| 3 | R4 independent reconstruction | R4 equation ledger | Both branches, QES, Page condition, and \(t_P\) independently tested |
| 4 | OB-1 through OB-10 deep dives | Obligation matrix | Each obligation has evidence, status, consequence, correction |
| 5 | R1–R4 interoperability | Shared-object and dependency matrix | \(\kappa=k\), ten bridges, and one-way dependency tested |
| 6 | Register audit | Claim / assumption / bridge / obligation dispositions | Every C-, ASM-, BR-, OB-item accounted for |
| 7 | Three Reality Checks | RC1/RC2/RC3 assessment | Independent removal tests, not heading checks |
| 8 | Phys-Pseudo conformance | Conformance matrix | All seven-construct rules checked against `PSEUDO.md` |
| 9 | Limiting cases and falsification | Limit/counterexample table | L1–L6 plus additional adversarial tests |
| 10 | Findings, verdicts, report, QA | Final report | Every finding has all required fields; verdict consistent with findings |

---

## 5. Load-bearing reconstruction matrix

| Work package | Source anchors | Independent reconstruction | Pass criterion |
|---|---|---|---|
| R1.1 Action and variation | AEMM(6)–(8), (9), (11a–c); AP(5.3),(5.5) | Vary \(I_0+I_G+I_M\) with respect to dilaton and metric; derive \(R+2=0\), metric equations, \(16\pi\) vs \(8\pi\) pattern, signs, boundary terms | Every coefficient and sign reproduced from the action; no reliance on Muse's algebra |
| R1.2 Stress and anomaly | AP(5.1)–(5.2); AEMM(42),(44)–(46),(61) | Derive \(\langle T_{\mu\nu}\rangle\), trace anomaly, Ward identity, light-cone conservation, chiral flux; separate anomaly from diffeomorphism Ward identity | Anomaly and conservation shown as independent inputs; convention factor \(-2\pi\) pinned |
| R1.3 Boundary dynamics | AEMM(13)–(25) | Derive Schwarzian action, ADM energy sign, \(E_0>0\), energy balance, evaporation law, \(k\), \(T(t)\) | Sign of \(E(u)\) and flux orientation verified; \(k\) dimensionally checked |
| R1.4 Interface and regulator | AEMM(30),(44)–(46),(64) | Derive \(t=f(u)\), cutoff \(z=\epsilon f'(u)\), \(\phi_b=\bar\phi_r/\epsilon\), shock \(E_S\), smooth-limit conditions | Transparent-interface setup consistent; shock-vs-smooth branch identified |
| R4.1 \(S_{\text{gen}}\) and endpoints | AMMZ(10),(13); Decision 2 | Derive generalized entropy, count actual QES endpoints, check one-endpoint versus two-endpoint normalization and \(2S_0\) price | Endpoint count matches the actual adopted configuration and source |
| R4.2 No-island branch | AMMZ(14); AEMM(112) | Derive two-interval CFT entropy, log form, late-time slope; compare fixed-\(\beta\) log with the integral over decaying \(T(t)\); check saturation | Branch formula valid in its stated regime; slope and saturation independently reproduced |
| R4.3 Island branch and QES | AMMZ(12),(13); AEMM(75)–(77),(92)–(93),(104)–(107) | Derive \(S_{\text{gen}}\) stationarity, QES position, on-shell island entropy; check imported hypotheses and approximation regime | QES import not promoted to derivation; one-endpoint value correctly normalized |
| R4.4 Page condition and \(t_P\) | AMMZ §3.3; Penington Eq.(2) | Solve \(S_{\text{no}}(t_P)=S_{\text{isl}}(t_P)\); derive leading estimate \(6\beta S/\pi c\); test F1–F5 and slow-evaporation self-consistency | Result classified exactly as exact / leading / approximate / conditional |
| R4.5 Extremization vs minimization | AMMZ(15), §3.3 | Check ordering of QES extremization and branch minimization; test island-rule status and large-\(S_0\) branch | Operations are distinct and correctly ordered; no unsupported minimality claim |
| Interop | AEMM(20),(25); AMMZ(13); Part C | Derive \(\kappa=k\); build shared-object matrix; test dependency graph for cycles and hidden assumptions | Identity exact or conditional with stated conditions; R1 outputs genuinely consumed by R4 |

---

## 6. Finding and verdict rules

Every finding uses the prompt10 nine-field format:

```text
ID:
Severity: BLOCKER | MATERIAL | MINOR
Affected file and exact line/section:
Mathematical or physical issue:
Independent evidence or derivation:
Primary source and exact locator:
Consequence for the research conclusion:
Minimum correction required:
Affected area: R1 | R4 | interoperability | conformance | source/epistemic
```

Add one extra field, **Disclosure**, recording whether the issue is already disclosed as an OB item. This prevents double-counting known limitations while still allowing a disclosed load-bearing gap to be classified as MATERIAL or BLOCKER if it undermines a central claim.

**Verdict logic**

- **PASS:** no unresolved BLOCKER or MATERIAL findings.
- **CONDITIONAL:** no unresolved BLOCKERs, but MATERIAL qualifications or scope restrictions remain; state exactly what is permitted and what is restricted.
- **BLOCKED:** at least one unresolved BLOCKER prevents reliable approval of the affected specification or central conclusion.

Separate verdicts for:

1. R1 mathematical validity.
2. R4 mathematical validity.
3. R1–R4 interoperability.
4. Phys-Pseudo conformance.
5. Source fidelity and epistemic integrity.
6. Overall JT-ISLAND-001 readiness.

A `?` label alone does not resolve an obligation. A source-reported result does not become derived because the package says so. A disclosed obligation can still be load-bearing and therefore MATERIAL or BLOCKER if the central claim depends on it and the disclosure does not narrow the claim.

---

## 7. Report structure

Follow prompt10 §10 exactly:

1. Executive verdict.
2. Review methodology and source versions.
3. Independent mathematical reconstructions.
4. Findings classified BLOCKER/MATERIAL/MINOR.
5. Detailed assessment of OB-1 through OB-10.
6. R1 assessment.
7. R4 assessment.
8. Interoperability assessment.
9. Three Reality Checks assessment.
10. Phys-Pseudo conformance assessment.
11. Limiting-case and counterexample assessment.
12. Required corrections.
13. Accepted claims and their scope.
14. Rejected or narrowed claims.
15. Final verdict and approval recommendation.

Appendices: source hash manifest, exact locator map, symbol dictionary, equation ledger, bridge matrix, obligation matrix, finding index.

---

## 8. Preliminary fault lines to test

These are **hypotheses, not findings**.

| ID | Hypothesis | Test | Potential impact |
|---|---|---|---|
| H1 | AMMZ's island configuration may have two endpoints and a \(2S_0\) price, while the spec adopts one endpoint | Render AMMZ §3.1–3.3; parse (10),(13),(15); count endpoints on the actual \(I_t\); check Decision 2's condition | R4 normalization, Page time, branch structure; possible BLOCKER |
| H2 | The hierarchy contains an internal tension: \(\phi_0\gg\phi\) versus \(S_0\ll S_{\text{Bek}}\), given \(S_{\text{Bek}}=(\phi_0+\phi_{\text{hor}})/4G_N\) | Determine whether \(\phi\) means fluctuation or background/horizon value; compare baseline §7's \(\phi_0\ll\phi_{\text{hor}}\)-scale claim | Regime validity and island domination; possible BLOCKER/MATERIAL |
| H3 | The no-island branch uses a fixed-\(\beta\) log that grows forever, while AMMZ(14) is an integral over decaying \(T(t)\) that saturates at \(2S_{\text{Bek}}\) | Derive both; compare at \(t_P\); determine domain of validity of the fixed-\(\beta\) form | R4 central result; possible BLOCKER |
| H4 | Time-coordinate collision: AEMM uses boundary time \(u\) and Poincaré time \(t=f(u)\), while R4 formulas use \(t\) and \(T(t)\) while evaporation uses \(u\) | Reconstruct the coordinate map; check every occurrence of \(u\), \(t\), \(f(u)\), \(T(t)\), \(T_H\) | R1/R4 interoperability; possible BLOCKER/MATERIAL |
| H5 | The Page-time self-check \(t_P k\sim O(1)\) may fail if \(\phi_0\) is not small compared with \(\phi_{\text{hor}}\) | Derive \(t_P k\) from \(t_P\sim 6\beta S_{\text{Bek}}/\pi c\) and \(k=(c/12)(4G_N/\bar\phi_r)\); use the hierarchy | Classification of \(t_P\), footprint F1; possible BLOCKER/MATERIAL |
| H6 | The \(\kappa=k\) identity may be a form match rather than an exact identity; AMMZ's \(\kappa\) and AEMM's \(k\) may differ in definition or units | Reconstruct AEMM(20),(25) and AMMZ(13); compare factors, dimensions, and time coordinates | Interoperability; possible MATERIAL/BLOCKER |
| H7 | Essential symbols appear without explicit declaration in the spec: \(S_{\text{Bek}}, S_0, \beta, \epsilon, T_H, T_1, E_1, E_0, E_S, k, \kappa, \phi_{\text{hor}}, \bar\phi_r, y_e, a, t_P\) | Build a symbol table; classify each as standard notation, defined in content, or undeclared; apply PSEUDO §2 | Conformance; possible MATERIAL |
| H8 | Some obligations are not attached or traceable: OB-1 and OB-6 are absent from the spec; `?OB-8followup` is not in the register; OB-2/OB-3 use ambiguous "drop" language | Map every register obligation to a `?` line or a prose resolution; apply PSEUDO §6 | Conformance and traceability; possible MATERIAL |
| H9 | The `[SR]/[D]/[C]/[P]/[A]/[U]` tags may be author-typed status labels prohibited by PSEUDO §7 | Check whether they function as status claims "wherever they appear"; distinguish prose attribution from status taxonomy | Conformance; possible MATERIAL/MINOR |
| H10 | Interoperability links may be circular or restatements: `link_state_stress` uses `shock_stress` to conclude it is the stress-energy of `quench_state`; `link_sgen_qes` uses `qes_position` to conclude it extremizes \(S_{\text{gen}}\); `link_branches_rad` restates `page_curve`; `link_action_eqs` omits `stressT` | Audit each `derive` premise list against its actual chain; test for conclusion-as-premise | Interoperability and conformance; possible MATERIAL/BLOCKER |
| H11 | Source locators may be mismatched: `given shock_stress` cites AEMM(21)/(61) while baseline cites (42),(44)–(46),(61); `noisland_branch` cites AMMZ(14) for a log form while (14) is an integral; `island_value` cites AMMZ(13) for a one-endpoint value | Resolve every locator on rendered source pages; compare equation content, not just number | Source fidelity and epistemic integrity; possible MATERIAL |
| H12 | Imported results may be presented as derived: `island_value` mixes [SR] with a derived rate identity; `interop_complete` is conditional on four open carriers; `link_flux_mass` uses `global_balance` as a premise | Classify each link as import, specialization, derivation, or conditional; check status language | Epistemic integrity and R4/interop verdicts; possible MATERIAL/BLOCKER |

The hierarchy and Page-time hypotheses are linked: if \(\phi_0\ll\phi_{\text{hor}}\), then \(t_P k\sim 1+\phi_0/\phi_{\text{hor}}\sim O(1)\); if \(\phi_0\gg\phi_{\text{hor}}\), the slow-evaporation self-check fails. Test explicitly.

---

## 9. Quality gates and stop conditions

- Every finding has exact file/line, independent evidence, primary-source locator, consequence, minimum correction, and affected area.
- Every central equation is either independently derived, source-verified on a rendered page, or explicitly marked "not reconstructed."
- No imported result is promoted to derived.
- Every source locator is checked against the pinned PDF, not inherited from the package.
- The dependency graph is checked for cycles and hidden premises; conclusion-as-premise patterns are tested.
- Conformance is judged against `PSEUDO.md`, not against Muse's checklist.
- The overall verdict is logically consistent with the finding severities.
- Pre-existing package files and sources hash identically at start and end.
- No files were written under `/home/chaschel/Documents/go`.
- The report is self-contained enough for another reviewer to reproduce the conclusions.

If a BLOCKER emerges, do not stop early. Finish the relevant audit, determine the full blast radius, and state exactly what the minimum correction is. Stop after the report; no endless review cycles.

---

## 10. Defaults

- **Final report location:** `research/experiments/jt-island-001/reviews/opencode-review-v1.0.md` (exact prompt10 path).
- **Scratch:** `research/experiments/jt-island-001-review/work/`.
- **Source reading:** `pdftotext -layout` for navigation, `pdftoppm` page renders for displayed equations; HTML versions only as secondary indexes.
- **Calculation:** manual algebra plus ephemeral Python stdlib checks where useful; no sympy/numpy dependency, no deliverable code.
- **Review pass:** one full pass only; a targeted revision pass is a separate director-authorized action.
- **Judgment:** no subagent delegation for physics judgments; any mechanical extraction help is limited to locator/hash verification.
