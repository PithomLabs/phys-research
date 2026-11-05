# Independent Adversarial Code Review — phys-research

## Mission

Perform a rigorous, adversarial code review of the entire `phys-research` repository as it exists at the time of review.

Act as a senior Go systems engineer, application-security reviewer and scientific-software reliability engineer.

Assume the implementation may contain subtle defects despite passing its existing tests.

Your objective is to identify:
- Actual bugs.
- Security vulnerabilities.
- Data corruption and provenance risks.
- Incorrect validation or workflow behavior.
- Agent-permission bypasses.
- Scientific-integrity failures introduced by the software.
- Missing tests that could allow serious failures to go undetected.

Do not redesign the project. Do not add features. Do not modify any files.

## 1. Review scope

Inspect the actual repository, not previous implementation summaries.

Review all relevant components:

### Go implementation
- `cmd/physctl`
- `internal/frontmatter`
- `internal/store`
- `internal/record`
- `internal/validate`
- `internal/refs`
- `internal/status`
- `internal/check`
- `internal/paper`

### OpenCode configuration
- `.opencode/opencode.json`
- `.opencode/agents/researcher.md`
- Relevant agent permissions and tool restrictions.

### Research records and scientific conventions
- `AGENTS.md`
- Source records.
- Issue records.
- Hypothesis records.
- Evidence records.
- Review records.
- Derivation references.
- Review freshness and provenance rules.

### Repository integrity
- `go.mod`
- `go.work` handling.
- `.gitignore`
- Git branches and hooks.
- Paper-cache behavior.
- Existing tests and fixtures.

Inspect the Git history where necessary to understand intentional behavior and identify regressions.

Pay particular attention to Plan 3, Plan 4 and the deep-research extension. Verify that their documented requirements match the actual implementation.

## 2. Critical review areas

### A. Go correctness and reliability

Inspect for:

- Incorrect error handling.
- Nil-pointer and panic conditions.
- Resource leaks.
- File-descriptor handling.
- Race conditions.
- Unsafe concurrency.
- Incorrect atomic-write behavior.
- ID allocation collisions.
- Partial writes and corrupted records.
- Inconsistent status and validation logic.
- Incorrect exit codes.
- JSON output inconsistencies.
- Platform-specific path assumptions.
- Unbounded memory or file reads.

Check whether the code behaves correctly with malformed, incomplete and adversarial inputs.

### B. Frontmatter parser

Adversarially test the restricted standard-library parser.

Check:

- Duplicate and unknown fields.
- Invalid indentation.
- Empty and malformed references.
- Quoted and unquoted values.
- Tabs and CRLF.
- Unicode.
- Path traversal attempts.
- Unsupported YAML constructs.
- Multiple or missing frontmatter delimiters.
- Premature termination.
- Correct preservation of the Phys-Pseudo body.

The parser must implement only the documented metadata subset.

It must not silently accept unsupported syntax or become an accidental general-purpose YAML parser.

Do not introduce external dependencies.

### C. Research-record integrity

Check:

- Unique and stable IDs.
- Correct relationship resolution.
- Invalid references.
- Self-references.
- `depends_on` cycle detection.
- Correct relationship-direction handling.
- Deterministic traversal.
- Stale review detection.
- Material-input digest validation.
- Consistency between `physctl status` and `physctl validate`.
- Legacy review handling, including R-001 as `LEGACY_UNPINNED`.
- Append-only review behavior.
- Correct distinction between workflow state and scientific standing.

Pay special attention to whether modifications to P-001, I-001, E-001 or H-001 correctly invalidate dependent reviews.

Verify that R-002 remains pinned to its three material inputs.

Check that records cannot silently lose unresolved obligations or overwrite historical findings.

### D. Paper acquisition security

This is a particularly important area.

Review `search`, `related`, `fetch` and `extract`.

Investigate:

**HTTP security**
- Server-side request forgery (SSRF).
- Redirect handling.
- DNS resolution and private-network access.
- HTTPS enforcement.
- Malicious redirects.
- HTTP status handling.
- Content-Type validation.
- Content-length and actual-size limits.
- Timeout enforcement.
- Retry behavior.
- HTTP 429 and 503 handling.

