package auth

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

var (
	githubOAuthConfig = &oauth2.Config{
		ClientID:     os.Getenv("OAUTH2_GITHUB_CLIENT_ID"),
		ClientSecret: os.Getenv("OAUTH2_GITHUB_CLIENT_SECRET"),
		RedirectURL:  "http://localhost:3000/api/oauth/github/callback",
		Endpoint:     github.Endpoint,
	}
)

// /api/oauth/github/login
func Oauth2GithubLogin(c *fiber.Ctx) error {
	url := githubOAuthConfig.AuthCodeURL("state", oauth2.AccessTypeOffline)
	return c.Redirect(url)
}

// /api/oauth/github/callback
func Oauth2GithubCallback(c *fiber.Ctx) error {
	code := c.Query("code")
	token, err := githubOAuthConfig.Exchange(c.Context(), code)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	httpClient := githubOAuthConfig.Client(c.Context(), token)
	resp, err := httpClient.Get("https://api.github.com/user")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	defer resp.Body.Close()
	c.Set("Content-Type", "application/json")
	return c.Status(fiber.StatusOK).SendStream(resp.Body)
}
