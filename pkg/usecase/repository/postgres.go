package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/i-akbarshoh/task-manager/pkg/domain/models"
	"github.com/i-akbarshoh/task-manager/pkg/infrastructure/utils"
	"golang.org/x/crypto/bcrypt"
)

type postgres struct {
	db *sql.DB
}

func NewPostgres(db *sql.DB) *postgres {
	return &postgres{db: db}
}

func (p *postgres) Register(user models.User) error {
	_, err := p.db.Exec(utils.Register, user.ID, user.FullName, user.Password, user.Avatar, user.Role, user.BirthDate, user.Phone, user.Position)
	if err != nil {
		return err
	}
	return nil
}

func (p *postgres) Login(r utils.LoginRequestModel) error {
	var password string
	if err := p.db.QueryRow(utils.Login, r.Phone).Scan(&password); err != nil {
		return err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(password), []byte(r.Password)); err != nil {
		return err
	}

	return nil
}

func (p *postgres) GetProgrammers() ([]models.User, error) {
	rows, err := p.db.Query(utils.GetProgrammers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var programmers []models.User
	for rows.Next() {
		var programmer models.User
		if err := rows.Scan(&programmer.ID, &programmer.FullName, &programmer.Avatar, &programmer.Role, &programmer.BirthDate, &programmer.Phone, &programmer.Position, &programmer.Password); err != nil {
			return nil, err
		}
		programmers = append(programmers, programmer)
	}

	return programmers, nil
}

func (p *postgres) GetProgrammer(id string) (models.User, error) {
	var programmer models.User
	if err := p.db.QueryRow(utils.GetProgrammer, id).Scan(&programmer.ID, &programmer.FullName, &programmer.Avatar, &programmer.Role, &programmer.BirthDate, &programmer.Phone, &programmer.Position, &programmer.Password); err != nil {
		return programmer, err
	}

	return programmer, nil
}

func (p *postgres) UpdatePassword(r utils.UpdatePasswordRequestModel) error {
	if err := p.Login(utils.LoginRequestModel{Phone: r.Phone, Password: r.OldPassword}); err != nil {
		return err
	}
	password, err := utils.GeneratePasswordHash(r.NewPassword)
	if err != nil {
		return err
	}
	_, err = p.db.Exec(utils.UpdatePassword, password, r.Phone)
	if err != nil {
		return err
	}
	return nil
}

func (p *postgres) DeleteUser(r utils.DeleteUserRequestModel) error {
	var (
		password string
	)

	if err := p.db.QueryRow(utils.GetPassword, r.ID).Scan(&password); err != nil {
		return err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(password), []byte(r.Password)); err != nil {
		return errors.Join(err, fmt.Errorf("password is not correct"))
	}
	_, err := p.db.Exec(utils.DeleteUser, r.ID)
	if err != nil {
		return err
	}
	return nil
}

func (p *postgres) CreateProject(project models.Project) error {
	_, err := p.db.Exec(utils.CreateProject, project.Name, project.Status, project.TeamLeadID, project.Attachment)
	if err != nil {
		return err
	}
	return nil
}

func (p *postgres) GetProjects() ([]models.Project, error) {
	rows, err := p.db.Query(utils.GetProjects)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var projects []models.Project
	for rows.Next() {
		var (
			project      models.Project
			finishedDate sql.NullTime
		)
		if err := rows.Scan(&project.ID, &project.Name, &project.StartedDate, &finishedDate, &project.Status, &project.TeamLeadID, &project.Attachment); err != nil {
			return nil, err
		}
		project.FinishedDate = finishedDate.Time.String()
		projects = append(projects, project)
	}

	return projects, nil
}

func (p *postgres) UpdateProjectStatus(r utils.UpdateProjectStatusRequestModel) error {
	_, err := p.db.Exec(utils.UpdateProjectStatus, r.Status, r.ID)
	if err != nil {
		return err
	}
	return nil
}

func (p *postgres) DeleteProject(id int) error {
	_, err := p.db.Exec(utils.DeleteProject, id)
	if err != nil {
		return err
	}
	return nil
}

func (p *postgres) GetProject(id int) (models.Project, error) {
	var (
		project      models.Project
		finishedDate sql.NullTime
	)
	if err := p.db.QueryRow(utils.GetProject, id).Scan(&project.ID, &project.Name, &project.StartedDate, &finishedDate, &project.Status, &project.TeamLeadID, &project.Attachment); err != nil {
		return project, err
	}
	project.FinishedDate = finishedDate.Time.String()
	return project, nil
}

func (p *postgres) AddAttendance(r utils.AddAttendanceRequestModel) error {
	_, err := p.db.Exec(utils.AddAttendance, r.Type, r.UserID, r.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (p *postgres) GetUserAttendances(r utils.GetUserAttendanceRequestModel) ([]models.Attendance, error) {
	rows, err := p.db.Query(utils.GetUserAttendance, r.UserID, r.Type)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var attendances []models.Attendance
	for rows.Next() {
		var attendance models.Attendance
		if err := rows.Scan(&attendance.Type, &attendance.UserID, &attendance.CreatedAt); err != nil {
			return nil, err
		}
		attendances = append(attendances, attendance)
	}

	return attendances, nil
}

func (p *postgres) CreateTask(r utils.CreateTaskRequestModel) (utils.CreateTaskResponseModel, error) {
	var (
		id int
	)

	if err := p.db.QueryRow(utils.CreateTask, r.Title, r.Description, r.StartAt, r.FinishAt, r.ProjectID, r.ProgrammerID, r.Attachment).Scan(&id); err != nil {
		return utils.CreateTaskResponseModel{}, err
	}
	return utils.CreateTaskResponseModel{
		ID: id,
	}, nil
}

func (p *postgres) UpdateTaskStatus(r utils.UpdateTaskStatusRequestModel) error {
	_, err := p.db.Exec(utils.UpdateTaskStatus, r.Status, r.ID)
	if err != nil {
		return err
	}
	return nil
}

func (p *postgres) GetTasks(r utils.GetTasksRequestModel) ([]models.Task, error) {
	rows, err := p.db.Query(utils.GetTasks, r.ProjectID, r.ProgrammerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var (
			task        models.Task
			finishedAt  sql.NullTime
			startedAt   sql.NullTime
			attachment  sql.NullString
			description sql.NullString
		)
		if err := rows.Scan(&task.ID, &task.Title, &description, &task.StartAt, &task.FinishAt, &startedAt, &finishedAt, &task.Status, &task.ProgrammerID, &task.ProjectID, &attachment); err != nil {
			return nil, err
		}
		task.FinishedAt = finishedAt.Time.String()
		task.StartedAt = startedAt.Time.String()
		task.Attachment = attachment.String
		task.Description = description.String
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (p *postgres) GetTask(id int) (models.Task, error) {
	var (
		task        models.Task
		finishedAt  sql.NullTime
		startedAt   sql.NullTime
		attachment  sql.NullString
		description sql.NullString
	)
	if err := p.db.QueryRow(utils.GetTask, id).Scan(&task.ID, &task.Title, &description, &task.StartAt, &task.FinishAt, &startedAt, &finishedAt, &task.Status, &task.ProgrammerID, &task.ProjectID, &attachment); err != nil {
		return task, err
	}
	task.FinishedAt = finishedAt.Time.String()
	task.StartedAt = startedAt.Time.String()
	task.Attachment = attachment.String
	task.Description = description.String
	return task, nil
}

func (p *postgres) DeleteTask(r utils.DeleteTaskRequestModel) error {
	if err := p.Login(utils.LoginRequestModel{
		Phone:    r.Phone,
		Password: r.Password,
	}); err != nil {
		return err
	}
	_, err := p.db.Exec(utils.DeleteTask, r.ID)
	if err != nil {
		return err
	}
	return nil
}

func (p *postgres) StartOrFinishTask(r utils.StartOrFinishTaskRequestModel) error {
	switch r.Type {
	case "start":
		_, err := p.db.Exec(utils.StartTask, r.Date, r.ID)
		if err != nil {
			return err
		}
	case "finish":
		_, err := p.db.Exec(utils.FinishTask, r.Date, r.ID)
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *postgres) WriteComment(r utils.WriteCommentRequestModel) (utils.WriteCommentResponseModel, error) {
	var (
		id int
	)

	if err := p.db.QueryRow(utils.WriteComment, r.TaskID, r.ProgrammerID, r.Text).Scan(&id); err != nil {
		return utils.WriteCommentResponseModel{}, err
	}
	return utils.WriteCommentResponseModel{
		ID: id,
	}, nil
}

func (p *postgres) GetComments(r utils.GetCommentsRequestModel) ([]models.Comment, error) {
	rows, err := p.db.Query(utils.GetComments, r.TaskID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []models.Comment
	for rows.Next() {
		var (
			comment models.Comment
		)
		if err := rows.Scan(&comment.ID, &comment.Text, &comment.CreatedAt, &comment.ProgrammerID, &comment.TaskID); err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}

	return comments, nil
}

func (p *postgres) DeleteComment(r utils.DeleteCommentRequestModel) error {
	if err := p.Login(utils.LoginRequestModel{
		Phone:    r.Phone,
		Password: r.Password,
	}); err != nil {
		return err
	}
	_, err := p.db.Exec(utils.DeleteComment, r.ID)
	if err != nil {
		return err
	}
	return nil
}
