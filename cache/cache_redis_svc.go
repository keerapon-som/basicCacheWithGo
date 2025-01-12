package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
)

type cache struct {
	redisClient *redis.Client
	ctx         context.Context
}

func NewRedisCache(redisClient *redis.Client) Cache {
	return &cache{
		redisClient: redisClient,
		ctx:         context.Background(),
	}
}

func (c *cache) Set(key string, value interface{}, duration int) error {
	jsonData, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return c.redisClient.Set(c.ctx, key, jsonData, time.Duration(duration)*time.Second).Err()
}

func (c *cache) Get(key string, receiver interface{}) error {
	val, err := c.redisClient.Get(c.ctx, key).Result()
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(val), receiver)
}
