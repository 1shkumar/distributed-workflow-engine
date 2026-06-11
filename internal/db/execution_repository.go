package db

import (
	"context"
	"encoding/json"

	"github.com/google/uuid"

	"github.com/1shkumar/mini-orchestrator/internal/models"
)

func GetWorkflowByID(
	workflowID string,
) (*models.Workflow, error) {

	query := `
	SELECT definition
	FROM workflows
	WHERE id = $1
	`

	var definition []byte

	err := DB.QueryRow(
		context.Background(),
		query,
		workflowID,
	).Scan(&definition)

	if err != nil {
		return nil, err
	}

	var workflow models.Workflow

	err = json.Unmarshal(
		definition,
		&workflow,
	)

	if err != nil {
		return nil, err
	}

	return &workflow, nil
}

func CreateWorkflowRun(
	workflowID string,
) (string, error) {

	runID := uuid.New().String()

	query := `
	INSERT INTO workflow_runs (
		id,
		workflow_id,
		status
	)
	VALUES ($1, $2, $3)
	`

	_, err := DB.Exec(
		context.Background(),
		query,
		runID,
		workflowID,
		"RUNNING",
	)

	if err != nil {
		return "", err
	}

	return runID, nil
}

func CreateTaskRun(
	workflowRunID string,
	taskName string,
	status string,
) error {

	taskRunID := uuid.New().String()

	query := `
	INSERT INTO task_runs (
		id,
		workflow_run_id,
		task_name,
		status
	)
	VALUES ($1, $2, $3, $4)
	`

	_, err := DB.Exec(
		context.Background(),
		query,
		taskRunID,
		workflowRunID,
		taskName,
		status,
	)

	return err
}
