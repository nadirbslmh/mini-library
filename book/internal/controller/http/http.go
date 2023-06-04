package http

import (
	"book-service/internal/service/book"
	"book-service/pkg/model"
	"net/http"
	responsemodel "pkg-service/model"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	service *book.BookService
}

func New(service *book.BookService) *Controller {
	return &Controller{
		service: service,
	}
}

func (ctrl *Controller) GetAll(c echo.Context) error {
	books, err := ctrl.service.GetAll()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, responsemodel.Response[any]{
			Status:  "failed",
			Message: "failed to fetch books",
		})
	}

	return c.JSON(http.StatusOK, responsemodel.Response[[]model.Book]{
		Status:  "success",
		Message: "all books",
		Data:    books,
	})
}

func (ctrl *Controller) GetByID(c echo.Context) error {
	bookId := c.Param("id")

	book, err := ctrl.service.GetByID(bookId)

	if err != nil {
		return c.JSON(http.StatusNotFound, responsemodel.Response[any]{
			Status:  "failed",
			Message: "book not found",
		})
	}

	return c.JSON(http.StatusOK, responsemodel.Response[model.Book]{
		Status:  "success",
		Message: "book data",
		Data:    book,
	})
}

func (ctrl *Controller) Create(c echo.Context) error {
	var bookInput model.BookInput = model.BookInput{}

	if err := c.Bind(&bookInput); err != nil {
		return c.JSON(http.StatusBadRequest, responsemodel.Response[any]{
			Status:  "failed",
			Message: "invalid request",
		})
	}

	err := bookInput.Validate()

	if err != nil {
		return c.JSON(http.StatusBadRequest, responsemodel.Response[any]{
			Status:  "failed",
			Message: "invalid request",
		})
	}

	book, err := ctrl.service.Create(bookInput)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, responsemodel.Response[any]{
			Status:  "failed",
			Message: "failed to create book",
		})
	}

	return c.JSON(http.StatusCreated, responsemodel.Response[model.Book]{
		Status:  "success",
		Message: "book created",
		Data:    book,
	})
}
