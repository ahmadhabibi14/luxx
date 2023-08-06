package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

	"luxx/routes"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env files")
	}
}

func main() {
	webdomain := fmt.Sprintf("localhost:%s", os.Getenv("WEB_PORT"))
	app := fiber.New()
	app.Static("/", "./client")

	api := app.Group("/api") // All Backend services in /api endpoints
	routes.Oauth2(api)

	log.Fatal(app.Listen(webdomain))
}
