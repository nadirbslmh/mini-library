package routes

import (
	"minilib/rent/internal/controller/http"
	"minilib/rent/internal/repository/mysql"
	"minilib/rent/internal/service/rent"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func SetupRoutes(e *echo.Echo, db *gorm.DB) {
	repository := mysql.New(db)
	service := rent.New(repository)
	controller := http.New(service)

	e.GET("/rents", controller.GetAll)
	e.POST("/rents", controller.Create)
}
