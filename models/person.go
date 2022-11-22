package models

import "time"

// Person ...
type Person struct {
	ID          string     `json:"id" binding:"required"`
	Firstname   string     `json:"firstname" binding:"required" minLength:"4" maxLength:"50" example:"Islombek--example"`
	Lastname    string     `json:"lastname" binding:"required" minLength:"4" maxLength:"50" example:"Oripov--example"`
	Middlename  string     `json:"middlename" binding:"required" minLength:"4" maxLength:"50" example:"Erkinjon_o'g'li--example"`
	Birthday    string     `json:"birthday" binding:"required" minLength:"4" maxLength:"50" example:"09-05-1998--example"`
	Job         string     `json:"job" binding:"required" minLength:"4" maxLength:"50" example:"Administrator--example"`
	PhoneNumber string     `json:"phoneNumber" binding:"required" minLength:"4" maxLength:"50" example:"94-650-61-84--example"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}

// CreatePersonModul ...
type CreatePersonModul struct {
	Firstname   string `json:"firstname" binding:"required" minLength:"4" maxLength:"50" example:"Islombek--example"`
	Lastname    string `json:"lastname" binding:"required" minLength:"4" maxLength:"50" example:"Oripov--example"`
	Middlename  string `json:"middlename" binding:"required" minLength:"4" maxLength:"50" example:"Erkinjon_o'g'li--example"`
	Birthday    string `json:"birthday" binding:"required" minLength:"4" maxLength:"50" example:"09-05-1998--example"`
	Job         string `json:"job" binding:"required" minLength:"4" maxLength:"50" example:"Administrator--example"`
	PhoneNumber string `json:"phoneNumber" binding:"required" minLength:"4" maxLength:"50" example:"94-650-61-84--example"`
}

// UpdatePersonModel ...
type UpdatePersonModel struct {
	ID          string     `json:"id" binding:"required"`
	Firstname   string     `json:"firstname" binding:"required" minLength:"4" maxLength:"50" example:"Islombek--example"`
	Lastname    string     `json:"lastname" binding:"required" minLength:"4" maxLength:"50" example:"Oripov--example"`
	Middlename  string     `json:"middlename" binding:"required" minLength:"4" maxLength:"50" example:"Erkinjon_o'g'li--example"`
	Birthday    string     `json:"birthday" binding:"required" minLength:"4" maxLength:"50" example:"09-05-1998--example"`
	Job         string     `json:"job" binding:"required" minLength:"4" maxLength:"50" example:"Administrator--example"`
	PhoneNumber string     `json:"phoneNumber" binding:"required" minLength:"4" maxLength:"50" example:"94-650-61-84--example"`
}

// sql ok
