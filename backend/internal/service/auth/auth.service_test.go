package auth

import (
	"reflect"
	"testing"
	"time"

	"github.com/Noblefel/ManorTalk/backend/internal/config"
	"github.com/Noblefel/ManorTalk/backend/internal/database"
	"github.com/Noblefel/ManorTalk/backend/internal/models"
	"github.com/Noblefel/ManorTalk/backend/internal/repository"
	"github.com/Noblefel/ManorTalk/backend/internal/repository/postgres"
	"github.com/Noblefel/ManorTalk/backend/internal/repository/redis"
	"github.com/Noblefel/ManorTalk/backend/internal/utils/token"
)

func TestNewAuthService(t *testing.T) {
	var db *database.DB
	var c *config.AppConfig
	cr := redis.NewRepo(db)
	ur := postgres.NewUserRepo(db)
	service := NewAuthService(c, cr, ur)

	typeString := reflect.TypeOf(service).String()

	if typeString != "*auth.authService" {
		t.Error("NewAuthService() did not get the correct type, wanted *auth.authService")
	}
}

func TestNewMockAuthService(t *testing.T) {
	service := NewMockAuthService()

	typeString := reflect.TypeOf(service).String()

	if typeString != "*auth.mockAuthService" {
		t.Error("NewMockAuthService() did not get the correct type, wanted *auth.mockAuthService")
	}
}

var tc = config.AppConfig{
	AccessTokenKey:  "access_key",
	AccessTokenExp:  1 * time.Minute,
	RefreshTokenKey: "refresh_key",
	RefreshTokenExp: 1 * time.Minute,
}

func newTestService() AuthService {
	cr := redis.NewMockRepo()
	ur := postgres.NewMockUserRepo()

	service := NewAuthService(&tc, cr, ur)

	return service
}

var s = newTestService()

func TestAuthService_Register(t *testing.T) {
	var tests = []struct {
		name    string
		payload models.UserRegisterInput
		isError bool
	}{
		{
			name: "register-ok",
			payload: models.UserRegisterInput{
				Email:    "test@example.com",
				Password: "password",
			},
			isError: false,
		},
		{
			name: "register-error-generate-password",
			payload: models.UserRegisterInput{
				Email: "test@example.com",
				Password: `apcsjdpcojpwejrcpajsdpjascjpjpewjcjpojacpsjcdpsjdcpa
				ojsdpcjasdcpaocsjdpcajprjcqwpojaposjdpasjdpcajsdcpaodcspajdpcaj
				spcsd`,
			},
			isError: true,
		},
		{
			name: "register-error-duplicate-key",
			payload: models.UserRegisterInput{
				Email:    repository.ErrDuplicateKeyString,
				Password: "password",
			},
			isError: true,
		},
		{
			name: "register-error-creating-user",
			payload: models.UserRegisterInput{
				Email:    repository.ErrUnexpectedKeyString,
				Password: "password",
			},
			isError: true,
		},
	}

	for _, tt := range tests {
		err := s.Register(tt.payload)

		if err != nil && !tt.isError {
			t.Errorf("%s should not return error, but got %s", tt.name, err)
		}

		if err == nil && tt.isError {
			t.Errorf("%s should return error", tt.name)
		}
	}
}

func TestAuthService_Login(t *testing.T) {
	var tests = []struct {
		name    string
		payload models.UserLoginInput
		isError bool
	}{
		{
			name: "login-ok",
			payload: models.UserLoginInput{
				Email:    "test@example.com",
				Password: "password",
			},
			isError: false,
		},
		{
			name: "login-error-no-user",
			payload: models.UserLoginInput{
				Email: repository.ErrNotFoundKeyString,
			},
			isError: true,
		},
		{
			name: "login-error-getting-user",
			payload: models.UserLoginInput{
				Email: repository.ErrUnexpectedKeyString,
			},
			isError: true,
		},
		{
			name: "login-error-invalid-credentials",
			payload: models.UserLoginInput{
				Password: "x",
			},
			isError: true,
		},
		{
			name: "login-error-setting-refresh-token",
			payload: models.UserLoginInput{
				Email:    "get-invalid-user",
				Password: "password",
			},
			isError: true,
		},
	}

	for _, tt := range tests {
		_, _, _, err := s.Login(tt.payload)

		if err != nil && !tt.isError {
			t.Errorf("%s should not return error, but got %s", tt.name, err)
		}

		if err == nil && tt.isError {
			t.Errorf("%s should return error", tt.name)
		}
	}
}

func TestAuthService_Refresh(t *testing.T) {
	// Sample refresh tokens
	var refreshToken, _ = token.Generate(token.Details{
		UserId:    1,
		UniqueId:  "uuid",
		SecretKey: tc.RefreshTokenKey,
		Duration:  1 * time.Minute,
	})

	var refreshTokenInvalid, _ = token.Generate(token.Details{
		UserId:   1,
		UniqueId: "uuid",
	})

	var refreshTokenInvalid2, _ = token.Generate(token.Details{
		UserId:    -1,
		UniqueId:  repository.ErrIncorrectKey,
		SecretKey: tc.RefreshTokenKey,
		Duration:  1 * time.Minute,
	})

	var refreshTokenUserNotFound, _ = token.Generate(token.Details{
		UserId:    repository.ErrNotFoundKeyInt,
		UniqueId:  "uuid",
		SecretKey: tc.RefreshTokenKey,
		Duration:  1 * time.Minute,
	})

	var refreshTokenUserUnexpectedError, _ = token.Generate(token.Details{
		UserId:    repository.ErrUnexpectedKeyInt,
		UniqueId:  "uuid",
		SecretKey: tc.RefreshTokenKey,
		Duration:  1 * time.Minute,
	})

	var tests = []struct {
		name         string
		refreshToken string
		isError      bool
	}{
		{
			name:         "refresh-ok",
			refreshToken: refreshToken,
			isError:      false,
		},
		{
			name:         "refresh-error-parsing-token",
			refreshToken: refreshTokenInvalid,
			isError:      true,
		},
		{
			name:         "refresh-error-getting-refresh-token",
			refreshToken: refreshTokenInvalid2,
			isError:      true,
		},
		{
			name:         "refresh-error-parsing-token",
			refreshToken: refreshTokenInvalid,
			isError:      true,
		},
		{
			name:         "refresh-error-no-user",
			refreshToken: refreshTokenUserNotFound,
			isError:      true,
		},
		{
			name:         "refresh-error-getting-user",
			refreshToken: refreshTokenUserUnexpectedError,
			isError:      true,
		},
	}

	for _, tt := range tests {
		_, _, err := s.Refresh(tt.refreshToken)

		if err != nil && !tt.isError {
			t.Errorf("%s should not return error, but got %s", tt.name, err)
		}

		if err == nil && tt.isError {
			t.Errorf("%s should return error", tt.name)
		}
	}
}
