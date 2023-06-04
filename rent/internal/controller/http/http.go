package http

import (
	"net/http"
	responsemodel "pkg-service/model"
	"rent-service/internal/service/rent"
	"rent-service/pkg/model"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	service *rent.Service
}

func New(service *rent.Service) *Controller {
	return &Controller{
		service: service,
	}
}

func (ctrl *Controller) GetAll(c echo.Context) error {
	userId := c.Param("id")

	rents, err := ctrl.service.GetAll(userId)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, responsemodel.Response[any]{
			Status:  "failed",
			Message: "failed to fetch book rents data",
		})
	}

	return c.JSON(http.StatusOK, responsemodel.Response[[]model.Rent]{
		Status:  "success",
		Message: "all book rents data",
		Data:    rents,
	})
}

func (ctrl *Controller) Create(c echo.Context) error {
	var rentInput model.RentInput

	if err := c.Bind(&rentInput); err != nil {
		return c.JSON(http.StatusBadRequest, responsemodel.Response[any]{
			Status:  "failed",
			Message: "invalid request",
		})
	}

	err := rentInput.Validate()

	if err != nil {
		return c.JSON(http.StatusBadRequest, responsemodel.Response[any]{
			Status:  "failed",
			Message: "invalid request",
		})
	}

	rent, err := ctrl.service.Create(rentInput)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, responsemodel.Response[any]{
			Status:  "failed",
			Message: "failed to create a book rent",
		})
	}

	return c.JSON(http.StatusCreated, responsemodel.Response[model.Rent]{
		Status:  "success",
		Message: "book rent created",
		Data:    rent,
	})
}
