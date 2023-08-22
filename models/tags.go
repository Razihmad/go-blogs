package models

import "time"

// Tag struct
type Tag struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Post      int       `json:"post"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
