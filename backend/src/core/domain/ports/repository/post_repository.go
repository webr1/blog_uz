package repository

import (
	"blogapp/src/core/domain/entity"
	"context"
)



type  PostRepository interface{
	Create(ctx context.Context, post *entity.PostEntity) (*entity.PostEntity, error)
	GetByID(ctx context.Context, id uint) (*entity.PostEntity,error)
	GetAll(ctx context.Context) ([]*entity.PostEntity,error)
	GetByUserID(ctx context.Context, userID uint) ([]*entity.PostEntity,error)
	Update(ctx context.Context, post *entity.PostEntity) (*entity.PostEntity,error)
	Delete(ctx context.Context, id uint ) error

}