package models

import "time"

// Teacher ...
type Teacher struct {
	ID        string     `json:"id" binding:"required"`
	PersonID  string     `json:"PersonID" binding:"required"`
	FullName  string     `json:"fullname" binding:"required" minLength:"4" maxLength:"50" example:"Islombek Erkinjon o'g'li Oripov--example"`
	LessonID  string     `json:"LessonID" binding:"required"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

// sql ok
// GroupsID  string     `json:"GroupsID" binding:"required"`
