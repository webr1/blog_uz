package post

import (
	"blogapp/src/core/domain/application/usecases/postusecases"
	"net/http"

	"github.com/labstack/echo/v4"
)


type PostListHandler struct{
	postListUC *postusecases.PostListUseCase
}

func NewPostListHandler(postListUC *postusecases.PostListUseCase) *PostListHandler{
	return &PostListHandler{
		postListUC: postListUC,
	}
}


// @Summary      List posts
// @Tags         posts
// @Produce      json
// @Success      200  {array}   object
// @Failure      500  {object}  object
// @Router       /posts [get]
func (h *PostListHandler) Invoke(c echo.Context) error{

	posts,err := h.postListUC.Invoke(c.Request().Context())
	if err != nil{
		return c.JSON(http.StatusInternalServerError,map[string]string{"error":err.Error()})
	}
	return c.JSON(http.StatusOK,posts)

}