package routes

import (
	"github.com/gofiber/fiber/v2"

	"luxx/handlers/users"
	"luxx/middlewares"
)

// Common auth
func Users(api fiber.Router) {
	user := api.Group("/user")

	user.Post("/user-data", middlewares.AuthJWT, users.UserData)
}
