package dbconn

import (
	"time"

	"github.com/go-redis/redis"
)

var (
	redisClient *redis.Client
)

// RedisDial create redis client pool and ping redis server
func RedisDial(addr string) {
	redisClient = redis.NewClient(&redis.Options{
		Addr:         addr,
		Password:     "",
		DB:           0,
		DialTimeout:  10 * time.Second,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		PoolSize:     100,
		PoolTimeout:  15 * time.Second,
	})

	_, err := redisClient.Ping().Result()
	if err != nil {
		panic(err.Error())
	}
}

// RedisClient return redis client
func RedisClient() *redis.Client {
	return redisClient
}
