package database

import (
	"fmt"
	"log"
	"minilib/rent/pkg/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

var (
	DB_USERNAME string = "root"
	DB_PASSWORD string = ""
	DB_NAME     string = "lib_rent"
	DB_HOST     string = "localhost"
	DB_PORT     string = "3306"
)

// connect to the database
func InitDatabase() (*gorm.DB, error) {
	var err error

	var dsn string = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		DB_USERNAME,
		DB_PASSWORD,
		DB_HOST,
		DB_PORT,
		DB_NAME,
	)

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("error when creating a connection to the database: %s\n", err)
		return nil, err
	}

	log.Println("connected to the database")

	migrate()

	return db, nil
}

// perform migration
func migrate() {
	err := db.AutoMigrate(&model.Rent{})

	if err != nil {
		log.Fatalf("failed to perform database migration: %s\n", err)
	}
}
