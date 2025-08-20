package http

import (
	"net/http"

	"rapi-pedidos/src/internal/product/application"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	createProduct   application.CreateProductCommand
	findAllProducts application.FindAllProductsCommand
	findProductById application.FindProductByIdCommand
	updateProduct   application.UpdateProductCommand
	deleteProduct   application.DeleteProductCommand
}

func NewProductHandler(
	createProduct application.CreateProductCommand,
	findAllProducts application.FindAllProductsCommand,
	findProductById application.FindProductByIdCommand,
	updateProduct application.UpdateProductCommand,
	deleteProduct application.DeleteProductCommand,
) *ProductHandler {
	return &ProductHandler{
		createProduct:   createProduct,
		findAllProducts: findAllProducts,
		findProductById: findProductById,
		updateProduct:   updateProduct,
		deleteProduct:   deleteProduct,
	}
}

func (h *ProductHandler) CreateProduct(ctx *gin.Context) {
	var input struct {
		Name        string  `json:"name" binding:"required"`
		Description string  `json:"description" binding:"required"`
		Price       float32 `json:"price" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, err := h.createProduct.Execute(ctx, input.Name, input.Description, input.Price); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "product created"})
}

func (h *ProductHandler) FindAllProducts(ctx *gin.Context) {
	products, err := h.findAllProducts.Execute(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, products)
}

func (h *ProductHandler) FindProductById(ctx *gin.Context) {
	id := ctx.Param("id")
	product, err := h.findProductById.Execute(ctx, id)
	if err != nil {
		if err.Error() == "record not found" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, product)
}

func (h *ProductHandler) UpdateProduct(ctx *gin.Context) {
	id := ctx.Param("id")
	var input struct {
		Name        string  `json:"name"`
		Description string  `json:"description"`
		Price       float32 `json:"price"`
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := h.updateProduct.Execute(ctx, id, input.Name, input.Description, input.Price)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

func (h *ProductHandler) DeleteProduct(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := h.deleteProduct.Execute(ctx, id); err != nil {
		if err.Error() == "record not found" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
