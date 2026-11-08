# BHC + ALG cores and MRC coverage matrix, v0.2 (phys-pseudo draft)

Status: PROPOSED. Nothing below is asserted true. Representation is not endorsement.
Syntax: the seven constructs only (given, define, assume, claim, derive, if/then/else, ?).
The surface form is provisional; conform it to PSEUDO.md before agents consume it.
Items marked [RATIFY] are draft-chosen values. Items marked [RECALL] are literature claims
made from memory and not yet through source intake. Items marked [SEARCHED] were checked
against retrieved sources in the 2026-11-06 session.

Branches:
- BHC (Branch C): Bohmian matter + classical (unquantized) gravity.
- ALG (Branch D): observer-dressed operator-algebra construction, gravity quantum perturbatively.
Both run through the same five pilots, the same rungs and the same audit.

---

## Change log (PR3)

```
claim V1 : v0.1 had BHC only. v0.2 adds ALG as a second branch (reason: ALG covers parts of R1, R2, R4
           from one construction but denies P3 and has no known place for P1; the comparison is informative)
claim V2 : R3 gains a Born-route table and a relocation inventory (reason: the Born "wall" is two walls,
           weights and outcomes; Gleason bears only on weights)
claim V3 : R2 gains a domain check and an explicit BLOCKED(domain) outcome (reason: no self-adjoint
           constraint operator means no spectral theorem)
claim V4 : R4 gains scope limits: islands are derived in AdS-type or flat-with-bath setups; unitarity is a
           declared premise, not a derived result (reason: Wigner presupposes bijectivity)
claim V5 : R5 gains a conditional falsifier tied to evolving dark energy (reason: a pure integration-constant
           Lambda cannot be the full story if w != -1 is confirmed)
claim V6 : R1 gains a gravitationally-induced-entanglement identifiability table (reason: the BMV inference
           "entanglement implies nonclassical gravity" is contested; BHC is a model neither side covers)
claim V7 : all pilots gain footprint, assumption-graph, silence-class and interface-contract fields
claim V8 : PR6 added: every proposed discriminator gets a one-line sanity check before entering the matrix
```

---

## 0. Pre-registration rules (fixed BEFORE any pilot computation)

```
assume PR1: each branch's core is its postulate list in sections 2 and 3 and nothing else
assume PR2: any further premise a pilot needs is an AUX; it is declared at the point of use,
            tagged with the pilot and branch ids, and counted in section 5
assume PR3: rung targets and failure conditions in section 4 are not edited after a pilot starts;
            a change is logged as a new version with the reason
assume PR4: a pilot ends in exactly one terminal state per branch:
            DERIVES | EXPLAINS | ACCOMMODATES | POSTULATES | REFUTES | BLOCKED
            UNDECLARED is not a terminal state; it is a defect that fails the audit
assume PR5: REFUTES names the core postulate contradicted; BLOCKED names the obstruction.
            Neither deletes a branch.
assume PR6: before a proposed discriminator enters a matrix cell it must pass a one-line sanity check:
            (a) units and dimensions; (b) does it survive the classical-equivalence check (is the model
            already equivalent to a standard one on this observable?); (c) is any key object trivial in
            the limit used (for example a modular flow of a tracial state)? A failure is logged, not hidden.
assume PR7: literature claims are tagged [RECALL] until source intake; a claim contested in the literature
            is recorded with its status "contested" and not used as a premise
assume PR8: both branches use the same rungs, thresholds and audit; neither is given extra AUX allowance
```

Rung definitions, shared by all pilots and branches:

```
define DERIVES(X)      := a derive of X from the core alone, zero AUX, passing the section 6 debt list
define EXPLAINS(X)     := X follows from the core plus at most one AUX, where the AUX is used in at least one
                          other pilot or is a standard piece of the cited framework; OR the core fixes the
                          qualitative structure of X but not its numerical value
define ACCOMMODATES(X) := X is reproduced only after choosing a free element the core does not fix
define POSTULATES(X)   := X is added as an assumption
define REFUTES(P)      := a derive, counterexample or toy check contradicts core postulate P
define BLOCKED(why)    := the required derivation is not completed; why names the specific obstruction
```

