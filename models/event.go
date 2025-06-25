package models

import (
	"telu-event-apps/db"
	"time"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int
}

var events = []Event{}

func (e Event) Save() error {
	// Query to insert a new event into the database. Use the ? placeholder to prevent SQL Injection. The value will be safely inserted later via stmt.Exec().
	query := `
	INSERT INTO events (name, description, location, dateTime, user_id) 
	VALUES (?, ?, ?, ?, ?)`

	// Prepare the statement to prevent SQL injection and prepare SQL statements so that they can be executed many times efficiently
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	// Execute the prepared statement, with the values ​​from the Event struct. The values ​​(e.Name, etc.) will be safely bound to ? in the query.
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err
	}

	// Get the last inserted ID from the result of the Exec() method. This is useful if you want to retrieve the ID of the newly created event.
	id, err := result.LastInsertId()
	e.ID = id

	// Closes the statement after the function completes (to free resources).
	defer stmt.Close()

	return err
}

func GetAllEvents() []Event {
	return events
}
