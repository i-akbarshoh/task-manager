package datastore

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/i-akbarshoh/task-manager/pkg/config"
	_ "github.com/lib/pq"
	"os"
)

func NewDB() *sql.DB {
	DBMS := "postgres"
	db, err := sql.Open(DBMS, fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.C.Database.Host, config.C.Database.Port, config.C.Database.User, config.C.Database.Password, config.C.Database.DBName,
	))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if err := db.Ping(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// migrate
	migrateUp(db)

	return db
}

func migrateUp(db *sql.DB) {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://database/migrations",
		config.C.Database.DBName, driver)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		fmt.Println(err)
		os.Exit(1)
	}
}
