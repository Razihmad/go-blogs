package dto

import "github.com/Razihmad/blog_posts/models"

type UserResponse struct {
	User  models.User `json:"user"`
	Token string      `json:"token"`
}
