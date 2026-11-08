# 08 — Normalization repair record

Audit: QSG-T001. Author: Agent A (normalization repair per `plans/prompt5.md`,
amended by `plans/plan6_review.md`). Date: 2026-11-06.

This document is the correction record for the normalization defect found in
adversarial review (`plans/plan5_imp_review.md` §1): the subtracted remainder
\(S=\tfrac12(G^{(1)}-H^{(1)})\) of doc 03 (0.5) is singular, so premise P2
(smoothness) was never established and every number downstream of it —
(1.8), (1.9), (1.3), (2.1), (2.3), (3.4), the \(A\)-constraint, doc 04 §4.1 —
is void. The locked scopes of docs 03–04 are untouched; short errata E3/E4/E5
appended to docs 03/04/05 point here for the full record.

Dispositions follow `01`. `PASS` is not used. "REPAIRED" below means
re-derived under the frozen corrected prescription and recorded here; it
does not mean independently replicated (see N5/O1 for the replication
record).

---

## N0. Repair protocol (execution decisions)

D1. One correction record (this file) plus short errata appends (E3: doc 03,
E4: doc 04, E5: doc 05). No new math or Go infrastructure (prompt5
mission). The seven-deliverable structure is preserved; this file is a
correction record, not an eighth deliverable.
D2. Alternative A adopted (N1.6): \(S:=\tfrac12G^{(1)}-H^{(1)}\) with the
standard \(H^{(1)}\) unchanged. Every \(H^{(1)}\)-only derivation stands
as written; every \(S\)-dependent number is re-derived, never imported.
D3. Blind sequencing. N1 uses only frozen §0 conventions, the verified
definitional record of `01` (§3.2: \(\alpha_D\), prefactor, \(\sigma\);
§4.1 navigation), and a fresh Minkowski computation. N2–N3 derive the
remainder/divergence/correction values from the frozen prescription
before any comparison. DF's equation chain (48)–(71) is opened only in
N4 (reconciliation phase), where comparison is the task. The DF-derived
numbers of `01` are never used as derivation targets.
D4. Outcome-neutral re-audit (N5): no prior finding is presumed to reverse
or survive; each is re-derived or re-checked and documented with its
dependency chain.
D5. This record establishes the corrected prescription and the derivations
in it. It does not by itself establish DF-equivalence (N4) or independent
replication (N5/O1); those carry their own verdicts.

---

## N1. Distribution-level normalization audit (REPAIRED)

### N1.1 Minkowski verification of the Wightman singularity (DERIVED)

