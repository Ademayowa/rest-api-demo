package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	// Define routes
	server.POST("/jobs", createJob)
	server.GET("/jobs", getJobs)
}
