package likeusecases

import (
	"blogapp/src/core/domain/entity"
	"blogapp/src/core/domain/ports/repository"
	"context"
)


type LikeUseCase struct {
	likeRepo repository.LikeRepository

}



func NewLikeUseCase(likeRepo repository.LikeRepository) *LikeUseCase{
	return &LikeUseCase{
		likeRepo: likeRepo,
	}
}


func (uc *LikeUseCase) Invoke(ctx context.Context,userID,postID uint) (*entity.LikeEntity,error){
	like := entity.NewLikeEntity(userID, postID)

	var err error
	like, err = uc.likeRepo.Create(ctx, like)
	if err != nil {
		return nil, err
	}
	return like, nil

}