package redis

import (
	"github.com/go-redis/redis/v8"
	"os"
	"sync"
)

var client *redis.Client

var once sync.Once

func GetClient() *redis.Client {
	return client
}

func NewRedisClient() *redis.Client {
	opt, _ := redis.ParseURL(os.Getenv("REDIS_CONNECTION_STRING"))
	c := redis.NewClient(opt)
	once.Do(func() {
		client = c
	})
	return c
}
