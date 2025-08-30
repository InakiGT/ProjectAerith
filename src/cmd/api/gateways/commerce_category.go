package gateways

import (
	"rapi-pedidos/src/internal/commerce_category/application"
	"rapi-pedidos/src/internal/commerce_category/infrastructure/http"
	"rapi-pedidos/src/internal/commerce_category/infrastructure/persistence"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CommerceCategoryBindRoutes(dbConn *gorm.DB, router *gin.Engine) {
	categoryRepo := persistence.NewGormRepository(dbConn)
	create := application.NewCreateCommerceCategory(categoryRepo)
	findAll := application.NewFindAllCommerceCategories(categoryRepo)
	findByID := application.NewFindCommerceCategoryByID(categoryRepo)
	update := application.NewUpdateCommerceCategory(categoryRepo)
	delete := application.NewDeleteCommerceCategory(categoryRepo)
	handler := http.NewCommerceCategoryHandler(*create, *findAll, *findByID, *update, *delete)

	router.POST("/commerce_categories", handler.CreateCategory)
	router.GET("/commerce_categories", handler.FindAllCategories)
	router.GET("/commerce_categories/:id", handler.FindCategoryByID)
	router.PUT("/commerce_categories/:id", handler.UpdateCategory)
	router.DELETE("/commerce_categories/:id", handler.DeleteCategory)
}
