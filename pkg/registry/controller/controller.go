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
}
