package models

import "time"

// LearingCenter center info
type LearingCenter struct {
	ID           string     `json:"id" binding:"required"`
	Direktor     person     `json:"direktor" binding:"required"`
	ListWorkers  string     `json:"listWorkers" binding:"required" minLength:"4" maxLength:"500000" example:"list of Cource--example"`
	ListCource   string     `json:"listCourse" binding:"required" minLength:"4" maxLength:"500000" example:"list of Cource--example"`
	ListTeachers string     `json:"teachersGroup" binding:"required" minLength:"4" maxLength:"500000" example:"List of --example"`
	Address      string     `json:"address" binding:"required" minLength:"4" maxLength:"50" example:"Toshkent,Yunusobod,Xasanboy_111--example"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at"`
}
