package model

import "time"

type Repository struct {
	UUID       string    `json:"uuid"`
	Name       string    `json:"name"`
	Visibility string    `json:"visibility"`
	Star       int       `json:"star"`
	CreatedAt  time.Time `json:"createdAt"`
}
