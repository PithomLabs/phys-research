# JT-ISLAND-002 Canonical Equation Ledger v1.1

Changelog vs frozen v1.0 (see work/freeze-v1.0.manifest): implements
ledger_review.md corrections A–H (Z/DeepSeek/Claude consolidated adjudication):
§6 factor chain quarantined as [U] + anchor identity corrected, then OQ-4
CLOSED by independent reconstruction (convention route + thermal cross-check;
CGHS [14] pinned); exactness language qualified (F2); regime corrected to
ε≪1, c≫24 (F3); Ω_crit lemma added to owed verifications (F4); R6 conceptual
risk added (F5); (5.13) typo demoted to unverified note (F6); B3 re-anchored
→PASS (OQ-3 fenced); B2 window predicated; finite-anchor sensitivity
UNVERIFIED; OQ-3 sharpened to located bridge-law inconsistency (open); §14
replaced by consolidated handoff. No other content changed.

RST gravity + holographic-CFT matter (GSST), evaporating case.
Canonical source: JHEP 05 (2020) 091 (pp. cited as Jp.N); cross-check
arXiv:2004.00598v2. Locator rule: JHEP numbers always; §3 arXiv equivalent in
brackets (JHEP(3.k)=arXiv(3.k+1), k=12..35). Evidence classes: [S]ource result,
[R]econstructed here, [I]mported, [A]pproximation, [U]nverified, [X]contradicted.

## 1. Executive adjudication and status

Model selection (RST–GSST holographic CFT) was settled by the director and is
not reopened. This ledger establishes the canonical equations: every load-bearing
equation below was verified against rendered JHEP pages; independent
re-derivations are shown where claimed; imports, approximations, and the residual
items (OQ-3 bridge-law audit; unconsumed eternal-island eqs) are explicit.

Obligation grades: O1 action PASS; O2 anomaly/stress PASS (OQ-3 carried);
O3 energy PASS (OQ-4 closed; flux values 4× corroborated); O4 state/frame PASS;
O5 entropy PASS; O6 QES/Page-time PASS. Bridge verdicts: B1 CONDITIONAL
(coherent import [32]); B2 CONDITIONAL (classical-ρ window); B3 PASS at leading
order (re-anchored per OQ-4; OQ-3 fenced to bridge law, not values). Overall:
CONDITIONALLY READY — see §14 for the exact conditions Muse must preserve.

## 2. Model identity and validity domain

Semi-classical RST black hole from thin-shell collapse, matter = strongly
coupled holographic 2D CFT (large c, AdS₃ dual) [S §3.1]. NOT free scalars
(paper: "instead", Jp.6). Regime: ε≡c/48M≪1 (3.11) [S] with c≫24 (Polyakov
dominance, text Jp.9); this implies M≫c/48 — the stronger shorthand M≫c is NOT
assumed unless separately sourced. Leading non-trivial ε order (text Jp.7);
post-scrambling, pre-endpoint (text Jp.17–19). Excluded: endpoint microphysics,
4D gravity, unitarity proof, exact-in-ε claims. QRT assumed (text §6) [I].

## 3. Conventions and symbols

