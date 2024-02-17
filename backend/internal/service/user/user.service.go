package user

import (
	"errors"

	"github.com/Noblefel/ManorTalk/backend/internal/config"
	"github.com/Noblefel/ManorTalk/backend/internal/models"
	"github.com/Noblefel/ManorTalk/backend/internal/repository"
	"github.com/Noblefel/ManorTalk/backend/internal/repository/postgres"
	"github.com/Noblefel/ManorTalk/backend/internal/repository/redis"
)

var (
	ErrDuplicateUsername = errors.New("Username already taken")
	ErrNoUser            = errors.New("User not found")
	ErrUnauthorized      = errors.New("You have no permission to do that")
	ErrAvatarTooLarge    = errors.New("Avatar image is too large (2MB max)")
	ErrAvatarInvalid     = errors.New("Invalid type, avatar should be jpg/jpeg/png")
)

type UserService interface {
	CheckUsername(username string) error
	Get(username string) (models.User, error)
	UpdateProfile(payload models.UpdateProfileInput, username string, authId int) (string, error)
}

type userService struct {
	c         *config.AppConfig
	cacheRepo repository.CacheRepo
	userRepo  repository.UserRepo
}

func NewUserService(c *config.AppConfig, cr repository.CacheRepo, ur repository.UserRepo) UserService {
	return &userService{
		c:         c,
		cacheRepo: cr,
		userRepo:  ur,
	}
}

// mockUserService is a lightweight replicate of the user service used inside handler tests
type mockUserService struct {
	cacheRepo repository.CacheRepo
	userRepo  repository.UserRepo
}

func NewMockUserService() UserService {
	return &mockUserService{
		cacheRepo: redis.NewMockRepo(),
		userRepo:  postgres.NewMockUserRepo(),
	}
}
