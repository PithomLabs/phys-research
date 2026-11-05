# Implementation Task: Autonomous Literature Research for phys-research

## Objective

Extend the existing `phys-research` project so its Researcher agent can independently investigate a specified physics topic.

The Researcher must be able to:

1. Discover relevant scientific literature.
2. Search multiple scholarly sources.
3. Follow references, citations, rebuttals and related publications.
4. Download legally accessible papers into a local reference library.
5. Extract page-aware text for detailed investigation.
6. Compare competing theoretical claims and experimental results.
7. Identify outstanding mathematical and physical problems.
8. Formulate competing hypotheses using the existing Phys-Pseudo language.
9. Save traceable research records for subsequent human investigation and mathematical formalization.

The intended experience is that the user supplies a research topic to OpenCode and the Researcher performs the literature investigation without requiring manual paper collection or a separate deep-research agent.

This is an extension of the existing Researcher, not a redesign of phys-research.

## 1. Reconnaissance first

Before changing anything, inspect:

- `AGENTS.md`
- `.opencode/agents/researcher.md`
- Existing `physctl` commands and packages
- Existing research-record metadata and templates
- `phys-pseudo/PSEUDO.md`
- Existing source-record conventions
- Existing `.gitignore`
- Installed OpenCode configuration and available tools

Verify the current implementation rather than relying on this task description as proof that a feature exists.

Preserve:
- The existing Go module and standard-library-only convention.
- The seven Phys-Pseudo constructs exactly.
- Existing research-record types and metadata schema.
- Existing ID allocation and relationship conventions.
- The human-controlled acceptance boundary.
- The distinction between structural validation and scientific correctness.

Do not modify unrelated Phys repositories.

## 2. Paper discovery and retrieval

Add a `paper` command group to the existing `physctl` executable.

Required operations:

```text
physctl paper search
physctl paper related
physctl paper fetch
physctl paper extract
```

### A. Search

Support discovery through:

1. arXiv API — preprints and version metadata.
2. OpenAlex API — scholarly works, open-access locations, references and citing works.
3. Crossref REST API — DOI and publication metadata.
4. Unpaywall DOI lookup — optional open-access location discovery.

Use official APIs and documented access methods. Verify current API documentation, authentication requirements, rate limits and usage conditions during implementation.

Requirements:

- Search by topic, keywords, authors, DOI and arXiv identifier where supported.
- Return machine-readable JSON.
- Deduplicate results by DOI and arXiv identifier.
- Preserve original publication and preprint identifiers.
- Distinguish peer-reviewed publications, preprints and other source types.
- Preserve publication dates and available version information.
- Return source URLs and available abstract/metadata.
- Identify open-access availability.
- Never interpret citation count as proof of scientific correctness.
- Do not silently discard conflicting metadata.

Do not scrape Google Scholar or bypass access restrictions.

### B. Related literature

Support citation expansion:

- References cited by a paper.
- Papers citing a paper.
- Related works identified by scholarly indexes.
- Relevant replies, rebuttals, corrections and subsequent versions.

Deduplicate results and preserve the relationship and discovery source.

Use bounded traversal. Do not recursively download an unlimited citation network.

### C. Download

Download openly accessible full texts when permitted by their license and access conditions.

Prioritize:
- arXiv PDFs.
- Institutional repositories.
- Publisher-provided open-access copies.
- Legitimate locations identified by OpenAlex or Unpaywall.

Do not bypass paywalls, authentication or access controls.

For unavailable full texts, preserve metadata and abstract and report the limitation.

Requirements:
- Verify HTTP status and content type.
- Apply request timeouts and reasonable file-size limits.
- Calculate SHA-256 of the downloaded file.
- Preserve exact source URL, retrieval date, paper version and license information where available.
- Avoid redownloading unchanged files.
- Handle redirects safely.
- Restrict downloads to valid HTTPS sources and prevent requests to localhost, private networks and other unsafe destinations.
- Treat downloaded documents as untrusted input.

Respect API rate limits, especially arXiv's sequential request guidance. Use an identifiable User-Agent.

### D. Local paper extraction

Store downloaded papers in an ignored local cache:

```text
paper-cache/
├── <paper-id>/
│   └── <version>/
│       ├── source.pdf
│       ├── fulltext.txt
│       └── manifest.json
```

Do not commit original PDFs or extracted full texts to Git.

Implement PDF text extraction by invoking the optional local `pdftotext` utility through Go's `os/exec`.

Do not add a Go PDF dependency.

Requirements:

- Preserve page boundaries.
- Preserve section headings where possible.
- Retain the original PDF as the primary source.
- Record the extraction method and status.
- Detect missing or empty extracted text.
- Clearly identify extraction limitations.
- Return useful page-specific text to the Researcher.
- Never silently invent page numbers or equation locators.

If `pdftotext` is unavailable, report `UNAVAILABLE`. Do not claim successful extraction.

The Researcher must inspect the original PDF pages whenever mathematical notation, equations, tables, diagrams or text extraction is ambiguous.

## 3. Local source records

