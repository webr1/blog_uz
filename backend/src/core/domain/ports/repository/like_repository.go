package repository


import (
	"context"
	"blogapp/src/core/domain/entity"
)


type LikeRepository interface {
	Create(ctx context.Context, post *entity.LikeEntity) (*entity.LikeEntity, error)
	Delete(ctx context.Context, id uint) error
	GetByPostID(ctx context.Context, postID uint) ([]*entity.LikeEntity,error)
	GetByUserAndPost(ctx context.Context, userID,postID uint) (*entity.LikeEntity,error)

}