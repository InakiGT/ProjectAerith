package gateways

import (
	"rapi-pedidos/src/internal/client_card/application"
	"rapi-pedidos/src/internal/client_card/infrastructure/http"
	"rapi-pedidos/src/internal/client_card/infrastructure/persistence"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ClientCardBindRoutes(dbConn *gorm.DB, router *gin.Engine) {
	clientCardRepo := persistence.NewGormRepository(dbConn)
	createClientCard := application.NewCreateClientCard(clientCardRepo)
	findAllClientCards := application.NewFindAllClientCards(clientCardRepo)
	findClientCardByID := application.NewFindClientCardByID(clientCardRepo)
	findClientCardsByClientID := application.NewFindClientCardsByClientID(clientCardRepo)
	updateClientCard := application.NewUpdateClientCard(clientCardRepo)
	deleteClientCard := application.NewDeleteClientCard(clientCardRepo)
	clientCardHandler := http.NewClientCardHanlder(*createClientCard, *findAllClientCards, *findClientCardByID, *findClientCardsByClientID, *updateClientCard, *deleteClientCard)

	router.POST("/client_cards", clientCardHandler.CreateClientCard)
	router.GET("/client_cards", clientCardHandler.FindAllClientCards)
	router.GET("/client_cards/:id", clientCardHandler.FindClientCardById)
	router.GET("/client_cards/user/:id", clientCardHandler.FindClientCardsByClientId)
	router.PUT("/client_cards/:id", clientCardHandler.UpdateClientCard)
	router.DELETE("/client_cards/:id", clientCardHandler.DeleteClietCard)
}
