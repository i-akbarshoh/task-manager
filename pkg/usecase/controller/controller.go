package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/i-akbarshoh/task-manager/pkg/domain/enums"
	"github.com/i-akbarshoh/task-manager/pkg/domain/models"
	"github.com/i-akbarshoh/task-manager/pkg/infrastructure/utils"
	"github.com/i-akbarshoh/task-manager/pkg/registry/repository"
	"strconv"
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
	var (
		body utils.DeleteUserRequestModel
	)

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{
			"message": "cannot bind body, " + err.Error(),
		})
		return
	}
	if err := con.r.DeleteUser(body); err != nil {
		c.JSON(500, gin.H{
			"message": "cannot delete user, " + err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
	})
}

func (con *controller) CreateProject(c *gin.Context) {
	var (
		body models.Project
	)

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{
			"message": "cannot bind body, " + err.Error(),
		})
		return
	}
	if err := con.r.CreateProject(body); err != nil {
		c.JSON(500, gin.H{
			"message": "cannot create project, " + err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
	})
}

func (con *controller) GetProjects(c *gin.Context) {
	projects, err := con.r.GetProjects()
	if err != nil {
		c.JSON(500, gin.H{
			"message": "cannot get projects, " + err.Error(),
		})
		return
	}
	c.JSON(200, projects)
}

func (con *controller) UpdateProjectStatus(c *gin.Context) {
	var (
		body utils.UpdateProjectStatusRequestModel
	)

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{
			"message": "cannot bind body, " + err.Error(),
		})
		return
	}
	if !utils.IsEnumValue(enums.StatusType, body.Status) {
		c.JSON(400, gin.H{
			"message": "invalid status",
		})
		return
	}
	if err := con.r.UpdateProjectStatus(body); err != nil {
		c.JSON(500, gin.H{
			"message": "cannot update project status, " + err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
	})
}

func (con *controller) DeleteProject(c *gin.Context) {
	var (
		body utils.DeleteProjectRequestModel
	)

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{
			"message": "cannot bind body, " + err.Error(),
		})
		return
	}
	if err := con.r.DeleteProject(body.ID); err != nil {
		c.JSON(500, gin.H{
			"message": "cannot delete project, " + err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
	})
}

func (con *controller) GetProject(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "invalid id, " + err.Error(),
		})
		return
	}
	project, err := con.r.GetProject(idInt)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "cannot get project, " + err.Error(),
		})
		return
	}
	c.JSON(200, project)
}

func (con *controller) AddAttendance(c *gin.Context) {
	var (
		body utils.AddAttendanceRequestModel
	)

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{
			"message": "cannot bind body, " + err.Error(),
		})
		return
	}
	if !utils.IsEnumValue(enums.AttendanceType, body.Type) {
		c.JSON(400, gin.H{
			"message": "invalid attendance type",
		})
		return
	}
	if err := con.r.AddAttendance(body); err != nil {
		c.JSON(500, gin.H{
			"message": "cannot add attendance, " + err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
	})
}

func (con *controller) GetUserAttendances(c *gin.Context) {
	var (
		body utils.GetUserAttendanceRequestModel
	)
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{
			"message": "cannot bind body, " + err.Error(),
		})
		return
	}
	if !utils.IsEnumValue(enums.AttendanceType, body.Type) {
		c.JSON(400, gin.H{
			"message": "invalid attendance type",
		})
		return
	}

	attendances, err := con.r.GetUserAttendances(body)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "cannot get user attendances, " + err.Error(),
		})
		return
	}
	c.JSON(200, attendances)
}

func (con *controller) CreateTask(c *gin.Context) {
	var (
		body utils.CreateTaskRequestModel
	)

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{
			"message": "cannot bind body, " + err.Error(),
		})
		return
	}
	res, err := con.r.CreateTask(body)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "cannot create task, " + err.Error(),
		})
		return
	}

	c.JSON(200, res)
}

func (con *controller) UpdateTaskStatus(c *gin.Context) {
	var (
		body utils.UpdateTaskStatusRequestModel
	)

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{
			"message": "cannot bind body, " + err.Error(),
		})
		return
	}
	if !utils.IsEnumValue(enums.StatusType, body.Status) {
		c.JSON(400, gin.H{
			"message": "invalid status",
		})
		return
	}
	if err := con.r.UpdateTaskStatus(body); err != nil {
		c.JSON(500, gin.H{
			"message": "cannot update task status, " + err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
	})
}

func (con *controller) GetTasks(c *gin.Context) {
	var (
		body utils.GetTasksRequestModel
	)
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{
			"message": "cannot bind body, " + err.Error(),
		})
		return
	}

	tasks, err := con.r.GetTasks(body)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "cannot get tasks, " + err.Error(),
		})
		return
	}
	c.JSON(200, tasks)
}

func (con *controller) GetTask(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "invalid id, " + err.Error(),
		})
		return
	}
	task, err := con.r.GetTask(idInt)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "cannot get task, " + err.Error(),
		})
		return
	}
	c.JSON(200, task)
}

func (con *controller) DeleteTask(c *gin.Context) {
	var (
		body utils.DeleteTaskRequestModel
	)

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{
			"message": "cannot bind body, " + err.Error(),
		})
		return
	}
	if err := con.r.DeleteTask(body); err != nil {
		c.JSON(500, gin.H{
			"message": "cannot delete task, " + err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
	})
}

func (con *controller) UpdateTask(c *gin.Context) {
	var (
		body utils.StartOrFinishTaskRequestModel
	)

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{
			"message": "cannot bind body, " + err.Error(),
		})
		return
	}
	if !utils.IsEnumValue([]string{"start", "finish"}, body.Type) {
		c.JSON(400, gin.H{
			"message": "invalid status",
		})
		return
	}

	if err := con.r.StartOrFinishTask(body); err != nil {
		c.JSON(500, gin.H{
			"message": "cannot update task, " + err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
	})
}

func (con *controller) WriteComment(c *gin.Context) {
	var (
		body utils.WriteCommentRequestModel
	)

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{
			"message": "cannot bind body, " + err.Error(),
		})
		return
	}

	res, err := con.r.WriteComment(body)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "cannot write comment, " + err.Error(),
		})
		return
	}

	c.JSON(200, res)
}

func (con *controller) GetComments(c *gin.Context) {
	var (
		body utils.GetCommentsRequestModel
	)

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{
			"message": "cannot bind body, " + err.Error(),
		})
		return
	}

	res, err := con.r.GetComments(body)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "cannot get comments, " + err.Error(),
		})
		return
	}

	c.JSON(200, res)
}

func (con *controller) DeleteComment(c *gin.Context) {
	var (
		body utils.DeleteCommentRequestModel
	)

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{
			"message": "cannot bind body, " + err.Error(),
		})
		return
	}

	if err := con.r.DeleteComment(body); err != nil {
		c.JSON(500, gin.H{
			"message": "cannot delete comment, " + err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
	})
}
