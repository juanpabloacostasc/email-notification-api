package domain

type RateLimiter interface {
	Allow(userID, notificationType string) bool
}

type NotificationRepository interface {
	Send(userID, message string)
}

type NotificationService interface {
	Send(notification Notification) error
}
