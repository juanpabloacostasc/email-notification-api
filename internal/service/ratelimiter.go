package service

import (
	"sync"
	"time"

	"notification-service/internal/domain"
)

type (
	sentRecord struct {
		timestamp time.Time
	}

	RateLimiter struct {
		mu      sync.Mutex
		records map[string][]sentRecord
		Rules   map[string]domain.RateLimitRule
	}
)

func NewRateLimiter(rules []domain.RateLimitRule) *RateLimiter {
	ruleMap := make(map[string]domain.RateLimitRule)
	for _, rule := range rules {
		ruleMap[rule.Type] = rule
	}
	return &RateLimiter{
		records: make(map[string][]sentRecord),
		Rules:   ruleMap,
	}
}

func (rl *RateLimiter) Allow(userID, notificationType string) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	rule, ok := rl.Rules[notificationType]
	if !ok {
		return true
	}

	now := time.Now()
	key := userID + "_" + notificationType

	records := rl.records[key]
	n := 0
	for _, record := range records {
		if now.Sub(record.timestamp) < rule.Duration {
			records[n] = record
			n++
		}
	}
	rl.records[key] = records[:n]

	if len(rl.records[key]) >= rule.Limit {
		return false
	}

	rl.records[key] = append(rl.records[key], sentRecord{timestamp: now})

	return true
}
