package model

import "time"

type Notification struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Strategy    string    `json:"strategy"`
	NotifyAt    time.Time `json:"notify_at"`
	UserID      string    `json:"user_id"`
}
