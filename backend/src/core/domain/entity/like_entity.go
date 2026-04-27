package entity

import (
	"time"
)

type LikeEntity struct{
	ID uint 
	UserID uint
	PostID uint
	CreatedAt time.Time
}

func NewLikeEntity(userID,postID uint) *LikeEntity{
	return &LikeEntity{
		UserID: userID,
		PostID: postID,
		CreatedAt: time.Now(),
	}
}

