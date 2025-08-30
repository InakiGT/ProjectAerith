package http

import (
	"net/http"
	"time"

	"rapi-pedidos/src/internal/commerce/application"

	"github.com/gin-gonic/gin"
)

type CommerceHandler struct {
	createCommerce   application.CreateCommercesCommand
	findAllComerces  application.FindAllCommercesCommand
	findCommerceByID application.FindCommerceByIDCommand
	updateCommerce   application.UpdateCommercesCommand
	deleteCommerce   application.DeleteCommerceCommand
}

func NewCommerceHandler(
	createCommerce application.CreateCommercesCommand,
	findAllComerces application.FindAllCommercesCommand,
	findCommerceByID application.FindCommerceByIDCommand,
	updateCommerce application.UpdateCommercesCommand,
	deleteCommerce application.DeleteCommerceCommand,
) *CommerceHandler {
	return &CommerceHandler{
		createCommerce:   createCommerce,
		findAllComerces:  findAllComerces,
		findCommerceByID: findCommerceByID,
		updateCommerce:   updateCommerce,
		deleteCommerce:   deleteCommerce,
	}
}

func (h *CommerceHandler) CreateCommerce(ctx *gin.Context) {
	var input struct {
		MainAddressId      uint      `json:"address_id" binding:"required"`
		CommerceCategoryId uint      `json:"category_id" binding:"required"`
		Banner             string    `json:"banner" binding:"required"`
		Status             string    `json:"status" binding:"required"`
		OpenTime           time.Time `json:"open_time" binding:"required"`
		CloseTime          time.Time `json:"close_time" binding:"required"`
		BaseCommision      float32   `json:"commision"`
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, err := h.createCommerce.Execute(ctx, input.MainAddressId, input.CommerceCategoryId, input.Banner, input.Status, input.OpenTime, input.CloseTime, input.BaseCommision); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "address created"})
}

func (h *CommerceHandler) FindAllComerces(ctx *gin.Context) {
	addresses, err := h.findAllComerces.Execute(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, addresses)
}

func (h *CommerceHandler) FindCommerceByID(ctx *gin.Context) {
	id := ctx.Param("id")
	address, err := h.findCommerceByID.Execute(ctx, id)
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

func (h *CommerceHandler) UpdateCommerce(ctx *gin.Context) {
	id := ctx.Param("id")
	var input struct {
		MainAddressId      uint      `json:"address_id"`
		CommerceCategoryId uint      `json:"category_id"`
		Banner             string    `json:"banner"`
		Status             string    `json:"status"`
		OpenTime           time.Time `json:"open_time"`
		CloseTime          time.Time `json:"close_time"`
		BaseCommision      float32   `json:"commision"`
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := h.updateCommerce.Execute(ctx, id, input.MainAddressId, input.CommerceCategoryId, input.Banner, input.Status, input.OpenTime, input.CloseTime, input.BaseCommision)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

func (h *CommerceHandler) DeleteCommerce(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := h.deleteCommerce.Execute(ctx, id); err != nil {
		if err.Error() == "record not found" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
