package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

const (
	redisHost = "redis"
	redisPort = "6379"
)

var (
	Client redisClientInterface = &redisClient{}
)

type redisClientInterface interface {
	setClient(client *redis.Client)
	Set(string, string, time.Duration) error
	Get(string) (string, error)
	Incr(string) (int64, error)
}

type redisClient struct {
	client *redis.Client
}

func (c *redisClient) setClient(client *redis.Client) {
	c.client = client
}

func (c *redisClient) Set(key string, value string, expiration time.Duration) error {
	ctx := context.Background()
	return c.client.Set(ctx, key, value, expiration).Err()
}

func (c *redisClient) Get(key string) (string, error) {
	ctx := context.Background()
	return c.client.Get(ctx, key).Result()
}

func (c *redisClient) Incr(key string) (int64, error) {
	ctx := context.Background()
	return c.client.Incr(ctx, key).Result()
}

func Init() {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redisHost, redisPort),
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}

	Client.setClient(rdb)
}
