package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// Connection to the database
	// dbConn := db.NewPostgresConnection()

	// Router
	router := gin.Default()
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "OK",
		})
	})

	router.Run()
}
