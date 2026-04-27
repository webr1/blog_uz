package mapper

import (
	"blogapp/src/core/domain/entity"
	"blogapp/src/infrastructure/persistence/models"
)

func ToUserModel(user *entity.UserEntity) *models.UserModel {
	return &models.UserModel{
		Username:     user.Username,
		Email:        user.Email,
		PasswordHash: user.PasswordHash,
	}
}

func ToUserEntity(model *models.UserModel) *entity.UserEntity {
	return &entity.UserEntity{
		ID: model.ID,
		Username:     model.Username,
		Email:        model.Email,
		PasswordHash: model.PasswordHash,
		CreatedAt: model.CreatedAt,

	}
}
