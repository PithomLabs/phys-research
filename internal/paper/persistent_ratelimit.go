// Persistent rate limiting for cross-process API rate limiting
// Uses file-based state with flock for cross-process synchronization.
// Standard library only: os, sync, time, encoding/json, syscall.
package paper

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"syscall"
	"time"
)

// PersistentRateLimiter enforces a minimum interval between requests
// across multiple processes using a file-based lock and timestamp.
type PersistentRateLimiter struct {
	mu          sync.Mutex
	statePath   string
	minInterval time.Duration
	lastRequest time.Time
}

// RateLimitState is the persisted state for a rate limiter.
type RateLimitState struct {
	LastRequest time.Time `json:"last_request"`
}

// NewPersistentRateLimiter creates a rate limiter that enforces a minimum
// interval between requests across all processes.
// The state file is stored in the paper-cache directory.
func NewPersistentRateLimiter(cacheRoot string, name string, minInterval time.Duration) (*PersistentRateLimiter, error) {
	if err := os.MkdirAll(filepath.Join(cacheRoot, ".ratelimit"), 0o755); err != nil {
		return nil, fmt.Errorf("create ratelimit dir: %w", err)
	}
	statePath := filepath.Join(cacheRoot, ".ratelimit", name+".json")

	prl := &PersistentRateLimiter{
		statePath:   statePath,
		minInterval: minInterval,
	}

	// Load existing state
	if err := prl.load(); err != nil {
		// Corrupt or missing state - start fresh but don't immediately allow requests
		// This is conservative: if state is corrupt, we wait the full interval
		prl.lastRequest = time.Time{}
	}

	return prl, nil
}

// load loads the rate limiter state from disk.
func (rl *PersistentRateLimiter) load() error {
	data, err := os.ReadFile(rl.statePath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil // No state file yet
		}
		return fmt.Errorf("read ratelimit state: %w", err)
	}

	var state RateLimitState
	if err := json.Unmarshal(data, &state); err != nil {
		return fmt.Errorf("unmarshal ratelimit state: %w", err)
	}

	// If the timestamp is too old (> 24h), treat as stale
	if time.Since(state.LastRequest) > 24*time.Hour {
		rl.lastRequest = time.Time{}
	} else {
		rl.lastRequest = state.LastRequest
	}
	return nil
}

// save persists the rate limiter state to disk.
// Must be called with rl.mu held.
func (rl *PersistentRateLimiter) save() error {
	state := RateLimitState{LastRequest: rl.lastRequest}
	data, err := json.Marshal(state)
	if err != nil {
		return fmt.Errorf("marshal ratelimit state: %w", err)
	}

	// Write to temp file and atomically rename
	tmpPath := rl.statePath + ".tmp"
	if err := os.WriteFile(tmpPath, data, 0o644); err != nil {
		return fmt.Errorf("write ratelimit temp: %w", err)
	}
	if err := os.Rename(tmpPath, rl.statePath); err != nil {
		return fmt.Errorf("rename ratelimit state: %w", err)
	}
	return nil
}

// Wait blocks until the minimum interval has elapsed since the last request.
// This method is safe for concurrent use across multiple processes.
func (rl *PersistentRateLimiter) Wait() error {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	// Open the lock file
	lockPath := rl.statePath + ".lock"
	f, err := os.OpenFile(lockPath, os.O_CREATE|os.O_RDWR, 0o644)
	if err != nil {
		return fmt.Errorf("open lock file: %w", err)
	}
	defer f.Close()

	// Acquire exclusive lock
	if err := syscall.Flock(int(f.Fd()), syscall.LOCK_EX); err != nil {
		return fmt.Errorf("flock: %w", err)
	}
	defer syscall.Flock(int(f.Fd()), syscall.LOCK_UN)

	// Reload state in case another process updated it
	if err := rl.load(); err != nil {
		return fmt.Errorf("reload ratelimit state: %w", err)
	}

	now := time.Now()
	if !rl.lastRequest.IsZero() {
		elapsed := now.Sub(rl.lastRequest)
		if elapsed < rl.minInterval {
			wait := rl.minInterval - elapsed
			time.Sleep(wait)
			now = time.Now()
		}
	}

	// Update last request time and persist
	rl.lastRequest = now
	if err := rl.save(); err != nil {
		return fmt.Errorf("save ratelimit state: %w", err)
	}

	return nil
}

// arXivPersistentLimiter enforces arXiv's 1 request per 3 seconds limit across all CLI invocations.
var arXivPersistentLimiter *PersistentRateLimiter

// OpenAlexPersistentLimiter enforces OpenAlex's rate limits across CLI invocations.
var OpenAlexPersistentLimiter *PersistentRateLimiter

// InitPersistentRateLimiters initializes the persistent rate limiters.
// Must be called before using the paper search/fetch functions.
// cacheRoot is the repository root (where paper-cache lives).
func InitPersistentRateLimiters(cacheRoot string) error {
	var err error
	arXivPersistentLimiter, err = NewPersistentRateLimiter(cacheRoot, "arxiv", 3*time.Second)
	if err != nil {
		return fmt.Errorf("init arXiv rate limiter: %w", err)
	}
	OpenAlexPersistentLimiter, err = NewPersistentRateLimiter(cacheRoot, "openalex", 100*time.Millisecond) // 10 req/s
	if err != nil {
		return fmt.Errorf("init OpenAlex rate limiter: %w", err)
	}
	return nil
}

// WaitForArXivRateLimit waits for the arXiv rate limiter.
// This should be called before making arXiv API requests.
func WaitForArXivRateLimit() error {
	if arXivPersistentLimiter == nil {
		return fmt.Errorf("arXiv rate limiter not initialized")
	}
	return arXivPersistentLimiter.Wait()
}

// WaitForOpenAlexRateLimit waits for the OpenAlex rate limiter.
// This should be called before making OpenAlex API requests.
func WaitForOpenAlexRateLimit() error {
	if OpenAlexPersistentLimiter == nil {
		return fmt.Errorf("OpenAlex rate limiter not initialized")
	}
	return OpenAlexPersistentLimiter.Wait()
}
