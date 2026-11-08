Adversarial review: Plan 5 — QSG-T001 resumption
Strong foundation
Amend before execution

Verdict: Approve the approach, but make five important corrections before Agent A executes it.

Plan 5 is substantially better than the previous workflow. It understands that the infrastructure work is finished, that the classical foundation has already been reconstructed, and that the remaining bottleneck is the Hadamard point-splitting conservation calculation.

Most importantly, it introduces a proper stop gate before implementation. 
plan5(9).md

However, I found two problems that could produce another round of wasted work, and three that could undermine the correctness or acceptance of the eventual result.
1. Findings by severity
Finding	Severity	Required action
Expected correction is disclosed before independent derivation	Critical	Introduce a blind mathematical derivation
Standalone Go fixture has no specified module/build boundary	Critical	Resolve packaging before implementation
Agent A and Agent C responsibilities are blurred	Major	Preserve independent review provenance
Two-point-function conventions are not sufficiently frozen	Major	Specify exact mathematical objects and normalizations
Validation and non-modification checks are incomplete	Major	Add hashes, explicit command results and acceptance boundaries

Let's examine each.
2. Critical finding: mathematical anchoring
Plan 5 says:
The previous audit identified a candidate correction proportional to \(g_{\mu\nu}v_1\).

It also identifies Décanini–Folacci's Equation (70) as the source covering the needed coefficient.

These are useful research leads, but they create a risk.

Agent A already knows what mathematical answer it is expected to find. It could unconsciously work backwards from that answer instead of independently deriving the correction.

That would reproduce the central weakness of Agent B's approach.

My amendment: use two distinct stages.

First, derive the point-split divergence from the frozen definitions without being given the expected correction tensor or coefficient.

Second, compare the independently obtained result against Décanini–Folacci and the other primary sources.

The order matters.

The agent must establish:
\[
W=G-H
\]

and explicitly calculate the field equations satisfied by both sides.

It must not infer:
\[
(\Box_g-m^2)W=0
\]

merely because \(W\) is smooth.

It must calculate the actual residual and then propagate that residual through the point-split stress-energy operator.

Only after that calculation should it compare the resulting geometric term with the literature.

Recommended rule

Do not disclose the expected correction until the independent residual-divergence calculation has been recorded and locked.

This is the most important correction to Plan 5.
3. Critical finding: the Go implementation has an unresolved location problem
Plan 5 correctly chooses:

research/experiments/qsg-preflight-001/theorems/qsg-t001/

instead of adding theory-specific material to shared phys-lib/gr code.

This is the correct decision. It follows the existing implementation contract, which explicitly requires a standalone theorem area and prohibits changes to shared modules. 
07-implementation-contract.md

However, there is a practical problem.

The proposed package contains:
qsg-t001/
  AGENTS.md
  manifest.json
  research_metadata.json
  fixture.go
  protocol/
    derivation_test.go


But there is no go.mod in this layout.

The workspace is located at:

/home/chaschel/Documents/go/go.work

The theorem directory is located under:

/home/chaschel/Desktop/biz/TOE-v6/

These are different locations, and the existing workspace lists five other modules.

The plan does not explain how Go will discover and compile the standalone fixture.

Required amendment: add a packaging preflight before implementation.

Agent A must establish:
1. What Go module owns the theorem fixture?
2. How will it resolve phys-math, phys-artifact and the existing libraries?
3. Can the standalone theorem directory execute its build and test commands using the existing workspace?
4. Does this require an additional workspace entry or a separate module?
Do not create a new module or modify go.work automatically.

First establish the smallest valid arrangement. If a module-boundary change is necessary, document it and stop for approval.

This prevents Agent A from spending hours writing tests that cannot compile within the existing architecture.
4. Major finding: independence and responsibility
There are three distinct responsibilities in this programme:
Agent	Responsibility
Agent A	Architecture, infrastructure and implementation
Agent B	Original QFTCS research
Agent C	Independent mathematical audit

The original Agent C instructions explicitly prohibit accepting Agent B's conclusions and require an independent reconstruction. They also prohibit beginning theorem implementation before the audit is complete. 
session-c4a1a27d-d96d-4646-80c2-734cec802d77.md

Plan 5 now assigns Agent A the remaining mathematical audit work in Agent C's directory.

This is understandable as a continuation strategy, but it creates a provenance problem.

