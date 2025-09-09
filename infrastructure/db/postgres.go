package db

import (
	"log"
	"trello-services/infrastructure/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresDB() *gorm.DB {
	dsn := config.GetEnv("DATABASE_URL", "")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Connected to Neon Postgres")
	return db
}
