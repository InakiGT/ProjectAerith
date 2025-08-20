package http

import (
	"net/http"

	"rapi-pedidos/src/internal/vehicle/application"

	"github.com/gin-gonic/gin"
)

type VehicleHandler struct {
	createVehicle   application.CreateVehicleCommand
	findAllVehicles application.FindAllVehiclesCommand
	findVehicleByID application.FindVehicleByIDCommand
	updateVehicle   application.UpdateVehicleCommand
	deleteVehicle   application.DeleteVehicleCommand
}

func NewVehicleHandler(
	createVehicle application.CreateVehicleCommand,
	findAllVehicles application.FindAllVehiclesCommand,
	findVehicleByID application.FindVehicleByIDCommand,
	updateVehicle application.UpdateVehicleCommand,
	deleteVehicle application.DeleteVehicleCommand,
) *VehicleHandler {
	return &VehicleHandler{
		createVehicle:   createVehicle,
		findAllVehicles: findAllVehicles,
		findVehicleByID: findVehicleByID,
		updateVehicle:   updateVehicle,
		deleteVehicle:   deleteVehicle,
	}
}

func (h *VehicleHandler) CreateVehicle(ctx *gin.Context) {
	var input struct {
		Color            string `json:"color"`
		Plate            string `json:"plate"`
		CardID           string `json:"card_id"`
		Type             string `json:"type"`
		DeliveryPersonID uint   `json:"delivery_person_id"`
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, err := h.createVehicle.Execute(ctx, input.Color, input.Type, input.Plate, input.CardID, input.DeliveryPersonID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "vehicle created"})
}

func (h *VehicleHandler) FindAllVehicles(ctx *gin.Context) {
	addresses, err := h.findAllVehicles.Execute(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, addresses)
}

func (h *VehicleHandler) FindVehicleByID(ctx *gin.Context) {
	id := ctx.Param("id")
	address, err := h.findVehicleByID.Execute(ctx, id)
	if err != nil {
		if err.Error() == "record not found" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, address)
}

func (h *VehicleHandler) UpdateVehicle(ctx *gin.Context) {
	id := ctx.Param("id")
	var input struct {
		Color            string `json:"color"`
		Plate            string `json:"plate"`
		CardID           string `json:"card_id"`
		Type             string `json:"type"`
		DeliveryPersonID uint   `json:"delivery_person_id"`
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.updateVehicle.Execute(ctx, id, input.Color, input.Type, input.Plate, input.CardID, input.DeliveryPersonID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

func (h *VehicleHandler) DeleteVehicle(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := h.deleteVehicle.Execute(ctx, id); err != nil {
		if err.Error() == "record not found" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
