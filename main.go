package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/abe/erp.api/configs"
	"github.com/abe/erp.api/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
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
	configs.DB_SSLMODE = os.Getenv("DB_SSLMODE")
	configs.DB_TZNAME = os.Getenv("DB_TZNAME")
	configs.APP_SECRET_KEY = os.Getenv("APP_SECRET_KEY")
	configs.APP_TRIGGER_API = os.Getenv("APP_TRIGGER_API")

	dns := fmt.Sprintf("host=%s user=%s dbname=%s port=%d password=%s sslmode=%s TimeZone=%s", configs.DB_HOST, configs.DB_USER, configs.DB_NAME, configs.DB_PORT, configs.DB_PASSWORD, configs.DB_SSLMODE, configs.DB_TZNAME)
	configs.Store, err = gorm.Open(postgres.Open(dns), &gorm.Config{
		DisableAutomaticPing:                     true,
		DisableForeignKeyConstraintWhenMigrating: false,
		SkipDefaultTransaction:                   true,
		NowFunc: func() time.Time {
			return time.Now().Local()
		},
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "tbt_", // table name prefix, table for `User` would be `t_users`
			SingularTable: false,  // use singular table name, table for `User` would be `user` with this option enabled
			NoLowerCase:   false,  // skip the snake_casing of names
			NameReplacer:  strings.NewReplacer("CID", "Cid"),
		},
	})

	if err != nil {
		panic("Failed to connect to database")
	}

	// Auto Migration DB
	configs.SetDB()
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
