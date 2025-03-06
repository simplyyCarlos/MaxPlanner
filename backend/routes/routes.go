package routes

import (
	"backend/controllers"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes sets up API routes
func RegisterRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.GET("/sncf", controllers.GetSNCFData)
	}
}
