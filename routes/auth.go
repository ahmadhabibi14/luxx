package routes

import (
	"github.com/gofiber/fiber/v2"

	"luxx/handlers/auth"
)

// Common auth
func Auth(api fiber.Router) {
	cAuth := api.Group("/auth")

	cAuth.Post("/register", auth.Register)
}
