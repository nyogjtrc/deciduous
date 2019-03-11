package config

import (
	"errors"

	"github.com/spf13/viper"
)

// config keys
const (
	KeyDebug      string = "debug"
	KeyMariaRead  string = "maria.read"
	KeyMariaWrite string = "maria.write"

	KeyRedisAddress string = "redis.addr"
	KeyRedisDB      string = "redis.db"

	KeyServicePort string = "service.port"
)

var mustHaveKeys = []string{
	KeyMariaRead,
	KeyMariaWrite,
	KeyRedisAddress,
	KeyRedisDB,
}

// Load config.yaml
func Load() error {
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	return checkValue()
}

func checkValue() error {
	for e := range mustHaveKeys {
		if !viper.IsSet(mustHaveKeys[e]) {
			return errors.New("config error: missing key [" + mustHaveKeys[e] + "]")
		}
	}

	return nil
}
