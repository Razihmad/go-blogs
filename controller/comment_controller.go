package controller

import (
	"fmt"

	"github.com/Razihmad/blog_posts/models"
	"github.com/Razihmad/blog_posts/usecase"
	"github.com/gofiber/fiber/v2"
)

func CommentToPost(c *fiber.Ctx) error {
	var comment *models.Comment
	if err := c.BodyParser(&comment); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Cannot parse JSON",
		})
	}
	post_id := c.Params("post_id")
	fmt.Println("post id:", post_id)
	comment, err := usecase.CommentToPost(c, *comment)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": fmt.Sprintf("Cannot comment to post: %v", err),
			"success": false,
			"data":    nil,
		})
	}
	return c.Status(201).JSON(fiber.Map{
		"message": "Commented to post successfully",
		"success": true,
		"data":    comment,
	})

}

func GetAllCommentsOfPost(c *fiber.Ctx) error {
	postId := c.Params("postId")
	fmt.Printf(postId)
	comments, err := usecase.GetAllCommentsOfPost(c, postId)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": fmt.Sprintf("Cannot get all comments of post: %v", err),
			"success": false,
			"data":    nil,
		})
	}
	fmt.Print("comments: ", comments)
	return c.Status(200).JSON(fiber.Map{
		"message": "Comments retrieved successfully",
		"success": true,
		"data":    comments,
	})

}
