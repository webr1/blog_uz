package like

import (
	"blogapp/src/core/domain/application/usecases/likeusecases"
	"net/http"
	"strconv"
	"github.com/labstack/echo/v4"
)


type LikeDeleteHandler struct {
	likeDeleteUC *likeusecases.LikeDeleteUseCase
}

func NewLikeDeleteHandler(likeDeleteUC *likeusecases.LikeDeleteUseCase) *LikeDeleteHandler {
	return &LikeDeleteHandler{likeDeleteUC: likeDeleteUC}
}


// @Summary      Unlike post
// @Tags         likes
// @Produce      json
// @Param        id   path      int    true  "Like ID"
// @Success      200  {object}  object
// @Failure      400  {object}  object
// @Router       /likes/{id} [delete]
func (h *LikeDeleteHandler) Invoke(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
	}
	if err := h.likeDeleteUC.Invoke(c.Request().Context(), uint(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "deleted"})
}