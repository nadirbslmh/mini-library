package http

import (
	"minilib/rent/internal/service/rent"
	"minilib/rent/pkg/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	service *rent.Service
}

func New() *Controller {
	return &Controller{
		service: rent.New(),
	}
}

func (h *Controller) GetAll(c echo.Context) error {
	rents := h.service.GetAll()

	return c.JSON(http.StatusOK, echo.Map{
		"data": rents,
	})
}

func (h *Controller) Create(c echo.Context) error {
	var rentInput model.Rent = model.Rent{}

	c.Bind(&rentInput)

	rent := h.service.Create(rentInput)

	return c.JSON(http.StatusCreated, echo.Map{
		"data": rent,
	})
}