Silence classes (recorded per pilot beside the rung):

```
define RELOCATED(S)      := the silence S is replaced by a named, smaller assumption inventory
define HALF_RELOCATED(S) := part of S is relocated and a named remainder is untouched
define GENUINE(S)        := no mapping or intersection touches S; it is an input or boundary datum
```

---

## 1. Interface contracts (the common-language check)

Three languages, one contract per pair. These are claims to be tested, not settled facts.

```
define LANG_M := matter: a *-algebra of observables A with states as positive normalized functionals;
                 probabilities as measures on the projection lattice
define LANG_S := spacetime: a Lorentzian manifold (M, g) with causal structure
define LANG_D := dynamics: automorphism flows on A (Heisenberg picture) or constraints on a phase space

define I_MS := matter <-> spacetime
  preserved : locality and covariance, if the map from spacetimes to algebras is a functor (locally covariant QFT)
  lost      : the state-dependence needed for backreaction
  roundtrip : free scalar field on a fixed globally hyperbolic background reproduces ordinary QFT
  breaks    : backreaction; needs the source rule (BHC P4; ALG expansion in G_N)

define I_SD := spacetime <-> dynamics
  preserved : covariance; evolution as gauge, via the hypersurface-deformation algebra of the constraints
  lost      : a privileged time
  roundtrip : classical GR in ADM form reproduces the Einstein equations on a chosen foliation
  breaks    : quantum time; the foliation or clock choice (R2)

define I_MD := matter <-> dynamics
  preserved : probabilities and symmetries (Stone: self-adjoint generator <-> unitary flow)
  lost      : anything not generated by a self-adjoint operator
  roundtrip : harmonic oscillator, Hamiltonian <-> unitary group <-> spectrum
  breaks    : no self-adjoint generator (Wheeler-DeWitt)

assume G1 : every quantity in a pilot declares its space or algebra, units and dimension
assume G2 : overlap limits must agree: G -> 0 gives flat-space QFT; hbar -> 0 gives GR on the same matter
assume G3 : each interface runs its roundtrip case before the pilots that use it
? G4 : is there a single substrate (states on algebras over a causal category, flows as automorphisms,
       probabilities as measures) in which all three interfaces are functors beyond the perturbative regime?

claim I_note : the substrate above is a candidate; it is established in the perturbative and
               curved-background regime [RECALL]; no accepted nonperturbative gravity version is known
```

---

## 2. Branch C: BHC core

```
given M, g   : manifold with Lorentzian metric, a classical field with no wave function
given Q      : actual matter configuration in configuration space CQ
given Psi    : wave function on CQ, defined relative to a foliation F
given hbar, G

assume P1 (ontology)     : matter has a definite actual configuration Q(t); Psi is not the whole story
assume P2 (dynamics)     : Psi obeys a linear wave equation on (M, g); Q obeys a guidance equation from Psi;
                           the Psi-density is equivariant under the flow
assume P2r (declared)    : the evolution of Psi is reversible (bijective on states); not derived from Wigner
assume P3 (gravity)      : g is classical, with no wave function and no superposition
assume P4 (source rule)  : the source of g is S_mu_nu[Q, Psi], the stress-energy of the ACTUAL configuration,
                           including any guidance contribution; NOT the expectation <T_mu_nu>
                           field equation: G_mu_nu + Lambda g_mu_nu = 8 pi G S_mu_nu
assume P5 (clock/foliation) : equilibrium statistics of records are independent of the foliation F and of the
                           clock subsystem within a declared class K; trajectories may depend on F

define CONSISTENT := the divergence of S_mu_nu with respect to g vanishes on solutions of P2

? C1 : is CONSISTENT true for general (Psi, Q, g), or only on a state-dependent subclass?
? C2 : which class K of clocks makes P5 true?
? C3 : is the pair (P3, P4) compatible with an island-type Page curve?
? C4 : is BHC LOCC-type, and does it entangle two masses? (distinct from mean-field sourcing)
```

Frame register (which framing assumptions BHC keeps):

