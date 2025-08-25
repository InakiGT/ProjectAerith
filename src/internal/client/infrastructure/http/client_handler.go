package http

import (
	"net/http"

	"rapi-pedidos/src/internal/client/application"

	"github.com/gin-gonic/gin"
)

type ClientHandler struct {
	createClient   application.CreateClientCommand
	findAllClients application.FindAllClientsCommand
	findClientByID application.FindClientByIDCommand
	updateClient   application.UpdateClientCommand
	deleteClient   application.DeleteClientCommand
}

func NewClientHandler(
	createClient application.CreateClientCommand,
	findAllClients application.FindAllClientsCommand,
	findClientByID application.FindClientByIDCommand,
	updateClient application.UpdateClientCommand,
	deleteClient application.DeleteClientCommand,
) *ClientHandler {
	return &ClientHandler{
		createClient:   createClient,
		findAllClients: findAllClients,
		findClientByID: findClientByID,
		updateClient:   updateClient,
		deleteClient:   deleteClient,
	}
}

func (h *ClientHandler) CreateClient(ctx *gin.Context) {
	var input struct {
		UserID        uint `json:"street" binding:"required"`
		MainAddressID uint `json:"city" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, err := h.createClient.Execute(ctx, input.UserID, input.MainAddressID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "address created"})
}

func (h *ClientHandler) FindAllClients(ctx *gin.Context) {
	addresses, err := h.findAllClients.Execute(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, addresses)
}

func (h *ClientHandler) FindClientByID(ctx *gin.Context) {
	id := ctx.Param("id")
	address, err := h.findClientByID.Execute(ctx, id)
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

func (h *ClientHandler) UpdateClient(ctx *gin.Context) {
	id := ctx.Param("id")
	var input struct {
		MainAddressID uint `json:"city" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := h.updateClient.Execute(ctx, id, input.MainAddressID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

func (h *ClientHandler) DeleteClient(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := h.deleteClient.Execute(ctx, id); err != nil {
		if err.Error() == "record not found" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
