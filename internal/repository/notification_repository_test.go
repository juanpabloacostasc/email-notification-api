package repository

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNotificationRepository_Send(t *testing.T) {
	tests := []struct {
		name     string
		userID   string
		message  string
		expected string
	}{
		{
			name:     "Test with simple message",
			userID:   "user1",
			message:  "Hello",
			expected: "Sending message to user user1: Hello\n",
		},
		{
			name:     "Test with empty message",
			userID:   "user2",
			message:  "",
			expected: "Sending message to user user2: \n",
		},
		{
			name:     "Test with message with spaces",
			userID:   "user3",
			message:  "Hello World",
			expected: "Sending message to user user3: Hello World\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			old := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			repo := NewNotificationRepository()
			repo.Send(tt.userID, tt.message)

			w.Close()
			os.Stdout = old

			var buf bytes.Buffer
			io.Copy(&buf, r)
			r.Close()

			assert.Equal(t, tt.expected, buf.String())
		})
	}
}