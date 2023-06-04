package routes

import (
	"rent-service/internal/controller/http"
	"rent-service/internal/repository/mysql"
	"rent-service/internal/service/rent"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func SetupRoutes(e *echo.Echo, db *gorm.DB) {
	repository := mysql.New(db)
	service := rent.New(repository)
	controller := http.New(service)

	e.GET("/rents/:id", controller.GetAll)
	e.POST("/rents", controller.Create)
}
