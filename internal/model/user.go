package model

type User struct {
	ID   string `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}
