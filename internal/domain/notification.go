package domain

type Notification struct {
	Type    string `json:"type"`
	UserID  string `json:"user_id"`
	Message string `json:"message"`
}
