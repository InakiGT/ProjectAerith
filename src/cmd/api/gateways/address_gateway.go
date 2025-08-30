package gateways

import (
	"rapi-pedidos/src/internal/address/application"
	"rapi-pedidos/src/internal/address/infrastructure/http"
	"rapi-pedidos/src/internal/address/infrastructure/persistence"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AddressBindRoutes(dbConn *gorm.DB, router *gin.Engine) {
	addressRepo := persistence.NewGormRepository(dbConn)
	create := application.NewCreateAddress(addressRepo)
	findAll := application.NewFindAllAddresses(addressRepo)
	findById := application.NewFindAddressById(addressRepo)
	update := application.NewUpdateAddress(addressRepo)
	delete := application.NewDeleteAddress(addressRepo)
	addressHandler := http.NewAddressHandler(*create, *findAll, *findById, *update, *delete)

	router.POST("/addresses", addressHandler.CreateAddress)
	router.GET("/addresses", addressHandler.FindAllAddresses)
	router.GET("/addresses/:id", addressHandler.FindAddressByID)
	router.PUT("/addresses/:id", addressHandler.UpdateAddress)
	router.DELETE("/addresses/:id", addressHandler.DeleteAddress)
}
