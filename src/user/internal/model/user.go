package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       string
	Name     string
	Email    string `gorm:"index:idx_user_email,unique"`
	Password string
}
