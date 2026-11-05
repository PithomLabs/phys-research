# Remediation Task — phys-research Adversarial Code Review

Implement the necessary corrections identified in `adv_review.md`.

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
