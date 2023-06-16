package util

import (
	"errors"
	"log"
	"pkg-service/constant"
	"pkg-service/util"

	"github.com/hashicorp/consul/api"
	"github.com/spf13/viper"
)

func GetConfig(key string) string {
	viper.AddConfigPath("./library")
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

	authErr := util.SetBatchConfigValues(client, authServiceConfig)
	bookErr := util.SetBatchConfigValues(client, bookServiceConfig)
	rentErr := util.SetBatchConfigValues(client, rentServiceConfig)

	isFailed := authErr != nil || bookErr != nil || rentErr != nil

	if isFailed {
		return errors.New("config initialization failed")
	}

	return nil
}
