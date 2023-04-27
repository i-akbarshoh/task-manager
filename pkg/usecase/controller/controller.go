package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/i-akbarshoh/task-manager/pkg/domain/models"
	"github.com/i-akbarshoh/task-manager/pkg/infrastructure/utils"
	"github.com/i-akbarshoh/task-manager/pkg/registry/repository"
)

type controller struct {
	r repository.Repository
}

func New(r repository.Repository) *controller {
	return &controller{r: r}
}

func (con *controller) Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func (con *controller) Register(c *gin.Context) {
	var (
		body models.User
	)
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{
			"message": "cannot bind body, " + err.Error(),
		})
		return
	}
	id := uuid.New().String()
	tokens, err := utils.GenerateNewTokens(id, map[string]string{"role": body.Position})
	if err != nil {
		c.JSON(500, gin.H{
			"message": "cannot generate tokens, " + err.Error(),
		})
		return
	}

	password, err := utils.GeneratePasswordHash(body.Password)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "cannot generate password, " + err.Error(),
		})
		return
	}
	body.ID = id
	body.Password = string(password)

	if err := con.r.Register(body); err != nil {
		c.JSON(500, gin.H{
			"message": "cannot register user, " + err.Error(),
		})
		return
	}

	c.JSON(200, utils.RegisterResponseModel{
		ID:           id,
		AccessToken:  tokens.Access,
		ExpiresIn:    tokens.AccExpire,
		RefreshToken: tokens.Refresh,
	})
}

func (con *controller) Login(c *gin.Context) {
	var (
		body utils.LoginRequestModel
	)
	meta, err := utils.ExtractTokenMetadata(c)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "cannot parse token, " + err.Error(),
		})
		return
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{
			"message": "cannot bind body, " + err.Error(),
		})
		return
	}

	if err := con.r.Login(body); err != nil {
		c.JSON(500, gin.H{
			"message": "cannot login, " + err.Error(),
		})
		return
	}

	c.JSON(200, utils.LoginResponseModel{
		ID: meta.UserID.String(),
	})
}

func (con *controller) GetProgrammers(c *gin.Context) {
	programmers, err := con.r.GetProgrammers()
	if err != nil {
		c.JSON(500, gin.H{
			"message": "cannot get programmers, " + err.Error(),
		})
		return
	}
	c.JSON(200, programmers)
}

func (con *controller) GetProgrammer(c *gin.Context) {
	id := c.Param("id")
	programmer, err := con.r.GetProgrammer(id)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "cannot get programmer, " + err.Error(),
		})
		return
	}
	c.JSON(200, programmer)
}

func (con *controller) UpdatePassword(c *gin.Context) {
	var (
		body utils.UpdatePasswordRequestModel
	)

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{
			"message": "cannot bind body, " + err.Error(),
		})
		return
	}

	if err := con.r.UpdatePassword(body); err != nil {
		c.JSON(500, gin.H{
			"message": "cannot update password, " + err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
	})
}

func (con *controller) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	if err := con.r.DeleteUser(id); err != nil {
		c.JSON(500, gin.H{
			"message": "cannot delete user, " + err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
	})
}
