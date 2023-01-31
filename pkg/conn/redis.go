package conn

import (
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

// OpenRedis will create redis connection and ping
func OpenRedis(addr, password string, db int) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	err := client.Ping().Err()
	return client, err
}

// OpenRedisFromViper will open redis connection from viper config data
func OpenRedisFromViper() (*redis.Client, error) {
	viper.SetDefault("redis.addr", "")
	viper.SetDefault("redis.password", "")
	viper.SetDefault("redis.db", 0)

	return OpenRedis(
		viper.GetString("redis.addr"),
		viper.GetString("redis.password"),
		viper.GetInt("redis.db"),
	)
}
