package delivery

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

func SetEndpoints(group *gin.RouterGroup, db *sql.DB) {
	// handle := NewHnadler(db)
	{
		group.GET("/cities")
		group.POST("/cities")
		group.GET("/cities/:id")
		group.PUT("/cities/:id")
		group.DELETE("/cities/:id")
	}
}
