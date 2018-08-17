package main

import (
	"fmt"

	"github.com/nyogjtrc/deciduous/cmd"
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