```
claim BHC_F1 : keeps a primitive state (Psi) and adds a primitive configuration (Q)
claim BHC_F2 : keeps time as a parameter via a foliation; P5 is the relaxation of it
claim BHC_F3 : keeps a fixed stage at the level of g classical
claim BHC_F4 : retains probability as typicality weights on Q, not as primitive
claim BHC_F5 : keeps unitarity as a declared premise (P2r)
claim BHC_F6 : keeps the continuum
claim BHC_profile : conservative: it discards none of the six; its novelty is P1 and P4
```

---

## 3. Branch D: ALG core

```
given A_loc     : the local QFT algebra of an observer's causal region, assumed type III
given clock     : an observer clock with energy bounded below
given Gdress    : the dressing group of gravitational constraints
given G_N       : Newton's constant; the construction is an expansion in G_N -> 0

assume Q1 (observables)  : invariant observables are the crossed product of A_loc with Gdress including the
                           observer clock; local observables without a dressing are not invariant
assume Q2 (entropy)      : in the crossed product the algebra is type II, so it carries a trace; entropy is
                           defined by the trace, and the generalized-entropy formula is recovered up to a
                           constant, to leading order in G_N
assume Q3 (probability)  : states are normalized positive functionals; Gleason-type uniqueness holds for von
                           Neumann algebras with no I_2 summand; type II and III have none
assume Q4 (time)         : evolution is conditioning on the internal clock (Page-Wootters structure in
                           operator-algebra form)
assume Q5 (metric)       : the classical metric labels a semiclassical sector in the G_N expansion; gravity
                           is quantum perturbatively (this DENIES BHC P3)

define INPUT_ALG := [QFT algebras exist (AQFT axioms hold); G_N expansion converges to the needed order;
                     clock energy is bounded below; an observer exists]

claim D_note1 : the maximum-entropy state in the de Sitter case is tracial; the modular flow of a tracial
                state is trivial; so no arrow of time can be drawn from that state's modular flow
? D1 : does the type II conclusion depend on how the clock's energy bound is imposed?
? D2 : do different clock choices give non-isomorphic invariant algebras?
? D3 : on a type II or type III algebra there are no pure states or minimal projections; where, if anywhere,
       can a Bohmian beable (BHC P1) attach?
? D4 : which results transfer between de Sitter and AdS settings?
```

Frame register for ALG:

```
claim ALG_F1 : discards a primitive state; the algebra and its states are primitive
claim ALG_F2 : discards time as a parameter; the clock is internal
claim ALG_F3 : discards the fixed stage for local observables (dressing); keeps a background expansion
claim ALG_F4 : keeps probability from the Gleason-type structure
claim ALG_F5 : keeps unitarity via the algebra automorphisms
claim ALG_F6 : keeps the continuum
claim ALG_profile : discards three of six; silent on outcomes
```

---

## 4. Coverage matrix: five problems, two branches

Per pilot: one shared question, branch-specific rung rules, and the forecast terminal state.
Forecasts are not commitments.

### R1. Semiclassical gravity (pilot S)

```
given R1_question : a state-dependent criterion for when quantum matter can consistently source classical spacetime

BHC:
  if CONSISTENT holds on a nontrivial class and a criterion C(Psi) is explicit and checkable
    then DERIVES(R1)
  else if it holds only with one AUX then EXPLAINS(R1)
  else if no source in the P4 family has zero divergence on any nontrivial class then REFUTES(P4)
  else BLOCKED(named obstruction)
  forecast: EXPLAINS; silence class: RELOCATED (to the form of S_mu_nu)

ALG:
  if the validity criterion follows from the G_N expansion plus a semiclassical clock
    then EXPLAINS(R1)
  else BLOCKED(named obstruction)
  forecast: EXPLAINS; silence class: RELOCATED (to convergence of the expansion)

R1_extra (gravitationally-induced-entanglement identifiability, shared):
define GIE_models := [perturbative quantum gravity; mean-field classical-metric with QFT matter
                      (Aziz-Howl; status: contested); Diosi-Penrose; stochastic or hybrid classical gravity;
                      BHC actual-configuration sourcing]
claim GIE_status : the Aziz-Howl claim is contested in the literature [SEARCHED]; rebuttals include an exact
                   Heisenberg-picture argument, a tunneling-channel argument with binding-potential
                   suppression, and an argument that any entanglement would be matter-mediated;
                   do not use Aziz-Howl as a premise (PR7)
claim GIE_gap  : no source covers BHC actual-configuration sourcing; BHC is a distinct row
? R1_q1 : in a two-branch toy, what does the empty branch contribute to S_mu_nu?
? R1_q2 : does each model entangle two masses, and with what dependence on M, R, d, dx, t and composition?
? R1_q3 : composite matter: does a realistic constituent description change a model's prediction?
          (novelty unverified; needs prior-art review)
```

