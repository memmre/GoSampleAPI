package models

import (
	"example.com/sample-api/database"
	"time"
)

type Event struct {
	ID          int64
	UserID      int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
}

func (event Event) Create() error {
	query := `
		INSERT INTO events (userID, name, description, location, dateTime) 
		VALUES (?, ?, ?, ?, ?)
	`
	statement, err := database.DATABASE.Prepare(query)
	if err != nil {
		return err
	}
	defer statement.Close()

	result, err := statement.Exec(event.UserID, event.Name, event.Description, event.Location, event.DateTime)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	event.ID = id
	return err
}

func (event Event) Delete() error {
	query := "DELETE FROM events WHERE id = ?"
	statement, err := database.DATABASE.Prepare(query)
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(event.ID)
	return err
}

func GetEvent(id int64) (*Event, error) {
	query := "SELECT * FROM events WHERE id = ?"
	row := database.DATABASE.QueryRow(query, id)

	var event Event
	err := row.Scan(&event.ID, &event.UserID, &event.Name, &event.Description, &event.Location, &event.DateTime)
	if err != nil {
		return nil, err
	}

	return &event, nil
}

func GetEvents() ([]Event, error) {
	query := "SELECT * FROM events ORDER BY dateTime DESC"

	rows, err := database.DATABASE.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event
	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.UserID, &event.Name, &event.Description, &event.Location, &event.DateTime)
		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}

func (event Event) Update() error {
	query := `
		UPDATE events 
		SET name = ?, description = ?, location = ?, dateTime = ?
		WHERE id = ?
	`
	statement, err := database.DATABASE.Prepare(query)
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(event.Name, event.Description, event.Location, event.DateTime, event.ID)
	return err
}
