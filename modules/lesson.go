package models

import "time"

type lesson struct {
	ID        string
	Course    string
	Teacher   person
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
