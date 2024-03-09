package post

import (
	"bytes"
	"io"
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
		t.Error("NewPostService() get incorrect type, wanted *post.PostService")
	}
}

func TestNewMockPostService(t *testing.T) {
	service := NewMockPostService()

	typeString := reflect.TypeOf(service).String()

	if typeString != "*post.mockPostService" {
		t.Error("NewMockPostService() get incorrect type, wanted *post.mockPostService")
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
		name       string
		title      string
		categoryId int
		image      io.ReadSeeker
		isError    bool
	}{
		{"success", "", 1, nil, false},
		{"no category", "", repository.NotFoundKeyInt, nil, true},
		{"error getting category", "", repository.UnexpectedKeyInt, nil, true},
		{"image invalid type", "", 0, bytes.NewReader(make([]byte, 1)), true},
		{"image too large", "", 0, bytes.NewReader(make([]byte, 2*1024*1024+2)), true},
		{"error verifying image", "", 0, &bytes.Reader{}, true},
		{"duplicate title", repository.DuplicateKey, 1, nil, true},
		{"error creating post", repository.UnexpectedKey, 1, nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			payload := models.PostCreateInput{
				Title:      tt.title,
				CategoryId: tt.categoryId,
				Image:      tt.image,
			}
			_, err := s.Create(payload, 1)

			if err != nil && !tt.isError {
				t.Errorf("expecting no error, got %v", err)
			}

			if err == nil && tt.isError {
				t.Error("expecting error")
			}
		})
	}
}

func TestPostService_Get(t *testing.T) {
	var tests = []struct {
		name    string
		slug    string
		isError bool
	}{
		{"success", "example", false},
		{"no post", repository.NotFoundKey, true},
		{"error getting post", repository.UnexpectedKey, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := s.Get(tt.slug)

			if err != nil && !tt.isError {
				t.Errorf("expecting no error, got %v", err)
			}

			if err == nil && tt.isError {
				t.Error("expecting error")
			}
		})
	}
}

func TestPostService_GetMany(t *testing.T) {
	var tests = []struct {
		name    string
		q       url.Values
		isError bool
	}{
		{"success", url.Values{"page": {"1"}, "total": {"10"}}, false},
		{"no category", url.Values{"category": {repository.NotFoundKey}}, true},
		{"error getting category", url.Values{"category": {repository.UnexpectedKey}}, true},
		{"error creating pagination meta", url.Values{"page": {"-1"}}, true},
		{"error counting posts", url.Values{"page": {"1"}, "order": {repository.UnexpectedKey}}, true},
		{"error getting posts", url.Values{"total": {"1"}, "order": {repository.UnexpectedKey}}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, _, err := s.GetMany(tt.q)

			if err != nil && !tt.isError {
				t.Errorf("expecting no error, got %v", err)
			}

			if err == nil && tt.isError {
				t.Error("expecting error")
			}
		})
	}
}

func TestPostService_Update(t *testing.T) {
	var tests = []struct {
		name       string
		title      string
		categoryId int
		image      io.ReadSeeker
		urlSlug    string
		authId     int
		isError    bool
	}{
		{"success", "", 0, nil, "", 0, false},
		{"no post", "", 0, nil, repository.NotFoundKey, 0, true},
		{"error getting post", "", 0, nil, repository.UnexpectedKey, 0, true},
		{"unauthorized", "", 0, nil, "", -1, true},
		{"no category", "", repository.NotFoundKeyInt, nil, "", 0, true},
		{"error getting category", "", repository.UnexpectedKeyInt, nil, "", 0, true},
		{"image invalid type", "", 0, bytes.NewReader(make([]byte, 1)), "", 0, true},
		{"image too large", "", 0, bytes.NewReader(make([]byte, 2*1024*1024+2)), "", 0, true},
		{"error verifying image", "", 0, &bytes.Reader{}, "", 0, true},
		{"duplicate title", repository.DuplicateKey, 0, nil, "", 0, true},
		{"error updating post", repository.UnexpectedKey, 0, nil, "", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			payload := models.PostUpdateInput{
				Title:      tt.title,
				CategoryId: tt.categoryId,
				Image:      tt.image,
			}
			err := s.Update(payload, tt.urlSlug, tt.authId)

			if err != nil && !tt.isError {
				t.Errorf("expecting no error, got %v", err)
			}

			if err == nil && tt.isError {
				t.Error("expecting error")
			}
		})
	}
}

func TestPostService_Delete(t *testing.T) {
	var tests = []struct {
		name    string
		slug    string
		authId  int
		isError bool
	}{
		{"success", "sample", 0, false},
		{"post not found", repository.NotFoundKey, 0, true},
		{"error getting post", repository.UnexpectedKey, 0, true},
		{"unauthorized", "", -1, true},
		{"error deleting post", "get-invalid-post", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := s.Delete(tt.slug, tt.authId)

			if err != nil && !tt.isError {
				t.Errorf("expecting no error, got %v", err)
			}

			if err == nil && tt.isError {
				t.Error("expecting error")
			}
		})
	}
}
