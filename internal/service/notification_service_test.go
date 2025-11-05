package service_test

import (
	"errors"
	"testing"

	"notification-service/internal/domain"
	"notification-service/internal/mocks"
	"notification-service/internal/service"

	"github.com/stretchr/testify/assert"
)

func TestNotificationService_Send(t *testing.T) {
	tests := []struct {
		name          string
		notification  domain.Notification
		setupMocks    func(*mocks.MockRateLimiter, *mocks.MockNotificationRepository)
		expectedError error
	}{
		{
			name: "success - should send notification when rate limit is not exceeded",
			notification: domain.Notification{
				UserID:  "user1",
				Type:    "status",
				Message: "Your order has been shipped",
			},
			setupMocks: func(rateLimiter *mocks.MockRateLimiter, repo *mocks.MockNotificationRepository) {
				rateLimiter.On("Allow", "user1", "status").Return(true)
				repo.On("Send", "user1", "Your order has been shipped").Return()
			},
			expectedError: nil,
		},
		{
			name: "error when rate limit is exceeded",
			notification: domain.Notification{
				UserID:  "user2",
				Type:    "status",
				Message: "Your order has been delivered",
			},
			setupMocks: func(rateLimiter *mocks.MockRateLimiter, repo *mocks.MockNotificationRepository) {
				rateLimiter.On("Allow", "user2", "status").Return(false)
			},
			expectedError: errors.New("rate limit exceeded"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rateLimiterMock := new(mocks.MockRateLimiter)
			notificationRepoMock := new(mocks.MockNotificationRepository)

			tt.setupMocks(rateLimiterMock, notificationRepoMock)

			service := service.NewNotificationService(rateLimiterMock, notificationRepoMock)

			err := service.Send(tt.notification)

			assert.Equal(t, tt.expectedError, err)

			rateLimiterMock.AssertExpectations(t)
			notificationRepoMock.AssertExpectations(t)
		})
	}
}