c matter central charge; λ length scale, =1 by rescaling except where restored
[S Jp.5]; φ dilaton; ρ conformal factor (ds²=−e^{2ρ}dx⁺dx⁻) [S (3.3)];
M dimensionless mass parameter; ℳ=λM/π physical mass [S (3.7) — the ledger's
form; Claude's inverted "M=λM_phys/π" corrected]; T^μν=−(4π/√−g)δS/δg_μν,
×2 vs CGHS [S fn.4]; N=c, κ=c/12 comparison-only identifications [S fn.3];
G_N never appears in GSST §3 (any G_N-bearing expression is non-source).
Coordinates: Kruskal x±; rescaled (v,u) (3.17); flat ω-frame (3.19);
asymptotic σ-frame (3.20); asymptotic time t=(σ⁺+σ⁻)/2 (3.21) [all S].

## 4. Action ledger (O1, PASS)

A-I_grav (3.1, Jp.5) [S]: I=(1/2π)∫d²x√−g e^{−2φ}[R+4(∇φ)²+4λ²], covariant.
Varies to (3.12) (metric) and the dilaton e.o.m. (matter-independent) [S text].
A-I_Q (3.28, Jp.9) [S]: −(c/12π)∫dx⁺dx⁻∂₊ρ∂₋ρ, conformal-gauge Polyakov;
nonlocality via t± (fn.5); dominant for c≫24. A-I_RST (3.29, Jp.9–10) [S]:
+(c/48π)∫d²x√−g φR, covariant, defining choice (follows [15]), O(c) like I_Q,
ε→0 limit undisturbed. Relative sign/normalization: −c/12π (gauge) vs +c/48π
(covariant) [S, verified]. Matter: no Lagrangian given — none invented
(Correction A); collapse pulse = compact-support T₊₊ + thin-shock patching
(3.16) + full evolution in [15] [S/I]. Gemini S_RST/S_PL/S_shock/free-scalar
matter/κ: contradicted (see §12).

## 5. R1 field equations, anomaly, stress (O2, PASS)

Classical: (3.4) vacuum/Kruskal e.o.m.; (3.5) e^{−2φ}=M−x⁺x⁻; (3.6) cigar,
horizon vu=0, singularity vu=1 [S Jp.6]. Coupled: (3.12) covariant metric eq;
(3.13) components; (3.14–3.15) F(x⁺)/G(x⁺) solution (both functions of x⁺ —
verified on render); (3.16–3.18) shock patching + (v,u) forms [S Jp.7–8].
Anomaly (3.22): ⟨T^μ_μ⟩=(c/12)R [S; CFT basis I]. Conservation: invoked
continuity input via Christensen–Fulling [31] [I + implicit no-diffeo-anomaly
assumption, recorded]; no standalone Ward identity is printed. Stress (3.23)
[S]: ⟨T₊₋⟩=−(c/6)∂₊∂₋ρ, ⟨T±±⟩=(c/12)(2∂±²ρ−2(∂±ρ)²−t±) — INDEPENDENTLY
RE-DERIVED here from (3.22)+conservation [R §B.2]. t±(x±) = state data.
Transformations (3.24–3.26) [S, verified]; (3.26) additionally CONFIRMED via the
eternal t±(σ±)=+1/2 cross-check [R]. Flux (3.27) [S]: T_−−(σ⁻)=(c/24)×
(1−1/(1+e^{σ⁻})²) from t_−(ω⁻)=0 — retained as source; independent sign audit
open (OQ-3, §12). Lifetime t=1/ε leading order (text Jp.9) [S; conversion
reconstructed §6/OQ-4].

## 6. Collapse solution, flux, energy balance (O3, CONDITIONAL)

Vacuum (3.32) [S], eternal (3.33–3.34) [S], collapse exterior (3.35) [S]:
ds²=−Me^{2φ}dvdu, Ω=M[1−v(u+1)−εlog(−Mvu)], v>1 only. Substitution of (3.32)
and (3.35) into (3.31) HAND-VERIFIED [R]. Strong-coupling boundary Ω=Ω_crit
(Sect. 3.30 text) bounds the spacetime; evaporating apparent-horizon location
has no explicit §3 equation (limitation recorded). Anchor identity (corrected
per F1): σ_A⁻=t_A−σ_A at fixed anchor position, so dσ_A⁻/dt_A=1 [R, from
(3.20)+(5.18)] — the v1.0 "dσ⁻/dt=2 from (3.21)" advanced time at fixed σ⁺ (a
null direction) and does not follow for anchor evolution. The v1.0 §6 chain
(2 × 1/4π) stays SUPERSEDED (error mechanisms: wrong time section;
δg_μν-vs-δg^μν form confusion) — Muse must NOT formalize it. RECONSTRUCTED
CHAIN [R, OQ-4 CLOSED]: (3.27) c/24 per σ⁻ [S] → T_GSST=2·T_CGHS (coefficient
ratios verified vs pinned [14]) → physical flux c/48 per σ⁻ (CGHS text: its
T_−− IS f-particle energy flux) → dσ_A⁻/dt_A=1 → dM/dt=−c/48 → lifetime
48M/c=1/ε (matches source statement — now a check, not input); thermal
cross-check πcT²/12=c/48π (ℳ units) independently agrees. First law holds in ℳ
(S=2πℳ/λ, T=λ/2π); dS/dt=−c/24. Each factor independently established; no
compensating pair. Claude's inversion corrected; his M-only numerics unaffected.

## 7. State and coordinate frames (O4, PASS)

Initial vacuum (3.32) + flat ω-frame (3.19) [S]. Vacuum data t_−(ω⁻)=0 [S text
Jp.9]; t₊(ω⁺)≠0 from incoming flux [S text Jp.15] — ω-frame is a PARTIAL-vacuum
frame (asymmetry recorded). Incoming = thin-shell T₊₊ + [15] interior [S/I].
Coherent state on inertial ℐ⁻ vacuum per [32], S_vN=S_vacuum [I, Jp.15] —
licenses RT (4.7) in vacuum coords. Map SHOWN [R]: ω⁺=σ⁺,
ω⁻=−log(1+e^{−σ⁻}), Jacobian 1/(1+e^{σ⁻}); Schwarzian verified four ways.
"Hartle–Hawking" is eternal-only (§5.1); "Unruh" appears NOWHERE in GSST
(grep-verified) — generic vacuum renaming forbidden (Gemini's label unsourced).

## 8. Generalized entropy (O5, PASS)

R: exterior of constant-Ω anchor curve, Ω_A≫M; drift negligible (text Jp.11)
[S]. (4.1) S_gen=Area/4G_N+S_bulk [S, QRT formula I]. (4.2)
Area/4G_N=2(Ω−Ω_crit), "natural extension", zero for I=∅ [S]. (4.3–4.5) RT via
pure AdS₃ after t±(y)=0 frame + Weyl strip [S/I: follows [11]]. (4.6) cutoff
surface z=δe^{−ρ}, δ=UV cutoff [S]. (4.7) S_bulk=(c/6)log[d²e^{ρ(A)}e^{ρ(I)}]
evaluated at t±=0, d in flat metric, δ dropped as additive const [S — note +ρ
signs]. Endpoints: ONE gravitational point (v_I,u_I) (evaporating); matter
A+I (island) or A+fixed (v₀,u₀) on Ω=Ω_crit with u₀=−ε/v₀ (no-island);
mirrors eternal-only. No S₀ in model. (5.12) island S_gen [S, assumptions:
outside shell, classical ρ at A always / at I pre-endpoint]. (5.21)
no-island (c/12)(t_A−σ_A)+…, v₀-free at leading order [S]. Page time in
RETARDED anchor time t_A−σ_A (not absolute t). Scrambling t_s=log(1/4ε) (5.11).

## 9. QES, branches, Page time (O6, PASS)

Stationarity (5.13–5.14) over (v_I,u_I) [S] ("Extremizing (5.13)" possibly
referring to (5.12): UNVERIFIED editorial observation [U], not a blocker).
Approximation log(v_A/v_I)≫1 → (5.15) quadratic +
log(u_A/u_I)=4(1+u_I) [S]. QES (5.16) u_I=−1+ε/v_I+O(ε²), near-horizon inside;
second root (near singularity) discarded [S]. (5.17–5.18) anchor relations [S].
Self-consistency (5.19) log(v_A/v_I)≈2σ_A+log(1/ε)≫1 ✓ [S]; validity
v_I>1 ⟺ t_A−σ_A≳log(1/ε), from scrambling time after first radiation [S].
(5.20) island 2M−(c/24)(t_A−σ_A)+… [S]. Branch rule = min after extremization
(text Jp.16) [S]. (5.22) t_Page=1/3ε=16M/c [S] — crossing algebra established
within the displayed leading-order semiclassical expressions [R §F.1; this
algebraic exactness does not make the physical Page time exact at arbitrary ε].
Leading order only; endpoint beyond control (§6); unitarity never claimed.

## 10. R1/R4 bridges

B1 state/frame — CONDITIONAL (import [32]). Contract: source R1 t-data+(3.26);
target (4.7) vacuum evaluation; mapping (ω⁺,ω⁻) instruction (Jp.15) + coherent
license; assumptions [32], t_−(ω⁻)=0; preserved vacuum-form RT; lost
non-coherent generality + t₊ asymmetry noted; regime ε≪1 post-scrambling;
obligation OQ-3 (sign caveat, not touching the frame instruction); failure =
non-coherent state; evidence Jp.15+(4.7).
B2 geometry — CONDITIONAL (approximation window). Contract: source (3.35)+ρ;
target (5.12) area + e^{ρ} factors; mapping direct substitution (paper-stated);
assumptions classical ρ at A/I with stated domains; preserved leading logs;
lost subleading/O(ε²)/endpoint; regime M≫c/48 pre-endpoint; failure = island
leaving classical domain; evidence Jp.16–17. VALIDITY PREDICATE (Muse owed):
formalize "island in classical-ρ domain ∧ log(v_A/v_I)≫1 holds" and test it
across the post-scrambling, pre-endpoint Page-time range; (5.16) coefficient
UNVERIFIED off the large-log limit (finite-anchor sensitivity per review5).
B3 energy/entropy — PASS (leading order; re-anchored per OQ-4). Contract:
source = source-stated (3.27) + fn.4 + anchor identity + pinned [14] + thermal
cross-check (v1.0 chain superseded, never consumed); target slopes −c/24/+c/12,
ratio 1/3; mapping first law + (6.1) 2:1 reasoning; non-circularity preserved
(flux-side conversion vs entropy-side slopes independently established).
OQ-3 fenced: it audits the (3.25)/(3.26) bridge law, not the flux values B3
consumes (4× corroborated).

## 11. Reality checks

RC1 inventory: 2D dilaton gravity, variational calculus, Polyakov/anomaly
theory, 2D CFT state data, holographic RT/AdS₃ (imported), QES theory,
saddle ODEs. Regimes §2; excluded endpoint/4D/unitarity/greybody (absent per §6).
RC2 intersections: R1∩R4={Ω,ρ,t±+frame,M,flux}. Removal probes (all pass):
no anomaly → no flux/lifetime/competitor; no RST → no Ω-linearity; no vacuum
frame → (4.7) undefined; no QRT → R4 undefined; no −Ω_crit → normalization
breaks. Necessity only; minimality unclaimed. Per-phase RC records:
work/phaseX-checks.md.
RC3 bridges: B1–B3 above add information (instruction+license; substitution;
numerical identity). Imported vs derived vs approximated separated throughout;
nothing promoted by repetition (RC-CONT-01).

## 12. Disputed, corrected, unresolved

Gemini mathematical specification: quarantined — S_CGHS faithful; S_RST,
S_PL, S_shock, free-scalar matter, κ, Ω̃, stress t±-sign, constant flux, dM/dt,
S_grav, S_matter, β-QES, island slope, "exact"/"unitary" claims contradicted
or unsourced with decisive evidence (work/phaseA-ledger.md §A.5, work/phaseE-ledger.md §E.6). Coincidences:
S_gen form, one endpoint, no-island slope value (λ=ℏ=1).
Claude: "M=λM_phys/π" inverted vs (3.7) — corrected; Bridge 3 requirement
sustained and satisfied (§6); baseline otherwise confirmed.
ChatGPT GSST analysis: confirmed (§2.4 trio + conditional bridge all verified);
DREH assessment carried as review-record (papers unpinned).
Muse HSS: fixture retained; "manufactured interoperability" answered for the
specified construction (imports stay tagged); HSS-paper content unverified here.
DeepSeek DREH: deferred; in-record objections preserved. DeepSeek modular
Hamiltonian: separate research lead ONLY — must not become a mandatory POC
dependency (director ruling); B2's contract stays substitution-based.
Z review (16 citations confirmed; crosswalk corroborated from arXiv side):
F1 energy chain quarantined [U] (OQ-4); F2 exactness language qualified; F3
regime corrected; F4 Ω_crit lemma owed; F5 R6 added; F6 typo unverified.
Claude ledger review: mass/lifetime/slopes/Page-numbers/(3.7)/flux/Ω_crit/island
slope (−1.994 vs −2) independently reproduced; sign conflict made concrete
(folded into OQ-3); (5.16) coefficient anchor-sensitive → UNVERIFIED.
Resolved this revision: OQ-4 CLOSED (convention route + thermal agreement;
[14] pinned); B3 re-anchored → PASS; O3 → PASS.
Unresolved: OQ-3 (sharpened: factor-2 (3.25)/(3.26) inconsistency + sign vs
CGHS; flux values fenced out as 4× corroborated); eternal (5.6–5.10)
unconsumed, scoped out.

## 13. Risk register

R1 OQ-3 sign audit: bounded verification residual; mitigation = preserve
(3.27)-sign as source-stated; Muse must not "fix" it.
R2 Import stack (QRT, RT, [15] interior, [31] method, [32] coherent, anomaly
basis): load-bearing imports; mitigation = tags preserved, no promotion.
R3 Leading-order scope: O(ε) corrections + endpoint unknown; mitigation = no
exactness/unitarity claims (Gemini failure mode fenced).
R4 Anchor/reference prescriptions: leading-order-insensitive per paper
(qualitative argument); mitigation = retarded-time form preserved.
R5 Locator drift: §3 JHEP↔arXiv shift; mitigation = canonical JHEP + index
(work/locator-index.md).
R6 Conceptual status of islands in long-range gravity (Geng et al., JHEP
01(2022)182): objection to islands in asymptotically flat space. Island
calculation conditional on adopted QRT prescription; debate explicitly NOT
resolved by JT-ISLAND-002; interpretation-level, non-blocking for internal
mathematical consistency.

## 14. Consolidated Muse handoff (CONDITIONALLY READY; supersedes prior notes)

MAY PROCEED on R1, state/frame, R4 core, AND the re-anchored energy chain
(§6/OQ-4), with all tags explicit. MUST NOT formalize the SUPERSEDED v1.0 §6
chain (2 × 1/4π); must not silently flip the (3.27) sign (record and resolve
against the rendered page); must not derive QRT/RT/[32]; no DREH/HSS equations;
no Phys-Pseudo/code beyond separately approved scope; apply RC-CONT-01
throughout.
MUST PRESERVE: all [I]/[A]/[U] tags; OQ-3 flag; leading-order scope;
retarded-time Page time; import list (QRT, RT, [15] interior, [31] method, [32]
license, anomaly basis, [14] convention pin); no S₀; no free-scalar
substitution; no invented matter Lagrangian; regime ε≪1, c≫24 (never silently
strengthened).
OWED VERIFICATIONS: (3.1)→(3.12) variation; OQ-3 bridge-law audit; Ω_crit lemma
(role of subtraction; empty-island vanishing; prescription-vs-derivation
distinction); B2 validity predicate over the Page-time range; eternal (5.6–5.10)
if ever consumed. STATUS: ledger remains CONDITIONALLY READY until obligations
discharge or are retained as conditional in the final formalization.

## 15. Bibliography (primary, pinned in work where consumed)

GSST JHEP 05(2020)091 [canonical; SHA-256 in sources/README.md];
arXiv:2004.00598v2 [cross-check]. Imported authorities (cited, not pinned):
CGHS [14]; RST [15]; Brown–Henneaux [BH]; Christensen–Fulling [31];
Fiola et al. [32]; Page-curve reasoning [33]; Almheiri et al. [11]; QRT/RT
prescription [2–4]. Geng et al., JHEP 01(2022)182 (R6 objection; cited, not
consumed). Agent reports: Claude (baseline), ChatGPT, Muse, DeepSeek, Gemini
(quarantined math), Qwen (out of scope) + JT-ISLAND-002_review.md (director
record). Reviews of this ledger: JT-ISLAND-002_Ledger_Review.pdf (Z, 16
citations confirmed + F1–F6), plans/review5.md (DeepSeek + Claude numeric
checks), ledger_review.md (director consolidated adjudication — governing).
Investigation record: work/execution-brief.md, work/locator-index.md,
work/phase{0,A,B,C,D,E,F}-ledger.md, work/phase{0,A,B,C,D,E,F}-checks.md,
work/phaseG-adjudication.md, work/open-questions.md, work/risk-register.md,
work/freeze-v1.0.manifest, work/oq3-probe.md, work/oq4-probe.md, work/log.md
(T0–T6), work/pages/ (rendered evidence).

