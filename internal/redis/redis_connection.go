package redis

import (
	"log"

	"github.com/go-redis/redis/v8"
)

var Redis *redis.Client

func RedisConnection() {
	redis := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
	})

	Redis = redis
	log.Print("Connect redis success")
}
