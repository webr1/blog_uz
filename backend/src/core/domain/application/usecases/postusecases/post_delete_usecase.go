package postusecases

import (
	"blogapp/src/core/domain/ports/repository"
	"context"
)

type PostDeleteUseCase struct {
	postRepo repository.PostRepository
}

func NewPostDeleteUseCase(postRepo repository.PostRepository) *PostDeleteUseCase {
    return &PostDeleteUseCase{
        postRepo: postRepo,
    }
}

func (uc *PostDeleteUseCase) Invoke(ctx context.Context,postID uint) (error){
	return uc.postRepo.Delete(ctx, postID)
}


