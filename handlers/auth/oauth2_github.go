package auth

import (
	"fmt"
	"os"

	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env files")
	}
}

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
	fmt.Println("to github")
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

func getGithubClientID() string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env files")
	}
	githubClientID, exists := os.LookupEnv("OAUTH2_GITHUB_CLIENT_ID")
	if !exists {
		log.Fatal("Github Client ID not defined in .env file")
	}
	return githubClientID
}

func getGithubClientSecret() string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env files")
	}
	githubClientSecret, exists := os.LookupEnv("OAUTH2_GITHUB_CLIENT_SECRET")
	if !exists {
		log.Fatal("Github Client ID not defined in .env file")
	}
	return githubClientSecret
}
