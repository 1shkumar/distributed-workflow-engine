package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/1shkumar/mini-orchestrator/internal/db"
	"github.com/1shkumar/mini-orchestrator/internal/kafka"
	"github.com/1shkumar/mini-orchestrator/internal/models"
)

func StartWorkflow(c *gin.Context) {

	workflowID := c.Param("id")

	workflow, err :=
		db.GetWorkflowByID(workflowID)

	if err != nil {
		c.JSON(
			http.StatusNotFound,
			gin.H{
				"error": "workflow not found",
			},
		)
		return
	}

	runID, err :=
		db.CreateWorkflowRun(workflowID)

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	for _, task := range workflow.Tasks {

		status := "PENDING"

		if len(task.DependsOn) == 0 {
			status = "QUEUED"
			event := models.TaskEvent{
				WorkflowRunID: runID,
				TaskName:      task.Name,
			}
			err :=
				kafka.PublishTask(event)
			if err != nil {
				c.JSON(
					http.StatusInternalServerError,
					gin.H{
						"error": err.Error(),
					},
				)
				return
			}
		}

		err :=
			db.CreateTaskRun(
				runID,
				task.Name,
				status,
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
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"workflow_run_id": runID,
			"status":          "RUNNING",
		},
	)
}
