# 02 — Classical independent derivation

Audit: QSG-T001, covariant conservation of the Hadamard-renormalized
stress-energy tensor. Author: Agent C, independent conservation audit.
Date: 2026-11-06.

This document reconstructs the classical foundation independently from the
frozen action. It does not use Agent B's intermediate steps. The final
classical results are then compared with Agent B's pp. 4–6.

Dispositions follow `01-primary-source-verification.md`:
DERIVED, SOURCE-REPORTED, INDEPENDENTLY REPRODUCED, CONDITIONAL, FALSIFIED,
BLOCKED. `PASS` is not used.

---

## 1. Conventions and action

### 1.1 Geometry

- \((M,g)\) is a four-dimensional smooth Lorentzian manifold.
- Signature is \((-+++)\). In particular \(g^{00}<0\) in adapted coordinates,
  and the d'Alembertian is
  \[
  \Box_g:=g^{\mu\nu}\nabla_\mu\nabla_\nu=\nabla^\mu\nabla_\mu .
  \]
- The connection is the Levi-Civita connection; it is torsion-free, so
  \(\nabla_\mu\nabla_\nu f=\nabla_\nu\nabla_\mu f\) for every scalar \(f\).
- Natural units \(c=\hbar=1\). The mass dimension of the scalar is
  \([\phi]=1\), of the metric is \([g_{\mu\nu}]=0\), and of the
  stress-energy tensor is \([T_{\mu\nu}]=4\).
- The background is fixed. No backreaction, no semiclassical Einstein
  equation, and no variation of the gravitational action are used.

### 1.2 Action

The frozen matter action is
\[
S_m[\phi,g]
=
-\frac12\int_M d^4x\,\sqrt{-g}
\left[
g^{\mu\nu}\nabla_\mu\phi\,\nabla_\nu\phi
+
m^2\phi^2
\right],
\qquad m>0,\quad \xi=0 .
\tag{2.1}
\]
The overall minus sign is the one that makes the Hamiltonian positive for
the \((-+++)\) convention. The scalar is real and minimally coupled. There is
no curvature coupling \(\xi R\phi^2\) and no self-interaction.

Because \(\phi\) is a scalar, \(\nabla_\mu\phi=\partial_\mu\phi\). The
covariant form is used so that the metric variation and the Noether argument
are manifestly geometric.

### 1.3 Boundary assumptions

\(M\) is globally hyperbolic and has no boundary (or the variations have
compact support in the interior). All integrations by parts below drop
boundary terms. For the field equation, \(\delta\phi\in C_c^\infty(M)\). For
the diffeomorphism argument, the generating vector field \(\xi^\mu\) is
smooth and compactly supported. These are the standard assumptions that make
the action functional differentiable and the Noether identity local.

---

## 2. Field equation from the action

Vary \(\phi\to\phi+\delta\phi\) with \(g\) fixed:
\[
\delta_\phi S_m
=
-\int_M d^4x\,\sqrt{-g}
\left[
g^{\mu\nu}\nabla_\mu\phi\,\nabla_\nu\delta\phi
+
m^2\phi\,\delta\phi
\right].
\]
Integrating the first term by parts,
\[
\delta_\phi S_m
=
\int_M d^4x\,\sqrt{-g}
\left[
\nabla^\mu\nabla_\mu\phi-m^2\phi
\right]\delta\phi
+
\text{boundary}.
\]
The boundary term vanishes for compactly supported \(\delta\phi\). Therefore
\[
\frac{1}{\sqrt{-g}}\frac{\delta S_m}{\delta\phi}
=
(\Box_g-m^2)\phi .
\tag{2.2}
\]
Setting \(\delta_\phi S_m=0\) for arbitrary \(\delta\phi\) gives the
Klein–Gordon equation
\[
\boxed{(\Box_g-m^2)\phi=0.}
\tag{2.3}
\]
This is the frozen field equation. The sign is checked by the flat-space
limit: \(\Box=-\partial_t^2+\Delta\), so (2.3) is
\((\partial_t^2-\Delta+m^2)\phi=0\), the standard massive Klein–Gordon
equation.

---

## 3. Stress-energy tensor through metric variation

### 3.1 Hilbert definition

