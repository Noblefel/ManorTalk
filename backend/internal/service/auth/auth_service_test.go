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
		t.Error("NewAuthService() get incorrect type, wanted *auth.authService")
	}
}

func TestNewMockAuthService(t *testing.T) {
	service := NewMockAuthService()

	typeString := reflect.TypeOf(service).String()

	if typeString != "*auth.mockAuthService" {
		t.Error("NewMockAuthService() get incorrect type, wanted *auth.mockAuthService")
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
		name     string
		email    string
		password string
		isError  bool
	}{
		{"success", "test@example.com", "password", false},
		{"error generate password", "test@example.com",
			`apcsjdpcojpwejrcpajsdpjascjpjpewjcjpojacpsjcdpsjdcpa
			ojsdpcjasdcpaocsjdpcajprjcqwpojaposjdpasjdpcajsdcpaodcspajdpcaj
			spcsd`, true,
		},
		{"duplicate key", repository.DuplicateKey, "password", true},
		{"error creating user", repository.UnexpectedKey, "password", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := models.UserRegisterInput{Email: tt.email, Password: tt.password}
			err := s.Register(p)

			if err != nil && !tt.isError {
				t.Errorf("expecting no error, got %v", err)
			}

			if err == nil && tt.isError {
				t.Error("expecting error")
			}
		})
	}
}

func TestAuthService_Login(t *testing.T) {
	var tests = []struct {
		name     string
		email    string
		password string
		isError  bool
	}{
		{"success", "test@example.com", "password", false},
		{"no user", repository.NotFoundKey, "", true},
		{"error getting user", repository.UnexpectedKey, "", true},
		{"invalid credentials", "", "x", true},
		{"error setting refresh token", "get-invalid-user", "password", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := models.UserLoginInput{Email: tt.email, Password: tt.password}
			_, _, _, err := s.Login(p)

			if err != nil && !tt.isError {
				t.Errorf("expecting no error, got %v", err)
			}

			if err == nil && tt.isError {
				t.Error("expecting error")
			}
		})
	}
}

func TestAuthService_Refresh(t *testing.T) {
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
		UniqueId:  repository.IncorrectKey,
		SecretKey: tc.RefreshTokenKey,
		Duration:  1 * time.Minute,
	})

	var refreshTokenUserNotFound, _ = token.Generate(token.Details{
		UserId:    repository.NotFoundKeyInt,
		UniqueId:  "uuid",
		SecretKey: tc.RefreshTokenKey,
		Duration:  1 * time.Minute,
	})

	var refreshTokenUserUnexpectedError, _ = token.Generate(token.Details{
		UserId:    repository.UnexpectedKeyInt,
		UniqueId:  "uuid",
		SecretKey: tc.RefreshTokenKey,
		Duration:  1 * time.Minute,
	})

	var tests = []struct {
		name         string
		refreshToken string
		isError      bool
	}{
		{"success", refreshToken, false},
		{"error parsing token", refreshTokenInvalid, true},
		{"error getting refresh token", refreshTokenInvalid2, true},
		{"no user", refreshTokenUserNotFound, true},
		{"error getting user", refreshTokenUserUnexpectedError, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, _, err := s.Refresh(tt.refreshToken)

			if err != nil && !tt.isError {
				t.Errorf("expecting no error, got %v", err)
			}

			if err == nil && tt.isError {
				t.Error("expecting error")
			}
		})
	}
}

func TestAuthService_Logout(t *testing.T) {
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
		UniqueId:  repository.IncorrectKey,
		SecretKey: tc.RefreshTokenKey,
		Duration:  1 * time.Minute,
	})

	var refreshTokenUnexpectedError, _ = token.Generate(token.Details{
		UserId:    repository.UnexpectedKeyInt,
		UniqueId:  "uuid",
		SecretKey: tc.RefreshTokenKey,
		Duration:  1 * time.Minute,
	})

	var tests = []struct {
		name         string
		refreshToken string
		isError      bool
	}{
		{"success", refreshToken, false},
		{"error parsing token", refreshTokenInvalid, true},
		{"error getting refresh token", refreshTokenInvalid2, true},
		{"error deleting refresh token", refreshTokenUnexpectedError, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := s.Logout(tt.refreshToken)

			if err != nil && !tt.isError {
				t.Errorf("expecting no error, got %v", err)
			}

			if err == nil && tt.isError {
				t.Error("expecting error")
			}
		})
	}
}
