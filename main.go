package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/jobs", getJobs)
	router.Run(":8080")
}

func getJobs(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "Hello Job board!"})
}
