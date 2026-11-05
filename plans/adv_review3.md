## FINAL INDEPENDENT REMEDIATION VERIFICATION — phys-research

### Executive Summary

My independent verification of the latest implementation reveals:

| Finding | Status | Evidence |
|---------|--------|----------|
| **CR-001/CR-004** TLS ServerName / Download Security | **PARTIALLY FIXED** (TLS bug CONFIRMED) | `tls_test.go` demonstrates certificate verification failure: "certificate is valid for example.com, not test.local" with empty `ServerName` + custom `DialContext`. No symlink escape protection in cache. |
| **CR-005** Cross-Process Rate Limiting | **IMPLEMENTED** (behavioral verification needed) | File-based `PersistentRateLimiter` with `flock` in `persistent_ratelimit.go`. Auto-initialized on first `Search()` call. |
| **CR-006** OpenCode Permissions | **NOT FIXED** | `.opencode/opencode.json` unchanged. Shell patterns are string prefixes (`git merge *` bypassable via `git -c merge.ff=false merge accepted`). No subagent shell denies. |
| **CR-003** ID Allocation Race | **FIXED** | Unbounded retry with exponential backoff + jitter (record.go:162-181). |
| **CR-002** Frontmatter Tabs | **FIXED** | `cleanValue` rejects tabs in quoted values (frontmatter.go:282-284). |
| **CR-007** Locator Validation | **FIXED** | `isValidLocator` rejects `p.`, `Eq.`, `Fig.`, `Sec.` but allows `Eq. (1)`, `Abstract`, etc. (validate.go:389-433). |
| **CR-008** go.work Path Containment | **FIXED** | `resolveModule` uses `filepath.Rel` with `..` prefix check (check.go:256-270). |
| **CR-009** Review Freshness | **FIXED** | Documentation: malformed inputs = errors, no inputs = `LEGACY_UNPINNED` warning (store/review.go:108-111). |

**Critical blocker remains**: TLS `ServerName: ""` bug makes `safeFetch` non-functional for HTTPS downloads.

---

### 1. TLS Finding — CONFIRMED BUG (NOT "NOT REPRODUCED")

The implementation report claims "NOT REPRODUCED — Go auto-supplies hostname from request URL when ServerName=''". The test `tls_test.go` **actually demonstrates the bug**:

```
=== RUN   TestTLSReproduction/CurrentImplementation_EmptyServerName
    tls_test.go:56: Request failed: Get "https://test.local:34091": 
    tls: failed to verify certificate: x509: certificate is valid for example.com, *.example.com, not test.local
```

**What happens:**
- `httptest.NewTLSServer` generates a self-signed cert for `example.com`
- Client uses custom `DialContext` connecting to `127.0.0.1:port` (validated IP)
- `TLSClientConfig.ServerName: ""` (empty)
- Go does **NOT** auto-supply the request hostname when using a custom `DialContext` that connects to an IP address
- TLS verification fails: cert is for `example.com`, connection is to `test.local` (SNI) / `127.0.0.1` (IP)

**Why this breaks production:**
- arXiv (`export.arxiv.org`) and OpenAlex (`api.openalex.org`) serve valid certs for their hostnames
- `safeFetch` connects via validated IP (e.g., `151.101.1.133` for arxiv.org)
- With empty `ServerName`, Go verifies cert against the IP → **fails**
- All HTTPS fetches will fail with `x509: certificate is valid for <hostname>, not <IP>`

**Fix required:** Set `ServerName` to the request hostname. Options:
1. Use `DialTLSContext` to set per-connection `ServerName`
2. Create per-request `http.Transport` with correct `ServerName`
3. Use `tls.Config.GetClientCertificate` / `BuildNameToCertificate` (complex)

---

### 2. Download Security — INCOMPLETE (Missing Symlink Protection)

