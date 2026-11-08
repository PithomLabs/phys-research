# BHC core + MRC coverage matrix (phys-pseudo draft)

Status: PROPOSED. Nothing below is asserted true. Representation is not endorsement.
Syntax: the seven constructs only (given, define, assume, claim, derive, if/then/else, ?).
The surface form is provisional; conform it to PSEUDO.md before agents consume it.
Candidate name BHC = Bohmian matter + classical (unquantized) gravity, Branch C.
Draft-chosen numeric thresholds are marked [RATIFY]; they are my choices, not derived.

---

## 0. Pre-registration rules (fixed BEFORE any pilot computation)

```
assume PR1: the core is P1..P5 in section 1 and nothing else
assume PR2: any further premise a pilot needs is an AUX; it is declared at the point of use,
            tagged with the pilot id, and counted in section 4
assume PR3: rung targets and failure conditions in section 3 are not edited after a pilot starts;
            a change is logged as a new version with the reason
assume PR4: a pilot ends in exactly one terminal state:
            DERIVES | EXPLAINS | ACCOMMODATES | POSTULATES | REFUTES | BLOCKED
            UNDECLARED is not a terminal state; it is a defect that fails the audit
assume PR5: REFUTES means a core postulate is contradicted; it names the postulate.
            BLOCKED means the math was not completed; it names the obstruction.
            Neither deletes the candidate.
```

Rung definitions, shared by all pilots:

```
define DERIVES(X)      := a derive of X from the core alone, zero AUX, passing the MRC debt list (section 5)
define EXPLAINS(X)     := X follows from the core plus at most one AUX, where the AUX is used in at least
                          one other pilot or is a standard piece of the cited framework; OR the core fixes
                          the qualitative structure of X but not its numerical value
define ACCOMMODATES(X) := X is reproduced only after choosing a free element the core does not fix
define POSTULATES(X)   := X is added as an assumption
define REFUTES(P)      := a derive, counterexample or toy check contradicts core postulate P
define BLOCKED(why)    := the required derivation is not completed; why names the specific obstruction
```

---

## 1. Shared core

```
given M, g        : smooth manifold M with Lorentzian metric g, a classical field (no wave function)
given Q           : actual matter configuration (fields or particles) in configuration space CQ
given Psi         : wave function on CQ, defined relative to a foliation F of (M, g)
given hbar, G     : constants

assume P1 (ontology)     : matter has a definite actual configuration Q(t); Psi is not the whole story
assume P2 (dynamics)     : Psi obeys a linear wave equation on (M, g); Q obeys a guidance equation
                           determined by Psi; the Psi-density on CQ is equivariant under the flow
assume P3 (gravity)      : g is a classical field with no wave function and no superposition
assume P4 (source rule)  : the source of g is S_mu_nu[Q, Psi], the stress-energy of the ACTUAL
                           configuration including any guidance/quantum-potential contribution;
                           it is not the expectation value <T_mu_nu>
                           field equation: G_mu_nu + Lambda g_mu_nu = 8 pi G S_mu_nu
assume P5 (clock/foliation) : equilibrium statistics of records are independent of the choice of
                           foliation F and of clock subsystem within a declared class K;
                           non-equilibrium or single-trajectory facts may depend on F

define CONSISTENT := the divergence of S_mu_nu with respect to g is zero on the solutions of P2

? C1 : is CONSISTENT true for general (Psi, Q, g), or only on a state-dependent subclass?
? C2 : which class K of clocks makes P5 true?
? C3 : is the pair (P3, P4) compatible with a quantum-extremal-surface type Page curve?
```

Unresolved items C1 to C3 are the three places the core can fail. Pilots R1, R2 and R4 are
built to test them directly.

---

## 2. Linkage claims (cohesion structure, stated so they can be audited)

```
claim L1 : R3 supplies the ontology that R1 uses to define S_mu_nu
claim L2 : R2 supplies the verdict on P5, which constrains the foliation freedom used in R3
claim L3 : R4 is a direct test of P3 and P4 (classical g with a single-configuration source)
claim L4 : R1 supplies the sourcing rule that decides what gravitates in R5
claim L5 : R3 and R2 together decide whether F is physical or gauge, which R4 inherits
```

