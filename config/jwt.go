package config

import (
	"os"
	"time"

	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func GenerateJWT(user_id string, expirationTime time.Time) (tokenString string, err error) {
	claims := jwt.MapClaims{
		"authorized": true,
		"user_id":    user_id,
		"exp":        expirationTime.Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return "Generate JWT Error", err
	}
	return tokenString, nil
}

func TokenValid(c *fiber.Ctx) error {
	tokenString := ExtractToken(c)
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})
	if err != nil {
		return err
	}
	return nil
}

func ExtractToken(c *fiber.Ctx) string {
	bearerToken := c.Cookies("auth")
	if bearerToken == "" {
		return ""
	}
	return bearerToken
}

func ExtractTokenID(c *fiber.Ctx) (interface{}, error) {
	tokenString := ExtractToken(c)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})
	if err != nil {
		return "", err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		uid := claims["user_id"]
		return uid, nil
	}
	return "", nil
}
