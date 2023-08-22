package usecase

import (
	"strconv"

	db "github.com/Razihmad/blog_posts/config"
	"github.com/Razihmad/blog_posts/models"
	baseValidator "github.com/Razihmad/blog_posts/utils"
	"github.com/gofiber/fiber/v2"
)

func CommentToPost(c *fiber.Ctx, comment models.Comment) (*models.Comment, error) {
	err := baseValidator.ValidateStruct(comment)
	userId := c.Locals("userId")
	if userId == 0 {
		return nil, c.Status(400).JSON(fiber.Map{
			"message": "Cannot comment to post",
			"success": false,
			"data":    nil,
		})

	}
	if err != nil {
		return nil, err
	}
	comment.UserId = userId.(int)
	postId := c.Query("post_id")
	comment.PostId, _ = strconv.Atoi(postId)

	var user models.User
	var post models.Post
	db.DB.First(&user, userId)
	db.DB.First(&post, postId)
	if user.Id == 0 || post.Id == 0 {
		return nil, fiber.ErrNotFound
	}
	db.DB.Create(&comment)
	return &comment, nil
}

func GetAllCommentsOfPost(c *fiber.Ctx, postId string) ([]models.Comment, error) {
	var post models.Post
	db.DB.First(&post, postId)
	if post.Id == 0 {
		return nil, fiber.ErrNotFound
	}

	var comments []models.Comment
	db.DB.Where("post_id = ?", postId).Find(&comments)
	return comments, nil
}
