package domain

//go:generate mockery generate -name RateLimiter -filename rate_limiter_mock.go -output mocks
type RateLimiter interface {
	Allow(userID, notificationType string) bool
}

//go:generate mockery generate -name NotificationRepository -filename notification_repository_mock.go -output mocks
type NotificationRepository interface {
	Send(userID, message string)
}

//go:generate mockery generate -name NotificationService -filename notification_service_mock.go -output mocks
type NotificationService interface {
	Send(notification Notification) error
}
