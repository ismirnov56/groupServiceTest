package models

type User struct {
	Id        int    `json:"id" db:"id"`
	FirstName string `json:"first_name" binding:"required" db:"first_name"`
	LastName  string `json:"last_name" binding:"required" db:"last_name"`
	BirthYear int    `json:"birth_year" binding:"required" db:"birth_year"`
	GroupId   int    `json:"group_id" binding:"required" db:"group_id"`
}

type UserGroup struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	BirthYear int    `json:"birth_year"`
	MyGroup   Group  `json:"group"`
}
