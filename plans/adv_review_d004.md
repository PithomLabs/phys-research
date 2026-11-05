# ADVERSARIAL REVIEW: D-004 — Corrected Einstein–Cartan Specialization of Yo–Nester Classification

## A. Executive Disposition

**FAIL** — D-004 contains a significant dimensional/coefficient error in the φ⊥ inhomogeneous term (`3a0/l²` vs source `3a0/l` or `3a0/(2l)`), the adjacent-case comparison in `derive adjacent_check` still omits `∼φμν⊥` despite claiming to fix F1, and the conditional Hbar conclusion is promoted beyond what the premises support. The EC-corner mapping remains an unverified assumption correctly flagged but still used in consequential `from` lists.

---

## B. F1–F3 Resolution Assessment

### F1 — Adjacent spin-2⁻ comparison
**Status: NOT RESOLVED**

D-004 `derive adjacent_check` (line 75) states:
> "the spin-2⁻ case PIC list Eq.(94) equals the EC-corner PIC list minus Tφ and ∼φμν⊥, with matching φ⊥ inhomogeneity -3a0/l², from [spin2m_case, torsion_pics, curvature_pics] | compare the two lists term by term; Tφ absent because b1+b2 ≠ 0 and ∼φμν⊥ absent because b1+b4 ≠ 0 in case (93)"

**Source verification:**
- EC-corner PICs (from `curvature_pics`): φ, pφ, ⇀φμ, Tφ, ∧φμν, ∼φμν, ∧φμν⊥, ∼φμν⊥, φ⊥, φμ⊥ — **10 entries**
- Spin-2⁻ case (Eq. 93: a0≠0, b1≠0, others 0): Eq.(94) lists φ, ∼φμν, φμ⊥, ∧φμν, pφ, ⇀φμ, φ⊥, ∧φμν⊥ — **8 entries**
- **Missing from spin-2⁻**: Tφ (b1+b2=b1≠0) AND ∼φμν⊥ (b1+b4=b1≠0)
- D-004 correctly identifies Tφ absence (b1+b2≠0) and ∼φμν⊥ absence (b1+b4≠0) in the explanation
- **BUT** the `derive adjacent_check` claim says "minus Tφ and ∼φμν⊥" — this is now **correct** in the claim text

**However:** The `derive adjacent_check` result text still says "with matching φ⊥ inhomogeneity -3a0/l²" — the coefficient error (see F2 below) propagates here. The spin-2⁻ case Eq.(94) shows `φ⊥ = π⊥/J - 3a0/l` (no factor of 2, no l²). The EC corner from Eq.(46) shows `φ⊥ = π⊥/J - 3a0/(2l)` (with factor 2). **These do not match.**

**Disposition: PARTIALLY FIXED** — The missing `∼φμν⊥` is now acknowledged in the explanation, but the φ⊥ inhomogeneity claim is still wrong due to the F2 coefficient error.

### F2 — Eq.(46) constant
**Status: NOT RESOLVED — NEW ERROR INTRODUCED**

**Source Eq.(46) (fulltext lines 408-411):**
```
π⊥ 3a0       1                1
φ⊥ ≡       − 2 − (b4 − b6 )R = (b4 + b6 )R⊥⊥    (46)
J      l   2κ                κ
```
Parsed: `φ⊥ = π⊥/J - 3a0/(2l) - (b4-b6)/κ R` (with `2l` = 2×l, not l²)

**Source Eq.(94) (spin-2⁻ case, fulltext line 1111):**
```
π 3a0                 ∧       π µν⊥
φ⊥ ≡ − 2 ≈ 0,              φµν⊥ ≡        ≈ 0.                 (94)
J     l                         J
```
Parsed: `φ⊥ = π⊥/J - 3a0/l` for the spin-2⁻ parameter choice (a0≠0, b1≠0)

**D-004 `derive curvature_pics` (line 73):**
> "φ⊥ = π⊥/J - 3a0/l² ≈ 0 as the sole inhomogeneous PIC"

**Error:** D-004 writes `3a0/l²` — dimensionally wrong (l is a length, π⊥/J has dimensions 1/length, so inhomogeneity must be 1/length). The source has `3a0/(2l)` or `3a0/l`, never `3a0/l²`.