The classical stress-energy tensor is defined by the metric variation at
fixed \(\phi\):
\[
T_{\mu\nu}
:=
-\frac{2}{\sqrt{-g}}\frac{\delta S_m}{\delta g^{\mu\nu}}
=
\frac{2}{\sqrt{-g}}\frac{\delta S_m}{\delta g_{\mu\nu}} .
\tag{3.1}
\]
The two expressions are equivalent because
\(\delta g_{\mu\nu}=-g_{\mu\alpha}g_{\nu\beta}\delta g^{\alpha\beta}\).
The sign is the standard Hilbert convention: it gives positive energy for
the action (2.1).

### 3.2 Explicit variation

Use
\[
\delta\sqrt{-g}=-\frac12\sqrt{-g}\,g_{\mu\nu}\,\delta g^{\mu\nu}.
\tag{3.2}
\]
Since \(\phi\) is fixed and \(\nabla_\alpha\phi=\partial_\alpha\phi\) does
not depend on the metric at first order,
\[
\delta\left(g^{\alpha\beta}\nabla_\alpha\phi\nabla_\beta\phi\right)
=
\delta g^{\alpha\beta}\,\nabla_\alpha\phi\nabla_\beta\phi .
\]
Thus
\[
\begin{aligned}
\delta_g S_m
&=
-\frac12\int_M d^4x\,
\Big[
\delta\sqrt{-g}\,
\left(g^{\alpha\beta}\nabla_\alpha\phi\nabla_\beta\phi+m^2\phi^2\right)
+
\sqrt{-g}\,
\delta g^{\alpha\beta}\,
\nabla_\alpha\phi\nabla_\beta\phi
\Big]\\
&=
-\frac12\int_M d^4x\,\sqrt{-g}\,
\delta g^{\mu\nu}
\left[
\nabla_\mu\phi\nabla_\nu\phi
-
\frac12 g_{\mu\nu}
\left(
\nabla_\alpha\phi\nabla^\alpha\phi+m^2\phi^2
\right)
\right].
\end{aligned}
\tag{3.3}
\]
Therefore
\[
\frac{\delta S_m}{\delta g^{\mu\nu}}
=
-\frac12\sqrt{-g}
\left[
\nabla_\mu\phi\nabla_\nu\phi
-
\frac12 g_{\mu\nu}
\left(
\nabla_\alpha\phi\nabla^\alpha\phi+m^2\phi^2
\right)
\right],
\]
and (3.1) gives
\[
\boxed{
T_{\mu\nu}
=
\nabla_\mu\phi\nabla_\nu\phi
-
\frac12 g_{\mu\nu}
\left(
\nabla_\alpha\phi\nabla^\alpha\phi+m^2\phi^2
\right).
}
\tag{3.4}
\]
This is symmetric, \(T_{\mu\nu}=T_{\nu\mu}\), and has mass dimension four.
The trace is
\[
T^\mu{}_\mu
=
-\left(\nabla_\alpha\phi\nabla^\alpha\phi+2m^2\phi^2\right).
\tag{3.5}
\]
For \(m>0\) and \(\xi=0\) this is not zero even on shell. The classical trace
vanishes on shell only in the conformal case \(m=0\), \(\xi=1/6\), which is
not the frozen model. This point matters for the trace-anomaly discussion in
`04`.

### 3.3 Flat-space check

In Minkowski coordinates \(g_{\mu\nu}=\mathrm{diag}(-1,+1,+1,+1)\),
\[
T_{00}
=
\frac12\left[
(\partial_0\phi)^2+(\partial_i\phi)^2+m^2\phi^2
\right]\ge 0,
\]
so the energy density is positive. This fixes the overall sign of (3.4).

---

## 4. Covariant divergence and off-shell identity

Take the divergence of (3.4) directly:
\[
\begin{aligned}
\nabla^\mu T_{\mu\nu}
&=
\nabla^\mu\left(\nabla_\mu\phi\nabla_\nu\phi\right)
-
\frac12\nabla_\nu\left(
\nabla_\alpha\phi\nabla^\alpha\phi+m^2\phi^2
\right)\\
&=
(\Box_g\phi)\nabla_\nu\phi
+
\nabla^\mu\phi\,\nabla_\mu\nabla_\nu\phi
-
\nabla^\alpha\phi\,\nabla_\nu\nabla_\alpha\phi
-
m^2\phi\,\nabla_\nu\phi\\
&=
(\Box_g\phi-m^2\phi)\nabla_\nu\phi
+
\nabla^\mu\phi
\left(
\nabla_\mu\nabla_\nu-\nabla_\nu\nabla_\mu
\right)\phi .
\end{aligned}
\tag{4.1}
\]
The commutator on the scalar \(\phi\) vanishes:
\[
[\nabla_\mu,\nabla_\nu]\phi=0
\tag{4.2}
\]
because the Levi-Civita connection is torsion-free. Hence
\[
\boxed{
\nabla^\mu T_{\mu\nu}
=
\left(\Box_g\phi-m^2\phi\right)\nabla_\nu\phi .
}
\tag{4.3}
\]
This is the off-shell identity. It holds for every smooth \(\phi\), whether
or not \(\phi\) solves (2.3).

