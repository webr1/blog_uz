package post

import (
	"blogapp/src/core/domain/application/usecases/postusecases"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type PostDeleteHandler struct {
	postDeleteUC *postusecases.PostDeleteUseCase
}

func NewPostDeleteHandler(postDeleteUC *postusecases.PostDeleteUseCase) *PostDeleteHandler {
	return &PostDeleteHandler{postDeleteUC: postDeleteUC}
}


// @Summary      Delete post
// @Tags         posts
// @Produce      json
// @Param        id   path      int    true  "Post ID"
// @Success      200  {object}  object
// @Failure      400  {object}  object
// @Router       /posts/{id} [delete]
func (h *PostDeleteHandler) Invoke(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
	}
	if err := h.postDeleteUC.Invoke(c.Request().Context(), uint(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "deleted"})
}