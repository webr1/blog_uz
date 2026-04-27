package models

import (
	"gorm.io/gorm"
)


type ProfileModel struct{
	gorm.Model
	UserID uint
	FullName string
	Bio string
	Avatar string
}
