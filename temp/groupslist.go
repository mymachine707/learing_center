package models

import "time"

// Groupslist list
type Groupslist struct {
	ID        string     `json:"id" binding:"required"`
	GroupName string     `json:"groupname" binding:"required" minLength:"4" maxLength:"50" example:"ALBION"`
	LCenterID string     `json:"learing_center_id" binding:"required" minLength:"4" maxLength:"50" example:"UUID_generator"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

// Teacher   Teacher    `json:"teacher_ID" binding:"required"`
