# JT-ISLAND-001 Audited Working Baseline v0.1 (PROVISIONAL — not frozen)

**Status:** working foundation for Phys-Pseudo v0.1. Every equation below is
re-sourced to a primary locator or derived in-note; dossier printings are
never premises (dossier equation locators are VOID per audit v1.0 §5).
Approval and freezing remain the director's, after OpenCode review.

**Authority:** audit v1.0 (findings) + addendum v1.1 (applied director
rulings). Open obligations are marked `?OB-n`; they become Phys-Pseudo `?`
items, never silent assumptions.

---

## 1. Adopted canonical model (Decision 1)

**AMMZ holographic-matter extension of the AEMM two-sided evaporating JT +
bath setup.** Two-sided AdS₂ JT black hole in Hartle-Hawking state; right
boundary coupled at u = 0 to an auxiliary half-line CFT bath (same CFT);
right side evaporates. Matter is a holographic CFT₂ with large central
charge c (AMMZ specialization of AEMM's general conformal matter).

Continuity witness (AMMZ §2.1 close): AMMZ states its formalism "can be
directly applied to the case considered in [9] by specializing to
Jackiw-Teitelboim gravity, with the only difference being that we consider
matter composed of a holographic CFT₂, rather than general matter."
Compatibility is therefore asserted by AMMZ itself; §8 below still
justifies it per equation class (director's critical rule).

NOT adopted: single-sided collapse (auditor's paused recommendation,
corrected by director); dossier's blended narrative; any equation without a
primary locator.

## 2. Coordinate dictionary (one map, used everywhere)

- x^± = t ± z: Poincaré null coords; AdS₂ metric
  ds² = −4dx^+dx^−/(x^+ − x^−)² (AEMM(10)). Boundary at z → 0 (x^+ → x^−).
- u: boundary proper (physical) time. t = f(u): Poincaré time via
  boundary reparametrization f (AEMM §2.1); f(0) = 0, f′(0) = 1.
- y^±: coords with x^± = f(y^±); metric AEMM(15); cutoff at
  (y^+ − y^−)/2 = ϵ. The y-frame flat metric −dy^+dy^− is where the
  post-shock ingoing <T> vanishes (AEMM(21) context).
- w, w̄: half-plane coords mapping the quench state to half-line vacuum
  (AEMM §3.2, map (43)/(47)).
- z = ϵf′(u): proper-distance cutoff surface (AEMM(64)ff). The dossier's
  constant z = ϵ is REJECTED (missing f′ factor).
- σ: bath spatial coordinate (bath = half-line). σ₀: radiation-region
  inner endpoint (AMMZ §3.1.1-3.1.2).

## 3. Convention pins (binding; dossier never stated these)

- Signature (−,+). AdS₂ radius = 1; on-shell R = −2 (AEMM(9)).
- Stress tensor: T_μν = −(2/√−g)δI_M/δg^μν (AEMM normalization; NOTE AEMM
  warns it differs from common 2D-CFT literature by −2π — conversions
  across literature must show the factor).
- Schwarzian: {f,u} = f'''/f' − (3/2)(f''/f')² (standard).
- G_N: 2D Newton constant (AMMZ's G_N^(2) = AEMM's G_N — same quantity).
- c: CFT central charge. Large-c + strong coupling required for Einstein
  3D dual (AMMZ §2: "to have an Einstein gravity dual we need that the CFT
  is suitably strongly coupled") — RECORDED as extra R4-route hypothesis.

## 4. Symbol dictionary (single meaning each; resolves dossier collisions)

- φ₀: extremal/topological dilaton constant. S₀ = φ₀/4G_N.
- φ: dynamical dilaton field. φ₀ ≫ φ (AEMM).
- φ_b = φ̄_r/ϵ: boundary dilaton value (RENAMED from dossier's φ_h; the
  dossier used φ_h for this on p.3).
- φ_hor(t): dynamical dilaton at the horizon (RENAMED; the dossier's p.9
  φ_h). S_BH(t) = (φ₀ + φ_hor(t))/4G_N.
- k = (c/12)(4G_N/φ̄_r) ≪ 1: evaporation rate (AEMM(24)).
- κ_surf = 2πT: surface gravity (RENAMED; dossier's p.8 κ collided with
  AMMZ/AEMM's evaporation-rate κ/k).
- T₀/T₁: pre/post-shock temperatures; T(t): instantaneous; T_H: Hawking.
- E_S: shock energy from turning on the coupling (AEMM(46),(30)):
  E_S > cT₀/24 (causality); E_S → ∞ limit discussed AEMM §3.2 (regulator
  subtlety RETAINED, not smoothed over).
- t_∞ = lim_{u→∞}f(u): Poincaré time of the boundary endpoint (AEMM(29)).

## 5. R1 corrected equations

### 5.1 Action (AEMM(6)-(8); dossier 4× REJECTED)

I₀ = (φ₀/16πG_N)[∫_M d²x√−g R + 2∫_∂M K] (topological; gives S₀ = φ₀/4G_N).
I_G = (1/16πG_N)[∫_M d²x√−g φ(R+2) + 2∫_∂M φ_b K].
I_M = I_CFT[g] (holographic CFT₂; see §8 compatibility).
The dossier's (K−1) counterterm variant is NOT adopted (unexplained
deviation from AEMM; ?OB-1: revisit only if on-shell-action finiteness is
needed — not needed for equations of motion).

### 5.2 Geometric constraint (AEMM(9))

δI/δφ = 0 ⟹ R + 2 = 0. Bulk is locally AdS₂. VERIFIED CLEAN (audit §6).

### 5.3 Metric equations (AEMM(11a-c); dossier sign/coefficient REJECTED)

Covariant form (AEMM-consistent; trace-verified + vacuum-tested in audit):
(1/8πG_N)(g_μν∇²φ − ∇_μ∇_νφ − g_μνφ) = <T_μν>. NOTE MINUS g_μνφ
(dossier printed plus — sign defect; dossier's own p.16 form agrees here).
Coordinate forms (AEMM(11a-c), x^± Poincaré):
(11a) 2∂_+∂_−φ + [4/(x^+−x^−)²]φ = 16πG_N T_+− (NOTE 16π — dossier's 8π REJECTED).
(11b) −[1/(x^+−x^−)²]∂_+((x^+−x^−)²∂_+φ) = 8πG_N T_++ (dossier form VERIFIED ✓).
(11c) same with +↔− (VERIFIED ✓).
Semiclassical replacement T_ab = <T_ab> (AEMM).

### 5.4 Stress tensor: anomaly + Ward identity as INDEPENDENT inputs (A + AP)

Trace (AEMM fn.2; AP context): <T^μ_μ> = cR/24π for conformal matter;
absorbed into φ₀ renormalization φ₀ → φ₀ + cG_N/3 (dossier's R-form equality
<T_+−> = (c/24π)Rg_+− is SUSPENDED — 2×-suspect under standard conventions;
?OB-2: pin or drop; NOT used downstream).
Components (VERIFIED against AP(5.1)-(5.2), notation mapped N↔c, ω↔ρ):
<T_+−> = −(c/12π)∂_+∂_−ρ.
<T_±±> = (c/12π)(∂²_±ρ − (∂_±ρ)²) + <:T_±±:>_state.
Transformation law with Schwarzian (AEMM(42)); post-quench shock form
(AEMM(44)-(46)) and AdS form (AEMM(61)):
<T_−−(x^−)> = E_S δ(x^−) − (c/24π){y^−,x^−}Θ(x^−).
Ward identity ∇^μ<T_μν> = 0 from diffeomorphism invariance (standard;
dossier's protocol.md citation VOID). Light-cone expansion
∂_−<T_++> + ∂_+<T_+−> − 2<T_+−>∂_+ρ = 0 (+↔− twin): DERIVED-IN-BASELINE
(direct computation from Γ^+_{++} = 2∂_+ρ, all other Γ = 0 in conformal
gauge — shown in audit working notes). Substituting the trace relation
gives chirality ∂_−<:T_++:> = ∂_+<:T_−−:> = 0 (dossier p.13 argument
structure RETAINED with verified inputs).
S_Polyakov nonlocal formula: SUSPENDED as unverified (?OB-3); baseline uses
the anomaly-transformation route (AEMM/AP), which needs no explicit S_Pol.

### 5.5 Boundary dynamics + evaporation (AEMM §2.1-2.2)

Schwarzian action (AEMM(13)); ADM energy E(u) = −(φ̄_r/8πG_N){f(u),u}
(AEMM(14); dossier's missing minus REJECTED — positive-mass check via
E₀ = πφ̄_rT₀²/4G_N (AEMM(20)) with {tanh} < 0).
Energy balance ∂_uE = f′(u)²(T_−− − T_++) (AEMM(22)).
Evaporation law {f,u} ∝ e^{−ku} (AEMM(23)); E(u) = E₁e^{−ku}, E₁ = E₀ + E_S
(AEMM(25)); f-equation + Bessel solution (AEMM(27)-(29)); working
approximations (AEMM(32)-(34)); dilaton integral solution (AEMM(35)-(37)).
2D flux form dM/dt = −(πc/12)T_H² + F_in: RETAINED as standard result
(AEMM(22) structure); full derivation from E(u) + (21) is OWED in
Phys-Pseudo Part A (?OB-4 if any step resists pinning).

### 5.6 Coupling + interface (AEMM §3.1-3.2; shock RETAINED)

Quench coupling at t = 0 with time-matching t = f(u); E_S shock
(AEMM(44)-(46)); E_S → ∞ regulator limit + E_S ∼ ϵ^{−1} lattice picture
(AEMM §3.2 — RETAINED as stated subtlety); causality bound E_S > cT₀/24
(AEMM(30)). AMMZ-side corroboration: AMMZ §3.2.2 couples over Δt with pulse
E ∝ c/Δt (same transient physics, holographic dual language). The dossier's
smooth T^grav = T^bath matching is SUPERSEDED by the quench + shock
treatment; any smooth-matching limit must be derived as Δt/E_S^{−1} → 0
with stated regulator, not assumed (?OB-5).

## 6. R4 corrected inputs

### 6.1 Generalized entropy (AMMZ(10); AEMM(69)-(73))

One-endpoint form: S_gen(a) = (φ₀ + φ(a))/4G_N + S_matter(R ∪ I(a)).
Renormalization split RETAINED: S_gen = (φ₀^Ren + φ^Ren)/4G_N + S^Ren
(AEMM(69)); φ₀^Ren = φ₀^Bare − (2/3)cG_N log ϵ (AEMM(72));
S^Bare = S^Ren − N(c/6)log ϵ (AEMM(73)). Regulator-independence of S_gen
follows AEMM §4.1 (imported result; derivation NOT repeated).
Multi-endpoint: sum one term per endpoint (Decision 2).

### 6.2 Regions (AMMZ §3.1.1-3.1.2)

Radiation region R_t: bath CFT outside σ₀ (σ₀ small, includes QM boundary
d.o.f. in AMMZ's description). Candidate island I_t: interior region bounded
by QES point a(t) = (x^+_r, x^−_r) near the horizon. One-endpoint
configuration adopted (dossier's stated I_t = (−∞, a(t)]).

### 6.3 No-island branch (VERIFIED log form + AMMZ(14) cross-check)

S_no-island(R_t) = (c/6)ln[(β/πϵ)sinh(πt/β)] (+S₀ DROPPED — unsourced in
dossier; ?OB-6: re-add only with sourced initial-entropy treatment).
Late-time slope πcT/6 (derived from the log form; cross-checked against
AMMZ(14) S_Rad = (πc/6)∫T dt′ and AEMM(112)). Dossier linearizations VOID.

### 6.4 QES (AMMZ(12) ← Penington + AEMM §4; dossier formula VOID)

Adopted position: y_e^+ = t − (1/2πT(t))log[(S_Bek(T(t)) − S₀)/c] + …
(AMMZ(12)), i.e. scrambling-time delay behind the horizon. Status:
SOURCE-REPORTED import (AMMZ recalls it from [8,9]); JT-specific derivation
lives in AEMM §4 (Eqs. (75)-(77) early, (92)-(93), (104)-(107) late).
Stationarity conditions ∂S_gen/∂a^± = 0 RETAINED in form; solution obligation
carried as import (no independent re-derivation of AEMM §4 in v0.1 —
recorded, not hidden).

### 6.5 Island branch value (Decision 2: 1×/endpoint)

Late-time island-branch entropy tracks the instantaneous BH entropy:
S_island ≈ S_BH(T(t)) + logs (AMMZ(13) applied to radiation via the shared
QES surface, AMMZ §3.1.2). Dossier's 2× VOID. Bulk-matter O(c·log)
contributions subleading per AEMM (line 1105 context); dossier's O(c⁰)
REJECTED as misstated order.

### 6.6 Branch selection (imported, NOT postulated)

Minimization S(R_t) = min{S_no-island, S_island} imported from RT/HRT
minimality (AMMZ §3.3: "The minimality condition in the RT/HRT prescription
leads to…"). Island rule S(A) = Min Ext[S_eff + Area/4G_N] (AMMZ(15)) is an
EFFECTIVE PRESCRIPTION: derived-from-RT/HRT within the holographic setup
(conditional on RT/HRT correctness, AMMZ line 469) and proposed beyond it
(AMMZ §4; "progress on direct derivation [46,47]"). Dossier's rationale.md
self-citation VOID. Extremization and minimization are SEPARATE operations
(prompt9 §4B requirement preserved).

### 6.7 Page time: OPEN (derived-not-frozen per Decision 2)

Definition ADOPTED: t_P where corrected branches intersect (AMMZ §3.3;
Penington Eq. (2) S_rad = A_hor/4G_N equivalent). VALUE open pending Step 5
independent reconstruction. Neither dossier 3βS/πc nor audit-note 6βS/πc is
adopted as a premise.

### 6.8 Entropy-computation route: DECIDED IN STEP 4 (issue D)

Candidates: (i) AMMZ 3D-RT route under holographic matter (no 2D CFT
computation needed); (ii) AEMM replica/G(η) machinery (Eqs. (56)-(60),
(66)-(68)) with explicit η/channel/state hypotheses. The dossier's
single-interval-on-complement formula is ADMISSIBLE ONLY with a purity proof
unavailable in the adopted two-sided setup — hence route (i) or (ii), decided
in Step 4 with the rejected alternative recorded.

## 7. Corrected hierarchy and regime (ASM-02 REWRITTEN)

1 ≪ c (large-c holographic CFT) ≪ S₀ = φ₀/4G_N ≪ S^i_Bek
(dynamical BH entropy); equivalently c ≪ φ₀/G_N with φ₀ ≪ φ_hor-scale and
k = (c/12)(4G_N/φ̄_r) ≪ 1 (AEMM(24)). φ₀ ≫ φ (AEMM dilaton assumption).
S₀ ≪ S^i_Bek REQUIRED for island domination (AMMZ 2S₀-price argument:
large S₀ kills the island). Dossier's φ₀ ∼ φ_h REJECTED (contradicts AEMM
and tensions AMMZ's island conclusion). Consequences (dossier bullets,
RETAINED in substance): S_CFT ∼ O(c) ≪ S_grav (leading quantum correction);
metric/dilaton fluctuations suppressed (semiclassical background valid);
higher-derivative corrections negligible. Breakdown: k^{−1}log k timescales
(near-extremal, AEMM), c ∼ S₀ (QES delocalization), G_N → ∞ (strong quantum
gravity — dossier RC2's "removal" limit direction CORRECTED: G_N → 0 is the
classical/no-island limit, G_N → ∞ is breakdown, not removal).

## 8. AEMM↔AMMZ per-class compatibility (Decision 1 critical rule)

(a) Pure JT gravity sector (action, R+2=0, metric eqs, Schwarzian/ADM):
matter-independent → IMPORTS CLEANLY (matter enters only via <T>).
(b) CFT anomaly/transformation laws (AEMM(42),(44)-(45); AP(5.1)-(5.2)):
hold for ANY 2D CFT incl. holographic → IMPORTS CLEANLY at fixed c.
(c) Quench/coupling + E_S + f-trajectory: AEMM's model needs conformal
matter + conformal BCs — holographic CFT qualifies; AMMZ §3.2.2 describes
the same transient (E ∝ c/Δt) → COMPATIBLE, shock RETAINED.
(d) Evaporation law/decay (AEMM(23)-(25)): follows from (a)-(c) → IMPORTS
with (a)-(c) justifications.
(e) Entropy computations: METHODS DIVERGE (AEMM replica/G(η) vs AMMZ 3D-RT)
→ NOT auto-compatible; Step 4 selects ONE route with hypotheses. NOTE:
AEMM's entropy formulas remain VALID under holographic specialization (a
holographic CFT is a CFT) but AMMZ's R4 chain does not use them — the spec
must not mix-and-match without the Step-4 decision.
(f) Extra R4 hypotheses from AMMZ: strong coupling (Einstein dual),
S₀ ≪ S^i_Bek, RT/HRT correctness — RECORDED as route-(i) assumptions.

## 9. Open obligations carried into Phys-Pseudo

?OB-1: (K−1) counterterm variant — shortlisted as not-needed; revisit only
for on-shell-action finiteness. ?OB-2: R-form trace equality — pin
conventions or drop (dropped from downstream use). ?OB-3: S_Polyakov
coefficient — suspended; anomaly-transformation route used instead.
?OB-4: full dM/dt derivation from E(u)+(21) — owed in Part A. ?OB-5:
smooth-matching limit from quench — derive or drop. ?OB-6: initial-entropy
S₀ term in no-island branch — re-add only if sourced. ?OB-7: Page-time value
— Step 5 derivation (unfrozen). ?OB-8: entropy route — Step 4 decision.
?OB-9: AEMM §4 QES re-derivation — explicitly NOT performed in v0.1
(import chain AMMZ(12)←[8,9] recorded instead).

## 10. Issue-D resolution: entropy-computation route (Step 4 decision)

DECISION: route (i) — R4 radiation entropy follows AMMZ's 3D-RT/HRT
construction (AMMZ(11), §3) under the holographic-matter specialization. NO
2D-CFT entropy computation is used in the R4 chain.

Reasons: (1) The R4 deliverable is the RADIATION Page curve; only AMMZ
constructs it. AEMM's own bath-entropy result (AEMM(112)) grows forever —
AEMM does not produce a radiation Page curve, so route (ii) cannot deliver
R4. (2) The dossier's complement formula requires pure-state S(R∪I) =
S([a,b]), unavailable in the adopted two-sided setup (R∪bath mixed,
entangled with L) — REJECTED. (3) AMMZ's route requirements (holographic
CFT, large c, strong coupling for Einstein dual, RT/HRT correctness) are
all satisfiable within the adopted model and are RECORDED as route
assumptions (not proven).

REJECTED alternative (ii): AEMM replica/G(η) machinery (AEMM(56)-(60),
(66)-(68)) — VALID mathematics under holographic specialization (a
holographic CFT is a CFT) but computes one-sided/BH entropies, not the
radiation Page curve. RETAINED ONLY as consistency cross-checks:
early-time slope πcT/6 agreement (AMMZ(14) vs AEMM(112)), QES derivations
(AEMM §4) behind the AMMZ(12) import. Explicitly NOT mixed into the island
derivation. If the holographic specialization is ever dropped, route (ii)
with full η/channel/state hypotheses is the documented fallback (?OB-8
CLOSED by this decision; fallback G(η) hypotheses recorded in §6.8 context
above).