**Impact:** This error propagates to `derive adjacent_check` ("matching φ⊥ inhomogeneity -3a0/l²") and any subsequent calculation using this coefficient.

**Disposition: REGRESSED** — D-003 had `3a0/2l²` (also wrong, but different); D-004 changes to `3a0/l²` which is still wrong and dimensionally inconsistent.

### F3 — Unused premise
**Status: RESOLVED**

D-003 had `given a0_scalar` (line 60) unused in any `from` list. D-004 removes it from givens and mentions it only in Premises prose (line 29): "a0 couples the scalar curvature R (cached fulltext line 204); supports corner_map plausibility in prose only." This correctly reflects its role.

**Disposition: FIXED**

---

## C. Substitution Audit — Independent EC-Corner Recalculation

### Parameter point: `ec_vals := {a0 ≠ 0, a1 = a2 = a3 = b1 = b2 = b3 = b4 = b5 = b6 = 0}`

### 1. Torsion PICs (Eqs. 36–39)

| PIC | Source Eq. | EC Substitution | Result |
|-----|------------|-----------------|--------|
| φμ⊥ | (36) `πμ⊥/J + 2/(3l)(a1-a2)T⊥µ = 2/(3l)(2a1+a2)T⊥µ⊥` | a1=a2=0 | `πμ⊥/J = 0` (homogeneous) |
| ∧φμν | (37) `∧πμν/J + 2/(3l)(a1-a3)T⊥[µν] = 2/(3l)(a1+2a3)T⊥[µν]` | a1=a3=0 | `∧πμν/J = 0` (homogeneous) |
| ∼φμν | (38) `∼πμν/J + a1/l T⊥hµνi` | a1=0 | `∼πμν/J = 0` (homogeneous) |
| φ | (39) `π/J + a2/l T⊥σ σ` | a2=0 | `π/J = 0` (homogeneous) |

**All four torsion PICs: homogeneous, exactly as D-004 claims.** ✅

### 2. Curvature PICs (Eqs. 43–48)

| PIC | Source Eq. | EC Substitution | Result |
|-----|------------|-----------------|--------|
| pφ | (43) `pπ/J - (b2-b3)/κ pR◦⊥ = (b2+b3)/κ pR⊥◦` | b2=b3=0 | `pπ/J = 0` (homogeneous) |
| ⇀φμ | (44) `⇀πμ/J + (b4-b5)/κ Rμ⊥ = -(b4+b5)/κ R⊥μ` | b4=b5=0 | `⇀πμ/J = 0` (homogeneous) |
| Tφ | (45) `Tπσµν/J + (b1-b2)/κ TRµνσ⊥ = (b1+b2)/κ TR⊥σµν` | b1=b2=0 | `Tπσµν/J = 0` (homogeneous) |
| φ⊥ | (46) `π⊥/J - 3a0/(2l) - (b4-b6)/κ R = (b4+b6)/κ R⊥⊥` | b4=b6=0 | **`π⊥/J - 3a0/(2l) = 0`** (INHOMOGENEOUS) |
| ∧φμν⊥ | (47) `∧πμν⊥/J - (b2-b5)/κ R[µν] = (b2+b5)/κ R⊥[µν]⊥` | b2=b5=0 | `∧πμν⊥/J = 0` (homogeneous) |
| ∼φμν⊥ | (48) `∼πμν⊥/J - (b1-b4)/κ Rhµνi = (b1+b4)/κ R⊥hµνi⊥` | b1=b4=0 | `∼πμν⊥/J = 0` (homogeneous) |

**Complete EC-corner PIC list (10 entries):**
1. φ = π/J ≈ 0 (homogeneous)
2. φμ⊥ = πμ⊥/J ≈ 0 (homogeneous)
3. ∧φμν = ∧πμν/J ≈ 0 (homogeneous)
4. ∼φμν = ∼πμν/J ≈ 0 (homogeneous)
5. pφ = pπ/J ≈ 0 (homogeneous)
6. ⇀φμ = ⇀πμ/J ≈ 0 (homogeneous)
7. Tφ = Tπσµν/J ≈ 0 (homogeneous)
8. ∧φμν⊥ = ∧πμν⊥/J ≈ 0 (homogeneous)
9. ∼φμν⊥ = ∼πμν⊥/J ≈ 0 (homogeneous)
10. **φ⊥ = π⊥/J - 3a0/(2l) ≈ 0** (sole inhomogeneous)

