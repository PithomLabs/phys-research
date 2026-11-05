Understood. This is the implementation agent's remediation report, not the independent adversarial review. I would treat it as a claim that the nine findings have been addressed, pending independent verification.

The report is encouraging, but I see one important omission.

## 1. Remediation status

| Finding                       | Implementation report | My assessment      |
| ----------------------------- | --------------------- | ------------------ |
| CR-001 — Redirect SSRF        | Fixed                 | Re-review required |
| CR-004 — Download security    | Fixed                 | Re-review required |
| CR-005 — API reliability      | Fixed                 | Re-review required |
| CR-003 — ID allocation        | Fixed                 | Re-review required |
| CR-002 — Frontmatter tabs     | Fixed                 | Re-review required |
| CR-007 — Locators             | Fixed                 | Re-review required |
| CR-008 — Workspace paths      | Fixed                 | Re-review required |
| CR-009 — Review freshness     | Clarified             | Likely resolved    |
| CR-006 — OpenCode permissions | Not reported          | Still open         |

plan_imp(1).md

## 2. Two things I would investigate particularly carefully

DNS rebinding protection

The implementation says it resolves and validates all IP addresses before connecting.

That alone is insufficient if the HTTP transport subsequently performs another DNS resolution when establishing the connection.

The adversary must verify that the actual connection uses a validated IP address, while preserving TLS hostname verification and SNI. It should also investigate proxy behavior and IPv6 redirects.

Rate limiting

The implementation reports in-memory token buckets. Because `physctl` is a command-line application, the adversary should check whether the arXiv three-second delay is respected across separate process invocations, not just multiple requests within one invocation.

These are verification questions, not claims that the fixes are defective.

## 3. What I recommend next

Do not ask the implementation agent to review its own changes.

Return the repository to the independent adversarial reviewer in a fresh session. Give it a focused remediation verification assignment.

# Independent Remediation Verification — phys-research

The implementation agent reports that CR-001 through CR-009 have been addressed.

Conduct a fresh, read-only verification. Do not trust the implementation report without examining the actual code and tests.

### 1. Verify every original finding

For each CR-001 through CR-009, assign exactly one disposition:

- FIXED
- PARTIALLY FIXED
- NOT FIXED
- NOT REPRODUCED
- DEFERRED

Provide code locations and reproducible evidence.

### 2. Highest priority: download security

Specifically investigate:

- Redirect validation on every hop.
- DNS rebinding and whether connections actually use validated IP addresses.
- TLS hostname verification after IP pinning.
- IPv4 and IPv6 private-address restrictions.
- HTTP proxy behavior.
- Path traversal and symlink-based cache escapes.
- PDF signature verification.
- Content-Type handling.
- Streaming size limits.
- Temporary-file cleanup and atomic promotion.

Use controlled local HTTP test servers. Do not conduct bulk downloads.

### 3. API reliability

Verify:

- arXiv's three-second sequential delay across separate CLI invocations and concurrent processes.
- OpenAlex API-key transmission.
- Cursor pagination without silent truncation.
- Retry-After handling for 429 and 503.
- Bounded retries.
- DOI normalization and deduplication.

### 4. Revisit CR-006 explicitly

The implementation report does not mention CR-006.

Verify actual OpenCode permissions for every agent independently, including shell permission behavior.

Test command variations involving Git merge, push, config modification and attempts to edit protected files.

Do not assume local hooks or prompts provide security.

### 5. Check remaining corrections

Verify:

- ID allocation under concurrent processes.
- Quoted frontmatter tab rejection.
- Locator validation without rejecting legitimate named sections.
- go.work path containment, including symlink escapes.
- Consistent review-freshness behavior.

### 6. Regression verification

Run:

gofmt -l .\
go vet ./...\
go build ./...\
go list ./...\
go test -count=1 ./...

Also run the relevant existing physctl validation and allowlisted derivation checks.

Report the exact output and package counts. Do not claim a test ran if it did not.

### 7. Final disposition

Determine whether autonomous external paper fetching can safely resume.

Use the original severity classification. Identify every remaining blocker.

STRICTLY READ-ONLY. No fixes, commits, configuration changes or scientific-record modifications.

The goal is independent verification of remediation, not another implementation exercise.

## 4. My recommendation

Keep the Researcher in restricted testing mode until this independent review is complete.

The implementation report is a good sign: it says the critical downloader changes, API reliability work and existing scientific provenance mechanisms were preserved.&#x20;

plan_imp(1).md



But I would not yet declare the downloader secure or the agent permissions reliable.

The next milestone is not another feature. It is independently demonstrating that the fixes actually close the vulnerabilities.
