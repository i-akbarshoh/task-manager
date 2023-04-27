package repository

import (
	"database/sql"
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
	_, err := p.db.Exec(utils.UpdatePassword, r.NewPassword, r.Phone)
	if err != nil {
		return err
	}
	return nil
}

func (p *postgres) DeleteUser(id string) error {
	_, err := p.db.Exec(utils.DeleteUser, id)
	if err != nil {
		return err
	}
	return nil
}
