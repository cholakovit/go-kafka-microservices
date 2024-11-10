package config

import (
	"worker/kafka"

	"github.com/IBM/sarama"
)

func NewKafkaConsumer(brokerUrl []string) (*kafka.Consumer, error) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	conn, err := sarama.NewConsumer(brokerUrl, config)
	if err != nil {
		return nil, err
	}

	return kafka.NewConsumer(conn), nil
}
