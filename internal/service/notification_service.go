package service

import (
	"errors"

	"notification-service/internal/domain"
)

type NotificationService struct {
	rateLimiter            domain.RateLimiter
	notificationRepository domain.NotificationRepository
}

func NewNotificationService(rateLimiter domain.RateLimiter, notificationRepository domain.NotificationRepository) *NotificationService {
	return &NotificationService{
		rateLimiter:            rateLimiter,
		notificationRepository: notificationRepository,
	}
}

func (s *NotificationService) Send(notification domain.Notification) error {
	if !s.rateLimiter.Allow(notification.UserID, notification.Type) {
		return errors.New("rate limit exceeded")
	}

	s.notificationRepository.Send(notification.UserID, notification.Message)
	return nil
}
