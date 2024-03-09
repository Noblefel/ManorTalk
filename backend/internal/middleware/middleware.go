package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/Noblefel/ManorTalk/backend/internal/config"
	res "github.com/Noblefel/ManorTalk/backend/internal/utils/response"
	"github.com/Noblefel/ManorTalk/backend/internal/utils/token"
)

type Middleware struct {
	c *config.AppConfig
}

func New(c *config.AppConfig) *Middleware { return &Middleware{c} }

func (m *Middleware) Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		accessToken := r.Header.Get("Authorization")

		if accessToken == "" {
			res.Message(w, http.StatusUnauthorized, "You need to login first")
			return
		}

		tokenDetails, err := token.Parse(m.c.AccessTokenKey, accessToken)
		if err != nil {
			if strings.Contains(err.Error(), "expired") {
				res.Message(w, http.StatusUnauthorized, "Token Expired")
				return
			}

			res.Message(w, http.StatusUnauthorized, "Invalid Token")
			return
		}

		ctx := context.WithValue(r.Context(), "user_id", tokenDetails.UserId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
