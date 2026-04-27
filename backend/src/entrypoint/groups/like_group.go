package groups

import (
	"blogapp/src/entrypoint/http/handlers/like"
	"github.com/labstack/echo/v4"
)

type LikeGroup struct {
	likeCreateHandler *like.LikeCreateHandler
	likeDeleteHandler *like.LikeDeleteHandler
}

func NewLikeGroup(
	likeCreateHandler *like.LikeCreateHandler,
	likeDeleteHandler *like.LikeDeleteHandler,
) *LikeGroup {
	return &LikeGroup{
		likeCreateHandler: likeCreateHandler,
		likeDeleteHandler: likeDeleteHandler,
	}
}

func (g *LikeGroup) RegisterRoutes(e *echo.Group) {
    e.POST("", g.likeCreateHandler.Invoke)
    e.DELETE("/:id", g.likeDeleteHandler.Invoke)
}