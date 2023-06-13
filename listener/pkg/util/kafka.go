package util

import (
	"context"
	"encoding/json"
	"listener/internal/handler/logging"
	logmodel "logging-service/pkg/model"
	"pkg-service/discovery"
	"strings"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func CreateConsumer(config *kafka.ConfigMap) (*kafka.Consumer, error) {
	consumer, err := kafka.NewConsumer(config)

	if err != nil {
		return nil, err
	}

	return consumer, nil
}

func HandleMessage(key string, value []byte, registry discovery.Registry) {
	logHandler := logging.New(registry)

	messageKey := strings.Split(key, ":")[0]

	switch messageKey {
	case "log":
		var logInput logmodel.LogInput

		err := json.Unmarshal(value, &logInput)

		if err != nil {
			panic(err)
		}

		logHandler.Write(context.TODO(), logInput)
	}
}
