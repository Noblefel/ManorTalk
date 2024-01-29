package post

import (
	"net/url"
	"reflect"
	"testing"

	"github.com/Noblefel/ManorTalk/backend/internal/config"
	"github.com/Noblefel/ManorTalk/backend/internal/database"
	"github.com/Noblefel/ManorTalk/backend/internal/models"
	"github.com/Noblefel/ManorTalk/backend/internal/repository"
	"github.com/Noblefel/ManorTalk/backend/internal/repository/postgres"
	"github.com/Noblefel/ManorTalk/backend/internal/repository/redis"
)

func TestNewPostService(t *testing.T) {
	var db *database.DB
	var c *config.AppConfig
	cr := redis.NewRepo(db)
	pr := postgres.NewPostRepo(db)
	service := NewPostService(c, cr, pr)

	typeString := reflect.TypeOf(service).String()

	if typeString != "*post.postService" {
		t.Error("NewPostService() did not get the correct type, wanted *post.PostService")
	}
}

func TestNewMockPostService(t *testing.T) {
	service := NewMockPostService()

	typeString := reflect.TypeOf(service).String()

	if typeString != "*post.mockPostService" {
		t.Error("NewMockPostService() did not get the correct type, wanted *post.mockPostService")
	}
}

func newTestService() PostService {
	var tc config.AppConfig
	cr := redis.NewMockRepo()
	ur := postgres.NewMockPostRepo()

	service := NewPostService(&tc, cr, ur)

	return service
}

var s = newTestService()

func TestPostService_Create(t *testing.T) {
	var tests = []struct {
		name    string
		payload models.PostCreateInput
		isError bool
	}{
		{
			name: "create-ok",
			payload: models.PostCreateInput{
				Title: "A sample title",
			},
			isError: false,
		},
		{
			name: "create-error-no-category",
			payload: models.PostCreateInput{
				CategoryId: repository.ErrNotFoundKeyInt,
			},
			isError: true,
		},
		{
			name: "create-error-getting-category",
			payload: models.PostCreateInput{
				CategoryId: repository.ErrUnexpectedKeyInt,
			},
			isError: true,
		},
		{
			name: "create-error-duplicate-title",
			payload: models.PostCreateInput{
				Title: repository.ErrDuplicateKeyString,
			},
			isError: true,
		},
		{
			name: "create-error-creating-post",
			payload: models.PostCreateInput{
				Title: repository.ErrUnexpectedKeyString,
			},
			isError: true,
		},
	}

	for _, tt := range tests {
		_, err := s.Create(tt.payload)

		if err != nil && !tt.isError {
			t.Errorf("%s should not return error, but got %s", tt.name, err)
		}

		if err == nil && tt.isError {
			t.Errorf("%s should return error", tt.name)
		}
	}
}

func TestPostService_Get(t *testing.T) {
	var tests = []struct {
		name    string
		slug    string
		isError bool
	}{
		{
			name:    "get-ok",
			slug:    "example",
			isError: false,
		},
		{
			name:    "create-error-no-post",
			slug:    repository.ErrNotFoundKeyString,
			isError: true,
		},
		{
			name:    "create-error-getting-post",
			slug:    repository.ErrUnexpectedKeyString,
			isError: true,
		},
	}

	for _, tt := range tests {
		_, err := s.Get(tt.slug)

		if err != nil && !tt.isError {
			t.Errorf("%s should not return error, but got %s", tt.name, err)
		}

		if err == nil && tt.isError {
			t.Errorf("%s should return error", tt.name)
		}
	}
}

func TestPostService_GetMany(t *testing.T) {
	var tests = []struct {
		name    string
		q       url.Values
		isError bool
	}{
		{
			name: "getMany-ok",
			q: url.Values{
				"page":  {"1"},
				"total": {"10"},
			},
			isError: false,
		},
		{
			name: "getMany-error-no-category",
			q: url.Values{
				"category": {repository.ErrNotFoundKeyString},
			},
			isError: true,
		},
		{
			name: "getMany-error-getting-category",
			q: url.Values{
				"category": {repository.ErrUnexpectedKeyString},
			},
			isError: true,
		},
		{
			name: "getMany-error-creating-pagination-meta",
			q: url.Values{
				"page": {"-1"},
			},
			isError: true,
		},
		{
			name: "getMany-error-counting-posts",
			q: url.Values{
				"page":  {"1"},
				"order": {repository.ErrUnexpectedKeyString},
			},
			isError: true,
		},
		{
			name: "getMany-error-getting-posts",
			q: url.Values{
				"total": {"1"},
				"order": {repository.ErrUnexpectedKeyString},
			},
			isError: true,
		},
	}

	for _, tt := range tests {
		_, _, err := s.GetMany(tt.q)

		if err != nil && !tt.isError {
			t.Errorf("%s should not return error, but got %s", tt.name, err)
		}

		if err == nil && tt.isError {
			t.Errorf("%s should return error", tt.name)
		}
	}
}

func TestPostService_Update(t *testing.T) {
	var tests = []struct {
		name    string
		payload models.PostUpdateInput
		urlSlug string
		isError bool
	}{
		{
			name: "update-ok",
			payload: models.PostUpdateInput{
				Title: "A sample title",
			},
			urlSlug: "example",
			isError: false,
		},
		{
			name:    "update-error-no-post",
			payload: models.PostUpdateInput{},
			urlSlug: repository.ErrNotFoundKeyString,
			isError: true,
		},
		{
			name:    "update-error-getting-post",
			payload: models.PostUpdateInput{},
			urlSlug: repository.ErrUnexpectedKeyString,
			isError: true,
		},
		{
			name: "update-error-duplicate-title",
			payload: models.PostUpdateInput{
				Title: repository.ErrDuplicateKeyString,
			},
			urlSlug: "example",
			isError: true,
		},
		{
			name: "update-error-updating-post",
			payload: models.PostUpdateInput{
				Title: repository.ErrUnexpectedKeyString,
			},
			urlSlug: "example",
			isError: true,
		},
	}

	for _, tt := range tests {
		err := s.Update(tt.payload, tt.urlSlug)

		if err != nil && !tt.isError {
			t.Errorf("%s should not return error, but got %s", tt.name, err)
		}

		if err == nil && tt.isError {
			t.Errorf("%s should return error", tt.name)
		}
	}
}

func TestPostService_Delete(t *testing.T) {
	var tests = []struct {
		name    string
		slug    string
		isError bool
	}{
		{
			name:    "delete-ok",
			slug:    "sample",
			isError: false,
		},
		{
			name:    "delete-error-no-post",
			slug:    repository.ErrNotFoundKeyString,
			isError: true,
		},
		{
			name:    "delete-error-getting-post",
			slug:    repository.ErrUnexpectedKeyString,
			isError: true,
		},
		{
			name:    "delete-error-deleting-post",
			slug:    "get-invalid-post",
			isError: true,
		},
	}

	for _, tt := range tests {
		err := s.Delete(tt.slug)

		if err != nil && !tt.isError {
			t.Errorf("%s should not return error, but got %s", tt.name, err)
		}

		if err == nil && tt.isError {
			t.Errorf("%s should return error", tt.name)
		}
	}
}
