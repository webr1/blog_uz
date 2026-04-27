package postusecases

import (
	"blogapp/src/core/domain/entity"
	"blogapp/src/core/domain/ports/repository"
	"context"
)

type PostUpdateUseCase struct {
	updateRepo repository.PostRepository
}

func NewPostUpdateUseCase(updateID repository.PostRepository) *PostUpdateUseCase{
	return &PostUpdateUseCase{
		updateRepo: updateID,
	}
}


func (uc *PostUpdateUseCase) Invoke(ctx context.Context,postID uint,title,content string) (*entity.PostEntity,error){
	post, err := uc.updateRepo.GetByID(ctx, postID)
	if err != nil {
		return nil, err
	}

	post.Title = title
	post.Content = content

	return uc.updateRepo.Update(ctx, post)


}



