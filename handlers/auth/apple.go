package auth

import (
	"os"

	"golang.org/x/oauth2"
)

var (
	appleOAuthConfig = &oauth2.Config{
		ClientID:     os.Getenv("OAUTH2_GITHUB_CLIENT_ID"),
		ClientSecret: os.Getenv("OAUTH2_GITHUB_CLIENT_SECRET"),
		RedirectURL:  "http://localhost:3000/api/oauth/github/callback",
		Scopes:       []string{"email"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://appleid.apple.com/auth/keys",
			TokenURL: "https://appleid.apple.com/auth/token",
		},
	}
)
