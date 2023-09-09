package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"

	"luxx/middlewares"
	"luxx/routes"
	"runtime"
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalln("Error loading .env files")
	}
	if cpu := runtime.NumCPU(); cpu == 1 {
		runtime.GOMAXPROCS(2)
	} else {
		runtime.GOMAXPROCS(cpu)
	}
}

func main() {
	webdomain := fmt.Sprintf("localhost:%s", os.Getenv("WEB_PORT"))
	file, err := os.OpenFile("./tmp/web-server.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer file.Close()
	// Define file to logs
	app := fiber.New()
	app.Static("/", "web/dist/")
	app.Get("*", func(c *fiber.Ctx) error {
		return c.SendFile("./web/dist/index.html")
	})
	app.Use(logger.New(middlewares.LoggerConfig(file)))
	app.Use(cors.New(middlewares.CORSConfig))

	// All Backend services in /api endpoints
	api := app.Group("/api")
	routes.Auth(api)
	routes.OAuth2(api)
	routes.Users(api)

	log.Fatal(app.Listen(webdomain))
}
