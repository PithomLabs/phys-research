# 04 — Conservation and counterterm derivation

Audit: QSG-T001. Author: Agent A (resumption, blind protocol per
`plans/plan5_review.md`). Date: 2026-11-06.

This document derives the divergence of the regularized stress-energy tensor,
identifies the residual, derives the conservation-restoring correction, shows
the cancellation, and classifies finite ambiguities. Sections marked BLIND use
only the locked doc-03 results, the geometric lemmas stated below, and their
own equations.

Locked carry-in from doc 03 (all under Lock 1): the remainder equation
\((\Box_x-m^2)S=J\) with \([J]=-3[v_1]/(8\pi^2)\) (1.8); \([v_0]=m^2/2-R/12\)
(V0); \([v_1]=-\tfrac14[(\Box-m^2)v_0]\) (V1a); \(s\)-tensor notation and
Synge consequences (2.1)–(2.2); the prescription (0.11) with operator (0.10).

Dispositions follow `01`: DERIVED, SOURCE-REPORTED, INDEPENDENTLY
REPRODUCED, CONDITIONAL, FALSIFIED, BLOCKED. `PASS` is not used.

---

## 1. BLIND: divergence of the regularized tensor

### 1.1 Geometric lemmas (imported standard theory)

Lemma L3 (Ricci identity, DERIVED from doc-03 (0.1)): for any covector
\(\alpha\),
\[
[\nabla^\mu,\nabla_\nu]\alpha_\mu = +R_\nu{}^\sigma\alpha_\sigma .
\tag{L3}
\]
Proof: \([\nabla_\alpha,\nabla_\beta]\alpha_\rho=-R^\sigma{}_{\rho\alpha\beta}
\alpha_\sigma\) (Leibniz on \(V^\rho\alpha_\rho\) with (0.1)), so
\([\nabla^\mu,\nabla_\nu]\alpha_\mu=-Q^\sigma{}_\nu\alpha_\sigma\) with
\(Q^\sigma{}_\nu=g^{\mu\lambda}R^\sigma{}_{\mu\lambda\nu}\).
Now \(Q^\sigma{}_\nu=g^{\sigma\rho}g^{\mu\lambda}R_{\rho\mu\lambda\nu}
=g^{\sigma\rho}g^{\mu\lambda}R_{\lambda\nu\rho\mu}\) (pair symmetry)
\(=g^{\sigma\rho}R_{\lambda\nu\rho}{}^\lambda
=-g^{\sigma\rho}R^\lambda{}_{\nu\lambda\rho}=-R^\sigma{}_\nu\)
(antisymmetry in the last pair; Ricci \(R_{\nu\rho}=R^\lambda{}_{\nu\lambda
\rho}\)). Hence the \(+\) sign. In particular
\([\Box,\nabla_\nu]f=+R_\nu{}^\sigma\nabla_\sigma f\) for scalar \(f\).

Lemma L4 (contracted Bianchi, SOURCE-REPORTED): \(\nabla^\mu R_{\mu\nu}
=\tfrac12\nabla_\nu R\). Standard; pure background geometry.

Boundary of the import: L3/L4 are background Riemannian geometry with no
renormalization content.

### 1.2 The raw tensor in \(s\)-tensors

