package main

import (
	"minilib/book/internal/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	routes.SetupRoutes(e)

	e.Logger.Fatal(e.Start(":8081"))
}
