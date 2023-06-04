package routes

import (
	"book-service/internal/controller/http"
	"book-service/internal/repository/mysql"
	"book-service/internal/service/book"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func SetupRoutes(e *echo.Echo, db *gorm.DB) {
	repository := mysql.New(db)
	service := book.New(repository)
	controller := http.New(service)

	e.GET("/books", controller.GetAll)
	e.GET("/books/:id", controller.GetByID)
	e.POST("/books", controller.Create)
}
