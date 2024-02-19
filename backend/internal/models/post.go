package models

import "time"

type Post struct {
	Id         int       `json:"id"`
	UserId     int       `json:"user_id"`
	Title      string    `json:"title"`
	Slug       string    `json:"slug"`
	Excerpt    string    `json:"excerpt"`
	Content    string    `json:"content"`
	CategoryId int       `json:"category_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	Category   Category  `json:"category"`
	User       User      `json:"user"`
}

type PostCreateInput struct {
	Title      string `json:"title" validate:"required,min=10,max=255,excludesall=~%^;'<>"`
	Slug       string `json:"slug" validate:"required,min=10,max=255"`
	Excerpt    string `json:"excerpt" validate:"max=255"`
	Content    string `json:"content" validate:"required,min=50"`
	CategoryId int    `json:"category_id" validate:"required"`
}

type PostUpdateInput struct {
	Title      string `json:"title" validate:"required,min=10,max=255,excludesall=~%^;'<>"`
	Slug       string `json:"slug" validate:"required,min=10,max=255"`
	Excerpt    string `json:"excerpt" validate:"max=255"`
	Content    string `json:"content" validate:"required,min=50"`
	CategoryId int    `json:"category_id" validate:"required"`
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
