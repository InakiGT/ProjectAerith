package http

import (
	"net/http"

	"rapi-pedidos/src/internal/client_card/application"

	"github.com/gin-gonic/gin"
)

type ClientCardHanlder struct {
	createClientCard          application.CreateClientCardCommand
	findAllClientCards        application.FindAllClientCardsCommand
	findClientCardByID        application.FindClientCardByIDCommand
	findClientCardsByClientID application.FindClientCardsByClientIDCommand
	updateClientCard          application.UpdateClientCardCommand
	deleteClientCard          application.DeleteClientCardCommand
}

func NewClientCardHanlder(
	createClientCard application.CreateClientCardCommand,
	findAllClientCards application.FindAllClientCardsCommand,
	findClientCardByID application.FindClientCardByIDCommand,
	findClientCardsByClientID application.FindClientCardsByClientIDCommand,
	updateClientCard application.UpdateClientCardCommand,
	deleteClientCard application.DeleteClientCardCommand,
) *ClientCardHanlder {
	return &ClientCardHanlder{
		createClientCard:          createClientCard,
		findAllClientCards:        findAllClientCards,
		findClientCardByID:        findClientCardByID,
		findClientCardsByClientID: findClientCardsByClientID,
		updateClientCard:          updateClientCard,
		deleteClientCard:          deleteClientCard,
	}
}

func (h *ClientCardHanlder) CreateClientCard(ctx *gin.Context) {
	var input struct {
		UserId   uint   `json:"user_id" binding:"required"`
		Provider string `json:"provider" binding:"required"`
		ExpYear  string `json:"exp_year" binding:"required"`
		ExpMonth string `json:"exp_month" binding:"required"`
		Last4    string `json:"last_4" binding:"required"`
		Brand    string `json:"brand" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: Servicio manejador de tarjetas
	if _, err := h.createClientCard.Execute(ctx, input.UserId, input.Provider, input.ExpYear, input.ExpMonth, input.Last4, input.Brand, ""); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "client card created"})
}

func (h *ClientCardHanlder) FindAllClientCards(ctx *gin.Context) {
	clientCards, err := h.findAllClientCards.Execute(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, clientCards)
}

func (h *ClientCardHanlder) FindClientCardById(ctx *gin.Context) {
	id := ctx.Param("id")
	clientCard, err := h.findClientCardByID.Execute(ctx, id)
	if err != nil {
		if err.Error() == "record not found" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, clientCard)
}

func (h *ClientCardHanlder) FindClientCardsByClientId(ctx *gin.Context) {
	id := ctx.Param("id")
	clientCard, err := h.findClientCardsByClientID.Execute(ctx, id)
	if err != nil {
		if err.Error() == "record not found" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, clientCard)
}

func (h *ClientCardHanlder) UpdateClientCard(ctx *gin.Context) {
	id := ctx.Param("id")
	var input struct {
		UserId   uint   `json:"user_id"`
		Provider string `json:"provider"`
		ExpYear  string `json:"exp_year"`
		ExpMonth string `json:"exp_month"`
		Last4    string `json:"last_4"`
		Brand    string `json:"brand"`
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := h.updateClientCard.Execute(ctx, id, input.Provider, input.ExpYear, input.ExpMonth, input.Last4, input.Brand, "")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

func (h *ClientCardHanlder) DeleteClietCard(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := h.deleteClientCard.Execute(ctx, id); err != nil {
		if err.Error() == "record not found" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
