package auth

import (
	"errors"

	"github.com/Noblefel/ManorTalk/backend/internal/config"
	"github.com/Noblefel/ManorTalk/backend/internal/models"
	"github.com/Noblefel/ManorTalk/backend/internal/repository"
	"github.com/Noblefel/ManorTalk/backend/internal/repository/postgres"
	"github.com/Noblefel/ManorTalk/backend/internal/repository/redis"
)

var (
	ErrDuplicateEmail     = errors.New("Email already in use")
	ErrInvalidCredentials = errors.New("Invalid credentials")
	ErrNoUser             = errors.New("User not found")
	ErrUnauthorized       = errors.New("Session invalid or expired, please login first")
)

type AuthService interface {
	Register(payload models.UserRegisterInput) error
	Login(payload models.UserLoginInput) (models.User, string, string, error)
	Refresh(refreshToken string) (models.User, string, error)
	Logout(refreshToken string) error
}

type authService struct {
	c         *config.AppConfig
	cacheRepo repository.CacheRepo
	userRepo  repository.UserRepo
}

func NewAuthService(c *config.AppConfig, cr repository.CacheRepo, ur repository.UserRepo) AuthService {
	return &authService{
		c:         c,
		cacheRepo: cr,
		userRepo:  ur,
	}
}

// mockAuthService is a lightweight replicate of the auth service used inside handler tests
type mockAuthService struct {
	cacheRepo repository.CacheRepo
	userRepo  repository.UserRepo
}

func NewMockAuthService() AuthService {
	return &mockAuthService{
		cacheRepo: redis.NewMockRepo(),
		userRepo:  postgres.NewMockUserRepo(),
	}
}
