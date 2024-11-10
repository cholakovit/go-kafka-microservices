package kafka

import "github.com/IBM/sarama"

type Consumer struct {
	consumer sarama.Consumer
}

func NewConsumer(consumer sarama.Consumer) *Consumer {
	return &Consumer{consumer: consumer}
}

func (c *Consumer) Close() error {
	return c.consumer.Close()
}

func (c *Consumer) ConsumePartition(topic string, partition int32) (sarama.PartitionConsumer, error) {
	return c.consumer.ConsumePartition(topic, partition, sarama.OffsetNewest)
}