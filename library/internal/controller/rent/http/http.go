package http

import (
	"context"
	"fmt"
	"minilib/library/internal/service/library"
	"minilib/library/pkg/constant"
	"minilib/pkg/auth"
	"net/http"
	"rent-service/pkg/model"
	"strconv"

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
	user, err := auth.GetUser(c)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to get all book rents data")
	}

	userId := strconv.Itoa(user.ID)

	ctxKey := constant.USER_ID_KEY

	ctx := context.WithValue(c.Request().Context(), ctxKey, userId)

	fmt.Println("new context: ", ctx.Value(ctxKey))

	rents, err := h.service.GetAllRents(ctx)

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
