# Targeted Remediation Plan — phys-research

Based on independent review `adv_review2.md` findings.

## Current Disposition (from independent review)

| Finding | Disposition | Status |
|---------|-------------|--------|
| CR-001/CR-004 | TLS ServerName empty breaks HTTPS | **PARTIALLY FIXED** - Critical blocker |
| CR-004 | Symlink escape in paper-cache | **PARTIALLY FIXED** - Gap remains |
| CR-005 | Cross-process rate limiting | **PARTIALLY FIXED** - In-memory only |
| CR-006 | OpenCode permissions | **NOT FIXED** |
| CR-003 | ID allocation race | FIXED |
| CR-002 | Frontmatter tabs | FIXED |
| CR-007 | Locator validation | FIXED |
| CR-008 | go.work path containment | FIXED |
| CR-009 | Review freshness docs | FIXED |

## Phase 1: Reproduce TLS Finding (Priority 1)

**Objective**: Verify if `ServerName: ""` actually breaks HTTPS in Go's net/http.

### Test Plan:
1. Create a controlled HTTPS test server with a valid hostname certificate (e.g., using `httptest.NewTLSServer` or a local test server with a self-signed cert for a specific hostname).
2. Implement the current `safeFetch` logic with `ServerName: ""` and a custom `DialContext` that connects to a validated IP.
3. Make an HTTPS request using the original hostname.
4. Observe if certificate verification fails.

**Expected outcomes**:
- If Go's net/http correctly supplies the request hostname as ServerName when empty → Finding is **NOT REPRODUCED**.
- If certificate verification fails with "certificate not valid for IP address" or similar → Finding **CONFIRMED**.

**Decision**: Only fix if reproduced. Do not blindly apply fixes.

## Phase 2: Cross-Process Rate Limiting (CR-005)

**Objective**: Persistent rate limiting across CLI invocations.

### Design:
- File-based token bucket with timestamps (JSON or plain text)
- Location: `paper-cache/.ratelimit/` (within existing cache dir)
- Separate files per endpoint: `arxiv.json`, `openalex.json`
- Format: `{ "tokens": float, "last_refill": timestamp }`
- Use file locking (`flock`/`fcntl`) for cross-process safety
- Standard library only: `os`, `sync`, `time`, `encoding/json`, `syscall` for `flock`

### Implementation:
1. `internal/paper/ratelimit.go` - add `PersistentRateLimiter` with file persistence
2. Persist state on every token consumption
3. On startup, load state; if file missing/corrupt, start fresh
4. Handle interrupted processes: on read, if timestamp is stale (> 24h), reset
5. Test concurrent processes: run multiple `physctl paper search` simultaneously, verify rate limit respected

### OpenAlex:
- Use same persistent mechanism
- Apply API key when configured (env var `OPENALEX_API_KEY`)

## Phase 3: OpenCode Permission Hardening (CR-006)

**Objective**: Harden `.opencode/opencode.json` permissions.

### Analysis Required:
1. Inspect actual OpenCode V2 permission evaluation semantics (read docs or test)
2. Test actual command execution in disposable checkout

### Changes to `.opencode/opencode.json`:
1. Replace string-prefix shell patterns with regex-based (if supported) or add explicit deny patterns for variations:
   - `git merge *` → also deny `git -c * merge *`, `git merge --*`
   - `git push *` → also deny `git -c * push *`
   - `git config *` → already covers variations
2. Add explicit `shell` deny rules for subagents (inherit from parent + explicit allows)
3. Explicitly deny `sed`, `cat`, `tee`, `mv`, `rm`, `chmod` on protected paths
3. Verify subagent permission isolation: test that researcher cannot run `git merge`, adversary cannot edit `research/sources/`, etc.

### Testing:
- Create disposable checkout
- Spawn each subagent type
- Attempt forbidden commands (Git variations, file edits, config changes)
- Verify denials work

**Note**: Do not claim accepted-branch protection exists — no remote configured.

## Phase 4: Cache Symlink Protection (CR-004)

**Objective**: Prevent symlink-based cache escapes.

### Implementation:
1. In `fetch.go` and `paper.go` cache path resolution:
   - After `filepath.Join(root, paperID, version)`, call `filepath.EvalSymlinks`
   - Verify result is within `paper-cache/` root using `filepath.Rel` + `..` check
2. Apply to:
   - `paperDirFor` (before `MkdirAll`)
   - `ManifestPath`, `SourcePDFPath`, `FullTextPath`
   - `Fetch` (before `MkdirAll` and write)
2. Add test: create symlink in cache dir pointing outside, verify fetch fails

## Phase 5: TLS ServerName Fix (CR-001/CR-004)

**Contingent on Phase 1 reproduction**.

If confirmed:
- Use `DialTLSContext` or per-request `tls.Config` with `ServerName` set to request hostname
- In `safeFetch`, create per-request `tls.Config` with `ServerName: parsedURL.Host`
- Or use `DialTLSContext` to set ServerName per connection

## Phase 6: Cache Symlink Protection Tests

Add tests in `internal/paper/*_test.go`:
1. Symlink escape attempt (symlink in cache dir → outside)
2. Redirect to localhost/private IP
3. HTML error page masquerading as PDF
4. Cross-process rate limiter concurrent access

## Verification Checklist

After all fixes:
- `gofmt -l .` → clean
- `go vet ./...` → clean
- `go build ./...` → clean
- `go test -count=1 ./...` → all pass
- `physctl validate` → 0 errors, 1 LEGACY_UNPINNED warning
- `physctl check` → profiles listed
- Live download test with controlled HTTPS server
- Concurrent rate limiter test (5 parallel processes)

## Disposition Reporting

After remediation, report each CR as:
- **FIXED** - verified with tests
- **NOT REPRODUCED** - TLS issue not reproduced
- **PARTIALLY FIXED** - significant gaps remain
- **UNRESOLVED** - not addressed

Request independent review after completion before autonomous operation.

---

**Next Step**: Begin Phase 1 (TLS reproduction test). Need clarification on whether to proceed with implementation or if the TLS test should be written first as a standalone verification.