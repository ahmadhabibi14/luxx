package middlewares

import (
	"time"

	"github.com/gofiber/fiber/v2/middleware/csrf"
)

var CSRFConfig = csrf.Config{
	KeyLookup:      "header:X-Csrf-Token",
	CookieName:     "csrf_",
	CookieSameSite: "Lax",
	Expiration:     1 * time.Hour,
}
