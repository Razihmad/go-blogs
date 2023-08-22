package usecase

import (
	"time"

	db "github.com/Razihmad/blog_posts/config"
	"github.com/Razihmad/blog_posts/models"
	baseValidator "github.com/Razihmad/blog_posts/utils"
	"github.com/gofiber/fiber/v2"
)

func CreatePost(post models.Post) (*models.Post, error) {
	err := baseValidator.ValidateStruct(post)
	if err != nil {
		return &post, err
	}

	authorId := post.AuthorId
	var author models.Author
	db.DB.First(&author, authorId)
	if author.Id == 0 {
		return nil, fiber.ErrNotFound
	}
	post.CreatedAt = time.Time{}
	post.UpdatedAt = time.Time{}
	db.DB.Create(&post)
	return &post, nil
}

func GetAllPosts() ([]models.Post, error) {
	var posts []models.Post
	db.DB.Find(&posts)
	return posts, nil
}

func GetPostById(id string) (models.Post, error) {
	var post models.Post
	db.DB.First(&post, id)
	return post, nil
}
