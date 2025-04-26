package models

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/Ademayowa/rest-api-demo/db"
)

type Job struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Location    string   `json:"location"`
	Salary      float64  `json:"salary"`
	Duties      []string `json:"duties"`
	Url         string   `json:"url"`
	CreatedAt   string   `json:"created_at"`
}

// Save job into the database
func (job *Job) Save() error {
	job.ID = uuid.New().String()

	// Serialize Duties field to JSON
	dutiesJSON, err := json.Marshal(job.Duties)
	if err != nil {
		return err
	}

	query := `
		INSERT INTO jobs(id, title, description, location, salary, duties, url, created_at)
		VALUES(?, ?, ?, ?, ?, ?, ?, ?)
	`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	job.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	_, err = stmt.Exec(
		job.ID,
		job.Title,
		job.Description,
		job.Location,
		job.Salary,
		string(dutiesJSON),
		job.Url,
		job.CreatedAt,
	)
	return err
}

// Get all jobs (with optional filtering)
func GetAllJobs(filterTitle string) ([]Job, error) {
	query := "SELECT * FROM jobs WHERE 1=1"
	args := []interface{}{}

	// Filter jobs by the title
	if strings.TrimSpace(filterTitle) != "" {
		query += " AND LOWER(title) LIKE ?"
		args = append(args, "%"+strings.ToLower(filterTitle)+"%")
	}

	// Fetch jobs
	rows, err := db.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

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
			&job.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		// Convert Duties field to []string
		if err := json.Unmarshal([]byte(dutiesJSON), &job.Duties); err != nil {
			return nil, err
		}

		jobs = append(jobs, job)
	}
	return jobs, nil
}
