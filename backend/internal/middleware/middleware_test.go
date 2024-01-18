package middleware

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"github.com/Noblefel/ManorTalk/backend/internal/config"
	"github.com/Noblefel/ManorTalk/backend/internal/database"
	"github.com/Noblefel/ManorTalk/backend/internal/utils/token"
)

func TestNewMiddleware(t *testing.T) {
	var db *database.DB
	var c *config.AppConfig
	middleware := New(c, db)

	typeString := reflect.TypeOf(middleware).String()
	if typeString != "*middleware.Middleware" {
		t.Error("middleware.New() did not get the correct type, wanted *middleware.Middleware")
	}
}

var m = NewTest(&config.AppConfig{
	AccessTokenKey: "test",
	AccessTokenExp: 5 * time.Minute,
})

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

var middlewareAuthTests = []struct {
	name               string
	authorization      string
	expectedUserId     int
	expectedErrMessage string
	statusCode         int
}{
	{
		name:           "middlewareAuth-ok",
		authorization:  sampleToken,
		expectedUserId: 1,
		statusCode:     http.StatusOK,
	},
	{
		name:           "middlewareAuth-ok-2",
		authorization:  sampleToken2,
		expectedUserId: 2,
		statusCode:     http.StatusOK,
	},
	{
		name:          "middlewareAuth-empty-authorization-header",
		authorization: "",
		// expectedErrMessage: "You need to login first",
		statusCode: http.StatusUnauthorized,
	},
	{
		name:          "middlewareAuth-expired-token",
		authorization: sampleToken3,
		// expectedErrMessage: "Token Expired",
		statusCode: http.StatusUnauthorized,
	},
	{
		name:          "middlewareAuth-invalid-token",
		authorization: "asdcapsdjapcjsdpoajd",
		// expectedErrMessage: "Invalid Token",
		statusCode: http.StatusUnauthorized,
	},
}

func TestMiddleware_Auth(t *testing.T) {
	for _, tt := range middlewareAuthTests {
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

		// jsonResp := new(struct {
		// 	Message string `json:"message"`
		// })

		// json.NewDecoder(w.Body).Decode(jsonResp)

		if w.Code != tt.statusCode {
			t.Errorf("%s returned response code of %d, wanted %d", tt.name, w.Code, tt.statusCode)
		}

		// if jsonResp.Message != tt.expectedErrMessage {
		// 	t.Errorf("%s returned wrong error message \n Wanted: %s \n Got: %s", tt.name, tt.expectedErrMessage, jsonResp.Message)
		// }
	}
}
