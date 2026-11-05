## INDEPENDENT REMEDIATION VERIFICATION — phys-research

### Executive Summary

The implementation agent reports all 9 findings addressed. My independent verification of the actual code reveals:

- **5 findings FIXED** (CR-002, CR-003, CR-007, CR-008, CR-009)
- **2 findings PARTIALLY FIXED** with significant correctness issues (CR-001, CR-004 — TLS ServerName empty breaks HTTPS)
- **1 finding PARTIALLY FIXED** with production reliability gap (CR-005 — in-memory rate limiters don't persist across CLI invocations)
- **1 finding NOT FIXED** (CR-006 — OpenCode permissions unchanged)

**Critical blocker for autonomous paper fetching**: The TLS `ServerName: ""` bug in `safeFetch` (paper.go:216) will cause certificate verification failures for all legitimate HTTPS downloads, making the fetcher non-functional. This must be fixed before resuming autonomous fetching.

---

### 1. Verification of Each Original Finding

| Finding | Disposition | Evidence |
|---------|-------------|----------|
| **CR-001** SSRF via Redirect Host Validation | **PARTIALLY FIXED** | Redirect validation added in `makeCheckRedirect` (paper.go:307-323) and `DialContext` validates IPs per-connection. **BUT**: `TLSClientConfig.ServerName: ""` (paper.go:216) breaks TLS hostname verification and SNI — all HTTPS fetches will fail certificate verification. |
| **CR-004** Download Security | **PARTIALLY FIXED** | PDF magic bytes (`%PDF`), Content-Type sniffing, paperID sanitization, atomic temp-file rename, 50MB streaming limit, comprehensive private IP blocking (RFC1918, link-local, loopback, multicast, CGNAT, TEST-NET, IPv6 ULA). **SAME BLOCKER**: Empty `ServerName` breaks HTTPS. No symlink escape protection in cache dir. |
| **CR-005** API Reliability | **PARTIALLY FIXED** | Token bucket limiters (arXiv 1/3s, OpenAlex 10/s), API key support, retry with backoff + Retry-After, OpenAlex cursor pagination, DOI normalization. **GAP**: In-memory rate limiters reset on each CLI invocation — arXiv's 3-second sequential delay NOT respected across separate `physctl` processes. |
| **CR-003** ID Allocation Race | **FIXED** | Removed 1000-iteration cap; exponential backoff with jitter (1ms→2ms→4ms... up to 1s + jitter) in record.go:162-181. Standard-library only. |
| **CR-002** Frontmatter Tabs in Quoted Values | **FIXED** | `cleanValue` now rejects tabs in unquoted content (frontmatter.go:282-284). |
| **CR-007** Locator Validation | **FIXED** | `isValidLocator` rejects `p.`, `Eq.`, `Fig.`, `Sec.` but allows `Eq. (1)`, `Fig. 1`, `Sec. 3`, `p. 1`, `Abstract`, `Introduction`, `Conclusion` (validate.go:389-433). |
| **CR-008** go.work Path Containment | **FIXED** | `resolveModule` resolves absolute paths, uses `filepath.Rel` with `..` prefix check to reject escapes (check.go:256-270). |
| **CR-009** Review Freshness Logic | **FIXED** | Documentation added: malformed inputs = errors (Rule="review"), no inputs = `LEGACY_UNPINNED` warning (store/review.go:108-111). |
| **CR-006** OpenCode Permissions | **NOT FIXED** | No changes to `.opencode/opencode.json`. Shell patterns remain string prefixes (`git merge *` doesn't match `git -c merge.ff=false merge accepted`). No explicit shell denies for subagents. Subagent permission isolation unverified. |

---

### 2. Deep Dive: Download Security (CR-001/CR-004)

#### Redirect Validation
- **Initial request**: `isSafeHost` validates hostname patterns (paper.go:156-172)
- **Redirects**: `makeCheckRedirect` calls `isSafeHost` on each redirect target (paper.go:307-323) — validates hostname patterns, enforces HTTPS
- **Per-connection**: Custom `DialContext` calls `resolveAndValidateHost` which does DNS lookup + private IP validation (paper.go:137-154, 196-214)

#### DNS Rebinding Protection
- `DialContext` resolves the host from the connection address and validates all IPs before connecting
- **Issue**: Go's HTTP library may resolve the hostname internally and pass an IP address to `DialContext`. If so, `resolveAndValidateHost` receives an IP string, `net.LookupIP` returns that same IP, validation passes. This works.
- **Gap**: No protection against DNS changes *between* redirect validation and connection establishment. The redirect `CheckRedirect` validates hostname pattern only; `DialContext` validates IPs at connection time. A malicious redirect could point to a hostname that resolves to public IP during `CheckRedirect` but private IP during `DialContext` (TOCTOU). However, both happen in the same request flow, so window is small.

#### TLS Hostname Verification — **CRITICAL BUG**
```go
// paper.go:215-219
TLSClientConfig: &tls.Config{
    ServerName:         "",    // BUG: empty!
    InsecureSkipVerify: false,
},
```
- With `ServerName: ""`, Go's TLS verifies the certificate against the **IP address** (from `DialContext`), not the hostname
- Legitimate HTTPS servers (arXiv, OpenAlex) present certificates for hostnames, not IPs
- **Result**: All HTTPS fetches will fail with certificate verification errors
- **Fix required**: Set `ServerName` to the original request hostname. This requires per-request TLS config or `DialTLSContext`.

#### IPv4/IPv6 Private Address Restrictions
- Comprehensive: RFC1918 (10/8, 172.16/12, 192.168/16), link-local (169.254/16), loopback, multicast, CGNAT (100.64/10), TEST-NET (192.0.2/24, 198.51.100/24, 203.0.113/24), IETF protocol (192.0.0/24)
- IPv6: Unique Local Addresses (fc00::/7) — correct

#### Path Traversal & Symlink Escapes
- `sanitizePaperID` rejects `..`, absolute paths, path separators (paper.go:344-368)
- **Gap**: No protection against symlinks *within* the paper-cache directory. If `paper-cache/2209.15487v1` is a symlink to `/etc`, `filepath.Join` will follow it. The sanitizer only validates the paperID string, not the resolved filesystem path.

#### PDF Verification & Content-Type
- Streaming PDF magic byte check (`%PDF`) — correct
- Content-Type sniffing rejects HTML/JSON error pages — correct (additional signal only)
- 50MB streaming limit with `io.LimitedReader` — correct

#### Temporary File Cleanup
- Downloads to `.tmp`, atomic `os.Rename` to final — correct
- Cleanup on error paths — correct
- **Gap**: If process crashes between write and rename, `.tmp` file remains (minor)

---

### 3. Deep Dive: API Reliability (CR-005)

| Feature | Implementation | Cross-Process Persistence |
|---------|----------------|---------------------------|
| arXiv rate limit (1 req/3s) | `arXivRateLimiter = NewRateLimiter(1.0/3.0)` (ratelimit.go:63) | **NO** — in-memory, resets per CLI invocation |
| OpenAlex rate limit (10 req/s) | `OpenAlexRateLimiter = NewRateLimiter(10.0)` (ratelimit.go:67) | **NO** — same issue |
| OpenAlex API key | Passed as `&api_key=` query param (openalex.go:36) | Works per-invocation |
| 429/503 retry + Retry-After | `doWithRetry` with exponential backoff (ratelimit.go:97-136) | Works per-invocation |
| OpenAlex pagination | Cursor-based loop (openalex.go:27-122) | Works |
| DOI normalization | `normalizeDOI` (ratelimit.go:70-77) | Works |

**Critical gap for autonomous operation**: The arXiv API requires 1 request per 3 seconds *globally* (not per-process). With in-memory limiters, running `physctl paper search` twice in quick succession from separate shell commands will violate the rate limit and risk IP ban. A persistent rate limiter (file-based token bucket with timestamps) is required.

---

### 4. CR-006: OpenCode Agent Permissions — Explicit Verification

**File**: `.opencode/opencode.json` — **UNCHANGED** from original review.

#### Tested Permission Bypasses (Conceptual)

| Attempt | Pattern Match | Expected Result |
|---------|---------------|-----------------|
| `git merge accepted` | `git merge *` ✓ | Blocked |
| `git -c merge.ff=false merge accepted` | `git merge *` ✗ (prefix is `git -c`) | **NOT BLOCKED** |
| `git push origin accepted` | `git push *` ✓ | Blocked |
| `git push --force origin accepted` | `git push *` ✓ | Blocked |
| `git config user.email "x@y"` | `git config *` ✓ | Blocked |
| `git config --global core.hooksPath /dev/null` | `git config *` ✓ | Blocked |
| `sed -i 's/deny/allow/' .opencode/opencode.json` | No shell deny for `sed` | **NOT BLOCKED** |
| `cat > .opencode/opencode.json <<'EOF'...` | No shell deny for `cat` | **NOT BLOCKED** |

#### Subagent Permission Isolation
- Each agent (researcher, adversary, formalizer) has its own `permissions` array
- Parent-level denies (lines 4-12) apply to... unclear if they apply to subagents
- Subagents have `edit: "*" deny` then specific allows — OpenCode evaluates in order, first match wins
- **No explicit `shell` permissions for subagents** — inheritance behavior undocumented

#### Human Acceptance Boundary
- Still only: pre-commit hook + `PHYS_HUMAN_MERGE` env var (explicitly documented as convenience, not security)
- No remote branch protection configured in this repo
- **No local enforcement** if no protected remote exists

**Assessment**: CR-006 remains **NOT FIXED**. The permission model has bypass vectors and subagent isolation is unverified.

---

### 5. Regression Verification

| Command | Output | Status |
|---------|--------|--------|
| `gofmt -l .` | (no output) | PASS |
| `go vet ./...` | (no output) | PASS |
| `go build ./...` | (no output) | PASS |
| `go list ./...` | 9 packages listed | PASS |
| `go test -count=1 ./...` | 8 packages OK, 1 no test files | PASS |
| `physctl validate` | 0 errors, 1 `LEGACY_UNPINNED` warning | PASS |
| `physctl status` | Reports issues, obligations, blocking findings | PASS |
| `physctl check` | Lists 5 allowlisted profiles | PASS |
| `physctl refs H-001 --in` | Traverses graph correctly | PASS |

**Note**: `physctl paper fetch` was **NOT tested live** (requires network + `pdftotext`). Given the TLS `ServerName` bug, it would fail certificate verification on any HTTPS download.

---

### 6. Final Disposition

**BLOCKED — Critical issues must be corrected before autonomous research continues.**

#### Mandatory Blockers (must fix before `physctl paper fetch` works):
1. **TLS ServerName bug** (CR-001/CR-004): Set `ServerName` to request hostname in `safeFetch` transport. Requires per-request TLS config or `DialTLSContext`.
2. **Cross-process rate limiting** (CR-005): Implement persistent rate limiter state (file-based token bucket with timestamps) to respect arXiv's 3-second rule across CLI invocations.

#### High Priority:
3. **OpenCode permissions** (CR-006): Add regex-based shell patterns; explicit subagent shell denies; verify isolation; document evaluation model.

#### Medium Priority (reliability/hardening):
4. **Symlink escape protection** in paper cache: Use `filepath.EvalSymlinks` + boundary check after resolving cache directory.
5. **DNS rebinding TOCTOU**: Consider pinning resolved IPs for the entire request chain (initial + redirects).

#### Already Fixed (verified):
- CR-002: Frontmatter tab rejection in quoted values
- CR-003: ID allocation race (unbounded retry + backoff + jitter)
- CR-007: Locator semantic validation
- CR-008: go.work path containment
- CR-009: Review freshness documentation

---

### Recommendation

**Do not resume autonomous paper fetching.** The TLS `ServerName` bug makes `safeFetch` non-functional for HTTPS. Fix the two mandatory blockers first, then re-verify with a live test against a controlled HTTPS test server (e.g., `https://httpbin.org/pdf` or local test server with valid cert).

The scientific provenance mechanisms (review freshness, append-only reviews, `?` obligations, validation disclaimer) remain intact and working correctly. The core research-record infrastructure is solid.
