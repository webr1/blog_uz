package models

import (
	"gorm.io/gorm"
)


type PostModel struct{
	gorm.Model
	UserID uint 
	Title   string `gorm:"size:255;not null"`
	Content string `gorm:"type:text"`
}