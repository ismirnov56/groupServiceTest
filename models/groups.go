package models

type Group struct {
	ID       int    `json:"id" db:"id"`
	Name     string `json:"name" binding:"required" db:"name"`
	ParentID *int64 `json:"parent_id" db:"parent_id"`
}

type GroupInfo struct {
	ID                 int    `json:"id" db:"id"`
	Name               string `json:"name" binding:"required" db:"name"`
	ParentID           *int64 `json:"parent_id" db:"parent_id"`
	UserCount          int    `json:"user_count" db:"user_count"`
	UserCountRecursive int    `json:"user_count_recursive" db:"user_count_recursive"`
}
