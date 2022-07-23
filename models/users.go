package models

type User struct {
	ID        int    `json:"id" db:"id"`
	FirstName string `json:"first_name" binding:"required" db:"first_name"`
	LastName  string `json:"last_name" binding:"required" db:"last_name"`
	BirthYear int    `json:"birth_year" binding:"required" db:"birth_year"`
	GroupID   int    `json:"group_id" binding:"required" db:"group_id"`
}

type UserGroup struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	BirthYear int    `json:"birth_year"`
	MyGroup   Group  `json:"group"`
}
