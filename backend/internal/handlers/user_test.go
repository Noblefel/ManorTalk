package handlers

import (
	"bytes"
	"context"
	"encoding/json"
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
		payload    *models.CheckUsernameInput
		statusCode int
	}{
		{
			name:       "userCheckUsername-ok",
			payload:    &models.CheckUsernameInput{Username: "test"},
			statusCode: http.StatusOK,
		},
		{
			name:       "userCheckUsername-error-decode-json",
			payload:    nil,
			statusCode: http.StatusBadRequest,
		},
		{
			name:       "userCheckUsername-error-validation",
			payload:    &models.CheckUsernameInput{Username: "t"},
			statusCode: http.StatusBadRequest,
		},
		{
			name:       "userCheckUsername-error-duplicate-username",
			payload:    &models.CheckUsernameInput{Username: service.ErrDuplicateUsername.Error()},
			statusCode: http.StatusConflict,
		},
		{
			name:       "userCheckUsername-error-unexpected",
			payload:    &models.CheckUsernameInput{Username: "unexpected error"},
			statusCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		var r *http.Request
		if tt.payload == nil {
			r = httptest.NewRequest("POST", "/users/check-username", nil)
		} else {
			jsonBytes, _ := json.Marshal(tt.payload)
			r = httptest.NewRequest("POST", "/users/check-username", bytes.NewBuffer(jsonBytes))
		}

		r.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		handler := http.HandlerFunc(h.user.CheckUsername)
		handler.ServeHTTP(w, r)

		if w.Code != tt.statusCode {
			t.Errorf("%s returned response code of %d, wanted %d", tt.name, w.Code, tt.statusCode)
		}
	}
}

func TestUser_Get(t *testing.T) {
	var tests = []struct {
		name       string
		username   string
		statusCode int
	}{
		{
			name:       "userGet-ok",
			username:   "example",
			statusCode: http.StatusOK,
		},
		{
			name:       "userGet-error-no-user",
			username:   service.ErrNoUser.Error(),
			statusCode: http.StatusNotFound,
		},
		{
			name:       "userGet-error-unexpected",
			username:   "unexpected error",
			statusCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		r := httptest.NewRequest("GET", "/users/{username}", nil)
		ctx := getCtxWithParam(r, params{"username": tt.username})
		r = r.WithContext(ctx)
		r.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		handler := http.HandlerFunc(h.user.Get)
		handler.ServeHTTP(w, r)

		if w.Code != tt.statusCode {
			t.Errorf("%s returned response code of %d, wanted %d", tt.name, w.Code, tt.statusCode)
		}
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
		{
			name:       "updateProfile-ok",
			username:   "test",
			statusCode: http.StatusOK,
		},
		{
			name:       "updateProfile-error-parsing-form",
			noForm:     true,
			statusCode: http.StatusBadRequest,
		},
		{
			name:           "updateProfile-error-validation",
			failValidation: true,
			statusCode:     http.StatusBadRequest,
		},
		{
			name:       "updateProfile-error-no-user",
			username:   service.ErrNoUser.Error(),
			statusCode: http.StatusNotFound,
		},
		{
			name:       "updateProfile-error-unauthorized",
			username:   service.ErrUnauthorized.Error(),
			statusCode: http.StatusUnauthorized,
		},
		{
			name:       "updateProfile-error-avatar-too-large",
			username:   service.ErrAvatarTooLarge.Error(),
			statusCode: http.StatusBadRequest,
		},
		{
			name:       "updateProfile-error-avatar-invalid",
			username:   service.ErrAvatarInvalid.Error(),
			statusCode: http.StatusBadRequest,
		},
		{
			name:       "updateProfile-error-duplicate-username",
			username:   service.ErrDuplicateUsername.Error(),
			statusCode: http.StatusConflict,
		},
		{
			name:       "updateProfile-error-unexpected",
			username:   "unexpected error",
			statusCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		var b bytes.Buffer
		writer := multipart.NewWriter(&b)
		v := "test"
		if tt.failValidation {
			v = ""
		}
		writer.WriteField("username", v)
		writer.Close()

		var r *http.Request
		if tt.noForm {
			r = httptest.NewRequest("PATCH", "/users/{username}", nil)
		} else {
			r = httptest.NewRequest("PATCH", "/users/{username}", &b)
		}

		ctx := getCtxWithParam(r, params{"username": tt.username})
		ctx = context.WithValue(ctx, "user_id", 1)
		r = r.WithContext(ctx)
		r.Header.Set("Content-Type", writer.FormDataContentType())
		w := httptest.NewRecorder()
		handler := http.HandlerFunc(h.user.UpdateProfile)
		handler.ServeHTTP(w, r)

		if w.Code != tt.statusCode {
			t.Errorf("%s returned response code of %d, wanted %d", tt.name, w.Code, tt.statusCode)
		}
	}
}
