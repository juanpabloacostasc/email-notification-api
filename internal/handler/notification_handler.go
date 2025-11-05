package handler

import (
	"encoding/json"
	"net/http"

	"notification-service/internal/domain"
)

type NotificationHandler struct {
	service domain.NotificationService
}

func NewNotificationHandler(service domain.NotificationService) *NotificationHandler {
	return &NotificationHandler{service: service}
}

func (h *NotificationHandler) SendNotification(w http.ResponseWriter, r *http.Request) {
	var notification domain.Notification
	if err := json.NewDecoder(r.Body).Decode(&notification); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.service.Send(notification); err != nil {
		http.Error(w, err.Error(), http.StatusTooManyRequests)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Notification sent successfully"))
}
