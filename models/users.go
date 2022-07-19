package models

type User struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	BirthYear int    `json:"birth_year" binding:"required"`
	GroupId   int    `json:"group_id" binding:"required"`
}

type UserGroup struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	BirthYear int    `json:"birth_year"`
	MyGroup   Group  `json:"group"`
}
