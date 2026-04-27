package postusecases

import (
	"blogapp/src/core/domain/entity"
	"blogapp/src/core/domain/ports/repository"
	"context"
)


type PostListUseCase struct {
	postRepo repository.PostRepository
}

func NewPostListUseCase(postRepo repository.PostRepository) *PostListUseCase{
	return &PostListUseCase{
		postRepo: postRepo,
	}
} 


func (uc *PostListUseCase) Invoke(ctx context.Context) ([]*entity.PostEntity,error){
	return uc.postRepo.GetAll(ctx)
}
