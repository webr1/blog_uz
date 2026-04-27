package entity

import (
	"time"
)

type CommentEntity struct{
	ID uint
	UserID uint
	PostID uint
	Content string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewCommentEntity(userID, postID uint, content string) *CommentEntity{
	return &CommentEntity{
		UserID: userID,
		PostID: postID,
		Content: content,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}