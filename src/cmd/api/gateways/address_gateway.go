package gateways

import (
	"rapi-pedidos/src/internal/address/application"
	"rapi-pedidos/src/internal/address/infraestructure/http"
	"rapi-pedidos/src/internal/address/infraestructure/persistence"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AddressBindRoutes(dbConn *gorm.DB, router *gin.Engine) {
	addressRepo := persistence.NewGormRepository(dbConn)
	createAddress := application.NewCreateAddress(addressRepo)
	findAllAddresses := application.NewFindAllAddresses(addressRepo)
	findAddresByID := application.NewFindAddressById(addressRepo)
	updateAddress := application.NewUpdateAddress(addressRepo)
	deleteAddress := application.NewDeleteAddress(addressRepo)
	addressHandler := http.NewAddressHandler(*createAddress, *findAllAddresses, *findAddresByID, *updateAddress, *deleteAddress)

	router.POST("addresses", addressHandler.CreateAddress)
	router.GET("addresses", addressHandler.FindAllAddresses)
	router.GET("addresses/:id", addressHandler.FindAddresByID)
	router.PUT("addresses/:id", addressHandler.UpdateAddress)
	router.DELETE("addresses/:id", addressHandler.DeleteAddress)
}
