package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/1shkumar/mini-orchestrator/internal/api"
	"github.com/1shkumar/mini-orchestrator/internal/db"
)

func main() {

	err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "healthy",
		})
	})

	r.POST("/workflow", api.CreateWorkflow)

	r.Run(":8080")
}
