package models

import (
	"gorm.io/gorm"
)

type CommentModel struct{
	gorm.Model
	PostID uint
	UserID uint
	Content string `gorm:"size:1000;not null"`
}