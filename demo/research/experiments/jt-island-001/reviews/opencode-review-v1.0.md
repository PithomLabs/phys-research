# OpenCode Independent Review v1.0 — JT-ISLAND-001

**Role:** Independent adversarial physics and mathematics reviewer  
**Primary target:** R4 — Page-curve reconstruction  
**Cross-validation target:** R1 — Semiclassical matter–geometry consistency  
**Specification under review:** `phys-pseudo-r1r4-v0.1.md`  
**Package root:** `research/experiments/jt-island-001/`  
**Governing authority:** Phys Research Workflow Protocol v1.0; director-approved rulings; `plans/prompt10.md` and the approved review plan  
**Review mode:** Independent reconstruction and attempted falsification. Review only; no implementation, no specification freeze, no revision cycle.  
**Report path:** `research/experiments/jt-island-001/reviews/opencode-review-v1.0.md`

---

## Document control

- This review was performed as one full adversarial pass. No pre-existing package file, source PDF, or source HTML was modified. All package and source hashes at the end of the review match the start manifest in `research/experiments/jt-island-001-review/work/package-hashes-start.txt`.
- The review wrote only to `research/experiments/jt-island-001-review/work/` (scratch extraction, page renders, hash manifest) and to this `reviews/` report. Nothing was written under `/home/chaschel/Documents/go`.
- The authority order used was: (1) Phys Research Workflow Protocol v1.0; (2) director-approved rulings; (3) `plans/prompt10.md` and the approved plan; (4) `PSEUDO.md` and `MRC_REVIEW_PROTOCOL.md`; (5) pinned primary literature; (6) Muse's artifacts; (7) the dossier and the prior audit note.
- Primary literature takes precedence over Muse and over the dossier. Where the package and a primary source disagree, the primary source controls.
- This report does not propose architectural expansion, implementation contracts, Go code, or novelty claims. It identifies defects and the minimum corrections needed to remove them.

---

## 1. Executive verdict

### 1.1 Overall outcome

**BLOCKED.**

The R1 equation set is, in the main, a faithful and mathematically correct reproduction of the AEMM JT action and field equations, the AP stress-tensor forms, the Ward identity and chiral-flux structure, the Schwarzian/ADM dynamics, and the evaporation law. Those results survive independent reconstruction. The failure is not the R1 algebra. The failure is the physical model and the R4 reconstruction built on top of it.

Three independent blockers prevent reliable approval:

1. **F-01 — the hierarchy is internally contradictory.** The package imposes both AEMM's \(\phi_0\gg\phi\) and \(S_0\ll S_{\rm Bek}\), with \(S_0=\phi_0/4G_N\) and \(S_{\rm Bek}=(\phi_0+\phi_{\rm hor})/4G_N\). The second condition is \(\phi_0\ll\phi_{\rm hor}\), the first is \(\phi_0\gg\phi_{\rm hor}\). There is no overlap. The package's own baseline states both conditions on adjacent lines (lines 223–224).
2. **F-04 — the endpoint count and the \(2S_0\) price are inconsistent.** The package adopts a one-endpoint island \(I_t=(-\infty,a(t)]\) with one area term \((\phi_0+\phi(a))/4G_N\), but its large-\(S_0\) branch uses AMMZ's two-endpoint \(2S_0\) island-creation price. One endpoint prices \(S_0\); two endpoints price \(2S_0\). The package cannot use both.
3. **F-09 — several interoperability links are circular or restatements.** `link_sgen_qes` uses `qes_position` to conclude that `qes_position` extremizes \(S_{\rm gen}\); `link_branches_rad` uses `page_curve` to conclude `page_curve`; `link_flux_mass` uses `global_balance` to conclude `global_balance`; `link_matter_sgen` restates the definition of `sgen_functional`; `link_dilaton_area` restates `dilaton_term` plus the endpoint postulate. The claimed ten-link interoperability is therefore not established.

The R4 Page curve is a conditional/structural expectation, not a derived result. The Page time \(6\beta S_i/\pi c=1/\kappa\) is a leading estimate; the exact intersections with the source branches are \(0.703/\kappa\) (fixed-\(\beta\) log branch) or \(0.811/\kappa\) (AMMZ(14)/(13)), and the package's F1 self-check fails catastrophically under the AEMM hierarchy. The spec is not Phys-Pseudo conforming, and the summary overclaims its epistemic status. The package is repairable, but not approvable as it stands.

### 1.2 Separate verdicts

| Area | Verdict | Basis |
|---|---|---|
| R1 mathematical validity | **CONDITIONAL** | Action, dilaton equation, metric equations, signs, \(16\pi/8\pi\) pattern, stress forms, Ward identity, chiral flux, ADM sign, \(E_0>0\), energy balance and evaporation law independently verified. The regime is not the one claimed (F-01, F-17); the time coordinate is conflated (F-05); the shock-vs-pulse state map is open (F-18). |
| R4 mathematical validity | **BLOCKED** | The no-island branch is not AMMZ(14) (F-02); the Page-time estimate fails the F1 self-check and differs from the exact intersections (F-03); the endpoint count/\(2S_0\) price is inconsistent (F-04); the Page curve is qualitative/conditional (F-16); QES/RT are imports (F-11). |
| R1–R4 interoperability | **BLOCKED** | The model hierarchy is incompatible (F-01); time coordinates and cutoffs are conflated (F-05, F-10); the state is a \(\delta\)-shock vs a finite-width pulse (F-18); several Part C links are restatements or circular (F-09); \(\kappa=k\) is a form match conditional on an unstated time identification. |
| Phys-Pseudo conformance | **BLOCKED / NON-CONFORMING (repairable)** | Status-tag taxonomy violates PSEUDO §7 (F-08); declaration-before-use failures (F-07); missing/untraceable `?` items (F-15); circular/restatement `derive` links (F-09); false header claim "per PSEUDO.md." |
| Source fidelity and epistemic integrity | **BLOCKED (repairable)** | Equation-level source fidelity is generally good; metadata/status claims are not (F-08, F-12, F-14, F-16, F-20, F-21). The "single-model, traceable" claim is contradicted by F-01. |
| Overall JT-ISLAND-001 readiness | **BLOCKED** | At least one unresolved BLOCKER prevents reliable approval of the affected specification and central conclusion. |

### 1.3 Finding severity summary

| Severity | Findings | Count |
|---|---|---|
| BLOCKER | F-01, F-04, F-09 | 3 |
| MATERIAL | F-02, F-03, F-05, F-06, F-07, F-08, F-10, F-11, F-13, F-14, F-15, F-16, F-17, F-18 | 14 |
| MINOR | F-12, F-19, F-20, F-21, F-22, F-23, F-24 | 7 |

The finding index is in Appendix G.

### 1.4 What survives and what does not

**Survives independent reconstruction:** the AEMM action (2.1)–(2.3); the dilaton equation \(R+2=0\); the metric equations (2.6a–c) including the minus sign on the \(g_{\mu\nu}\phi\) term and the \(16\pi\) versus \(8\pi\) pattern; the AP stress forms (5.1)–(5.2); the trace anomaly and its absorption into \(\phi_0\); the Ward identity and the light-cone conservation identity; the chiral-flux conclusion; the Schwarzian action and ADM energy with the correct minus sign; \(E_0=\pi\bar\phi_rT_0^2/4G_N>0\); the energy balance and evaporation law \(E(u)=E_1e^{-ku}\), \(k=(c/12)(4G_N/\bar\phi_r)\); the no-island slope \(\pi cT/6\) and additive constant \((c/6)\ln(\beta/2\pi\epsilon)\); the structural branch intersection; and the formal statement that AMMZ's formalism can be directly applied to the AEMM case.

**Does not survive:** the claim that the combined model is one coherent regime; the use of the fixed-\(\beta\) log as AMMZ's growth branch; the \(1/\kappa\) Page time as a derived value; the one-endpoint normalization together with the \(2S_0\) price; the ten-link interoperability as established; the "DERIVED" status of the Page curve; the "single-model, traceable" summary; and the claim of PSEUDO conformance.

---

## 2. Review methodology and source versions

### 2.1 Scope and procedure

The review followed the approved plan in `plans/jt-island-001-review-plan.md` and the eighteen-item checklist in `opencode-checklist.md`, but did not treat the checklist as an authority. The work was performed in one full pass:

1. **Provenance and hashes.** Recorded SHA-256 for every pre-existing package file and pinned source at the start (`work/package-hashes-start.txt`); confirmed the source PDFs match `sources/README.md`; re-checked all hashes at the end. No pre-existing file changed.
2. **Source extraction.** Extracted all four PDFs with `pdftotext -layout`; recorded `pdfinfo` for page counts and versions; rendered the key AMMZ and AP pages to PNG for displayed equations. Used the arXiv HTML versions only as secondary indexes for exact equation content, never as a substitute for the pinned PDF where a displayed equation was load-bearing.
3. **Independent reconstruction.** Re-derived the R1 action variation, the metric equations, the stress/anomaly/Ward/chiral structure, the Schwarzian/ADM sign, the evaporation law, the no-island log and slope, the island branch, the exact branch intersections, the F1 self-check, the endpoint-count comparison, the time-map and cutoff-frame counterexamples, and the \(\kappa=k\) form match. The algebra is recorded in §3.
4. **Locator verification.** Checked every load-bearing locator against the rendered source: AMMZ(1)–(15), AEMM printed (2.1)–(5.10) and the package's group-ordinal convention, AP(5.1)–(5.2), Penington Eq. (2). The AEMM group-ordinal mapping is in Appendix B.
5. **Register audit.** Audited C-R1-01–C-R1-04, C-R4-01–C-R4-05, C-INT-01, ASM-01–ASM-06, BR-01–BR-10, and OB-1–OB-10 against the spec, the baseline, the registers, and the primary sources.
6. **Reality checks and conformance.** Applied RC1–RC3 independently; judged the spec against `PSEUDO.md`, not against Muse's checklist.
7. **Falsification attempts.** Tested the hierarchy, the endpoint count, the Page-time algebra, the time-coordinate map, the cutoff frame, the state map, the limiting cases, and the dependency graph. Concrete counterexamples are in §11.
8. **Report and QA.** Every finding has the nine required fields plus a disclosure field; the overall verdict is checked against the severity logic; the package hashes are re-checked; the report is self-contained.

### 2.2 Source versions

| Source | Version / identifier | Pages | SHA-256 (PDF) | Role |
|---|---|---|---|---|
| AMMZ | arXiv:1908.10996v2 [hep-th], 4 Nov 2019 | 22 | `f68be764aa6e21b79d3095401a9434c8e02047d5e6bc327a3c7b370afca9a905` | R4 island construction, Page curve, holographic matter |
| AEMM | arXiv:1905.08762v3 [hep-th], 4 Nov 2019 | 55 | `f51a59cbf6d660471965c6eca0920b70eb598652af366e65b60050edeabc3575` | R1 JT action, field equations, shock, evaporation, QES |
| AP | arXiv:1402.6334v3 [hep-th], 12 Aug 2015 | 23 | `37e78a45c5e573098b9c1457597e04cd36c336eabed8e4d389923884ec1aabb4` | Anomaly and stress-tensor forms; \(16\pi/8\pi\) tiebreak |
| Penington | arXiv:1905.08255v3 [hep-th], 29 Aug 2020 | 75 | `fa7cf02db1814f4b55e72f33f2e817c5001645b5c4f1b4ea77444b38ee03738c` | QES-origin chain; Eq. (2); Page-time definition |

The AMMZ HTML, AEMM HTML, and AP HTML hashes are also recorded in the start manifest and match. The package files (spec, baseline, branch reconstruction, registers, reality checks, audit corrections, dossier audit note, summary, index, checklist) are byte-identical to the start manifest; the source PDFs are byte-identical to `sources/README.md`.

### 2.3 Locator conventions

- **AMMZ\((n)\)** means the printed equation number \((n)\), \(n=1,\dots,15\). AMMZ uses a true single sequence; the mapping is direct.
- **AEMM\((n)\)** in the package means the ordinal of the \(n\)-th numbered equation group in arXiv:1905.08762v3, counting subequations as one group. The printed PDF uses sectional numbers \((2.1)\)–\((5.10)\). This convention is undocumented in the package (F-12); the mapping for the load-bearing equations is in Appendix B.
- **AP\((5.n)\)** is the printed sectional number in arXiv:1402.6334v3; the load-bearing equations are (5.1) and (5.2) on PDF p.15.
- **Penington Eq. (2)** is on PDF p.5; the Page-time definition \(S_{\rm rad}=A_{\rm hor}/4G_N\) is on PDF p.6.
- Page numbers are PDF pages unless marked "printed p." The AMMZ §3.3 passage and the \(2S_0\) price are on PDF p.14; the AMMZ(15) island rule is on PDF p.15; the AMMZ near-extremal condition is on PDF p.4. The AEMM near-extremal restriction \(\phi_0\gg\phi\) is in the text after (2.3) on PDF p.8 (printed p.7); AEMM(2.1)–(2.3) are on PDF p.8; (2.6a–c) on PDF p.9; (2.15)–(2.20) on PDF pp.11–12; (4.40) on PDF p.35 (printed p.34); (4.44) on PDF p.37 (printed p.36).

---

## 3. Independent mathematical reconstructions

This section records the load-bearing reconstructions. It is deliberately independent of Muse's algebra: the equations below were re-derived from the action and the primary-source forms, not copied from the package. Where the package's result is correct, that is stated; where it is not, the discrepancy is stated and is carried into §4.

### 3.1 R1.1 — Action, dilaton equation, and metric equations

The adopted action is

\[
I=I_0+I_G+I_M,
\]
\[
I_0=\frac{\phi_0}{16\pi G_N}\left[\int_M d^2x\sqrt{-g}\,R+2\int_{\partial M}K\right],
\]
\[
I_G=\frac{1}{16\pi G_N}\left[\int_M d^2x\sqrt{-g}\,\phi(R+2)+2\int_{\partial M}\phi_bK\right],
\qquad
I_M=I_{\rm CFT}[g].
\]

This matches AEMM printed (2.1)–(2.3) exactly (PDF p.8; package AEMM(6)–(8)) and the spec line 11. The normalization is not the dossier's \(4\times\) form; the topological coefficient is \(\phi_0/16\pi G_N\), not \(\phi_0/4\pi G_N\). The boundary term is \(K\), not \(K-1\). The package's baseline §5.1 records this correctly.

Varying with respect to the dilaton gives

\[
\frac{1}{16\pi G_N}(R+2)=0
\quad\Longrightarrow\quad
R=-2,
\]

which is AEMM printed (2.4) (package AEMM(9)). The bulk is locally AdS\(_2\). Verified clean.

Varying with respect to \(g^{\mu\nu}\), using AEMM's stress-tensor convention \(T_{\mu\nu}=-(2/\sqrt{-g})\delta I_M/\delta g^{\mu\nu}\), gives the covariant equation

\[
\frac{1}{8\pi G_N}\left(g_{\mu\nu}\nabla^2\phi-\nabla_\mu\nabla_\nu\phi-g_{\mu\nu}\phi\right)
=\langle T_{\mu\nu}\rangle .
\]

The minus sign on \(g_{\mu\nu}\phi\) and the \(16\pi\) versus \(8\pi\) pattern are correct. In Poincaré conformal coordinates \(ds^2=-4dx^+dx^-/(x^+-x^-)^2\), the components are

\[
2\partial_+\partial_-\phi+\frac{4}{(x^+-x^-)^2}\phi
=16\pi G_N\,T_{+-},
\tag{3.1a}
\]
\[
-\frac{1}{(x^+-x^-)^2}\partial_+\left((x^+-x^-)^2\partial_+\phi\right)
=8\pi G_N\,T_{++},
\tag{3.1b}
\]
\[
-\frac{1}{(x^+-x^-)^2}\partial_-\left((x^+-x^-)^2\partial_-\phi\right)
=8\pi G_N\,T_{--}.
\tag{3.1c}
\]

These match AEMM printed (2.6a)–(2.6c) (PDF p.9) and the spec line 28.

**Why the \(16\pi/8\pi\) pattern is not an error.** In two dimensions the \(+-\)-component is the trace part of the metric equation; the \(++\) and \(--\) components are the transverse constraints. The variation of \(\sqrt{-g}\phi(R+2)\) produces the \(g_{\mu\nu}\nabla^2\phi-\nabla_\mu\nabla_\nu\phi\) tensor, whose \(+-\)-component is twice the \(++\)-component after the \(g_{\mu\nu}\) trace contraction; the \(2\phi\) term contributes only to the \(+-\)-component because \(g_{++}=g_{--}=0\). The factor two between (3.1a) and (3.1b,c) is therefore structural, not a typo. The package's baseline §5.3 and spec line 28 record this correctly. The dossier's \(8\pi\) in the \(+-\)-equation and its wrong sign on \(g\phi\) are rejected.

**Independent check with the black-hole solution.** The dilaton profile

\[
\phi=2\bar\phi_r\,\frac{1-(\pi T_0)^2x^+x^-}{x^+-x^-}
\]

solves (3.1a) with \(T_{+-}=0\):
\[
2\partial_+\partial_-\phi+\frac{4}{(x^+-x^-)^2}\phi=0.
\]
This confirms the relative sign and coefficient of the \(\phi\) term. A wrong sign would not admit the AdS\(_2\) black-hole solution. Verified.

### 3.2 R1.2 — Stress tensor, trace anomaly, Ward identity, and chiral flux

The AP forms are, with \(N\leftrightarrow c\) and \(\omega\leftrightarrow\rho\),

\[
\langle T_{+-}\rangle=-\frac{c}{12\pi}\partial_+\partial_-\rho,
\qquad
\langle T_{\pm\pm}\rangle=\frac{c}{12\pi}\left(\partial_\pm^2\rho-(\partial_\pm\rho)^2\right)
+\langle :T_{\pm\pm}:\rangle_{\rm state}.
\]

This matches AP printed (5.1)–(5.2) (PDF p.15) and the spec line 25. The trace anomaly is

\[
\langle T^\mu_{\ \mu}\rangle=\frac{c}{24\pi}R,
\]
which on shell \(R=-2\) gives \(\langle T^\mu_{\ \mu}\rangle=-c/12\pi\). The anomaly is absorbed into the extremal dilaton,
\[
\phi_0\to\phi_0^{\rm Ren}=\phi_0+\frac{cG_N}{3},
\]
as in AEMM footnote 2 (PDF p.11; package AEMM(42) region). The package's `trace_R_form` claim (spec line 23) is a separate, suspect \(R\)-form equality; it is correctly marked `[U]` and is not used downstream. The AP route does not need it.

**Ward identity.** Diffeomorphism invariance of the matter action gives
\[
\nabla^\mu\langle T_{\mu\nu}\rangle=0.
\]
This is independent of the trace anomaly: the anomaly fixes the trace, the Ward identity fixes the divergence. Removing the anomaly changes the solution (it removes the Hawking-flux terms), not the identity. This is the corrected RC2 statement (audit corrections Decision 3; reality-checks lines 35–38). Verified against AP §5.1, where the anomaly supplies \(T_{+-}\) and conservation is imposed independently to obtain \(T_{\pm\pm}\).

**Light-cone conservation and chiral flux.** In conformal gauge \(ds^2=-e^{2\rho}dy^+dy^-\), the nonzero connections are
\[
\Gamma^+_{++}=2\partial_+\rho,\qquad
\Gamma^-_{--}=2\partial_-\rho,
\]
and the mixed connections vanish in null coordinates. The conservation equation \(\nabla^\mu T_{\mu+}=0\) becomes
\[
\partial_-\langle T_{++}\rangle+\partial_+\langle T_{+-}\rangle
-2\langle T_{+-}\rangle\partial_+\rho=0,
\tag{3.2}
\]
with the \(+ \leftrightarrow -\) twin. Substitute
\[
\langle T_{+-}\rangle=-\frac{c}{12\pi}\partial_+\partial_-\rho,
\qquad
\langle T_{++}\rangle=\frac{c}{12\pi}\left(\partial_+^2\rho-(\partial_+\rho)^2\right)
+\langle :T_{++}:\rangle .
\]
The \(\partial_+^2\partial_-\rho\) terms cancel, and the \(\partial_+\rho\,\partial_+\partial_-\rho\) terms cancel:
\[
-2\frac{c}{12\pi}+\frac{c}{6\pi}=0.
\]
The result is
\[
\partial_-\langle :T_{++}(y^+):\rangle=0,
\qquad
\partial_+\langle :T_{--}(y^-):\rangle=0.
\tag{3.3}
\]
This independently verifies the package's `chiral_flux` (spec line 31) and the baseline §5.4. The anomaly and the Ward identity are joint, independent inputs; the geometric pieces cancel, leaving the state piece chiral. The package's structure is correct.

### 3.3 R1.3 — Schwarzian action, ADM energy, and evaporation

The boundary reduction gives
\[
S_G=\frac{\bar\phi_r}{8\pi G_N}\int du\,\{f,u\},
\qquad
E(u)=-\frac{\bar\phi_r}{8\pi G_N}\{f,u\},
\]
matching AEMM printed (2.8)–(2.9) (package AEMM(13)–(14)). The minus sign is load-bearing. For the black-hole solution
\[
f(u)=\frac{1}{\pi T_0}\tanh(\pi T_0 u),
\]
the Schwarzian is
\[
\{f,u\}=-2(\pi T_0)^2.
\]
Therefore
\[
E_0=-\frac{\bar\phi_r}{8\pi G_N}\left(-2\pi^2T_0^2\right)
=\frac{\pi\bar\phi_rT_0^2}{4G_N}>0,
\]
matching AEMM printed (2.15) (package AEMM(20)). The package's ADM sign and positive-mass check are correct.

