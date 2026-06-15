package models

type TaskStatus struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}

type WorkflowStatusResponse struct {
	WorkflowStatus string       `json:"workflow_status"`
	Tasks          []TaskStatus `json:"tasks"`
}
