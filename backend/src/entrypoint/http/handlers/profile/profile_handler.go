package profile


import (
	"blogapp/src/core/domain/application/usecases/profileusecases"
	"net/http"
	"strconv"
	"github.com/labstack/echo/v4"
)

type ProfileUpdateHandler struct {
	profileUC *profileusecases.ProfileUseCase
}

func NewProfileUpdateHandler(profileUC *profileusecases.ProfileUseCase) *ProfileUpdateHandler {
	return &ProfileUpdateHandler{profileUC: profileUC}
}


// @Summary      Update profile
// @Tags         profile
// @Accept       json
// @Produce      json
// @Param        id   path      int    true  "User ID"
// @Param        body body object{full_name=string,bio=string,avatar=string} true "Profile"
// @Success      200  {object}  object
// @Failure      400  {object}  object
// @Router       /profile/{id} [put]
func (h *ProfileUpdateHandler) Invoke(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
	}
	var req struct {
		FullName string `json:"full_name"`
		Bio      string `json:"bio"`
		Avatar   string `json:"avatar"`
	}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}
	profile, err := h.profileUC.Invoke(c.Request().Context(), uint(id), req.FullName, req.Bio, req.Avatar)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, profile)
}