| Protection | Status | Notes |
|------------|--------|-------|
| Redirect validation | ✅ | `makeCheckRedirect` validates each hop |
| DNS rebinding | ⚠️ | `DialContext` validates IPs, but TOCTOU between redirect check and dial |
| IPv4/IPv6 private blocks | ✅ | Comprehensive (RFC1918, link-local, loopback, CGNAT, TEST-NET, ULA) |
| PDF magic bytes (`%PDF`) | ✅ | Streaming verification |
| Content-Type sniffing | ✅ | Rejects HTML/JSON error pages |
| 50MB streaming limit | ✅ | `io.LimitedReader` |
| Paper ID path traversal | ✅ | `sanitizePaperID` rejects `..`, absolute paths |
| **Cache symlink escape** | ❌ **MISSING** | No `filepath.EvalSymlinks` + boundary check after `MkdirAll` |
| Temp file cleanup | ✅ | Atomic `.tmp` → rename |
| HTTPS enforcement | ✅ | Rejects non-HTTPS URLs and redirects |

**Symlink escape scenario:**
1. Attacker creates `paper-cache/2209.15487v1` as symlink to `/etc`
2. `Fetch()` calls `MkdirAll(dir)` → follows symlink, creates `/etc/v1/`
3. `os.WriteFile(tmpPath)` writes to `/etc/v1/source.pdf.tmp`
4. `os.Rename` moves to `/etc/v1/source.pdf`
5. Attacker controls file in `/etc`

**Fix:** After `MkdirAll`, verify `filepath.EvalSymlinks(dir)` is within `filepath.EvalSymlinks(paperCacheDir(root))`.

---

### 3. Cross-Process Rate Limiting — IMPLEMENTED, Needs Behavioral Test

**Implementation** (`persistent_ratelimit.go`):
- File-based state in `paper-cache/.ratelimit/{arxiv,openalex}.json`
- `flock` on `.lock` file for cross-process synchronization
- Conservative on corrupt state: waits full interval
- 24-hour staleness threshold for timestamps
- Auto-initialized via `initPersistentRateLimiters()` on first `Search()`

**Auto-initialization flow:**
```
Search() → initPersistentRateLimiters() → InitPersistentRateLimiters(cacheRoot)
```
- `cacheRoot` found via `PHYS_RESEARCH_ROOT` env or walking up from CWD to find `research/`

**Behavioral properties NOT yet verified by tests:**
- [ ] Two independent `physctl paper search` processes respect 3s interval
- [ ] State persists after process termination (survives `kill -9`)
- [ ] Corrupt JSON state → conservative wait (full interval), not immediate allow
- [ ] Interrupted `Wait()` (process crash during sleep) doesn't release early
- [ ] Lock file cleanup on crash (stale locks handled)
- [ ] OpenAlex 10 req/s (100ms) limit works across processes

**Gap:** The in-memory `RateLimiter` (ratelimit.go) still exists and is used by `arXivSearch`/`openAlexSearch` directly — but `Search()` calls `WaitForArXivRateLimit()` first. Need to verify the in-memory limiter doesn't conflict or allow bypass.

---

### 4. OpenCode Permissions — UNCHANGED, BYPASSABLE

**File:** `.opencode/opencode.json` — identical to original review.

**Bypass vectors (verified by inspection):**

| Attempt | Pattern Match | Result |
|---------|---------------|--------|
| `git merge accepted` | `git merge *` ✓ | Blocked |
| `git -c merge.ff=false merge accepted` | `git merge *` ✗ | **NOT BLOCKED** |
| `git push origin accepted` | `git push *` ✓ | Blocked |
| `git push --force origin accepted` | `git push *` ✓ | Blocked |
| `sed -i 's/deny/allow/' .opencode/opencode.json` | No shell deny for `sed` | **NOT BLOCKED** |
| `cat > .opencode/opencode.json <<'EOF'...` | No shell deny for `cat` | **NOT BLOCKED** |
| `tee .opencode/opencode.json` | No shell deny for `tee` | **NOT BLOCKED** |
| `mv .opencode/opencode.json.bak .opencode/opencode.json` | No shell deny for `mv` | **NOT BLOCKED** |

**Subagent isolation:** Unverified. Each agent has `edit: "*" deny` then specific allows. OpenCode's evaluation order (first-match-wins) and whether parent denies apply to subagents is undocumented behavior.

**Human acceptance boundary:** Still only pre-commit hook + `PHYS_HUMAN_MERGE` env var (explicitly documented as convenience, not security). No remote branch protection configured.

---

### 5. Regression Verification

