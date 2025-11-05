package server

import "net/http"

func MapURLs() {
	notificationHandler := buildNotificationHandler()

	http.HandleFunc("/send", notificationHandler.SendNotification)
}
