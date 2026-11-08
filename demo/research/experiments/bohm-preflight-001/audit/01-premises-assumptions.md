# 01 — Premises, system freeze, assumptions register

Programme: Bohmian equivariance, Benchmark 2 (plan7 + plan7.1, approval
safeguards S1/S2). Author: executor. Date: 2026-11-06.

Dispositions follow the QSG convention: DERIVED (shown here from stated
premises), SOURCE-REPORTED (standard result cited with inspected source),
CONDITIONAL (holds under an explicit undischarged premise), UNSUPPORTED
(not established here), INDEPENDENTLY-REPRODUCED (blind re-derivation
agrees — not claimed unless performed).

---

## 1. System freeze (S1 — binding on docs 02–03 and the fixture)

- Finite collection of \(N \ge 1\) spinless, distinguishable particles in
  three-dimensional Euclidean space. Configuration space
  \(\mathcal Q = \mathbb R^{3N}\), points \(q=(q_1,\dots,q_N)\).
- Masses \(m_1,\dots,m_N > 0\) (stated, distinct symbols; no
  equal-mass assumption).
- Real-valued potential \(V:\mathcal Q\to\mathbb R\) (regularity per B3).
- Wavefunction \(\psi:\mathbb R\times\mathcal Q\to\mathbb C\) governed by
  the standard nonrelativistic Schrödinger equation
  \[ i\hbar\,\partial_t\psi = H\psi,\qquad
     H = \sum_{k=1}^N T_k + V,\qquad
     T_k = -\frac{\hbar^2}{2m_k}\Delta_k, \]
  with \(\Delta_k\) the Laplacian in \(q_k\). \(\hbar\) and \(m_k\) stay
  explicit (no \(\hbar=1\)); the Bohmian current carries \(\hbar/m_k\).
- Density \(\rho = |\psi|^2 = \psi^*\psi\).
- Current (per particle) \(j_k = \frac{\hbar}{m_k}\,
  \mathrm{Im}(\psi^*\nabla_k\psi)\), \(\nabla_k\) the gradient in \(q_k\).
- Guidance velocity \(v_k = j_k/\rho\) on the open set \(\{\rho>0\}\);
  undefined at nodes \(\{\rho=0\}\) (B5).
- Flow \(\Phi_{t,t_0}\) of \(\dot q_k = v_k(t,q)\) where it exists (B6/B7).

## 2. Central result (frozen statement; proof in doc 03)

