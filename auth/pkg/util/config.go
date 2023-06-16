package util

import (
	"pkg-service/constant"
	"pkg-service/util"

	"github.com/hashicorp/consul/api"
)

func GetDBConfigs(client *api.Client) (map[string]string, error) {
	keys := []string{
		constant.AUTH_DB_USERNAME,
		constant.AUTH_DB_PASSWORD,
		constant.AUTH_DB_NAME,
		constant.AUTH_DB_HOST,
		constant.AUTH_DB_PORT,
	}

	configs, err := util.GetBatchConfigValues(client, keys)

	if err != nil {
		return nil, err
	}

	return configs, nil
}
