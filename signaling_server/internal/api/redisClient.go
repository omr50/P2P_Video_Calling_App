package api

import (
	"context"
	"fmt"
	"time"

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

func MarkUserOnline(ctx context.Context, rdb *redis.Client, email string) error {

	key := "online:user:" + email
	err := rdb.HSet(ctx, key, map[string]interface{}{
		"connected_at": time.Now().Unix(),
		"last_seen":    time.Now().Unix(),
	}).Err()
	if err != nil {
		return err
	}

	// set expiration to keep cehcking, WS pings will keep it going
	err = rdb.Expire(ctx, key, 120*time.Second).Err()
	if err != nil {
		return err
	}
	fmt.Println("Marked user online in redis")
	return nil
}

func MarkUserOffline(ctx context.Context, rdb *redis.Client, email string) error {
	key := "online:user:" + email
	err := rdb.Del(ctx, key).Err()
	if err != nil {
		return err
	}
	fmt.Println("marked user offline", err)
	return nil

}

func SetClientExpiration(ctx context.Context, rdb *redis.Client, email string, t time.Duration) error {
	return rdb.Expire(ctx, email, t).Err()
}