### R2. Problem of time (pilot T)

```
given R2_question : for a constraint C = H_clock + H_sys, is the map from the physical inner product to the
                    clock-conditioned Hilbert space unitary, and do clock changes preserve it?
define R2_cases   := [ideal clock, spectrum = R; semi-bounded or POVM clock]

Both branches share the test; BHC reads the result as the verdict on P5, ALG as the verdict on Q4 and D2.

if unitary equivalence is derived for ideal clocks AND a deviation bound for semi-bounded clocks
  then DERIVES(R2)
else if only for ideal clocks then EXPLAINS(R2) with K = ideal clocks
else if the clock must be chosen case by case then ACCOMMODATES(R2)
else if an equilibrium-level prediction depends on the clock inside K then REFUTES(P5) [BHC] or REFUTES(Q4) [ALG]
else if the constraint operator is not essentially self-adjoint in the chosen toy then BLOCKED(domain)
else BLOCKED(named obstruction)

claim R2_target : BHC: EXPLAINS (ideal clocks); ALG: EXPLAINS to DERIVES (semi-bounded clock is built in)
claim R2_class  : HALF_RELOCATED: the role confusion (generator vs constraint) dissolves; the clock or
                  foliation choice remains
? R2_q1 : are Bohmian trajectories clock-covariant, with only the statistics invariant?
? R2_q2 : for a real wave function the Bohmian velocity vanishes; what is the timeless guidance on a
          constraint surface? [RECALL; verify]
```

### R3. Quantum measurement (pilot M)

```
given R3_question : reconstruct the measurement problem independently; compare Bohmian, collapse and Everett
define R3_triad := {A: linear unitary evolution; B: definite outcomes; C: Psi is complete}
derive R3_trilemma from [R3_triad] : not all three hold for general measurement interactions
                    result: at least one of A, B, C must be denied
define R3_routes := [Bohmian denies C; collapse denies A; Everett denies B as fundamental]

define Born_routes := [Gleason; Bohmian equivariance plus typicality; Deutsch-Wallace; envariance]
claim Born_routes_note : the last two are contested [RECALL]; record each route's assumption set and the
                    assumption common to all routes; that common assumption marks the wall

define Gleason_scope := [dim >= 3 (effect-level extensions for dim 2); countable additivity; a measure on
                    ALL projections; noncontextuality]
define Relocation_inventory := [noncontextuality; sigma-additivity; the dim-2 exception; outcome definiteness]

claim R3_two_walls : weights: Gleason bears on the form of the rule; outcomes: no measure on any lattice
                     touches them
claim R3_Bohm_note : Bohmian values are contextual; the outcome statistics are Born-form under equivariance,
                     so Gleason and Kochen-Specker do not exclude P1

BHC:
  if Born statistics follow from equivariance plus typicality with the typicality measure as the only AUX
    then EXPLAINS(R3 weights)
  and if outcomes follow from P1 (the conditional wave function gives effective collapse)
    then EXPLAINS(R3 outcomes) with P1 as the single added postulate
  else if the statistics are inserted as a postulate then POSTULATES(R3)
  else if a nonequilibrium audit (rho != |Psi|^2) shows a violation outside the relocated assumptions then REFUTES(P2)
  else BLOCKED(foliation)
  forecast: EXPLAINS (weights), EXPLAINS via P1 (outcomes); class: RELOCATED (weights), GENUINE (outcomes, as P1)

ALG:
  if Q3 gives the Born form on the invariant algebra then EXPLAINS(R3 weights)
  outcomes: BLOCKED unless an ontology is added; if D3 has no resolution then BLOCKED(no pure states)
  forecast: EXPLAINS (weights), GENUINE silence (outcomes)

? R3_q1 : does the Bohmian foliation requirement survive once R2 shows clock-independence of statistics?
? R3_q2 : audit: for nonequilibrium Bohmian statistics, which of the Gleason_scope assumptions is violated?
```

