package models

type Workflow struct {
	Name  string `json:"name" binding:"required"`
	Tasks []Task `json:"tasks" binding:"required"`
}

type Task struct {
	Name      string   `json:"name" binding:"required"`
	DependsOn []string `json:"depends_on"`
}

type WorkflowResponse struct {
	ID      string `json:"id"`
	Message string `json:"message"`
}
