package repository

import (
	"app/models"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type GroupInfoPostgres struct {
	db *sqlx.DB
}

func NewGroupInfoPostgres(db *sqlx.DB) *GroupInfoPostgres {
	return &GroupInfoPostgres{db: db}
}

func (r *GroupInfoPostgres) GetGroupInfoById(groupId int) (models.GroupInfo, error) {
	var groupInfo models.GroupInfo

	query := fmt.Sprintf(
		`WITH RECURSIVE r as (
				  SELECT id, parent_id
				  FROM %[1]s
				  WHERE id = $1
				  UNION
				  SELECT g.id, g.parent_id
				  FROM %[1]s g,
					   r
				  WHERE r.id = g.parent_id
				  ),
			count_group_with_recursive as (SELECT count(*) as user_count_recursive
										   FROM %[2]s
										   WHERE group_id in (SELECT id FROM r)),
			count_in_group as (SELECT count(*) as user_count
							   FROM %[2]s
							   WHERE group_id = $1)
		SELECT id, name, parent_id, count_in_group.user_count, count_group_with_recursive.user_count_recursive
		FROM %[1]s,
			 count_in_group,
			 count_group_with_recursive
		WHERE id = $1;`,
		groupsTable, usersTable)

	if err := r.db.Get(&groupInfo, query, groupId); err != nil {
		return groupInfo, err
	}

	return groupInfo, nil
}
