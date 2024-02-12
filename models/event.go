package models

import "time"

type Event struct {
	ID          int
	Name        string	`binding:"required"` // gin tags to enforce key presence in post body
	Description string	`binding:"required"`
	Location    string	`binding:"required"`
	DateTime    time.Time	`binding:"required"`
	UserID      int
}

var events = []Event{}

func (e Event) Save() {
	//Later: add it to a database
	events = append(events, e)
}

func GetAllEvents() []Event {
	return events
}