package entity

import (
	"time"
)

type ProfileEntity struct{
	ID uint
	UserID uint
	FullName string
	Bio string
	Avatar string
	CreatedAt time.Time
	UpdatedAt time.Time
}


func NewProfileEntity(userID uint, fullName, bio, avatar string) *ProfileEntity{
	return &ProfileEntity{
		UserID: userID,
		FullName: fullName,
		Bio: bio,
		Avatar: avatar,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}