### 4.1 What the identity says

- Off shell, the divergence is not zero in general. It is proportional to the
  field equation, with the gradient \(\nabla_\nu\phi\) as the coefficient.
- On shell, \((\Box_g-m^2)\phi=0\), so
  \[
  \boxed{\nabla^\mu T_{\mu\nu}=0.}
  \tag{4.4}
  \]
- The identity is a differential identity of the matter action. It is not an
  independent dynamical equation and it is not a consequence of the
  Einstein equations. It follows from the fact that \(S_m[\phi,g]\) is a
  scalar under simultaneous diffeomorphisms of \(\phi\) and \(g\).

---

## 5. Noether derivation from diffeomorphism invariance

The direct calculation (4.3) is already a complete derivation. The following
Noether argument identifies its geometric origin and is the form that
generalizes to the quantum Ward identity.

### 5.1 Transformation laws

Under an infinitesimal diffeomorphism generated by a compactly supported
vector field \(\xi^\mu\),
\[
\delta_\xi\phi=\xi^\nu\nabla_\nu\phi,
\qquad
\delta_\xi g^{\mu\nu}
=
-\nabla^\mu\xi^\nu-\nabla^\nu\xi^\mu .
\tag{5.1}
\]
The second formula is the Lie derivative of the inverse metric:
\(\mathcal L_\xi g^{\mu\nu}
=\xi^\rho\partial_\rho g^{\mu\nu}
-g^{\rho\nu}\partial_\rho\xi^\mu
-g^{\mu\rho}\partial_\rho\xi^\nu
=-\nabla^\mu\xi^\nu-\nabla^\nu\xi^\mu\).

### 5.2 Invariance and the identity

Because \(S_m[\phi,g]\) is a geometric scalar integral, it is invariant under
active diffeomorphisms:
\[
0=\delta_\xi S_m
=
\int_M d^4x
\left[
\frac{\delta S_m}{\delta\phi}\delta_\xi\phi
+
\frac{\delta S_m}{\delta g^{\mu\nu}}\delta_\xi g^{\mu\nu}
\right].
\tag{5.2}
\]
Substitute (2.2), (3.1) and (5.1):
\[
0=
\int_M d^4x\,\sqrt{-g}\,
\left[
(\Box_g-m^2)\phi\;\xi^\nu\nabla_\nu\phi
-
\frac12 T_{\mu\nu}
\left(
-\nabla^\mu\xi^\nu-\nabla^\nu\xi^\mu
\right)
\right].
\]
The symmetry \(T_{\mu\nu}=T_{\nu\mu}\) gives
\[
-\frac12 T_{\mu\nu}
\left(
-\nabla^\mu\xi^\nu-\nabla^\nu\xi^\mu
\right)
=
T_{\mu\nu}\nabla^\mu\xi^\nu .
\]
Integrating by parts and dropping the boundary term,
\[
0=
\int_M d^4x\,\sqrt{-g}\,
\xi^\nu
\left[
(\Box_g-m^2)\phi\,\nabla_\nu\phi
-
\nabla^\mu T_{\mu\nu}
\right].
\tag{5.3}
\]
Since \(\xi^\nu\) is arbitrary and compactly supported, the bracket vanishes
pointwise:
\[
\boxed{
\nabla^\mu T_{\mu\nu}
=
(\Box_g-m^2)\phi\,\nabla_\nu\phi .
}
\tag{5.4}
\]
This reproduces (4.3). On shell it gives (4.4).

### 5.3 Off-shell versus on-shell

The distinction is essential:

| Statement | Assumption | Content |
|---|---|---|
| \(\nabla^\mu T_{\mu\nu}=(\Box_g-m^2)\phi\,\nabla_\nu\phi\) | none beyond smoothness and compact support | off-shell Noether identity |
| \(\nabla^\mu T_{\mu\nu}=0\) | \((\Box_g-m^2)\phi=0\) | on-shell dynamical conservation |

