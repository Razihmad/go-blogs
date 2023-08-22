package models

import "time"

// Author struct
type Author struct {
	Id        int       `json:"id"`
	Name      string    `json:"name" validate:"required,min=3,max=20"`
	Email     string    `json:"email" validate:"required,email"`
	Password  string    `json:"password" validate:"required,min=6,max=20"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
