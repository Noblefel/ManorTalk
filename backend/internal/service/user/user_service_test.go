package user

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/Noblefel/ManorTalk/backend/internal/config"
	"github.com/Noblefel/ManorTalk/backend/internal/database"
	"github.com/Noblefel/ManorTalk/backend/internal/models"
	"github.com/Noblefel/ManorTalk/backend/internal/repository"
	"github.com/Noblefel/ManorTalk/backend/internal/repository/postgres"
	"github.com/Noblefel/ManorTalk/backend/internal/repository/redis"
)

func TestNewUserService(t *testing.T) {
	var db *database.DB
	var c *config.AppConfig
	cr := redis.NewRepo(db)
	ur := postgres.NewUserRepo(db)
	service := NewUserService(c, cr, ur)

	typeString := reflect.TypeOf(service).String()
	if typeString != "*user.userService" {
		t.Error("NewUserService() get incorrect type, wanted *user.userService")
	}
}

func TestNewMockUserService(t *testing.T) {
	service := NewMockUserService()

	typeString := reflect.TypeOf(service).String()

	if typeString != "*user.mockUserService" {
		t.Error("NewMockUserService() get incorrect type, wanted *user.mockUserService")
	}
}

func newTestService() UserService {
	var tc config.AppConfig
	cr := redis.NewMockRepo()
	ur := postgres.NewMockUserRepo()

	service := NewUserService(&tc, cr, ur)

	return service
}

var s = newTestService()

func TestUserService_CheckUsername(t *testing.T) {
	var tests = []struct {
		name     string
		username string
		isError  bool
	}{
		{"success", repository.NotFoundKey, false},
		{"duplicate username", "test", true},
		{"error getting user by username", repository.UnexpectedKey, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := s.CheckUsername(tt.username)

			if err != nil && !tt.isError {
				t.Errorf("expecting no error, got: %v", err)
			}

			if err == nil && tt.isError {
				t.Error("expecting error")
			}
		})
	}
}

func TestUserService_Get(t *testing.T) {
	var tests = []struct {
		name     string
		username string
		isError  bool
	}{
		{"success", "test", false},
		{"user not found", repository.NotFoundKey, true},
		{"error getting user by username", repository.UnexpectedKey, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := s.Get(tt.username)

			if err != nil && !tt.isError {
				t.Errorf("expecting no error, got: %v", err)
			}

			if err == nil && tt.isError {
				t.Error("expecting error")
			}
		})
	}
}

func TestUserService_UpdateProfile(t *testing.T) {
	var tests = []struct {
		name     string
		payload  models.UpdateProfileInput
		username string
		authId   int
		isError  bool
	}{
		{
			name: "success",
			payload: models.UpdateProfileInput{
				Name: "test",
			},
		},
		{
			name:     "user not found",
			username: repository.NotFoundKey,
			isError:  true,
		},
		{
			name:     "error getting user",
			username: repository.UnexpectedKey,
			isError:  true,
		},
		{
			name:    "unauthorized",
			authId:  -1,
			isError: true,
		},
		{
			name:    "avatar invalid type",
			payload: models.UpdateProfileInput{Avatar: bytes.NewReader(make([]byte, 1))},
			isError: true,
		},
		{
			name:    "avatar too large",
			payload: models.UpdateProfileInput{Avatar: bytes.NewReader(make([]byte, 2*1024*1024+2))},
			isError: true,
		},
		{
			name:    "error verifying image",
			payload: models.UpdateProfileInput{Avatar: &bytes.Reader{}},
			isError: true,
		},
		{
			name:    "duplicate username",
			payload: models.UpdateProfileInput{Name: repository.DuplicateKey},
			isError: true,
		},
		{
			name:    "error updating user",
			payload: models.UpdateProfileInput{Name: repository.UnexpectedKey},
			isError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := s.UpdateProfile(tt.payload, tt.username, tt.authId)

			if err != nil && !tt.isError {
				t.Errorf("expecting no error, got: %v", err)
			}

			if err == nil && tt.isError {
				t.Errorf("expecting error")
			}
		})
	}
}
