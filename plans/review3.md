The third review is useful, but its TLS conclusion is likely a false positive. The PDF-fetching hold should remain in place for other reasons.

## My assessment

| Finding                           | Assessment                             | Decision                                    |
| --------------------------------- | -------------------------------------- | ------------------------------------------- |
| CR-001 — TLS `ServerName`         | Test appears invalid                   | Do not change TLS code based on this report |
| CR-004 — Download security        | Unresolved                             | Keep PDF fetching disabled                  |
| CR-005 — Persistent rate limiting | Implemented, not behaviorally verified | Require integration tests                   |
| CR-006 — OpenCode permissions     | Unresolved                             | Do not claim agent isolation                |
| CR-002, 003, 007, 008, 009        | Reported fixed; regression checks pass | Accept provisionally                        |

### 1. TLS: challenge the adversarial finding

The report says its test requests `https://test.local` using a certificate valid for `example.com`. That certificate should fail hostname verification regardless of whether Go correctly supplies the TLS server name. The test does not isolate the suspected bug.&#x20;

adv_review3(5).md



I independently reproduced the relevant scenario using Go 1.23.2:

- Certificate has `test.local` as its DNS SAN.
- Request goes to `https://test.local`.
- Custom `DialContext` connects to `127.0.0.1`.
- `TLSClientConfig.ServerName` remains empty.
- HTTPS succeeds.

This is consistent with `net/http.Transport` supplying the URL hostname for TLS when `ServerName` is empty.

Disposition: CR-001 is not established. Require a corrected test using a certificate valid for the requested hostname, then test the actual `safeFetch` implementation. Do not introduce a TLS workaround merely to satisfy a broken test.

### 2. Download security: still a blocker

The cache symlink escape is a legitimate concern. The report describes a path in which a symlink inside `paper-cache` could redirect writes outside the cache.&#x20;

adv_review3(5).md



However, I would strengthen the proposed fix: checking `EvalSymlinks` after directory creation is useful, but by itself may leave a time-of-check/time-of-use race if another process can replace the path afterward.

The implementation should enforce cache-root containment and prevent symlink traversal during writes, with an integration test attempting the escape.

The DNS-rebinding concern also needs a test proving that the IP actually used for the connection is the IP that passed validation—not merely that DNS was checked earlier.

Disposition: keep `physctl paper fetch` disabled.

### 3. Rate limiting: promising, not accepted yet

The persistent `flock`-based limiter is the right direction for coordinating independent CLI processes. But the report explicitly says that its cross-process behavior has not been tested.&#x20;

adv_review3(5).md



Require tests for simultaneous invocations, process termination, corrupt state, interrupted waits, and the OpenAlex limit. Until then, treat the feature as implemented but unverified.

### 4. OpenCode permissions: stop treating configuration as a security boundary

The review identifies bypasses through shell commands such as `git -c`, `sed`, `cat`, `tee`, and `mv`, while subagent permission inheritance remains unverified.&#x20;

adv_review3(5).md



My recommendation is not to keep adding shell-pattern rules in an attempt to create a security sandbox.

Use operating-system isolation instead:

- Run the agent in a separate user or container.
- Give it write access only to the designated research workspace.
- Make protected configuration and implementation files read-only to the agent.
- Keep credentials and human acceptance controls outside its environment.

A local Git hook and an environment variable are workflow safeguards, not a trustworthy human-only merge boundary. A protected remote branch with separate credentials is still needed for that.

## Next implementation milestone

Do not add features. Give the implementation agent this narrowly scoped task:

# Phys Research — Final Security Remediation

Perform only the following work. Do not add features, change scientific records, alter research hypotheses, introduce dependencies, or modify the architecture.

1. **TLS verification**
   - Reproduce the TLS issue using a trusted test certificate whose DNS SAN matches the requested URL hostname.
   - Use a custom `DialContext` that connects to a validated IP.
   - Keep `TLSClientConfig.ServerName` empty for the initial test.
   - Test the actual `safeFetch` implementation.
   - Do not change TLS code unless the correctly constructed test demonstrates a failure.
2. **Cache path protection**
   - Prevent symlink traversal and writes outside the canonical paper-cache root.
   - Account for symlinked parent directories and concurrent path replacement.
   - Add an adversarial integration test attempting to write outside the cache.
3. **Persistent rate limiting**
   - Add behavioral tests using independent processes.
   - Verify the arXiv three-second interval across processes.
   - Test process termination, corrupt state, interrupted waits, and OpenAlex throttling.
   - Verify the in-memory limiter cannot bypass or conflict with the persistent limiter.
4. **OpenCode security**
   - Do not rely on shell string patterns as a security boundary.
   - Determine the actual permission-inheritance behavior for subagents.
   - Document the limitations of OpenCode permissions.
   - Recommend an OS-level isolation boundary; do not claim secure isolation without demonstrating it.
5. **Regression and preservation**
   - Run gofmt, go vet, go build, go test -count=1 ./..., and all existing physctl validation commands.
   - Preserve all existing research records and their hashes.
   - Preserve the standard-library-only requirement.
   - Report each finding as FIXED, NOT REPRODUCED, UNRESOLVED, or NOT TESTED.
   - Provide test names, commands, and actual results.

Stop after this remediation. Do not enable autonomous PDF fetching or merge changes into an acceptance branch.

Bottom line: The scientific provenance infrastructure remains in good shape according to the reported regression checks.&#x20;

adv_review3(5).md

&#x20;The security review has correctly kept two important boundaries open, but it has not correctly established the TLS bug. Resolve the cache-write risk, verify the limiter, and establish real OS isolation before authorizing autonomous paper acquisition.

