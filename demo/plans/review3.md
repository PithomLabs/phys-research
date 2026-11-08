The 24-page Quantum Stress Energy Conservation Research.pdf is Agent B's substantive QSG research result. It is important, and I would preserve it as the foundation of QSG-T001.

My assessment is more favorable than my assessment of the Bohmian report, but there are still mathematical corrections required before we promote its theorem.

Agent B — QSG research

Strong foundation

Research direction: Correct.

Classical derivation: Correct in its central conservation identity.

Quantum renormalization: Substantive, but the conservation proof needs correction.

Backreaction: Supported in a restricted FLRW setting.

Phys-Pseudo: Requires conformance correction.

Theorem readiness: Not yet accepted.

## 1. What Agent B got right

There are four meaningful contributions.

A. The classical foundation is properly reconstructed.

On pages 4–6, Agent B derives the scalar-field stress-energy tensor and obtains:

\\[ \nabla^\mu T\_{\mu\nu} = (\Box_g\phi-m^2\phi)\nabla\_\nu\phi. \\]

It correctly distinguishes the off-shell identity from on-shell conservation.

This is an important correction to the earlier Bohmian report, which confused the two.&#x20;

Quantum Stress Energy Conservation Research.pdf



B. The Hadamard framework is appropriately selected.

Pages 7–11 develop the quantum field, two-point function, Hadamard condition, point-splitting prescription and finite curvature ambiguities.

This is the right mathematical foundation for our initial problem.&#x20;

Quantum Stress Energy Conservation Research.pdf

Quantum Stress Energy Conservation Research.pdf



C. It correctly recognizes that conservation and the trace anomaly are different issues.

This is important.

A renormalized stress-energy tensor can have a nonzero trace while remaining covariantly conserved.

Agent B explicitly investigates this on page 13.&#x20;

Quantum Stress Energy Conservation Research.pdf



D. It identifies an existing existence-and-uniqueness result.

Meda, Pinamonti and Siemssen (2021) prove local existence and uniqueness results for semiclassical Einstein equations in spatially flat cosmological models with a massive quantum scalar field.

This gives us an established mathematical benchmark for a later backreaction theorem.&#x20;

