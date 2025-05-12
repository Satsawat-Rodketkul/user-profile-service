package redis

import (
	"fmt"
	"log"
	"user-profile-service/internal/config"

	"github.com/go-redis/redis/v8"
)

var Redis *redis.Client

func RedisConnection() {
	redisHost := config.GetValue("REDIS_HOST")
	redisPort := config.GetValue("REDIS_PORT")

	dsn := fmt.Sprintf("%v:%v", redisHost, redisPort)

	redis := redis.NewClient(&redis.Options{
		Addr: dsn,
	})

	Redis = redis
	log.Print("Connect redis success")
}
