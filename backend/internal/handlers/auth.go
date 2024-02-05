package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/Noblefel/ManorTalk/backend/internal/models"
	service "github.com/Noblefel/ManorTalk/backend/internal/service/auth"
	res "github.com/Noblefel/ManorTalk/backend/internal/utils/response"
	"github.com/Noblefel/ManorTalk/backend/internal/utils/validate"
)

type AuthHandlers struct {
	service service.AuthService
}

func NewAuthHandlers(s service.AuthService) *AuthHandlers {
	return &AuthHandlers{
		service: s,
	}
}

func (h *AuthHandlers) Register(w http.ResponseWriter, r *http.Request) {
	var payload models.UserRegisterInput

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		res.MessageJSON(w, http.StatusBadRequest, "Error decoding json")
		return
	}

	if err := validate.Struct(payload); err != nil {
		res.JSON(w, http.StatusBadRequest, res.Response{
			Message: "Some fields are invalid",
			Errors:  err,
		})
		return
	}

	err := h.service.Register(payload)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrDuplicateEmail):
			res.MessageJSON(w, http.StatusConflict, "Username or "+err.Error())
			return
		default:
			log.Println(err)
			res.MessageJSON(w, http.StatusInternalServerError, "Sorry, we had some problems creating your account")
			return
		}
	}

	res.MessageJSON(w, http.StatusOK, "User succesfully registered")
}

func (h *AuthHandlers) Login(w http.ResponseWriter, r *http.Request) {
	var payload models.UserLoginInput

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		res.MessageJSON(w, http.StatusBadRequest, "Error decoding json")
		return
	}

	if err := validate.Struct(payload); err != nil {
		res.JSON(w, http.StatusBadRequest, res.Response{
			Message: "Some fields are invalid",
			Errors:  err,
		})
		return
	}

	user, accessToken, refreshToken, err := h.service.Login(payload)
	if err != nil {
		switch {
		case errors.Is(service.ErrInvalidCredentials, err), errors.Is(service.ErrNoUser, err):
			res.MessageJSON(w, http.StatusUnauthorized, service.ErrInvalidCredentials.Error())
			return
		default:
			log.Println(err)
			res.MessageJSON(w, http.StatusInternalServerError, "Sorry, we had some problems when authenticating")
			return
		}
	}

	http.SetCookie(w, &http.Cookie{
		Name:  "refresh_token",
		Value: refreshToken,
		Path:  "/",
	})

	res.JSON(w, http.StatusOK, res.Response{
		Data: map[string]interface{}{
			"access_token": accessToken,
			"user":         user,
		},
	})
}

func (h *AuthHandlers) Refresh(w http.ResponseWriter, r *http.Request) {
	refreshToken, err := r.Cookie("refresh_token")
	if err != nil {
		res.MessageJSON(w, http.StatusUnauthorized, service.ErrUnauthorized.Error())
		return
	}

	user, accessToken, err := h.service.Refresh(refreshToken.Value)
	if err != nil {
		switch {
		case errors.Is(service.ErrUnauthorized, err), errors.Is(service.ErrNoUser, err):
			res.MessageJSON(w, http.StatusUnauthorized, service.ErrUnauthorized.Error())
			return
		default:
			log.Println(err)
			res.MessageJSON(w, http.StatusInternalServerError, "Sorry, we had some problems verifying your request")
			return
		}
	}

	res.JSON(w, http.StatusOK, res.Response{
		Data: map[string]interface{}{
			"user":         user,
			"access_token": accessToken,
		},
	})
}

func (h *AuthHandlers) Logout(w http.ResponseWriter, r *http.Request) {
	refreshToken, err := r.Cookie("refresh_token")
	if err != nil {
		log.Println(err)
		res.MessageJSON(w, http.StatusUnauthorized, service.ErrUnauthorized.Error())
		return
	}

	err = h.service.Logout(refreshToken.Value)
	if err != nil {
		log.Println(err)
		res.MessageJSON(w, http.StatusUnauthorized, service.ErrUnauthorized.Error())
		return
	}

	res.MessageJSON(w, http.StatusOK, "Logged out")
}
