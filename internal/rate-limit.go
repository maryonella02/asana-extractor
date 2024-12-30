package internal

import (
	"fmt"
	"time"
)

// FixedWindowRateLimiter implements rate limiting using the Fixed Window algorithm.
type FixedWindowRateLimiter struct {
	WindowSize   time.Duration
	MaxRequests  int
	LastReset    time.Time
	RequestCount int
}

func NewFixedWindowRateLimiter(interval time.Duration, maxRequests int) *FixedWindowRateLimiter {
	return &FixedWindowRateLimiter{
		WindowSize:  interval,
		MaxRequests: maxRequests,
	}
}

// Allow checks if the request is allowed based on Fixed Window rate limiting rules.
func (limiter *FixedWindowRateLimiter) Allow() error {
	currentTime := time.Now()

	if currentTime.Sub(limiter.LastReset) < limiter.WindowSize {
		if limiter.RequestCount >= limiter.MaxRequests {
			return fmt.Errorf("rate limit exceeded")
			//better to add as const
		}

		limiter.RequestCount++
	} else {
		limiter.RequestCount = 1
		limiter.LastReset = currentTime
	}

	return nil
}
