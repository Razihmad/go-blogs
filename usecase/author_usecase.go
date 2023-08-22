package usecase

import (
	"fmt"
	"time"

	db "github.com/Razihmad/blog_posts/config"
	"github.com/Razihmad/blog_posts/dto"
	"github.com/Razihmad/blog_posts/middleware"
	"github.com/Razihmad/blog_posts/models"
	baseValidator "github.com/Razihmad/blog_posts/utils"
	"github.com/gofiber/fiber/v2"
)

func CreateAuthor(c *fiber.Ctx, author models.Author) (*dto.AuthorResponse, error) {
	err := baseValidator.ValidateStruct(author)
	fmt.Println("err: ", err)
	if err != nil {
		return nil, err
	}
	author.CreatedAt = time.Time{}
	author.UpdatedAt = time.Time{}
	db.DB.Create(&author)
	token, _, err := middleware.GenerateJWTAuthor(author)
	if err != nil {
		return nil, err
	}
	return &dto.AuthorResponse{
		Author: author,
		Token:  token,
	}, nil

}

func GetAllAuthors(c *fiber.Ctx) ([]models.Author, error) {
	var authors []models.Author
	db.DB.Find(&authors)
	return authors, nil
}

func GetAuthorById(c *fiber.Ctx, id string) (*models.Author, error) {
	var author models.Author
	db.DB.First(&author, id)
	if author.Id == 0 {
		return nil, fiber.ErrNotFound
	}
	return &author, nil
}