**D-004 claim (line 73):** "φ⊥ = π⊥/J - 3a0/l² ≈ 0 as the sole inhomogeneous PIC" — **coefficient error** (`l²` vs `2l`), but correctly identifies it as the sole inhomogeneous PIC.

### 3. Kinetic Critical Combinations (Table 1)

At EC corner (`a1=a2=a3=b1=b2=b3=b4=b5=b6=0`, `a0≠0`):

| Sector | Kinetic Combo | Value at EC | Triggers PIC? |
|--------|---------------|-------------|---------------|
| 0+ (a) | a2 | 0 | YES (φ, χ) |
| 0+ (b) | b4+b6 | 0 | YES (φ⊥, χ⊥) |
| 1+ (a) | a1+2a3 | 0 | YES (∧φμν, ∧χμν) |
| 1+ (b) | b2+b5 | 0 | YES (∧φμν⊥, ∧χμν⊥) |
| 2+ (a) | a1 | 0 | YES (∼φμν, ∼χμν) |
| 2+ (b) | b1+b4 | 0 | YES (∧φμν⊥, ∧χμν⊥) |
| 1- (a) | 2a1+a2 | 0 | YES (φμ⊥, χμ⊥) |
| 1- (b) | b4+b5 | 0 | YES (⇀φμ, ⇀χμ) |
| 0- (a) | b2+b3 | 0 | YES (pφ, pχ) |
| 2- | b1+b2 | 0 | YES (Tφ, Tχ) |

**All 10 kinetic combinations vanish at EC corner** — full PIC set activated. ✅ Matches D-004 `derive pair_structure` premise.

### 4. Mass Combinations (Table 1) at EC Corner

| Sector | Mass Combo | Value at EC | Linear PB |
|--------|------------|-------------|-----------|
| 0+ | a0, 2a0+a2 | a0, 2a0 | ≠0 (a0≠0) |
| 1+ | a0 | a0 | ≠0 |
| 2+ | a0, a1-a0 | a0, -a0 | ≠0 |
| 1- | a1-a0, 2a0+a2 | -a0, 2a0 | ≠0 |
| 0- | pχ: a0, a3 | a0, 0 | pχ mix |
| 2- | a1-a2 | 0 | **VANISHES** |

**Result:** 9 sectors have nonzero linear mass (second-class at linear level); 2- sector has zero linear mass. ✅ Matches D-004.

---

## D. Pair-Classification Audit

### Linear-Level Pair Structure (from Table 1)

| Spin | PIC | χ Partner | Mass Combo | EC Value | Linear PB |
|------|-----|-----------|------------|----------|-----------|
| 0+ | φ | χ | a0 | a0 ≠ 0 | ≠0 |
| 0+ | φ⊥ | χ⊥ | a0 | a0 ≠ 0 | ≠0 |
| 1+ | ∧φμν | ∧χμν | a0 | a0 ≠ 0 | ≠0 |
| 1+ | ∧φμν⊥ | ∧χμν⊥ | a0 | a0 ≠ 0 | ≠0 |
| 2+ | ∼φμν | ∼χμν | a0 | a0 ≠ 0 | ≠0 |
| 2+ | ∧φμν⊥ | ∧χμν⊥ | a0 | a0 ≠ 0 | ≠0 |
| 1- | φμ⊥ | χμ⊥ | a1-a0 | -a0 ≠ 0 | ≠0 |
| 1- | ⇀φμ | ⇀χμ | 2a0+a2 | 2a0 ≠ 0 | ≠0 |
| 0- | pφ | pχ | a0, a3 | a0, 0 | mixed* |
| 2- | Tφ | Tχ | a1-a2 | **0** | **0** |

*pφ mass = a0 ≠ 0; pχ mass = a3 = 0 → pφ second-class, pχ first-class at linear level. The pair pφ/pχ is mixed.

