package auth

import (
	"io"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gofiber/fiber/v2"
)

func TestRegister(t *testing.T) {
	app := fiber.New()
	app.Post("/api/auth/register", Register)

	payload := `{"email":"mbahlorem@proton.me", "passsword":"mb4hl0r3m@##", "username":"mbahlorem63", "fullname":"Mbah Lorem"}`
	req := httptest.NewRequest(fiber.MethodPost, "/api/auth/register", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")

	// Http response
	resp, _ := app.Test(req)

	assert.Equal(t, fiber.StatusOK, resp.StatusCode, "Register failed")
	// Do something with results:
	body, _ := io.ReadAll(resp.Body) // Get all response body
	if resp.StatusCode == fiber.StatusOK {
		t.Logf("Response Body : %s\n", string(body))
	} else {
		t.Errorf("Error: Invalid request\n")
	}
}
