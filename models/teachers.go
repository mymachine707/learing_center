package models

import "time"

// Teacher ...
type Teacher struct {
	ID        string     `json:"id" binding:"required"`
	FullName  string     `json:"fullname" binding:"required" minLength:"4" maxLength:"50" example:"Islombek Erkinjon o'g'li Oripov--example"`
	LessonID  Lesson     `json:"LessonID" binding:"required"`
	GroupsID  string     `json:"GroupsID" binding:"required"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
