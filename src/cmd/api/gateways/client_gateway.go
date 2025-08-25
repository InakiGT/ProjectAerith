package gateways

import (
	"rapi-pedidos/src/internal/client/application"
	"rapi-pedidos/src/internal/client/infrastructure/http"
	"rapi-pedidos/src/internal/client/infrastructure/persistence"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ClientBindRoutes(dbConn *gorm.DB, router *gin.Engine) {
	clientRepo := persistence.NewGormRepository(dbConn)
	createclient := application.NewCreateClient(clientRepo)
	findAllclientes := application.NewFindAllClients(clientRepo)
	findAddresByID := application.NewFindClientByID(clientRepo)
	updateclient := application.NewUpdateClient(clientRepo)
	deleteclient := application.NewDeleteClient(clientRepo)
	clientHandler := http.NewClientHandler(*createclient, *findAllclientes, *findAddresByID, *updateclient, *deleteclient)

	router.POST("/clients", clientHandler.CreateClient)
	router.GET("/clients", clientHandler.FindAllClients)
	router.GET("/clients/:id", clientHandler.FindClientByID)
	router.PUT("/clients/:id", clientHandler.UpdateClient)
	router.DELETE("/clients/:id", clientHandler.DeleteClient)
}
