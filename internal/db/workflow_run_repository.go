package db

import "context"

func AreAllTasksCompleted(
	workflowRunID string,
) bool {

	query := `
	SELECT COUNT(*)
	FROM task_runs
	WHERE workflow_run_id = $1
	AND status != 'SUCCESS'
	`

	var remaining int

	err := DB.QueryRow(
		context.Background(),
		query,
		workflowRunID,
	).Scan(&remaining)

	if err != nil {
		return false
	}

	return remaining == 0
}

func MarkWorkflowCompleted(
	workflowRunID string,
) error {

	query := `
	UPDATE workflow_runs
	SET status = 'COMPLETED'
	WHERE id = $1
	`

	_, err := DB.Exec(
		context.Background(),
		query,
		workflowRunID,
	)

	return err
}

func MarkWorkflowFailed(
	workflowRunID string,
) error {

	query := `
	UPDATE workflow_runs
	SET status = 'FAILED'
	WHERE id = $1
	`

	_, err := DB.Exec(
		context.Background(),
		query,
		workflowRunID,
	)

	return err
}
