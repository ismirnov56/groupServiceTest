package repository

import (
	"app/models"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (r *UserPostgres) CreateUser(user models.User) (models.User, error) {
	var result models.User

	query := fmt.Sprintf("INSERT INTO %s (first_name, last_name, birth_year, group_id) "+
		"values ($1, $2, $3, $4) RETURNING id, first_name, last_name, birth_year, group_id", usersTable)

	row := r.db.QueryRow(query, user.FirstName, user.LastName, user.BirthYear, user.GroupId)

	if err := row.Scan(&result.Id, &result.FirstName, &result.LastName, &result.BirthYear, &result.GroupId); err != nil {
		return result, err
	}

	return result, nil
}
