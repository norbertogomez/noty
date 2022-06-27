package service

import (
	"noty-users/internal/model"
	"noty-users/internal/repository"
	"noty-users/internal/utils"
)

type LoginService interface {
	LoginUser(email string, password string) bool
}

type LoginUserService struct {
	repository.UserRepository
}

func NewLoginUserService(userRepository repository.UserRepository) *LoginUserService {
	return &LoginUserService{UserRepository: userRepository}
}

func (db LoginUserService) LoginUser(email string, password string) bool {
	var user model.User
	err := db.FindOneByEmail(email, &user)
	if err != nil {
		return false
	}

	if utils.CheckPasswordHash(password, user.Password) {
		return true
	}

	return false
}
