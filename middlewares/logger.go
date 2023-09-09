package middlewares

import (
	"os"

	"github.com/gofiber/fiber/v2/middleware/logger"
)

func LoggerConfig(file *os.File) logger.Config {
	return logger.Config{
		Format:     "${time} >> ${latency} | HTPP Status: ${status} - ${method} ${path}\n",
		TimeFormat: "2006-01-02 03:04:05 PM",
		TimeZone:   "Asia/Makassar",
		Output:     file,
	}
}
