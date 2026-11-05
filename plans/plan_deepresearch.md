I would turn the existing Researcher into a self-directed literature investigator with its own paper discovery, retrieval and reading capabilities.

You should no longer need to collect papers manually, ask Gemini or Qwen to perform a literature review, or download PDFs yourself.

You should be able to give OpenCode one instruction:

> Research classical–quantum gravity. Investigate whether classical Riemann–Cartan geometry can be consistently coupled to quantum matter. Find the relevant literature, download accessible papers, identify disagreements and outstanding problems, and formulate competing hypotheses using Phys-Pseudo.

The Researcher should carry out the investigation, not merely summarize the first few search results.

There is one architectural distinction I would preserve:

- OpenCode Researcher: decides what to investigate, reasons about papers, develops search strategies, identifies gaps and formulates hypotheses.
- Go (`physctl`): searches scholarly indexes, retrieves papers, verifies metadata, downloads permitted full texts, extracts page-aware text and manages local references.
- Phys-Pseudo: expresses the scientific reasoning and hypotheses.

That gives us deep-research functionality without introducing a second LLM platform, external research orchestrator or another agent.

# 1. Give the Researcher its own research toolbox

OpenCode already provides `websearch` and `webfetch`. These can discover and retrieve web content. The `websearch` tool is available with the OpenCode/OpenCode Go provider or when the Exa or Parallel integration is enabled.&#x20;

