package models

import (
	"time"

	"example.com/event-booking-app/db"
)

type Event struct {
	ID          int64
	Name        string	`binding:"required"` // gin tags to enforce key presence in post body
	Description string	`binding:"required"`
	Location    string	`binding:"required"`
	DateTime    time.Time	`binding:"required"`
	UserID      int
}

func (event *Event) Save() error {
	//Later: add it to a database
	query := `
	INSERT into events (name, description, location, dateTime, user_id) 
	VALUES (?, ?, ?, ?, ?)`

	stmt, err := db.DB.Prepare(query) // stores the query in memory to be usable over and over with different values in a highly efficient way
	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.UserID)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	event.ID = id

	return err
}

func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM events"
	rows, err := db.DB.Query(query) //Prepare() could have been used here as well but was not intentionally - Used query instead of Exec because its just fetching data instead of modifying a record in the database

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var events []Event

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
	query := "SELECT * FROM events WHERE id = ?"
	row := db.DB.QueryRow(query, id)

	var event Event

	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)

	if err != nil {
		return nil, err // null value for pointer is nil hence the return type is a pointer for this function so that we can use nil here
	}

	return &event, nil
}

func (event Event) Update() error {
	query := `
	UPDATE events
	SET	name = ?, description = ?, location = ?, dateTime = ?
	WHERE id = ?
	`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.ID)
	return err
}

func (event Event) Delete() error {
	query := "DELETE FROM events WHERE id = ?"
	stmt, err := db.DB.Prepare(query)

	if err!= nil {
		return err
	}

	defer stmt.Close()

	_, err =stmt.Exec(event.ID)
	return err
}