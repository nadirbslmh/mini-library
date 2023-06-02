package http

import (
	"minilib/book/internal/service/book"
	"minilib/book/pkg/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	service *book.Service
}

func New() *Controller {
	return &Controller{
		service: book.New(),
	}
}

func (h *Controller) GetAll(c echo.Context) error {
	books := h.service.GetAll()

	return c.JSON(http.StatusOK, echo.Map{
		"data": books,
	})
}

func (h *Controller) Create(c echo.Context) error {
	var bookInput model.Book = model.Book{}

	c.Bind(&bookInput)

	book := h.service.Create(bookInput)

	return c.JSON(http.StatusCreated, echo.Map{
		"data": book,
	})
}
