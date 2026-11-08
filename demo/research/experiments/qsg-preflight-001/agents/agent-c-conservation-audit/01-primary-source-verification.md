# 01 — Primary-source verification

Audit: QSG-T001, covariant conservation of the Hadamard-renormalized
stress-energy tensor. Author: Agent C, independent conservation audit.
Date: 2026-11-06.

This document verifies the primary bibliography, conventions, equations and
page locations used in the audit. It does not accept Agent B's citations,
classifications or equations. Where the full text of a source could not be
inspected, the item is marked **BLOCKED** rather than inferred.

Disposition vocabulary used throughout the seven deliverables:

| Disposition | Meaning |
|---|---|
| DERIVED | Obtained independently here from stated assumptions and explicit calculation. |
| SOURCE-REPORTED | Asserted by a primary source; not independently re-derived in full. |
| INDEPENDENTLY REPRODUCED | A source result was checked by an independent calculation or by direct equation-level comparison. |
| CONDITIONAL | Holds only under an additional assumption or criterion stated explicitly. |
| FALSIFIED | Shown false, inconsistent, or non-conforming by explicit counter-calculation or source comparison. |
| BLOCKED | Cannot be verified with the accessible material; the exact missing item is stated. |

`PASS` is not used as a disposition.

---

## 1. Verification method

Three levels were used.

1. **Metadata level.** DOI, journal, volume, page range, year and arXiv ID
   were checked against the publisher page, INSPIRE, or the arXiv abstract
   page. This is sufficient to catch wrong page numbers, wrong journals and
   misattributed references.
2. **Equation level.** For arXiv-accessible sources, the full text was
   downloaded and the relevant equations were read in the PDF, not only in
   an extracted-text rendering. For Décanini–Folacci, pages containing
   Eqs. (52), (55), (56), (57), (61), (64), (67), (68), (70), (71), (74),
   (108), (109) and (111) were rendered as page images and inspected.
3. **Page level.** Agent B's PDF was rendered page by page. The claims
   audited in `05-agent-b-findings.md` carry the printed page number of the
   supplied PDF.

The two Wald papers are paywalled at the full-text level. Their abstracts
were verified directly (INSPIRE for Wald 1977; APS for Wald 1978). The
abstracts are sufficient to establish the structural claims used in this
audit, but not the detailed page-level equations. Those details are marked
BLOCKED.

---

## 2. Verified bibliography

