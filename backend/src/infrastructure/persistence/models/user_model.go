package models

import (
	"gorm.io/gorm"
)


type UserModel struct{
	gorm.Model
	Username string `gorm:"size:255;not null;unique"`
	Email    string `gorm:"size:255;not null;unique"`
	PasswordHash string `gorm:"size:255;not null"`
}