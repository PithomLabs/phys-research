# Phase C Ledger (draft §6 of final ledger) — Collapse, flux, energy balance

## C.1 Geometry (Source results)

- Vacuum (3.32, Jp.10): Ω = −x⁺x⁻ − (c/48)log(−x⁺x⁻). SUBSTITUTION CHECKED
  (this investigation): ∂₊∂₋Ω = −1 ✓ (log mixed derivative vanishes);
  −∂±²Ω = (c/24)t± gives t±(x±) = −1/[2(x±)²] ≠ 0, consistent with text
  ("t±(x±)≠0... metric not manifestly flat in Kruskal").
- Shock: thin-shell T₊₊ compact support + patching (3.16); full evolution cited
  to [15] (Imported). GSST uses v>1 exterior only.
- Exterior (3.35, Jp.10): ds² = −Me^{2φ}dvdu,
  Ω = M[1−v(u+1) − εlog(−Mvu)]. SUBSTITUTION CHECKED: via (3.17),
  ∂₊∂₋ = (1/M)∂_v∂_u; ∂_v∂_uΩ = −M (log term splits, mixed derivative zero)
  ⟹ ∂₊∂₋Ω = −1 ✓. Second eq. of (3.31) fixes nonzero t± (consistent with flux).
- Strong-coupling boundary: Ω = Ω_crit (3.30 text, Jp.10); spacetime ends there
  (analogue of radial origin). Endpoint beyond semiclassical control (see §6/Discussion).
- Horizons: classical bifurcate horizon vu=0 (3.6 text, Jp.6). Evaporating-case
  apparent-horizon location NOT given as an equation in §3 — INFERRED quantities
  only; the entropy calculation uses near-horizon QES (Phase E), not an
  explicit horizon formula. Recorded as limitation, not gap.

## C.2 Energy accounting — OQ-2 (Independently reconstructed; closes OQ-2)

1. Flux component: T_−−(σ⁻) (3.27, Jp.9); late value c/24 per σ⁻. [Source]
2. Time: asymptotic t = (σ⁺+σ⁻)/2 (3.21); at ℐ⁺ (fixed σ⁺): dσ⁻/dt = 2. [Derived]
3. Physical convention: T^μν = −(4π/√−g)δS/δg_μν (fn.4) ⟹ physical flux
   density = T_−−/(4π) per σ⁻. Sign: paper's stated mass-loss direction;
   convention-sign audit carried as OQ-3 (see Phase D).
4. Mass loss: dℳ/dt = −(c/24)·2/(4π) = −c/48π (λ=1);
   dM/dt = (π/λ)dℳ/dt = −c/48. [Derived]
5. Lifetime: M/|dM/dt| = 48M/c = 1/ε ✓ = paper's t_lifetime (text, Jp.9).
   Loop CLOSES (magnitudes exact).
6. First law: S = 2πℳ/λ (3.9), T = λ/2π (3.8) ⟹ dS = −dℳ_lost/T EXACT in
   physical ℳ (matches §6 text). dS/dt = 2dM/dt = −c/24 (slope magnitude for
   Phase E cross-check). Radiation/b.h. entropy-rate ratio 2:1 (text, §6)
   consistent with no-island +c/12 vs island −c/24 (slopes verified Phase E).

OQ-2 DISPOSITION: CLOSED — all six steps shown with explicit conversion factors
(2 from dσ⁻/dt, 1/4π from T-definition, λ/π from (3.7)). Residual: OQ-3
(independent sign-convention audit) stays OPEN; it questions my verification,
not the paper's self-consistent chain (flux→lifetime→Page time all use +(3.27)).

## C.3 Claude mass-normalization concern (adjudication input)

Claude's "M = λM_phys/π" is inverted vs (3.7) ℳ = λM/π — but Claude's Bridge 3
numerics (lifetime 1/ε, |dS/dt| = c/24, crossing at lifetime/3) use only M and
are UNAFFECTED by the inversion. The inversion matters solely for physical-mass
statements. Full Bridge 3 ruling in Phase F; concern partially sustained (units
needed this reconstruction) but numerically harmless as used by Claude.
