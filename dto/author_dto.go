package dto

import "github.com/Razihmad/blog_posts/models"

type AuthorResponse struct {
	Author models.Author `json:"author"`
	Token  string        `json:"token"`
}
