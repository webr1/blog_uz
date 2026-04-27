package repository

import (
	"blogapp/src/core/domain/entity"
	"blogapp/src/core/domain/ports/repository"
	"blogapp/src/infrastructure/persistence/mapper"
	"blogapp/src/infrastructure/persistence/models"
	"context"

	"gorm.io/gorm"
)


type LikeRepositoryImpl struct{
	db *gorm.DB
}

func NewLikeRepositoryImpl(db *gorm.DB) repository.LikeRepository{
	return &LikeRepositoryImpl{
		db: db,
	}
}

func (r *LikeRepositoryImpl) Create(ctx context.Context,post *entity.LikeEntity) (*entity.LikeEntity,error){
	like := mapper.ToLikeModel(post)

	if err := r.db.Create(like).Error;err != nil{
		return nil,err
	}

	return mapper.ToLikeEntity(like),nil
}

func (r *LikeRepositoryImpl) Delete(ctx context.Context,id uint) error{
	return r.db.Delete(&models.LikeModel{},id).Error
}

func (r *LikeRepositoryImpl) GetByPostID(ctx context.Context, id uint) ([]*entity.LikeEntity,error){
	var likeModel []models.LikeModel
	if err := r.db.Where("post_id = ?", id).Find(&likeModel).Error;err != nil{
		return nil,err
	}
	likes := make([]*entity.LikeEntity,len(likeModel))
	for i,model := range likeModel{
		likes[i] = mapper.ToLikeEntity(&model)
	}
	return likes,nil
}
func (r *LikeRepositoryImpl) GetByUserAndPost(ctx context.Context, userID, postID uint) (*entity.LikeEntity, error) {
    var likeModel models.LikeModel
    if err := r.db.Where("user_id = ? AND post_id = ?", userID, postID).First(&likeModel).Error; err != nil {
        return nil, err
    }
    return mapper.ToLikeEntity(&likeModel), nil
}