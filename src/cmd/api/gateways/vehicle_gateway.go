package gateways

import (
	"rapi-pedidos/src/internal/vehicle/application"
	"rapi-pedidos/src/internal/vehicle/infrastructure/http"
	"rapi-pedidos/src/internal/vehicle/infrastructure/persistence"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func VehicleBindRoutes(dbConn *gorm.DB, router *gin.Engine) {
	vehicleRepo := persistence.NewGormRepository(dbConn)
	createVehicle := application.NewCreateVehicle(vehicleRepo)
	updateVehicle := application.NewUpdateVehicle(vehicleRepo)
	deleteVehicle := application.NewDeleteVehicle(vehicleRepo)
	findAllVehicles := application.NewFindAllVehicles(vehicleRepo)
	findVehicleByID := application.NewFindVehicleByID(vehicleRepo)
	// application.NewFindVehicleByCardID(vehicleRepo)
	// application.NewFindVehicleByPlate(vehicleRepo)
	vehicleHandler := http.NewVehicleHandler(*createVehicle, *findAllVehicles, *findVehicleByID, *updateVehicle, *deleteVehicle)

	router.POST("/vehicles", vehicleHandler.CreateVehicle)
	router.GET("/vehicles", vehicleHandler.FindAllVehicles)
	router.GET("/vehicles/:id", vehicleHandler.FindVehicleByID)
	router.PUT("/vehicles/:id", vehicleHandler.UpdateVehicle)
	router.DELETE("/vehicles/:id", vehicleHandler.DeleteVehicle)
}
