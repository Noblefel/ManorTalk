package repository

import (
	"github.com/Noblefel/ManorTalk/backend/internal/models"
	"github.com/Noblefel/ManorTalk/backend/internal/utils/token"
)

// Default values to maintain consistency throughout test repos.
// Will use the equal "==" operator
const (
	ErrNotFoundKeyInt      = -1
	ErrNotFoundKeyString   = "not-found"
	ErrUnexpectedKeyInt    = -2
	ErrUnexpectedKeyString = "unexpected-error"
	ErrIncorrectKey        = "something-incorrect"
	ErrDuplicateKeyString  = "already-exists"
)

type RedisRepo interface {
	SetRefreshToken(td token.Details) error
	GetRefreshToken(td token.Details) (string, error)
}

type UserRepo interface {
	// Register is a small wrapper around CreateUser and bcrypt.GenerateFromPassword
	Register(payload models.UserRegisterInput) (models.User, error)
	// Login is a small wrapper around GetUserByEmail and bcrypt.CompareHashAndPassword
	Login(payload models.UserLoginInput) (models.User, error)

	CreateUser(email, password string) (models.User, error)
	GetUserById(id int) (models.User, error)
	GetUserByEmail(email string) (models.User, error)
}

type PostRepo interface {
	CreatePost(p models.Post) (models.Post, error)
	GetPostBySlug(slug string) (models.Post, error)
	GetPostsByCategory(slug string) ([]models.Post, error)
	UpdatePost(p models.Post) error
	DeletePost(id int) error

	GetCategories() ([]models.Category, error)
	GetCategoryById(id int) (models.Category, error)
}
