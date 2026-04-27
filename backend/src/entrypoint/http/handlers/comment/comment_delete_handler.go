package comment

import (
	"blogapp/src/core/domain/application/usecases/commentusecases"
	"net/http"
	"strconv"
	"github.com/labstack/echo/v4"
)

type CommentDeleteHandler struct {
	commentDeleteUC *commentusecases.CommentsDeleteUseCase
}

func NewCommentDeleteHandler(commentDeleteUC *commentusecases.CommentsDeleteUseCase) *CommentDeleteHandler {
	return &CommentDeleteHandler{commentDeleteUC: commentDeleteUC}
}


// @Summary      Delete comment
// @Tags         comments
// @Produce      json
// @Param        id   path      int    true  "Comment ID"
// @Success      200  {object}  object
// @Failure      400  {object}  object
// @Router       /comments/{id} [delete]
func (h *CommentDeleteHandler) Invoke(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
	}
	if err := h.commentDeleteUC.Invoke(c.Request().Context(), uint(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "deleted"})
}