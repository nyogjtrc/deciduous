package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nyogjtrc/deciduous/conn"
	"github.com/nyogjtrc/deciduous/logging"
)

// DBnow response db time now
func DBnow(c *gin.Context) {
	var result string
	err := conn.DBRead().QueryRow("SELECT NOW()").Scan(&result)
	if err != nil {
		logging.L().Error(err.Error())
	}
	c.JSON(http.StatusOK, gin.H{
		"now": result,
	})
}
