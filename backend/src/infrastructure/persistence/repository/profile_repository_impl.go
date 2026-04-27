package repository

import (
	"blogapp/src/core/domain/entity"
	"blogapp/src/core/domain/ports/repository"
	"blogapp/src/infrastructure/persistence/mapper"
	"blogapp/src/infrastructure/persistence/models"
	"context"

	"gorm.io/gorm"
)


type ProfileRepositoryImpl struct{
	db *gorm.DB
}

func NewProfileRepositoryImpl(db *gorm.DB) repository.ProfileRepository{
	return &ProfileRepositoryImpl{
		db: db,
	}
}


func (r *ProfileRepositoryImpl) Create(ctx context.Context,post *entity.ProfileEntity) (*entity.ProfileEntity,error){
	profileModel := mapper.ToProfileModel(post)
	if err := r.db.Create(profileModel).Error; err != nil{
		return nil,err
	}
	return mapper.ToProfileEntity(profileModel),nil
}


func (r *ProfileRepositoryImpl) Update(ctx context.Context, post *entity.ProfileEntity) (*entity.ProfileEntity,error){
	postUpdate := mapper.ToProfileModel(post)
	if err := r.db.Save(postUpdate).Error; err != nil{
		return nil,err
	}
	return mapper.ToProfileEntity(postUpdate),nil
}


func (r *ProfileRepositoryImpl) GetByUserID(ctx context.Context,id uint) (*entity.ProfileEntity,error){
	var profileModel models.ProfileModel
	if err := r.db.Where("user_id = ?", id).First(&profileModel).Error; err != nil {
		return nil,err
	}
	return mapper.ToProfileEntity(&profileModel),nil
}