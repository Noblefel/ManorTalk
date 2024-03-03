package models

import (
	"mime/multipart"
	"time"
)

type Post struct {
	Id         int       `json:"id,omitempty"`
	UserId     int       `json:"user_id,omitempty"`
	Title      string    `json:"title,omitempty"`
	Slug       string    `json:"slug,omitempty"`
	Excerpt    string    `json:"excerpt,omitempty"`
	Image      string    `json:"image,omitempty"`
	Content    string    `json:"content,omitempty"`
	CategoryId int       `json:"category_id,omitempty"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	Category   Category  `json:"category,omitempty"`
	User       User      `json:"user,omitempty"`
}

type PostCreateInput struct {
	Title      string `json:"title" validate:"required,min=10,max=50,excludesall=~%^;'<>"`
	Slug       string `json:"slug" validate:"required,min=10,max=255"`
	Excerpt    string `json:"excerpt" validate:"max=255"`
	Content    string `json:"content" validate:"required,min=50"`
	CategoryId int    `json:"category_id" validate:"required"`
	Files      map[string][]*multipart.FileHeader
}

type PostUpdateInput struct {
	Title      string `json:"title" validate:"required,min=10,max=50,excludesall=~%^;'<>"`
	Slug       string `json:"slug" validate:"required,min=10,max=255"`
	Excerpt    string `json:"excerpt" validate:"max=255"`
	Content    string `json:"content" validate:"required,min=50"`
	CategoryId int    `json:"category_id" validate:"required"`
	Files      map[string][]*multipart.FileHeader
}

type Category struct {
	Id        *int       `json:"id,omitempty"`
	Name      string     `json:"name"`
	Slug      string     `json:"slug"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

type PostsFilters struct {
	Order    string
	Category string
	Search   string
	Cursor   int
	UserId   int
	Limit    int
}
