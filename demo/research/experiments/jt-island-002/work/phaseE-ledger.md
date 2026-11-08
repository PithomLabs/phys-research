# Phase E Ledger (draft §8 of final ledger) — Generalized entropy

## E.1 Regions and anchor (Source results)

- Radiation region R: exterior of timelike anchor curve (Figs.2,5; §4–5).
- Anchor curve: constant Ω, Ω_A ≫ M, far outside hole (text, Jp.11). Drift in
  (3.20) coords negligible on lifetime (text). Eternal: anchor per side + mirror
  anchor (v,u)^m=(u,v) (text, Jp.13).
- S_AI: spacelike surface A↔I; result independent of choice (text, Jp.11).
  No-island: A↔fixed reference (v₀,u₀) on Ω=Ω_crit, u₀=−ε/v₀, v₀<1 (text, Jp.18);
  prescription-change independence claimed for leading late-time result.

## E.2 Functionals (Source results, Jp.11–12)

- (4.1): S_gen = Area(I)/4G_N + S_bulk[S_AI]. QRT adaptation (imported formula).
- (4.2): Area(I)/4G_N = 2(Ω(I)−Ω_crit). "Natural semiclassical extension";
  zero for I=∅ (empty-island normalization ✓ by construction).
- (4.3): S_bulk ≃ Length/4G₃ (≃: holographic step, following [11]).
- (4.4–4.5): y-frame (t±(y±)=0) + Weyl strip → pure Poincaré AdS₃.
- (4.6): regulated boundary z = δe^{−ρ(y⁺,y⁻)}; δ = UV cutoff.
- (4.7): S_bulk = (c/6)log[d(A,I)²e^{ρ(A)}e^{ρ(I)}]|_{t±=0}; d in flat
  ds²=−dy⁺dy⁻; UV cutoff dropped as additive constant (text).
  NOTE the +ρ(A)+ρ(I) signs (contrast Gemini's inverted form).

## E.3 Endpoint taxonomy (verified)

- Gravitational island endpoints (evaporating): ONE point (v_I,u_I); single
  area term 2(Ω(I)−Ω_crit) in (5.12). No S₀ anywhere in GSST (grep-confirmed
  absence of topological term; 2D CGHS/RST has none).
- Matter interval endpoints: A+I (island config); A+(v₀,u₀) (no-island).
- Mirror endpoints: eternal case only (A^m, island per side).
- Reference endpoint: (v₀,u₀) regulator for no-island; dependence subleading
  (text, Jp.18).

## E.4 Branches (Source results, Jp.16–18)

- (5.12): S_gen^island = 2M(1−v_I(1+u_I)−εlog(−Mv_Iu_I))
  + (c/12)log[(log(v_A/v_I)log(u_A/u_I))²·v_Iu_I/((1−v_A(1+u_A))(1−v_I(1+u_I)))].
  Area via (3.35); distance d(A,I)² in ω coords expressed in (v,u).
  Assumptions (text, Jp.17): island outside shell; classical ρ at A (always)
  and at I (most of lifetime; fails near endpoint).
- (5.21): S_gen^no-island = (c/12)(t_A−σ_A) + … (log corrections; v₀-free at
  leading order). Valid from scrambling time onwards.

## E.5 Cutoff and anchor sensitivity

- UV δ → additive constant, dropped (4.7 text). Physical-regulator independent
  at leading order; coordinate-frame dependence carried by e^{ρ} factors.
- Anchor: Page time expressed in RETARDED time t_A−σ_A (5.20–5.22) — explicitly
  anchor-shifted, not absolute physical time. Scrambling t_s=log(1/4ε) (5.11).
- t_A−σ_A≈0 = first Hawking radiation through anchor (from (3.27), text Jp.17).

## E.6 Gemini entropy adjudication (prompt §7)

| Gemini claim | GSST source | Ruling |
|---|---|---|
| S_grav = 2e^{−2φ}−(cℏ/12)φ | (4.2): 2e^{−2φ}+(c/12)φ−2Ω_crit | CONTRADICTED (φ-term sign; missing −Ω_crit) |
| S_matter = (c/6)ln[d²/(εε e^{ρ}e^{ρ})] | (4.7): (c/6)log[d²e^{ρ(A)}e^{ρ(I)}] | CONTRADICTED (inverted ρ dependence) |
| No-island ≈ (cℏλ²/12)t | (5.21): (c/12)(t_A−σ_A)+… | Slope value coincides (ℏ=λ=1); structure missing (no anchor shift, no reference point, no logs) |
| One endpoint, just inside horizon | (5.12) single point; text Jp.17–18 | Count/location AGREE |
| QES x⁻_I=x⁻_H−O(e^{−2πt/β}) | (5.13–5.17), no β anywhere | UNSOURCED generic ansatz; no extremization shown |
| S_island = 2M₀−(cℏλ²/24π)t | (5.20): 2M−(c/24)(t_A−σ_A) | CONTRADICTED (π in denominator) |
| "Exact Page time"; 12-vs-24π algebra ⟹ 1/3 | (5.22) leading-order; consistent-coefficient algebra | CONTRADICTED (internally inconsistent algebra + exactness overclaim) |
| "Unitary Page curve... to zero" | Leading-order; endpoint beyond control (§6) | CONTRADICTED (unitarity + endpoint overclaims) |
