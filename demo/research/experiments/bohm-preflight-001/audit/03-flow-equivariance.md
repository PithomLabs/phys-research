# 03 — Flow, transport, equivariance (trajectory identity DERIVED; local pushforward CONDITIONAL; global SOURCED)

Parents: doc 01 (freeze, B-register), doc 02 (continuity lemma (C)).
Proves the frozen pushforward statement locally from stated premises;
global conclusion via one inspected primary source (B7). Local and
global are established separately, as frozen.

---

## 1. Guidance field and local flow (DERIVED modulo standard ODE)

On the open set \(\{\rho>0\}\), \(v_k = j_k/\rho\) with \(j_k\) from
doc 02. By B2, \(j_k\) is \(C^1\) and \(\rho\) is \(C^1\) and
nonvanishing there, so each \(v_k\) is \(C^1\), hence locally
Lipschitz (quotient rule; DERIVED from calculus). By B6 (standard
ODE theory, Picard–Lindelöf, SOURCE-REPORTED): for every
\((t_0,q_0)\) with \(\rho(t_0,q_0)>0\) there exist a maximal open
interval \(I\ni t_0\) and a unique \(C^1\) solution
\(t\mapsto Q(t;t_0,q_0)\) of \(\dot q_k = v_k(t,q)\),
\(Q(t_0)=q_0\), defined as long as the trajectory stays in
\(\{\rho>0\}\) and bounded. The local flow map
\(\Phi_{t,t_0}(q_0) = Q(t;t_0,q_0)\) is a \(C^1\) diffeomorphism
onto its image (standard ODE flow theorem, SOURCE-REPORTED).
Nothing is asserted about arrival at nodes, blowup, or escape:
those are the obstructions B7 addresses globally.

## 2. Transport identity (DERIVED)

Let \(\sigma(t,q)\) be \(C^1\) and consider its values along flow
trajectories: \(\frac{d}{dt}\sigma(t,\Phi_{t,t_0}(q_0))
= \partial_t\sigma + v\!\cdot\!\nabla\sigma\) (chain rule).
Conservation of \(\sigma\)-mass in comoving volumes is equivalent
to the Eulerian transport equation
\[ \partial_t\sigma + \nabla\!\cdot\!(\sigma v) = 0,
\tag{T} \]
by the Reynolds transport theorem for \(C^1\) fields (standard
calculus, DERIVED from the divergence theorem under B4-type decay
for the integrated form; pointwise form needs no decay).
KEY STEP: the Schrödinger density \(\rho_t = |\psi_t|^2\) satisfies
(T) with the guidance \(v\), because \(\rho_t v = j_t\) HOLDS BY
CONSTRUCTION of \(v=j/\rho\) on \(\{\rho>0\}\), and doc-02 (C)
reads \(\partial_t\rho + \nabla\!\cdot\! j = 0\). So (C) IS the
transport equation for \(\rho\) along the Bohmian flow wherever
the flow is defined. Consumes: doc 02 + S1 definitions. No new
premise.

## 3. Uniqueness lemma (DERIVED)

Premise U1 (stated): on the time interval of interest, \(v\) is
\(C^1\) with bounded \(\nabla\!\cdot\! v\), and solutions decay
enough that boundary flux vanishes (follows from B4 for
Schrödinger densities).
Lemma: under U1, (T) with given \(C^1\) initial data \(\sigma_0\)
has at most one \(C^1\) solution. Proof (energy estimate):
for \(w=\sigma^1-\sigma^2\) solving \(\partial_t w+\nabla\!\cdot\!(wv)
=0\) with \(w(0)=0\),
\(\frac{d}{dt}\int w^2 = 2\int w\,\partial_t w
= -2\int w\,\nabla\!\cdot\!(wv)
= -2\int w(v\!\cdot\!\nabla w + w\,\nabla\!\cdot\! v)
= -\int v\!\cdot\!\nabla(w^2) - 2\int w^2\nabla\!\cdot\! v
= -\int w^2\nabla\!\cdot\! v\)
(integration by parts; boundary term vanishes by U1-decay).
Hence \(|\frac{d}{dt}\|w\|_2^2| \le \sup|\nabla\!\cdot\! v|\,
\|w\|_2^2\) and Grönwall gives \(\|w\|_2\equiv0\).
Disposition: DERIVED from U1 + calculus. (U1 is the precise
analytic price of uniqueness; it is NOT theorem-free.)
FLAG (correction of 2026-11-06): U1's decay premise and B4
cover boundary flux at spatial infinity ONLY. On the
node-punctured domain \(\{\rho>0\}\) the energy estimate's
integration by parts additionally needs flux control on the
internal nodal boundaries, which neither U1 nor B4 supplies.
The lemma stands as proved; its APPLICATION to the Bohmian
punctured domain is no longer claimed — §4 routes the local
pushforward through the trajectory identity + B8/B9 instead.

