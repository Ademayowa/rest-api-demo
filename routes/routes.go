package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	// Define routes
	router.POST("/jobs", createJob)
}
