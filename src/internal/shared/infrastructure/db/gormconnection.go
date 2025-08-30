package db

import (
	"fmt"
	"log"
	"os"

	addresspersistence "rapi-pedidos/src/internal/address/infrastructure/persistence"
	clientpersistence "rapi-pedidos/src/internal/client/infrastructure/persistence"
	clientcardpersistence "rapi-pedidos/src/internal/client_card/infrastructure/persistence"
	commercepersistence "rapi-pedidos/src/internal/commerce/infrastructure/persistence"
	commercecategorypersistence "rapi-pedidos/src/internal/commerce_category/infrastructure/persistence"
	deliverypersonpersistence "rapi-pedidos/src/internal/delivery_person/infrastructure/persistence"
	productpersistence "rapi-pedidos/src/internal/product/infrastructure/persistence"
	userpersistence "rapi-pedidos/src/internal/user/infrastructure/persistence"
	vehiclepersistence "rapi-pedidos/src/internal/vehicle/infrastructure/persistence"

	_ "github.com/lib/pq"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewGormConnection() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to open database connection: %v", err)
	}

	// Migraciones de los Schemas
	if err := db.AutoMigrate(&userpersistence.User{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	if err := db.AutoMigrate(&addresspersistence.Address{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	if err := db.AutoMigrate(&productpersistence.Product{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	if err := db.AutoMigrate(&deliverypersonpersistence.DeliveryPerson{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	if err := db.AutoMigrate(&vehiclepersistence.Vehicle{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	if err := db.AutoMigrate(&clientpersistence.Client{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	if err := db.AutoMigrate(&clientcardpersistence.ClientCard{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	if err := db.AutoMigrate(&commercecategorypersistence.CommerceCategory{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	if err := db.AutoMigrate(&commercepersistence.Commerce{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	return db
}
