package groups

import (
	"blogapp/src/entrypoint/http/handlers/post"

	"github.com/labstack/echo/v4"
)

type PostGroup struct {
	postCreateHandler *post.PostHandler
	postListHandler   *post.PostListHandler
	postUpdateHandler *post.PostUpdateHandler
	postDeleteHandler *post.PostDeleteHandler
}

func NewPostGroup(
	postCreateHandler *post.PostHandler,
	postListHandler *post.PostListHandler,
	postUpdateHandler *post.PostUpdateHandler,
	postDeleteHandler *post.PostDeleteHandler,
) *PostGroup {
	return &PostGroup{
		postCreateHandler: postCreateHandler,
		postListHandler:   postListHandler,
		postUpdateHandler: postUpdateHandler,
		postDeleteHandler: postDeleteHandler,
	}
}

func (g *PostGroup) RegisterRoutes(e *echo.Group) {
	e.POST("", g.postCreateHandler.Invoke)
	e.GET("", g.postListHandler.Invoke)
	e.PUT("/:id", g.postUpdateHandler.Invoke)
	e.DELETE("/:id", g.postDeleteHandler.Invoke)
}
