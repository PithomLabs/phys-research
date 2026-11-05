I would approve Plan 3 with six corrections. The reconnaissance is good, the selected paper is suitable, and the implementation is now specific enough to begin.

There are a few scientific and architectural details I would correct before authorizing implementation.
1. Decisions: approve these
Agent's question	My decision
MVP 1 paper	Approve  MICROSCOPE 2022
Seventh metadata field	Approve  class: physical\|mathematical, required only for hypotheses
Go module path	github.com/PithomLabs/phys-research
Priority enum	high, medium, low
Human setup	Approve separate working and accepted branches, with stronger acceptance protection described below

Why MICROSCOPE?

It is an excellent first workflow test because the theoretical null prediction and experimental measurement are clearly distinguishable, the result is quantitative, and the paper exposes real questions about systematic uncertainties and data processing.

The published result is:
\[
\eta(\mathrm{Ti},\mathrm{Pt})
=
[-1.5\pm2.3_{\rm stat}\pm1.5_{\rm syst}]
\times10^{-15}.
\]

The experiment used titanium and platinum alloy test masses in low-Earth orbit. 

Phys. Rev. Lett.+1

It also gives the Adversary real material to examine: thermal systematics, glitch removal, M-ECM data reconstruction and the approximation identifying the measured \(\delta\) with the Eötvös parameter \(\eta\).
2. Six corrections before implementation
Correction 1 — Remove a circular assumption.

The proposed hypothesis says:
assume wep: universality of free fall holds...


That assumes the very principle the experiment is testing.

Replace it with:
claim wep_prediction: under WEP, eta(Ti,Pt) = 0


Theoretical expectation is not an assumption that the experiment's conclusion is true.

Correction 2 — Narrow the long-range force claim.

This claim is too broad:
claim bound: any long-range mediator producing
a Ti/Pt differential acceleration must yield
|eta| below the reported sensitivity


MICROSCOPE constrains the differential acceleration of a particular composition pair in a specified experimental regime. Translating that result into mediator parameters requires a particular interaction model, composition-dependent couplings, source and interaction range.

The paper does discuss long-range interactions, but does not establish a universal bound on every possible mediator. 

La physique autrement+1

Keep model-specific force constraints as a subsequent research issue.

Correction 3 — Fix the numerical falsifier.

The proposed 4.5 × 10⁻¹⁵ as the combined 1σ envelope is incorrect.

The companion MICROSCOPE analysis reports approximately \(2.7\times10^{-15}\) when statistical and systematic uncertainties are combined quadratically. 

TU Delft Research Portal

Retain the original statistical and systematic uncertainties separately. Define any 3σ falsification threshold using a declared total-uncertainty prescription, rather than hardcoding 4.5.

A reproducible, independently validated nonzero result exceeding 3σ total significance is a reasonable initial contradiction criterion, not automatic proof of WEP violation.

Correction 4 — Do not manufacture an independent theoretical result.

The proposed issue contains four expectation/actual sections. Keep them, but allow:

NOT APPLICABLE — no independent theoretical derivation is required for this experimental result.

For MVP 1, the independent theoretical result is not a missing derivation that the agent needs to invent.

Correction 5 — Scope graph-cycle validation.

Only check cycles in the depends_on relationship graph.

A paper challenging another paper while being challenged in return is a perfectly legitimate scientific relationship. Rejecting every cycle across all relationships would corrupt the research graph.

Correction 6 — Strengthen human acceptance.

I approve main as the working branch and accepted as the human-controlled branch.

However, this proposed hook:
PHYS_HUMAN_MERGE=1


is not a genuine security boundary. An agent with unrestricted shell access can set that environment variable.

Use remote branch protection and separate human-controlled credentials. Agents may commit to the working branch but must not possess credentials capable of updating accepted. If there is no protected remote, perform the acceptance merge from a separate human-controlled clone or environment.

Also, keep go.work local and untracked unless the relative workspace layout is deliberately standardized for every checkout. The verified local test setup should not become an accidental machine-specific repository dependency.
3. Revised MVP 1 hypothesis
I would use this as the starting artifact:
# H-001 — MICROSCOPE WEP null result

given eta_definition:
  eta(A,B) = 2(a_A - a_B)/(a_A + a_B)

claim wep_prediction:
  under WEP, eta(Ti,Pt) = 0

given reported_measurement:
  eta(Ti,Pt) =
  [-1.5 +/- 2.3(stat) +/- 1.5(syst)] x 10^-15

given experimental_regime:
  MICROSCOPE SUEP sensor, TiAlV/PtRh test masses,
  low-Earth orbit, final 2022 mission analysis

claim reported_result:
  the reported measurement is consistent with
  the WEP null prediction within the stated uncertainties

assume reported_systematics:
  published thermal and instrumental systematic
  uncertainty estimates are applicable

? whether the extrapolation of the thermal attenuation
  factor to the EP measurement frequency is sufficiently
  justified by the calibration data

? whether glitch masking and M-ECM reconstruction
  preserve sensitivity to a genuine WEP-violating signal


The falsification condition belongs in the accompanying research-record prose, not as an invented Phys-Pseudo construct.

The Adversary should independently examine the two unresolved questions, reproduce the uncertainty interpretation, and check that the hypothesis does not claim more than this experiment establishes.

The original paper identifies the thermal attenuation extrapolation and glitch handling explicitly. 

La physique autrement+1

4. Exact authorization to the implementation agent
Approved. Proceed with Plan 3, incorporating these amendments:
1. Paper: approve MICROSCOPE 2022, DOI 10.1103/PhysRevLett.129.121102, arXiv:2209.15487v1.
2. Metadata: add class: physical|mathematical, required only for hypothesis records. A physical hypothesis requires an explicit falsification condition. A mathematical hypothesis requires premises, conclusion, domain and proof obligations.
3. Module: github.com/PithomLabs/phys-research.
4. Priority: high|medium|low.
5. Git: main is the working branch; accepted is protected and human-controlled. Do not treat a user-set environment variable or local Git hook as a security boundary. Keep agent credentials separate from human acceptance credentials. Agents must not be able to update the accepted remote branch.
Additional corrections:
6. Remove the circular assume wep statement. WEP is the theoretical null prediction, not an assumed experimental result.
7. Narrow the long-range mediator claim. General force constraints require a specified interaction model, coupling, range and source.
8. Remove the incorrect 4.5 × 10^-15 combined 1σ falsification threshold. Preserve statistical and systematic uncertainty separately. Any 3σ criterion must use an explicitly declared total-uncertainty prescription.
9. Permit NOT APPLICABLE for independent theoretical derivation when the issue is purely experimental.
10. Restrict cycle rejection to depends_on relationships, not the entire research relationship graph.
11. Reviews must record the exact hypothesis digest and become stale when their target changes.
12. Make go.work developer-local and untracked unless the relative workspace layout is intentionally standardized.
13. Preserve the seven Phys-Pseudo constructs exactly. No invented falsification syntax, no general-purpose YAML dependency, and no claim that structural validation establishes scientific correctness.
Proceed with MVP 1 only. Do not begin Branch C mathematical implementation or expand the scope.

Run the full specified test suite and report actual results, limitations and human decisions required. Do not claim a feature or test is complete unless verified.




My final recommendation: authorize implementation. The most important thing is now to get one complete, trustworthy research record through the system. Resist adding features until that works.