[image](https://www.google.com/s2/favicons?domain=https://link.springer.com\&sz=32)

Springer Nature Link+1



Their result is restricted: it uses selected compatible initial states and geometric initial conditions, and the demonstrated solution is initially mild and local in time. It is not a general existence theorem for arbitrary curved spacetimes.&#x20;

[image](https://www.google.com/s2/favicons?domain=https://link.springer.com\&sz=32)

Springer Nature Link+1



## 2. The most important defect: the conservation derivation

Requires correction

This is the main reason I would not yet accept QSG-T001.

On pages 12–13, Agent B constructs:

\\[ W(x,x')=G^+(x,x')-H(x,x') \\]

where \\(G^+\\) is the state two-point function and \\(H\\) is the Hadamard subtraction parametrix.

The report then claims that because both terms satisfy the Klein–Gordon equation up to smooth terms, \\(W\\) satisfies the homogeneous equation identically.

That step is not justified.

The Hadamard parametrix is generally a local approximate solution, not an exact bisolution of the Klein–Gordon operator.

Subtracting it leaves a smooth remainder, but that remainder need not satisfy the homogeneous equation.

This matters because its smooth differential remainder contributes to the divergence of the point-split stress-energy expression.

There is a particularly relevant historical correction.

In his 1978 paper, Robert Wald explicitly identified an error in an earlier point-separation conservation argument. He showed that the naive prescription produces a divergence equal to the gradient of a local curvature term, which must be corrected. That correction also has consequences for the trace anomaly.&#x20;

[image](https://www.google.com/s2/favicons?domain=https://journals.aps.org\&sz=32)

Phys. Rev. D



Décanini and Folacci provide a useful explicit treatment. Their renormalized stress-energy expression includes a state-independent geometric correction tensor \\(\widetilde{\Theta}\_{\mu\nu}\\), whose divergence is chosen to cancel the residual from point splitting. See their equations (64), (68) and (69).&#x20;

[image](https://www.google.com/s2/favicons?domain=https://arxiv.org\&sz=32)

arXiv



Agent B is approaching the right mathematical construction, but its Step 4 and Step 5 must be reconciled.

There is also an internal wording inconsistency: page 13 describes its correction tensor as conserved while simultaneously using its nonzero divergence to cancel the residual. Those cannot both be true of the same tensor.&#x20;

Quantum Stress Energy Conservation Research.pdf



This is exactly the kind of defect our MRC and independent derivation process should catch.

## 3. Other corrections required

| Location    | Issue                                                               | Required correction                                                                                                 |
| ----------- | ------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------- |
| Page 11     | Wald's axioms summarized as four conditions                         | Reconcile with Wald's 1978 discussion of the fifth condition and trace anomaly.                                     |
| Page 11     | Renormalization ambiguity coefficients called dimensionless         | Restore the appropriate mass or reference scales for dimensional consistency.                                       |
| Page 15     | Backreaction existence and order reduction treated too broadly      | State the exact FLRW theorem assumptions and distinguish established results from additional approximation methods. |
| Page 16     | Noise-kernel positivity described as an ordinary pointwise property | Specify its distributional and smeared interpretation.                                                              |
| Page 19     | Variance divided by expectation                                     | Dimensionally invalid as written; the intended relative standard deviation must be specified.                       |
| Pages 19–20 | Phys-Pseudo specification                                           | Rewrite using the authoritative seven-construct grammar.                                                            |

The fluctuation discussion is useful, particularly its recognition that a relative variance criterion becomes problematic when the mean is zero and that macroscopic superpositions require special treatment. But Kuo and Ford's criterion is not a universal theorem declaring a threshold at which semiclassical gravity necessarily transitions into stochastic gravity.&#x20;

[image](https://www.google.com/s2/favicons?domain=https://journals.aps.org\&sz=32)

Phys. Rev. D

Quantum Stress Energy Conservation Research.pdf



One easy bibliography correction: the report lists Wald's 1977 Communications in Mathematical Physics paper as volume 54, page 11. The correct reference is volume 54, pages 1–19.&#x20;

[image](https://www.google.com/s2/favicons?domain=https://www.scienceopen.com\&sz=32)

ScienceOpen



## 4. What I would separate into theorem targets

I would not try to make this entire 24-page report into one theorem.

Instead:

QSG-T001 · First theorem

Conservation of Hadamard-renormalized stress-energy

Prove, under specified assumptions and with the correct local geometric correction:

\\[ \nabla^\mu \langle\hat T\_{\mu\nu}\rangle\_{\rm ren}=0. \\]

Scope: free, minimally coupled, massive real scalar field on a smooth globally hyperbolic four-dimensional spacetime.

QSG-T002 · Later theorem

Local existence of semiclassical backreaction solutions

Reconstruct the restricted FLRW existence-and-uniqueness result of Meda, Pinamonti and Siemssen, and distinguish it from any new generalization.

QSG-T003 · Research hypothesis

State-dependent semiclassical validity

Investigate whether smeared stress-energy correlations and their gravitational response provide a defensible validity criterion.

Do not declare a universal threshold until it is derived.

The first theorem is primarily a reconstruction of an established mathematical result, not a claim of novelty. That is acceptable: it exercises our complete Phys derivation workflow against a serious mathematical target.

## 5. What I recommend doing now

I would not immediately send the PDF to the coding implementation agent.

Instead, commission Agent C — Independent Conservation Audit.

Its assignment should be extremely narrow:

1. Independently derive the classical stress-energy divergence identity.
2. Reconstruct the four-dimensional Hadamard point-splitting prescription.
3. Independently derive the local conservation-restoring geometric term.
4. Compare its result with Wald (1978) and Décanini–Folacci.
5. Identify every missing premise or invalid step in Agent B's pages 12–13.
6. Establish the exact theorem statement, including all assumptions and renormalization ambiguities.

Agent C should receive the original PDF and primary sources, but should not be instructed to accept Agent B's claimed DERIVED status.

Once that review is complete, Agent B can revise the derivation.

Only then should we prepare the RED-phase theorem contract and implement QSG-T001 using the GR reference structure.

## 6. My final verdict

The PDF confirms that our pivot to conventional QFT on classical curved spacetime was productive.

We now have:

- A correct classical mathematical foundation.
- An appropriate quantum-field framework.
- A concrete renormalization prescription to investigate.
- An established conservation theorem to reconstruct.
- A restricted backreaction existence result for a later benchmark.
- A promising direction for studying semiclassical validity.

But the central conservation proof still needs independent reconstruction. The report's current declaration that QSG-T001 is DERIVED is premature.&#x20;

Quantum Stress Energy Conservation Research.pdf

Quantum Stress Energy Conservation Research.pdf



My decision: retain Agent B's PDF as the QSG research baseline, send its conservation derivation for independent audit, and do not change our infrastructure.

We are now one properly verified mathematical result away from our first QSG theorem formalization.