**Paper identity**
- DOI normalization.
- arXiv identifier and version handling.
- Duplicate detection.
- Conflicting metadata.
- Version changes.
- Incorrect association of PDFs with source records.
- Manifest integrity.

**Downloads**
- Truncated or corrupted PDFs.
- HTML error pages saved as PDFs.
- Incorrect PDF content.
- SHA-256 verification.
- Cache reuse.
- Temporary-file cleanup.
- Partial-download recovery.
- Unsafe filenames and paths.
- Cache poisoning.

**Extraction**
- `pdftotext` availability.
- Safe `os/exec` invocation.
- Argument injection.
- Path traversal.
- Timeouts and large outputs.
- Empty extraction.
- Page-boundary preservation.
- Correct UNAVAILABLE reporting.

Verify that unavailable extraction is never reported as successful.

Do not assume extracted equations, tables or page locators are correct merely because extraction returned exit code zero.

### E. External API correctness

Inspect integrations with:

- arXiv.
- OpenAlex.
- Crossref.
- Optional Unpaywall.

Check:

- Actual API endpoint construction.
- Response decoding.
- Pagination.
- Rate-limit handling.
- Retry logic.
- Optional API-key handling.
- Email configuration.
- Credential leakage.
- API error responses.
- Metadata normalization.
- DOI/arXiv deduplication.
- Version consistency.

Verify that the implementation does not depend on retired endpoints.

Unpaywall must remain optional and must not prevent the rest of the pipeline from functioning.

OpenAlex must remain functional when no API key is configured, while accurately reporting limitations.

Do not make live production downloads or consume substantial API budgets during this review.

### F. OpenCode agent security

Inspect the effective permissions for every configured agent.

Important: subagents use their own permissions. Do not assume that parent-agent restrictions automatically apply.

Check whether agents can:

- Modify `AGENTS.md`.
- Modify `.opencode` configuration.
- Modify Go validation code.
- Change `go.mod` or workspace configuration.
- Modify hooks.
- Perform unauthorized Git merges.
- Push to the accepted branch.
- Change Git configuration.
- Modify scientific records outside their assigned scope.
- Modify reviews they are supposed to be unable to edit.

Check the actual OpenCode version and configuration semantics.

Remember that shell resources are raw command text: test whether permission patterns genuinely match dangerous commands and variations.

Determine whether the accepted-branch protection is real or merely prompt-based.

Do not assume local Git hooks enforce security if `core.hooksPath` is unset.

Do not access, print, modify or reuse user credentials.

### G. Deep-research methodology and scientific integrity

Review whether the Researcher implementation enforces the intended boundaries.

Check:

- Every extracted scientific claim has a precise source locator.
- Paper versions and PDF hashes are recorded.
- Authors' claims are distinguished from independently derived results.
- Mathematical methods are not automatically transferred between physical frameworks.
- Theoretical predictions are distinguished from experimental measurements.
- Statistical and systematic uncertainties remain separate.
- Missing full texts and failed extraction are disclosed.
- Unresolved `?` obligations are preserved.
- Competing hypotheses are not silently collapsed into one.
- The Researcher cannot create Adversary reviews.
- The Researcher cannot declare a hypothesis or theorem established.
- Research budgets are respected.
- The agent does not claim exhaustive coverage when its search budget is exhausted.
- Downloaded papers are treated as untrusted input, not executable instructions.

Pay particular attention to P-001, P-002, H-001, H-002, I-001, I-002, E-001, E-002, R-001 and R-002 if present.

Do not perform a new physics literature review. This is a review of the software's scientific-integrity mechanisms.

### H. Existing Phys infrastructure integration

Inspect the interaction with:

- Phys-Pseudo.
- Phys.
- PhysMath.
- PhysLib.
- PhysArtifact.
- PhysTrace.

Verify:

