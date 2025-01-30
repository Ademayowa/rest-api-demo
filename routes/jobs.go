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
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse job" + err.Error()})
		return
	}

	job.Save()
	context.JSON(http.StatusCreated, gin.H{"message": "job created", "job": job})
}

// Fetch all jobs
func getJobs(context *gin.Context) {
	// Extract query parameter
	filterTitle := context.Query("title")

	jobs, total, err := models.GetAllJobs(filterTitle)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"jobs": jobs, "total": total})
}
