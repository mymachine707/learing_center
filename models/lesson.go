package models

import "time"

type lesson struct {
	ID        string     `json:"id" binding:"required" `
	Course    string     `json:"course" binding:"required" minLength:"4" maxLength:"50" example:"GOLANG"`
	Teacher   person     `json:"teacher" binding:"required"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
