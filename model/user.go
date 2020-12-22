package model

import "time"

type User struct {
	ID           string `json:"id"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	Repositories []string `json:"repositories"`
	CreatedAt    time.Time `json:"createdAt"`
}