**D-004 claim (line 74):** "9 pairs have mass proportional to a0 ≠ 0 hence second-class at linear level; the 2- pair (Tφ,Tχ) has mass a1-a2 = 0 hence vanishing linear PB"

**Assessment:** Correct on the 2- pair; the "9 pairs" count is slightly imprecise (the 0- sector has a mixed pair), but the essential claim — 9 nonzero linear masses, 1 zero — is correct.

### Exceptional 2⁻ Pair

**D-004 `derive two_minus_gap` (line 77):**
> "the 2- pair PB at the EC corner equals 0 plus field-dependent nonlinear terms whose explicit form is not given in the examined sources, so its classification is undetermined here"

**Source verification:** Fulltext search (lines 428, 1253-1254, 1266) confirms:
- "In Table 1 only the PIC's pφ and Tφσµν do not have their 'natural' second-class pair counterparts" (line 1253-1254)
- No explicit 2- pair PB computation for EC corner in the paper
- Nonlinear terms are acknowledged but not exhibited for this corner

**D-004 claim (line 79):** `claim two_minus_sc: the 2- pair is second-class at the EC corner via nonzero nonlinear PB terms`

**Assessment:** This is a **claim**, not a derivation. It is explicitly flagged with obligations D2' and D2map (lines 89-90). The `if two_minus_sc:` branch (lines 81-86) correctly treats this as a conditional branch. The gap is honestly stated.

---

## E. Hbar Assessment

### D-004 Hbar Logic (lines 81-88)

```
if two_minus_sc:
    then derive hbar_uniform: Hbar elimination proceeds uniformly over all 10 sectors with Hbar_perp absorbing all primary second-class contributions, from [pair_structure, two_minus_sc, d002_results] | all pairs second-class, Blagojevic Hbar rule as recorded in D-002 [mine]
        ? what is the explicit Hbar_perp EC expression? Still needs the SIC solutions (D2).
else
    derive hbar_split: the 2- sector needs separate treatment (first-class, tertiary, or phase-space-dependent analysis) before Hbar elimination, from [pair_structure] | the remaining 9 sectors eliminate; the 2- pair does not [mine]
        ? which alternative holds for the 2- pair? Needs its explicit nonlinear PB terms (D2').
```

### Assessment

