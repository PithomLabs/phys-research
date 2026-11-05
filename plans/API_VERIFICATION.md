API_VERIFICATION.md — Phase 0 Reconnaissance Evidence
Location: /home/chaschel/.opencode/plan/API_VERIFICATION.md
Status: Phase 0 complete; evidence recorded; ready for user review before Phase 1
Environment: Linux; OpenCode agent session; Go module at /home/chaschel/Desktop/biz/TOE-v5/phys-research (Plan 4 A/B/C completed)
Date: 2026-11-04 (UTC)

=== Verification method ===
Actual HTTP requests from the implementation environment using Go-standard-equivalent `curl` with identifying User-Agent, bounded timeouts (10-20s), and safe redirect handling (`-L`). No production bulk downloads. No invented credentials.

=== Service classifications ===

1. arXiv API
Classification: VERIFIED
Endpoint tested: https://export.arxiv.org/api/query?search_query=all:quantum+gravity+cartan&start=0&max_results=3
Method: GET; User-Agent: opencode-research/1.0 (phys-research; phase-0)
Status: 200
Content-Type: application/atom+xml; charset=utf-8
Latency: 0.713047 s
Response size: 6017 bytes
Results: 3 entries returned (Atom `<entry>` elements present)
Rate limits / retry: arXiv documentation recommends 3-second sequential delay, one connection at a time; no Retry-After header observed in response body; no HTTP 503 during test.
Authentication: None required.
Implementation recommendation: VERIFIED — primary physics-preprint source. Implement sequential request delay (>= 3 s) in Go client. Cache Atom results. Deduplicate by arXiv identifier.

2. OpenAlex API
Classification: VERIFIED_WITH_LIMITATIONS
Endpoint tested: https://api.openalex.org/works?search=quantum+gravity+cartan&per_page=5&filter=open_access.is_oa:true
Method: GET; User-Agent: opencode-research/1.0; Mailto: research-test@example.org (polite pool requirement documented but not enforced for basic access)
Status: 200
Content-Type: application/json
Latency: 2.346052 s
Response size: 92272 bytes
Results: 5 returned; total count: 13450; per_page: 5
Response fields verified: meta.count, meta.per_page, meta.notice, meta.db_response_time_ms, meta.cost_usd; results[].id (openalex.org/...), results[].title, results[].doi, results[].open_access (is_oa, oa_status, oa_url, any_repository_has_fulltext)
Rate limits / budget: Response notice: "No API key, so this request shares one free daily budget ($0.10) with everyone on your network's IP address." Documentation: 10 requests/second without key; 30 with free API key; $1/day budget per key; query complexity affects budget consumption.
Authentication: Optional free API key available at https://openalex.org/settings/api; endpoint also allows POST /api-keys with {"email":"..."} to obtain key without manual registration.
Implementation recommendation: VERIFIED_WITH_LIMITATIONS — primary literature discovery source. Recommend obtaining a free API key (email-based) before implementing bulk discovery to avoid shared-budget exhaustion. Implement bounded search (limit initial results to 30 or less per strategy) and respect rate limits. Deduplicate by DOI and arXiv identifier. Treat citation counts as discovery aids, never as correctness proof.

3. Crossref REST API
Classification: VERIFIED
Endpoint tested: https://api.crossref.org/works/10.1103/PhysRevLett.129.121102
Method: GET; User-Agent: opencode-research/1.0 (phys-research; phase-0; mailto:research-test@example.org)
Status: 200
Content-Type: application/json
Latency: 1.060977 s
Response size: 19452 bytes
Response fields verified: status ("ok"), message-type ("work"), message.indexed.date-time, message.publisher, message.license[].start.date, message.reference-count, message.issued.date-parts
DOI verification: Confirmed DOI resolves to American Physical Society (APS) publication; metadata matches expected bibliographic identity.
Rate limits / concurrency: Crossref documentation recommends polite pool with identifying User-Agent and `mailto` parameter; individual record lookups have different limits from list queries. Response headers: no HTTP 429 observed during single test; must implement caching (do not repeat identical DOI lookups within session) and respect any 429 responses with retry/backoff.
Authentication: None required; polite pool requires User-Agent + mailto.
Implementation recommendation: VERIFIED — DOI metadata verification source. Use for confirming publication identity of discovered papers, not for full-text retrieval. Cache results by DOI. Implement retry with exponential backoff on HTTP 429. Do not rely on Crossref for paper search; use for verification only.

