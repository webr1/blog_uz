package postusecases

import (
	"blogapp/src/core/domain/entity"
	"blogapp/src/core/domain/ports/repository"
	"context"
)


type PostUseCase struct {
	postRepo repository.PostRepository
}

func NewPostUseCase(postRepo repository.PostRepository) *PostUseCase{
	return &PostUseCase{
		postRepo: postRepo,
	}
} 


func (uc *PostUseCase) Invoke(ctx context.Context, userID uint, title,content string) (*entity.PostEntity,error){
	
	post := entity.NewPostEntity(userID,title,content)
	var err error
	post,err = uc.postRepo.Create(ctx,post)
	if err != nil {
		return nil,err
	}
	return post,nil

}
