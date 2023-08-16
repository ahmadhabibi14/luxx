package users

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"luxx/config"

	"github.com/gofiber/fiber/v2"
)

type (
	userDataOut struct {
		Username string    `json:"username"`
		Fullname string    `json:"fullname"`
		Email    string    `json:"email"`
		Avatar   string    `json:"avatar"`
		JoinAt   time.Time `json:"join_at"`
	}
	userDataError struct {
		ErrorMsg string `json:"error"`
	}
)

func UserData(c *fiber.Ctx) error {
	var db *sql.DB = config.ConnectDB()
	var out userDataOut
	var errmsg userDataError

	id, err := config.ExtractTokenID(c)
	if err != nil {
		errmsg.ErrorMsg = "Unexpected error occurs"
		errorResp, _ := json.Marshal(errmsg)
		return c.Status(fiber.StatusBadRequest).JSON(string(errorResp))
	}

	user_id := fmt.Sprintf("%v", id)
	var get_user_data error
	get_user_data = db.QueryRow(
		"SELECT username, full_name, email, avatar, join_at FROM Users WHERE user_id = ?",
		user_id,
	).Scan(
		&out.Username,
		&out.Fullname,
		&out.Email,
		&out.Avatar,
		&out.JoinAt,
	)
	if get_user_data == nil {
		errmsg.ErrorMsg = "Cannot get user"
		errorResp, _ := json.Marshal(errmsg)
		return c.Status(fiber.StatusBadRequest).JSON(string(errorResp))
	}

	successResp, _ := json.Marshal(out)

	defer db.Close()
	return c.Status(fiber.StatusOK).JSON(string(successResp))
}
