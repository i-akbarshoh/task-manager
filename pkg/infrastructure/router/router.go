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
	r.GET("/programmer/:id", c.GetProgrammer) // TODO: authorization is not working
	r.PUT("/programmer", c.UpdatePassword)
	r.DELETE("/programmer/:id", c.DeleteUser)
	return r
}
