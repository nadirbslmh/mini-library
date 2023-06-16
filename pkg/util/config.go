package util

import (
	"fmt"

	"github.com/hashicorp/consul/api"
)

func SetConfigValue(client *api.Client, key, value string) error {
	pair := &api.KVPair{
		Key:   key,
		Value: []byte(value),
	}

	_, err := client.KV().Put(pair, nil)
	return err
}

func GetConfigValue(client *api.Client, key string) (string, error) {
	kvPair, _, err := client.KV().Get(key, nil)
	if err != nil {
		return "", err
	}

	if kvPair == nil {
		return "", fmt.Errorf("key '%s' not found in Consul", key)
	}

	return string(kvPair.Value), nil
}

func SetBatchConfigValues(client *api.Client, pairs map[string]string) error {
	for key, val := range pairs {
		if err := SetConfigValue(client, key, val); err != nil {
			return err
		}
	}

	return nil
}

func GetBatchConfigValues(client *api.Client, keys []string) (map[string]string, error) {
	values := map[string]string{}

	for _, key := range keys {
		val, err := GetConfigValue(client, key)
		if err != nil {
			return nil, err
		}
		values[key] = val
	}

	return values, nil
}
