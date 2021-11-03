package delivery

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

func SetEndpoints(group *gin.RouterGroup, db *sql.DB) {
	handle := NewHandlers(db)
	{
		group.GET("/cities", handle.GetCities)
		group.POST("/cities", handle.CreateCity)
		group.GET("/cities/:id", handle.GetCity)
		group.PUT("/cities/:id", handle.UpdateCity)
		group.DELETE("/cities/:id", handle.DeleteCity)
	}
}
