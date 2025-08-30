package http

import (
	"net/http"

	"rapi-pedidos/src/internal/commerce_category/application"

	"github.com/gin-gonic/gin"
)

type CommerceCategoryHandler struct {
	createCategory    application.CreateCommerceCategoryCommand
	findAllCategories application.FindAllCommerceCategoriesCommand
	findCategoryByID  application.FindCommerceCategoryByIDCommand
	updateCategory    application.UpdateCommerceCategoryCommand
	deleteCategory    application.DeleteCommerceCategoryCommand
}

func NewCommerceCategoryHandler(
	createCategory application.CreateCommerceCategoryCommand,
	findAllCategories application.FindAllCommerceCategoriesCommand,
	findAddressByID application.FindCommerceCategoryByIDCommand,
	updateCategory application.UpdateCommerceCategoryCommand,
	deleteCategory application.DeleteCommerceCategoryCommand,
) *CommerceCategoryHandler {
	return &CommerceCategoryHandler{
		createCategory:    createCategory,
		findAllCategories: findAllCategories,
		findCategoryByID:  findAddressByID,
		updateCategory:    updateCategory,
		deleteCategory:    deleteCategory,
	}
}

func (h *CommerceCategoryHandler) CreateCategory(ctx *gin.Context) {
	var input struct {
		Name string `json:"name" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, err := h.createCategory.Execute(ctx, input.Name); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "address created"})
}

func (h *CommerceCategoryHandler) FindAllCategories(ctx *gin.Context) {
	addresses, err := h.findAllCategories.Execute(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, addresses)
}

func (h *CommerceCategoryHandler) FindCategoryByID(ctx *gin.Context) {
	id := ctx.Param("id")
	address, err := h.findCategoryByID.Execute(ctx, id)
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

func (h *CommerceCategoryHandler) UpdateCategory(ctx *gin.Context) {
	id := ctx.Param("id")
	var input struct {
		Name string `json:"name"`
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := h.updateCategory.Execute(ctx, id, input.Name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

func (h *CommerceCategoryHandler) DeleteCategory(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := h.deleteCategory.Execute(ctx, id); err != nil {
		if err.Error() == "record not found" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
