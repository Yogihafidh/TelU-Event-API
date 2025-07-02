package models

import (
	"telu-event-apps/db"
	"telu-event-apps/utils"
)

type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

func (u *User) Save() error {
	// SQL query to insert a new user into the database. Use the ? placeholder to prevent SQL Injection.
	query := `
	INSERT INTO users (email, password) 
	VALUES (?, ?)
	`

	// Prepare the SQL statement to prevent SQL injection and prepare SQL statements so that they can be executed many times efficiently.
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	// Hash password sebelum disimpan
	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	// Save hash password to the User struct
	u.Password = hashedPassword

	// Execute the prepared statement, with the values from the User struct. The values (u.Email, u.Password) will be safely bound to ? in the query.
	result, err := stmt.Exec(u.Email, hashedPassword)
	if err != nil {
		return err
	}

	// Get the last inserted ID from the result of the Exec() method. This is useful if you want to retrieve the ID of the newly created user.
	id, err := result.LastInsertId()
	u.ID = id

	// Closes the statement after the function completes (to free resources).
	defer stmt.Close()

	// Return any error that occurred during the execution of the statement.
	return err
}

func (u User) ExistsByEmail() (bool, error) {
	query := `SELECT COUNT(*) FROM users WHERE email = ?`
	var count int
	err := db.DB.QueryRow(query, u.Email).Scan(&count)
	return count > 0, err
}

func (u *User) ValidateCredentials() (bool, error) {
	// SQL query to select the password from the users table where the email matches the provided email.
	query := `SELECT id, password FROM users WHERE email = ?`

	// Method QueryRow menjalankan query dan mengambil satu baris hasil. lalu method Scann akan membaca kolom hasil query (password) dari baris tersebut lalu enyalin nilainya ke variabel retrievedPassword.
	var retrievedPassword string
	err := db.DB.QueryRow(query, u.Email).Scan(&u.ID, &retrievedPassword)
	if err != nil {
		return false, err
	}

	// Verify the provided password against the retrieved password from the database.
	isValid, err := utils.VerifyPassword(u.Password, retrievedPassword)
	if err != nil {
		return false, err
	}

	return isValid, nil
}
