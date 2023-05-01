package repository

import (
	"github.com/i-akbarshoh/task-manager/pkg/domain/models"
	"github.com/i-akbarshoh/task-manager/pkg/infrastructure/utils"
)

type Repository interface {
	Register(user models.User) error
	Login(r utils.LoginRequestModel) error
	GetProgrammers() ([]models.User, error)
	GetProgrammer(id string) (models.User, error)
	UpdatePassword(r utils.UpdatePasswordRequestModel) error
	DeleteUser(r utils.DeleteUserRequestModel) error
	CreateProject(project models.Project) error
	GetProjects() ([]models.Project, error)
	UpdateProjectStatus(r utils.UpdateProjectStatusRequestModel) error
	DeleteProject(id int) error
	GetProject(id int) (models.Project, error)
	AddAttendance(r utils.AddAttendanceRequestModel) error
	GetUserAttendances(r utils.GetUserAttendanceRequestModel) ([]models.Attendance, error)
	CreateTask(r utils.CreateTaskRequestModel) (utils.CreateTaskResponseModel, error)
	UpdateTaskStatus(r utils.UpdateTaskStatusRequestModel) error
	GetTasks(r utils.GetTasksRequestModel) ([]models.Task, error)
	GetTask(id int) (models.Task, error)
	DeleteTask(r utils.DeleteTaskRequestModel) error
	StartOrFinishTask(r utils.StartOrFinishTaskRequestModel) error
	WriteComment(r utils.WriteCommentRequestModel) (utils.WriteCommentResponseModel, error)
	GetComments(r utils.GetCommentsRequestModel) ([]models.Comment, error)
	DeleteComment(r utils.DeleteCommentRequestModel) error
}
