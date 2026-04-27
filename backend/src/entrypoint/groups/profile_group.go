package groups

import (
	"blogapp/src/entrypoint/http/handlers/profile"
	"github.com/labstack/echo/v4"
)

type ProfileGroup struct {
	profileUpdateHandler *profile.ProfileUpdateHandler
}

func NewProfileGroup(profileUpdateHandler *profile.ProfileUpdateHandler) *ProfileGroup {
	return &ProfileGroup{profileUpdateHandler: profileUpdateHandler}
}

func (g *ProfileGroup) RegisterRoutes(e *echo.Group) {
    e.PUT("/:id", g.profileUpdateHandler.Invoke)
}