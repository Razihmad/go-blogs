package usecase

import (
	"time"

	db "github.com/Razihmad/blog_posts/config"
	"github.com/Razihmad/blog_posts/dto"
	"github.com/Razihmad/blog_posts/middleware"
	"github.com/Razihmad/blog_posts/models"
	baseValidator "github.com/Razihmad/blog_posts/utils"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(c *fiber.Ctx, user models.User) (*dto.UserResponse, error) {
	err := baseValidator.ValidateStruct(user)
	if err != nil {
		return nil, err
	}
	user.CreatedAt = time.Time{}
	user.UpdatedAt = time.Time{}
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user.Password = string(hash)
	db.DB.Create(&user)
	token, _, err := middleware.GenerateJWTUser(user)
	if err != nil {
		return nil, err
	}
	return &dto.UserResponse{
		User:  user,
		Token: token,
	}, nil
}

func UpdateUser(c *fiber.Ctx, updateUser models.User, userId string) (*models.User, error) {
	err := baseValidator.ValidateStruct(updateUser)
	if err != nil {
		return nil, err
	}
	var user models.User
	db.DB.First(&user, userId)
	if user.Id == 0 {
		return nil, fiber.ErrNotFound
	}
	db.DB.Model(&user).Updates(updateUser)
	return &user, nil
}

func GetAllUsers(c *fiber.Ctx) ([]models.User, error) {
	var users []models.User
	db.DB.Find(&users)
	return users, nil
}
