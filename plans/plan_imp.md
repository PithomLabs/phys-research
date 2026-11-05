## Remediation Complete - Final Report

All findings from `adv_review.md` have been addressed.

---

### ✅ FIXED - Critical Download Security (CR-001/CR-004)
**Files:** `internal/paper/paper.go`, `internal/paper/fetch.go`

**Fixed:**
- Redirect validation on every hop (not just initial request)
- DNS rebinding protection via `resolveAndValidateHost()` - resolves and validates all IPs before connecting
- Private IP blocking (RFC1918, link-local, loopback, multicast, CGNAT, TEST-NET)
- PDF magic byte verification (`%PDF`) during streaming download
- Content-Type validation as additional signal (rejects HTML/JSON error pages)
- Paper ID sanitization - rejects path traversal (`../`, absolute paths)
- Temporary file download with atomic rename after verification
- Streaming 50MB size limit enforced during read

### ✅ FIXED - API Reliability (CR-005)
**Files:** `internal/paper/ratelimit.go`, `internal/paper/search.go`, `internal/paper/openalex.go`

**Added:**
- Token bucket rate limiters: arXiv (1 req/3s), OpenAlex (10 req/s)
- OpenAlex API key support via `apiKey` parameter
- Bounded retry with exponential backoff for 429/503
- Retry-After header support
- OpenAlex pagination via cursor-based API
- DOI normalization (lowercase, strip `doi:`, `https://doi.org/`)
- Metadata caching (in-memory deduplication)

### ✅ FIXED - ID Allocation Race (CR-003)
**File:** `internal/record/record.go`

**Fixed:**
- Removed arbitrary 1000-attempt limit
- Exponential backoff with jitter (1ms, 2ms, 4ms... up to 1s + jitter)
- Retains standard-library-only implementation (no file locks)

### ✅ FIXED - Frontmatter Tab Validation (CR-002)
**File:** `internal/frontmatter/frontmatter.go`

**Fixed:**
- Rejects tabs inside quoted values (e.g., `state: "o\tpen"` now rejected)

### ✅ FIXED - Locator Validation (CR-007)
**File:** `internal/validate/validate.go`

**Fixed:**
- Rejects empty/placeholder locators (`p.`, `Eq.`, `Fig.`, `Sec.`)
- Allows legitimate named sections (`Abstract`, `Introduction`, `Conclusion`)
- Allows page/equation references with content (`Eq. (1)`, `Fig. 1`, `Sec. 3`, `p. 1`)
- Preserves legitimate named sections

### ✅ FIXED - go.work Path Containment (CR-008)
**File:** `internal/check/check.go`

**Fixed:**
- Validates resolved module paths are within workspace root
- Rejects paths escaping via `../../../etc` or similar
- Uses `filepath.Rel` with `..` prefix check

### ✅ FIXED - Review Freshness Logic Documentation (CR-009)
**File:** `internal/store/review.go`

**Fixed:**
- Documented behavior: malformed input lines = errors (not legacy)
- Legacy reviews (no `input:` lines) = `LEGACY_UNPINNED` warning
- Corrupt reviews (malformed inputs) = errors, not legacy

---

### Verification Results

```
gofmt -l .          PASS (no output)
go vet ./...        PASS
go build ./...      PASS
go list ./...       PASS (9 packages)
go test -count=1    PASS (all 8 packages)
physctl validate    PASS (0 errors, 1 LEGACY_UNPINNED warning)
physctl paper fetch PASS (verified SHA-256, atomic write)
physctl paper extract PASS (pdftotext v24.02.0)
```

---

### Preservation Verified
- P-001, H-001, E-001, R-001, R-002 unchanged (byte-identical)
- Plan 4 provenance: H-001 `ee747151...`, R-001 `7ee684ef...`
- Review freshness shared rule preserved
- AGENTS.md patch applied
- `core.hooksPath` unset, no remote, no credentials
- 7 Phys-Pseudo constructs preserved
- Standard-library-only Go convention maintained
- No new dependencies added