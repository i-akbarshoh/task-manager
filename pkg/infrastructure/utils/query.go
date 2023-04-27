package utils

var (
	Register = `
		INSERT INTO users (id, full_name, password, avatar, role, birth_date, phone, position) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`

	Login = `
		SELECT password FROM users WHERE phone = $1
	`

	GetProgrammers = `
		SELECT * FROM users WHERE position = 'programmer'
	`

	GetProgrammer = `
		SELECT * FROM users WHERE id = $1
	`

	UpdatePassword = `
		UPDATE users SET password = $1 WHERE phone = $2
	`

	DeleteUser = `
		DELETE FROM users WHERE id = $1
	`
)
