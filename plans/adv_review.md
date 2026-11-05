## ADVERSARIAL CODE REVIEW — phys-research

### Executive Assessment

**Overall Engineering Condition:** The codebase is well-structured, passes all tests, and follows Go best practices. The restricted frontmatter parser correctly rejects unsupported YAML constructs. The review provenance system (pinned material inputs, shared freshness rule between `validate` and `status`) is soundly implemented. The OpenCode agent permissions are reasonably designed.

**Most Significant Risks:**
1. **Critical SSRF vulnerability** in paper download (`safeFetch` does not validate redirect targets)
2. **PDF content verification missing** — HTML error pages could be saved as PDFs
3. **ID allocation race condition** under high concurrency (theoretical but real)
4. **OpenCode shell permission patterns** may not match all dangerous command variations
5. **arXiv/OpenAlex API rate limiting completely absent** — will cause IP bans in production

**Disposition: READY WITH FINDINGS** — Continued Researcher-only testing is acceptable with the documented risks. The Critical SSRF (CR-001/CR-004) must be fixed before any autonomous paper fetching runs.

---

### Ranked Findings

#### CR-001: SSRF via Redirect Host Validation Bypass (Critical)
**Where:** `internal/paper/paper.go:81-135` (`safeFetch`)
**What:** Private-network/localhost host check only runs on initial request, not on redirects. An attacker controlling a mirror could redirect to `http://169.254.169.254/latest/meta-data/` (AWS metadata) or `http://localhost:port`.
**Fix:** Move host validation into `CheckRedirect` callback; add CIDR checks for RFC1918, link-local, loopback.
**Test:** Mock server redirecting to `127.0.0.1` → verify rejection.

#### CR-004: Multiple Paper Download Safety Failures (Critical)
**Where:** `internal/paper/paper.go`, `internal/paper/fetch.go`
**What:** 
- No DNS rebinding protection (DNS re-resolved on redirects)
- No PDF magic byte verification (`%PDF-` header)
- `paperID` used directly in `filepath.Join` — path traversal via `../../../etc/passwd`
- Content-Type check accepts empty/`binary` — HTML error pages saved as PDFs
- 50MB limit enforced during read, not before download starts
**Fix:** Add PDF magic check; sanitize `paperID` (hash or strict regex); enforce Content-Type; add `DialContext` with fixed IPs; verify `Content-Length` before read.
**Test:** Mock server returning HTML with `Content-Type: application/pdf`; paperID with `../`; redirect to localhost.

#### CR-005: External API Reliability Failures (High)
**Where:** `internal/paper/search.go`, `internal/paper/openalex.go`
**What:** 
- No rate limiting for arXiv (required: 1 req/3s, email in UA)
- OpenAlex `apiKey` parameter accepted but never used
- No pagination — silent truncation
- No retry on 429/503
- `httpQuery` function defined but unused
- DOI normalization not implemented
**Fix:** Token bucket rate limiter; use `apiKey` in requests; implement pagination; add retry with backoff for 429/503; DOI normalization function.
**Test:** Mock 429 response → verify retry; pagination aggregation.

#### CR-003: ID Allocation Race Condition (High)
**Where:** `internal/record/record.go:160-177` (`Create`)
**What:** Fixed 1000-iteration retry loop with `os.Link` + `O_EXCL`. Under high concurrency, two processes can both scan for next ID, get same value, first wins, second fails with `ErrExist`, retries... 1000 limit could exhaust.
**Fix:** Remove arbitrary limit; add exponential backoff; or use directory lock file (`flock`).
**Test:** 50 concurrent `physctl new` goroutines → all succeed with distinct IDs.

