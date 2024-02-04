package models

import "time"

type User struct {
	Id        int       `json:"id"`
	Name      string    `json:"name,omitempty"`
	Username  string    `json:"username"`
	Avatar    string    `json:"avatar,omitempty"`
	Email     string    `json:"email"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserRegisterInput struct {
	Username string `json:"username" validate:"required,min=3,max=40,excludesall=~%^;'<>()[]@!#/&*"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=72"`
}

type UserLoginInput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=72"`
}

type CheckUsernameInput struct {
	Username string `json:"username" validate:"required,min=3,max=40,excludesall=~%^;'<>()[]@!#/&*"`
}
