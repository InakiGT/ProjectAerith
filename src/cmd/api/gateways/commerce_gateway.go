package gateways

import (
	"rapi-pedidos/src/internal/commerce/application"
	"rapi-pedidos/src/internal/commerce/infrastructure/http"
	"rapi-pedidos/src/internal/commerce/infrastructure/persistence"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CommerceBindRoutes(dbConn *gorm.DB, router *gin.Engine) {
	comerceRepo := persistence.NewGormRepository(dbConn)
	create := application.NewCreateCommerce(comerceRepo)
	findAll := application.NewFindAllCommerces(comerceRepo)
	findByID := application.NewFindCommerceByID(comerceRepo)
	update := application.NewUpdateCommerce(comerceRepo)
	delete := application.NewDeleteCommerce(comerceRepo)
	handler := http.NewCommerceHandler(*create, *findAll, *findByID, *update, *delete)

	router.POST("/commerces", handler.CreateCommerce)
	router.GET("/commerces", handler.FindAllComerces)
	router.GET("/commerces/:id", handler.FindCommerceByID)
	router.PUT("/commerces/:id", handler.UpdateCommerce)
	router.DELETE("/commerces/:id", handler.DeleteCommerce)
}
