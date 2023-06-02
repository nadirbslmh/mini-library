package main

import (
	"minilib/rent/internal/controller/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	handler := http.New()

	e.GET("/rents", handler.GetAll)
	e.POST("/rents", handler.Create)

	e.Logger.Fatal(e.Start(":8082"))
}
