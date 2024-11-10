package main

import (
	"fmt"
	"worker/config"
	"worker/processor"

	"os"
	"os/signal"
	"syscall"
)

func main() {
	consumer, err := config.NewKafkaConsumer([]string{"localhost:29092"})
	if err != nil {
		panic(err)
	}
	defer consumer.Close()

	fmt.Println("Consumer started")
	msgCount := processor.StartCommentWorker(consumer, "comments")

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)
	<-sigchan

	fmt.Printf("Processed %d messages\n", msgCount)
}
