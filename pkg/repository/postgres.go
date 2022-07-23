package repository

import (
	"app/pkg"
	"fmt"
	"github.com/jmoiron/sqlx"
)

const (
	groupsTable = "groups"
	usersTable  = "users"
)

func NewPostgresDB(cfg pkg.Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.DBHost, cfg.DBPort, cfg.DBUsername, cfg.DBPassword, cfg.DBName, cfg.DBSSLMode,
	),
	)

	if err != nil {
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		return nil, err
	}

	return db, nil
}
