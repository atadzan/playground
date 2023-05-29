package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
)

func main() {
	cache := redis.NewClient(&redis.Options{
		Addr:     "localhost:6380",
		Password: "",
		DB:       1,
	})
	ctx := context.Background()
	err := cache.Ping(ctx)
	if err.Err() != nil {
		log.Fatal("can't ping redis. Error: ", err.Err())
	}
	fmt.Println("Connceted")
}
