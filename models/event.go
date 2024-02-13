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

var events = []Event{}

func (e Event) Save() error {
	//Later: add it to a database
	query := `
	INSERT into events (name, description, location, dateTime, user_id) 
	VALUES (?, ?, ?, ?, ?)`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	e.ID = id

	return err
}

func GetAllEvents() []Event {
	return events
}