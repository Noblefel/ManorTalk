package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/Noblefel/ManorTalk/backend/internal/config"
	"github.com/Noblefel/ManorTalk/backend/internal/database"
	res "github.com/Noblefel/ManorTalk/backend/internal/utils/response"
	"github.com/Noblefel/ManorTalk/backend/internal/utils/token"
)

type Middleware struct {
	c  *config.AppConfig
	db *database.DB
}

func New(c *config.AppConfig, db *database.DB) *Middleware {
	return &Middleware{
		c:  c,
		db: db,
	}
}

func NewTest(c *config.AppConfig) *Middleware {
	return &Middleware{
		c: c,
	}
}

func (m *Middleware) Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		accessToken := r.Header.Get("Authorization")

		if accessToken == "" {
			res.JSON(w, r, http.StatusUnauthorized, res.Response{
				Message: "You need to login first",
			})
			return
		}

		tokenDetails, err := token.Parse(m.c.AccessTokenKey, accessToken)
		if err != nil {
			// if errors.Is(jwt.ErrTokenExpired, err) {
			if strings.Contains(err.Error(), "expired") {
				res.JSON(w, r, http.StatusUnauthorized, res.Response{
					Message: "Token Expired",
				})
				return
			}

			res.JSON(w, r, http.StatusUnauthorized, res.Response{
				Message: "Invalid Token",
			})
			return
		}

		ctx := context.WithValue(r.Context(), "user_id", tokenDetails.UserId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
