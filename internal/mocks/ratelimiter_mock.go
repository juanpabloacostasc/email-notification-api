package mocks

import (
	"github.com/stretchr/testify/mock"
)

type MockRateLimiter struct {
	mock.Mock
}

func (m *MockRateLimiter) Allow(userID, notificationType string) bool {
	args := m.Called(userID, notificationType)
	return args.Bool(0)
}