package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/i-akbarshoh/task-manager/pkg/infrastructure/router/middleware"
	"github.com/i-akbarshoh/task-manager/pkg/registry/controller"
)

func NewRouter(c controller.Controller) *gin.Engine {
	r := gin.Default()
	//r.Use(gin.Logger())
	//r.Use(gin.Recovery())
	r.Use(cors.Default())
	r.Use(middleware.Authorizer())
	r.GET("/ping", c.Ping)
	r.POST("/register", c.Register)
	r.POST("/login", c.Login)
	r.GET("/programmers", c.GetProgrammers)
	r.GET("/programmer", c.GetProgrammer) // TODO: authorization is not working
	r.PUT("/programmer", c.UpdatePassword)
	r.DELETE("/programmer", c.DeleteUser)
	r.POST("/create-project", c.CreateProject)
	r.GET("/projects", c.GetProjects)
	r.PUT("/update-project-status", c.UpdateProjectStatus)
	r.DELETE("/delete-project", c.DeleteProject)
	r.GET("/project", c.GetProject)
	r.POST("/add-attendance", c.AddAttendance)
	r.GET("/user-attendances", c.GetUserAttendances)
	r.POST("/create-task", c.CreateTask)
	r.PUT("/update-task-status", c.UpdateTaskStatus)
	r.GET("/tasks", c.GetTasks)
	r.GET("/task", c.GetTask)
	r.DELETE("/delete-task", c.DeleteTask)
	r.PUT("/update-task", c.UpdateTask)
	r.POST("/write-comment", c.WriteComment)
	r.GET("/comments", c.GetComments)
	r.DELETE("/delete-comment", c.DeleteComment)
	return r
}
