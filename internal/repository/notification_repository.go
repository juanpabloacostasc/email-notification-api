package repository

import "fmt"

type NotificationRepository struct{}

func NewNotificationRepository() *NotificationRepository {
	return &NotificationRepository{}
}

func (g *NotificationRepository) Send(userID, message string) {
	fmt.Printf("Sending message to user %s: %s\n", userID, message)
}
