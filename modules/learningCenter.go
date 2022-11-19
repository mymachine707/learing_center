package modules

import "time"

type lCenter struct {
	id            string
	Direktor      person
	Administrator person
	CoursesGroup  string
	TeachersGroup string
	Address       string
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     *time.Time `json:"updated_at"`
	DeletedAt     *time.Time `json:"deleted_at"`
}
