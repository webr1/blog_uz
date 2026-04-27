package auth

import (
	"blogapp/src/core/domain/application/usecases/authusecases"
	"net/http"

	"github.com/labstack/echo/v4"
)



type RegisterHandler struct{
	registerUC *authusecases.RegisterUseCase
}

func NewRegisterHandler(registerUC *authusecases.RegisterUseCase) *RegisterHandler{
	return &RegisterHandler{registerUC: registerUC}
}

// @Summary      Register user
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        body body object{username=string,email=string,password=string} true "Register"
// @Success      201  {object}  object
// @Failure      400  {object}  object
// @Failure      500  {object}  object
// @Router       /auth/register [post]
func (h *RegisterHandler) Invoke(c echo.Context) error{
	var req struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.Bind(&req);err != nil{
		return c.JSON(http.StatusBadRequest, map[string]string{"request":"invalid request"})
	}
	user, err := h.registerUC.Invoke(c.Request().Context(),req.Username,req.Email,req.Password)
	if err != nil{
		return c.JSON(http.StatusInternalServerError,map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated,user)
}
