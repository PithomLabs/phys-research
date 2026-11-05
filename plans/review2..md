This is a useful independent review, but I found an important error in its most serious finding. I would not blindly authorize the proposed TLS fix.

The reviewer has correctly identified two significant remaining problems, but its claim that an empty TLSClientConfig.ServerName necessarily breaks HTTPS is incorrect for the standard Go net/http.Transport configuration described in this report.
1. My assessment of the findings
Consolidated disposition
Finding	My assessment
CR-001: TLS / redirect security	Reproduce first
CR-004: PDF and cache safety	Partial
CR-005: Cross-process rate limiting	Unresolved
CR-006: OpenCode permissions	Unresolved
CR-003: ID allocation	Fixed
CR-002: Frontmatter	Fixed
CR-007: Source locators	Fixed
CR-008: Workspace paths	Fixed
CR-009: Review freshness	Fixed


adv_review2(7).md
2. Important correction: TLS ServerName
The reviewer claims:
TLSClientConfig: &tls.Config{
    ServerName: "",
    InsecureSkipVerify: false,
}


causes Go to verify the certificate against the IP address rather than the original hostname.

That conclusion overlooks the standard library's behavior.

Go's net/http.Transport automatically fills an empty TLS ServerName using the target hostname. Its connectMethod retains the original URL hostname even when a custom DialContext connects to a validated IP address. 

Git at Google+1

The relevant implementation is:
- connectMethod.targetAddr is populated from the request URL.
- connectMethod.tlsHost() derives the certificate hostname from that target address.
- persistConn.addTLS() fills cfg.ServerName when it is empty.
The TLS documentation also confirms that ServerName is used for hostname verification and SNI. 

Go Packages+1

Therefore:

CR-001's alleged TLS failure is not established.

I would mark it NOT REPRODUCED, pending a controlled HTTPS test. Do not introduce DialTLS, disable certificate verification, or hardcode a hostname just to fix a problem that may not exist.
3. What remains genuinely concerning
CR-005 — Cross-process rate limiting

This is a valid concern.

The limiter is in memory. Each invocation of physctl starts a new limiter, so repeated CLI processes can issue arXiv requests less than three seconds apart.

The fix should preserve your no-database architecture: use a small persistent rate-limit state file with cross-process coordination and crash-safe handling.

Test multiple processes running concurrently, not just goroutines within one process.

CR-006 — OpenCode permissions

This remains unresolved.

The official OpenCode V2 documentation confirms that shell permission resources use whole-value wildcard matching, not arbitrary command-semantic matching. The last matching rule wins, and subagents use their own permission configurations. Global permissions are combined with agent-specific rules. 

OpenCode

Consequently, a rule matching git merge * does not automatically match:
git -c merge.ff=false merge accepted


And denying the edit operation does not prevent an agent from changing a file with unrestricted shell redirection or another shell utility.

The appropriate solution is a broad shell denial with narrowly permitted commands per agent, followed by empirical testing. Do not try to solve this with increasingly elaborate regex patterns.

CR-004 — Symlink escape

The reviewer also identified missing protection against symlinks inside paper-cache.

I would retain this as a medium-priority hardening issue. Verify that cache operations cannot escape the canonical cache root, including through existing symlinked directories.

The review itself distinguishes these remaining weaknesses from the five findings it verified as fixed. 
adv_review2(7).md

4. What I would do next
I would authorize one final, narrowly scoped remediation. No redesign.
Targeted Remediation — phys-research
Use the independent review in adv_review2(7).md.

Implement only confirmed defects. Do not blindly apply every recommendation.
1. Reproduce the TLS finding first
The reviewer claims an empty TLSClientConfig.ServerName breaks HTTPS.

Check this against the actual Go net/http.Transport behavior.

Write a controlled HTTPS integration test using:
- A TLS server with a hostname certificate.
- A custom DialContext that connects to a validated IP.
- An HTTPS request using the original hostname.
- Normal certificate verification enabled.
If standard net/http correctly supplies the request hostname when ServerName is empty, classify the finding as NOT REPRODUCED.

Do not introduce custom DialTLS or disable certificate verification unnecessarily.
2. Cross-process API rate limiting
Fix CR-005.

Enforce arXiv's three-second sequential request interval across:
- Multiple requests in one process.
- Separate physctl invocations.
- Concurrent physctl processes.
Use a minimal standard-library-only persistent mechanism. Do not introduce SQLite, a service or a daemon.

Test concurrent process behavior and recovery after an interrupted process.

Apply appropriate limits and retry behavior to OpenAlex as well.
3. OpenCode permission hardening
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
4. Cache symlink protection
Investigate CR-004's remaining symlink risk.

Ensure paper-cache operations cannot escape the intended cache root through symlinked directories or files.

Add controlled tests.

Keep the implementation portable and standard-library-only.
5. Preserve everything else
- Do not change scientific records.
- Do not alter P-001, H-001, E-001, R-001 or R-002.
- Preserve P-002 through P-006, H-002, I-002, E-002 and D-001.
- Preserve all review-pinning and freshness rules.
- Do not modify Phys-Pseudo.
- No new dependencies, agents, databases or services.
- No Branch C implementation.
6. Verification
Run all existing Go tests and validation.

Add reproducible tests for cross-process rate limiting, HTTPS hostname verification, redirect SSRF and cache symlink containment.

Report the disposition of every original CR finding, distinguishing FIXED, NOT REPRODUCED, PARTIALLY FIXED and UNRESOLVED.

Do not run bulk paper downloads.

Stop after remediation. Request another independent review before declaring the downloader ready for autonomous operation.




5. Overall judgment
The independent reviewer has done something valuable: it tested the implementation's assumptions rather than accepting the implementation agent's success report.

However, adversarial review itself must be subjected to adversarial review. The TLS claim illustrates why.

My current disposition is:
- Research-record system: suitable for continued development.
- Paper downloader: pending targeted security verification.
- Autonomous literature discovery: can continue cautiously with bounded API use.
- Autonomous fetching: hold until the remaining security checks pass.
- Scientific acceptance: remains human-controlled.
We are getting close to a useful research platform. The next improvement should increase trustworthiness, not add functionality.
