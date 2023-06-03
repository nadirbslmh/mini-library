package http

import (
	"minilib/auth/internal/service/auth"
	"minilib/auth/pkg/model"
	responsemodel "minilib/pkg/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	service *auth.AuthService
}

func New(service *auth.AuthService) *Controller {
	return &Controller{
		service: service,
	}
}

func (ctrl *Controller) Register(c echo.Context) error {
	var userInput model.UserInput = model.UserInput{}

	if err := c.Bind(&userInput); err != nil {
		return c.JSON(http.StatusBadRequest, responsemodel.Response[any]{
			Status:  "failed",
			Message: "invalid request",
		})
	}

	err := userInput.Validate()

	if err != nil {
		return c.JSON(http.StatusBadRequest, responsemodel.Response[any]{
			Status:  "failed",
			Message: "invalid request",
		})
	}

	auth, err := ctrl.service.Register(userInput)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, responsemodel.Response[any]{
			Status:  "failed",
			Message: "failed to create user",
		})
	}

	return c.JSON(http.StatusCreated, responsemodel.Response[model.User]{
		Status:  "success",
		Message: "user created",
		Data:    auth,
	})
}

func (ctrl *Controller) Login(c echo.Context) error {
	var userInput model.UserInput = model.UserInput{}

	if err := c.Bind(&userInput); err != nil {
		return c.JSON(http.StatusBadRequest, responsemodel.Response[any]{
			Status:  "failed",
			Message: "invalid request",
		})
	}

	err := userInput.Validate()

	if err != nil {
		return c.JSON(http.StatusBadRequest, responsemodel.Response[any]{
			Status:  "failed",
			Message: "invalid request",
		})
	}

	token, err := ctrl.service.Login(userInput)

	if err != nil {
		return c.JSON(http.StatusUnauthorized, responsemodel.Response[any]{
			Status:  "failed",
			Message: "username or password is invalid",
		})
	}

	return c.JSON(http.StatusOK, responsemodel.Response[string]{
		Status:  "success",
		Message: "login success",
		Data:    token,
	})
}
