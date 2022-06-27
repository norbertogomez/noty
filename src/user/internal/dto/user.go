package dto

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"noty-users/internal/model"
)

type UserCreationInfo struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (dto *UserCreationInfo) ToModel() model.User {
	return model.User{
		Model:    gorm.Model{},
		ID:       uuid.New().String(),
		Name:     dto.Name,
		Email:    dto.Email,
		Password: dto.Password,
	}
}
