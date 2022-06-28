package dto

type NotificationInfo struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Strategy    string `json:"strategy"`
	NotifyAt    string `json:"notify_at"`
	UserID      string `json:"user_id"`
}

func NewNotificationInfo(ID string, name string, description string, strategy string, notifyAt string, userID string) *NotificationInfo {
	return &NotificationInfo{ID: ID, Name: name, Description: description, Strategy: strategy, NotifyAt: notifyAt, UserID: userID}
}