| Source | Verified bibliographic data | What was verified | Disposition |
|---|---|---|---|
| R. M. Wald, “The back reaction effect in particle creation in curved spacetime” | Commun. Math. Phys. **54** (1977) 1–19. DOI 10.1007/BF01609833. | INSPIRE record and Springer abstract. Abstract states five conditions, uniqueness (at most one renormalized stress-energy operator), and existence investigated via an explicit point-splitting-type prescription. Full text not accessible. | SOURCE-REPORTED at abstract level; detailed five-condition list **BLOCKED**. |
| R. M. Wald, “Trace anomaly of a conformally invariant quantum field in curved spacetime” | Phys. Rev. D **17** (1978) 1477–1484. DOI 10.1103/PhysRevD.17.1477. | APS abstract and article metadata. Abstract states: the naive point-separation divergence is the gradient of a local curvature term; subtracting that term times the metric restores conservation and makes the trace nonvanishing; the trace is fixed modulo a conserved local curvature term; the conformal-scalar trace is \((2880\pi^2)^{-1}(C_{\alpha\beta\delta\gamma}C^{\alpha\beta\delta\gamma}+R_{\alpha\beta}R^{\alpha\beta}-\tfrac13R^2)\) plus an arbitrary constant times \(\nabla_\alpha\nabla^\alpha R\); axiom 5 forbids third or higher derivatives. Full text not accessible. | SOURCE-REPORTED at abstract level; detailed equations **BLOCKED**. |
| P. Décanini and A. Folacci, “Hadamard renormalization of the stress-energy tensor for a quantized scalar field in a general spacetime of arbitrary dimension” | Phys. Rev. D **78** (2008) 044025. arXiv:gr-qc/0512118v2. DOI 10.1103/PhysRevD.78.044025. | Full arXiv text and rendered page images. Conventions: Hawking–Ellis, signature \((-+++)\), \(\Box=g^{\mu\nu}\nabla_\mu\nabla_\nu\), units \(\hbar=c=1\), general \(\xi\), dimension \(D\). Equations (52), (55), (56a,b), (57), (61), (64), (67), (68), (69), (70), (71), (73), (74), (79), (80), (83), (108a), (109), (110), (111) read and checked. | INDEPENDENTLY REPRODUCED for the equations used in the conservation calculation; the full renormalization construction is SOURCE-REPORTED. |
| S. Hollands and R. M. Wald, “Conservation of the stress tensor in perturbative interacting quantum field theory in curved spacetimes” | Rev. Math. Phys. **17** (2005) 227–312. arXiv:gr-qc/0404074v2. DOI 10.1142/S0129055X05002340. | Full arXiv text. Axioms T1–T11, the microlocal spectrum condition (Eq. (31)), the local Hadamard parametrix and local Wick expansion, Theorem 5.2 (trace cocycle) and Theorem 5.3 (interacting conservation, statement (4): \(\nabla^a\Theta^L_{ab}=0\) under T1–T10 and T11b). Footnote 14: in even dimensions a symmetric local Hadamard parametrix cannot satisfy the wave equation to arbitrarily high order; the failure produces the anomalies. | SOURCE-REPORTED for the perturbative theorem; the free-field conservation calculation here is INDEPENDENTLY REPRODUCED. |
| M. J. Radzikowski, “Micro-local approach to the Hadamard condition in quantum field theory on curved space-time” | Commun. Math. Phys. **179** (1996) 529–553. DOI 10.1007/BF02100096. | Bibliographic metadata verified. The equivalence between the Hadamard condition and the microlocal spectrum condition is used as a standard result. | SOURCE-REPORTED. |
| W. Tichy and É. É. Flanagan, “How unique is the expected stress-energy tensor of a massive scalar field?” | Phys. Rev. D **58** (1998) 124007. arXiv:gr-qc/9807015. DOI 10.1103/PhysRevD.58.124007. | Full arXiv text. Eq. (1.7): \(t_{ab}=\frac{1}{\sqrt{-g}}\frac{\delta}{\delta g^{ab}}\int d^4x\sqrt{-g}\,m^4F(R/m^2)\). Eq. (1.10): the tensor written in terms of \(F'\), \(F''\), \(F^{(3)}\). The paper states that the massive-field ambiguity allowed by the Wald axioms is an infinite-parameter family (a free function \(F\)), not a two-parameter family. | SOURCE-REPORTED; the structure is INDEPENDENTLY REPRODUCED (the metric variation of a local invariant is conserved by construction). |
| L. L. Salcedo, “On the uniqueness of the expected stress-energy tensor in renormalizable field theories” | Phys. Rev. D **60** (1999) 107502. arXiv:gr-qc/9903057. DOI 10.1103/PhysRevD.60.107502. | Full arXiv text. The paper argues that the effective-action ambiguity in a four-dimensional renormalizable theory is at most a local polynomial of canonical dimension four; for a massive scalar without self-interaction the allowed terms are \(m^4\), \(m^2R\), \(R^2\), \(R_{\mu\nu}R^{\mu\nu}\) and the Gauss–Bonnet density (topological, no stress-tensor contribution). It argues that arbitrary functions of \(R/m^2\) are excluded by the renormalizability/effective-action criterion. | SOURCE-REPORTED; the four-term minimal basis is INDEPENDENTLY REPRODUCED. |
| S. Hollands and R. M. Wald, “Local Wick polynomials and time ordered products of quantum fields in curved spacetime” | Commun. Math. Phys. **223** (2001) 289–326. arXiv:gr-qc/0103074. DOI 10.1007/s002200100540. | Bibliographic data verified in the reference list of the 2005 paper. Used only as the local-covariance background. | SOURCE-REPORTED. |
| S. Hollands and R. M. Wald, “Existence of local covariant time ordered products of quantum fields in curved spacetime” | Commun. Math. Phys. **231** (2002) 309–345. arXiv:gr-qc/0111108. DOI 10.1007/s00220-002-0719-y. | Bibliographic data verified in the reference list of the 2005 paper. Used only as the local-covariance background. | SOURCE-REPORTED. |

---

## 3. Conventions

### 3.1 Frozen model

The audit uses the frozen model of the assignment:

- four-dimensional, smooth, globally hyperbolic Lorentzian manifold
  \((M,g)\);
- signature \((-+++)\);
- natural units \(c=\hbar=1\);
- real scalar field \(\phi\), mass \(m>0\), minimal coupling \(\xi=0\);
- fixed classical background; no backreaction.

### 3.2 Décanini–Folacci

Décanini–Folacci state explicitly that they use the geometrical conventions of
Hawking–Ellis for \(R\), \(R_{\mu\nu}\) and \(R_{\mu\nu\rho\sigma}\), units
\(\hbar=c=1\), and \(\Box=g^{\mu\nu}\nabla_\mu\nabla_\nu\). The signature is
\((-+++)\). The scalar field action is
\[
S=-\frac12\int d^Dx\sqrt{-g}
\left(g^{\mu\nu}\nabla_\mu\phi\nabla_\nu\phi
+\xi R\phi^2+m^2\phi^2\right),
\]
and the wave equation is
\[
(\Box-m^2-\xi R)\phi=0.
\]
At \(\xi=0\) this is the frozen equation \((\Box-m^2)\phi=0\).

The dimension-dependent constant is
\[
\alpha_D=\frac{\Gamma(D/2-1)}{(2\pi)^{D/2}}\qquad(D\ne2),
\]
so
\[
\alpha_4=\frac{1}{4\pi^2}.
\]
The Hadamard singular part of the Feynman propagator is written with the
prefactor \(i\alpha_D/2\); in four dimensions this is \(i/(8\pi^2)\), which is
the standard \(1/(8\pi^2)\) coefficient when the world function is
\(\sigma=\tfrac12s^2\).

### 3.3 Hollands–Wald

Hollands–Wald use abstract index notation and state that their conventions
generally follow their earlier papers. The relevant statements used here are
the microlocal spectrum condition (Eq. (31)), the local Hadamard parametrix,
and the conservation theorem (Theorem 5.3). The conservation identity is
stated as \(\nabla^a\Theta^L_{ab}=0\); it is a tensor identity and does not
depend on the sign of the Riemann tensor.

### 3.4 Sign-convention check that matters

The only curvature-sign-sensitive step in the conservation calculation is the
commutator
\[
\nabla^\mu\nabla_\nu\nabla_\mu w
=\nabla_\nu\Box w+R_\nu{}^\mu\nabla_\mu w,
\]
which is used together with the transport equation
\[
\nabla^\mu w_{\mu\nu}
=\frac14\nabla_\nu\Box w
+\frac12R_\nu{}^\mu\nabla_\mu w
-\nabla_\nu v_1 .
\]
The \(R_\nu{}^\mu\nabla_\mu w\) terms enter with opposite signs and cancel in
\(\nabla^\mu\lim T_{\mu\nu}W\). The final residual
\(-\tfrac{D}{4}\alpha_D\nabla_\nu v_1\) is therefore independent of the
Riemann sign convention, provided the commutator and the transport equation
are derived in the same convention. This was checked explicitly.

---

## 4. Exact equations verified

### 4.1 Décanini–Folacci 2008

The following equations were read from the rendered pages and checked
algebraically.

**Hadamard remainder equation.** For even \(D\ne2\),
\[
(\Box-m^2-\xi R)W
=-(D+2)V_1-2V_1^{;\mu}\sigma_{;\mu}+O(\sigma),
\tag{52}
\]
and, using the expansion of \(V_1\),
\[
(\Box-m^2-\xi R)W
=-(D+2)v_1+\frac D2 v_1^{;\mu}\sigma_{;\mu}+O(\sigma).
\tag{55}
\]
The coincidence limit at \(\xi=0\) is therefore
\[
\lim_{x'\to x}(\Box-m^2)W=-(D+2)v_1,
\]
i.e. \(-6v_1\) in four dimensions. This is the equation Agent B replaces by
the false statement \((\Box-m^2)W=0\).

**Transport equations.** At general \(\xi\),
\[
w^\rho{}_\rho=(m^2+\xi R)w-(D+2)v_1,
\tag{56a}
\]
\[
w^\rho{}_{a;\rho}
=\frac14(\Box w)_{;a}
+\frac12w^\rho{}_{\rho;a}
+\frac12R^\rho{}_a w_{;\rho}
-\frac12(m^2+\xi R)w_{;a}
+\frac D2 v_{1;a}.
\tag{56b}
\]
Combining them gives the useful form
\[
w^\rho{}_{a;\rho}
=\frac14(\Box w)_{;a}
+\frac12R^\rho{}_a w_{;\rho}
+\frac12\xi R_{;a}w
-v_{1;a}.
\tag{57}
\]
At \(\xi=0\),
\[
w^\rho{}_\rho=m^2w-(D+2)v_1,
\qquad
w^\rho{}_{a;\rho}
=\frac14(\Box w)_{;a}
+\frac12R^\rho{}_a w_{;\rho}
-v_{1;a}.
\]

**Point-splitting operator.** For general \(\xi\), the bi-differential
operator is
\[
\begin{aligned}
T_{\mu\nu}(x,x')
={}&(1-2\xi)g_\nu{}^{\nu'}\nabla_\mu\nabla_{\nu'}
+\left(2\xi-\frac12\right)g_{\mu\nu}g^{\rho\sigma}\nabla_\rho\nabla_{\sigma'}\\
&-2\xi g_\mu{}^{\mu'}g_\nu{}^{\nu'}\nabla_{\mu'}\nabla_{\nu'}
+2\xi g_{\mu\nu}\nabla^\rho\nabla_\rho\\
&+\xi\left(R_{\mu\nu}-\frac12g_{\mu\nu}R\right)
-\frac12g_{\mu\nu}m^2 .
\end{aligned}
\tag{61}
\]
At \(\xi=0\) this reduces to the symmetrized operator
\[
T_{\mu\nu}(x,x')
=\frac12\left(g_\nu{}^{\nu'}\nabla_\mu\nabla_{\nu'}
+g_\mu{}^{\mu'}\nabla_\nu\nabla_{\mu'}\right)
-\frac12g_{\mu\nu}\left(g^{\rho\sigma}\nabla_\rho\nabla_{\sigma'}+m^2\right).
\]

**Coincidence-limit expression.** For even \(D\),
\[
\begin{aligned}
\langle T_{\mu\nu}\rangle_{\rm ren}
={}&\frac{\alpha_D}{2}\Big[
-w_{\mu\nu}-\frac12g_{\mu\nu}w^\rho{}_\rho
+\frac12(1-2\xi)w_{;\mu\nu}
+\left(2\xi-\frac12\right)\frac12g_{\mu\nu}\Box w\\
&\qquad\qquad
+\xi\left(R_{\mu\nu}-\frac12g_{\mu\nu}R\right)w
-\frac12g_{\mu\nu}m^2w
\Big]+\tilde\Theta_{\mu\nu}.
\end{aligned}
\tag{67}
\]
Conservation of (67) requires
\[
\left[\tilde\Theta_{\mu\nu}-\frac D4\alpha_D g_{\mu\nu}v_1\right]^{;\nu}=0
\tag{68}
\]
for even \(D\). Thus one may choose
\[
\tilde\Theta_{\mu\nu}
=\frac D4\alpha_D g_{\mu\nu}v_1+\Theta_{\mu\nu},
\qquad
\Theta_{\mu\nu}{}^{;\nu}=0,
\tag{69}
\]
and the renormalized tensor is
\[
\boxed{
\langle T_{\mu\nu}\rangle_{\rm ren}
=\frac{\alpha_D}{2}\left[
\lim_{x'\to x}T_{\mu\nu}(x,x')W(x,x')
+\frac D2 g_{\mu\nu}v_1
\right]
+\Theta_{\mu\nu}
}
\tag{70}
\]
with
\[
\Theta_{\mu\nu}{}^{;\nu}=0.
\tag{74}
\]
For even \(D\), the equivalent \(w\)-form is
\[
\langle T_{\mu\nu}\rangle_{\rm ren}
=\frac{\alpha_D}{2}\Big[
-w_{\mu\nu}
+\frac12(1-2\xi)w_{;\mu\nu}
+\left(2\xi-\frac12\right)\frac12g_{\mu\nu}\Box w
+\xi R_{\mu\nu}w
-g_{\mu\nu}v_1
\Big]
+\Theta_{\mu\nu}.
\tag{71}
\]
At \(D=4\), \(\alpha_4=1/(4\pi^2)\), and the local correction in (70) is
\[
\frac D4\alpha_D g_{\mu\nu}v_1
=\frac{1}{4\pi^2}g_{\mu\nu}v_1 .
\]

**Hadamard coefficients at \(\xi=0\).** From Eq. (108a) and Eq. (109),
\[
v_0=\frac12m^2-\frac1{12}R,
\tag{108a'}
\]
\[
v_1=\frac18m^4-\frac1{24}m^2R
+\frac1{120}\Box R
+\frac1{288}R^2
-\frac1{720}R_{\mu\nu}R^{\mu\nu}
+\frac1{720}R_{\mu\nu\rho\sigma}R^{\mu\nu\rho\sigma}.
\tag{109'}
\]
The general-\(\xi\) forms are
\[
v_0=\frac12m^2+\frac12\left(\xi-\frac16\right)R,
\]
\[
v_1=\frac18m^4+\frac14\left(\xi-\frac16\right)m^2R
-\frac1{24}\left(\xi-\frac15\right)\Box R
+\frac18\left(\xi-\frac16\right)^2R^2
-\frac1{720}R_{\mu\nu}R^{\mu\nu}
+\frac1{720}R_{\mu\nu\rho\sigma}R^{\mu\nu\rho\sigma}.
\]

**Trace anomaly for the conformal scalar.** At \(m=0\), \(\xi=1/6\), Eq. (111)
gives
\[
\langle T^\mu{}_\mu\rangle_{\rm ren}
=\frac{1}{(2\pi)^2}\left[
\frac1{720}\Box R
-\frac1{720}R_{\mu\nu}R^{\mu\nu}
+\frac1{720}R_{\mu\nu\rho\sigma}R^{\mu\nu\rho\sigma}
\right].
\tag{111}
\]
Equivalently, using \(C^2=R_{\mu\nu\rho\sigma}^2-2R_{\mu\nu}^2+\tfrac13R^2\),
\[
\langle T^\mu{}_\mu\rangle_{\rm ren}
=\frac{1}{2880\pi^2}
\left(C_{\mu\nu\rho\sigma}C^{\mu\nu\rho\sigma}
+R_{\mu\nu}R^{\mu\nu}-\frac13R^2\right)
+\frac{1}{2880\pi^2}\Box R .
\]
The first line is exactly Wald's 1978 abstract formula up to the coefficient
of the \(\Box R\) term, which Wald leaves arbitrary at the abstract level.

### 4.2 Hollands–Wald 2005

**Microlocal spectrum condition (axiom T3, Eq. (31)).**
\[
C^+(M,g)
=\left\{(x_1,k_1;x_2,-k_2)\in T^*M^2\setminus\{0\}
\;\middle|\;
(x_1,k_1)\sim(x_2,k_2),\;
k_1\in(V^+)_{x_1}\right\}.
\]
The Hadamard condition is \(\mathrm{WF}(\omega_2)\subset C^+(M,g)\).

**Local Hadamard parametrix.** Hollands–Wald define a symmetric, locally
constructed Hadamard parametrix \(H(x,y)\) and the local Wick expansion
(Eq. (59)). Footnote 14 is the decisive statement for this audit:

> “in even dimensions, it is impossible to construct a local and covariant
> Hadamard parametrix \(H(x,y)\), that is symmetric in \(x\) and \(y\) and
> satisfies the wave equation to arbitrarily high order in the geodesic
> distance between \(x\) and \(y\). Consequently, if one requires \(H(x,y)\)
> to be symmetric ... then it will fail to satisfy the wave equation, and
> ‘anomalies’ will occur for the regularized quantities.”

This directly contradicts Agent B's assertion that \(W=G^+-H\) obeys the
homogeneous Klein–Gordon equation. In even dimensions the failure is
structural, not a technicality.

**Conservation theorem (Theorem 5.3, statement (4)).** Under axioms T1–T10
and T11b, the interacting stress-energy tensor is conserved:
\[
\nabla^a\Theta^L_{ab}(x)=0.
\]
The theorem is perturbative and interacting; the free-field point-splitting
calculation in `04` is a direct realization of the same Ward identity in the
one-point function.

**Trace cocycle (Theorem 5.2).** For a massless conformally coupled scalar,
\(T^a{}_a(x)=C(x)\mathbb 1\), where \(C(x)\) is a local curvature term of
mass dimension \(D\) satisfying the cocycle condition. This is the
anomaly statement, distinct from conservation.

### 4.3 Wald 1978 abstract

The APS abstract was read directly. The load-bearing sentences are:

> “It is now seen that this prescription yields a stress-energy tensor whose
> divergence is not zero but is the gradient of a local curvature term.
> However, this deficiency can be corrected by subtracting off this local
> curvature term times the metric tensor; as a direct consequence the trace
> of \(T_{\mu\nu}\) becomes nonvanishing.”

This is the exact structural statement that Agent B's derivation is missing:
the residual is a gradient of a local curvature term, and the correction is
that term times the metric. The abstract also gives the conformal trace
\[
T^\mu{}_\mu
=\frac{1}{2880\pi^2}
\left(
C_{\alpha\beta\delta\gamma}C^{\alpha\beta\delta\gamma}
+R_{\alpha\beta}R^{\alpha\beta}
-\frac13R^2
\right)
+c\,\nabla_\alpha\nabla^\alpha R,
\]
with \(c\) arbitrary at the abstract level, and states that axiom 5
(no local curvature terms containing third or higher derivatives of the
metric) sets \(c=0\) if it can be imposed. The detailed equations are
BLOCKED.

### 4.4 Wald 1977 abstract

The INSPIRE/Springer abstract was read directly. It states that **five**
conditions are listed, that they uniquely determine \(T_{\mu\nu}\) (at most
one renormalized stress-energy operator), and that existence is investigated
via an explicit point-splitting-type prescription. The exact list of the five
conditions is not in the abstract and the full text is paywalled; that list
is **BLOCKED**. The abstract alone is sufficient to falsify Agent B's
“four axioms” statement on p. 11.

### 4.5 Tichy–Flanagan 1998

Eq. (1.7) is
\[
t_{ab}
=\frac{1}{\sqrt{-g}}\frac{\delta}{\delta g^{ab}}
\int d^4x\sqrt{-g}\,m^4F(R/m^2),
\]
where \(F\) is any dimensionless function of a dimensionless argument. The
paper states that this is an infinite-parameter ambiguity for the massive
field, “rather than a two parameter ambiguity” in the massless case. This is
the primary source for the claim that the Wald axioms alone do not fix a
four-parameter family. The paper also notes that \(t_{ab}\to0\) as
\(m^2\to0\) and as \(R\to0\) for suitable \(F\).

### 4.6 Salcedo 1999

Salcedo argues that the effective-action ambiguity in a four-dimensional
renormalizable theory is at most a local polynomial of canonical dimension
four. For the massive scalar without self-interaction, the allowed terms are
\(m^4\), \(m^2R\), \(R^2\), \(R_{\mu\nu}R^{\mu\nu}\), and the Gauss–Bonnet
density (which is topological and does not contribute to the stress tensor).
He explicitly argues that arbitrary functions of \(R/m^2\) of the
Tichy–Flanagan type are excluded by the requirement that the effective action
have a consistent ultraviolet-finite part. This is an additional criterion
beyond the Wald axioms; it is the reason the finite-ambiguity classification
in this audit is CONDITIONAL.

---

## 5. Source-supported claim matrix

| Claim | Source support | Independent check | Disposition |
|---|---|---|---|
| Classical action, EOM and \(T_{\mu\nu}\) of the minimally coupled scalar | Standard; Décanini–Folacci Eq. (9) for \(T_{\mu\nu}\) | Full derivation in `02` | DERIVED |
| Off-shell identity \(\nabla^\mu T_{\mu\nu}=(\Box-m^2)\phi\,\nabla_\nu\phi\) | Standard Noether/diffeo argument | Full derivation in `02` | DERIVED |
| Hadamard form of \(G^+\) and smoothness of \(W=G^+-H\) | Décanini–Folacci §II; Hollands–Wald §2 and footnote 14; Wald 1978 abstract | Definitional; smoothness is the Hadamard condition | SOURCE-REPORTED; definitional part DERIVED |
| \(W\) does **not** satisfy \((\Box-m^2)W=0\) | Décanini–Folacci Eq. (52)/(55); Hollands–Wald footnote 14; Wald 1978 abstract | Explicit finite-part calculation in `03` | DERIVED / INDEPENDENTLY REPRODUCED |
| \(\lim(\Box-m^2)W=-(D+2)v_1=-6v_1\) at \(D=4\), \(\xi=0\) | Décanini–Folacci Eq. (55) | Explicit expansion of the \(v_1\sigma\ln|\sigma|\) term in `03` | INDEPENDENTLY REPRODUCED |
| \(v_0=\tfrac12m^2-\tfrac1{12}R\), \(v_1\) as in Eq. (109′) | Décanini–Folacci Eqs. (108a), (109) | Direct substitution at \(\xi=0\) in `01` and `03` | INDEPENDENTLY REPRODUCED |
| Raw point-split divergence is a gradient of a local curvature term | Wald 1978 abstract; Décanini–Folacci Eq. (68) | Explicit calculation in `04` | DERIVED |
| Correction is the local curvature term times the metric | Wald 1978 abstract; Décanini–Folacci Eq. (70) | Explicit calculation in `04` | DERIVED |
| Raw divergence \(=-\tfrac D4\alpha_D\nabla_\nu v_1\); correction divergence \(=+\tfrac D4\alpha_D\nabla_\nu v_1\) | Décanini–Folacci Eqs. (68)–(70) | Explicit calculation in `04` | DERIVED |
| Trace anomaly of the conformal scalar | Wald 1978 abstract; Décanini–Folacci Eq. (111) | Direct comparison of the two formulas | INDEPENDENTLY REPRODUCED |
| Trace nonvanishing does not imply conservation failure | Wald 1978 abstract; Hollands–Wald Theorem 5.2 and discussion | Explicit tensor argument in `04` | DERIVED |
| Massive-field finite ambiguity is infinite under the Wald axioms alone | Tichy–Flanagan Eq. (1.7) and surrounding text | Metric variation of a local invariant is conserved by construction | SOURCE-REPORTED |
| Additional renormalizability criterion reduces the ambiguity to the dimension-four local basis | Salcedo 1999 | Independent construction of the four-term basis in `04` | SOURCE-REPORTED / CONDITIONAL |
| Perturbative interacting conservation under T1–T11 | Hollands–Wald Theorem 5.3 | Free-field case independently derived; interacting case not reproduced | SOURCE-REPORTED |

---

## 6. Agent B's bibliography and attribution: verified defects

The following are direct comparisons between the supplied PDF and the
verified sources. Page numbers are those of the supplied PDF.

| Page | Agent B's assertion | Verified result | Disposition |
|---|---|---|---|
| p. 1 table | “Wald (1977, 1978), Commun. Math. Phys. 54, 11; Phys. Rev. D 17, 14772” | Wald 1977 is Commun. Math. Phys. **54**, 1–19; Wald 1978 is Phys. Rev. D **17**, 1477–1484. | FALSIFIED (bibliographic). |
| p. 1 table | Wald “formulated four fundamental axioms” | Wald 1977 abstract lists **five** conditions. Wald 1978 abstract refers to axioms 1, 3, 4 and 5. | FALSIFIED. |
| p. 11 | “Wald proved ... four axioms” and cites ref. 1 | Ref. 1 in Agent B's bibliography is Salcedo, arXiv:gr-qc/9903057, not Wald. | FALSIFIED (misattribution). |
| p. 11 | Conservation is cited to ref. 7 | Ref. 7 is “Quantum fields on self-similar spacetimes,” arXiv:2509.10360, not a Wald paper and not the conservation theorem. | FALSIFIED (misattribution). |
| pp. 22–24 | “Works cited” list | There is no Wald 1977 or Wald 1978 entry. The closest entries are [1] Salcedo, [3] Décanini–Folacci and [9] Hollands–Wald. | FALSIFIED (missing primary sources). |
| p. 18 | “Radzikowski's microlocal spectrum condition ...12” | Ref. 12 is arXiv:1911.01829, not Radzikowski 1996. Radzikowski's paper is absent from the bibliography. | FALSIFIED (misattribution). |
| p. 11 | Finite ambiguity \(c_1g_{\mu\nu}+c_2G_{\mu\nu}+c_3H^{(1)}_{\mu\nu}+c_4H^{(2)}_{\mu\nu}\), all \(c_i\) dimensionless | \(g_{\mu\nu}\) has dimension 0, so \(c_1\) must have dimension 4; \(G_{\mu\nu}\) has dimension 2, so \(c_2\) must have dimension 2. The dimensionless coefficients belong on \(m^4g_{\mu\nu}\) and \(m^2G_{\mu\nu}\). | FALSIFIED (dimensional). |
| p. 11 | \(H^{(1)}_{\mu\nu}=2\nabla_\mu\nabla_\nu R-2g_{\mu\nu}\Box R+2RR_{\mu\nu}-\tfrac12g_{\mu\nu}R^2\) | Correct variation of \(\int\sqrt{-g}R^2\): \(H^{(1)}_{\mu\nu}=2RR_{\mu\nu}-\tfrac12g_{\mu\nu}R^2+2(g_{\mu\nu}\Box-\nabla_\mu\nabla_\nu)R\). The derivative terms have the opposite sign. | FALSIFIED (sign). |
| p. 13 | \(\langle T^\mu{}_\mu\rangle=\frac{1}{16\pi^2}a_2=\frac{1}{2880\pi^2}[C^2-E_4+\tfrac13\Box R]\) | Wald 1978 abstract: \((2880\pi^2)^{-1}(C^2+R_{\alpha\beta}R^{\alpha\beta}-\tfrac13R^2)+c\Box R\). Décanini–Folacci Eq. (111): \((2880\pi^2)^{-1}(R_{\mu\nu\rho\sigma}^2-R_{\mu\nu}^2+\Box R)\). These are equivalent up to the \(\Box R\) coefficient; Agent B's \(C^2-E_4\) combination is not. | FALSIFIED. |
| p. 13 | Uses the minimal-coupling \(a_2\) for the conformally coupled (\(m=0,\xi=1/6\)) case | The heat-kernel coefficient for a conformal scalar is not the minimal-coupling \(a_2\); the \(R^2\) and \(\Box R\) terms differ. | FALSIFIED. |

The two most consequential attribution defects are:

1. **The five-condition vs four-axiom error.** Agent B omits axiom 5, which is
   precisely the condition that controls the \(\Box R\) ambiguity in the trace
   anomaly. His finite-ambiguity discussion is therefore not Wald's.
2. **The Salcedo/Wald substitution.** The uniqueness argument he attributes to
   Wald is actually the Salcedo argument, and Salcedo's conclusion is
   conditional on the renormalizability/effective-action criterion, not a
   consequence of the Wald axioms alone. Tichy–Flanagan show the opposite
   conclusion under the Wald axioms alone.

---

## 7. Blocked items

The following cannot be verified from the accessible material and are not
used as load-bearing premises.

| Item | Status | Consequence |
|---|---|---|
| Exact list of Wald 1977's five conditions | BLOCKED: full text paywalled; abstract does not list them. | The “five conditions” count is verified; the exact wording is not. Agent B's four-axiom list is falsified by the count alone. |
| Detailed page-level equations of Wald 1978 | BLOCKED: full text paywalled; abstract verified. | The structural claim (gradient of local curvature term; subtract metric times term; nonzero trace) is SOURCE-REPORTED from the abstract. The explicit coefficient is verified instead against Décanini–Folacci Eq. (70), which is fully accessible. |
| Exact journal page numbers of individual Hollands–Wald statements | BLOCKED: arXiv version used; journal pagination not inspected. | Citations use section numbers, theorem numbers and equation numbers of arXiv:gr-qc/0404074v2. |
| Content of Agent B's ref. 7, “Quantum fields on self-similar spacetimes” (arXiv:2509.10360) | BLOCKED: not inspected beyond title/arXiv ID. | The reference is irrelevant to the conservation theorem; it is cited for a claim it does not support. |
| Full text of Radzikowski 1996 | BLOCKED: only metadata verified. | The Hadamard/\(\mu\)SC equivalence is used as a standard SOURCE-REPORTED result; the audit does not depend on the proof. |
| Full text of Salcedo 1999 beyond the arXiv version | Not needed: arXiv version read. | The argument is quoted from the arXiv version; the published version may differ in formatting only. |

---

## 8. What this verification establishes

1. Agent B's **headline equation** \(\nabla^\mu\langle T_{\mu\nu}\rangle_{\rm ren}=0\)
   is not supported by the derivation he gives. The primary sources support a
   corrected prescription with an explicit local \(g_{\mu\nu}v_1\) term.
2. The **central structural claim** of Wald 1978 and Décanini–Folacci is
   verified: the uncorrected point-split divergence is the gradient of a local
   curvature term, and subtracting that term times the metric restores
   conservation while producing a nonzero trace.
3. The **five-condition vs four-axiom** error and the **Salcedo/Wald**
   misattribution are verified bibliographic defects, not matters of taste.
4. The **finite-ambiguity** statement is CONDITIONAL: the four-term basis is
   the minimal local dimension-four basis under Salcedo's renormalizability
   criterion; the Wald axioms alone allow the infinite Tichy–Flanagan family.
5. The **trace anomaly** formula in Agent B's p. 13 is FALSIFIED by direct
   comparison with Wald 1978 and Décanini–Folacci Eq. (111). The correct
   formula is the standard conformal-scalar anomaly up to the \(\Box R\)
   ambiguity.

The load-bearing equations for the conservation proof are now fixed. The
next document derives the classical foundation independently.