Let \(\rho_t(q) = |\psi_t(q)|^2\). Local form: on any time interval
where the Bohmian flow exists and \(\rho_t\) is sufficiently regular,
\[ (\Phi_{t,t_0})_\#\,\rho_{t_0} = \rho_t, \]
i.e. the pushforward of the initial distribution under the
configuration-space flow equals the Schrödinger-evolved distribution.
Global form: the same identity for all \(t\in\mathbb R\) under global
flow-existence premises. The continuity equation
\(\partial_t\rho + \sum_k\nabla_k\!\cdot\! j_k = 0\) is a LEMMA, not the
result: equivariance additionally requires the flow, the transport
identity, and control of probability flux (uniqueness for the
continuity equation, or the trajectory route of doc 03 §4). Local
and global conclusions are established (or not) separately. NOTE
(correction of 2026-11-06): the local form is CONDITIONAL on B8+B9
per corrected doc 03 §4 — the freeze states the claim, §4 states
its disposition.

## 3. Assumptions register

- B1 (dynamics): \(\psi\) satisfies the frozen Schrödinger equation
  strongly enough that \(\partial_t\rho\) exists pointwise
  (\(\psi,\partial_t\psi\) continuous). Premise.
- B2 (regularity): \(\nabla_k\psi\), \(\Delta_k\psi\) continuous on
  \(\mathbb R\times\mathcal Q\) (so \(\rho, j_k\) are \(C^1\) and the
  continuity computation is classical), AND the full spatial Hessian
  of \(\psi\) (all second partials
  \(\partial_{q_k^a}\partial_{q_l^b}\psi\)) continuous. The Hessian
  clause is what makes \(j_k\) itself \(C^1\) (\(j\) is built from
  \(\nabla\psi\), so \(\nabla j\) needs second derivatives of
  \(\psi\)), hence \(v=j/\rho\) locally Lipschitz on
  \(\{\rho>0\}\) for the doc-03 flow argument. Continuity of
  \(\nabla\psi\) and \(\Delta\psi\) alone suffices for the doc-02
  continuity computation but not for local-flow existence. Premise;
  minimal sufficient regularity recorded in doc 02 if weaker
  conditions suffice.
- B3 (potential): \(V\) real-valued; sufficient regularity for B1–B2 to
  hold (smooth bounded \(V\) certainly suffices; singular/Coulomb-type
  \(V\) NOT included unless doc 02 derives the needed domain facts —
  default excluded).
- B4 (decay/boundary): \(\psi,\nabla_k\psi\to 0\) at infinity fast
  enough that boundary flux vanishes when integrated conservation is
  asserted. Needed ONLY for global/integrated statements, never for
  the local continuity identity. Premise where used, flagged at use.
- B5 (nodes): \(v_k\) is defined only on \(\{\rho>0\}\); no assertion
  is made about trajectories hitting, leaving, or crossing the nodal
  set \(\{\rho=0\}\) unless an explicit node premise is stated and
  sourced. Standing restriction.
- B6 (local flow): on \(\{\rho>0\}\) with B2, \(v\) is locally
  Lipschitz (quotient of \(C^1\) functions over nonvanishing
  denominator), hence the ODE \(\dot q=v(t,q)\) has unique local
  solutions (standard ODE theory, SOURCE-REPORTED at use).
- B7 (global flow + transport): existence of \(\Phi_{t,t_0}\) for all
  time, non-arrival at nodes in finite time, and uniqueness for the
  continuity equation. NOT derived from B1–B6 here; either imported
  with an inspected primary source (Dürr–Goldstein–Zanghì-type
  results) or the global conclusion stays CONDITIONAL/unsupported.
  Decision recorded in doc 03 with evidence.
- B8 (a.e. local flow on a common interval): on some nontrivial
  interval \([t_0,t_1]\), the Bohmian flow \(\Phi_{t,t_0}(q_0)\)
  exists for all \(t\in[t_0,t_1]\) for \(\mu_{t_0}\)-almost every
  initial \(q_0\), where \(\mu_{t_0}\) is the measure with density
  \(\rho_{t_0}\). Strictly stronger than B6 (single-trajectory
  local existence): it supplies the common interval and the
  full-measure set over which trajectory identities assemble into
  a pushforward statement. Premise where used (doc 03 §4).
- B9 (no probability loss through nodes/escape): over the B8
  interval, no \(\mu\)-mass reaches the nodal set \(\{\rho=0\}\)
  or escapes to infinity/any finite-time blowup boundary, so the
  trajectory-level density identity integrates to the full
  pushforward without boundary-loss terms. B4 covers only decay
  at spatial infinity for classical densities and does NOT
  supply nodal-flux control. Premise where used (doc 03 §4).

## 4. Exclusions (stop conditions, binding)

Spin; relativistic Bohmian mechanics; quantum field theory;
measurement theory, collapse, or Born-rule justification
(typicality/BHC-R3 routes out); interpretation comparison;
backreaction/self-gravity; singular potentials (default, per B3).

## 5. Recon decisions carried in (Step 1 record)

- Placement: Go artifact in `phys-lib/bohm/equivariance/` (mission
  constraint + minisuperspace precedent); math record here
  (`bohm-preflight-001/audit/`, single-executor layout rather than
  the QSG multi-agent tree).
- Placement ADDENDUM (same session, 2026-11-06): a direct user order
  ("do not modify any phys codebase") superseded the mission
  placement mid-Step-4. The artifact was removed from phys-lib
  (worktree restored clean; recovery ref
  `refs/tbh/recovery/before-discard/20261106T182751Z-39065` retained
  in phys-lib git storage) and rebuilt as the standalone TOE-v6
  module `theorems/bohm-equiv/` (`github.com/PithomLabs/bohm-equiv`,
  colocated go.work, phys-artifact + phys-math read-only). All
  Step-4/5 evidence below refers to the standalone module; the
  shared repos carry zero diff from this programme.
- Reuse map: physmath expression/function/relation/condition/
  derivation/equation/definition families; rules CANON, SIMPLIFY,
  DIFFERENTIATE-001 (+ TOTAL-DERIVATIVE/SUBSTITUTE/RAT-NORMALIZE
  only if the derivation consumes them); Encode/Validate/Replay;
  NEQ conditions; minisuperspace household (Build/Verify, Derivs
  replay, anti-goldens incl. historical-wrong-path pattern,
  limitation-record + coverage-audit tests, source_mapping sidecar
  + structural test). Expected gaps (verify at implementation): no
  integration rule, no ODE/flow execution, no measure/pushforward
  execution, PDE existence unrepresentable → sketch/limitations.
- Manifest: GR 16-key shape mirrored from minisuperspace; gaps
  registered (expected_result + falsification absent; source
  relocated to sidecar per maxwellbox precedent — frozen
  phys-trace reads ONLY the minisuperspace manifest path, so no
  in-place consumer constrains us); NO research_metadata.json in
  repo (minisuperspace precedent; epistemic tracking lives here).
- gr/equivariance (algebraic \(f(\rho_X)=\rho_Y(f)\)) is the §3.3
  comparison foil for doc 04, not a premise.
- No subagent cross-check: mission neither requires nor forbids;
  the chain is standard textbook material whose risks are
  assumption-explicitness (handled by B-register + dispositions),
  not coincidence calculus; recorded limitation.
- phys-math untouched expected (deficiency-driven only); shared
  go commands need writable GOCACHE (sandbox default read-only);
  phys-lib git baseline clean except pre-existing untracked
  `.kilo/`; frozen `phys`/`phys-gr` + root go.work never touched;
  nothing committed.
