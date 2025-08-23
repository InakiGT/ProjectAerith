package http

import (
	"net/http"

	"rapi-pedidos/src/internal/address/application"

	"github.com/gin-gonic/gin"
)

type AddresHandler struct {
	createAddress    application.CreateAddressCommand
	findAllAddresses application.FindAllAddressesCommand
	findAddressByID  application.FindAddresByIdCommand
	updateAddress    application.UpadteAddressCommand
	deleteAddress    application.DeleteAddressCommand
}

func NewAddressHandler(
	createAddress application.CreateAddressCommand,
	findAllAddresses application.FindAllAddressesCommand,
	findAddressByID application.FindAddresByIdCommand,
	updateAddress application.UpadteAddressCommand,
	deleteAddress application.DeleteAddressCommand,
) *AddresHandler {
	return &AddresHandler{
		createAddress:    createAddress,
		findAllAddresses: findAllAddresses,
		findAddressByID:  findAddressByID,
		updateAddress:    updateAddress,
		deleteAddress:    deleteAddress,
	}
}

func (h *AddresHandler) CreateAddress(ctx *gin.Context) {
	var input struct {
		Street     string `json:"street" binding:"required"`
		City       string `json:"city" binding:"required"`
		Country    string `json:"country" binding:"required"`
		Number     string `json:"number" binding:"required"`
		PostalCode string `json:"postal_code" binding:"required"`
		Cologne    string `json:"cologne" binding:"required"`
		UserID     uint   `json:"user_id" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, err := h.createAddress.Execute(ctx, input.City, input.Country, input.Number, input.Street, input.PostalCode, input.Cologne, input.UserID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "address created"})
}

func (h *AddresHandler) FindAllAddresses(ctx *gin.Context) {
	addresses, err := h.findAllAddresses.Execute(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, addresses)
}

func (h *AddresHandler) FindAddresByID(ctx *gin.Context) {
	id := ctx.Param("id")
	address, err := h.findAddressByID.Execute(ctx, id)
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

func (h *AddresHandler) UpdateAddress(ctx *gin.Context) {
	id := ctx.Param("id")
	var input struct {
		Street     string `json:"street"`
		City       string `json:"city"`
		Country    string `json:"country"`
		Number     string `json:"number"`
		PostalCode string `json:"postal_code"`
		Cologne    string `json:"cologne"`
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := h.updateAddress.Execute(ctx, id, input.City, input.Country, input.Number, input.Street, input.PostalCode, input.Cologne)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

func (h *AddresHandler) DeleteAddress(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := h.deleteAddress.Execute(ctx, id); err != nil {
		if err.Error() == "record not found" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