The energy balance is
\[
\partial_uE(u)=f'(u)^2\left(T_{--}-T_{++}\right),
\]
matching AEMM printed (2.17) (package AEMM(22)). The shock stress is
\[
\langle T_{--}(x^-)\rangle=E_S\delta(x^-)
-\frac{c}{24\pi}\{y^-,x^-\}\Theta(x^-),
\]
matching AEMM printed (2.16) (package AEMM(21)). The evaporation law is
\[
\{f,u\}\propto e^{-ku},
\qquad
E(u)=E_1e^{-ku},
\qquad
E_1=E_0+E_S,
\qquad
k=\frac{c}{12}\frac{4G_N}{\bar\phi_r}\ll1,
\]
matching AEMM printed (2.18)–(2.20), (2.24)–(2.25) (package AEMM(23)–(25), (24)). Verified.

**Independent closure of OB-4.** The standard 2D flux form is
\[
\frac{dE}{du}=-kE=-\frac{\pi c}{12}T^2.
\]
Using \(E=\pi\bar\phi_rT^2/(4G_N)\) and \(k=cG_N/(3\bar\phi_r)\),
\[
\frac{dE}{du}=\frac{\pi\bar\phi_r}{2G_N}T\dot T
=-\frac{\pi c}{12}T^2
\quad\Longrightarrow\quad
\dot T=-\frac{k}{2}T,
\]
so \(T(u)=T_1e^{-ku/2}\). This is consistent with the evaporation law and with the package's \(T(t)=T_1e^{-kt/2}\) once the time coordinate is identified. The formula in `global_balance` is therefore correct; the package has not exhibited the line-by-line derivation, but the obligation is closable. (The general form with incoming flux is \(dE/du=-(\pi c/12)T^2+F_{\rm in}\).)

### 3.4 R1.4 — Interface, cutoff, and shock

The boundary condition is
\[
g_{uu}\big|_{\rm bdy}=\frac{1}{\epsilon^2}
=\frac{-t'^2+z'^2}{z^2},
\qquad
\phi=\phi_b=\frac{\bar\phi_r}{\epsilon},
\]
matching AEMM printed (2.12) (PDF p.9). The physical cutoff is
\[
z=\epsilon f'(u),
\]
not the constant \(z=\epsilon\). The \(f'\) factor is load-bearing in the Weyl transformation and in the stress-tensor transformation. The package's baseline §2 records this correctly and rejects the dossier's constant cutoff.

The quench state is the Hartle–Hawking/eternal state plus a right-boundary coupling at \(t=0\), with the \(E_S\) shock transient. The shock is instantaneous in AEMM; AMMZ's §3.2.2 describes a finite-width coupling over \(\Delta t\) with \(E\sim c/\Delta t\). The two are not identical states; the \(\delta\)-shock is the \(\Delta t\to0\) limit. The package records the shock and the AMMZ pulse as "same transient physics, holographic dual language" (baseline lines 144–145), but does not derive the map. This is F-18.

### 3.5 R4.1 — Generalized entropy and endpoint counting

AMMZ printed (10) is
\[
S_{\rm gen}(y)=\frac{\phi(y)}{4G_N}+S_{\rm Bulk\text{-}2d}[I_y],
\]
and AMMZ printed (15) is the island rule
\[
S(A)=\operatorname*{Min}_{I_g}\operatorname*{Ext}_{I_g}
\left[S_{\rm eff}(A\cup I_g)+\frac{\operatorname{Area}[\partial I_g]}{4G_N}\right].
\]

In JT, the total dilaton is \(\phi_0+\phi\), so the package's area term
\[
S_{\rm grav}(a)=\frac{\phi_0+\phi(a)}{4G_N}
\]
is the correct one-endpoint contribution. For a configuration with endpoints \(a_i\),
\[
S_{\rm grav}=\sum_i\frac{\phi_0+\phi(a_i)}{4G_N}.
\]
The extremal part is \(S_0=\phi_0/4G_N\) per endpoint. One endpoint prices \(S_0\); two endpoints price \(2S_0\). The dynamical part is \(S_{\rm Bek}-S_0=\phi_{\rm hor}/4G_N\).

AMMZ §3.3 (PDF p.14) explicitly states that the island has **two endpoints** and that the island-creation price is **\(2S_0\)**:
> "we need to pay a \(2S_0\) price to create the 'island' (because of the two endpoints of the island)"

The package adopts a **one-endpoint** island \(I_t=(-\infty,a(t)]\) with \(\partial I=\{a\}\) (spec lines 52–54, 57). It then uses the \(2S_0\) price in `island_disfavored` (spec line 74). This is the F-04 inconsistency. The one-endpoint normalization and the two-endpoint price cannot both be correct.

The package's `dilaton_term` cites AMMZ(10) for "per QES endpoint, summed over endpoints" (spec line 54). AMMZ(10) is a single dilaton term; the multi-endpoint area is AMMZ(15). This is the F-20 citation mismatch.

### 3.6 R4.2 — No-island branch and the AMMZ growth branch

The fixed-\(\beta\) 2D CFT single-interval entropy is
\[
S_{\rm no}(t)=\frac{c}{6}\ln\left[\frac{\beta}{\pi\epsilon}\sinh\left(\frac{\pi t}{\beta}\right)\right]+\log g,
\]
where \(\log g\) is the boundary entropy. For the Cardy brane, \(\log g=0\); the package omits it without stating the condition (F-19). At late times,
\[
S_{\rm no}(t)\to\frac{\pi cT}{6}t+\frac{c}{6}\ln\frac{\beta}{2\pi\epsilon}.
\tag{3.4}
\]
The slope is \(\pi cT/6\), and the additive constant is UV-sensitive. This is correct as a fixed-temperature CFT formula.

AMMZ printed (14) is a different function:
\[
S_{\rm Rad}(t)=\frac{\pi c}{6}\int_0^t dt'\,T(t')
=2S_{\rm Bek}^i\left(1-e^{-\frac{\kappa}{2}t}\right).
\tag{3.5}
\]
It is the early-time growth branch, computed with the decaying temperature \(T(t)=T_ie^{-\kappa t/2}\). It saturates at \(2S_{\rm Bek}^i\). AEMM printed (4.44) has the same saturating form:
\[
S_{\rm bath}\sim\frac{\bar\phi}{4G_N}4\pi T_1\left(1-e^{-\frac{k}{2}u}\right).
\tag{3.6}
\]
The package's `noisland_branch` (spec line 58) uses (3.4), not (3.5), but cites AMMZ(14) as a cross-check. The slopes agree at early times; the functions do not. The log branch grows without bound; the AMMZ branch saturates. This is F-02.

### 3.7 R4.3 — Island branch and QES

AMMZ printed (12) gives the QES position
\[
y_e^+=t-\frac{1}{2\pi T(t)}
\log\frac{S_{\rm Bek}(T(t))-S_0}{c}+\cdots,
\tag{3.7}
\]
which is a source-reported import (AMMZ cites [8,9]). The JT-specific derivation is in AEMM §4. AEMM printed (4.40) gives the Hayden–Preskill/scrambling time
\[
t_{HP}=\frac{\beta}{2\pi}\log\left[\frac{16}{c}\left(S-S_0\right)\right],
\tag{3.8}
\]
with an explicit factor \(16\). The package's `qes_position` uses (3.7) and cites AEMM §4; the factor \(16\) is not reconciled. This is F-21 (MINOR because the leading Page time is unaffected, but the import is not fully specified).

The island branch value is
\[
S_{\rm isl}(t)\approx S_{\rm BH}(T(t))+\text{logs},
\qquad
S_{\rm BH}(T)=\frac{\phi_0+\phi_{\rm hor}(T)}{4G_N}.
\tag{3.9}
\]
For JT,
\[
\phi_{\rm hor}(T)=2\pi\bar\phi_rT,
\qquad
S_{\rm BH}(T)=S_0+\frac{\pi\bar\phi_rT}{2G_N}.
\tag{3.10}
\]
The package's `island_value` (spec line 63) and `bh_entropy_fn` (spec line 55) have the correct form. The \(O(c\log)\) matter and log-\(S\) corrections are subleading under the hierarchy, but their coefficient and order should be stated; the package does.

### 3.8 R4.4 — Branch intersection and Page time

Let \(x=e^{-\kappa t/2}\), and write \(S_i=S_{\rm BH}(T_i)=S_0+\pi\bar\phi_rT_i/(2G_N)\). The island branch is
\[
S_{\rm isl}(t)=S_0+(S_i-S_0)x.
\tag{3.11}
\]

**Case A: exact AMMZ(14) growth branch.** Set (3.5) equal to (3.11):
\[
2S_i(1-x)=S_0+(S_i-S_0)x.
\]
Solving,
\[
x=\frac{2S_i-S_0}{3S_i-S_0},
\qquad
t_P=\frac{2}{\kappa}\ln\frac{3S_i-S_0}{2S_i-S_0}.
\tag{3.12}
\]
For \(S_0\ll S_i\),
\[
x\to\frac{2}{3},
\qquad
t_P=\frac{2}{\kappa}\ln\frac{3}{2}
=\frac{0.81093}{\kappa}.
\tag{3.13}
\]

**Case B: fixed-\(\beta\) log branch, leading late-time form.** Set (3.4) without the constant equal to (3.11):
\[
-2(S_i-S_0)\ln x=S_0+(S_i-S_0)x.
\]
For \(S_0\ll S_i\),
\[
-2\ln x=x
\quad\Longrightarrow\quad
x=2W(1/2)=0.70347,
\qquad
t_P=\frac{2W(1/2)}{\kappa}
=\frac{0.70347}{\kappa}.
\tag{3.14}
\]
Including the constant shifts \(x\) at \(O(k\log)\), as recorded in F2.

**Case C: the package's estimate.**
\[
t_P\approx\frac{6\beta S_i}{\pi c}.
\]
Using \(S_i=\pi\bar\phi_rT_i/(2G_N)\), \(\beta=1/T_i\), and \(\kappa=cG_N/(3\bar\phi_r)\),
\[
\frac{6\beta S_i}{\pi c}
=\frac{1}{\kappa}.
\tag{3.15}
\]
So the package's value is \(1/\kappa\), which is 23% above (3.13) and 42% above (3.14). The package calls it a leading estimate with footprint F1–F5; that is honest in part, but the estimate is not the exact intersection, and the spread is not reported.

**F1 self-check.** The package's branch reconstruction (lines 46–48) writes
\[
t_Pk\sim\frac{\phi_0+\phi_{\rm hor}}{2\pi T\bar\phi_r}
\sim O(1)+S_0\text{-piece}.
\]
Using \(t_P=6\beta S_i/(\pi c)\), \(S_i=(\phi_0+\phi_{\rm hor})/(4G_N)\), \(k=cG_N/(3\bar\phi_r)\), and \(\phi_{\rm hor}=2\pi\bar\phi_rT\),
\[
t_Pk=\frac{\phi_0+\phi_{\rm hor}}{2\pi T\bar\phi_r}
=1+\frac{\phi_0}{\phi_{\rm hor}}.
\tag{3.16}
\]
Under the package's island-domination hierarchy \(S_0\ll S_i\), \(\phi_0\ll\phi_{\rm hor}\), so \(t_Pk\approx1\): the evaporation during growth is an \(O(1)\) correction, not a small one. Under AEMM's \(\phi_0\gg\phi_{\rm hor}\), \(t_Pk\gg1\): the slow-evaporation approximation fails catastrophically. The F1 self-check is not a pass; it is a diagnostic that exposes F-01. This is F-03.

### 3.9 R4.5 — Extremization and minimization

The package correctly separates the two operations:
- `qes_stationarity` (spec line 60): \(\partial S_{\rm gen}/\partial a^\pm=0\) fixes the QES.
- `branch_selection` (spec line 64): \(S(R_t)=\min\{S_{\rm no},S_{\rm isl}\}\) after extremizing each candidate.
- `extremization_then_minimization` (spec line 65): the ordered operations are distinct.

This matches AMMZ §3.3 and (15). The package's `island_rule_status` (spec line 70) correctly says the island rule is an effective prescription, derived from RT/HRT only within the holographic setup and conditional on RT/HRT correctness. No unsupported minimality claim is made in the spec. This part is structurally sound.

The defect is not the ordering; it is the inputs: the one-endpoint normalization (F-04), the wrong no-island branch (F-02), and the imported QES/RT (F-11). The ordering is accepted; the branch values are not.

### 3.10 Interoperability — the \(\kappa=k\) identity

AEMM gives
\[
E(u)=\frac{\pi\bar\phi_r}{4G_N}T(u)^2,
\qquad
E(u)=E_1e^{-ku}
\quad\Longrightarrow\quad
T(u)=T_1e^{-ku/2}.
\tag{3.17}
\]
AMMZ gives
\[
T(t)\sim T_ie^{-\kappa t/2}.
\tag{3.18}
\]
The functional forms match if \(\kappa=k\) and if AMMZ's \(t\) is AEMM's \(u\). The package's branch reconstruction (lines 28–31) calls this a term-by-term derivation. It is a **form match**, not an exact identity: the two papers use different time coordinates, and the package never states the identification. If \(t_{\rm AMMZ}=u_{\rm AEMM}\), the identity is correct; if \(t_{\rm AMMZ}=t_{\rm AEMM}=f(u)\), the decay is \(T(f(u))\), not \(T(u)\), and the Page-time integral is in a different variable. This is F-05.

The identity also inherits the near-extremal regime. AEMM's \(k\ll1\) and AMMZ's \(\kappa\propto c\) are compatible in the overlap where the dynamical entropy exceeds \(c\), but the package's \(S_0\ll S_{\rm Bek}\) pushes the model to the opposite hierarchy. So even the form match is conditional on F-01.

### 3.11 Independent verification ledger

| Result | Independent status | Evidence |
|---|---|---|
| JT action (2.1)–(2.3), normalization, \(K\) boundary term | **Verified** | AEMM PDF p.8; direct comparison |
| Dilaton equation \(R+2=0\) | **Verified** | AEMM (2.4); direct variation |
| Metric equations (2.6a–c), minus \(g\phi\), \(16\pi/8\pi\) | **Verified** | AEMM PDF p.9; black-hole solution check |
| AP stress forms (5.1)–(5.2) | **Verified** | AP PDF p.15; mapping \(N\leftrightarrow c\), \(\omega\leftrightarrow\rho\) |
| Trace anomaly and \(\phi_0\) shift | **Verified** | AEMM footnote 2; AP §5 |
| Ward identity and light-cone conservation | **Verified** | Direct conformal-gauge computation |
| Chiral flux \(\partial_-\langle :T_{++}:\rangle=0\) | **Verified** | Geometric cancellation shown in §3.2 |
| Schwarzian action and ADM sign | **Verified** | AEMM (2.8)–(2.9); \(\tanh\) check |
| \(E_0=\pi\bar\phi_rT_0^2/4G_N>0\) | **Verified** | AEMM (2.15); explicit Schwarzian |
| Energy balance and evaporation law | **Verified** | AEMM (2.17)–(2.25); \(dE/du=-kE\) |
| No-island slope \(\pi cT/6\) and constant | **Verified** | Late-time expansion of the log |
| AMMZ(14) saturating branch | **Verified** | AMMZ PDF p.13; AEMM (4.44) |
| Exact AMMZ branch intersection | **Derived** | §3.8, \(0.811/\kappa\) |
| Exact log-branch intersection | **Derived** | §3.8, \(0.703/\kappa\) |
| Package \(t_P=1/\kappa\) | **Leading estimate only** | §3.8, 23–42% high |
| F1 self-check \(t_Pk=1+\phi_0/\phi_{\rm hor}\) | **Derived; fails under AEMM** | §3.8 |
| One-endpoint area term | **Correct given the postulate** | AMMZ(10); Decision 2 |
| Two-endpoint \(2S_0\) price | **Correct for AMMZ's two-endpoint island** | AMMZ §3.3, PDF p.14 |
| \(\kappa=k\) | **Form match, conditional** | §3.10; time identification unstated |
| QES position AMMZ(12) | **Source-reported import** | AMMZ PDF p.9; AEMM §4 |
| RT/HRT island rule | **Source-reported import/effective prescription** | AMMZ(15), §4 |

---

## 4. Findings classified BLOCKER / MATERIAL / MINOR

Each finding uses the prompt10 nine-field format plus a disclosure field. The disclosure field records whether the package already carries the issue as an obligation; a disclosed obligation can still be MATERIAL or BLOCKER if a central claim depends on it.

---

### F-01 — BLOCKER — The hierarchy and the adopted model are mutually incompatible

**Affected file and exact line/section:** `phys-pseudo-r1r4-v0.1.md` line 12 (`given regime_hierarchy`); `working-baseline.md` lines 220–226 (§7), especially lines 223–224; `registers.md` lines 21, 27–28; `branch-reconstruction.md` lines 42–53 (F1/F5); `summary.md` lines 3–11.

**Mathematical or physical issue:** The package imposes two conditions that cannot both hold in the adopted AEMM model:
\[
\phi_0\gg\phi \quad\text{(AEMM near-extremal restriction)}
\]
and
\[
S_0\ll S_{\rm Bek} \quad\text{(island domination)},
\]
with
\[
S_0=\frac{\phi_0}{4G_N},\qquad
S_{\rm Bek}=\frac{\phi_0+\phi_{\rm hor}}{4G_N}.
\]
The second condition is equivalent to \(\phi_0\ll\phi_{\rm hor}\). The first requires \(\phi_0\gg\phi_{\rm hor}\) if \(\phi\) is the dynamical/horizon dilaton. There is no overlap. The spec line 12 asserts both in one line; the baseline asserts both on adjacent lines: line 223 says "\(\phi_0\ll\phi_{\rm hor}\)-scale" and line 224 says "\(\phi_0\gg\phi\)."

**Independent evidence or derivation:** Set \(4G_N=1\) for the counting argument. Then \(S_0=\phi_0\) and \(S_{\rm Bek}=\phi_0+\phi_{\rm hor}\).
- Choose \(\phi_0=100\), \(\phi_{\rm hor}=900\). Then \(S_0=100\), \(S_{\rm Bek}=1000\), so \(S_0\ll S_{\rm Bek}\), but \(\phi_0\ll\phi_{\rm hor}\), violating AEMM's \(\phi_0\gg\phi\).
- Choose \(\phi_0=1000\), \(\phi_{\rm hor}=10\). Then \(\phi_0\gg\phi_{\rm hor}\), but \(S_0=1000\), \(S_{\rm Bek}=1010\), so \(S_0\ll S_{\rm Bek}\) fails.
No choice satisfies both. AMMZ's near-extremal condition is not \(c\ll S_0\); it is
\[
c\ll\frac{\phi-\phi_0}{4G_N}=S_{\rm Bek}-S_0,
\]
which requires the **dynamical** entropy \(S_{\rm Bek}-S_0\) to exceed \(c\). That condition is compatible with \(\phi_0\gg\phi_{\rm hor}\) as long as \(\phi_{\rm hor}/4G_N\gg c\), but then \(S_0\approx S_{\rm Bek}\), so the island-domination condition fails. The package's hierarchy \(c\ll S_0\) is a different condition (F-17).

**Primary source and exact locator:** AEMM printed (2.1)–(2.3), PDF p.8; the text after (2.3) on PDF p.8 (printed p.7) says the dimensional reduction "implicitly imposing the restriction \(\phi_0\gg\phi\)." AMMZ §2, PDF p.4: \(c\ll(\phi-\phi_0)/(4G_N^{(2)})\). AMMZ §3.3, PDF p.14: \(S_0\ll S^i_{\rm Bek}\) and the \(2S_0\) price.

**Consequence for the research conclusion:** The adopted model is not one coherent regime. The island branch's existence, the Page transition, and the slow-evaporation approximation are not established under the AEMM hierarchy. The F1 self-check fails catastrophically: \(t_Pk=1+\phi_0/\phi_{\rm hor}\gg1\) under AEMM, so the evaporation during growth is not a small correction. This blocks the central R4 conclusion.

**Minimum correction required:** Choose one regime and re-derive R1/R4 in it:
- (a) **AEMM near-extremal:** \(\phi_0\gg\phi_{\rm hor}\). Then \(S_0\approx S_{\rm Bek}\), the island is not favored by the \(2S_0\) price, and R4 yields no Page transition in the adopted two-sided model. This is the honest consequence of AEMM's hierarchy.
- (b) **Island domination:** \(S_0\ll S_{\rm Bek}\). Then AEMM's \(\phi_0\gg\phi\) must be dropped; the topological term is no longer the AEMM near-extremal term, and the action/regime must be restated (a different JT model with a small extremal entropy, or a different dilaton normalization). This is not AEMM's model as written.
- (c) **Different model:** adopt a model in which the extremal entropy is small and the dynamical dilaton is large, and re-source the action and the near-extremal condition accordingly.

No silent compromise is acceptable. The spec must state which regime is adopted and remove the contradictory line.

**Affected area:** R1 / R4 / interoperability.

**Disclosure:** Not disclosed as a contradiction. The two conditions are separately labeled (`[A/C]` on line 12; F5 in the branch reconstruction), but the incompatibility is not carried as a `?` or an OB. The F1 footprint does not cover it.

---

### F-02 — MATERIAL — The no-island branch is not AMMZ(14); only the slope matches

**Affected file and exact line/section:** `phys-pseudo-r1r4-v0.1.md` line 58 (`claim noisland_branch`) and line 59 (`derive noisland_slope`); `branch-reconstruction.md` lines 10–22; `working-baseline.md` §6.3 lines 168–173.

