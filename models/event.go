package models

import (
	"telu-event-apps/db"
	"time"
)

// Structs in Go function as temporary containers for data being processed – they can be used to store data from requests, fill in data from databases, and be used for operations (queries).
type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int64
}

func (e *Event) Save() error {
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

func GetAllEvents() ([]Event, error) {
	query := `SELECT * FROM events`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}

	// Close the query result connection once the function completes.
	defer rows.Close()

	// Empty slice to store all the events that we will take from the database
	var events = []Event{}
	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}

func GetEventByID(id int64) (*Event, error) {
	query := `SELECT * FROM events WHERE id = ?`

	// Query to select an event by ID from the database. Use the ? placeholder to prevent SQL Injection. The value will be safely inserted later via row.Scan().
	row := db.DB.QueryRow(query, id)

	// Empty slice to store the events that we will take from the database
	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
	if err != nil {
		return nil, err
	}

	// return pointer because if error occurs, we want to return nil
	return &event, nil

}

func (e Event) Update() error {
	// Query to update an existing event in the database. Use the ? placeholder to prevent SQL Injection. The value will be safely inserted later via stmt.Exec().
	query := `
	UPDATE events 
	SET name = ?, description = ?, location = ?, dateTime = ?
	WHERE id = ?`

	// Prepare the statement to prevent SQL injection and prepare SQL statements so that they can be executed many times efficiently
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	// Closes the statement after the function completes (to free resources).
	defer stmt.Close()

	// Execute the prepared statement, with the values ​​from the Event struct. The values ​​(e.Name, etc.) will be safely bound to ? in the query.
	_, err = stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.ID)
	return err
}

func (e Event) Delete() error {
	query := `DELETE FROM events WHERE id = ?`

	// Prepare the statement to prevent SQL injection and prepare SQL statements so that they can be executed many times efficiently
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	// Closes the statement after the function completes (to free resources).
	defer stmt.Close()

	// Execute the prepared statement, with the ID of the event to be deleted.
	_, err = stmt.Exec(e.ID)
	return err
}
