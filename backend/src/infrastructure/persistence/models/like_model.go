package models

import (
	"gorm.io/gorm"
)

type LikeModel struct{
	gorm.Model	
	PostID uint
	UserID uint
}