package redisdb

import (
	"github.com/go-redis/redis"
	"os"
)

var Client *redis.Client

func init() {
	Client = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST"),
		Password: os.Getenv("REDIS_AUTH"),
		DB:       0,
	})
}
