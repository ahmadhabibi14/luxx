package auth

import (
	"context"
	"database/sql"
	"encoding/json"
	"luxx/config"
	"luxx/database/sqlc"

	"github.com/gofiber/fiber/v2"
)

type (
	loginInput struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}
	loginOut struct {
		Token    string `json:"token"`
		Username string `json:"username"`
		Message  string `json:"message"`
	}
	loginError struct {
		ErrorMsg string `json:"error"`
	}
)

func Login(c *fiber.Ctx) error {
	var db *sql.DB = config.ConnectDB()
	queries := sqlc.New(db)
	ctx := context.Background()
	var (
		in     loginInput
		out    loginOut
		errmsg loginError
	)
	_, isEmailExist := queries.GetUserByEmail(ctx, in.Email)
	if isEmailExist == nil {
		errmsg.ErrorMsg = "Username already exists"
		errorResp, _ := json.Marshal(errmsg)
		return c.Status(fiber.StatusBadRequest).JSON(string(errorResp))
	}
}
