package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

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
	app := fiber.New()
	// app.Use(cors.New(middlewares.CORSConfig))
	// app.Use(csrf.New(middlewares.CSRFConfig))

	// All Backend services in /api endpoints
	api := app.Group("/api")
	routes.Auth(api)
	routes.OAuth2(api)
	routes.Users(api)

	log.Fatal(app.Listen(webdomain))
}
