# Targeted Remediation — phys-research

Use the independent review in adv_review2.md.

Implement only confirmed defects. Do not blindly apply every recommendation.

## 1. Reproduce the TLS finding first

The reviewer claims an empty TLSClientConfig.ServerName breaks HTTPS.

Check this against the actual Go net/http.Transport behavior.

Write a controlled HTTPS integration test using:
- A TLS server with a hostname certificate.
- A custom DialContext that connects to a validated IP.
- An HTTPS request using the original hostname.
- Normal certificate verification enabled.

If standard net/http correctly supplies the request hostname when ServerName is empty, classify the finding as NOT REPRODUCED.

Do not introduce custom DialTLS or disable certificate verification unnecessarily.

## 2. Cross-process API rate limiting

Fix CR-005.

Enforce arXiv's three-second sequential request interval across:
- Multiple requests in one process.
- Separate physctl invocations.
- Concurrent physctl processes.

Use a minimal standard-library-only persistent mechanism. Do not introduce SQLite, a service or a daemon.

Test concurrent process behavior and recovery after an interrupted process.

Apply appropriate limits and retry behavior to OpenAlex as well.

## 3. OpenCode permission hardening

Fix CR-006.

Inspect the actual installed OpenCode V2 configuration semantics.

Ensure:
- Researcher shell access is limited to required physctl paper operations.
- Adversary has no shell access unless specifically justified.
- Formalizer receives only narrowly scoped validation/test commands.
- Global and agent-specific permissions are consistent.
- No agent can change AGENTS.md, OpenCode configurations, Go validation code, Git configuration or protected acceptance state through shell commands.

Use catch-all shell denial followed by narrowly scoped allowances where supported.

Test actual execution, including command variations, redirection, compound commands and alternative Git invocation forms.

Do not rely solely on the prompt or edit-tool restrictions.

Do not claim accepted-branch protection exists: there is still no protected remote.

## 4. Cache symlink protection

Investigate CR-004's remaining symlink risk.

Ensure paper-cache operations cannot escape the intended cache root through symlinked directories or files.

Add controlled tests.

Keep the implementation portable and standard-library-only.

## 5. Preserve everything else

- Do not change scientific records.
- Do not alter P-001, H-001, E-001, R-001 or R-002.
- Preserve P-002 through P-006, H-002, I-002, E-002 and D-001.
- Preserve all review-pinning and freshness rules.
- Do not modify Phys-Pseudo.
- No new dependencies, agents, databases or services.
- No Branch C implementation.

## 6. Verification

Run all existing Go tests and validation.

Add reproducible tests for cross-process rate limiting, HTTPS hostname verification, redirect SSRF and cache symlink containment.

Report the disposition of every original CR finding, distinguishing FIXED, NOT REPRODUCED, PARTIALLY FIXED and UNRESOLVED.

Do not run bulk paper downloads.

Stop after remediation. Request another independent review before declaring the downloader ready for autonomous operation.
