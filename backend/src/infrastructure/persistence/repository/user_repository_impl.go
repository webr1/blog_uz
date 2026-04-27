package repository

import (
	"context"
	"blogapp/src/core/domain/entity"
	"blogapp/src/core/domain/ports/repository"
	"blogapp/src/infrastructure/persistence/mapper"
	"blogapp/src/infrastructure/persistence/models"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &UserRepositoryImpl{db: db}
}

func (r *UserRepositoryImpl) Create(ctx context.Context, user *entity.UserEntity) (*entity.UserEntity, error) {
	userModel := mapper.ToUserModel(user)
	if err := r.db.Create(userModel).Error; err != nil {
		return nil, err
	}
	return mapper.ToUserEntity(userModel), nil
}


func (r *UserRepositoryImpl) GetByID(ctx context.Context, id uint) (*entity.UserEntity,error){
	var userModel models.UserModel
	if err := r.db.First(&userModel, id).Error;err != nil{
		return nil,err
	}
	return mapper.ToUserEntity(&userModel),nil
}


func (r *UserRepositoryImpl) GetByEmail(ctx context.Context, email string) (*entity.UserEntity,error){
	var userModel models.UserModel
	if err := r.db.Where("email = ?", email).First(&userModel).Error; err != nil {
		return nil, err
	}
	return mapper.ToUserEntity(&userModel),nil
}

func (r *UserRepositoryImpl) GetByUsername(ctx context.Context, username string) (*entity.UserEntity,error){
	var userModel models.UserModel
	if err := r.db.Where("username = ?",username).First(&userModel).Error;err != nil{
		return  nil,err
	}
	return mapper.ToUserEntity(&userModel),nil
}
