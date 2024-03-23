package repository

import (
	"github.com/Noblefel/ManorTalk/backend/internal/models"
	"github.com/Noblefel/ManorTalk/backend/internal/utils/pagination"
	"github.com/Noblefel/ManorTalk/backend/internal/utils/token"
)

// Default values to maintain consistency throughout mock repos.
const (
	NotFoundKeyInt   = -1
	NotFoundKey      = "not-found"
	UnexpectedKeyInt = -2
	UnexpectedKey    = "unexpected-error"
	IncorrectKey     = "something-incorrect"
	DuplicateKey     = "already-exists"
)

type CacheRepo interface {
	SetRefreshToken(td token.Details) error
	GetRefreshToken(td token.Details) (string, error)
	DelRefreshToken(td token.Details) error
}

type UserRepo interface {
	CreateUser(username, email, password string) (int, error)
	GetUser(filters models.UserFilters) (models.User, error)
	UpdateUser(u models.User) error
}

type PostRepo interface {
	CreatePost(p models.Post) (models.Post, error)
	GetPosts(pgMeta *pagination.Meta, filters models.PostsFilters) ([]models.Post, error)
	GetPostBySlug(slug string) (models.Post, error)
	UpdatePost(p models.Post) error
	DeletePost(id int) error
	CountPosts(filters models.PostsFilters) (int, error)

	GetCategories() ([]models.Category, error)
	GetCategoryById(id int) (models.Category, error)
	GetCategoryBySlug(slug string) (models.Category, error)
}
