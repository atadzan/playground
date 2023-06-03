package main

import (
	"github.com/atadzan/playground/centrifugo"
	"github.com/atadzan/playground/centrifugo/events"
	"github.com/centrifugal/centrifuge-go"
	"log"
)

func main() {
	client := centrifugo.NewCentrifugoConnection("user-1")
	defer client.Close()

	events.ClientEvents(client)

	if err := client.Connect(); err != nil {
		log.Println("can't connect to centrifugo client.Error:", err.Error())
		return
	}
	sub, err := client.NewSubscription(centrifugo.Channel, centrifuge.SubscriptionConfig{
		Recoverable: true,
		JoinLeave:   true,
	})
	if err != nil {
		log.Println("cant create new subscription", err.Error())
		return
	}

	events.SubscriptionEvents(sub)

	if err = sub.Subscribe(); err != nil {
		log.Println("subscribe error", err.Error())
		return
	}
	defer sub.Unsubscribe()

	select {}
}
