package models

import (
	"encoding/json"
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
