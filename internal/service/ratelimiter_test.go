package service

import (
	"testing"
	"time"

	"notification-service/internal/domain"

	"github.com/stretchr/testify/assert"
)

func TestRateLimiter_Allow(t *testing.T) {
	rules := []domain.RateLimitRule{
		{Type: "status", Limit: 2, Duration: time.Minute},
		{Type: "news", Limit: 1, Duration: 24 * time.Hour},
		{Type: "marketing", Limit: 3, Duration: time.Hour},
	}

	rateLimiter := NewRateLimiter(rules)

	tests := []struct {
		name             string
		userID           string
		notificationType string
		setup            func()
		expectedResult   bool
	}{
		{
			name:             "true - no rule is defined for the notification type",
			userID:           "user1",
			notificationType: "unknown_type",
			expectedResult:   true,
		},
		{
			name:             "true - limit is not exceeded",
			userID:           "user2",
			notificationType: "status",
			expectedResult:   true,
		},
		{
			name:             "false - limit is exceeded",
			userID:           "user3",
			notificationType: "status",
			setup: func() {
				rateLimiter.Allow("user3", "status")
				rateLimiter.Allow("user3", "status")
			},
			expectedResult: false,
		},
		{
			name:             "true - allow again after the time window has passed",
			userID:           "user4",
			notificationType: "status",
			setup: func() {
				key := "user4_status"
				rateLimiter.records[key] = []sentRecord{
					{timestamp: time.Now().Add(-2 * time.Minute)},
					{timestamp: time.Now().Add(-2 * time.Minute)},
				}
			},
			expectedResult: true,
		},
		{
			name:             "true - allow a single notification for news type",
			userID:           "user5",
			notificationType: "news",
			expectedResult:   true,
		},
		{
			name:             "false - deny a second notification for news type within 24 hours",
			userID:           "user6",
			notificationType: "news",
			setup: func() {
				rateLimiter.Allow("user6", "news")
			},
			expectedResult: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				tt.setup()
			}
			assert.Equal(t, tt.expectedResult, rateLimiter.Allow(tt.userID, tt.notificationType))
		})
	}
}
