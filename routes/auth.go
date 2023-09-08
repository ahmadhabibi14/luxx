package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"luxx/handlers/auth"
	"luxx/middlewares"
)

// Common auth
func Auth(api fiber.Router) {
	cAuth := api.Group("/auth")
	cAuth.Use(cors.New(middlewares.CORSConfig))

	cAuth.Post("/register", auth.Register)
}
