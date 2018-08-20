package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/nyogjtrc/deciduous/core/logging"
	"github.com/nyogjtrc/deciduous/routes"
	"github.com/spf13/cobra"
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

	engine := gin.Default()

	// sets routes
	routes.API(engine)
	routes.Websocket(engine)

	engine.Run() // listen and serve on 0.0.0.0:8080
}
