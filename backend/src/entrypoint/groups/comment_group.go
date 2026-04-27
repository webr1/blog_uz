package groups

import (
	"blogapp/src/entrypoint/http/handlers/comment"
	"github.com/labstack/echo/v4"
)

type CommentGroup struct {
	commentCreateHandler *comment.CommentCreateHandler
	commentDeleteHandler *comment.CommentDeleteHandler
}

func NewCommentGroup(
	commentCreateHandler *comment.CommentCreateHandler,
	commentDeleteHandler *comment.CommentDeleteHandler,
) *CommentGroup {
	return &CommentGroup{
		commentCreateHandler: commentCreateHandler,
		commentDeleteHandler: commentDeleteHandler,
	}
}

func (g *CommentGroup) RegisterRoutes(e *echo.Group) {
    e.POST("", g.commentCreateHandler.Invoke)
    e.DELETE("/:id", g.commentDeleteHandler.Invoke)
}