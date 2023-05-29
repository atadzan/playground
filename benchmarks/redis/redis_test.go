package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"math/rand"
	"testing"
)

func BenchmarkSetRandomRedisParallel(b *testing.B) {
	client2 := redis.NewClient(&redis.Options{Addr: "127.0.0.1:6380", Password: "", DB: 0})
	if _, err := client2.Ping(context.Background()).Result(); err != nil {
		log.Fatal(err)
	}

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			key := fmt.Sprintf("bench-%d", rand.Int31())
			_, err := client2.Set(context.Background(), key, rand.Int31(), 0).Result()
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}
