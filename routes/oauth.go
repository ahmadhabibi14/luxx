package routes

import (
	"github.com/gofiber/fiber/v2"

	"luxx/handlers/auth"
)

func OAuth2(api fiber.Router) {
	oauth := api.Group("/oauth")

	// GitHub Oauth2 provider
	oauth.Post("/github/login", auth.Oauth2GithubLogin)
	oauth.Post("/github/callback", auth.Oauth2GithubCallback)
}
