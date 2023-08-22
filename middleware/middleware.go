package middleware

import (
	"fmt"
	"time"

	"github.com/Razihmad/blog_posts/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func GenerateJWTAuthor(author models.Author) (string, int64, error) {
	exp := time.Now().Add(time.Hour * 24 * 7).Unix()
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorid"] = author.Id
	claims["exp"] = exp
	s, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", 0, err
	}
	return s, exp, nil

}
func GenerateJWTUser(user models.User) (string, int64, error) {
	exp := time.Now().Add(time.Hour * 24 * 7).Unix()
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["userid"] = user.Id
	claims["exp"] = exp
	s, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", 0, err
	}
	return s, exp, nil

}

func ParseJWT(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if len(authHeader) < 7 || authHeader[:7] != "Bearer " {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid or missing token",
		})
	}
	tokenString := authHeader[7:]
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid tokens",
		})
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid token claims",
		})
	}
	var authorId, userId float64
	if claims["authorid"] == nil {
		fmt.Println("authorid is nil")
		authorId = 0
	} else {
		authorId = claims["authorid"].(float64)
	}

	if claims["userid"] == nil {
		fmt.Println("userid is nil")
		userId = 0
	} else {
		userId = claims["userid"].(float64)
	}

	c.Locals("authorId", int(authorId))
	c.Locals("userId", int(userId))
	fmt.Println("userId", userId)
	fmt.Println("authorId", authorId)
	return c.Next()
}
