package http

import (
	"minilib/book/pkg/model"
	"minilib/library/internal/service/library"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	service *library.Service
}

func New(service *library.Service) *Controller {
	return &Controller{
		service: service,
	}
}

func (h *Controller) GetAll(c echo.Context) error {
	books, err := h.service.GetAll(c.Request().Context())

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to get all books")
	}

	return c.JSON(http.StatusOK, books)
}

func (h *Controller) GetByID(c echo.Context) error {
	bookId := c.Param("id")

	book, err := h.service.GetByID(c.Request().Context(), bookId)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to get book data")
	}

	return c.JSON(http.StatusOK, book)
}

func (h *Controller) Create(c echo.Context) error {
	var bookInput model.BookInput

	c.Bind(&bookInput)

	book, err := h.service.Create(c.Request().Context(), bookInput)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to create book")
	}

	return c.JSON(http.StatusCreated, book)
}
