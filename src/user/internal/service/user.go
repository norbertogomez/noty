package service

import (
	"noty-users/internal/dto"
	"noty-users/internal/repository"
	"noty-users/internal/utils"
)

type UserCreator interface {
	CreateUser(info dto.UserCreationInfo) (bool, error)
}

type UserService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return UserService{userRepo: userRepository}
}

func (svc *UserService) CreateUser(info dto.UserCreationInfo) (bool, error) {
	pwd, err := utils.HashPassword(info.Password)
	if err != nil {
		return false, err
	}

	info.Password = pwd
	user := info.ToModel()
	if err := svc.userRepo.CreateUser(user); err != nil {
		return false, err
	}

	return true, nil
}
