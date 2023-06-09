package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
)

type processCache struct {
	Type   string
	Status string
}

type detailedInfoCache struct {
	Total       int
	Completed   int
	Description string
}

type channelCache struct {
	ChannelId string `redis:"channelId"`
	IsForeign bool   `redis:"isForeign"`
	Username  string `redis:"username"`
}

func main() {
	cache := redis.NewClient(&redis.Options{
		Addr:     "localhost:6385",
		Password: "",
		DB:       1,
	})
	ctx := context.Background()
	err := cache.Ping(ctx).Err()
	if err != nil {
		log.Fatal("can't ping redis. Error: ", err)
	}
	fmt.Println("Connected")
	channelId := "UCKT1kXUTbb2FivzJ8bvQ2YA"

	// Set values
	if baseErr := cache.HSet(ctx, channelId+":baseInfo", "channelId", channelId, "isForeign", true, "username", "user-1").Err(); baseErr != nil {
		log.Println("can't set", err.Error())
		return
	}
	if baseErr := cache.HSet(ctx, channelId+":processInfo", "type", "parse", "status", "active").Err(); baseErr != nil {
		log.Println("can't set", err.Error())
		return
	}
	if baseErr := cache.HSet(ctx, channelId+":detailedInfo", "total", "8", "completed", 4, "description", "видео из").Err(); baseErr != nil {
		log.Println("can't set", err.Error())
		return
	}
	fmt.Println("Success")

	// get values and scan into struct
	var channel channelCache
	if err = cache.HMGet(ctx, channelId+":baseInfo", "channelId", "isForeign", "username").Scan(&channel); err != nil {
		log.Println("can't get ", err.Error())
		return
	}
	fmt.Println(channel)

	// update specific values
	if err = cache.HSet(ctx, channelId+":"+"detailedInfo", "completed", 5).Err(); err != nil {
		log.Println("can't get ", err.Error())
		return
	}
	fmt.Println("Success")
}
