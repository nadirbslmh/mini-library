package util

import (
	"pkg-service/constant"
	"pkg-service/util"

	"github.com/hashicorp/consul/api"
)

func GetDBConfigs(client *api.Client) (map[string]string, error) {
	keys := []string{
		constant.BOOK_DB_USERNAME,
		constant.BOOK_DB_PASSWORD,
		constant.BOOK_DB_NAME,
		constant.BOOK_DB_HOST,
		constant.BOOK_DB_PORT,
	}

	configs, err := util.GetBatchConfigValues(client, keys)

	if err != nil {
		return nil, err
	}

	return configs, nil
}