Downloaded material must not become an independent source of truth separate from the existing research records.

Use the existing source-record template and ID allocator.

A source record should retain, where available:

- Paper ID.
- DOI.
- arXiv identifier and version.
- Full title.
- Authors.
- Journal or publication venue.
- Publication date.
- Source URL.
- Open-access location.
- License or access information.
- PDF SHA-256.
- Retrieval date.
- Extraction status.
- Relevant claims and precise page, section, equation, figure or table locators.

Keep metadata in the existing research-record format.

The local `manifest.json` is a rebuildable cache manifest, not a new research artifact type.

Do not create a new artifact family or identity convention.

## 4. Deep-research protocol for the Researcher

Extend `.opencode/agents/researcher.md`.

The Researcher must accept a topic and central research question, then perform an iterative investigation.

### Stage A — Scope

Identify:
- Central scientific question.
- Relevant physical concepts.
- Competing theoretical frameworks.
- Synonyms and alternative terminology.
- Relevant mathematical structures.
- Applicable physical regimes and scales.

Do not assume the user's preferred hypothesis is correct.

### Stage B — Literature discovery

Generate multiple targeted search queries, including:

- Foundational literature.
- Important reviews.
- Recent theoretical developments.
- Alternative explanations.
- Technical objections and no-go results.
- Replies, rebuttals and corrections.
- Relevant experiments and observational constraints.

Do not rely on a single query or publication.

### Stage C — Citation expansion

For important papers:

- Examine their reference lists.
- Identify subsequent citing papers.
- Search for independent reproductions.
- Search for objections and rebuttals.
- Identify changes across preprint and publication versions.
- Identify relevant corrections and errata.

Prioritize papers by scientific relevance, not popularity alone.

### Stage D — Paper analysis

For each relevant paper, extract:

- Main scientific question.
- Physical assumptions.
- Mathematical definitions.
- Important equations.
- Derivations and stated results.
- Physical regime and effective-theory scale.
- Predictions.
- Experimental methods and measurements.
- Limitations explicitly acknowledged by the authors.

Distinguish:
- Authors' claims.
- Mathematically derived results.
- Independently reproduced results.
- Experimental observations.
- Speculative interpretations.

Every extracted substantive claim must have a precise source locator.

### Stage E — Literature synthesis

Compare papers and identify:

- Agreements.
- Contradictions.
- Different assumptions producing different conclusions.
- Results that have been independently reproduced.
- Results challenged by subsequent work.
- Missing derivations.
- Unresolved mathematical questions.
- Experimental gaps.
- Questions that existing literature does not answer.

Do not treat a disagreement as resolved merely because one paper is newer.

Assess publication status and evidential strength.

### Stage F — Outstanding problems

Convert genuine unresolved questions into linked issue records.

Each issue should state:

- The precise unresolved question.
- Relevant sources.
- Existing results.
- Missing mathematical or physical steps.
- Applicable regime.
- Why the question matters.
- What result would resolve or narrow it.

Avoid creating duplicate issues.

### Stage G — Hypothesis formulation

Develop competing explanations where scientifically justified.

Use the existing seven Phys-Pseudo constructs:

```text
given
define
assume
claim
derive
if/then/else
?
```

Do not invent additional constructs.

Each physical hypothesis must have:
- Explicit assumptions.
- Defined physical regime.
- Expected consequences.
- Relevant competing explanation or null model.
- Falsification condition in ordinary research-record prose.
- Unresolved `?` obligations.

Mathematical hypotheses must instead specify premises, conclusions, domains and proof obligations.

Do not write an independent derivation unless one has actually been performed.

Do not invent equations to fill gaps.

### Stage H — Expected versus actual

Maintain separate comparisons:

**Theory**
- Paper-claimed expectation.
- Independent theoretical result, where available.
- Assumptions and domain.
- Derivation source.
- Difference or unresolved issue.

**Experiment**
- Predicted observable.
- Measurement conditions.
- Actual measured value.
- Units and statistical/systematic uncertainties.
- Methodology and source locator.
- Difference and resulting constraint.

Never represent an author's assertion as an independently derived result.

Never represent a prediction or target as an experimental measurement.

### Stage I — Iterative investigation

Use discovered gaps and disagreements to formulate new search queries.

Continue investigating while new searches uncover materially relevant explanations, objections, derivations or evidence.

Stop when:
- The configured research budget is exhausted.
- Consecutive search rounds produce no materially new findings.
- Relevant full text is unavailable.
- A decisive mathematical or empirical question requires human work.

Never claim that a literature search is exhaustive merely because its budget is exhausted.

### Stage J — Research report

Produce a compact report containing:

1. Research question and scope.
2. Literature discovered.
3. Established results.
4. Contested results and competing interpretations.
5. Outstanding problems.
6. Competing Phys-Pseudo hypotheses.
7. Expected versus actual comparisons.
8. Most important unresolved mathematical question.
9. Recommended next research task.
10. Missing sources or limitations.

Save the supporting records in Git. Avoid lengthy narrative summaries when structured scientific artifacts communicate the result more directly.

