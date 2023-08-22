package routes

import (
	"github.com/Razihmad/blog_posts/controller"
	"github.com/Razihmad/blog_posts/middleware"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

const secretKey = "secret"

func SetUp(app *fiber.App) {
	private := app.Group("/api")
	private.Use(func(c *fiber.Ctx) error {
		middleware.ParseJWT(c)
		return nil
	})
	private.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("secret"),
	}))
	private.Get("/test", controller.Test)
	private.Post("/posts", controller.CreatePost)
	private.Post("/comments", controller.CommentToPost)
	private.Get("/comments/:postId", controller.GetAllCommentsOfPost)
	private.Put("/users", controller.UpdateUser)
	private.Get("/users", controller.GetAllUsers)
	app.Get("/posts", controller.GetAllPosts)
	app.Get("/post/:id", controller.GetPostById)
	app.Post("/authors", controller.CreateAuthor)
	app.Post("/users", controller.CreateUser)
	app.Get("/authors", controller.GetAllAuthors)
	app.Get("/author/:id", controller.GetAuthorById)

}
