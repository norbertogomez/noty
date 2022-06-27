package repository

import (
	"gorm.io/gorm"
	"noty-users/internal/database"
	"noty-users/internal/model"
)

type UserFinder interface {
	FindOneByEmail(email string, user *model.User) error
}

type UserCreator interface {
	CreateUser(user model.User) error
}

type UserRepository interface {
	UserFinder
	UserCreator
}

type GormUserRepository struct {
	Db *gorm.DB
}

func NewUserRepository() *GormUserRepository {
	db := database.InitDb()
	db.AutoMigrate(&model.User{})
	return &GormUserRepository{Db: db}
}

func (rp *GormUserRepository) CreateUser(user model.User) error {
	if err := rp.Db.Create(&user).Error; err != nil {
		return err
	}

	return nil
}

func (rp *GormUserRepository) FindOneByEmail(email string, user *model.User) error {
	err := rp.Db.Where("email = ?", email).First(user).Error
	if err != nil {
		return err
	}
	return nil
}
