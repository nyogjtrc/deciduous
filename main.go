package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/nyogjtrc/deciduous/cmd"
	"github.com/nyogjtrc/deciduous/core/dbconn"
	"github.com/nyogjtrc/deciduous/logging"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	version string
	date    string
	commit  string
)

func main() {
	logging.L().Info(
		"info",
		zap.String("version", version),
		zap.String("commit:", commit),
		zap.String("date:", date),
	)

	readconfig()
	connectDB()

	cmd.Execute()
}

func readconfig() {
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}
}

func connectDB() {
	DSN := "developer:developer@/"

	key := "maria.read"
	if !viper.IsSet(key) {
		logging.L().Fatal("config key not found", zap.String("key", key))
	}

	DSN = viper.GetString(key)
	dbconn.OpenRead(DSN)

	key = "maria.write"
	if !viper.IsSet(key) {
		logging.L().Fatal("config key not found", zap.String("key", key))
	}

	DSN = viper.GetString(key)
	dbconn.OpenWrite(DSN)
}
