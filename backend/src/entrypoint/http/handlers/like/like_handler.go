package like

import (
	"blogapp/src/core/domain/application/usecases/likeusecases"
	"net/http"
	"github.com/labstack/echo/v4"
)

type LikeCreateHandler struct {
	likeCreateUC *likeusecases.LikeUseCase
}

func NewLikeCreateHandler(likeCreateUC *likeusecases.LikeUseCase) *LikeCreateHandler {
	return &LikeCreateHandler{likeCreateUC: likeCreateUC}
}


// @Summary      Like post
// @Tags         likes
// @Accept       json
// @Produce      json
// @Param        body body object{user_id=integer,post_id=integer} true "Like"
// @Success      201  {object}  object
// @Failure      400  {object}  object
// @Router       /likes [post]
func (h *LikeCreateHandler) Invoke(c echo.Context) error {
	var req struct {
		UserID uint `json:"user_id"`
		PostID uint `json:"post_id"`
	}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}
	like, err := h.likeCreateUC.Invoke(c.Request().Context(), req.UserID, req.PostID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, like)
}
