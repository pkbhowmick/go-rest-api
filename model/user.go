package model

import "time"

type User struct {
	UUID         string       `json:"uuid"`
	FirstName    string       `json:"firstName"`
	LastName     string       `json:"lastName"`
	Email        string       `json:"email"`
	Password     string       `json:"password"`
	Repositories []Repository `json:"repositories"`
	CreatedAt    time.Time    `json:"createdAt"`
}