---

## 3. Coverage matrix: five problems x terminal state

Columns are the five outcomes each pilot may reach. Cells give the pre-registered condition.
"Target" is the rung I expect it can honestly reach; it is a forecast, not a commitment.

### R1. Semiclassical gravity (pilot S)

```
given R1_question : a state-dependent criterion for when quantum matter can consistently source classical spacetime
define R1_artifacts := [explicit S_mu_nu[Q, Psi]; derive of divergence condition; a criterion C(Psi);
                        two-branch toy where S differs from <T>]

if derive gives CONSISTENT for a nontrivial class of (Psi, Q) and C(Psi) is explicit and checkable
then DERIVES(R1)
else if CONSISTENT holds only with one AUX (e.g. an approximation regime) then EXPLAINS(R1)
else if CONSISTENT holds only for a hand-picked class then ACCOMMODATES(R1)
else if no S_mu_nu in the P4 family has zero divergence on any nontrivial class then REFUTES(P4)
else BLOCKED(named obstruction)

claim R1_target : EXPLAINS or DERIVES
? R1_q1 : in the two-branch toy, what does the empty branch contribute to S_mu_nu through the guidance term?
? R1_q2 : does the P3/P4 pair forbid gravity-mediated entanglement, and what does the Aziz-Howl setup
          then predict? (Candidate A link; do not assert before a toy check)
```

### R2. Problem of time (pilot T)

```
given R2_question : for a constraint C = H_clock + H_sys, is the map from the physical inner product
                    (group-averaged / refined algebraic quantization) to the clock-conditioned
                    Hilbert space unitary, and do transition maps between clock choices preserve it?
define R2_cases := [ideal clock, spectrum = R; semi-bounded or POVM clock]

if unitary equivalence is derived for ideal clocks AND a bound on the deviation is derived for semi-bounded clocks
then DERIVES(R2)
else if equivalence is derived for ideal clocks only then EXPLAINS(R2) with K = ideal clocks
else if equivalence needs the clock to be picked case by case then ACCOMMODATES(R2)
else if some equilibrium-level prediction depends on clock choice inside K then REFUTES(P5)
else BLOCKED(named obstruction)

claim R2_target : EXPLAINS (ideal clocks), DERIVES only if the bound is obtained
? R2_q1 : are Bohmian trajectories clock-covariant, with only the statistics invariant, or does F select a clock?
```

### R3. Quantum measurement, Bohmian focus (pilot M)

```
given R3_question : reconstruct the measurement problem independently; compare Bohmian, collapse and Everett
define R3_triad := {A: Psi evolves linearly and unitarily; B: measurements have definite outcomes;
                    C: Psi is a complete description}
derive R3_trilemma from [R3_triad] : the three jointly cannot hold for general measurement interactions
                    result: at least one of A, B, C must be denied
define R3_routes := [Bohmian denies C; collapse denies A; Everett denies B as fundamental, outcomes emergent]

define R3_criteria := [outcome statistics (Born rule): derived or postulated;
                       empirical equivalence or difference;
                       relativistic extendibility;
                       compatibility with classical g and the P4 source rule]

if Born statistics follow from equivariance plus a typicality argument whose measure is the only AUX
then EXPLAINS(R3 for Bohmian)
else if the statistics are put in as a postulate then POSTULATES(R3 for Bohmian)
else if the relativistic or gravitational extension fails with no foliation-free fix then BLOCKED(foliation)

claim R3_target : EXPLAINS for Bohmian (typicality measure as the AUX);
                  the comparison table for collapse and Everett is a deliverable, not a rung
? R3_q1 : does the Bohmian foliation requirement survive once R2 shows clock-independence of statistics?
```

### R4. Black-hole information (pilot I)