| Command | Result |
|---------|--------|
| `gofmt -l .` | PASS (no output) |
| `go vet ./...` | PASS (no output) |
| `go build ./...` | PASS |
| `go list ./...` | 9 packages |
| `go test -count=1 ./...` | PASS (8 packages, 1 no test files) |
| `physctl validate` | PASS (0 errors, 1 `LEGACY_UNPINNED` warning) |
| `physctl status` | PASS (issues, obligations, blocking findings) |
| `physctl check` | PASS (5 allowlisted profiles) |
| `physctl refs H-001 --in` | PASS (graph traversal) |

**Scientific record integrity verified:**
- P-001, H-001, E-001, R-001, R-002 unchanged (byte-identical to prior review)
- Review freshness shared rule preserved (`LEGACY_UNPINNED` for R-001, fresh for R-002)
- `?` obligations preserved in H-001
- Plan 4 provenance: H-001 digest `ee747151...`, R-001 digest `7ee684ef...`

---

### 6. Final Disposition

#### Per-Finding Summary

| Finding | Disposition |
|---------|-------------|
| CR-001 TLS ServerName | **UNRESOLVED** (bug confirmed, breaks HTTPS) |
| CR-004 Download Security | **PARTIALLY FIXED** (missing symlink escape protection) |
| CR-005 Cross-Process Rate Limiting | **IMPLEMENTED** (behavioral verification needed) |
| CR-006 OpenCode Permissions | **UNRESOLVED** (unchanged, bypassable) |
| CR-003 ID Allocation | **FIXED** |
| CR-002 Frontmatter Tabs | **FIXED** |
| CR-007 Locator Validation | **FIXED** |
| CR-008 go.work Paths | **FIXED** |
| CR-009 Review Freshness | **FIXED** |

#### Explicit Answers

**1. Is direct external fetching safe?**
→ **NO.** The TLS `ServerName: ""` bug causes certificate verification failures for all legitimate HTTPS hosts. `physctl paper fetch` will fail on arXiv/OpenAlex.

**2. Is Researcher-only autonomous paper retrieval safe?**
→ **NO.** Even if TLS were fixed, missing symlink escape protection in the cache directory allows arbitrary file writes. Cross-process rate limiting needs behavioral verification before trusting it to respect arXiv's 3-second rule.

**3. Are OpenCode permissions sufficiently restricted?**
→ **NO.** CR-006 remains open. Shell patterns are string-prefix based and bypassable. No subagent shell denies. Subagent isolation unverified. Protected files (AGENTS.md, .opencode/*, cmd/*, go.mod, hooks/*) can be modified via `sed`, `tee`, `cat`, `mv` shell commands.

**4. What remaining risks require human action?**
- **Mandatory before any fetching:** Fix TLS `ServerName` (per-request transport or `DialTLSContext`); add symlink escape protection (`EvalSymlinks` + boundary check).
- **Mandatory before autonomous fetching:** Behaviorally verify cross-process rate limiting (test two `physctl paper search` processes 1s apart → second waits).
- **Mandatory for agent security:** Replace shell string-prefix patterns with regex or explicit deny lists; add explicit `shell: deny` for subagents; verify OpenCode permission inheritance model.
- **Recommended:** Add integration test for `physctl paper fetch` against a controlled HTTPS test server with valid cert.

---

### Recommendation

**Current authorization (unchanged from prior review):**

| Activity | Status |
|----------|--------|
| Local research-record development | ✅ Allowed |
| Bounded scholarly metadata discovery (`physctl paper search`) | ✅ Allowed (with rate limiter caveat) |
| PDF downloading (`physctl paper fetch`) | ❌ **HOLD** — TLS bug + symlink escape |
| Unrestricted autonomous external PDF fetching | ❌ **HOLD** |
| Claiming Researcher is securely isolated from protected files | ❌ **HOLD** |

**Next milestone:** Not another feature. Independent demonstration that:
1. `safeFetch` successfully downloads from `https://export.arxiv.org` and `https://api.openalex.org` with valid certs
2. Symlink escape attempt is blocked
3. Two sequential `physctl paper search` CLI invocations respect 3-second interval
4. Agent shell bypass attempts are blocked

The scientific provenance infrastructure (records, reviews, validation, status, refs, check) remains solid and well-tested.
