package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/abe/erp.api/configs"
	"github.com/abe/erp.api/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}

	configs.APP_NAME = os.Getenv("APP_NAME")
	configs.APP_VERSION = os.Getenv("APP_VERSION")
	configs.APP_DESCRIPTION = os.Getenv("APP_DESCRIPTION")
	configs.APP_BODY_LIMIT, _ = strconv.Atoi(os.Getenv("APP_BODY_LIMIT"))
	configs.APP_PORT, _ = strconv.Atoi(os.Getenv("APP_PORT"))
	configs.DB_HOST = os.Getenv("DB_HOST")
	configs.DB_PORT, _ = strconv.Atoi(os.Getenv("DB_PORT"))
	configs.DB_NAME = os.Getenv("DB_NAME")
	configs.DB_USER = os.Getenv("DB_USER")
	configs.DB_PASSWORD = os.Getenv("DB_PASSWORD")
	configs.APP_SECRET_KEY = os.Getenv("APP_SECRET_KEY")
	configs.APP_TRIGGER_API = os.Getenv("APP_TRIGGER_API")
}

func main() {
	// Create config variable
	config := fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  configs.APP_DESCRIPTION, // add custom server header
		AppName:       configs.APP_NAME,
		BodyLimit:     configs.APP_BODY_LIMIT, // this is the default limit of 10MB
	}

	app := fiber.New(config)
	app.Use(cors.New())
	app.Use(requestid.New())
	app.Use(logger.New())
	routes.Router(app)
	app.Static("/", "/public")
	app.Listen(fmt.Sprintf(":%d", configs.APP_PORT))
}