## 5. Research budget

Use configurable defaults to prevent unbounded searches and downloads.

Initial defaults:

| Operation | Default limit |
|---|---:|
| Initial search results | 30 |
| Full texts downloaded | 20 |
| Citation-expansion rounds | 3 |
| Distinct search strategies | 5 |
| No-new-findings stopping condition | 2 consecutive rounds |

These are operational limits, not scientific completeness criteria.

The Researcher may recommend expanding them when a major controversy or missing result justifies it, but must disclose the additional scope.

## 6. OpenCode integration

OpenCode remains the orchestrator and LLM runtime.

Do not add:
- Another LLM client.
- Another research agent.
- A custom orchestrator.
- A background scheduler.
- A message bus.
- An MCP server.

Expand only the existing Researcher role.

Ensure it has the required permissions to:
- Use available OpenCode websearch/webfetch tools when configured.
- Execute the specific `physctl paper` operations.
- Write source, issue, hypothesis and evidence records.
- Read the local paper cache.

Do not broaden permissions to unrelated paths or arbitrary system operations.

Check the actual OpenCode configuration version. Do not assume a tool or permission exists without verifying it.

The Researcher may identify competing explanations and contrary literature, but must not create Adversary review records or imply an independent adversarial review has occurred.

## 7. Go implementation constraints

- Standard library only.
- Reuse existing `physctl`.
- Reuse existing ID allocation and record-writing operations.
- Reuse existing metadata parser.
- Use `net/http`, `encoding/json`, `encoding/xml`, `crypto/sha256` and other standard-library packages where appropriate.
- Apply bounded retries, timeouts and response-size limits.
- Use atomic writes.
- Use deterministic JSON output.
- Return explicit diagnostic codes for partial or unavailable operations.
- Do not silently convert failed retrieval or extraction into success.
- Do not execute downloaded document contents.
- Do not permit arbitrary URL fetching through agent-controlled shell arguments.

Do not refactor unrelated packages.

## 8. Testing

Test:

- arXiv search and version handling.
- OpenAlex search and citation expansion.
- Crossref DOI metadata verification.
- Optional Unpaywall lookup.
- Duplicate-paper detection.
- Source-version changes.
- Download restrictions and redirects.
- PDF size and content-type validation.
- SHA-256 verification.
- Cache reuse and redownload behavior.
- PDF extraction success and unavailable states.
- Page-boundary preservation.
- Malformed metadata and API responses.
- Network timeout and rate-limit handling.
- JSON output.
- Research-record provenance.
- Existing `physctl` functionality.

Use `httptest` or equivalent controlled HTTP test servers. Do not make production web requests inside unit tests.

Document which tests use live services and do not claim they were performed unless actually run.

## 9. Implementation sequence

### Phase 0 — Reconnaissance

Confirm existing behavior and identify the smallest required changes.

### Phase 1 — Paper acquisition

Implement and test:

- `physctl paper search`
- `physctl paper related`
- `physctl paper fetch`
- `physctl paper extract`

Demonstrate a complete retrieval cycle using one openly accessible paper.

### Phase 2 — Researcher enhancement

Update only `researcher.md` and necessary shared instructions.

Add the iterative deep-research protocol, citation expansion, source comparison, bounded investigation and hypothesis formulation.

### Phase 3 — End-to-end demonstration

Run a topic-level investigation using:

**Topic:** Classical Einstein–Cartan geometry coupled to quantum matter through stochastic dynamics.

Demonstrate:

1. Multiple scholarly search strategies.
2. Source discovery and deduplication.
3. Local download and text extraction of accessible papers.
4. Citation expansion.
5. Identification of outstanding issues.
6. Competing Phys-Pseudo hypotheses.
7. Research-record validation.
8. A compact literature synthesis with source locators.

Do not attempt to solve Branch C during this implementation.

## 10. Completion criteria

The implementation is complete when a user can start OpenCode, invoke the Researcher with a physics topic, and have it autonomously:

- Find relevant papers.
- Verify their bibliographic metadata.
- Download accessible full texts.
- Maintain local paper references.
- Read and compare relevant passages.
- Follow citations and objections.
- Identify outstanding problems.
- Formulate competing hypotheses using Phys-Pseudo.
- Record results with precise provenance.
- Report unresolved questions and research limitations.

All existing MVP 1 tests and protections must remain intact.

## 11. Final report

Report only:

1. New Go functionality.
2. Changes to the Researcher instructions.
3. Sources and API integrations actually implemented.
4. Local download and extraction results.
5. End-to-end research demonstration.
6. Tests actually executed and outcomes.
7. Limitations, unavailable services and remaining human decisions.

Do not claim that the Researcher has performed deep research unless the complete discovery, retrieval, citation-expansion, analysis and record-writing cycle has actually been demonstrated.

Do not declare generated hypotheses established.

**Primary engineering principle: make the Researcher independently capable of discovering, retrieving, reading, connecting and formulating scientific knowledge, while keeping the entire process small, reproducible, source-grounded and human-supervised.**
