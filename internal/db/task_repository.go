package db

import (
	"context"
)

func MarkTaskSuccess(
	workflowRunID string,
	taskName string,
) error {

	query := `
	UPDATE task_runs
	SET
		status = 'SUCCESS',
		completed_at = NOW()
	WHERE workflow_run_id = $1
	AND task_name = $2
	`

	_, err := DB.Exec(
		context.Background(),
		query,
		workflowRunID,
		taskName,
	)

	return err
}

func IsTaskCompleted(
	workflowRunID string,
	taskName string,
) bool {

	query := `
	SELECT COUNT(*)
	FROM task_runs
	WHERE workflow_run_id = $1
	AND task_name = $2
	AND status = 'SUCCESS'
	`

	var count int

	err := DB.QueryRow(
		context.Background(),
		query,
		workflowRunID,
		taskName,
	).Scan(&count)

	if err != nil {
		return false
	}

	return count > 0
}

func MarkTaskQueued(
	workflowRunID string,
	taskName string,
) error {

	query := `
	UPDATE task_runs
	SET status = 'QUEUED'
	WHERE workflow_run_id = $1
	AND task_name = $2
	`

	_, err := DB.Exec(
		context.Background(),
		query,
		workflowRunID,
		taskName,
	)

	return err
}
