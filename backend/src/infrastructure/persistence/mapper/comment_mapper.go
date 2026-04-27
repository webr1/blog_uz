package mapper

import (
	"blogapp/src/core/domain/entity"
	"blogapp/src/infrastructure/persistence/models"
)

func ToCommentModel(comment *entity.CommentEntity) *models.CommentModel {
	return &models.CommentModel{
		PostID:  comment.PostID,
		UserID:  comment.UserID,
		Content: comment.Content,
	}
}

func ToCommentEntity(model *models.CommentModel) *entity.CommentEntity {
	return &entity.CommentEntity{
		ID:        model.ID,
		PostID:    model.PostID,
		UserID:    model.UserID,
		Content:   model.Content,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}
}