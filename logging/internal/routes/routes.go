package routes

import (
	"logging-service/internal/controller/http"
	"logging-service/internal/database"
	"logging-service/internal/repository/mongo"
	"logging-service/internal/service/logging"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo, instance database.MongoInstance) {
	repository := mongo.New(instance)
	service := logging.New(repository)
	controller := http.New(service)

	e.POST("/logs", controller.Write)
}
