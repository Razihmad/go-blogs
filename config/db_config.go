package config

import (
	"fmt"
	"os"

	"github.com/Razihmad/blog_posts/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	// Name is a variable that represents the name of the application
	DB *gorm.DB
)

func ConnectDB() error {
	godotenv.Load()
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbUri := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbName)
	db, err := gorm.Open(mysql.Open(dbUri), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return err
	}
	DB = db
	AutoMigrate(db)

	return nil
}

func AutoMigrate(db *gorm.DB) {
	db.Debug().AutoMigrate(
		&models.Post{},
		&models.Author{},
		&models.User{},
		&models.Comment{},
		&models.Tag{},
	)
}
