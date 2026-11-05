I reviewed `remediation_plan.md`. I approve its sequential approach, but recommend three amendments before implementation.

The plan is correctly structured: reproduce the TLS issue first, then address rate limiting, permissions and cache containment.&#x20;

remediation_plan.md



## 1. Important correction: TLS ServerName

The standalone TLS test should come first.

However, there is a strong reason to expect the adversarial review's finding will not reproduce.

Go's standard `net/http.Transport` automatically assigns the request hostname to `TLSClientConfig.ServerName` when it is empty. This behavior is present in the Go transport implementation.&#x20;

[image](https://www.google.com/s2/favicons?domain=https://go.googlesource.com\&sz=32)

Git at Google+1



Therefore:

- Do not blindly set `ServerName`.
- Test the actual `safeFetch` implementation.
- Use a hostname certificate, custom IP dialer and normal certificate verification.
- Check the original hostname is preserved through redirects.
- If the test passes, classify the finding as `NOT REPRODUCED`.

The reviewer may have mistaken custom IP dialing for changing the TLS verification hostname.

## 2. Amend the rate limiter design

The proposed persistent token bucket is reasonable, but I would make two corrections.

For arXiv: Prefer a persistent minimum-interval limiter rather than a burst-capable token bucket.

Store a `next_allowed_at` timestamp and serialize requests across processes using a stable lock file.

Important safeguards:

- Hold the lock while reserving the request slot and making the request.
- Persist the reservation before issuing the request.
- Never reset a corrupt state file and immediately start sending requests.
- Use conservative recovery after interruption.
- Keep the implementation standard-library-only.

`syscall.Flock` is platform-specific, so explicitly declare the supported OS or provide platform-specific implementations.

The current plan's suggestion to start fresh when the state file is corrupt is unsafe for a rate-limited service.&#x20;

remediation_plan.md



## 3. Amend permissions and symlink protection

OpenCode: Do not assume regex-based shell restrictions exist. Test the actual version and permission matching behavior. Prefer deny-by-default shell access with narrowly justified command permissions.

Cache: `filepath.EvalSymlinks` before `MkdirAll` is insufficient when the destination does not yet exist. Resolve and validate existing parent directories, then revalidate after directory creation. Explicitly document any remaining local filesystem race assumptions.

The proposed symlink check needs stronger tests than simply checking a pre-existing symlink.&#x20;

## Authorization to the implementation agent

Plan remediation approved with three amendments.

Execute sequentially.

### Phase 1 — TLS reproduction

Write and run the standalone controlled HTTPS test against the actual safeFetch implementation.

Test:

- Hostname certificate.
- Custom DialContext connecting to a validated IP.
- Empty TLSClientConfig.ServerName.
- Original hostname verification and SNI.
- Normal certificate verification.
- Redirect behavior.

Go net/http.Transport automatically supplies the target hostname when ServerName is empty. Account for this behavior.

If reproduced, fix it. If not, classify CR-001 TLS finding as NOT REPRODUCED. Do not make unnecessary TLS changes.

### Phase 2 — Persistent rate limiting

Fix cross-process rate limiting.

For arXiv, use a persistent serialized minimum-request interval, not a burst-capable token bucket.

Requirements:

- Shared state across physctl invocations.
- Stable cross-process lock file.
- Persist request reservation before network execution.
- Hold serialization through the request.
- Conservative recovery if state is corrupt or interrupted.
- Never silently reset corrupt state and immediately issue a request.

Maintain standard-library-only constraints. Document OS portability.

### Phase 3 — OpenCode permissions

Test actual OpenCode V2 permission behavior.

Use deny-by-default shell access and only narrowly justified command allowances.

Test shell variations, compound commands, redirection and subagent isolation.

Do not claim remote branch protection.

### Phase 4 — Cache symlink protection

Ensure cache operations cannot escape the canonical cache directory through symlinks.

Account for nonexistent directories, symlinked parent directories and revalidation after directory creation.

Keep the solution standard-library-only and document remaining threat-model assumptions.

### Preservation

No scientific-record changes. No Branch C work. No dependencies, databases, agents or services. Preserve all nine original review findings and report their updated dispositions.

### Verification

Run all existing Go tests, validation and allowlisted checks. Add controlled HTTPS, redirect, cross-process rate-limit and symlink-escape regression tests.

Perform one live open-access download only after the security checks pass.

Stop after remediation and request another independent adversarial review.

## My recommendation

Proceed with the standalone TLS test first, exactly as Plan 4 proposes.

The priorities are:

1. Establish whether the TLS problem actually exists.
2. Fix cross-process rate limiting.
3. Establish genuine agent permission restrictions.
4. Protect local cache paths.

Do not expand the research functionality yet. Once these are independently verified, we can safely resume autonomous paper acquisition and topic-level research.
