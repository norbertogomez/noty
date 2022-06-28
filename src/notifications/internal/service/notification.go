package service

import (
	"github.com/google/uuid"
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
	err := s.notificationRepo.Create(notification)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (s NotificationService) UpdateNotification(notificationDTO dto.NotificationUpdateInfo) (bool, error) {
	notification := notificationDTO.ToModel()
	err := s.notificationRepo.Update(notification)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (s NotificationService) FindByID(id string) (dto.NotificationInfo, error) {
	var notification dto.NotificationInfo
	notificationUUID, err := uuid.Parse(id)
	if err != nil {
		return notification, err
	}

	if err := s.notificationRepo.FindByID(&notification, notificationUUID.String()); err != nil {
		return notification, err
	}

	return notification, nil
}