#### CR-006: OpenCode Agent Permission Bypass Vectors (High)
**Where:** `.opencode/opencode.json`
**What:** 
- Shell deny patterns (`git merge *`, `git push *`) are string prefixes — may not match `git -c merge.ff=false merge accepted`
- No explicit `shell` deny rules for subagents (inherit undefined behavior)
- Subagent permission isolation not clearly documented — parent denies may not apply
- Local `accepted` branch protection is only hook + env var (admittedly not a security boundary)
**Fix:** Document OpenCode permission evaluation order; add regex-based shell patterns; explicit subagent shell denies; verify top-level denies propagate.
**Test:** Attempt `git checkout accepted`, `git merge accepted`, `sed -i .opencode/opencode.json` from researcher agent.

#### CR-002: Frontmatter Parser Allows Tabs in Quoted Values (Medium)
**Where:** `internal/frontmatter/frontmatter.go:109-111, 265-283`
**What:** Tab check runs on raw line (tabs inside quotes pass), then `cleanValue` strips quotes but doesn't re-check for tabs. Value `"foo\tbar"` becomes `foo\tbar` in parsed header.
**Fix:** In `cleanValue`, after quote stripping, reject values containing `\t`.
**Test:** `id: "P-001"`, `state: "o\tpen"` → reject.

#### CR-007: Source Locator Validation Too Permissive (Medium)
**Where:** `internal/validate/validate.go:313-375` (`validateSource`)
**What:** `valueOK` rejects placeholders but accepts `locator: p.`, `locator: Eq.`, `locator: Fig.` — semantically empty but structurally valid.
**Fix:** Require locator to contain a digit (page/eq/figure/table number) via regex.
**Test:** Claims with `locator: p.` → reject.

#### CR-008: go.work Parsing Fragility (Medium)
**Where:** `internal/check/check.go:202-271` (`resolveModule`)
**What:** Regex-only parsing ignores `replace`/`toolchain` directives; no boundary check — `use ../../../etc` escapes workspace; no syntax validation of `go.work`.
**Fix:** Use `golang.org/x/mod/modfile` for parsing; verify resolved path under workspace root; validate `go` directive present.
**Test:** `go.work` with `use ../../../etc` → reject; malformed `go.work` → clear error.

#### CR-009: Review Freshness Logic Edge Case (Medium)
**Where:** `internal/store/review.go:107-155` (`ReviewIssues`)
**What:** Review with malformed `input:` lines (e.g., `input: P-001` missing digest) has `len(Inputs) > 0` → not `LEGACY_UNPINNED`, but generates "review" errors for malformed inputs. Inconsistent: legacy review (no inputs) = warning; review with broken inputs = errors.
**Fix:** Document intent clearly; consider treating malformed inputs as "effectively unpinned" (warning).
**Test:** Review with `input: P-001` → verify error type and no `LEGACY_UNPINNED`.

---

### Test Results

| Command | Result |
|---------|--------|
| `gofmt -l .` | PASS (no output) |
| `go vet ./...` | PASS (no output) |
| `go build ./...` | PASS |
| `go test -count=1 ./...` | PASS (all 8 packages) |
| `physctl validate` | PASS (1 warning: `LEGACY_UNPINNED` for R-001) |
| `physctl status` | PASS (reports issues, obligations, stale reviews, blocking findings) |
| `physctl refs H-001 --in` | PASS (traverses graph correctly) |
| `physctl check` | PASS (lists 5 allowlisted profiles) |

**Tests not executable:** 
- `physctl paper search/fetch/extract` — require network access and `pdftotext` binary (not installed in review environment). These were not tested live per review constraints.

---

### Security Assessment

**Agent Permission Boundaries:**
- Researcher: Can edit `research/sources/*`, `issues/*`, `hypotheses/*`, `evidence/*`. Denied on `*`, `AGENTS.md`, `.opencode/*`, `cmd/*`, `go.mod`, `go.work`, `hooks/*`. **Gap:** No explicit shell permissions defined — inheritance unclear.
- Adversary: Can edit only `research/reviews/*`. Denied on `*`. **Gap:** Same shell permission gap.
- Formalizer: Can edit only `derivations/*`. Denied on `*`. **Gap:** Same shell permission gap.
- Top-level: Denies `git merge *`, `git push *`, `git config *` via shell patterns. **Gap:** Patterns are string prefixes, not semantic command matching.

