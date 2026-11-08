# 03 — Hadamard and point-splitting audit

Audit: QSG-T001, covariant conservation of the Hadamard-renormalized
stress-energy tensor. Author: Agent A (resumption, blind protocol per
`plans/plan5_review.md`). Date: 2026-11-06.

This document reconstructs the quantum Hadamard foundations from the frozen
definitions. Sections marked BLIND are derived without using the expected
correction, the residual value, or any numbered conservation result from the
literature. Comparison with primary sources follows the recorded locks.

Dispositions follow `01-primary-source-verification.md`:
DERIVED, SOURCE-REPORTED, INDEPENDENTLY REPRODUCED, CONDITIONAL, FALSIFIED,
BLOCKED. `PASS` is not used.

---

## 0. Frozen conventions (Step 2 of the execution plan)

Everything in §0 is definition, premise, or standard geometric identity. No
result about the remainder equation, the residual, or the correction is stated
or assumed here. The quarantine manifest (§0.7) lists what is excluded.

### 0.1 Geometry

- \((M,g)\): four-dimensional, smooth, globally hyperbolic Lorentzian manifold,
  no boundary. Signature \((-+++)\). Natural units \(c=\hbar=1\).
- \(\Box_g := g^{\mu\nu}\nabla_\mu\nabla_\nu\), Levi-Civita connection.
- Real scalar field, mass \(m > 0\), minimal coupling \(\xi = 0\).
- Fixed classical background; no backreaction.
- Riemann tensor (Hawking–Ellis slot convention), stated as the definition
  from which all commutators below follow:
  \[
  R^\rho{}_{\sigma\mu\nu}V^\sigma
  = (\nabla_\mu\nabla_\nu - \nabla_\nu\nabla_\mu)V^\rho .
  \tag{0.1}
  \]
  Ricci tensor: \(R_{\mu\nu} := R^\rho{}_{\mu\rho\nu}\).
  Consequences used verbatim later (definitional, no derivation needed):
  \[
  [\nabla_\mu,\nabla_\nu]V^\rho = R^\rho{}_{\sigma\mu\nu}V^\sigma,
  \qquad
  [\nabla_\mu,\nabla_\nu]\alpha_\rho = -R^\sigma{}_{\rho\mu\nu}\alpha_\sigma,
  \qquad
  [\nabla_\mu,\nabla_\nu]f = 0
  \]
  for vector \(V\), covector \(\alpha\), scalar \(f\).

### 0.2 Two-point functions (objects 1–3, 5)

Work in a geodesically convex neighborhood \(\mathcal C\) of the diagonal.

