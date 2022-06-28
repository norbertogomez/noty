package dto

import (
	"noty-notifications/internal/model"
	"time"
)

type NotificationUpdateInfo struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Strategy    string    `json:"strategy"`
	NotifyAt    time.Time `json:"notify_at"`
}

func (dto *NotificationUpdateInfo) ToModel() model.Notification {
	return model.Notification{
		ID:          dto.ID,
		Name:        dto.Name,
		Description: dto.Description,
		Strategy:    dto.Strategy,
		NotifyAt:    dto.NotifyAt,
	}
}
