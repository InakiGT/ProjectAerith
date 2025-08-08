package main

import (
	"log"
	"rapi-pedidos/src/internal/shared/infrastructure/db"
	"rapi-pedidos/src/internal/user/application"
	"rapi-pedidos/src/internal/user/infraestructure/http"
	"rapi-pedidos/src/internal/user/infraestructure/persistence"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: Error loading .env file: %v", err)
	}

	// Connection to the database
	dbConn := db.NewPostgresConnection()

	userRepo := persistence.NewPgRepository(dbConn)
	createUser := application.NewCreateUser(userRepo)
	userHandler := http.NewUserHandler(*createUser)

	// Router
	router := gin.Default()
	router.POST("/users", userHandler.CreateUser)

	router.Run()
}
