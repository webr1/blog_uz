package post

import (
	"blogapp/src/core/domain/application/usecases/postusecases"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type PostUpdateHandler struct {
	postUpdateUC *postusecases.PostUpdateUseCase
}

func NewPostUpdateHandler(postUpdateUC *postusecases.PostUpdateUseCase) *PostUpdateHandler {
	return &PostUpdateHandler{postUpdateUC: postUpdateUC}
}


// @Summary      Update post
// @Tags         posts
// @Accept       json
// @Produce      json
// @Param        id   path      int    true  "Post ID"
// @Param        body body object{title=string,content=string} true "Post"
// @Success      200  {object}  object
// @Failure      400  {object}  object
// @Router       /posts/{id} [put]
func (h *PostUpdateHandler) Invoke(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
	}
	var req struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}
	post, err := h.postUpdateUC.Invoke(c.Request().Context(), uint(id), req.Title, req.Content)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, post)
}