package http

import (
	responsemodel "minilib/pkg/model"
	"minilib/rent/internal/service/rent"
	"minilib/rent/pkg/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	service *rent.RentService
}

func New(service *rent.RentService) *Controller {
	return &Controller{
		service: rent.New(service),
	}
}

func (ctrl *Controller) GetAll(c echo.Context) error {
	rents, err := ctrl.service.GetAll()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, responsemodel.Response{
			Status:  "failed",
			Message: "failed to fetch book rents",
		})
	}

	return c.JSON(http.StatusOK, responsemodel.Response{
		Status:  "success",
		Message: "all book rents",
		Data:    rents,
	})
}

func (ctrl *Controller) Create(c echo.Context) error {
	var rentInput model.RentInput = model.RentInput{}

	if err := c.Bind(&rentInput); err != nil {
		return c.JSON(http.StatusBadRequest, responsemodel.Response{
			Status:  "failed",
			Message: "invalid request",
		})
	}

	err := rentInput.Validate()

	if err != nil {
		return c.JSON(http.StatusBadRequest, responsemodel.Response{
			Status:  "failed",
			Message: "invalid request",
		})
	}

	rent, err := ctrl.service.Create(rentInput)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, responsemodel.Response{
			Status:  "failed",
			Message: "failed to create a book rent data",
		})
	}

	return c.JSON(http.StatusCreated, responsemodel.Response{
		Status:  "success",
		Message: "rent created",
		Data:    rent,
	})
}
