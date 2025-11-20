package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/shadykip/go-worker/internal/handlers"
	"github.com/shadykip/go-worker/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// DB setup
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		dbURL = "host=localhost user=dev password=linspace dbname=worker port=5432 sslmode=disable"
	}

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatal("DB connection failed")
	}
	db.AutoMigrate(&models.Job{})

	// Server
	r := gin.Default()
	r.POST("/jobs", handlers.EnqueueJob(db))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
