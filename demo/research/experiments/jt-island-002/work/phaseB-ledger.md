# Phase B Ledger (draft §5 of final ledger) — R1 field equations → anomaly → stress

## B.1 Classical equations (Source results)

- (3.4, Jp.6): vacuum/Kruskal-gauge e.o.m. from (3.1); (3.5) general solution.
- (3.12, Jp.7): covariant metric eq 4e^{−2φ}[∇∇φ − g(□φ−(∇φ)²+1)] = T_μν.
  Dilaton e.o.m. unaffected by matter (text). Status: source equations;
  variation from (3.1) independently repeatable (not repeated here — obligation
  recorded for Muse formalization, not a gap in the source).
- (3.13, Jp.7): components −∂₊∂₋e^{−2φ}=1, −2∂±²e^{−2φ}=T±±.

## B.2 Anomaly and conservation (kept distinct)

- (3.22, Jp.8): ⟨T^μ_μ⟩ = (c/12)R. Source equation; CFT-theorem basis imported
  (not derived in GSST). Applies to the holographic CFT (c) — the paper's §3.1
  assumption carries in.
- Conservation: invoked as "the continuity equation expressing
  energy-momentum conservation," integrated per Christensen–Fulling [31] (text,
  Jp.8). GSST prints no standalone ∇^μT_μν=0 or Ward identity; absence of a
  diffeomorphism anomaly is an implicit assumption. Status: imported identity +
  implicit assumption (flagged, not hidden).
- (3.23, Jp.8): ⟨T₊₋⟩=−(c/6)∂₊∂₋ρ, ⟨T±±⟩=(c/12)(2∂±²ρ−2(∂±ρ)²−t±).
  INDEPENDENTLY RECONSTRUCTED CHECK (this investigation): in
  ds²=−e^{2ρ}dx⁺dx⁻, R=8e^{−2ρ}∂₊∂₋ρ and T^μ_μ=−4e^{−2ρ}T₊₋ give
  T₊₋=−(c/6)∂₊∂₋ρ from (3.22); conservation
  ∂∓T±±+∂±T₊₋−2T₊₋∂±ρ=0 then forces ∂₋[T₊₊−(c/12)(2∂₊²ρ−2(∂₊ρ)²)]=0,
  i.e. the bracket equals −(c/12)t₊(x⁺). ✓ Coefficients consistent within
  GSST's T-normalization (fn.4); cross-literature use needs the ×2 factor.
- t±(x±): integration functions = state/boundary data, not derived quantities.
  Fixed by physical BCs per configuration (vacuum, eternal, collapse).

## B.3 Transformations and flux (Source results + paper derivation)

- (3.24–3.26, Jp.9): ρ, T±± (Schwarzian), t± transformation laws. Verified on
  render; standard 2D CFT form adapted to GSST's T-normalization.
- (3.27, Jp.9): T_−−(σ⁻) = −(c/12)t_−(σ⁻) = (c/24)(1−1/(1+e^{σ⁻})²), derived in
  paper from t_−(ω⁻)=0 via (3.26). Re-derivation owed in Phase D (show the map).
- t_lifetime = 1/ε at leading order (text, Jp.9). Stated, not derived there;
  energy-accounting derivation owed in Phase C (OQ-2).

## B.4 Semiclassical Ω equations (Source results)

- (3.31, Jp.10): ∂₊∂₋Ω+1=0, −∂±²Ω=(c/24)t±, with t± the same boundary functions.
  Linear in Ω — the RST solvability payoff. Follows from I_grav+I_Q+I_RST
  variation (paper asserts; substitution check: (3.32)/(3.35) satisfy (3.31) —
  recorded as Muse-level verification, quick hand-check in Phase C).
- (3.32) vacuum; (3.33)+(3.34) eternal with t±(σ±)=1/2 thermal gas at T=1/2π
  (verifies temperature consistency: matches (3.8) at λ=1 ✓).
