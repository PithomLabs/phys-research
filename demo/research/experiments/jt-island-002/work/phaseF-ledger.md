# Phase F Ledger (draft §§9–11 of final ledger) — QES, Page time, bridges, RCs

## F.1 QES reconstruction (10 steps, prompt §F)

1. S_gen before extremization: (5.12) ✓ (Phase E).
2. Independent coords: (v_I, u_I) ✓.
3. Stationarity: (5.13–5.14) ✓ (paper writes "Extremizing (5.13)" meaning (5.12)
   — trivial typo, equations themselves correct).
4. Approximation: log(v_A/v_I)≫1 → drop last term of (5.13) → (5.15) ✓.
5. QES position: (5.16) u_I=−1+ε/v_I+O(ε²), near-horizon, inside horizon;
   second root near singularity discarded as unphysical (text) ✓.
6. Validity: (5.19) log(v_A/v_I)≈2σ_A+log(1/ε)≫1 SELF-CONSISTENT ✓ (large
   anchor + small ε both push the same direction); v_I>1 ⟺ t_A−σ_A≳log(1/ε);
   island valid from scrambling time after first radiation ✓. Classical-ρ at
   island: valid most of lifetime, fails near endpoint (text Jp.17) ✓.
7. Island entropy: (5.20) 2M−(c/24)(t_A−σ_A)+… ✓ (time piece grows to O(M):
   paper explicit).
8. No-island entropy: (5.21) (c/12)(t_A−σ_A)+… ✓.
9. Branch rule: min (text Jp.16: "the one that gives a smaller value") ✓;
   extremization-then-minimization ordering explicit in §5.2 procedure ✓.
10. Crossing: (5.22) t_Page=1/3ε=16M/c for t_A−σ_A>1/3ε ✓.
    INDEPENDENT ALGEBRA CHECK (this investigation): 2M−(c/24)τ=(c/12)τ ⟹
    2M=(c/8)τ ⟹ τ=16M/c; 1/3ε=16M/c since ε=c/48M ✓ EXACT.
    Retarded-time offset: t_Page is in t_A−σ_A (anchor-shifted), NOT absolute t.

Status: leading-order semiclassical Page transition (paper: abstract, §6).
No unitarity proof (paper claims none); endpoint beyond control (text Jp.19).

## F.2 Bridge verdicts (ten-field condensates; full contracts in final §10)

**Bridge 1 — State/frame: CONDITIONAL (import [32]).**
Source: R1 t± + (3.26) + t_−(ω⁻)=0. Target: (4.7) vacuum-frame evaluation.
Mapping: (ω⁺,ω⁻) instruction (Jp.15) + coherent-state license [32].
Assumptions: [32] (S_vN coherent = S_vN vacuum); t_−(ω⁻)=0 (source).
Preserved: vacuum-form RT applicability. Lost: non-coherent-state generality;
t₊(ω⁺)≠0 asymmetry recorded. Regime: ε≪1, post-scrambling. Obligation: none
open except OQ-3's sign caveat (does not touch the frame instruction).
Failure: non-coherent incoming state. Evidence: Jp.15 text + (4.7) condition.

**Bridge 2 — Geometry: CONDITIONAL (approximation window).**
Source: (3.35) Ω + classical ρ. Target: (5.12) area + e^{ρ} factors.
Mapping: direct substitution (paper states). Assumptions: classical ρ at A
(always) and I (most of lifetime). Preserved: leading logs. Lost: subleading
logs, O(ε²), endpoint domain. Regime: M≫c, pre-endpoint. Obligation: none.
Failure: island leaving classical-ρ domain (late times). Evidence: Jp.16–17.

**Bridge 3 — Energy/entropy: PASS (leading order).**
Source: dM/dt=−c/48, lifetime 1/ε (Phase C). Target: island slope −c/24 (5.20),
no-island +c/12 (5.21), ratio 1/3 (5.22). Mapping: first law (S=2M; ℳ-form
exact) + (6.1) 2:1 reasoning. Check: 2M=(c/12+c/24)t ⟹ t=16M/c=lifetime/3 ✓
from INDEPENDENT sides (flux-side lifetime vs entropy-side slopes; no circular
use of (5.22) as premise — (5.22) re-derived here from (5.20)+(5.21)).
Assumptions: leading order; OQ-3 sign caveat attached (magnitudes exact).
Preserved: numerical cross-identity. Lost: O(ε) corrections. Failure: any
normalization-factor change (all pinned Phase C). Genuine cross-check, not
shared definition.

## F.3 RC1–RC3 (consolidated; per-phase records in phaseX-checks.md)

RC1: 2D dilaton gravity (dynamics), variational calculus (eqs + extremization),
Polyakov/anomaly (backreaction), 2D CFT (state data t±, Schwarzian), holographic
RT/AdS₃ (matter entropy, imported), QES theory (saddles), ODEs (saddle
conditions). Regimes: ε≪1, M≫c≫(24), post-scrambling, pre-endpoint. Excluded:
endpoint microphysics, 4D gravity, greybody factors (absent in 2D per §6).
RC2: R1∩R4 = {Ω, ρ, t±+frame, M, flux}. Removal: no anomaly → no flux/lifetime,
island has no competitor slope; no RST → no Ω-linearity/exact solution; no
vacuum frame → (4.7) undefined; no QRT → R4 undefined. Necessity only;
minimality NOT claimed.
RC3: Bridges 1–3 add information (frame instruction + license; substitution
fixing area; numerical identity). Imported (QRT, RT, [32], anomaly basis, [15]
interior, [31] method) vs derived (chains, slopes, crossing) vs approximated
(leading ε, classical ρ, fixed reference) explicitly separated throughout.
