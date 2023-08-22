package main_test

import (
	"bytes"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/Razihmad/blog_posts/models"
	"github.com/gavv/httpexpect"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-testfixtures/testfixtures/v3"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	app        *fiber.App
	db         *sql.DB
	fixtures   *testfixtures.Loader
	httpExpect *httpexpect.Expect
)

func TestMain(m *testing.M) {
	var err error
	// app = setupTest()
	godotenv.Load()
	dbHost := os.Getenv("DB_HOST")
	dbName := "test"
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbUri := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbName)
	DB, err := gorm.Open(mysql.Open(dbUri), &gorm.Config{})
	if err != nil {
		fmt.Printf("Error opening database: %q", err)
	}
	DB.AutoMigrate(&models.User{}, &models.Author{}, &models.Post{}, &models.Comment{})
	db, err = sql.Open("mysql", dbUri)
	if err != nil {
		fmt.Printf("Error opening database: %q", err)
	}
	fixtures, err = testfixtures.New(
		testfixtures.Database(db),          // You database connection
		testfixtures.Dialect("mysql"),      // Available: "postgresql", "timescaledb", "mysql", "mariadb", "sqlite" and "sqlserver"
		testfixtures.Directory("fixtures"), // The directory containing the YAML files
	)
	if err != nil {
		fmt.Printf("Error loading fixtures: %q", err)
	}
	if err := fixtures.Load(); err != nil {
		fmt.Printf("Error loading fixtures: %q", err)
	}
	os.Exit(m.Run())
}

func TestGetAllPosts(t *testing.T) {
	t.Run("Get all posts", func(t *testing.T) {
		resp, err := http.Get("http://localhost:3000/posts")
		if err != nil {
			t.Error(err)
		}

		if resp.StatusCode != 200 {
			t.Errorf("Expected status code 200, got %d", resp.StatusCode)
		}
	})
}

func TestCreatePost(t *testing.T) {
	t.Run("Create post", func(t *testing.T) {
		requestData := []byte(`{"title":"test","content":"test"}`)
		req, err := http.NewRequest("POST", "http://localhost:3000/api/posts", bytes.NewBuffer(requestData))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpZCI6MSwiZXhwIjoxNjkzMzA5NTI5fQ.aXCUtCQsoHDBc-PMg3hyPEkBBEQTUXFf2b5VX8d87J8")
		client := http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			t.Error(err)
		}
		if resp.StatusCode != 201 {
			t.Errorf("Expected status code 201, got %d", resp.StatusCode)
		}
	})

}