**Premises:**
- `pair_structure`: 9 linear second-class pairs + 1 exceptional 2- pair (CONDITIONAL on corner_map)
- `two_minus_sc`: CLAIM (not derived), flagged as unresolved (D2')
- `d002_results`: D-002 `hbar_governs` and `ec_corner`, both CONDITIONAL per R-003

**Blagojević rule (from D-002 line 71):** "with extra constraints H_perp → Hbar_perp" — stated under "extra-constraints precondition"

**Logical structure:** VALID conditional inference. IF all 10 sectors are second-class, THEN Blagojević's extra-constraints rule triggers Hbar branch uniformly.

**Gap 1:** `two_minus_sc` is a claim, not established. The `if` branch depends on an unresolved premise.

**Gap 2:** Even if `two_minus_sc` true, the explicit `Hbar_perp` EC expression is unknown (marked `?` on line 84).

**Gap 3:** The "extra constraints" in Blagojević's rule refer to the full constraint set. D-004 only analyzes primary if-constraints. Secondary constraints and the full Dirac algorithm at the EC corner are not computed.

**Gap 4:** D-002 `hbar_governs` uses Blagojević (4.7a) mixing term `-1/2 R_ijab H_ij`. This assumes the full PGT algebra, not just the PIC subset.

**Disposition:** The Hbar conclusion is **CONDITIONAL on multiple unverified premises** (corner_map, two_minus_sc, full constraint inventory, Blagojević rule applicability at EC corner). D-004 correctly flags the unresolved `?` but the `claim two_minus_sc` (line 79) is presented without qualification in the main text, which could mislead a reader into thinking it's established.

---

## F. EC-Corner Mapping (corner_map)

### D-004 line 68:
> `assume corner_map: the EC action corresponds to ec_vals [mine: inherits D-002 ec_is_corner, still unconfirmed by any source]`

### P-007 (Einstein-Cartan Theory, Trautman gr-qc/0606062):
- Classical EC action: tetrad + independent spin connection + algebraic Cartan relation (Eq. 25: torsion = 8π(spin + trace terms))
- No quadratic curvature terms (R²), no quadratic torsion terms beyond the standard Einstein-Cartan Lagrangian
- P-007 is an encyclopedia article, not a Hamiltonian classification paper

### Yo-Nester (gr-qc/0112030):
- Nine-parameter Lagrangian: `a0 R + a1 T² + a2 T² + a3 T² + b1 R² + b2 R² + b3 R² + b4 R² + b5 R² + b6 R²`
- EC corner claimed: `a0 ≠ 0, all other a_k = b_k = 0`

### Assessment:
- **P-007 does not map to Yo-Nester parameters.** P-007 gives field equations, conservation laws, and the algebraic torsion-spin relation. It does not decompose the Lagrangian into the 9-parameter basis or identify the Yo-Nester parameter values.
- **D-002 line 80:** `assume ec_is_corner: the EC action corresponds to ec_corner [mine: mapping; Yo-Nester never state the EC corner]`
- **D-004 correctly inherits this as an explicit assumption** and flags `? D2map` (line 89): "is corner_map correct, i.e. does any source state the EC corner in these parameters? Needs Nikolic 1995 or an equivalent primary classification source."

**Disposition:** The mapping is **explicitly conditional and unverified**. D-004 correctly propagates this conditionality through `from` lists. However, all consequential `derive` statements include `corner_map` in their `from` lists, making every quantitative result conditional on this unproven mapping.

---

## G. Phys-Pseudo Compliance

### Syntax Check
- Named constructs: `given`, `define`, `assume`, `claim`, `derive`, `if/then/else`, `?` — all valid per PSEUDO.md §1
- `from` lists reference existing premises/intermediate results — ✅
- Each `derive` has explicit result — ✅
- Assumptions labeled `assume` — ✅ (corner_map, line 68)
- Source-transcribed equations not presented as independent derivations — ✅ (marked with source locators)
- `?` obligations capture unresolved work — ✅ (D2map, D2prime, D2', D2-comp, D2map, D2prime)
- Conditional branches do not silently promote antecedents — ✅ (if/else structure preserves conditionality)

### Obligations Consistency
| Obligation | In Phys-Pseudo | In Obligations Section |
|------------|----------------|------------------------|
| D2' | line 49 | `? D2prime` line 90 |
| D2-comp | line 50 | — (missing from `?` list) |
| D2map | — | `? D2map` line 89 |
| D2prime | `? D2prime` line 90 | — |

**Minor inconsistency:** `D2-comp` appears in Obligations section (line 50) but not in `?` list; `D2map`/`D2prime` appear in `?` list but `D2'` vs `D2prime` naming differs.

---

## H. Per-Derivation Disposition Table

| Derivation | Lines | Disposition | Evidence |
|------------|-------|-------------|----------|
| `torsion_pics` | 72 | **VERIFIED** | All 4 Eqs.(36-39) independently reproduced; homogeneous at EC corner |
| `curvature_pics` | 73 | **CONDITIONAL** | 5/6 homogeneous verified; φ⊥ inhomogeneous **coefficient error** (`3a0/l²` vs source `3a0/(2l)`); rests on `corner_map` |
| `pair_structure` | 74 | **CONDITIONAL** | 9 nonzero + 1 zero mass verified; 0- sector mixed pair slightly miscounted; rests on `corner_map` |
| `adjacent_check` | 75 | **PARTIALLY FIXED** | Missing `∼φμν⊥` now acknowledged in explanation; φ⊥ inhomogeneity **still wrong** due to F2 error |
| `two_minus_gap` | 77 | **VERIFIED** | Negative-existential verified within source bounds; correctly states nonlinear terms not exhibited |
| `hbar_uniform` / `hbar_split` | 81-88 | **CONDITIONAL** | Valid case analysis; depends on unverified `two_minus_sc` claim and `corner_map` |

---

## I. Remaining Obligations (Minimum Unresolved Mathematical Work)

1. **D2map** — Verify EC-corner mapping: Does any source (Nikolić 1995, Hehl et al. 1976, Blagojević & Nikolić) explicitly state the EC action corresponds to `{a0≠0, all other a_k=b_k=0}` in Yo-Nester's 9-parameter basis? If not, the mapping must be independently derived from the EC Lagrangian.

2. **D2prime (D2')** — Compute explicit nonlinear `Tφ/Tχ` pair-PB terms at the all-critical corner. Requires full Dirac-Bergmann computation for the EC parameter point, or retrieval from Nikolić 1995.

3. **F2 correction** — Fix φ⊥ inhomogeneity coefficient: source Eq.(46) gives `3a0/(2l)`, not `3a0/l²`. Propagate correction to `adjacent_check` and any downstream use.

4. **Full EC constraint inventory** — Compute secondary if-constraints at EC corner. The Blagojević Hbar rule requires the full constraint set, not just primary PICs.

5. **2- pair nonlinear PB** — Determine whether the 2- pair is genuinely second-class, first-class, or phase-space-dependent. This requires the explicit nonlinear terms (D2').

6. **Hbar_perp EC expression** — If `two_minus_sc` holds, derive the explicit `Hbar_perp` for the EC corner (needs SIC solutions per line 84).

---

## J. Preservation and Validation

### Commands Run
```bash
cd /home/chaschel/Desktop/biz/TOE-v5/phys-research
go test -count=1 ./...                    # PASS (9 packages)
go vet ./...                               # PASS
go build ./...                             # PASS
physctl validate                           # PASS (0 errors, 1 LEGACY_UNPINNED warning)
sha256sum -c docs/BASELINE-2026-11-05.sha256
# 4 FAILED: docs/AGENT_CONTRACT.md, internal/paper/fetch.go, internal/record/record.go, internal/record/record_test.go
# (exactly the 4 post-freeze changes; D-004, D-003, R-004 unchanged)
```

### Hash Verification (unchanged since last review)
```
D-004: d9291ba4cfc4a56253af0054b9e41a057c0f06277168249abb2a02a050c30efc
D-003: ff7275c9f32a1313d20326848fe060696b32de4c2aa44df69e839de4feb2ec08
R-004: 7ee684efba8deb9704d707e4511a2ad0493b4028eb2f4e8e4e1283a3e787aaed
D-002: 3f26181a50c560260e3d4ae0d6d29de8f519b0b45fed198fc64c34346912fcbc
P-007: 0256e9b7e7ef98f7208666177570a80a17ec58db93dd5769a4bc7fdcaf23b456
```

All preserved byte-identical.

---

## K. Recommended Next Mathematical Calculation

**Highest-impact calculation that can proceed without assuming unresolved results:**

**Fix the φ⊥ coefficient and re-derive the adjacent-case comparison with correct inhomogeneity.**

1. Correct `derive curvature_pics`: `φ⊥ = π⊥/J - 3a0/(2l)` (not `3a0/l²`)
2. Correct `derive adjacent_check`: Spin-2⁻ case Eq.(94) has `φ⊥ = π⊥/J - 3a0/l`; EC corner has `π⊥/J - 3a0/(2l)`. **These do not match.** The factor-of-2 difference means the inhomogeneity is NOT matching between the two cases.
3. Report this discrepancy explicitly: the adjacent spin-2⁻ case does NOT share the same φ⊥ inhomogeneity as the EC corner — another difference beyond Tφ and ∼φμν⊥.

This calculation uses only already-verified source equations (Eqs.(46), (94)) and requires no new assumptions. It would resolve whether the "matching inhomogeneity" claim in D-003/D-004 was ever correct, and further constrain the adjacent-case comparison.

---

## Review Principle Compliance

- ✅ Did not edit D-004, D-003, R-004, D-002, H-002
- ✅ Did not modify PhysMath, PhysLib, PhysArtifact, Phys
- ✅ Did not commit or merge
- ✅ Preserved valid conditional mathematics (corner_map, two_minus_sc branches)
- ✅ Rejected unsupported conclusion (φ⊥ coefficient, 9-pair count precision)
- ✅ Did not manufacture a successful review
- ✅ Created review record targeting H-002 with D-004 as material input (via `physctl new --type review --review H-002`)