### R4. Black-hole information (pilot I)

```
given R4_setting : a black hole coupled to a bath in a solvable model (JT gravity plus a CFT is the candidate)
define R4_steps  := [S1 reproduce the island Page curve and Page time by generalized-entropy extremization;
                     S2 list the steps that use fluctuating geometry (replica wormholes);
                     S3 state the missing microscopic reconstruction; do not solve it;
                     S4 test whether the branch core gives a Page curve by another route]
claim R4_scope : the island derivations are physics-level controlled calculations in AdS-type or
                 flat-with-bath setups [RECALL]; transfer to de Sitter is unestablished (D4)
claim R4_premise : unitarity enters as a declared premise (P2r), not as a derived result

BHC:
  if S4 gives a Page curve from the core with zero AUX then DERIVES(R4)
  else if only with the quantum-extremal-surface formula as AUX then POSTULATES(R4)
  else if a fluctuating-geometry sum is indispensable and the core cannot supply it then REFUTES(P3)
  else BLOCKED(named obstruction)
  forecast: POSTULATES or REFUTES(P3); a REFUTES(P3) is recorded as evidence against the classical-gravity branch

ALG:
  if the generalized-entropy formula follows from Q2 and the island structure follows at leading order
    then EXPLAINS(R4 entropy)
  Page curve: POSTULATES or BLOCKED unless the replica-wormhole structure follows from the crossed product
  forecast: EXPLAINS (entropy), POSTULATES or BLOCKED (Page curve)
```

### R5. Cosmological constant (pilot L)

```
given R5_question : test the equations of ONE adjustment proposal (default: unimodular gravity)
define R5_steps   := [write trace-free field equations; show Lambda is an integration constant;
                      derive that a constant vacuum energy density does not source curvature;
                      check what each branch's source rule does to this]
claim R5_note : classically, unimodular gravity is equivalent to GR with Lambda as an integration constant;
                it makes no distinct lensing prediction; quantum equivalence is open

BHC:
  if the constant vacuum energy drops out AND P4 respects it
    then EXPLAINS(R5 non-gravitation) and POSTULATES(R5 value)
  else if P4 reintroduces a gravitating trace then REFUTES(R5 default proposal) under the core
  else BLOCKED
ALG:
  Lambda sets the horizon scale and enters as an input; the de Sitter algebra's type does not determine it
  forecast: POSTULATES (value); EXPLAINS only the structure of the entropy bound

claim R5_class : GENUINE for the value of Lambda in both branches
claim R5_falsifier (conditional): if evolving dark energy is confirmed at the 5-sigma level, a pure
        integration-constant Lambda cannot be the full account and the default proposal is replaced
        [SEARCHED: current combinations give roughly 2.8 to 4.2 sigma and are dataset-dependent; not a detection]
? R5_q1 : does the actual-configuration vacuum field produce a fluctuating S_mu_nu whose trace sources g?
? R5_q2 : swap to vacuum-energy sequestering if the default fails; log the swap per PR3
```

---

## 5. Cohesion audit and assumption graph (per branch, then compared)

```
define AUX_TOTAL     := distinct AUX across R1..R5 for the branch
define AUX_SHARED(a) := a is used in at least 2 pilots
define UNDECLARED_N  := problems with no terminal state
define STRONG_N      := pilots ending in DERIVES or EXPLAINS using only the core

if UNDECLARED_N = 0 and STRONG_N >= 2 and AUX_TOTAL <= 3 and every private-use AUX is logged with its cost
   and no REFUTES targets a core postulate unrecorded
then AUDIT = COHESIVE_SO_FAR
else if UNDECLARED_N > 0 then AUDIT = DEFECT
else AUDIT = NOT_COHESIVE

? AUDIT_thresholds : STRONG_N >= 2 and AUX_TOTAL <= 3 are draft values [RATIFY]

define ASSUMPTION_GRAPH := nodes = assumptions used in any derive; edges = co-occurrence in one derive;
                           an assumption load-bearing in two or more pilots is a shared AUX
define FOOTPRINT(pilot) := the minimal set of structures and assumptions the terminal state used
define BRANCH_COMPARE   := compare footprints and silence classes; a branch is not "ahead" if it covers a
                           different set of silences; report both profiles
```