The off-shell identity is the classical shadow of the quantum Ward identity.
The quantum conservation theorem is not obtained by setting the classical
field equation to zero inside an expectation value; it requires the Hadamard
subtraction and the local correction, as shown in `03` and `04`.

---

## 6. Checks

### 6.1 Signature

With \((-+++)\), \(g^{00}<0\) and the kinetic term in (2.1) is positive for
time-dependent fields. The d'Alembertian is \(\Box=g^{\mu\nu}\nabla_\mu\nabla_\nu\).
The flat-space check in §3.3 confirms the sign of \(T_{00}\).

### 6.2 Signs

- The action has the overall minus sign. Flipping it would flip the sign of
  \(T_{\mu\nu}\) and of the off-shell identity.
- The Hilbert definition (3.1) has the minus sign on
  \(\delta S/\delta g^{\mu\nu}\). With the action (2.1), this gives the
  positive-energy \(T_{\mu\nu}\) of (3.4).
- The Noether identity (5.4) has the same sign as the direct divergence
  (4.3). This is a nontrivial check of the variation conventions.

### 6.3 Index placement

\(T_{\mu\nu}\) is a symmetric \((0,2)\) tensor. \(\nabla^\mu T_{\mu\nu}\) is a
covector. The right-hand side \((\Box_g-m^2)\phi\,\nabla_\nu\phi\) is a
covector. The contracted index in the divergence is the first index of
\(T_{\mu\nu}\), consistent with the conservation law of a symmetric tensor.

### 6.4 Mass dimensions

\[
[\phi]=1,\quad [m]=1,\quad [\nabla]=1,\quad
[T_{\mu\nu}]=4,\quad
[\nabla^\mu T_{\mu\nu}]=5,\quad
[(\Box_g-m^2)\phi]=3,\quad
[\nabla_\nu\phi]=2.
\]
The two sides of (4.3) both have dimension \(5\). The action (2.1) is
dimensionless because \([d^4x]=-4\) and the Lagrangian density has dimension
\(4\).

### 6.5 Variation conventions

The two common conventions are
\[
T_{\mu\nu}=-\frac{2}{\sqrt{-g}}\frac{\delta S_m}{\delta g^{\mu\nu}}
\qquad\text{and}\qquad
T^{\mu\nu}=\frac{2}{\sqrt{-g}}\frac{\delta S_m}{\delta g_{\mu\nu}} .
\]
They are equivalent. The important point is to use the same convention in the
classical derivation and in the point-splitting operator. The operator (7.1)
below is constructed from (3.4), so the conventions are consistent.

### 6.6 Boundary assumptions

Global hyperbolicity gives a well-posed Cauchy problem and excludes timelike
boundaries. The integrations by parts use compact support of \(\delta\phi\)
and \(\xi^\mu\). If the spacetime had a boundary, or if the variations did
not vanish on it, (2.3) and (5.4) would acquire boundary terms. The frozen
model excludes this.

---

## 7. Classical point-splitting operator (needed in `03` and `04`)

