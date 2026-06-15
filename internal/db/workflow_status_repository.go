package db

import (
	"context"

	"github.com/1shkumar/mini-orchestrator/internal/models"
)

func GetWorkflowStatus(
	workflowID string,
) (
	string,
	[]models.TaskStatus,
	error,
) {

	query := `
	SELECT wr.status,
		tr.task_name,
		tr.status
	FROM workflow_runs wr
	INNER JOIN task_runs tr
	ON wr.id = tr.workflow_run_id
	WHERE wr.workflow_id = $1
	AND wr.created_at = (
		SELECT MAX(created_at)
		FROM workflow_runs
		WHERE workflow_id = $1
	)
	`

	rows, err := DB.Query(
		context.Background(),
		query,
		workflowID,
	)

	if err != nil {
		return "", nil, err
	}

	defer rows.Close()

	var workflowStatus string
	var tasks []models.TaskStatus

	for rows.Next() {

		var task models.TaskStatus

		err := rows.Scan(
			&workflowStatus,
			&task.Name,
			&task.Status,
		)

		if err != nil {
			return "", nil, err
		}

		tasks = append(
			tasks,
			task,
		)
	}

	return workflowStatus,
		tasks,
		nil
}
