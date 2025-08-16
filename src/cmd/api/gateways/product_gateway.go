package gateways

import (
	"rapi-pedidos/src/internal/product/application"
	"rapi-pedidos/src/internal/product/infraestructure/http"
	"rapi-pedidos/src/internal/product/infraestructure/persistence"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ProductBindRoutes(dbConn *gorm.DB, router *gin.Engine) {
	productRepo := persistence.NewGormRepository(dbConn)
	createProduct := application.NewCreateProduct(productRepo)
	findAllProducts := application.NewFindAllProducts(productRepo)
	findProductById := application.NewFindProductById(productRepo)
	updateProduct := application.NewUpdateProduct(productRepo)
	deleteProduct := application.NewDeleteProduct(productRepo)
	productHandler := http.NewProductHandler(*createProduct, *findAllProducts, *findProductById, *updateProduct, *deleteProduct)

	router.POST("/products", productHandler.CreateProduct)
	router.GET("/products", productHandler.FindAllProducts)
	router.GET("/products/:id", productHandler.FindProductById)
	router.PUT("/products/:id", productHandler.UpdateProduct)
	router.DELETE("/products/:id", productHandler.DeleteProduct)
}
