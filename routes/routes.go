package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	// Define routes
	router.POST("/jobs", createJob)
	router.GET("/jobs", getJobs)
	router.GET("/jobs/:id", getJob)
	router.DELETE("/jobs/:id", deleteJob)
}
