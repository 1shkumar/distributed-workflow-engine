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

func IncrementTaskAttempt(
	workflowRunID string,
	taskName string,
) error {

	query := `
	UPDATE task_runs
	SET attempt = attempt + 1
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

func GetTaskAttempt(
	workflowRunID string,
	taskName string,
) (
	int,
	error,
) {

	query := `
	SELECT attempt
	FROM task_runs
	WHERE workflow_run_id = $1
	AND task_name = $2
	`

	var attempt int

	err := DB.QueryRow(
		context.Background(),
		query,
		workflowRunID,
		taskName,
	).Scan(&attempt)

	return attempt,
		err
}

func MarkTaskFailed(
	workflowRunID string,
	taskName string,
) error {

	query := `
	UPDATE task_runs
	SET status = 'FAILED'
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
