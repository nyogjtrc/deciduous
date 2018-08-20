package cmd

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/nyogjtrc/deciduous/core/logging"
	"github.com/nyogjtrc/deciduous/routes"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(serviceCmd)
}

var serviceCmd = &cobra.Command{
	Use: "service",
	Run: func(cmd *cobra.Command, args []string) {
		service()
	},
}

func service() {
	logging.L().Info("run service")

	gin.SetMode(gin.DebugMode)

	engine := gin.New()
	engine.Use(gin.Recovery())
	engine.Use(gin.Logger())

	// sets routes
	routes.API(engine)
	routes.Websocket(engine)

	port := "8080"
	if value := viper.GetString("service"); value != "" {
		port = value
	}

	engine.Run(fmt.Sprintf(":%s", port))
}