The quantum construction uses the bi-differential operator obtained by
promoting (3.4) to split points. At \(\xi=0\), in the symmetrized form used in
`03`,
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
\tag{7.1}
\]
where \(g_\mu{}^{\mu'}(x,x')\) is the bivector of parallel transport from
\(x'\) to \(x\), and the derivatives act on the bivector arguments. The
coincidence limit of (7.1) acting on a symmetric biscalar is the classical
\(T_{\mu\nu}\). This operator is used in `03` to compute the coincidence
limit of \(T_{\mu\nu}W\) and in `04` to compute its divergence. It is not an
independent assumption; it is fixed by (3.4) and the parallel-transport
structure.

---

## 8. Audit of Agent B's classical reconstruction (pp. 4–6)

### 8.1 What is correct

Agent B's final classical results on pp. 4–6 are correct:

- the Euler–Lagrange equation \((\Box_g-m^2)\phi=0\);
- the symmetric stress tensor (3.4);
- the off-shell identity \(\nabla^\mu T_{\mu\nu}=(\Box_g-m^2)\phi\nabla_\nu\phi\);
- the on-shell conservation \(\nabla^\mu T_{\mu\nu}=0\);
- the Noether derivation from diffeomorphism invariance.

These results are INDEPENDENTLY REPRODUCED in §§2–5.

### 8.2 Defects

| Location | Agent B's assertion | Independent result | Disposition |
|---|---|---|---|
| p. 4, canonical momentum | \(\pi=\dfrac{\delta\mathcal L_m}{\delta(\partial_0\phi)}=-\sqrt{-g}\,g^{0\nu}\partial_\nu\phi=+\sqrt{-g}\,\nabla^0\phi\) | The first equality gives \(-\sqrt{-g}\,\nabla^0\phi\), not \(+\sqrt{-g}\,\nabla^0\phi\). With \((-+++)\), \(\nabla^0\phi=g^{0\nu}\partial_\nu\phi\) and the correct momentum is \(\pi=-\sqrt{-g}\,\nabla^0\phi\). | FALSIFIED (sign); non-load-bearing for the conservation theorem, but it is a real error in the canonical setup. |
| p. 4, metric variation intermediate | \(\delta_{g^{\mu\nu}}S_m=-\frac12\int d^4x[\sqrt{-g}\,g^{\mu\nu}\nabla_\mu\phi\nabla_\nu\phi-\frac12\sqrt{-g}\,g_{\mu\nu}\delta g^{\mu\nu}(\cdots)]\) | The first term must contain \(\delta g^{\mu\nu}\), not \(g^{\mu\nu}\); the second term already contains \(\delta g^{\mu\nu}\) and should not be written as if it were part of the same factor. The correct intermediate is (3.3). | FALSIFIED (intermediate); the final \(T_{\mu\nu}\) is correct. |
| p. 5, cancellation of mixed second derivatives | “Because covariant derivatives commute when acting on scalar field gradients (\([\nabla_\mu,\nabla_\nu]\phi=0\)), index symmetry establishes \(\nabla_\mu\phi\nabla^\mu\nabla_\nu\phi=\nabla_\mu\phi\nabla_\nu\nabla^\mu\phi\).” | \([\nabla_\mu,\nabla_\nu]\phi=0\) is true, but the cancellation uses the symmetry of the Hessian of a scalar, \(\nabla_\mu\nabla_\nu\phi=\nabla_\nu\nabla_\mu\phi\). The statement “covariant derivatives commute when acting on scalar field gradients” is false: \(\nabla_\nu\phi\) is a covector, and \([\nabla_\mu,\nabla_\nu]\nabla_\alpha\phi=R^\beta{}_{\alpha\mu\nu}\nabla_\beta\phi\neq0\) in general. The displayed identity itself is correct. | FALSIFIED (justification); the result is INDEPENDENTLY REPRODUCED. |
| p. 6, symmetry statement | “Exploiting the symmetry \(T_{\mu\nu}=\ \) [blank]” | The intended statement is \(T_{\mu\nu}=T_{\nu\mu}\). The blank is a typographical defect, not a mathematical error. | FALSIFIED (typographical); the Noether derivation is correct. |
| pp. 4–6, classification | Classical reconstruction is presented as part of a DERIVED chain | The classical results are correct, but they do not establish the quantum conservation theorem. The quantum step is a separate argument and is invalid as written (`03`–`05`). | CONDITIONAL: classical foundation INDEPENDENTLY REPRODUCED; quantum extension FALSIFIED. |

### 8.3 Assessment

Agent B's classical reconstruction is **not** the source of the failure of
QSG-T001. The final classical equations are correct. The defects are
intermediate sign/justification/typographical errors. The fatal errors are in
the quantum Hadamard step (p. 12) and the local correction (pp. 10, 13), not
in the classical foundation.

**Disposition of Audit A:** the classical action, field equation, stress
tensor, off-shell identity and on-shell conservation are DERIVED and
INDEPENDENTLY REPRODUCED. Agent B's classical final results are correct; his
intermediate defects are FALSIFIED but non-load-bearing.

---

## 9. What this document establishes

1. The frozen action (2.1) gives the Klein–Gordon equation (2.3).
2. The Hilbert stress tensor is (3.4), with positive energy and the correct
   sign.
3. The off-shell identity is (4.3), and on-shell conservation is (4.4).
4. The identity follows from diffeomorphism invariance of the matter action;
   it is not an independent axiom.
5. The classical point-splitting operator (7.1) is fixed by (3.4).
6. Agent B's final classical results are correct; his intermediate defects do
   not affect the classical conclusion.

The next document reconstructs the quantum Hadamard foundations and shows
precisely why the smooth remainder \(W\) is not a homogeneous solution.
