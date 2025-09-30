package repository

import (
	"database/sql"
	"embed"

	"github.com/pressly/goose/v3"
)

//go:embed migrations/*.sql
var migrationsFS embed.FS

func RunMigrations(db *sql.DB) error {
	goose.SetBaseFS(migrationsFS)

	return goose.Up(db, "migrations")
}
