package config

import (
	"strings"

	"github.com/spf13/viper"
)

// LoadEnv from os env
func LoadEnv() {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("DECIDUOUS")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
}

// LoadFile from config.yml
func LoadFile() error {
	viper.SetConfigType("yml")
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	return nil
}

// LoadRun will load config from env and file
func LoadRun() error {
	LoadEnv()
	return LoadFile()
}

// LoadTest will load config from env for testing
func LoadTest() {
	LoadEnv()
}
