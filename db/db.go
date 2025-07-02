package db

import (
	"database/sql"

	// Import  driver (the SQLite)
	_ "github.com/glebarez/go-sqlite"
)

// DB global variable is of type pointer to sql.DB.
var DB *sql.DB

func InitDB() {
	// Open a database connection. sql.Open does not directly open a connection, but sets up the connection configuration.
	var err error
	DB, err = sql.Open("sqlite", "events.db")
	if err != nil {
		panic("could not connect to database")
	}

	// Connection performance settings
	DB.SetMaxOpenConns(10) // Limits the maximum total number of connections that may be open to a database at any one time.
	DB.SetMaxIdleConns(5)  // Sets the number of connections that can remain active but idle in the pool.

	createTables()
}

func createTables() {
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	);
	`
	_, err := DB.Exec(createUsersTable)
	if err != nil {
		panic("could not create users table: " + err.Error())
	}

	// Set Query to create the events table if it does not exist
	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		user_id INTEGER,
		FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
	);`

	// Execute the query to create the table
	_, err = DB.Exec(createEventsTable)
	if err != nil {
		panic("could not create events table: " + err.Error())
	}

}