1. **Wightman function.**
   \[
   G^+(x,x') := \langle \phi(x)\phi(x') \rangle .
   \]
   Premise P1 (free field in a Hadamard state): \(G^+\) satisfies the
   Klein–Gordon equation in **each** argument:
   \[
   (\Box_x - m^2)G^+(x,x') = 0,
   \qquad
   (\Box_{x'} - m^2)G^+(x,x') = 0.
   \tag{0.2}
   \]

2. **Reversed Wightman function.**
   \[
   G^-(x,x') := G^+(x',x) = \langle \phi(x')\phi(x) \rangle .
   \]
   By (0.2) applied with arguments exchanged, \(G^-\) also satisfies
   \((\Box_x - m^2)G^- = 0\) and \((\Box_{x'} - m^2)G^- = 0\).

3. **Symmetric two-point function.** Exact sum, no factor of one-half:
   \[
   \boxed{
   G^{(1)}(x,x') := G^+(x,x') + G^-(x,x')
   = \langle\{\phi(x),\phi(x')\}\rangle .
   }
   \tag{0.3}
   \]
   Symmetric: \(G^{(1)}(x,x') = G^{(1)}(x',x)\). From P1, term by term:
   \[
   \boxed{
   (\Box_x - m^2)G^{(1)}(x,x') = 0,
   \qquad
   (\Box_{x'} - m^2)G^{(1)}(x,x') = 0.
   }
   \tag{0.4}
   \]
   Disposition: DERIVED (one line from P1).

4. **Smooth remainder.** With \(H^{(1)}\) the parametrix of §0.3:
   \[
   \boxed{
   S(x,x') := \tfrac12\big[G^{(1)}(x,x') - H^{(1)}(x,x')\big].
   }
   \tag{0.5}
   \]
   Premise P2 (Hadamard state): \(S\) is smooth (\(C^\infty\); at least \(C^4\),
   the order consumed by the stress-energy operator plus one divergence) on
   \(\mathcal C\). Writing \(H_{\rm sub} := \tfrac12H^{(1)}\), (0.4) gives the
   exact relation
   \[
   \boxed{
   (\Box_x - m^2)S = -(\Box_x - m^2)H_{\rm sub},
   \qquad
   (\Box_{x'} - m^2)S = -(\Box_{x'} - m^2)H_{\rm sub}.
   }
   \tag{0.6}
   \]
   Disposition: DERIVED from P1+P2. The VALUE of the right-hand side is the
   blind computation of §1; nothing about it is assumed here.

### 0.3 Hadamard parametrix (object 4)

Synge's world function \(\sigma(x,x')\) (half the squared geodesic distance)
and the Van Vleck–Morette determinant \(\Delta(x,x')\). Frozen Synge-calculus
identities (standard geometric facts, used verbatim):
\[
[\sigma]=0,\quad
[\sigma_{;\mu}]=0,\quad
[\sigma_{;\mu'}]=0,\quad
[\sigma_{;\mu\nu}]=g_{\mu\nu},\quad
[\sigma_{;\mu\nu'}]=-g_{\mu\nu},\quad
\sigma^{;\mu}\sigma_{;\mu}=2\sigma,
\]
\[
[\Box_x\sigma]=4,\qquad [\Delta^{1/2}]=1,\qquad [\nabla_\mu\Delta^{1/2}]=0,
\tag{0.7}
\]
where \([\cdot]\) is the coincidence limit \(x'\to x\) with primed indices
parallel-transported to \(x\).

Parametrix ansatz at \(D=4\), with explicit prefactor:
\[
\boxed{
H^{(1)}(x,x')
= \frac{1}{8\pi^2}\left[
\frac{U(x,x')}{\sigma} + V(x,x')\ln|\bar\sigma|
\right],
\qquad \bar\sigma := \sigma/\ell^2,
}
\tag{0.8}
\]
where \(\ell > 0\) is a fixed renormalization length (its dependence is
tracked as finite ambiguity; nothing below fixes \(\ell\)), and
\[
U = \Delta^{1/2}\sum_{n\ge0}u_n\sigma^n,\quad u_0=1,
\qquad
V = \sum_{n\ge0}v_n\sigma^n,
\tag{0.9}
\]
are smooth biscalars whose coefficients \(u_n, v_n\) are local curvature
polynomials fixed by the following defining property DP: the coefficients are
the unique local expressions such that \((\Box_x-m^2)H^{(1)}\) contains no
\(\sigma^{-2}\), \(\sigma^{-1}\), or \(\ln|\bar\sigma|\) terms (i.e. it is
smooth on \(\mathcal C\) minus the diagonal, extending smoothly to it).

**No explicit coefficient values are assumed in this freeze.** In particular
no explicit \(u_n\) or \(v_n\) appears anywhere in §0. Whatever coefficient
values the blind derivation needs are derived there from DP and (0.7).

### 0.4 Bi-differential operator (object 6, with recheck)

Frozen operator, quoted from `02` Eq. (7.1):
\[
\boxed{
\begin{aligned}
T_{\mu\nu}(x,x')
={}&
\frac12\left(
g_\nu{}^{\nu'}\nabla_\mu\nabla_{\nu'}
+
g_\mu{}^{\mu'}\nabla_\nu\nabla_{\mu'}
\right)\\
&-\frac12 g_{\mu\nu}
\left(
g^{\rho\sigma'}\nabla_\rho\nabla_{\sigma'}+m^2
\right),
\end{aligned}}
\tag{0.10}
\]
with \(g_\mu{}^{\mu'}\) the parallel-transport bivector. \(T_{\mu\nu}=T_{\nu\mu}\)
term by term. Recheck V-02-71 (carried-dependency verification, classical,
non-blind): for \(F(x,x')=\phi(x)\phi(x')\),
\[
\nabla_\mu\nabla_{\nu'}F = \nabla_\mu\phi(x)\,\nabla_{\nu'}\phi(x'),
\]
and \([g_\nu{}^{\nu'}\nabla_{\nu'}\phi]=\nabla_\nu\phi\), so the first line of
(0.10) has coincidence limit
\(\tfrac12(\nabla_\mu\phi\nabla_\nu\phi+\nabla_\nu\phi\nabla_\mu\phi)
=\nabla_\mu\phi\nabla_\nu\phi\);
the second gives
\(-\tfrac12g_{\mu\nu}(\nabla^\rho\phi\nabla_\rho\phi+m^2\phi^2)\).
Hence \(\lim_{x'\to x}T_{\mu\nu}(x,x')F\) is exactly the classical (3.4).
For general symmetric biscalar \(B\), \(\lim T_{\mu\nu}B\) is the same form
evaluated on \(B\)'s coincidence data:
\([\nabla_\mu\nabla_{\nu'}B]-\tfrac12g_{\mu\nu}
([\nabla^\rho\nabla_{\rho'}B]+m^2[B])\).
Disposition: DERIVED (recheck of `02` §7; the only inputs are smoothness and
\([g_\mu{}^{\mu'}]=\delta_\mu^\mu\)).

### 0.5 Subtraction prescription (structural)

\[
\boxed{
\langle T_{\mu\nu}\rangle_{\rm raw}
:= \lim_{x'\to x} T_{\mu\nu}(x,x')\,S(x,x').
}
\tag{0.11}
\]
The full renormalized tensor is \(\langle T_{\mu\nu}\rangle_{\rm raw}\) plus a
correction to be derived. **No form of the correction is assumed.**

### 0.6 Derivative and coincidence conventions (objects 7–8)

- \(\nabla_\mu\): unprimed (\(x\)) derivative; \(\nabla_{\mu'}\): primed
  (\(x'\)) derivative. \(\Box\) acts unprimed unless marked \(\Box'\).
- \([\cdot]\): coincidence limit with primed indices parallel-transported.
- Synge's rule (frozen coincidence-derivative convention): for any biscalar
  \(B\),
  \[
  \boxed{
  \nabla_\mu[B] = [\nabla_\mu B + g_\mu{}^{\mu'}\nabla_{\mu'}B],
  }
  \tag{0.12}
  \]
  and correspondingly for higher derivatives (each total derivative is the
  coincidence limit of the sum of the unprimed and transported primed
  derivative). A total derivative of a coincidence limit is NEVER replaced by
  differentiation in one argument alone.
- Curvature-sign convention: (0.1). All commutators derive from it.

### 0.7 Quarantine manifest

The following appear NOWHERE in §0 and are assumed nowhere in the blind
sections (§1–§2): any value for \((\Box-m^2)H^{(1)}\), \((\Box-m^2)S\), or
their coincidence limits; any explicit \(u_n\) or \(v_n\); any residual,
correction, coefficient, or sign for the conservation argument; any numbered
conservation equation from any literature source; any content of Agent B
pp. 10–13. The blind sections cite only §0, premises P1/P2/DP, and their own
previous equations.

---

## 1. BLIND: the remainder equation

This section uses only §0, premises P1/P2/DP, the geometric lemmas L1–L2
stated below, and its own equations.

### 1.1 Geometric lemmas (imported standard theory)

Lemma L1 (Van Vleck–Morette transport, SOURCE-REPORTED): on \(\mathcal C\),
\[
\sigma^{;\mu}\nabla_\mu\Delta^{1/2}
+ \tfrac12(\Box_x\sigma - 4)\Delta^{1/2} = 0.
\tag{L1}
\]
Standard Synge identity; pure background geometry. Findability pointer: any
account of the Hadamard construction (e.g. Fulling, Aspects of QFT in Curved
Spacetime; Birrell–Davies §6.2). Only the verbatim statement is used.

Lemma L2 (leading-transport consequence, SOURCE-REPORTED): let
\(M := [(\Box_x-m^2)\Delta^{1/2} - [(\Box_x-m^2)\Delta^{1/2}]]
+ 2(v_0-[v_0]) + \tau v_0 + 2D_\sigma v_0\),
where \(\tau := \Box_x\sigma - 4\), \(D_\sigma := \sigma^{;\mu}\nabla_\mu\),
and \(v_0\) is the leading coefficient of (0.9) satisfying the standard
\(v_0\)-transport equation. Then \(M/\sigma\) extends smoothly to the diagonal
with value \(0\):
\[
\lim_{x'\to x} M/\sigma = 0.
\tag{L2}
\]
This is the coincidence content of the \(v_0\)-transport equation. Only the
verbatim statement is used.

Geometric data G2 (SOURCE-REPORTED): \([(\Box_x-m^2)\Delta^{1/2}] = R/6 - m^2\).
Standard Van Vleck coincidence expansion. Only the verbatim value is used.

Boundary of the import: L1, L2, G2 contain no statement about
\((\Box-m^2)S\), any divergence, any residual, or any correction.

### 1.2 The operator applied to the ansatz

For \(f(\sigma)\) smooth away from \(\sigma=0\) and smooth biscalar \(A\),
Leibniz plus \(\sigma^{;\mu}\sigma_{;\mu}=2\sigma\) gives
\[
\Box_x(fA) = f''(2\sigma)A + f'(\Box_x\sigma)A
+ 2f'\sigma^{;\mu}\nabla_\mu A + f\Box_x A.
\tag{1.1}
\]
DERIVED. With \(f=1/\sigma\) (\(f'=-1/\sigma^2\), \(f''=2/\sigma^3\)):
\[
\boxed{
\Box_x(U/\sigma)
= \frac{(4-\Box_x\sigma)U - 2\sigma^{;\mu}U_{;\mu}}{\sigma^2}
+ \frac{\Box_x U}{\sigma}.
}
\tag{1.2}
\]
With \(f=\ln|\bar\sigma|\) (\(f'=1/\sigma\), \(f''=-1/\sigma^2\); the scale
\(\ell\) drops out of every derivative):
\[
\boxed{
\Box_x(V\ln|\bar\sigma|)
= \frac{(\Box_x\sigma-2)V + 2\sigma^{;\mu}V_{;\mu}}{\sigma}
+ (\Box_x V)\ln|\bar\sigma|.
}
\tag{1.3}
\]
Hence, with \(\tau := \Box_x\sigma-4\),
\[
(\Box_x-m^2)H^{(1)}
= \frac{1}{8\pi^2}\left[
\frac{N_2}{\sigma^2} + \frac{N_1}{\sigma} + N_L\ln|\bar\sigma|
\right],
\tag{1.4}
\]
\[
\begin{aligned}
N_2 &:= -\tau U - 2\sigma^{;\mu}U_{;\mu},\\
N_1 &:= (\Box_x-m^2)U + (2+\tau)V + 2D_\sigma V,\\
N_L &:= (\Box_x-m^2)V.
\end{aligned}
\tag{1.5}
\]
DERIVED from (0.8) and (1.1)–(1.3).

### 1.3 Minimal-\(U\) adoption

Write \(U=\Delta^{1/2}\bar u\) with \(\bar u = 1+u_1\sigma+\cdots\).
By L1, \(-\tau\Delta^{1/2}-2D_\sigma\Delta^{1/2}=0\) exactly, so
\(N_2 = -2\Delta^{1/2}D_\sigma\bar u\).
Adopt the minimal representative \(U := \Delta^{1/2}\) (i.e. \(\bar u\equiv1\)):
then \(N_2 \equiv 0\). Justification: the leading singularity
\(\Delta^{1/2}/(8\pi^2\sigma)\) is universal (fixed by the flat-space
singularity, part of ansatz normalization); the \(\sigma^{-1}\) obstruction
that \(u_{n\ge1}\) would address is instead absorbed by the \(V\ln|\bar\sigma|\)
term, whose coefficients are then uniquely fixed by DP order by order
(§1.4–§1.5); two DP-satisfying representatives differ by terms smooth enough
to be absorbed into \(S\). Disposition: DERIVED-choice.

With \(U=\Delta^{1/2}\) and \(V=v_0+v_1\sigma+R_2\)
(\(R_2=O(\sigma^2)\) smooth):
\[
N_1 = (\Box_x-m^2)\Delta^{1/2} + (2+\tau)V + 2D_\sigma V,
\qquad
N_L = (\Box_x-m^2)V.
\tag{1.6}
\]

### 1.4 Pole cancellation: \([v_0]\)

DP requires \(N_1/\sigma\) smooth, hence \([N_1]=0\) (else a \(\sigma^{-1}\)
pole). Since \([\tau]=[D_\sigma V]=0\):
\[
\boxed{
[v_0] = -\tfrac12[(\Box_x-m^2)\Delta^{1/2}]
= \frac{m^2}{2} - \frac{R}{12}.
}
\tag{V0}
\]
DERIVED from DP + (1.6) + G2. In particular \([v_0]\) is a local scalar of
mass dimension 2. No other coefficient value is used or derived here.

### 1.5 Log cancellation: the \([v_1]\) relation

DP requires \(N_L\ln|\bar\sigma|\) bounded, hence \([N_L]=0\) (else a
logarithmic blow-up). Now
\((\Box_x-m^2)(v_1\sigma)
= \sigma(\Box_x-m^2)v_1 + 2D_\sigma v_1 + v_1(4+\tau)\),
whose coincidence value is \(4[v_1]\), while
\([(\Box_x-m^2)R_2]=0\). Therefore
\[
\boxed{
[N_L] = [(\Box_x-m^2)v_0] + 4[v_1] = 0,
\qquad\text{i.e.}\quad
[v_1] = -\tfrac14[(\Box_x-m^2)v_0].
}
\tag{V1a}
\]
DERIVED from DP + Leibniz. So \([v_1]\) is a local scalar of mass dimension 4
(\([v_0]\) has dimension 2, \(\Box-m^2\) adds 2). Its curvature expansion is
NOT derived here and is NOT needed for the conservation argument.

Moreover \(N_L\ln|\bar\sigma|\to0\) at coincidence (\(N_L\) smooth and
vanishing there, times \(\ln|\bar\sigma|\)).

### 1.6 The remainder limit

Expand \(N_1/\sigma\) using \(V=v_0+v_1\sigma+R_2\):
\[
\frac{N_1}{\sigma}
= \frac{M}{\sigma}
+ 2v_1 + 2D_\sigma v_1 + 4v_1 + o(1),
\]
where \(M\) is exactly the combination of L2 and \(o(1)\) collects
\(R_2\)-terms vanishing at coincidence. The terms: \(2v_1\sigma/\sigma\to
2[v_1]\); \(2D_\sigma(v_1\sigma)/\sigma = 2D_\sigma v_1 + 4v_1 \to 4[v_1]\).
By L2, \(M/\sigma\to0\). Hence
\[
\boxed{
\lim_{x'\to x} N_1/\sigma = 6[v_1].
}
\tag{1.7}
\]
DERIVED from (1.6) + V0 + L2.

The smooth source term. From (0.6), (1.4), \(N_2\equiv0\), (1.7), and
\(N_L\ln|\bar\sigma|\to0\):
\[
\boxed{
(\Box_x-m^2)S = J,\qquad
J := -\frac{1}{16\pi^2}\left[\frac{N_1}{\sigma}+N_L\ln|\bar\sigma|\right]
\ \text{smooth},\qquad
[J] = -\frac{6[v_1]}{16\pi^2} = -\frac{3[v_1]}{8\pi^2}.
}
\tag{1.8}
\]
DERIVED. By \(x\leftrightarrow x'\) symmetry (\(S\), \(H^{(1)}\) symmetric),
\((\Box_{x'}-m^2)S\) has the same coincidence limit.

### 1.7 Falsification of the homogeneous equation

Flat-space lemma (DERIVED by Poincaré invariance of the Minkowski Hadamard
construction): in Minkowski space \(v_0\equiv[v_0]=m^2/2\) is constant, so
\([(\Box-m^2)v_0] = -m^4/2\) and by (V1a) \([v_1] = m^4/8 \ne 0\) (since
\(m>0\)). Therefore, in flat space already,
\[
[(\Box-m^2)S] = -\frac{3m^4}{64\pi^2} \ne 0.
\tag{1.9}
\]
**Disposition: the assertion that a smooth Hadamard remainder satisfies the
homogeneous Klein–Gordon equation is FALSIFIED.** Smoothness (\(S\in
C^\infty\)) does not imply \((\Box-m^2)S=0\); the Minkowski vacuum is an
explicit counterexample. The remainder obeys the inhomogeneous equation (1.8)
with a generically nonvanishing smooth source.

### 1.8 Lock-1 scope checkpoint

Derived under lock 1 (§0 + L1/L2/G2 + this section): the exact relation
(0.6); the ansatz calculus (1.2)–(1.5); minimal-\(U\) adoption; (V0); (V1a);
the remainder equation and limit (1.8); the falsification (1.9). Assumed from
outside: P1, P2, DP, L1, L2, G2 (all recorded with dispositions; none
contains the residual, divergence, or correction).

---

## 2. BLIND: smoothness, coincidence limits, and prescription audit

### 2.1 What smoothness gives and does not give

P2 gives: \(S\) smooth, hence (0.11) well-defined; all \(s\)-tensors (§2.2)
exist; \(J\) smooth with limit (1.8). DP gives: \((\Box-m^2)H^{(1)}\) smooth,
hence (0.6) is an equation between smooth objects. Neither gives a
homogeneous equation for \(S\): \(H^{(1)}\) is built from local geometry to
match the universal singularity, and the \(\ln|\bar\sigma|\) term forced by
the \(\sigma^{-1}\) obstruction leaves the smooth leftover (1.8). Regularity
is not dynamics. Disposition: DERIVED (conceptual consequence of §1).

### 2.2 \(s\)-tensor notation and Synge consequences

For the divergence computation, freeze the notation (definitional):
\[
s:=[S],\quad s_\mu:=[\nabla_\mu S],\quad
s_{\mu\nu}:=[\nabla_\mu\nabla_\nu S],\quad
s_{\mu\nu'}:=[\nabla_\mu\nabla_{\nu'}S]\ (\nu'\ \text{transported}),
\]
\[
t := g^{\mu\nu}s_{\mu\nu'} = [\nabla^\mu\nabla_{\mu'}S].
\tag{2.1}
\]
\(S\) symmetric \(\Rightarrow\) \(s_\mu=[\nabla_{\mu'}S]\),
\(s_{\mu\nu'}=s_{\nu\mu'}\). From Synge's rule (0.12), DERIVED:
\[
\nabla_\mu s = 2s_\mu,
\qquad
\nabla_\mu s_\nu = s_{\mu\nu} + s_{\nu\mu'},
\qquad
\Box s = [\Box_x S] + t.
\tag{2.2}
\]
The last: \(\Box s = g^{\mu\nu}\nabla_\mu s_\nu
= g^{\mu\nu}s_{\mu\nu} + g^{\mu\nu}s_{\nu\mu'} = [\Box_x S] + t\).
Every coincidence-limit derivative in the divergence computation cites
(0.12); total derivatives are never replaced by single-argument derivatives.

### 2.3 Failure-mode analysis (abstract)

Replacing \(\nabla_\mu[\cdot]\) by \([\nabla_\mu\cdot]\) drops the transported
primed term (\(s'\)-type halves above): factors of 2 are lost and
curvature/transport pieces carried by the primed derivative go missing. In a
divergence computation this corrupts both the \(s\)-transport terms and the
residual identification. Audit rule enforced in the sequel: each step moving
\(\nabla\) across \([\cdot]\) names the Synge instance. Page-anchored
verification of prior work against this rule is post-lock (doc 05).

### 2.4 Prescription consistency checks

(C1) Well-definedness: (0.11) exists because \(S\) is \(C^4\) by P2 and
\(T_{\mu\nu}(x,x')\) has smooth coefficients on \(\mathcal C\). DERIVED.
(C2) Symmetry: \(T_{\mu\nu}=T_{\nu\mu}\) (§0.4) and \(S\) symmetric imply
\(\langle T_{\mu\nu}\rangle_{\rm raw}\) symmetric. DERIVED.
(C3) Dimensions: \([S]=2\) (\([G^{(1)}]=[H^{(1)}]=2\)), \(T\) carries
dimension 2 (two derivatives or \(m^2\)), so \([\langle
T_{\mu\nu}\rangle_{\rm raw}]=4\). DERIVED.
Conservation is NOT checked here; it is the blind computation of doc 04.

---

## Lock 1

- Timestamp: 2026-11-06T15:54:55Z.
- Scope: file lines 1–510 (§0 preamble through §2.4), i.e. everything above
  this header.
- SHA-256 of scope bytes:
  `2e861438fc6247061fa1e509ab2f815e317a910b84a018a3e301d9e5a7c30db0`.
- Quote check on scope: no `Eq. (70)`-type literature citations, no
  Décanini/Wald content citations, no Agent B content. Mechanical grep hits
  were: (i) quarantine-manifest meta-mentions (§0.7, structural, not content);
  (ii) `` `02` Eq. (7.1) `` — the frozen classical operator, explicitly
  allowed non-blind material; (iii) LaTeX `v_1`/`[v_1]` — this document's own
  formal coefficient notation, never an imported value (no explicit curvature
  polynomial for any \(v_n\) appears in scope; \([v_0]\) in (V0) is derived,
  not imported).
- Quarantine statement: §§0–2 cite only §0, premises P1/P2/DP, lemmas
  L1/L2/G2 (pure background geometry, no residual/divergence/correction
  content), and their own equations. Comparison with the literature begins in
  §3, after this lock.

---

## 3. POST-LOCK: comparison with primary sources

Comparison basis: Décanini–Folacci 2008 full text, inspected this session via
the ar5iv HTML rendering of arXiv:gr-qc/0512118
(`https://ar5iv.labs.arxiv.org/html/gr-qc/0512118`, redirect-resolved and
read at equation level). Equation numbers below are DF's. Dispositions apply
to the agreement between the locked blind result and DF.

### 3.1 Convention reconciliation (all inspected)

(R1) World function: DF Eq. (23) gives \(2\sigma=\sigma^{;\mu}\sigma_{;\mu}\)
with "\(2\sigma\) the square of the geodesic distance", i.e. DF's \(\sigma\)
is Synge's world function — identical to §0.3. INDEPENDENTLY REPRODUCED
(convention agreement).
(R2) Propagator: DF define \(G^F=i\langle T\phi\phi'\rangle\) (their §II.2);
hence \(-iG^F=\langle T\phi\phi'\rangle\) and
\(\mathrm{Re}\langle T\phi\phi'\rangle=\tfrac12G^{(1)}\). Their
renormalization replacement (their Eq. (63)),
\(\lim\mathcal T[-iG^F]\to(\alpha_D/2)\lim\mathcal TW\), therefore
regularizes the time-ordered function, whose real symmetric part is the
framework of §0. DERIVED (reconciliation).
(R3) Operator: DF Eq. (61) at \(\xi=0\),
\(\mathcal T_{\mu\nu}=g_\nu{}^{\nu'}\nabla_\mu\nabla_{\nu'}
-\tfrac12g_{\mu\nu}(g^{\rho\sigma'}\nabla_\rho\nabla_{\sigma'}+m^2)\),
is the unsymmetrized form of (0.10); on symmetric biscalars both have
identical coincidence limits (one-line symmetry argument). INDEPENDENTLY
REPRODUCED.
(R4) Coefficients: DF Eqs. (108a)/(109) give \(v_0=(1/2)m^2+
(1/2)(\xi-1/6)R\) and \(v_1=(1/8)m^4+\cdots\); at \(\xi=0\),
\([v_0]=m^2/2-R/12\) (matches (V0) exactly) and flat-space
\([v_1]=m^4/8\) (matches §1.7 exactly). INDEPENDENTLY REPRODUCED
(\(v_0\) fully; \(v_1\) in flat space; curved-space \(v_1\) polynomial
not derived here).
(R5) Series structure: DF Eq. (29a), \(U=\sum_{n=0}^{D/2-2}U_n\sigma^n\),
gives \(U=U_0=\Delta^{1/2}\) at \(D=4\) — identical to the §1.3 minimal
choice; \(V=\sum_{n\ge0}V_n\sigma^n\) matches (0.9). INDEPENDENTLY
REPRODUCED.

### 3.2 Remainder equation: structural agreement, coefficient conditional

DF Eqs. (52)/(55) (inspected verbatim, confirming the `01` transcription):
\((\Box-m^2-\xi R)W=-(D+2)V_1-2V_1{}^{;\mu}\sigma_{;\mu}+O(\sigma)\),
i.e. \(-6v_1\) at coincidence (\(D=4\), \(\xi=0\)).
The blind result (1.8): \([(\Box-m^2)S]=-3[v_1]/(8\pi^2)\).

Structural agreement — INDEPENDENTLY REPRODUCED: both sides derive an
inhomogeneous remainder equation with a \(v_1\)-driven smooth source, and
both falsify the homogeneous claim (DF's RHS \(\ne0\) coincides with the
blind (1.9)).

Coefficient reconciliation — CONDITIONAL: from DF's (63)+(48),
\(W=8\pi^2S+A\) with \(A\) smooth (log-scale \(V\ln\ell^2\) plus the
\(i\epsilon\)/direction-averaging terms of the Feynman route). The two
independently verified limits (\(-6v_1\) and \(-3[v_1]/(8\pi^2)\)) then
jointly DERIVE the constraint \([(\Box-m^2)A]=-3[v_1]\): DF's \(W\) and
\(8\pi^2S\) differ by a smooth term with computable nonvanishing
equation-limit. The exact determination of \(A\) (from DF's averaging
terms) is a bounded open computation. It is NOT a blocker for doc 04:
doc 04 uses (1.8) self-consistently (residual and correction derive from
the same remainder equation, so their cancellation is exact in §0
conventions regardless of \(A\)), and carries this conditional into its
own literature comparison.

Record: this subtlety (two valid parametrix routes differing by smooth
terms with nonvanishing equation-limits) is precisely what the blind
protocol exists to catch; targeting DF's \(-6v_1\) directly would have
laundered Feynman-route averaging terms into the symmetric framework.

### 3.3 `01` transcription check

The DF equations recorded in `01` §4.1 and used as navigation —
(52), (55), (61), (67)–(70), (108a), (109) — were verified verbatim
against the inspected full text (numbers and content match).
Disposition: `01`'s DF record is INDEPENDENTLY REPRODUCED at every
equation reinspected in §3.1–§3.2.

---

## 4. What this document establishes

1. Remainder equation (1.8): DERIVED (blind) — \((\Box_x-m^2)S=J\),
   \(J\) smooth, \([J]=-3[v_1]/(8\pi^2)\).
2. Homogeneous-remainder claim: FALSIFIED (§1.7, Minkowski counterexample).
3. Coefficient \([v_0]=m^2/2-R/12\): DERIVED (V0); matches DF (108a).
4. \([v_1]\): formal local scalar of dimension 4 (V1a); flat value
   \(m^4/8\) DERIVED; curved polynomial not needed, not derived.
5. Prescription (0.11): well-defined, symmetric, dimension-4 (C1–C3).
6. Synge calculus (2.2) and audit rule (§2.3): DERIVED, binding on doc 04.
7. DF comparison: conventions + structure INDEPENDENTLY REPRODUCED;
   remainder-coefficient match CONDITIONAL on the smooth-difference term
   \(A\) (§3.2); `01` DF transcription verified.
8. Carries into doc 04: (1.8), (2.1)–(2.2), (V0), (V1a), L2 (all locked).

Open (non-blocking): explicit determination of \(A\) (§3.2).

---

## Erratum E1 (post-lock correction, 2026-11-06)

§2.2's `Box s = [Box_x S] + t` is wrong by a factor of 2. Since
\(\nabla_\nu s = 2s_\nu\), \(\Box s = \nabla^\mu(2s_\mu)
= 2([\Box_x S] + t)\). Correct formula:
\(\boxed{\Box s = 2[\Box_x S] + 2t}\).

Found by the Step-6 independent cross-check (quarantined agent without
audit-directory access; it derived the factor independently and flagged
the inconsistency — the cross-check prompt itself contained the error,
which the agent caught rather than parroted). Verified by direct
recomputation.

Impact: the \(\Box s\) formula is used only in doc-04 §3.5 (trace (3.4));
corrected there by Erratum E2. Docs 03 §§0–1 and 04 §§1–2, §3.1–3.4 use
no \(\Box s\) and are unaffected. Lock 1 scope (lines 1–510) is unaltered;
this erratum is appended post-lock.

---

## Erratum E3 (normalization repair, 2026-11-06)

Adversarial review (`plans/plan5_imp_review.md` §1) found (0.5)
\(S=\tfrac12(G^{(1)}-H^{(1)})\) singular (\(1/(16\pi^2\sigma)\):
\(G^{(1)}\sim1/(4\pi^2\sigma)\) vs \(H^{(1)}\sim1/(8\pi^2\sigma)\)),
voiding P2 and every \(S\)-dependent number: (0.6)-with-\(H_{\rm sub}\),
(1.8)/(1.9)-values, and all of doc 04's numbers. Full record: doc 08
N1 (Minkowski verification, convention table, A-vs-B identity,
prescription freeze \(S:=\tfrac12G^{(1)}-H^{(1)}\)) and N2
(remainder reconstruction: \([J]=-3[v_1]/(4\pi^2)\), derived
\(v_0\)-transport + L2, re-derived \(\Box s\), flat falsification
re-run). STAND as written: Synge data, ansatz + DP, \(N\)-calculus,
minimal-\(U\), (V0), (V1a), (1.7)|L2, L1/G2, E1, and the §3 DF
convention reconciliation (R1–R5); §3.2's \(A\)-constraint is
superseded by N4.5 (\([(\Box-m^2)A]=0\), \([\nabla'J_A]=0\)).
Lock 1 scope (lines 1–510) is unaltered; this erratum is appended
post-lock. E1's \(\Box s\) correction is re-confirmed by fresh
derivation (N2.4).
