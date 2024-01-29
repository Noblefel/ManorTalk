package handlers

import (
	"bytes"
	"encoding/json"
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
			name: "authRegister-ok",
			payload: &models.UserRegisterInput{
				Email:    "test@example.com",
				Password: "password123",
			},
			statusCode: http.StatusOK,
		},
		{
			name:       "authRegister-error-decode-json",
			payload:    nil,
			statusCode: http.StatusBadRequest,
		},
		{
			name: "authRegister-error-validation",
			payload: &models.UserRegisterInput{
				Email:    "not-an-email",
				Password: "",
			},
			statusCode: http.StatusBadRequest,
		},
		{
			name: "authRegister-error-duplicate-email",
			payload: &models.UserRegisterInput{
				Email:    "test@example.com",
				Password: service.ErrDuplicateEmail.Error(),
			},
			statusCode: http.StatusConflict,
		},
		{
			name: "authRegister-error-unexpected",
			payload: &models.UserRegisterInput{
				Email:    "test@example.com",
				Password: http.StatusText(http.StatusInternalServerError),
			},
			statusCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		var r *http.Request
		if tt.payload == nil {
			r = httptest.NewRequest("POST", "/auth/register", nil)
		} else {
			jsonBytes, _ := json.Marshal(tt.payload)
			r = httptest.NewRequest("POST", "/auth/register", bytes.NewBuffer(jsonBytes))
		}

		r.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		handler := http.HandlerFunc(h.auth.Register)
		handler.ServeHTTP(w, r)

		if w.Code != tt.statusCode {
			t.Errorf("%s returned response code of %d, wanted %d", tt.name, w.Code, tt.statusCode)
		}
	}
}

func TestAuth_Login(t *testing.T) {
	var tests = []struct {
		name       string
		payload    *models.UserLoginInput
		statusCode int
	}{
		{
			name: "authLogin-ok",
			payload: &models.UserLoginInput{
				Email:    "test@example.com",
				Password: "password",
			},
			statusCode: http.StatusOK,
		},
		{
			name:       "authLogin-error-decode-json",
			payload:    nil,
			statusCode: http.StatusBadRequest,
		},
		{
			name: "authLogin-error-validation",
			payload: &models.UserLoginInput{
				Email:    "not-an-email",
				Password: "",
			},
			statusCode: http.StatusBadRequest,
		},
		{
			name: "authLogin-error-invalid-credentials",
			payload: &models.UserLoginInput{
				Email:    "test@example.com",
				Password: service.ErrInvalidCredentials.Error(),
			},
			statusCode: http.StatusUnauthorized,
		},
		{
			name: "authLogin-error-no-user",
			payload: &models.UserLoginInput{
				Email:    "test@example.com",
				Password: service.ErrNoUser.Error(),
			},
			statusCode: http.StatusUnauthorized,
		},
		{
			name: "authLogin-error-unexpected",
			payload: &models.UserLoginInput{
				Email:    "test@example.com",
				Password: http.StatusText(http.StatusInternalServerError),
			},
			statusCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		var r *http.Request
		if tt.payload == nil {
			r = httptest.NewRequest("POST", "/auth/login", nil)
		} else {
			jsonBytes, _ := json.Marshal(tt.payload)
			r = httptest.NewRequest("POST", "/auth/login", bytes.NewBuffer(jsonBytes))
		}

		r.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		handler := http.HandlerFunc(h.auth.Login)
		handler.ServeHTTP(w, r)

		if w.Code != tt.statusCode {
			t.Errorf("%s returned response code of %d, wanted %d", tt.name, w.Code, tt.statusCode)
		}
	}
}

func TestAuth_Refresh(t *testing.T) {
	var tests = []struct {
		name       string
		cookie     *http.Cookie
		statusCode int
	}{
		{
			name: "authRefresh-ok",
			cookie: &http.Cookie{
				Name:  "refresh_token",
				Value: "refresh_token",
			},
			statusCode: http.StatusOK,
		},
		{
			name:       "authRefresh-error-missing-cookie",
			cookie:     nil,
			statusCode: http.StatusUnauthorized,
		},
		{
			name: "authRefresh-error-unauthorized",
			cookie: &http.Cookie{
				Name:  "refresh_token",
				Value: service.ErrUnauthorized.Error(),
			},
			statusCode: http.StatusUnauthorized,
		},
		{
			name: "authRefresh-error-no-user",
			cookie: &http.Cookie{
				Name:  "refresh_token",
				Value: service.ErrNoUser.Error(),
			},
			statusCode: http.StatusUnauthorized,
		},
		{
			name: "authRefresh-error-unexpected",
			cookie: &http.Cookie{
				Name:  "refresh_token",
				Value: http.StatusText(http.StatusInternalServerError),
			},
			statusCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		r := httptest.NewRequest("POST", "/auth/refresh", nil)
		r.Header.Set("Content-Type", "application/json")

		if tt.cookie != nil {
			r.AddCookie(tt.cookie)
		}

		w := httptest.NewRecorder()
		handler := http.HandlerFunc(h.auth.Refresh)
		handler.ServeHTTP(w, r)

		if w.Code != tt.statusCode {
			t.Errorf("%s returned response code of %d, wanted %d", tt.name, w.Code, tt.statusCode)
		}
	}
}
