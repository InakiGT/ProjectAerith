package main

import (
	"log"
	"rapi-pedidos/src/internal/shared/infrastructure/db"
	"rapi-pedidos/src/internal/user/application"
	"rapi-pedidos/src/internal/user/infraestructure/hashing"
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
	dbConn := db.NewGormConnection()

	userRepo := persistence.NewGormRepository(dbConn)
	hasher := hashing.NewBcryptHasher()
	createUser := application.NewCreateUser(userRepo, hasher)
	findAllUsers := application.NewFindAllUsers(userRepo)
	userHandler := http.NewUserHandler(*createUser, *findAllUsers)

	// Router
	router := gin.Default()
	router.POST("/users", userHandler.CreateUser)
	router.GET("/users", userHandler.FindAllUsers)

	router.Run()
}
