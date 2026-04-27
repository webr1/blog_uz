package comment

import (
	"blogapp/src/core/domain/application/usecases/commentusecases"
	"net/http"
	"github.com/labstack/echo/v4"
)

type CommentCreateHandler struct {
	commentCreateUC *commentusecases.CommentsCreateUseCase
}

func NewCommentCreateHandler(commentCreateUC *commentusecases.CommentsCreateUseCase) *CommentCreateHandler {
	return &CommentCreateHandler{commentCreateUC: commentCreateUC}
}


// @Summary      Create comment
// @Tags         comments
// @Accept       json
// @Produce      json
// @Param        body body object{post_id=integer,user_id=integer,content=string} true "Comment"
// @Success      201  {object}  object
// @Failure      400  {object}  object
// @Router       /comments [post]
func (h *CommentCreateHandler) Invoke(c echo.Context) error {
	var req struct {
		PostID  uint   `json:"post_id"`
		UserID  uint   `json:"user_id"`
		Content string `json:"content"`
	}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}
	comment, err := h.commentCreateUC.Invoke(c.Request().Context(), req.PostID, req.UserID, req.Content)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, comment)
}