4. Unpaywall DOI lookup
Classification: REQUIRES_CONFIGURATION (endpoint verified; full use blocked by missing email parameter)
Endpoint tested: https://api.unpaywall.org/v2/10.1038/nature12373 (without email parameter, to confirm endpoint behavior)
Method: GET; User-Agent: opencode-research/1.0
Status: 422
Content-Type: application/json
Latency: 0.895010 s
Response size: 142 bytes
Response content: {"HTTP_status_code":422,"error":true,"message":"Email address required in API call, see http://unpaywall.org/products/api"}
Endpoint verification: Confirmed endpoint exists, responds to HTTPS, requires `email` parameter in query string (`?email=<email>`). No redirection to localhost or private networks.
Rate limits: Documentation: 100,000 calls/day per email; search endpoint retired 2026-09-18 (use OpenAlex for search instead). DOI lookup remains available.
Authentication: Required — email parameter mandatory; no key/token required beyond valid email.
Environment state: UNPAYWALL_EMAIL environment variable NOT SET; MAIL and EMAIL variables NOT SET; no valid email configured in this session.
Implementation recommendation: REQUIRES_CONFIGURATION — optional supplementary open-access discovery. Before using in production: user must supply a valid email via environment variable (UNPAYWALL_EMAIL). If unavailable, report UNAVAILABLE for full-text discovery; do not invent an email. Once configured, use DOI lookup for open-access location only when OpenAlex open_access data is insufficient. Do not make Unpaywall mandatory in the acquisition pipeline; OpenAlex already provides open-access locations for most open-access papers.

5. Optional local PDF extraction (`pdftotext`)
Classification: VERIFIED
Binary: /usr/bin/pdftotext
Version: 24.02.0 (Poppler developers; Copyright 2005-2024)
Status: Binary responds to `-v`; executable and available.
Implementation recommendation: VERIFIED — optional system utility. Call via `os/exec` with `-layout` option. Never a required Go dependency. If unavailable at runtime, `physctl paper extract` reports UNAVAILABLE (explicit diagnostic). Never claim successful extraction without actually invoking the utility and verifying non-empty output.

=== API reliability assessment ===
Verified combination (minimum reliable pipeline):
- arXiv (primary physics preprints, version metadata, open-access PDF download)
- OpenAlex (primary discovery, citation relationships, open-access location) — REQUIRES free API key for reliable bulk use
- Crossref (DOI verification) — reliable with polite pool
- Unpaywall (optional DOI-based OA lookup) — REQUIRES_CONFIGURATION (email parameter)
- `pdftotext` (optional extraction) — available in this environment

=== Important notes ===
- No production bulk downloads performed during reconnaissance. No papers downloaded in this phase.
- No credentials invented or borrowed. No remote/credential changes made.
- No `git config`, no `core.hooksPath` changes, no remote creation.
- Existing Plan 4 commits (13d76b8, 3415b6a, df41122) and AGENTS.md patch preserved untouched.
- All HTTP requests made with bounded timeouts (10-20s) and identifying User-Agent.
- Rate limits observed in response bodies (arXiv docs only; OpenAlex budget notice; Crossref polite-pool recommendation; Unpaywall 100k/day doc). No HTTP 429 encountered during minimal tests.
- `UNPAYWALL_EMAIL` not configured: user decision required before using Unpaywall in production pipeline.
- Research budget defaults (30/20/3/5/2) not changed; these remain operational limits for Phase 2 implementation.

=== Next step recommendation ===
Phase 0 complete. Ready for user review of this evidence. If approved, proceed to Phase 1 (`physctl paper` implementation). Before Phase 1:
- Confirm whether user wants to configure `UNPAYWALL_EMAIL` for full Unpaywall verification.
- Confirm whether user wants an OpenAlex free API key obtained (optional but recommended for reliable bulk discovery).
- Confirm `P-` namespace reuse for paper source records (user clarification: reuse existing prefix).
- Confirm Phase 1 should reuse existing record-writing templates (`physctl new`) rather than inventing a new ID system.
