package handlers

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	auth_service "github.com/Noblefel/ManorTalk/backend/internal/service/auth"
	post_service "github.com/Noblefel/ManorTalk/backend/internal/service/post"
	"github.com/go-chi/chi/v5"
)

type testHandlers struct {
	auth *AuthHandlers
	post *PostHandlers
}

func newTestHandlers() *testHandlers {
	authMock := auth_service.NewMockAuthService()
	postMock := post_service.NewMockPostService()

	return &testHandlers{
		auth: NewAuthHandlers(authMock),
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
		{
			name:       "not-found",
			url:        "/xmo02v3o2cm3ro",
			method:     "GET",
			handler:    NotFound,
			statusCode: http.StatusNotFound,
		},
		{
			name:       "method-not-allowed",
			url:        "/users",
			method:     "asjcaosjdcoa",
			handler:    MethodNotAllowed,
			statusCode: http.StatusMethodNotAllowed,
		},
	}

	for _, tt := range tests {
		r := httptest.NewRequest(tt.method, tt.url, nil)
		w := httptest.NewRecorder()

		tt.handler.ServeHTTP(w, r)

		if w.Code != tt.statusCode {
			t.Errorf("%s returned response code of %d, wanted %d", tt.name, w.Code, tt.statusCode)
		}
	}
}
