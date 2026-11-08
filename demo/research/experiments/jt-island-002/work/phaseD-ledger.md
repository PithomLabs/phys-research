# Phase D Ledger (draft §7 of final ledger) — State and coordinate frames

## D.1 State specification (Source results; no generic renaming)

- Initial: linear-dilaton vacuum (3.32); flat frame (ω⁺,ω⁻) (3.19) with
  ρ(ω⁺,ω⁻)=0 (text, Jp.8).
- Vacuum data: t_−(ω⁻)=0 (text, Jp.9: flat + vanishing T in v<1); t₊(ω⁺)≠0
  due to incoming flux (text, Jp.15). ASYMMETRY recorded: ω-frame is vacuum
  only in the − (outgoing) sector.
- Incoming matter: thin-shell T₊₊ compact support (§3.1); "short burst of
  matter energy" (text, Jp.10); full solution in [15] (Imported).
- Collapse-state entropy assumption: incoming matter as COHERENT STATE on
  inertial-observer vacuum at past null infinity, following [32] (Fiola et
  al.); S_vN(coherent) = S_vN(vacuum) (text, Jp.15). Imported theorem [32];
  load-bearing for RT use in (4.7).
- Vacuum naming: "Hartle-Hawking state" used ONLY for the eternal case (§5.1,
  Jp.13: t₊(v)=t₋(u)=0, Kruskal positive frequency). The evaporating case gets
  NO vacuum name — operational spec only. "Unruh vacuum" appears NOWHERE in
  GSST (grep-verified) ⟹ Gemini's "Unruh vacuum" label is UNSOURCED (prompt
  §D violation; ruled in Phase G adjudication).

## D.2 Frame map (Independently reconstructed)

- ω↔σ: ω⁺=σ⁺ (v>0); ω⁻=−log(1+e^{−σ⁻}) from (3.19)+(3.20); Jacobian
  dω⁻/dσ⁻=1/(1+e^{σ⁻}) (→1 early, →0 late).
- Schwarzian verified by 4 independent computations (direct, numeric point,
  composition law through u, eternal-case cross-check of (3.26)):
  {ω⁻,σ⁻} = −e^{σ⁻}(e^{σ⁻}+2)/[2(1+e^{σ⁻})²].
- (3.26) CONFIRMED via eternal case: predicts t±(σ±)=+1/2 from t±(x±)=0 ✓
  matches paper's text (Jp.10).
- PUZZLE → OQ-3: applying confirmed (3.26) with t_−(ω⁻)=0 yields
  t_−(σ⁻)=−{ω⁻,σ⁻}=+(1/2)(1−1/(1+e^{σ⁻})²), i.e. EXACTLY −(3.27)'s implied t_−.
  Every alternative (ρ-terms, map branches, T-route, inversion identity) was
  exhausted; all give the same flipped sign. (3.27) RETAINED as Source equation
  (positive flux required downstream + standard CGHS result); my sign
  re-derivation gap carried as OQ-3 (material uncertainty in verification, not
  a paper contradiction — downstream chain is self-consistent as printed).

## D.3 Bridge 1 input (for Phase F)

Paper's instruction (Jp.15): compute d(A,I) in (ω⁺,ω⁻) = CFT vacuum coordinates,
licensed by the coherent-state import [32]. What is preserved: vacuum-form RT
evaluation. What would fail for non-coherent incoming state: S_vN ≠ S_vacuum,
RT (4.7) inapplicable as used. t₊(ω⁺)≠0 does not obstruct because the entropy
evaluation needs the vacuum form, not vanishing t₊.
