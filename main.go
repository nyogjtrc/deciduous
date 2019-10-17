package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/nyogjtrc/deciduous/config"
	"github.com/nyogjtrc/deciduous/conn"
	"github.com/nyogjtrc/deciduous/logging"
	"github.com/nyogjtrc/deciduous/routes"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	version string
	date    string
	commit  string
)

func connectDB() {
	dbRead, err := conn.DBReadOpen(viper.GetString(config.KeyMariaRead))
	if err != nil {
		zap.L().Panic(err.Error())
	}
	conn.SetDBRead(dbRead)

	dbWrite, err := conn.DBWriteOpen(viper.GetString(config.KeyMariaWrite))
	if err != nil {
		zap.L().Panic(err.Error())
	}
	conn.SetDBWrite(dbWrite)
}

func connectRedis() {
	rClient, err := conn.DialRedis(
		viper.GetString(config.KeyRedisAddress),
		"",
		viper.GetInt(config.KeyRedisDB),
	)
	if err != nil {
		zap.L().Panic(err.Error())
	}
	conn.SetRedis(rClient)
}

func service() {
	engine := gin.New()
	engine.Use(gin.Recovery())
	engine.Use(gin.Logger())

	// sets routes
	routes.API(engine)
	routes.Websocket(engine)

	port := "8080"
	if value := viper.GetString(config.KeyServicePort); value != "" {
		port = value
	}

	engine.Run(fmt.Sprintf(":%s", port))
}

func main() {
	// setting logger
	logger := logging.NewRollingLogger()
	zap.ReplaceGlobals(logger)

	zap.L().Info("info",
		zap.String("version", version),
		zap.String("commit:", commit),
		zap.String("date:", date),
	)

	err := config.Load()
	if err != nil {
		zap.L().Panic(err.Error())
	}

	connectDB()
	connectRedis()

	service()
}
