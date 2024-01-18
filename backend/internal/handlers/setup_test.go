package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/Noblefel/ManorTalk/backend/internal/config"
	"github.com/go-chi/chi/v5"
)

type testHandlers struct {
	auth *AuthHandlers
	post *PostHandlers
}

func newTestHandlers() *testHandlers {
	testConfig := config.AppConfig{
		AccessTokenKey:  "access_key",
		AccessTokenExp:  1 * time.Minute,
		RefreshTokenKey: "refresh_key",
		RefreshTokenExp: 1 * time.Minute,
	}

	return &testHandlers{
		auth: NewTestAuthHandlers(&testConfig),
		post: NewTestPostHandlers(&testConfig),
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
