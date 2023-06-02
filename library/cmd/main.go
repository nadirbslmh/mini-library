package main

import (
	"minilib/library/internal/controller/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	handler := http.New()

	e.GET("/books", handler.GetAll)
	e.POST("/books", handler.Create)

	e.Logger.Fatal(e.Start(":8080"))
}
