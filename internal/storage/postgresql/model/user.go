package model

type User struct {
	id       int64
	username string
}

func NewUser() *User {
	return &User{
		id:       0,
		username: "",
	}
}
