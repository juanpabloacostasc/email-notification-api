package handler_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"notification-service/internal/domain"
	"notification-service/internal/handler"
	"notification-service/internal/mocks"

	"github.com/stretchr/testify/assert"
)

func TestNotificationHandler_SendNotification(t *testing.T) {
	tests := []struct {
		name           string
		requestBody    []byte
		setupMocks     func(*mocks.MockNotificationService)
		expectedStatus int
		expectedBody   string
	}{
		{
			name:        "success",
			requestBody: marshal(domain.Notification{UserID: "user1", Type: "status", Message: "Hello"}),
			setupMocks: func(service *mocks.MockNotificationService) {
				service.On("Send", domain.Notification{UserID: "user1", Type: "status", Message: "Hello"}).Return(nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   "Notification sent successfully",
		},
		{
			name:           "fail - invalid request body",
			requestBody:    []byte(`{"invalid json"}`),
			setupMocks:     func(service *mocks.MockNotificationService) {},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "Invalid request body\n",
		},
		{
			name:        "fail - rate limit exceeded",
			requestBody: marshal(domain.Notification{UserID: "user2", Type: "news", Message: "News update"}),
			setupMocks: func(service *mocks.MockNotificationService) {
				service.On("Send", domain.Notification{UserID: "user2", Type: "news", Message: "News update"}).Return(errors.New("rate limit exceeded"))
			},
			expectedStatus: http.StatusTooManyRequests,
			expectedBody:   "rate limit exceeded\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService := new(mocks.MockNotificationService)
			tt.setupMocks(mockService)

			h := handler.NewNotificationHandler(mockService)

			req, err := http.NewRequest("POST", "/send", bytes.NewBuffer(tt.requestBody))
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(h.SendNotification)

			handler.ServeHTTP(rr, req)

			assert.Equal(t, tt.expectedStatus, rr.Code)
			assert.Equal(t, tt.expectedBody, rr.Body.String())

			mockService.AssertExpectations(t)
		})
	}
}

func marshal(v interface{}) []byte {
	bytes, _ := json.Marshal(v)
	return bytes
}
