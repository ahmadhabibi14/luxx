package auth

import (
	"context"
	"database/sql"
	"luxx/config"
	"luxx/database/sqlc"
	"time"

	json "github.com/goccy/go-json"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type (
	loginInput struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}
	loginOut struct {
		Token   string `json:"token"`
		UserID  string `json:"user_id"`
		Message string `json:"message"`
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

	if err := c.BodyParser(&in); err != nil {
		errmsg.ErrorMsg = invalid_input
		errorResp, _ := json.Marshal(errmsg)
		return c.Status(fiber.StatusBadRequest).JSON(string(errorResp))
	}

	userLoginRow, isUserExist := queries.UserLogin(ctx, in.Email)
	if isUserExist != nil {
		errmsg.ErrorMsg = user_not_found
		errorResp, _ := json.Marshal(errmsg)
		return c.Status(fiber.StatusBadRequest).JSON(string(errorResp))
	}

	err := verifyPassword(in.Password, userLoginRow.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		errmsg.ErrorMsg = invalid_password
		errorResp, _ := json.Marshal(errmsg)
		return c.Status(fiber.StatusBadRequest).JSON(string(errorResp))
	}

	token, err := config.GenerateJWT(userLoginRow.UserID, time.Now().AddDate(0, 2, 0))
	if err != nil {
		errmsg.ErrorMsg = error_generate_token
		errorResp, _ := json.Marshal(errmsg)
		return c.Status(fiber.StatusInternalServerError).JSON(string(errorResp))
	}

	out = loginOut{
		Token:   token,
		UserID:  userLoginRow.UserID,
		Message: login_success,
	}
	successResp, _ := json.Marshal(out)
	config.SetJWTasCookie(c, token, time.Now().AddDate(0, 2, 0))
	defer db.Close()
	return c.Status(fiber.StatusOK).JSON(string(successResp))
}

func verifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
