package database

import (
	"github.com/thitiphongD/go-backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() (*gorm.DB, error) {
	dsn := "host=127.0.0.1 user=myuser password=mypassword dbname=mydb port=5432 sslmode=disable TimeZone=Asia/Bangkok"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	DB = db
	err = db.AutoMigrate(
		&models.User{},
		&models.Manga{},
	)
	if err != nil {
		return nil, err
	}
	return db, nil
}
