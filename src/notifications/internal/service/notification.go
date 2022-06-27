package service

import (
	"noty-notifications/internal/dto"
	"noty-notifications/internal/repository"
)

type NotificationService struct {
	notificationRepo repository.NotificationRepository
}

func NewNotificationService(notificationRepo repository.NotificationRepository) NotificationService {
	return NotificationService{notificationRepo: notificationRepo}
}

func (s NotificationService) CreateNotification(notificationDTO dto.NotificationCreationInfo) (bool, error) {
	notification := notificationDTO.ToModel()
	err := s.notificationRepo.CreateNotification(notification)
	if err != nil {
		return false, err
	}

	return true, nil
}