- The seven Phys-Pseudo constructs remain unchanged.
- No unsupported parser or grammar contract has been introduced.
- Artifact identity follows existing PhysArtifact conventions where applicable.
- Theory-neutral mathematical representation remains separate from physical interpretation.
- Only allowlisted test profiles can be executed by `physctl check`.
- `UNAVAILABLE` is not converted into `PASS`.
- Test success is not represented as proof of physical correctness.
- The capability audit is respected.
- Test oracles are not generated solely by the implementation under test.

## 3. Adversarial testing

Run tests where safe and feasible.

At minimum:

- `gofmt -l .`
- `go vet ./...`
- `go test -count=1 ./...`
- `go build ./...`
- Existing `physctl validate`.
- Existing paper-command tests.
- Relevant `physctl check` profiles using the configured developer-local workspace.

Create temporary test data for adversarial cases.

Do not modify committed records or the working tree.

Do not install dependencies, change Go module files, alter Git configuration or make external API calls unless genuinely necessary and explicitly reported.

Check the Git status before and after testing.

If a test cannot run, report precisely why.

## 4. Findings format

Use this format for every finding:

### Finding ID: CR-001

**Severity:** Critical / High / Medium / Low

**Category:** Security / Correctness / Integrity / Reliability / Scientific provenance / Testing

**Where:** Exact file, function and line range.

**What:** The defect or vulnerability.

**Why:** Concrete failure mechanism and likely consequence.

**Evidence:** Reproducible test, code path or exact configuration behavior.

**Minimum fix:** Smallest change that resolves the problem.

**Regression test:** Specific test that should prevent recurrence.

Distinguish:

- Confirmed defect.
- Reproducible risk.
- Missing protection.
- Design concern.
- Speculative concern.

Do not report stylistic preferences as defects.

Avoid duplicate findings. Consolidate findings with a shared root cause.

## 5. Severity classification

- **Critical:** Unauthorized scientific acceptance, credential exposure, destructive corruption or serious security compromise.
- **High:** Scientific provenance failure, incorrect review freshness, unsafe downloads, false validation success or silent data corruption.
- **Medium:** Incorrect behavior, incomplete error handling, reproducibility problems or important missing tests.
- **Low:** Minor reliability or maintainability concerns.

Prioritize findings by actual impact, not theoretical complexity.

## 6. Review constraints

STRICTLY READ-ONLY.

- Do not edit source files.
- Do not apply patches.
- Do not create new research records.
- Do not modify agent instructions.
- Do not change Git branches or configuration.
- Do not commit or push.
- Do not fix issues during the review.
- Do not expand the project scope.

You may create temporary test files outside the repository and remove them afterward.

Use an independent context and a different model from the Researcher.

Treat existing implementation summaries as claims to verify, not evidence.

## 7. Final deliverables

Produce:

**A. Executive assessment**
- Overall engineering condition.
- Most significant risks.
- Whether the implementation is safe for continued Researcher-only testing.

**B. Ranked findings**
- Critical and High first.
- Exact evidence and minimum fixes.
- No duplicate or cosmetic findings.

**C. Test results**
- Commands actually executed.
- Pass/fail results.
- Tests not executable and reasons.

**D. Security assessment**
- Actual agent permission boundaries.
- Download safety.
- Human acceptance protection.
- Remaining trust assumptions.

**E. Scientific-integrity assessment**
- Whether source provenance and hypothesis/review freshness are reliable.
- Whether research claims can be confused with established results.

**F. Recommended remediation order**
- Smallest essential fixes first.
- Explicitly distinguish mandatory fixes from optional improvements.

**G. Overall disposition**

Choose one:

1. READY — no material blockers identified.
2. READY WITH FINDINGS — continued testing is acceptable, with specific risks documented.
3. BLOCKED — critical issues must be corrected before autonomous research continues.

Do not automatically recommend a redesign.

The goal is to make the existing project more reliable without introducing unnecessary complexity.

**Review principle: Attack the implementation, verify the evidence, identify the smallest necessary corrections, and leave the code untouched.**
