package models

type Group struct {
	Id       int    `json:"id" db:"id"`
	Name     string `json:"name" binding:"required" db:"name"`
	ParentId *int64 `json:"parent_id" db:"parent_id"`
}

type GroupInfo struct {
	Id                 int    `json:"id" db:"id"`
	Name               string `json:"name" binding:"required" db:"name"`
	ParentId           *int64 `json:"parent_id" db:"parent_id"`
	UserCount          int    `json:"user_count" db:"user_count"`
	UserCountRecursive int    `json:"user_count_recursive" db:"user_count_recursive"`
}
