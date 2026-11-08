# R4 Branch Reconstruction (Step 5, in-note derivation)

**Purpose:** independently reconstruct both entropy branches from corrected
baseline inputs and derive their intersection. The Page-time VALUE is
DERIVED-NOT-FROZEN per Decision 2 (neither 3βS/πc nor 6βS/πc is adopted as
a premise). All inputs below come from `working-baseline.md` §§6-7.

## 1. No-island branch

Premise (baseline §6.3, verified log form; S₀ term dropped per ?OB-6):
S_no(t) = (c/6)·ln[(β/πϵ)·sinh(πt/β)].

Late-time expansion (t ≫ β), line by line:
sinh(πt/β) = (1/2)e^{πt/β}(1 − e^{−2πt/β}) → (1/2)e^{πt/β} + exp. suppressed.
ln[(β/πϵ)·sinh(πt/β)] → πt/β + ln(β/2πϵ).
S_no(t) → (cπ/6β)·t + (c/6)·ln(β/2πϵ).
With β = 1/T_H: slope = πcT_H/6. The additive constant is UV-sensitive
(ϵ-dependent) and non-universal — RECORDED, not dropped silently (it enters
the intersection at subleading order).

Cross-check: AMMZ(14) gives S_Rad = (πc/6)∫T dt′ → same slope πcT/6 at
fixed T. AEMM(112) same. Three-way agreement on the slope. ✓

## 2. Island branch

Premise (baseline §6.5, Decision 2, AMMZ(13) via §3.1.2 shared surface):
S_isl(t) ≈ S_BH(T(t)) + logs, S_BH(T) = (φ₀ + φ_hor(T))/4G_N, one endpoint.
Evaporation law: T(t) = T₁·e^{−kt/2}, where the AMMZ↔AEMM rate
identification κ_AMMZ = k_AEMM is DERIVED (not assumed): AEMM(20)+(25) give
E₁e^{−ku} = (πφ̄_r/4G_N)T(t)², hence T(t) = T₁e^{−ku/2}, matching AMMZ's
T(t) ∼ T_i e^{−κt/2} form term by term.
"logs": O(c·log) matter + log-S_BH terms (AMMZ(13)); subleading vs
S_BH ∼ φ/G_N under the hierarchy. Dossier's O(c⁰) REJECTED (audit).

## 3. Intersection (Page condition)

Page definition (adopted, AMMZ §3.3; Penington Eq.(2) equivalent):
S_no(t_P) = S_isl(t_P).
Full equation:
(πcT_H/6)·t_P + (c/6)ln(β/2πϵ) + [no-island exp-suppressed] =
(φ₀ + φ_hor(T(t_P)))/4G_N + [O(c·log) island corrections].
Leading large-entropy estimate (S₀ ≪ S^i_Bek, drop const/logs, evaluate
S_BH at T₁): t_P ≈ 6βS^i_Bek/(πc).

ASSUMPTION FOOTPRINT (why the value does not freeze):
(F1) Slow-evaporation: S_BH(T(t_P)) ≈ S^i_Bek. Self-check: t_P·k ≈
(φ₀+φ_hor)/(2πTφ̄_r) ∼ O(1) + S₀-piece — t_P is ORDER the evaporation
timescale, so evaporation-during-growth is an O(1) correction, computed
nowhere yet.
(F2) UV constant (c/6)ln(β/2πϵ) dropped — shifts t_P at O(β/c × logs).
(F3) Island O(c·log) terms dropped.
(F4) Fixed-β no-island form (ignores T(t) drift inside the growth phase).
(F5) S₀ ≪ S^i_Bek (hierarchy; else AMMZ 2S₀-price changes branch structure).
Each Fn is a Phys-Pseudo `?`/CONDITIONAL carrier in Part B. Closing any Fn
requires a calculation owed to OpenCode's checklist, not an assumption.

## 4. Consistency cross-checks (pass)

(C1) Intersection precedes AMMZ(14) saturation: S^i_Bek < 2S^i_Bek ✓
(Page transition cuts the growth before asymptote — matches AMMZ §3.3).
(C2) Penington's Page definition (S_rad = A_hor/4G_N) is the same condition
in covariant form ✓. (C3) Dossier's 3βS/πc reproduces NOTHING in this
derivation (it combined the wrong slope with 1× island) — confirmed void;
the audit-note 6βS/πc reproduces as the LEADING estimate only, hence also
not freezable as an exact value. (C4) Monotonic decrease post-t_P follows
S_BH(T(t)) with falling T(t) ✓ (dossier's "saturates" language rejected).
