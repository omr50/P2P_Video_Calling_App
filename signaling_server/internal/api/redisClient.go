package api

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

// redis.NewClient(&redis.Opitons{
// 	Addr: "localhost:6379",
// 	DB: 0,
// })

var RedisClient *redis.Client

func InitRedisClient() {

	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})

	if err := rdb.Ping(ctx).Err(); err != nil {
		panic(err)
	}

	fmt.Println("Redis client created!")
	RedisClient = rdb
}
