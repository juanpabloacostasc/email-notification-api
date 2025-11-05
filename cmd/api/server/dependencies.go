package server

import (
	"log"

	"notification-service/config"
	"notification-service/internal/domain"
	"notification-service/internal/handler"
	"notification-service/internal/repository"
	"notification-service/internal/service"
)

func buildNotificationHandler() *handler.NotificationHandler {
	return handler.NewNotificationHandler(buildNotificationService())
}

func buildNotificationService() *service.NotificationService {
	cfg, err := config.LoadConfig("config/config.json")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	rules := cfg.ToDomainRateLimitRules()
	return service.NewNotificationService(buildRateLimiter(rules), buildNotificationGateway())
}

func buildRateLimiter(rules []domain.RateLimitRule) *service.RateLimiter {
	return service.NewRateLimiter(rules)
}

func buildNotificationGateway() *repository.NotificationRepository {
	return repository.NewNotificationRepository()
}
