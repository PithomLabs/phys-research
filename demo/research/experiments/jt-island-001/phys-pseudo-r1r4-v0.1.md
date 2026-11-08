# Phys-Pseudo JT-ISLAND-001 v0.1 — R1 (Part A), R4 (Part B), interop (Part C)
# Model: AMMZ holographic-matter extension of AEMM two-sided evaporating JT+bath.
# Premises: working-baseline.md only. Dossier printings are never premises.
# Status tags (prose only, per PSEUDO.md): [SR]=SOURCE-REPORTED [D]=DERIVED-in-spec
# [C]=CONDITIONAL [P]=POSTULATED [A]=APPROXIMATE [U]=UNSUPPORTED.
# Primary locators: AMMZ(1)-(15); AEMM(6)-(112); AP(5.1)-(5.6); Penington Eq.(2).
# Page-time value: DERIVED-NOT-FROZEN (Part B). No implementation content.

# ============ PART A — R1: semiclassical matter-geometry consistency ============

given jt_action: total action I=I_0+I_G+I_M with I_0=(phi0/16πG_N)[∫R+2∫K], I_G=(1/16πG_N)[∫phi(R+2)+2∫phi_b K], I_M=I_CFT[g] for holographic CFT2 [SR: AEMM(6)-(8); holographic specialization AMMZ §2]
given regime_hierarchy: 1<<c<<S_0<<S^i_Bek, k=(c/12)(4G_N/phibar_r)<<1, phi0>>phi, S_0<<S^i_Bek required for island domination [A/C: baseline §7; AEMM(24); AMMZ §2, 2S_0-price]
define ricci_ads2 := 2D Ricci scalar R with AdS on-shell value R=-2
define coords_poincare := x^±=t±z with ds²=-4dx^+dx^-/(x^+-x^-)², boundary at z→0 [SR: AEMM(10)]
define coords_physical := boundary time u with t=f(u), y^± via x^±=f(y^±), cutoff z=eps f'(u) [SR: AEMM §2.1, §3.1]
define conformal_gauge := ds²=-e^{2rho}dy^+dy^- form of the metric
define stressT := <T_μν>=-(2/√-g)δI_M/δg^μν in AEMM normalization (differs from common 2D-CFT literature by -2π) [SR: AEMM]
define schwarzian := standard {f,u}=f'''/f'-(3/2)(f''/f')²
assume no_diffeo_anomaly: matter effective action is diffeomorphism invariant with no diffeomorphism anomaly [P: standard; BR-03 re-sourced]
assume conformal_matter: matter is 2D CFT of central charge c, holographic and strongly coupled [P: AMMZ §2 specialization]
assume quench_state: Hartle-Hawking/eternal state plus right quench coupling at t=0 with E_S shock transient [P: AEMM §3.1-3.2]
given trace_anomaly: <T^μ_μ>=cR/24π for conformal matter, absorbed into phi0 renormalization phi0→phi0+cG_N/3 [SR: AEMM fn.2]
claim trace_R_form: equality <T_+->=(c/24π)Rg_+- in stated conventions [U: 2×-suspect under standard conventions; conventions never pinned by dossier]
  ? ?OB-2: pin R/□ conventions to verify trace_R_form, or drop it permanently; currently dropped from all downstream use
given stress_forms: <T_+->=-(c/12π)∂_+∂_-rho and <T_±±>=(c/12π)(∂²_±rho-(∂_±rho)²)+<:T_±±:>_state [SR: AP(5.1)-(5.2)]
given shock_stress: <T_--(x^-)>=E_S δ(x^-)-(c/24π){y^-,x^-}Θ(x^-) after scaling f'(0)=1 [SR: AEMM(21)/(61)]
derive curvature_constraint: R+2=0, bulk locally AdS2, from [jt_action] | vary I_0+I_G wrt dilaton phi; delta(R+2) (AEMM(9))
derive metric_eqs: (g∇²phi-∇∇phi-gphi)/8πG_N=<T> with (11a) 2∂∂phi+[4/(Δx)²]phi=16πG_N T_+- and (11b-c) constraints at 8πG_N from [jt_action, stressT] | vary wrt metric; NOTE minus gphi and 16π-vs-8π pattern (AEMM(11a-c); cross-checked AP(5.3)/(5.5))
derive ward_identity: ∇^μ<T_μν>=0 from [no_diffeo_anomaly, stressT] | diffeomorphism invariance of I_M; independent of trace_anomaly
derive lightcone_conservation: ∂_-<T_++>+∂_+<T_+->-2<T_+->∂_+rho=0 and +↔- twin from [ward_identity, conformal_gauge] | Γ^+_{++}=2∂_+rho, all other Γ vanish in conformal gauge
derive chiral_flux: ∂_-<:T_++:(y^+):>=0 and ∂_+<:T_--(y^-):>=0 from [lightcone_conservation, trace_anomaly, stress_forms] | substitute trace relation into conservation; geometric pieces cancel leaving state piece chiral
derive schwarzian_dynamics: boundary action S_G=(phibar_r/8πG_N)∫{f,u} and ADM energy E(u)=-(phibar_r/8πG_N){f,u} from [jt_action, coords_physical, schwarzian] | boundary reduction (AEMM(13)-(14)); minus sign load-bearing, positive-mass check E_0=πphibar_rT_0²/4G_N
derive energy_balance: ∂_uE(u)=f'(u)²(T_--−T_++) from [schwarzian_dynamics, stressT] | vary I_G+I_CFT wrt boundary time; ingoing minus outgoing flux (AEMM(22))
derive evap_law: E(u)=E_1 e^{-ku} with E_1=E_0+E_S from [energy_balance, shock_stress, schwarzian_dynamics] | Schwarzian decay equation (AEMM(23)-(25))
derive bulk_conservation: local ∇^μ<T_μν>=0 holds throughout M_grav bulk from [ward_identity] | Ward identity is local; bulk/bath/interface accounted separately below
claim interface_quench: coupling matches grav/bath times via t=f(u) with E_S shock transient; smooth T^grav=T^bath continuity NOT assumed [C: AEMM §3.1-3.2; AMMZ §3.2.2 pulse corroboration]
claim global_balance: BH mass decay balances net bath flux; 2D form dM/dt=-(πc/12)T_H²+F_in [C: standard form, AEMM(22) structure; full derivation from E(u)+(21) owed]
  ? ?OB-4: exhibit dM/dt line-by-line from E(u) and shock_stress, or downgrade global_balance to source-reported-only
claim shock_regime: E_S shock transient present at the quench interface [C: AEMM §3.1-3.2]
if shock_regime:
  then
    claim shock_retained: E_S with E_S>cT_0/24 causality bound and E_S→∞ regulator subtlety retained [SR: AEMM(30), §3.2]
  else
    claim smooth_limit_unproven: smooth-matching limit requires Δt/E_S^{-1}→0 derivation with stated regulator [U]
? ?OB-5: determine which branch of the shock_regime conditional holds; derive the smooth limit or keep quench-only matching
claim polyakov_explicit: explicit nonlocal S_Polyakov[g] with verified coefficient and stated conventions [U: dossier form unverified against inspected sources]
  ? ?OB-3: verify polyakov_explicit via Polyakov-original/textbook check, or drop permanently; baseline uses anomaly-transformation route (trace_anomaly, stress_forms) instead

# ============ PART B — R4: Page-curve reconstruction ============

given radiation_region: R_t is bath CFT outside sigma_0 at boundary time t [SR: AMMZ §3.1.1-3.1.2]
given island_region: I_t=(-infinity,a(t)] in M_grav with one-endpoint boundary ∂I={a}, a=(x^+_r,x^-_r) near horizon [P: configuration adopted; Decision 2]
define sgen_functional := S_gen(a)=(phi0+phi(a))/4G_N+S_matter(R∪I(a)) with bare/Ren split per AEMM(69)-(73) [SR: AMMZ(10)]
define dilaton_term := (phi0+phi(a))/4G_N per QES endpoint, summed over endpoints [SR: AMMZ(10); Decision 2]
define bh_entropy_fn := S_BH(T)=(phi0+phi_hor(T))/4G_N with phi_hor the dynamical horizon dilaton (renamed; dossier phi_h collision resolved)
assume holographic_entropy_route: radiation entropy computed via 3D-RT/HRT under holographic matter; needs large c, strong coupling, RT/HRT correctness [P/A: AMMZ(11), §3; Step-4 route (i)]
assume endpoint_count_one: island has exactly one QES endpoint in the adopted configuration [P: dossier stated config, retained]
claim noisland_branch: S_no(t)=(c/6)ln[(beta/πeps)sinh(πt/beta)] without S_0 term [SR: verified log form; AMMZ(14) cross-check]
derive noisland_slope: late-time S_no→(πcT_H/6)t+(c/6)ln(beta/2πeps) from [noisland_branch] | expand sinh at t>>beta; slope πcT_H/6, UV constant recorded not dropped
claim qes_stationarity: ∂S_gen/∂a^+=0 and ∂S_gen/∂a^-=0 fix the QES [SR: stationarity form; extremization step]
claim qes_position: y_e^+=t-(1/2πT(t))log[(S_Bek(T(t))-S_0)/c]+… behind horizon [SR: AMMZ(12) import ←Penington + AEMM §4]
  ? ?OB-9: AEMM §4 QES re-derivation explicitly not performed in v0.1; close by re-deriving (75)-(77)/(92)-(93)/(104)-(107) or keep import chain with fresh verification
claim island_value: S_isl(t)≈S_BH(T(t))+logs with T(t)=T_1 e^{-kt/2} and AMMZ↔AEMM rate identity kappa=k [SR: AMMZ(13) via shared surface; rate identity derived in branch-reconstruction §2]
claim branch_selection: S(R_t)=min{S_no-island(R_t),S_island(R_t)} [SR: RT/HRT minimality, AMMZ §3.3; minimization step, separate from extremization]
derive extremization_then_minimization: QES extremization (qes_stationarity) and branch minimization (branch_selection) are distinct ordered operations from [qes_stationarity, branch_selection] | extremize S_gen per candidate first, then minimize across candidates (AMMZ(15))
derive page_condition: S_no(t_P)=S_isl(t_P) from [noisland_branch, island_value, branch_selection] | branches intersect; equivalent to Penington S_rad=A_hor/4G_N
claim pagetime_value: t_P≈6betaS^i_Bek/(πc) as LEADING estimate with footprint F1-F5 [C: DERIVED-NOT-FROZEN; branch-reconstruction §3]
  ? ?OB-7: close footprint F1 (O(1) evaporation-during-growth), F2 (UV const), F3 (island logs), F4 (T drift in growth), F5 (S_0) before any freeze
derive page_curve: S(R_t) rises as noisland_branch then falls tracking island_value after page_condition from [noisland_slope, island_value, page_condition, branch_selection] | min-selects growing branch early, decreasing BH-tracking branch late
claim island_rule_status: S(A)=Min Ext[S_eff+Area/4G_N] is an effective prescription, derived-from-RT/HRT only within holographic setup [SR/C: AMMZ(15), §4; conditional on RT/HRT correctness]
claim large_S0_regime: S_0 not negligible vs S^i_Bek [C: hierarchy violation case]
if large_S0_regime:
  then
    claim island_disfavored: 2S_0 island-creation price makes no-island branch win at all times [SR: AMMZ §3.3 2S_0 argument]
  else
    claim island_competes: island branch can dominate after page_condition under S_0<<S^i_Bek [C: regime_hierarchy]
claim entropy_route_fallback: G(eta) machinery is the documented fallback if holographic specialization drops [C: baseline §10]
  ? ?OB-8followup: state/channel/eta hypotheses owed if the fallback is ever activated

# ============ PART C — Shared R1/R4 interoperability ============

derive link_action_eqs: metric_eqs and curvature_constraint used across Parts A/B come from the single jt_action, not separate assumptions from [jt_action] | one variational origin matching AEMM(9),(11a-c) (BR-01 rebuilt)
derive link_anomaly_stress: stress_forms plus trace_anomaly jointly determine chiral_flux through ward_identity from [trace_anomaly, ward_identity, stress_forms] | anomaly-transformation route AP(5.1)-(5.2); independent-inputs structure (BR-02/03 rebuilt; Polyakov-explicit leg suspended per polyakov_explicit)
derive link_state_stress: shock_stress is the stress-energy OF quench_state, fixed mutually with no independent choice from [quench_state, shock_stress, conformal_matter] | AEMM §3.2 chain (43)-(46): quench map plus half-plane vacuum fix E_S and flux
derive link_state_entropy: holographic_entropy_route computes radiation entropy for the quench-evolved state on the quench geometry from [quench_state, holographic_entropy_route, radiation_region] | 3D-RT on quench background (AMMZ §3); state enters via geometry and w-map
  ? ?OB-10: 3D-RT extremal-surface computation imported wholesale from AMMZ §3; close by reconstructing the RT surface or keep import with fresh verification
derive link_geometry_factors: entropy expressions use conformal factors of the coords_poincare/coords_physical metrics with cutoff z=eps f'(u) from [coords_poincare, coords_physical, noisland_branch] | Weyl rule AEMM(53); dossier constant-eps cutoff rejected
derive link_dilaton_area: dilaton_term with bh_entropy_fn gives the one-endpoint area law summed per endpoint from [dilaton_term, bh_entropy_fn, endpoint_count_one] | AMMZ(10); Decision 2
derive link_matter_sgen: sgen_functional equals dilaton_term plus matter entropy on R∪I with AEMM Ren split from [sgen_functional, dilaton_term, holographic_entropy_route] | AMMZ(10) plus AEMM(69)-(73)
derive link_sgen_qes: qes_position extremizes sgen_functional through qes_stationarity from [sgen_functional, qes_stationarity, qes_position] | AMMZ(12)←Penington+AEMM §4 chain (?OB-9 carried)
derive link_branches_rad: page_curve gives S(R_t) via branch_selection over noisland/island branches meeting at page_condition from [page_curve, branch_selection, page_condition] | AMMZ §3.3 minimization
derive link_flux_mass: global_balance connects evap_law decay to bath flux through energy_balance from [energy_balance, evap_law, global_balance] | AEMM(22)+(25) structure (?OB-4 carried: full dM/dt line-by-line owed)
claim interop_complete: all ten prompt9 §4C connections established above or explicitly obligated [C: carriers OB-4, OB-5, OB-9, OB-10]
  ? interop_carriers: close OB-4 (dM/dt), OB-5 (shock-vs-smooth branch), OB-9 (QES re-derivation), OB-10 (RT reconstruction) or keep as explicit BLOCKED/? inputs to implementation
