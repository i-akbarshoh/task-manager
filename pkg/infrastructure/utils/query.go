package utils

var (
	Register = `
		INSERT INTO users (id, full_name, password, avatar, role, birth_date, phone, position) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`

	Login = `
		SELECT password FROM users WHERE phone = $1
	`

	GetProgrammers = `
		SELECT * FROM users WHERE position = 'programmer'
	`

	GetProgrammer = `
		SELECT * FROM users WHERE id = $1
	`

	UpdatePassword = `
		UPDATE users SET password = $1 WHERE phone = $2
	`

	DeleteUser = `
		DELETE FROM users WHERE id = $1
	`

	GetPassword = `
		SELECT password FROM users WHERE id = $1
	`

	CreateProject = `
		INSERT INTO projects (name, status, teamlead_id, attachment) VALUES ($1, $2, $3, $4)
	`

	GetProjects = `
		SELECT * FROM projects
	`

	UpdateProjectStatus = `
		UPDATE projects SET status = $1 WHERE id = $2
	`

	DeleteProject = `
		DELETE FROM projects WHERE id = $1
	`

	GetProject = `
		SELECT * FROM projects WHERE id = $1
	`

	AddAttendance = `
		INSERT INTO attendance (type, user_id, date) VALUES ($1, $2, $3)
	`

	GetUserAttendance = `
		SELECT * FROM attendance WHERE user_id = $1 AND type = $2
	`

	CreateTask = `
		INSERT INTO tasks (title, description, start_at, finish_at, project_id, programmer_id, attachment) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id
	`

	UpdateTaskStatus = `
		UPDATE tasks SET status = $1 WHERE id = $2
	`

	GetTasks = `
		SELECT * FROM tasks WHERE project_id = $1 AND programmer_id = $2
	`

	GetTask = `
		SELECT * FROM tasks WHERE id = $1
	`

	DeleteTask = `
		DELETE FROM tasks WHERE id = $1
	`

	StartTask = `
		UPDATE tasks SET started_at = $1 WHERE id = $2
	`

	FinishTask = `
		UPDATE tasks SET finished_at = $1 WHERE id = $2
	`

	WriteComment = `
		INSERT INTO comments (task_id, programmer_id, text) VALUES ($1, $2, $3) RETURNING id
	`

	GetComments = `
		SELECT * FROM comments WHERE task_id = $1
	`

	DeleteComment = `
		DELETE FROM comments WHERE id = $1
	`
)
