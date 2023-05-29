package main

import (
	"encoding/json"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type Message struct {
	Title       string
	Description string
}

func main() {
	//admin, err := kafka.NewAdminClient(&kafka.ConfigMap{"bootstrap.servers": "localhost"})
	//if err != nil {
	//	log.Println("cant create admin. Error: ", err.Error())
	//	return
	//}

	//_, err = admin.CreateTopics(context.Background(),
	//	[]kafka.TopicSpecification{kafka.TopicSpecification{
	//		Topic:         "myTopic",
	//		NumPartitions: 1}})
	//if err != nil {
	//	log.Println("can't create topic. Error: ", err.Error())
	//	return
	//}

	//defer admin.Close()
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost"})
	if err != nil {
		panic(err)
	}

	defer p.Close()

	// Deliver report handler gor produced messages
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	// Produce message to topic (async)
	topic := "myTopic"
	//for _, word := range []string{"hello", "world"} {
	msg := &Message{Title: "Hello-title-1", Description: "World-description-1"}
	raw, _ := json.Marshal(msg)
	p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          raw,
	}, nil)
	//}

	// wait for message deliveries before shutdown
	p.Flush(2 * 1000)
}
