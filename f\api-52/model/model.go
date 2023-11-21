package model

type Owners struct {
	ID uint
	Name string
	Email string
}

type User struct {
	ID uint
	Name string
	Role string
	password string 
}