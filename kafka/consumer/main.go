package main

import (
	"encoding/json"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"log"
)

type Message struct {
	Title       string
	Description string
}

func main() {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost",
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		log.Println("cant create consumer", err.Error())
		panic(err)
	}
	c.Subscribe("myTopic", nil)
	// A signal handler or similar could be used to set this to false to break the loop.
	run := true
	var income Message
	for run {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			fmt.Printf("Message on %s \n", string(msg.Value))
			if err := json.Unmarshal(msg.Value, &income); err != nil {
				log.Println(err.Error())
			}
		} else if !err.(kafka.Error).IsTimeout() {
			// The client will automatically try to recover from all errors.
			// Timeout is not considered an error because it is raised by
			// ReadMessage in absence of messages
			fmt.Printf("Consumer error: %v (%v\n)", err, msg)
		}
	}
	c.Close()
}
