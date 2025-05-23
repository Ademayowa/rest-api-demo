package routes

import (
	"encoding/json"
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

// Fetch a single job
func getJob(context *gin.Context) {
	jobId := context.Param("id")

	job, err := models.GetJobByID(jobId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch job"})
		return
	}

	context.JSON(http.StatusOK, job)
}

// Delete a job
func deleteJob(context *gin.Context) {
	jobId := context.Param("id")

	job, err := models.GetJobByID(jobId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch job"})
		return
	}

	err = job.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "could not delete job"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "job deleted"})
}

// Update a job
func updateJob(context *gin.Context) {
	// Extract job ID from the URL
	jobId := context.Param("id")

	// Parse the request body to get the updated job data
	var updatedJob models.Job
	if err := context.ShouldBindJSON(&updatedJob); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}
	// Convert Duties field to JSON
	dutiesJSON, err := json.Marshal(updatedJob.Duties)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "failure at duties field"})
		return
	}
	// Update job in the database
	err = models.UpdateJobByID(jobId, updatedJob, string(dutiesJSON))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "could not update job"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "job updated"})
}
