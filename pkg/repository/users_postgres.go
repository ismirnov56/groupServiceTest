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

	row := r.db.QueryRow(query, user.FirstName, user.LastName, user.BirthYear, user.GroupID)

	if err := row.Scan(&result.ID, &result.FirstName, &result.LastName, &result.BirthYear, &result.GroupID); err != nil {
		return result, err
	}

	return result, nil
}

func (r *UserPostgres) GetAllUsers() ([]models.User, error) {
	var users []models.User

	query := fmt.Sprintf("select * FROM %s", usersTable)

	if err := r.db.Select(&users, query); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserPostgres) GetUserByID(userID int) (models.User, error) {
	var user models.User

	query := fmt.Sprintf("select * FROM %s WHERE id = $1", usersTable)

	if err := r.db.Get(&user, query, userID); err != nil {
		return user, err
	}

	return user, nil
}

func (r *UserPostgres) DeleteUser(userID int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", usersTable)

	_, err := r.db.Exec(query, userID)

	return err
}

func (r *UserPostgres) UpdateUser(userID int, user models.User) (models.User, error) {
	var resultUser models.User

	query := fmt.Sprintf("UPDATE %s SET first_name = $1, last_name = $2, birth_year = $3, group_id = $4"+
		"WHERE id = $5 RETURNING id, first_name, last_name, birth_year, group_id", usersTable)

	row := r.db.QueryRow(query, user.FirstName, user.LastName, user.BirthYear, user.GroupID, userID)

	if err := row.Scan(
		&resultUser.ID,
		&resultUser.FirstName,
		&resultUser.LastName,
		&resultUser.BirthYear,
		&resultUser.GroupID,
	); err != nil {
		return resultUser, err
	}

	return resultUser, nil
}
