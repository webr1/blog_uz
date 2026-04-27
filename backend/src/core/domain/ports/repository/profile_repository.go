package repository


import (
	"context"
	"blogapp/src/core/domain/entity"

)



type ProfileRepository interface{
	Create(ctx context.Context, post *entity.ProfileEntity) (*entity.ProfileEntity,error)
	GetByUserID(ctx context.Context,id uint) (*entity.ProfileEntity,error)
	Update(ctx context.Context,profile *entity.ProfileEntity) (*entity.ProfileEntity, error)
}


