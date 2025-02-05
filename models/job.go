package models

import (
	"encoding/json"
	"strings"

	"github.com/Ademayowa/rest-api-demo/db"
)

type Job struct {
	ID          int64    `json:"id"`
	Title       string   `json:"title" binding:"required"`
	Description string   `json:"description" binding:"required"`
	Location    string   `json:"location" binding:"required"`
	Salary      string   `json:"salary" binding:"required"`
	Duties      []string `json:"duties" binding:"required"`
	Url         string   `json:"url"`
}

// Save into database
func (job Job) Save() error {
	query := `
    INSERT INTO jobs(title, description, location, salary, duties, url)
    VALUES (?, ?, ?, ?, ?, ?)
  `

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}
	defer stmt.Close()

	// Serialize Duties field to JSON
	dutiesJSON, err := json.Marshal(job.Duties)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(
		job.Title,
		job.Description,
		job.Location,
		job.Salary,
		string(dutiesJSON),
		job.Url,
	)
	if err != nil {
		return err
	}

	// Add the auto generated ID from the database
	id, err := result.LastInsertId()
	job.ID = id

	return err
}

// Get all jobs (with optional filtering)
func GetAllJobs(filterTitle string) ([]Job, int, error) {
	query := "SELECT * FROM jobs WHERE 1=1"
	args := []interface{}{}

	if strings.TrimSpace(filterTitle) != "" {
		query += " title LIKE ?"
		args = append(args, "%"+strings.ToLower(filterTitle)+"%")
	}

	// Count total jobs that matches the filter from the database
	countQuery := "SELECT COUNT(*) FROM (" + query + ") AS count_query"

	var total int
	err := db.DB.QueryRow(countQuery, args...).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	// Fetch all jobs with filter from the database
	rows, err := db.DB.Query(query, args...)
	if err != nil {
		return nil, 0, err
	}

	var jobs []Job

	for rows.Next() {
		var job Job
		var dutiesJSON string

		err := rows.Scan(
			&job.ID,
			&job.Title,
			&job.Description,
			&job.Location,
			&job.Salary,
			&dutiesJSON,
			&job.Url,
		)
		if err != nil {
			return nil, 0, err
		}

		// Convert duties field back to []string
		err = json.Unmarshal([]byte(dutiesJSON), &job.Duties)
		if err != nil {
			return nil, 0, err
		}

		jobs = append(jobs, job)
	}

	return jobs, total, err
}

// Get a single job
func GetJobByID(id int64) (*Job, error) {
	query := "SELECT * FROM jobs WHERE id = ?"
	row := db.DB.QueryRow(query, id)

	var job Job
	var dutiesJSON string

	err := row.Scan(
		&job.ID,
		&job.Title,
		&job.Description,
		&job.Location,
		&job.Salary,
		&dutiesJSON,
		&job.Url,
	)
	if err != nil {
		return nil, err
	}

	// Convert Duties field from JSON to []string
	err = json.Unmarshal([]byte(dutiesJSON), &job.Duties)
	if err != nil {
		return nil, err
	}

	return &job, nil
}

// Delete job
func (job Job) Delete() error {
	query := "DELETE FROM jobs WHERE id =?"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(job.ID)

	return err
}

// Update a job
func UpdateJobByID(id int64, updatedJob Job, dutiesJSON string) error {
	query := `
		UPDATE jobs
		SET title = ?, description = ?, location = ?, salary = ?, duties = ?, url = ?
		WHERE id = ?
	`
	_, err := db.DB.Exec(query,
		updatedJob.Title,
		updatedJob.Description,
		updatedJob.Location,
		updatedJob.Salary,
		dutiesJSON,
		updatedJob.Url,
		id,
	)

	return err
}
