package routes

import (
	"minilib/book/internal/controller/http"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
	controller := http.New()

	e.GET("/books", controller.GetAll)
	e.POST("/books", controller.Create)
}
