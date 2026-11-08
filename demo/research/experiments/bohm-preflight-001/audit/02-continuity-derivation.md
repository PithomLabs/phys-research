# 02 — Continuity equation (DERIVED lemma)

Parent: doc 01 (S1 freeze, B-register). This document derives the
continuity equation from the frozen Schrödinger dynamics. It is a
LEMMA for equivariance, not equivariance itself: no flow, no
transport, no pushforward appears here.

---

## 1. Current-definition identities (DERIVED)

From S1: \(j_k = (\hbar/m_k)\,\mathrm{Im}(\psi^*\nabla_k\psi)\).
(i) \(j_k\) is real-valued (imaginary part of a complex field).
(ii) Equivalent form:
\[ j_k = \frac{\hbar}{2im_k}
       (\psi^*\nabla_k\psi - \psi\nabla_k\psi^*), \]
since \(z - z^* = 2i\,\mathrm{Im}\,z\) with \(z=\psi^*\nabla_k\psi\).
Both forms are used below; they are algebraically identical, no
approximation. Consumes: definitions only.

## 2. Continuity derivation (DERIVED)

\(\rho = \psi^*\psi\). By B1–B2 all derivatives below exist
classically and the product/chain rules apply pointwise.

Step 1 — time derivative:
\[ \partial_t\rho
   = (\partial_t\psi^*)\psi + \psi^*(\partial_t\psi). \]

Step 2 — Schrödinger equation and its conjugate. From S1,
\(i\hbar\,\partial_t\psi = H\psi\), i.e.
\(\partial_t\psi = -(i/\hbar)H\psi\).
\(H\) has real coefficients (\(T_k\) real differential operator,
\(V\) real by B3), so conjugating gives
\(-i\hbar\,\partial_t\psi^* = (H\psi)^*\), i.e.
\(\partial_t\psi^* = +(i/\hbar)(H\psi)^*\).
The relative sign is the load-bearing step; it is checked twice
in §4.

Step 3 — potential cancellation. Substituting:
\[ \partial_t\rho
   = \frac{i}{\hbar}\big[(H\psi)^*\psi - \psi^*H\psi\big]. \]
Write \(H = T + V\) (\(T=\sum_k T_k\)). The potential pieces:
\((V\psi)^*\psi - \psi^*V\psi = V|\psi|^2 - V|\psi|^2 = 0\)
using \(V\) real (B3). Only kinetic terms survive:
\[ \partial_t\rho
   = \frac{i}{\hbar}\sum_k\big[(T_k\psi)^*\psi - \psi^*T_k\psi\big].
\]

Step 4 — kinetic term as divergence. Per particle,
\((T_k\psi)^*\psi - \psi^*T_k\psi
= -(\hbar^2/2m_k)[(\Delta_k\psi)^*\psi - \psi^*\Delta_k\psi]\).
Divergence identity (product rule, B2):
\[ \nabla_k\!\cdot\!(\psi^*\nabla_k\psi - \psi\nabla_k\psi^*)
   = \psi^*\Delta_k\psi - \psi\Delta_k\psi^*, \]
because the cross terms \((\nabla_k\psi^*)\!\cdot\!(\nabla_k\psi)\)
cancel. Hence
\((\Delta_k\psi)^*\psi - \psi^*\Delta_k\psi
= -\nabla_k\!\cdot\!(\psi^*\nabla_k\psi - \psi\nabla_k\psi^*)\).
With \(\psi^*\nabla_k\psi - \psi\nabla_k\psi^*
= 2i\,\mathrm{Im}(\psi^*\nabla_k\psi)\):
\[ (T_k\psi)^*\psi - \psi^*T_k\psi
   = -\frac{\hbar^2}{2m_k}\cdot\big[-2i\,
      \nabla_k\!\cdot\!\mathrm{Im}(\psi^*\nabla_k\psi)\big]
   = \frac{i\hbar^2}{m_k}\,
      \nabla_k\!\cdot\!\mathrm{Im}(\psi^*\nabla_k\psi). \]

Step 5 — assembly:
\[ \partial_t\rho
   = \frac{i}{\hbar}\sum_k \frac{i\hbar^2}{m_k}\,
      \nabla_k\!\cdot\!\mathrm{Im}(\psi^*\nabla_k\psi)
   = -\sum_k \nabla_k\!\cdot\!
      \left[\frac{\hbar}{m_k}\mathrm{Im}(\psi^*\nabla_k\psi)\right], \]
using \(i\cdot i = -1\) and \(\hbar^2/\hbar = \hbar\). That is
\[ \boxed{
   \partial_t\rho + \sum_{k=1}^N \nabla_k\!\cdot\! j_k = 0.
   }
\tag{C}
\]
Disposition: DERIVED from S1+B1+B2+B3. Consumed exactly: B1
(\(\partial_t\) exists), B2 (classical product rule + \(\Delta_k\)),
B3 (\(V\) real, Step 3). NOT consumed: B4 (no integration), B5–B7
(no flow, no nodes, no transport).

## 3. What (C) does and does not establish

ESTABLISHED: pointwise differential identity for every classical
solution; \(j_k\) real; \(V\)-independence of the current (any real
\(V\) gives the same \((C)\) — the potential affects \(\psi\), not
the identity's form). NOT established: conservation of total
probability (needs B4 integration — stated, not derived, until
doc 03 needs it); anything about trajectories, flow, nodes, or
equivariance; anything about non-solutions (for arbitrary test
\(\psi\), (C) need not hold — this becomes a fixture NEGATIVE
check, doc 03 §5).

## 4. Self-checks (both performed)

(i) Sign audit: \(\partial_t\psi\) carries \(-i/\hbar\);
\(\partial_t\psi^*\) carries \(+i/\hbar\); Step 5 gives
\(i\cdot i = -1\), matching the \(+\sum\nabla\!\cdot\! j\) on the
left of (C). Recomputed independently from (ii)-form of \(j_k\):
identical result.
(ii) Free-particle plane wave \(\psi = e^{i(p\cdot q - Et)/\hbar}\):
\(\rho = 1\), \(j_k = p_k/m_k\) constant, \(\partial_t\rho = 0\),
\(\nabla_k\!\cdot\! j_k = 0\): (C) holds \(0=0\). (Plane wave is
not \(L^2\); used only as an algebraic check of coefficients,
never as a physical state.)
(iii) Dimensional audit: \([\rho]=L^{-3N}\),
\([j_k]=[\hbar/m]\cdot L^{-3N-1} = (L^2T^{-1})(L^{-3N-1})\),
\([\nabla_k\!\cdot\! j_k]=L^{-3N}T^{-1}=[\partial_t\rho]\). ✓
