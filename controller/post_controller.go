package controller

import (
	"fmt"

	"github.com/Razihmad/blog_posts/models"
	"github.com/Razihmad/blog_posts/usecase"
	"github.com/gofiber/fiber/v2"
)

// CreatePost godoc
// @Summary Create a new post
// @Description Create a new post
// @Tags posts
// @Accept  json

func CreatePost(c *fiber.Ctx) error {
	var post *models.Post
	authorId := c.Locals("authorId")
	fmt.Println("authorId: ", authorId)
	if authorId == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid author id",
		})
	}

	if err := c.BodyParser(&post); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Cannot parse JSON",
		})
	}
	post.AuthorId = authorId.(int)

	post, err := usecase.CreatePost(*post)
	if err != nil {
		message := fmt.Sprintf("Cannot create post: %v", err)
		fmt.Println(message)
		return c.Status(400).JSON(fiber.Map{
			"message": message,
			"success": false,
			"data":    nil,
		})
	}
	return c.Status(201).JSON(fiber.Map{
		"message": "Post created successfully",
		"success": true,
		"data":    post,
	})
}

func GetAllPosts(c *fiber.Ctx) error {
	posts, err := usecase.GetAllPosts()
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": fmt.Errorf("Cannot get all posts: %w", err),
			"success": false,
			"data":    nil,
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"message": "Posts retrieved successfully",
		"success": true,
		"data":    posts,
	})
}

func GetPostById(c *fiber.Ctx) error {
	id := c.Params("id")
	post, err := usecase.GetPostById(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": fmt.Errorf("Cannot get post by id: %w", err),
			"success": false,
			"data":    nil,
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"message": "Post retrieved successfully",
		"success": true,
		"data":    post,
	})
}
