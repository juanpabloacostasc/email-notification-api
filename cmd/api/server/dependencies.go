package server

import (
	"time"

	"notification-service/internal/domain"
	"notification-service/internal/handler"
	"notification-service/internal/repository"
	"notification-service/internal/service"
)

func buildNotificationHandler() *handler.NotificationHandler {
	return handler.NewNotificationHandler(buildNotificationService())
}

func buildNotificationService() *service.NotificationService {
	rules := []domain.RateLimitRule{
		{Type: "status", Limit: 2, Duration: time.Minute},
		{Type: "news", Limit: 1, Duration: 24 * time.Hour},
		{Type: "marketing", Limit: 3, Duration: time.Hour},
		{Type: "project_invitations", Limit: 1, Duration: 24 * time.Hour},
	}
	return service.NewNotificationService(buildRateLimiter(rules), buildNotificationGateway())
}

func buildRateLimiter(rules []domain.RateLimitRule) *service.RateLimiter {
	return service.NewRateLimiter(rules)
}

func buildNotificationGateway() *repository.NotificationRepository {
	return repository.NewNotificationRepository()
}
