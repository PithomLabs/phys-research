package paper

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

// RateLimiter implements a token bucket rate limiter for API requests.
type RateLimiter struct {
	mu         sync.Mutex
	tokens     float64
	maxTokens  float64
	refillRate float64 // tokens per second
	lastRefill time.Time
}

// NewRateLimiter creates a new rate limiter with the given rate (requests per second).
func NewRateLimiter(requestsPerSecond float64) *RateLimiter {
	return &RateLimiter{
		tokens:     requestsPerSecond,
		maxTokens:  requestsPerSecond,
		refillRate: requestsPerSecond,
		lastRefill: time.Now(),
	}
}

// Wait blocks until a token is available, then consumes it.
func (rl *RateLimiter) Wait() {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	for {
		now := time.Now()
		elapsed := now.Sub(rl.lastRefill).Seconds()
		rl.tokens = minFloat(rl.maxTokens, rl.tokens+elapsed*rl.refillRate)
		rl.lastRefill = now

		if rl.tokens >= 1 {
			rl.tokens--
			return
		}

		// Wait for the next token
		waitTime := time.Duration((1 - rl.tokens) / rl.refillRate * float64(time.Second))
		rl.mu.Unlock()
		time.Sleep(waitTime)
		rl.mu.Lock()
	}
}

func minFloat(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

// arXivRateLimiter enforces arXiv's 1 request per 3 seconds limit across all calls.
var arXivRateLimiter = NewRateLimiter(1.0 / 3.0)

// OpenAlexRateLimiter enforces OpenAlex's rate limits.
// Without API key: 10 req/s (shared budget). With API key: higher limits.
var OpenAlexRateLimiter = NewRateLimiter(10.0)

// DOI normalization: lowercases and trims whitespace/URL prefixes.
func normalizeDOI(doi string) string {
	d := strings.TrimSpace(doi)
	d = strings.TrimPrefix(d, "https://doi.org/")
	d = strings.TrimPrefix(d, "http://doi.org/")
	d = strings.TrimPrefix(d, "doi:")
	d = strings.ToLower(d)
	return d
}

// RetryConfig defines retry behavior for HTTP requests.
type RetryConfig struct {
	MaxRetries    int
	BaseDelay     time.Duration
	MaxDelay      time.Duration
	BackoffFactor float64
}

// DefaultRetryConfig returns a sensible default for API retries.
func DefaultRetryConfig() RetryConfig {
	return RetryConfig{
		MaxRetries:    3,
		BaseDelay:     1 * time.Second,
		MaxDelay:      30 * time.Second,
		BackoffFactor: 2.0,
	}
}

// doWithRetry executes an HTTP request with retries for 429/503 responses.
func doWithRetry(req *http.Request, client *http.Client, cfg RetryConfig) (*http.Response, error) {
	var lastErr error
	for attempt := 0; attempt <= cfg.MaxRetries; attempt++ {
		if attempt > 0 {
			delay := cfg.BaseDelay
			for i := 1; i < attempt; i++ {
				delay = time.Duration(float64(delay) * cfg.BackoffFactor)
				if delay > cfg.MaxDelay {
					delay = cfg.MaxDelay
				}
			}
			time.Sleep(delay)
		}

		resp, err := client.Do(req)
		if err != nil {
			lastErr = err
			continue
		}

		// Check for retryable status codes
		if resp.StatusCode == 429 || resp.StatusCode == 503 {
			resp.Body.Close()
			// Check for Retry-After header
			if retryAfter := resp.Header.Get("Retry-After"); retryAfter != "" {
				if secs, err := strconv.Atoi(retryAfter); err == nil {
					time.Sleep(time.Duration(secs) * time.Second)
					continue
				}
				// If it's an HTTP-date, we could parse it, but for simplicity just use backoff
			}
			continue
		}

		return resp, nil
	}

	return nil, fmt.Errorf("after %d retries: %w", cfg.MaxRetries, lastErr)
}
