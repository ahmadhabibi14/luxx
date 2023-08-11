package middlewares

import (
	"luxx/config"

	"github.com/gofiber/fiber/v2"
)

func AuthJWT(c *fiber.Ctx) error {
	err := config.TokenValid(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized Access",
		})
	}
	return c.Next()
}
