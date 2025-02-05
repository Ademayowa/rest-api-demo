package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Ademayowa/rest-api-demo/models"
	"github.com/gin-gonic/gin"
)

// Create a job
func createJob(context *gin.Context) {
	var job models.Job

	err := context.ShouldBindJSON(&job)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "could not create job",
			"error":   err.Error(),
		})
		return
	}

	job.Save()
	context.JSON(http.StatusCreated, gin.H{"message": "job created", "data": job})
}

// Fetch all jobs
func getJobs(context *gin.Context) {
	// Extract the query parameter
	filterTitle := context.Query("title")

	jobs, total, err := models.GetAllJobs(filterTitle)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "could not fetch all jobs",
			"error":   err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "jobs fetch successfully",
		"data":    jobs,
		"total":   total,
	})
}

// Fetch a single job
func getJob(context *gin.Context) {
	// Extract the id from the URL & convert to int64
	jobId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "could not parse job id",
			"error":   err.Error(),
		})
		return
	}

	job, err := models.GetJobByID(jobId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "could not fetch job",
			"error":   err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "job fetch successfully", "data": job})
}

// Delete a job
func deleteJob(context *gin.Context) {
	jobId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "could not parse job id",
			"error":   err.Error(),
		})
		return
	}

	job, err := models.GetJobByID(jobId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "could not fetch job",
			"error":   err.Error(),
		})
		return
	}

	err = job.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "could not delete job",
			"error":   err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "job deleted successfully"})
}

func updateJob(context *gin.Context) {
	// Parse job ID into the URL & convert to int64
	jobId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid job id",
			"error":   err.Error(),
		})
		return
	}

	// Parse the request body to get the updated job data
	var updatedJob models.Job

	err = context.ShouldBindJSON(&updatedJob)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request body",
			"error":   err.Error(),
		})
		return
	}

	// Serialize the Duties field to JSON for database storage
	dutiesJSON, err := json.Marshal(updatedJob.Duties)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "error processing duties field",
			"error":   err.Error(),
		})
		return
	}

	// Update job in the database
	err = models.UpdateJobByID(jobId, updatedJob, string(dutiesJSON))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "could not update job",
			"error":   err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "job updated successfully"})
}
