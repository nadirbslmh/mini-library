package routes

import (
	authcontroller "library-service/internal/controller/auth/http"
	bookcontroller "library-service/internal/controller/book/http"
	rentcontroller "library-service/internal/controller/rent/http"
	"log"
	"pkg-service/discovery"

	authgateway "library-service/internal/gateway/auth/grpc"
	bookgateway "library-service/internal/gateway/book/grpc"
	rentgateway "library-service/internal/gateway/rent/grpc"
	"library-service/internal/service/library"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/labstack/echo-contrib/echoprometheus"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus"

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

	// setup prometheus
	customCounter := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "custom_requests_total",
			Help: "How many HTTP requests processed, partitioned by path and HTTP method.",
		},
		[]string{"path", "method"},
	)
	if err := prometheus.Register(customCounter); err != nil {
		log.Fatal(err)
	}

	apiRequestDuration := prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "api_request_duration_seconds",
			Help:    "Histogram for the request duration of the public API, partitioned by path and HTTP method.",
			Buckets: prometheus.ExponentialBuckets(0.1, 1.5, 5),
		},
		[]string{"path", "method"},
	)

	if err := prometheus.Register(apiRequestDuration); err != nil {
		log.Fatal(err)
	}

	e.Use(echoprometheus.NewMiddlewareWithConfig(echoprometheus.MiddlewareConfig{
		AfterNext: func(c echo.Context, err error) {
			customCounter.WithLabelValues(c.Path(), c.Request().Method).Inc()
		},
	}))

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			timer := prometheus.NewTimer(prometheus.ObserverFunc(func(v float64) {
				apiRequestDuration.WithLabelValues(c.Path(), c.Request().Method).Observe(v)
			}))
			defer timer.ObserveDuration()

			return next(c)
		}
	})

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

	e.GET("/metrics", echoprometheus.NewHandler())
}