**Mathematical or physical issue:** The package uses the fixed-\(\beta\) 2D CFT log
\[
S_{\rm no}(t)=\frac{c}{6}\ln\left[\frac{\beta}{\pi\epsilon}\sinh\left(\frac{\pi t}{\beta}\right)\right]
\tag{4.1}
\]
and cites AMMZ(14) as a cross-check. AMMZ(14) is not this function. It is
\[
S_{\rm Rad}(t)=\frac{\pi c}{6}\int_0^t dt'\,T(t')
=2S_{\rm Bek}^i\left(1-e^{-\frac{\kappa}{2}t}\right),
\tag{4.2}
\]
the early-time growth branch computed with the decaying temperature \(T(t)=T_ie^{-\kappa t/2}\). The two functions agree only in their early-time slope \(\pi cT/6\); the log grows without bound, while (4.2) saturates at \(2S_{\rm Bek}^i\). AEMM(4.44) is the same saturating form.

**Independent evidence or derivation:** Differentiate:
\[
\frac{dS_{\rm no}}{dt}=\frac{\pi cT}{6}\coth\left(\frac{\pi t}{\beta}\right)
\to\frac{\pi cT}{6},
\qquad
\frac{dS_{\rm Rad}}{dt}=\frac{\pi cT_i}{6}e^{-\kappa t/2}
\to0.
\]
The second derivative of the log is negative but the function never saturates; the AMMZ branch has a finite asymptote. At the Page time, the two differ by an O(1) factor in the entropy. The exact intersections are different: using (4.2) with the island branch gives \(t_P=0.811/\kappa\) (F-03); using (4.1) gives \(t_P=0.703/\kappa\). The package's \(1/\kappa\) is a third value.

**Primary source and exact locator:** AMMZ printed (14), PDF p.13; AMMZ printed (13), PDF p.9; AEMM printed (4.44), PDF p.37 (printed p.36).

**Consequence for the research conclusion:** The no-island branch used in the intersection is not the source's growth branch. The Page-time value and the branch comparison are not derived from the adopted AMMZ construction. The slope \(\pi cT/6\) survives; the branch formula does not.

**Minimum correction required:** Either (a) use AMMZ(14) as the no-island branch and derive its intersection with AMMZ(13); or (b) state explicitly that the fixed-\(\beta\) log is a separate approximation, valid only while \(t\) is below the saturation time, and quantify the domain and the error. Do not cite AMMZ(14) as an exact cross-check. Recompute \(t_P\) with the chosen branch.

**Affected area:** R4.

**Disclosure:** The package labels `noisland_branch` [SR] and `pagetime_value` [C], but the formula mismatch is not disclosed. The branch reconstruction §4 says "same slope" and "three-way agreement on the slope"; that is true but does not justify using the log as the branch.

---

### F-03 — MATERIAL — The Page-time estimate fails the F1 self-check and differs from the exact intersection

**Affected file and exact line/section:** `phys-pseudo-r1r4-v0.1.md` line 67 (`claim pagetime_value`) and line 68 (`?OB-7`); `branch-reconstruction.md` lines 42–53 (F1); `working-baseline.md` lines 223–224.

**Mathematical or physical issue:** The package's estimate is
\[
t_P\approx\frac{6\beta S_i}{\pi c}.
\tag{4.3}
\]
Using \(S_i=\pi\bar\phi_rT_i/(2G_N)\), \(\beta=1/T_i\), and \(k=cG_N/(3\bar\phi_r)\), this is exactly \(1/\kappa\). The exact intersections are \(0.703/\kappa\) (fixed-\(\beta\) log branch) and \(0.811/\kappa\) (AMMZ(14)/(13)), so (4.3) is 23–42% high. More seriously, the F1 self-check is not a pass:
\[
t_Pk=\frac{\phi_0+\phi_{\rm hor}}{2\pi T\bar\phi_r}
=1+\frac{\phi_0}{\phi_{\rm hor}}.
\tag{4.4}
\]
Under the package's island-domination hierarchy \(S_0\ll S_i\), \(\phi_0\ll\phi_{\rm hor}\), so \(t_Pk\approx1\): the evaporation during growth is an \(O(1)\) correction, not a small one. Under AEMM \(\phi_0\gg\phi_{\rm hor}\), \(t_Pk\gg1\): the slow-evaporation approximation fails. The package writes "\(\sim O(1)+S_0\)-piece" but does not state that the same hierarchy that makes it \(O(1)\) is the one that contradicts AEMM.

**Independent evidence or derivation:** The algebra is in §3.8. The exact branch equations are:
- AMMZ(14) vs (13): \(x=(2S_i-S_0)/(3S_i-S_0)\), \(t_P=(2/\kappa)\ln(3/2)=0.81093/\kappa\).
- Log branch vs (13): \(x=2W(1/2)=0.70347\), \(t_P=0.70347/\kappa\).
- Package: \(1/\kappa\).
The F1 ratio \(t_Pk=1+\phi_0/\phi_{\rm hor}\) follows from the same identities. The package's claim that F1 is an "O(1) correction, computed nowhere yet" is an admission that the value is not frozen; it is not a derivation.

**Primary source and exact locator:** AMMZ printed (13)–(14), PDF pp.9, 13; AEMM printed (2.20), (2.24)–(2.25), PDF pp.11–13; Penington Eq. (2), PDF p.5; Page-time definition, PDF p.6.

**Consequence for the research conclusion:** The Page time is a leading heuristic, not a derived value. The package's \(1/\kappa\) is not the exact intersection, and the F1 self-check fails in the AEMM regime. The R4 conclusion "the Page time is \(6\beta S/\pi c\)" is not established.

**Minimum correction required:** Report the exact intersections and the spread; keep \(t_P\) as a leading estimate with F1–F5 explicitly quantified or bounded; do not call it derived. Resolve F-01 first: the F1 self-check cannot pass in the AEMM regime. If the model is changed to the island-domination regime, recompute F1 and the exact intersection; if it stays AEMM, state that no Page transition occurs.

**Affected area:** R4.

**Disclosure:** `?OB-7` and F1–F5 are disclosed. The disclosure is honest, but it does not narrow the claim enough: the spec's `[C]` tag still presents the value as a conditional result, while the text and summary call it a derived Page time.

---

### F-04 — BLOCKER — The one-endpoint island and the \(2S_0\) price are inconsistent

**Affected file and exact line/section:** `phys-pseudo-r1r4-v0.1.md` lines 52–54 (`island_region`, `sgen_functional`, `dilaton_term`), line 57 (`assume endpoint_count_one`), line 74 (`claim island_disfavored`), line 88 (`derive link_dilaton_area`); `working-baseline.md` §6.2 lines 163–166, §6.5 line 187; `audit-corrections-v1.1.md` lines 40–47 (Decision 2).

**Mathematical or physical issue:** The package adopts a one-endpoint island
\[
I_t=(-\infty,a(t)],\qquad \partial I=\{a\},
\]
with the one-endpoint area term
\[
S_{\rm grav}(a)=\frac{\phi_0+\phi(a)}{4G_N}.
\]
It then uses AMMZ's **two-endpoint** \(2S_0\) island-creation price in the large-\(S_0\) branch. One endpoint prices \(S_0\); two endpoints price \(2S_0\). The package cannot use both. AMMZ §3.3 explicitly says the island has two endpoints and the price is \(2S_0\). If the package's one-endpoint configuration is the actual one, the price is \(S_0\), and the large-\(S_0\) threshold is different. If the two-endpoint configuration is the actual one, the area term is not the package's one-endpoint form and the island branch value changes.

**Independent evidence or derivation:** For a set of endpoints \(a_i\),
\[
S_{\rm grav}=\sum_i\frac{\phi_0+\phi(a_i)}{4G_N}.
\]
The extremal part is \(S_0\) per endpoint. One endpoint gives \(S_0\); two endpoints give \(2S_0\). The package's `dilaton_term` line 54 says "per QES endpoint, summed over endpoints," but the adopted geometry has only one endpoint. The `island_disfavored` line 74 uses the \(2S_0\) argument from AMMZ §3.3, which is derived for the two-endpoint island. The mismatch is not a convention; it changes the threshold for island domination by a factor of two and changes the area term.

**Primary source and exact locator:** AMMZ §3.3, PDF p.14: "we need to pay a \(2S_0\) price to create the 'island' (because of the two endpoints of the island)." AMMZ printed (10), PDF pp.7–8: single dilaton term. AMMZ printed (13), PDF p.9: \(S_{\rm Bek}(T)+\) logs. AMMZ printed (15), PDF p.15: \(\operatorname{Area}[\partial I_g]/4G_N\).

**Consequence for the research conclusion:** The R4 normalization and the large-\(S_0\) branch structure are not internally consistent. The Page transition is not established for the adopted configuration. This is a central conclusion, hence BLOCKER.

**Minimum correction required:** Decide the actual endpoint count from the AMMZ geometry and use the corresponding area term and \(S_0\) price consistently:
- **One endpoint:** use \((\phi_0+\phi(a))/4G_N\), replace the \(2S_0\) argument with the correct one-endpoint \(S_0\) argument, and re-derive the large-\(S_0\) branch.
- **Two endpoints:** restore the two-endpoint sum \(2(\phi_0+\phi_{\rm hor})/4G_N\), re-derive the island branch and the \(S_{\rm Bek}\) normalization, and keep the \(2S_0\) price.

Remove the AMMZ(10) "summed over endpoints" citation; cite AMMZ(15) for the multi-endpoint area. State the endpoint count in `island_region` and in the QES stationarity conditions.

**Affected area:** R4.

**Disclosure:** Decision 2 is disclosed as a director ruling, but the inconsistency with AMMZ's \(2S_0\) argument is not disclosed. `?OB-7` (F5) mentions \(S_0\ll S^i_{\rm Bek}\) but not the endpoint-count mismatch.

---

### F-05 — MATERIAL — Time-coordinate collision between AEMM's \(u\) and Poincaré \(t=f(u)\)

**Affected file and exact line/section:** `phys-pseudo-r1r4-v0.1.md` lines 14–15 (`coords_poincare`, `coords_physical`), line 34 (`evap_law` \(E(u)\)), line 61 (`qes_position`), line 63 (`island_value` \(T(t)=T_1e^{-kt/2}\)), line 66 (`page_condition`), line 67 (`pagetime_value`); `branch-reconstruction.md` lines 28–31.

**Mathematical or physical issue:** AEMM distinguishes the physical boundary time \(u\) from the Poincaré time \(t=f(u)\). The evaporation law is \(E(u)=E_1e^{-ku}\), hence \(T(u)=T_1e^{-ku/2}\). AMMZ's formulas use \(t\), which in AMMZ is the boundary time at which the entropy is computed. The spec uses \(t\) for the Poincaré time in line 14 and for the AMMZ boundary time in Part B, and writes the AEMM decay as \(T(t)=T_1e^{-kt/2}\). The rate identity \(\kappa=k\) is derived in \(u\) but applied in \(t\). This is a coordinate collision, not a notation choice: the map \(f(u)\) is nontrivial and load-bearing in the cutoff \(z=\epsilon f'(u)\) and in the stress transformation.

