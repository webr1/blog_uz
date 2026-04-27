package auth

import (
	"blogapp/src/core/domain/application/usecases/authusecases"
	"github.com/labstack/echo/v4"
)
type LoginHandler struct{
	loginUC *authusecases.LoginUseCase
}
func NewLoginHandler(loginUC *authusecases.LoginUseCase) *LoginHandler{
	return &LoginHandler{
		loginUC: loginUC,
	}
}


// @Summary      Login user
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        body body object{email=string,password=string} true "Login"
// @Success      200  {object}  object
// @Failure      400  {object}  object
// @Failure      401  {object}  object
// @Router       /auth/login [post]
func (h *LoginHandler) Login(c echo.Context) error {
    var req struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    }

    if err := c.Bind(&req); err != nil {
        return c.JSON(400, map[string]string{"error": "invalid request"})
    }

    token, user, err := h.loginUC.Invoke(c.Request().Context(), req.Email, req.Password)
    if err != nil {
        if err.Error() == "invalid password" {
            return c.JSON(401, map[string]string{"error": "invalid email or password"})
        }
        return c.JSON(500, map[string]string{"error": err.Error()})
    }

    return c.JSON(200, map[string]interface{}{
        "token": token,
        "user":  user,
    })
}