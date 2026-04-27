package mapper

import (
	"blogapp/src/core/domain/entity"
	"blogapp/src/infrastructure/persistence/models"
)

func ToProfileModel(profile *entity.ProfileEntity) *models.ProfileModel {
	return &models.ProfileModel{
		UserID:   profile.UserID,
		FullName: profile.FullName,
		Bio:      profile.Bio,
		Avatar:   profile.Avatar,
	}
}

func ToProfileEntity(model *models.ProfileModel) *entity.ProfileEntity {
	return &entity.ProfileEntity{
		ID:        model.ID,
		UserID:    model.UserID,
		FullName:  model.FullName,
		Bio:       model.Bio,
		Avatar:    model.Avatar,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}
}