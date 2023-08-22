package controller

import (
	"fmt"

	"github.com/Razihmad/blog_posts/models"
	"github.com/Razihmad/blog_posts/usecase"
	"github.com/gofiber/fiber/v2"
)

func CreateAuthor(c *fiber.Ctx) error {
	var author *models.Author
	if err := c.BodyParser(&author); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Cannot parse JSON",
		})
	}

	author_res, err := usecase.CreateAuthor(c, *author)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": fmt.Errorf("Cannot create author: %w", err),
			"success": false,
			"data":    nil,
		})
	}
	fmt.Println("author", author)
	return c.Status(201).JSON(fiber.Map{
		"message": "Author created successfully",
		"success": true,
		"data":    author_res,
	})

}

func GetAllAuthors(c *fiber.Ctx) error {
	authors, err := usecase.GetAllAuthors(c)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": fmt.Errorf("Cannot get all authors: %w", err),
			"success": false,
			"data":    nil,
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"message": "Authors retrieved successfully",
		"success": true,
		"data":    authors,
	})
}

func GetAuthorById(c *fiber.Ctx) error {
	id := c.Params("id")
	author, err := usecase.GetAuthorById(c, id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": fmt.Errorf("Cannot get author by id: %w", err),
			"success": false,
			"data":    nil,
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"message": "Author retrieved successfully",
		"success": true,
		"data":    author,
	})
}