## 4. Local equivariance (CONDITIONAL on B8+B9; trajectory identity DERIVED)

Trajectory identity (DERIVED modulo standard ODE). Let
\(Q(t) = \Phi_{t,t_0}(q_0)\) be a trajectory that exists on an
interval and stays in \(\{\rho>0\}\) (B6-local, single
trajectory — no a.e. or common-interval premise needed here).
Let \(J(t) = \det D\Phi_{t,t_0}(q_0)\) be the Jacobian
determinant of the flow map. By Liouville's formula
(\(\dot J = (\nabla\!\cdot\! v)\,J\); standard ODE theory,
SOURCE-REPORTED, same grade as B6) and the chain rule:
\[ \frac{d}{dt}\big[\rho(t,Q(t))\,J(t)\big]
   = J\,\big[\partial_t\rho + v\!\cdot\!\nabla\rho
     + \rho\,\nabla\!\cdot\! v\big]
   = J\,\big[\partial_t\rho + \nabla\!\cdot\!(\rho v)\big]
   = 0, \]
where the last equality is the transport equation (T) of §2,
which holds on \(\{\rho>0\}\) because \(\rho v = j\) there by
CONSTRUCTION of \(v=j/\rho\). Since \(J(t_0) = 1\):
\[ \boxed{
   \rho_t(\Phi_{t,t_0}(q_0))\,
   \big|\det D\Phi_{t,t_0}(q_0)\big| = \rho_{t_0}(q_0).
   } \tag{D-traj} \]
Consumes: §2 + chain rule + Liouville import. Nothing about
nodes, escape, or other trajectories enters: (D-traj) is a
per-trajectory identity.

Theorem (local, CONDITIONAL). Assume B1–B6 + B8 + B9 on
\([t_0,t_1]\). Then
\[ \boxed{ (\Phi_{t,t_0})_\#\,\rho_{t_0} = \rho_t,
   \qquad t\in[t_0,t_1]. } \tag{E-loc} \]
Assembly: B8 supplies the common interval and the
\(\mu_{t_0}\)-full-measure set of initial configurations whose
trajectories exist throughout \([t_0,t_1]\); (D-traj) holds on
each of them; B9 guarantees no \(\mu\)-mass is lost through
the nodal set or escape/blowup boundaries, so the
trajectory identities integrate to the full pushforward
identity with no boundary-loss terms. ∎
CORRECTION NOTE (2026-11-06, review-driven): the previous
version of this section proved (E-loc) via the §3 uniqueness
lemma applied on the node-punctured domain. That route is
WITHDRAWN: the energy estimate's integration by parts needs
flux control on the internal nodal boundaries, which U1/B4
do not supply (see §3 flag). The trajectory route above is
the corrected proof; its undischarged premises are exactly
B8+B9. What makes it work: (C) + the DEFINITION
\(v=j/\rho\) + per-trajectory flow + a.e. assembly without
loss. The continuity equation alone proves nothing about
distributions; the flow + B8/B9 do the rest. This is the
deliverable floor: trajectory identity derived, pushforward
conditional on stated premises, no other imports beyond
standard ODE/calculus.

## 5. Global equivariance (SOURCE-REPORTED import)

