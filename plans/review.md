The adversarial review found a real security problem. I would not authorize autonomous paper downloading until the critical findings are fixed.

The good news is that the core research-record and provenance mechanisms appear sound. The problems are concentrated in paper acquisition, API reliability and OpenCode permission enforcement.

## Independent code-review verdict

READY WITH FINDINGS

Nine findings reported

- 2 critical download-security findings.
- 3 high-priority reliability and permission findings.
- 3 medium-priority correctness and validation findings.
- 1 medium-priority review-freshness edge case.

Source: independent adversarial review, not independently reproduced by me against the codebase.&#x20;

adv_review(20261104-134054).md



# 1. What must be fixed first

## Critical: paper-download security

CR-001 and CR-004 should be treated as one blocking security workstream, although their individual defects need separate tests.

The reviewer reports:

- Redirects can bypass private-network and localhost checks.
- DNS rebinding is not prevented.
- Paper IDs can potentially escape the intended cache directory.
- HTML error pages can be saved as PDFs.
- PDF content is not verified.
- Content-Type validation is insufficient.

adv_review(20261104-134054).md

These defects matter because the Researcher will download URLs obtained from external scholarly metadata. Those URLs must be treated as untrusted.

Required minimum correction:

1. Validate every redirect destination, not just the original URL.
2. Reject loopback, private, link-local and otherwise prohibited IP addresses, including IPv6.
3. Prevent DNS rebinding by validating resolved addresses and ensuring the connection uses a validated address.
4. Validate paper IDs before constructing filesystem paths.
5. Verify actual PDF content, including its magic signature.
6. Reject HTML or JSON error pages masquerading as PDFs.
7. Download to a temporary file, enforce a streaming size limit, verify the PDF and only then move it into the cache atomically.

One nuance: `Content-Length` preflight is useful, but it must not replace the existing streaming size limit. Servers can omit or misreport it. Likewise, accept legitimate PDF responses served as `application/octet-stream` when the actual PDF signature is valid.

Until this is fixed, do not run `physctl paper fetch` against externally supplied URLs.

## High: API rate limits and pagination

CR-005 is also important.

The report says:

- arXiv requests have no enforced three-second sequential delay.
- OpenAlex accepts an API key but never uses it.
- Citation/search results can be silently truncated.
- HTTP 429 and 503 responses are not retried.
- DOI normalization is missing.

adv_review(20261104-134054).md

The Researcher performs multiple search strategies and citation-expansion rounds. Without rate limiting, autonomous investigation could exhaust the OpenAlex budget or cause service restrictions.

Required:

- Enforce arXiv's sequential request delay across invocations, not merely inside one process.
- Actually use `OPENALEX_API_KEY` when configured.
- Implement bounded retries with appropriate `Retry-After` handling.
- Make truncation explicit, rather than silently presenting incomplete results as complete.
- Normalize DOI values before deduplication.

I would fix CR-001/004 and CR-005 before enabling the Researcher's full autonomous acquisition workflow.

# 2. Other findings: what I would accept and how I would resolve them

| Finding                        | My decision                                                        |
| ------------------------------ | ------------------------------------------------------------------ |
| CR-003 ID allocation race      | Fix, but verify with concurrent processes before introducing locks |
| CR-006 OpenCode permissions    | Verify empirically; do not trust configuration patterns alone      |
| CR-002 tabs in quoted metadata | Small fix                                                          |
| CR-007 locator validation      | Improve, but reject the proposed digit-only rule                   |
| CR-008 go.work parsing         | Fix path containment without adding a dependency                   |
| CR-009 malformed review inputs | Likely correct existing behavior; clarify rather than weaken       |

Some important qualifications:

CR-003 — ID allocation: The race is credible, but the reviewer has not demonstrated that the 1,000-retry limit actually fails. Test multiple concurrent processes allocating IDs. Use atomic exclusive creation and retry by rescanning. Do not introduce platform-specific file locking without evidence it is necessary.

CR-006 — agent permissions: The report correctly warns that OpenCode subagents have their own permission configurations.&#x20;

adv_review(20261104-134054).md



Test actual command execution in a disposable checkout. In particular, test Git options, alternate command forms and shell-based file edits.

Also, the report's claim that remote branch protection exists conflicts with our previously recorded state that no remote was configured. Until you independently verify a protected remote, treat `accepted` as unprotected.

CR-007 — source locators: Do not require every locator to contain a number. A valid locator can be `Abstract`, `Conclusion` or `Introduction`. Reject empty and meaningless placeholders, but preserve legitimate named sections. The reviewer identified a validation weakness, but its suggested fix is overly restrictive.

CR-008 — go.work: Do not add `golang.org/x/mod` because that would break our standard-library-only requirement. Keep a restricted parser or use the Go toolchain to validate the file. Always canonicalize resolved paths and reject paths escaping the permitted workspace.

CR-009 — review freshness: A review with no input digests is legitimately legacy. A review containing a malformed input digest is corrupt. It should produce an error, not be reclassified as legacy. I would retain the current distinction and document it.

adv_review(20261104-134054).md

