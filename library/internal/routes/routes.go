package routes

import (
	authcontroller "minilib/library/internal/controller/auth/http"
	bookcontroller "minilib/library/internal/controller/book/http"
	rentcontroller "minilib/library/internal/controller/rent/http"
	"minilib/pkg/discovery"

	authgateway "minilib/library/internal/gateway/auth/http"
	bookgateway "minilib/library/internal/gateway/book/http"
	rentgateway "minilib/library/internal/gateway/rent/http"
	"minilib/library/internal/service/library"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo, registry discovery.Registry) {
	bookGateway := bookgateway.New(registry)
	bookService := library.NewBookService(*bookGateway)
	bookController := bookcontroller.New(bookService)

	rentGateway := rentgateway.New(registry, bookGateway)
	rentService := library.NewRentService(*rentGateway)
	rentController := rentcontroller.New(rentService)

	authGateway := authgateway.New(registry)
	authService := library.NewAuthService(*authGateway)
	authController := authcontroller.New(authService)

	endpoints := e.Group("/api/v1")

	endpoints.POST("/register", authController.Register)
	endpoints.POST("/login", authController.Login)

	endpoints.GET("/books", bookController.GetAll)
	endpoints.GET("/books/:id", bookController.GetByID)
	endpoints.POST("/books", bookController.Create)

	endpoints.GET("/rents", rentController.GetAll)
	endpoints.POST("/rents", rentController.Create)
}