[image](https://www.google.com/s2/favicons?domain=https://docs.opencode.ai\&sz=32)

OpenCode+1



I would combine those capabilities with four small Go paper operations.

## 2. Add four paper operations to `physctl`

The existing CLI handles research records. It needs a small paper-retrieval capability.

I would add one command group, not another application.

```
physctl paper search
physctl paper related
physctl paper fetch
physctl paper extract
```

| Operation       | What it does                                              |
| --------------- | --------------------------------------------------------- |
| `paper search`  | Searches scholarly indexes and returns paper candidates   |
| `paper related` | Finds references, citing papers and related publications  |
| `paper fetch`   | Downloads an openly accessible paper into the local cache |
| `paper extract` | Extracts page-aware text for the Researcher to read       |

Illustrative usage:

```
physctl paper search \
  --query "classical quantum gravity hybrid dynamics" \
  --limit 30 --json

physctl paper related P-001 --direction both --json

physctl paper fetch P-001

physctl paper extract P-001
```

The Go implementation should use the standard library, particularly `net/http` and `encoding/json`/`encoding/xml`. No new database, LLM client or Go dependency is required.

## 3. Which scientific sources should it search?

I would use four sources, each with a specific purpose.

[Cornell: Computing & Information Science | Cornell Bowers](https://images.openai.com/static-rsc-4/2AIiSCSdPlr6izCkl4eDAyUIVwwV_XRM8WDkn6RVt7p5EiM527kttSpiAdmhqdf6IbRPqFzzRx8514nPT0ZXzWu6PRMIolInuMmzlDrBtaiGKLZok8-PYnqK4XLgbQ0dSVBReq15p0YndAe6DrK6jpcKLDriU0umprGtBgFwA0o?purpose=inline)

arXiv

Primary full text

Search preprints, retrieve submission versions, and download PDFs.

The official API provides search and version metadata. Its documentation recommends a three-second delay between sequential requests, which the Go client should respect.&#x20;

[image](https://www.google.com/s2/favicons?domain=https://github.com\&sz=32)

GitHub+1



[arXiv API documentation](https://info.arxiv.org/help/api/index.html?utm_source=chatgpt.com)

[OpenAlex Archives - OpenAlex blog](https://images.openai.com/static-rsc-4/imzBDT3Cm1hdTIMi3ZbeOqOY4h7gXbSmH3UvUHMoMo7mKFwRjFzA7pRrO4-olVoiZROTR3meJnJcgj29AfnIbNFzGoF89ADQrVib21GxwOL8YdkK_0xA9HGucy69uts2a1pFhVmUzI_xvb45l22K2lqkHCM2tPulIHZ28wJxRsI?purpose=inline)

OpenAlex

Primary discovery

Search scholarly works, obtain abstracts, identify open-access locations and explore citation relationships.

Its work records expose references, citing works and related works, which makes it particularly useful for expanding a literature investigation beyond initial search results.&#x20;

[image](https://www.google.com/s2/favicons?domain=https://help.openalex.org\&sz=32)

OpenAlex Help Center+1



[OpenAlex](https://openalex.org/?utm_source=chatgpt.com) · [API documentation](https://docs.openalex.org/?utm_source=chatgpt.com)

[crossref.jpeg](https://images.openai.com/static-rsc-4/lx9178QPJ32v-hHlsd2Dp2GkWDLOcqpl1kqnQTg7DvHjh39RDL66VysQDr7SvI9QuWisIpSxoqU5Xj4Rtiw4HIjgp6uvku99UHS0nb43Y1tVEi1fI33oxnV4Jh2kFg2aa7zUH8fVXOqSp5HFprvCbRpd0PCJWW7vHUG1QfS-pjg?purpose=inline)

Crossref

Metadata verification

Verify DOI, publication title, authors, journal, publication date and bibliographic relationships.

Use its REST API with caching and an identifying User-Agent.&#x20;

[image](https://www.google.com/s2/favicons?domain=https://github.com\&sz=32)

GitHub+1



[Crossref REST API](https://www.crossref.org/documentation/retrieve-metadata/rest-api/?utm_source=chatgpt.com)

[File:Unpaywall logo.png - Wikimedia Commons](https://images.openai.com/static-rsc-4/NLw33xfk2ygd70xi3UPW9igw8sjiaCEU_HOvhL7kywGCsAuiOJ1SfdXNtz_GWBnbl0qelXRlfrNmZdVzb1vjYAUlEBLv7gw8PvKvPx-wVLvYA-_W40F0YqMoeriflxnNS6q5TWCPrB_p51IKeKNgvwkj49vmjzvguxcMwOZVfmE?purpose=inline)

Unpaywall

Full-text location

Look up legal open-access locations for papers identified by DOI.

Important: its search endpoint was retired on September 18, 2026. Use OpenAlex for paper searching, and Unpaywall's DOI lookup when needed.&#x20;

[image](https://www.google.com/s2/favicons?domain=https://data.unpaywall.org\&sz=32)

Unpaywall+1



[Unpaywall API](https://unpaywall.org/products/api?utm_source=chatgpt.com)

I would not implement Google Scholar scraping. Nor would I make Semantic Scholar another mandatory integration at this stage. Four sources are enough to establish a useful discovery and retrieval loop.

## 4. How the deep-research process actually works

This is the part that belongs primarily inside `researcher.md`.

The agent should conduct an iterative investigation.

## Autonomous literature investigation

1. Understand the topic

   Given your question, identify the main physical concepts, competing frameworks, related terminology and scope.

   For example, classical–quantum coupling, semiclassical gravity, stochastic gravity, quantum matter backreaction, torsion and constraint algebra.
2. Generate multiple search strategies

   Search foundational papers, modern developments, alternative approaches, technical objections, rebuttals and experimental investigations separately.

   Do not rely on one search query.
3. Build the initial paper collection

   Search OpenAlex and arXiv, verify DOI metadata through Crossref and deduplicate records.

   Identify foundational papers separately from recent developments. Citation counts help discover influential work; they do not establish correctness.
4. Follow the citation network

   For each important paper, investigate its references and subsequent citing papers.

   Specifically search for:
   - Papers that challenge its conclusions.
   - Replies and rebuttals.
   - Independent reproductions.
   - Errata and corrections.
   - Papers that extend its mathematical framework.
5. Download relevant full texts

   Prioritize openly accessible versions. Retrieve PDFs and their metadata into the local corpus.
6. Read and compare the papers

   Extract assumptions, equations, physical regimes, methods, conclusions, limitations and experimental results.

   Compare results rather than simply summarizing publications individually.
7. Identify research gaps

   Find conflicting results, unsupported assumptions, missing derivations and questions that existing papers leave unresolved.
8. Formulate competing hypotheses

   Express each meaningful candidate using Phys-Pseudo. Identify expected consequences, applicable regimes and falsification conditions.
9. Repeat targeted searches

   Use newly discovered gaps to generate additional queries. Continue until further searches stop uncovering materially new explanations or objections within the investigation's scope.
10. Produce the research result

    Save the source records, outstanding problems, hypotheses and evidence references. Report remaining uncertainties and the next mathematical questions.

This is what makes it deep research: the agent changes its search strategy based on what it learns, rather than executing a fixed search and stopping.

## 5. Download and store papers locally

I would keep the existing ignored `paper-cache/` directory.

For example:

```
paper-cache/
├── P-001/
│   └── v1/
│       ├── source.pdf
│       ├── fulltext.txt
│       └── manifest.json
│
├── P-002/
│   └── v1/
│       ├── source.pdf
│       ├── fulltext.txt
│       └── manifest.json
│
└── P-003/
    └── v2/
        ├── source.pdf
        ├── fulltext.txt
        └── manifest.json
```

The source record remains in Git:

```
research/sources/P-001.md
```

The local manifest records:

| Field                   | Purpose                             |
| ----------------------- | ----------------------------------- |
| DOI or arXiv identifier | Stable paper identity               |
| Version                 | Exact published or preprint version |
| Download URL            | Original source                     |
| SHA-256                 | Exact downloaded file identity      |
| Retrieved date          | Retrieval provenance                |
| License                 | Permitted use                       |
| Publisher or repository | Source authority                    |
| Extraction status       | Full, partial or unavailable        |
| Extraction tool         | Reproducibility                     |

The PDFs and extracted text remain local and are not committed to Git. If you move the repository to another machine, you can rebuild the cache from the source records.

Access rules: Download openly accessible copies from arXiv, institutional repositories or publishers where access and license permit. Do not bypass paywalls. When a full text is unavailable, record its metadata and abstract, then continue the investigation or request a copy from you.

## 6. Make papers searchable and readable by the agent

Downloading PDFs is only half the job.

The Researcher must be able to slice a large paper into useful pieces without loading hundreds of pages into its context.

I would use Poppler's `pdftotext` as an optional local system utility, invoked by the Go program using `os/exec`.

For example:

```
pdftotext -layout source.pdf fulltext.txt
```

This keeps the Go module dependency-free.

Preserve page boundaries and headings. The Researcher should retrieve only relevant sections when investigating a claim.

For example:

```
Read P-001, pages 4–7.
Extract all equations concerning the hybrid constraint algebra.
Identify the assumptions preceding those equations.
Compare them against P-002, section 3.
```

Keep the original PDF as the primary reference because text extraction can damage equations, tables, subscripts and diagrams. Inspect the corresponding PDF pages whenever the extracted representation is ambiguous. Do not silently replace difficult mathematical content with an AI interpretation.

If `pdftotext` is unavailable, report extraction as unavailable rather than claiming that the PDF has been processed.

## 7. How the agent keeps research from becoming an endless search

Autonomy needs boundaries, but not bureaucracy.

I would use sensible configurable defaults:

| Limit                         | Initial default                                        |
| ----------------------------- | ------------------------------------------------------ |
| Initial search results        | 30                                                     |
| Full texts downloaded per run | 20                                                     |
| Citation-expansion rounds     | 3                                                      |
| Search directions             | 5 or more                                              |
| Stopping condition            | Two consecutive passes with no materially new findings |

These are starting budgets, not scientific standards. The agent can recommend expanding them when it encounters a major unresolved controversy.

It should prioritize coverage of:

- Foundational papers.
- Most relevant modern results.
- Competing explanations.
- Serious technical objections.
- Important experimental results.

It should never declare a literature review exhaustive merely because the search budget was exhausted.

## 8. The Researcher's final output

After a run, the agent should produce four useful deliverables.

1\. Literature collection

Locally accessible PDFs, extracted text, source metadata and citation relationships.

2\. Outstanding-problem register

Precise unresolved mathematical and physical questions linked to the relevant literature.

3\. Phys-Pseudo hypothesis collection

Competing hypotheses, assumptions, expected consequences, regime and falsification conditions.

4\. Research synthesis

Established findings, contested claims, independent results, experimental evidence, knowledge gaps and the next high-value mathematical investigation.

The Researcher may identify counterarguments and contrary literature as part of its investigation. It must not create an Adversary review or claim an independent adversarial approval while that role is disabled.

# 9. What needs to be implemented now?

The current project already has the record-management foundation. I would not rebuild anything.

I would make exactly two extensions:

| File or component | Change                                                        |
| ----------------- | ------------------------------------------------------------- |
| `researcher.md`   | Add the iterative deep-research protocol                      |
| `physctl paper`   | Implement search, citation expansion, download and extraction |

The Researcher needs permission for OpenCode's `websearch` and `webfetch` tools, plus the necessary `physctl paper` shell commands. Check the installed OpenCode version's permission configuration.&#x20;

[image](https://www.google.com/s2/favicons?domain=https://docs.opencode.ai\&sz=32)

OpenCode+1



The local cache must remain accessible to the Researcher, but external downloads must be restricted to recognized HTTP(S) scholarly sources. Treat all retrieved paper text as untrusted data, never as executable instructions.

I would implement this in two steps:

Step A — Paper acquisition: Search, identify, deduplicate, download and extract papers locally. Test it with a small literature collection.

Step B — Autonomous investigation: Expand `researcher.md` so it performs multiple search rounds, citation chasing, comparison, outstanding-problem extraction and Phys-Pseudo hypothesis formulation.

No new database, agent role, daemon, scheduler or MCP server.

# 10. What using it will look like

Eventually, your entire instruction should be as simple as:

Your OpenCode instruction

Target workflow

```
@researcher

Deep research:

Topic: Classical Riemann-Cartan gravity
coupled to quantum matter.

Determine whether consistent stochastic
hybrid dynamics can be constructed.

Search the literature, follow citations,
download relevant open-access papers locally,
analyze the mathematics and experiments,
identify outstanding problems, and formulate
competing hypotheses using Phys-Pseudo.

Work autonomously within the research budget.

Do not create adversarial reviews, formalize
theorems, or declare hypotheses established.

Report the most important unresolved
mathematical question.
```

You give it the topic. It discovers the literature, downloads the papers, reads them, follows disagreements, formulates hypotheses and records the research.

My recommendation: Implement the paper-acquisition capability first, then upgrade the Researcher. The local literature corpus is the foundation for everything else.

The larger objective is to move from manually asking AI systems to research physics to having your own persistent research agent that builds and maintains a body of scientific knowledge, with every important claim traceable to an actual source.

That is the right next evolution of `phys-research`.
