# JT-ISLAND-001 audit report (pause-and-report)

**Top finding: the dossier does not pass audit — 8 BLOCKER-class contradictions confirmed, and per the approved pause-and-report decision, Parts B/C and interop drafting are paused until the director rules on the model fork below.**

Full evidence: [dossier-audit-note-v1.0.md](/home/chaschel/Desktop/biz/TOE-v6/research/experiments/jt-island-001/dossier-audit-note-v1.0.md). Method: every dossier equation read from rendered PDF pages (text extraction loses them all); every locator checked against AMMZ v2, AEMM v3, AP 1402.6334v3, and Penington (all fetched, hashed, in `research/experiments/jt-island-001/sources/`).

## The load-bearing defects

- **All equation locators are fabricated numbering.** AMMZ uses single-sequence (1)–(15); AEMM single-sequence; AP's (X.Y) contents don't match. Every dossier locator of the form "AMMZ Eq. (2.9)/(3.4)/(3.15)/(3.18)" fails to resolve. A corrected locator map was rebuilt (e.g. QES position is AMMZ(12), late entropy is (13), growth branch is (14), island rule is (15)).
- **Issue C — island 2× is a denominator error.** P.10 writes S_island = (φ₀+φ_h)/2G_N = 2S_BH, contradicting its own pp.8–9 (4G_N, one endpoint), AMMZ(13) (1×), and AEMM (1×). Same page drops φ₀ in the Page-time equality.
- **Issue B — three mutually inconsistent numbers.** The no-island log formula is correct (implies slope πcT/6, verified), but both printed linearizations are wrong (2×, 4×), and t_P ≈ 3βS/πc matches neither its own branches nor the independent 6βS/πc. Neither source contains it.
- **Issue F — field-equation sign/coefficient defects.** Action topological term 4× off vs AEMM(6); covariant metric equation has a gφ sign error (4 witnesses: AEMM trace, vacuum-solution test, AP 16π/8π pattern, the dossier's own p.16 which prints the correct form); +− expansion matches neither its own covariant form nor AEMM(11a).
- **Issue A — RC2 removal claim is false.** "Removing the trace anomaly breaks conservation" contradicts the dossier's own p.13 derivation and AP §5.1 (conservation and anomaly are independent joint inputs).
- **Issue E — no single consistent model.** Single-sided-collapse narrative (AMMZ) blended with two-sided+shock equations (AEMM): no E_S shock, no holographic-matter mention, no t=f(u), constant-z vs ϵf'(u) cutoff, φ_h in 3 meanings, unmapped coordinates.
- **Issue G — ADM mass sign error** (gives negative BH mass); matching oversimplified; dM/dt standard but underived. **Issue D** — two-interval formula needs an unstated purity hypothesis, invalid in AEMM's setup. **Issue H** — pervasive mislabeling + circular self-citation (rationale.md, protocol.md); Penington never cited; dossier's "AP 1405.4251" resolves to an unrelated statistics paper (true ID 1402.6334, resolved via AEMM ref 35).

## Fairness record

Genuinely clean items include the Poincaré metric, R+2=0, the ++/−− constraints (exact AEMM match), the no-island log formula, AP-verified stress-tensor components, and the dossier's scope discipline (N0/N1 honesty, exclusions). Its equations fail; its scope honesty holds.

## Decisions needed (audit §8)

1. **Model fork:** adopt AMMZ single-sided-collapse as the working model (recommendation — matches the dossier narrative and is the only regime where its R4 formulas are valid), importing AEMM equations only with per-equation compatibility justification? Or redirect?
2. **Island normalization:** acknowledge adopting 1× (voiding C-R4-04/05 as printed, Page time → 6βS/πc) before formalizing Part B on corrected values?
3. Noting only (unambiguous, 4 witnesses): adopting the AEMM-consistent metric-equation sign and replacing the RC2 removal claim — proceeding unless directed otherwise.

## Work status

Continuing independently meanwhile: register scaffolding, corrected-R1 baseline skeleton (provisional), OpenCode checklist seeds. No Part B/C derivations on disputed normalizations.
