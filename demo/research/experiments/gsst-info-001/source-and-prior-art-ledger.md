# Source and Prior-Art Ledger

Format per source: ID | full citation | DOI/arXiv | version/date | access status | exact locators | claim supported | assumptions/regime | limitations/competing results.
Rule: pinned file + hash in `sources/` for every consumed primary; inaccessible = [U], never quoted.

## Pinned primaries

- P-001 | Gautason et al., JHEP 05 (2020) 091 | 10.1007/JHEP05(2020)091, arXiv:2004.00598v2 | JHEP published | pinned (jt-island-002/sources/, hashes in its README) | §§3–6 | GSST model + entropy branches | leading order, ε≪1 | endpoint beyond control.
- P-002 | Callan–Giddings–Harvey–Strominger, PRD 45 (1992) R1005 | hep-th/9111056v1 | v1 | pinned (jt-island-002/sources/, SHA-256 3a239311…) | (16–19) | T-convention ×2, flux=energy identification | N free scalars (comparison only) | conventions differ from GSST (fn.4).

## Pinned this experiment

- P-003 | Fiola–Preskill–Strominger–Trivedi, "Black hole thermodynamics and information loss in two dimensions", PRD 50 (1994) 3987 | hep-th/9403137v1 | v1 | pinned (sources/FPST-hepth-9403137.pdf, SHA-256 45d45d9d86466a07f7f384638084cb3e3eb0e4b8a2a13dd9c9aaec876b25e2fb, 89pp) | coherent-state scope TBD (lines ~279, 661–672, 924 flagged) | [32] license scope = H6 crux | scope adjudication pending (Step 3).
- P-004 | Haocheng Zhong, "An algebraic description of the Page transition", JHEP 04(2026)160 | arXiv:2601.11363v1 | v1 16 Jan 2026 | pinned (sources/Zhong-2601.11363.pdf, SHA-256 128ecac875b4ba2a39f8bd25969621fbdac4226379913a7530a2865e7a8a4b83, 22pp) | Thm 1.1: finite-dim H=HA⊗HĀ + code space + isometric encoding → entropy formula ⟺ relative-entropy equality ⟺ operator reconstruction; Thm 1.2: relative-entropy saturation ⟺ exact channel reversibility; probe (1.10): Srel-difference vs δ across Page time; §4 type I/II algebraic description (content TBD) | hypotheses about GSST only | MAPPED §§1–2 (2026-11-08); §§3–5 read owed. Applicability flags (UNVERIFIED): finite-dim + tensor factorization + code space + isometry have NO established GSST counterparts (type-III CFT!); §4 covers type I/II — type-III coverage unknown.
- P-005 | Haocheng Zhong, "Probing the Page transition via approximate QEC", JHEP 01(2025)086 | arXiv:2408.15104v3 | v3 Jan 2025 | pinned (sources/Zhong-2408.15104.pdf, SHA-256 295fa5fa3045e10469917d1ba2b9ff6dbe1de7dac79b3feef6d4dbe7318f3a3c, 29pp) | precursor: island formula in AQEC language; Page transition as AQEC property | hypotheses about GSST only | content mapping pending (Step 3).

- P-006 | Pastawski–Yoshida–Harlow–Preskill, "Holographic quantum error-correcting codes", JHEP 06 (2015) 149 | arXiv:1503.06237v2 | v2 | pinned (sources/DHW-1503.06237.pdf — FILENAME MISNOMER, kept for ref stability; SHA-256 55572e7133514f4adc54f5bfb025111fc3fcb911c87ca22bae0a9e30955d1e05, 65pp) | AdS/CFT reconstruction needs AdS boundary + code subspace + bulk EFT (verified V-002, NOT as JLMS attribution) | caught by V-002 J-attribution flag | None.
- P-007 | Jafferis–Lewkowycz–Maldacena–Suh, "Relative entropy equals bulk relative entropy", JHEP 06 (2016) 004 | arXiv:1512.06431 | fetched 2026-11-08 | pinned (sources/JLMS-1512.06431.pdf, SHA-256 04198f2251ee5d14074b00ea4decc5ba856c758e0d81d96313457112bc99e08a, 24pp) | boundary=bulk relative entropy + area; code-subspace assumptions | attribution + assumption check in V-003 | None.
- P-008 | Blanco–Casini–Hung–Myers, "Relative entropy and holography", JHEP 08 (2013) 060 | arXiv:1305.3182 | fetched 2026-11-08 | pinned (sources/BCHM-1305.3182.pdf, SHA-256 6eceb7f4f0bdf5ccb2a39c9847850be747bf36de5ae077c8692461f3644f4542, 76pp) | first law Δ⟨H⟩=ΔS for ball-region perturbations (usage: H1 prerequisite context; H1 reduction verified V-002 item R as logic) | assumption mismatch recorded (balls-in-vacuum vs half-line-R-in-evaporating-BH) | None.

## Further (fetch only as hypothesis cycles require)

Tomita–Takesaki; Fisher/info-geometry texts; type III/II texts; Twirled Petz/Fawzi–Renner; AMPS/Harlow–Hayden.
