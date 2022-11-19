package models

import "time"

type Teacher struct {
	ID        string
	FullName  string
	CoursInfo lesson
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
