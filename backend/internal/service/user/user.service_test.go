package user

import (
	"reflect"
	"testing"

	"github.com/Noblefel/ManorTalk/backend/internal/config"
	"github.com/Noblefel/ManorTalk/backend/internal/database"
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
		t.Error("NewUserService() did not get the correct type, wanted *user.userService")
	}
}

func TestNewMockUserService(t *testing.T) {
	service := NewMockUserService()

	typeString := reflect.TypeOf(service).String()

	if typeString != "*user.mockUserService" {
		t.Error("NewMockUserService() did not get the correct type, wanted *user.mockUserService")
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
		{
			name:     "checkUsername-ok",
			username: repository.ErrNotFoundKeyString,
			isError:  false,
		},
		{
			name:     "checkUsername-error-duplicate-username",
			username: "test",
			isError:  true,
		},
		{
			name:     "checkUsername-error-getting-post-by-username",
			username: repository.ErrUnexpectedKeyString,
			isError:  true,
		},
	}

	for _, tt := range tests {
		err := s.CheckUsername(tt.username)

		if err != nil && !tt.isError {
			t.Errorf("%s should not return error, but got %s", tt.name, err)
		}

		if err == nil && tt.isError {
			t.Errorf("%s should return error", tt.name)
		}
	}
}

func TestUserService_Get(t *testing.T) {
	var tests = []struct {
		name     string
		username string
		isError  bool
	}{
		{
			name:     "get-ok",
			username: "test",
			isError:  false,
		},
		{
			name:     "get-error-not-found",
			username: repository.ErrNotFoundKeyString,
			isError:  true,
		},
		{
			name:     "get-error-getting-post-by-username",
			username: repository.ErrUnexpectedKeyString,
			isError:  true,
		},
	}

	for _, tt := range tests {
		_, err := s.Get(tt.username)

		if err != nil && !tt.isError {
			t.Errorf("%s should not return error, but got %s", tt.name, err)
		}

		if err == nil && tt.isError {
			t.Errorf("%s should return error", tt.name)
		}
	}
}
