package dto

import (
	"github.com/google/uuid"
	"noty-notifications/internal/model"
	"time"
)

type NotificationCreationInfo struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Strategy    string    `json:"strategy"`
	NotifyAt    time.Time `json:"notify_at"`
	UserID      string    `json:"user_id"`
}

func (dto *NotificationCreationInfo) ToModel() model.Notification {
	return model.Notification{
		ID:          uuid.New().String(),
		Name:        dto.Name,
		Description: dto.Description,
		Strategy:    dto.Strategy,
		NotifyAt:    dto.NotifyAt,
		UserID:      dto.UserID,
	}
}
