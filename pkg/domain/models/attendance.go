package models

type Attendance struct {
	Type      string `json:"type"`
	UserID    string `json:"user_id"`
	CreatedAt string `json:"created_at"`
}
