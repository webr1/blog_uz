package repository



import (
	"context"
	"blogapp/src/core/domain/entity"
)


type CommentRepository interface {
	Create(ctx context.Context , post *entity.CommentEntity) (*entity.CommentEntity, error)
	GetByPostID(ctx context.Context,postID uint) ([]*entity.CommentEntity,error)
	Delete(ctx context.Context,id uint) error
}
