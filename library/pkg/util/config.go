package util

import (
	"errors"
	"log"
	"os"
	"pkg-service/constant"
	"pkg-service/util"

	"github.com/hashicorp/consul/api"
	"github.com/spf13/viper"
)

func GetConfig(key string) string {
	var fromEnv string = os.Getenv(key)

	if fromEnv != "" {
		return fromEnv
	}

	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("error when reading configuration file: %s\n", err)
	}

	return viper.GetString(key)
}

func InitializeConfigs(client *api.Client) error {
	authServiceConfig := map[string]string{
		constant.AUTH_DB_USERNAME: GetConfig(constant.AUTH_DB_USERNAME),
		constant.AUTH_DB_PASSWORD: GetConfig(constant.AUTH_DB_PASSWORD),
		constant.AUTH_DB_NAME:     GetConfig(constant.AUTH_DB_NAME),
		constant.AUTH_DB_HOST:     GetConfig(constant.AUTH_DB_HOST),
		constant.AUTH_DB_PORT:     GetConfig(constant.AUTH_DB_PORT),
	}

	bookServiceConfig := map[string]string{
		constant.BOOK_DB_USERNAME: GetConfig(constant.BOOK_DB_USERNAME),
		constant.BOOK_DB_PASSWORD: GetConfig(constant.BOOK_DB_PASSWORD),
		constant.BOOK_DB_NAME:     GetConfig(constant.BOOK_DB_NAME),
		constant.BOOK_DB_HOST:     GetConfig(constant.BOOK_DB_HOST),
		constant.BOOK_DB_PORT:     GetConfig(constant.BOOK_DB_PORT),
	}

	rentServiceConfig := map[string]string{
		constant.RENT_DB_USERNAME: GetConfig(constant.RENT_DB_USERNAME),
		constant.RENT_DB_PASSWORD: GetConfig(constant.RENT_DB_PASSWORD),
		constant.RENT_DB_NAME:     GetConfig(constant.RENT_DB_NAME),
		constant.RENT_DB_HOST:     GetConfig(constant.RENT_DB_HOST),
		constant.RENT_DB_PORT:     GetConfig(constant.RENT_DB_PORT),
	}

	logServiceConfig := map[string]string{
		constant.LOG_DB_NAME:   GetConfig(constant.LOG_DB_NAME),
		constant.LOG_MONGO_URI: GetConfig(constant.LOG_MONGO_URI),
	}

	authErr := util.SetBatchConfigValues(client, authServiceConfig)
	bookErr := util.SetBatchConfigValues(client, bookServiceConfig)
	rentErr := util.SetBatchConfigValues(client, rentServiceConfig)
	logErr := util.SetBatchConfigValues(client, logServiceConfig)

	isFailed := authErr != nil || bookErr != nil || rentErr != nil || logErr != nil

	if isFailed {
		log.Println("error occurred in auth: ", authErr)
		log.Println("error occurred in book: ", bookErr)
		log.Println("error occurred in rent: ", rentErr)
		log.Println("error occurred in log: ", logErr)

		return errors.New("config initialization failed")
	}

	return nil
}
