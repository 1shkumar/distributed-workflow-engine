package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/1shkumar/mini-orchestrator/internal/db"
	"github.com/1shkumar/mini-orchestrator/internal/models"
)

func CreateWorkflow(c *gin.Context) {

	var workflow models.Workflow

	err := c.ShouldBindJSON(&workflow)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	id, err := db.CreateWorkflow(workflow)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, models.WorkflowResponse{
		ID:      id,
		Message: "workflow created successfully",
	})
}
