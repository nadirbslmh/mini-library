package routes

import (
	"minilib/library/internal/controller/http"
	"minilib/pkg/discovery"

	bookgateway "minilib/library/internal/gateway/book/http"
	"minilib/library/internal/service/library"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo, registry discovery.Registry) {
	bookGateway := bookgateway.New(registry)
	bookService := library.New(*bookGateway)
	bookController := http.New(bookService)

	endpoints := e.Group("/api/v1")

	endpoints.GET("/books", bookController.GetAll)
	endpoints.POST("/books", bookController.Create)
}
