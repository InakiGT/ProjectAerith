package gateways

import (
	"rapi-pedidos/src/internal/user/application"
	"rapi-pedidos/src/internal/user/infrastructure/hashing"
	"rapi-pedidos/src/internal/user/infrastructure/http"
	"rapi-pedidos/src/internal/user/infrastructure/persistence"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserBindRoutes(dbConn *gorm.DB, router *gin.Engine) {
	userRepo := persistence.NewGormRepository(dbConn)
	hasher := hashing.NewBcryptHasher()
	createUser := application.NewCreateUser(userRepo, hasher)
	findAllUsers := application.NewFindAllUsers(userRepo)
	findUserById := application.NewFindUserById(userRepo)
	findUserByEmail := application.NewFindUserByEmail(userRepo)
	updateUser := application.NewUpdateUser(userRepo, hasher)
	deleteUser := application.NewDeleteUser(userRepo)
	userHandler := http.NewUserHandler(*createUser, *findAllUsers, *findUserById, *findUserByEmail, *updateUser, *deleteUser)

	router.POST("/users", userHandler.CreateUser)
	router.GET("/users", userHandler.FindAllUsers)
	router.GET("/users/:id", userHandler.FindUserById)
	router.GET("/users/email/:email", userHandler.FindUserByEmail)
	router.PUT("/users/:id", userHandler.UpdateUser)
	router.DELETE("/users/:id", userHandler.DeleteUser)
}
