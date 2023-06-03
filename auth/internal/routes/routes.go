package routes

import (
	"minilib/auth/internal/controller/http"
	"minilib/auth/internal/repository/mysql"
	"minilib/auth/internal/service/auth"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func SetupRoutes(e *echo.Echo, db *gorm.DB) {
	repository := mysql.New(db)
	service := auth.New(repository)
	controller := http.New(service)

	e.POST("/auth/register", controller.Register)
	e.POST("/auth/login", controller.Login)
}
