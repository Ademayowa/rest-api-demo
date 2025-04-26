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

// Fetch all jobs
func getJobs(context *gin.Context) {
	// Extract job query parameter from the URL
	filterTitle := context.Query("query")

	// Get all jobs with filter
	jobs, err := models.GetAllJobs(filterTitle)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch jobs"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"jobs": jobs})
}
