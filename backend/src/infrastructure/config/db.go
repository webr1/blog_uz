package config

import (
	"fmt"
	"log"
	"os"

	"blogapp/src/infrastructure/persistence/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase() *gorm.DB {
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

	return db
}

func RunMigrations(db *gorm.DB) {
	err := db.AutoMigrate(
		&models.UserModel{},
		&models.PostModel{},
		&models.CommentModel{},
		&models.LikeModel{},
		&models.ProfileModel{},
	)
	if err != nil {
		log.Fatal("Failed to run migrations:", err)
	}
}