package repository

import (
	"blogapp/src/core/domain/entity"
	"blogapp/src/core/domain/ports/repository"
	"blogapp/src/infrastructure/persistence/mapper"
	"blogapp/src/infrastructure/persistence/models"
	"context"

	"gorm.io/gorm"
)


type CommentRepositoryImpl struct{
	db *gorm.DB
}

func NewCommentRepositoryImpl(db *gorm.DB) repository.CommentRepository{
	return &CommentRepositoryImpl{
		db: db,
	}
}

func (r *CommentRepositoryImpl) Create(ctx context.Context, post *entity.CommentEntity) (*entity.CommentEntity,error){
	commentPost := mapper.ToCommentModel(post)

	if err := r.db.Create(commentPost).Error;err != nil{
		return nil,err
	}

	return mapper.ToCommentEntity(commentPost),nil
}

func (r *CommentRepositoryImpl) GetByPostID(ctx context.Context,id uint) ([]*entity.CommentEntity,error){
	var commModel []models.CommentModel
	if err := r.db.Where("post_id = ?",id).Find(&commModel).Error; err != nil{
		return nil,err
	}
	comment := make([]*entity.CommentEntity, len(commModel))
	for i,model := range commModel{
		comment[i] = mapper.ToCommentEntity(&model)
	}
	return comment,nil
}

func (r *CommentRepositoryImpl) Delete(ctx context.Context, id uint) error{
	return r.db.Delete(&models.CommentModel{},id).Error
}