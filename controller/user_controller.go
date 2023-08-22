package controller

import (
	"fmt"

	"github.com/Razihmad/blog_posts/models"
	"github.com/Razihmad/blog_posts/usecase"
	"github.com/gofiber/fiber/v2"
)

func CreateUser(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Cannot parse JSON",
		})
	}
	user_response, err := usecase.CreateUser(c, user)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": fmt.Errorf("Cannot create user: %w", err),
			"success": false,
			"data":    nil,
		})
	}
	return c.Status(201).JSON(fiber.Map{
		"message": "User created successfully",
		"success": true,
		"data":    user_response,
	})

}

func UpdateUser(c *fiber.Ctx) error {
	var updateUser *models.User
	if err := c.BodyParser(&updateUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Cannot parse JSON",
		})
	}
	userId := c.Params("id")
	user, err := usecase.UpdateUser(c, *updateUser, userId)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": fmt.Errorf("Cannot update user: %w", err),
			"success": false,
			"data":    nil,
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"message": "User updated successfully",
		"success": true,
		"data":    user,
	})

}

func GetAllUsers(c *fiber.Ctx) error {
	users, err := usecase.GetAllUsers(c)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": fmt.Errorf("Cannot get all users: %w", err),
			"success": false,
			"data":    nil,
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"message": "Users retrieved successfully",
		"success": true,
		"data":    users,
	})
}

func Test(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}
