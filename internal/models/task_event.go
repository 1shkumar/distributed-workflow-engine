package models

type TaskEvent struct {
	WorkflowRunID string `json:"workflow_run_id"`
	TaskName      string `json:"task_name"`
}
