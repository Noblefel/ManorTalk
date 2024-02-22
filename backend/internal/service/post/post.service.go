package post

import (
	"errors"
	"net/url"

	"github.com/Noblefel/ManorTalk/backend/internal/config"
	"github.com/Noblefel/ManorTalk/backend/internal/models"
	"github.com/Noblefel/ManorTalk/backend/internal/repository"
	"github.com/Noblefel/ManorTalk/backend/internal/repository/postgres"
	"github.com/Noblefel/ManorTalk/backend/internal/repository/redis"
	"github.com/Noblefel/ManorTalk/backend/internal/utils/pagination"
)

var (
	ErrDuplicateTitle = errors.New("Title has already been used")
	ErrNoCategory     = errors.New("Category not found")
	ErrNoPost         = errors.New("Post not found")
	ErrUnauthorized   = errors.New("You have no permission to do that")
)

type PostService interface {
	Create(payload models.PostCreateInput, authId int) (models.Post, error)
	Get(slug string) (models.Post, error)
	GetMany(q url.Values) ([]models.Post, *pagination.Meta, error)
	Update(payload models.PostUpdateInput, urlSlug string, authId int) error
	Delete(slug string, authId int) error
	GetCategories() ([]models.Category, error)
}

type postService struct {
	c         *config.AppConfig
	cacheRepo repository.CacheRepo
	postRepo  repository.PostRepo
}

func NewPostService(c *config.AppConfig, cr repository.CacheRepo, pr repository.PostRepo) PostService {
	return &postService{
		c:         c,
		cacheRepo: cr,
		postRepo:  pr,
	}
}

// mockPostService is a lightweight replicate of the post service used inside handler tests
type mockPostService struct {
	cacheRepo repository.CacheRepo
	postRepo  repository.PostRepo
}

func NewMockPostService() PostService {
	return &mockPostService{
		cacheRepo: redis.NewMockRepo(),
		postRepo:  postgres.NewMockPostRepo(),
	}
}
