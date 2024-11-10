package kafka

import (
	"log"

	"github.com/IBM/sarama"
)

type Producer struct {
	Producer sarama.SyncProducer
}

func (p *Producer) Close() error {
	return p.Producer.Close()
}

func (p *Producer) PushComment(topic string, message []byte) {
	go func() {
		msg := &sarama.ProducerMessage{
			Topic:     topic,
			Value:     sarama.ByteEncoder(message),
		}

		partition, offset, err := p.Producer.SendMessage(msg)
		if err != nil {
			log.Printf("Failed to send message to Kafka: %s", err.Error())
			return
		}
		log.Printf("Message stored in topic %s partition %d at offset %d\n", topic, partition, offset+1)
	}()
}