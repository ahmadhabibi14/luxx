package auth

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"

	"luxx/config"
	"luxx/database/sqlc"
	"luxx/utils"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type (
	registerInput struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=8,containsany=!@#?*%&>_<}-{+"`
		Username string `json:"username" validate:"required,omitempty,min=4"`
		Fullname string `json:"fullname" validate:"required"`
	}
	registerOut struct {
		Token    string `json:"token"`
		Username string `json:"username"`
		Message  string `json:"message"`
	}
	registerError struct {
		ErrorMsg string `json:"error"`
	}
)

func Register(c *fiber.Ctx) error {
	var db *sql.DB = config.ConnectDB()
	queries := sqlc.New(db)
	ctx := context.Background()
	var (
		in     registerInput
		out    registerOut
		errmsg registerError
	)

	if err := c.BodyParser(&in); err != nil {
		errmsg.ErrorMsg = "invalid input"
		errorResp, _ := json.Marshal(errmsg)
		return c.Status(fiber.StatusBadRequest).JSON(string(errorResp))
	}

	isValid := utils.ValidateStruct(in)
	if isValid != nil {
		errmsg.ErrorMsg = isValid.Error()
		errorResp, _ := json.Marshal(errmsg)
		return c.Status(fiber.StatusBadRequest).JSON(string(errorResp))
	}

	_, isUsernameExist := queries.GetUserByUsername(ctx, in.Username)
	if isUsernameExist == nil {
		errmsg.ErrorMsg = "Username already exists"
		errorResp, _ := json.Marshal(errmsg)
		return c.Status(fiber.StatusBadRequest).JSON(string(errorResp))
	}

	_, isEmailExist := queries.GetUserByEmail(ctx, in.Email)
	if isEmailExist == nil {
		errmsg.ErrorMsg = "Username already exists"
		errorResp, _ := json.Marshal(errmsg)
		return c.Status(fiber.StatusBadRequest).JSON(string(errorResp))
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		errmsg.ErrorMsg = "Unable to register user"
		errorResp, _ := json.Marshal(errmsg)
		return c.Status(fiber.StatusBadRequest).JSON(string(errorResp))
	}
	in.Password = string(hashedPassword)
	user_id := utils.GenerateRandomID(10)
	_, err = db.Exec(
		"INSERT INTO Users (user_id, username, full_name, email, password) VALUES (?, ?, ?, ?, ?)",
		user_id,
		in.Username,
		in.Fullname,
		in.Email,
		in.Password,
	)
	if err != nil {
		errmsg.ErrorMsg = "Unable to register user"
		errorResp, _ := json.Marshal(errmsg)
		return c.Status(fiber.StatusInternalServerError).JSON(string(errorResp))
	}

	token, err := config.GenerateJWT(user_id, time.Now().AddDate(0, 2, 0))
	if err != nil {
		errmsg.ErrorMsg = "Unable to generate session token"
		errorResp, _ := json.Marshal(errmsg)
		return c.Status(fiber.StatusInternalServerError).JSON(string(errorResp))
	}

	out = registerOut{
		Token:    token,
		Username: in.Username,
		Message:  "User created successfully!",
	}
	successResp, _ := json.Marshal(out)
	config.SetJWTasCookie(c, token, time.Now().AddDate(0, 2, 0))

	defer db.Close()
	return c.Status(fiber.StatusOK).JSON(string(successResp))
}
