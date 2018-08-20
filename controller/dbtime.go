package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nyogjtrc/deciduous/core/dbconn"
	"github.com/nyogjtrc/deciduous/core/logging"
)

// DBnow response db time now
func DBnow(c *gin.Context) {
	var result string
	err := dbconn.DBRead().QueryRow("SELECT NOW()").Scan(&result)
	if err != nil {
		logging.L().Error(err.Error())
	}
	c.JSON(http.StatusOK, gin.H{
		"now": result,
	})
}
