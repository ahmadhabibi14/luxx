package users

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"luxx/config"
	"luxx/database/sqlc"

	"github.com/gofiber/fiber/v2"
)

type (
	userDataError struct {
		ErrorMsg string `json:"error"`
	}
)

func UserData(c *fiber.Ctx) error {
	var db *sql.DB = config.ConnectDB()
	var errmsg userDataError
	ctx := context.Background()

	id, err := config.ExtractTokenID(c)
	if err != nil {
		errmsg.ErrorMsg = "Unexpected error occurs"
		errorResp, _ := json.Marshal(errmsg)
		return c.Status(fiber.StatusBadRequest).JSON(string(errorResp))
	}

	user_id := fmt.Sprintf("%v", id)
	queries := sqlc.New(db)

	// Get user data
	user_data, err := queries.GetUserDataByUserId(ctx, user_id)
	if err != nil {
		errmsg.ErrorMsg = "Cannot get user data"
		errorResp, _ := json.Marshal(errmsg)
		return c.Status(fiber.StatusBadRequest).JSON(string(errorResp))
	}

	defer db.Close()
	return c.Status(fiber.StatusOK).JSON(user_data)
}
