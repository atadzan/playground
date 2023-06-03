package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/atadzan/playground/centrifugo"
	"github.com/atadzan/playground/centrifugo/events"
	"github.com/centrifugal/centrifuge-go"
	"log"
)

func main() {
	client := centrifugo.NewCentrifugoConnection("user-3")

	events.ClientEvents(client)
	if err := client.Connect(); err != nil {
		log.Println("cant connect to centrifugo client.Error:", err.Error())
		return
	}
	sub, err := client.NewSubscription(centrifugo.Channel, centrifuge.SubscriptionConfig{
		Recoverable: true,
		JoinLeave:   true,
	})
	if err != nil {
		log.Println("sub error.", err.Error())
	}
	events.SubscriptionEvents(sub)

	if err := sub.Subscribe(); err != nil {
		log.Println("subscribe error", err.Error())
		return
	}
	defer sub.Unsubscribe()
	ctx := context.Background()
	var msg string
	go func() {
		for {
			fmt.Scan(&msg)
			rawMsg, _ := json.Marshal(msg)
			_, err := client.Publish(ctx, centrifugo.Channel, rawMsg)
			if err != nil {
				log.Println("publish error", err.Error())
			}
			if msg == "exit" {
				break
			}
		}

	}()
	select {}
}
