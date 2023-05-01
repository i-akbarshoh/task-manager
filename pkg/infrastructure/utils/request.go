package utils

type LoginRequestModel struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type UpdatePasswordRequestModel struct {
	Phone       string `json:"phone"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

type DeleteUserRequestModel struct {
	ID       string `json:"id"`
	Password string `json:"password"`
}

type UpdateProjectStatusRequestModel struct {
	ID     int    `json:"id"`
	Status string `json:"status"`
}

type DeleteProjectRequestModel struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type AddAttendanceRequestModel struct {
	Type      string `json:"type"`
	UserID    string `json:"user_id"`
	CreatedAt string `json:"created_at"`
}

type GetUserAttendanceRequestModel struct {
	UserID string `json:"user_id"`
	Type   string `json:"type"`
}

type CreateTaskRequestModel struct {
	Title        string `json:"title"`
	Description  string `json:"description"`
	StartAt      string `json:"start_at"`
	FinishAt     string `json:"finish_at"`
	ProjectID    int    `json:"project_id"`
	ProgrammerID string `json:"programmer_id"`
	Attachment   string `json:"attachment"`
}

type UpdateTaskStatusRequestModel struct {
	ID     int    `json:"id"`
	Status string `json:"status"`
}

type GetTasksRequestModel struct {
	ProjectID    int    `json:"project_id"`
	ProgrammerID string `json:"programmer_id"`
}

type DeleteTaskRequestModel struct {
	ID       int    `json:"id"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type StartOrFinishTaskRequestModel struct {
	ID   int    `json:"id"`
	Date string `json:"date"`
	Type string `json:"type"`
}

type WriteCommentRequestModel struct {
	Text         string `json:"text"`
	ProgrammerID string `json:"programmer_id"`
	TaskID       int    `json:"task_id"`
}

type GetCommentsRequestModel struct {
	TaskID int `json:"task_id"`
}

type DeleteCommentRequestModel struct {
	ID       int    `json:"id"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}
