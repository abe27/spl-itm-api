package configs

import (
	"github.com/abe/erp.api/models"
	"gorm.io/gorm"
)

var (
	Store           *gorm.DB
	APP_NAME        string
	APP_VERSION     string
	APP_DESCRIPTION string
	APP_BODY_LIMIT  int
	APP_PORT        int
	APP_PUBLIC_DIRS string
	DB_HOST         string
	DB_PORT         int
	DB_NAME         string
	DB_USER         string
	DB_PASSWORD     string
	DB_SSLMODE      string
	DB_TZNAME       string
	APP_SECRET_KEY  string
	APP_TRIGGER_API string
)

func SetDB() {
	Store.AutoMigrate(
		&models.User{},
		&models.JwtToken{},
		&models.Administrator{},
		&models.Position{},
		&models.Section{},
		&models.Department{},
		&models.Whs{},
		&models.Factory{},
		&models.Unit{},
		&models.Shipment{},
		&models.SystemLogger{},
	)
}
