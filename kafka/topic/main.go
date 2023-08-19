package main

import (
	"context"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"log"
	"time"
)

func main() {
	a, err := kafka.NewAdminClient(&kafka.ConfigMap{"bootstrap.servers": "localhost:9092"})
	if err != nil {
		log.Println("error", err)
	}

	maxDuration := 60 * time.Second

	results, err := a.CreateTopics(context.Background(), []kafka.TopicSpecification{
		{
			Topic:         "test-topic",
			NumPartitions: 3,
		},
		{
			Topic:         "test-topic-2",
			NumPartitions: 3,
		},
		{
			Topic:         "test-topic-3",
			NumPartitions: 3,
		}},
		// admin operation option
		kafka.SetAdminOperationTimeout(maxDuration))

	if err != nil {
		log.Println("error-2", err)
	}

	for _, result := range results {
		log.Println("topic", result.Topic, "\n error:", result.Error)
	}
	a.Close()
}