From (0.10)–(0.11), taking the coincidence limit (all primed indices
transported per §0.6; no differentiated-transport terms arise since limits
are taken before differentiation):
\[
R_{\mu\nu} := \langle T_{\mu\nu}\rangle_{\rm raw}
= s_{\mu\nu'} - \tfrac12g_{\mu\nu}(t + m^2s),
\tag{1.1}
\]
with \(s_{\mu\nu'}=s_{\nu\mu'}\) (symmetry of \(S\)).
Transport note: the parallel-transport bivector enters only through
\([\cdot]\); every surviving object is an unprimed tensor at \(x\).

### 1.3 Divergence: reduction to \(s\)-transport

Differentiating (1.1) and using \(\nabla_\nu s=2s_\nu\) (2.2):
\[
\nabla^\mu R_{\mu\nu}
= \nabla^\mu s_{\mu\nu'} - \tfrac12\nabla_\nu t - m^2s_\nu .
\tag{1.2}
\]
Lemma T1 (DERIVED): \(\nabla_\nu t
= \nabla^\mu s_{\mu\nu'} + \nabla^\mu s_{\mu\nu} - 2m^2s_\nu - \nabla_\nu[J]
- R_\nu{}^\sigma s_\sigma\).
Proof: Synge on \(t=[\nabla^\mu\nabla_{\mu'}S]\) gives
\(\nabla_\nu t=[\nabla_\nu\nabla^\mu\nabla_{\mu'}S]
+[\nabla_{\nu'}\nabla^\mu\nabla_{\mu'}S]\).
First bracket: unprimed pair commutes on the primed covector
(\([\nabla_\nu,\nabla^\mu]\) acts as scalar commutator \(=0\));
\([\nabla^\mu\nabla_{\mu'}\nabla_\nu S]
=\nabla^\mu s_{\nu\mu'}-[\nabla^{\mu'}\nabla_{\mu'}\nabla_\nu S]\)
(Synge)
\(=\nabla^\mu s_{\nu\mu'}-[\nabla_\nu\Box' S]\)
(primed/unprimed commute)
\(=\nabla^\mu s_{\nu\mu'}-m^2s_\nu-\tfrac12\nabla_\nu[J]\)
(equation \((\Box'-m^2)S=J\); \(J\) symmetric so
\([\nabla_\nu J]=\tfrac12\nabla_\nu[J]\)).
Second bracket: \([\nabla^\mu\nabla_{\nu'}\nabla_{\mu'}S]
=\nabla^\mu s_{\mu\nu}-[\nabla^{\mu'}\nabla_{\nu'}\nabla_{\mu'}S]\)
(Synge; \([\nabla_{\nu'}\nabla_{\mu'}S]=s_{\mu\nu}\) by symmetry).
Here \([\nabla^{\mu'}\nabla_{\nu'}\nabla_{\mu'}S]\) is all-primed: by
\(x\leftrightarrow x'\) symmetry it equals \([\nabla^\mu\nabla_\nu\nabla_\mu S]
=[\nabla_\nu\Box S]+R_\nu{}^\sigma s_\sigma\) (L3:
\([\nabla^\mu,\nabla_\nu]\nabla_\mu S=+R_\nu{}^\sigma\nabla_\sigma S\))
\(=m^2s_\nu+\tfrac12\nabla_\nu[J]+R_\nu{}^\sigma s_\sigma\).
So the second bracket is
\(\nabla^\mu s_{\mu\nu}-m^2s_\nu-\tfrac12\nabla_\nu[J]-R_\nu{}^\sigma s_\sigma\).
Sum, with \(s_{\nu\mu'}=s_{\mu\nu'}\), is T1. ∎

Insert T1 in (1.2) (mass terms cancel):
\[
\boxed{
\nabla^\mu R_{\mu\nu}
= \tfrac12\nabla^\mu s_{\mu\nu'} - \tfrac12\nabla^\mu s_{\mu\nu}
+ \tfrac12\nabla_\nu[J] + \tfrac12R_\nu{}^\sigma s_\sigma.
}
\tag{D1}
\]
DERIVED. Every \(\nabla\)-across-\([\cdot]\) step above cites Synge (0.12).

### 1.4 The transport identity

Lemma TI (DERIVED): \(\nabla^\mu s_{\mu\nu'}
= \nabla_\mu s^\mu{}_\nu - R_\nu{}^\sigma s_\sigma\).
Proof: Synge on \(s_{\mu\nu'}=[\nabla_\mu\nabla_{\nu'}S]\):
\(\nabla^\mu s_{\mu\nu'}=[\Box_x\nabla_{\nu'}S]
+[\nabla^{\mu'}\nabla_\mu\nabla_{\nu'}S]\).
First: \([\nabla_{\nu'}\Box_x S]=m^2s_\nu+\tfrac12\nabla_\nu[J]\)
(different-argument commute; equation; \(J\) symmetric).
Second: \([\nabla_\mu\nabla^{\mu'}\nabla_{\nu'}S]
=[\nabla_\mu V_{\nu'}]\) with \(V_{\nu'}=\nabla^{\mu'}\nabla_{\nu'}S\);
Synge gives \(\nabla_\mu[V_{\nu'}]-[\nabla_{\mu'}V_{\nu'}]
=\nabla_\mu s^\mu{}_\nu-[\Box'\nabla_{\nu'}S]\)
(\([V_{\nu'}]=s^\mu{}_\nu\) by \(x\leftrightarrow x'\) symmetry).
Now \([\Box'\nabla_{\nu'}S]=[\nabla_{\nu'}\Box' S]+R_\nu{}^\sigma s_\sigma\):
the primed commutator \([\Box',\nabla_{\nu'}]\) on \(S\) contributes
\(+R_\nu{}^\sigma s_\sigma\) by L3 in primed slots (transported), since
\([\nabla^{\mu'},\nabla_{\nu'}]\nabla_{\mu'}S
=+R_{\nu'}{}^{\sigma'}\nabla_{\sigma'}S\).
So the second term is
\(\nabla_\mu s^\mu{}_\nu-m^2s_\nu-\tfrac12\nabla_\nu[J]-R_\nu{}^\sigma s_\sigma\).
The \(m^2s_\nu\) and \(\tfrac12\nabla_\nu[J]\) pieces cancel in the sum,
leaving the Ricci term. ∎

Since \(\nabla_\mu s^\mu{}_\nu=\nabla^\mu s_{\mu\nu}\) (metric
compatibility), TI gives
\(\tfrac12\nabla^\mu s_{\mu\nu'}-\tfrac12\nabla^\mu s_{\mu\nu}
=-\tfrac12R_\nu{}^\sigma s_\sigma\), which cancels the explicit
\(+\tfrac12R_\nu{}^\sigma s_\sigma\) in (D1): the curvature terms enter
with opposite signs and cancel.

### 1.5 The residual

From (D1) and TI:
\[
\boxed{
\nabla^\mu\langle T_{\mu\nu}\rangle_{\rm raw}
= \tfrac12\nabla_\nu[J]
= -\frac{3}{16\pi^2}\,\nabla_\nu[v_1].
}
\tag{1.3}
\]
DERIVED from locked (1.8). The raw point-split divergence is the gradient
of the local curvature scalar \(-3[v_1]/(32\pi^2)\) (i.e. a gradient of local
geometry, with coefficient fixed by the blind remainder computation).
Flat-space check: \([J]\) constant \(\Rightarrow\) residual zero, consistent
with translation invariance (the residual is curvature-driven). The exchange
of \(\nabla^\mu_x\) with \(\lim_{x'\to x}\) is justified by smoothness on
\(\mathcal C\) (doc 03, P2).

---

## 2. BLIND: correction and cancellation

### 2.1 The correction

(1.3) is a pure gradient of a local scalar. It is cancelled by a local
metric term. With the ansatz \(C_{\mu\nu}=K\,g_{\mu\nu}[v_1]\) (\(K\)
constant; \([v_1]\) the locked local scalar of dimension 4):
\(\nabla^\mu C_{\mu\nu}=K\,\nabla_\nu[v_1]\) (metric compatibility).
Cancellation of (1.3) requires \(K=+3/(16\pi^2)\):
\[
\boxed{
C_{\mu\nu} = +\frac{3}{16\pi^2}\,g_{\mu\nu}[v_1],
\qquad
\nabla^\mu C_{\mu\nu} = +\frac{3}{16\pi^2}\,\nabla_\nu[v_1].
}
\tag{2.1}
\]
DERIVED. \(C_{\mu\nu}\) is local, state-independent, geometric, of mass
dimension 4.

### 2.2 The correction is not independently conserved

From (2.1): \(\nabla^\mu C_{\mu\nu}=+3\nabla_\nu[v_1]/(16\pi^2)\), which is
generically nonzero (it must be: its nonvanishing is what cancels (1.3)).
**\(C_{\mu\nu}\) is a prescription-restoring term, not a separately
conserved tensor.** Any description of the conservation-restoring correction
as independently conserved while its divergence cancels the residual is
contradictory. DERIVED (explicit).

### 2.3 Renormalized tensor and conservation

\[
\boxed{
\langle T_{\mu\nu}\rangle_{\rm ren}
:= R_{\mu\nu} + C_{\mu\nu} + \Theta_{\mu\nu},
\qquad \nabla^\mu\Theta_{\mu\nu}=0,
}
\tag{2.2}
\]
where \(\Theta_{\mu\nu}\) is the finite conserved ambiguity (§3).
Theorem (blind conservation): in the frozen model,
\[
\boxed{
\nabla^\mu\langle T_{\mu\nu}\rangle_{\rm ren}
= -\frac{3\nabla_\nu[v_1]}{16\pi^2}
+ \frac{3\nabla_\nu[v_1]}{16\pi^2} + 0 = 0.
}
\tag{2.3}
\]
DERIVED. The residual and the correction divergence are computed from the
same locked remainder equation (1.8); their cancellation is exact in §0
conventions. No separate conservation of \(C_{\mu\nu}\) is assumed or holds.

---

## 3. BLIND: ambiguities, scale dependence, and trace

### 3.1 Variation implies conservation

Theorem V (DERIVED, Noether): let \(I[g]=\int_M\sqrt{-g}\,L\) with \(L\) a
local scalar of the metric/curvature, and
\(H_{\mu\nu}:=(1/\sqrt{-g})\,\delta I/\delta g^{\mu\nu}\).
Then \(\nabla^\mu H_{\mu\nu}=0\).
Proof: diffeomorphism invariance, \(\delta_\xi I=0\) for compactly supported
\(\xi\), with \(\delta_\xi g^{\mu\nu}=-\nabla^\mu\xi^\nu-\nabla^\nu\xi^\mu\)
(doc 02 §5.1):
\(0=\int(\delta I/\delta g^{\mu\nu})\delta_\xi g^{\mu\nu}
=-2\int\sqrt{-g}\,H_{\mu\nu}\nabla^\mu\xi^\nu
=+2\int\sqrt{-g}\,(\nabla^\mu H_{\mu\nu})\xi^\nu\)
(\(H\) symmetric since \(\delta g^{\mu\nu}\) symmetric; boundary dropped).
\(\xi\) arbitrary \(\Rightarrow\) \(\nabla^\mu H_{\mu\nu}=0\). ∎
Convention: ambiguity tensors use \((1/\sqrt{-g})\delta/\delta g^{\mu\nu}\)
(no \(-2\)); normalizations are absorbed in coefficients.

### 3.2 Curvature variations

Needed identities: Palatini
\(\delta R_{\alpha\beta}=\nabla_\lambda\delta\Gamma^\lambda{}_{\alpha\beta}
-\nabla_\beta\delta\Gamma^\lambda{}_{\alpha\lambda}\)
(DERIVED in 3 lines from the definition of \(R^\rho{}_{\sigma\mu\nu}\):
\(\delta\Gamma\) is tensorial, so \(\delta R\) is its antisymmetrized
covariant derivative); \(\delta g_{\alpha\beta}
=-g_{\alpha\mu}g_{\beta\nu}\delta g^{\mu\nu}\);
\(\delta\sqrt{-g}/\sqrt{-g}=+\tfrac12g^{\mu\nu}\delta g_{\mu\nu}\);
\(g^{\alpha\beta}\delta R_{\alpha\beta}
=\nabla^\mu\nabla^\nu\delta g_{\mu\nu}-\Box(g^{\mu\nu}\delta g_{\mu\nu})\)
(DERIVED by contracting Palatini with the explicit \(\delta\Gamma\));
conversion \(\delta/\delta g^{\mu\nu}=-g_{\mu\alpha}g_{\nu\beta}\,
\delta/\delta g_{\alpha\beta}\).

\(H^{(1)}_{\mu\nu}\) from \(I_1=\int R^2\sqrt{-g}\) (DERIVED):
\(\delta(R^2)=2R(R_{\alpha\beta}\delta g^{\alpha\beta}
+g^{\alpha\beta}\delta R_{\alpha\beta})\); integrating by parts twice:
\[
\boxed{
H^{(1)}_{\mu\nu}
= 2RR_{\mu\nu} - \tfrac12g_{\mu\nu}R^2
+ 2(g_{\mu\nu}\Box - \nabla_\mu\nabla_\nu)R.
}
\tag{3.1}
\]
Self-check (DERIVED, uses L3 with the verified \(+\) sign and L4):
\(\nabla^\mu H^{(1)}_{\mu\nu}
=2R_{\nu\sigma}\nabla^\sigma R+R\nabla_\nu R-R\nabla_\nu R
-2\nabla_\nu\Box R-2R_\nu{}^\sigma\nabla_\sigma R+2\nabla_\nu\Box R=0\). ✓

\(H^{(2)}_{\mu\nu}\) from \(I_2=\int R_{\alpha\beta}R^{\alpha\beta}\sqrt{-g}\)
(DERIVED):
\(\delta(R_{\alpha\beta}R^{\alpha\beta})
=2R^{\alpha\beta}\delta R_{\alpha\beta}
+2R_\alpha{}^\lambda R_{\gamma\lambda}\,\delta g^{\alpha\gamma}\)
(the second term from \(\delta g^{\cdot\cdot}\) inside \(R^{\alpha\beta}\));
Palatini + double integration by parts gives, with
\(M^{\mu\nu}\) the coefficient of \(\delta g_{\mu\nu}\),
\(M^{\mu\nu}=[\nabla_\alpha\nabla_\nu R^{\alpha\mu}
+\nabla_\alpha\nabla_\mu R^{\alpha\nu}-\Box R^{\mu\nu}]
-\tfrac12g^{\mu\nu}\Box R+\tfrac12g^{\mu\nu}R_{\alpha\beta}^2
-2R^{\mu\lambda}R^\nu{}_\lambda\),
where the \(-\tfrac12g^{\mu\nu}\Box R\) comes from the
\(\delta\Gamma^\lambda{}_{\alpha\lambda}\) Palatini piece via contracted
Bianchi (\(\nabla_\beta R^{\alpha\beta}=\tfrac12\nabla^\alpha R\)).
Hence, with \(H^{(2)}_{\mu\nu}=-g_{\mu\alpha}g_{\nu\beta}M^{\alpha\beta}\):
\[
\boxed{
\begin{aligned}
H^{(2)}_{\mu\nu}
={}&-\nabla^\alpha\nabla_\nu R_{\alpha\mu}
-\nabla^\alpha\nabla_\mu R_{\alpha\nu}
+\Box R_{\mu\nu} + \tfrac12g_{\mu\nu}\Box R\\
&-\tfrac12g_{\mu\nu}R_{\alpha\beta}R^{\alpha\beta}
+2R_\mu{}^\lambda R_{\nu\lambda}.
\end{aligned}}
\tag{3.2}
\]
Self-checks: (i) linearized (flat background, keep linear terms):
\(\partial^\mu H^{(2){\rm lin}}_{\mu\nu}
=-\tfrac12\partial_\nu\Box R-\tfrac12\Box\partial_\nu R
+\tfrac12\Box\partial_\nu R+\tfrac12\partial_\nu\Box R=0\)
(DERIVED, uses linearized Bianchi; validates all derivative-term
coefficients); (ii) the \(R^2\)-terms
(\(-\tfrac12gR_{\alpha\beta}^2\) from \(\delta\sqrt{-g}\),
\(+2R_\mu{}^\lambda R_{\nu\lambda}\) from \(\delta g\) inside
\(R^{\alpha\beta}\)) are two-line variations; (iii) full nonlinear
conservation is guaranteed by Theorem V, of which (3.2) is an instance.
The \(\nabla^\alpha\nabla R\)-form is kept (no Bianchi reshuffle needed).

\(m^4g_{\mu\nu}\) and \(m^2G_{\mu\nu}\) (DERIVED, one line each):
\((1/\sqrt{-g})\delta\int m^4\sqrt{-g}/\delta g^{\mu\nu}
=-\tfrac12m^4g_{\mu\nu}\) (absorbing \(-\tfrac12\) in the coefficient,
basis element \(m^4g_{\mu\nu}\), conserved since \(\nabla g=0\));
\((1/\sqrt{-g})\delta\int m^2R\sqrt{-g}/\delta g^{\mu\nu}
=-m^2G_{\mu\nu}\) up to the standard Einstein-tensor variation
(same computation as (3.1) at linear order in \(R\); conserved by L4).

Gauss–Bonnet (Lovelock lemma, SOURCE-REPORTED): in four dimensions the
variation of \(\int(R^2-4R_{\alpha\beta}^2+R_{\alpha\beta\gamma\delta}^2)
\sqrt{-g}\) vanishes identically, so the \(R_{\alpha\beta\gamma\delta}^2\)
variation reduces to a combination of (3.1)–(3.2). Standard (Lovelock's
theorem); pure geometry. \(\Box R\) is a total derivative (no contribution).

### 3.3 Ambiguity basis (polynomial)

Local scalars of mass dimension 4: \(m^4\), \(m^2R\), \(R^2\),
\(R_{\alpha\beta}^2\), \(R_{\alpha\beta\gamma\delta}^2\), \(\Box R\).
Modulo total derivatives and Gauss–Bonnet, every polynomial
dimension-4 ambiguity is, for dimensionless constants \(c_i\):
\[
\boxed{
\Theta_{\mu\nu}
= c_0\,m^4g_{\mu\nu} + c_1\,m^2G_{\mu\nu}
+ c_2\,H^{(1)}_{\mu\nu} + c_3\,H^{(2)}_{\mu\nu},
}
\tag{3.3}
\]
each basis tensor dimension 4, conserved (Theorem V + self-checks).
DERIVED modulo Lovelock. Scope note: (3.3) is the polynomial basis;
non-polynomial extensions (e.g. arbitrary functions of \(R/m^2\)) are a
post-lock literature question, not derived here.

### 3.4 Scale dependence is pure ambiguity

The renormalization length \(\ell\) enters only through
\(\ln|\bar\sigma|=\ln|\sigma|-\ln\ell^2\) in (0.8). Under
\(\ell\to\ell'\): \(S\to S+\Delta S\),
\(\Delta S=(V\ln(\ell'/\ell))/(8\pi^2)\)
(\(\Delta H^{(1)}=-(1/8\pi^2)V\cdot2\ln(\ell'/\ell)\);
\(\Delta S=-\tfrac12\Delta H^{(1)}\)).
\(\Delta S\) is smooth, symmetric, local-geometric. Redoing §1.3–§1.5 with
\(S\to\Delta S\) (the derivation uses only smoothness, symmetry, and the
field equation with symmetric source \(J_{\Delta S}=(\Box-m^2)\Delta S\)):
\(\nabla^\mu\lim T_{\mu\nu}\Delta S=\tfrac12\nabla_\nu[J_{\Delta S}]
=\tfrac12\nabla_\nu[(N_L\cdot{\rm const})]=0\),
since \([N_L]\equiv0\) at every base point (V1a holds everywhere by DP).
Hence \(\ell\)-changes shift \(\langle T_{\mu\nu}\rangle_{\rm raw}\) by a
conserved local geometric tensor — absorbable into \(\Theta_{\mu\nu}\).
DERIVED. The residual (1.3) and correction (2.1) are \(\ell\)-independent
(\([v_1]\) carries no \(\ell\)).

### 3.5 Trace

From (1.1): \(g^{\mu\nu}R_{\mu\nu}=t-2(t+m^2s)\), i.e.
\(\mathrm{tr}_{\rm raw}=-t-2m^2s\).
With \(t=\Box s-m^2s-[J]\) (from \(\Box s=[\Box_xS]+t\) (2.2) and
\([\Box_xS]=m^2s+[J]\)):
\[
\mathrm{tr}_{\rm raw} = -\Box s - m^2s + [J],
\qquad
\mathrm{tr}_{\rm ren} = -\Box s - m^2s + \frac{3[v_1]}{8\pi^2}
+ \Theta^\mu{}_\mu,
\tag{3.4}
\]
using \(g^{\mu\nu}C_{\mu\nu}=4\cdot3[v_1]/(16\pi^2)\) and locked (1.8).
DERIVED. The trace is generically nonzero (state pieces \(-\Box s-m^2s\)
plus geometric \(+3[v_1]/(8\pi^2)\) plus \(\Theta^\mu{}_\mu\)).
A nonzero trace coexists with exact conservation (2.3): trace and
divergence are independent projections (the Exhibit is \(C_{\mu\nu}\)
itself: \(\mathrm{tr}\,C=3[v_1]/(4\pi^2)\ne0\) while \(\nabla C\) cancels
the residual). No anomaly confusion: in this massive minimal model the
trace carries state-dependent pieces; the conformal anomaly
(\(m=0\), \(\xi=1/6\), a different model) is post-lock comparison only.

### 3.6 Lock-2 scope checkpoint

Derived under lock 2 (locked doc-03 items + L3/L4 + this section):
reduction (D1); transport identity TI; residual (1.3); correction (2.1)
with non-conservation (2.2); conservation theorem (2.3); Theorem V;
\(H^{(1)}\) (3.1) with full self-check; \(H^{(2)}\) (3.2) with linearized
self-check; basis (3.3); scale absorption (§3.4); trace (3.4).
Assumed from outside: locked doc-03 items, L3 (derived), L4, Palatini
(derived), Lovelock (SOURCE-REPORTED geometry). Nothing from the
Feynman-route literature, Agent B, or any expected correction shape is
used: \(C_{\mu\nu}\) is forced by (1.3), not imported.

---

## Lock 2

- Timestamp: 2026-11-06T16:03:29Z.
- Scope: file lines 1–397 (header through §3.6), i.e. everything above
  this header.
- SHA-256 of scope bytes:
  `2ee966a80d1e4e3257fcc2887e129acb91e7ec8ed6ba52fc5b7a0f95b4dde5c8`.
- Quote check on scope: no literature equation citations, no Décanini/Wald
  content, no Agent B content. The two mechanical grep hits are the §2.2
  section title (this document's own derived result) and the §3.6
  scope-checkpoint meta-mention. No explicit \(v_1\) curvature polynomial
  appears; \([v_1]\) stays formal throughout.
- Note: during blind writing, self-review caught and fixed two errors
  before locking (all-primed commutator curvature terms in T1/TI proofs,
  now explicit with Ricci cancellation; a factor 2 in the §3.4 scale-shift
  formula). Both fixes use only locked §0–§2 material. The residual (1.3)
  and correction (2.1) are unchanged.
- Quarantine statement: §§1–3 cite only locked doc-03 items, L3 (derived),
  L4/Palatini/Lovelock (pure geometry), and their own equations. The
  correction \(C_{\mu\nu}\) is forced by the derived residual (1.3), not
  imported. Literature/Agent B comparison begins in §4, after this lock.

---

## 4. POST-LOCK: comparison

### 4.1 Feynman-route literature (DF, Wald abstracts)

DF Eqs. (67)–(70), (74) (full text inspected in doc 03 §3; `01`
transcription verified): DF's residual is \(-(D/4)\alpha_D\nabla v_1\)
(\(-\nabla v_1/(4\pi^2)\) at \(D=4\)) with correction
\(+(D/4)\alpha_D\,g_{\mu\nu}v_1\); their \(\Theta_{\mu\nu}\) is conserved.
Blind (1.3)/(2.1): \(-3\nabla[v_1]/(16\pi^2)\),
\(+3g_{\mu\nu}[v_1]/(16\pi^2)\), \(\Theta\) conserved (3.3).
Structural agreement — INDEPENDENTLY REPRODUCED on both sides: raw
divergence is a local-curvature gradient; correction is a local scalar
times the metric; ambiguities are separately conserved; internal
cancellations are exact in each framework. Coefficient cross-match —
CONDITIONAL on the smooth framework-difference term \(A\) (doc 03 §3.2):
the Feynman route (time-ordered, \(i\epsilon\)/averaging) and the symmetric
route (\((1/2)G^{(1)}\), minimal subtraction) split "raw vs correction"
differently by smooth terms; the constraint \([(\Box-m^2)A]=-3[v_1]\) is
DERIVED from the two verified limits. Exact \(A\) remains a bounded open
computation; it does not affect (2.3), which is exact in §0 conventions.
Wald 1978 abstract (verified in `01`): "divergence is the gradient of a
local curvature term; subtract that term times the metric" — matches
(1.3)+(2.1) structurally. INDEPENDENTLY REPRODUCED (structure);
SOURCE-REPORTED (abstract level).

### 4.2 Agent B, pp. 11 and 13 (rendered pages inspected post-lock)

Evidence: `plans/Quantum Stress Energy Conservation Research.pdf` pp. 11,
13 rendered to PNG and read this session (first inspection of these
formulas; text extraction blanks equations).

(B1) \(H^{(1)}_{\mu\nu}\) (p. 11): Agent B gives
\(2\nabla_\mu\nabla_\nu R-2g_{\mu\nu}\Box R+2RR_{\mu\nu}
-\tfrac12g_{\mu\nu}R^2\).
Blind (3.1): \(-2\nabla_\mu\nabla_\nu R+2g_{\mu\nu}\Box R+2RR_{\mu\nu}
-\tfrac12g_{\mu\nu}R^2\).
The derivative terms have opposite signs. FALSIFIED: moreover
\(\nabla^\mu(B^{(1)}_{\mu\nu})-\nabla^\mu(H^{(1)}_{\mu\nu})
=4[\Box,\nabla_\nu]R=+4R_\nu{}^\sigma\nabla_\sigma R\ne0\), so Agent B's
tensor is not conserved, contradicting its "conserved tensors" label.
Minimum correction: adopt (3.1).
(B2) \(H^{(2)}_{\mu\nu}\) (p. 11): Agent B gives
\(\nabla_\mu\nabla_\nu R-g_{\mu\nu}\Box R-\Box R_{\mu\nu}
+2R^{\alpha\beta}R_{\alpha\mu\beta\nu}
-\tfrac12g_{\mu\nu}R_{\alpha\beta}^2\).
FALSIFIED: at linearized order,
\(\partial^\mu(B^{(2){\rm lin}}_{\mu\nu})
=\Box\partial_\nu R-\partial_\nu\Box R-\tfrac12\Box\partial_\nu R
=-\tfrac12\Box\partial_\nu R\ne0\)
(linearized Bianchi), while blind (3.2) passes the same check with 0.
The derivative structure is wrong, not merely unsigned. Minimum
correction: adopt (3.2).
(B3) Ambiguity dimensions (p. 11): "\(c_1,c_2,c_3,c_4\) real
dimensionless" for \(c_1g+c_2G+c_3H^{(1)}+c_4H^{(2)}\).
FALSIFIED: \([g_{\mu\nu}]=0\) so \([c_1]=4\); \([G_{\mu\nu}]=2\) so
\([c_2]=2\); only \(c_3,c_4\) are dimensionless. Blind (3.3) carries the
correct mass factors. (Confirms `01` §6.)
(B4) "Four axioms" (p. 11): FALSIFIED (Wald 1977 abstract lists five;
`01` §4.4). Carried, not re-derived.
(B5) Step-4 homogeneous-\(W\) cancellation (p. 13):
"\(\Box_x\nabla_{\nu'}W=\nabla_{\nu'}\Box_xW=m^2\nabla_{\nu'}W\)" giving
\(\tfrac12\lim[g_\nu{}^{\nu'}\Box_x\nabla_{\nu'}W-m^2\nabla_\nu W]=0\).
FALSIFIED: uses the false homogeneous equation (doc 03 §1.7); the correct
\(\Box_x\nabla_{\nu'}S=\nabla_{\nu'}\Box_xS
=m^2\nabla_{\nu'}S+\nabla_{\nu'}J\) leaves \(\tfrac12[\nabla_{\nu'}J]
=\tfrac14\nabla_\nu[J]\ne0\). Agent B's zero drops the residual source.
(B6) Residual mechanism (p. 13, step 5): residual attributed to
"derivatives of parallel transport bivectors \(\nabla^\mu g_\mu{}^{\mu'}\)".
FALSIFIED: \([\nabla g]=0\) at coincidence (smooth transport); had the
residual come solely from these terms it would vanish. The true residual
comes from the remainder source \(J\) (1.3). No valid derivation exists.
(B7) Correction \(C^{(0)}_{\mu\nu}\) (p. 13): called "a local conserved
curvature tensor" while the proof requires
\(\nabla^\mu C^{(0)}_{\mu\nu}=+\)(residual) \(\ne0\) to cancel
\(-\nabla^\mu C^{(0)}_{\mu\nu}\). FALSIFIED (logical contradiction;
compare blind §2.2, which proves the correction cannot be independently
conserved). Additionally no explicit formula for \(C^{(0)}\) is ever
given ("constructed from derivatives of \([V_0]\) and \([V_1]\)" is not a
derivation): the correction claim is unsubstantiated as stated.
(B8) Scale shift "\(\propto[V_1]g_{\mu\nu}\)" (p. 11): FALSIFIED as a
specific form — the shift \(\lim T_{\mu\nu}\Delta S\) involves \(V\)'s
full coincidence data (including \(v_0\)-Hessians with
\(\nabla_\mu\nabla_\nu R\) pieces), not a pure-trace \([V_1]g\) term, and
no derivation is given. The structural claim (scale shift is a conserved
ambiguity-basis element) is INDEPENDENTLY REPRODUCED by blind §3.4;
Agent B's mapping to the basis is unproven (CONDITIONAL as its claim,
and undermined by (B1)–(B2)).
(B9) Trace anomaly (p. 13):
\(\langle T^\mu{}_\mu\rangle=(1/16\pi^2)a_2
=(1/2880\pi^2)[C^2-E_4+\tfrac13\Box R]\).
FALSIFIED three ways: (i) \(C^2-E_4=2R_{\alpha\beta}^2-\tfrac23R^2\),
so the formula equals \((1/2880\pi^2)[2R_{\alpha\beta}^2-\tfrac23R^2
+\tfrac13\Box R]\), matching neither Wald 1978
(\((2880\pi^2)^{-1}[C^2+R_{\alpha\beta}^2-\tfrac13R^2]+c\Box R\))
nor DF (111); (ii) it uses minimal-coupling \(a_2\) content for the
conformal (\(m=0\), \(\xi=1/6\)) case; (iii) it does not apply to this
programme's frozen model (\(m>0\), \(\xi=0\)), whose trace (3.4) carries
state-dependent pieces. (Confirms `01` §8.)
(B10) "Non-zero trace does not imply non-conservation" (p. 13 prose):
statement INDEPENDENTLY REPRODUCED (blind §3.5 exhibits the coexistence
explicitly); Agent B's argument for it FALSIFIED (invalid proof (B5)–(B7)
plus wrong formula (B9)).

### 4.3 Finite-ambiguity literature scope

`01` §§4.5–4.6 (readback-verified): Tichy–Flanagan derive an
infinite-parameter massive-field family \(m^4F(R/m^2)\) under the Wald
axioms alone; Salcedo argues renormalizability restricts to the
polynomial dimension-four basis. Blind (3.3) is the polynomial basis:
CONSISTENT with Salcedo's criterion, CONDITIONAL on accepting it; the
Tichy–Flanagan non-polynomial extension is beyond blind scope (flagged
for doc 06's assumption register, not derived here). Dispositions:
SOURCE-REPORTED (`01` record); DERIVED ((3.3) as the polynomial basis).

---

## 5. What this document establishes

1. Raw divergence (1.3): DERIVED (blind) — gradient of local geometry,
   \(-3\nabla[v_1]/(16\pi^2)\), via (D1)+TI with Ricci cancellation.
2. Correction (2.1): DERIVED (blind) — \(+3g[v_1]/(16\pi^2)\), forced by
   (1.3); explicitly not independently conserved (§2.2).
3. Conservation (2.3): DERIVED (blind) — exact cancellation in §0
   conventions.
4. Ambiguity basis (3.3): DERIVED (blind) modulo Lovelock — with
   \(H^{(1)}\) (3.1, full self-check) and \(H^{(2)}\) (3.2, linearized
   self-check + Theorem V).
5. Scale absorption (§3.4): DERIVED — \(\ell\)-shifts are conserved
   ambiguities; residual/correction \(\ell\)-independent.
6. Trace (3.4): DERIVED — generically nonzero, compatible with (2.3).
7. DF/Wald comparison: structure INDEPENDENTLY REPRODUCED; coefficient
   cross-match CONDITIONAL on \(A\) (§4.1, doc 03 §3.2).
8. Agent B (B1)–(B10): two structural statements reproduced ((B8)-basis
   idea, (B10)-coexistence); every formula and derivation step FALSIFIED.
9. Carries into doc 05: (B1)–(B10) with rendered-page evidence; into
   doc 06: (2.2)–(2.3), (3.3), (3.4), TF/Salcedo scope flag.

Open (non-blocking): explicit \(A\) (§4.1); full nonlinear \(\nabla H^{(2)}\)
line-by-line check (covered by linearized check + Theorem V + (B2)
external contrast).
---

## Erratum E2 (post-lock correction, 2026-11-06)

§3.5's (3.4) used the erroneous \(\Box s\)-formula (doc-03 Erratum E1).
With \(t=\tfrac12\Box s-m^2s-[J]\), the corrected trace is:
\(\mathrm{tr}_{\rm raw} = -\tfrac12\Box s - m^2s + [J]\);
\(\boxed{\mathrm{tr}_{\rm ren}
= -\tfrac12\Box s - m^2s + 3[v_1]/(8\pi^2) + \Theta^\mu{}_\mu}\).
Only the \(\Box s\) coefficient changes (\(-1\to-\tfrac12\));
qualitative conclusions (generically nonzero; compatible with (2.3))
are unchanged. (1.3), (2.1)–(2.3), (3.1)–(3.3), §3.4 use no \(\Box s\)
and are unaffected. Lock 2 scope (lines 1–397) is unaltered; this
erratum is appended post-lock.

---

## Erratum E4 (normalization repair, 2026-11-06)

E3 (doc 03) voids this document's \(S\)-dependent numbers:
(1.3), (2.1)–(2.3), (3.4)-numbers, §4.1-numbers. Two successive
corrections are recorded in doc 08: N3 re-derives the chain under
the repaired prescription as the symmetric-part computation
(residual \(\tfrac12\nabla[J]\), correction \(+3g[v_1]/(8\pi^2)\));
N4 then corrects TI itself — the locked T1/TI proofs conflate
\(J=(\Box-m^2)S\) with \(J'=(\Box'-m^2)S\) (both \([\nabla J]\)-halves
in TI's proof should distinguish \(J/J'\)); corrected TI gains
\(-2\delta\) and the residual is \([\nabla'J]=\tfrac13\nabla[J]\)
(two routes: revised-T1/TI + checker-#1 Synge-A). Final DERIVED
numbers (N4.5): residual \(-\nabla[v_1]/(4\pi^2)\), correction
\(+g[v_1]/(4\pi^2)\) (matches DF exactly, not imported),
cancellation (N4.7), trace (N4.8). STAND as written: (D1) (exact;
uses \(a+a'=\nabla[J]\)), L3, Theorem V, \(H^{(1)}\)/\(H^{(2)}\)
(3.1)–(3.3), ambiguity/scale/trace STRUCTURES (§3.4 re-proved via
v1-transport in N5.3-D5), §4.2 Agent-B falsifications B1–B10
(superseded only in the cited blind numbers — see E5), §4.3 scope.
Lock 2 scope (lines 1–397) is unaltered; this erratum is appended
post-lock.
