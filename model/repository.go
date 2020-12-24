package model

import "time"

type Repository struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	Visibility string    `json:"visibility"`
	Star       int       `json:"star"`
	CreatedAt  time.Time `json:"createdAt"`
}
