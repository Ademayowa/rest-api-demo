package models

import (
	"encoding/json"

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
