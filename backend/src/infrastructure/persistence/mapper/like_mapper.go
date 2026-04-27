package mapper

import (
	"blogapp/src/core/domain/entity"
	"blogapp/src/infrastructure/persistence/models"
)

func ToLikeModel(like *entity.LikeEntity) *models.LikeModel {
	return &models.LikeModel{
		PostID: like.PostID,
		UserID: like.UserID,
	}
}

func ToLikeEntity(model *models.LikeModel) *entity.LikeEntity {
	return &entity.LikeEntity{
		ID:        model.ID,
		PostID:    model.PostID,
		UserID:    model.UserID,
		CreatedAt: model.CreatedAt,
	}
}