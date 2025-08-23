package gateways

import (
	"rapi-pedidos/src/internal/delivery_person/application"
	"rapi-pedidos/src/internal/delivery_person/infrastructure/http"
	"rapi-pedidos/src/internal/delivery_person/infrastructure/persistence"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DeliveryPersonBindRoutes(dbConn *gorm.DB, router *gin.Engine) {
	deliveryPersonRepo := persistence.NewGormRepository(dbConn)
	createDeliveryPerson := application.NewCreateDeliveryPerson(deliveryPersonRepo)
	findAllDeliveryPersons := application.NewFindAllDeliveryPersons(deliveryPersonRepo)
	findDeliveryPersonByID := application.NewFindDeliveryPersonByID(deliveryPersonRepo)
	updateDeliveryPerson := application.NewUpdateDeliveryPerson(deliveryPersonRepo)
	deleteDeliveryPerson := application.NewDeleteDeliveryPerson(deliveryPersonRepo)
	deliveryPersonHandler := http.NewDeliveryPersonHandler(*createDeliveryPerson, *findAllDeliveryPersons, *findDeliveryPersonByID, *updateDeliveryPerson, *deleteDeliveryPerson)

	router.POST("/delivery_persons", deliveryPersonHandler.CreateDeliveryPerson)
	router.GET("/delivery_persons", deliveryPersonHandler.FindAllDeliveryPersons)
	router.GET("/delivery_persons/:id", deliveryPersonHandler.FindDeliveryPersonByID)
	router.PUT("/delivery_persons/:id", deliveryPersonHandler.UpdateDeliveryPerson)
	router.DELETE("/delivery_persons/:id", deliveryPersonHandler.DeleteDeliveryPerson)
}
