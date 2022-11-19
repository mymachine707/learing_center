package models

import "time"

type Groups struct {
	ID          string
	GroupName   string
	lCenterInfo lCenter
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}
