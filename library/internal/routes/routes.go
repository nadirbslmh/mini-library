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

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"

	"minilib/pkg/auth"
)

func SetupRoutes(e *echo.Echo, registry discovery.Registry) {
	jwtConfig := auth.NewDefaultConfig()
	authMiddlewareConfig := jwtConfig.Init()

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

	protectedEndpoints := e.Group("/api/v1", echojwt.WithConfig(authMiddlewareConfig))
	protectedEndpoints.Use(auth.VerifyToken)

	protectedEndpoints.GET("/books", bookController.GetAll)
	protectedEndpoints.GET("/books/:id", bookController.GetByID)
	protectedEndpoints.POST("/books", bookController.Create)

	protectedEndpoints.GET("/rents", rentController.GetAll)
	protectedEndpoints.POST("/rents", rentController.Create)
}
