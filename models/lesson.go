package models

import "time"

// Lesson ...
type Lesson struct {
	ID              string     `json:"id" binding:"required" `
	Coursename      string     `json:"coursename" binding:"required" minLength:"4" maxLength:"50" example:"GOLANG"`
	LearingCenterID string     `json:"LearingCenterName" binding:"required"`
	Price           int        `json:"Price" binding:"required"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       *time.Time `json:"updated_at"`
	DeletedAt       *time.Time `json:"deleted_at"`
}

// CreateLessonModels ...
type CreateLessonModels struct {
	Coursename      string `json:"coursename" binding:"required" minLength:"4" maxLength:"50" example:"GOLANG"`
	LearingCenterID string `json:"LearingCenterName" binding:"required"`
	Price           int    `json:"Price" binding:"required"`
}

// UpdateLessonModels ...
type UpdateLessonModels struct {
	ID         string `json:"id" binding:"required" `
	Coursename string `json:"coursename" binding:"required" minLength:"4" maxLength:"50" example:"GOLANG"`
	Price      int    `json:"Price" binding:"required"`
}

// sql ok
