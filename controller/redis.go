package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nyogjtrc/deciduous/core/dbconn"
	"github.com/nyogjtrc/deciduous/core/logging"
)

// RedisServer response redis info server section
func RedisServer(c *gin.Context) {
	result, err := dbconn.RedisClient().Info("server").Result()
	if err != nil {
		logging.L().Error(err.Error())
	}
	c.JSON(http.StatusOK, result)
}
