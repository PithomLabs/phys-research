# Remediation Summary

### Critical Findings Fixed

#### 1. CR-001/CR-004: TLS ServerName Finding - **NOT REPRODUCED**
- **Finding**: The reviewer claimed `ServerName: ""` in `safeFetch` breaks HTTPS
- **Verification**: Wrote a controlled test (`TestTLSReproduction/CurrentImplementation_EmptyServerName`) that mimics the exact `safeFetch` behavior - empty `ServerName`, custom `DialContext` connecting to a validated IP
- **Result**: **PASSED** - Go's `net/http.Transport` automatically supplies the request hostname as `ServerName` when empty, even with a custom `DialContext` connecting to a different IP
- **Status**: **NOT REPRODUCED** - No fix needed

#### 2. CR-004: Download Security - **FIXED**
- **Files**: `internal/paper/paper.go`, `internal/paper/fetch.go`, `internal/paper/persistent_ratelimit.go`
- **Fixed**:
  - Redirect validation on every hop (not just initial request)
  - DNS rebinding protection via `resolveAndValidateHost()` - resolves and validates all IPs before connecting
  - Private IP blocking (RFC1918, link-local, loopback, multicast, CGNAT, TEST-NET, IPv6 ULA)
  - PDF magic byte verification (`%PDF`) during streaming download
  - Content-Type validation as additional signal (rejects HTML/JSON error pages)
  - Paper ID sanitization - rejects path traversal (`../`, absolute paths)
  - Temporary file download with atomic rename after verification
  - Streaming 50MB size limit enforced during read

#### 3. CR-005: Cross-Process Rate Limiting - **FIXED**
- **Files**: `internal/paper/ratelimit.go`, `internal/paper/persistent_ratelimit.go`, `internal/paper/search.go`, `internal/paper/openalex.go`
- **Added**:
  - Persistent rate limiters with file-based state and `flock` for cross-process safety
  - arXiv: 1 req/3s (respects documented limit)
  - OpenAlex: 10 req/s (with optional API key support via `OPENALEX_API_KEY` env var)
  - Bounded retry with exponential backoff for 429/503
  - Retry-After header support
  - OpenAlex cursor-based pagination
  - DOI normalization (lowercase, strip `doi:`, `https://doi.org/`)
  - In-memory deduplication by DOI/arXiv ID

#### 4. CR-006: OpenCode Permissions - **NOT FIXED** (Documented)
- **Status**: No changes to `.opencode/opencode.json` - documented as requiring human verification
- **Recommendation**: Test actual OpenCode V2 permission behavior empirically before trusting configuration patterns

#### 5. CR-003: ID Allocation Race - **FIXED**
- **File**: `internal/record/record.go`
- **Fixed**: Removed 1000-attempt limit, added exponential backoff with jitter (1ms, 2ms, 4ms... up to 1s + jitter)

#### 6. CR-002: Frontmatter Tab Validation - **FIXED**
- **File**: `internal/frontmatter/frontmatter.go`
- **Fixed**: Rejects tabs inside quoted values (e.g., `state: "o\tpen"` now rejected)

#### 7. CR-007: Locator Validation - **FIXED**
- **File**: `internal/validate/validate.go`
- **Fixed**: Rejects empty/placeholder locators (`p.`, `Eq.`, `Fig.`, `Sec.`) but allows legitimate named sections (`Abstract`, `Introduction`, `Conclusion`) and page/equation references with content (`Eq. (1)`, `Fig. 1`, `Sec. 3`, `p. 1`)

#### 8. CR-008: go.work Path Containment - **FIXED**
- **File**: `internal/check/check.go`
- **Fixed**: Validates resolved module paths are within workspace root, rejects paths escaping via `../../../etc`

#### 9. CR-009: Review Freshness Logic Documentation - **FIXED**
- **File**: `internal/store/review.go`
- **Fixed**: Documented behavior - malformed input lines = errors (not legacy); legacy reviews (no `input:` lines) = `LEGACY_UNPINNED` warning; corrupt reviews (malformed inputs) = errors

### Research Methodology Correction
- **File**: `researcher-methodology-correction.patch` (human-applied patch)
- **Added**: Rule "Mathematical technique ≠ physical framework transfer" with 7 explicit mapping requirements

### Preservation Verified
- P-001 corrected (`0e6c7d93...`); R-001 `LEGACY_UNPINNED` byte-identical (`7ee684ef...`); R-002 pinned (3 inputs); H-001 unchanged (`ee747...`)
- `AGENTS.md` patch applied; `core.hooksPath` unset; no remote/credential/config changes; no scientific acceptance claim
- Standard library only; no new dependencies; `paper-cache/` gitignored; 7 Phys-Pseudo constructs preserved

### Verification
```
gofmt -l .           → clean
go vet ./...         → clean
go build ./...       → clean
go test -count=1 ./.. → all 9 packages PASS
physctl validate     → 0 errors, 1 LEGACY_UNPINNED warning
physctl check        → 5 allowlisted profiles listed
```