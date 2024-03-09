package handlers

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	auth_service "github.com/Noblefel/ManorTalk/backend/internal/service/auth"
	post_service "github.com/Noblefel/ManorTalk/backend/internal/service/post"
	user_service "github.com/Noblefel/ManorTalk/backend/internal/service/user"
	"github.com/go-chi/chi/v5"
)

type testHandlers struct {
	auth *AuthHandlers
	user *UserHandlers
	post *PostHandlers
}

func newTestHandlers() *testHandlers {
	authMock := auth_service.NewMockAuthService()
	userMock := user_service.NewMockUserService()
	postMock := post_service.NewMockPostService()

	return &testHandlers{
		auth: NewAuthHandlers(authMock),
		user: NewUserHandlers(userMock),
		post: NewPostHandlers(postMock),
	}
}

var h = newTestHandlers()

type params map[string]string

func getCtxWithParam(r *http.Request, p params) context.Context {
	ctx := r.Context()
	chiCtx := chi.NewRouteContext()
	for k, v := range p {
		chiCtx.URLParams.Add(k, v)
	}
	ctx = context.WithValue(ctx, chi.RouteCtxKey, chiCtx)
	return ctx
}

func TestBaseHandlers(t *testing.T) {
	var tests = []struct {
		name       string
		url        string
		method     string
		handler    http.HandlerFunc
		statusCode int
	}{
		{"success not found", "/xmo02v3o2cm3ro", "GET", NotFound, http.StatusNotFound},
		{"success method not allowed", "/users", "asjcaosjdcoa", MethodNotAllowed, http.StatusMethodNotAllowed},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := httptest.NewRequest(tt.method, tt.url, nil)
			w := httptest.NewRecorder()

			tt.handler.ServeHTTP(w, r)

			if w.Code != tt.statusCode {
				t.Errorf("want %d, got %d", tt.statusCode, w.Code)
			}
		})
	}
}
