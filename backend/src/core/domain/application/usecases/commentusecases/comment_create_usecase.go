package commentusecases

import (
	"blogapp/src/core/domain/entity"
	"blogapp/src/core/domain/ports/repository"
	"context"
)

type CommentsCreateUseCase struct{
	commentRepo repository.CommentRepository
}

func NewCommentsCreateUseCase(commentRepo repository.CommentRepository) *CommentsCreateUseCase{
	return	&CommentsCreateUseCase{
		commentRepo: commentRepo,
	}
}

func (uc *CommentsCreateUseCase) Invoke(ctx context.Context, postID, userID uint, content string)(*entity.CommentEntity,error){
	comment := entity.NewCommentEntity(postID, userID, content)
	return uc.commentRepo.Create(ctx, comment)
}