Inspected primary source: Teufel–Tumulka, "Simple Proof for
Global Existence of Bohmian Trajectories", arXiv:math-ph/0406030
(CMP 258 (2005) 349–365), full text inspected 2026-11-06:
Theorem 1 (current-field existence + equivariance on \([0,T]\)
under integrability conditions (15)–(17): a.s.-no node approach,
no escape to infinity, no singular-set approach); Theorem 3 +
Corollary 4 (spinless Schrödinger case: \(V=V_1+V_2\) smooth,
\(V_1\) bounded below, \(V_2\) \(-\tfrac12\Delta\)-form-bounded
relative bound <1; \(\psi(0)\in C^\infty(H)\) normalized
\(\Rightarrow\) trajectories exist globally for
\(\mu_0\)-almost every \(q\), and \(|\psi_t|^2\) equivariant).
Imported statement (B7, closed as SOURCE-REPORTED): under the
TT potential/wavefunction class (which CONTAINS our S1 freeze
with B3 strengthened to smooth + bounded-below \(V\) — recorded
B3-final), the Bohmian flow exists for all \(t\) for
\(\mu_0\)-almost every initial \(q\), and
\[ \boxed{ (\Phi_{t,t_0})_\#\,\rho_{t_0} = \rho_t
   \quad\text{for all }t,\ \mu_{t_0}\text{-a.e.} }
\tag{E-glob} \]
Qualifications carried verbatim from the source: ALMOST ALL
(nodes genuinely kill some trajectories in finite time —
Holland p.85's proof to the contrary is incorrect per TT §1);
self-adjoint \(H\) with \(\psi(0)\in C^\infty(H)\); singular
(Coulomb-type) potentials need the excised-configuration
formulation (outside our freeze; NOT claimed). Our B5 (no node
assertions) is consistent with the a.e. qualification. The TT
proof itself is not reproduced; its statement is imported with
the inspected citation above. B3-final: smooth + bounded-below
\(V\) (import-grade); singular \(V\) excluded. SCOPE BAR (kept
visible per review): the TT import covers ONLY the TT
potential/wavefunction class — smooth \(V=V_1+V_2\) as above,
self-adjoint \(H\), \(\psi(0)\in C^\infty(H)\) normalized. It
does NOT automatically generalize to every classical solution
of the original S1 assumptions; any wider claim would need its
own source or derivation, and none is made here.

## 6. MRC regime/scope + dispositions

ESTABLISHED (this record): system freeze S1; continuity (C)
DERIVED (doc 02); local flow + transport identity + uniqueness
lemma + trajectory identity (D-traj) DERIVED modulo standard
ODE/calculus (B6, U1-proof, Liouville ingredients); (E-loc)
CONDITIONAL on B8+B9 (§4, corrected 2026-11-06); global
(E-glob) SOURCE-REPORTED via inspected TT with a.e.
qualification. REGIME: finite spinless nonrelativistic
particles, smooth bounded-below potential, classical (\(C^2\))
solutions, local-in-time trajectory identity derived and
pushforward conditional, global a.s. imported. NOT
established: node-hit behavior,
singular potentials, spin/magnetic/relativistic cases (TT cover
them; we do not import what we do not need), typicality,
Born-rule justification, measurement, interpretation
comparison — all excluded per doc 01 §4. MRC debt: B8+B9
undischarged (the precise price of the local pushforward, per
the 2026-11-06 correction); every other B-premise is either
consumed by a derivation step above or explicitly imported
(B6/U1-calculus/Liouville standard, B7-TT).

## 7. Fixture handoff (binding on Step 4)

MAY PIN (executable on explicit test inputs via DIFFERENTIATE-001
+ CANON/SIMPLIFY, minisuperspace precedent): current-assembly
identities (\(j\) from \(\psi^*\nabla\psi\) parts, both §1 forms);
continuity-residual computation on test wavefunctions (solution
test input → residual zero; NON-solution test input → residual
nonzero as an honest NEGATIVE check); guidance quotient
construction with NEQ(\(\rho\),0) conditions (node-blind assembly
without the condition is the ANTI-GOLDEN); coefficient identities
(\(i\cdot i=-1\) chain, \(\hbar^2/\hbar\) chain as rational
skeletons). MUST SKETCH/limit (unexecutable: no integration rule,
no ODE/flow execution, no measure/pushforward, PDE existence
unrepresentable): flow ODE existence, pushforward definition,
uniqueness lemma, global existence, node analysis, (E-loc)/(E-glob)
as theorems. Anti-golden ideas: wrong-sign current
(\(j\to -j\) breaks (C)); dropped-conjugate residual; guidance
without NEQ condition. No golden may assume the conclusion: every
executed identity must hold-or-fail on its test inputs
independently of (E-loc)/(E-glob).
