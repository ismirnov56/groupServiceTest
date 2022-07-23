package repository

import (
	"app/models"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type GroupPostgres struct {
	db *sqlx.DB
}

func NewGroupPostgres(db *sqlx.DB) *GroupPostgres {
	return &GroupPostgres{db: db}
}

func (r *GroupPostgres) CreateGroup(group models.Group) (models.Group, error) {
	var resultGroup models.Group

	query := fmt.Sprintf("INSERT INTO %s (name, parent_id) "+
		"values ($1, $2) RETURNING id, name, parent_id", groupsTable)

	row := r.db.QueryRow(query, group.Name, group.ParentID)

	if err := row.Scan(&resultGroup.ID, &resultGroup.Name, &resultGroup.ParentID); err != nil {
		return resultGroup, err
	}

	return resultGroup, nil
}

func (r *GroupPostgres) GetAllGroups() ([]models.Group, error) {
	var groups []models.Group

	query := fmt.Sprintf("select * FROM %s", groupsTable)

	if err := r.db.Select(&groups, query); err != nil {
		return nil, err
	}

	return groups, nil
}

func (r *GroupPostgres) GetGroupByID(groupId int) (models.Group, error) {
	var group models.Group

	query := fmt.Sprintf("select * FROM %s WHERE id = $1", groupsTable)

	if err := r.db.Get(&group, query, groupId); err != nil {
		return group, err
	}

	return group, nil
}

func (r *GroupPostgres) DeleteGroup(groupID int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", groupsTable)

	_, err := r.db.Exec(query, groupID)

	return err
}

func (r *GroupPostgres) UpdateGroup(groupID int, group models.Group) (models.Group, error) {
	var resultGroup models.Group

	query := fmt.Sprintf("UPDATE %s SET name = $1, parent_id = $2 "+
		"WHERE id = $3 RETURNING id, name, parent_id", groupsTable)

	row := r.db.QueryRow(query, group.Name, group.ParentID, groupID)

	if err := row.Scan(&resultGroup.ID, &resultGroup.Name, &resultGroup.ParentID); err != nil {
		return resultGroup, err
	}

	return resultGroup, nil
}