Interpretation guard: COHESIVE_SO_FAR is not "the theory is true". A REFUTES(P3) or REFUTES(P4) is a valid
result: it falsifies that form of the branch.

---

## 6. Per-pilot close-out (MRC debt list)

```
? needMap               : which branch of mathematics applies, and where it does not
? needInvariant         : what stays fixed under the relevant change (clock, foliation, branch)
? needToyCheck          : a solvable model reproducing the claim
? needNullModel         : what a structure-free alternative predicts for the same observable
? needObstruction       : the strongest known obstruction, stated, not dismissed
? needFaithfulnessReview: adversarial pass for overclaiming; no final-truth language
? needSanityCheck       : PR6 one-liner passed (units, classical equivalence, triviality)
? needInterfaceRoundtrip: the section 1 roundtrip for each interface used
```

Pilot order: R3 + R1 together, then R2, then R4, then R5. R1 uses R3's ontology (BHC); R2 settles P5 (BHC) and
Q4/D2 (ALG); R4 is the most discriminating for P3; R5 depends on the source rule from R1.

---

## 7. Series-silence register

What the Math Reality Check series says about each problem, and what is external to it.

```
claim S_R1 : the series uses "semiclassical" almost only for the hbar -> 0 stationary-phase limit; it has no
             entry on quantum matter sourcing a classical metric; all content is external
claim S_R2 : one passing mention of Wheeler-DeWitt; no Page-Wootters or problem-of-time map; all external
claim S_R3 : about twenty "can X solve the measurement problem? NO" rows; no Gleason, Everett, collapse
             models or contextuality; one passing Bohmian line; the positive content is external
claim S_R4 : the best covered: Harlow-Hayden, islands and replica wormholes (as topology sectors) appear,
             verdict partial
claim S_R5 : every map has a "can X fix Lambda? NO" row; no specific adjustment proposal is analysed
             (unimodular gravity does not appear)
claim S_all: the series is a record of what structures cannot do; positive content for these pilots is external
```

---

## 8. Source-intake register

```
[SEARCHED] Aziz-Howl, Nature 646, 813 (2025): contested; rebuttals: arXiv 2511.00852, 2511.07348, 2512.13675
[SEARCHED] DESI combinations ~2.8 to 4.2 sigma preference for evolving dark energy; not 5 sigma; dataset-dependent
[UNVERIFIED, from an attached model response] Trillo-Navascues 2025; Di Biagio Oct 2026; Lin-Mondal 2026;
           Howl-Cooper-Hackermuller Aug 2026; Miki-Li-Chen Sep 2026
[RECALL] Gleason (1957) and effect/algebra extensions (Busch; Christensen, Yeadon)
[RECALL] Chandrasekaran-Longo-Penington-Witten (2022); Witten crossed-product (2021); Jensen-Sorce-Speranza
[RECALL] Hoehn-Smith-Lock on Dirac, reduced and Page-Wootters equivalence
[RECALL] Goldstein-Teufel; Duerr-Goldstein-Zanghi typicality; Duerr et al. on relativistic Bohmian mechanics
[RECALL] Almheiri-Engelhardt-Marolf-Maxfield; Penington; replica wormholes; Harlow-Hayden
```

---

## 9. Known weak points

```
? W1 : BHC P4 may be divergence-free only on a narrow class; R1 tests this first
? W2 : BHC P5 asserts what R2 must prove; if R2 fails, P5 and the relativistic Bohmian extension are redone
? W3 : ALG takes QFT algebras as given; the 4D QFT existence gap is inherited
? W4 : ALG is a G_N expansion; every result carries that restriction
? W5 : the fork: a Bohmian ontology has no known place on a type II or III algebra (D3)
? W6 : the rung ladder is a classification, not a computed status
? W7 : G4 (a single nonperturbative substrate) is open; the interface contracts are the substitute
```
