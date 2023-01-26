package cmd

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/nyogjtrc/deciduous/internal/ver"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

func init() {
	rootCmd.AddCommand(apiCmd)
}

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "REST API service",
	Long:  `REST API service`,
	Run: func(cmd *cobra.Command, args []string) {
		zap.ReplaceGlobals(zap.Must(zap.NewProduction()))

		r := gin.Default()
		r.Use(gzip.Gzip(gzip.BestSpeed))
		r.GET("/api/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "pong"})
		})

		ver.Router(r)

		go func() {
			if err := r.Run(); err != nil {
				zap.S().Fatal(err)
			}
		}()

		// listen os.Signal
		termChan := make(chan os.Signal, 1)
		signal.Notify(termChan, syscall.SIGTERM, syscall.SIGINT)

		// close goroutine when catch interrupt signal
		<-termChan
		zap.L().Info("SIGTERM received. close goroutine\n")
	},
}