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

	MAIL_SMTP      string
	MAIL_SMTP_PORT string
	MAIL_USERNAME  string
	MAIL_PASSWORD  string
)

func SetDB() {
	Store.AutoMigrate(
		&models.User{},
		&models.SystemLogger{},
		&models.JwtToken{},
		&models.Administrator{},
		&models.Position{},
		&models.Section{},
		&models.Department{},
		&models.Area{},
		&models.Whs{},
		&models.Factory{},
		&models.Unit{},
		&models.ItemType{},
		&models.Shipment{},
		&models.MailBox{},
		&models.MailType{},
		&models.DownloadMailBox{},
		&models.Part{},
	)
}
