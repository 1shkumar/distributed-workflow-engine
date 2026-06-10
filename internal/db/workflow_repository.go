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
