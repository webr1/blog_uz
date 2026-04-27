package groups

import (
	"blogapp/src/entrypoint/http/handlers/auth"
	"github.com/labstack/echo/v4"
)

type AuthGroup struct {
	registerHandler *auth.RegisterHandler
	loginHandler    *auth.LoginHandler
}

func NewAuthGroup(
	registerHandler *auth.RegisterHandler,
	loginHandler *auth.LoginHandler,
) *AuthGroup {
	return &AuthGroup{
		registerHandler: registerHandler,
		loginHandler:    loginHandler,
	}
}

func (g *AuthGroup) RegisterRoutes(e *echo.Echo) {
    authGroup := e.Group("/auth")
    authGroup.POST("/register", g.registerHandler.Invoke)
    authGroup.POST("/login", g.loginHandler.Login)
}