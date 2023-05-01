package models

type Task struct {
	ID           int    `json:"id"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	StartAt      string `json:"start_at"`
	FinishAt     string `json:"finish_at"`
	StartedAt    string `json:"started_at"`
	FinishedAt   string `json:"finished_at"`
	Status       string `json:"status"`
	ProgrammerID string `json:"programmer_id"`
	ProjectID    int    `json:"project_id"`
	Attachment   string `json:"attachment"`
}

type Comment struct {
	ID           int    `json:"id"`
	Text         string `json:"text"`
	TaskID       int    `json:"task_id"`
	ProgrammerID string `json:"programmer_id"`
	CreatedAt    string `json:"created_at"`
}
