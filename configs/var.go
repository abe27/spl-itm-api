package configs

import "gorm.io/gorm"

var (
	Store           *gorm.DB
	APP_NAME        string
	APP_VERSION     string
	APP_DESCRIPTION string
	APP_BODY_LIMIT  int
	APP_PORT        int
	DB_HOST         string
	DB_PORT         int
	DB_NAME         string
	DB_USER         string
	DB_PASSWORD     string
	APP_SECRET_KEY  string
	APP_TRIGGER_API string
)

func SetDB() error {
	var err error
	return err
}
