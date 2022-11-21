package models

import "time"

// Lesson ...
type Lesson struct {
	ID              string        `json:"id" binding:"required" `
	CourseName      string        `json:"courseName" binding:"required" minLength:"4" maxLength:"50" example:"GOLANG"`
	LearingCenterID LearingCenter `json:"LearingCenterName" binding:"required"`
	Price           int           `json:"Price" binding:"required"`
	CreatedAt       time.Time     `json:"created_at"`
	UpdatedAt       *time.Time    `json:"updated_at"`
	DeletedAt       *time.Time    `json:"deleted_at"`
}

// sql ok
