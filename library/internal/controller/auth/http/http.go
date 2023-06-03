package http

import (
	"minilib/auth/pkg/model"
	"minilib/library/internal/service/library"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	service *library.AuthService
}

func New(service *library.AuthService) *Controller {
	return &Controller{
		service: service,
	}
}

func (h *Controller) Register(c echo.Context) error {
	var UserInput model.UserInput

	c.Bind(&UserInput)

	user, err := h.service.Register(c.Request().Context(), UserInput)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to create a user")
	}

	return c.JSON(http.StatusCreated, user)
}

func (h *Controller) Login(c echo.Context) error {
	var UserInput model.UserInput

	c.Bind(&UserInput)

	data, err := h.service.Login(c.Request().Context(), UserInput)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to login, wrong username or password")
	}

	return c.JSON(http.StatusOK, data)
}
