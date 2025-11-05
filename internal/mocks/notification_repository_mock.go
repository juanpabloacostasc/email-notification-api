package mocks

import (
	"github.com/stretchr/testify/mock"
)

type MockNotificationRepository struct {
	mock.Mock
}

func (m *MockNotificationRepository) Send(userID, message string) {
	m.Called(userID, message)
}