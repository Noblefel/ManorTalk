package middleware

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"github.com/Noblefel/ManorTalk/backend/internal/config"
	"github.com/Noblefel/ManorTalk/backend/internal/utils/token"
)

func TestNewMiddleware(t *testing.T) {
	var c *config.AppConfig
	middleware := New(c)

	typeString := reflect.TypeOf(middleware).String()
	if typeString != "*middleware.Middleware" {
		t.Error("middleware.New() get incorrect type, wanted *middleware.Middleware")
	}
}

var m = New(&config.AppConfig{
	AccessTokenKey: "test",
	AccessTokenExp: 5 * time.Minute,
})

func TestMiddleware_Auth(t *testing.T) {
	var sampleToken, _ = token.Generate(token.Details{
		SecretKey: m.c.AccessTokenKey,
		UserId:    1,
		UniqueId:  "sxzcro2nrondoaisncd",
		Duration:  m.c.AccessTokenExp,
	})

	var sampleToken2, _ = token.Generate(token.Details{
		SecretKey: "test",
		UserId:    2,
		Duration:  m.c.AccessTokenExp,
	})

	var sampleToken3, _ = token.Generate(token.Details{
		SecretKey: m.c.AccessTokenKey,
		UserId:    2,
	})

	var tests = []struct {
		name           string
		authorization  string
		expectedUserId int
		statusCode     int
	}{
		{"success", sampleToken, 1, http.StatusOK},
		{"success 2", sampleToken2, 2, http.StatusOK},
		{"empty authorization header", "", 0, http.StatusUnauthorized},
		{"expired token", sampleToken3, 0, http.StatusUnauthorized},
		{"invalid token", "asdcapsdjapcjsdpoajd", 0, http.StatusUnauthorized},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				userId := r.Context().Value("user_id")
				if userId == nil {
					t.Error("User id not in context")
					return
				}

				if userId != tt.expectedUserId {
					t.Error("Expected user id does not match")
					return
				}
			})

			r := httptest.NewRequest("GET", "/", nil)
			r.Header.Set("Authorization", tt.authorization)
			w := httptest.NewRecorder()
			h := m.Auth(next)
			h.ServeHTTP(w, r)

			if w.Code != tt.statusCode {
				t.Errorf("want %d, got %d", tt.statusCode, w.Code)
			}
		})
	}
}
