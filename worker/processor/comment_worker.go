package processor

import (
	"fmt"
	"worker/kafka"

	"log"
	"os"
	"os/signal"
	"syscall"
)

func StartCommentWorker(consumer *kafka.Consumer, topic string) int {
	partitionConsumer, err := consumer.ConsumePartition(topic, 0)
	if err != nil {
		log.Fatalf("Failed to start consumer: %s", err)
	}
	defer partitionConsumer.Close()

	msgCount := 0
	doneCh := make(chan struct{})
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	go func ()  {
		for {
			select {
			case err := <-partitionConsumer.Errors():
				fmt.Println("Error:", err)

			case msg := <-partitionConsumer.Messages():
				msgCount++
				fmt.Printf("Received message Count: %d | Topic (%s) | Message (%s)\n", msgCount, msg.Topic, string(msg.Value))
			

			case <-sigchan:
				fmt.Println("Interruption detected")
				doneCh <- struct{}{}
				return
			}
		}
	}()

	<-doneCh
	return msgCount
}