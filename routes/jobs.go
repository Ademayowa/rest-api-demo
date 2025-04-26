package routes

import (
	"net/http"

	"github.com/Ademayowa/rest-api-demo/models"
	"github.com/gin-gonic/gin"
)

// Create a job
func createJob(context *gin.Context) {
	var job models.Job

	err := context.ShouldBindJSON(&job)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "could not parse job data"})
		return
	}

	job.Save()
	context.JSON(http.StatusCreated, gin.H{"message": "job created", "job": job})
}
