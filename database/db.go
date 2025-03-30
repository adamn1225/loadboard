package database

import (
	"fmt"
	"log"
	"os"

	"loadboard/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	log.Println("Connected to the database!")

	// Auto-migrate tables
	db.AutoMigrate(&models.Load{}, &models.LoadClaim{}, &models.User{})
	if err != nil {
		log.Fatal("Failed auto-migrating:", err)
	}

	DB = db
}
