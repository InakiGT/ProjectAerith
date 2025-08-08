package db

import (
	"fmt"
	"log"
	"os"

	"rapi-pedidos/src/internal/user/infraestructure/persistence"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewGormConnection() *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to open database connection: %v", err)
	}

	// Migraciones de los Schemas
	if err := db.AutoMigrate(&persistence.User{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	return db
}
