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
	DeleteUser(id string) error
}
