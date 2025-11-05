package mocks

import (
	"notification-service/internal/domain"

	"github.com/stretchr/testify/mock"
)

type MockNotificationService struct {
	mock.Mock
}

func (m *MockNotificationService) Send(notification domain.Notification) error {
	args := m.Called(notification)
	return args.Error(0)
}