package auth

import (
	"database/sql"
	"encoding/json"

	"luxx/config"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type (
	registerInput struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Username string `json:"username"`
		Fullname string `json:"fullname"`
	}
	registerOut struct {
		Token  string `json:"token"`
		UserId string `json:"userId"`
	}
	registerError struct {
		ErrorMsg string `json:"error"`
	}
)

func Register(c *fiber.Ctx) error {
	var db *sql.DB = config.ConnectDB()
	var in registerInput
	var out registerOut
	var errmsg registerError

	if err := c.BodyParser(&in); err != nil {
		errmsg.ErrorMsg = "invalid input"
		errorResp, _ := json.Marshal(errmsg)
		return c.Status(fiber.StatusBadRequest).JSON(errorResp)
	}

	var checkUsername error
	checkUsername = db.QueryRow("SELECT username FROM Users WHERE username = ?", in.Username).Scan(&in.Username)
	if checkUsername == nil {
		errmsg.ErrorMsg = "Username already exists"
		errorResp, _ := json.Marshal(errmsg)
		return c.Status(fiber.StatusBadRequest).JSON(errorResp)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		errmsg.ErrorMsg = "Unable to register user"
		errorResp, _ := json.Marshal(errmsg)
		return c.Status(fiber.StatusBadRequest).JSON(errorResp)
	}
	in.Password = string(hashedPassword)
	_, err = db.Exec("INSERT INTO Users (username, email, password, fullname) VALUES (?, ?, ?, ?)", in.Username, in.Email, in.Password, in.Fullname)
	if err != nil {
		errmsg.ErrorMsg = "Unable to register user"
		errorResp, _ := json.Marshal(errmsg)
		return c.Status(fiber.StatusInternalServerError).JSON(errorResp)
	}

	defer db.Close()
	return c.JSON(out)
}
