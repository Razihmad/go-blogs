package models

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type Post struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	AuthorId  int       `json:"author_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
