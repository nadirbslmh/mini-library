package routes

import (
	authcontroller "library-service/internal/controller/auth/http"
	bookcontroller "library-service/internal/controller/book/http"
	rentcontroller "library-service/internal/controller/rent/http"
	"pkg-service/discovery"

	authgateway "library-service/internal/gateway/auth/grpc"
	bookgateway "library-service/internal/gateway/book/grpc"
	rentgateway "library-service/internal/gateway/rent/grpc"
	"library-service/internal/service/library"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"

	"pkg-service/auth"
)

func SetupRoutes(e *echo.Echo, registry discovery.Registry, producer *kafka.Producer) {
	jwtConfig := auth.NewDefaultConfig()
	authMiddlewareConfig := jwtConfig.Init()

	bookGateway := bookgateway.New(registry)
	bookService := library.NewBookService(*bookGateway)
	bookController := bookcontroller.New(bookService)

	rentGateway := rentgateway.New(registry, bookGateway, producer)
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
