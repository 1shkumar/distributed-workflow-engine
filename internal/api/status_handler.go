package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/1shkumar/mini-orchestrator/internal/db"
	"github.com/1shkumar/mini-orchestrator/internal/models"
)

func GetWorkflowStatus(
	c *gin.Context,
) {

	workflowID :=
		c.Param("id")

	status, tasks, err :=
		db.GetWorkflowStatus(
			workflowID,
		)

	if err != nil {

		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error": err.Error(),
			},
		)

		return
	}

	response :=
		models.WorkflowStatusResponse{
			WorkflowStatus: status,
			Tasks:          tasks,
		}

	c.JSON(
		http.StatusOK,
		response,
	)
}
