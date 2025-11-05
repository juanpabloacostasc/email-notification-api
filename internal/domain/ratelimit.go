package domain

import "time"

type RateLimitRule struct {
	Type     string
	Limit    int
	Duration time.Duration
}
