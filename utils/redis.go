package utils

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

var redisClient *redis.Client

func GetRedisClient() *redis.Client {
	if redisClient == nil {
		redisClient = redis.NewClient(&redis.Options{
			Addr: "localhost:6379",
		})
	}

	return redisClient
}

func Set(key string, value interface{}, duration int) error {
	rdb := GetRedisClient()
	err := rdb.Set(ctx, key, value, time.Duration(duration)*time.Second).Err()
	if err != nil {
		return err
	}

	return nil
}

func Client() {
	// Connect to Redis
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	redisClient = rdb

	// Set a value
	err := rdb.Set(ctx, "key", "value", 30*time.Second).Err()
	if err != nil {
		panic(err)
	}

	// Get a value
	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val) // Output: key value
}
