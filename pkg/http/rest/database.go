package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// DatabaseHandler API handler
type DatabaseHandler struct {
	db *gorm.DB
}

// NewDatabaseHandler return instance
func NewDatabaseHandler(db *gorm.DB) *DatabaseHandler {
	return &DatabaseHandler{db}
}

// Router for gin
func (h *DatabaseHandler) Router(r *gin.Engine) {
	r.GET("api/database/now", func(c *gin.Context) {
		var result string

		mysqlDB, err := h.db.DB()
		if err != nil {
			zap.L().Error(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}

		err = mysqlDB.QueryRow("SELECT NOW()").Scan(&result)
		if err != nil {
			zap.L().Error(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}

		c.JSON(http.StatusOK, gin.H{"now": result})
	})
}
