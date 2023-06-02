package routes

import (
	"minilib/library/internal/controller/http"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
	handler := http.New()

	book := e.Group("/api/v1")

	book.GET("/books", handler.GetAll)
	book.POST("/books", handler.Create)
}
