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

func (ctrl *Controller) GetAll(c echo.Context) error {
	rents := ctrl.service.GetAll()

	return c.JSON(http.StatusOK, echo.Map{
		"data": rents,
	})
}

func (ctrl *Controller) Create(c echo.Context) error {
	var rentInput model.Rent = model.Rent{}

	c.Bind(&rentInput)

	rent := ctrl.service.Create(rentInput)

	return c.JSON(http.StatusCreated, echo.Map{
		"data": rent,
	})
}