**Independent evidence or derivation:** AEMM printed (2.5)–(2.9) define \(x^\pm=t\pm z\), \(t=f(u)\), and the boundary condition. The decay law is in \(u\). If \(t_{\rm AMMZ}=u_{\rm AEMM}\), then \(T(t)=T_1e^{-kt/2}\) is correct. If \(t_{\rm AMMZ}=t_{\rm AEMM}=f(u)\), the decay is \(T(f(u))\), not \(T(u)\); the Page-time integral is in \(t\), and the relation to \(u\) is nonlinear. The package never states which identification is used. The cutoff \(z=\epsilon f'(u)\) is in the \(x\)-frame; the no-island formula's \(\epsilon\) is the constant \(y\)-frame cutoff. The same collision appears in `link_geometry_factors` (F-10).

**Primary source and exact locator:** AEMM printed (2.5)–(2.9), PDF pp.9–10; AEMM printed (2.15), (2.20), (2.25), PDF pp.11–13; AMMZ printed (13)–(14), PDF pp.9, 13.

**Consequence for the research conclusion:** The \(\kappa=k\) identity and the Page-time comparisons are conditional on an unstated time identification. The R1/R4 interop is not complete. The F1 self-check uses \(t\) ambiguously.

**Minimum correction required:** Introduce two distinct symbols (for example \(u_{\rm bdy}\) and \(t_{\rm Poinc}\)), state the identification \(t_{\rm AMMZ}=u_{\rm bdy}\), and rewrite every Part B formula in the chosen time. Re-derive the rate identity and the cutoff transformation in that variable. If the identification is \(t_{\rm AMMZ}=f(u)\), re-derive the decay and the Page condition accordingly.

**Affected area:** R1 / R4 / interoperability.

**Disclosure:** Not disclosed. `?OB-5` concerns shock-vs-smooth, not the time-coordinate collision.

---

### F-06 — MATERIAL — Route (i) claims "NO 2D-CFT entropy computation," but the no-island branch is a 2D CFT computation

**Affected file and exact line/section:** `working-baseline.md` §10 lines 269–271 ("NO 2D-CFT entropy computation is used in the R4 chain"); `phys-pseudo-r1r4-v0.1.md` line 58 (`noisland_branch`), line 59 (`noisland_slope`); `branch-reconstruction.md` §1.

**Mathematical or physical issue:** The route decision says R4 follows AMMZ's 3D-RT/HRT construction and uses no 2D-CFT entropy computation. But the no-island branch is exactly the standard 2D CFT single-interval entropy, and its slope is derived from that formula. The island branch may use 3D RT; the no-island branch does not. The claim is false as written. It also obscures the fact that R4 mixes a 2D CFT no-island branch with a 3D RT island branch, which requires a matching argument.

**Independent evidence or derivation:** The log formula (4.1) is the standard 2D CFT Cardy–Tonni result; it is not a 3D RT computation. AMMZ(14) is also a 2D CFT entropy result (the integral of \(T(t)\)). The route decision conflates "the radiation entropy calculation" with "the island branch." The package's own `noisland_branch` is a 2D CFT computation. The fact that the island branch uses 3D RT does not make the whole R4 chain route (i).

**Primary source and exact locator:** AMMZ printed (14), PDF p.13; AMMZ §2 (holographic CFT\(_2\)), PDF p.4; AEMM printed (4.44), PDF p.37.

**Consequence for the research conclusion:** The route decision's scope is overstated. The R4 entropy chain is a hybrid, not a single route. This weakens the claim that one route computes the whole Page curve and complicates the compatibility argument.

**Minimum correction required:** Restate route (i) as applying to the island branch (3D RT) while the no-island branch is the 2D CFT formula; state the matching conditions and the regime where both are valid. Alternatively, derive the no-island branch from the 3D RT construction (if possible) and remove the 2D CFT formula. Do not claim "NO 2D-CFT entropy computation" while using one.

**Affected area:** R4 / interoperability.

**Disclosure:** The route decision is disclosed, but the contradiction is not. `?OB-10` covers the RT import, not the no-island branch.

---

### F-07 — MATERIAL — Essential quantities are undeclared; the spec is not self-contained

**Affected file and exact line/section:** `phys-pseudo-r1r4-v0.1.md` throughout, especially line 12 (`S_0`, `S^i_{\rm Bek}`, `phibar_r`, `c`), line 22 (`phi0`), line 53 (`S_matter`), line 54 (`phi(a)`), line 55 (`S_BH`, `phi_hor`), line 58 (`beta`, `eps`, `T_H`), line 61 (`T`, `S_Bek`), line 63 (`T_1`, `k`), line 67 (`t_P`), line 70 (`S_eff`), line 74 (`S_0`), line 77 (`G(eta)`), line 84 (`E_S`).

**Mathematical or physical issue:** PSEUDO §2.1 requires declaration before use: every name in a premise list must be introduced by an earlier statement. PSEUDO §2.6 requires every **essential** quantity to be introduced by an earlier `given`/`define`/`assume` and to appear among the step's premises when it carries the step. The spec introduces many symbols only in prose or not at all. The following are essential and undeclared in the spec: \(S_{\rm Bek}\), \(S_0\), \(\beta\), \(\epsilon\), \(T_H\), \(T_1\), \(E_1\), \(E_0\), \(E_S\), \(k\), \(\kappa\), \(y_e\), \(a\), \(t_P\), \(S_{\rm matter}\), \(S_{\rm eff}\), \(S_{\rm no}\), \(S_{\rm isl}\), \(M_{\rm grav}\), \(\phi_b\), \(\bar\phi_r\), \(\phi_{\rm hor}\), \(T\), \(T_0\), \(c\), \(S^i_{\rm Bek}\). Some are in `working-baseline.md`, but the spec is the document under conformance review, and the baseline is a separate file. The spec's header says "Premises: working-baseline.md only," which does not satisfy declaration-before-use inside the spec.

**Independent evidence or derivation:** Check the premise lists:
- `derive noisland_slope` from `[noisland_branch]` uses \(T_H\), \(\beta\), \(\epsilon\), which are not names.
- `derive page_condition` uses \(t_P\), which is not declared before.
- `claim pagetime_value` uses \(S^i_{\rm Bek}\), which is not declared.
- `derive link_dilaton_area` uses \(\phi_0\), \(\phi(a)\), \(G_N\), which are not introduced by a `define`/`given` in the spec.
The baseline dictionary is a separate document; the spec does not import it as a named `given`. A reader cannot resolve the symbols from the spec alone. This is a declaration-before-use failure and an essential-quantity failure.

**Primary source and exact locator:** PSEUDO.md §2.1, §2.6, §4.

**Consequence for the research conclusion:** The specification is not self-contained or conforming. The derivation premise lists are incomplete, and the "single-model, traceable" claim is undermined. This is a MATERIAL conformance defect, not a mathematical error in the equations.

**Minimum correction required:** Add a symbol dictionary to the spec (as `define`/`given` lines or a prose appendix) and declare each essential quantity before first use. Add the symbols to the relevant premise lists, or state explicitly that they are background notation. Alternatively, make the baseline a normative part of the spec by reference—but PSEUDO conformance still requires the spec itself to be readable.

**Affected area:** Conformance.

**Disclosure:** Not disclosed. The baseline has a dictionary, but the spec does not.

---

### F-08 — MATERIAL — The `[SR]/[D]/[C]/[P]/[A]/[U]` status taxonomy violates PSEUDO §7

**Affected file and exact line/section:** `phys-pseudo-r1r4-v0.1.md` lines 4–6 (header) and the tags throughout, e.g. line 11 `[SR: ...]`, line 12 `[A/C: ...]`, line 19 `[P: ...]`, line 23 `[U: ...]`, line 25 `[SR: ...]`, line 32 `[SR: ...]`, line 58 `[SR: ...]`, line 67 `[C: ...]`, line 69 `[D]`, line 74 `[SR: ...]`.

**Mathematical or physical issue:** PSEUDO §1 says the seven constructs are the complete vocabulary. PSEUDO §7 says: "No author-typed status. Words like `PROVEN`, `VERIFIED`, `SUPPORTED`, `CONFIRMED` as labels or claims of status are non-conforming wherever they appear; the vocabulary licenses no such marker." PSEUDO §8 says: "Do not add provenance operators, evidence fields, citation registries, or an evidence-status taxonomy to the language." The spec's `[SR]/[D]/[C]/[P]/[A]/[U]` tags are an evidence-status taxonomy. The header's claim "Status tags (prose only, per PSEUDO.md)" is false: PSEUDO permits prose attribution, but not a status taxonomy. The tags are used as labels on claims and derivations, not merely as source notes.

**Independent evidence or derivation:** PSEUDO §7 lines 239–243; §8 lines 268–269. The tags `[D]=DERIVED-in-spec`, `[C]=CONDITIONAL`, `[P]=POSTULATED`, `[A]=APPROXIMATE`, `[U]=UNSUPPORTED` are exactly status labels. The `[SR]` tag could be recast as prose attribution ("AEMM eq. (2.1)"), but the rest cannot. The header's "per PSEUDO.md" is a false claim of conformance. The tags also mask the epistemic gaps: `[D]` on `page_curve` overstates a qualitative/conditional result (F-16).

**Primary source and exact locator:** PSEUDO.md §1, §7, §8.

**Consequence for the research conclusion:** The spec is non-conforming. The status labels make the document look more authoritative than its content supports, and they conflict with the review's obligation to judge epistemic status independently.

**Minimum correction required:** Remove the status-tag taxonomy. Express source attribution in prose ("AEMM printed (2.1)–(2.3)") and express uncertainty with `?` where genuinely unresolved. Do not use `[D]/[C]/[P]/[A]/[U]`. If the tags are retained as prose abbreviations, they must not function as status claims, and the header must not claim PSEUDO conformance.

**Affected area:** Conformance / source-epistemic.

**Disclosure:** Not disclosed; the header claims conformance.

---

### F-09 — BLOCKER — Several interoperability links are circular or restatements

**Affected file and exact line/section:** `phys-pseudo-r1r4-v0.1.md` line 84 (`link_state_stress`), line 87 (`link_geometry_factors`), line 88 (`link_dilaton_area`), line 89 (`link_matter_sgen`), line 90 (`link_sgen_qes`), line 91 (`link_branches_rad`), line 92 (`link_flux_mass`), line 93 (`interop_complete`).

**Mathematical or physical issue:** Part C claims ten shared links that establish R1/R4 interoperability. Several of the `derive` statements conclude their own premise:
- `link_sgen_qes` derives "qes_position extremizes sgen_functional" from `[sgen_functional, qes_stationarity, qes_position]`. The conclusion is a restatement of the premise `qes_position` plus the definition of `qes_stationarity`; no new connection is established.
- `link_branches_rad` derives "page_curve gives S(R_t) via branch_selection" from `[page_curve, branch_selection, page_condition]`. This is `page_curve` restated.
- `link_flux_mass` derives "global_balance connects evap_law decay to bath flux through energy_balance" from `[energy_balance, evap_law, global_balance]`. This is `global_balance` restated.
- `link_matter_sgen` derives "sgen_functional equals dilaton_term plus matter entropy" from `[sgen_functional, dilaton_term, holographic_entropy_route]`. If `sgen_functional` is defined as that sum, the link is a definitional restatement.
- `link_dilaton_area` derives the one-endpoint area law from `[dilaton_term, bh_entropy_fn, endpoint_count_one]`, which already contain the answer.
- `link_state_stress` derives the mutual fixing of `shock_stress` and `quench_state` from `[quench_state, shock_stress, conformal_matter]`; the conclusion is in the premises.
- `link_geometry_factors` cites `noisland_branch` as a premise while using a cutoff \(z=\epsilon f'(u)\) that `noisland_branch` does not use (see F-10). This is not a circularity but a false dependency.

**Independent evidence or derivation:** A non-circular `derive` must have a conclusion not present (explicitly or by definition) in its premise set. For each link above, the conclusion follows by substitution or by the definition of one of the premises. The only links that add real content are `link_action_eqs` (single variational origin, verified in §3.1), `link_anomaly_stress` (anomaly + Ward → chiral flux, verified in §3.2), and, conditionally, `link_state_entropy` (the 3D-RT route, but that is an import; see F-11). That is at most three of ten, not ten.

**Primary source and exact locator:** PSEUDO.md §3 (derivation must rest on the listed premises and add a step), §4 (dependency graph); AMMZ §3.3, PDF p.14 (branch minimization); AEMM printed (2.17), (2.22), (2.25), PDF pp.11–13 (energy balance and flux).

**Consequence for the research conclusion:** The claimed ten-link interoperability is not established. `interop_complete` line 93 asserts "all ten prompt9 §4C connections established above or explicitly obligated," but the circular links establish nothing, and the obligating `?` at line 94 covers only OB-4, OB-5, OB-9, OB-10, not the circularity. The interoperability central claim is therefore unsupported. This is a BLOCKER because interoperability completeness is a central conclusion of the package.

**Minimum correction required:** For each link, either (a) state the substantive derivation step and premises that do not already contain the conclusion, or (b) downgrade it from `derive` to a definition/restatement and remove it from the count of established links. Recount the interoperability links honestly. If the count falls below the prompt's requirement, say so and mark the gap as a `?`.

**Affected area:** Interoperability / conformance.

**Disclosure:** `?OB-4`, `?OB-5`, `?OB-9`, `?OB-10` are disclosed as carriers. The circularity is not disclosed; `interop_complete` is labeled `[C]` but claims the links are "established above," which is false for the circular ones.

---

### F-10 — MATERIAL — The cutoff frame is inconsistent between `link_geometry_factors` and `noisland_branch`

**Affected file and exact line/section:** `phys-pseudo-r1r4-v0.1.md` line 58 (`noisland_branch`, constant \(\epsilon\)), line 87 (`link_geometry_factors`, cutoff \(z=\epsilon f'(u)\)); `working-baseline.md` §2 (physical cutoff), §6.3.

**Mathematical or physical issue:** `link_geometry_factors` uses the physical cutoff \(z=\epsilon f'(u)\), correctly derived from the boundary condition \(g_{uu}|_{\rm bdy}=1/\epsilon^2\). `noisland_branch` uses the fixed-\(\beta\) CFT formula with a constant cutoff \(\epsilon\) in the \(y\)-frame. Under the Weyl transformation from the \(y\)-frame to the \(x\)-frame, the entropy shifts by a \(-\frac{c}{6}\ln f'(u)\) term (the standard anomalous transformation of the interval entropy). The package uses both cutoffs without relating them. The additive constant in `noisland_slope` is therefore frame-dependent and not the constant that enters the intersection.

**Independent evidence or derivation:** The 2D CFT interval entropy transforms under \(y\to x\) as
\[
S_x=S_y+\frac{c}{6}\ln\left|\frac{dy}{dx}\right|_{\rm endpoints}.
\]
For the boundary cutoff, \(z=\epsilon f'(u)\) versus \(z=\epsilon\), the difference is
\[
\Delta S=\frac{c}{6}\ln f'(u).
\]
At \(u=0\), \(f'(0)=1\), so the shift vanishes at the quench; it grows at late times. The physical entropy must be computed with the same cutoff in both branches. The package's `noisland_slope` constant \((c/6)\ln(\beta/2\pi\epsilon)\) is the constant-cutoff result; `link_geometry_factors` claims the \(f'\)-cutoff. These are not the same observable.

**Primary source and exact locator:** AEMM printed (2.12)–(2.14), PDF pp.9–10 (boundary condition and cutoff); AEMM printed (2.16) and §3.2 (stress transformation); AMMZ printed (9), PDF p.8 (cutoff relation \(z_w=\epsilon e^{-\rho}\sqrt{dw^+/dy^+\,dw^-/dy^-}\)).

**Consequence for the research conclusion:** The additive constant entering the branch intersection is not consistently defined. The intersection (and the Page time at subleading order) is frame-dependent as written. The leading slope is unaffected; the `pagetime_value` and the `island_disfavored` threshold at \(O(1)\) are affected.

**Minimum correction required:** State the frame of every entropy formula. Use \(z=\epsilon f'(u)\) consistently and include the \(-\frac{c}{6}\ln f'(u)\) transformation in the no-island branch, or compute both branches in the constant-cutoff frame and state why the \(f'\) factor drops. Record the additive constant with its frame label.

**Affected area:** R4 / interoperability.

**Disclosure:** The cutoff issue is not disclosed. The baseline §2 records the physical cutoff correctly, but the spec mixes frames.

---

### F-11 — MATERIAL — QES and RT/HRT are disclosed imports, but the R4 text presents them as derived

**Affected file and exact line/section:** `phys-pseudo-r1r4-v0.1.md` line 61 (`qes_position`), line 64 (`branch_selection`), line 65 (`extremization_then_minimization`), line 69 (`page_curve`), line 70 (`island_rule_status`), line 86 (`?OB-10`), line 90 (`link_sgen_qes`); `registers.md` §5 (`?OB-9`, `?OB-10`); `summary.md` lines 8–14.

**Mathematical or physical issue:** The QES position AMMZ(12) is imported from AMMZ's citations [8,9]; the JT-specific derivation is in AEMM §4. The 3D-RT island computation is imported wholesale from AMMZ §3. Both are disclosed as `?OB-9` and `?OB-10`. The problem is not the import per se; it is that the spec and summary present the resulting R4 claims as derived ("DERIVED WITHIN THIS SPECIFICATION: ... branch intersection with footprint F1–F5; ten R1/R4 links"; `page_curve` tagged `[D]`). An imported result carried as an open obligation cannot simultaneously be a derived result of the specification. The status is inconsistent.

**Independent evidence or derivation:** `qes_position` line 61 cites "AMMZ(12)←Penington+AEMM §4 chain (`?OB-9` carried)." The `?` is carried, so the result is conditional. `page_curve` line 69 is tagged `[D]`, and `summary.md` lists the branch intersection as derived. The two statuses conflict. The import is central: without AMMZ(12) and the 3D-RT construction, there is no island branch and no Page curve. A disclosed `?` on a central result can still be MATERIAL (per the approved governance amendment); the disclosure does not cure the status conflict.

**Primary source and exact locator:** AMMZ printed (12), PDF p.9; AMMZ §3, PDF pp.12–15; AEMM §4, PDF pp.30–37; Penington Eq. (2), PDF p.5; PSEUDO.md §4–§5 (a `?` is an open question, not a derived result).

**Consequence for the research conclusion:** The R4 result is conditional on two unverified imports. The summary's "DERIVED WITHIN THIS SPECIFICATION" overstates the epistemic status. The Page curve is not independently derived in the package.

**Minimum correction required:** Tag every result that depends on `?OB-9`/`?OB-10` as conditional/imported, not derived. Either reconstruct the QES/RT steps or state plainly that the R4 chain is "conditional on AMMZ(12)/§3." Remove `page_curve` from the "derived" list in the summary. Do not claim both "derived" and "carried ?" for the same result.

**Affected area:** R4 / source-epistemic.

**Disclosure:** The imports are disclosed (`?OB-9`, `?OB-10`). The status conflict is not.

---

### F-12 — MINOR — The AEMM locator convention is undocumented and does not match the printed equation numbers

**Affected file and exact line/section:** `phys-pseudo-r1r4-v0.1.md` lines 6, 11, 25, 28, 34, 83, 90, 92 (AEMM(6)–(112) locators); `dossier-audit-note-v1.0.md` lines 12–13 ("127 numbered eqs, single-sequence"); `sources/README.md` AEMM entry.

**Mathematical or physical issue:** The AEMM PDF prints sectional equation numbers \((2.1)\)–\((5.10)\), not one global sequence. The package's AEMM(\(n\)) is the ordinal of the \(n\)-th numbered equation group, counting subequations as one group. Under that convention the audit note's "single-sequence" statement is defensible, but the convention itself is nowhere documented. A reader checking "AEMM(11a–c)" against the PDF finds printed \((2.6a–c)\); a reader checking "AEMM(24)" finds printed \((2.20)\); a reader checking "AEMM(112)" finds printed \((4.44)\). Without the mapping table, the locators appear wrong even when the content is right.

**Independent evidence or derivation:** The verified group-ordinal mapping is:
(6)–(8) = (2.1)–(2.3); (9) = (2.4); (10) = (2.5); (11a–c) = (2.6a–c); (12) = (2.7); (13)–(14) = (2.8)–(2.9); (20) = (2.15); (21)–(25) = (2.16)–(2.20); (42) = (3.4); (44)–(46) = (3.6)–(3.8); (53) = (3.15); (56)–(60) = (3.18)–(3.22); (61)–(68) = (3.23)–(3.30); (69)–(73) = (4.1)–(4.5); (75)–(77) = (4.7)–(4.9); (92)–(93) = (4.24)–(4.25); (104)–(107) = (4.36)–(4.39); (108) = (4.40); (112) = (4.44). The mapping is bijective over the checked set. The content-level locators are therefore faithful under the convention; the defect is the undocumented convention, not a false equation claim.

**Primary source and exact locator:** AEMM arXiv:1905.08762v3 PDF, printed equations throughout; Appendix B of this report.

**Consequence for the research conclusion:** A verification reader cannot reproduce the locators without reconstructing the convention. This is a traceability defect, not a mathematical error. It is MINOR because the equations themselves are correct and the mapping is recoverable.

**Minimum correction required:** Document the convention in the spec header and in the registers ("AEMM(\(n\)) = ordinal of the \(n\)-th numbered equation group, counting subequations as one group"), or replace every locator with the printed sectional number. Reconcile the audit note's "single-sequence" wording.

**Affected area:** Source-epistemic / traceability.

**Disclosure:** The audit note claims "single-sequence" without defining it. The convention is implicit in the package.

---

### F-13 — MATERIAL — RC1's "away from extremality" regime contradicts the adopted near-extremal AEMM model

**Affected file and exact line/section:** `reality-checks.md` lines 8–11 (JT dilaton gravity "Regime: semiclassical, \(k\ll1\), away from extremality"); `phys-pseudo-r1r4-v0.1.md` line 12 (AEMM hierarchy and near-extremal condition); `working-baseline.md` lines 220–226.

**Mathematical or physical issue:** RC1 states that the JT dilaton-gravity landscape applies "away from extremality." The adopted AEMM model is explicitly near-extremal: AEMM imposes \(\phi_0\gg\phi\), and the dimensional reduction is the near-extremal throat. The two statements are not the same regime. The RC1 entry may be describing the general JT framework, but the package applies it to the near-extremal AEMM model without noting the transition. The mismatch matters because the near-extremal condition is exactly the one that conflicts with the island-domination hierarchy (F-01).

**Independent evidence or derivation:** AEMM printed (2.1)–(2.3) and the text after (2.3), PDF p.8, state the near-extremal restriction \(\phi_0\gg\phi\). The near-extremal throat is the setting in which the Schwarzian dynamics and the \(k\ll1\) evaporation law are derived. RC1's "away from extremality" is a valid general statement about JT gravity, but it is not the regime of the adopted model. A landscape entry can be broader than the model; the defect is that the entry is presented as the model's regime without qualification.

**Primary source and exact locator:** AEMM PDF p.8 (printed p.7), text after (2.3); AEMM printed (2.24), PDF p.12; AMMZ §2, PDF p.4 (near-extremal condition \(c\ll(\phi-\phi_0)/4G_N\)).

**Consequence for the research conclusion:** RC1 gives a misleading regime statement for the adopted model. A reader applying RC1's validity claim to the package would import the wrong extremality assumption. This is a MATERIAL reality-check defect because it obscures the F-01 conflict.

**Minimum correction required:** Rewrite the RC1 JT entry to distinguish the general JT framework from the adopted near-extremal AEMM regime, and state the extremality condition explicitly. Do not say "away from extremality" as the model's regime.

**Affected area:** R1 / conformance / reality checks.

**Disclosure:** Not disclosed. The near-extremal condition is in the baseline but not connected to RC1.

---

### F-14 — MATERIAL — The summary overclaims "MATHEMATICALLY ESTABLISHED," "DERIVED," and "single-model, traceable"

**Affected file and exact line/section:** `summary.md` lines 3–11 ("MATHEMATICALLY ESTABLISHED (in sources, imported)"; "DERIVED WITHIN THIS SPECIFICATION: ... branch intersection with footprint F1–F5; ten R1/R4 links"); lines 12–14 (imports); lines 15–18 (conditional/unresolved); `phys-pseudo-r1r4-v0.1.md` line 93.

**Mathematical or physical issue:** The summary's status claims are stronger than the content supports:
- "MATHEMATICALLY ESTABLISHED (in sources, imported)" is acceptable for source facts but is applied to the package's own R1/R4 reconstruction ("corrected R1 equation set with verified signs/coefficients").
- "DERIVED WITHIN THIS SPECIFICATION" includes the one-endpoint normalization (a postulate, Decision 2), the branch intersection with F1–F5 (an estimate with a failed self-check; F-03), and the ten links (at most three non-circular; F-09).
- "single-model, traceable" is contradicted by F-01 (the two hierarchies do not describe one model) and F-05 (the time coordinate is not one variable).
- The Page-time value is called conditional in one line and derived in another.

**Independent evidence or derivation:** Compare the summary's status list with §3 of this report: the R1 algebra is verified; the R4 branch formulas are not; the links are not ten. The summary's own "UNRESOLVED" line admits `?OB-1..10` and the interop carriers, which is incompatible with "MATHEMATICALLY ESTABLISHED" for the package's reconstruction. The overclaim is not a mathematical error; it is an epistemic error that affects how the package is used.

**Primary source and exact locator:** PSEUDO.md §7 (no author-typed status); §8 (no evidence-status taxonomy); the source statuses themselves (AEMM/AEMM/AMMZ) do not establish the package's reconstruction.

**Consequence for the research conclusion:** The summary would lead a reader to treat the R4 Page curve and the interop as established. They are not. This is MATERIAL because it affects the approval decision.

**Minimum correction required:** Rewrite the summary to separate (a) source-established results, (b) independently reconstructed results, (c) conditional/imported results, and (d) open obligations, with no author-typed status vocabulary. State the one-endpoint postulate as a postulate, the Page time as a leading estimate, and the links as a reduced count.

**Affected area:** Source-epistemic / overall readiness.

**Disclosure:** The summary discloses the unresolved list, but the headline statuses contradict it.

---

### F-15 — MATERIAL — `?OB-6` is dropped from the spec and `?OB-8followup` is absent from the register

**Affected file and exact line/section:** `phys-pseudo-r1r4-v0.1.md` line 58 (`noisland_branch`, \(S_0\) term dropped without a `?`); `registers.md` §5 (OB-1–OB-10; `?OB-8followup` not listed); `phys-pseudo-r1r4-v0.1.md` line 78 (`?OB-8followup` appears in the spec).

**Mathematical or physical issue:** `branch-reconstruction.md` line 10 says the \(S_0\) term is dropped from the no-island branch "per ?OB-6." The spec's `noisland_branch` does not carry `?OB-6`; the dropped term is silently absent. Separately, the spec line 78 introduces `?OB-8followup`, but the registers' obligation list contains no such ID. The obligation tracking is therefore not closed: one disclosed obligation disappears, and one spec obligation has no register entry.

**Independent evidence or derivation:** Read the spec line 58: no `?OB-6` and no \(S_0\) term. Read the registers §5: OB-1–OB-10, no OB-8followup. The branch reconstruction explicitly invokes `?OB-6`. The traceability chain spec ↔ registers ↔ branch reconstruction is broken in both directions.

**Primary source and exact locator:** Not a source issue; PSEUDO.md §5–§6 (obligations must be traceable) and the package's own register.

**Consequence for the research conclusion:** A dropped entropy term and an untracked obligation are invisible to a reader of the spec. The no-island branch's normalization is not fully specified. This is MATERIAL because the dropped \(S_0\) affects the intersection at \(O(1)\) in the large-\(S_0\) regime and because the register is the package's accountability mechanism.

**Minimum correction required:** Add `?OB-6` to the spec at `noisland_branch` and state the dropped \(S_0\) term and its justification; add `?OB-8followup` to the registers with its content and closure condition, or remove it from the spec if it is not a real obligation. Reconcile the three documents.

**Affected area:** R4 / conformance.

**Disclosure:** OB-6 is disclosed in the branch reconstruction, not in the spec; OB-8followup is disclosed in the spec, not in the registers.

---

### F-16 — MATERIAL — `page_curve` and `pagetime_value` statuses overstate a qualitative/conditional result

**Affected file and exact line/section:** `phys-pseudo-r1r4-v0.1.md` line 69 (`claim page_curve` tagged `[D]`), line 67 (`claim pagetime_value`), line 93 (`interop_complete`); `summary.md` lines 8–11.

**Mathematical or physical issue:** `page_curve` is tagged `[D]` (derived in spec). What is actually established is a qualitative/structural expectation: under the island-domination hierarchy, the island branch can dominate after a crossing, so the entropy follows a rising-then-falling curve of the Page form. The quantitative curve is not derived: the no-island branch is not AMMZ(14) (F-02), the exact intersection is not \(1/\kappa\) (F-03), the endpoint normalization is inconsistent (F-04), and the QES/RT inputs are imports (F-11). `pagetime_value` is a leading estimate, not a derived value, and the F1 self-check fails. The `[D]` tag and the summary's "DERIVED" claim overstate both.

**Independent evidence or derivation:** The branch comparison in §3.8 gives three different \(t_P\) values depending on the branch; the package's value is the middle heuristic. The curve's shape (rise then fall) is robust under the island rule; the value and the branch are not. "Qualitative Page-like curve" is the correct status; "derived Page curve" is not.

**Primary source and exact locator:** AMMZ §3.3, PDF p.14 (branch minimization); AMMZ(13)–(14), PDF pp.9, 13; Penington Eq. (2), PDF p.5.

**Consequence for the research conclusion:** The central R4 claim is presented at a higher epistemic level than the evidence supports. A reviewer or implementer could treat the curve and time as derived. This is MATERIAL.

**Minimum correction required:** Downgrade `page_curve` to a conditional/structural result with the conditions stated (island domination, one/two-endpoint choice, branch formula). Keep `pagetime_value` as a leading estimate with the spread \(0.703/\kappa\)–\(0.811/\kappa\) and the F1 caveat. Remove `[D]` from both.

**Affected area:** R4 / source-epistemic.

**Disclosure:** `?OB-7` and F1–F5 are disclosed; the `[D]` tag contradicts them.

---

### F-17 — MATERIAL — The hierarchy \(c\ll S_0\) is not AMMZ's condition, and the "equivalently" step has a factor-4 error

**Affected file and exact line/section:** `working-baseline.md` lines 222–224; `phys-pseudo-r1r4-v0.1.md` line 12; `registers.md` line 21 (ASM-02); `audit-corrections-v1.1.md` (ASM-02 disposition).

**Mathematical or physical issue:** The baseline writes
\[
1\ll c\ll S_0=\frac{\phi_0}{4G_N}\ll S^i_{\rm Bek},
\qquad\text{"equivalently }c\ll\frac{\phi_0}{G_N}\text{."}
\]
The equivalence is wrong by a factor of four: \(c\ll S_0\) is \(c\ll\phi_0/(4G_N)\), not \(c\ll\phi_0/G_N\). More importantly, AMMZ's near-extremal condition is not \(c\ll S_0\). It is
\[
c\ll\frac{\phi-\phi_0}{4G_N}=S_{\rm Bek}-S_0,
\]
i.e. the **dynamical** entropy \(S_{\rm Bek}-S_0\) must exceed \(c\). The baseline calls the full \(S_{\rm Bek}\) the "dynamical BH entropy" and uses \(c\ll S_0\) as the hierarchy. These are different regimes:
- \(c\ll S_0\) says the extremal entropy is large compared with \(c\).
- AMMZ's condition says the entropy above extremality is large compared with \(c\).
Under AEMM's \(\phi_0\gg\phi\), \(S_0\) is large and \(S_{\rm Bek}-S_0=\phi_{\rm hor}/4G_N\) can be small; AMMZ's condition then fails even though \(c\ll S_0\) holds. Conversely, under island domination \(S_0\ll S_{\rm Bek}\), the dynamical entropy is large, but \(S_0\) itself may be comparable to \(c\).

**Independent evidence or derivation:** AMMZ §2, PDF p.4, states verbatim: "In certain theories corresponding to near-extremal black holes, we also require that \(c\ll(\phi-\phi_0)/(4G_N^{(2)})\), where the \(\phi_0\) piece corresponds to the extremal entropy." The condition is on the difference. The baseline's \(c\ll S_0\) is not that inequality. The factor-4 error is elementary: dividing \(\phi_0\) by \(4G_N\) versus \(G_N\).

**Primary source and exact locator:** AMMZ §2, PDF p.4 (the \(c\ll(\phi-\phi_0)/4G_N\) condition); AMMZ printed (12), PDF p.9 (\(S_{\rm Bek}-S_0\) in the logarithm); AEMM printed (2.24), PDF p.12 (the \(k\) hierarchy).

**Consequence for the research conclusion:** The stated hierarchy does not guarantee the validity of the holographic/QES approximations the package relies on. The "equivalently" step is numerically wrong. This compounds F-01: the package's hierarchy is neither AEMM's nor AMMZ's.

**Minimum correction required:** Replace the hierarchy with the correct AMMZ condition \(c\ll S_{\rm Bek}-S_0\), state separately whether \(S_0\) is large or small, and remove the factor-4 "equivalently" statement. Distinguish \(S_{\rm Bek}\) (total) from \(S_{\rm Bek}-S_0\) (dynamical) throughout the baseline.

**Affected area:** R1 / R4 / source fidelity.

**Disclosure:** ASM-02 is marked "REWRITTEN," but the rewrite retains the error. Not disclosed.

---

### F-18 — MATERIAL — R1 uses a \(\delta\)-shock quench while R4 uses AMMZ's finite-width pulse; the state map is not established

**Affected file and exact line/section:** `phys-pseudo-r1r4-v0.1.md` line 84 (`link_state_stress`), line 85 (`link_state_entropy`), line 94 (`interop_carriers`); `working-baseline.md` lines 144–145 ("same transient physics, holographic dual language"); `registers.md` §5 (`?OB-5`).

**Mathematical or physical issue:** AEMM's quench is an instantaneous shock at \(u=0\) with stress \(E_S\delta(x^-)\) (printed (2.16)). AMMZ's §3.2.2 coupling is a finite-width pulse over \(\Delta t\), with energy \(E\sim c/\Delta t\), applied to the bath. The two are different states: the shock is the \(\Delta t\to0\) limit of the pulse, but the limit is singular (the injected energy diverges as \(1/\Delta t\) at fixed \(c\)), and the entropy evolution at early times depends on the width. The package asserts the two are "same transient physics, holographic dual language" without deriving the map. `?OB-5` is open, and `link_state_stress` is a restatement (F-09), so the state compatibility is not established anywhere.

**Independent evidence or derivation:** AEMM (2.16) is a delta function; AMMZ §3.2.2 describes a coupling turned on over a finite interval and computes the resulting stress. A delta shock has infinite instantaneous energy at fixed \(E_S\); a finite pulse has finite energy. The entropy produced depends on the pulse shape at \(O(c\log)\); the late-time slope may be universal, but the package uses the early-time branch (AMMZ(14)) where the difference matters. A state map would require showing that the \(\Delta t\to0\) limit of the AMMZ pulse reproduces the AEMM shock with \(E_S\) fixed, or that both flows converge to the same late-time attractor.

**Primary source and exact locator:** AEMM printed (2.16), PDF p.11; AEMM §3.2, PDF pp.13–16; AMMZ §3.2.2, PDF p.12; AMMZ §3.1, PDF pp.9–10.

**Consequence for the research conclusion:** The R1 state and the R4 state are not shown to be the same. The interop `link_state_stress`/`link_state_entropy` are conditional on `?OB-5`. Since the Page curve depends on the state, this is MATERIAL.

**Minimum correction required:** Either derive the map (show the pulse-to-shock limit and its effect on the entropy) or state explicitly that R1 and R4 use different quench models and that the interop is valid only for observables insensitive to the difference. Carry `?OB-5` on the relevant R4 claims. Quantify the \(O(c\log)\) difference if it is claimed subleading.

**Affected area:** R1 / R4 / interoperability.

**Disclosure:** `?OB-5` is disclosed as a carrier. The baseline's "same transient physics" assertion is not supported.

---

### F-19 — MINOR — The no-island formula omits the boundary entropy \(\log g\) without stating the Cardy-brane condition

**Affected file and exact line/section:** `phys-pseudo-r1r4-v0.1.md` line 58 (`noisland_branch`); `branch-reconstruction.md` line 11; `working-baseline.md` §6.3.

**Mathematical or physical issue:** The standard 2D CFT single-interval entropy is
\[
S=\frac{c}{6}\ln\left[\frac{\beta}{\pi\epsilon}\sinh\left(\frac{\pi t}{\beta}\right)\right]+\log g,
\]
where \(\log g\) is the boundary entropy (Affleck–Ludwig). The package omits \(\log g\). For the Cardy brane used by AMMZ, \(\log g=0\), so the omission is correct **for that brane**; but the spec does not state the condition. In a general holographic CFT with a nonzero boundary entropy, the additive constant shifts and the intersection moves.

**Independent evidence or derivation:** The boundary-entropy term is a standard contribution to the entanglement entropy of an interval ending on a boundary. The AMMZ Cardy brane construction has \(\log g=0\) by the choice of boundary state. The package's baseline does not record that choice. Since the additive constant already enters the intersection at subleading order (F-02, F-10), the omitted \(\log g\) is the same order.

**Primary source and exact locator:** AMMZ §2 and §3.1 (Cardy brane), PDF pp.4–10; standard 2D CFT result (not in the pinned sources; the package should state the condition from its own setup).

**Consequence for the research conclusion:** The no-island normalization is correct only under an unstated condition. The effect is a constant shift, so the leading slope and the leading Page time are unaffected; the exact intersection is. MINOR because the leading result survives.

**Minimum correction required:** State that \(\log g=0\) for the adopted Cardy brane, or include \(+\log g\) in the formula and carry it through the intersection.

**Affected area:** R4.

**Disclosure:** Not disclosed.

---

### F-20 — MINOR — AMMZ(10) is cited for a multi-endpoint sum it does not contain

**Affected file and exact line/section:** `phys-pseudo-r1r4-v0.1.md` line 54 (`dilaton_term`), line 88 (`link_dilaton_area`); `branch-reconstruction.md` (island normalization).

**Mathematical or physical issue:** The spec's `dilaton_term` cites AMMZ(10) for "per QES endpoint, summed over endpoints." AMMZ(10) is a single dilaton term \(S_{\rm gen}(y)=\phi(y)/4G_N+S_{\rm bulk}[I_y]\). The multi-endpoint area term is AMMZ(15), \(S(A)=\operatorname{Min}\operatorname{Ext}[S_{\rm eff}(A\cup I_g)+\operatorname{Area}[\partial I_g]/4G_N]\). The citation is therefore to the wrong equation, even though the physics (sum over endpoints) is correct.

**Independent evidence or derivation:** AMMZ(10) contains one \(\phi(y)\); there is no sum. AMMZ(15) contains \(\operatorname{Area}[\partial I_g]\), which is the sum over boundary points. The correct citation for the endpoint sum is (15). This is a citation defect, not a mathematical error.

**Primary source and exact locator:** AMMZ printed (10), PDF pp.7–8; AMMZ printed (15), PDF p.15.

**Consequence for the research conclusion:** Traceability only. The formula is correct given the endpoint postulate; the citation misleads a verifier. MINOR.

**Minimum correction required:** Cite AMMZ(15) for the multi-endpoint area and AMMZ(10) for the single-endpoint generalized entropy. Correct both line 54 and line 88.

**Affected area:** R4 / source fidelity.

**Disclosure:** Not disclosed.

---

### F-21 — MINOR — The AMMZ(12) versus AEMM(4.40) logarithmic factor is unreconciled

**Affected file and exact line/section:** `phys-pseudo-r1r4-v0.1.md` line 61 (`qes_position`), line 90 (`link_sgen_qes`); `registers.md` line 41.

**Mathematical or physical issue:** AMMZ(12) gives
\[
y_e^+=t-\frac{1}{2\pi T(t)}\log\frac{S_{\rm Bek}(T(t))-S_0}{c}+\cdots,
\]
while AEMM(4.40) gives
\[
t_{HP}=\frac{\beta}{2\pi}\log\left[\frac{16}{c}(S-S_0)\right].
\]
The two logarithms differ by an additive \(\log 16\) at fixed \(S-S_0\) (and the identification of \(S\) with \(S_{\rm Bek}\)). The package cites both but does not reconcile the factor. The factor is subleading in the large-\(S\) limit, so the leading Page time is unaffected; but the package treats the QES position as an import with a definite coefficient, and the coefficient is not definite.

**Independent evidence or derivation:** Compare the two expressions at \(T(t)=T\) and \(S=S_{\rm Bek}\): AMMZ's argument is \((S-S_0)/c\); AEMM's is \(16(S-S_0)/c\). The difference is \(\log 16\) in the scrambling time. This can arise from different conventions for the UV cutoff or the central charge normalization; the package does not state which. It is a MINOR discrepancy because the leading large-\(S\) behavior is the same.

**Primary source and exact locator:** AMMZ printed (12), PDF p.9; AEMM printed (4.40), PDF p.35 (printed p.34).

**Consequence for the research conclusion:** The QES position's subleading constant is ambiguous. The leading Page time is unaffected. MINOR.

**Minimum correction required:** State which convention is adopted and reconcile the \(\log 16\), or mark the constant as convention-dependent and not used.

**Affected area:** R4 / source fidelity.

**Disclosure:** Not disclosed; `?OB-9` covers the derivation, not the factor mismatch.

---

### F-22 — MINOR — `ward_identity` is `SOURCE-REPORTED` in the register but derived in the spec

**Affected file and exact line/section:** `registers.md` line 36 (`ward_identity` status `SOURCE-REPORTED`); `phys-pseudo-r1r4-v0.1.md` line 30 (`derive ward_identity` region); `registers.md` line 54 (dependency graph).

**Mathematical or physical issue:** The register lists `ward_identity` as `SOURCE-REPORTED` with limitation "Dossier citation voided," while the spec presents it as a derived step and the dependency graph uses it as an internal node. The two statuses are inconsistent. The physics is standard and the identity is correct (verified in §3.2), so the defect is bookkeeping, not mathematics.

**Independent evidence or derivation:** `registers.md` line 36 says `SOURCE-REPORTED`; the spec's `derive ward_identity` statement (line 30) is a derivation. A claim cannot be both source-reported and derived in the same package without explanation. The correct status is "standard result, independently verified in review" or "source-reported and re-derived"; the package should pick one.

**Primary source and exact locator:** PSEUDO.md §4 (consistent status); AP §5.1, PDF p.15 (the Ward identity is imposed, not derived from the anomaly).

**Consequence for the research conclusion:** Status bookkeeping only. MINOR.

**Minimum correction required:** Reconcile the status: either mark it `derived` (standard diffeomorphism invariance) or mark the spec step as an import. Update the dependency graph accordingly.

**Affected area:** Conformance / source-epistemic.

**Disclosure:** The inconsistency is visible in the package; not flagged.

---

### F-23 — MINOR — The dependency graph routes R4 objects into `noisland_branch` and the "acyclic" claim is undermined by the circular links

**Affected file and exact line/section:** `registers.md` lines 49–64; especially line 58 (`radiation_region + island_region + sgen_functional ─→ noisland_branch`) and lines 62–64 ("No edge points backward ... acyclic").

**Mathematical or physical issue:** The dependency graph makes `noisland_branch` depend on `island_region` and `sgen_functional`, which are island-branch objects. The no-island branch should not depend on the island region; it is the complementary candidate. The edge is wrong. Separately, the graph is called "acyclic," but the Part C links include the circular/restatement edges of F-09 (`link_sgen_qes`, `link_branches_rad`, `link_flux_mass`), which the graph does not display. If displayed, those edges would be self-loops. The acyclicity claim is therefore true only for the reduced graph actually drawn, not for the full Part C graph.

**Independent evidence or derivation:** `noisland_branch` is computed on the radiation region \(R_t\) without an island; it needs the metric, the state, and the cutoff, not `island_region` or `sgen_functional`. The correct graph edge is `radiation_region + coords + cutoff → noisland_branch`. The Part C self-loops are documented in F-09.

**Primary source and exact locator:** Package-internal; PSEUDO.md §4 (dependency graph must reflect the actual derivation).

**Consequence for the research conclusion:** The graph misstates the derivation structure and overstates acyclicity. MINOR because the mathematics of the no-island branch is unaffected; the defect is in the accountability artifact.

**Minimum correction required:** Remove the backward edges from `noisland_branch`; redraw the graph with the Part C links and either eliminate the self-loops (F-09) or mark them as restatements. Re-state the acyclicity claim for the corrected graph.

**Affected area:** Conformance / interoperability.

**Disclosure:** Not disclosed.

---

### F-24 — MINOR — \(S_{\rm BH}\) and \(S_{\rm Bek}\) are used interchangeably without definition

**Affected file and exact line/section:** `phys-pseudo-r1r4-v0.1.md` line 55 (`bh_entropy_fn` uses \(S_{\rm BH}\)), line 61 (`qes_position` uses \(S_{\rm Bek}\)), line 63 (`island_value`); `working-baseline.md` line 222 ("\(S^i_{\rm Bek}\) (dynamical BH entropy)").

**Mathematical or physical issue:** The package uses \(S_{\rm BH}\) and \(S_{\rm Bek}\) as if they are the same quantity. AEMM/AMMZ use \(S_{\rm Bek}\) for the Bekenstein–Hawking entropy of the black hole, and \(S_{\rm BH}\) in the package appears to mean the same. But the package also calls \(S_{\rm Bek}\) the "dynamical BH entropy" while \(S_0\) is the extremal entropy, and uses \(S_{\rm BH}\) in the island branch in a way that includes \(S_0\). The notation should be defined once: \(S_{\rm Bek}=(\phi_0+\phi_{\rm hor})/4G_N\), \(S_0=\phi_0/4G_N\), \(S_{\rm dyn}=S_{\rm Bek}-S_0=\phi_{\rm hor}/4G_N\). The interchangeable use invites the F-17 error.

**Independent evidence or derivation:** AMMZ defines \(S_{\rm Bek}(T)\) as the Bekenstein–Hawking entropy and \(S_0\) as the extremal entropy (PDF p.9). AEMM uses \(S\) for the density of states in (4.40). The package's `bh_entropy_fn` should state whether it returns \(S_{\rm Bek}\) or \(S_{\rm dyn}\). The island branch's \(S_0+(S_i-S_0)x\) uses both. A one-line dictionary would remove the ambiguity.

**Primary source and exact locator:** AMMZ printed (12), PDF p.9; AMMZ §3.3, PDF p.14; AEMM printed (4.40), PDF p.35.

**Consequence for the research conclusion:** Notation ambiguity; contributes to F-17. MINOR as a standalone finding.

**Minimum correction required:** Define \(S_{\rm Bek}\), \(S_0\), and \(S_{\rm dyn}\) once and use them consistently; replace "dynamical BH entropy" for \(S_{\rm Bek}\) with \(S_{\rm dyn}\).

**Affected area:** R4 / source fidelity.

**Disclosure:** Not disclosed.

---

## 5. Assessment of the open obligations OB-1–OB-10

The register in `registers.md` §5 lists ten obligations plus the interop carriers. The table below gives the independent disposition of each: whether it is load-bearing, whether it can be closed from the pinned sources, and what the package must do.

| ID | Obligation (register) | Load-bearing? | Independent assessment | Required action |
|---|---|---|---|---|
| OB-1 | \((K-1)\) counterterm: "shortlisted not-needed; revisit for on-shell finiteness" | No | The AEMM action (2.1)–(2.3) uses the plain Gibbons–Hawking–York term \(2\int K\), not \(K-1\). The package adopts AEMM's action and does not need the counterterm for the equations used. The obligation can be closed by citing AEMM(2.1)–(2.3) and noting that the \(K-1\) variant is not adopted. | Close as "not adopted; AEMM(2.1)–(2.3) is the action." |
| OB-2 | \(R\)-form trace equality: "pin or drop (dropped downstream)" | No | The suspect equality is correctly not used. It can be closed as dropped. | Close as dropped; keep the `[U]` marker off any downstream claim. |
| OB-3 | \(S_{\rm Polyakov}\) coefficient: "suspended; anomaly route used" | No | The anomaly route is independently sufficient (§3.2 of this report). The Polyakov-explicit leg is not needed for R1. | Close as "not needed for the adopted route; anomaly route verified." |
| OB-4 | \(dM/dt\) line-by-line: "OWED in Part A scope" | **Yes** for the evaporation law | The formula \(dE/du=-kE=-(\pi c/12)T^2\) is correct and independently closable (§3.3). The package has not exhibited the line-by-line derivation, so the obligation is real but low-risk. | Add the three-line derivation (energy density, flux, \(T(u)=T_1e^{-ku/2}\)); then close. |
| OB-5 | Shock-vs-smooth branch: "decide via derivation" | **Yes** for R1/R4 state compatibility | Not closed. AEMM's \(\delta\)-shock and AMMZ's finite-width pulse are different states; the map is not derived (F-18). The late-time slope may be universal, but the package uses the early-time branch. | Derive the pulse-to-shock limit or state the interop is limited to universal late-time observables; carry the `?` on the branch claims. |
| OB-6 | \(S_0\) term: "re-add only if sourced" | **Yes** for the no-island normalization | The term is dropped in `branch-reconstruction.md` line 10 and absent from the spec line 58, without a `?` (F-15). The fixed-\(\beta\) log formula indeed has no \(S_0\) term; the issue is that the package does not state where the extremal term went and how it cancels between branches. | State explicitly that the no-island branch has no \(S_0\) term and why; reconcile with the island branch's \(S_0\). Keep the `?` in the spec. |
| OB-7 | \(t_P\) footprint F1–F5: "close before any freeze" | **Yes** for the central R4 value | Not closed. The F1 self-check fails under the AEMM hierarchy (F-03); the exact intersections differ from \(1/\kappa\) by 23–42% (F-02, F-03). | Report the exact intersections and the F1 failure; do not freeze \(t_P\) until F-01 is resolved. |
| OB-8 | Route decision: "CLOSED (route i adopted; fallback documented)" | No (closure itself is the issue) | The route is closed by decision, but its scope is overstated: route (i) claims no 2D-CFT entropy computation while the no-island branch is a 2D CFT computation (F-06). The fallback `?OB-8followup` (spec line 78) is not in the register (F-15). | Restate the route scope; register `?OB-8followup`. |
| OB-9 | AEMM §4 re-derivation: "explicitly deferred, import recorded" | **Yes** for the QES branch | The import is real and central. `qes_position` and the QES stationarity are conditional on it (F-11). The AMMZ(12)/AEMM(4.40) factor-16 mismatch is a further sub-issue (F-21). | Either reconstruct the AEMM §4 QES step or tag every dependent result as imported/conditional. |
| OB-10 | RT reconstruction: "deferred, import recorded" | **Yes** for the island branch | The import is real and central. The 3D-RT construction is not reconstructed; the island branch is conditional on it (F-11). | Same as OB-9: reconstruct or tag as conditional. |
| Carriers | OB-4/5/9/10 | — | The carrier list omits OB-6/OB-7 and the F-09 circularity, so `interop_complete` is not carried by the listed obligations alone. | Extend the carrier list to include OB-6, OB-7, and the circular-link repair. |

**Summary.** OB-1, OB-2, OB-3, and OB-8 are non-load-bearing and closable by statement. OB-4 is independently closable and the formula is correct. OB-5, OB-6, OB-7, OB-9, and OB-10 are load-bearing and remain open; OB-5, OB-9, and OB-10 are central. The disclosed obligations do not, by themselves, block acceptance of the R1 algebra, but they do block the claim that R4 is derived. The register's carrier list is incomplete.

---

## 6. R1 assessment — semiclassical matter–geometry consistency

### 6.1 Claim-by-claim disposition

| Claim | Content | Independent verdict | Basis |
|---|---|---|---|
| C-R1-01 | JT action and field equations | **ACCEPTED, with the adopted-model caveat** | Action (2.1)–(2.3), \(R+2=0\), and (2.6a–c) independently verified (§3.1). The near-extremal restriction is not the one stated (F-01, F-13, F-17). |
| C-R1-02 | Anomaly and stress-tensor forms | **ACCEPTED** | AP(5.1)–(5.2) verified; trace anomaly and \(\phi_0\) shift verified; the \(R\)-form equality is correctly unused (OB-2). |
| C-R1-03 | Conservation/Ward structure and chiral flux | **ACCEPTED** | Ward identity independent of the anomaly; conformal-gauge cancellation verified (§3.2). |
| C-R1-04 | Mass/flux, Schwarzian, ADM, evaporation | **ACCEPTED for the formulas; CONDITIONAL for the derivation** | ADM sign and \(E_0>0\) verified; energy balance and \(E_1e^{-ku}\) verified; \(dM/dt\) line-by-line owed (OB-4), but the formula is correct. |
| Quench state | HH/eternal + right coupling + \(E_S\) shock | **CONDITIONAL** | The AEMM shock is correct, but its relation to AMMZ's pulse is not established (F-18). |
| Time coordinate | \(u\) vs \(t=f(u)\) | **DEFECTIVE** | The spec conflates them (F-05); the decay law is in \(u\), the entropy formulas in \(t\). |
| Cutoff | \(z=\epsilon f'(u)\) | **CORRECT in R1, INCONSISTENT in R4** | AEMM(2.12)–(2.14) verified; F-10 is the R4 mismatch. |

### 6.2 Assumption audit (ASM-01–ASM-06, mapped to the current register)

The dossier-era IDs ASM-01–ASM-06 were rewritten in `audit-corrections-v1.1.md` lines 104–111. The current register uses A-DEF/A-PHY/A-MATH/A-BC/A-SC/A-REN. The mapping and disposition:

| Old ID | Current register | Statement | Status after review |
|---|---|---|---|
| ASM-01 | A-IMP-1 + A-DEF-1 | JT action postulate | **Corrected and verified**; the action is AEMM(2.1)–(2.3), not the dossier's \(4\times\) form. |
| ASM-02 | A-BC-2 + baseline §7 | Hierarchy \(c\ll S_0\ll S^i_{\rm Bek}\) | **Incorrect/incompatible**: wrong condition (F-17) and contradicts AEMM's \(\phi_0\gg\phi\) (F-01). |
| ASM-03 | A-PHY-1 | No diffeomorphism anomaly; Ward identity | **Retained; verified** (§3.2). The anomaly and the Ward identity are independent inputs. |
| ASM-04 | A-MATH-2 + A-IMP-4/5 | QES/RT minimality and stationarity | **Narrowed to a single interval; imported, not derived** (F-11); stationarity is meaningful under smoothness (A-MATH-2). |
| ASM-05 | A-PHY-2 | Holographic CFT matter (large \(c\), strong coupling) | **Retained as a postulate**; the AMMZ near-extremal condition is misstated (F-17). |
| ASM-06 | A-PHY-4 + A-DEF-1 | Two-sided HH + quench state; shared state | **Rewritten for the two-sided model, but the R1/R4 state map is open** (F-18). |

### 6.3 R1 verdict

**CONDITIONAL.** The R1 equation set is mathematically sound and faithful to AEMM/AP. The conditions under which it is claimed to hold are not: the hierarchy is internally contradictory (F-01), the near-extremal condition is misstated (F-17), the time coordinate is conflated (F-05), and the quench state's relation to R4 is not established (F-18). The R1 algebra can be accepted as a reproduction of the sources; the R1 physical model cannot be accepted as stated.

---

## 7. R4 assessment — Page-curve reconstruction

### 7.1 Claim-by-claim disposition

| Claim | Content | Independent verdict | Basis |
|---|---|---|---|
| C-R4-01 | Generalized entropy \(S_{\rm gen}=\phi/4G_N+S_{\rm bulk}\) | **ACCEPTED as a formula; CONDITIONAL on the endpoint count and the Ren split** | AMMZ(10)/(11) verified. The area term is correct for a single endpoint; the endpoint count is inconsistent (F-04). The Ren split is inherited from AEMM(69)–(73). |
| C-R4-02 | No-island branch \(S_{\rm no}=(c/6)\ln[(\beta/\pi\epsilon)\sinh(\pi t/\beta)]\) and slope \(\pi cT/6\) | **Slope ACCEPTED; branch identification REJECTED** | The log formula is a correct fixed-\(\beta\) CFT result; its slope is correct. It is not AMMZ(14) (F-02), it omits \(\log g\) (F-19), and its cutoff frame is inconsistent with `link_geometry_factors` (F-10). |
| C-R4-03 | QES position AMMZ(12) and stationarity | **CONDITIONAL IMPORT** | AMMZ(12) is source-reported; the JT derivation is in AEMM §4 and is not reconstructed (F-11). The log factor versus AEMM(4.40) is unreconciled (F-21). |
| C-R4-04 | Island branch \(S_{\rm isl}=S_0+(S_i-S_0)e^{-\kappa t/2}\) | **CONDITIONAL: correct form under the one-endpoint postulate, but the \(2S_0\) price is inconsistent** | The form follows from AMMZ(13) and the one-endpoint area term. The package also uses the two-endpoint \(2S_0\) price (F-04), which changes the large-\(S_0\) branch. |
| C-R4-05 | Page condition and Page time \(t_P\approx6\beta S_i/\pi c\) | **REJECTED as derived; leading estimate only** | The exact intersections are \(0.703/\kappa\) and \(0.811/\kappa\), not \(1/\kappa\) (F-03); the F1 self-check fails under AEMM (F-01, F-03); `page_curve` is qualitative/conditional (F-16). |

### 7.2 Branch reconstruction F1–F5

`branch-reconstruction.md` defines five footprint items. The independent disposition:

| Footprint | Package statement | Independent disposition |
|---|---|---|
| F1 | \(t_Pk\sim O(1)+S_0\)-piece; slow evaporation assumed | **FAILS.** The correct identity is \(t_Pk=1+\phi_0/\phi_{\rm hor}\). Under island domination it is \(O(1)\) (not small); under AEMM it is \(\gg1\). The footprint is not a small correction. |
| F2 | Additive constant \(\propto\ln(\beta/\epsilon)\) enters at subleading order | **Correct in form**, but the constant is frame-dependent (F-10) and the exact intersection shifts at \(O(k\log)\). |
| F3 | \(S_0\ll S_i\) expansion | **Correct algebraically** for the island branch; the no-island branch it is compared to is the wrong function (F-02). |
| F4 | \(\kappa=k\) form match | **Conditional** on the time identification (F-05); not an exact identity. |
| F5 | Large-\(S_0\) branch: island disfavored | **Correct physics for the two-endpoint \(2S_0\) price**, but the package adopts one endpoint (F-04). The branch cannot be applied to the adopted configuration without re-deriving the price. |

### 7.3 The Page curve

What the package actually establishes, after independent reconstruction:

1. Under the island rule, if the island branch crosses the no-island branch from below, the entropy follows a rising-then-falling curve of the Page form. This is a structural consequence of min-selection, conditional on RT/HRT (BR-09).
2. The rising branch in the package is a fixed-\(\beta\) CFT log, not AMMZ(14); the falling branch is the one-endpoint \(S_0+(S_i-S_0)e^{-\kappa t/2}\).
3. The crossing is at \(t_P=0.703/\kappa\) for the log branch (leading, without the constant) or \(0.811/\kappa\) for the AMMZ(14) branch; the package's \(1/\kappa\) is a heuristic that is 23–42% high.
4. The F1 self-check does not pass. It cannot pass in the AEMM regime, where \(\phi_0\gg\phi_{\rm hor}\) makes \(t_Pk\gg1\).
5. The large-\(S_0\) branch uses a two-endpoint price that is inconsistent with the adopted one-endpoint island.

Therefore the Page curve is **not derived** as claimed. It is a conditional/structural expectation, valid in the island-domination regime, with a branch formula and a Page-time value that are not the source's.

### 7.4 R4 verdict

**BLOCKED.** The no-island branch is not AMMZ(14) (F-02), the Page time is not the exact intersection and fails the F1 self-check (F-03), the endpoint count and the \(S_0\) price are inconsistent (F-04), the curve is qualitative/conditional (F-16), and the QES/RT inputs are imports (F-11). The R4 conclusion is not established.

---

## 8. Interoperability assessment — C-INT-01 and the ten links

### 8.1 C-INT-01

The package asserts one model, one state, one geometry, and a shared stress tensor across R1 and R4. The independent finding is:
- **One action, one variational origin:** supported (BR-01; §3.1). The R1 and R4 equations come from the same JT action.
- **One hierarchy:** not supported. The two conditions do not overlap (F-01).
- **One time coordinate:** not supported. \(u\) and \(t=f(u)\) are conflated (F-05).
- **One state:** not established. AEMM's \(\delta\)-shock and AMMZ's finite-width pulse are different states (F-18).
- **One cutoff:** not supported. The \(f'\)-cutoff and the constant cutoff are mixed (F-10).
- **One shared stress tensor:** supported in form (the same \(T_{\mu\nu}\) enters both sides), conditional on the state map.

C-INT-01 is therefore **partially supported**: the variational origin is shared; the state, hierarchy, time, and cutoff are not.

### 8.2 Link-by-link disposition

| Link (spec line) | Content | Independent disposition |
|---|---|---|
| `link_action_eqs` (82) | Metric equations and \(R+2=0\) from one action | **Established.** Verified in §3.1. |
| `link_anomaly_stress` (83) | Anomaly + Ward → chiral flux | **Established.** Verified in §3.2. |
| `link_state_stress` (84) | Shock stress is the stress of the quench state | **Restatement** of the premises; the state map is open (F-09, F-18). |
| `link_state_entropy` (85) | 3D-RT computes the radiation entropy of the quench state | **Conditional import** (`?OB-10`); not derived (F-11). |
| `link_geometry_factors` (87) | Entropy formulas use the \(f'\)-cutoff | **False dependency / frame mismatch** (F-09, F-10). |
| `link_dilaton_area` (88) | One-endpoint area law summed per endpoint | **Restatement plus a citation mismatch** (F-04, F-09, F-20). |
| `link_matter_sgen` (89) | \(S_{\rm gen}\) = dilaton term + matter entropy | **Definitional restatement** (F-09). |
| `link_sgen_qes` (90) | QES position extremizes \(S_{\rm gen}\) | **Circular restatement** (F-09); import `?OB-9` (F-11). |
| `link_branches_rad` (91) | Page curve from branch selection | **Circular restatement** (F-09). |
| `link_flux_mass` (92) | Global balance connects decay to bath flux | **Circular restatement** (F-09); OB-4 closable. |

Of the ten claimed links, two are established, three are conditional imports/restatements, and five are circular/definitional or have a false dependency. The claim `interop_complete` (line 93) is false.

### 8.3 The \(\kappa=k\) identity

The identity is a form match, not an exact identity. AEMM's decay is in the physical boundary time \(u\); AMMZ's formulas use \(t\). If \(t_{\rm AMMZ}=u_{\rm AEMM}\), the identity holds. The package never states the identification (F-05). The identity also inherits the near-extremal regime: \(\kappa=k\) is compatible with both hierarchies only in the overlap where the dynamical entropy exceeds \(c\), which the package's \(S_0\ll S_{\rm Bek}\) does not guarantee in the AEMM regime.

### 8.4 Bridge register BR-01–BR-10

The bridge dispositions are in Appendix E. The summary:
- **BR-01** (action → field equations): established (verified).
- **BR-02** (anomaly → stress forms): established in form; convention pinned.
- **BR-03** (diffeomorphism invariance → Ward identity): established (standard).
- **BR-04** (gravitational/bath interface): conditional; OB-5 open.
- **BR-05** (state → matter entropy): conditional import; OB-10 open.
- **BR-06** (dilaton → gravitational entropy): correct for one endpoint; the two-endpoint price is inconsistent.
- **BR-07** (area + matter → \(S_{\rm gen}\)): conditional on the Ren split; the split is stated but scheme dependence remains.
- **BR-08** (\(S_{\rm gen}\) → QES): conditional import; OB-9 open.
- **BR-09** (candidates → \(S(R)\)): structural; conditional on RT/HRT and on the branch values.
- **BR-10** (R1 ↔ R4 integration): not established; F-01, F-05, F-09, F-10, F-18.

### 8.5 Interoperability verdict

**BLOCKED.** The model hierarchy is incompatible (F-01); the time and cutoff frames are conflated (F-05, F-10); the state is a \(\delta\)-shock versus a finite-width pulse (F-18); several Part C links are circular or restatements (F-09); and the \(\kappa=k\) identity is conditional on an unstated time identification. The variational-origin link and the anomaly–stress link survive; the rest do not establish interoperability.

---

## 9. Independent application of RC1–RC3

### 9.1 RC1 — Mathematical landscape

**What is correct.** The landscape entries correctly identify what each structure establishes and what it does not: differential geometry gives causal structure and \(R+2=0\), not dynamics; JT dilaton gravity gives the field equations, Schwarzian dynamics, and ADM energy, not matter entropy; QFT in curved spacetime gives the stress forms and Hawking flux, not fine-grained entropy; RT/HRT gives \(S(R)\) only under holographic assumptions; extremization plus minimization gives the Page transition structurally, not a numerical \(t_P\); the ODE gives \(f(u)\) and the evaporation law, not shock microphysics. The separation of "establishes" from "does not establish" is the right shape and is largely accurate.

**What is defective.** The JT entry says the regime is "semiclassical, \(k\ll1\), away from extremality." The adopted AEMM model is near-extremal, with \(\phi_0\gg\phi\). This is a MATERIAL defect (F-13): the entry describes a different regime from the one the package uses, and the near-extremal condition is precisely the one that conflicts with the island hierarchy.

**What is missing.** The landscape does not include the near-extremal throat as a separate entry, and it does not state that the island-domination condition \(S_0\ll S_{\rm Bek}\) is a *different* regime from AEMM's near-extremality. The coordinate entry "ODE/boundary dynamics: Establishes \(T(t)\), \(E(u)\)" also mixes the two time variables without comment (F-05).

**RC1 verdict:** Adequate in structure, materially incomplete on the adopted regime. Accept with F-13.

### 9.2 RC2 — Mathematical intersection and removal tests

**What is correct.**
- "Remove dilaton cost (\(G_N\to\infty\)): no finite extremum" is correct, with the stated caveat that the regime breaks.
- "Remove CFT (\(c\to0\)): extremum → classical horizon/singularity, no island" is correct.
- "Remove anomaly: flux loses Hawking terms but conservation still holds" is the key correction, and it is correct: the anomaly and the Ward identity are independent joint inputs (AP §5.1; §3.2 of this report).
- "Remove Ren split: regulator-dependent pieces" is correct: AEMM(69)–(73) is required.
- "Remove minimization: unbounded Hawking growth" is correct.
- The no-minimality disclaimer is appropriate.

**What is defective.** "Page curve = branches ∩ minimization ∩ hierarchy (\(S_0\ll S^i_{\rm Bek}\)). Violate \(S_0\ll S^i_{\rm Bek}\): \(2S_0\) price blocks transition" states the condition but does not test it against the adopted model. The condition \(S_0\ll S^i_{\rm Bek}\) is exactly what conflicts with AEMM's \(\phi_0\gg\phi\) (F-01). RC2 performs the removal test on every ingredient except the one that fails. It also uses the two-endpoint \(2S_0\) price without noticing the package's one-endpoint configuration (F-04).

**What is missing.** The R1/R4 integration entry says "shared state/geometry/cutoffs + per-equation compatibility." The package does not have shared state (F-18), shared time (F-05), or shared cutoff (F-10). RC2 names the failure mode but does not apply it to the package. The intersection is asserted, not demonstrated.

**RC2 verdict:** Conceptually strong and correct on the anomaly/Ward structure; materially incomplete on the hierarchy and integration tests. Accept the method; reject the claim that the tests were passed.

### 9.3 RC3 — Mathematical interoperability

**What is correct.** The distinction between identities (\(R+2=0\), Ward), transformations (Weyl/CFT maps), physical postulates (holographic matter, quench state, no diffeomorphism anomaly), imported theorems (AEMM/AMMZ/AP), approximations (semiclassical, slow evaporation, logs dropped), and open obligations is exactly the right taxonomy for a bridge register. The ten-field bridge format is appropriate.

**What is defective.** The closing claim "No two quantities equated by shared symbols (\(\phi_h/\kappa\) collisions resolved in baseline §4)" is false:
- \(t\) is used for both the Poincaré time and the boundary time (F-05).
- \(S_{\rm BH}\) and \(S_{\rm Bek}\) are used interchangeably (F-24).
- \(S_{\rm Bek}\) is called the "dynamical BH entropy" while the dynamical entropy is \(S_{\rm Bek}-S_0\) (F-17).
- \(\phi_h\) is renamed \(\phi_{\rm hor}\), but \(\phi_0\gg\phi\) and \(\phi_0\ll\phi_{\rm hor}\) both appear (F-01).
The "no shared symbols" claim is therefore contradicted by the package itself.

**What is missing.** RC3 says each bridge "records ... status," but several statuses are overclaimed: BR-01 is correct; BR-02 is conditional; BR-03 is source-reported; BR-04 is conditional; BR-05 is conditional; BR-06 is source-reported but inconsistent with the endpoint count; BR-07 is conditional; BR-08 is conditional; BR-09 is source-reported/structural; BR-10 is conditional and, as F-09 shows, not established. The bridge register does not carry the circularity of the Part C links.

**RC3 verdict:** The taxonomy is right; the "no shared symbols" and "interoperability" claims are not. Accept the format; reject the completion claim.

---

## 10. Phys-Pseudo conformance

The specification was judged against `PSEUDO.md` as the sole authority, not against Muse's checklist.

### 10.1 Construct and role conformance

| Rule | Requirement | Spec behavior | Verdict |
|---|---|---|---|
| §1 | Seven constructs only; no status or provenance construct | Uses the seven constructs, but appends `[SR]/[D]/[C]/[P]/[A]/[U]` status tags to statements | **Non-conforming** (F-08) |
| §2.1 | Declaration before use; every premise name introduced earlier | Many essential symbols are undeclared or introduced only in prose/baseline (F-07) | **Non-conforming** |
| §2.2 | Premise lists contain only exact names | Lists contain names, but some are definitions used as premises (`stressT`, `conformal_gauge`, `coords_physical`, `schwarzian`, `sgen_functional`, `dilaton_term`, `bh_entropy_fn`) | **Non-conforming in role** (definitions used as premises) |
| §2.3 | No empty premise lists | None empty | Conforming |
| §2.4 | One name per construct | No duplicate declarations found | Conforming |
| §2.5 | Branch scope | The `if large_S0_regime` and `if shock_regime` blocks are properly scoped; no branch name is referenced outside | Conforming |
| §2.6 | Essential quantities declared and in premises | Essential quantities are missing (F-07) | **Non-conforming** |
| §3 | Roles not interchangeable; definitions not used as premises | `given regime_hierarchy` is really an assumption; definitions are used as premises; `claim noisland_branch` is used as a premise for a `derive` (allowed, but the result rests on an unestablished claim and the text does not say so) | **Non-conforming in role** |
| §4 | `derive` needs name, explicit result, non-empty premise list, reasoning | The derive steps have names, results, and reasoning; several results restate their premises (F-09) | **Non-conforming in substance** |
| §5 | Conditionals explicit; branch statements scoped | `if shock_regime` uses a claim name as a condition rather than an explicit comparison; borderline but the condition is a declared proposition | **Minor** |
| §6 | `?` attaches to the statement one level above; content says what is missing and what resolves it | Most `?` attachments are correct; `?OB-6` is missing from the spec (F-15); `?OB-8followup` is untracked | **Partially non-conforming** |
| §7 | No author-typed status | The tags are an author-typed status taxonomy; the header claims "per PSEUDO.md" | **Non-conforming** (F-08) |
| §8 | Source fidelity in prose; no provenance operators/evidence taxonomy | The `[SR: ...]` attributions are acceptable as prose source notes, but the status taxonomy is prohibited; the summary uses status vocabulary | **Non-conforming** (F-08, F-14) |

### 10.2 What is conforming

- The document uses the seven constructs in the right general shape: `given`, `define`, `assume`, `claim`, `derive`, `if/then/else`, `?`.
- The `derive` steps have names, explicit results, premise lists, and reasoning sketches. The template is complete.
- The branch blocks are indented and scoped correctly.
- Most `?` items state what is missing and what would resolve it (for example `?OB-9` names the AEMM equations to be re-derived).
- The document is compact and readable as pseudocode; the form is close to conforming.

### 10.3 What must be repaired

1. Remove the status-tag taxonomy (F-08). Replace it with prose source notes and `?` items.
2. Add a symbol dictionary and declare essential quantities before use (F-07).
3. Move assumption-like `given` statements to `assume` (regime, hierarchy) or restate them as inputs with their source.
4. Do not use `define` names as premises; substitute the definitions or make the needed statements explicit.
5. Carry `?OB-6` in the spec and register `?OB-8followup` (F-15).
6. Replace the circular `derive` links with genuine derivations or restatements (F-09).
7. Correct the header claim "per PSEUDO.md."

### 10.4 Conformance verdict

**BLOCKED / NON-CONFORMING (repairable).** The spec is structurally close to PSEUDO — the seven constructs, the branch layout, and most `?` attachments are right — but it violates the central invariant (§7) with an author-typed status taxonomy, violates declaration-before-use (§2.1, §2.6), and contains circular `derive` steps (§4). The defects are repairable without changing the physics, except where the physics findings (F-01, F-04) require substantive revision.

---

## 11. Limiting cases and counterexamples

### 11.1 Hierarchy no-overlap

Set \(4G_N=1\), so \(S_0=\phi_0\) and \(S_{\rm Bek}=\phi_0+\phi_{\rm hor}\).
- **Case A:** \(\phi_0=100\), \(\phi_{\rm hor}=900\). Then \(S_0=100\), \(S_{\rm Bek}=1000\). Island domination holds (\(100\ll1000\)), but \(\phi_0\ll\phi_{\rm hor}\). AEMM's near-extremal restriction fails.
- **Case B:** \(\phi_0=1000\), \(\phi_{\rm hor}=10\). Then \(S_0=1000\), \(S_{\rm Bek}=1010\). AEMM's restriction holds (\(\phi_0\gg\phi_{\rm hor}\)), but \(S_0\ll S_{\rm Bek}\) fails (\(1000\not\ll1010\)).
- **Case C:** \(\phi_0=100\), \(\phi_{\rm hor}=100\). Then \(S_0=100\), \(S_{\rm Bek}=200\): \(S_0\) is half of \(S_{\rm Bek}\), not negligible; and \(\phi_0\) is not \(\gg\phi_{\rm hor}\).
No assignment satisfies both conditions. The package's model has no overlap regime. This is F-01.

### 11.2 Endpoint count: \(S_0\) versus \(2S_0\)

Take \(S_i=1\) (in units where \(S_0\) is measured) and \(S_0=0.6\).
- **One endpoint:** the island-creation price is \(S_0=0.6<1\), so the island can compete at late times.
- **Two endpoints:** the price is \(2S_0=1.2>1\), so the island is disfavored at all times.
The same \(S_0\) gives opposite branch conclusions depending on the endpoint count. The package adopts one endpoint (line 52) and uses the two-endpoint price (line 74). This is F-04, and it is not a small correction: it can remove the Page transition entirely.

### 11.3 Time map: \(T(u)\) versus \(T(f(u))\)

Take \(f(u)=\frac{1}{\pi T_0}\tanh(\pi T_0u)\). Then \(f'(u)=\operatorname{sech}^2(\pi T_0u)\), and as \(u\to\infty\), \(f(u)\to1/(\pi T_0)\) while \(f'(u)\to0\).
- AEMM in physical time: \(T(u)=T_1e^{-ku/2}\to0\).
- AMMZ in Poincaré time: \(T(f(u))=T_1e^{-kf(u)/2}\to T_1e^{-k/(2\pi T_0)}\), a nonzero constant.
The two decays are qualitatively different: one vanishes, the other saturates. The Page-time integral \(\int T\,dt\) gives different answers. The package writes \(T(t)=T_1e^{-kt/2}\) without stating which time is meant (F-05). This is a counterexample to the "form match" claim: the functional forms match only after the time identification is imposed.

### 11.4 Cutoff frame: constant \(\epsilon\) versus \(\epsilon f'(u)\)

Under a conformal map, the interval entropy shifts by
\[
\Delta S=\frac{c}{6}\ln\left|\frac{dy}{dx}\right|.
\]
For the boundary cutoff \(z=\epsilon f'(u)\) versus \(z=\epsilon\),
\[
\Delta S=\frac{c}{6}\ln f'(u)
=\frac{c}{6}\ln\operatorname{sech}^2(\pi T_0u)
\simeq-\frac{c\pi T_0}{3}u
\quad (u\to\infty).
\]
The shift grows linearly with \(u\) at late times. The package uses the constant-cutoff log in `noisland_branch` and the \(f'\)-cutoff in `link_geometry_factors` (F-10). The difference is not a constant; it changes the late-time slope of the no-island branch and therefore the intersection. This is a counterexample to the claim that the two frames are interchangeable.

### 11.5 Exact Page-time values

For the same \(S_i\gg S_0\):
- Fixed-\(\beta\) log branch: \(t_P=2W(1/2)/\kappa=0.70347/\kappa\).
- AMMZ(14) branch: \(t_P=(2/\kappa)\ln(3/2)=0.81093/\kappa\).
- Package estimate: \(t_P=1/\kappa\).
The package value is 42% above the first and 23% above the second. The spread is not a small correction at the level of a "derived" value. This is F-03.

### 11.6 F1 self-check

With \(t_P=6\beta S_i/(\pi c)\), \(S_i=(\phi_0+\phi_{\rm hor})/(4G_N)\), \(\beta=1/T\), and \(k=cG_N/(3\bar\phi_r)\),
\[
t_Pk=\frac{\phi_0+\phi_{\rm hor}}{2\pi T\bar\phi_r}
=1+\frac{\phi_0}{\phi_{\rm hor}}.
\]
Under \(\phi_0\gg\phi_{\rm hor}\), \(t_Pk\gg1\), so the slow-evaporation approximation used to derive \(t_P\) is invalid. The self-check is a diagnostic of F-01, not a confirmation.

### 11.7 Large-\(S_0\) branch

If \(S_0\) is comparable to \(S_{\rm Bek}\), the \(2S_0\) price blocks the island (AMMZ §3.3). The package's `large_S0_regime` branch captures this correctly. But the package's main branch assumes \(S_0\ll S_{\rm Bek}\), which is incompatible with AEMM. The two branches are therefore not two regimes of one model; they are two different models. This is F-01 and F-04.

### 11.8 The \(G_N\to0\) limit

As \(G_N\to0\) with fixed \(\phi\), the area term \(\phi/4G_N\) diverges and the island is disfavored at all times; the no-island branch wins and there is no Page transition. The package's limitations section records this (L1). The limit is a correct sanity check on the min-selection: it shows that the Page transition requires the gravitational area term to be comparable to the matter entropy, which is the island-domination regime. The package's main hierarchy is meant to be that regime; AEMM's is not.

---

## 12. Required corrections

The corrections are ordered by priority. A correction is "required" if the affected finding is MATERIAL or BLOCKER; minor corrections are listed separately.

### 12.1 Blocking corrections

**C1 — Resolve the hierarchy (F-01).** Choose one regime and state it once:
- (a) AEMM near-extremal, \(\phi_0\gg\phi_{\rm hor}\): then \(S_0\approx S_{\rm Bek}\), the island is not favored by the \(2S_0\) price, and the R4 Page transition does not occur in the adopted model. State this consequence.
- (b) Island domination, \(S_0\ll S_{\rm Bek}\): then AEMM's \(\phi_0\gg\phi\) must be dropped, the action's near-extremal interpretation must be restated, and the model is no longer AEMM as written.
- (c) A different model with a small extremal entropy and a large dynamical dilaton, with the action and near-extremal condition re-sourced.

Remove the contradictory line (spec line 12) and the adjacent contradictory lines (baseline 223–224). Do not present (a) and (b) as one model.

**C2 — Resolve the endpoint count and the \(S_0\) price (F-04).** State the island's actual endpoint count and use the corresponding area term and creation price:
- One endpoint: area \((\phi_0+\phi(a))/4G_N\), price \(S_0\); re-derive the large-\(S_0\) branch.
- Two endpoints: area \(2(\phi_0+\phi_{\rm hor})/4G_N\), price \(2S_0\); re-derive the island branch.
Remove the AMMZ(10) "summed over endpoints" citation; cite AMMZ(15) for the multi-endpoint area.

**C3 — Repair the interoperability links (F-09).** For each of `link_state_stress`, `link_geometry_factors`, `link_dilaton_area`, `link_matter_sgen`, `link_sgen_qes`, `link_branches_rad`, `link_flux_mass`:
- If the conclusion is already in the premises, reclassify it as a definition/restatement and remove it from the count of established links.
- If it is a genuine link, state the inference step and the premises that do not contain the conclusion.
Recount the established links; do not claim ten unless ten are established.

### 12.2 Material corrections

**C4 — Use the correct no-island branch or state the approximation (F-02).** Either use AMMZ(14) with its saturating form, or keep the fixed-\(\beta\) log and state its domain of validity, the omitted constant, and the error at the intersection. Remove the "AMMZ(14) cross-check" claim for the log.

**C5 — Report the exact Page time and the F1 failure (F-03, F-16).** Report \(0.703/\kappa\) and \(0.811/\kappa\) alongside the \(1/\kappa\) estimate; state that F1 fails under AEMM and is \(O(1)\) under island domination. Downgrade `page_curve` from `[D]` to conditional/structural.

**C6 — Separate the time coordinates (F-05).** Introduce \(u_{\rm bdy}\) and \(t_{\rm Poinc}\), state the identification \(t_{\rm AMMZ}=u_{\rm bdy}\), and rewrite every Part B formula in the chosen variable. Re-derive the rate identity and the cutoff transformation.

**C7 — Use one cutoff frame (F-10).** State the frame of every entropy formula; use \(z=\epsilon f'(u)\) consistently and include the \(-\frac{c}{6}\ln f'(u)\) transformation, or compute both branches in the constant-cutoff frame and state why the \(f'\) factor drops.

**C8 — Establish or limit the state map (F-18).** Derive the pulse-to-shock limit, or state that the interop is valid only for observables insensitive to the pulse width and quantify the difference. Carry `?OB-5` on the affected claims.

**C9 — Tag the imports consistently (F-11).** Every result depending on `?OB-9`/`?OB-10` is conditional/imported, not derived. Either reconstruct the QES/RT steps or state the dependency plainly. Remove `page_curve` from the "derived" list.

**C10 — Repair Phys-Pseudo conformance (F-07, F-08, F-15).**
- Remove the status taxonomy; use prose source notes and `?`.
- Add a symbol dictionary and declare essential quantities before use.
- Carry `?OB-6` in the spec; register `?OB-8followup`.
- Do not use definitions as premises.
- Correct the header claim "per PSEUDO.md."
- Replace circular `derive` links (C3).

**C11 — Correct the status claims (F-14, F-16).** Rewrite the summary to separate source-established, independently reconstructed, conditional/imported, and open results. Remove "MATHEMATICALLY ESTABLISHED" for the package's own reconstruction, "DERIVED" for the one-endpoint postulate/branch intersection, and "single-model, traceable" until F-01/F-05 are resolved.

**C12 — Correct the hierarchy statement and the factor-4 error (F-17).** Replace \(c\ll S_0\) with the AMMZ condition \(c\ll S_{\rm Bek}-S_0\); remove "equivalently \(c\ll\phi_0/G_N\)"; distinguish \(S_{\rm Bek}\) from \(S_{\rm dyn}\).

**C13 — Correct RC1's regime statement (F-13).** Distinguish general JT from the near-extremal AEMM regime; remove "away from extremality" as the model's regime.

**C14 — Correct route (i)'s scope (F-06).** State that the no-island branch is a 2D CFT computation and the island branch is the 3D RT import; do not claim "NO 2D-CFT entropy computation."

### 12.3 Minor corrections

- **F-12:** Document the AEMM group-ordinal convention or replace locators with printed numbers.
- **F-19:** State that \(\log g=0\) for the adopted Cardy brane, or include \(\log g\).
- **F-20:** Cite AMMZ(15) for the multi-endpoint area.
- **F-21:** Reconcile the AMMZ(12)/AEMM(4.40) \(\log16\) or mark it convention-dependent.
- **F-22:** Reconcile `ward_identity` status.
- **F-23:** Remove the backward edge from `noisland_branch`; redraw the graph; qualify the acyclicity claim.
- **F-24:** Define \(S_{\rm Bek}\), \(S_0\), and \(S_{\rm dyn}\) once and use them consistently.

---

## 13. Accepted claims

The following claims are accepted, with the stated scope:

1. **JT action and dilaton equation.** AEMM printed (2.1)–(2.3), \(R+2=0\). The normalization is \(\phi_0/16\pi G_N\), the boundary term is \(K\), and the sign of the \(g_{\mu\nu}\phi\) term is minus. Accepted as a faithful reproduction of AEMM. (C-R1-01, BR-01.)
2. **Metric equations.** AEMM printed (2.6a–c), including the \(16\pi/8\pi\) pattern. Accepted. (C-R1-01, BR-01.)
3. **Stress-tensor forms.** AP printed (5.1)–(5.2), with \(N\leftrightarrow c\) and \(\omega\leftrightarrow\rho\). Accepted. (C-R1-02, BR-02.)
4. **Trace anomaly and \(\phi_0\) shift.** \(\langle T^\mu_{\ \mu}\rangle=cR/24\pi\) and \(\phi_0\to\phi_0+cG_N/3\). Accepted. (C-R1-02.)
5. **Ward identity and chiral flux.** The Ward identity is independent of the anomaly; the light-cone conservation and the cancellation leaving \(\partial_-\langle :T_{++}:\rangle=0\) are correct. Accepted. (C-R1-03, BR-03.)
6. **Schwarzian action and ADM energy.** \(S_G=(\bar\phi_r/8\pi G_N)\int\{f,u\}\), \(E(u)=-(\bar\phi_r/8\pi G_N)\{f,u\}\), and \(E_0=\pi\bar\phi_rT_0^2/4G_N>0\). Accepted. (C-R1-04.)
7. **Energy balance and evaporation.** \(\partial_uE=f'^2(T_{--}-T_{++})\) and \(E=E_1e^{-ku}\), with \(k=(c/12)(4G_N/\bar\phi_r)\). Accepted; the line-by-line \(dM/dt\) remains owed (OB-4) but the formula is correct. (C-R1-04.)
8. **No-island slope.** \(\pi cT/6\) and the additive constant \((c/6)\ln(\beta/2\pi\epsilon)\) for the fixed-\(\beta\) log. Accepted as a fixed-\(\beta\) CFT result. (C-R4-02, partial.)
9. **Island branch form.** \(S_0+(S_i-S_0)e^{-\kappa t/2}\) given the one-endpoint postulate. Accepted as a conditional form. (C-R4-04, partial.)
10. **Extremization/minimization ordering.** QES stationarity then branch minimization. Accepted as structure. (C-R4-05, partial.)
11. **Island rule status.** The island rule is an effective prescription derived from RT/HRT only within the holographic setup. Accepted. (C-R4-05, partial.)
12. **Large-\(S_0\) physics.** A \(2S_0\) price blocks the island for a two-endpoint island when \(S_0\) is comparable to \(S_{\rm Bek}\). Accepted for the two-endpoint configuration. (C-R4-04, partial.)
13. **AMMZ formalism compatibility.** AMMZ's formalism "can be directly applied" to the AEMM case, with the matter changed to a holographic CFT. Accepted as a source statement; it does not resolve the hierarchy conflict. (C-INT-01, partial.)
14. **\(G_N\to0\) limit.** The area term diverges and the island is disfavored; no Page transition. Accepted as a sanity check. (RC2/L1.)

---

## 14. Rejected or narrowed claims

**Rejected:**
1. **"The combined model is one coherent regime."** The two hierarchies do not overlap (F-01).
2. **"The no-island branch is AMMZ(14)."** It is a fixed-\(\beta\) log; AMMZ(14) is a saturating integral (F-02).
3. **"\(t_P=6\beta S_i/\pi c\) is derived."** It is a leading estimate; exact intersections are \(0.703/\kappa\) and \(0.811/\kappa\) (F-03).
4. **"The one-endpoint island and the \(2S_0\) price are compatible."** They are not (F-04).
5. **"The ten interoperability links are established."** At most three are; several are circular/restatements (F-09).
6. **"The spec conforms to PSEUDO.md."** It does not (F-07, F-08, F-15).
7. **"MATHEMATICALLY ESTABLISHED / DERIVED / single-model, traceable"** as summary statuses. They overstate the content (F-14, F-16).
8. **"\(\kappa=k\) is a term-by-term identity."** It is a form match conditional on the time identification (F-05).
9. **"Route (i) uses no 2D-CFT entropy computation."** It does (F-06).
10. **"RC3's no-shared-symbols claim."** \(t\), \(S_{\rm BH}/S_{\rm Bek}\), and \(\phi_0/\phi_{\rm hor}\) are shared or confused (F-05, F-17, F-24).

**Narrowed:**
1. **C-R4-01** (\(S_{\rm gen}\)): accepted as a formula; narrowed by the endpoint count and the Ren split.
2. **C-R4-02** (no-island): slope accepted; branch identification narrowed to a fixed-\(\beta\) approximation.
3. **C-R4-03** (QES): narrowed to a source-reported import.
4. **C-R4-04** (island): narrowed to the one-endpoint form, conditional on the \(S_0\) treatment.
5. **C-R4-05** (Page): narrowed to a structural/conditional expectation.
6. **C-INT-01** (shared model): narrowed to the shared variational origin; the state/time/cutoff/hierarchy are not shared.
7. **BR-01–BR-10**: narrowed per Appendix E; BR-01 and BR-03 stand; BR-02, BR-04–BR-10 are conditional or source-reported.

---

## 15. Final verdict

### 15.1 Overall verdict

**BLOCKED.**

The package contains a correct R1 equation set built on a model that is not internally consistent. The R4 Page-curve reconstruction is not derived: its no-island branch is the wrong function, its Page time is a heuristic that fails the package's own F1 self-check, and its endpoint normalization is inconsistent with the \(2S_0\) price it uses. The interoperability claim is not established: several links are circular, the time and cutoff frames are conflated, and the state map is open. The specification is non-conforming to Phys-Pseudo, and the summary overclaims the epistemic status.

The package is repairable. Section 12 gives the minimum corrections. The repairs require substantive physics decisions (C1, C2), not merely editorial changes.

### 15.2 Separate verdicts

| Area | Verdict |
|---|---|
| R1 mathematical validity | **CONDITIONAL** — algebra verified; regime and time coordinate defective. |
| R4 mathematical validity | **BLOCKED** — branch, Page time, and endpoint count defective. |
| R1–R4 interoperability | **BLOCKED** — hierarchy, time, cutoff, state, and links. |
| Phys-Pseudo conformance | **BLOCKED / NON-CONFORMING (repairable)** — status taxonomy, declaration-before-use, circular derives. |
| Source fidelity and epistemic integrity | **BLOCKED (repairable)** — equation-level fidelity good; metadata/status claims not. |
| Overall readiness | **BLOCKED** — at least one unresolved BLOCKER. |

### 15.3 Conditions for a re-review

A targeted revision pass requires separate director authorization (the one-pass governance rule). If authorized, the re-review should check at minimum:
1. The chosen hierarchy (C1) and its consequences for the island branch.
2. The endpoint count and the \(S_0\) price (C2).
3. The no-island branch and the exact Page time (C4, C5).
4. The time-coordinate and cutoff maps (C6, C7).
5. The state map (C8).
6. The reclassified links (C3).
7. PSEUDO conformance (C10).
8. The corrected summary statuses (C11).

### 15.4 Limitations of this review

- This was one full pass, not a line-by-line re-derivation of every equation in AEMM §4 or AMMZ §3; the QES/RT imports were assessed as imports, not reconstructed. That is the same limitation the package carries (`?OB-9`, `?OB-10`).
- The review checked the load-bearing locators against the pinned PDFs and the rendered pages; it did not re-check every citation in the package.
- The review did not assess the `G(\eta)` fallback beyond noting that it is documented and not activated.
- The review did not test any implementation, because the assignment forbids implementation content.
- The Page-time numerical values are leading-order; the \(O(k\log)\) corrections were not computed beyond the additive constant.
- The review does not claim novelty for any result; all accepted results are source results or standard algebra.

### 15.5 Closing statement

The R1 mathematics is largely right. The R4 physics is not yet right. The package's own audit corrected many dossier errors, but it did not catch the hierarchy contradiction, the endpoint/price inconsistency, or the circular interoperability links. Those are the three blockers. Until they are resolved, the package cannot be approved, frozen, or implemented.

---

## Appendix A — Source and package hash manifest

**Start and end check.** All pre-existing package files and pinned sources were hashed at the start of the review (`work/package-hashes-start.txt`) and re-hashed at the end with `sha256sum -c`. All 17 entries returned `OK`. No pre-existing file was modified. The only additions are this report under `reviews/` and the scratch files under `jt-island-001-review/work/`.

| # | File (relative to package root) | SHA-256 | End check |
|---|---|---|---|
| 1 | `audit-corrections-v1.1.md` | `fd719be24e936165f9332dca9a3fa46e7873af659d78affb1199ab92e7a5e86b` | OK |
| 2 | `branch-reconstruction.md` | `6b7ac65210c30ccc3dd4a46cc30d4c07c143cf6f29289fb76b8a9f08e9747d10` | OK |
| 3 | `dossier-audit-note-v1.0.md` | `c18c6002a3fe54a9fb1b7fdd0ff83750a9238db57801d3008bf76ba83503b24c` | OK |
| 4 | `index.md` | `8e9eb66f9558f98017493f2be2d651fcf348f0549460ea857f4bdeb7cd594af6` | OK |
| 5 | `opencode-checklist.md` | `bc56d7aa111e2a288671659cc7e612a45b19cf3210c9cb87efeb38c9340dee87` | OK |
| 6 | `phys-pseudo-r1r4-v0.1.md` | `34020bba8cef54dc8e50e0599a8630aa2420fc2c3fad99ebf050e02fdbeb8b9a` | OK |
| 7 | `reality-checks.md` | `56ee4521be7fac73a2b57458336a1c6060ce2183e41be2acf674f33afd908c36` | OK |
| 8 | `registers.md` | `f4eb867a786e692da89c6a032d511064c0a77a422beffebeda8f071d5ddd2d79` | OK |
| 9 | `summary.md` | `cd2e930de6ac54c2dca0effaf54a57573fceb40731c9ff03afbb39bb1e95e118` | OK |
| 10 | `working-baseline.md` | `2006554c31076166d6eb61f9ad6be36f7551810df1917380daee416363d01364` | OK |
| 11 | `sources/AEMM-1905.08762v3.pdf` | `f51a59cbf6d660471965c6eca0920b70eb598652af366e65b60050edeabc3575` | OK |
| 12 | `sources/AMMZ-1908.10996v2.pdf` | `f68be764aa6e21b79d3095401a9434c8e02047d5e6bc327a3c7b370afca9a905` | OK |
| 13 | `sources/AP-1402.6334.pdf` | `37e78a45c5e573098b9c1457597e04cd36c336eabed8e4d389923884ec1aabb4` | OK |
| 14 | `sources/Penington-1905.08255.pdf` | `fa7cf02db1814f4b55e72f33f2e817c5001645b5c4f1b4ea77444b38ee03738c` | OK |
| 15 | `sources/AEMM-1905.08762v3.html` | `acdf1d10afb74c0a455ad79ea452be59bc6790159d38128c34fbe6e002f199ff` | OK |
| 16 | `sources/AMMZ-1908.10996v2.html` | `7d0a34485d3d3954e02d00ef96a13284de30821497632b91dc65015552f5f7e8` | OK |
| 17 | `sources/AP-1402.6334.html` | `1fd84c04a937bbaac2d92187fd175443808cc81f9ddf725ab3cf72013f779747` | OK |

**Source metadata.** AMMZ: 22 PDF pages, arXiv:1908.10996v2 [hep-th], 4 Nov 2019. AEMM: 55 PDF pages, arXiv:1905.08762v3 [hep-th], 4 Nov 2019. AP: 23 PDF pages, arXiv:1402.6334v3 [hep-th], 12 Aug 2015. Penington: 75 PDF pages, arXiv:1905.08255v3 [hep-th], 29 Aug 2020.

---

## Appendix B — Exact locator map

### B.1 AMMZ printed equations (single sequence; PDF page numbers)

| Eq. | Content | PDF p. |
|---|---|---|
| (1) | \(I_{\rm grav}=\int d^2y\sqrt{-g}\left(\phi R/16\pi G_N^{(2)}+U(\phi)\right)\) | 4 |
| (2) | \(I=I_{\rm grav}+I_{\rm CFT}\) | 4 |
| (3) | \(g^{(3)}_{ij}|_{\rm bdy}=\epsilon^{-2}g^{(2)}_{ij}\) | 5 |
| (4) | \(ds^2=-e^{2\rho}dy^+dy^-\), chiral stress components | 5 |
| (5) | \(T_{y^\pm y^\pm}=-(c/24\pi)\{w^\pm,y^\pm\}\) | 5 |
| (6) | \(ds^2=-dw^+dw^-\), \(T_{ww}=0\) | 6 |
| (7) | AdS\(_3\) Poincaré metric | 6 |
| (8) | metric matching \( -dw^+dw^-/z_w^2=-\epsilon^{-2}e^{2\rho}dy^+dy^-\) | 6 |
| (9) | \(z_w=\epsilon e^{-\rho}\sqrt{(dw^+/dy^+)(dw^-/dy^-)}\) | 6 |
| (10) | \(S_{\rm gen}(y)=\phi(y)/4G_N^{(2)}+S_{\rm bulk\text{-}2d}[I_y]\) | 7–8 |
| (11) | \(S_{\rm gen}\approx\phi/4G_N+{\rm Area}^{(3)}[\Sigma_y]/4G_N^{(3)}\) | 8 |
| (12) | \(y_e^+=t-\frac{1}{2\pi T(t)}\log\frac{S_{\rm Bek}(T(t))-S_0}{c}+\cdots\) | 9 |
| (13) | \(S_{\rm BH}(t)=S_{\rm Bek}(T(t))+{\rm logs}\) | 9 |
| (14) | \(S_{\rm BH}(t)\sim S_{\rm Rad}(t)=\frac{\pi c}{6}\int_0^t dt'\,T(t')=2S^i_{\rm Bek}(1-e^{-\frac{\kappa}{2}t})\) | 13 |
| (15) | \(S(A)={\rm Min}_{I_g}{\rm Ext}_{I_g}\left[S_{\rm eff}(A\cup I_g)+{\rm Area}[\partial I_g]/4G_N\right]\) | 15 |

Key passages: AMMZ §2 near-extremal condition \(c\ll(\phi-\phi_0)/(4G_N^{(2)})\) on PDF p.4; AMMZ §2.1 compatibility statement ("This formalism can be directly applied to the case considered in [9] ... the only difference being that we consider matter composed of a holographic CFT\(_2\)") on PDF p.7; AMMZ §3.3 two-endpoint/\(2S_0\) passage and Fig. 11 on PDF p.14; island rule (15) on PDF p.15.

### B.2 AEMM printed equations and the package's group-ordinal convention

The AEMM PDF uses sectional numbers \((2.1)\)–\((5.10)\). The package's AEMM(\(n\)) is the ordinal of the \(n\)-th numbered equation group, counting subequations as one group. The mapping for the load-bearing equations:

| Package AEMM(\(n\)) | Printed AEMM | Content | PDF p. |
|---|---|---|---|
| (6)–(8) | (2.1)–(2.3) | Action \(I_0\), \(I_G\), \(I_M\) | 8 |
| (9) | (2.4) | \(R+2=0\) | 8 |
| (10) | (2.5) | Poincaré metric / coordinates | 8–9 |
| (11a–c) | (2.6a–c) | Metric equations, \(16\pi/8\pi\) | 9 |
| (12) | (2.7) | Boundary condition \(g_{uu}=1/\epsilon^2\), \(\phi_b=\bar\phi_r/\epsilon\) | 9 |
| (13)–(14) | (2.8)–(2.9) | Schwarzian action, ADM energy | 9–10 |
| (20) | (2.15) | \(E_0=\pi\bar\phi_rT_0^2/4G_N\) | 11 |
| (21)–(25) | (2.16)–(2.20) | Shock stress, energy balance, decay, \(k\), \(E(u)\) | 11–12 |
| (42) | (3.4) | Stress transformation / Schwarzian | 17 |
| (44)–(46) | (3.6)–(3.8) | Quench map and shock | 17–18 |
| (53) | (3.15) | Weyl/cutoff relation | 21 |
| (56)–(60) | (3.18)–(3.22) | Cross-ratio and \(G(\eta)\) | 22–24 |
| (61)–(68) | (3.23)–(3.30) | AdS entropy formulas | 24–27 |
| (69)–(73) | (4.1)–(4.5) | \(S_{\rm gen}\) renormalization | 29–30 |
| (75)–(77) | (4.7)–(4.9) | QES equations | 30–31 |
| (92)–(93) | (4.24)–(4.25) | QES equations | 33 |
| (104)–(107) | (4.36)–(4.39) | QES/scrambling chain | 35 |
| (108) | (4.40) | \(t_{HP}=\frac{\beta}{2\pi}\log[\frac{16}{c}(S-S_0)]\) | 35 |
| (112) | (4.44) | Bath entropy \(S_{\rm bath}\sim\frac{\bar\phi}{4G_N}4\pi T_1(1-e^{-\frac{k}{2}u})\) | 37 |

Additional AEMM locators: near-extremal restriction \(\phi_0\gg\phi\) in the text after (2.3), PDF p.8 (printed p.7); footnote 2 on the anomaly shift, PDF p.11; §4 QES derivation, PDF pp.30–37.

### B.3 AP and Penington

| Source | Locator | Content | PDF p. |
|---|---|---|---|
| AP | (5.1) | \(\langle T_{+-}\rangle=-\frac{c}{12\pi}\partial_+\partial_-\rho\) | 15 |
| AP | (5.2) | \(\langle T_{\pm\pm}\rangle=\frac{c}{12\pi}(\partial_\pm^2\rho-(\partial_\pm\rho)^2)+\langle :T_{\pm\pm}:\rangle\) | 15 |
| Penington | Eq. (2) | \(S=\min(S_{\rm rad},A_{\rm hor}/4G_N)\) | 5 |
| Penington | text after Eq. (2) | Page time \(S_{\rm rad}=A_{\rm hor}/4G_N\) | 6 |

---

## Appendix C — Symbol dictionary

| Symbol | Meaning | Defining statement |
|---|---|---|
| \(G_N\) | 2D Newton constant | AEMM(2.1) |
| \(\phi\) | dynamical dilaton | AEMM(2.2) |
| \(\phi_0\) | extremal/topological dilaton | AEMM(2.1) |
| \(\phi_b\) | boundary dilaton, \(\bar\phi_r/\epsilon\) | AEMM(2.7) |
| \(\bar\phi_r\) | boundary dilaton scale | AEMM(2.7) |
| \(\phi_{\rm hor}\) | dynamical horizon dilaton, \(2\pi\bar\phi_rT\) | AMMZ(13); AEMM §2 |
| \(S_0\) | extremal entropy, \(\phi_0/4G_N\) | AMMZ(12)–(13) |
| \(S_{\rm Bek}\) | Bekenstein–Hawking entropy, \((\phi_0+\phi_{\rm hor})/4G_N\) | AMMZ(12)–(13) |
| \(S_{\rm dyn}\) | dynamical entropy, \(S_{\rm Bek}-S_0=\phi_{\rm hor}/4G_N\) | derived |
| \(S_{\rm BH}\) | package symbol for \(S_{\rm Bek}\) (should be defined once) | `bh_entropy_fn` |
| \(S_i\) | initial \(S_{\rm Bek}(T_i)\) | branch reconstruction |
| \(S_{\rm no}\) | no-island radiation entropy | `noisland_branch` |
| \(S_{\rm isl}\) | island branch entropy | `island_value` |
| \(S_{\rm matter}\) | matter entropy on \(R\cup I\) | `sgen_functional` |
| \(S_{\rm eff}\) | effective entropy in the island rule | AMMZ(15) |
| \(T\) | black-hole temperature | AEMM §2 |
| \(T_0\) | initial temperature | AEMM(2.15) |
| \(T_1\) | post-quench temperature | AEMM(2.20) |
| \(T_H\) | package symbol for \(T\) (should be defined once) | `noisland_slope` |
| \(\beta\) | inverse temperature, \(1/T\) | standard |
| \(\epsilon\) | UV cutoff | AEMM(2.7) |
| \(c\) | central charge | AMMZ §2 |
| \(k\) | AEMM decay rate, \(\frac{c}{12}\frac{4G_N}{\bar\phi_r}\) | AEMM(2.20) |
| \(\kappa\) | AMMZ decay rate | AMMZ(14) |
| \(E_0\) | initial ADM energy, \(\pi\bar\phi_rT_0^2/4G_N\) | AEMM(2.15) |
| \(E_1\) | post-quench energy, \(E_0+E_S\) | AEMM(2.20) |
| \(E_S\) | shock energy | AEMM(2.16) |
| \(f(u)\) | boundary trajectory, \(t=f(u)\) | AEMM(2.5) |
| \(u\) | physical boundary time | AEMM §2 |
| \(t\) | Poincaré time (and AMMZ boundary time; must be disambiguated) | AEMM(2.5); AMMZ(12) |
| \(z\) | Poincaré radial coordinate | AEMM(2.5) |
| \(a\) | island endpoint | `island_region` |
| \(y_e\) | QES position | AMMZ(12) |
| \(t_P\) | Page time | package |
| \(R_t\) | radiation region (bath outside \(\sigma_0\)) | AMMZ §3.1 |
| \(I_t\) | island region \((-\infty,a(t)]\) | Decision 2 |
| \(g\) | boundary entropy (omitted; \(\log g=0\) for Cardy brane) | standard |

---

## Appendix D — Equation ledger

| Package claim | Printed source | Independent status |
|---|---|---|
| `jt_action` \(I_0,I_G,I_M\) | AEMM(2.1)–(2.3) | Verified exact; normalization \(\phi_0/16\pi G_N\), \(K\) term. |
| `curvature_constraint` \(R+2=0\) | AEMM(2.4) | Verified. |
| `metric_eqs` (11a–c) | AEMM(2.6a–c) | Verified; minus \(g\phi\), \(16\pi/8\pi\) pattern. |
| `coords_poincare` | AEMM(2.5) | Verified. |
| `coords_physical` \(t=f(u)\), \(z=\epsilon f'(u)\) | AEMM(2.7), §2.1 | Verified for R1; mixed in R4 (F-05, F-10). |
| `stress_forms` | AP(5.1)–(5.2) | Verified with \(N\leftrightarrow c\), \(\omega\leftrightarrow\rho\). |
| `trace_anomaly` | AEMM fn.2 | Verified; \(\phi_0\to\phi_0+cG_N/3\). |
| `ward_identity` | AP §5.1 (imposed) | Verified; independent of anomaly. |
| `lightcone_conservation` | Derived | Verified by direct computation. |
| `chiral_flux` | Derived | Verified; geometric pieces cancel. |
| `schwarzian_dynamics` | AEMM(2.8)–(2.9) | Verified; minus sign; \(E_0>0\). |
| `energy_balance` | AEMM(2.17) | Verified. |
| `evap_law` | AEMM(2.18)–(2.20), (2.24)–(2.25) | Verified; \(k\) formula. |
| `shock_stress` | AEMM(2.16) | Verified. |
| `noisland_branch` | Fixed-\(\beta\) CFT log (not AMMZ(14)) | Correct as a log; wrong branch identification (F-02); \(\log g\) omitted (F-19). |
| `noisland_slope` | Late-time expansion | Verified slope \(\pi cT/6\); constant frame-dependent (F-10). |
| `qes_position` | AMMZ(12); AEMM §4 | Import; factor 16 vs AEMM(4.40) (F-21). |
| `island_value` | AMMZ(13) | Form conditional on one endpoint; \(2S_0\) price inconsistent (F-04). |
| `branch_selection` | AMMZ §3.3 | Verified structure. |
| `page_condition` | Branch intersection | Verified equation; inputs defective. |
| `pagetime_value` | Heuristic | Leading estimate; exact \(0.703/\kappa\) or \(0.811/\kappa\) (F-03). |
| `page_curve` | AMMZ §3.3 | Qualitative/conditional; not derived (F-16). |
| `island_rule_status` | AMMZ(15) | Verified as conditional prescription. |
| `global_balance` | Standard 2D flux | Formula correct; derivation owed (OB-4). |
| `k`/`kappa` identity | AEMM(2.20); AMMZ(14) | Form match; time identification unstated (F-05). |

---

## Appendix E — Bridge matrix BR-01–BR-10

| ID | Bridge (source → target) | Package status | Independent disposition | Finding |
|---|---|---|---|---|
| BR-01 | JT action → field equations (variation) | DERIVED (rebuilt) | **Established.** Action, dilaton equation, metric equations verified. Boundary terms are bulk-only in the bridge record; no defect for the equations used. | None |
| BR-02 | Anomaly → stress forms (AP/AEMM transformation) | CONDITIONAL | **Established in form.** AP(5.1)–(5.2) verified; the \(R\)-form equality is dropped (OB-2). | None (OB-2 closed as dropped) |
| BR-03 | Diffeomorphism invariance → Ward identity (Noether) | SOURCE-REPORTED | **Established.** Standard identity; independent of the anomaly. | F-22 (status bookkeeping) |
| BR-04 | Gravitational/bath interface (quench, \(t=f(u)\), shock) | CONDITIONAL | **Conditional.** The shock is verified; the pulse-to-shock map is not (OB-5). The \(t=f(u)\) map is stated but not used consistently. | F-05, F-18 |
| BR-05 | State → matter entropy (3D RT, route i) | CONDITIONAL | **Conditional import.** The 3D-RT computation is not reconstructed (OB-10). Route scope overstated. | F-06, F-11 |
| BR-06 | Dilaton → gravitational entropy (area law/endpoint) | SOURCE-REPORTED | **Correct for one endpoint; inconsistent with the \(2S_0\) price.** AMMZ(10) is a single term; the multi-endpoint area is AMMZ(15). | F-04, F-20 |
| BR-07 | Area + matter → \(S_{\rm gen}\) (additivity + Ren split) | CONDITIONAL | **Conditional.** The Ren split is stated (AEMM(69)–(73)); scheme dependence remains, and the endpoint count changes the area term. | F-04, F-10 |
| BR-08 | \(S_{\rm gen}\) → QES (stationarity) | CONDITIONAL | **Conditional import.** The stationarity condition is correct in form; the QES position is imported (OB-9). | F-11, F-21 |
| BR-09 | Candidates → \(S(R)\) (min-selection) | SOURCE-REPORTED | **Structural.** The min-selection is correct; the branch values are wrong (log vs AMMZ(14); one vs two endpoints). | F-02, F-04, F-16 |
| BR-10 | R1 ↔ R4 integration (Part C links) | CONDITIONAL | **Not established.** Of ten links, two stand, three are conditional/restatements, five are circular/definitional or have a false dependency. | F-01, F-05, F-09, F-10, F-18 |

---

## Appendix F — Obligation matrix OB-1–OB-10

| ID | Load-bearing | Closable from pinned sources? | Independent status | Action |
|---|---|---|---|---|
| OB-1 | No | Yes | AEMM action uses plain \(K\); the \(K-1\) variant is not adopted. | Close. |
| OB-2 | No | Yes | Suspect \(R\)-form equality is dropped. | Close as dropped. |
| OB-3 | No | Yes | Anomaly route is sufficient; Polyakov-explicit leg not needed. | Close as not needed. |
| OB-4 | Yes (evaporation) | Yes | Formula verified; line-by-line derivation owed. | Add three-line derivation; close. |
| OB-5 | Yes (state) | No (requires derivation) | Shock vs pulse map not established. | Derive or limit; carry `?`. |
| OB-6 | Yes (normalization) | Partially | \(S_0\) dropped from no-island branch; no `?` in spec. | State and carry `?`. |
| OB-7 | Yes (central value) | No (requires resolution of F-01) | F1 self-check fails; exact intersections differ. | Report exact values; do not freeze. |
| OB-8 | No (closure scope) | Yes (restate) | Route (i) scope overstated; `?OB-8followup` untracked. | Restate; register follow-up. |
| OB-9 | Yes (QES) | Yes in principle (AEMM §4) | Import not reconstructed; factor 16 unreconciled. | Reconstruct or tag conditional. |
| OB-10 | Yes (island) | Yes in principle (AMMZ §3) | 3D-RT import not reconstructed. | Reconstruct or tag conditional. |

**Carrier completeness.** The register's carrier list (OB-4/5/9/10) omits OB-6, OB-7, and the F-09 circular-link repair. `interop_complete` is not carried by the listed obligations alone.

---

## Appendix G — Finding index

| ID | Severity | Area | One-line summary | Correction |
|---|---|---|---|---|
| F-01 | BLOCKER | R1/R4/interop | \(\phi_0\gg\phi\) and \(S_0\ll S_{\rm Bek}\) have no overlap. | C1 |
| F-02 | MATERIAL | R4 | No-island branch is a fixed-\(\beta\) log, not AMMZ(14); only the slope matches. | C4 |
| F-03 | MATERIAL | R4 | \(t_P=1/\kappa\) is not the exact intersection; F1 self-check fails under AEMM. | C5 |
| F-04 | BLOCKER | R4 | One-endpoint island with a two-endpoint \(2S_0\) price. | C2 |
| F-05 | MATERIAL | R1/R4/interop | AEMM \(u\) and Poincaré \(t=f(u)\) are conflated. | C6 |
| F-06 | MATERIAL | R4/interop | Route (i) claims no 2D-CFT computation while using one. | C14 |
| F-07 | MATERIAL | Conformance | Essential quantities undeclared; spec not self-contained. | C10 |
| F-08 | MATERIAL | Conformance/epistemic | Status-tag taxonomy violates PSEUDO §7; false header claim. | C10 |
| F-09 | BLOCKER | Interop | Several Part C links are circular or restatements; "ten links" not established. | C3 |
| F-10 | MATERIAL | R4/interop | Constant cutoff vs \(z=\epsilon f'(u)\); frame mismatch. | C7 |
| F-11 | MATERIAL | R4/epistemic | QES/RT imports presented as derived. | C9 |
| F-12 | MINOR | Traceability | AEMM group-ordinal locator convention undocumented. | §12.3 |
| F-13 | MATERIAL | R1/conformance | RC1 says "away from extremality"; model is near-extremal. | C13 |
| F-14 | MATERIAL | Epistemic | Summary overclaims "ESTABLISHED/DERIVED/single-model." | C11 |
| F-15 | MATERIAL | Conformance | `?OB-6` missing from spec; `?OB-8followup` unregistered. | C10 |
| F-16 | MATERIAL | R4/epistemic | `page_curve`/`pagetime_value` statuses overstate. | C5, C11 |
| F-17 | MATERIAL | R1/R4/source | Hierarchy \(c\ll S_0\) is not AMMZ's condition; factor-4 "equivalently." | C12 |
| F-18 | MATERIAL | R1/R4/interop | \(\delta\)-shock vs finite-width pulse; state map not established. | C8 |
| F-19 | MINOR | R4 | No-island formula omits \(\log g\) without stating Cardy-brane condition. | §12.3 |
| F-20 | MINOR | Source fidelity | AMMZ(10) cited for a multi-endpoint sum; correct citation is AMMZ(15). | §12.3 |
| F-21 | MINOR | Source fidelity | AMMZ(12)/AEMM(4.40) \(\log16\) unreconciled. | §12.3 |
| F-22 | MINOR | Conformance | `ward_identity` status inconsistent between register and spec. | §12.3 |
| F-23 | MINOR | Conformance/interop | Dependency graph routes island objects into `noisland_branch`; acyclicity overstated. | §12.3 |
| F-24 | MINOR | R4/source | \(S_{\rm BH}\) and \(S_{\rm Bek}\) used interchangeably. | §12.3 |

**Severity count:** 3 BLOCKER, 14 MATERIAL, 7 MINOR.

---

## Completion record

- **Review pass:** one full adversarial pass, per the approved plan and the one-pass governance rule.
- **Report status:** complete; all 15 required sections and Appendices A–G are present.
- **Pre-existing files:** all 17 package/source hashes match the start manifest (`sha256sum -c`: all OK).
- **Writes outside the report:** only the scratch workspace `jt-island-001-review/work/` (extracted text, page renders, hash manifest). Nothing was written under `/home/chaschel/Documents/go`.
- **Implementation:** none; no Go code, no implementation contract, no architectural proposal.
- **Novelty:** no novelty claims are made. Accepted results are source results or standard algebra.
- **No freeze or approval:** this report does not freeze, approve, or authorize implementation. It records the independent verdict and stops.

*End of report.*

<!-- END-OF-REPORT -->