Free massless field in Minkowski space, signature \((-+++)\).
\(G^+(x,x')=\int\!\frac{d^3k}{(2\pi)^3 2\omega}\,
e^{-i\omega\Delta t+i\mathbf k\cdot\Delta\mathbf x}\) with
\(\omega=|\mathbf k|\), \(\Delta t=t-t'\), \(r=|\Delta\mathbf x|\).
Spherical integration:
\[
G^+=\frac{1}{8\pi^2}\int_0^\infty\!dk\,k\,e^{-ik\Delta t}
\frac{e^{ikr}-e^{-ikr}}{ikr}
=\frac{1}{8\pi^2 ir}\int_0^\infty\!dk\,
\big[e^{-ik(\Delta t-r)}-e^{-ik(\Delta t+r)}\big].
\]
With \(\int_0^\infty dk\,e^{-ikb}=1/(ib)\) (\(\mathrm{Im}\,b<0\)):
\[
G^+=\frac{-i}{8\pi^2r}\cdot(-i)\cdot
\frac{2r}{(\Delta t-i\epsilon)^2-r^2}
=\frac{-1}{4\pi^2}\cdot\frac{1}{(\Delta t-i\epsilon)^2-r^2}.
\]
Since \((\Delta t-i\epsilon)^2-r^2=-s^2-2i\epsilon\Delta t\) with
\(s^2=-(\Delta t)^2+r^2=2\sigma\):
\[
\boxed{
G^+(x,x')=\frac{1}{8\pi^2}\cdot\frac{1}{\sigma+i\epsilon\Delta t}
+O(m^2,\ \text{regular}).
}
\tag{N1.1}
\]
For spacelike separation (\(\sigma>0\)) the \(i\epsilon\) is irrelevant:
\[
\boxed{G^+\sim\frac{1}{8\pi^2\sigma},\qquad
G^-\sim\frac{1}{8\pi^2\sigma}\quad\text{(spacelike)}.}
\tag{N1.2}
\]
\(G^-\) follows by \(x\leftrightarrow x'\) symmetry (\(\sigma\) symmetric;
the \(i\epsilon\) sign flips, immaterial at \(\sigma>0\)). The mass term
adds \(O(m^2\ln|\sigma|)\)-type subleading corrections, not a leading-pole
change. Disposition: DERIVED from the mode integral. The reviewer's
premise is independently confirmed; the sign is \(+\) (an earlier
scratch attempt dropped the angular-measure \(i\); recomputation fixes
it and the spacelike reality check \(G^+\in\mathbb R\) passes).

### N1.2 Convention table (frozen for the repair)

All entries at \(D=4\), signature \((-+++)\), \(\sigma=\tfrac12s^2\)
(Synge; DF Eq. (23) per `01`/doc-03-R1: "\(2\sigma\) the square of the
geodesic distance"). "Sing." = leading spacelike singularity.

| Symbol | Definition | Sing. | Distributional prescription | Scope |
|---|---|---|---|---|
| \(G^+\) | \(\langle\phi(x)\phi(x')\rangle\) | \(1/(8\pi^2\sigma)\) | \(\sigma+i\epsilon\Delta t\) boundary value (N1.1) | complex (timelike); real spacelike |
| \(G^-\) | \(G^+(x',x)\) | \(1/(8\pi^2\sigma)\) | \(\sigma-i\epsilon\Delta t\) | conjugate-related to \(G^+\) |
| \(G^{(1)}\) | \(G^++G^-\), exact sum (0.3) | \(1/(4\pi^2\sigma)\) | sum: \(i\pi\delta\) parts cancel, \(=2P(1/\sigma)/(8\pi^2)\) | real, symmetric |
| \(\tfrac12G^{(1)}\) | half the sum | \(1/(8\pi^2\sigma)\) | \(=P(1/\sigma)/(8\pi^2)\) spacelike | real, symmetric; \(=\mathrm{Re}\langle T\phi\phi'\rangle\) sing. part |
| \(H^{(1)}\) | \((1/8\pi^2)[U/\sigma+V\ln\|\bar\sigma\|]\) (0.8) | \(1/(8\pi^2\sigma)\) | \(P(1/\sigma)\), \(\ln\|\cdot\|\); \(\bar\sigma=\sigma/\ell^2\) | real, symmetric: the \(\tfrac12G^{(1)}\)-parametrix |
| \(G^F\) (DF) | \(i\langle T\phi\phi'\rangle\) (doc-03-R2) | \(i/(8\pi^2\sigma)\) | \(\sigma+i\epsilon\) (Feynman) | complex; \(-iG^F=\langle T\phi\phi'\rangle\) |
| \(S_{\rm old}\) | \(\tfrac12(G^{(1)}-H^{(1)})\) (0.5) | \(1/(16\pi^2\sigma)\) | — | SINGULAR: defect, void |
| \(S\) (new) | \(\tfrac12G^{(1)}-H^{(1)}\) (N1.5) | none (N2) | — | smooth real (P2-repaired, N2) |

Symbol hygiene (binding on N2–N5 and the fixture): \(H\) without a
superscript never appears; \(H^{(1)}\) always means the table entry
(the real symmetric \(\tfrac12G^{(1)}\)-parametrix); DF's Feynman
parametrix is always called out as such, never abbreviated \(H\).

### N1.3 Complete singular-structure check (DERIVED)

Leading pole. From the table:
\(\tfrac12G^{(1)}\sim1/(8\pi^2\sigma)=H^{(1)}_{\rm lead}\).
The leading singularities match exactly. The old \(S\) retained
\([(1/4\pi^2)-(1/8\pi^2)]/2=1/(16\pi^2)\) per \(\sigma\): singular.
Logarithm. Both \(\tfrac12G^{(1)}\) and \(H^{(1)}\) carry
\((1/8\pi^2)V\ln|\sigma|\)-type terms with the same coefficient
biscalar \(V\) (the Hadamard recursion fixes \(V\) from the universal
singularity; \(U=\Delta^{1/2}\) both sides per R5). The scale
difference \(\ln|\sigma|-\ln|\bar\sigma|=\ln\ell^2\) is smooth and
\(\ell\)-dependent (finite ambiguity, doc-04-§3.4 structure, N3.5).
Subleading smooth pieces differ — that difference is the remainder.

### N1.4 Distributional prescriptions and real-vs-complex scoping

- Wightman boundary values: \(G^\pm\) approach the light cone as
\((\sigma\pm i\epsilon\Delta t)^{-1}/(8\pi^2)\) (N1.1). The symmetric
combination cancels the \(\mp i\pi\delta(\sigma)\) halves, leaving the
real principal value: \(\tfrac12G^{(1)}\sim P(1/\sigma)/(8\pi^2)\).
- Feynman object: \(G^F\sim(i/8\pi^2)(\sigma+i\epsilon)^{-1}\)
(\(i\alpha_4/2\), `01` §3.2). Its real-time-ordered content:
\(\mathrm{Re}\langle T\phi\phi'\rangle
=\mathrm{Re}(-iG^F)\sim(1/8\pi^2)P(1/\sigma)\) — the same real
symmetric singularity as \(\tfrac12G^{(1)}\). This is why the real
symmetric parametrix \(H^{(1)}\) (not a complex Feynman parametrix)
is the correct subtraction for \(\tfrac12G^{(1)}\).
- Scoping rule: repaired claims about \(S\) concern the real smooth
biscalar. Smoothness of the real symmetric remainder does not by
itself assert anything about the full complex Wightman difference
\(G^+-H^+\); the commutator \((G^+-G^-)/2=(i/2)\Delta\) and the
\(i\epsilon\) terms are state-independent distribution theory outside
\(S\). N2's smoothness argument respects this boundary explicitly.

### N1.5 The A-vs-B algebraic identity (DERIVED)

\[
S_A:=\tfrac12G^{(1)}-H_A^{(1)},\qquad
S_B:=\tfrac12\big(G^{(1)}-H_B^{(1)}\big).
\]
\(S_A=S_B\iff H_B^{(1)}=2H_A^{(1)}\) as biscalars — the factor 2 must
cover the entire parametrix (pole term, log term, prefactor, and
every \(u_n\), \(v_n\) it multiplies), not merely the leading pole.
Adopted: Alternative A with \(H_A^{(1)}\) = the frozen (0.8). Reasons:
(i) (0.8) already carries DF's standard \(\alpha_4/2=1/(8\pi^2)\)
normalization (`01` §3.2) and matches the Feynman real-part
correspondence (N1.4); (ii) A moves only the \(S\)-definition, B would
redefine the verified \(H^{(1)}\) into a non-standard doubled object
\(H_B^{(1)}=(1/4\pi^2)[U/\sigma+V\ln|\bar\sigma|]\) contradicting the
DF record. The equivalence \(S_A=S_B\iff H_B=2H_A\) is recorded as the
verified identity above, not as an expectation.

### N1.6 Prescription freeze (REPAIRED)

\[
\boxed{
S(x,x'):=\tfrac12G^{(1)}(x,x')-H^{(1)}(x,x'),\qquad
(\Box_x-m^2)S=-(\Box_x-m^2)H^{(1)},
}
\tag{N1.3}
\]
the equation DERIVED in one line from P1 ((0.4) halved) with no value
assumed for the right-hand side. The value, the transport closure,
and smoothness are N2's blind computation. The subtraction
prescription (0.11) and operator (0.10) are unchanged in form and now
act on the redefined \(S\).

### N1.7 Stand/fall register (frozen §0 + locked derivations)

STAND as written (no \(S\)-dependence): Synge data (0.7); ansatz
(0.8)–(0.9) with DP; \(N\)-calculus (1.1)–(1.5); minimal-\(U\) (R5);
(V0); (V1a); (1.7) conditional on L2 (N2 derives the transport
content); L1/G2; E1/E2 (Synge arithmetic); (D1)+TI as structural
lemmas (re-applied in N3 after smoothness); §3 ambiguity/scale/trace
structures (numbers re-derived in N3); doc-05 entries without
\(S\)-dependence (adjudicated item-by-item in N5, presumed nothing).
FALL (void, re-derived in N2–N4): P2 as stated; (0.5); (0.6) with
\(H_{\rm sub}\); the (1.8)/(1.9) values; the (1.3)/(2.1)/(2.3)/(3.4)
numbers; the \([(\Box-m^2)A]=-3[v_1]\) constraint; doc-04-§4.1
numbers; doc-05 \(S\)-dependent applications (N5 register); doc-06
(T)/gate/O-register (revised 06); doc-07 verdict numbers; fixture
goldens (Step 7).

Gate N1: prescription frozen only with the table (N1.2), the
Minkowski check (N1.1), the full singular match (N1.3), the scoping
rule (N1.4), and the A/B identity (N1.5) all recorded. PASSED.

---

## N2. Remainder reconstruction (REPAIRED)

Blind computation from the frozen prescription (N1.6), DP, Synge data,
L1/G2, and P1. No value from (1.8)-old, no literature number, and no
cross-checker result is used or quoted below (comparison is N2.6,
after both derivations exist).

### N2.1 Smoothness of the redefined remainder

\(\tfrac12G^{(1)}\) carries the real symmetric singularity
\((1/8\pi^2)[U/\sigma+V\ln|\sigma|]\) (N1.2–N1.3: the two Wightman
boundary values \((\sigma\pm i\epsilon\Delta t)^{-1}\) sum to twice
the principal value; the \(\mp i\pi\delta(\sigma)\) halves cancel).
\(H^{(1)}\) carries exactly this singular structure (same \(U\),
same \(V\), same \(1/(8\pi^2)\) prefactor) by the Hadamard recursion.
Their difference is therefore free of pole and log singularities.
Premise P2-repaired (Hadamard state, assumptions A2/A7): \(S\) is
\(C^\infty\) on \(\mathcal C\), hence \(C^4\), the order consumed by
\(T_{\mu\nu}\) (two derivatives) plus one divergence plus continuity
of the result. Epistemic status, stated exactly: smoothness is the
Hadamard-state premise, not a theorem — but unlike the void P2,
which contradicted the \(1/(16\pi^2\sigma)\) singularity, P2-repaired
is coherent with the verified singular match (N1.3) and is exactly
the standard symmetric-form Hadamard condition. Real-vs-complex
boundary (N1.4): \(S\) is real; no claim is made about the full
complex Wightman difference \(G^+-H^+\). The commutator part
\((G^+-G^-)/2\) never enters \(S\). Disposition: subtraction
normalization DERIVED; smoothness PREMISE (A2), now coherent.

### N2.2 Remainder equation and its coincidence value (DERIVED)

Exact operator, argument, sign: \((\Box_x-m^2)\) acting on \(S\),
\((\Box_x-m^2)S=-(\Box_x-m^2)H^{(1)}\) (N1.3, one line from P1).
With \(N_2\equiv0\) (L1, stands) and the \(N\)-calculus (stands):
\[
\boxed{
J:=(\Box_x-m^2)S
=-\frac{1}{8\pi^2}\left[\frac{N_1}{\sigma}
+N_L\ln|\bar\sigma|\right]\ \text{smooth},
}
\tag{N2.1}
\]
\(N_1\), \(N_L\) as in locked (1.5)–(1.6). The prefactor is
\(1/(8\pi^2)\), twice the void (1.8)'s \(1/(16\pi^2)\), because the
frozen prescription subtracts the full \(H^{(1)}\), not half.
\(N_L\ln|\bar\sigma|\to0\) at coincidence (\([N_L]=0\) by (V1a),
stands). With (1.7) re-established in N2.3 (\(\lim N_1/\sigma
=6[v_1]\)):
\[
\boxed{
[J]=-\frac{6[v_1]}{8\pi^2}=-\frac{3[v_1]}{4\pi^2}.
}
\tag{N2.2}
\]
By \(x\leftrightarrow x'\) symmetry, \((\Box_{x'}-m^2)S\) has the
same coincidence limit. Disposition: DERIVED from N1.6 + locked
\(N\)-calculus + N2.3.

### N2.3 Transport equations and L2 (DERIVED; reviewer's §3 closed)

DP requires \(N_1/\sigma\) smooth. Split \(N_1\) with
\(V=v_0+v_1\sigma+R_2\):
\[
N_1=Z+\sigma\tilde K,\qquad
Z:=Q+2v_0+\tau v_0+2D_\sigma v_0,
\]
\(Q:=(\Box_x-m^2)\Delta^{1/2}\), where the \(v_1\sigma\)- and
\(R_2\)-pieces are manifestly \(\sigma\cdot(\text{smooth})\)
(\(2D_\sigma(v_1\sigma)=2\sigma D_\sigma v_1+4v_1\sigma\);
\(R_2=O(\sigma^2)\) gives \(\sigma^2\)-contributions). DP forces
\(Z/\sigma\) smooth. The Hadamard construction satisfies this via
the \(v_0\)-transport equation
\[
\boxed{
2D_\sigma v_0+(\tau+2)v_0+Q=0,\qquad [v_0]=-\tfrac12[Q],
}
\tag{N2.3}
\]
an ODE along each geodesic with initial data (V0) (uniqueness of the
Hadamard coefficients: SOURCE-REPORTED standard theory, same
standing as L1). Under (N2.3), \(Z\equiv0\), and L2's combination
\(M=(Q-[Q])+2(v_0-[v_0])+\tau v_0+2D_\sigma v_0\) satisfies
\(M\equiv-[Q]-2[v_0]=0\) by (V0). Hence
\[
\boxed{\lim_{x'\to x}M/\sigma=0,}
\tag{N2.4}
\]
i.e. L2 is DERIVED (strengthened to \(M\equiv0\)), no longer
imported. Then \(\lim N_1/\sigma=[\tilde K]=6[v_1]\) (terms:
\(2v_1\sigma/\sigma\to2[v_1]\);
\(2D_\sigma(v_1\sigma)/\sigma\to4[v_1]\); \(\tau v_1\),
\(2D_\sigma v_1\), \(R_2\)-terms \(\to0\)): (1.7) REPAIRED.
Recorded honestly: a non-standard representative with
\(Z=\sigma\cdot(\text{smooth}')\ne0\) would shift \([J]\) by
\([\text{smooth}']\); the transport equation (uniqueness) fixes the
standard representative. Order analysis: \([J]\) consumes \(V\) to
\(O(\sigma)\) (\(v_0\)-biscalar via (N2.3) \(+\) the \([v_1]\) value
via (V1a)); higher \(v_n\) never enter; \(N_L\ln\to0\) consumes
\([N_L]=0\) + \(V\in C^2\); the point-split divergence consumes
\(S\in C^4\) (N2.1). All finite.

### N2.4 Coincidence calculus (re-derived)

From Synge (0.12) with \(s=[S]\), \(s_\mu=[\nabla_\mu S]\):
\(\nabla_\mu s=2s_\mu\);
\(\nabla_\mu s_\nu=s_{\mu\nu}+s_{\nu\mu'}\);
\[
\boxed{\Box s=\nabla^\mu(2s_\mu)=2[\Box_xS]+2t,}
\tag{N2.5}
\]
confirming the E1-corrected form by fresh derivation (the void
(2.2) dropped the factor 2 from \(\nabla s=2s_\mu\)). DERIVED.

### N2.5 Flat-space falsification re-run (DERIVED)

Flat: \(v_0\equiv[v_0]=m^2/2\) constant, so
\([(\Box-m^2)v_0]=-m^4/2\) and (V1a) gives \([v_1]=m^4/8\ne0\).
Hence \([J]=-3m^4/(32\pi^2)\ne0\): the homogeneous-remainder
assertion is FALSIFIED again (Minkowski counterexample, doubled
numerical value). This re-runs doc-03-§1.7 under the repaired
prescription; the conclusion (inhomogeneous, nonvanishing source)
is unchanged, the number is doubled.

### N2.6 Executor/cross-checker comparison (checker #1)

Quarantined checker #1 (brief-only, no audit/plan reads, compliance
confirmed) returned full derivations (a)–(e). Agreement: \([J]
=-3[v_1]/(4\pi^2)\) (integer \(6=4+2\)); (V0); (V1a); exact
\(v_0\)-transport with \(M\equiv0\); \(\Box s=2([J]+m^2s+t)\)
(E1-form); flat \([v_1]=m^4/8\), \([J]=-3m^4/(32\pi^2)\ne0\).
AGREEMENT on all of N2. O1a (remainder limit + transport + L2):
CLOSED by two independent derivations.

DISAGREEMENT on the residual factor — the checker's adversarial
finding: via an independent Synge-A route (not T1/TI) it derived
\(\nabla^\mu R_{\mu\nu}=[\nabla_{\nu'}J]\) EXACTLY (needs only
\([\nabla g]=0\), flag F1), and \([\nabla'J]=\tfrac13\nabla[J]\)
given v1-symmetry-halves (F2) + \([\nabla\tau]=[\nabla'\tau]=0\)
(F3). I.e. residual factor \(1/3\), not the \(1/2\) of N3.1. It
further flagged F4: exact DP (\(R_0\equiv0\)) is needed for
\(\nabla[J]\) (else the \(N_L\ln\)-gradient diverges logarithmically).
The executor VERIFIED the refutation (N4.2: the locked T1/TI proofs
conflate \(J=(\Box-m^2)S\) with \(J'=(\Box'-m^2)S\); corrected TI
gains \(-2\delta\) and gives \([\nabla'J]\), two routes agreeing)
and completed closure (N4.3 v1-transport, N4.4 F2/F3/split).
N3.1/N3.2/N3.6-numbers are SUPERSEDED by N4.5; N2 stands intact.
O1b (\(\delta\)/\(J'\)-structure, \(1/3\), v1-transport): checker #2
RETURNED full agreement — residual \([\nabla'J]\) via a third route
(B/C bitensor divergence), \(R_0\equiv0\) derived with
\([\nabla(N_L\ln)]=[\nabla'(N_L\ln)]=0\),
\([\nabla\tau]=[\nabla'\tau]=0\) derived from
\(s^{;u}s_{;u}=2s\) (no \(\sigma\)-expansion import), Q1-split
\(8w\)/\(4w\)/\(12w\) with the F-sym1-necessity note, residual
\(-\nabla[v_1]/(4\pi^2)\), correction \(+1/(4\pi^2)\). Its flags
all close: F1\(\to\)DF-(66b), F-sym\(\to\)frozen (0.7),
F-sym1\(\to\)A9, R2\(\to\)definitional, UNI\(\to\)standard
(same standing as N2.3-uniqueness), F-JK\(\to\)exact identity,
EXT\(\to\)corroborated-not-load-bearing. O1b CLOSED.

---

## N3. Divergence reconstruction (REPAIRED)

### N3.1 Residual (DERIVED)

(D1)+TI are structural lemmas (re-verified line by line in this
repair: every Synge instance, both Ricci terms with the L3 \(+\)
sign, the \(m^2s_\nu\) and \(\tfrac12\nabla_\nu[J]\) cancellations).
For any smooth symmetric \(S\) with symmetric source \(J\):
\(\nabla^\mu R_{\mu\nu}=\tfrac12\nabla_\nu[J]\). Under P2-repaired
(N2.1) with (N2.2):
\[
\boxed{
\nabla^\mu\langle T_{\mu\nu}\rangle_{\rm raw}
=\tfrac12\nabla_\nu[J]
=-\frac{3}{8\pi^2}\,\nabla_\nu[v_1].
}
\tag{N3.1}
\]
The raw divergence is the gradient of the local curvature scalar
\(-3[v_1]/(16\pi^2)\). Flat check: \([J]\) constant \(\Rightarrow\)
residual zero (translation invariance). DERIVED from N2.1+N2.2.

### N3.2 Correction and cancellation (DERIVED)

(N3.1) is a pure gradient of a local scalar. With the ansatz
\(C_{\mu\nu}=K\,g_{\mu\nu}[v_1]\) (\(K\) constant):
\(\nabla^\mu C_{\mu\nu}=K\,\nabla_\nu[v_1]\) (metric compatibility).
Cancellation of (N3.1) forces \(K=+3/(8\pi^2)\):
\[
\boxed{
C_{\mu\nu}=+\frac{3}{8\pi^2}\,g_{\mu\nu}[v_1],\qquad
\nabla^\mu C_{\mu\nu}=+\frac{3}{8\pi^2}\,\nabla_\nu[v_1].
}
\tag{N3.2}
\]
With \(\langle T_{\mu\nu}\rangle_{\rm ren}:=R_{\mu\nu}+C_{\mu\nu}
+\Theta_{\mu\nu}\) (\(\nabla^\mu\Theta_{\mu\nu}=0\)):
\[
\boxed{
\nabla^\mu\langle T_{\mu\nu}\rangle_{\rm ren}
=-\frac{3\nabla_\nu[v_1]}{8\pi^2}
+\frac{3\nabla_\nu[v_1]}{8\pi^2}+0=0.
}
\tag{N3.3}
\]
DERIVED. The residual and the correction divergence are computed
from the same repaired remainder equation (N2.1); cancellation is
exact in N1 conventions. No separate conservation of \(C_{\mu\nu}\)
is assumed or holds (N3.4).

### N3.3 Uniqueness limits (Amendment 4; binding language)

A computed divergence does not by itself determine a tensor. The
repair distinguishes three layers and bans the word "uniquely"
unless layer (c) is established in the same document:
- (a) Minimal term: (N3.2) is the minimal local metric-proportional
conservation-restoring term — DERIVED (forced by (N3.1) given the
metric-proportional ansatz, which Wald 1978's abstract structure
"subtract that term times the metric" motivates and which the
gradient form of (N3.1) makes sufficient).
- (b) Conserved-addition freedom: any separately conserved local
curvature tensor may be added without changing the divergence.
Under the Wald axioms alone this freedom is the infinite
Tichy–Flanagan family (`01` §4.5, SOURCE-REPORTED) — DERIVED
(structural: metric variations of local invariants are conserved
by Theorem V).
- (c) Restriction assumptions: the polynomial dimension-four basis
(3.3) holds CONDITIONAL on Salcedo's renormalizability criterion
(A6, `01` §4.6); it is not derived from the Wald axioms.
The repaired theorem claims (a)+(b) with (c) conditional. The
coefficient \(3/(8\pi^2)\) is derived, not imported; it is the
coefficient of the minimal term, not a uniqueness claim.

### N3.4 The correction is not independently conserved

From (N3.2): \(\nabla^\mu C_{\mu\nu}=+3\nabla_\nu[v_1]/(8\pi^2)\),
generically nonzero — its nonvanishing is what cancels (N3.1).
Re-asserted under repair (was doc-04-§2.2). DERIVED (explicit).

### N3.5 Scale absorption re-run (DERIVED)

\(\ell\to\ell'\): \(\Delta H^{(1)}=(1/8\pi^2)V\cdot(-2\ln(\ell'/\ell))\),
so \(\Delta S=-\Delta H^{(1)}=(1/8\pi^2)V\cdot2\ln(\ell'/\ell)\)
(no \(\tfrac12\): doubled vs the void §3.4, smooth symmetric
local-geometric). \(J_{\Delta S}=(\ln(\ell'/\ell)/4\pi^2)N_L\),
\([J_{\Delta S}]\equiv0\) ((V1a) everywhere by DP), hence
\(\nabla^\mu\lim T_{\mu\nu}\Delta S=\tfrac12\nabla_\nu[J_{\Delta S}]
=0\): \(\ell\)-changes shift the raw tensor by a conserved local
geometric tensor, absorbable into \(\Theta_{\mu\nu}\). Residual
(N3.1) and correction (N3.2) are \(\ell\)-independent. DERIVED.

### N3.6 Trace re-run (DERIVED)

From (1.1) with (N2.5): \(t=\tfrac12\Box s-m^2s-[J]\), so
\(\mathrm{tr}_{\rm raw}=-t-2m^2s
=-\tfrac12\Box s-m^2s+[J]\). With (N2.2) and
\(\mathrm{tr}\,C=4\cdot3[v_1]/(8\pi^2)\):
\[
\boxed{
\mathrm{tr}_{\rm ren}
=-\tfrac12\Box s-m^2s+\frac{3[v_1]}{4\pi^2}
+\Theta^\mu{}_\mu.
}
\tag{N3.4}
\]
Generically nonzero; compatible with (N3.3) (trace and divergence
are independent projections). DERIVED. Fixture-consumed rational
cores (hand-derived here, pinned in Step 7): TRACE
\(-3/4+3/2=+3/4\) (units \([v_1]/\pi^2\)); FLAT-CHAIN
\((1/8)\cdot(-3/4)=-3/32\); NONZERO \(-3/32\ne0\). COMBO-6
(\(2+4=6\)), HALF (\(\tfrac12\cdot6=3\)), CANCEL (\(-3+3=0\))
skeletons are structurally unchanged (the residual/correction
numerators stay \(\mp3\); the shared symbolic denominator moves
\(16\pi^2\to8\pi^2\)).

Gate N3: residual DERIVED from N2, correction forced (not
imported), cancellation demonstrated, uniqueness layered, scale
and trace re-run. PASSED (as the symmetric-part computation;
see N4.2/N4.5 for the TI correction that supersedes N3.1/N3.2/
N3.6-numbers — the J/J' distinction).

---

## N4. DF reconciliation, TI correction, derived coefficient

### N4.0 Verdict up front

DF's (48)–(111) chain was re-inspected verbatim from the full text
(ar5iv gr-qc/0512118, fetched 2026-11-06). Findings: (i) `01`'s
transcription of DF's coincidence equation carries a sign error
(N4.1; corrected here and in revised-06, `01` itself byte-identical);
(ii) the locked T1/TI proofs conflate \(J=(\Box-m^2)S\) with
\(J'=(\Box'-m^2)S\) (N4.2); corrected TI gains \(-2\delta\) and the
residual is \([\nabla'J]\), confirmed by two routes; (iii) with
v1-transport derived (N4.3), F2-weak/F3 established (N4.4), and the
Q1-split verified, the residual is \(-\nabla[v_1]/(4\pi^2)\) and the
correction \(+g[v_1]/(4\pi^2)\) — DERIVED, matching DF exactly, not
imported (N4.5); (iv) the A-term carries \(([(\Box-m^2)A],
[\nabla'J_A])=(0,0)\), i.e. a conserved contribution (N4.5); explicit
\(A\) stays open (O2 refined). DF-equivalence: RECONCILED at every
number; full framework identification CONDITIONAL on explicit \(A\).

### N4.1 DF re-inspection record (INDEPENDENTLY REPRODUCED)

Verified verbatim from the fetched full text: (48)
\(G^F_{\rm reg}=(i\alpha_D/2)W\); (51a) \(w_a=\tfrac12w_{;a}\) plus
the text "\(\ldots\) constrained by the symmetry of \(W(x,x')\)"
(DF's \(W\) IS symmetric); (52)/(55) (as in `01`); (56a/b)/(57) (as
in `01`); (60) \(\langle T\rangle=\lim\mathcal T[-iG^F]\); (61)
\(\mathcal T_{\mu\nu}\) (as in `01`); (63) replacement \(\to
(\alpha_D/2)\lim\mathcal TW\); (64); (66b)
\(g_\nu{}^{\nu'}g_{\mu\nu';\rho}=-\tfrac12R_{\mu\nu\rho a}\sigma^{;a}
+O(\sigma)\) (hence F1: \([\nabla g]=0\)); the unnumbered coincidence
equation between (66c) and (68) (navigation label "(67')" — the paper
numbers no equation (67)); (68)/(69)/(70); the unnumbered (71');
(74); (75a/b) \(w=[W]\), \(w_{\mu\nu}=\lim W_{;\mu\nu}\) (both
derivatives unprimed); (76)/(77) scale terms; (108a)/(109)
(\(v_0\), \(v_1\) exactly as in `01`).

CORRECTION to `01` §4.1 (recorded here; `01` frozen): DF's (67')
reads \(-\,(w_{\mu\nu}-\tfrac12g_{\mu\nu}w^\rho{}_\rho)\), i.e. the
\(w^\rho{}_\rho\) term is \(+\tfrac12g\,w^\rho{}_\rho\), not
\(-\tfrac12g\,w^\rho{}_\rho\) as transcribed. With the corrected
sign: (67') reproduces \(\lim\mathcal T^{DF}W\) exactly in
\(s\)-language (verified: \(-s_{\mu\nu}+\tfrac12g[\Box W]
+\tfrac12w_{;\mu\nu}-\tfrac14g\Box w-\tfrac12gm^2w
=s_{\mu\nu'}-\tfrac12g(t+m^2w)\) using
\(w_{;\mu\nu}=2s_{\mu\nu}+2s_{\nu\mu'}\) and E1); the (68)-algebra
gives \(-2\nabla v_1\) (all \(\Box w\)/Ricci/\(m^2\) pieces cancel);
and (71')\(\to\)(111) reproduces the universal conformal anomaly
\((2880\pi^2)^{-1}(\Box R-R_{ab}^2+R_{abcd}^2)\), anomaly-locking
DF's correction \(+gv_1/(4\pi^2)\). DF's chain is internally
consistent. (75b) was additionally verified from
(50)+(51a)+(54): \([\nabla_\mu\nabla_\nu W]=\nabla_\mu\nabla_\nu w
-\nabla_\nu w_\mu-\nabla_\mu w_\nu+w_{\mu\nu}=w_{\mu\nu}\).

### N4.2 J/J' distinction and TI correction (DERIVED, two routes)

\(J:=(\Box_x-m^2)S\), \(J':=(\Box_{x'}-m^2)S\). From \(S\) symmetric,
\(J'(x,x')=J(x',x)\) (definitional swap, one line). \(J\) is NOT
symmetric as a biscalar in generic curved spacetime:
\((\Box'-m^2)H^{(1)}-(\Box-m^2)H^{(1)}
=(8\pi^2)^{-1}[(N_1-N_1')/\sigma+(N_L-N_L')\ln|\bar\sigma|]\ne0\)
off-diagonal (unprimed/primed derivatives of \(\sigma\),
\(\Delta^{1/2}\), \(V\) differ; they agree only in maximally
symmetric cases). Write \(a=[\nabla J]\), \(a'=[\nabla'J]\),
\(b=[\nabla J']\), \(b'=[\nabla'J']\); swap gives \(b=a'\), \(b'=a\).

Re-examined T1: the first bracket uses \([\nabla_\nu\Box'S]
=m^2s_\nu+b=m^2s_\nu+a'\); the second uses
\([\nabla_\nu\Box S]=m^2s_\nu+a\). So
\(\nabla t=\nabla s_{\mu\nu'}+\nabla s_{\mu\nu}-2m^2s-(a+a')
-Rs\) with \(a+a'=\nabla[J]\) EXACT (Synge, no symmetry needed):
T1 and (D1) are UNCHANGED. Re-derived TI: the first bracket uses
\([\nabla_{\nu'}\Box S]=m^2s+a'\); the second uses
\([\nabla_{\nu'}\Box'S]=m^2s+b'=m^2s+a\) (plus \(R_ss\) via primed
L3). Hence the CORRECTED transport identity
\[
\boxed{
\nabla^\mu s_{\mu\nu'}=\nabla_\mu s^\mu{}_\nu
-R_\nu{}^\sigma s_\sigma-2\delta,\qquad
\delta:=\tfrac12(a-a')=\tfrac12([\nabla J]-[\nabla'J]).
}
\tag{N4.1}
\]
(The locked TI proof wrote \([\nabla_{\nu'}J]\) for both halves —
valid only under J-symmetric.) With (D1):
\[
\boxed{
\nabla^\mu R_{\mu\nu}
=\tfrac12\nabla[J]-\delta=[\nabla'J].
}
\tag{N4.2}
\]
Exact given Synge + L3 + F1 (\([\nabla g]=0\), DF-(66b)) +
\(J/J'\)-swap. INDEPENDENT corroboration: checker #1's Synge-A
route (differentiating \(A_{\mu\nu}=T_{\mu\nu}S\) directly, never
via T1/TI) gives \(\nabla R=\tfrac12([\nabla'J]+[\nabla J'])
=[\nabla'J]\) by the same swap. Two routes, one result. The naive
\(\tfrac12\nabla[J]\) (N3.1) is the symmetric part of the true
\([\nabla'J]\); it missed \(-\delta\).

### N4.3 v1-transport and J-C1 closure (DERIVED)

\(N_L=R_0+\sigma\tilde K^R\) with
\(R_0:=(\Box-m^2)v_0+4v_1+2D_\sigma v_1+\tau v_1\) (from
\((\Box-m^2)(v_1\sigma)=\sigma(\Box-m^2)v_1+2D_\sigma v_1
+v_1(4+\tau)\) and \((\Box-m^2)R_2=O(\sigma)\)). DP requires
\(N_L\ln|\bar\sigma|\) smooth; \([N_L]=0\) ((V1a)) gives only
\(C^0\) (\(\nabla(R_0\ln)\) diverges as \((\nabla R_0)\ln\) unless
\(R_0\) vanishes identically). By the same DP+uniqueness argument
as N2.3, the Hadamard construction satisfies the v1-transport
equation
\[
\boxed{
R_0\equiv0,\qquad\text{i.e.}\quad
2D_\sigma v_1+(\tau+4)v_1+(\Box-m^2)v_0=0,
}
\tag{N4.3}
\]
a first-order linear ODE along geodesics with data (V1a)
(\([R_0]=[(\Box-m^2)v_0]+4[v_1]=0\)); unique regular solution.
Corroboration: P2-repaired (\(S\) smooth) forces \(J=(\Box-m^2)S\)
smooth, hence \(R_0\equiv0\) by consistency (else \(J\notin C^1\)).
Then \(N_L=\sigma\tilde K^R\), \(N_L\ln|\bar\sigma|=O(\sigma\ln\sigma)\),
and \([\nabla(N_L\ln|\bar\sigma|)]=[\nabla'(N_L\ln|\bar\sigma|)]=0\)
(\((\nabla N_L)\ln=O(d\ln d)\to0\) since \([\nabla N_L]=0\);
\(N_L\nabla\sigma/\sigma=O(d)\to0\)). So with \(N_1/\sigma=\tilde K\)
(N2.3): \(J=-(8\pi^2)^{-1}[\tilde K+O(\sigma\ln\sigma)]\),
\([\nabla J]=-(8\pi^2)^{-1}[\nabla\tilde K]\),
\([\nabla'J]=-(8\pi^2)^{-1}[\nabla'\tilde K]\). \([J]\) (N2.2) is
unaffected (it used only \([R_0]=0\)). Scope: v1-transport is the
LAST required order (\(J\in C^1\) suffices for \([\nabla'J]\);
\(R_1\equiv0\) etc. are not consumed).

### N4.4 F2-weak, F3, and the Q1-split (DERIVED)

F2-weak: \([\nabla v_1]=[\nabla'v_1]\) (\(=w\), so
\(\nabla[v_1]=2w\)): SOURCE-REPORTED standard Hadamard theory
(symmetric \(V\)-coefficients; cf. DF §II.2, Christensen 1976/78,
Brown–Ottewill 1986) — same standing as L1. Only this weak
(first-derivative-coincidence) form is consumed, not full biscalar
symmetry. F3: \([\nabla\tau]=[\nabla'\tau]=0\): DERIVED from
\(\tau=\Box_x\sigma-4=-(1/3)R_{ab}\sigma^{;a}\sigma^{;b}
+O(\sigma^{3/2})\) (DF (54) contracted; standard), whose every
first derivative is \(O(d)\to0\). Q1-split (independently verified
here): with \(Q_1:=\tilde K=6v_1+2D_\sigma v_1+\tau v_1\),
\([\nabla_\nu D_\sigma v_1]=[\nabla_\nu v_1]=w_\nu\) (via
\([s^{;a}{}_{;\nu}]=\delta^a_\nu\)),
\([\nabla_{\nu'}D_\sigma v_1]=-[\nabla_\nu v_1]=-w_\nu\) (via
\([s^{;a}{}_{;\nu'}]=-\delta^a_\nu\)), so
\([\nabla Q_1]=8w\), \([\nabla'Q_1]=4w\), \(\nabla[Q_1]=12w\):
\[
\boxed{[\nabla'Q_1]=\tfrac13\nabla[Q_1].}
\tag{N4.4}
\]
(F2-weak is necessary: without \(u=u'\) the ratio is
\((3u'-u)/(3u+3u')\ne\tfrac13\).) Hence
\([\nabla'J]=\tfrac13\nabla[J]\) and
\(\delta=\tfrac12\nabla[J]-[\nabla'J]=\tfrac16\nabla[J]
=-\nabla[v_1]/(8\pi^2)\).

### N4.5 Residual, correction, cancellation, A-data (DERIVED)

From (N4.2), (N4.4), (N2.2):
\[
\boxed{
\nabla^\mu\langle T_{\mu\nu}\rangle_{\rm raw}
=[\nabla'J]=\tfrac13\nabla[J]
=-\frac{1}{4\pi^2}\,\nabla_\nu[v_1].
}
\tag{N4.5}
\]
Forced correction (\(C_{\mu\nu}=Kg_{\mu\nu}[v_1]\),
\(\nabla C=K\nabla[v_1]\)):
\[
\boxed{
C_{\mu\nu}=+\frac{1}{4\pi^2}\,g_{\mu\nu}[v_1],\qquad
\nabla^\mu C_{\mu\nu}=+\frac{1}{4\pi^2}\,\nabla_\nu[v_1],
}
\tag{N4.6}
\]
\[
\boxed{
\nabla^\mu\langle T_{\mu\nu}\rangle_{\rm ren}
=-\frac{\nabla_\nu[v_1]}{4\pi^2}
+\frac{\nabla_\nu[v_1]}{4\pi^2}+0=0.
}
\tag{N4.7}
\]
DERIVED — the coefficient \(1/(4\pi^2)\) is forced by the repaired
residual through the verified \(1/3\) split, not imported from DF
(no DF number is a derivation input anywhere in N1–N4; DF enters
only as the N4.1 comparison, which now matches exactly). The
uniqueness layers and non-conservation (N3.3/N3.4) carry over
verbatim with the new numbers. Trace (N3.6 re-run):
\(\mathrm{tr}_{\rm ren}=-\tfrac12\Box s-m^2s+[J]
+4[v_1]/(4\pi^2)+\Theta^\mu{}_\mu\), i.e.
\[
\boxed{
\mathrm{tr}_{\rm ren}
=-\tfrac12\Box s-m^2s+\frac{[v_1]}{4\pi^2}
+\Theta^\mu{}_\mu,
}
\tag{N4.8}
\]
whose geometric piece \([v_1]/(4\pi^2)\) formally reproduces DF
(111) under conformal \([v_1]\) — independent corroboration.
A-data: \(W=8\pi^2S+A\) (definitional; \(A\) smooth symmetric)
gives \([(\Box-m^2)A]=-6v_1-8\pi^2[J]=0\) (DERIVED from N2.2 +
DF-(55)) and \([\nabla'J_A]=[\nabla'J_W]-8\pi^2[\nabla'J_S]
=-2\nabla v_1-8\pi^2(-\nabla[v_1]/(4\pi^2))=0\) (DERIVED-combination
from DF's \(\nabla\lim\mathcal TW=-2\nabla v_1\) via the structural
(N4.2) applied to \(W\)); hence \(\nabla\lim TA=[\nabla'J_A]=0\):
\(A\) contributes a conserved local tensor (absorbed into
\(\Theta\)). Explicit \(A\): still open (O2 refined — blocks only
full framework identification, not the internal theorem).

Gate N4: DF chain re-inspected; 01-(67) corrected; TI corrected
(three routes: revised-T1/TI, Synge-A, B/C-bitensor); v1-transport
derived (two derivations); F2/F3 established; Q1-split verified
(twice); coefficient DERIVED matching DF; A-data derived.
PASSED (O1b closed; checker #2 returned full agreement, N2.6).

---

## N5. Neutral finding re-audit (REPAIRED)

### N5.1 Provisional register

At repair start, every prior finding depending directly or
indirectly on the defective \(S\) was marked PROVISIONAL with no
presumed outcome (Amendment 3): doc-05 C2, E1, E4, E5, E6, E7,
F2, G2-first-premise, G3-raw-cite, H1, H3, H4-verdict, H8,
J3, L9, D5-structural, K1; doc-06 (T)/gate/O-register/Phys-Pseudo;
doc-07 verdict/answers; all fixture goldens. Each is re-adjudicated
below with its dependency chain. Outcome: most FALSIFIED verdicts
are CONFIRMED (several with corrected reasoning/numbers); L9 is
UPGRADED (conclusion now DERIVED with DF-matched coefficient);
K1 is AMENDED (01-(67) sign); the (T)/gate/O-register/07/fixture
revisions are recorded in Steps 6–7, not presumed here.

### N5.2 Four separate questions on Agent B's subtraction

Agent B's \(W_B=G^+-H^{(1)}\) (C2/I10 intent). Re-derived
decomposition: \(G^+=G^{(1)}/2+(i/2)\Delta\) (\(\Delta\) the real
Pauli-Jordan commutator function), so \(W_B=S+(i/2)\Delta\) with
\(S\) the repaired remainder.
1. Singularity removal? NO — \((i/2)\Delta\) is light-cone
singular (standard distribution theory) and \(W_B\) is complex
non-symmetric. DERIVED (N2.1 + \(\Delta\)-theory).
2. Real part smooth? YES — \(\mathrm{Re}\,W_B=S\) (smooth by
N2.1; \(\Delta\) real). DERIVED. (This sharpens C2: the real
part is exactly the repaired remainder.)
3. Full complex remainder smooth? NO — (1). DERIVED. (2)\(\not
\Rightarrow\)(3), as Amendment 1 requires.
4. Field equation? \((\Box-m^2)W_B=J+0=J\) (inhomogeneous;
\((\Box-m^2)\Delta=0\) standard): neither Agent B's claimed
homogeneous-0 nor smooth-irrelevant. DERIVED (N2.2).
Each answer cites separate evidence; none establishes another.

### N5.3 Item re-adjudication (doc-05 entries)

C2: FALSIFIED (CONFIRMED, reasoning replaced by N5.2:
\(W_B=S+(i/2)\Delta\), light-cone-singular + complex +
non-symmetric; \(\lim TW_B\) ill-defined as real symmetric).
Depends: N2.1. E1/E2: FALSIFIED-as-applied (CONFIRMED; depend:
re-derived C2). E4: FALSIFIED-central (CONFIRMED; refutation now
N2.5 \([J]=-3m^4/(32\pi^2)\ne0\)). E5: FALSIFIED (CONFIRMED;
corrected number: drops \(\tfrac12[\nabla'J]
=\tfrac16\nabla[J]\ne0\), not \(\tfrac14\nabla[J]\); depends:
N4.2/N4.4). E6: FALSIFIED (CONFIRMED; true residual \([\nabla'J]\)
(N4.5)). E7: FALSIFIED-contradiction + unsubstantiated
(CONFIRMED; minimum correction now (N4.6) with non-conservation
explicit; depends: N3.4/N4.5). F2: statement
INDEPENDENTLY-REPRODUCED with (N4.7)+(N4.8), argument FALSIFIED
(CONFIRMED). G2: FALSIFIED-as-Agent-B's-argument (CONFIRMED;
true premises now (N4.7) + correct \(H\)-tensors). G3: FALSIFIED
(CONFIRMED; raw (1.1) still has (N4.5) divergence). H1/H3:
FALSIFIED (CONFIRMED; depend: E4–E7/C2 re-derived). H4-verdict:
FALSIFIED (CONFIRMED; no subtraction analysis + C2). H8:
CONDITIONAL (CONFIRMED; condition now the N1–N4 subtraction).
J3: FALSIFIED-incomplete (CONFIRMED; omits (N4.6); depends: N4.5).
L9: TRUE (UPGRADED: conclusion DERIVED with DF-matched
coefficient (N4.7), not merely conditional). D5-structural:
INDEPENDENTLY-REPRODUCED (CONFIRMED; re-proof:
\(\nabla\lim T\Delta S=[\nabla'J_{\Delta S}]
=\mathrm{const}\cdot[\nabla'N_L]=0\) since \(N_L=O(\sigma)\) by
(N4.3); specific \([V_1]g\)-form still FALSIFIED). K1: AMENDED
(DF navigation stands EXCEPT the (67)-sign, N4.1; re-verified
verbatim). All other doc-05 entries (A1–A6, B, C1, C3–C7, D1–D4,
F1, G1, G4–G6, H2, H5–H7, I1–I12, J1–J2, J4–J5, K2–K10, L1–L8):
S-INDEPENDENT on inspection (bibliography, classical core,
\(T\)-operator, \([U]\), \(V\)-coefficient/a2 comparisons,
\(H\)-tensors/dimensions, trace-formula, Friedmann, syntax,
summaries, dependency records) — verdicts STAND as written, each
checked for \(S\)-dependence rather than presumed (I9/I10 re-point
to N2.5/re-derived-C2 respectively).

### N5.4 O1 decision (refined, per review §4/O-table)

O1 splits by what the two independent derivations closed. O1a
(remainder limit \([J]\) + \(v_0\)-transport + L2): CLOSED —
executor N2 + checker #1 (a–c,e) agree exactly, both blind.
O1b (\(\delta\)/\(J'\)-structure, \([\nabla'J]=\tfrac13\nabla[J]\),
v1-transport): CLOSED — checker #1 (d) found \([\nabla'J]\)
+ \(1/3\) + flags F1–F4 independently; executor verified via
revised-T1/TI + Q1-split (N4.2/N4.4, different route); checker #2
returned full independent agreement on TI-correction, v1-transport
\(R_0\equiv0\), F3 (derived from \(s^{;u}s_{;u}=2s\)), Q1-split,
and the \(1/(4\pi^2)\) coefficient (N2.6). O1 CLOSED in full.
O2: refined — explicit-\(A\) determination open; blocks ONLY full
framework identification with DF, NOT the internal theorem (N4.7
is exact in N1 conventions with DF-matched numbers; A-data
\((0,0)\) derived). O3: unchanged (blocks shape-claims only).
O4: unchanged (blocked-external, unneeded). O5: updated — fixture
repaired to Step-7 scope; machine-verification limits still bound
it (SK-records). 06/07/JSON agreement: Step 6.

Gate N5: provisional register cleared item-by-item with chains;
four questions answered separately; O-register decided. PASSED.
