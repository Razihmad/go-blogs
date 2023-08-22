package models

import "time"

// Comment struct
type Comment struct {
	Id        int       `json:"id"`
	Content   string    `json:"content"`
	UserId    int       `json:"user_id"`
	PostId    int       `json:"post_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
