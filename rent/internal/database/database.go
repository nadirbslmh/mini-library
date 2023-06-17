package database

import (
	"fmt"
	"log"
	"pkg-service/constant"
	"rent-service/pkg/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// connect to the database
func InitDatabase(configs map[string]string) (*gorm.DB, error) {
	var err error

	var (
		DB_USERNAME string = configs[constant.RENT_DB_USERNAME]
		DB_PASSWORD string = configs[constant.RENT_DB_PASSWORD]
		DB_NAME     string = configs[constant.RENT_DB_NAME]
		DB_HOST     string = configs[constant.RENT_DB_HOST]
		DB_PORT     string = configs[constant.RENT_DB_PORT]
	)

	var dsn string = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		DB_USERNAME,
		DB_PASSWORD,
		DB_HOST,
		DB_PORT,
		DB_NAME,
	)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("error when creating a connection to the database: %s\n", err)
		return nil, err
	}

	log.Println("connected to the database")

	migrate()

	return DB, nil
}

// perform migration
func migrate() {
	err := DB.AutoMigrate(&model.Rent{})

	if err != nil {
		log.Fatalf("failed to perform database migration: %s\n", err)
	}
}
