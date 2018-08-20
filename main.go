package main

import (
	_ "github.com/go-sql-driver/mysql"

	"github.com/nyogjtrc/deciduous/cmd"
	"github.com/nyogjtrc/deciduous/core/dbconn"
	"github.com/nyogjtrc/deciduous/core/logging"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	version string
	date    string
	commit  string
)

func readconfig() {
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		logging.L().Panic(err.Error())
	}
}

func connectDB() {
	key := "maria.read"
	if !viper.IsSet(key) {
		logging.L().Fatal("config key not found", zap.String("key", key))
	}
	dbconn.OpenRead(viper.GetString(key))

	key = "maria.write"
	if !viper.IsSet(key) {
		logging.L().Fatal("config key not found", zap.String("key", key))
	}
	dbconn.OpenWrite(viper.GetString(key))
}

func connectRedis() {
	key := "redis.addr"
	if !viper.IsSet(key) {
		logging.L().Fatal("config key not found", zap.String("key", key))
	}

	dbconn.RedisDial(viper.GetString(key))
}

func main() {
	logging.L().Info("info",
		zap.String("version", version),
		zap.String("commit:", commit),
		zap.String("date:", date),
	)

	readconfig()

	connectDB()
	connectRedis()

	cmd.Execute()
}