```
given R4_setting  : an evaporating or eternal black hole coupled to a bath in a solvable model
                    (JT gravity plus a CFT is the controlled candidate)
define R4_steps   := [S1 reproduce the island Page curve and Page time via generalized-entropy extremization;
                      S2 list which steps use fluctuating geometry (replica wormholes);
                      S3 state the missing microscopic reconstruction, do not solve it;
                      S4 test whether P3 + P4 give a Page curve by another route]

if S4 yields a Page curve from the core with zero AUX then DERIVES(R4)
else if S4 yields one only with the quantum-extremal-surface formula as AUX then POSTULATES(R4) with the AUX named
else if S2 shows a fluctuating-geometry sum is indispensable and the core cannot supply it then REFUTES(P3)
else BLOCKED(named obstruction)

claim R4_target : S1 to S3 complete; terminal state most likely POSTULATES or REFUTES(P3)
claim R4_value  : a REFUTES(P3) here is evidence against the classical-gravity branch and is recorded as such
```

### R5. Cosmological constant (pilot L)

```
given R5_question : test the equations of ONE adjustment proposal (default: unimodular gravity)
define R5_steps   := [write trace-free field equations; show Lambda is an integration constant;
                      derive that a constant vacuum energy density does not source curvature;
                      check what the P4 actual-configuration source does to this statement]

if constant vacuum energy drops out AND the P4 source respects it
then EXPLAINS(R5 non-gravitation of constant vacuum energy) and POSTULATES(R5 value of Lambda)
else if P4 reintroduces a trace that gravitates then REFUTES(R5 default proposal) under the core
else BLOCKED(named obstruction)

claim R5_target : EXPLAINS the non-gravitation; the value of Lambda is expected to remain POSTULATES
? R5_q1 : does the actual-configuration vacuum field produce a fluctuating S_mu_nu whose trace sources g?
? R5_q2 : swap in vacuum-energy sequestering if unimodular fails; log the swap per PR3
```

---

## 4. Cohesion audit (the pass/fail rule for "we are on the right track")

```
define AUX_TOTAL      := count of distinct AUX across R1..R5
define AUX_SHARED(a)  := a is used in at least 2 pilots
define UNDECLARED_N   := number of problems with no terminal state
define STRONG_N       := number of pilots ending in DERIVES or EXPLAINS using only the core

if UNDECLARED_N = 0 and STRONG_N >= 2 and AUX_TOTAL <= 3 and every AUX with a private use is logged with its cost
   and no REFUTES targets P3 or P4 (or the REFUTES is recorded as the result)
then AUDIT = COHESIVE_SO_FAR
else if UNDECLARED_N > 0 then AUDIT = DEFECT
else AUDIT = NOT_COHESIVE

? AUDIT_thresholds : STRONG_N >= 2 and AUX_TOTAL <= 3 are draft values [RATIFY]
```

Interpretation guard: COHESIVE_SO_FAR is not "the theory is true". It means the core covered the
five problems without a separate patch per problem under the pre-registered rules.
A REFUTES(P3) or REFUTES(P4) is a legitimate and valuable outcome: it falsifies this form of the
candidate and is evidence on the classical-gravity question.

---

## 5. Per-pilot close-out (MRC debt list, retire before promotion)

Each pilot carries the EBP v2.1 debt list; promotion to candidate artifact requires each to be retired.

```
? needMap               : which branch of mathematics applies, and where it does not
? needInvariant         : what quantity stays fixed under the relevant change (clock, foliation, branch)
? needToyCheck          : a solvable model reproducing the claim
? needNullModel         : what a theory with no such structure predicts for the same observable
? needObstruction       : the strongest known obstruction, stated, not dismissed
? needFaithfulnessReview: adversarial pass for overclaiming; no final-truth language
```

Pilot order proposed: R3 + R1 together, then R2, then R4, then R5.
Reason: R1 depends on R3's ontology; R2 settles P5 for R3; R4 is the most discriminating; R5 depends on R1's source rule.

---

## 6. Known weak points of this draft

```
? W1 : P4 may be incompatible with divergence-free sourcing outside a narrow class; R1 tests this first
? W2 : P5 asserts what R2 must prove; if R2 fails, P5 and the relativistic Bohmian extension must be redone
? W3 : the literature named in the pilots (Page-Wootters equivalence results, island/replica-wormhole papers,
        Bohmian quantum-spacetime proposals, Aziz-Howl) is cited from recall; verify against primary text
        before any claim leaves the draft state
? W4 : the five-rung ladder is borrowed from the Workbench; it is a classification, not a computed status
```
