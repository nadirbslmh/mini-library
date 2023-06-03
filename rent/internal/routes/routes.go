package routes

import (
	"minilib/rent/internal/controller/http"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
	controller := http.New()

	e.GET("/rents", controller.GetAll)
	e.POST("/rents", controller.Create)
}
