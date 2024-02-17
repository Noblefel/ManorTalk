package models

import (
	"mime/multipart"
	"time"
)

type User struct {
	Id         int       `json:"id,omitempty"`
	Name       string    `json:"name,omitempty"`
	Username   string    `json:"username,omitempty"`
	Avatar     string    `json:"avatar,omitempty"`
	Email      string    `json:"email,omitempty"`
	Password   string    `json:"password,omitempty"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	PostsCount int       `json:"posts_count,omitempty"`
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

type UpdateProfileInput struct {
	Name     string `json:"name" validate:"max=255"`
	Username string `json:"username" validate:"required,min=3,max=40,excludesall=~%^;'<>()[]@!#/&*"`
	Files    map[string][]*multipart.FileHeader
}