**Download Safety:** 
- **CRITICAL FAILURES:** SSRF via redirects (CR-001), no PDF verification (CR-004), path traversal in cache dir (CR-004), no DNS rebinding protection (CR-004).
- **Mitigations present:** 50MB size limit, 5-redirect hop limit, User-Agent identification, localhost block on initial request only.

**Human Acceptance Protection:**
- Remote branch protection on `accepted` with credentials agents lack (documented as real boundary)
- Local pre-commit hook + `PHYS_HUMAN_MERGE` env var (explicitly documented as convenience only)
- No local enforcement if no protected remote exists

**Remaining Trust Assumptions:**
- `pdftotext` binary is safe (no argument injection — `exec.Command` used correctly with args array)
- `go` toolchain in PATH is trusted
- Paper cache directory is not shared with untrusted processes
- OpenCode permission evaluation matches documented behavior

---

### Scientific-Integrity Assessment

**Source Provenance:** 
- ✅ Every claim requires `locator:` (enforced by validation)
- ⚠️ Locator format not validated for semantic content (CR-007)
- ✅ SHA-256 of exact PDF recorded in source record
- ✅ Paper version (arXiv v1, v2, etc.) tracked in cache and source record

**Hypothesis/Review Freshness:**
- ✅ Single shared freshness rule in `store/review.go` used by both `validate` and `status` (safeguard 6)
- ✅ Material inputs pinned at review creation (one-hop neighbors, both directions, reviews excluded)
- ✅ `LEGACY_UNPINNED` warning for pre-pinning reviews (R-001)
- ✅ Stale detection names the changed record (target or specific material input)
- ✅ Reviews are append-only (validated via git history + working tree check)

**Research Claim Boundaries:**
- ✅ Researcher cannot write reviews (enforced by permissions)
- ✅ Researcher cannot declare established (enforced by AGENTS.md + agent instructions)
- ✅ `?` obligations preserved (reported by status, never errors)
- ✅ Validation disclaimer: "Structural validation only: no mathematical or physical correctness is implied"
- ✅ Check profiles allowlisted, `UNAVAILABLE` for missing capabilities
- ⚠️ Researcher agent instructions permit `websearch`/`webfetch` — could retrieve unverified content

**Competing Hypotheses:** 
- Multiple hypotheses can exist (H-001, H-002) with independent reviews (R-001, R-002)
- No mechanism to silently collapse them — each reviewed independently

---

### Recommended Remediation Order

**Mandatory (before autonomous paper fetching):**
1. **CR-001/CR-004:** Fix SSRF in `safeFetch` (redirect host validation, DNS rebinding protection, PDF magic bytes, paperID sanitization, Content-Type enforcement)
2. **CR-005:** Add arXiv rate limiting and OpenAlex API key usage

**High Priority:**
3. **CR-003:** Fix ID allocation race (remove 1000 limit, add backoff)
4. **CR-006:** Harden OpenCode shell permissions (regex patterns, explicit subagent denies)

**Medium Priority:**
5. **CR-002:** Reject tabs in quoted frontmatter values
6. **CR-007:** Validate locator semantic content (require digit)
7. **CR-008:** Use proper `go.work` parser (`golang.org/x/mod/modfile`)
8. **CR-009:** Document/clarify malformed-input review behavior

---

### Overall Disposition

**READY WITH FINDINGS** — The implementation is structurally sound for Researcher-only testing with the following conditions:
1. **Do not run `physctl paper fetch`** until CR-001/CR-004 are fixed (SSRF + download safety)
2. **Monitor arXiv/OpenAlex usage** — current code will get IP banned (CR-005)
3. **Verify OpenCode permission behavior** empirically before trusting agent isolation (CR-006)

The review provenance system, frontmatter parser, and validation logic are correctly implemented and well-tested. The scientific integrity mechanisms (freshness, append-only reviews, `?` obligations, disclaimer) work as designed.
