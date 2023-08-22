package main

import (
	"fmt"

	"github.com/Razihmad/blog_posts/config"
	"github.com/Razihmad/blog_posts/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	config.ConnectDB()
	fmt.Println("connection successfully")
	app := fiber.New()
	routes.SetUp(app)
	app.Listen(":3000")

}
