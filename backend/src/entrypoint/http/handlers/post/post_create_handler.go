package post

import (
	"blogapp/src/core/domain/application/usecases/postusecases"
	"net/http"

	"github.com/labstack/echo/v4"
)


type PostHandler struct{
	postUC *postusecases.PostUseCase
}

func NewPostHandler(postUC *postusecases.PostUseCase) *PostHandler{
	return &PostHandler{
		postUC: postUC,
	}
}

// @Summary      Create post
// @Tags         posts
// @Accept       json
// @Produce      json
// @Param        body body object{title=string,content=string,user_id=integer} true "Post"
// @Success      200  {object}  object
// @Failure      400  {object}  object
// @Router       /posts [post]
func (h *PostHandler) Invoke(c echo.Context) error {
	var req struct{
		Title   string `json:"title"`
		Content string `json:"content"`
		UserID  uint   `json:"user_id"`
	}
	if err := c.Bind(&req); err != nil{
		return c.JSON(http.StatusBadRequest,map[string]string{"error":"invalid content"})
	}
	post, err :=h.postUC.Invoke(c.Request().Context(),req.UserID,req.Title,req.Content)
	if err != nil{
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK,post)

}