package config

import (
	"log"
	"producer/kafka"

	"github.com/IBM/sarama"
)

func NewKafkaProducer() *kafka.Producer {
	brokerUrls := []string{"localhost:29092"}
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5

	saramaProducer, err := sarama.NewSyncProducer(brokerUrls, config)
	if err != nil {
		log.Fatalf("Failed to start Sarama producer: %s", err.Error())
	}

	return &kafka.Producer{
		Producer: saramaProducer,
	}
}