# Phase A Ledger (draft §4 of final ledger) — Action + conventions

Canonical locators: JHEP printed pages. Evidence classes per prompt §7.

## A.1 Verified action terms (Source results, JHEP §3)

**A-I_grav (3.1, Jp.5).** I_grav = (1/2π)∫d²x√−g e^{−2φ}[R + 4(∇φ)² + 4λ²].
Covariant. Symbols: φ dilaton; λ characteristic length scale; R 2D Ricci scalar.
λ = 1 by 2D coordinate rescaling (text, Jp.5); λ restored only where stated
(3.7–3.9, 6.1). Variation wrt g_μν → (3.12); wrt φ → dilaton e.o.m.
(unaffected by matter coupling); in Kruskal gauge → (3.4)/(3.13).

**A-I_Q (3.28, Jp.9).** I_Q = −(c/12π)∫dx⁺dx⁻ ∂₊ρ∂₋ρ. Conformal-gauge form of
the nonlocal Polyakov term; nonlocality enters via t± in (3.23) (fn.5).
Dominant over dilaton-gravity loops for c ≫ 24 (text). Variation → anomaly-fed
stress tensor (3.23). Source result.

**A-I_RST (3.29, Jp.9–10).** I_RST = +(c/48π)∫d²x√−g φR. Covariant, local.
"Allowed by general covariance, does not disturb the classical (ε→0) limit";
enters at same order as I_Q. Defining model choice (follows [15] = original RST),
not derived. Variation → Ω-linearized equations (3.31). Source result.

**Relative sign/normalization (prompt §A audit point):** Polyakov −c/12π
(conformal-gauge) vs RST +c/48π (covariant): opposite signs, 4:1 magnitude
ratio in the forms printed. No G_N appears anywhere in GSST §3.

## A.2 Matter sector (Assumption, §3.1 Jp.6–7)

Strongly coupled holographic 2D CFT, large c, AdS₃ dual; Brown–Henneaux
c = 3L₃/2G₃ (3.10, Jp.7). "Important technical assumption" enabling von Neumann
entropy evaluation; "does not affect the gravitational sector" beyond T_μν.
Original-CGHS N free scalars explicitly NOT adopted. N = c, κ = c/12 only as
comparison identifications to [14,15] (fn.3). NO explicit matter Lagrangian is
given in GSST — per Correction A, none is invented here.

## A.3 Collapse pulse (Source/BC, Jp.7–8)

"Thin shell of infalling matter energy" = compact-support T₊₊(x⁺) + idealised
null-shock patching (3.16); full evolution cited to [15] (Imported). It is a
stress-energy source + matching condition, NOT an action term. S_shock-style
action terms have no source support.

## A.4 Convention pins (Phase A)

- T^μν = −(4π/√−g)δS/δg_μν (fn.4, Jp.7); ×2 vs [14]. (OQ-2 input.)
- ℳ = λM/π (3.7, Jp.6); S = 2M = 2πℳ/λ (3.9); T = λ/2π (3.8).
  NOTE: Claude's "M = λM_phys/π" is INVERTED vs source; corrected here.
- ε ≡ c/48M ≪ 1 (3.11); working order: leading non-trivial ε (text Jp.7).

## A.5 Gemini action adjudication (prompt §7, first dispute)

| Gemini claim (gem-05/06) | GSST source | Ruling |
|---|---|---|
| S_CGHS = (1/2π)∫e^{−2φ}[R+4(∇φ)²+4λ²] | (3.1) identical | Source-faithful |
| S_RST = −(κ/8π)∫φR, κ=cG_N/12 | (3.29) +(c/48π)∫φR | CONTRADICTED (sign; coeff needs G_N=−2) |
| S_PL = −(κ/16π)∫R□⁻¹R | (3.28) gauge form only; no G_N in §3 | Inconsistent (no single G_N reconciles with S_RST; G_N extraneous) |
| S_shock action term | shock is T-source + patching (3.16) | No source support (Correction A violation) |
| Matter = c free scalars | §3.1 holographic CFT "instead" | CONTRADICTED |
| κ = cG_N/12 | fn.3: κ = c/12 | CONTRADICTED (extra G_N) |
| Ω̃ = e^{−2φ}+(κ/2)φ; λ-form solution, no εlog | (3.30) Ω = e^{−2φ}+(c/24)φ; (3.35) with −εlog | Different object; matches only at unstated G_N=1 |
| "Exact" solution/language | Leading non-trivial ε (text Jp.7) | Overclaim |

Quarantine upheld: only S_CGHS may be cited from Gemini; all else excluded
from the canonical ledger.
