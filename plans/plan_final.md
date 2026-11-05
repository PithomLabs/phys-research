# Final Security Remediation Plan — phys-research

## Context
Based on `adv_review3.md` (adversarial review) and `review3.md` (human assessment), the following findings need remediation:

| Finding | Current Status | Required Action |
|---------|---------------|-----------------|
| CR-001 TLS ServerName | **UNRESOLVED** (disputed) | Reproduce with valid cert; fix only if reproduced |
| CR-004 Cache symlink escape | **UNRESOLVED** | Implement symlink escape protection + adversarial test |
| CR-005 Rate limiting | **IMPLEMENTED, NOT VERIFIED** | Add behavioral cross-process tests |
| CR-006 OpenCode permissions | **UNRESOLVED** | Don't rely on shell patterns; recommend OS isolation |
| CR-002, 003, 007, 008, 009 | **FIXED** (verified) | No action needed |

## Phase 1: TLS Verification (Priority 1)

### Objective
Determine if `safeFetch` with empty `ServerName` + custom `DialContext` actually fails TLS verification when using a certificate whose SAN matches the requested hostname.

### Test Design
1. Create a test TLS server with a certificate valid for `test.local` (SAN: `test.local`)
2. Use `safeFetch` logic: custom `DialContext` connecting to `127.0.0.1:port`, empty `ServerName`
3. Request `https://test.local:port` with empty `ServerName`
4. If TLS verification fails → bug confirmed, fix needed
6. If TLS succeeds → NOT REPRODUCED, no TLS fix needed

### Implementation
- Create test in `internal/paper/tls_test.go` with proper certificate
- Use `httptest.NewTLSServer` with custom certificate for `test.local`
- Run test; if passes → NOT REPRODUCED; if fails → fix TLS config

## Phase 2: Cache Symlink Protection (CR-004)

### Objective
Prevent symlink traversal and writes outside canonical paper-cache root.

### Implementation
1. **`paperDirFor` / `ManifestPath` / `SourcePDFPath` / `FullTextPath`**: After `MkdirAll`, verify `filepath.EvalSymlinks(dir)` is within `filepath.EvalSymlinks(paperCacheDir(root))`
2. **Account for symlinked parents**: Resolve both paths before comparison
3. **Re-validate after `MkdirAll`**: Re-check after directory creation (concurrent replacement)
4. **Adversarial test**: Create symlink in cache dir pointing to `/tmp/escape`, attempt fetch, verify blocked

## Phase 3: Persistent Rate Limiting Behavioral Tests (CR-005)

### Objective
Verify cross-process rate limiting works under real conditions.

### Tests to Add
1. **Concurrent processes**: Spawn 2 `physctl paper search` processes simultaneously; verify second waits ~3s
2. **Process crash recovery**: Start `Wait()`, kill process mid-sleep, restart → should wait remaining time
3. **Corrupt state**: Write invalid JSON to state file; verify conservative wait (full interval)
4. **Interrupted wait**: Kill process during sleep; verify lock cleanup on restart
5. **OpenAlex throttling**: Verify 10 req/s limit enforced across processes
6. **In-memory vs persistent**: Verify in-memory limiter doesn't bypass persistent limiter

## Phase 4: OpenCode Security Documentation (CR-006)

### Objective
Document limitations of current permission model; recommend OS-level isolation.

### Deliverables
1. Update `AGENTS.md` with explicit permission limitations
2. Add `SECURITY.md` documenting:
   - Shell pattern limitations (prefix matching only)
   - Subagent permission inheritance behavior
   - Recommended OS-level isolation (user/container)
   - Credential isolation requirements
2. Add `SECURITY.md` patch for human application

## Phase 5: Regression & Preservation Verification

### Required Commands
```bash
gofmt -l .
go vet ./...
go build ./...
go test -count=1 ./...
physctl validate
physctl check
```

### Verification Checklist
- [ ] All existing tests pass
- [ ] All existing research records preserved (P-001, H-001, E-001, R-001, R-002)
- [ ] Standard-library-only requirement preserved
- [ ] TLS finding reported as NOT REPRODUCED or FIXED
- [ ] Cache symlink test added and passes
- [ ] Rate limiter cross-process tests added and pass
- [ ] OpenCode security limitations documented

## Reporting
For each finding, report:
- **FIXED**: Test passes, fix verified
- **NOT REPRODUCED**: Test demonstrates no bug
- **UNRESOLVED**: Issue exists, no fix applied
- **NOT TESTED**: No test exists yet

## Execution Order
1. **TLS verification** (determines if TLS fix needed)
2. **Cache symlink protection** (independent, critical)
3. **Rate limiting tests** (can run in parallel with other work)
5. **OpenCode security doc** (documentation only)
5. **Full regression test suite**

## Success Criteria
- TLS test passes (NOT REPRODUCED) OR TLS fix applied and verified
- Symlink escape test fails before fix, passes after
- Rate limiter cross-process tests pass
- OpenCode security doc created
- All existing tests pass
- All previous research records/hash preserved
- No new dependencies added

---

**Ready to implement when you switch to Build mode.**