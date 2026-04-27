package entity

import (
	"time"
)


type UserEntity struct {
	ID uint
	Username string
	Email string
	PasswordHash string
	CreatedAt time.Time
}


func NewUserEntity(username,email, password string) *UserEntity{
	return &UserEntity{
		Username: username,
		Email: email,
		PasswordHash: password,
		CreatedAt: time.Now(),
	}
}



