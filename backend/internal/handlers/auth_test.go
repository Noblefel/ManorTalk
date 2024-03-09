package handlers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/Noblefel/ManorTalk/backend/internal/config"
	"github.com/Noblefel/ManorTalk/backend/internal/database"
	"github.com/Noblefel/ManorTalk/backend/internal/models"
	"github.com/Noblefel/ManorTalk/backend/internal/repository/postgres"
	"github.com/Noblefel/ManorTalk/backend/internal/repository/redis"
	service "github.com/Noblefel/ManorTalk/backend/internal/service/auth"
)

func TestNewAuthHandlers(t *testing.T) {
	var db *database.DB
	var c *config.AppConfig
	cr := redis.NewRepo(db)
	ur := postgres.NewUserRepo(db)
	s := service.NewAuthService(c, cr, ur)
	auth := NewAuthHandlers(s)

	typeString := reflect.TypeOf(auth).String()

	if typeString != "*handlers.AuthHandlers" {
		t.Error("NewAuthHandlers() did not get the correct type, wanted *handlers.AuthHandlers")
	}
}

func TestAuth_Register(t *testing.T) {
	var tests = []struct {
		name       string
		payload    *models.UserRegisterInput
		statusCode int
	}{
		{
			name:       "success",
			payload:    &models.UserRegisterInput{"test", "test@example.com", "password123"},
			statusCode: http.StatusOK,
		},
		{"error decoding json", nil, http.StatusBadRequest},
		{
			name:       "error validation",
			payload:    &models.UserRegisterInput{"test", "not-an-email", ""},
			statusCode: http.StatusBadRequest,
		},
		{
			name:       "duplicate email",
			payload:    &models.UserRegisterInput{"test", "test@example.com", service.ErrDuplicateEmail.Error()},
			statusCode: http.StatusConflict,
		},
		{
			name:       "unexpected error",
			payload:    &models.UserRegisterInput{"test", "test@example.com", "unexpected error"},
			statusCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var body io.Reader
			if tt.payload != nil {
				b, _ := json.Marshal(tt.payload)
				body = bytes.NewBuffer(b)
			}

			r := httptest.NewRequest("POST", "/auth/register", body)
			r.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			handler := http.HandlerFunc(h.auth.Register)
			handler.ServeHTTP(w, r)

			if w.Code != tt.statusCode {
				t.Errorf("want %d, got %d", tt.statusCode, w.Code)
			}
		})
	}
}

func TestAuth_Login(t *testing.T) {
	var tests = []struct {
		name       string
		email      string
		password   string
		statusCode int
	}{
		{"success", "test@example.com", "password", http.StatusOK},
		{"error decoding json", "", "", http.StatusBadRequest},
		{"error validation", "not-an-email", "", http.StatusBadRequest},
		{"invalid credentials", "test@example.com", service.ErrInvalidCredentials.Error(), http.StatusUnauthorized},
		{"no user", "test@example.com", service.ErrNoUser.Error(), http.StatusUnauthorized},
		{"unexpected error", "test@example.com", "unexpected error", http.StatusInternalServerError},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var body io.Reader
			if tt.email != "" {
				b, _ := json.Marshal(models.UserLoginInput{Email: tt.email, Password: tt.password})
				body = bytes.NewBuffer(b)
			}

			r := httptest.NewRequest("POST", "/auth/login", body)
			r.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			handler := http.HandlerFunc(h.auth.Login)
			handler.ServeHTTP(w, r)

			if w.Code != tt.statusCode {
				t.Errorf("want %d, got %d", tt.statusCode, w.Code)
			}
		})
	}
}

func TestAuth_Refresh(t *testing.T) {
	var tests = []struct {
		name       string
		cookie     *http.Cookie
		statusCode int
	}{
		{"success", &http.Cookie{Name: "refresh_token", Value: "refresh_token"}, http.StatusOK},
		{"missing cookie", nil, http.StatusUnauthorized},
		{"unauthorized", &http.Cookie{Name: "refresh_token", Value: service.ErrUnauthorized.Error()}, http.StatusUnauthorized},
		{"no user", &http.Cookie{Name: "refresh_token", Value: service.ErrNoUser.Error()}, http.StatusUnauthorized},
		{"unexpected error", &http.Cookie{Name: "refresh_token", Value: "unexpected error"}, http.StatusInternalServerError},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := httptest.NewRequest("POST", "/auth/refresh", nil)
			r.Header.Set("Content-Type", "application/json")

			if tt.cookie != nil {
				r.AddCookie(tt.cookie)
			}

			w := httptest.NewRecorder()
			handler := http.HandlerFunc(h.auth.Refresh)
			handler.ServeHTTP(w, r)

			if w.Code != tt.statusCode {
				t.Errorf("want %d, got %d", tt.statusCode, w.Code)
			}
		})
	}
}

func TestAuth_Logout(t *testing.T) {
	var tests = []struct {
		name       string
		cookie     *http.Cookie
		statusCode int
	}{
		{"success", &http.Cookie{Name: "refresh_token", Value: "refresh_token"}, http.StatusOK},
		{"missing cookie", nil, http.StatusUnauthorized},
		{"unauthorized", &http.Cookie{Name: "refresh_token", Value: service.ErrUnauthorized.Error()}, http.StatusUnauthorized},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := httptest.NewRequest("POST", "/auth/logout", nil)
			r.Header.Set("Content-Type", "application/json")

			if tt.cookie != nil {
				r.AddCookie(tt.cookie)
			}

			w := httptest.NewRecorder()
			handler := http.HandlerFunc(h.auth.Logout)
			handler.ServeHTTP(w, r)

			if w.Code != tt.statusCode {
				t.Errorf("want %d, got %d", tt.statusCode, w.Code)
			}
		})
	}
}
