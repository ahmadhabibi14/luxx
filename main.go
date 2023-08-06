package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"

	"log"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env files")
	}
}

func main() {
	webdomain := fmt.Sprintf("localhost:%s", os.Getenv("WEB_PORT"))
	app := fiber.New()
	app.Static("/", "./client")

	api := app.Group("/api") // All Backend services in /api endpoints

	api.Get("/hello", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Hello World",
		})
	})

	app.Listen(webdomain)
}
