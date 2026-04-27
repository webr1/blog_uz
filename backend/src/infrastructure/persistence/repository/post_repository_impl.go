package repository

import (
	"blogapp/src/core/domain/entity"
	"blogapp/src/core/domain/ports/repository"
	"blogapp/src/infrastructure/persistence/mapper"
	"blogapp/src/infrastructure/persistence/models"
	"context"

	"gorm.io/gorm"
)


type PostRepositoryImpl struct{
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) repository.PostRepository{
	return &PostRepositoryImpl{
		db: db,
	}
}

func (r *PostRepositoryImpl) Create(ctx context.Context, post *entity.PostEntity) (*entity.PostEntity,error){
	postModel := mapper.ToPostModel(post)
	if err := r.db.Create(postModel).Error;err != nil{
		return nil,err
	}
	return mapper.ToPostEntity(postModel),nil
}

func (r *PostRepositoryImpl) GetByID(ctx context.Context, id uint) (*entity.PostEntity,error){
	var postModel models.PostModel
	if err := r.db.First(&postModel,id).Error;err != nil{
		return nil,err
	}
	return mapper.ToPostEntity(&postModel),nil
}

func (r *PostRepositoryImpl) GetAll(ctx context.Context) ([]*entity.PostEntity,error){
	var postModel []models.PostModel
	if err := r.db.Find(&postModel).Error;err != nil{
		return nil,err
	}
	posts := make([]*entity.PostEntity, len(postModel))
	for i,model := range postModel{
		posts[i] = mapper.ToPostEntity(&model)
	}
	return posts,nil
	
}

func (r *PostRepositoryImpl) GetByUserID(ctx context.Context, userID uint) ([]*entity.PostEntity, error) {
    var postModels []models.PostModel
    if err := r.db.Where("user_id = ?", userID).Find(&postModels).Error; err != nil {
        return nil, err
    }
    posts := make([]*entity.PostEntity, len(postModels))
    for i, model := range postModels {
        posts[i] = mapper.ToPostEntity(&model)
    }
    return posts, nil
}




func (r * PostRepositoryImpl) Update(ctx context.Context, post *entity.PostEntity) (*entity.PostEntity,error){
	postUpdate := mapper.ToPostModel(post)
	if err := r.db.Save(postUpdate).Error; err != nil {
		return nil, err
	}
	return mapper.ToPostEntity(postUpdate), nil
}

func (r *PostRepositoryImpl) Delete(ctx context.Context, id uint) error {
    return r.db.Delete(&models.PostModel{}, id).Error
}