We must distinguish between:
- Agent A independently reconstructing the mathematics.
- Agent A merely accepting Agent C's existing documents.
- An independent second reviewer checking Agent A's final correction.
My recommendation:

Keep documents 01 and 02 frozen, but do not treat their conclusions as unquestionable mathematical facts.

Agent A should use them as research records and source-navigation aids, independently checking every mathematical dependency that carries into documents 03 and 04.

For the most delicate calculation—the local correction and its divergence—I recommend a fresh-context independent mathematical review before the theorem is accepted.

That reviewer should receive the model, conventions and primary sources, but not Agent A's expected answer.

If no genuinely independent reviewer is available, record that limitation. Do not claim independent cross-review.
5. Major finding: freeze the exact two-point-function convention
Plan 5 correctly recognizes that the Wightman and symmetric two-point functions must be reconciled.

But merely saying that their normalization will be frozen is insufficient.

The derivation must explicitly declare:
- The Wightman function \(G^+(x,x')\).
- The reversed Wightman function.
- The symmetric two-point function and its normalization.
- The exact Hadamard parametrix \(H(x,x')\).
- The definition of the smooth remainder.
- The bi-differential stress-energy operator.
- The curvature-sign convention.
- The coincidence-limit derivative convention.
In particular, define the relationship between the Wightman function and the symmetric function explicitly, including any factor of two.

For a real scalar field, a commonly used symmetric function is proportional to
\[
G^{(1)}(x,x')
=
G^+(x,x')+G^+(x',x).
\]

The factor used in the actual subtraction prescription must be fixed, not assumed.

Also distinguish differentiation in the first argument from differentiation in the second argument.

This is important because the correction coefficient and sign cannot safely be transferred between papers until their conventions have been reconciled.
6. Three additional efficiency improvements
These are not mathematical blockers, but I would incorporate them.

A. Do not repeatedly extract the entire PDF.

The previous session already extracted all 24 pages.

Use that material as the index for Agent B's defect table. Inspect individual rendered pages only where mathematical notation is ambiguous.

This avoids repeating an expensive operation merely to retrieve information already available.

B. Strengthen the non-modification audit.

Before and after execution, record SHA-256 hashes for:
- Agent B's PDF.
- Agent A's eight completed infrastructure reports.
- Agent C's existing documents 01 and 02.
Record git status --porcelain for affected repositories.

This makes the promise not to modify protected material independently checkable.

C. Improve command validation.

The proposed command:
go vet ./... ; gofmt -l .


should not be treated as a single successful validation.

Record each command's exit status separately. For gofmt -l, require empty output, not merely exit code zero.

The implementation contract also requires a second-checkout PhysTrace verification. Do not silently make this optional. If the facility is unavailable, record it as BLOCKED and distinguish local acceptance from independent artifact acceptance. 
07-implementation-contract.md


7. What Plan 5 gets right — preserve these decisions
I would not change the following parts.
Keep the existing Agent A reconnaissance

No repeated repository mapping, architecture discovery or GR fixture investigation.

Keep documents 01 and 02 frozen

Reuse them after readback and integrity verification. No rewriting.

Keep Décanini–Folacci as the primary working source

Use Wald's inaccessible full texts honestly as source-reported evidence, not independently reconstructed evidence.

Keep the six mathematical dispositions

DERIVED, SOURCE-REPORTED, INDEPENDENTLY REPRODUCED, CONDITIONAL, FALSIFIED and BLOCKED.

Keep human mathematics separate from machine verification

The feasibility study explicitly says that quantum operator semantics and general renormalization procedures are not currently executable in PhysMath. 
04-qsg-mathematical-feasibility.md


I particularly support the insistence that a Go fixture must never pretend to execute an analytical proof that the existing engine cannot perform.

The current infrastructure supports mathematical representation, pinned results, per-instance obligations and provenance. Universal conservation remains a human mathematical proof obligation. 
08-handoff.md

8. Recommended execution sequence
I would make the final workflow slightly more precise:
#chatgpt-mermaid-_r_sq_{font-family:-apple-system-body,ui-sans-serif,-apple-system,system-ui,"Segoe UI","Helvetica","Apple Color Emoji","Arial",sans-serif,"Segoe UI Emoji","Segoe UI Symbol";font-size:16px;fill:rgb(237, 237, 237);}@keyframes edge-animation-frame{from{stroke-dashoffset:0;}}@keyframes dash{to{stroke-dashoffset:0;}}#chatgpt-mermaid-_r_sq_ .edge-animation-slow{stroke-dasharray:9,5!important;stroke-dashoffset:900;animation:dash 50s linear infinite;stroke-linecap:round;}#chatgpt-mermaid-_r_sq_ .edge-animation-fast{stroke-dasharray:9,5!important;stroke-dashoffset:900;animation:dash 20s linear infinite;stroke-linecap:round;}#chatgpt-mermaid-_r_sq_ .error-icon{fill:rgb(48, 48, 48);}#chatgpt-mermaid-_r_sq_ .error-text{fill:rgb(237, 237, 237);stroke:rgb(237, 237, 237);}#chatgpt-mermaid-_r_sq_ .edge-thickness-normal{stroke-width:1px;}#chatgpt-mermaid-_r_sq_ .edge-thickness-thick{stroke-width:3.5px;}#chatgpt-mermaid-_r_sq_ .edge-pattern-solid{stroke-dasharray:0;}#chatgpt-mermaid-_r_sq_ .edge-thickness-invisible{stroke-width:0;fill:none;}#chatgpt-mermaid-_r_sq_ .edge-pattern-dashed{stroke-dasharray:3;}#chatgpt-mermaid-_r_sq_ .edge-pattern-dotted{stroke-dasharray:2;}#chatgpt-mermaid-_r_sq_ .marker{fill:rgb(175, 175, 175);stroke:rgb(175, 175, 175);}#chatgpt-mermaid-_r_sq_ .marker.cross{stroke:rgb(175, 175, 175);}#chatgpt-mermaid-_r_sq_ svg{font-family:-apple-system-body,ui-sans-serif,-apple-system,system-ui,"Segoe UI","Helvetica","Apple Color Emoji","Arial",sans-serif,"Segoe UI Emoji","Segoe UI Symbol";font-size:16px;}#chatgpt-mermaid-_r_sq_ p{margin:0;}#chatgpt-mermaid-_r_sq_ .label{font-family:-apple-system-body,ui-sans-serif,-apple-system,system-ui,"Segoe UI","Helvetica","Apple Color Emoji","Arial",sans-serif,"Segoe UI Emoji","Segoe UI Symbol";color:rgb(237, 237, 237);}#chatgpt-mermaid-_r_sq_ .cluster-label text{fill:rgb(237, 237, 237);}#chatgpt-mermaid-_r_sq_ .cluster-label span{color:rgb(237, 237, 237);}#chatgpt-mermaid-_r_sq_ .cluster-label span p{background-color:transparent;}#chatgpt-mermaid-_r_sq_ .label text,#chatgpt-mermaid-_r_sq_ span{fill:rgb(237, 237, 237);color:rgb(237, 237, 237);}#chatgpt-mermaid-_r_sq_ .node rect,#chatgpt-mermaid-_r_sq_ .node circle,#chatgpt-mermaid-_r_sq_ .node ellipse,#chatgpt-mermaid-_r_sq_ .node polygon,#chatgpt-mermaid-_r_sq_ .node path{fill:rgb(36, 22, 63);stroke:rgb(120, 73, 209);stroke-width:1px;}#chatgpt-mermaid-_r_sq_ .rough-node .label text,#chatgpt-mermaid-_r_sq_ .node .label text,#chatgpt-mermaid-_r_sq_ .image-shape .label,#chatgpt-mermaid-_r_sq_ .icon-shape .label{text-anchor:middle;}#chatgpt-mermaid-_r_sq_ .node .katex path{fill:#000;stroke:#000;stroke-width:1px;}#chatgpt-mermaid-_r_sq_ .rough-node .label,#chatgpt-mermaid-_r_sq_ .node .label,#chatgpt-mermaid-_r_sq_ .image-shape .label,#chatgpt-mermaid-_r_sq_ .icon-shape .label{text-align:center;}#chatgpt-mermaid-_r_sq_ .node.clickable{cursor:pointer;}#chatgpt-mermaid-_r_sq_ .root .anchor path{fill:rgb(175, 175, 175)!important;stroke-width:0;stroke:rgb(175, 175, 175);}#chatgpt-mermaid-_r_sq_ .arrowheadPath{fill:rgb(175, 175, 175);}#chatgpt-mermaid-_r_sq_ .edgePath .path{stroke:rgb(175, 175, 175);stroke-width:1px;}#chatgpt-mermaid-_r_sq_ .flowchart-link{stroke:rgb(175, 175, 175);fill:none;}#chatgpt-mermaid-_r_sq_ .edgeLabel{background-color:rgb(0, 0, 0);text-align:center;}#chatgpt-mermaid-_r_sq_ .edgeLabel p{background-color:rgb(0, 0, 0);}#chatgpt-mermaid-_r_sq_ .edgeLabel rect{opacity:0.5;background-color:rgb(0, 0, 0);fill:rgb(0, 0, 0);}#chatgpt-mermaid-_r_sq_ .labelBkg{background-color:rgba(0, 0, 0, 0.5);}#chatgpt-mermaid-_r_sq_ .cluster rect{fill:rgb(48, 48, 48);stroke:rgba(255, 255, 255, 0.15);stroke-width:1px;}#chatgpt-mermaid-_r_sq_ .cluster text{fill:rgb(237, 237, 237);}#chatgpt-mermaid-_r_sq_ .cluster span{color:rgb(237, 237, 237);}#chatgpt-mermaid-_r_sq_ div.mermaidTooltip{position:absolute;text-align:center;max-width:200px;padding:2px;font-family:-apple-system-body,ui-sans-serif,-apple-system,system-ui,"Segoe UI","Helvetica","Apple Color Emoji","Arial",sans-serif,"Segoe UI Emoji","Segoe UI Symbol";font-size:12px;background:rgb(48, 48, 48);border:1px solid rgba(255, 255, 255, 0.15);border-radius:2px;pointer-events:none;z-index:100;}#chatgpt-mermaid-_r_sq_ .flowchartTitleText{text-anchor:middle;font-size:18px;fill:rgb(237, 237, 237);}#chatgpt-mermaid-_r_sq_ rect.text{fill:none;stroke-width:0;}#chatgpt-mermaid-_r_sq_ .icon-shape,#chatgpt-mermaid-_r_sq_ .image-shape{background-color:rgb(0, 0, 0);text-align:center;}#chatgpt-mermaid-_r_sq_ .icon-shape p,#chatgpt-mermaid-_r_sq_ .image-shape p{background-color:rgb(0, 0, 0);padding:2px;}#chatgpt-mermaid-_r_sq_ .icon-shape .label rect,#chatgpt-mermaid-_r_sq_ .image-shape .label rect{opacity:0.5;background-color:rgb(0, 0, 0);fill:rgb(0, 0, 0);}#chatgpt-mermaid-_r_sq_ .label-icon{display:inline-block;height:1em;overflow:visible;vertical-align:-0.125em;}#chatgpt-mermaid-_r_sq_ .node .label-icon path{fill:currentColor;stroke:revert;stroke-width:revert;}#chatgpt-mermaid-_r_sq_ .node .neo-node{stroke:rgb(120, 73, 209);}#chatgpt-mermaid-_r_sq_ [data-look="neo"].node rect,#chatgpt-mermaid-_r_sq_ [data-look="neo"].cluster rect,#chatgpt-mermaid-_r_sq_ [data-look="neo"].node polygon{stroke:url(#chatgpt-mermaid-_r_sq_-gradient);filter:drop-shadow( 1px 2px 2px rgba(185,185,185,1));}#chatgpt-mermaid-_r_sq_ [data-look="neo"].swimlane.cluster rect{filter:none;}#chatgpt-mermaid-_r_sq_ [data-look="neo"].node path{stroke:url(#chatgpt-mermaid-_r_sq_-gradient);stroke-width:1px;}#chatgpt-mermaid-_r_sq_ [data-look="neo"].node .outer-path{filter:drop-shadow( 1px 2px 2px rgba(185,185,185,1));}#chatgpt-mermaid-_r_sq_ [data-look="neo"].node .neo-line path{stroke:rgb(120, 73, 209);filter:none;}#chatgpt-mermaid-_r_sq_ [data-look="neo"].node circle{stroke:url(#chatgpt-mermaid-_r_sq_-gradient);filter:drop-shadow( 1px 2px 2px rgba(185,185,185,1));}#chatgpt-mermaid-_r_sq_ [data-look="neo"].node circle .state-start{fill:#000000;}#chatgpt-mermaid-_r_sq_ [data-look="neo"].icon-shape .icon{fill:url(#chatgpt-mermaid-_r_sq_-gradient);filter:drop-shadow( 1px 2px 2px rgba(185,185,185,1));}#chatgpt-mermaid-_r_sq_ [data-look="neo"].icon-shape .icon-neo path{stroke:url(#chatgpt-mermaid-_r_sq_-gradient);filter:drop-shadow( 1px 2px 2px rgba(185,185,185,1));}#chatgpt-mermaid-_r_sq_ .node text{font-size:14px;font-weight:600;letter-spacing:normal;fill:rgb(206, 176, 251);}#chatgpt-mermaid-_r_sq_ .edgeLabels text{font-size:13px;font-weight:600;letter-spacing:-0.08px;fill:rgb(206, 176, 251);}#chatgpt-mermaid-_r_sq_ .node tspan[font-weight="normal"],#chatgpt-mermaid-_r_sq_ .edgeLabels tspan[font-weight="normal"]{font-weight:600;}#chatgpt-mermaid-_r_sq_ .edgeLabel .label rect{opacity:1;rx:13px;ry:13px;fill:rgb(23, 12, 38);stroke:rgb(65, 47, 92);stroke-width:1px;}#chatgpt-mermaid-_r_sq_ .node rect,#chatgpt-mermaid-_r_sq_ .node circle,#chatgpt-mermaid-_r_sq_ .node ellipse,#chatgpt-mermaid-_r_sq_ .node polygon,#chatgpt-mermaid-_r_sq_ .node path{fill:rgb(44, 24, 74);stroke:rgba(255, 255, 255, 0.1);stroke-width:1px;}#chatgpt-mermaid-_r_sq_ .node rect{rx:16px;ry:16px;}#chatgpt-mermaid-_r_sq_ .mermaid-decision .label-container{fill:rgb(23, 12, 38);stroke:rgb(65, 47, 92);stroke-dasharray:2px,2px;}#chatgpt-mermaid-_r_sq_ .edgePaths .flowchart-link{stroke:rgb(175, 175, 175);stroke-width:1px;stroke-linecap:round;stroke-linejoin:round;}#chatgpt-mermaid-_r_sq_ .marker{fill:rgb(175, 175, 175);stroke:rgb(175, 175, 175);}#chatgpt-mermaid-_r_sq_ :root{--mermaid-font-family:-apple-system-body,ui-sans-serif,-apple-system,system-ui,"Segoe UI","Helvetica","Apple Color Emoji","Arial",sans-serif,"Segoe UI Emoji","Segoe UI Symbol";}Verify existing files andhashesFreeze exact mathematicalconventionsBlind independent HadamardcalculationDerive residual point-splitdivergenceDerive local correctionindependentlyCompare against primarysourcesIndependent mathematicalcross-checkConservation theoremjustified?Document BLOCKED orFALSIFIEDComplete theorem contractand Phys-PseudoResolve Go module boundaryMinimal standalone fixtureBuild, tests, vet, formattingIndependent PhysTraceverificationFinal evidence-based handoffNoYes





Notice two deliberate changes:
- The independent calculation comes before comparison with the expected correction.
- The Go packaging problem is resolved only after the mathematical gate.
This avoids premature engineering.

9. Final verdict
Recommended after amendments
Area	Decision
Overall direction	Approve
Existing reconnaissance	Freeze
Existing classical derivation	Freeze
Primary-source strategy	Approve with blind derivation
Conservation calculation	Highest priority
Independent review	Strengthen
Phys-Pseudo	Retain existing specification
Go implementation	Conditional on mathematical and packaging gates
Architectural redesign	Not justified

One final scientific point: Do not make Agent A's success dependent on proving Agent B wrong in every respect.

There are two separate questions:
1. Is Agent B's derivation mathematically valid?
2. Is the conservation theorem itself true under the properly specified renormalization prescription?
The first can be false while the second is true.

That distinction must remain intact throughout the investigation.
My recommendation
Revise Plan 5 with the five amendments above, particularly the blind correction derivation and the Go module-boundary preflight. Then authorize Agent A to execute it.

The goal is not another impressive research report. It is one independently defensible mathematical theorem, represented faithfully by our existing infrastructure, without expanding that infrastructure unnecessarily.
