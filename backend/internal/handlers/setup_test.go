package handlers

import (
	"time"

	"github.com/Noblefel/ManorTalk/backend/internal/config"
)

type testHandlers struct {
	auth *AuthHandlers
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
	}
}

var h = newTestHandlers()
