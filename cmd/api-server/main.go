package main

import (
	"log"

	"github.com/1shkumar/mini-orchestrator/internal/api"
	"github.com/1shkumar/mini-orchestrator/internal/db"
	"github.com/1shkumar/mini-orchestrator/internal/kafka"
	"github.com/gin-gonic/gin"
)

func main() {

	err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}

	kafka.InitProducer()

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "healthy",
		})
	})

	r.POST("/workflow", api.CreateWorkflow)

	r.POST(
		"/workflow/:id/start",
		api.StartWorkflow,
	)

	r.Run(":8080")
}
