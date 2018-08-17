package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nyogjtrc/deciduous/controller"
)

// API routes
func API(router *gin.Engine) {
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/db/now", controller.DBnow)
}
