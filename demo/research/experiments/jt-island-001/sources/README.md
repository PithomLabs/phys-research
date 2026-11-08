# JT-ISLAND-001 primary sources (Step 1 record)

Fetched 2026-11-07 from arXiv canonical. All files live in this `sources/`
directory (TOE-v6 only); nothing was written to any phys repository.

## AMMZ (canonical R4 + R1 starting point)

- File: `AMMZ-1908.10996v2.pdf`
- arXiv:1908.10996v2 [hep-th], updated 2019-11-04 (latest version)
- Title: "The Page curve of Hawking radiation from semiclassical geometry"
- Authors: Almheiri, Mahajan, Maldacena, Zhao
- Pages: 22 (pdfinfo; the `file` page count is unreliable for this PDF)
- SHA-256: f68be764aa6e21b79d3095401a9434c8e02047d5e6bc327a3c7b370afca9a905
- Verified: header page reads "The Page curve of Hawking radiation from
  semiclassical geometry / arXiv:1908.10996v2 [hep-th] 4 Nov 2019".

## AEMM (canonical supporting point)

- File: `AEMM-1905.08762v3.pdf` (renamed to pin the resolved version)
- arXiv:1905.08762v3 [hep-th], updated 2019-11-04 (latest version; the
  dossier cites AEMM without a version pin, so latest-published is used)
- Title: "The entropy of bulk quantum fields and the entanglement wedge
  of an evaporating black hole"
- Authors: Almheiri, Engelhardt, Marolf, Maxfield
- Pages: 55 (pdfinfo)
- SHA-256: f51a59cbf6d660471965c6eca0920b70eb598652af366e65b60050edeabc3575
- Verified: header page reads "arXiv:1905.08762v3 [hep-th] 4 Nov 2019".

## Polyakov 1981 (issue-A support, canonical citation)

- Canonical: A.M. Polyakov, "Quantum geometry of bosonic strings",
  Phys. Lett. B103 (1981) 207-210. DOI: 10.1016/0370-2693(81)90743-7.
- Fetch status: NOT OBTAINED. No trusted open-access copy found; the only
  indexed full-text hits are untrusted proxy mirrors, which are rejected as
  a provenance risk. The audit (Step 2, issue A) will first determine from
  AMMZ/AEMM + first-principles 2D CFT derivation whether the original is
  necessary; if it is, escalate for access rather than citing a mirror.

## AP Almheiri-Polchinski (tiebreak source)

- File: `AP-1402.6334.pdf` (+ `.html`)
- arXiv:1402.6334v3 [hep-th], updated 2015-08-12
- Title: "Models of AdS2 Backreaction and Holography"
- SHA-256: 37e78a45c5e573098b9c1457597e04cd36c336eabed8e4d389923884ec1aabb
- NOTE: dossier cites "AP arXiv:1405.4251" — that ID is an unrelated
  statistics paper. True ID resolved via AEMM ref (35). Dossier's AP leg VOID.
- Used for: AP (5.1)-(5.2) verifies dossier's T_+−/T_±± forms; AP (5.3)/(5.5)
  confirms 16π-vs-8π pattern against dossier's +− equation.

## Penington (QES-origin chain)

- File: `Penington-1905.08255.pdf`
- arXiv:1905.08255 (latest), "Entanglement wedge reconstruction and the
  information paradox", 75 pp, header-verified
- SHA-256: fa7cf02db1814f4b55e72f33f2e817c5001645b5c4f1b4ea77444b38ee03738c
- Status: BOUNDED READ done (Step 2, 2026-11-07). Setup: one-sided BH formed
  from collapse, radiation into reservoir H_rad via absorbing BCs (higher-D
  AdS/CFT, not JT-specific). Results used: non-empty QES just inside horizon
  at scrambling-time delay (leading order); S = min(S_rad, A_hor/4G_N)
  Eq. (2); Page time defined by S_rad = A_hor/4G_N. Scope caveats recorded:
  RT-based Page curve alone "not entirely satisfactory" (does not explain
  entropy decrease/firewall — needs entanglement-wedge reconstruction).
  Compatibility: general-D conceptual origin imported by AMMZ; JT-specific
  realization comes from AEMM §4. Not a source of JT equations.

## Dossier equation access

- `pdftotext` extraction of "JT Black Hole Research Dossier.pdf" (1078 lines)
  loses all displayed equations (rendered images). Equation-level audit uses
  the rendered PDF pages plus the arXiv sources above — never the text
  extraction. AMMZ arXiv HTML v2 remains available as a cross-check for exact
  equation forms.
