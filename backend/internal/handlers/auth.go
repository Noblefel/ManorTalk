package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/Noblefel/ManorTalk/backend/internal/config"
	"github.com/Noblefel/ManorTalk/backend/internal/database"
	"github.com/Noblefel/ManorTalk/backend/internal/models"
	"github.com/Noblefel/ManorTalk/backend/internal/repository"
	"github.com/Noblefel/ManorTalk/backend/internal/repository/postgres"
	"github.com/Noblefel/ManorTalk/backend/internal/repository/redis"
	res "github.com/Noblefel/ManorTalk/backend/internal/utils/response"
	"github.com/Noblefel/ManorTalk/backend/internal/utils/token"
	"github.com/Noblefel/ManorTalk/backend/internal/utils/validate"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandlers struct {
	c         *config.AppConfig
	redisRepo repository.RedisRepo
	userRepo  repository.UserRepo
}

func NewAuthHandlers(c *config.AppConfig, db *database.DB) *AuthHandlers {
	return &AuthHandlers{
		c:         c,
		redisRepo: redis.NewRepo(db),
		userRepo:  postgres.NewUserRepo(db),
	}
}

func NewTestAuthHandlers(c *config.AppConfig) *AuthHandlers {
	return &AuthHandlers{
		c:         c,
		redisRepo: redis.NewTestRepo(),
		userRepo:  postgres.NewTestUserRepo(),
	}
}

func (h *AuthHandlers) Register(w http.ResponseWriter, r *http.Request) {
	var payload models.UserRegisterInput

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		res.JSON(w, r, http.StatusBadRequest, res.Response{
			Message: "Error decoding json",
		})
		return
	}

	if err := validate.Struct(payload); err != nil {
		res.JSON(w, r, http.StatusBadRequest, res.Response{
			Message: "Some fields are invalid",
			Errors:  err,
		})
		return
	}

	user, err := h.userRepo.Register(payload)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value") {
			res.JSON(w, r, http.StatusConflict, res.Response{
				Message: "Email already in use",
			})
			return
		}

		res.JSON(w, r, http.StatusInternalServerError, res.Response{
			Message: "Unexpected error when registering user",
		})
		return
	}

	res.JSON(w, r, http.StatusOK, res.Response{
		Message: "User succesfully registered",
		Data:    user,
	})
}

func (h *AuthHandlers) Login(w http.ResponseWriter, r *http.Request) {
	var payload models.UserLoginInput

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		res.JSON(w, r, http.StatusBadRequest, res.Response{
			Message: "Error decoding json",
		})
		return
	}

	if err := validate.Struct(payload); err != nil {
		res.JSON(w, r, http.StatusBadRequest, res.Response{
			Message: "Some fields are invalid",
			Errors:  err,
		})
		return
	}

	user, err := h.userRepo.Login(payload)
	if err != nil {
		if errors.Is(bcrypt.ErrMismatchedHashAndPassword, err) || errors.Is(sql.ErrNoRows, err) {
			res.JSON(w, r, http.StatusUnauthorized, res.Response{
				Message: "Invalid credentials",
			})
			return
		}

		res.JSON(w, r, http.StatusInternalServerError, res.Response{
			Message: "Unexpected error when getting user",
		})
		return
	}

	accessTD := token.Details{
		UserId:    user.Id,
		SecretKey: h.c.AccessTokenKey,
		Duration:  h.c.AccessTokenExp,
	}

	accessToken, err := token.Generate(accessTD)
	if err != nil {
		res.JSON(w, r, http.StatusInternalServerError, res.Response{
			Message: "Something went wrong",
		})
		return
	}

	refreshTD := token.Details{
		UserId:    user.Id,
		SecretKey: h.c.RefreshTokenKey,
		UniqueId:  uuid.NewString(),
		Duration:  h.c.RefreshTokenExp,
	}

	refreshToken, err := token.Generate(refreshTD)
	if err != nil {
		res.JSON(w, r, http.StatusInternalServerError, res.Response{
			Message: "Something went wrong",
		})
		return
	}

	if err = h.redisRepo.SetRefreshToken(refreshTD); err != nil {
		res.JSON(w, r, http.StatusInternalServerError, res.Response{
			Message: "Unexpected error when saving token",
		})
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:  "refresh_token",
		Value: refreshToken,
	})

	res.JSON(w, r, http.StatusOK, res.Response{
		Data: map[string]string{
			"access_token": accessToken,
		},
	})
}

func (h *AuthHandlers) Refresh(w http.ResponseWriter, r *http.Request) {
	refreshToken, err := r.Cookie("refresh_token")
	if err != nil {
		res.JSON(w, r, http.StatusUnauthorized, res.Response{
			Message: "You need to log in first",
		})
		return
	}

	tokenDetails, err := token.Parse(h.c.RefreshTokenKey, refreshToken.Value)
	if err != nil {
		res.JSON(w, r, http.StatusUnauthorized, res.Response{
			Message: "Unauthorized",
		})
		return
	}

	uuid, err := h.redisRepo.GetRefreshToken(*tokenDetails)
	if err != nil || uuid != tokenDetails.UniqueId {
		res.JSON(w, r, http.StatusUnauthorized, res.Response{
			Message: "Unauthorized",
		})
		return
	}

	user, err := h.userRepo.GetUserById(tokenDetails.UserId)
	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			res.JSON(w, r, http.StatusNotFound, res.Response{
				Message: "User not found",
			})
			return
		}

		res.JSON(w, r, http.StatusInternalServerError, res.Response{
			Message: "Unexpected error when getting user",
		})
		return
	}

	accessToken, err := token.Generate(token.Details{
		SecretKey: h.c.AccessTokenKey,
		UserId:    user.Id,
		Duration:  h.c.AccessTokenExp,
	})

	if err != nil {
		res.JSON(w, r, http.StatusInternalServerError, res.Response{
			Message: "Something went wrong",
		})
		return
	}

	user.Password = ""

	res.JSON(w, r, http.StatusOK, res.Response{
		Data: map[string]interface{}{
			"user":         user,
			"access_token": accessToken,
		},
	})
}
