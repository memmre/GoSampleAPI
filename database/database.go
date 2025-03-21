package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var DATABASE *sql.DB

func InitializeDatabase() {
	var err error
	DATABASE, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("Could not connect to database.")
	}

	DATABASE.SetMaxOpenConns(10)
	DATABASE.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createUsersTable := `
		CREATE TABLE IF NOT EXISTS users(
		    id INTEGER PRIMARY KEY AUTOINCREMENT,
		    emailAddress TEXT NOT NULL UNIQUE,
		    password TEXT NOT NULL
		)
	`
	_, err := DATABASE.Exec(createUsersTable)
	if err != nil {
		panic("Could not create users table.")
	}

	createEventsTable := `
		CREATE TABLE IF NOT EXISTS events(
		    id INTEGER PRIMARY KEY AUTOINCREMENT,
		    userID INTEGER NOT NULL,
		    name TEXT NOT NULL,
		    description TEXT NOT NULL,
		    location TEXT NOT NULL,
		    dateTime DATETIME NOT NULL,
		    FOREIGN KEY(userID) REFERENCES users(id)
		)
	`
	_, err = DATABASE.Exec(createEventsTable)
	if err != nil {
		panic("Could not create events table.")
	}
}
