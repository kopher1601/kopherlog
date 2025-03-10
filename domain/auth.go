package domain

import "time"

type Login struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type SignIn struct {
	Email    string
	Password string
}

type User struct {
	ID        int
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}
