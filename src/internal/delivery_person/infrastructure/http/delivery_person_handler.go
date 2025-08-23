package http

import (
	"net/http"
	"time"

	"rapi-pedidos/src/internal/delivery_person/application"

	"github.com/gin-gonic/gin"
)

type DeliveryPersonHandler struct {
	createDeliveryPerson   application.CreateDeliveryPersonCommand
	findAllDeliveryPersons application.FindAllDeliveryPersonCommand
	findDeliveryPersonByID application.FindDeliveryPersonByIDCommand
	updateDeliveryPerson   application.UpdateDeliveryPersonCommand
	deleteDeliveryPerson   application.DeleteDeliveryPersonCommand
}

func NewDeliveryPersonHandler(
	createDeliveryPerson application.CreateDeliveryPersonCommand,
	findAllDeliveryPersons application.FindAllDeliveryPersonCommand,
	findDeliveryPersonByID application.FindDeliveryPersonByIDCommand,
	updateDeliveryPerson application.UpdateDeliveryPersonCommand,
	deleteDeliveryPerson application.DeleteDeliveryPersonCommand,
) *DeliveryPersonHandler {
	return &DeliveryPersonHandler{
		createDeliveryPerson:   createDeliveryPerson,
		findAllDeliveryPersons: findAllDeliveryPersons,
		findDeliveryPersonByID: findDeliveryPersonByID,
		updateDeliveryPerson:   updateDeliveryPerson,
		deleteDeliveryPerson:   deleteDeliveryPerson,
	}
}

func (h *DeliveryPersonHandler) CreateDeliveryPerson(ctx *gin.Context) {
	var input struct {
		UserId     uint      `json:"user_id"`
		PersonalID string    `json:"personal_id"`
		Birthday   time.Time `json:"birthday"`
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, err := h.createDeliveryPerson.Execute(ctx, input.UserId, input.Birthday, input.PersonalID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "delivery person created"})
}

func (h *DeliveryPersonHandler) FindAllDeliveryPersons(ctx *gin.Context) {
	addresses, err := h.findAllDeliveryPersons.Execute(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, addresses)
}

func (h *DeliveryPersonHandler) FindDeliveryPersonByID(ctx *gin.Context) {
	id := ctx.Param("id")
	address, err := h.findDeliveryPersonByID.Execute(ctx, id)
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

func (h *DeliveryPersonHandler) UpdateDeliveryPerson(ctx *gin.Context) {
	id := ctx.Param("id")
	var input struct {
		PersonalId string    `json:"personal_id"`
		Birthday   time.Time `json:"birthday"`
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, err := h.updateDeliveryPerson.Execute(ctx, id, input.PersonalId, input.Birthday); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

func (h *DeliveryPersonHandler) DeleteDeliveryPerson(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := h.deleteDeliveryPerson.Execute(ctx, id); err != nil {
		if err.Error() == "record not found" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
