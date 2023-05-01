package controller

import "github.com/gin-gonic/gin"

type Controller interface {
	Ping(c *gin.Context)
	Register(c *gin.Context)
	Login(c *gin.Context)
	GetProgrammers(c *gin.Context)
	GetProgrammer(c *gin.Context)
	UpdatePassword(c *gin.Context)
	DeleteUser(c *gin.Context)
	CreateProject(c *gin.Context)
	GetProjects(c *gin.Context)
	UpdateProjectStatus(c *gin.Context)
	DeleteProject(c *gin.Context)
	GetProject(c *gin.Context)
	AddAttendance(c *gin.Context)
	GetUserAttendances(c *gin.Context)
	CreateTask(c *gin.Context)
	UpdateTaskStatus(c *gin.Context)
	GetTasks(c *gin.Context)
	GetTask(c *gin.Context)
	DeleteTask(c *gin.Context)
	UpdateTask(c *gin.Context)
	WriteComment(c *gin.Context)
	GetComments(c *gin.Context)
	DeleteComment(c *gin.Context)
}
