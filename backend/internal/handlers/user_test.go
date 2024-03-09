package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/Noblefel/ManorTalk/backend/internal/config"
	"github.com/Noblefel/ManorTalk/backend/internal/database"
	"github.com/Noblefel/ManorTalk/backend/internal/models"
	"github.com/Noblefel/ManorTalk/backend/internal/repository/postgres"
	"github.com/Noblefel/ManorTalk/backend/internal/repository/redis"
	service "github.com/Noblefel/ManorTalk/backend/internal/service/user"
)

func TestNewUserHandlers(t *testing.T) {
	var db *database.DB
	var c *config.AppConfig
	cr := redis.NewRepo(db)
	ur := postgres.NewUserRepo(db)
	s := service.NewUserService(c, cr, ur)
	user := NewUserHandlers(s)

	typeString := reflect.TypeOf(user).String()

	if typeString != "*handlers.UserHandlers" {
		t.Error("NewUserHandlers() did not get the correct type, wanted *handlers.UserHandlers")
	}
}

func TestUser_CheckUsername(t *testing.T) {
	var tests = []struct {
		name       string
		username   string
		statusCode int
	}{
		{"success", "test", http.StatusOK},
		{"error decoding json", "", http.StatusBadRequest},
		{"error validation", "t", http.StatusBadRequest},
		{"duplicate username", service.ErrDuplicateUsername.Error(), http.StatusConflict},
		{"unexpected error", "unexpected error", http.StatusInternalServerError},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var body io.Reader
			if tt.username != "" {
				jsonBytes, _ := json.Marshal(models.CheckUsernameInput{Username: tt.username})
				body = bytes.NewReader(jsonBytes)
			}

			r := httptest.NewRequest("POST", "/users/check-username", body)
			r.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			handler := http.HandlerFunc(h.user.CheckUsername)
			handler.ServeHTTP(w, r)

			if w.Code != tt.statusCode {
				t.Errorf("want %d, got %d", tt.statusCode, w.Code)
			}
		})
	}
}

func TestUser_Get(t *testing.T) {
	var tests = []struct {
		name       string
		username   string
		statusCode int
	}{
		{"success", "example", http.StatusOK},
		{"no user", service.ErrNoUser.Error(), http.StatusNotFound},
		{"unexpected error", "unexpected error", http.StatusInternalServerError},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := httptest.NewRequest("GET", "/users/{username}", nil)
			ctx := getCtxWithParam(r, params{"username": tt.username})
			r = r.WithContext(ctx)
			r.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			handler := http.HandlerFunc(h.user.Get)
			handler.ServeHTTP(w, r)

			if w.Code != tt.statusCode {
				t.Errorf("want %d, got %d", tt.statusCode, w.Code)
			}
		})
	}
}

func TestUser_UpdateProfile(t *testing.T) {
	var tests = []struct {
		name           string
		noForm         bool
		failValidation bool
		username       string
		statusCode     int
	}{
		{"success", false, false, "test", http.StatusOK},
		{"error parsing form", true, false, "", http.StatusBadRequest},
		{"error validation", false, true, "", http.StatusBadRequest},
		{"no user", false, false, service.ErrNoUser.Error(), http.StatusNotFound},
		{"unauthorized", false, false, service.ErrUnauthorized.Error(), http.StatusUnauthorized},
		{"avatar invalid type", false, false, service.ErrAvatarInvalid.Error(), http.StatusBadRequest},
		{"avatar too large", false, false, service.ErrAvatarTooLarge.Error(), http.StatusBadRequest},
		{"duplicate username", false, false, service.ErrDuplicateUsername.Error(), http.StatusConflict},
		{"unexpected error", false, false, "unexpected error", http.StatusInternalServerError},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var b bytes.Buffer
			fw := multipart.NewWriter(&b)
			v := "test"
			if tt.failValidation {
				v = ""
			}
			fw.WriteField("username", v)
			fw.CreateFormFile("avatar", "x")
			fw.Close()

			var r *http.Request
			if tt.noForm {
				r = httptest.NewRequest("PATCH", "/users/{username}", nil)
			} else {
				r = httptest.NewRequest("PATCH", "/users/{username}", &b)
			}

			ctx := getCtxWithParam(r, params{"username": tt.username})
			ctx = context.WithValue(ctx, "user_id", 1)
			r = r.WithContext(ctx)
			r.Header.Set("Content-Type", fw.FormDataContentType())
			w := httptest.NewRecorder()
			handler := http.HandlerFunc(h.user.UpdateProfile)
			handler.ServeHTTP(w, r)

			if w.Code != tt.statusCode {
				t.Errorf("want %d, got %d", tt.statusCode, w.Code)
			}
		})
	}
}
