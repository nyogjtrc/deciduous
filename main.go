package main

import (
	"fmt"

	"github.com/nyogjtrc/deciduous/cmd"
	"github.com/spf13/viper"
)

var (
	version string
	date    string
	commit  string
)

func main() {
	fmt.Println("version:", version)
	fmt.Println("commit:", commit)
	fmt.Println("date:", date)

	readconfig()

	fmt.Println(viper.GetBool("debug"))

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
