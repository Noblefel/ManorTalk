package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"github.com/Noblefel/ManorTalk/backend/internal/config"
	"github.com/Noblefel/ManorTalk/backend/internal/database"
	"github.com/Noblefel/ManorTalk/backend/internal/models"
	"github.com/Noblefel/ManorTalk/backend/internal/utils/token"
)

func TestNewAuthHandlers(t *testing.T) {
	var db *database.DB
	var c *config.AppConfig
	auth := NewAuthHandlers(c, db)

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
				Email:    "alreadyexists@error.com",
				Password: "password123",
			},
			statusCode: http.StatusConflict,
		},
		{
			name: "authRegister-error-creating-user",
			payload: &models.UserRegisterInput{
				Email:    "unexpected@error.com",
				Password: "password123",
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
				Password: "incorrectpassword",
			},
			statusCode: http.StatusUnauthorized,
		},
		{
			name: "authLogin-error-user-not-found",
			payload: &models.UserLoginInput{
				Email:    "notfound@error.com",
				Password: "password",
			},
			statusCode: http.StatusUnauthorized,
		},
		{
			name: "authLogin-error-authenticating",
			payload: &models.UserLoginInput{
				Email:    "unexpected@error.com",
				Password: "password",
			},
			statusCode: http.StatusInternalServerError,
		},
		{
			name: "authLogin-error-saving-token",
			payload: &models.UserLoginInput{
				Email:    "invaliduser@example.com",
				Password: "password",
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

// Sample refresh tokens
var refreshToken, _ = token.Generate(token.Details{
	UserId:    1,
	UniqueId:  "uuid",
	SecretKey: h.auth.c.RefreshTokenKey,
	Duration:  1 * time.Minute,
})

var refreshTokenInvalid, _ = token.Generate(token.Details{
	UserId:   1,
	UniqueId: "uuid",
})

var refreshTokenInvalid2, _ = token.Generate(token.Details{
	UserId:    -1,
	UniqueId:  "incorrect",
	SecretKey: h.auth.c.RefreshTokenKey,
	Duration:  1 * time.Minute,
})

var refreshTokenUserNotFound, _ = token.Generate(token.Details{
	UserId:    9999999,
	UniqueId:  "uuid",
	SecretKey: h.auth.c.RefreshTokenKey,
	Duration:  1 * time.Minute,
})

var refreshTokenUserUnexpectedError, _ = token.Generate(token.Details{
	UserId:    -1,
	UniqueId:  "uuid",
	SecretKey: h.auth.c.RefreshTokenKey,
	Duration:  1 * time.Minute,
})

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
				Value: refreshToken,
			},
			statusCode: http.StatusOK,
		},
		{
			name:       "authRefresh-error-missing-cookie",
			cookie:     nil,
			statusCode: http.StatusUnauthorized,
		},
		{
			name: "authRefresh-error-parsing-token",
			cookie: &http.Cookie{
				Name:  "refresh_token",
				Value: refreshTokenInvalid,
			},
			statusCode: http.StatusUnauthorized,
		},
		{
			name: "authRefresh-error-getting-token-redis",
			cookie: &http.Cookie{
				Name:  "refresh_token",
				Value: refreshTokenInvalid2,
			},
			statusCode: http.StatusUnauthorized,
		},
		{
			name: "authRefresh-error-user-not-found",
			cookie: &http.Cookie{
				Name:  "refresh_token",
				Value: refreshTokenUserNotFound,
			},
			statusCode: http.StatusNotFound,
		},
		{
			name: "authRefresh-error-getting-user",
			cookie: &http.Cookie{
				Name:  "refresh_token",
				Value: refreshTokenUserUnexpectedError,
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
