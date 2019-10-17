package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nyogjtrc/deciduous/conn"
	"go.uber.org/zap"
)

// RedisServer response redis info server section
func RedisServer(c *gin.Context) {
	result, err := conn.Redis().Info("server").Result()
	if err != nil {
		zap.L().Error(err.Error())
	}
	c.JSON(http.StatusOK, result)
}
