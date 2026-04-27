package mapper

import (
	"blogapp/src/core/domain/entity"
	"blogapp/src/infrastructure/persistence/models"
)

func ToPostModel(post *entity.PostEntity) *models.PostModel {
    return &models.PostModel{
        UserID:  post.UserID,
        Title:   post.Title,
        Content: post.Content,
    }
}

func ToPostEntity(model *models.PostModel) *entity.PostEntity{
	return &entity.PostEntity{
		ID:        model.ID,
		UserID:    model.UserID,
		Title:     model.Title,
		Content:   model.Content,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}
}


