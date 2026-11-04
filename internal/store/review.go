package store

import (
	"fmt"
	"regexp"
	"strings"
)

// This file holds the single review-freshness rule. The validate and status
// commands both call it, so they can never disagree about whether a review is
// current: a changed target digest or a changed material-input digest is
// STALE in both (plan4.md, safeguard 6).

var sha256Hex = regexp.MustCompile(`^[0-9a-f]{64}$`)

// InputRef is one material-input digest recorded in a review body.
type InputRef struct {
	ID     string // record ID, e.g. P-001
	Digest string // as recorded, "sha256:<64 hex>"; "" when the line is malformed
	Line   int    // file-absolute line number of the input: line
	Err    string // non-empty when the line itself is malformed
}

// PinnedInput is a material input digest recorded when a review is created.
type PinnedInput struct {
	ID     string
	Digest string // SHA-256 hex of the record's exact bytes
}

// ReviewProvenance is what a review records about the versions it examined.
type ReviewProvenance struct {
	Reviewed     string // reviewed: <ID>
	ReviewedLine int
	Digest       string // digest: sha256:<hex> of the reviewed target
	DigestLine   int
	Inputs       []InputRef // input: <ID> sha256:<hex> lines, in file order
}

// ParseReview extracts the provenance block from a review record's body.
// Line numbers are file-absolute.
func ParseReview(r *Record) ReviewProvenance {
	var p ReviewProvenance
	for i, ln := range strings.Split(string(r.Body), "\n") {
		line := r.BodyLine + i
		t := strings.TrimSpace(ln)
		switch {
		case strings.HasPrefix(t, "reviewed:") && p.Reviewed == "":
			p.Reviewed = strings.TrimSpace(strings.TrimPrefix(t, "reviewed:"))
			p.ReviewedLine = line
		case strings.HasPrefix(t, "digest:") && p.Digest == "":
			p.Digest = strings.TrimSpace(strings.TrimPrefix(t, "digest:"))
			p.DigestLine = line
		case strings.HasPrefix(t, "input:"):
			rest := strings.TrimSpace(strings.TrimPrefix(t, "input:"))
			ref := InputRef{Line: line}
			fields := strings.Fields(rest)
			switch {
			case len(fields) == 2:
				ref.ID, ref.Digest = fields[0], fields[1]
				if !sha256Hex.MatchString(strings.TrimPrefix(ref.Digest, "sha256:")) ||
					!strings.HasPrefix(ref.Digest, "sha256:") {
					ref.Err = "input digest must be 'sha256:<64 hex>'"
					ref.Digest = ""
				}
			case len(fields) == 1:
				ref.ID = fields[0]
				ref.Err = "input line must be 'input: <ID> sha256:<64 hex>'"
			default:
				ref.Err = "input line must be 'input: <ID> sha256:<64 hex>'"
			}
			p.Inputs = append(p.Inputs, ref)
		}
	}
	return p
}

// ReviewIssueKind classifies a provenance problem.
type ReviewIssueKind string

const (
	IssueTargetMissing   ReviewIssueKind = "target-missing"
	IssueTargetMalformed ReviewIssueKind = "target-malformed"
	IssueTargetStale     ReviewIssueKind = "target-stale"
	IssueInputMalformed  ReviewIssueKind = "input-malformed"
	IssueInputMissing    ReviewIssueKind = "input-missing"
	IssueInputStale      ReviewIssueKind = "input-stale"
)

// ReviewIssue is one provenance problem. Line numbers are file-absolute.
type ReviewIssue struct {
	Kind ReviewIssueKind
	Line int    // the recorded line the problem was found on
	ID   string // target or input record ID
	Want string // digest as recorded by the review, "" when absent
	Have string // digest currently on disk, "" when the record does not resolve
	Msg  string // extra detail (e.g. why an input line is malformed)
}

