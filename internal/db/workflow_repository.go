package db

import (
	"context"
	"encoding/json"

	"github.com/google/uuid"

	"github.com/1shkumar/mini-orchestrator/internal/models"
)

func CreateWorkflow(
	workflow models.Workflow,
) (string, error) {

	id := uuid.New().String()

	definition, err := json.Marshal(workflow)
	if err != nil {
		return "", err
	}

	query := `
	INSERT INTO workflows (
		id,
		name,
		definition
	)
	VALUES ($1, $2, $3)
	`

	_, err = DB.Exec(
		context.Background(),
		query,
		id,
		workflow.Name,
		definition,
	)

	if err != nil {
		return "", err
	}

	return id, nil
}

func GetWorkflowDefinitionByRunID(
	workflowRunID string,
) (
	models.Workflow,
	error,
) {

	query := `
	SELECT w.definition
	FROM workflows w
	INNER JOIN workflow_runs wr
	ON wr.workflow_id = w.id
	WHERE wr.id = $1
	`

	var rawDefinition []byte

	err := DB.QueryRow(
		context.Background(),
		query,
		workflowRunID,
	).Scan(&rawDefinition)

	if err != nil {
		return models.Workflow{}, err
	}

	var workflow models.Workflow

	err = json.Unmarshal(
		rawDefinition,
		&workflow,
	)

	if err != nil {
		return models.Workflow{}, err
	}

	return workflow, nil
}
