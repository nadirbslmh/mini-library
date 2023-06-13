package util

import "github.com/confluentinc/confluent-kafka-go/v2/kafka"

func CreateProducer(config *kafka.ConfigMap) (*kafka.Producer, error) {
	producer, err := kafka.NewProducer(config)

	if err != nil {
		return nil, err
	}

	return producer, nil
}

func SendMessage(producer *kafka.Producer, message *kafka.Message) error {
	err := producer.Produce(message, nil)

	if err != nil {
		return err
	}

	producer.Flush(15 * 1000)

	return nil
}
