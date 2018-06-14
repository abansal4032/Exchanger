package models

import "database/sql"

type User struct {
	UserID            string `json:"userId"`
	Name              string `json:"name"`
	Contact           string `json:"contact"`
	Email             string `json:"email"`
	Location          int    `json:"location"`
	Credits           int    `json:"credits"`
	RegistrationToken string `json:"registrationToken"`
}

type Entity struct {
	EntityID   string            `json:"entityId"`
	Name       string            `json:"name"`
	Type       string            `json:"type"`
	Owner      string            `json:"owner"`
	Action     string            `json:"actionType"`
	Status     string            `json:"status"`
	Price      int               `json:"price"`
	Borrower   sql.NullString    `json:"borrower"`
	Location   int               `json:"location"`
	Attributes map[string]string `json:"attributes"`
}

type Requests struct {
	RequestID        string `json:"requestId"`
	EntityID         string `json:"entityId"`
	Requester        string `json:"requester"`
	Intent           string `json:"intent"`
	DurationInDays   int    `json:"durationInDays"`
	Status           string `json:"status"`
	RequesterComment string `json:"requesterComment"`
	OwnerComment     string `json:"ownerComment"`
}
