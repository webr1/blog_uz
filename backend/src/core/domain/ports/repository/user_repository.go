package repository
import (
	"context"
	"blogapp/src/core/domain/entity"
)

type UserRepository interface{
	Create(ctx context.Context, user *entity.UserEntity) (*entity.UserEntity, error)
	GetByID(ctx context.Context, id uint) (*entity.UserEntity, error)
	GetByEmail(ctx context.Context, email string) (*entity.UserEntity, error)
	GetByUsername(ctx context.Context, username string) (*entity.UserEntity, error)
}




