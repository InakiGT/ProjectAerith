package main

import (
	"log"

	"rapi-pedidos/src/cmd/api/gateways"
	"rapi-pedidos/src/internal/shared/infrastructure/db"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: Error loading .env file: %v", err)
	}

	// Connection to the database
	dbConn := db.NewGormConnection()

	// Router
	router := gin.Default()

	gateways.UserBindRoutes(dbConn, router)
	gateways.AddressBindRoutes(dbConn, router)
	gateways.ProductBindRoutes(dbConn, router)
	gateways.VehicleBindRoutes(dbConn, router)
	gateways.ClientCardBindRoutes(dbConn, router)

	router.Run()
}