# 3. One reporting inconsistency to resolve

The code-review report says all eight Go packages passed. The earlier Phase 2 implementation report said nine packages passed.

This is not necessarily a defect, but the next verification run should explicitly report:

```
go list ./...
go test -count=1 ./...
```

The review also could not execute paper search, fetch or extraction in its own environment, so its assessment of those functions is primarily static code inspection. It did not reproduce the SSRF exploit.&#x20;

adv_review(20261104-134054).md



That does not make the findings unimportant. It means the remediation agent should reproduce them with controlled tests.

# 4. Recommended remediation order

1. Secure paper downloading

   CR-001 and CR-004. Blocking before external PDF fetch.
2. Fix API reliability

   CR-005. Required before autonomous multi-query investigation.
3. Verify OpenCode permission boundaries

   CR-006. Establish what agents can actually execute.
4. Fix metadata and ID correctness

   CR-002, CR-003 and CR-007.
5. Harden workspace resolution and document review behavior

   CR-008 and CR-009.

The paper downloader and API reliability fixes should receive their own commits and regression tests. Avoid refactoring unrelated code.

# 5. Prompt for the remediation agent

I recommend sending this to a different implementation agent, with the code-review findings attached.

# Remediation Task — phys-research Adversarial Code Review

Implement the necessary corrections identified in `adv_review(20261104-134054).md`.

## Priority 1 — Blocking download-security defects

Fix CR-001 and CR-004.

Requirements:

- Validate every redirect destination.
- Block loopback, private, link-local, multicast and prohibited IPv4/IPv6 destinations.
- Prevent DNS rebinding by connecting only to validated resolved addresses.
- Ensure TLS hostname verification remains correct.
- Validate paper IDs before filesystem operations.
- Prevent path traversal and symlink-based cache escapes.
- Verify actual PDF content and magic signature.
- Reject HTML/JSON error pages saved as PDFs.
- Treat Content-Type as an additional validation signal, not the sole PDF test.
- Enforce streaming size limits even when Content-Length is absent or incorrect.
- Download to temporary files and promote atomically only after verification.
- Clean up failed downloads.

Add adversarial tests for redirects, private addresses, invalid PDF content, malicious IDs, oversized downloads and corrupted responses.

Do not permit autonomous external paper fetching until these tests pass.

## Priority 2 — API reliability

Fix CR-005.

- Enforce arXiv's documented sequential request delay across separate CLI invocations.
- Correct OpenAlex API-key usage.
- Implement bounded retry/backoff for 429 and 503 responses.
- Respect Retry-After.
- Make pagination and truncation explicit.
- Normalize DOI identifiers before deduplication.
- Cache repeated metadata lookups.

Do not invent rate-limit values or disregard provider documentation.

## Priority 3 — OpenCode permission verification

Investigate CR-006.

Test the installed OpenCode version in a disposable checkout.

Verify effective permissions for every subagent independently.

Test Git command variations and shell-based modification of protected paths.

Do not claim the accepted branch is protected without a credential-separated remote and verified branch restrictions.

Do not access or modify credentials.

## Priority 4 — Remaining correctness issues

- CR-003: test concurrent process ID allocation; retain portable standard-library implementation.
- CR-002: reject tabs inside quoted frontmatter values if prohibited by the documented subset.
- CR-007: reject empty and placeholder locators, but allow legitimate named-section locators.
- CR-008: enforce workspace path containment without adding dependencies.
- CR-009: document malformed review inputs as errors; do not classify corrupt pins as legacy.

## Preservation requirements

- No new dependencies.
- Preserve standard-library-only Go convention.
- Preserve all existing research records and their history.
- Preserve P-001, H-001, E-001, R-001 and R-002 exactly.
- Preserve the shared review-freshness implementation.
- Preserve all seven Phys-Pseudo constructs.
- Do not alter scientific conclusions.
- Do not add agents, databases, services or orchestration.
- Do not start new physics research.

## Verification

Run:

- gofmt -l .
- go vet ./...
- go build ./...
- go list ./...
- go test -count=1 ./...
- physctl validate
- Existing allowlisted physctl check profiles.
- New adversarial security regression tests.

Use controlled local HTTP test servers. Do not run bulk external downloads.

After remediation, conduct a live download verification using a known open-access arXiv paper.

Report each original finding as FIXED, NOT REPRODUCED, DEFERRED or UNRESOLVED, with test evidence.

Inspect the final Git diff. Report limitations honestly.

The goal is targeted remediation, not redesign.

## My final recommendation

Keep the Researcher in testing mode, but separate its capabilities:

- Literature metadata discovery: may continue under bounded API usage.
- External PDF downloads: suspend until CR-001 and CR-004 are fixed.
- Full autonomous research loops: suspend until CR-005 is fixed and the API budgets are respected.
- Scientific acceptance: remains human-controlled.

The encouraging part is that the reviewer found the existing provenance, review freshness and scientific-boundary mechanisms largely sound.&#x20;

adv_review(20261104-134054).md



We do not need another architecture. We need to harden the one we have built.
