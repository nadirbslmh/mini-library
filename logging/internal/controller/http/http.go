package http

import (
	"logging-service/internal/service/logging"
	"logging-service/pkg/model"
	"net/http"
	responsemodel "pkg-service/model"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	service *logging.LogService
}

func New(service *logging.LogService) *Controller {
	return &Controller{
		service: service,
	}
}

func (ctrl *Controller) Write(c echo.Context) error {
	var logInput model.LogInput = model.LogInput{}

	if err := c.Bind(&logInput); err != nil {
		return c.JSON(http.StatusBadRequest, responsemodel.Response[any]{
			Status:  "failed",
			Message: "invalid request",
		})
	}

	err := logInput.Validate()

	if err != nil {
		return c.JSON(http.StatusBadRequest, responsemodel.Response[any]{
			Status:  "failed",
			Message: "invalid request",
		})
	}

	log, err := ctrl.service.Write(logInput)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, responsemodel.Response[any]{
			Status:  "failed",
			Message: "failed to write log",
		})
	}

	return c.JSON(http.StatusCreated, responsemodel.Response[model.Log]{
		Status:  "success",
		Message: "log written",
		Data:    *log,
	})
}
