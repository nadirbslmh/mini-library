package http

import (
	"minilib/library/internal/service/library"
	"minilib/pkg/auth"
	"minilib/rent/pkg/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	service *library.RentService
}

func New(service *library.RentService) *Controller {
	return &Controller{
		service: service,
	}
}

func (h *Controller) GetAll(c echo.Context) error {
	rents, err := h.service.GetAllRents(c.Request().Context())

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to get all book rents data")
	}

	return c.JSON(http.StatusOK, rents)
}

func (h *Controller) Create(c echo.Context) error {
	var rentInput model.RentInput

	c.Bind(&rentInput)

	user, err := auth.GetUser(c)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to create book rent")
	}

	rentInput.UserID = user.ID

	rent, err := h.service.CreateRent(c.Request().Context(), rentInput)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to create book rent")
	}

	return c.JSON(http.StatusCreated, rent)
}
