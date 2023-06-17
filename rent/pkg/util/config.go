package util

import (
	"pkg-service/constant"
	"pkg-service/util"

	"github.com/hashicorp/consul/api"
)

func GetDBConfigs(client *api.Client) (map[string]string, error) {
	keys := []string{
		constant.RENT_DB_USERNAME,
		constant.RENT_DB_PASSWORD,
		constant.RENT_DB_NAME,
		constant.RENT_DB_HOST,
		constant.RENT_DB_PORT,
	}

	configs, err := util.GetBatchConfigValues(client, keys)

	if err != nil {
		return nil, err
	}

	return configs, nil
}
