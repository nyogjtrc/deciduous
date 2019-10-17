package conn

import (
	"github.com/go-redis/redis"
)

var redisClient *redis.Client

// SetRedis will replace redis client
func SetRedis(client *redis.Client) {
	redisClient = client
}

// Redis getter
func Redis() *redis.Client {
	return redisClient
}

// DialRedis create redis connection
func DialRedis(addr, password string, db int) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	err := client.Ping().Err()
	return client, err
}

// DialTestRedis create testing redis connection
func DialTestRedis() (*redis.Client, error) {
	return DialRedis("", "", 1)
}
