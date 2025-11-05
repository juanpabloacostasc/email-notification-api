package service

import (
	"sync"
	"time"

	"notification-service/internal/domain"
)

type (
	SentRecord struct {
		Timestamp time.Time
	}

	RateLimiter struct {
		mu      sync.Mutex
		Records map[string][]SentRecord
		Rules   map[string]domain.RateLimitRule
	}
)

func NewRateLimiter(rules []domain.RateLimitRule) *RateLimiter {
	ruleMap := make(map[string]domain.RateLimitRule)
	for _, rule := range rules {
		ruleMap[rule.Type] = rule
	}
	return &RateLimiter{
		Records: make(map[string][]SentRecord),
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

	records := rl.Records[key]
	n := 0
	for _, record := range records {
		if now.Sub(record.Timestamp) < rule.Duration {
			records[n] = record
			n++
		}
	}
	rl.Records[key] = records[:n]

	if len(rl.Records[key]) >= rule.Limit {
		return false
	}

	rl.Records[key] = append(rl.Records[key], SentRecord{Timestamp: now})

	return true
}
