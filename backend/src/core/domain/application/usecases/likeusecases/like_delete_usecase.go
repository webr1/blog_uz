package likeusecases

import (
	"blogapp/src/core/domain/ports/repository"
	"context"
)


type LikeDeleteUseCase struct{
	likeRepo repository.LikeRepository
}


func NewLikeDeleteUseCase(likeRepo repository.LikeRepository) *LikeDeleteUseCase{
	return  &LikeDeleteUseCase{
		likeRepo: likeRepo,
	}
}

func (uc *LikeDeleteUseCase) Invoke(ctx context.Context, likeID uint) error{
	return uc.likeRepo.Delete(ctx,likeID)
}

