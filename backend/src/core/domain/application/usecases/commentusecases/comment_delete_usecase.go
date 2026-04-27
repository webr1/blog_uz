package commentusecases

import (
	"blogapp/src/core/domain/ports/repository"
	"context"
)

type CommentsDeleteUseCase struct{
	commentRepo repository.CommentRepository
}


func NewCommentsDeleteUseCase(commentRepo repository.CommentRepository) *CommentsDeleteUseCase{
	return &CommentsDeleteUseCase{
		commentRepo:commentRepo,
	}
}

func (uc*CommentsDeleteUseCase) Invoke(ctx context.Context,commentID uint) error{
	return uc.commentRepo.Delete(ctx,commentID)
}