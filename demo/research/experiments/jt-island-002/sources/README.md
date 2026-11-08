# JT-ISLAND-002 Primary Sources (Phase 0 record)

Fetched 2026-11-08. All files live in this `sources/` directory; nothing was
written to any other repository.

## GSST canonical (JHEP published — canonical locator system)

- File: `GSST-JHEP05-2020-091.pdf` (+ `.txt` extraction)
- Gautason, Schneiderbauer, Sybesma, Thorlacius, *Page Curve for an
  Evaporating Black Hole*, JHEP 05 (2020) 091, DOI 10.1007/JHEP05(2020)091
- Received Apr 22, accepted Apr 29, published May 20, 2020. Open access (CC-BY 4.0).
- Pages: 23 (pdfinfo)
- SHA-256: `e1200b84c53dd44d83654e3c3905bf0b7a20b41c7442fbcb06ba1a9a7322a68f`
- Verified: header page reads "JHEP05(2020)091", "Published: May 20, 2020".

## GSST arXiv (version cross-check)

- File: `GSST-2004.00598.pdf` (+ `.txt` extraction)
- arXiv:2004.00598v2 [hep-th], 21 Apr 2020 (v1: 1 Apr 2020; v2: typo/reference edits)
- Pages: 25 (pdfinfo; the `file` page count is unreliable for this PDF)
- SHA-256: `808d6b6e7d89ddb2b394cfd611e6c25a0669b956fff5b574409d9e5496ea6047`
- Verified: header reads "arXiv:2004.00598v2 [hep-th] 21 Apr 2020".

## Version difference (RECORDED — do not silently transfer numbers)

- §3: JHEP has 35 numbered equations (3.1–3.35); arXiv v2 has 36 (3.1–3.36).
  The arXiv T_μν-definition equation (3.12) is unnumbered in JHEP, shifting all
  later §3 numbers down by one: **JHEP (3.k) = arXiv (3.k+1) for k = 12..35**.
  Verified instances: field equation JHEP(3.12) = arXiv(3.13);
  ω-frame JHEP(3.19) = arXiv(3.20); σ-frame JHEP(3.20) = arXiv(3.21);
  collapse solution JHEP(3.35) = arXiv(3.36).
- §§2, 4, 5, 6: numbering identical in both versions
  (2.1–2.3; 4.1–4.7; 5.1–5.22; 6.1). Content equivalence per equation is
  verified during Phases A–F, not assumed from numbering.
- Footnote positions differ (JHEP footnote 3 vs arXiv footnote 3 adjacent text);
  footnote content checked per use.

## Auxiliary primaries

- CGHS `CGHS-hepth-9111056.pdf` (+ `.txt` extraction): Callan–Giddings–Harvey–
  Strominger, "Evanescent black holes", PRD 45 (1992) R1005, hep-th/9111056v1.
  12pp. SHA-256: `3a239311da3b2b4de86f74899989f170789705f9c9e562b1886eb6c1c32532d9`.
  Consumed: (16) trace −1/12 (N=1) confirming fn.4 ×2; (18–19) flux +λ²/48 with
  CGHS text "limiting value of T^f_{−−} is the flux of f-particle energy".
  Fetched for OQ-4 (ledger_review Track 2a).
- Not fetched: original RST, HSS, Christensen–Fulling, Fiola et al., Hirano,
  Brown–Henneaux (cited, unpinned).