// ReviewIssues compares recorded provenance against the current records.
// lookup resolves a record ID to its current state.
//
// It returns every problem it finds, plus unpinned, which reports a review
// that records no material inputs at all. An unpinned review is a valid
// historical review of the exact target version it examined whose original
// source context was not pinned; it is NOT stale and must never be described
// as fully fresh.
func (p ReviewProvenance) ReviewIssues(lookup func(id string) (*Record, bool)) (issues []ReviewIssue, unpinned bool) {
	unpinned = len(p.Inputs) == 0

	targetDigestOK := strings.HasPrefix(p.Digest, "sha256:") &&
		sha256Hex.MatchString(strings.TrimPrefix(p.Digest, "sha256:"))
	if !targetDigestOK {
		issues = append(issues, ReviewIssue{
			Kind: IssueTargetMalformed, Line: p.DigestLine, ID: p.Reviewed, Want: p.Digest,
		})
	}

	if p.Reviewed != "" {
		target, ok := lookup(p.Reviewed)
		switch {
		case !ok:
			issues = append(issues, ReviewIssue{
				Kind: IssueTargetMissing, Line: p.ReviewedLine, ID: p.Reviewed, Want: p.Digest,
			})
		case targetDigestOK && p.Digest != "sha256:"+target.Digest:
			issues = append(issues, ReviewIssue{
				Kind: IssueTargetStale, Line: p.DigestLine, ID: p.Reviewed,
				Want: p.Digest, Have: target.Digest,
			})
		}
	}

	for _, in := range p.Inputs {
		if in.Err != "" {
			issues = append(issues, ReviewIssue{
				Kind: IssueInputMalformed, Line: in.Line, ID: in.ID, Msg: in.Err,
			})
			continue
		}
		rec, ok := lookup(in.ID)
		if !ok {
			issues = append(issues, ReviewIssue{
				Kind: IssueInputMissing, Line: in.Line, ID: in.ID, Want: in.Digest,
			})
			continue
		}
		if in.Digest != "sha256:"+rec.Digest {
			issues = append(issues, ReviewIssue{
				Kind: IssueInputStale, Line: in.Line, ID: in.ID,
				Want: in.Digest, Have: rec.Digest,
			})
		}
	}
	return issues, unpinned
}

// Reason renders a human-readable staleness reason for status output.
func (i ReviewIssue) Reason() string {
	switch i.Kind {
	case IssueTargetMissing:
		return "reviewed target no longer exists"
	case IssueTargetMalformed:
		return "review does not record a well-formed target digest"
	case IssueTargetStale:
		return "target changed since review (digest mismatch)"
	case IssueInputMalformed:
		return fmt.Sprintf("material input %s is recorded malformed", i.ID)
	case IssueInputMissing:
		return fmt.Sprintf("material input %s no longer exists", i.ID)
	case IssueInputStale:
		return fmt.Sprintf("material input %s changed since review (digest mismatch)", i.ID)
	}
	return "unknown provenance problem"
}

// Message renders the validation message for this issue, preserving the
// wording validate has always reported for the target checks.
func (i ReviewIssue) Message() string {
	switch i.Kind {
	case IssueTargetMissing:
		return fmt.Sprintf("reviewed target %s not found", i.ID)
	case IssueTargetMalformed:
		return "review must record 'digest: sha256:<64 hex>' of the reviewed version"
	case IssueTargetStale:
		return fmt.Sprintf("STALE: reviewed digest %s does not match current %s (sha256:%s); write a new review",
			i.Want, i.ID, i.Have)
	case IssueInputMalformed:
		return fmt.Sprintf("malformed material input line: %s", i.Msg)
	case IssueInputMissing:
		return fmt.Sprintf("STALE: material input %s recorded in this review no longer exists; write a new review", i.ID)
	case IssueInputStale:
		return fmt.Sprintf("STALE: material input %s digest %s does not match current sha256:%s; write a new review",
			i.ID, i.Want, i.Have)
	}
	return "unknown provenance problem"
}

// Rule returns the validate rule name this issue is reported under.
func (i ReviewIssue) Rule() string {
	switch i.Kind {
	case IssueTargetStale, IssueInputMissing, IssueInputStale:
		return "stale"
	}
	return "review"
}

// LegacyUnpinnedWarning is the validate warning for a review that pins no
// material inputs. It is a warning, never a clean result.
const LegacyUnpinnedWarning = "LEGACY_UNPINNED: review records only the target digest; material input versions are not pinned"
