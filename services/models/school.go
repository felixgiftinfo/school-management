package models

import "github.com/google/uuid"

type School struct {
	Id      uuid.UUID `json:"id,omitempty"`
	Name    string    `json:"name,omitempty"`
	Address string    `json:"address,omitempty"`
	LGA     string    `json:"lga,omitempty"`
	State   string    `json:"state,omitempty"`
	Country string    `json:"country,omitempty"`
}
